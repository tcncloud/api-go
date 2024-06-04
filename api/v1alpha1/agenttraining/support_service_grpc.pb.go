// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: api/v1alpha1/agenttraining/support_service.proto

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
	AgentTrainingSupportService_ListLearningOpportunitiesByOrgId_FullMethodName = "/api.v1alpha1.agenttraining.AgentTrainingSupportService/ListLearningOpportunitiesByOrgId"
	AgentTrainingSupportService_DeleteLearningOpportunityByOrgId_FullMethodName = "/api.v1alpha1.agenttraining.AgentTrainingSupportService/DeleteLearningOpportunityByOrgId"
)

// AgentTrainingSupportServiceClient is the client API for AgentTrainingSupportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AgentTrainingSupportService provides customer support
// endpoints for agent training.
type AgentTrainingSupportServiceClient interface {
	// ListLearningOpportunitiesByOrgId lists learning opportunities for a specific org.
	ListLearningOpportunitiesByOrgId(ctx context.Context, in *ListLearningOpportunitiesByOrgIdRequest, opts ...grpc.CallOption) (*ListLearningOpportunitiesResponse, error)
	// DeleteLearningOpportunityByOrgId deletes a learning opportunity in a specific org.
	DeleteLearningOpportunityByOrgId(ctx context.Context, in *DeleteLearningOpportunityByOrgIdRequest, opts ...grpc.CallOption) (*DeleteLearningOpportunityResponse, error)
}

type agentTrainingSupportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentTrainingSupportServiceClient(cc grpc.ClientConnInterface) AgentTrainingSupportServiceClient {
	return &agentTrainingSupportServiceClient{cc}
}

func (c *agentTrainingSupportServiceClient) ListLearningOpportunitiesByOrgId(ctx context.Context, in *ListLearningOpportunitiesByOrgIdRequest, opts ...grpc.CallOption) (*ListLearningOpportunitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLearningOpportunitiesResponse)
	err := c.cc.Invoke(ctx, AgentTrainingSupportService_ListLearningOpportunitiesByOrgId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTrainingSupportServiceClient) DeleteLearningOpportunityByOrgId(ctx context.Context, in *DeleteLearningOpportunityByOrgIdRequest, opts ...grpc.CallOption) (*DeleteLearningOpportunityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteLearningOpportunityResponse)
	err := c.cc.Invoke(ctx, AgentTrainingSupportService_DeleteLearningOpportunityByOrgId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentTrainingSupportServiceServer is the server API for AgentTrainingSupportService service.
// All implementations must embed UnimplementedAgentTrainingSupportServiceServer
// for forward compatibility
//
// AgentTrainingSupportService provides customer support
// endpoints for agent training.
type AgentTrainingSupportServiceServer interface {
	// ListLearningOpportunitiesByOrgId lists learning opportunities for a specific org.
	ListLearningOpportunitiesByOrgId(context.Context, *ListLearningOpportunitiesByOrgIdRequest) (*ListLearningOpportunitiesResponse, error)
	// DeleteLearningOpportunityByOrgId deletes a learning opportunity in a specific org.
	DeleteLearningOpportunityByOrgId(context.Context, *DeleteLearningOpportunityByOrgIdRequest) (*DeleteLearningOpportunityResponse, error)
	mustEmbedUnimplementedAgentTrainingSupportServiceServer()
}

// UnimplementedAgentTrainingSupportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentTrainingSupportServiceServer struct {
}

func (UnimplementedAgentTrainingSupportServiceServer) ListLearningOpportunitiesByOrgId(context.Context, *ListLearningOpportunitiesByOrgIdRequest) (*ListLearningOpportunitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLearningOpportunitiesByOrgId not implemented")
}
func (UnimplementedAgentTrainingSupportServiceServer) DeleteLearningOpportunityByOrgId(context.Context, *DeleteLearningOpportunityByOrgIdRequest) (*DeleteLearningOpportunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLearningOpportunityByOrgId not implemented")
}
func (UnimplementedAgentTrainingSupportServiceServer) mustEmbedUnimplementedAgentTrainingSupportServiceServer() {
}

// UnsafeAgentTrainingSupportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentTrainingSupportServiceServer will
// result in compilation errors.
type UnsafeAgentTrainingSupportServiceServer interface {
	mustEmbedUnimplementedAgentTrainingSupportServiceServer()
}

func RegisterAgentTrainingSupportServiceServer(s grpc.ServiceRegistrar, srv AgentTrainingSupportServiceServer) {
	s.RegisterService(&AgentTrainingSupportService_ServiceDesc, srv)
}

func _AgentTrainingSupportService_ListLearningOpportunitiesByOrgId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLearningOpportunitiesByOrgIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTrainingSupportServiceServer).ListLearningOpportunitiesByOrgId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentTrainingSupportService_ListLearningOpportunitiesByOrgId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTrainingSupportServiceServer).ListLearningOpportunitiesByOrgId(ctx, req.(*ListLearningOpportunitiesByOrgIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTrainingSupportService_DeleteLearningOpportunityByOrgId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLearningOpportunityByOrgIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTrainingSupportServiceServer).DeleteLearningOpportunityByOrgId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentTrainingSupportService_DeleteLearningOpportunityByOrgId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTrainingSupportServiceServer).DeleteLearningOpportunityByOrgId(ctx, req.(*DeleteLearningOpportunityByOrgIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentTrainingSupportService_ServiceDesc is the grpc.ServiceDesc for AgentTrainingSupportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentTrainingSupportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.agenttraining.AgentTrainingSupportService",
	HandlerType: (*AgentTrainingSupportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLearningOpportunitiesByOrgId",
			Handler:    _AgentTrainingSupportService_ListLearningOpportunitiesByOrgId_Handler,
		},
		{
			MethodName: "DeleteLearningOpportunityByOrgId",
			Handler:    _AgentTrainingSupportService_DeleteLearningOpportunityByOrgId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1alpha1/agenttraining/support_service.proto",
}
