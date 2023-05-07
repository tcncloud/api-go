// Copyright (c) 2020, TCN Inc.
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
// Source: api/v1alpha1/ghostnotifier/service.proto

package ghostnotifierconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	commons "github.com/tcncloud/api-go/commons"
	ghostnotifier "github.com/tcncloud/api-go/v1alpha1/ghostnotifier"
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
	// GhostNotifierApiName is the fully-qualified name of the GhostNotifierApi service.
	GhostNotifierApiName = "api.v1alpha1.ghostnotifier.GhostNotifierApi"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GhostNotifierApiListNotificationsProcedure is the fully-qualified name of the GhostNotifierApi's
	// ListNotifications RPC.
	GhostNotifierApiListNotificationsProcedure = "/api.v1alpha1.ghostnotifier.GhostNotifierApi/ListNotifications"
)

// GhostNotifierApiClient is a client for the api.v1alpha1.ghostnotifier.GhostNotifierApi service.
type GhostNotifierApiClient interface {
	// Opens a server side stream that will forward and ghost notifications to the client for the given user
	ListNotifications(context.Context, *connect_go.Request[ghostnotifier.ListNotificationsReq]) (*connect_go.ServerStreamForClient[commons.GhostNotification], error)
}

// NewGhostNotifierApiClient constructs a client for the api.v1alpha1.ghostnotifier.GhostNotifierApi
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGhostNotifierApiClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) GhostNotifierApiClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &ghostNotifierApiClient{
		listNotifications: connect_go.NewClient[ghostnotifier.ListNotificationsReq, commons.GhostNotification](
			httpClient,
			baseURL+GhostNotifierApiListNotificationsProcedure,
			opts...,
		),
	}
}

// ghostNotifierApiClient implements GhostNotifierApiClient.
type ghostNotifierApiClient struct {
	listNotifications *connect_go.Client[ghostnotifier.ListNotificationsReq, commons.GhostNotification]
}

// ListNotifications calls api.v1alpha1.ghostnotifier.GhostNotifierApi.ListNotifications.
func (c *ghostNotifierApiClient) ListNotifications(ctx context.Context, req *connect_go.Request[ghostnotifier.ListNotificationsReq]) (*connect_go.ServerStreamForClient[commons.GhostNotification], error) {
	return c.listNotifications.CallServerStream(ctx, req)
}

// GhostNotifierApiHandler is an implementation of the api.v1alpha1.ghostnotifier.GhostNotifierApi
// service.
type GhostNotifierApiHandler interface {
	// Opens a server side stream that will forward and ghost notifications to the client for the given user
	ListNotifications(context.Context, *connect_go.Request[ghostnotifier.ListNotificationsReq], *connect_go.ServerStream[commons.GhostNotification]) error
}

// NewGhostNotifierApiHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGhostNotifierApiHandler(svc GhostNotifierApiHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(GhostNotifierApiListNotificationsProcedure, connect_go.NewServerStreamHandler(
		GhostNotifierApiListNotificationsProcedure,
		svc.ListNotifications,
		opts...,
	))
	return "/api.v1alpha1.ghostnotifier.GhostNotifierApi/", mux
}

// UnimplementedGhostNotifierApiHandler returns CodeUnimplemented from all methods.
type UnimplementedGhostNotifierApiHandler struct{}

func (UnimplementedGhostNotifierApiHandler) ListNotifications(context.Context, *connect_go.Request[ghostnotifier.ListNotificationsReq], *connect_go.ServerStream[commons.GhostNotification]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.ghostnotifier.GhostNotifierApi.ListNotifications is not implemented"))
}
