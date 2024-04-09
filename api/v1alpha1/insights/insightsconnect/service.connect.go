// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/insights/service.proto

package insightsconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	insights "github.com/tcncloud/api-go/api/v1alpha1/insights"
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
	// InsightsName is the fully-qualified name of the Insights service.
	InsightsName = "api.v1alpha1.insights.Insights"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// InsightsCreateInsightProcedure is the fully-qualified name of the Insights's CreateInsight RPC.
	InsightsCreateInsightProcedure = "/api.v1alpha1.insights.Insights/CreateInsight"
	// InsightsListInsightsProcedure is the fully-qualified name of the Insights's ListInsights RPC.
	InsightsListInsightsProcedure = "/api.v1alpha1.insights.Insights/ListInsights"
	// InsightsListOrgInsightsProcedure is the fully-qualified name of the Insights's ListOrgInsights
	// RPC.
	InsightsListOrgInsightsProcedure = "/api.v1alpha1.insights.Insights/ListOrgInsights"
	// InsightsUpdateInsightProcedure is the fully-qualified name of the Insights's UpdateInsight RPC.
	InsightsUpdateInsightProcedure = "/api.v1alpha1.insights.Insights/UpdateInsight"
	// InsightsDeleteInsightProcedure is the fully-qualified name of the Insights's DeleteInsight RPC.
	InsightsDeleteInsightProcedure = "/api.v1alpha1.insights.Insights/DeleteInsight"
	// InsightsGetInsightProcedure is the fully-qualified name of the Insights's GetInsight RPC.
	InsightsGetInsightProcedure = "/api.v1alpha1.insights.Insights/GetInsight"
	// InsightsCreateCommonsInsightProcedure is the fully-qualified name of the Insights's
	// CreateCommonsInsight RPC.
	InsightsCreateCommonsInsightProcedure = "/api.v1alpha1.insights.Insights/CreateCommonsInsight"
	// InsightsUpdateCommonsInsightProcedure is the fully-qualified name of the Insights's
	// UpdateCommonsInsight RPC.
	InsightsUpdateCommonsInsightProcedure = "/api.v1alpha1.insights.Insights/UpdateCommonsInsight"
	// InsightsDeleteCommonsInsightProcedure is the fully-qualified name of the Insights's
	// DeleteCommonsInsight RPC.
	InsightsDeleteCommonsInsightProcedure = "/api.v1alpha1.insights.Insights/DeleteCommonsInsight"
	// InsightsGetVfsSchemaProcedure is the fully-qualified name of the Insights's GetVfsSchema RPC.
	InsightsGetVfsSchemaProcedure = "/api.v1alpha1.insights.Insights/GetVfsSchema"
	// InsightsListVfsesProcedure is the fully-qualified name of the Insights's ListVfses RPC.
	InsightsListVfsesProcedure = "/api.v1alpha1.insights.Insights/ListVfses"
	// InsightsListVfsSchemasProcedure is the fully-qualified name of the Insights's ListVfsSchemas RPC.
	InsightsListVfsSchemasProcedure = "/api.v1alpha1.insights.Insights/ListVfsSchemas"
	// InsightsPublishInsightProcedure is the fully-qualified name of the Insights's PublishInsight RPC.
	InsightsPublishInsightProcedure = "/api.v1alpha1.insights.Insights/PublishInsight"
)

// InsightsClient is a client for the api.v1alpha1.insights.Insights service.
type InsightsClient interface {
	// CreateInsight creates a new insight
	CreateInsight(context.Context, *connect_go.Request[insights.CreateInsightRequest]) (*connect_go.Response[insights.CreateInsightResponse], error)
	// ListInsights lists insights
	ListInsights(context.Context, *connect_go.Request[insights.ListInsightsRequest]) (*connect_go.Response[insights.ListInsightsResponse], error)
	// ListOrgInsights lists insights for an org. Used for support app.
	ListOrgInsights(context.Context, *connect_go.Request[insights.ListOrgInsightsRequest]) (*connect_go.Response[insights.ListOrgInsightsResponse], error)
	// UpdateInsight updates an existing insight
	UpdateInsight(context.Context, *connect_go.Request[insights.UpdateInsightRequest]) (*connect_go.Response[insights.UpdateInsightResponse], error)
	// DeleteInsight deletes a insight
	DeleteInsight(context.Context, *connect_go.Request[insights.DeleteInsightRequest]) (*connect_go.Response[insights.DeleteInsightResponse], error)
	// GetInsight gets a insight by id
	GetInsight(context.Context, *connect_go.Request[insights.GetInsightRequest]) (*connect_go.Response[insights.GetInsightResponse], error)
	// CreateCommonsInsight is deprecated.
	CreateCommonsInsight(context.Context, *connect_go.Request[insights.CreateInsightRequest]) (*connect_go.Response[insights.CreateInsightResponse], error)
	// UpdateCommonsInsight is deprecated.
	UpdateCommonsInsight(context.Context, *connect_go.Request[insights.UpdateInsightRequest]) (*connect_go.Response[insights.UpdateInsightResponse], error)
	// DeleteCommonsInsight is deprecated.
	DeleteCommonsInsight(context.Context, *connect_go.Request[insights.DeleteInsightRequest]) (*connect_go.Response[insights.DeleteInsightResponse], error)
	// GetVfsSchema gets schema for a vfs
	GetVfsSchema(context.Context, *connect_go.Request[insights.GetVfsSchemaRequest]) (*connect_go.Response[insights.GetVfsSchemaResponse], error)
	// ListVfses lists exported vfs aliases
	ListVfses(context.Context, *connect_go.Request[insights.ListVfsesRequest]) (*connect_go.Response[insights.ListVfsesResponse], error)
	// ListVfses lists exported vfs aliases
	ListVfsSchemas(context.Context, *connect_go.Request[insights.ListVfsSchemasRequest]) (*connect_go.Response[insights.ListVfsSchemasResponse], error)
	// PublishInsight publishes an insight
	PublishInsight(context.Context, *connect_go.Request[insights.PublishInsightRequest]) (*connect_go.Response[insights.PublishInsightResponse], error)
}

// NewInsightsClient constructs a client for the api.v1alpha1.insights.Insights service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewInsightsClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) InsightsClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &insightsClient{
		createInsight: connect_go.NewClient[insights.CreateInsightRequest, insights.CreateInsightResponse](
			httpClient,
			baseURL+InsightsCreateInsightProcedure,
			opts...,
		),
		listInsights: connect_go.NewClient[insights.ListInsightsRequest, insights.ListInsightsResponse](
			httpClient,
			baseURL+InsightsListInsightsProcedure,
			opts...,
		),
		listOrgInsights: connect_go.NewClient[insights.ListOrgInsightsRequest, insights.ListOrgInsightsResponse](
			httpClient,
			baseURL+InsightsListOrgInsightsProcedure,
			opts...,
		),
		updateInsight: connect_go.NewClient[insights.UpdateInsightRequest, insights.UpdateInsightResponse](
			httpClient,
			baseURL+InsightsUpdateInsightProcedure,
			opts...,
		),
		deleteInsight: connect_go.NewClient[insights.DeleteInsightRequest, insights.DeleteInsightResponse](
			httpClient,
			baseURL+InsightsDeleteInsightProcedure,
			opts...,
		),
		getInsight: connect_go.NewClient[insights.GetInsightRequest, insights.GetInsightResponse](
			httpClient,
			baseURL+InsightsGetInsightProcedure,
			opts...,
		),
		createCommonsInsight: connect_go.NewClient[insights.CreateInsightRequest, insights.CreateInsightResponse](
			httpClient,
			baseURL+InsightsCreateCommonsInsightProcedure,
			opts...,
		),
		updateCommonsInsight: connect_go.NewClient[insights.UpdateInsightRequest, insights.UpdateInsightResponse](
			httpClient,
			baseURL+InsightsUpdateCommonsInsightProcedure,
			opts...,
		),
		deleteCommonsInsight: connect_go.NewClient[insights.DeleteInsightRequest, insights.DeleteInsightResponse](
			httpClient,
			baseURL+InsightsDeleteCommonsInsightProcedure,
			opts...,
		),
		getVfsSchema: connect_go.NewClient[insights.GetVfsSchemaRequest, insights.GetVfsSchemaResponse](
			httpClient,
			baseURL+InsightsGetVfsSchemaProcedure,
			opts...,
		),
		listVfses: connect_go.NewClient[insights.ListVfsesRequest, insights.ListVfsesResponse](
			httpClient,
			baseURL+InsightsListVfsesProcedure,
			opts...,
		),
		listVfsSchemas: connect_go.NewClient[insights.ListVfsSchemasRequest, insights.ListVfsSchemasResponse](
			httpClient,
			baseURL+InsightsListVfsSchemasProcedure,
			opts...,
		),
		publishInsight: connect_go.NewClient[insights.PublishInsightRequest, insights.PublishInsightResponse](
			httpClient,
			baseURL+InsightsPublishInsightProcedure,
			opts...,
		),
	}
}

// insightsClient implements InsightsClient.
type insightsClient struct {
	createInsight        *connect_go.Client[insights.CreateInsightRequest, insights.CreateInsightResponse]
	listInsights         *connect_go.Client[insights.ListInsightsRequest, insights.ListInsightsResponse]
	listOrgInsights      *connect_go.Client[insights.ListOrgInsightsRequest, insights.ListOrgInsightsResponse]
	updateInsight        *connect_go.Client[insights.UpdateInsightRequest, insights.UpdateInsightResponse]
	deleteInsight        *connect_go.Client[insights.DeleteInsightRequest, insights.DeleteInsightResponse]
	getInsight           *connect_go.Client[insights.GetInsightRequest, insights.GetInsightResponse]
	createCommonsInsight *connect_go.Client[insights.CreateInsightRequest, insights.CreateInsightResponse]
	updateCommonsInsight *connect_go.Client[insights.UpdateInsightRequest, insights.UpdateInsightResponse]
	deleteCommonsInsight *connect_go.Client[insights.DeleteInsightRequest, insights.DeleteInsightResponse]
	getVfsSchema         *connect_go.Client[insights.GetVfsSchemaRequest, insights.GetVfsSchemaResponse]
	listVfses            *connect_go.Client[insights.ListVfsesRequest, insights.ListVfsesResponse]
	listVfsSchemas       *connect_go.Client[insights.ListVfsSchemasRequest, insights.ListVfsSchemasResponse]
	publishInsight       *connect_go.Client[insights.PublishInsightRequest, insights.PublishInsightResponse]
}

// CreateInsight calls api.v1alpha1.insights.Insights.CreateInsight.
func (c *insightsClient) CreateInsight(ctx context.Context, req *connect_go.Request[insights.CreateInsightRequest]) (*connect_go.Response[insights.CreateInsightResponse], error) {
	return c.createInsight.CallUnary(ctx, req)
}

// ListInsights calls api.v1alpha1.insights.Insights.ListInsights.
func (c *insightsClient) ListInsights(ctx context.Context, req *connect_go.Request[insights.ListInsightsRequest]) (*connect_go.Response[insights.ListInsightsResponse], error) {
	return c.listInsights.CallUnary(ctx, req)
}

// ListOrgInsights calls api.v1alpha1.insights.Insights.ListOrgInsights.
func (c *insightsClient) ListOrgInsights(ctx context.Context, req *connect_go.Request[insights.ListOrgInsightsRequest]) (*connect_go.Response[insights.ListOrgInsightsResponse], error) {
	return c.listOrgInsights.CallUnary(ctx, req)
}

// UpdateInsight calls api.v1alpha1.insights.Insights.UpdateInsight.
func (c *insightsClient) UpdateInsight(ctx context.Context, req *connect_go.Request[insights.UpdateInsightRequest]) (*connect_go.Response[insights.UpdateInsightResponse], error) {
	return c.updateInsight.CallUnary(ctx, req)
}

// DeleteInsight calls api.v1alpha1.insights.Insights.DeleteInsight.
func (c *insightsClient) DeleteInsight(ctx context.Context, req *connect_go.Request[insights.DeleteInsightRequest]) (*connect_go.Response[insights.DeleteInsightResponse], error) {
	return c.deleteInsight.CallUnary(ctx, req)
}

// GetInsight calls api.v1alpha1.insights.Insights.GetInsight.
func (c *insightsClient) GetInsight(ctx context.Context, req *connect_go.Request[insights.GetInsightRequest]) (*connect_go.Response[insights.GetInsightResponse], error) {
	return c.getInsight.CallUnary(ctx, req)
}

// CreateCommonsInsight calls api.v1alpha1.insights.Insights.CreateCommonsInsight.
func (c *insightsClient) CreateCommonsInsight(ctx context.Context, req *connect_go.Request[insights.CreateInsightRequest]) (*connect_go.Response[insights.CreateInsightResponse], error) {
	return c.createCommonsInsight.CallUnary(ctx, req)
}

// UpdateCommonsInsight calls api.v1alpha1.insights.Insights.UpdateCommonsInsight.
func (c *insightsClient) UpdateCommonsInsight(ctx context.Context, req *connect_go.Request[insights.UpdateInsightRequest]) (*connect_go.Response[insights.UpdateInsightResponse], error) {
	return c.updateCommonsInsight.CallUnary(ctx, req)
}

// DeleteCommonsInsight calls api.v1alpha1.insights.Insights.DeleteCommonsInsight.
func (c *insightsClient) DeleteCommonsInsight(ctx context.Context, req *connect_go.Request[insights.DeleteInsightRequest]) (*connect_go.Response[insights.DeleteInsightResponse], error) {
	return c.deleteCommonsInsight.CallUnary(ctx, req)
}

// GetVfsSchema calls api.v1alpha1.insights.Insights.GetVfsSchema.
func (c *insightsClient) GetVfsSchema(ctx context.Context, req *connect_go.Request[insights.GetVfsSchemaRequest]) (*connect_go.Response[insights.GetVfsSchemaResponse], error) {
	return c.getVfsSchema.CallUnary(ctx, req)
}

// ListVfses calls api.v1alpha1.insights.Insights.ListVfses.
func (c *insightsClient) ListVfses(ctx context.Context, req *connect_go.Request[insights.ListVfsesRequest]) (*connect_go.Response[insights.ListVfsesResponse], error) {
	return c.listVfses.CallUnary(ctx, req)
}

// ListVfsSchemas calls api.v1alpha1.insights.Insights.ListVfsSchemas.
func (c *insightsClient) ListVfsSchemas(ctx context.Context, req *connect_go.Request[insights.ListVfsSchemasRequest]) (*connect_go.Response[insights.ListVfsSchemasResponse], error) {
	return c.listVfsSchemas.CallUnary(ctx, req)
}

// PublishInsight calls api.v1alpha1.insights.Insights.PublishInsight.
func (c *insightsClient) PublishInsight(ctx context.Context, req *connect_go.Request[insights.PublishInsightRequest]) (*connect_go.Response[insights.PublishInsightResponse], error) {
	return c.publishInsight.CallUnary(ctx, req)
}

// InsightsHandler is an implementation of the api.v1alpha1.insights.Insights service.
type InsightsHandler interface {
	// CreateInsight creates a new insight
	CreateInsight(context.Context, *connect_go.Request[insights.CreateInsightRequest]) (*connect_go.Response[insights.CreateInsightResponse], error)
	// ListInsights lists insights
	ListInsights(context.Context, *connect_go.Request[insights.ListInsightsRequest]) (*connect_go.Response[insights.ListInsightsResponse], error)
	// ListOrgInsights lists insights for an org. Used for support app.
	ListOrgInsights(context.Context, *connect_go.Request[insights.ListOrgInsightsRequest]) (*connect_go.Response[insights.ListOrgInsightsResponse], error)
	// UpdateInsight updates an existing insight
	UpdateInsight(context.Context, *connect_go.Request[insights.UpdateInsightRequest]) (*connect_go.Response[insights.UpdateInsightResponse], error)
	// DeleteInsight deletes a insight
	DeleteInsight(context.Context, *connect_go.Request[insights.DeleteInsightRequest]) (*connect_go.Response[insights.DeleteInsightResponse], error)
	// GetInsight gets a insight by id
	GetInsight(context.Context, *connect_go.Request[insights.GetInsightRequest]) (*connect_go.Response[insights.GetInsightResponse], error)
	// CreateCommonsInsight is deprecated.
	CreateCommonsInsight(context.Context, *connect_go.Request[insights.CreateInsightRequest]) (*connect_go.Response[insights.CreateInsightResponse], error)
	// UpdateCommonsInsight is deprecated.
	UpdateCommonsInsight(context.Context, *connect_go.Request[insights.UpdateInsightRequest]) (*connect_go.Response[insights.UpdateInsightResponse], error)
	// DeleteCommonsInsight is deprecated.
	DeleteCommonsInsight(context.Context, *connect_go.Request[insights.DeleteInsightRequest]) (*connect_go.Response[insights.DeleteInsightResponse], error)
	// GetVfsSchema gets schema for a vfs
	GetVfsSchema(context.Context, *connect_go.Request[insights.GetVfsSchemaRequest]) (*connect_go.Response[insights.GetVfsSchemaResponse], error)
	// ListVfses lists exported vfs aliases
	ListVfses(context.Context, *connect_go.Request[insights.ListVfsesRequest]) (*connect_go.Response[insights.ListVfsesResponse], error)
	// ListVfses lists exported vfs aliases
	ListVfsSchemas(context.Context, *connect_go.Request[insights.ListVfsSchemasRequest]) (*connect_go.Response[insights.ListVfsSchemasResponse], error)
	// PublishInsight publishes an insight
	PublishInsight(context.Context, *connect_go.Request[insights.PublishInsightRequest]) (*connect_go.Response[insights.PublishInsightResponse], error)
}

// NewInsightsHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewInsightsHandler(svc InsightsHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	insightsCreateInsightHandler := connect_go.NewUnaryHandler(
		InsightsCreateInsightProcedure,
		svc.CreateInsight,
		opts...,
	)
	insightsListInsightsHandler := connect_go.NewUnaryHandler(
		InsightsListInsightsProcedure,
		svc.ListInsights,
		opts...,
	)
	insightsListOrgInsightsHandler := connect_go.NewUnaryHandler(
		InsightsListOrgInsightsProcedure,
		svc.ListOrgInsights,
		opts...,
	)
	insightsUpdateInsightHandler := connect_go.NewUnaryHandler(
		InsightsUpdateInsightProcedure,
		svc.UpdateInsight,
		opts...,
	)
	insightsDeleteInsightHandler := connect_go.NewUnaryHandler(
		InsightsDeleteInsightProcedure,
		svc.DeleteInsight,
		opts...,
	)
	insightsGetInsightHandler := connect_go.NewUnaryHandler(
		InsightsGetInsightProcedure,
		svc.GetInsight,
		opts...,
	)
	insightsCreateCommonsInsightHandler := connect_go.NewUnaryHandler(
		InsightsCreateCommonsInsightProcedure,
		svc.CreateCommonsInsight,
		opts...,
	)
	insightsUpdateCommonsInsightHandler := connect_go.NewUnaryHandler(
		InsightsUpdateCommonsInsightProcedure,
		svc.UpdateCommonsInsight,
		opts...,
	)
	insightsDeleteCommonsInsightHandler := connect_go.NewUnaryHandler(
		InsightsDeleteCommonsInsightProcedure,
		svc.DeleteCommonsInsight,
		opts...,
	)
	insightsGetVfsSchemaHandler := connect_go.NewUnaryHandler(
		InsightsGetVfsSchemaProcedure,
		svc.GetVfsSchema,
		opts...,
	)
	insightsListVfsesHandler := connect_go.NewUnaryHandler(
		InsightsListVfsesProcedure,
		svc.ListVfses,
		opts...,
	)
	insightsListVfsSchemasHandler := connect_go.NewUnaryHandler(
		InsightsListVfsSchemasProcedure,
		svc.ListVfsSchemas,
		opts...,
	)
	insightsPublishInsightHandler := connect_go.NewUnaryHandler(
		InsightsPublishInsightProcedure,
		svc.PublishInsight,
		opts...,
	)
	return "/api.v1alpha1.insights.Insights/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case InsightsCreateInsightProcedure:
			insightsCreateInsightHandler.ServeHTTP(w, r)
		case InsightsListInsightsProcedure:
			insightsListInsightsHandler.ServeHTTP(w, r)
		case InsightsListOrgInsightsProcedure:
			insightsListOrgInsightsHandler.ServeHTTP(w, r)
		case InsightsUpdateInsightProcedure:
			insightsUpdateInsightHandler.ServeHTTP(w, r)
		case InsightsDeleteInsightProcedure:
			insightsDeleteInsightHandler.ServeHTTP(w, r)
		case InsightsGetInsightProcedure:
			insightsGetInsightHandler.ServeHTTP(w, r)
		case InsightsCreateCommonsInsightProcedure:
			insightsCreateCommonsInsightHandler.ServeHTTP(w, r)
		case InsightsUpdateCommonsInsightProcedure:
			insightsUpdateCommonsInsightHandler.ServeHTTP(w, r)
		case InsightsDeleteCommonsInsightProcedure:
			insightsDeleteCommonsInsightHandler.ServeHTTP(w, r)
		case InsightsGetVfsSchemaProcedure:
			insightsGetVfsSchemaHandler.ServeHTTP(w, r)
		case InsightsListVfsesProcedure:
			insightsListVfsesHandler.ServeHTTP(w, r)
		case InsightsListVfsSchemasProcedure:
			insightsListVfsSchemasHandler.ServeHTTP(w, r)
		case InsightsPublishInsightProcedure:
			insightsPublishInsightHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedInsightsHandler returns CodeUnimplemented from all methods.
type UnimplementedInsightsHandler struct{}

func (UnimplementedInsightsHandler) CreateInsight(context.Context, *connect_go.Request[insights.CreateInsightRequest]) (*connect_go.Response[insights.CreateInsightResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.CreateInsight is not implemented"))
}

func (UnimplementedInsightsHandler) ListInsights(context.Context, *connect_go.Request[insights.ListInsightsRequest]) (*connect_go.Response[insights.ListInsightsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.ListInsights is not implemented"))
}

func (UnimplementedInsightsHandler) ListOrgInsights(context.Context, *connect_go.Request[insights.ListOrgInsightsRequest]) (*connect_go.Response[insights.ListOrgInsightsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.ListOrgInsights is not implemented"))
}

func (UnimplementedInsightsHandler) UpdateInsight(context.Context, *connect_go.Request[insights.UpdateInsightRequest]) (*connect_go.Response[insights.UpdateInsightResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.UpdateInsight is not implemented"))
}

func (UnimplementedInsightsHandler) DeleteInsight(context.Context, *connect_go.Request[insights.DeleteInsightRequest]) (*connect_go.Response[insights.DeleteInsightResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.DeleteInsight is not implemented"))
}

func (UnimplementedInsightsHandler) GetInsight(context.Context, *connect_go.Request[insights.GetInsightRequest]) (*connect_go.Response[insights.GetInsightResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.GetInsight is not implemented"))
}

func (UnimplementedInsightsHandler) CreateCommonsInsight(context.Context, *connect_go.Request[insights.CreateInsightRequest]) (*connect_go.Response[insights.CreateInsightResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.CreateCommonsInsight is not implemented"))
}

func (UnimplementedInsightsHandler) UpdateCommonsInsight(context.Context, *connect_go.Request[insights.UpdateInsightRequest]) (*connect_go.Response[insights.UpdateInsightResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.UpdateCommonsInsight is not implemented"))
}

func (UnimplementedInsightsHandler) DeleteCommonsInsight(context.Context, *connect_go.Request[insights.DeleteInsightRequest]) (*connect_go.Response[insights.DeleteInsightResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.DeleteCommonsInsight is not implemented"))
}

func (UnimplementedInsightsHandler) GetVfsSchema(context.Context, *connect_go.Request[insights.GetVfsSchemaRequest]) (*connect_go.Response[insights.GetVfsSchemaResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.GetVfsSchema is not implemented"))
}

func (UnimplementedInsightsHandler) ListVfses(context.Context, *connect_go.Request[insights.ListVfsesRequest]) (*connect_go.Response[insights.ListVfsesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.ListVfses is not implemented"))
}

func (UnimplementedInsightsHandler) ListVfsSchemas(context.Context, *connect_go.Request[insights.ListVfsSchemasRequest]) (*connect_go.Response[insights.ListVfsSchemasResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.ListVfsSchemas is not implemented"))
}

func (UnimplementedInsightsHandler) PublishInsight(context.Context, *connect_go.Request[insights.PublishInsightRequest]) (*connect_go.Response[insights.PublishInsightResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.insights.Insights.PublishInsight is not implemented"))
}
