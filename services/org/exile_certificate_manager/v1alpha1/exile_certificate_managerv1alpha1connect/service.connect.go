// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: services/org/exile_certificate_manager/v1alpha1/service.proto

package exile_certificate_managerv1alpha1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1alpha1 "github.com/tcncloud/api-go/services/org/exile_certificate_manager/v1alpha1"
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
	// ExileCertificateManagerServiceName is the fully-qualified name of the
	// ExileCertificateManagerService service.
	ExileCertificateManagerServiceName = "services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ExileCertificateManagerServiceCreateExileCertificateProcedure is the fully-qualified name of the
	// ExileCertificateManagerService's CreateExileCertificate RPC.
	ExileCertificateManagerServiceCreateExileCertificateProcedure = "/services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService/CreateExileCertificate"
	// ExileCertificateManagerServiceRevokeExileCertificateProcedure is the fully-qualified name of the
	// ExileCertificateManagerService's RevokeExileCertificate RPC.
	ExileCertificateManagerServiceRevokeExileCertificateProcedure = "/services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService/RevokeExileCertificate"
	// ExileCertificateManagerServiceAssignExileConfigurationProcedure is the fully-qualified name of
	// the ExileCertificateManagerService's AssignExileConfiguration RPC.
	ExileCertificateManagerServiceAssignExileConfigurationProcedure = "/services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService/AssignExileConfiguration"
	// ExileCertificateManagerServiceUnassignExileConfigurationProcedure is the fully-qualified name of
	// the ExileCertificateManagerService's UnassignExileConfiguration RPC.
	ExileCertificateManagerServiceUnassignExileConfigurationProcedure = "/services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService/UnassignExileConfiguration"
	// ExileCertificateManagerServiceListExileCertificatesProcedure is the fully-qualified name of the
	// ExileCertificateManagerService's ListExileCertificates RPC.
	ExileCertificateManagerServiceListExileCertificatesProcedure = "/services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService/ListExileCertificates"
	// ExileCertificateManagerServiceCreateExileConfigurationProcedure is the fully-qualified name of
	// the ExileCertificateManagerService's CreateExileConfiguration RPC.
	ExileCertificateManagerServiceCreateExileConfigurationProcedure = "/services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService/CreateExileConfiguration"
	// ExileCertificateManagerServiceUpdateExileConfigurationProcedure is the fully-qualified name of
	// the ExileCertificateManagerService's UpdateExileConfiguration RPC.
	ExileCertificateManagerServiceUpdateExileConfigurationProcedure = "/services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService/UpdateExileConfiguration"
	// ExileCertificateManagerServiceDeleteExileConfigurationProcedure is the fully-qualified name of
	// the ExileCertificateManagerService's DeleteExileConfiguration RPC.
	ExileCertificateManagerServiceDeleteExileConfigurationProcedure = "/services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService/DeleteExileConfiguration"
	// ExileCertificateManagerServiceListExileConfigurationsProcedure is the fully-qualified name of the
	// ExileCertificateManagerService's ListExileConfigurations RPC.
	ExileCertificateManagerServiceListExileConfigurationsProcedure = "/services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService/ListExileConfigurations"
)

// ExileCertificateManagerServiceClient is a client for the
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService service.
type ExileCertificateManagerServiceClient interface {
	// CreateExileCertificate creates a new TLS certificate and
	// returns the exile ceritificate for the current organization.
	CreateExileCertificate(context.Context, *connect_go.Request[v1alpha1.CreateExileCertificateRequest]) (*connect_go.Response[v1alpha1.CreateExileCertificateResponse], error)
	// RevokeExileCertificate deletes a exile ceritificate for the current organization.
	RevokeExileCertificate(context.Context, *connect_go.Request[v1alpha1.RevokeExileCertificateRequest]) (*connect_go.Response[v1alpha1.RevokeExileCertificateResponse], error)
	// AssignExileConfiguration assigns a configuration to a certificate.
	AssignExileConfiguration(context.Context, *connect_go.Request[v1alpha1.AssignExileConfigurationRequest]) (*connect_go.Response[v1alpha1.AssignExileConfigurationResponse], error)
	// UnassignExileConfiguration unassigns a configuration from a certificate.
	UnassignExileConfiguration(context.Context, *connect_go.Request[v1alpha1.UnassignExileConfigurationRequest]) (*connect_go.Response[v1alpha1.UnassignExileConfigurationResponse], error)
	// ListExileCertificates returns a list of exile ceritificate for the current organization.
	ListExileCertificates(context.Context, *connect_go.Request[v1alpha1.ListExileCertificatesRequest]) (*connect_go.Response[v1alpha1.ListExileCertificatesResponse], error)
	// CreateExileConfiguration creates a new exile configuration for the current organization.
	CreateExileConfiguration(context.Context, *connect_go.Request[v1alpha1.CreateExileConfigurationRequest]) (*connect_go.Response[v1alpha1.CreateExileConfigurationResponse], error)
	// UpdateExileConfiguration updates a exile configuration for the current organization.
	UpdateExileConfiguration(context.Context, *connect_go.Request[v1alpha1.UpdateExileConfigurationRequest]) (*connect_go.Response[v1alpha1.UpdateExileConfigurationResponse], error)
	// DeleteExileConfiguration deletes a exile configuration for the current organization.
	// The configuration can only be deleted if it is not assigned to any certificates.
	DeleteExileConfiguration(context.Context, *connect_go.Request[v1alpha1.DeleteExileConfigurationRequest]) (*connect_go.Response[v1alpha1.DeleteExileConfigurationResponse], error)
	// ListExileConfigurations returns a list of exile configurations for the current organization.
	ListExileConfigurations(context.Context, *connect_go.Request[v1alpha1.ListExileConfigurationsRequest]) (*connect_go.Response[v1alpha1.ListExileConfigurationsResponse], error)
}

// NewExileCertificateManagerServiceClient constructs a client for the
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewExileCertificateManagerServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ExileCertificateManagerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &exileCertificateManagerServiceClient{
		createExileCertificate: connect_go.NewClient[v1alpha1.CreateExileCertificateRequest, v1alpha1.CreateExileCertificateResponse](
			httpClient,
			baseURL+ExileCertificateManagerServiceCreateExileCertificateProcedure,
			opts...,
		),
		revokeExileCertificate: connect_go.NewClient[v1alpha1.RevokeExileCertificateRequest, v1alpha1.RevokeExileCertificateResponse](
			httpClient,
			baseURL+ExileCertificateManagerServiceRevokeExileCertificateProcedure,
			opts...,
		),
		assignExileConfiguration: connect_go.NewClient[v1alpha1.AssignExileConfigurationRequest, v1alpha1.AssignExileConfigurationResponse](
			httpClient,
			baseURL+ExileCertificateManagerServiceAssignExileConfigurationProcedure,
			opts...,
		),
		unassignExileConfiguration: connect_go.NewClient[v1alpha1.UnassignExileConfigurationRequest, v1alpha1.UnassignExileConfigurationResponse](
			httpClient,
			baseURL+ExileCertificateManagerServiceUnassignExileConfigurationProcedure,
			opts...,
		),
		listExileCertificates: connect_go.NewClient[v1alpha1.ListExileCertificatesRequest, v1alpha1.ListExileCertificatesResponse](
			httpClient,
			baseURL+ExileCertificateManagerServiceListExileCertificatesProcedure,
			opts...,
		),
		createExileConfiguration: connect_go.NewClient[v1alpha1.CreateExileConfigurationRequest, v1alpha1.CreateExileConfigurationResponse](
			httpClient,
			baseURL+ExileCertificateManagerServiceCreateExileConfigurationProcedure,
			opts...,
		),
		updateExileConfiguration: connect_go.NewClient[v1alpha1.UpdateExileConfigurationRequest, v1alpha1.UpdateExileConfigurationResponse](
			httpClient,
			baseURL+ExileCertificateManagerServiceUpdateExileConfigurationProcedure,
			opts...,
		),
		deleteExileConfiguration: connect_go.NewClient[v1alpha1.DeleteExileConfigurationRequest, v1alpha1.DeleteExileConfigurationResponse](
			httpClient,
			baseURL+ExileCertificateManagerServiceDeleteExileConfigurationProcedure,
			opts...,
		),
		listExileConfigurations: connect_go.NewClient[v1alpha1.ListExileConfigurationsRequest, v1alpha1.ListExileConfigurationsResponse](
			httpClient,
			baseURL+ExileCertificateManagerServiceListExileConfigurationsProcedure,
			opts...,
		),
	}
}

// exileCertificateManagerServiceClient implements ExileCertificateManagerServiceClient.
type exileCertificateManagerServiceClient struct {
	createExileCertificate     *connect_go.Client[v1alpha1.CreateExileCertificateRequest, v1alpha1.CreateExileCertificateResponse]
	revokeExileCertificate     *connect_go.Client[v1alpha1.RevokeExileCertificateRequest, v1alpha1.RevokeExileCertificateResponse]
	assignExileConfiguration   *connect_go.Client[v1alpha1.AssignExileConfigurationRequest, v1alpha1.AssignExileConfigurationResponse]
	unassignExileConfiguration *connect_go.Client[v1alpha1.UnassignExileConfigurationRequest, v1alpha1.UnassignExileConfigurationResponse]
	listExileCertificates      *connect_go.Client[v1alpha1.ListExileCertificatesRequest, v1alpha1.ListExileCertificatesResponse]
	createExileConfiguration   *connect_go.Client[v1alpha1.CreateExileConfigurationRequest, v1alpha1.CreateExileConfigurationResponse]
	updateExileConfiguration   *connect_go.Client[v1alpha1.UpdateExileConfigurationRequest, v1alpha1.UpdateExileConfigurationResponse]
	deleteExileConfiguration   *connect_go.Client[v1alpha1.DeleteExileConfigurationRequest, v1alpha1.DeleteExileConfigurationResponse]
	listExileConfigurations    *connect_go.Client[v1alpha1.ListExileConfigurationsRequest, v1alpha1.ListExileConfigurationsResponse]
}

// CreateExileCertificate calls
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.CreateExileCertificate.
func (c *exileCertificateManagerServiceClient) CreateExileCertificate(ctx context.Context, req *connect_go.Request[v1alpha1.CreateExileCertificateRequest]) (*connect_go.Response[v1alpha1.CreateExileCertificateResponse], error) {
	return c.createExileCertificate.CallUnary(ctx, req)
}

// RevokeExileCertificate calls
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.RevokeExileCertificate.
func (c *exileCertificateManagerServiceClient) RevokeExileCertificate(ctx context.Context, req *connect_go.Request[v1alpha1.RevokeExileCertificateRequest]) (*connect_go.Response[v1alpha1.RevokeExileCertificateResponse], error) {
	return c.revokeExileCertificate.CallUnary(ctx, req)
}

// AssignExileConfiguration calls
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.AssignExileConfiguration.
func (c *exileCertificateManagerServiceClient) AssignExileConfiguration(ctx context.Context, req *connect_go.Request[v1alpha1.AssignExileConfigurationRequest]) (*connect_go.Response[v1alpha1.AssignExileConfigurationResponse], error) {
	return c.assignExileConfiguration.CallUnary(ctx, req)
}

// UnassignExileConfiguration calls
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.UnassignExileConfiguration.
func (c *exileCertificateManagerServiceClient) UnassignExileConfiguration(ctx context.Context, req *connect_go.Request[v1alpha1.UnassignExileConfigurationRequest]) (*connect_go.Response[v1alpha1.UnassignExileConfigurationResponse], error) {
	return c.unassignExileConfiguration.CallUnary(ctx, req)
}

// ListExileCertificates calls
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.ListExileCertificates.
func (c *exileCertificateManagerServiceClient) ListExileCertificates(ctx context.Context, req *connect_go.Request[v1alpha1.ListExileCertificatesRequest]) (*connect_go.Response[v1alpha1.ListExileCertificatesResponse], error) {
	return c.listExileCertificates.CallUnary(ctx, req)
}

// CreateExileConfiguration calls
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.CreateExileConfiguration.
func (c *exileCertificateManagerServiceClient) CreateExileConfiguration(ctx context.Context, req *connect_go.Request[v1alpha1.CreateExileConfigurationRequest]) (*connect_go.Response[v1alpha1.CreateExileConfigurationResponse], error) {
	return c.createExileConfiguration.CallUnary(ctx, req)
}

// UpdateExileConfiguration calls
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.UpdateExileConfiguration.
func (c *exileCertificateManagerServiceClient) UpdateExileConfiguration(ctx context.Context, req *connect_go.Request[v1alpha1.UpdateExileConfigurationRequest]) (*connect_go.Response[v1alpha1.UpdateExileConfigurationResponse], error) {
	return c.updateExileConfiguration.CallUnary(ctx, req)
}

// DeleteExileConfiguration calls
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.DeleteExileConfiguration.
func (c *exileCertificateManagerServiceClient) DeleteExileConfiguration(ctx context.Context, req *connect_go.Request[v1alpha1.DeleteExileConfigurationRequest]) (*connect_go.Response[v1alpha1.DeleteExileConfigurationResponse], error) {
	return c.deleteExileConfiguration.CallUnary(ctx, req)
}

// ListExileConfigurations calls
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.ListExileConfigurations.
func (c *exileCertificateManagerServiceClient) ListExileConfigurations(ctx context.Context, req *connect_go.Request[v1alpha1.ListExileConfigurationsRequest]) (*connect_go.Response[v1alpha1.ListExileConfigurationsResponse], error) {
	return c.listExileConfigurations.CallUnary(ctx, req)
}

// ExileCertificateManagerServiceHandler is an implementation of the
// services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService service.
type ExileCertificateManagerServiceHandler interface {
	// CreateExileCertificate creates a new TLS certificate and
	// returns the exile ceritificate for the current organization.
	CreateExileCertificate(context.Context, *connect_go.Request[v1alpha1.CreateExileCertificateRequest]) (*connect_go.Response[v1alpha1.CreateExileCertificateResponse], error)
	// RevokeExileCertificate deletes a exile ceritificate for the current organization.
	RevokeExileCertificate(context.Context, *connect_go.Request[v1alpha1.RevokeExileCertificateRequest]) (*connect_go.Response[v1alpha1.RevokeExileCertificateResponse], error)
	// AssignExileConfiguration assigns a configuration to a certificate.
	AssignExileConfiguration(context.Context, *connect_go.Request[v1alpha1.AssignExileConfigurationRequest]) (*connect_go.Response[v1alpha1.AssignExileConfigurationResponse], error)
	// UnassignExileConfiguration unassigns a configuration from a certificate.
	UnassignExileConfiguration(context.Context, *connect_go.Request[v1alpha1.UnassignExileConfigurationRequest]) (*connect_go.Response[v1alpha1.UnassignExileConfigurationResponse], error)
	// ListExileCertificates returns a list of exile ceritificate for the current organization.
	ListExileCertificates(context.Context, *connect_go.Request[v1alpha1.ListExileCertificatesRequest]) (*connect_go.Response[v1alpha1.ListExileCertificatesResponse], error)
	// CreateExileConfiguration creates a new exile configuration for the current organization.
	CreateExileConfiguration(context.Context, *connect_go.Request[v1alpha1.CreateExileConfigurationRequest]) (*connect_go.Response[v1alpha1.CreateExileConfigurationResponse], error)
	// UpdateExileConfiguration updates a exile configuration for the current organization.
	UpdateExileConfiguration(context.Context, *connect_go.Request[v1alpha1.UpdateExileConfigurationRequest]) (*connect_go.Response[v1alpha1.UpdateExileConfigurationResponse], error)
	// DeleteExileConfiguration deletes a exile configuration for the current organization.
	// The configuration can only be deleted if it is not assigned to any certificates.
	DeleteExileConfiguration(context.Context, *connect_go.Request[v1alpha1.DeleteExileConfigurationRequest]) (*connect_go.Response[v1alpha1.DeleteExileConfigurationResponse], error)
	// ListExileConfigurations returns a list of exile configurations for the current organization.
	ListExileConfigurations(context.Context, *connect_go.Request[v1alpha1.ListExileConfigurationsRequest]) (*connect_go.Response[v1alpha1.ListExileConfigurationsResponse], error)
}

// NewExileCertificateManagerServiceHandler builds an HTTP handler from the service implementation.
// It returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewExileCertificateManagerServiceHandler(svc ExileCertificateManagerServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	exileCertificateManagerServiceCreateExileCertificateHandler := connect_go.NewUnaryHandler(
		ExileCertificateManagerServiceCreateExileCertificateProcedure,
		svc.CreateExileCertificate,
		opts...,
	)
	exileCertificateManagerServiceRevokeExileCertificateHandler := connect_go.NewUnaryHandler(
		ExileCertificateManagerServiceRevokeExileCertificateProcedure,
		svc.RevokeExileCertificate,
		opts...,
	)
	exileCertificateManagerServiceAssignExileConfigurationHandler := connect_go.NewUnaryHandler(
		ExileCertificateManagerServiceAssignExileConfigurationProcedure,
		svc.AssignExileConfiguration,
		opts...,
	)
	exileCertificateManagerServiceUnassignExileConfigurationHandler := connect_go.NewUnaryHandler(
		ExileCertificateManagerServiceUnassignExileConfigurationProcedure,
		svc.UnassignExileConfiguration,
		opts...,
	)
	exileCertificateManagerServiceListExileCertificatesHandler := connect_go.NewUnaryHandler(
		ExileCertificateManagerServiceListExileCertificatesProcedure,
		svc.ListExileCertificates,
		opts...,
	)
	exileCertificateManagerServiceCreateExileConfigurationHandler := connect_go.NewUnaryHandler(
		ExileCertificateManagerServiceCreateExileConfigurationProcedure,
		svc.CreateExileConfiguration,
		opts...,
	)
	exileCertificateManagerServiceUpdateExileConfigurationHandler := connect_go.NewUnaryHandler(
		ExileCertificateManagerServiceUpdateExileConfigurationProcedure,
		svc.UpdateExileConfiguration,
		opts...,
	)
	exileCertificateManagerServiceDeleteExileConfigurationHandler := connect_go.NewUnaryHandler(
		ExileCertificateManagerServiceDeleteExileConfigurationProcedure,
		svc.DeleteExileConfiguration,
		opts...,
	)
	exileCertificateManagerServiceListExileConfigurationsHandler := connect_go.NewUnaryHandler(
		ExileCertificateManagerServiceListExileConfigurationsProcedure,
		svc.ListExileConfigurations,
		opts...,
	)
	return "/services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ExileCertificateManagerServiceCreateExileCertificateProcedure:
			exileCertificateManagerServiceCreateExileCertificateHandler.ServeHTTP(w, r)
		case ExileCertificateManagerServiceRevokeExileCertificateProcedure:
			exileCertificateManagerServiceRevokeExileCertificateHandler.ServeHTTP(w, r)
		case ExileCertificateManagerServiceAssignExileConfigurationProcedure:
			exileCertificateManagerServiceAssignExileConfigurationHandler.ServeHTTP(w, r)
		case ExileCertificateManagerServiceUnassignExileConfigurationProcedure:
			exileCertificateManagerServiceUnassignExileConfigurationHandler.ServeHTTP(w, r)
		case ExileCertificateManagerServiceListExileCertificatesProcedure:
			exileCertificateManagerServiceListExileCertificatesHandler.ServeHTTP(w, r)
		case ExileCertificateManagerServiceCreateExileConfigurationProcedure:
			exileCertificateManagerServiceCreateExileConfigurationHandler.ServeHTTP(w, r)
		case ExileCertificateManagerServiceUpdateExileConfigurationProcedure:
			exileCertificateManagerServiceUpdateExileConfigurationHandler.ServeHTTP(w, r)
		case ExileCertificateManagerServiceDeleteExileConfigurationProcedure:
			exileCertificateManagerServiceDeleteExileConfigurationHandler.ServeHTTP(w, r)
		case ExileCertificateManagerServiceListExileConfigurationsProcedure:
			exileCertificateManagerServiceListExileConfigurationsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedExileCertificateManagerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedExileCertificateManagerServiceHandler struct{}

func (UnimplementedExileCertificateManagerServiceHandler) CreateExileCertificate(context.Context, *connect_go.Request[v1alpha1.CreateExileCertificateRequest]) (*connect_go.Response[v1alpha1.CreateExileCertificateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.CreateExileCertificate is not implemented"))
}

func (UnimplementedExileCertificateManagerServiceHandler) RevokeExileCertificate(context.Context, *connect_go.Request[v1alpha1.RevokeExileCertificateRequest]) (*connect_go.Response[v1alpha1.RevokeExileCertificateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.RevokeExileCertificate is not implemented"))
}

func (UnimplementedExileCertificateManagerServiceHandler) AssignExileConfiguration(context.Context, *connect_go.Request[v1alpha1.AssignExileConfigurationRequest]) (*connect_go.Response[v1alpha1.AssignExileConfigurationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.AssignExileConfiguration is not implemented"))
}

func (UnimplementedExileCertificateManagerServiceHandler) UnassignExileConfiguration(context.Context, *connect_go.Request[v1alpha1.UnassignExileConfigurationRequest]) (*connect_go.Response[v1alpha1.UnassignExileConfigurationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.UnassignExileConfiguration is not implemented"))
}

func (UnimplementedExileCertificateManagerServiceHandler) ListExileCertificates(context.Context, *connect_go.Request[v1alpha1.ListExileCertificatesRequest]) (*connect_go.Response[v1alpha1.ListExileCertificatesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.ListExileCertificates is not implemented"))
}

func (UnimplementedExileCertificateManagerServiceHandler) CreateExileConfiguration(context.Context, *connect_go.Request[v1alpha1.CreateExileConfigurationRequest]) (*connect_go.Response[v1alpha1.CreateExileConfigurationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.CreateExileConfiguration is not implemented"))
}

func (UnimplementedExileCertificateManagerServiceHandler) UpdateExileConfiguration(context.Context, *connect_go.Request[v1alpha1.UpdateExileConfigurationRequest]) (*connect_go.Response[v1alpha1.UpdateExileConfigurationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.UpdateExileConfiguration is not implemented"))
}

func (UnimplementedExileCertificateManagerServiceHandler) DeleteExileConfiguration(context.Context, *connect_go.Request[v1alpha1.DeleteExileConfigurationRequest]) (*connect_go.Response[v1alpha1.DeleteExileConfigurationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.DeleteExileConfiguration is not implemented"))
}

func (UnimplementedExileCertificateManagerServiceHandler) ListExileConfigurations(context.Context, *connect_go.Request[v1alpha1.ListExileConfigurationsRequest]) (*connect_go.Response[v1alpha1.ListExileConfigurationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.ListExileConfigurations is not implemented"))
}