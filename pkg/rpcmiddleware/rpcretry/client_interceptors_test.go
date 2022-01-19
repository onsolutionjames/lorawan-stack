// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpcretry_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ratelimit"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/rpcretry"
	"go.thethings.network/lorawan-stack/v3/pkg/util/rpctest"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	exaushtedErr = errors.DefineResourceExhausted("rpcretry_unary_resource_exausted", "mock error of a resource exaushted scenario")
	internalErr  = errors.DefineInternal("rpcretry_unary_internal", "mock error that represents an error that should not be retried")
)

type testService struct {
	rpctest.FooBarServer
	md metadata.MD

	reqCounter    uint
	unaryErr      error
	reqSleep      time.Duration
	contextFiller func(ctx context.Context) context.Context
}

func (fs *testService) Unary(ctx context.Context, foo *rpctest.Foo) (*rpctest.Bar, error) {
	fs.reqCounter++
	if fs.reqSleep > 0 {
		time.Sleep(fs.reqSleep)
	}
	if fs.contextFiller != nil {
		ctx = fs.contextFiller(ctx)
	}

	if err := grpc.SendHeader(ctx, fs.md); err != nil {
		return nil, err
	}
	return &rpctest.Bar{Message: "bar"}, fs.unaryErr
}

func Test_UnaryClientInterceptor(t *testing.T) {
	type Service struct {
		unaryErr      error
		sleep         time.Duration
		contextFiller func(ctx context.Context) context.Context
		md            metadata.MD
	}
	type Client struct {
		retries       uint
		retryTimeout  time.Duration
		useMetadata   bool
		jitter        float64
		contextFiller func(ctx context.Context) context.Context
	}

	for _, tt := range []struct {
		name              string
		service           Service
		client            Client
		errAssertion      func(error) bool
		expectedReqAmount int
	}{
		{
			name:              "no error",
			client:            Client{retries: 5, retryTimeout: 3 * test.Delay},
			service:           Service{sleep: 0},
			expectedReqAmount: 1,
		},
		{
			name:              "unretriable error",
			client:            Client{retries: 5, retryTimeout: 3 * test.Delay},
			service:           Service{unaryErr: internalErr.New()},
			errAssertion:      errors.IsInternal,
			expectedReqAmount: 1,
		},
		{
			name:              "retriable error",
			client:            Client{retries: 5, retryTimeout: 3 * test.Delay},
			service:           Service{exaushtedErr.New(), 0, nil, nil},
			errAssertion:      errors.IsResourceExhausted,
			expectedReqAmount: 6,
		},
		{
			name: "timeout error",
			client: Client{
				retries:      5,
				retryTimeout: 10 * test.Delay,
				contextFiller: func(ctx context.Context) context.Context {
					ctx, _ = context.WithTimeout(ctx, 5*test.Delay)
					return ctx
				},
			},
			service:           Service{sleep: 5 * test.Delay},
			errAssertion:      errors.IsDeadlineExceeded,
			expectedReqAmount: 1,
		},
		{
			name: "timeout from metadata rate-limiter",
			client: Client{
				retries:      1,
				retryTimeout: 5 * time.Second,
				contextFiller: func(ctx context.Context) context.Context {
					// Client timeout below default timeout, to trigger deadline error and fails test.
					ctx, _ = context.WithTimeout(ctx, 3*time.Second)
					return ctx
				},
				useMetadata: true,
			},
			service: Service{
				unaryErr: exaushtedErr.New(),
				contextFiller: func(ctx context.Context) context.Context {
					r := ratelimit.Result{Limit: 10, Remaining: 8, RetryAfter: time.Second, ResetAfter: time.Second}
					headers := r.GRPCHeaders()
					return metadata.NewOutgoingContext(ctx, headers)
				},
				md: ratelimit.Result{Limit: 10, Remaining: 8, RetryAfter: time.Second, ResetAfter: time.Second}.GRPCHeaders(),
			},
			errAssertion:      errors.IsResourceExhausted,
			expectedReqAmount: 2,
		},
	} {
		lis, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			panic(err)
		}

		server := grpc.NewServer()
		testService := &testService{
			unaryErr:      tt.service.unaryErr,
			reqSleep:      tt.service.sleep,
			contextFiller: tt.service.contextFiller,
			md:            tt.service.md,
		}
		rpctest.RegisterFooBarServer(server, testService)
		go server.Serve(lis)

		cc, err := grpc.DialContext(
			test.Context(), lis.Addr().String(), grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(rpcretry.UnaryClientInterceptor(
				rpcretry.WithMax(tt.client.retries),
				rpcretry.WithDefaultTimeout(tt.client.retryTimeout),
				rpcretry.UseMetadata(tt.client.useMetadata),
				rpcretry.WithJitter(tt.client.jitter),
			)),
		)
		if err != nil {
			t.Fail()
		}
		defer cc.Close()

		client := rpctest.NewFooBarClient(cc)
		t.Run(tt.name, func(t *testing.T) {
			a := assertions.New(t)

			ctx := test.Context()
			if tt.client.contextFiller != nil {
				ctx = tt.client.contextFiller(ctx)
			}

			resp, err := client.Unary(ctx, &rpctest.Foo{Message: "foo"})
			if tt.errAssertion != nil {
				a.So(tt.errAssertion(err), should.BeTrue)
			} else {
				a.So(err, should.BeNil)
				a.So(resp.Message, should.Equal, "bar")
			}
			a.So(testService.reqCounter, should.Equal, tt.expectedReqAmount)
		})
	}
}

func Test_StreamClientInterceptor(t *testing.T) {
	// test case - clientStreams not available
	{
		lis, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}

		server := grpc.NewServer()
		testService := &testService{}
		rpctest.RegisterFooBarServer(server, testService)
		go server.Serve(lis)

		cc, err := grpc.DialContext(
			test.Context(), lis.Addr().String(), grpc.WithInsecure(),
			grpc.WithStreamInterceptor(rpcretry.StreamClientInterceptor(
				rpcretry.WithMax(1),
				rpcretry.WithDefaultTimeout(100*time.Millisecond),
			)),
		)
		if err != nil {
			t.Fail()
		}
		defer cc.Close()

		a := assertions.New(t)
		client := rpctest.NewFooBarClient(cc)
		_, err = client.ClientStream(context.Background())
		a.So(errors.IsUnavailable(err), should.BeTrue)

	}

}
