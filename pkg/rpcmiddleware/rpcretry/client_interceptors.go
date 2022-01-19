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

package rpcretry

import (
	"context"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/backoffutils"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AttemptMetadataKey struct{}

const (
	headerXRateLimit     = "x-rate-limit-limit"
	headerXRateAvailable = "x-rate-limit-available"
	headerXRateReset     = "x-rate-limit-reset"
	headerXRateRetry     = "x-rate-limit-retry"
)

var errInvalidStreamType = errors.DefineUnavailable("unsupported_stream_type_grpcretry", "Client streams are not supported in the grpc retry interceptor")

// UnaryClientInterceptor returns a new unary client interceptor that retries the execution of external gRPC calls, the
// retry attempt will only occur if any of the validators define the error as a trigger.
func UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
	callOpts := evaluateClientOpt(opts...)
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		logger := log.FromContext(ctx)

		var md metadata.MD
		var err error
		err = invoker(ctx, method, req, reply, cc, append(opts, grpc.Header(&md))...)
		if err == nil || !isRetriable(err, callOpts) {
			return err
		}

		for attempt := uint(1); attempt <= callOpts.max; attempt++ {
			retryTimeout := callOpts.timeout
			if callOpts.enableXrateHeader {
				if headerTimeout := getHeaderLimiterTimeout(ctx, md, callOpts); headerTimeout > 0 {
					retryTimeout = headerTimeout
				}
			}

			logger.WithField("attempt", attempt).Infof("Failed request, waiting %v until next attempt", retryTimeout)
			err = waitRetryBackoff(ctx, retryTimeout)
			if err != nil {
				logger.WithError(err).Debug("An unexpected error occurred while in the timeout for the next request retry")
				return err
			}

			callCtx := context.WithValue(ctx, AttemptMetadataKey{}, attempt)
			err = invoker(callCtx, method, req, reply, cc, append(opts, grpc.Header(&md))...)
			if err == nil {
				return nil
			}
			logger.WithField("attempt", attempt).WithError(err).Debug("Failed request retry")

			if !isRetriable(err, callOpts) {
				return err
			}
		}

		return err
	}
}

// StreamClientInterceptor returns a new streaming client interceptor that retries the execution of external gRPC
// calls, the retry attempt will only occur if any of the validators define the error as a trigger.
func StreamClientInterceptor(opts ...Option) grpc.StreamClientInterceptor {
	callOpts := evaluateClientOpt(opts...)
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		if callOpts.max == 0 {
			return streamer(ctx, desc, cc, method, opts...)
		}

		if desc.ClientStreams {
			return nil, errInvalidStreamType.New()
		}

		logger := log.FromContext(ctx)
		var md metadata.MD
		var err error

		for attempt := uint(0); attempt <= callOpts.max; attempt++ {
			retryTimeout := callOpts.timeout
			if callOpts.enableXrateHeader {
				if headerTimeout := getHeaderLimiterTimeout(ctx, md, callOpts); headerTimeout > 0 {
					retryTimeout = headerTimeout
				}
			}

			// TODO: incorporate the attemp == 0 skip within the waitRetryBackoff
			if attempt > 0 {
				logger.WithField("attempt", attempt).Infof("Failed request, waiting %v until next attempt", retryTimeout)
				err = waitRetryBackoff(ctx, retryTimeout)
				if err != nil {
					logger.WithError(err).Debug("An unexpected error occurred while in the timeout for the next request retry")
					return nil, err
				}
			}

			var clientStream grpc.ClientStream
			callCtx := context.WithValue(ctx, AttemptMetadataKey{}, attempt)
			clientStream, err = streamer(callCtx, desc, cc, method, append(opts, grpc.Header(&md))...)
			if err == nil {
				retriableStream := &retriableStreamer{
					ClientStream: clientStream,
					callOpts:     callOpts,
					parentCtx:    ctx,
					streamerCall: func(ctx context.Context) (grpc.ClientStream, error) {
						return streamer(ctx, desc, cc, method, opts...)
					},
				}
				return retriableStream, nil
			}
			logger.WithField("attempt", attempt).WithError(err).Debug("Failed request")

			if !isRetriable(err, callOpts) {
				return clientStream, err
			}
		}
		return nil, err
	}
}

func evaluateClientOpt(opts ...Option) *options {
	optCopy := &options{}
	*optCopy = *defaultOptions

	for _, f := range opts {
		f(optCopy)
	}
	return optCopy
}

func isRetriable(err error, opt *options) bool {
	if err == nil {
		return false
	}
	for _, check := range opt.validators {
		if check(err) {
			return true
		}
	}
	return false
}

func waitRetryBackoff(ctx context.Context, timeout time.Duration) error {
	timer := time.NewTicker(timeout)
	select {
	case <-ctx.Done():
		timer.Stop()
		return ctx.Err()
	case <-timer.C:
	}
	return nil
}

func hasRateLimitHeaders(md metadata.MD) bool {
	return len(md.Get(headerXRateLimit)) > 0 &&
		len(md.Get(headerXRateAvailable)) > 0 &&
		len(md.Get(headerXRateReset)) > 0 &&
		len(md.Get(headerXRateRetry)) > 0
}

func getHeaderLimiterTimeout(ctx context.Context, md metadata.MD, opt *options) time.Duration {
	if !hasRateLimitHeaders(md) {
		return -1
	}
	var reset int
	reset, _ = strconv.Atoi(md.Get(headerXRateReset)[0])
	timeout := time.Duration(reset) * time.Second
	if opt.jitter > 0 {
		timeout = backoffutils.JitterUp(timeout, opt.jitter)
	}
	return timeout
}
