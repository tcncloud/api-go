// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: api/v1alpha1/agenttraining/service.proto

package agenttraining

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
	AgentTrainingService_CreateLearningOpportunity_FullMethodName        = "/api.v1alpha1.agenttraining.AgentTrainingService/CreateLearningOpportunity"
	AgentTrainingService_ListLearningOpportunities_FullMethodName        = "/api.v1alpha1.agenttraining.AgentTrainingService/ListLearningOpportunities"
	AgentTrainingService_ListAgentLearningOpportunities_FullMethodName   = "/api.v1alpha1.agenttraining.AgentTrainingService/ListAgentLearningOpportunities"
	AgentTrainingService_CompleteAgentLearningOpportunity_FullMethodName = "/api.v1alpha1.agenttraining.AgentTrainingService/CompleteAgentLearningOpportunity"
	AgentTrainingService_UpdateLearningOpportunity_FullMethodName        = "/api.v1alpha1.agenttraining.AgentTrainingService/UpdateLearningOpportunity"
	AgentTrainingService_DeleteLearningOpportunity_FullMethodName        = "/api.v1alpha1.agenttraining.AgentTrainingService/DeleteLearningOpportunity"
	AgentTrainingService_GetLearningOpportunity_FullMethodName           = "/api.v1alpha1.agenttraining.AgentTrainingService/GetLearningOpportunity"
)

// AgentTrainingServiceClient is the client API for AgentTrainingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentTrainingServiceClient interface {
	// CreateLearningOpportunity creates a new learning opportunity.
	CreateLearningOpportunity(ctx context.Context, in *CreateLearningOpportunityRequest, opts ...grpc.CallOption) (*CreateLearningOpportunityResponse, error)
	// ListLearningOpportunities lists learning opportunities.
	ListLearningOpportunities(ctx context.Context, in *ListLearningOpportunitiesRequest, opts ...grpc.CallOption) (*ListLearningOpportunitiesResponse, error)
	// ListAgentLearningOpportunities lists learning opportunities by agent.
	ListAgentLearningOpportunities(ctx context.Context, in *ListAgentLearningOpportunitiesRequest, opts ...grpc.CallOption) (*ListAgentLearningOpportunitiesResponse, error)
	// CompleteAgentLearningOpportunity completes an agent's learning opportunity.
	CompleteAgentLearningOpportunity(ctx context.Context, in *CompleteAgentLearningOpportunityRequest, opts ...grpc.CallOption) (*CompleteAgentLearningOpportunityResponse, error)
	// UpdateLearningOpportunity updates a learning opportunity.
	UpdateLearningOpportunity(ctx context.Context, in *UpdateLearningOpportunityRequest, opts ...grpc.CallOption) (*UpdateLearningOpportunityResponse, error)
	// DeleteLearningOpportunity deletes a learning opportunity.
	DeleteLearningOpportunity(ctx context.Context, in *DeleteLearningOpportunityRequest, opts ...grpc.CallOption) (*DeleteLearningOpportunityResponse, error)
	// GetLearningOpportunity gets a learning opportunity.
	GetLearningOpportunity(ctx context.Context, in *GetLearningOpportunityRequest, opts ...grpc.CallOption) (*GetLearningOpportunityResponse, error)
}

type agentTrainingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentTrainingServiceClient(cc grpc.ClientConnInterface) AgentTrainingServiceClient {
	return &agentTrainingServiceClient{cc}
}

func (c *agentTrainingServiceClient) CreateLearningOpportunity(ctx context.Context, in *CreateLearningOpportunityRequest, opts ...grpc.CallOption) (*CreateLearningOpportunityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLearningOpportunityResponse)
	err := c.cc.Invoke(ctx, AgentTrainingService_CreateLearningOpportunity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTrainingServiceClient) ListLearningOpportunities(ctx context.Context, in *ListLearningOpportunitiesRequest, opts ...grpc.CallOption) (*ListLearningOpportunitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLearningOpportunitiesResponse)
	err := c.cc.Invoke(ctx, AgentTrainingService_ListLearningOpportunities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTrainingServiceClient) ListAgentLearningOpportunities(ctx context.Context, in *ListAgentLearningOpportunitiesRequest, opts ...grpc.CallOption) (*ListAgentLearningOpportunitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAgentLearningOpportunitiesResponse)
	err := c.cc.Invoke(ctx, AgentTrainingService_ListAgentLearningOpportunities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTrainingServiceClient) CompleteAgentLearningOpportunity(ctx context.Context, in *CompleteAgentLearningOpportunityRequest, opts ...grpc.CallOption) (*CompleteAgentLearningOpportunityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompleteAgentLearningOpportunityResponse)
	err := c.cc.Invoke(ctx, AgentTrainingService_CompleteAgentLearningOpportunity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTrainingServiceClient) UpdateLearningOpportunity(ctx context.Context, in *UpdateLearningOpportunityRequest, opts ...grpc.CallOption) (*UpdateLearningOpportunityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateLearningOpportunityResponse)
	err := c.cc.Invoke(ctx, AgentTrainingService_UpdateLearningOpportunity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTrainingServiceClient) DeleteLearningOpportunity(ctx context.Context, in *DeleteLearningOpportunityRequest, opts ...grpc.CallOption) (*DeleteLearningOpportunityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteLearningOpportunityResponse)
	err := c.cc.Invoke(ctx, AgentTrainingService_DeleteLearningOpportunity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTrainingServiceClient) GetLearningOpportunity(ctx context.Context, in *GetLearningOpportunityRequest, opts ...grpc.CallOption) (*GetLearningOpportunityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLearningOpportunityResponse)
	err := c.cc.Invoke(ctx, AgentTrainingService_GetLearningOpportunity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentTrainingServiceServer is the server API for AgentTrainingService service.
// All implementations must embed UnimplementedAgentTrainingServiceServer
// for forward compatibility
type AgentTrainingServiceServer interface {
	// CreateLearningOpportunity creates a new learning opportunity.
	CreateLearningOpportunity(context.Context, *CreateLearningOpportunityRequest) (*CreateLearningOpportunityResponse, error)
	// ListLearningOpportunities lists learning opportunities.
	ListLearningOpportunities(context.Context, *ListLearningOpportunitiesRequest) (*ListLearningOpportunitiesResponse, error)
	// ListAgentLearningOpportunities lists learning opportunities by agent.
	ListAgentLearningOpportunities(context.Context, *ListAgentLearningOpportunitiesRequest) (*ListAgentLearningOpportunitiesResponse, error)
	// CompleteAgentLearningOpportunity completes an agent's learning opportunity.
	CompleteAgentLearningOpportunity(context.Context, *CompleteAgentLearningOpportunityRequest) (*CompleteAgentLearningOpportunityResponse, error)
	// UpdateLearningOpportunity updates a learning opportunity.
	UpdateLearningOpportunity(context.Context, *UpdateLearningOpportunityRequest) (*UpdateLearningOpportunityResponse, error)
	// DeleteLearningOpportunity deletes a learning opportunity.
	DeleteLearningOpportunity(context.Context, *DeleteLearningOpportunityRequest) (*DeleteLearningOpportunityResponse, error)
	// GetLearningOpportunity gets a learning opportunity.
	GetLearningOpportunity(context.Context, *GetLearningOpportunityRequest) (*GetLearningOpportunityResponse, error)
	mustEmbedUnimplementedAgentTrainingServiceServer()
}

// UnimplementedAgentTrainingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentTrainingServiceServer struct {
}

func (UnimplementedAgentTrainingServiceServer) CreateLearningOpportunity(context.Context, *CreateLearningOpportunityRequest) (*CreateLearningOpportunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLearningOpportunity not implemented")
}
func (UnimplementedAgentTrainingServiceServer) ListLearningOpportunities(context.Context, *ListLearningOpportunitiesRequest) (*ListLearningOpportunitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLearningOpportunities not implemented")
}
func (UnimplementedAgentTrainingServiceServer) ListAgentLearningOpportunities(context.Context, *ListAgentLearningOpportunitiesRequest) (*ListAgentLearningOpportunitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgentLearningOpportunities not implemented")
}
func (UnimplementedAgentTrainingServiceServer) CompleteAgentLearningOpportunity(context.Context, *CompleteAgentLearningOpportunityRequest) (*CompleteAgentLearningOpportunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteAgentLearningOpportunity not implemented")
}
func (UnimplementedAgentTrainingServiceServer) UpdateLearningOpportunity(context.Context, *UpdateLearningOpportunityRequest) (*UpdateLearningOpportunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLearningOpportunity not implemented")
}
func (UnimplementedAgentTrainingServiceServer) DeleteLearningOpportunity(context.Context, *DeleteLearningOpportunityRequest) (*DeleteLearningOpportunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLearningOpportunity not implemented")
}
func (UnimplementedAgentTrainingServiceServer) GetLearningOpportunity(context.Context, *GetLearningOpportunityRequest) (*GetLearningOpportunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLearningOpportunity not implemented")
}
func (UnimplementedAgentTrainingServiceServer) mustEmbedUnimplementedAgentTrainingServiceServer() {}

// UnsafeAgentTrainingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentTrainingServiceServer will
// result in compilation errors.
type UnsafeAgentTrainingServiceServer interface {
	mustEmbedUnimplementedAgentTrainingServiceServer()
}

func RegisterAgentTrainingServiceServer(s grpc.ServiceRegistrar, srv AgentTrainingServiceServer) {
	s.RegisterService(&AgentTrainingService_ServiceDesc, srv)
}

func _AgentTrainingService_CreateLearningOpportunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLearningOpportunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTrainingServiceServer).CreateLearningOpportunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentTrainingService_CreateLearningOpportunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTrainingServiceServer).CreateLearningOpportunity(ctx, req.(*CreateLearningOpportunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTrainingService_ListLearningOpportunities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLearningOpportunitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTrainingServiceServer).ListLearningOpportunities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentTrainingService_ListLearningOpportunities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTrainingServiceServer).ListLearningOpportunities(ctx, req.(*ListLearningOpportunitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTrainingService_ListAgentLearningOpportunities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentLearningOpportunitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTrainingServiceServer).ListAgentLearningOpportunities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentTrainingService_ListAgentLearningOpportunities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTrainingServiceServer).ListAgentLearningOpportunities(ctx, req.(*ListAgentLearningOpportunitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTrainingService_CompleteAgentLearningOpportunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteAgentLearningOpportunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTrainingServiceServer).CompleteAgentLearningOpportunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentTrainingService_CompleteAgentLearningOpportunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTrainingServiceServer).CompleteAgentLearningOpportunity(ctx, req.(*CompleteAgentLearningOpportunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTrainingService_UpdateLearningOpportunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLearningOpportunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTrainingServiceServer).UpdateLearningOpportunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentTrainingService_UpdateLearningOpportunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTrainingServiceServer).UpdateLearningOpportunity(ctx, req.(*UpdateLearningOpportunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTrainingService_DeleteLearningOpportunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLearningOpportunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTrainingServiceServer).DeleteLearningOpportunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentTrainingService_DeleteLearningOpportunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTrainingServiceServer).DeleteLearningOpportunity(ctx, req.(*DeleteLearningOpportunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTrainingService_GetLearningOpportunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLearningOpportunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTrainingServiceServer).GetLearningOpportunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentTrainingService_GetLearningOpportunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTrainingServiceServer).GetLearningOpportunity(ctx, req.(*GetLearningOpportunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentTrainingService_ServiceDesc is the grpc.ServiceDesc for AgentTrainingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentTrainingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.agenttraining.AgentTrainingService",
	HandlerType: (*AgentTrainingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLearningOpportunity",
			Handler:    _AgentTrainingService_CreateLearningOpportunity_Handler,
		},
		{
			MethodName: "ListLearningOpportunities",
			Handler:    _AgentTrainingService_ListLearningOpportunities_Handler,
		},
		{
			MethodName: "ListAgentLearningOpportunities",
			Handler:    _AgentTrainingService_ListAgentLearningOpportunities_Handler,
		},
		{
			MethodName: "CompleteAgentLearningOpportunity",
			Handler:    _AgentTrainingService_CompleteAgentLearningOpportunity_Handler,
		},
		{
			MethodName: "UpdateLearningOpportunity",
			Handler:    _AgentTrainingService_UpdateLearningOpportunity_Handler,
		},
		{
			MethodName: "DeleteLearningOpportunity",
			Handler:    _AgentTrainingService_DeleteLearningOpportunity_Handler,
		},
		{
			MethodName: "GetLearningOpportunity",
			Handler:    _AgentTrainingService_GetLearningOpportunity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1alpha1/agenttraining/service.proto",
}
