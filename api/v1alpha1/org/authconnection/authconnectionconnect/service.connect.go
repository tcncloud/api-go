// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/org/authconnection/service.proto

package authconnectionconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	authconnection "github.com/tcncloud/api-go/api/v1alpha1/org/authconnection"
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
	// AuthConnectionServiceName is the fully-qualified name of the AuthConnectionService service.
	AuthConnectionServiceName = "api.v1alpha1.org.authconnection.AuthConnectionService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthConnectionServiceCreateAuthConnectionProcedure is the fully-qualified name of the
	// AuthConnectionService's CreateAuthConnection RPC.
	AuthConnectionServiceCreateAuthConnectionProcedure = "/api.v1alpha1.org.authconnection.AuthConnectionService/CreateAuthConnection"
	// AuthConnectionServiceGetAuthConnectionSettingsProcedure is the fully-qualified name of the
	// AuthConnectionService's GetAuthConnectionSettings RPC.
	AuthConnectionServiceGetAuthConnectionSettingsProcedure = "/api.v1alpha1.org.authconnection.AuthConnectionService/GetAuthConnectionSettings"
	// AuthConnectionServiceDeleteAuthConnectionProcedure is the fully-qualified name of the
	// AuthConnectionService's DeleteAuthConnection RPC.
	AuthConnectionServiceDeleteAuthConnectionProcedure = "/api.v1alpha1.org.authconnection.AuthConnectionService/DeleteAuthConnection"
	// AuthConnectionServiceUpdateAuthConnectionSecretProcedure is the fully-qualified name of the
	// AuthConnectionService's UpdateAuthConnectionSecret RPC.
	AuthConnectionServiceUpdateAuthConnectionSecretProcedure = "/api.v1alpha1.org.authconnection.AuthConnectionService/UpdateAuthConnectionSecret"
	// AuthConnectionServiceUpdateAuthConnectionGroupsProcedure is the fully-qualified name of the
	// AuthConnectionService's UpdateAuthConnectionGroups RPC.
	AuthConnectionServiceUpdateAuthConnectionGroupsProcedure = "/api.v1alpha1.org.authconnection.AuthConnectionService/UpdateAuthConnectionGroups"
)

// AuthConnectionServiceClient is a client for the
// api.v1alpha1.org.authconnection.AuthConnectionService service.
type AuthConnectionServiceClient interface {
	// CreateAuthConnection creates a new auth0 connection.
	CreateAuthConnection(context.Context, *connect_go.Request[authconnection.CreateAuthConnectionRequest]) (*connect_go.Response[authconnection.CreateAuthConnectionResponse], error)
	// GetAuthConnectionSettings gets auth0 connection settings.
	GetAuthConnectionSettings(context.Context, *connect_go.Request[authconnection.GetAuthConnectionSettingsRequest]) (*connect_go.Response[authconnection.GetAuthConnectionSettingsResponse], error)
	// DeleteAuthConnection removes the current orgs auth settings.
	DeleteAuthConnection(context.Context, *connect_go.Request[authconnection.DeleteAuthConnectionRequest]) (*connect_go.Response[authconnection.DeleteAuthConnectionResponse], error)
	// UpdateAuthConnectionSecret updates a connections secret.
	UpdateAuthConnectionSecret(context.Context, *connect_go.Request[authconnection.UpdateAuthConnectionSecretRequest]) (*connect_go.Response[authconnection.UpdateAuthConnectionSecretResponse], error)
	// UpdateAuthConnectionGroups updates a connections groups.
	UpdateAuthConnectionGroups(context.Context, *connect_go.Request[authconnection.UpdateAuthConnectionGroupsRequest]) (*connect_go.Response[authconnection.UpdateAuthConnectionGroupsResponse], error)
}

// NewAuthConnectionServiceClient constructs a client for the
// api.v1alpha1.org.authconnection.AuthConnectionService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthConnectionServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AuthConnectionServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authConnectionServiceClient{
		createAuthConnection: connect_go.NewClient[authconnection.CreateAuthConnectionRequest, authconnection.CreateAuthConnectionResponse](
			httpClient,
			baseURL+AuthConnectionServiceCreateAuthConnectionProcedure,
			opts...,
		),
		getAuthConnectionSettings: connect_go.NewClient[authconnection.GetAuthConnectionSettingsRequest, authconnection.GetAuthConnectionSettingsResponse](
			httpClient,
			baseURL+AuthConnectionServiceGetAuthConnectionSettingsProcedure,
			opts...,
		),
		deleteAuthConnection: connect_go.NewClient[authconnection.DeleteAuthConnectionRequest, authconnection.DeleteAuthConnectionResponse](
			httpClient,
			baseURL+AuthConnectionServiceDeleteAuthConnectionProcedure,
			opts...,
		),
		updateAuthConnectionSecret: connect_go.NewClient[authconnection.UpdateAuthConnectionSecretRequest, authconnection.UpdateAuthConnectionSecretResponse](
			httpClient,
			baseURL+AuthConnectionServiceUpdateAuthConnectionSecretProcedure,
			opts...,
		),
		updateAuthConnectionGroups: connect_go.NewClient[authconnection.UpdateAuthConnectionGroupsRequest, authconnection.UpdateAuthConnectionGroupsResponse](
			httpClient,
			baseURL+AuthConnectionServiceUpdateAuthConnectionGroupsProcedure,
			opts...,
		),
	}
}

// authConnectionServiceClient implements AuthConnectionServiceClient.
type authConnectionServiceClient struct {
	createAuthConnection       *connect_go.Client[authconnection.CreateAuthConnectionRequest, authconnection.CreateAuthConnectionResponse]
	getAuthConnectionSettings  *connect_go.Client[authconnection.GetAuthConnectionSettingsRequest, authconnection.GetAuthConnectionSettingsResponse]
	deleteAuthConnection       *connect_go.Client[authconnection.DeleteAuthConnectionRequest, authconnection.DeleteAuthConnectionResponse]
	updateAuthConnectionSecret *connect_go.Client[authconnection.UpdateAuthConnectionSecretRequest, authconnection.UpdateAuthConnectionSecretResponse]
	updateAuthConnectionGroups *connect_go.Client[authconnection.UpdateAuthConnectionGroupsRequest, authconnection.UpdateAuthConnectionGroupsResponse]
}

// CreateAuthConnection calls
// api.v1alpha1.org.authconnection.AuthConnectionService.CreateAuthConnection.
func (c *authConnectionServiceClient) CreateAuthConnection(ctx context.Context, req *connect_go.Request[authconnection.CreateAuthConnectionRequest]) (*connect_go.Response[authconnection.CreateAuthConnectionResponse], error) {
	return c.createAuthConnection.CallUnary(ctx, req)
}

// GetAuthConnectionSettings calls
// api.v1alpha1.org.authconnection.AuthConnectionService.GetAuthConnectionSettings.
func (c *authConnectionServiceClient) GetAuthConnectionSettings(ctx context.Context, req *connect_go.Request[authconnection.GetAuthConnectionSettingsRequest]) (*connect_go.Response[authconnection.GetAuthConnectionSettingsResponse], error) {
	return c.getAuthConnectionSettings.CallUnary(ctx, req)
}

// DeleteAuthConnection calls
// api.v1alpha1.org.authconnection.AuthConnectionService.DeleteAuthConnection.
func (c *authConnectionServiceClient) DeleteAuthConnection(ctx context.Context, req *connect_go.Request[authconnection.DeleteAuthConnectionRequest]) (*connect_go.Response[authconnection.DeleteAuthConnectionResponse], error) {
	return c.deleteAuthConnection.CallUnary(ctx, req)
}

// UpdateAuthConnectionSecret calls
// api.v1alpha1.org.authconnection.AuthConnectionService.UpdateAuthConnectionSecret.
func (c *authConnectionServiceClient) UpdateAuthConnectionSecret(ctx context.Context, req *connect_go.Request[authconnection.UpdateAuthConnectionSecretRequest]) (*connect_go.Response[authconnection.UpdateAuthConnectionSecretResponse], error) {
	return c.updateAuthConnectionSecret.CallUnary(ctx, req)
}

// UpdateAuthConnectionGroups calls
// api.v1alpha1.org.authconnection.AuthConnectionService.UpdateAuthConnectionGroups.
func (c *authConnectionServiceClient) UpdateAuthConnectionGroups(ctx context.Context, req *connect_go.Request[authconnection.UpdateAuthConnectionGroupsRequest]) (*connect_go.Response[authconnection.UpdateAuthConnectionGroupsResponse], error) {
	return c.updateAuthConnectionGroups.CallUnary(ctx, req)
}

// AuthConnectionServiceHandler is an implementation of the
// api.v1alpha1.org.authconnection.AuthConnectionService service.
type AuthConnectionServiceHandler interface {
	// CreateAuthConnection creates a new auth0 connection.
	CreateAuthConnection(context.Context, *connect_go.Request[authconnection.CreateAuthConnectionRequest]) (*connect_go.Response[authconnection.CreateAuthConnectionResponse], error)
	// GetAuthConnectionSettings gets auth0 connection settings.
	GetAuthConnectionSettings(context.Context, *connect_go.Request[authconnection.GetAuthConnectionSettingsRequest]) (*connect_go.Response[authconnection.GetAuthConnectionSettingsResponse], error)
	// DeleteAuthConnection removes the current orgs auth settings.
	DeleteAuthConnection(context.Context, *connect_go.Request[authconnection.DeleteAuthConnectionRequest]) (*connect_go.Response[authconnection.DeleteAuthConnectionResponse], error)
	// UpdateAuthConnectionSecret updates a connections secret.
	UpdateAuthConnectionSecret(context.Context, *connect_go.Request[authconnection.UpdateAuthConnectionSecretRequest]) (*connect_go.Response[authconnection.UpdateAuthConnectionSecretResponse], error)
	// UpdateAuthConnectionGroups updates a connections groups.
	UpdateAuthConnectionGroups(context.Context, *connect_go.Request[authconnection.UpdateAuthConnectionGroupsRequest]) (*connect_go.Response[authconnection.UpdateAuthConnectionGroupsResponse], error)
}

// NewAuthConnectionServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthConnectionServiceHandler(svc AuthConnectionServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	authConnectionServiceCreateAuthConnectionHandler := connect_go.NewUnaryHandler(
		AuthConnectionServiceCreateAuthConnectionProcedure,
		svc.CreateAuthConnection,
		opts...,
	)
	authConnectionServiceGetAuthConnectionSettingsHandler := connect_go.NewUnaryHandler(
		AuthConnectionServiceGetAuthConnectionSettingsProcedure,
		svc.GetAuthConnectionSettings,
		opts...,
	)
	authConnectionServiceDeleteAuthConnectionHandler := connect_go.NewUnaryHandler(
		AuthConnectionServiceDeleteAuthConnectionProcedure,
		svc.DeleteAuthConnection,
		opts...,
	)
	authConnectionServiceUpdateAuthConnectionSecretHandler := connect_go.NewUnaryHandler(
		AuthConnectionServiceUpdateAuthConnectionSecretProcedure,
		svc.UpdateAuthConnectionSecret,
		opts...,
	)
	authConnectionServiceUpdateAuthConnectionGroupsHandler := connect_go.NewUnaryHandler(
		AuthConnectionServiceUpdateAuthConnectionGroupsProcedure,
		svc.UpdateAuthConnectionGroups,
		opts...,
	)
	return "/api.v1alpha1.org.authconnection.AuthConnectionService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthConnectionServiceCreateAuthConnectionProcedure:
			authConnectionServiceCreateAuthConnectionHandler.ServeHTTP(w, r)
		case AuthConnectionServiceGetAuthConnectionSettingsProcedure:
			authConnectionServiceGetAuthConnectionSettingsHandler.ServeHTTP(w, r)
		case AuthConnectionServiceDeleteAuthConnectionProcedure:
			authConnectionServiceDeleteAuthConnectionHandler.ServeHTTP(w, r)
		case AuthConnectionServiceUpdateAuthConnectionSecretProcedure:
			authConnectionServiceUpdateAuthConnectionSecretHandler.ServeHTTP(w, r)
		case AuthConnectionServiceUpdateAuthConnectionGroupsProcedure:
			authConnectionServiceUpdateAuthConnectionGroupsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthConnectionServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthConnectionServiceHandler struct{}

func (UnimplementedAuthConnectionServiceHandler) CreateAuthConnection(context.Context, *connect_go.Request[authconnection.CreateAuthConnectionRequest]) (*connect_go.Response[authconnection.CreateAuthConnectionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.authconnection.AuthConnectionService.CreateAuthConnection is not implemented"))
}

func (UnimplementedAuthConnectionServiceHandler) GetAuthConnectionSettings(context.Context, *connect_go.Request[authconnection.GetAuthConnectionSettingsRequest]) (*connect_go.Response[authconnection.GetAuthConnectionSettingsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.authconnection.AuthConnectionService.GetAuthConnectionSettings is not implemented"))
}

func (UnimplementedAuthConnectionServiceHandler) DeleteAuthConnection(context.Context, *connect_go.Request[authconnection.DeleteAuthConnectionRequest]) (*connect_go.Response[authconnection.DeleteAuthConnectionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.authconnection.AuthConnectionService.DeleteAuthConnection is not implemented"))
}

func (UnimplementedAuthConnectionServiceHandler) UpdateAuthConnectionSecret(context.Context, *connect_go.Request[authconnection.UpdateAuthConnectionSecretRequest]) (*connect_go.Response[authconnection.UpdateAuthConnectionSecretResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.authconnection.AuthConnectionService.UpdateAuthConnectionSecret is not implemented"))
}

func (UnimplementedAuthConnectionServiceHandler) UpdateAuthConnectionGroups(context.Context, *connect_go.Request[authconnection.UpdateAuthConnectionGroupsRequest]) (*connect_go.Response[authconnection.UpdateAuthConnectionGroupsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.org.authconnection.AuthConnectionService.UpdateAuthConnectionGroups is not implemented"))
}
