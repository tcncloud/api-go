// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v0alpha/dashboards.proto

package v0alphaconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v0alpha "github.com/tcncloud/api-go/api/v0alpha"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	// DashboardsName is the fully-qualified name of the Dashboards service.
	DashboardsName = "api.v0alpha.Dashboards"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DashboardsCreateDashboardProcedure is the fully-qualified name of the Dashboards's
	// CreateDashboard RPC.
	DashboardsCreateDashboardProcedure = "/api.v0alpha.Dashboards/CreateDashboard"
	// DashboardsGetDashboardProcedure is the fully-qualified name of the Dashboards's GetDashboard RPC.
	DashboardsGetDashboardProcedure = "/api.v0alpha.Dashboards/GetDashboard"
	// DashboardsGetDefaultDashboardProcedure is the fully-qualified name of the Dashboards's
	// GetDefaultDashboard RPC.
	DashboardsGetDefaultDashboardProcedure = "/api.v0alpha.Dashboards/GetDefaultDashboard"
	// DashboardsListDashboardsProcedure is the fully-qualified name of the Dashboards's ListDashboards
	// RPC.
	DashboardsListDashboardsProcedure = "/api.v0alpha.Dashboards/ListDashboards"
	// DashboardsListProductTypesProcedure is the fully-qualified name of the Dashboards's
	// ListProductTypes RPC.
	DashboardsListProductTypesProcedure = "/api.v0alpha.Dashboards/ListProductTypes"
	// DashboardsDeleteDashboardProcedure is the fully-qualified name of the Dashboards's
	// DeleteDashboard RPC.
	DashboardsDeleteDashboardProcedure = "/api.v0alpha.Dashboards/DeleteDashboard"
	// DashboardsSetDefaultDashboardProcedure is the fully-qualified name of the Dashboards's
	// SetDefaultDashboard RPC.
	DashboardsSetDefaultDashboardProcedure = "/api.v0alpha.Dashboards/SetDefaultDashboard"
	// DashboardsUpdateDashboardProcedure is the fully-qualified name of the Dashboards's
	// UpdateDashboard RPC.
	DashboardsUpdateDashboardProcedure = "/api.v0alpha.Dashboards/UpdateDashboard"
	// DashboardsUpdateDashboardTitleAndDescriptionProcedure is the fully-qualified name of the
	// Dashboards's UpdateDashboardTitleAndDescription RPC.
	DashboardsUpdateDashboardTitleAndDescriptionProcedure = "/api.v0alpha.Dashboards/UpdateDashboardTitleAndDescription"
	// DashboardsUpdateDashboardViewProcedure is the fully-qualified name of the Dashboards's
	// UpdateDashboardView RPC.
	DashboardsUpdateDashboardViewProcedure = "/api.v0alpha.Dashboards/UpdateDashboardView"
	// DashboardsUpdateDashboardLayoutProcedure is the fully-qualified name of the Dashboards's
	// UpdateDashboardLayout RPC.
	DashboardsUpdateDashboardLayoutProcedure = "/api.v0alpha.Dashboards/UpdateDashboardLayout"
)

// DashboardsClient is a client for the api.v0alpha.Dashboards service.
type DashboardsClient interface {
	// CreateDashboard creates a dashboard and associated panels
	CreateDashboard(context.Context, *connect_go.Request[v0alpha.CreateDashboardRequest]) (*connect_go.Response[v0alpha.CreateDashboardResponse], error)
	// GetDashboard retrieves a dashboard by the given ID and orgID and fetches its associated panels
	GetDashboard(context.Context, *connect_go.Request[v0alpha.GetDashboardRequest]) (*connect_go.Response[v0alpha.Dashboard], error)
	GetDefaultDashboard(context.Context, *connect_go.Request[v0alpha.GetDefaultDashboardRequest]) (*connect_go.Response[v0alpha.Dashboard], error)
	// ListDashboards retrieves all dashboard summaries for the given organization
	ListDashboards(context.Context, *connect_go.Request[v0alpha.ListDashboardsRequest]) (*connect_go.Response[v0alpha.ListDashboardsResponse], error)
	ListProductTypes(context.Context, *connect_go.Request[v0alpha.ListProductTypesRequest]) (*connect_go.Response[v0alpha.ListProductTypesResult], error)
	// Deletes a given dashboard
	DeleteDashboard(context.Context, *connect_go.Request[v0alpha.DeleteDashboardRequest]) (*connect_go.Response[emptypb.Empty], error)
	// Sets a dashboard as the user's default
	SetDefaultDashboard(context.Context, *connect_go.Request[v0alpha.SetDefaultDashboardRequest]) (*connect_go.Response[emptypb.Empty], error)
	// UpdateDashboard updates a dashboard
	UpdateDashboard(context.Context, *connect_go.Request[v0alpha.UpdateDashboardRequest]) (*connect_go.Response[emptypb.Empty], error)
	// UpdateDashboardTitleAndDescription updates a dashboards title and description
	UpdateDashboardTitleAndDescription(context.Context, *connect_go.Request[v0alpha.UpdateDashboardTitleAndDescriptionRequest]) (*connect_go.Response[emptypb.Empty], error)
	// UpdateDashboardView updates a dashboards view with the given view
	UpdateDashboardView(context.Context, *connect_go.Request[v0alpha.UpdateDashboardViewRequest]) (*connect_go.Response[emptypb.Empty], error)
	// UpdateDashboardLayout replaces a dashboards layout with a given layout
	UpdateDashboardLayout(context.Context, *connect_go.Request[v0alpha.UpdateDashboardLayoutRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewDashboardsClient constructs a client for the api.v0alpha.Dashboards service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDashboardsClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) DashboardsClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &dashboardsClient{
		createDashboard: connect_go.NewClient[v0alpha.CreateDashboardRequest, v0alpha.CreateDashboardResponse](
			httpClient,
			baseURL+DashboardsCreateDashboardProcedure,
			opts...,
		),
		getDashboard: connect_go.NewClient[v0alpha.GetDashboardRequest, v0alpha.Dashboard](
			httpClient,
			baseURL+DashboardsGetDashboardProcedure,
			opts...,
		),
		getDefaultDashboard: connect_go.NewClient[v0alpha.GetDefaultDashboardRequest, v0alpha.Dashboard](
			httpClient,
			baseURL+DashboardsGetDefaultDashboardProcedure,
			opts...,
		),
		listDashboards: connect_go.NewClient[v0alpha.ListDashboardsRequest, v0alpha.ListDashboardsResponse](
			httpClient,
			baseURL+DashboardsListDashboardsProcedure,
			opts...,
		),
		listProductTypes: connect_go.NewClient[v0alpha.ListProductTypesRequest, v0alpha.ListProductTypesResult](
			httpClient,
			baseURL+DashboardsListProductTypesProcedure,
			opts...,
		),
		deleteDashboard: connect_go.NewClient[v0alpha.DeleteDashboardRequest, emptypb.Empty](
			httpClient,
			baseURL+DashboardsDeleteDashboardProcedure,
			opts...,
		),
		setDefaultDashboard: connect_go.NewClient[v0alpha.SetDefaultDashboardRequest, emptypb.Empty](
			httpClient,
			baseURL+DashboardsSetDefaultDashboardProcedure,
			opts...,
		),
		updateDashboard: connect_go.NewClient[v0alpha.UpdateDashboardRequest, emptypb.Empty](
			httpClient,
			baseURL+DashboardsUpdateDashboardProcedure,
			opts...,
		),
		updateDashboardTitleAndDescription: connect_go.NewClient[v0alpha.UpdateDashboardTitleAndDescriptionRequest, emptypb.Empty](
			httpClient,
			baseURL+DashboardsUpdateDashboardTitleAndDescriptionProcedure,
			opts...,
		),
		updateDashboardView: connect_go.NewClient[v0alpha.UpdateDashboardViewRequest, emptypb.Empty](
			httpClient,
			baseURL+DashboardsUpdateDashboardViewProcedure,
			opts...,
		),
		updateDashboardLayout: connect_go.NewClient[v0alpha.UpdateDashboardLayoutRequest, emptypb.Empty](
			httpClient,
			baseURL+DashboardsUpdateDashboardLayoutProcedure,
			opts...,
		),
	}
}

// dashboardsClient implements DashboardsClient.
type dashboardsClient struct {
	createDashboard                    *connect_go.Client[v0alpha.CreateDashboardRequest, v0alpha.CreateDashboardResponse]
	getDashboard                       *connect_go.Client[v0alpha.GetDashboardRequest, v0alpha.Dashboard]
	getDefaultDashboard                *connect_go.Client[v0alpha.GetDefaultDashboardRequest, v0alpha.Dashboard]
	listDashboards                     *connect_go.Client[v0alpha.ListDashboardsRequest, v0alpha.ListDashboardsResponse]
	listProductTypes                   *connect_go.Client[v0alpha.ListProductTypesRequest, v0alpha.ListProductTypesResult]
	deleteDashboard                    *connect_go.Client[v0alpha.DeleteDashboardRequest, emptypb.Empty]
	setDefaultDashboard                *connect_go.Client[v0alpha.SetDefaultDashboardRequest, emptypb.Empty]
	updateDashboard                    *connect_go.Client[v0alpha.UpdateDashboardRequest, emptypb.Empty]
	updateDashboardTitleAndDescription *connect_go.Client[v0alpha.UpdateDashboardTitleAndDescriptionRequest, emptypb.Empty]
	updateDashboardView                *connect_go.Client[v0alpha.UpdateDashboardViewRequest, emptypb.Empty]
	updateDashboardLayout              *connect_go.Client[v0alpha.UpdateDashboardLayoutRequest, emptypb.Empty]
}

// CreateDashboard calls api.v0alpha.Dashboards.CreateDashboard.
func (c *dashboardsClient) CreateDashboard(ctx context.Context, req *connect_go.Request[v0alpha.CreateDashboardRequest]) (*connect_go.Response[v0alpha.CreateDashboardResponse], error) {
	return c.createDashboard.CallUnary(ctx, req)
}

// GetDashboard calls api.v0alpha.Dashboards.GetDashboard.
func (c *dashboardsClient) GetDashboard(ctx context.Context, req *connect_go.Request[v0alpha.GetDashboardRequest]) (*connect_go.Response[v0alpha.Dashboard], error) {
	return c.getDashboard.CallUnary(ctx, req)
}

// GetDefaultDashboard calls api.v0alpha.Dashboards.GetDefaultDashboard.
func (c *dashboardsClient) GetDefaultDashboard(ctx context.Context, req *connect_go.Request[v0alpha.GetDefaultDashboardRequest]) (*connect_go.Response[v0alpha.Dashboard], error) {
	return c.getDefaultDashboard.CallUnary(ctx, req)
}

// ListDashboards calls api.v0alpha.Dashboards.ListDashboards.
func (c *dashboardsClient) ListDashboards(ctx context.Context, req *connect_go.Request[v0alpha.ListDashboardsRequest]) (*connect_go.Response[v0alpha.ListDashboardsResponse], error) {
	return c.listDashboards.CallUnary(ctx, req)
}

// ListProductTypes calls api.v0alpha.Dashboards.ListProductTypes.
func (c *dashboardsClient) ListProductTypes(ctx context.Context, req *connect_go.Request[v0alpha.ListProductTypesRequest]) (*connect_go.Response[v0alpha.ListProductTypesResult], error) {
	return c.listProductTypes.CallUnary(ctx, req)
}

// DeleteDashboard calls api.v0alpha.Dashboards.DeleteDashboard.
func (c *dashboardsClient) DeleteDashboard(ctx context.Context, req *connect_go.Request[v0alpha.DeleteDashboardRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deleteDashboard.CallUnary(ctx, req)
}

// SetDefaultDashboard calls api.v0alpha.Dashboards.SetDefaultDashboard.
func (c *dashboardsClient) SetDefaultDashboard(ctx context.Context, req *connect_go.Request[v0alpha.SetDefaultDashboardRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.setDefaultDashboard.CallUnary(ctx, req)
}

// UpdateDashboard calls api.v0alpha.Dashboards.UpdateDashboard.
func (c *dashboardsClient) UpdateDashboard(ctx context.Context, req *connect_go.Request[v0alpha.UpdateDashboardRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.updateDashboard.CallUnary(ctx, req)
}

// UpdateDashboardTitleAndDescription calls
// api.v0alpha.Dashboards.UpdateDashboardTitleAndDescription.
func (c *dashboardsClient) UpdateDashboardTitleAndDescription(ctx context.Context, req *connect_go.Request[v0alpha.UpdateDashboardTitleAndDescriptionRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.updateDashboardTitleAndDescription.CallUnary(ctx, req)
}

// UpdateDashboardView calls api.v0alpha.Dashboards.UpdateDashboardView.
func (c *dashboardsClient) UpdateDashboardView(ctx context.Context, req *connect_go.Request[v0alpha.UpdateDashboardViewRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.updateDashboardView.CallUnary(ctx, req)
}

// UpdateDashboardLayout calls api.v0alpha.Dashboards.UpdateDashboardLayout.
func (c *dashboardsClient) UpdateDashboardLayout(ctx context.Context, req *connect_go.Request[v0alpha.UpdateDashboardLayoutRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.updateDashboardLayout.CallUnary(ctx, req)
}

// DashboardsHandler is an implementation of the api.v0alpha.Dashboards service.
type DashboardsHandler interface {
	// CreateDashboard creates a dashboard and associated panels
	CreateDashboard(context.Context, *connect_go.Request[v0alpha.CreateDashboardRequest]) (*connect_go.Response[v0alpha.CreateDashboardResponse], error)
	// GetDashboard retrieves a dashboard by the given ID and orgID and fetches its associated panels
	GetDashboard(context.Context, *connect_go.Request[v0alpha.GetDashboardRequest]) (*connect_go.Response[v0alpha.Dashboard], error)
	GetDefaultDashboard(context.Context, *connect_go.Request[v0alpha.GetDefaultDashboardRequest]) (*connect_go.Response[v0alpha.Dashboard], error)
	// ListDashboards retrieves all dashboard summaries for the given organization
	ListDashboards(context.Context, *connect_go.Request[v0alpha.ListDashboardsRequest]) (*connect_go.Response[v0alpha.ListDashboardsResponse], error)
	ListProductTypes(context.Context, *connect_go.Request[v0alpha.ListProductTypesRequest]) (*connect_go.Response[v0alpha.ListProductTypesResult], error)
	// Deletes a given dashboard
	DeleteDashboard(context.Context, *connect_go.Request[v0alpha.DeleteDashboardRequest]) (*connect_go.Response[emptypb.Empty], error)
	// Sets a dashboard as the user's default
	SetDefaultDashboard(context.Context, *connect_go.Request[v0alpha.SetDefaultDashboardRequest]) (*connect_go.Response[emptypb.Empty], error)
	// UpdateDashboard updates a dashboard
	UpdateDashboard(context.Context, *connect_go.Request[v0alpha.UpdateDashboardRequest]) (*connect_go.Response[emptypb.Empty], error)
	// UpdateDashboardTitleAndDescription updates a dashboards title and description
	UpdateDashboardTitleAndDescription(context.Context, *connect_go.Request[v0alpha.UpdateDashboardTitleAndDescriptionRequest]) (*connect_go.Response[emptypb.Empty], error)
	// UpdateDashboardView updates a dashboards view with the given view
	UpdateDashboardView(context.Context, *connect_go.Request[v0alpha.UpdateDashboardViewRequest]) (*connect_go.Response[emptypb.Empty], error)
	// UpdateDashboardLayout replaces a dashboards layout with a given layout
	UpdateDashboardLayout(context.Context, *connect_go.Request[v0alpha.UpdateDashboardLayoutRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewDashboardsHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDashboardsHandler(svc DashboardsHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(DashboardsCreateDashboardProcedure, connect_go.NewUnaryHandler(
		DashboardsCreateDashboardProcedure,
		svc.CreateDashboard,
		opts...,
	))
	mux.Handle(DashboardsGetDashboardProcedure, connect_go.NewUnaryHandler(
		DashboardsGetDashboardProcedure,
		svc.GetDashboard,
		opts...,
	))
	mux.Handle(DashboardsGetDefaultDashboardProcedure, connect_go.NewUnaryHandler(
		DashboardsGetDefaultDashboardProcedure,
		svc.GetDefaultDashboard,
		opts...,
	))
	mux.Handle(DashboardsListDashboardsProcedure, connect_go.NewUnaryHandler(
		DashboardsListDashboardsProcedure,
		svc.ListDashboards,
		opts...,
	))
	mux.Handle(DashboardsListProductTypesProcedure, connect_go.NewUnaryHandler(
		DashboardsListProductTypesProcedure,
		svc.ListProductTypes,
		opts...,
	))
	mux.Handle(DashboardsDeleteDashboardProcedure, connect_go.NewUnaryHandler(
		DashboardsDeleteDashboardProcedure,
		svc.DeleteDashboard,
		opts...,
	))
	mux.Handle(DashboardsSetDefaultDashboardProcedure, connect_go.NewUnaryHandler(
		DashboardsSetDefaultDashboardProcedure,
		svc.SetDefaultDashboard,
		opts...,
	))
	mux.Handle(DashboardsUpdateDashboardProcedure, connect_go.NewUnaryHandler(
		DashboardsUpdateDashboardProcedure,
		svc.UpdateDashboard,
		opts...,
	))
	mux.Handle(DashboardsUpdateDashboardTitleAndDescriptionProcedure, connect_go.NewUnaryHandler(
		DashboardsUpdateDashboardTitleAndDescriptionProcedure,
		svc.UpdateDashboardTitleAndDescription,
		opts...,
	))
	mux.Handle(DashboardsUpdateDashboardViewProcedure, connect_go.NewUnaryHandler(
		DashboardsUpdateDashboardViewProcedure,
		svc.UpdateDashboardView,
		opts...,
	))
	mux.Handle(DashboardsUpdateDashboardLayoutProcedure, connect_go.NewUnaryHandler(
		DashboardsUpdateDashboardLayoutProcedure,
		svc.UpdateDashboardLayout,
		opts...,
	))
	return "/api.v0alpha.Dashboards/", mux
}

// UnimplementedDashboardsHandler returns CodeUnimplemented from all methods.
type UnimplementedDashboardsHandler struct{}

func (UnimplementedDashboardsHandler) CreateDashboard(context.Context, *connect_go.Request[v0alpha.CreateDashboardRequest]) (*connect_go.Response[v0alpha.CreateDashboardResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Dashboards.CreateDashboard is not implemented"))
}

func (UnimplementedDashboardsHandler) GetDashboard(context.Context, *connect_go.Request[v0alpha.GetDashboardRequest]) (*connect_go.Response[v0alpha.Dashboard], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Dashboards.GetDashboard is not implemented"))
}

func (UnimplementedDashboardsHandler) GetDefaultDashboard(context.Context, *connect_go.Request[v0alpha.GetDefaultDashboardRequest]) (*connect_go.Response[v0alpha.Dashboard], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Dashboards.GetDefaultDashboard is not implemented"))
}

func (UnimplementedDashboardsHandler) ListDashboards(context.Context, *connect_go.Request[v0alpha.ListDashboardsRequest]) (*connect_go.Response[v0alpha.ListDashboardsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Dashboards.ListDashboards is not implemented"))
}

func (UnimplementedDashboardsHandler) ListProductTypes(context.Context, *connect_go.Request[v0alpha.ListProductTypesRequest]) (*connect_go.Response[v0alpha.ListProductTypesResult], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Dashboards.ListProductTypes is not implemented"))
}

func (UnimplementedDashboardsHandler) DeleteDashboard(context.Context, *connect_go.Request[v0alpha.DeleteDashboardRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Dashboards.DeleteDashboard is not implemented"))
}

func (UnimplementedDashboardsHandler) SetDefaultDashboard(context.Context, *connect_go.Request[v0alpha.SetDefaultDashboardRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Dashboards.SetDefaultDashboard is not implemented"))
}

func (UnimplementedDashboardsHandler) UpdateDashboard(context.Context, *connect_go.Request[v0alpha.UpdateDashboardRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Dashboards.UpdateDashboard is not implemented"))
}

func (UnimplementedDashboardsHandler) UpdateDashboardTitleAndDescription(context.Context, *connect_go.Request[v0alpha.UpdateDashboardTitleAndDescriptionRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Dashboards.UpdateDashboardTitleAndDescription is not implemented"))
}

func (UnimplementedDashboardsHandler) UpdateDashboardView(context.Context, *connect_go.Request[v0alpha.UpdateDashboardViewRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Dashboards.UpdateDashboardView is not implemented"))
}

func (UnimplementedDashboardsHandler) UpdateDashboardLayout(context.Context, *connect_go.Request[v0alpha.UpdateDashboardLayoutRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Dashboards.UpdateDashboardLayout is not implemented"))
}
