// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/explorer/service.proto

package explorerconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	explorer "github.com/tcncloud/api-go/api/v1alpha1/explorer"
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
	// ExplorerServiceName is the fully-qualified name of the ExplorerService service.
	ExplorerServiceName = "api.v1alpha1.explorer.ExplorerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ExplorerServiceListDatasourceSchemasProcedure is the fully-qualified name of the
	// ExplorerService's ListDatasourceSchemas RPC.
	ExplorerServiceListDatasourceSchemasProcedure = "/api.v1alpha1.explorer.ExplorerService/ListDatasourceSchemas"
	// ExplorerServiceQueryProcedure is the fully-qualified name of the ExplorerService's Query RPC.
	ExplorerServiceQueryProcedure = "/api.v1alpha1.explorer.ExplorerService/Query"
)

// ExplorerServiceClient is a client for the api.v1alpha1.explorer.ExplorerService service.
type ExplorerServiceClient interface {
	// ListDatasourceSchemas lists all accessible datasources and their schemas.
	ListDatasourceSchemas(context.Context, *connect_go.Request[explorer.ListDatasourceSchemasRequest]) (*connect_go.Response[explorer.ListDatasourceSchemasResponse], error)
	// Query queries a datasource.
	Query(context.Context, *connect_go.Request[explorer.QueryRequest]) (*connect_go.Response[explorer.QueryResponse], error)
}

// NewExplorerServiceClient constructs a client for the api.v1alpha1.explorer.ExplorerService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewExplorerServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ExplorerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &explorerServiceClient{
		listDatasourceSchemas: connect_go.NewClient[explorer.ListDatasourceSchemasRequest, explorer.ListDatasourceSchemasResponse](
			httpClient,
			baseURL+ExplorerServiceListDatasourceSchemasProcedure,
			opts...,
		),
		query: connect_go.NewClient[explorer.QueryRequest, explorer.QueryResponse](
			httpClient,
			baseURL+ExplorerServiceQueryProcedure,
			opts...,
		),
	}
}

// explorerServiceClient implements ExplorerServiceClient.
type explorerServiceClient struct {
	listDatasourceSchemas *connect_go.Client[explorer.ListDatasourceSchemasRequest, explorer.ListDatasourceSchemasResponse]
	query                 *connect_go.Client[explorer.QueryRequest, explorer.QueryResponse]
}

// ListDatasourceSchemas calls api.v1alpha1.explorer.ExplorerService.ListDatasourceSchemas.
func (c *explorerServiceClient) ListDatasourceSchemas(ctx context.Context, req *connect_go.Request[explorer.ListDatasourceSchemasRequest]) (*connect_go.Response[explorer.ListDatasourceSchemasResponse], error) {
	return c.listDatasourceSchemas.CallUnary(ctx, req)
}

// Query calls api.v1alpha1.explorer.ExplorerService.Query.
func (c *explorerServiceClient) Query(ctx context.Context, req *connect_go.Request[explorer.QueryRequest]) (*connect_go.Response[explorer.QueryResponse], error) {
	return c.query.CallUnary(ctx, req)
}

// ExplorerServiceHandler is an implementation of the api.v1alpha1.explorer.ExplorerService service.
type ExplorerServiceHandler interface {
	// ListDatasourceSchemas lists all accessible datasources and their schemas.
	ListDatasourceSchemas(context.Context, *connect_go.Request[explorer.ListDatasourceSchemasRequest]) (*connect_go.Response[explorer.ListDatasourceSchemasResponse], error)
	// Query queries a datasource.
	Query(context.Context, *connect_go.Request[explorer.QueryRequest]) (*connect_go.Response[explorer.QueryResponse], error)
}

// NewExplorerServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewExplorerServiceHandler(svc ExplorerServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	explorerServiceListDatasourceSchemasHandler := connect_go.NewUnaryHandler(
		ExplorerServiceListDatasourceSchemasProcedure,
		svc.ListDatasourceSchemas,
		opts...,
	)
	explorerServiceQueryHandler := connect_go.NewUnaryHandler(
		ExplorerServiceQueryProcedure,
		svc.Query,
		opts...,
	)
	return "/api.v1alpha1.explorer.ExplorerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ExplorerServiceListDatasourceSchemasProcedure:
			explorerServiceListDatasourceSchemasHandler.ServeHTTP(w, r)
		case ExplorerServiceQueryProcedure:
			explorerServiceQueryHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedExplorerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedExplorerServiceHandler struct{}

func (UnimplementedExplorerServiceHandler) ListDatasourceSchemas(context.Context, *connect_go.Request[explorer.ListDatasourceSchemasRequest]) (*connect_go.Response[explorer.ListDatasourceSchemasResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.explorer.ExplorerService.ListDatasourceSchemas is not implemented"))
}

func (UnimplementedExplorerServiceHandler) Query(context.Context, *connect_go.Request[explorer.QueryRequest]) (*connect_go.Response[explorer.QueryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.explorer.ExplorerService.Query is not implemented"))
}
