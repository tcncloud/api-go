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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: api/v1alpha1/workflows/service.proto

package workflows

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	WorkflowDefinitionPersistService_CreateWorkflowDefinition_FullMethodName   = "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/CreateWorkflowDefinition"
	WorkflowDefinitionPersistService_GetWorkflowDefinition_FullMethodName      = "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/GetWorkflowDefinition"
	WorkflowDefinitionPersistService_ListWorkflowDefinitions_FullMethodName    = "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/ListWorkflowDefinitions"
	WorkflowDefinitionPersistService_UpdateWorkflowDefinition_FullMethodName   = "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/UpdateWorkflowDefinition"
	WorkflowDefinitionPersistService_DeleteWorkflowDefinition_FullMethodName   = "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/DeleteWorkflowDefinition"
	WorkflowDefinitionPersistService_ValidateWorkflowDefinition_FullMethodName = "/api.v1alpha1.workflows.WorkflowDefinitionPersistService/ValidateWorkflowDefinition"
)

// WorkflowDefinitionPersistServiceClient is the client API for WorkflowDefinitionPersistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// WorkflowDefinitionPersistService is the service that provides CRUD operations for workflow definitions.
// PERMISSION_WORKFLOWS is required for all operations
type WorkflowDefinitionPersistServiceClient interface {
	// CreateWorkflowDefinition creates a new flow definition in the database
	CreateWorkflowDefinition(ctx context.Context, in *CreateWorkflowDefinitionRequest, opts ...grpc.CallOption) (*CreateWorkflowDefinitionResponse, error)
	// GetWorkflowDefinition retrieves a flow definition from the database
	GetWorkflowDefinition(ctx context.Context, in *GetWorkflowDefinitionRequest, opts ...grpc.CallOption) (*GetWorkflowDefinitionResponse, error)
	// ListWorkflowDefinitions retrieves a list of flow definitions from the database optionally filtered by the owning application
	// if application is not specified, all flow definitions for the org are returned
	ListWorkflowDefinitions(ctx context.Context, in *ListWorkflowDefinitionsRequest, opts ...grpc.CallOption) (WorkflowDefinitionPersistService_ListWorkflowDefinitionsClient, error)
	// UpdateWorkflowDefinition updates a flow definition in the database. Only the name, description and definition graph itself are updated
	UpdateWorkflowDefinition(ctx context.Context, in *UpdateWorkflowDefinitionRequest, opts ...grpc.CallOption) (*UpdateWorkflowDefinitionResponse, error)
	// DeleteWorkflowDefinition deletes a flow definition from the database
	DeleteWorkflowDefinition(ctx context.Context, in *DeleteWorkflowDefinitionRequest, opts ...grpc.CallOption) (*DeleteWorkflowDefinitionResponse, error)
	// ValidateWorkflowDefinition validates a flow definition in the database. Only the name, description and definition graph itself are updated
	ValidateWorkflowDefinition(ctx context.Context, in *ValidateWorkflowDefinitionRequest, opts ...grpc.CallOption) (*ValidateWorkflowDefinitionResponse, error)
}

type workflowDefinitionPersistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowDefinitionPersistServiceClient(cc grpc.ClientConnInterface) WorkflowDefinitionPersistServiceClient {
	return &workflowDefinitionPersistServiceClient{cc}
}

func (c *workflowDefinitionPersistServiceClient) CreateWorkflowDefinition(ctx context.Context, in *CreateWorkflowDefinitionRequest, opts ...grpc.CallOption) (*CreateWorkflowDefinitionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateWorkflowDefinitionResponse)
	err := c.cc.Invoke(ctx, WorkflowDefinitionPersistService_CreateWorkflowDefinition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowDefinitionPersistServiceClient) GetWorkflowDefinition(ctx context.Context, in *GetWorkflowDefinitionRequest, opts ...grpc.CallOption) (*GetWorkflowDefinitionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWorkflowDefinitionResponse)
	err := c.cc.Invoke(ctx, WorkflowDefinitionPersistService_GetWorkflowDefinition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowDefinitionPersistServiceClient) ListWorkflowDefinitions(ctx context.Context, in *ListWorkflowDefinitionsRequest, opts ...grpc.CallOption) (WorkflowDefinitionPersistService_ListWorkflowDefinitionsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WorkflowDefinitionPersistService_ServiceDesc.Streams[0], WorkflowDefinitionPersistService_ListWorkflowDefinitions_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &workflowDefinitionPersistServiceListWorkflowDefinitionsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WorkflowDefinitionPersistService_ListWorkflowDefinitionsClient interface {
	Recv() (*ListWorkflowDefinitionsResponse, error)
	grpc.ClientStream
}

type workflowDefinitionPersistServiceListWorkflowDefinitionsClient struct {
	grpc.ClientStream
}

func (x *workflowDefinitionPersistServiceListWorkflowDefinitionsClient) Recv() (*ListWorkflowDefinitionsResponse, error) {
	m := new(ListWorkflowDefinitionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workflowDefinitionPersistServiceClient) UpdateWorkflowDefinition(ctx context.Context, in *UpdateWorkflowDefinitionRequest, opts ...grpc.CallOption) (*UpdateWorkflowDefinitionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateWorkflowDefinitionResponse)
	err := c.cc.Invoke(ctx, WorkflowDefinitionPersistService_UpdateWorkflowDefinition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowDefinitionPersistServiceClient) DeleteWorkflowDefinition(ctx context.Context, in *DeleteWorkflowDefinitionRequest, opts ...grpc.CallOption) (*DeleteWorkflowDefinitionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteWorkflowDefinitionResponse)
	err := c.cc.Invoke(ctx, WorkflowDefinitionPersistService_DeleteWorkflowDefinition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowDefinitionPersistServiceClient) ValidateWorkflowDefinition(ctx context.Context, in *ValidateWorkflowDefinitionRequest, opts ...grpc.CallOption) (*ValidateWorkflowDefinitionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateWorkflowDefinitionResponse)
	err := c.cc.Invoke(ctx, WorkflowDefinitionPersistService_ValidateWorkflowDefinition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowDefinitionPersistServiceServer is the server API for WorkflowDefinitionPersistService service.
// All implementations must embed UnimplementedWorkflowDefinitionPersistServiceServer
// for forward compatibility
//
// WorkflowDefinitionPersistService is the service that provides CRUD operations for workflow definitions.
// PERMISSION_WORKFLOWS is required for all operations
type WorkflowDefinitionPersistServiceServer interface {
	// CreateWorkflowDefinition creates a new flow definition in the database
	CreateWorkflowDefinition(context.Context, *CreateWorkflowDefinitionRequest) (*CreateWorkflowDefinitionResponse, error)
	// GetWorkflowDefinition retrieves a flow definition from the database
	GetWorkflowDefinition(context.Context, *GetWorkflowDefinitionRequest) (*GetWorkflowDefinitionResponse, error)
	// ListWorkflowDefinitions retrieves a list of flow definitions from the database optionally filtered by the owning application
	// if application is not specified, all flow definitions for the org are returned
	ListWorkflowDefinitions(*ListWorkflowDefinitionsRequest, WorkflowDefinitionPersistService_ListWorkflowDefinitionsServer) error
	// UpdateWorkflowDefinition updates a flow definition in the database. Only the name, description and definition graph itself are updated
	UpdateWorkflowDefinition(context.Context, *UpdateWorkflowDefinitionRequest) (*UpdateWorkflowDefinitionResponse, error)
	// DeleteWorkflowDefinition deletes a flow definition from the database
	DeleteWorkflowDefinition(context.Context, *DeleteWorkflowDefinitionRequest) (*DeleteWorkflowDefinitionResponse, error)
	// ValidateWorkflowDefinition validates a flow definition in the database. Only the name, description and definition graph itself are updated
	ValidateWorkflowDefinition(context.Context, *ValidateWorkflowDefinitionRequest) (*ValidateWorkflowDefinitionResponse, error)
	mustEmbedUnimplementedWorkflowDefinitionPersistServiceServer()
}

// UnimplementedWorkflowDefinitionPersistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkflowDefinitionPersistServiceServer struct {
}

func (UnimplementedWorkflowDefinitionPersistServiceServer) CreateWorkflowDefinition(context.Context, *CreateWorkflowDefinitionRequest) (*CreateWorkflowDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflowDefinition not implemented")
}
func (UnimplementedWorkflowDefinitionPersistServiceServer) GetWorkflowDefinition(context.Context, *GetWorkflowDefinitionRequest) (*GetWorkflowDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowDefinition not implemented")
}
func (UnimplementedWorkflowDefinitionPersistServiceServer) ListWorkflowDefinitions(*ListWorkflowDefinitionsRequest, WorkflowDefinitionPersistService_ListWorkflowDefinitionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListWorkflowDefinitions not implemented")
}
func (UnimplementedWorkflowDefinitionPersistServiceServer) UpdateWorkflowDefinition(context.Context, *UpdateWorkflowDefinitionRequest) (*UpdateWorkflowDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkflowDefinition not implemented")
}
func (UnimplementedWorkflowDefinitionPersistServiceServer) DeleteWorkflowDefinition(context.Context, *DeleteWorkflowDefinitionRequest) (*DeleteWorkflowDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkflowDefinition not implemented")
}
func (UnimplementedWorkflowDefinitionPersistServiceServer) ValidateWorkflowDefinition(context.Context, *ValidateWorkflowDefinitionRequest) (*ValidateWorkflowDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateWorkflowDefinition not implemented")
}
func (UnimplementedWorkflowDefinitionPersistServiceServer) mustEmbedUnimplementedWorkflowDefinitionPersistServiceServer() {
}

// UnsafeWorkflowDefinitionPersistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowDefinitionPersistServiceServer will
// result in compilation errors.
type UnsafeWorkflowDefinitionPersistServiceServer interface {
	mustEmbedUnimplementedWorkflowDefinitionPersistServiceServer()
}

func RegisterWorkflowDefinitionPersistServiceServer(s grpc.ServiceRegistrar, srv WorkflowDefinitionPersistServiceServer) {
	s.RegisterService(&WorkflowDefinitionPersistService_ServiceDesc, srv)
}

func _WorkflowDefinitionPersistService_CreateWorkflowDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowDefinitionPersistServiceServer).CreateWorkflowDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowDefinitionPersistService_CreateWorkflowDefinition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowDefinitionPersistServiceServer).CreateWorkflowDefinition(ctx, req.(*CreateWorkflowDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowDefinitionPersistService_GetWorkflowDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowDefinitionPersistServiceServer).GetWorkflowDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowDefinitionPersistService_GetWorkflowDefinition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowDefinitionPersistServiceServer).GetWorkflowDefinition(ctx, req.(*GetWorkflowDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowDefinitionPersistService_ListWorkflowDefinitions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListWorkflowDefinitionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkflowDefinitionPersistServiceServer).ListWorkflowDefinitions(m, &workflowDefinitionPersistServiceListWorkflowDefinitionsServer{ServerStream: stream})
}

type WorkflowDefinitionPersistService_ListWorkflowDefinitionsServer interface {
	Send(*ListWorkflowDefinitionsResponse) error
	grpc.ServerStream
}

type workflowDefinitionPersistServiceListWorkflowDefinitionsServer struct {
	grpc.ServerStream
}

func (x *workflowDefinitionPersistServiceListWorkflowDefinitionsServer) Send(m *ListWorkflowDefinitionsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WorkflowDefinitionPersistService_UpdateWorkflowDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkflowDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowDefinitionPersistServiceServer).UpdateWorkflowDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowDefinitionPersistService_UpdateWorkflowDefinition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowDefinitionPersistServiceServer).UpdateWorkflowDefinition(ctx, req.(*UpdateWorkflowDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowDefinitionPersistService_DeleteWorkflowDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkflowDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowDefinitionPersistServiceServer).DeleteWorkflowDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowDefinitionPersistService_DeleteWorkflowDefinition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowDefinitionPersistServiceServer).DeleteWorkflowDefinition(ctx, req.(*DeleteWorkflowDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowDefinitionPersistService_ValidateWorkflowDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateWorkflowDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowDefinitionPersistServiceServer).ValidateWorkflowDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowDefinitionPersistService_ValidateWorkflowDefinition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowDefinitionPersistServiceServer).ValidateWorkflowDefinition(ctx, req.(*ValidateWorkflowDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkflowDefinitionPersistService_ServiceDesc is the grpc.ServiceDesc for WorkflowDefinitionPersistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkflowDefinitionPersistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.workflows.WorkflowDefinitionPersistService",
	HandlerType: (*WorkflowDefinitionPersistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkflowDefinition",
			Handler:    _WorkflowDefinitionPersistService_CreateWorkflowDefinition_Handler,
		},
		{
			MethodName: "GetWorkflowDefinition",
			Handler:    _WorkflowDefinitionPersistService_GetWorkflowDefinition_Handler,
		},
		{
			MethodName: "UpdateWorkflowDefinition",
			Handler:    _WorkflowDefinitionPersistService_UpdateWorkflowDefinition_Handler,
		},
		{
			MethodName: "DeleteWorkflowDefinition",
			Handler:    _WorkflowDefinitionPersistService_DeleteWorkflowDefinition_Handler,
		},
		{
			MethodName: "ValidateWorkflowDefinition",
			Handler:    _WorkflowDefinitionPersistService_ValidateWorkflowDefinition_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListWorkflowDefinitions",
			Handler:       _WorkflowDefinitionPersistService_ListWorkflowDefinitions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1alpha1/workflows/service.proto",
}
