// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/scorecards/support_service.proto

package scorecardsconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	scorecards "github.com/tcncloud/api-go/api/v1alpha1/scorecards"
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
	// ScorecardsSupportName is the fully-qualified name of the ScorecardsSupport service.
	ScorecardsSupportName = "api.v1alpha1.scorecards.ScorecardsSupport"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ScorecardsSupportListEvaluationsByOrgIdProcedure is the fully-qualified name of the
	// ScorecardsSupport's ListEvaluationsByOrgId RPC.
	ScorecardsSupportListEvaluationsByOrgIdProcedure = "/api.v1alpha1.scorecards.ScorecardsSupport/ListEvaluationsByOrgId"
	// ScorecardsSupportListAutoEvaluationsByOrgIdProcedure is the fully-qualified name of the
	// ScorecardsSupport's ListAutoEvaluationsByOrgId RPC.
	ScorecardsSupportListAutoEvaluationsByOrgIdProcedure = "/api.v1alpha1.scorecards.ScorecardsSupport/ListAutoEvaluationsByOrgId"
	// ScorecardsSupportBulkDeleteEvaluationsProcedure is the fully-qualified name of the
	// ScorecardsSupport's BulkDeleteEvaluations RPC.
	ScorecardsSupportBulkDeleteEvaluationsProcedure = "/api.v1alpha1.scorecards.ScorecardsSupport/BulkDeleteEvaluations"
	// ScorecardsSupportBulkDeleteAutoEvaluationsProcedure is the fully-qualified name of the
	// ScorecardsSupport's BulkDeleteAutoEvaluations RPC.
	ScorecardsSupportBulkDeleteAutoEvaluationsProcedure = "/api.v1alpha1.scorecards.ScorecardsSupport/BulkDeleteAutoEvaluations"
	// ScorecardsSupportDeleteEvaluationByOrgIdProcedure is the fully-qualified name of the
	// ScorecardsSupport's DeleteEvaluationByOrgId RPC.
	ScorecardsSupportDeleteEvaluationByOrgIdProcedure = "/api.v1alpha1.scorecards.ScorecardsSupport/DeleteEvaluationByOrgId"
	// ScorecardsSupportDeleteAutoEvaluationByOrgIdProcedure is the fully-qualified name of the
	// ScorecardsSupport's DeleteAutoEvaluationByOrgId RPC.
	ScorecardsSupportDeleteAutoEvaluationByOrgIdProcedure = "/api.v1alpha1.scorecards.ScorecardsSupport/DeleteAutoEvaluationByOrgId"
	// ScorecardsSupportListScorecardsByOrgIdProcedure is the fully-qualified name of the
	// ScorecardsSupport's ListScorecardsByOrgId RPC.
	ScorecardsSupportListScorecardsByOrgIdProcedure = "/api.v1alpha1.scorecards.ScorecardsSupport/ListScorecardsByOrgId"
	// ScorecardsSupportListCategoriesByOrgIdProcedure is the fully-qualified name of the
	// ScorecardsSupport's ListCategoriesByOrgId RPC.
	ScorecardsSupportListCategoriesByOrgIdProcedure = "/api.v1alpha1.scorecards.ScorecardsSupport/ListCategoriesByOrgId"
)

// ScorecardsSupportClient is a client for the api.v1alpha1.scorecards.ScorecardsSupport service.
type ScorecardsSupportClient interface {
	// ListEvaluationsByOrgId gets a list of evaluations by org id
	ListEvaluationsByOrgId(context.Context, *connect_go.Request[scorecards.ListEvaluationsByOrgIdRequest]) (*connect_go.Response[scorecards.ListEvaluationsResponse], error)
	// ListAutoEvaluationsByOrgId gets a list of auto evaluations
	ListAutoEvaluationsByOrgId(context.Context, *connect_go.Request[scorecards.ListAutoEvaluationsByOrgIdRequest]) (*connect_go.Response[scorecards.ListAutoEvaluationsResponse], error)
	// BulkDeleteEvaluations deletes a set of evaluations in a given org.
	BulkDeleteEvaluations(context.Context, *connect_go.Request[scorecards.BulkDeleteEvaluationsRequest]) (*connect_go.Response[scorecards.BulkDeleteEvaluationsResponse], error)
	// BulkDeleteAutoEvaluations deletes a set of auto evaluations in a given org.
	BulkDeleteAutoEvaluations(context.Context, *connect_go.Request[scorecards.BulkDeleteAutoEvaluationsRequest]) (*connect_go.Response[scorecards.BulkDeleteAutoEvaluationsResponse], error)
	// DeleteEvaluationByOrgId delete an evaluation in a specific org
	DeleteEvaluationByOrgId(context.Context, *connect_go.Request[scorecards.DeleteEvaluationByOrgIdRequest]) (*connect_go.Response[scorecards.DeleteEvaluationResponse], error)
	// DeleteAutoEvaluationByOrgId deletes an auto evaluations
	DeleteAutoEvaluationByOrgId(context.Context, *connect_go.Request[scorecards.DeleteAutoEvaluationByOrgIdRequest]) (*connect_go.Response[scorecards.DeleteAutoEvaluationResponse], error)
	// ListScorecardsByOrgId lists scorecards
	ListScorecardsByOrgId(context.Context, *connect_go.Request[scorecards.ListScorecardsByOrgIdRequest]) (*connect_go.Response[scorecards.ListScorecardsResponse], error)
	// ListCategoriesByOrgId lists categories
	ListCategoriesByOrgId(context.Context, *connect_go.Request[scorecards.ListCategoriesByOrgIdRequest]) (*connect_go.Response[scorecards.ListCategoriesResponse], error)
}

// NewScorecardsSupportClient constructs a client for the api.v1alpha1.scorecards.ScorecardsSupport
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewScorecardsSupportClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ScorecardsSupportClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &scorecardsSupportClient{
		listEvaluationsByOrgId: connect_go.NewClient[scorecards.ListEvaluationsByOrgIdRequest, scorecards.ListEvaluationsResponse](
			httpClient,
			baseURL+ScorecardsSupportListEvaluationsByOrgIdProcedure,
			opts...,
		),
		listAutoEvaluationsByOrgId: connect_go.NewClient[scorecards.ListAutoEvaluationsByOrgIdRequest, scorecards.ListAutoEvaluationsResponse](
			httpClient,
			baseURL+ScorecardsSupportListAutoEvaluationsByOrgIdProcedure,
			opts...,
		),
		bulkDeleteEvaluations: connect_go.NewClient[scorecards.BulkDeleteEvaluationsRequest, scorecards.BulkDeleteEvaluationsResponse](
			httpClient,
			baseURL+ScorecardsSupportBulkDeleteEvaluationsProcedure,
			opts...,
		),
		bulkDeleteAutoEvaluations: connect_go.NewClient[scorecards.BulkDeleteAutoEvaluationsRequest, scorecards.BulkDeleteAutoEvaluationsResponse](
			httpClient,
			baseURL+ScorecardsSupportBulkDeleteAutoEvaluationsProcedure,
			opts...,
		),
		deleteEvaluationByOrgId: connect_go.NewClient[scorecards.DeleteEvaluationByOrgIdRequest, scorecards.DeleteEvaluationResponse](
			httpClient,
			baseURL+ScorecardsSupportDeleteEvaluationByOrgIdProcedure,
			opts...,
		),
		deleteAutoEvaluationByOrgId: connect_go.NewClient[scorecards.DeleteAutoEvaluationByOrgIdRequest, scorecards.DeleteAutoEvaluationResponse](
			httpClient,
			baseURL+ScorecardsSupportDeleteAutoEvaluationByOrgIdProcedure,
			opts...,
		),
		listScorecardsByOrgId: connect_go.NewClient[scorecards.ListScorecardsByOrgIdRequest, scorecards.ListScorecardsResponse](
			httpClient,
			baseURL+ScorecardsSupportListScorecardsByOrgIdProcedure,
			opts...,
		),
		listCategoriesByOrgId: connect_go.NewClient[scorecards.ListCategoriesByOrgIdRequest, scorecards.ListCategoriesResponse](
			httpClient,
			baseURL+ScorecardsSupportListCategoriesByOrgIdProcedure,
			opts...,
		),
	}
}

// scorecardsSupportClient implements ScorecardsSupportClient.
type scorecardsSupportClient struct {
	listEvaluationsByOrgId      *connect_go.Client[scorecards.ListEvaluationsByOrgIdRequest, scorecards.ListEvaluationsResponse]
	listAutoEvaluationsByOrgId  *connect_go.Client[scorecards.ListAutoEvaluationsByOrgIdRequest, scorecards.ListAutoEvaluationsResponse]
	bulkDeleteEvaluations       *connect_go.Client[scorecards.BulkDeleteEvaluationsRequest, scorecards.BulkDeleteEvaluationsResponse]
	bulkDeleteAutoEvaluations   *connect_go.Client[scorecards.BulkDeleteAutoEvaluationsRequest, scorecards.BulkDeleteAutoEvaluationsResponse]
	deleteEvaluationByOrgId     *connect_go.Client[scorecards.DeleteEvaluationByOrgIdRequest, scorecards.DeleteEvaluationResponse]
	deleteAutoEvaluationByOrgId *connect_go.Client[scorecards.DeleteAutoEvaluationByOrgIdRequest, scorecards.DeleteAutoEvaluationResponse]
	listScorecardsByOrgId       *connect_go.Client[scorecards.ListScorecardsByOrgIdRequest, scorecards.ListScorecardsResponse]
	listCategoriesByOrgId       *connect_go.Client[scorecards.ListCategoriesByOrgIdRequest, scorecards.ListCategoriesResponse]
}

// ListEvaluationsByOrgId calls api.v1alpha1.scorecards.ScorecardsSupport.ListEvaluationsByOrgId.
func (c *scorecardsSupportClient) ListEvaluationsByOrgId(ctx context.Context, req *connect_go.Request[scorecards.ListEvaluationsByOrgIdRequest]) (*connect_go.Response[scorecards.ListEvaluationsResponse], error) {
	return c.listEvaluationsByOrgId.CallUnary(ctx, req)
}

// ListAutoEvaluationsByOrgId calls
// api.v1alpha1.scorecards.ScorecardsSupport.ListAutoEvaluationsByOrgId.
func (c *scorecardsSupportClient) ListAutoEvaluationsByOrgId(ctx context.Context, req *connect_go.Request[scorecards.ListAutoEvaluationsByOrgIdRequest]) (*connect_go.Response[scorecards.ListAutoEvaluationsResponse], error) {
	return c.listAutoEvaluationsByOrgId.CallUnary(ctx, req)
}

// BulkDeleteEvaluations calls api.v1alpha1.scorecards.ScorecardsSupport.BulkDeleteEvaluations.
func (c *scorecardsSupportClient) BulkDeleteEvaluations(ctx context.Context, req *connect_go.Request[scorecards.BulkDeleteEvaluationsRequest]) (*connect_go.Response[scorecards.BulkDeleteEvaluationsResponse], error) {
	return c.bulkDeleteEvaluations.CallUnary(ctx, req)
}

// BulkDeleteAutoEvaluations calls
// api.v1alpha1.scorecards.ScorecardsSupport.BulkDeleteAutoEvaluations.
func (c *scorecardsSupportClient) BulkDeleteAutoEvaluations(ctx context.Context, req *connect_go.Request[scorecards.BulkDeleteAutoEvaluationsRequest]) (*connect_go.Response[scorecards.BulkDeleteAutoEvaluationsResponse], error) {
	return c.bulkDeleteAutoEvaluations.CallUnary(ctx, req)
}

// DeleteEvaluationByOrgId calls api.v1alpha1.scorecards.ScorecardsSupport.DeleteEvaluationByOrgId.
func (c *scorecardsSupportClient) DeleteEvaluationByOrgId(ctx context.Context, req *connect_go.Request[scorecards.DeleteEvaluationByOrgIdRequest]) (*connect_go.Response[scorecards.DeleteEvaluationResponse], error) {
	return c.deleteEvaluationByOrgId.CallUnary(ctx, req)
}

// DeleteAutoEvaluationByOrgId calls
// api.v1alpha1.scorecards.ScorecardsSupport.DeleteAutoEvaluationByOrgId.
func (c *scorecardsSupportClient) DeleteAutoEvaluationByOrgId(ctx context.Context, req *connect_go.Request[scorecards.DeleteAutoEvaluationByOrgIdRequest]) (*connect_go.Response[scorecards.DeleteAutoEvaluationResponse], error) {
	return c.deleteAutoEvaluationByOrgId.CallUnary(ctx, req)
}

// ListScorecardsByOrgId calls api.v1alpha1.scorecards.ScorecardsSupport.ListScorecardsByOrgId.
func (c *scorecardsSupportClient) ListScorecardsByOrgId(ctx context.Context, req *connect_go.Request[scorecards.ListScorecardsByOrgIdRequest]) (*connect_go.Response[scorecards.ListScorecardsResponse], error) {
	return c.listScorecardsByOrgId.CallUnary(ctx, req)
}

// ListCategoriesByOrgId calls api.v1alpha1.scorecards.ScorecardsSupport.ListCategoriesByOrgId.
func (c *scorecardsSupportClient) ListCategoriesByOrgId(ctx context.Context, req *connect_go.Request[scorecards.ListCategoriesByOrgIdRequest]) (*connect_go.Response[scorecards.ListCategoriesResponse], error) {
	return c.listCategoriesByOrgId.CallUnary(ctx, req)
}

// ScorecardsSupportHandler is an implementation of the api.v1alpha1.scorecards.ScorecardsSupport
// service.
type ScorecardsSupportHandler interface {
	// ListEvaluationsByOrgId gets a list of evaluations by org id
	ListEvaluationsByOrgId(context.Context, *connect_go.Request[scorecards.ListEvaluationsByOrgIdRequest]) (*connect_go.Response[scorecards.ListEvaluationsResponse], error)
	// ListAutoEvaluationsByOrgId gets a list of auto evaluations
	ListAutoEvaluationsByOrgId(context.Context, *connect_go.Request[scorecards.ListAutoEvaluationsByOrgIdRequest]) (*connect_go.Response[scorecards.ListAutoEvaluationsResponse], error)
	// BulkDeleteEvaluations deletes a set of evaluations in a given org.
	BulkDeleteEvaluations(context.Context, *connect_go.Request[scorecards.BulkDeleteEvaluationsRequest]) (*connect_go.Response[scorecards.BulkDeleteEvaluationsResponse], error)
	// BulkDeleteAutoEvaluations deletes a set of auto evaluations in a given org.
	BulkDeleteAutoEvaluations(context.Context, *connect_go.Request[scorecards.BulkDeleteAutoEvaluationsRequest]) (*connect_go.Response[scorecards.BulkDeleteAutoEvaluationsResponse], error)
	// DeleteEvaluationByOrgId delete an evaluation in a specific org
	DeleteEvaluationByOrgId(context.Context, *connect_go.Request[scorecards.DeleteEvaluationByOrgIdRequest]) (*connect_go.Response[scorecards.DeleteEvaluationResponse], error)
	// DeleteAutoEvaluationByOrgId deletes an auto evaluations
	DeleteAutoEvaluationByOrgId(context.Context, *connect_go.Request[scorecards.DeleteAutoEvaluationByOrgIdRequest]) (*connect_go.Response[scorecards.DeleteAutoEvaluationResponse], error)
	// ListScorecardsByOrgId lists scorecards
	ListScorecardsByOrgId(context.Context, *connect_go.Request[scorecards.ListScorecardsByOrgIdRequest]) (*connect_go.Response[scorecards.ListScorecardsResponse], error)
	// ListCategoriesByOrgId lists categories
	ListCategoriesByOrgId(context.Context, *connect_go.Request[scorecards.ListCategoriesByOrgIdRequest]) (*connect_go.Response[scorecards.ListCategoriesResponse], error)
}

// NewScorecardsSupportHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewScorecardsSupportHandler(svc ScorecardsSupportHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	scorecardsSupportListEvaluationsByOrgIdHandler := connect_go.NewUnaryHandler(
		ScorecardsSupportListEvaluationsByOrgIdProcedure,
		svc.ListEvaluationsByOrgId,
		opts...,
	)
	scorecardsSupportListAutoEvaluationsByOrgIdHandler := connect_go.NewUnaryHandler(
		ScorecardsSupportListAutoEvaluationsByOrgIdProcedure,
		svc.ListAutoEvaluationsByOrgId,
		opts...,
	)
	scorecardsSupportBulkDeleteEvaluationsHandler := connect_go.NewUnaryHandler(
		ScorecardsSupportBulkDeleteEvaluationsProcedure,
		svc.BulkDeleteEvaluations,
		opts...,
	)
	scorecardsSupportBulkDeleteAutoEvaluationsHandler := connect_go.NewUnaryHandler(
		ScorecardsSupportBulkDeleteAutoEvaluationsProcedure,
		svc.BulkDeleteAutoEvaluations,
		opts...,
	)
	scorecardsSupportDeleteEvaluationByOrgIdHandler := connect_go.NewUnaryHandler(
		ScorecardsSupportDeleteEvaluationByOrgIdProcedure,
		svc.DeleteEvaluationByOrgId,
		opts...,
	)
	scorecardsSupportDeleteAutoEvaluationByOrgIdHandler := connect_go.NewUnaryHandler(
		ScorecardsSupportDeleteAutoEvaluationByOrgIdProcedure,
		svc.DeleteAutoEvaluationByOrgId,
		opts...,
	)
	scorecardsSupportListScorecardsByOrgIdHandler := connect_go.NewUnaryHandler(
		ScorecardsSupportListScorecardsByOrgIdProcedure,
		svc.ListScorecardsByOrgId,
		opts...,
	)
	scorecardsSupportListCategoriesByOrgIdHandler := connect_go.NewUnaryHandler(
		ScorecardsSupportListCategoriesByOrgIdProcedure,
		svc.ListCategoriesByOrgId,
		opts...,
	)
	return "/api.v1alpha1.scorecards.ScorecardsSupport/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ScorecardsSupportListEvaluationsByOrgIdProcedure:
			scorecardsSupportListEvaluationsByOrgIdHandler.ServeHTTP(w, r)
		case ScorecardsSupportListAutoEvaluationsByOrgIdProcedure:
			scorecardsSupportListAutoEvaluationsByOrgIdHandler.ServeHTTP(w, r)
		case ScorecardsSupportBulkDeleteEvaluationsProcedure:
			scorecardsSupportBulkDeleteEvaluationsHandler.ServeHTTP(w, r)
		case ScorecardsSupportBulkDeleteAutoEvaluationsProcedure:
			scorecardsSupportBulkDeleteAutoEvaluationsHandler.ServeHTTP(w, r)
		case ScorecardsSupportDeleteEvaluationByOrgIdProcedure:
			scorecardsSupportDeleteEvaluationByOrgIdHandler.ServeHTTP(w, r)
		case ScorecardsSupportDeleteAutoEvaluationByOrgIdProcedure:
			scorecardsSupportDeleteAutoEvaluationByOrgIdHandler.ServeHTTP(w, r)
		case ScorecardsSupportListScorecardsByOrgIdProcedure:
			scorecardsSupportListScorecardsByOrgIdHandler.ServeHTTP(w, r)
		case ScorecardsSupportListCategoriesByOrgIdProcedure:
			scorecardsSupportListCategoriesByOrgIdHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedScorecardsSupportHandler returns CodeUnimplemented from all methods.
type UnimplementedScorecardsSupportHandler struct{}

func (UnimplementedScorecardsSupportHandler) ListEvaluationsByOrgId(context.Context, *connect_go.Request[scorecards.ListEvaluationsByOrgIdRequest]) (*connect_go.Response[scorecards.ListEvaluationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.scorecards.ScorecardsSupport.ListEvaluationsByOrgId is not implemented"))
}

func (UnimplementedScorecardsSupportHandler) ListAutoEvaluationsByOrgId(context.Context, *connect_go.Request[scorecards.ListAutoEvaluationsByOrgIdRequest]) (*connect_go.Response[scorecards.ListAutoEvaluationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.scorecards.ScorecardsSupport.ListAutoEvaluationsByOrgId is not implemented"))
}

func (UnimplementedScorecardsSupportHandler) BulkDeleteEvaluations(context.Context, *connect_go.Request[scorecards.BulkDeleteEvaluationsRequest]) (*connect_go.Response[scorecards.BulkDeleteEvaluationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.scorecards.ScorecardsSupport.BulkDeleteEvaluations is not implemented"))
}

func (UnimplementedScorecardsSupportHandler) BulkDeleteAutoEvaluations(context.Context, *connect_go.Request[scorecards.BulkDeleteAutoEvaluationsRequest]) (*connect_go.Response[scorecards.BulkDeleteAutoEvaluationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.scorecards.ScorecardsSupport.BulkDeleteAutoEvaluations is not implemented"))
}

func (UnimplementedScorecardsSupportHandler) DeleteEvaluationByOrgId(context.Context, *connect_go.Request[scorecards.DeleteEvaluationByOrgIdRequest]) (*connect_go.Response[scorecards.DeleteEvaluationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.scorecards.ScorecardsSupport.DeleteEvaluationByOrgId is not implemented"))
}

func (UnimplementedScorecardsSupportHandler) DeleteAutoEvaluationByOrgId(context.Context, *connect_go.Request[scorecards.DeleteAutoEvaluationByOrgIdRequest]) (*connect_go.Response[scorecards.DeleteAutoEvaluationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.scorecards.ScorecardsSupport.DeleteAutoEvaluationByOrgId is not implemented"))
}

func (UnimplementedScorecardsSupportHandler) ListScorecardsByOrgId(context.Context, *connect_go.Request[scorecards.ListScorecardsByOrgIdRequest]) (*connect_go.Response[scorecards.ListScorecardsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.scorecards.ScorecardsSupport.ListScorecardsByOrgId is not implemented"))
}

func (UnimplementedScorecardsSupportHandler) ListCategoriesByOrgId(context.Context, *connect_go.Request[scorecards.ListCategoriesByOrgIdRequest]) (*connect_go.Response[scorecards.ListCategoriesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.scorecards.ScorecardsSupport.ListCategoriesByOrgId is not implemented"))
}
