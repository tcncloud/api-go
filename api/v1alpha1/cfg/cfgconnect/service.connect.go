// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/cfg/service.proto

package cfgconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	cfg "github.com/tcncloud/api-go/api/v1alpha1/cfg"
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
	CfgName = "api.v1alpha1.cfg.Cfg"
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
	CfgGetWebAgentConfigProcedure = "/api.v1alpha1.cfg.Cfg/GetWebAgentConfig"
)

// CfgClient is a client for the api.v1alpha1.cfg.Cfg service.
type CfgClient interface {
	// Get the configuration for a given web agent.
	GetWebAgentConfig(context.Context, *connect_go.Request[cfg.GetWebAgentConfigReq]) (*connect_go.Response[cfg.WebAgent], error)
}

// NewCfgClient constructs a client for the api.v1alpha1.cfg.Cfg service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCfgClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CfgClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &cfgClient{
		getWebAgentConfig: connect_go.NewClient[cfg.GetWebAgentConfigReq, cfg.WebAgent](
			httpClient,
			baseURL+CfgGetWebAgentConfigProcedure,
			opts...,
		),
	}
}

// cfgClient implements CfgClient.
type cfgClient struct {
	getWebAgentConfig *connect_go.Client[cfg.GetWebAgentConfigReq, cfg.WebAgent]
}

// GetWebAgentConfig calls api.v1alpha1.cfg.Cfg.GetWebAgentConfig.
func (c *cfgClient) GetWebAgentConfig(ctx context.Context, req *connect_go.Request[cfg.GetWebAgentConfigReq]) (*connect_go.Response[cfg.WebAgent], error) {
	return c.getWebAgentConfig.CallUnary(ctx, req)
}

// CfgHandler is an implementation of the api.v1alpha1.cfg.Cfg service.
type CfgHandler interface {
	// Get the configuration for a given web agent.
	GetWebAgentConfig(context.Context, *connect_go.Request[cfg.GetWebAgentConfigReq]) (*connect_go.Response[cfg.WebAgent], error)
}

// NewCfgHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCfgHandler(svc CfgHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	cfgGetWebAgentConfigHandler := connect_go.NewUnaryHandler(
		CfgGetWebAgentConfigProcedure,
		svc.GetWebAgentConfig,
		opts...,
	)
	return "/api.v1alpha1.cfg.Cfg/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CfgGetWebAgentConfigProcedure:
			cfgGetWebAgentConfigHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCfgHandler returns CodeUnimplemented from all methods.
type UnimplementedCfgHandler struct{}

func (UnimplementedCfgHandler) GetWebAgentConfig(context.Context, *connect_go.Request[cfg.GetWebAgentConfigReq]) (*connect_go.Response[cfg.WebAgent], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.cfg.Cfg.GetWebAgentConfig is not implemented"))
}
