// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/idp/service.proto

package idpconnect

import (
	connect_go "github.com/bufbuild/connect-go"
	_ "github.com/tcncloud/api-go/api/v1alpha1/idp"
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
	// IdentityProviderName is the fully-qualified name of the IdentityProvider service.
	IdentityProviderName = "api.v1alpha1.idp.IdentityProvider"
)

// IdentityProviderClient is a client for the api.v1alpha1.idp.IdentityProvider service.
type IdentityProviderClient interface {
}

// NewIdentityProviderClient constructs a client for the api.v1alpha1.idp.IdentityProvider service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIdentityProviderClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IdentityProviderClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &identityProviderClient{}
}

// identityProviderClient implements IdentityProviderClient.
type identityProviderClient struct {
}

// IdentityProviderHandler is an implementation of the api.v1alpha1.idp.IdentityProvider service.
type IdentityProviderHandler interface {
}

// NewIdentityProviderHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIdentityProviderHandler(svc IdentityProviderHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	return "/api.v1alpha1.idp.IdentityProvider/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIdentityProviderHandler returns CodeUnimplemented from all methods.
type UnimplementedIdentityProviderHandler struct{}
