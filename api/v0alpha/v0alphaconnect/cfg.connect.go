// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v0alpha/cfg.proto

package v0alphaconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v0alpha "github.com/tcncloud/api-go/v0alpha"
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
	// CfgName is the fully-qualified name of the Cfg service.
	CfgName = "api.v0alpha.Cfg"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CfgGetWebAgentConfigProcedure is the fully-qualified name of the Cfg's GetWebAgentConfig RPC.
	CfgGetWebAgentConfigProcedure = "/api.v0alpha.Cfg/GetWebAgentConfig"
)

// CfgClient is a client for the api.v0alpha.Cfg service.
type CfgClient interface {
	GetWebAgentConfig(context.Context, *connect_go.Request[v0alpha.GetConfigReq]) (*connect_go.Response[v0alpha.WebAgent], error)
}

// NewCfgClient constructs a client for the api.v0alpha.Cfg service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCfgClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CfgClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &cfgClient{
		getWebAgentConfig: connect_go.NewClient[v0alpha.GetConfigReq, v0alpha.WebAgent](
			httpClient,
			baseURL+CfgGetWebAgentConfigProcedure,
			opts...,
		),
	}
}

// cfgClient implements CfgClient.
type cfgClient struct {
	getWebAgentConfig *connect_go.Client[v0alpha.GetConfigReq, v0alpha.WebAgent]
}

// GetWebAgentConfig calls api.v0alpha.Cfg.GetWebAgentConfig.
func (c *cfgClient) GetWebAgentConfig(ctx context.Context, req *connect_go.Request[v0alpha.GetConfigReq]) (*connect_go.Response[v0alpha.WebAgent], error) {
	return c.getWebAgentConfig.CallUnary(ctx, req)
}

// CfgHandler is an implementation of the api.v0alpha.Cfg service.
type CfgHandler interface {
	GetWebAgentConfig(context.Context, *connect_go.Request[v0alpha.GetConfigReq]) (*connect_go.Response[v0alpha.WebAgent], error)
}

// NewCfgHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCfgHandler(svc CfgHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(CfgGetWebAgentConfigProcedure, connect_go.NewUnaryHandler(
		CfgGetWebAgentConfigProcedure,
		svc.GetWebAgentConfig,
		opts...,
	))
	return "/api.v0alpha.Cfg/", mux
}

// UnimplementedCfgHandler returns CodeUnimplemented from all methods.
type UnimplementedCfgHandler struct{}

func (UnimplementedCfgHandler) GetWebAgentConfig(context.Context, *connect_go.Request[v0alpha.GetConfigReq]) (*connect_go.Response[v0alpha.WebAgent], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Cfg.GetWebAgentConfig is not implemented"))
}
