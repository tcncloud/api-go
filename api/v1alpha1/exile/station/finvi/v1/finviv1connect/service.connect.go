// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/exile/station/finvi/v1/service.proto

package finviv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/tcncloud/api-go/api/v1alpha1/exile/station/finvi/v1"
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
	// GenericFinviServiceName is the fully-qualified name of the GenericFinviService service.
	GenericFinviServiceName = "api.v1alpha1.exile.station.finvi.v1.GenericFinviService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GenericFinviServicePopAccountProcedure is the fully-qualified name of the GenericFinviService's
	// PopAccount RPC.
	GenericFinviServicePopAccountProcedure = "/api.v1alpha1.exile.station.finvi.v1.GenericFinviService/PopAccount"
)

// GenericFinviServiceClient is a client for the
// api.v1alpha1.exile.station.finvi.v1.GenericFinviService service.
type GenericFinviServiceClient interface {
	// Displays a given record from a pool to the specified user.
	PopAccount(context.Context, *connect_go.Request[v1.PopAccountReq]) (*connect_go.Response[v1.PopAccountRes], error)
}

// NewGenericFinviServiceClient constructs a client for the
// api.v1alpha1.exile.station.finvi.v1.GenericFinviService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGenericFinviServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) GenericFinviServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &genericFinviServiceClient{
		popAccount: connect_go.NewClient[v1.PopAccountReq, v1.PopAccountRes](
			httpClient,
			baseURL+GenericFinviServicePopAccountProcedure,
			opts...,
		),
	}
}

// genericFinviServiceClient implements GenericFinviServiceClient.
type genericFinviServiceClient struct {
	popAccount *connect_go.Client[v1.PopAccountReq, v1.PopAccountRes]
}

// PopAccount calls api.v1alpha1.exile.station.finvi.v1.GenericFinviService.PopAccount.
func (c *genericFinviServiceClient) PopAccount(ctx context.Context, req *connect_go.Request[v1.PopAccountReq]) (*connect_go.Response[v1.PopAccountRes], error) {
	return c.popAccount.CallUnary(ctx, req)
}

// GenericFinviServiceHandler is an implementation of the
// api.v1alpha1.exile.station.finvi.v1.GenericFinviService service.
type GenericFinviServiceHandler interface {
	// Displays a given record from a pool to the specified user.
	PopAccount(context.Context, *connect_go.Request[v1.PopAccountReq]) (*connect_go.Response[v1.PopAccountRes], error)
}

// NewGenericFinviServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGenericFinviServiceHandler(svc GenericFinviServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	genericFinviServicePopAccountHandler := connect_go.NewUnaryHandler(
		GenericFinviServicePopAccountProcedure,
		svc.PopAccount,
		opts...,
	)
	return "/api.v1alpha1.exile.station.finvi.v1.GenericFinviService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GenericFinviServicePopAccountProcedure:
			genericFinviServicePopAccountHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGenericFinviServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGenericFinviServiceHandler struct{}

func (UnimplementedGenericFinviServiceHandler) PopAccount(context.Context, *connect_go.Request[v1.PopAccountReq]) (*connect_go.Response[v1.PopAccountRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.exile.station.finvi.v1.GenericFinviService.PopAccount is not implemented"))
}
