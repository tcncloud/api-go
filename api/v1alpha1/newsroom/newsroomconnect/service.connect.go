// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/newsroom/service.proto

package newsroomconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	newsroom "github.com/tcncloud/api-go/api/v1alpha1/newsroom"
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
	// NewsroomAPIName is the fully-qualified name of the NewsroomAPI service.
	NewsroomAPIName = "api.v1alpha1.newsroom.NewsroomAPI"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// NewsroomAPICreateNewsArticleProcedure is the fully-qualified name of the NewsroomAPI's
	// CreateNewsArticle RPC.
	NewsroomAPICreateNewsArticleProcedure = "/api.v1alpha1.newsroom.NewsroomAPI/CreateNewsArticle"
	// NewsroomAPIListNewsArticlesProcedure is the fully-qualified name of the NewsroomAPI's
	// ListNewsArticles RPC.
	NewsroomAPIListNewsArticlesProcedure = "/api.v1alpha1.newsroom.NewsroomAPI/ListNewsArticles"
	// NewsroomAPIGetNewsArticleByIdProcedure is the fully-qualified name of the NewsroomAPI's
	// GetNewsArticleById RPC.
	NewsroomAPIGetNewsArticleByIdProcedure = "/api.v1alpha1.newsroom.NewsroomAPI/GetNewsArticleById"
	// NewsroomAPIUpdateNewsArticleProcedure is the fully-qualified name of the NewsroomAPI's
	// UpdateNewsArticle RPC.
	NewsroomAPIUpdateNewsArticleProcedure = "/api.v1alpha1.newsroom.NewsroomAPI/UpdateNewsArticle"
	// NewsroomAPICreatePublishedArticleProcedure is the fully-qualified name of the NewsroomAPI's
	// CreatePublishedArticle RPC.
	NewsroomAPICreatePublishedArticleProcedure = "/api.v1alpha1.newsroom.NewsroomAPI/CreatePublishedArticle"
	// NewsroomAPIListPublishedArticlesProcedure is the fully-qualified name of the NewsroomAPI's
	// ListPublishedArticles RPC.
	NewsroomAPIListPublishedArticlesProcedure = "/api.v1alpha1.newsroom.NewsroomAPI/ListPublishedArticles"
	// NewsroomAPIGetPublishedArticleByIdProcedure is the fully-qualified name of the NewsroomAPI's
	// GetPublishedArticleById RPC.
	NewsroomAPIGetPublishedArticleByIdProcedure = "/api.v1alpha1.newsroom.NewsroomAPI/GetPublishedArticleById"
	// NewsroomAPIUserActivityProcedure is the fully-qualified name of the NewsroomAPI's UserActivity
	// RPC.
	NewsroomAPIUserActivityProcedure = "/api.v1alpha1.newsroom.NewsroomAPI/UserActivity"
	// NewsroomAPIGetNewsForUserProcedure is the fully-qualified name of the NewsroomAPI's
	// GetNewsForUser RPC.
	NewsroomAPIGetNewsForUserProcedure = "/api.v1alpha1.newsroom.NewsroomAPI/GetNewsForUser"
)

// NewsroomAPIClient is a client for the api.v1alpha1.newsroom.NewsroomAPI service.
type NewsroomAPIClient interface {
	// create news article
	CreateNewsArticle(context.Context, *connect_go.Request[newsroom.CreateNewsArticleRequest]) (*connect_go.Response[newsroom.CreateNewsArticleResponse], error)
	// list news articles
	ListNewsArticles(context.Context, *connect_go.Request[newsroom.ListNewsArticlesRequest]) (*connect_go.Response[newsroom.ListNewsArticlesResponse], error)
	// get news article details by the id
	GetNewsArticleById(context.Context, *connect_go.Request[newsroom.GetNewsArticleByIdRequest]) (*connect_go.Response[newsroom.GetNewsArticleByIdResponse], error)
	// update news article
	UpdateNewsArticle(context.Context, *connect_go.Request[newsroom.UpdateNewsArticleRequest]) (*connect_go.Response[newsroom.UpdateNewsArticleResponse], error)
	// create published article
	CreatePublishedArticle(context.Context, *connect_go.Request[newsroom.CreatePublishedArticleRequest]) (*connect_go.Response[newsroom.CreatePublishedArticleResponse], error)
	// list published articles
	ListPublishedArticles(context.Context, *connect_go.Request[newsroom.ListPublishedArticlesRequest]) (*connect_go.Response[newsroom.ListPublishedArticlesResponse], error)
	// get published article details by the id
	GetPublishedArticleById(context.Context, *connect_go.Request[newsroom.GetPublishedArticleByIdRequest]) (*connect_go.Response[newsroom.GetPublishedArticleByIdResponse], error)
	// user activity updates
	UserActivity(context.Context, *connect_go.Request[newsroom.UserActivityRequest]) (*connect_go.Response[newsroom.UserActivityResponse], error)
	// fetch the unseen articles for the user
	GetNewsForUser(context.Context, *connect_go.Request[newsroom.GetNewsForUserRequest]) (*connect_go.Response[newsroom.GetNewsForUserResponse], error)
}

// NewNewsroomAPIClient constructs a client for the api.v1alpha1.newsroom.NewsroomAPI service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNewsroomAPIClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) NewsroomAPIClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &newsroomAPIClient{
		createNewsArticle: connect_go.NewClient[newsroom.CreateNewsArticleRequest, newsroom.CreateNewsArticleResponse](
			httpClient,
			baseURL+NewsroomAPICreateNewsArticleProcedure,
			opts...,
		),
		listNewsArticles: connect_go.NewClient[newsroom.ListNewsArticlesRequest, newsroom.ListNewsArticlesResponse](
			httpClient,
			baseURL+NewsroomAPIListNewsArticlesProcedure,
			opts...,
		),
		getNewsArticleById: connect_go.NewClient[newsroom.GetNewsArticleByIdRequest, newsroom.GetNewsArticleByIdResponse](
			httpClient,
			baseURL+NewsroomAPIGetNewsArticleByIdProcedure,
			opts...,
		),
		updateNewsArticle: connect_go.NewClient[newsroom.UpdateNewsArticleRequest, newsroom.UpdateNewsArticleResponse](
			httpClient,
			baseURL+NewsroomAPIUpdateNewsArticleProcedure,
			opts...,
		),
		createPublishedArticle: connect_go.NewClient[newsroom.CreatePublishedArticleRequest, newsroom.CreatePublishedArticleResponse](
			httpClient,
			baseURL+NewsroomAPICreatePublishedArticleProcedure,
			opts...,
		),
		listPublishedArticles: connect_go.NewClient[newsroom.ListPublishedArticlesRequest, newsroom.ListPublishedArticlesResponse](
			httpClient,
			baseURL+NewsroomAPIListPublishedArticlesProcedure,
			opts...,
		),
		getPublishedArticleById: connect_go.NewClient[newsroom.GetPublishedArticleByIdRequest, newsroom.GetPublishedArticleByIdResponse](
			httpClient,
			baseURL+NewsroomAPIGetPublishedArticleByIdProcedure,
			opts...,
		),
		userActivity: connect_go.NewClient[newsroom.UserActivityRequest, newsroom.UserActivityResponse](
			httpClient,
			baseURL+NewsroomAPIUserActivityProcedure,
			opts...,
		),
		getNewsForUser: connect_go.NewClient[newsroom.GetNewsForUserRequest, newsroom.GetNewsForUserResponse](
			httpClient,
			baseURL+NewsroomAPIGetNewsForUserProcedure,
			opts...,
		),
	}
}

// newsroomAPIClient implements NewsroomAPIClient.
type newsroomAPIClient struct {
	createNewsArticle       *connect_go.Client[newsroom.CreateNewsArticleRequest, newsroom.CreateNewsArticleResponse]
	listNewsArticles        *connect_go.Client[newsroom.ListNewsArticlesRequest, newsroom.ListNewsArticlesResponse]
	getNewsArticleById      *connect_go.Client[newsroom.GetNewsArticleByIdRequest, newsroom.GetNewsArticleByIdResponse]
	updateNewsArticle       *connect_go.Client[newsroom.UpdateNewsArticleRequest, newsroom.UpdateNewsArticleResponse]
	createPublishedArticle  *connect_go.Client[newsroom.CreatePublishedArticleRequest, newsroom.CreatePublishedArticleResponse]
	listPublishedArticles   *connect_go.Client[newsroom.ListPublishedArticlesRequest, newsroom.ListPublishedArticlesResponse]
	getPublishedArticleById *connect_go.Client[newsroom.GetPublishedArticleByIdRequest, newsroom.GetPublishedArticleByIdResponse]
	userActivity            *connect_go.Client[newsroom.UserActivityRequest, newsroom.UserActivityResponse]
	getNewsForUser          *connect_go.Client[newsroom.GetNewsForUserRequest, newsroom.GetNewsForUserResponse]
}

// CreateNewsArticle calls api.v1alpha1.newsroom.NewsroomAPI.CreateNewsArticle.
func (c *newsroomAPIClient) CreateNewsArticle(ctx context.Context, req *connect_go.Request[newsroom.CreateNewsArticleRequest]) (*connect_go.Response[newsroom.CreateNewsArticleResponse], error) {
	return c.createNewsArticle.CallUnary(ctx, req)
}

// ListNewsArticles calls api.v1alpha1.newsroom.NewsroomAPI.ListNewsArticles.
func (c *newsroomAPIClient) ListNewsArticles(ctx context.Context, req *connect_go.Request[newsroom.ListNewsArticlesRequest]) (*connect_go.Response[newsroom.ListNewsArticlesResponse], error) {
	return c.listNewsArticles.CallUnary(ctx, req)
}

// GetNewsArticleById calls api.v1alpha1.newsroom.NewsroomAPI.GetNewsArticleById.
func (c *newsroomAPIClient) GetNewsArticleById(ctx context.Context, req *connect_go.Request[newsroom.GetNewsArticleByIdRequest]) (*connect_go.Response[newsroom.GetNewsArticleByIdResponse], error) {
	return c.getNewsArticleById.CallUnary(ctx, req)
}

// UpdateNewsArticle calls api.v1alpha1.newsroom.NewsroomAPI.UpdateNewsArticle.
func (c *newsroomAPIClient) UpdateNewsArticle(ctx context.Context, req *connect_go.Request[newsroom.UpdateNewsArticleRequest]) (*connect_go.Response[newsroom.UpdateNewsArticleResponse], error) {
	return c.updateNewsArticle.CallUnary(ctx, req)
}

// CreatePublishedArticle calls api.v1alpha1.newsroom.NewsroomAPI.CreatePublishedArticle.
func (c *newsroomAPIClient) CreatePublishedArticle(ctx context.Context, req *connect_go.Request[newsroom.CreatePublishedArticleRequest]) (*connect_go.Response[newsroom.CreatePublishedArticleResponse], error) {
	return c.createPublishedArticle.CallUnary(ctx, req)
}

// ListPublishedArticles calls api.v1alpha1.newsroom.NewsroomAPI.ListPublishedArticles.
func (c *newsroomAPIClient) ListPublishedArticles(ctx context.Context, req *connect_go.Request[newsroom.ListPublishedArticlesRequest]) (*connect_go.Response[newsroom.ListPublishedArticlesResponse], error) {
	return c.listPublishedArticles.CallUnary(ctx, req)
}

// GetPublishedArticleById calls api.v1alpha1.newsroom.NewsroomAPI.GetPublishedArticleById.
func (c *newsroomAPIClient) GetPublishedArticleById(ctx context.Context, req *connect_go.Request[newsroom.GetPublishedArticleByIdRequest]) (*connect_go.Response[newsroom.GetPublishedArticleByIdResponse], error) {
	return c.getPublishedArticleById.CallUnary(ctx, req)
}

// UserActivity calls api.v1alpha1.newsroom.NewsroomAPI.UserActivity.
func (c *newsroomAPIClient) UserActivity(ctx context.Context, req *connect_go.Request[newsroom.UserActivityRequest]) (*connect_go.Response[newsroom.UserActivityResponse], error) {
	return c.userActivity.CallUnary(ctx, req)
}

// GetNewsForUser calls api.v1alpha1.newsroom.NewsroomAPI.GetNewsForUser.
func (c *newsroomAPIClient) GetNewsForUser(ctx context.Context, req *connect_go.Request[newsroom.GetNewsForUserRequest]) (*connect_go.Response[newsroom.GetNewsForUserResponse], error) {
	return c.getNewsForUser.CallUnary(ctx, req)
}

// NewsroomAPIHandler is an implementation of the api.v1alpha1.newsroom.NewsroomAPI service.
type NewsroomAPIHandler interface {
	// create news article
	CreateNewsArticle(context.Context, *connect_go.Request[newsroom.CreateNewsArticleRequest]) (*connect_go.Response[newsroom.CreateNewsArticleResponse], error)
	// list news articles
	ListNewsArticles(context.Context, *connect_go.Request[newsroom.ListNewsArticlesRequest]) (*connect_go.Response[newsroom.ListNewsArticlesResponse], error)
	// get news article details by the id
	GetNewsArticleById(context.Context, *connect_go.Request[newsroom.GetNewsArticleByIdRequest]) (*connect_go.Response[newsroom.GetNewsArticleByIdResponse], error)
	// update news article
	UpdateNewsArticle(context.Context, *connect_go.Request[newsroom.UpdateNewsArticleRequest]) (*connect_go.Response[newsroom.UpdateNewsArticleResponse], error)
	// create published article
	CreatePublishedArticle(context.Context, *connect_go.Request[newsroom.CreatePublishedArticleRequest]) (*connect_go.Response[newsroom.CreatePublishedArticleResponse], error)
	// list published articles
	ListPublishedArticles(context.Context, *connect_go.Request[newsroom.ListPublishedArticlesRequest]) (*connect_go.Response[newsroom.ListPublishedArticlesResponse], error)
	// get published article details by the id
	GetPublishedArticleById(context.Context, *connect_go.Request[newsroom.GetPublishedArticleByIdRequest]) (*connect_go.Response[newsroom.GetPublishedArticleByIdResponse], error)
	// user activity updates
	UserActivity(context.Context, *connect_go.Request[newsroom.UserActivityRequest]) (*connect_go.Response[newsroom.UserActivityResponse], error)
	// fetch the unseen articles for the user
	GetNewsForUser(context.Context, *connect_go.Request[newsroom.GetNewsForUserRequest]) (*connect_go.Response[newsroom.GetNewsForUserResponse], error)
}

// NewNewsroomAPIHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNewsroomAPIHandler(svc NewsroomAPIHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	newsroomAPICreateNewsArticleHandler := connect_go.NewUnaryHandler(
		NewsroomAPICreateNewsArticleProcedure,
		svc.CreateNewsArticle,
		opts...,
	)
	newsroomAPIListNewsArticlesHandler := connect_go.NewUnaryHandler(
		NewsroomAPIListNewsArticlesProcedure,
		svc.ListNewsArticles,
		opts...,
	)
	newsroomAPIGetNewsArticleByIdHandler := connect_go.NewUnaryHandler(
		NewsroomAPIGetNewsArticleByIdProcedure,
		svc.GetNewsArticleById,
		opts...,
	)
	newsroomAPIUpdateNewsArticleHandler := connect_go.NewUnaryHandler(
		NewsroomAPIUpdateNewsArticleProcedure,
		svc.UpdateNewsArticle,
		opts...,
	)
	newsroomAPICreatePublishedArticleHandler := connect_go.NewUnaryHandler(
		NewsroomAPICreatePublishedArticleProcedure,
		svc.CreatePublishedArticle,
		opts...,
	)
	newsroomAPIListPublishedArticlesHandler := connect_go.NewUnaryHandler(
		NewsroomAPIListPublishedArticlesProcedure,
		svc.ListPublishedArticles,
		opts...,
	)
	newsroomAPIGetPublishedArticleByIdHandler := connect_go.NewUnaryHandler(
		NewsroomAPIGetPublishedArticleByIdProcedure,
		svc.GetPublishedArticleById,
		opts...,
	)
	newsroomAPIUserActivityHandler := connect_go.NewUnaryHandler(
		NewsroomAPIUserActivityProcedure,
		svc.UserActivity,
		opts...,
	)
	newsroomAPIGetNewsForUserHandler := connect_go.NewUnaryHandler(
		NewsroomAPIGetNewsForUserProcedure,
		svc.GetNewsForUser,
		opts...,
	)
	return "/api.v1alpha1.newsroom.NewsroomAPI/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case NewsroomAPICreateNewsArticleProcedure:
			newsroomAPICreateNewsArticleHandler.ServeHTTP(w, r)
		case NewsroomAPIListNewsArticlesProcedure:
			newsroomAPIListNewsArticlesHandler.ServeHTTP(w, r)
		case NewsroomAPIGetNewsArticleByIdProcedure:
			newsroomAPIGetNewsArticleByIdHandler.ServeHTTP(w, r)
		case NewsroomAPIUpdateNewsArticleProcedure:
			newsroomAPIUpdateNewsArticleHandler.ServeHTTP(w, r)
		case NewsroomAPICreatePublishedArticleProcedure:
			newsroomAPICreatePublishedArticleHandler.ServeHTTP(w, r)
		case NewsroomAPIListPublishedArticlesProcedure:
			newsroomAPIListPublishedArticlesHandler.ServeHTTP(w, r)
		case NewsroomAPIGetPublishedArticleByIdProcedure:
			newsroomAPIGetPublishedArticleByIdHandler.ServeHTTP(w, r)
		case NewsroomAPIUserActivityProcedure:
			newsroomAPIUserActivityHandler.ServeHTTP(w, r)
		case NewsroomAPIGetNewsForUserProcedure:
			newsroomAPIGetNewsForUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedNewsroomAPIHandler returns CodeUnimplemented from all methods.
type UnimplementedNewsroomAPIHandler struct{}

func (UnimplementedNewsroomAPIHandler) CreateNewsArticle(context.Context, *connect_go.Request[newsroom.CreateNewsArticleRequest]) (*connect_go.Response[newsroom.CreateNewsArticleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.newsroom.NewsroomAPI.CreateNewsArticle is not implemented"))
}

func (UnimplementedNewsroomAPIHandler) ListNewsArticles(context.Context, *connect_go.Request[newsroom.ListNewsArticlesRequest]) (*connect_go.Response[newsroom.ListNewsArticlesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.newsroom.NewsroomAPI.ListNewsArticles is not implemented"))
}

func (UnimplementedNewsroomAPIHandler) GetNewsArticleById(context.Context, *connect_go.Request[newsroom.GetNewsArticleByIdRequest]) (*connect_go.Response[newsroom.GetNewsArticleByIdResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.newsroom.NewsroomAPI.GetNewsArticleById is not implemented"))
}

func (UnimplementedNewsroomAPIHandler) UpdateNewsArticle(context.Context, *connect_go.Request[newsroom.UpdateNewsArticleRequest]) (*connect_go.Response[newsroom.UpdateNewsArticleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.newsroom.NewsroomAPI.UpdateNewsArticle is not implemented"))
}

func (UnimplementedNewsroomAPIHandler) CreatePublishedArticle(context.Context, *connect_go.Request[newsroom.CreatePublishedArticleRequest]) (*connect_go.Response[newsroom.CreatePublishedArticleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.newsroom.NewsroomAPI.CreatePublishedArticle is not implemented"))
}

func (UnimplementedNewsroomAPIHandler) ListPublishedArticles(context.Context, *connect_go.Request[newsroom.ListPublishedArticlesRequest]) (*connect_go.Response[newsroom.ListPublishedArticlesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.newsroom.NewsroomAPI.ListPublishedArticles is not implemented"))
}

func (UnimplementedNewsroomAPIHandler) GetPublishedArticleById(context.Context, *connect_go.Request[newsroom.GetPublishedArticleByIdRequest]) (*connect_go.Response[newsroom.GetPublishedArticleByIdResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.newsroom.NewsroomAPI.GetPublishedArticleById is not implemented"))
}

func (UnimplementedNewsroomAPIHandler) UserActivity(context.Context, *connect_go.Request[newsroom.UserActivityRequest]) (*connect_go.Response[newsroom.UserActivityResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.newsroom.NewsroomAPI.UserActivity is not implemented"))
}

func (UnimplementedNewsroomAPIHandler) GetNewsForUser(context.Context, *connect_go.Request[newsroom.GetNewsForUserRequest]) (*connect_go.Response[newsroom.GetNewsForUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.newsroom.NewsroomAPI.GetNewsForUser is not implemented"))
}
