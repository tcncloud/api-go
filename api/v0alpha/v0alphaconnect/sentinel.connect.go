// Copyright (c) 2019, TCN Inc.
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

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v0alpha/sentinel.proto

package v0alphaconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v0alpha "github.com/tcncloud/api-go/api/v0alpha"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// SentinelName is the fully-qualified name of the Sentinel service.
	SentinelName = "api.v0alpha.Sentinel"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SentinelSendEventsProcedure is the fully-qualified name of the Sentinel's SendEvents RPC.
	SentinelSendEventsProcedure = "/api.v0alpha.Sentinel/SendEvents"
)

// SentinelClient is a client for the api.v0alpha.Sentinel service.
type SentinelClient interface {
	// Send a json blob of ui events and logs.
	SendEvents(context.Context, *connect_go.Request[v0alpha.SendEventsReq]) (*connect_go.Response[v0alpha.SendEventsRes], error)
}

// NewSentinelClient constructs a client for the api.v0alpha.Sentinel service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSentinelClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SentinelClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &sentinelClient{
		sendEvents: connect_go.NewClient[v0alpha.SendEventsReq, v0alpha.SendEventsRes](
			httpClient,
			baseURL+SentinelSendEventsProcedure,
			opts...,
		),
	}
}

// sentinelClient implements SentinelClient.
type sentinelClient struct {
	sendEvents *connect_go.Client[v0alpha.SendEventsReq, v0alpha.SendEventsRes]
}

// SendEvents calls api.v0alpha.Sentinel.SendEvents.
func (c *sentinelClient) SendEvents(ctx context.Context, req *connect_go.Request[v0alpha.SendEventsReq]) (*connect_go.Response[v0alpha.SendEventsRes], error) {
	return c.sendEvents.CallUnary(ctx, req)
}

// SentinelHandler is an implementation of the api.v0alpha.Sentinel service.
type SentinelHandler interface {
	// Send a json blob of ui events and logs.
	SendEvents(context.Context, *connect_go.Request[v0alpha.SendEventsReq]) (*connect_go.Response[v0alpha.SendEventsRes], error)
}

// NewSentinelHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSentinelHandler(svc SentinelHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	sentinelSendEventsHandler := connect_go.NewUnaryHandler(
		SentinelSendEventsProcedure,
		svc.SendEvents,
		opts...,
	)
	return "/api.v0alpha.Sentinel/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SentinelSendEventsProcedure:
			sentinelSendEventsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSentinelHandler returns CodeUnimplemented from all methods.
type UnimplementedSentinelHandler struct{}

func (UnimplementedSentinelHandler) SendEvents(context.Context, *connect_go.Request[v0alpha.SendEventsReq]) (*connect_go.Response[v0alpha.SendEventsRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Sentinel.SendEvents is not implemented"))
}
