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
	"io"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type retriableStreamer struct {
	grpc.ClientStream
	mu sync.Mutex

	bufferedSends []interface{}
	wasClosedSend bool // indicates that CloseSend was closed
	callOpts      *options
	parentCtx     context.Context
	streamerCall  func(ctx context.Context) (grpc.ClientStream, error)
}

func (s *retriableStreamer) setStream(clientStream grpc.ClientStream) {
	s.mu.Lock()
	s.ClientStream = clientStream
	s.mu.Unlock()
}

func (s *retriableStreamer) getStream() grpc.ClientStream {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ClientStream
}

func (s *retriableStreamer) SendMsg(m interface{}) error {
	s.mu.Lock()
	s.bufferedSends = append(s.bufferedSends, m)
	s.mu.Unlock()
	return s.getStream().SendMsg(m)
}

func (s *retriableStreamer) CloseSend() error {
	s.mu.Lock()
	s.wasClosedSend = true
	s.mu.Unlock()
	return s.getStream().CloseSend()
}

func (s *retriableStreamer) Header() (metadata.MD, error) {
	return s.getStream().Header()
}

func (s *retriableStreamer) Trailer() metadata.MD {
	return s.getStream().Trailer()
}

func (s *retriableStreamer) RecvMsg(m interface{}) error {
	shouldRetry, err := s.receiveMsgAndIndicateRetry(m)
	if !shouldRetry {
		return err
	}

	for attempt := uint(1); attempt < s.callOpts.max; attempt++ {
		timeout := s.callOpts.timeout
		if err := waitRetryBackoff(s.parentCtx, timeout); err != nil {
			return err
		}

		newStream, err := s.reestablishStreamAndResendBuffer(s.parentCtx)
		if err != nil {
			if isRetriable(err, s.callOpts) {
				continue
			}
			return err
		}

		s.setStream(newStream)
		shouldRetry, err = s.receiveMsgAndIndicateRetry(m)
		if !shouldRetry {
			return err
		}
	}
	return err
}

func (s *retriableStreamer) receiveMsgAndIndicateRetry(m interface{}) (bool, error) {
	err := s.getStream().RecvMsg(m)
	if err == nil || err == io.EOF {
		return false, err
	}
	return isRetriable(err, s.callOpts), err
}

func (s *retriableStreamer) reestablishStreamAndResendBuffer(ctx context.Context) (grpc.ClientStream, error) {
	s.mu.Lock()
	bufferedSends = s.bufferedSends
	s.mu.Unlock()
	newStream, err := s.streamerCall(ctx)
	if err != nil {
		// TODO: add logging
		return nil, err
	}
	for _, msg := range bufferedSends {
		if err := newStream.SendMsg(msg); err != nil {
			// TODO: log failled to resend message to reconnect
			return nil, err
		}
	}
	if err := newStream.CloseSend(); err != nil {
		// TODO: log failed to closeSend on new Stream
		return nil, err
	}
	return newStream, nil
}
