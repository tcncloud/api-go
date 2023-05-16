// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v0alpha/cbs.proto

package v0alphaconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v0alpha "github.com/tcncloud/api-go/api/v0alpha"
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
	// CBSName is the fully-qualified name of the CBS service.
	CBSName = "api.v0alpha.CBS"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CBSCreateServiceIdProcedure is the fully-qualified name of the CBS's CreateServiceId RPC.
	CBSCreateServiceIdProcedure = "/api.v0alpha.CBS/CreateServiceId"
	// CBSCreateCallbackWithDetailsProcedure is the fully-qualified name of the CBS's
	// CreateCallbackWithDetails RPC.
	CBSCreateCallbackWithDetailsProcedure = "/api.v0alpha.CBS/CreateCallbackWithDetails"
	// CBSUpdateScheduledCallbackToReadyProcedure is the fully-qualified name of the CBS's
	// UpdateScheduledCallbackToReady RPC.
	CBSUpdateScheduledCallbackToReadyProcedure = "/api.v0alpha.CBS/UpdateScheduledCallbackToReady"
	// CBSUpdateScheduledCallbackToCanceledProcedure is the fully-qualified name of the CBS's
	// UpdateScheduledCallbackToCanceled RPC.
	CBSUpdateScheduledCallbackToCanceledProcedure = "/api.v0alpha.CBS/UpdateScheduledCallbackToCanceled"
	// CBSGetNextScheduledCallbackWithDetailsProcedure is the fully-qualified name of the CBS's
	// GetNextScheduledCallbackWithDetails RPC.
	CBSGetNextScheduledCallbackWithDetailsProcedure = "/api.v0alpha.CBS/GetNextScheduledCallbackWithDetails"
	// CBSUpdateScheduledCallbackToClosedProcedure is the fully-qualified name of the CBS's
	// UpdateScheduledCallbackToClosed RPC.
	CBSUpdateScheduledCallbackToClosedProcedure = "/api.v0alpha.CBS/UpdateScheduledCallbackToClosed"
	// CBSUpdateScheduledCallbackProcedure is the fully-qualified name of the CBS's
	// UpdateScheduledCallback RPC.
	CBSUpdateScheduledCallbackProcedure = "/api.v0alpha.CBS/UpdateScheduledCallback"
	// CBSGetScheduledCallbackWithDetailsProcedure is the fully-qualified name of the CBS's
	// GetScheduledCallbackWithDetails RPC.
	CBSGetScheduledCallbackWithDetailsProcedure = "/api.v0alpha.CBS/GetScheduledCallbackWithDetails"
	// CBSListScheduledCallbacksWithDetailsProcedure is the fully-qualified name of the CBS's
	// ListScheduledCallbacksWithDetails RPC.
	CBSListScheduledCallbacksWithDetailsProcedure = "/api.v0alpha.CBS/ListScheduledCallbacksWithDetails"
	// CBSListScheduledCallbacksWithDetailsBySkillsProcedure is the fully-qualified name of the CBS's
	// ListScheduledCallbacksWithDetailsBySkills RPC.
	CBSListScheduledCallbacksWithDetailsBySkillsProcedure = "/api.v0alpha.CBS/ListScheduledCallbacksWithDetailsBySkills"
)

// CBSClient is a client for the api.v0alpha.CBS service.
type CBSClient interface {
	CreateServiceId(context.Context, *connect_go.Request[v0alpha.CreateServiceIdReq]) (*connect_go.Response[v0alpha.CreateServiceIdRes], error)
	CreateCallbackWithDetails(context.Context, *connect_go.Request[v0alpha.CreateCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.CreateCallbackWithDetailsRes], error)
	UpdateScheduledCallbackToReady(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackToReadyReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToReadyRes], error)
	UpdateScheduledCallbackToCanceled(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackToCanceledReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToCanceledRes], error)
	// Skills that are a matching subset of the ones provided in the request.
	// If a callback is found the status of the callback is set to OPEN.
	// Required permissions:
	//
	//	NONE
	//
	// Errors:
	//   - grpc.Invalid: the service_id provided in the request is invalid.
	//   - grpc.NotFound: no matching service_id is found.
	//     callback is not found after updating the status of it to OPEN (shouldn't happen).
	GetNextScheduledCallbackWithDetails(context.Context, *connect_go.Request[v0alpha.GetNextScheduledCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.GetNextScheduledCallbackWithDetailsRes], error)
	UpdateScheduledCallbackToClosed(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackToClosedReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToClosedRes], error)
	// Updates a callback with the provided info, and replaces the details with the ones provided.
	UpdateScheduledCallback(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackRes], error)
	// Gets a scheduled callback's info and it's details for the given scheduled_callback_id.
	GetScheduledCallbackWithDetails(context.Context, *connect_go.Request[v0alpha.GetScheduledCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.GetScheduledCallbackWithDetailsRes], error)
	// Lists callbacks by phone number, caller id, or time range
	ListScheduledCallbacksWithDetails(context.Context, *connect_go.Request[v0alpha.ListScheduledCallbacksWithDetailsReq]) (*connect_go.ServerStreamForClient[v0alpha.ListScheduledCallbacksWithDetailsRes], error)
	// List callbacks by skills
	ListScheduledCallbacksWithDetailsBySkills(context.Context, *connect_go.Request[v0alpha.ListScheduledCallbacksWithDetailsBySkillsReq]) (*connect_go.Response[v0alpha.ListScheduledCallbacksWithDetailsRes], error)
}

// NewCBSClient constructs a client for the api.v0alpha.CBS service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCBSClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CBSClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &cBSClient{
		createServiceId: connect_go.NewClient[v0alpha.CreateServiceIdReq, v0alpha.CreateServiceIdRes](
			httpClient,
			baseURL+CBSCreateServiceIdProcedure,
			opts...,
		),
		createCallbackWithDetails: connect_go.NewClient[v0alpha.CreateCallbackWithDetailsReq, v0alpha.CreateCallbackWithDetailsRes](
			httpClient,
			baseURL+CBSCreateCallbackWithDetailsProcedure,
			opts...,
		),
		updateScheduledCallbackToReady: connect_go.NewClient[v0alpha.UpdateScheduledCallbackToReadyReq, v0alpha.UpdateScheduledCallbackToReadyRes](
			httpClient,
			baseURL+CBSUpdateScheduledCallbackToReadyProcedure,
			opts...,
		),
		updateScheduledCallbackToCanceled: connect_go.NewClient[v0alpha.UpdateScheduledCallbackToCanceledReq, v0alpha.UpdateScheduledCallbackToCanceledRes](
			httpClient,
			baseURL+CBSUpdateScheduledCallbackToCanceledProcedure,
			opts...,
		),
		getNextScheduledCallbackWithDetails: connect_go.NewClient[v0alpha.GetNextScheduledCallbackWithDetailsReq, v0alpha.GetNextScheduledCallbackWithDetailsRes](
			httpClient,
			baseURL+CBSGetNextScheduledCallbackWithDetailsProcedure,
			opts...,
		),
		updateScheduledCallbackToClosed: connect_go.NewClient[v0alpha.UpdateScheduledCallbackToClosedReq, v0alpha.UpdateScheduledCallbackToClosedRes](
			httpClient,
			baseURL+CBSUpdateScheduledCallbackToClosedProcedure,
			opts...,
		),
		updateScheduledCallback: connect_go.NewClient[v0alpha.UpdateScheduledCallbackReq, v0alpha.UpdateScheduledCallbackRes](
			httpClient,
			baseURL+CBSUpdateScheduledCallbackProcedure,
			opts...,
		),
		getScheduledCallbackWithDetails: connect_go.NewClient[v0alpha.GetScheduledCallbackWithDetailsReq, v0alpha.GetScheduledCallbackWithDetailsRes](
			httpClient,
			baseURL+CBSGetScheduledCallbackWithDetailsProcedure,
			opts...,
		),
		listScheduledCallbacksWithDetails: connect_go.NewClient[v0alpha.ListScheduledCallbacksWithDetailsReq, v0alpha.ListScheduledCallbacksWithDetailsRes](
			httpClient,
			baseURL+CBSListScheduledCallbacksWithDetailsProcedure,
			opts...,
		),
		listScheduledCallbacksWithDetailsBySkills: connect_go.NewClient[v0alpha.ListScheduledCallbacksWithDetailsBySkillsReq, v0alpha.ListScheduledCallbacksWithDetailsRes](
			httpClient,
			baseURL+CBSListScheduledCallbacksWithDetailsBySkillsProcedure,
			opts...,
		),
	}
}

// cBSClient implements CBSClient.
type cBSClient struct {
	createServiceId                           *connect_go.Client[v0alpha.CreateServiceIdReq, v0alpha.CreateServiceIdRes]
	createCallbackWithDetails                 *connect_go.Client[v0alpha.CreateCallbackWithDetailsReq, v0alpha.CreateCallbackWithDetailsRes]
	updateScheduledCallbackToReady            *connect_go.Client[v0alpha.UpdateScheduledCallbackToReadyReq, v0alpha.UpdateScheduledCallbackToReadyRes]
	updateScheduledCallbackToCanceled         *connect_go.Client[v0alpha.UpdateScheduledCallbackToCanceledReq, v0alpha.UpdateScheduledCallbackToCanceledRes]
	getNextScheduledCallbackWithDetails       *connect_go.Client[v0alpha.GetNextScheduledCallbackWithDetailsReq, v0alpha.GetNextScheduledCallbackWithDetailsRes]
	updateScheduledCallbackToClosed           *connect_go.Client[v0alpha.UpdateScheduledCallbackToClosedReq, v0alpha.UpdateScheduledCallbackToClosedRes]
	updateScheduledCallback                   *connect_go.Client[v0alpha.UpdateScheduledCallbackReq, v0alpha.UpdateScheduledCallbackRes]
	getScheduledCallbackWithDetails           *connect_go.Client[v0alpha.GetScheduledCallbackWithDetailsReq, v0alpha.GetScheduledCallbackWithDetailsRes]
	listScheduledCallbacksWithDetails         *connect_go.Client[v0alpha.ListScheduledCallbacksWithDetailsReq, v0alpha.ListScheduledCallbacksWithDetailsRes]
	listScheduledCallbacksWithDetailsBySkills *connect_go.Client[v0alpha.ListScheduledCallbacksWithDetailsBySkillsReq, v0alpha.ListScheduledCallbacksWithDetailsRes]
}

// CreateServiceId calls api.v0alpha.CBS.CreateServiceId.
func (c *cBSClient) CreateServiceId(ctx context.Context, req *connect_go.Request[v0alpha.CreateServiceIdReq]) (*connect_go.Response[v0alpha.CreateServiceIdRes], error) {
	return c.createServiceId.CallUnary(ctx, req)
}

// CreateCallbackWithDetails calls api.v0alpha.CBS.CreateCallbackWithDetails.
func (c *cBSClient) CreateCallbackWithDetails(ctx context.Context, req *connect_go.Request[v0alpha.CreateCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.CreateCallbackWithDetailsRes], error) {
	return c.createCallbackWithDetails.CallUnary(ctx, req)
}

// UpdateScheduledCallbackToReady calls api.v0alpha.CBS.UpdateScheduledCallbackToReady.
func (c *cBSClient) UpdateScheduledCallbackToReady(ctx context.Context, req *connect_go.Request[v0alpha.UpdateScheduledCallbackToReadyReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToReadyRes], error) {
	return c.updateScheduledCallbackToReady.CallUnary(ctx, req)
}

// UpdateScheduledCallbackToCanceled calls api.v0alpha.CBS.UpdateScheduledCallbackToCanceled.
func (c *cBSClient) UpdateScheduledCallbackToCanceled(ctx context.Context, req *connect_go.Request[v0alpha.UpdateScheduledCallbackToCanceledReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToCanceledRes], error) {
	return c.updateScheduledCallbackToCanceled.CallUnary(ctx, req)
}

// GetNextScheduledCallbackWithDetails calls api.v0alpha.CBS.GetNextScheduledCallbackWithDetails.
func (c *cBSClient) GetNextScheduledCallbackWithDetails(ctx context.Context, req *connect_go.Request[v0alpha.GetNextScheduledCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.GetNextScheduledCallbackWithDetailsRes], error) {
	return c.getNextScheduledCallbackWithDetails.CallUnary(ctx, req)
}

// UpdateScheduledCallbackToClosed calls api.v0alpha.CBS.UpdateScheduledCallbackToClosed.
func (c *cBSClient) UpdateScheduledCallbackToClosed(ctx context.Context, req *connect_go.Request[v0alpha.UpdateScheduledCallbackToClosedReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToClosedRes], error) {
	return c.updateScheduledCallbackToClosed.CallUnary(ctx, req)
}

// UpdateScheduledCallback calls api.v0alpha.CBS.UpdateScheduledCallback.
func (c *cBSClient) UpdateScheduledCallback(ctx context.Context, req *connect_go.Request[v0alpha.UpdateScheduledCallbackReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackRes], error) {
	return c.updateScheduledCallback.CallUnary(ctx, req)
}

// GetScheduledCallbackWithDetails calls api.v0alpha.CBS.GetScheduledCallbackWithDetails.
func (c *cBSClient) GetScheduledCallbackWithDetails(ctx context.Context, req *connect_go.Request[v0alpha.GetScheduledCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.GetScheduledCallbackWithDetailsRes], error) {
	return c.getScheduledCallbackWithDetails.CallUnary(ctx, req)
}

// ListScheduledCallbacksWithDetails calls api.v0alpha.CBS.ListScheduledCallbacksWithDetails.
func (c *cBSClient) ListScheduledCallbacksWithDetails(ctx context.Context, req *connect_go.Request[v0alpha.ListScheduledCallbacksWithDetailsReq]) (*connect_go.ServerStreamForClient[v0alpha.ListScheduledCallbacksWithDetailsRes], error) {
	return c.listScheduledCallbacksWithDetails.CallServerStream(ctx, req)
}

// ListScheduledCallbacksWithDetailsBySkills calls
// api.v0alpha.CBS.ListScheduledCallbacksWithDetailsBySkills.
func (c *cBSClient) ListScheduledCallbacksWithDetailsBySkills(ctx context.Context, req *connect_go.Request[v0alpha.ListScheduledCallbacksWithDetailsBySkillsReq]) (*connect_go.Response[v0alpha.ListScheduledCallbacksWithDetailsRes], error) {
	return c.listScheduledCallbacksWithDetailsBySkills.CallUnary(ctx, req)
}

// CBSHandler is an implementation of the api.v0alpha.CBS service.
type CBSHandler interface {
	CreateServiceId(context.Context, *connect_go.Request[v0alpha.CreateServiceIdReq]) (*connect_go.Response[v0alpha.CreateServiceIdRes], error)
	CreateCallbackWithDetails(context.Context, *connect_go.Request[v0alpha.CreateCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.CreateCallbackWithDetailsRes], error)
	UpdateScheduledCallbackToReady(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackToReadyReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToReadyRes], error)
	UpdateScheduledCallbackToCanceled(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackToCanceledReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToCanceledRes], error)
	// Skills that are a matching subset of the ones provided in the request.
	// If a callback is found the status of the callback is set to OPEN.
	// Required permissions:
	//
	//	NONE
	//
	// Errors:
	//   - grpc.Invalid: the service_id provided in the request is invalid.
	//   - grpc.NotFound: no matching service_id is found.
	//     callback is not found after updating the status of it to OPEN (shouldn't happen).
	GetNextScheduledCallbackWithDetails(context.Context, *connect_go.Request[v0alpha.GetNextScheduledCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.GetNextScheduledCallbackWithDetailsRes], error)
	UpdateScheduledCallbackToClosed(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackToClosedReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToClosedRes], error)
	// Updates a callback with the provided info, and replaces the details with the ones provided.
	UpdateScheduledCallback(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackRes], error)
	// Gets a scheduled callback's info and it's details for the given scheduled_callback_id.
	GetScheduledCallbackWithDetails(context.Context, *connect_go.Request[v0alpha.GetScheduledCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.GetScheduledCallbackWithDetailsRes], error)
	// Lists callbacks by phone number, caller id, or time range
	ListScheduledCallbacksWithDetails(context.Context, *connect_go.Request[v0alpha.ListScheduledCallbacksWithDetailsReq], *connect_go.ServerStream[v0alpha.ListScheduledCallbacksWithDetailsRes]) error
	// List callbacks by skills
	ListScheduledCallbacksWithDetailsBySkills(context.Context, *connect_go.Request[v0alpha.ListScheduledCallbacksWithDetailsBySkillsReq]) (*connect_go.Response[v0alpha.ListScheduledCallbacksWithDetailsRes], error)
}

// NewCBSHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCBSHandler(svc CBSHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(CBSCreateServiceIdProcedure, connect_go.NewUnaryHandler(
		CBSCreateServiceIdProcedure,
		svc.CreateServiceId,
		opts...,
	))
	mux.Handle(CBSCreateCallbackWithDetailsProcedure, connect_go.NewUnaryHandler(
		CBSCreateCallbackWithDetailsProcedure,
		svc.CreateCallbackWithDetails,
		opts...,
	))
	mux.Handle(CBSUpdateScheduledCallbackToReadyProcedure, connect_go.NewUnaryHandler(
		CBSUpdateScheduledCallbackToReadyProcedure,
		svc.UpdateScheduledCallbackToReady,
		opts...,
	))
	mux.Handle(CBSUpdateScheduledCallbackToCanceledProcedure, connect_go.NewUnaryHandler(
		CBSUpdateScheduledCallbackToCanceledProcedure,
		svc.UpdateScheduledCallbackToCanceled,
		opts...,
	))
	mux.Handle(CBSGetNextScheduledCallbackWithDetailsProcedure, connect_go.NewUnaryHandler(
		CBSGetNextScheduledCallbackWithDetailsProcedure,
		svc.GetNextScheduledCallbackWithDetails,
		opts...,
	))
	mux.Handle(CBSUpdateScheduledCallbackToClosedProcedure, connect_go.NewUnaryHandler(
		CBSUpdateScheduledCallbackToClosedProcedure,
		svc.UpdateScheduledCallbackToClosed,
		opts...,
	))
	mux.Handle(CBSUpdateScheduledCallbackProcedure, connect_go.NewUnaryHandler(
		CBSUpdateScheduledCallbackProcedure,
		svc.UpdateScheduledCallback,
		opts...,
	))
	mux.Handle(CBSGetScheduledCallbackWithDetailsProcedure, connect_go.NewUnaryHandler(
		CBSGetScheduledCallbackWithDetailsProcedure,
		svc.GetScheduledCallbackWithDetails,
		opts...,
	))
	mux.Handle(CBSListScheduledCallbacksWithDetailsProcedure, connect_go.NewServerStreamHandler(
		CBSListScheduledCallbacksWithDetailsProcedure,
		svc.ListScheduledCallbacksWithDetails,
		opts...,
	))
	mux.Handle(CBSListScheduledCallbacksWithDetailsBySkillsProcedure, connect_go.NewUnaryHandler(
		CBSListScheduledCallbacksWithDetailsBySkillsProcedure,
		svc.ListScheduledCallbacksWithDetailsBySkills,
		opts...,
	))
	return "/api.v0alpha.CBS/", mux
}

// UnimplementedCBSHandler returns CodeUnimplemented from all methods.
type UnimplementedCBSHandler struct{}

func (UnimplementedCBSHandler) CreateServiceId(context.Context, *connect_go.Request[v0alpha.CreateServiceIdReq]) (*connect_go.Response[v0alpha.CreateServiceIdRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.CBS.CreateServiceId is not implemented"))
}

func (UnimplementedCBSHandler) CreateCallbackWithDetails(context.Context, *connect_go.Request[v0alpha.CreateCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.CreateCallbackWithDetailsRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.CBS.CreateCallbackWithDetails is not implemented"))
}

func (UnimplementedCBSHandler) UpdateScheduledCallbackToReady(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackToReadyReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToReadyRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.CBS.UpdateScheduledCallbackToReady is not implemented"))
}

func (UnimplementedCBSHandler) UpdateScheduledCallbackToCanceled(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackToCanceledReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToCanceledRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.CBS.UpdateScheduledCallbackToCanceled is not implemented"))
}

func (UnimplementedCBSHandler) GetNextScheduledCallbackWithDetails(context.Context, *connect_go.Request[v0alpha.GetNextScheduledCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.GetNextScheduledCallbackWithDetailsRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.CBS.GetNextScheduledCallbackWithDetails is not implemented"))
}

func (UnimplementedCBSHandler) UpdateScheduledCallbackToClosed(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackToClosedReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackToClosedRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.CBS.UpdateScheduledCallbackToClosed is not implemented"))
}

func (UnimplementedCBSHandler) UpdateScheduledCallback(context.Context, *connect_go.Request[v0alpha.UpdateScheduledCallbackReq]) (*connect_go.Response[v0alpha.UpdateScheduledCallbackRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.CBS.UpdateScheduledCallback is not implemented"))
}

func (UnimplementedCBSHandler) GetScheduledCallbackWithDetails(context.Context, *connect_go.Request[v0alpha.GetScheduledCallbackWithDetailsReq]) (*connect_go.Response[v0alpha.GetScheduledCallbackWithDetailsRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.CBS.GetScheduledCallbackWithDetails is not implemented"))
}

func (UnimplementedCBSHandler) ListScheduledCallbacksWithDetails(context.Context, *connect_go.Request[v0alpha.ListScheduledCallbacksWithDetailsReq], *connect_go.ServerStream[v0alpha.ListScheduledCallbacksWithDetailsRes]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.CBS.ListScheduledCallbacksWithDetails is not implemented"))
}

func (UnimplementedCBSHandler) ListScheduledCallbacksWithDetailsBySkills(context.Context, *connect_go.Request[v0alpha.ListScheduledCallbacksWithDetailsBySkillsReq]) (*connect_go.Response[v0alpha.ListScheduledCallbacksWithDetailsRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.CBS.ListScheduledCallbacksWithDetailsBySkills is not implemented"))
}
