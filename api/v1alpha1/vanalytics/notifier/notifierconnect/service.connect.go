// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/vanalytics/notifier/service.proto

package notifierconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	notifier "github.com/tcncloud/api-go/api/v1alpha1/vanalytics/notifier"
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
	// NotifierName is the fully-qualified name of the Notifier service.
	NotifierName = "api.v1alpha1.vanalytics.notifier.Notifier"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// NotifierGetNotifyProcedure is the fully-qualified name of the Notifier's GetNotify RPC.
	NotifierGetNotifyProcedure = "/api.v1alpha1.vanalytics.notifier.Notifier/GetNotify"
)

// NotifierClient is a client for the api.v1alpha1.vanalytics.notifier.Notifier service.
type NotifierClient interface {
	// GetNotify returns a notify.
	GetNotify(context.Context, *connect_go.Request[notifier.GetNotifyRequest]) (*connect_go.Response[notifier.Notify], error)
}

// NewNotifierClient constructs a client for the api.v1alpha1.vanalytics.notifier.Notifier service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNotifierClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) NotifierClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &notifierClient{
		getNotify: connect_go.NewClient[notifier.GetNotifyRequest, notifier.Notify](
			httpClient,
			baseURL+NotifierGetNotifyProcedure,
			opts...,
		),
	}
}

// notifierClient implements NotifierClient.
type notifierClient struct {
	getNotify *connect_go.Client[notifier.GetNotifyRequest, notifier.Notify]
}

// GetNotify calls api.v1alpha1.vanalytics.notifier.Notifier.GetNotify.
func (c *notifierClient) GetNotify(ctx context.Context, req *connect_go.Request[notifier.GetNotifyRequest]) (*connect_go.Response[notifier.Notify], error) {
	return c.getNotify.CallUnary(ctx, req)
}

// NotifierHandler is an implementation of the api.v1alpha1.vanalytics.notifier.Notifier service.
type NotifierHandler interface {
	// GetNotify returns a notify.
	GetNotify(context.Context, *connect_go.Request[notifier.GetNotifyRequest]) (*connect_go.Response[notifier.Notify], error)
}

// NewNotifierHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNotifierHandler(svc NotifierHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	notifierGetNotifyHandler := connect_go.NewUnaryHandler(
		NotifierGetNotifyProcedure,
		svc.GetNotify,
		opts...,
	)
	return "/api.v1alpha1.vanalytics.notifier.Notifier/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case NotifierGetNotifyProcedure:
			notifierGetNotifyHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedNotifierHandler returns CodeUnimplemented from all methods.
type UnimplementedNotifierHandler struct{}

func (UnimplementedNotifierHandler) GetNotify(context.Context, *connect_go.Request[notifier.GetNotifyRequest]) (*connect_go.Response[notifier.Notify], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.vanalytics.notifier.Notifier.GetNotify is not implemented"))
}
