// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/org/idp/service.proto

package idpconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	idp "github.com/tcncloud/api-go/api/v1alpha1/org/idp"
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
	// IdpServiceName is the fully-qualified name of the IdpService service.
	IdpServiceName = "api.v1alpha1.org.idp.IdpService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IdpServiceCreateClientProcedure is the fully-qualified name of the IdpService's CreateClient RPC.
	IdpServiceCreateClientProcedure = "/api.v1alpha1.org.idp.IdpService/CreateClient"
	// IdpServiceUpdateClientProcedure is the fully-qualified name of the IdpService's UpdateClient RPC.
	IdpServiceUpdateClientProcedure = "/api.v1alpha1.org.idp.IdpService/UpdateClient"
	// IdpServiceDeleteClientProcedure is the fully-qualified name of the IdpService's DeleteClient RPC.
	IdpServiceDeleteClientProcedure = "/api.v1alpha1.org.idp.IdpService/DeleteClient"
	// IdpServiceListClientsProcedure is the fully-qualified name of the IdpService's ListClients RPC.
	IdpServiceListClientsProcedure = "/api.v1alpha1.org.idp.IdpService/ListClients"
)

// IdpServiceClient is a client for the api.v1alpha1.org.idp.IdpService service.
type IdpServiceClient interface {
	// CreateClient creates a client.
	CreateClient(context.Context, *connect_go.Request[idp.CreateClientRequest]) (*connect_go.Response[idp.CreateClientResponse], error)
	// UpdateClient updates an existing client
	UpdateClient(context.Context, *connect_go.Request[idp.UpdateClientRequest]) (*connect_go.Response[idp.UpdateClientResponse], error)
	// DeleteClient deletes the provided client.
	DeleteClient(context.Context, *connect_go.Request[idp.DeleteClientRequest]) (*connect_go.Response[idp.DeleteClientResponse], error)
	// ListClients returns all clients.
	ListClients(context.Context, *connect_go.Request[idp.ListClientsRequest]) (*connect_go.Response[idp.ListClientsResponse], error)
}

// NewIdpServiceClient constructs a client for the api.v1alpha1.org.idp.IdpService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIdpServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IdpServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &idpServiceClient{
		createClient: connect_go.NewClient[idp.CreateClientRequest, idp.CreateClientResponse](
			httpClient,
			baseURL+IdpServiceCreateClientProcedure,
			opts...,
		),
		updateClient: connect_go.NewClient[idp.UpdateClientRequest, idp.UpdateClientResponse](
			httpClient,
			baseURL+IdpServiceUpdateClientProcedure,
			opts...,
		),
		deleteClient: connect_go.NewClient[idp.DeleteClientRequest, idp.DeleteClientResponse](
			httpClient,
			baseURL+IdpServiceDeleteClientProcedure,
			opts...,
		),
		listClients: connect_go.NewClient[idp.ListClientsRequest, idp.ListClientsResponse](
			httpClient,
			baseURL+IdpServiceListClientsProcedure,
			opts...,
		),
	}
}

// idpServiceClient implements IdpServiceClient.
type idpServiceClient struct {
	createClient *connect_go.Client[idp.CreateClientRequest, idp.CreateClientResponse]
	updateClient *connect_go.Client[idp.UpdateClientRequest, idp.UpdateClientResponse]
	deleteClient *connect_go.Client[idp.DeleteClientRequest, idp.DeleteClientResponse]
	listClients  *connect_go.Client[idp.ListClientsRequest, idp.ListClientsResponse]
}

// CreateClient calls api.v1alpha1.org.idp.IdpService.CreateClient.
func (c *idpServiceClient) CreateClient(ctx context.Context, req *connect_go.Request[idp.CreateClientRequest]) (*connect_go.Response[idp.CreateClientResponse], error) {
	return c.createClient.CallUnary(ctx, req)
}

// UpdateClient calls api.v1alpha1.org.idp.IdpService.UpdateClient.
func (c *idpServiceClient) UpdateClient(ctx context.Context, req *connect_go.Request[idp.UpdateClientRequest]) (*connect_go.Response[idp.UpdateClientResponse], error) {
	return c.updateClient.CallUnary(ctx, req)
}

// DeleteClient calls api.v1alpha1.org.idp.IdpService.DeleteClient.
func (c *idpServiceClient) DeleteClient(ctx context.Context, req *connect_go.Request[idp.DeleteClientRequest]) (*connect_go.Response[idp.DeleteClientResponse], error) {
	return c.deleteClient.CallUnary(ctx, req)
}

// ListClients calls api.v1alpha1.org.idp.IdpService.ListClients.
func (c *idpServiceClient) ListClients(ctx context.Context, req *connect_go.Request[idp.ListClientsRequest]) (*connect_go.Response[idp.ListClientsResponse], error) {
	return c.listClients.CallUnary(ctx, req)
}

// IdpServiceHandler is an implementation of the api.v1alpha1.org.idp.IdpService service.
type IdpServiceHandler interface {
	// CreateClient creates a client.
	CreateClient(context.Context, *connect_go.Request[idp.CreateClientRequest]) (*connect_go.Response[idp.CreateClientResponse], error)
	// UpdateClient updates an existing client
	UpdateClient(context.Context, *connect_go.Request[idp.UpdateClientRequest]) (*connect_go.Response[idp.UpdateClientResponse], error)
	// DeleteClient deletes the provided client.
	DeleteClient(context.Context, *connect_go.Request[idp.DeleteClientRequest]) (*connect_go.Response[idp.DeleteClientResponse], error)
	// ListClients returns all clients.
	ListClients(context.Context, *connect_go.Request[idp.ListClientsRequest]) (*connect_go.Response[idp.ListClientsResponse], error)
}

// NewIdpServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIdpServiceHandler(svc IdpServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	idpServiceCreateClientHandler := connect_go.NewUnaryHandler(
		IdpServiceCreateClientProcedure,
		svc.CreateClient,
		opts...,
	)
	idpServiceUpdateClientHandler := connect_go.NewUnaryHandler(
		IdpServiceUpdateClientProcedure,
		svc.UpdateClient,
		opts...,
	)
	idpServiceDeleteClientHandler := connect_go.NewUnaryHandler(
		IdpServiceDeleteClientProcedure,
		svc.DeleteClient,
		opts...,
	)
	idpServiceListClientsHandler := connect_go.NewUnaryHandler(
		IdpServiceListClientsProcedure,
		svc.ListClients,
		opts...,
	)
	return "/api.v1alpha1.org.idp.IdpService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IdpServiceCreateClientProcedure:
			idpServiceCreateClientHandler.ServeHTTP(w, r)
		case IdpServiceUpdateClientProcedure:
			idpServiceUpdateClientHandler.ServeHTTP(w, r)
		case IdpServiceDeleteClientProcedure:
			idpServiceDeleteClientHandler.ServeHTTP(w, r)
		case IdpServiceListClientsProcedure:
			idpServiceListClientsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIdpServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIdpServiceHandler struct{}

func (UnimplementedIdpServiceHandler) CreateClient(context.Context, *connect_go.Request[idp.CreateClientRequest]) (*connect_go.Response[idp.CreateClientResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.idp.IdpService.CreateClient is not implemented"))
}

func (UnimplementedIdpServiceHandler) UpdateClient(context.Context, *connect_go.Request[idp.UpdateClientRequest]) (*connect_go.Response[idp.UpdateClientResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.idp.IdpService.UpdateClient is not implemented"))
}

func (UnimplementedIdpServiceHandler) DeleteClient(context.Context, *connect_go.Request[idp.DeleteClientRequest]) (*connect_go.Response[idp.DeleteClientResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.idp.IdpService.DeleteClient is not implemented"))
}

func (UnimplementedIdpServiceHandler) ListClients(context.Context, *connect_go.Request[idp.ListClientsRequest]) (*connect_go.Response[idp.ListClientsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.idp.IdpService.ListClients is not implemented"))
}
