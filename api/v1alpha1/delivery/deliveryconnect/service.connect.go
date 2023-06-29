// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/delivery/service.proto

package deliveryconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	delivery "github.com/tcncloud/api-go/api/v1alpha1/delivery"
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
	// DeliveryApiName is the fully-qualified name of the DeliveryApi service.
	DeliveryApiName = "api.v1alpha1.delivery.DeliveryApi"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DeliveryApiCreateTransferConfigProcedure is the fully-qualified name of the DeliveryApi's
	// CreateTransferConfig RPC.
	DeliveryApiCreateTransferConfigProcedure = "/api.v1alpha1.delivery.DeliveryApi/CreateTransferConfig"
	// DeliveryApiListTransferConfigsProcedure is the fully-qualified name of the DeliveryApi's
	// ListTransferConfigs RPC.
	DeliveryApiListTransferConfigsProcedure = "/api.v1alpha1.delivery.DeliveryApi/ListTransferConfigs"
	// DeliveryApiListTransferConfigsByCredentialIDProcedure is the fully-qualified name of the
	// DeliveryApi's ListTransferConfigsByCredentialID RPC.
	DeliveryApiListTransferConfigsByCredentialIDProcedure = "/api.v1alpha1.delivery.DeliveryApi/ListTransferConfigsByCredentialID"
	// DeliveryApiUpdateTransferConfigProcedure is the fully-qualified name of the DeliveryApi's
	// UpdateTransferConfig RPC.
	DeliveryApiUpdateTransferConfigProcedure = "/api.v1alpha1.delivery.DeliveryApi/UpdateTransferConfig"
	// DeliveryApiDeleteTransferConfigProcedure is the fully-qualified name of the DeliveryApi's
	// DeleteTransferConfig RPC.
	DeliveryApiDeleteTransferConfigProcedure = "/api.v1alpha1.delivery.DeliveryApi/DeleteTransferConfig"
	// DeliveryApiGetTransferConfigProcedure is the fully-qualified name of the DeliveryApi's
	// GetTransferConfig RPC.
	DeliveryApiGetTransferConfigProcedure = "/api.v1alpha1.delivery.DeliveryApi/GetTransferConfig"
	// DeliveryApiGetTransferConfigByNameProcedure is the fully-qualified name of the DeliveryApi's
	// GetTransferConfigByName RPC.
	DeliveryApiGetTransferConfigByNameProcedure = "/api.v1alpha1.delivery.DeliveryApi/GetTransferConfigByName"
	// DeliveryApiListHistoryProcedure is the fully-qualified name of the DeliveryApi's ListHistory RPC.
	DeliveryApiListHistoryProcedure = "/api.v1alpha1.delivery.DeliveryApi/ListHistory"
	// DeliveryApiListHistoryByTransferConfigProcedure is the fully-qualified name of the DeliveryApi's
	// ListHistoryByTransferConfig RPC.
	DeliveryApiListHistoryByTransferConfigProcedure = "/api.v1alpha1.delivery.DeliveryApi/ListHistoryByTransferConfig"
	// DeliveryApiListCredentialsProcedure is the fully-qualified name of the DeliveryApi's
	// ListCredentials RPC.
	DeliveryApiListCredentialsProcedure = "/api.v1alpha1.delivery.DeliveryApi/ListCredentials"
	// DeliveryApiGetCredentialProcedure is the fully-qualified name of the DeliveryApi's GetCredential
	// RPC.
	DeliveryApiGetCredentialProcedure = "/api.v1alpha1.delivery.DeliveryApi/GetCredential"
	// DeliveryApiCreateCredentialProcedure is the fully-qualified name of the DeliveryApi's
	// CreateCredential RPC.
	DeliveryApiCreateCredentialProcedure = "/api.v1alpha1.delivery.DeliveryApi/CreateCredential"
	// DeliveryApiDeleteCredentialProcedure is the fully-qualified name of the DeliveryApi's
	// DeleteCredential RPC.
	DeliveryApiDeleteCredentialProcedure = "/api.v1alpha1.delivery.DeliveryApi/DeleteCredential"
	// DeliveryApiUpdateCredentialProcedure is the fully-qualified name of the DeliveryApi's
	// UpdateCredential RPC.
	DeliveryApiUpdateCredentialProcedure = "/api.v1alpha1.delivery.DeliveryApi/UpdateCredential"
)

// DeliveryApiClient is a client for the api.v1alpha1.delivery.DeliveryApi service.
type DeliveryApiClient interface {
	CreateTransferConfig(context.Context, *connect_go.Request[delivery.CreateTransferConfigReq]) (*connect_go.Response[delivery.CreateTransferConfigRes], error)
	ListTransferConfigs(context.Context, *connect_go.Request[delivery.ListTransferConfigsReq]) (*connect_go.Response[delivery.ListTransferConfigsRes], error)
	ListTransferConfigsByCredentialID(context.Context, *connect_go.Request[delivery.ListTransferConfigsByCredentialIDReq]) (*connect_go.Response[delivery.ListTransferConfigsByCredentialIDRes], error)
	UpdateTransferConfig(context.Context, *connect_go.Request[delivery.UpdateTransferConfigReq]) (*connect_go.Response[delivery.UpdateTransferConfigRes], error)
	DeleteTransferConfig(context.Context, *connect_go.Request[delivery.DeleteTransferConfigReq]) (*connect_go.Response[delivery.DeleteTransferConfigRes], error)
	GetTransferConfig(context.Context, *connect_go.Request[delivery.GetTransferConfigReq]) (*connect_go.Response[delivery.GetTransferConfigRes], error)
	GetTransferConfigByName(context.Context, *connect_go.Request[delivery.GetTransferConfigByNameReq]) (*connect_go.Response[delivery.GetTransferConfigByNameRes], error)
	ListHistory(context.Context, *connect_go.Request[delivery.ListHistoryReq]) (*connect_go.Response[delivery.ListHistoryRes], error)
	ListHistoryByTransferConfig(context.Context, *connect_go.Request[delivery.ListHistoryByTransferConfigReq]) (*connect_go.Response[delivery.ListHistoryByTransferConfigRes], error)
	ListCredentials(context.Context, *connect_go.Request[delivery.ListCredentialsReq]) (*connect_go.Response[delivery.ListCredentialsRes], error)
	GetCredential(context.Context, *connect_go.Request[delivery.GetCredentialReq]) (*connect_go.Response[delivery.GetCredentialRes], error)
	CreateCredential(context.Context, *connect_go.Request[delivery.CreateCredentialReq]) (*connect_go.Response[delivery.CreateCredentialRes], error)
	DeleteCredential(context.Context, *connect_go.Request[delivery.DeleteCredentialReq]) (*connect_go.Response[delivery.DeleteCredentialRes], error)
	UpdateCredential(context.Context, *connect_go.Request[delivery.UpdateCredentialReq]) (*connect_go.Response[delivery.UpdateCredentialRes], error)
}

// NewDeliveryApiClient constructs a client for the api.v1alpha1.delivery.DeliveryApi service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDeliveryApiClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) DeliveryApiClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &deliveryApiClient{
		createTransferConfig: connect_go.NewClient[delivery.CreateTransferConfigReq, delivery.CreateTransferConfigRes](
			httpClient,
			baseURL+DeliveryApiCreateTransferConfigProcedure,
			opts...,
		),
		listTransferConfigs: connect_go.NewClient[delivery.ListTransferConfigsReq, delivery.ListTransferConfigsRes](
			httpClient,
			baseURL+DeliveryApiListTransferConfigsProcedure,
			opts...,
		),
		listTransferConfigsByCredentialID: connect_go.NewClient[delivery.ListTransferConfigsByCredentialIDReq, delivery.ListTransferConfigsByCredentialIDRes](
			httpClient,
			baseURL+DeliveryApiListTransferConfigsByCredentialIDProcedure,
			opts...,
		),
		updateTransferConfig: connect_go.NewClient[delivery.UpdateTransferConfigReq, delivery.UpdateTransferConfigRes](
			httpClient,
			baseURL+DeliveryApiUpdateTransferConfigProcedure,
			opts...,
		),
		deleteTransferConfig: connect_go.NewClient[delivery.DeleteTransferConfigReq, delivery.DeleteTransferConfigRes](
			httpClient,
			baseURL+DeliveryApiDeleteTransferConfigProcedure,
			opts...,
		),
		getTransferConfig: connect_go.NewClient[delivery.GetTransferConfigReq, delivery.GetTransferConfigRes](
			httpClient,
			baseURL+DeliveryApiGetTransferConfigProcedure,
			opts...,
		),
		getTransferConfigByName: connect_go.NewClient[delivery.GetTransferConfigByNameReq, delivery.GetTransferConfigByNameRes](
			httpClient,
			baseURL+DeliveryApiGetTransferConfigByNameProcedure,
			opts...,
		),
		listHistory: connect_go.NewClient[delivery.ListHistoryReq, delivery.ListHistoryRes](
			httpClient,
			baseURL+DeliveryApiListHistoryProcedure,
			opts...,
		),
		listHistoryByTransferConfig: connect_go.NewClient[delivery.ListHistoryByTransferConfigReq, delivery.ListHistoryByTransferConfigRes](
			httpClient,
			baseURL+DeliveryApiListHistoryByTransferConfigProcedure,
			opts...,
		),
		listCredentials: connect_go.NewClient[delivery.ListCredentialsReq, delivery.ListCredentialsRes](
			httpClient,
			baseURL+DeliveryApiListCredentialsProcedure,
			opts...,
		),
		getCredential: connect_go.NewClient[delivery.GetCredentialReq, delivery.GetCredentialRes](
			httpClient,
			baseURL+DeliveryApiGetCredentialProcedure,
			opts...,
		),
		createCredential: connect_go.NewClient[delivery.CreateCredentialReq, delivery.CreateCredentialRes](
			httpClient,
			baseURL+DeliveryApiCreateCredentialProcedure,
			opts...,
		),
		deleteCredential: connect_go.NewClient[delivery.DeleteCredentialReq, delivery.DeleteCredentialRes](
			httpClient,
			baseURL+DeliveryApiDeleteCredentialProcedure,
			opts...,
		),
		updateCredential: connect_go.NewClient[delivery.UpdateCredentialReq, delivery.UpdateCredentialRes](
			httpClient,
			baseURL+DeliveryApiUpdateCredentialProcedure,
			opts...,
		),
	}
}

// deliveryApiClient implements DeliveryApiClient.
type deliveryApiClient struct {
	createTransferConfig              *connect_go.Client[delivery.CreateTransferConfigReq, delivery.CreateTransferConfigRes]
	listTransferConfigs               *connect_go.Client[delivery.ListTransferConfigsReq, delivery.ListTransferConfigsRes]
	listTransferConfigsByCredentialID *connect_go.Client[delivery.ListTransferConfigsByCredentialIDReq, delivery.ListTransferConfigsByCredentialIDRes]
	updateTransferConfig              *connect_go.Client[delivery.UpdateTransferConfigReq, delivery.UpdateTransferConfigRes]
	deleteTransferConfig              *connect_go.Client[delivery.DeleteTransferConfigReq, delivery.DeleteTransferConfigRes]
	getTransferConfig                 *connect_go.Client[delivery.GetTransferConfigReq, delivery.GetTransferConfigRes]
	getTransferConfigByName           *connect_go.Client[delivery.GetTransferConfigByNameReq, delivery.GetTransferConfigByNameRes]
	listHistory                       *connect_go.Client[delivery.ListHistoryReq, delivery.ListHistoryRes]
	listHistoryByTransferConfig       *connect_go.Client[delivery.ListHistoryByTransferConfigReq, delivery.ListHistoryByTransferConfigRes]
	listCredentials                   *connect_go.Client[delivery.ListCredentialsReq, delivery.ListCredentialsRes]
	getCredential                     *connect_go.Client[delivery.GetCredentialReq, delivery.GetCredentialRes]
	createCredential                  *connect_go.Client[delivery.CreateCredentialReq, delivery.CreateCredentialRes]
	deleteCredential                  *connect_go.Client[delivery.DeleteCredentialReq, delivery.DeleteCredentialRes]
	updateCredential                  *connect_go.Client[delivery.UpdateCredentialReq, delivery.UpdateCredentialRes]
}

// CreateTransferConfig calls api.v1alpha1.delivery.DeliveryApi.CreateTransferConfig.
func (c *deliveryApiClient) CreateTransferConfig(ctx context.Context, req *connect_go.Request[delivery.CreateTransferConfigReq]) (*connect_go.Response[delivery.CreateTransferConfigRes], error) {
	return c.createTransferConfig.CallUnary(ctx, req)
}

// ListTransferConfigs calls api.v1alpha1.delivery.DeliveryApi.ListTransferConfigs.
func (c *deliveryApiClient) ListTransferConfigs(ctx context.Context, req *connect_go.Request[delivery.ListTransferConfigsReq]) (*connect_go.Response[delivery.ListTransferConfigsRes], error) {
	return c.listTransferConfigs.CallUnary(ctx, req)
}

// ListTransferConfigsByCredentialID calls
// api.v1alpha1.delivery.DeliveryApi.ListTransferConfigsByCredentialID.
func (c *deliveryApiClient) ListTransferConfigsByCredentialID(ctx context.Context, req *connect_go.Request[delivery.ListTransferConfigsByCredentialIDReq]) (*connect_go.Response[delivery.ListTransferConfigsByCredentialIDRes], error) {
	return c.listTransferConfigsByCredentialID.CallUnary(ctx, req)
}

// UpdateTransferConfig calls api.v1alpha1.delivery.DeliveryApi.UpdateTransferConfig.
func (c *deliveryApiClient) UpdateTransferConfig(ctx context.Context, req *connect_go.Request[delivery.UpdateTransferConfigReq]) (*connect_go.Response[delivery.UpdateTransferConfigRes], error) {
	return c.updateTransferConfig.CallUnary(ctx, req)
}

// DeleteTransferConfig calls api.v1alpha1.delivery.DeliveryApi.DeleteTransferConfig.
func (c *deliveryApiClient) DeleteTransferConfig(ctx context.Context, req *connect_go.Request[delivery.DeleteTransferConfigReq]) (*connect_go.Response[delivery.DeleteTransferConfigRes], error) {
	return c.deleteTransferConfig.CallUnary(ctx, req)
}

// GetTransferConfig calls api.v1alpha1.delivery.DeliveryApi.GetTransferConfig.
func (c *deliveryApiClient) GetTransferConfig(ctx context.Context, req *connect_go.Request[delivery.GetTransferConfigReq]) (*connect_go.Response[delivery.GetTransferConfigRes], error) {
	return c.getTransferConfig.CallUnary(ctx, req)
}

// GetTransferConfigByName calls api.v1alpha1.delivery.DeliveryApi.GetTransferConfigByName.
func (c *deliveryApiClient) GetTransferConfigByName(ctx context.Context, req *connect_go.Request[delivery.GetTransferConfigByNameReq]) (*connect_go.Response[delivery.GetTransferConfigByNameRes], error) {
	return c.getTransferConfigByName.CallUnary(ctx, req)
}

// ListHistory calls api.v1alpha1.delivery.DeliveryApi.ListHistory.
func (c *deliveryApiClient) ListHistory(ctx context.Context, req *connect_go.Request[delivery.ListHistoryReq]) (*connect_go.Response[delivery.ListHistoryRes], error) {
	return c.listHistory.CallUnary(ctx, req)
}

// ListHistoryByTransferConfig calls api.v1alpha1.delivery.DeliveryApi.ListHistoryByTransferConfig.
func (c *deliveryApiClient) ListHistoryByTransferConfig(ctx context.Context, req *connect_go.Request[delivery.ListHistoryByTransferConfigReq]) (*connect_go.Response[delivery.ListHistoryByTransferConfigRes], error) {
	return c.listHistoryByTransferConfig.CallUnary(ctx, req)
}

// ListCredentials calls api.v1alpha1.delivery.DeliveryApi.ListCredentials.
func (c *deliveryApiClient) ListCredentials(ctx context.Context, req *connect_go.Request[delivery.ListCredentialsReq]) (*connect_go.Response[delivery.ListCredentialsRes], error) {
	return c.listCredentials.CallUnary(ctx, req)
}

// GetCredential calls api.v1alpha1.delivery.DeliveryApi.GetCredential.
func (c *deliveryApiClient) GetCredential(ctx context.Context, req *connect_go.Request[delivery.GetCredentialReq]) (*connect_go.Response[delivery.GetCredentialRes], error) {
	return c.getCredential.CallUnary(ctx, req)
}

// CreateCredential calls api.v1alpha1.delivery.DeliveryApi.CreateCredential.
func (c *deliveryApiClient) CreateCredential(ctx context.Context, req *connect_go.Request[delivery.CreateCredentialReq]) (*connect_go.Response[delivery.CreateCredentialRes], error) {
	return c.createCredential.CallUnary(ctx, req)
}

// DeleteCredential calls api.v1alpha1.delivery.DeliveryApi.DeleteCredential.
func (c *deliveryApiClient) DeleteCredential(ctx context.Context, req *connect_go.Request[delivery.DeleteCredentialReq]) (*connect_go.Response[delivery.DeleteCredentialRes], error) {
	return c.deleteCredential.CallUnary(ctx, req)
}

// UpdateCredential calls api.v1alpha1.delivery.DeliveryApi.UpdateCredential.
func (c *deliveryApiClient) UpdateCredential(ctx context.Context, req *connect_go.Request[delivery.UpdateCredentialReq]) (*connect_go.Response[delivery.UpdateCredentialRes], error) {
	return c.updateCredential.CallUnary(ctx, req)
}

// DeliveryApiHandler is an implementation of the api.v1alpha1.delivery.DeliveryApi service.
type DeliveryApiHandler interface {
	CreateTransferConfig(context.Context, *connect_go.Request[delivery.CreateTransferConfigReq]) (*connect_go.Response[delivery.CreateTransferConfigRes], error)
	ListTransferConfigs(context.Context, *connect_go.Request[delivery.ListTransferConfigsReq]) (*connect_go.Response[delivery.ListTransferConfigsRes], error)
	ListTransferConfigsByCredentialID(context.Context, *connect_go.Request[delivery.ListTransferConfigsByCredentialIDReq]) (*connect_go.Response[delivery.ListTransferConfigsByCredentialIDRes], error)
	UpdateTransferConfig(context.Context, *connect_go.Request[delivery.UpdateTransferConfigReq]) (*connect_go.Response[delivery.UpdateTransferConfigRes], error)
	DeleteTransferConfig(context.Context, *connect_go.Request[delivery.DeleteTransferConfigReq]) (*connect_go.Response[delivery.DeleteTransferConfigRes], error)
	GetTransferConfig(context.Context, *connect_go.Request[delivery.GetTransferConfigReq]) (*connect_go.Response[delivery.GetTransferConfigRes], error)
	GetTransferConfigByName(context.Context, *connect_go.Request[delivery.GetTransferConfigByNameReq]) (*connect_go.Response[delivery.GetTransferConfigByNameRes], error)
	ListHistory(context.Context, *connect_go.Request[delivery.ListHistoryReq]) (*connect_go.Response[delivery.ListHistoryRes], error)
	ListHistoryByTransferConfig(context.Context, *connect_go.Request[delivery.ListHistoryByTransferConfigReq]) (*connect_go.Response[delivery.ListHistoryByTransferConfigRes], error)
	ListCredentials(context.Context, *connect_go.Request[delivery.ListCredentialsReq]) (*connect_go.Response[delivery.ListCredentialsRes], error)
	GetCredential(context.Context, *connect_go.Request[delivery.GetCredentialReq]) (*connect_go.Response[delivery.GetCredentialRes], error)
	CreateCredential(context.Context, *connect_go.Request[delivery.CreateCredentialReq]) (*connect_go.Response[delivery.CreateCredentialRes], error)
	DeleteCredential(context.Context, *connect_go.Request[delivery.DeleteCredentialReq]) (*connect_go.Response[delivery.DeleteCredentialRes], error)
	UpdateCredential(context.Context, *connect_go.Request[delivery.UpdateCredentialReq]) (*connect_go.Response[delivery.UpdateCredentialRes], error)
}

// NewDeliveryApiHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDeliveryApiHandler(svc DeliveryApiHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	deliveryApiCreateTransferConfigHandler := connect_go.NewUnaryHandler(
		DeliveryApiCreateTransferConfigProcedure,
		svc.CreateTransferConfig,
		opts...,
	)
	deliveryApiListTransferConfigsHandler := connect_go.NewUnaryHandler(
		DeliveryApiListTransferConfigsProcedure,
		svc.ListTransferConfigs,
		opts...,
	)
	deliveryApiListTransferConfigsByCredentialIDHandler := connect_go.NewUnaryHandler(
		DeliveryApiListTransferConfigsByCredentialIDProcedure,
		svc.ListTransferConfigsByCredentialID,
		opts...,
	)
	deliveryApiUpdateTransferConfigHandler := connect_go.NewUnaryHandler(
		DeliveryApiUpdateTransferConfigProcedure,
		svc.UpdateTransferConfig,
		opts...,
	)
	deliveryApiDeleteTransferConfigHandler := connect_go.NewUnaryHandler(
		DeliveryApiDeleteTransferConfigProcedure,
		svc.DeleteTransferConfig,
		opts...,
	)
	deliveryApiGetTransferConfigHandler := connect_go.NewUnaryHandler(
		DeliveryApiGetTransferConfigProcedure,
		svc.GetTransferConfig,
		opts...,
	)
	deliveryApiGetTransferConfigByNameHandler := connect_go.NewUnaryHandler(
		DeliveryApiGetTransferConfigByNameProcedure,
		svc.GetTransferConfigByName,
		opts...,
	)
	deliveryApiListHistoryHandler := connect_go.NewUnaryHandler(
		DeliveryApiListHistoryProcedure,
		svc.ListHistory,
		opts...,
	)
	deliveryApiListHistoryByTransferConfigHandler := connect_go.NewUnaryHandler(
		DeliveryApiListHistoryByTransferConfigProcedure,
		svc.ListHistoryByTransferConfig,
		opts...,
	)
	deliveryApiListCredentialsHandler := connect_go.NewUnaryHandler(
		DeliveryApiListCredentialsProcedure,
		svc.ListCredentials,
		opts...,
	)
	deliveryApiGetCredentialHandler := connect_go.NewUnaryHandler(
		DeliveryApiGetCredentialProcedure,
		svc.GetCredential,
		opts...,
	)
	deliveryApiCreateCredentialHandler := connect_go.NewUnaryHandler(
		DeliveryApiCreateCredentialProcedure,
		svc.CreateCredential,
		opts...,
	)
	deliveryApiDeleteCredentialHandler := connect_go.NewUnaryHandler(
		DeliveryApiDeleteCredentialProcedure,
		svc.DeleteCredential,
		opts...,
	)
	deliveryApiUpdateCredentialHandler := connect_go.NewUnaryHandler(
		DeliveryApiUpdateCredentialProcedure,
		svc.UpdateCredential,
		opts...,
	)
	return "/api.v1alpha1.delivery.DeliveryApi/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DeliveryApiCreateTransferConfigProcedure:
			deliveryApiCreateTransferConfigHandler.ServeHTTP(w, r)
		case DeliveryApiListTransferConfigsProcedure:
			deliveryApiListTransferConfigsHandler.ServeHTTP(w, r)
		case DeliveryApiListTransferConfigsByCredentialIDProcedure:
			deliveryApiListTransferConfigsByCredentialIDHandler.ServeHTTP(w, r)
		case DeliveryApiUpdateTransferConfigProcedure:
			deliveryApiUpdateTransferConfigHandler.ServeHTTP(w, r)
		case DeliveryApiDeleteTransferConfigProcedure:
			deliveryApiDeleteTransferConfigHandler.ServeHTTP(w, r)
		case DeliveryApiGetTransferConfigProcedure:
			deliveryApiGetTransferConfigHandler.ServeHTTP(w, r)
		case DeliveryApiGetTransferConfigByNameProcedure:
			deliveryApiGetTransferConfigByNameHandler.ServeHTTP(w, r)
		case DeliveryApiListHistoryProcedure:
			deliveryApiListHistoryHandler.ServeHTTP(w, r)
		case DeliveryApiListHistoryByTransferConfigProcedure:
			deliveryApiListHistoryByTransferConfigHandler.ServeHTTP(w, r)
		case DeliveryApiListCredentialsProcedure:
			deliveryApiListCredentialsHandler.ServeHTTP(w, r)
		case DeliveryApiGetCredentialProcedure:
			deliveryApiGetCredentialHandler.ServeHTTP(w, r)
		case DeliveryApiCreateCredentialProcedure:
			deliveryApiCreateCredentialHandler.ServeHTTP(w, r)
		case DeliveryApiDeleteCredentialProcedure:
			deliveryApiDeleteCredentialHandler.ServeHTTP(w, r)
		case DeliveryApiUpdateCredentialProcedure:
			deliveryApiUpdateCredentialHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDeliveryApiHandler returns CodeUnimplemented from all methods.
type UnimplementedDeliveryApiHandler struct{}

func (UnimplementedDeliveryApiHandler) CreateTransferConfig(context.Context, *connect_go.Request[delivery.CreateTransferConfigReq]) (*connect_go.Response[delivery.CreateTransferConfigRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.CreateTransferConfig is not implemented"))
}

func (UnimplementedDeliveryApiHandler) ListTransferConfigs(context.Context, *connect_go.Request[delivery.ListTransferConfigsReq]) (*connect_go.Response[delivery.ListTransferConfigsRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.ListTransferConfigs is not implemented"))
}

func (UnimplementedDeliveryApiHandler) ListTransferConfigsByCredentialID(context.Context, *connect_go.Request[delivery.ListTransferConfigsByCredentialIDReq]) (*connect_go.Response[delivery.ListTransferConfigsByCredentialIDRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.ListTransferConfigsByCredentialID is not implemented"))
}

func (UnimplementedDeliveryApiHandler) UpdateTransferConfig(context.Context, *connect_go.Request[delivery.UpdateTransferConfigReq]) (*connect_go.Response[delivery.UpdateTransferConfigRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.UpdateTransferConfig is not implemented"))
}

func (UnimplementedDeliveryApiHandler) DeleteTransferConfig(context.Context, *connect_go.Request[delivery.DeleteTransferConfigReq]) (*connect_go.Response[delivery.DeleteTransferConfigRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.DeleteTransferConfig is not implemented"))
}

func (UnimplementedDeliveryApiHandler) GetTransferConfig(context.Context, *connect_go.Request[delivery.GetTransferConfigReq]) (*connect_go.Response[delivery.GetTransferConfigRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.GetTransferConfig is not implemented"))
}

func (UnimplementedDeliveryApiHandler) GetTransferConfigByName(context.Context, *connect_go.Request[delivery.GetTransferConfigByNameReq]) (*connect_go.Response[delivery.GetTransferConfigByNameRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.GetTransferConfigByName is not implemented"))
}

func (UnimplementedDeliveryApiHandler) ListHistory(context.Context, *connect_go.Request[delivery.ListHistoryReq]) (*connect_go.Response[delivery.ListHistoryRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.ListHistory is not implemented"))
}

func (UnimplementedDeliveryApiHandler) ListHistoryByTransferConfig(context.Context, *connect_go.Request[delivery.ListHistoryByTransferConfigReq]) (*connect_go.Response[delivery.ListHistoryByTransferConfigRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.ListHistoryByTransferConfig is not implemented"))
}

func (UnimplementedDeliveryApiHandler) ListCredentials(context.Context, *connect_go.Request[delivery.ListCredentialsReq]) (*connect_go.Response[delivery.ListCredentialsRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.ListCredentials is not implemented"))
}

func (UnimplementedDeliveryApiHandler) GetCredential(context.Context, *connect_go.Request[delivery.GetCredentialReq]) (*connect_go.Response[delivery.GetCredentialRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.GetCredential is not implemented"))
}

func (UnimplementedDeliveryApiHandler) CreateCredential(context.Context, *connect_go.Request[delivery.CreateCredentialReq]) (*connect_go.Response[delivery.CreateCredentialRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.CreateCredential is not implemented"))
}

func (UnimplementedDeliveryApiHandler) DeleteCredential(context.Context, *connect_go.Request[delivery.DeleteCredentialReq]) (*connect_go.Response[delivery.DeleteCredentialRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.DeleteCredential is not implemented"))
}

func (UnimplementedDeliveryApiHandler) UpdateCredential(context.Context, *connect_go.Request[delivery.UpdateCredentialReq]) (*connect_go.Response[delivery.UpdateCredentialRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.delivery.DeliveryApi.UpdateCredential is not implemented"))
}
