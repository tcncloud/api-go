// Copyright (c) 2020, TCN Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/workflows/service.proto

package workflowsconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	workflows "github.com/tcncloud/api-go/api/v1alpha1/workflows"
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
	// WorkflowDefinitionPersistServiceName is the fully-qualified name of the
	// WorkflowDefinitionPersistService service.
	WorkflowDefinitionPersistServiceName = "api.v1alpha1.workflows.WorkflowDefinitionPersistService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// WorkflowDefinitionPersistServiceCreateWorkflowDefinitionProcedure is the fully-qualified name of
	// the WorkflowDefinitionPersistService's CreateWorkflowDefinition RPC.
	WorkflowDefinitionPersistServiceCreateWorkflowDefinitionProcedure = "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/CreateWorkflowDefinition"
	// WorkflowDefinitionPersistServiceGetWorkflowDefinitionProcedure is the fully-qualified name of the
	// WorkflowDefinitionPersistService's GetWorkflowDefinition RPC.
	WorkflowDefinitionPersistServiceGetWorkflowDefinitionProcedure = "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/GetWorkflowDefinition"
	// WorkflowDefinitionPersistServiceListWorkflowDefinitionsProcedure is the fully-qualified name of
	// the WorkflowDefinitionPersistService's ListWorkflowDefinitions RPC.
	WorkflowDefinitionPersistServiceListWorkflowDefinitionsProcedure = "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/ListWorkflowDefinitions"
	// WorkflowDefinitionPersistServiceUpdateWorkflowDefinitionProcedure is the fully-qualified name of
	// the WorkflowDefinitionPersistService's UpdateWorkflowDefinition RPC.
	WorkflowDefinitionPersistServiceUpdateWorkflowDefinitionProcedure = "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/UpdateWorkflowDefinition"
	// WorkflowDefinitionPersistServiceValidateWorkflowDefinitionProcedure is the fully-qualified name
	// of the WorkflowDefinitionPersistService's ValidateWorkflowDefinition RPC.
	WorkflowDefinitionPersistServiceValidateWorkflowDefinitionProcedure = "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/ValidateWorkflowDefinition"
)

// WorkflowDefinitionPersistServiceClient is a client for the
// api.v1alpha1.workflows.WorkflowDefinitionPersistService service.
type WorkflowDefinitionPersistServiceClient interface {
	// CreateWorkflowDefinition creates a new flow definition in the database
	CreateWorkflowDefinition(context.Context, *connect_go.Request[workflows.CreateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.CreateWorkflowDefinitionResponse], error)
	// GetWorkflowDefinition retrieves a flow definition from the database
	GetWorkflowDefinition(context.Context, *connect_go.Request[workflows.GetWorkflowDefinitionRequest]) (*connect_go.Response[workflows.GetWorkflowDefinitionResponse], error)
	// ListWorkflowDefinitions retrieves a list of flow definitions from the database optionally filtered by the owning application
	// if application is not specified, all flow definitions for the org are returned
	ListWorkflowDefinitions(context.Context, *connect_go.Request[workflows.ListWorkflowDefinitionsRequest]) (*connect_go.ServerStreamForClient[workflows.ListWorkflowDefinitionsResponse], error)
	// UpdateWorkflowDefinition updates a flow definition in the database. Only the name, description and definition graph itself are updated
	UpdateWorkflowDefinition(context.Context, *connect_go.Request[workflows.UpdateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.UpdateWorkflowDefinitionResponse], error)
	// ValidateWorkflowDefinition validates a flow definition in the database. Only the name, description and definition graph itself are updated
	ValidateWorkflowDefinition(context.Context, *connect_go.Request[workflows.ValidateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.ValidateWorkflowDefinitionResponse], error)
}

// NewWorkflowDefinitionPersistServiceClient constructs a client for the
// api.v1alpha1.workflows.WorkflowDefinitionPersistService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWorkflowDefinitionPersistServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) WorkflowDefinitionPersistServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &workflowDefinitionPersistServiceClient{
		createWorkflowDefinition: connect_go.NewClient[workflows.CreateWorkflowDefinitionRequest, workflows.CreateWorkflowDefinitionResponse](
			httpClient,
			baseURL+WorkflowDefinitionPersistServiceCreateWorkflowDefinitionProcedure,
			opts...,
		),
		getWorkflowDefinition: connect_go.NewClient[workflows.GetWorkflowDefinitionRequest, workflows.GetWorkflowDefinitionResponse](
			httpClient,
			baseURL+WorkflowDefinitionPersistServiceGetWorkflowDefinitionProcedure,
			opts...,
		),
		listWorkflowDefinitions: connect_go.NewClient[workflows.ListWorkflowDefinitionsRequest, workflows.ListWorkflowDefinitionsResponse](
			httpClient,
			baseURL+WorkflowDefinitionPersistServiceListWorkflowDefinitionsProcedure,
			opts...,
		),
		updateWorkflowDefinition: connect_go.NewClient[workflows.UpdateWorkflowDefinitionRequest, workflows.UpdateWorkflowDefinitionResponse](
			httpClient,
			baseURL+WorkflowDefinitionPersistServiceUpdateWorkflowDefinitionProcedure,
			opts...,
		),
		validateWorkflowDefinition: connect_go.NewClient[workflows.ValidateWorkflowDefinitionRequest, workflows.ValidateWorkflowDefinitionResponse](
			httpClient,
			baseURL+WorkflowDefinitionPersistServiceValidateWorkflowDefinitionProcedure,
			opts...,
		),
	}
}

// workflowDefinitionPersistServiceClient implements WorkflowDefinitionPersistServiceClient.
type workflowDefinitionPersistServiceClient struct {
	createWorkflowDefinition   *connect_go.Client[workflows.CreateWorkflowDefinitionRequest, workflows.CreateWorkflowDefinitionResponse]
	getWorkflowDefinition      *connect_go.Client[workflows.GetWorkflowDefinitionRequest, workflows.GetWorkflowDefinitionResponse]
	listWorkflowDefinitions    *connect_go.Client[workflows.ListWorkflowDefinitionsRequest, workflows.ListWorkflowDefinitionsResponse]
	updateWorkflowDefinition   *connect_go.Client[workflows.UpdateWorkflowDefinitionRequest, workflows.UpdateWorkflowDefinitionResponse]
	validateWorkflowDefinition *connect_go.Client[workflows.ValidateWorkflowDefinitionRequest, workflows.ValidateWorkflowDefinitionResponse]
}

// CreateWorkflowDefinition calls
// api.v1alpha1.workflows.WorkflowDefinitionPersistService.CreateWorkflowDefinition.
func (c *workflowDefinitionPersistServiceClient) CreateWorkflowDefinition(ctx context.Context, req *connect_go.Request[workflows.CreateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.CreateWorkflowDefinitionResponse], error) {
	return c.createWorkflowDefinition.CallUnary(ctx, req)
}

// GetWorkflowDefinition calls
// api.v1alpha1.workflows.WorkflowDefinitionPersistService.GetWorkflowDefinition.
func (c *workflowDefinitionPersistServiceClient) GetWorkflowDefinition(ctx context.Context, req *connect_go.Request[workflows.GetWorkflowDefinitionRequest]) (*connect_go.Response[workflows.GetWorkflowDefinitionResponse], error) {
	return c.getWorkflowDefinition.CallUnary(ctx, req)
}

// ListWorkflowDefinitions calls
// api.v1alpha1.workflows.WorkflowDefinitionPersistService.ListWorkflowDefinitions.
func (c *workflowDefinitionPersistServiceClient) ListWorkflowDefinitions(ctx context.Context, req *connect_go.Request[workflows.ListWorkflowDefinitionsRequest]) (*connect_go.ServerStreamForClient[workflows.ListWorkflowDefinitionsResponse], error) {
	return c.listWorkflowDefinitions.CallServerStream(ctx, req)
}

// UpdateWorkflowDefinition calls
// api.v1alpha1.workflows.WorkflowDefinitionPersistService.UpdateWorkflowDefinition.
func (c *workflowDefinitionPersistServiceClient) UpdateWorkflowDefinition(ctx context.Context, req *connect_go.Request[workflows.UpdateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.UpdateWorkflowDefinitionResponse], error) {
	return c.updateWorkflowDefinition.CallUnary(ctx, req)
}

// ValidateWorkflowDefinition calls
// api.v1alpha1.workflows.WorkflowDefinitionPersistService.ValidateWorkflowDefinition.
func (c *workflowDefinitionPersistServiceClient) ValidateWorkflowDefinition(ctx context.Context, req *connect_go.Request[workflows.ValidateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.ValidateWorkflowDefinitionResponse], error) {
	return c.validateWorkflowDefinition.CallUnary(ctx, req)
}

// WorkflowDefinitionPersistServiceHandler is an implementation of the
// api.v1alpha1.workflows.WorkflowDefinitionPersistService service.
type WorkflowDefinitionPersistServiceHandler interface {
	// CreateWorkflowDefinition creates a new flow definition in the database
	CreateWorkflowDefinition(context.Context, *connect_go.Request[workflows.CreateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.CreateWorkflowDefinitionResponse], error)
	// GetWorkflowDefinition retrieves a flow definition from the database
	GetWorkflowDefinition(context.Context, *connect_go.Request[workflows.GetWorkflowDefinitionRequest]) (*connect_go.Response[workflows.GetWorkflowDefinitionResponse], error)
	// ListWorkflowDefinitions retrieves a list of flow definitions from the database optionally filtered by the owning application
	// if application is not specified, all flow definitions for the org are returned
	ListWorkflowDefinitions(context.Context, *connect_go.Request[workflows.ListWorkflowDefinitionsRequest], *connect_go.ServerStream[workflows.ListWorkflowDefinitionsResponse]) error
	// UpdateWorkflowDefinition updates a flow definition in the database. Only the name, description and definition graph itself are updated
	UpdateWorkflowDefinition(context.Context, *connect_go.Request[workflows.UpdateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.UpdateWorkflowDefinitionResponse], error)
	// ValidateWorkflowDefinition validates a flow definition in the database. Only the name, description and definition graph itself are updated
	ValidateWorkflowDefinition(context.Context, *connect_go.Request[workflows.ValidateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.ValidateWorkflowDefinitionResponse], error)
}

// NewWorkflowDefinitionPersistServiceHandler builds an HTTP handler from the service
// implementation. It returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkflowDefinitionPersistServiceHandler(svc WorkflowDefinitionPersistServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	workflowDefinitionPersistServiceCreateWorkflowDefinitionHandler := connect_go.NewUnaryHandler(
		WorkflowDefinitionPersistServiceCreateWorkflowDefinitionProcedure,
		svc.CreateWorkflowDefinition,
		opts...,
	)
	workflowDefinitionPersistServiceGetWorkflowDefinitionHandler := connect_go.NewUnaryHandler(
		WorkflowDefinitionPersistServiceGetWorkflowDefinitionProcedure,
		svc.GetWorkflowDefinition,
		opts...,
	)
	workflowDefinitionPersistServiceListWorkflowDefinitionsHandler := connect_go.NewServerStreamHandler(
		WorkflowDefinitionPersistServiceListWorkflowDefinitionsProcedure,
		svc.ListWorkflowDefinitions,
		opts...,
	)
	workflowDefinitionPersistServiceUpdateWorkflowDefinitionHandler := connect_go.NewUnaryHandler(
		WorkflowDefinitionPersistServiceUpdateWorkflowDefinitionProcedure,
		svc.UpdateWorkflowDefinition,
		opts...,
	)
	workflowDefinitionPersistServiceValidateWorkflowDefinitionHandler := connect_go.NewUnaryHandler(
		WorkflowDefinitionPersistServiceValidateWorkflowDefinitionProcedure,
		svc.ValidateWorkflowDefinition,
		opts...,
	)
	return "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WorkflowDefinitionPersistServiceCreateWorkflowDefinitionProcedure:
			workflowDefinitionPersistServiceCreateWorkflowDefinitionHandler.ServeHTTP(w, r)
		case WorkflowDefinitionPersistServiceGetWorkflowDefinitionProcedure:
			workflowDefinitionPersistServiceGetWorkflowDefinitionHandler.ServeHTTP(w, r)
		case WorkflowDefinitionPersistServiceListWorkflowDefinitionsProcedure:
			workflowDefinitionPersistServiceListWorkflowDefinitionsHandler.ServeHTTP(w, r)
		case WorkflowDefinitionPersistServiceUpdateWorkflowDefinitionProcedure:
			workflowDefinitionPersistServiceUpdateWorkflowDefinitionHandler.ServeHTTP(w, r)
		case WorkflowDefinitionPersistServiceValidateWorkflowDefinitionProcedure:
			workflowDefinitionPersistServiceValidateWorkflowDefinitionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWorkflowDefinitionPersistServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkflowDefinitionPersistServiceHandler struct{}

func (UnimplementedWorkflowDefinitionPersistServiceHandler) CreateWorkflowDefinition(context.Context, *connect_go.Request[workflows.CreateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.CreateWorkflowDefinitionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.workflows.WorkflowDefinitionPersistService.CreateWorkflowDefinition is not implemented"))
}

func (UnimplementedWorkflowDefinitionPersistServiceHandler) GetWorkflowDefinition(context.Context, *connect_go.Request[workflows.GetWorkflowDefinitionRequest]) (*connect_go.Response[workflows.GetWorkflowDefinitionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.workflows.WorkflowDefinitionPersistService.GetWorkflowDefinition is not implemented"))
}

func (UnimplementedWorkflowDefinitionPersistServiceHandler) ListWorkflowDefinitions(context.Context, *connect_go.Request[workflows.ListWorkflowDefinitionsRequest], *connect_go.ServerStream[workflows.ListWorkflowDefinitionsResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.workflows.WorkflowDefinitionPersistService.ListWorkflowDefinitions is not implemented"))
}

func (UnimplementedWorkflowDefinitionPersistServiceHandler) UpdateWorkflowDefinition(context.Context, *connect_go.Request[workflows.UpdateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.UpdateWorkflowDefinitionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.workflows.WorkflowDefinitionPersistService.UpdateWorkflowDefinition is not implemented"))
}

func (UnimplementedWorkflowDefinitionPersistServiceHandler) ValidateWorkflowDefinition(context.Context, *connect_go.Request[workflows.ValidateWorkflowDefinitionRequest]) (*connect_go.Response[workflows.ValidateWorkflowDefinitionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.workflows.WorkflowDefinitionPersistService.ValidateWorkflowDefinition is not implemented"))
}
