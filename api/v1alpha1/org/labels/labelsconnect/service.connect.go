// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/org/labels/service.proto

package labelsconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	labels "github.com/tcncloud/api-go/api/v1alpha1/org/labels"
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
	// LabelsServiceName is the fully-qualified name of the LabelsService service.
	LabelsServiceName = "api.v1alpha1.org.labels.LabelsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LabelsServiceCreateLabelProcedure is the fully-qualified name of the LabelsService's CreateLabel
	// RPC.
	LabelsServiceCreateLabelProcedure = "/api.v1alpha1.org.labels.LabelsService/CreateLabel"
	// LabelsServiceGetLabelProcedure is the fully-qualified name of the LabelsService's GetLabel RPC.
	LabelsServiceGetLabelProcedure = "/api.v1alpha1.org.labels.LabelsService/GetLabel"
	// LabelsServiceUpdateLabelProcedure is the fully-qualified name of the LabelsService's UpdateLabel
	// RPC.
	LabelsServiceUpdateLabelProcedure = "/api.v1alpha1.org.labels.LabelsService/UpdateLabel"
	// LabelsServiceListLabelsProcedure is the fully-qualified name of the LabelsService's ListLabels
	// RPC.
	LabelsServiceListLabelsProcedure = "/api.v1alpha1.org.labels.LabelsService/ListLabels"
	// LabelsServiceDeleteLabelProcedure is the fully-qualified name of the LabelsService's DeleteLabel
	// RPC.
	LabelsServiceDeleteLabelProcedure = "/api.v1alpha1.org.labels.LabelsService/DeleteLabel"
	// LabelsServiceAttachLabelProcedure is the fully-qualified name of the LabelsService's AttachLabel
	// RPC.
	LabelsServiceAttachLabelProcedure = "/api.v1alpha1.org.labels.LabelsService/AttachLabel"
	// LabelsServiceDetachLabelProcedure is the fully-qualified name of the LabelsService's DetachLabel
	// RPC.
	LabelsServiceDetachLabelProcedure = "/api.v1alpha1.org.labels.LabelsService/DetachLabel"
	// LabelsServiceGetLabeledEntityMapProcedure is the fully-qualified name of the LabelsService's
	// GetLabeledEntityMap RPC.
	LabelsServiceGetLabeledEntityMapProcedure = "/api.v1alpha1.org.labels.LabelsService/GetLabeledEntityMap"
	// LabelsServiceAssignLabelsProcedure is the fully-qualified name of the LabelsService's
	// AssignLabels RPC.
	LabelsServiceAssignLabelsProcedure = "/api.v1alpha1.org.labels.LabelsService/AssignLabels"
	// LabelsServiceRevokeLabelsProcedure is the fully-qualified name of the LabelsService's
	// RevokeLabels RPC.
	LabelsServiceRevokeLabelsProcedure = "/api.v1alpha1.org.labels.LabelsService/RevokeLabels"
)

// LabelsServiceClient is a client for the api.v1alpha1.org.labels.LabelsService service.
type LabelsServiceClient interface {
	// CreateLabel creates a new label.
	CreateLabel(context.Context, *connect_go.Request[labels.CreateLabelRequest]) (*connect_go.Response[labels.CreateLabelResponse], error)
	// GetLabel gets a single label
	GetLabel(context.Context, *connect_go.Request[labels.GetLabelRequest]) (*connect_go.Response[labels.GetLabelResponse], error)
	// UpdateLabel updates a single label
	UpdateLabel(context.Context, *connect_go.Request[labels.UpdateLabelRequest]) (*connect_go.Response[labels.UpdateLabelResponse], error)
	// ListLabels lists all labels for a given organization
	ListLabels(context.Context, *connect_go.Request[labels.ListLabelsRequest]) (*connect_go.Response[labels.ListLabelsResponse], error)
	// DeleteLabel deletes a single label
	DeleteLabel(context.Context, *connect_go.Request[labels.DeleteLabelRequest]) (*connect_go.Response[labels.DeleteLabelResponse], error)
	// AttachLabel attaches a label to a given entity type
	AttachLabel(context.Context, *connect_go.Request[labels.AttachLabelRequest]) (*connect_go.Response[labels.AttachLabelResponse], error)
	// DetachLabel detaches a label from an entity based on an entity type
	DetachLabel(context.Context, *connect_go.Request[labels.DetachLabelRequest]) (*connect_go.Response[labels.DetachLabelResponse], error)
	// GetLabeledEntityMap gives back a map of entity Id to attached labels. The Entity type is specified on the request
	GetLabeledEntityMap(context.Context, *connect_go.Request[labels.GetLabeledEntityMapRequest]) (*connect_go.Response[labels.GetLabeledEntityMapResponse], error)
	// AssignLabels assigns labels to a specific permission group.
	AssignLabels(context.Context, *connect_go.Request[labels.AssignLabelsRequest]) (*connect_go.Response[labels.AssignLabelsResponse], error)
	// RevokeLabels revokes labels from a specific permission group.
	RevokeLabels(context.Context, *connect_go.Request[labels.RevokeLabelsRequest]) (*connect_go.Response[labels.RevokeLabelsResponse], error)
}

// NewLabelsServiceClient constructs a client for the api.v1alpha1.org.labels.LabelsService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLabelsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) LabelsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &labelsServiceClient{
		createLabel: connect_go.NewClient[labels.CreateLabelRequest, labels.CreateLabelResponse](
			httpClient,
			baseURL+LabelsServiceCreateLabelProcedure,
			opts...,
		),
		getLabel: connect_go.NewClient[labels.GetLabelRequest, labels.GetLabelResponse](
			httpClient,
			baseURL+LabelsServiceGetLabelProcedure,
			opts...,
		),
		updateLabel: connect_go.NewClient[labels.UpdateLabelRequest, labels.UpdateLabelResponse](
			httpClient,
			baseURL+LabelsServiceUpdateLabelProcedure,
			opts...,
		),
		listLabels: connect_go.NewClient[labels.ListLabelsRequest, labels.ListLabelsResponse](
			httpClient,
			baseURL+LabelsServiceListLabelsProcedure,
			opts...,
		),
		deleteLabel: connect_go.NewClient[labels.DeleteLabelRequest, labels.DeleteLabelResponse](
			httpClient,
			baseURL+LabelsServiceDeleteLabelProcedure,
			opts...,
		),
		attachLabel: connect_go.NewClient[labels.AttachLabelRequest, labels.AttachLabelResponse](
			httpClient,
			baseURL+LabelsServiceAttachLabelProcedure,
			opts...,
		),
		detachLabel: connect_go.NewClient[labels.DetachLabelRequest, labels.DetachLabelResponse](
			httpClient,
			baseURL+LabelsServiceDetachLabelProcedure,
			opts...,
		),
		getLabeledEntityMap: connect_go.NewClient[labels.GetLabeledEntityMapRequest, labels.GetLabeledEntityMapResponse](
			httpClient,
			baseURL+LabelsServiceGetLabeledEntityMapProcedure,
			opts...,
		),
		assignLabels: connect_go.NewClient[labels.AssignLabelsRequest, labels.AssignLabelsResponse](
			httpClient,
			baseURL+LabelsServiceAssignLabelsProcedure,
			opts...,
		),
		revokeLabels: connect_go.NewClient[labels.RevokeLabelsRequest, labels.RevokeLabelsResponse](
			httpClient,
			baseURL+LabelsServiceRevokeLabelsProcedure,
			opts...,
		),
	}
}

// labelsServiceClient implements LabelsServiceClient.
type labelsServiceClient struct {
	createLabel         *connect_go.Client[labels.CreateLabelRequest, labels.CreateLabelResponse]
	getLabel            *connect_go.Client[labels.GetLabelRequest, labels.GetLabelResponse]
	updateLabel         *connect_go.Client[labels.UpdateLabelRequest, labels.UpdateLabelResponse]
	listLabels          *connect_go.Client[labels.ListLabelsRequest, labels.ListLabelsResponse]
	deleteLabel         *connect_go.Client[labels.DeleteLabelRequest, labels.DeleteLabelResponse]
	attachLabel         *connect_go.Client[labels.AttachLabelRequest, labels.AttachLabelResponse]
	detachLabel         *connect_go.Client[labels.DetachLabelRequest, labels.DetachLabelResponse]
	getLabeledEntityMap *connect_go.Client[labels.GetLabeledEntityMapRequest, labels.GetLabeledEntityMapResponse]
	assignLabels        *connect_go.Client[labels.AssignLabelsRequest, labels.AssignLabelsResponse]
	revokeLabels        *connect_go.Client[labels.RevokeLabelsRequest, labels.RevokeLabelsResponse]
}

// CreateLabel calls api.v1alpha1.org.labels.LabelsService.CreateLabel.
func (c *labelsServiceClient) CreateLabel(ctx context.Context, req *connect_go.Request[labels.CreateLabelRequest]) (*connect_go.Response[labels.CreateLabelResponse], error) {
	return c.createLabel.CallUnary(ctx, req)
}

// GetLabel calls api.v1alpha1.org.labels.LabelsService.GetLabel.
func (c *labelsServiceClient) GetLabel(ctx context.Context, req *connect_go.Request[labels.GetLabelRequest]) (*connect_go.Response[labels.GetLabelResponse], error) {
	return c.getLabel.CallUnary(ctx, req)
}

// UpdateLabel calls api.v1alpha1.org.labels.LabelsService.UpdateLabel.
func (c *labelsServiceClient) UpdateLabel(ctx context.Context, req *connect_go.Request[labels.UpdateLabelRequest]) (*connect_go.Response[labels.UpdateLabelResponse], error) {
	return c.updateLabel.CallUnary(ctx, req)
}

// ListLabels calls api.v1alpha1.org.labels.LabelsService.ListLabels.
func (c *labelsServiceClient) ListLabels(ctx context.Context, req *connect_go.Request[labels.ListLabelsRequest]) (*connect_go.Response[labels.ListLabelsResponse], error) {
	return c.listLabels.CallUnary(ctx, req)
}

// DeleteLabel calls api.v1alpha1.org.labels.LabelsService.DeleteLabel.
func (c *labelsServiceClient) DeleteLabel(ctx context.Context, req *connect_go.Request[labels.DeleteLabelRequest]) (*connect_go.Response[labels.DeleteLabelResponse], error) {
	return c.deleteLabel.CallUnary(ctx, req)
}

// AttachLabel calls api.v1alpha1.org.labels.LabelsService.AttachLabel.
func (c *labelsServiceClient) AttachLabel(ctx context.Context, req *connect_go.Request[labels.AttachLabelRequest]) (*connect_go.Response[labels.AttachLabelResponse], error) {
	return c.attachLabel.CallUnary(ctx, req)
}

// DetachLabel calls api.v1alpha1.org.labels.LabelsService.DetachLabel.
func (c *labelsServiceClient) DetachLabel(ctx context.Context, req *connect_go.Request[labels.DetachLabelRequest]) (*connect_go.Response[labels.DetachLabelResponse], error) {
	return c.detachLabel.CallUnary(ctx, req)
}

// GetLabeledEntityMap calls api.v1alpha1.org.labels.LabelsService.GetLabeledEntityMap.
func (c *labelsServiceClient) GetLabeledEntityMap(ctx context.Context, req *connect_go.Request[labels.GetLabeledEntityMapRequest]) (*connect_go.Response[labels.GetLabeledEntityMapResponse], error) {
	return c.getLabeledEntityMap.CallUnary(ctx, req)
}

// AssignLabels calls api.v1alpha1.org.labels.LabelsService.AssignLabels.
func (c *labelsServiceClient) AssignLabels(ctx context.Context, req *connect_go.Request[labels.AssignLabelsRequest]) (*connect_go.Response[labels.AssignLabelsResponse], error) {
	return c.assignLabels.CallUnary(ctx, req)
}

// RevokeLabels calls api.v1alpha1.org.labels.LabelsService.RevokeLabels.
func (c *labelsServiceClient) RevokeLabels(ctx context.Context, req *connect_go.Request[labels.RevokeLabelsRequest]) (*connect_go.Response[labels.RevokeLabelsResponse], error) {
	return c.revokeLabels.CallUnary(ctx, req)
}

// LabelsServiceHandler is an implementation of the api.v1alpha1.org.labels.LabelsService service.
type LabelsServiceHandler interface {
	// CreateLabel creates a new label.
	CreateLabel(context.Context, *connect_go.Request[labels.CreateLabelRequest]) (*connect_go.Response[labels.CreateLabelResponse], error)
	// GetLabel gets a single label
	GetLabel(context.Context, *connect_go.Request[labels.GetLabelRequest]) (*connect_go.Response[labels.GetLabelResponse], error)
	// UpdateLabel updates a single label
	UpdateLabel(context.Context, *connect_go.Request[labels.UpdateLabelRequest]) (*connect_go.Response[labels.UpdateLabelResponse], error)
	// ListLabels lists all labels for a given organization
	ListLabels(context.Context, *connect_go.Request[labels.ListLabelsRequest]) (*connect_go.Response[labels.ListLabelsResponse], error)
	// DeleteLabel deletes a single label
	DeleteLabel(context.Context, *connect_go.Request[labels.DeleteLabelRequest]) (*connect_go.Response[labels.DeleteLabelResponse], error)
	// AttachLabel attaches a label to a given entity type
	AttachLabel(context.Context, *connect_go.Request[labels.AttachLabelRequest]) (*connect_go.Response[labels.AttachLabelResponse], error)
	// DetachLabel detaches a label from an entity based on an entity type
	DetachLabel(context.Context, *connect_go.Request[labels.DetachLabelRequest]) (*connect_go.Response[labels.DetachLabelResponse], error)
	// GetLabeledEntityMap gives back a map of entity Id to attached labels. The Entity type is specified on the request
	GetLabeledEntityMap(context.Context, *connect_go.Request[labels.GetLabeledEntityMapRequest]) (*connect_go.Response[labels.GetLabeledEntityMapResponse], error)
	// AssignLabels assigns labels to a specific permission group.
	AssignLabels(context.Context, *connect_go.Request[labels.AssignLabelsRequest]) (*connect_go.Response[labels.AssignLabelsResponse], error)
	// RevokeLabels revokes labels from a specific permission group.
	RevokeLabels(context.Context, *connect_go.Request[labels.RevokeLabelsRequest]) (*connect_go.Response[labels.RevokeLabelsResponse], error)
}

// NewLabelsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLabelsServiceHandler(svc LabelsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	labelsServiceCreateLabelHandler := connect_go.NewUnaryHandler(
		LabelsServiceCreateLabelProcedure,
		svc.CreateLabel,
		opts...,
	)
	labelsServiceGetLabelHandler := connect_go.NewUnaryHandler(
		LabelsServiceGetLabelProcedure,
		svc.GetLabel,
		opts...,
	)
	labelsServiceUpdateLabelHandler := connect_go.NewUnaryHandler(
		LabelsServiceUpdateLabelProcedure,
		svc.UpdateLabel,
		opts...,
	)
	labelsServiceListLabelsHandler := connect_go.NewUnaryHandler(
		LabelsServiceListLabelsProcedure,
		svc.ListLabels,
		opts...,
	)
	labelsServiceDeleteLabelHandler := connect_go.NewUnaryHandler(
		LabelsServiceDeleteLabelProcedure,
		svc.DeleteLabel,
		opts...,
	)
	labelsServiceAttachLabelHandler := connect_go.NewUnaryHandler(
		LabelsServiceAttachLabelProcedure,
		svc.AttachLabel,
		opts...,
	)
	labelsServiceDetachLabelHandler := connect_go.NewUnaryHandler(
		LabelsServiceDetachLabelProcedure,
		svc.DetachLabel,
		opts...,
	)
	labelsServiceGetLabeledEntityMapHandler := connect_go.NewUnaryHandler(
		LabelsServiceGetLabeledEntityMapProcedure,
		svc.GetLabeledEntityMap,
		opts...,
	)
	labelsServiceAssignLabelsHandler := connect_go.NewUnaryHandler(
		LabelsServiceAssignLabelsProcedure,
		svc.AssignLabels,
		opts...,
	)
	labelsServiceRevokeLabelsHandler := connect_go.NewUnaryHandler(
		LabelsServiceRevokeLabelsProcedure,
		svc.RevokeLabels,
		opts...,
	)
	return "/api.v1alpha1.org.labels.LabelsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LabelsServiceCreateLabelProcedure:
			labelsServiceCreateLabelHandler.ServeHTTP(w, r)
		case LabelsServiceGetLabelProcedure:
			labelsServiceGetLabelHandler.ServeHTTP(w, r)
		case LabelsServiceUpdateLabelProcedure:
			labelsServiceUpdateLabelHandler.ServeHTTP(w, r)
		case LabelsServiceListLabelsProcedure:
			labelsServiceListLabelsHandler.ServeHTTP(w, r)
		case LabelsServiceDeleteLabelProcedure:
			labelsServiceDeleteLabelHandler.ServeHTTP(w, r)
		case LabelsServiceAttachLabelProcedure:
			labelsServiceAttachLabelHandler.ServeHTTP(w, r)
		case LabelsServiceDetachLabelProcedure:
			labelsServiceDetachLabelHandler.ServeHTTP(w, r)
		case LabelsServiceGetLabeledEntityMapProcedure:
			labelsServiceGetLabeledEntityMapHandler.ServeHTTP(w, r)
		case LabelsServiceAssignLabelsProcedure:
			labelsServiceAssignLabelsHandler.ServeHTTP(w, r)
		case LabelsServiceRevokeLabelsProcedure:
			labelsServiceRevokeLabelsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLabelsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLabelsServiceHandler struct{}

func (UnimplementedLabelsServiceHandler) CreateLabel(context.Context, *connect_go.Request[labels.CreateLabelRequest]) (*connect_go.Response[labels.CreateLabelResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.labels.LabelsService.CreateLabel is not implemented"))
}

func (UnimplementedLabelsServiceHandler) GetLabel(context.Context, *connect_go.Request[labels.GetLabelRequest]) (*connect_go.Response[labels.GetLabelResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.labels.LabelsService.GetLabel is not implemented"))
}

func (UnimplementedLabelsServiceHandler) UpdateLabel(context.Context, *connect_go.Request[labels.UpdateLabelRequest]) (*connect_go.Response[labels.UpdateLabelResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.labels.LabelsService.UpdateLabel is not implemented"))
}

func (UnimplementedLabelsServiceHandler) ListLabels(context.Context, *connect_go.Request[labels.ListLabelsRequest]) (*connect_go.Response[labels.ListLabelsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.labels.LabelsService.ListLabels is not implemented"))
}

func (UnimplementedLabelsServiceHandler) DeleteLabel(context.Context, *connect_go.Request[labels.DeleteLabelRequest]) (*connect_go.Response[labels.DeleteLabelResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.labels.LabelsService.DeleteLabel is not implemented"))
}

func (UnimplementedLabelsServiceHandler) AttachLabel(context.Context, *connect_go.Request[labels.AttachLabelRequest]) (*connect_go.Response[labels.AttachLabelResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.labels.LabelsService.AttachLabel is not implemented"))
}

func (UnimplementedLabelsServiceHandler) DetachLabel(context.Context, *connect_go.Request[labels.DetachLabelRequest]) (*connect_go.Response[labels.DetachLabelResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.labels.LabelsService.DetachLabel is not implemented"))
}

func (UnimplementedLabelsServiceHandler) GetLabeledEntityMap(context.Context, *connect_go.Request[labels.GetLabeledEntityMapRequest]) (*connect_go.Response[labels.GetLabeledEntityMapResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.labels.LabelsService.GetLabeledEntityMap is not implemented"))
}

func (UnimplementedLabelsServiceHandler) AssignLabels(context.Context, *connect_go.Request[labels.AssignLabelsRequest]) (*connect_go.Response[labels.AssignLabelsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.labels.LabelsService.AssignLabels is not implemented"))
}

func (UnimplementedLabelsServiceHandler) RevokeLabels(context.Context, *connect_go.Request[labels.RevokeLabelsRequest]) (*connect_go.Response[labels.RevokeLabelsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.labels.LabelsService.RevokeLabels is not implemented"))
}
