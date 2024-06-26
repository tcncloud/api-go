// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: api/v1alpha1/insights/service.proto

package insights

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
	Insights_CreateInsight_FullMethodName        = "/api.v1alpha1.insights.Insights/CreateInsight"
	Insights_ListInsights_FullMethodName         = "/api.v1alpha1.insights.Insights/ListInsights"
	Insights_ListOrgInsights_FullMethodName      = "/api.v1alpha1.insights.Insights/ListOrgInsights"
	Insights_UpdateInsight_FullMethodName        = "/api.v1alpha1.insights.Insights/UpdateInsight"
	Insights_DeleteInsight_FullMethodName        = "/api.v1alpha1.insights.Insights/DeleteInsight"
	Insights_GetInsight_FullMethodName           = "/api.v1alpha1.insights.Insights/GetInsight"
	Insights_CreateCommonsInsight_FullMethodName = "/api.v1alpha1.insights.Insights/CreateCommonsInsight"
	Insights_UpdateCommonsInsight_FullMethodName = "/api.v1alpha1.insights.Insights/UpdateCommonsInsight"
	Insights_DeleteCommonsInsight_FullMethodName = "/api.v1alpha1.insights.Insights/DeleteCommonsInsight"
	Insights_GetVfsSchema_FullMethodName         = "/api.v1alpha1.insights.Insights/GetVfsSchema"
	Insights_ListVfses_FullMethodName            = "/api.v1alpha1.insights.Insights/ListVfses"
	Insights_ListVfsSchemas_FullMethodName       = "/api.v1alpha1.insights.Insights/ListVfsSchemas"
	Insights_PublishInsight_FullMethodName       = "/api.v1alpha1.insights.Insights/PublishInsight"
)

// InsightsClient is the client API for Insights service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InsightsClient interface {
	// CreateInsight creates a new insight
	CreateInsight(ctx context.Context, in *CreateInsightRequest, opts ...grpc.CallOption) (*CreateInsightResponse, error)
	// ListInsights lists insights
	ListInsights(ctx context.Context, in *ListInsightsRequest, opts ...grpc.CallOption) (*ListInsightsResponse, error)
	// ListOrgInsights lists insights for an org. Used for support app.
	ListOrgInsights(ctx context.Context, in *ListOrgInsightsRequest, opts ...grpc.CallOption) (*ListOrgInsightsResponse, error)
	// UpdateInsight updates an existing insight
	UpdateInsight(ctx context.Context, in *UpdateInsightRequest, opts ...grpc.CallOption) (*UpdateInsightResponse, error)
	// DeleteInsight deletes a insight
	DeleteInsight(ctx context.Context, in *DeleteInsightRequest, opts ...grpc.CallOption) (*DeleteInsightResponse, error)
	// GetInsight gets a insight by id
	GetInsight(ctx context.Context, in *GetInsightRequest, opts ...grpc.CallOption) (*GetInsightResponse, error)
	// CreateCommonsInsight is deprecated.
	CreateCommonsInsight(ctx context.Context, in *CreateInsightRequest, opts ...grpc.CallOption) (*CreateInsightResponse, error)
	// UpdateCommonsInsight is deprecated.
	UpdateCommonsInsight(ctx context.Context, in *UpdateInsightRequest, opts ...grpc.CallOption) (*UpdateInsightResponse, error)
	// DeleteCommonsInsight is deprecated.
	DeleteCommonsInsight(ctx context.Context, in *DeleteInsightRequest, opts ...grpc.CallOption) (*DeleteInsightResponse, error)
	// GetVfsSchema gets schema for a vfs
	GetVfsSchema(ctx context.Context, in *GetVfsSchemaRequest, opts ...grpc.CallOption) (*GetVfsSchemaResponse, error)
	// ListVfses lists exported vfs aliases
	ListVfses(ctx context.Context, in *ListVfsesRequest, opts ...grpc.CallOption) (*ListVfsesResponse, error)
	// ListVfses lists exported vfs aliases
	ListVfsSchemas(ctx context.Context, in *ListVfsSchemasRequest, opts ...grpc.CallOption) (*ListVfsSchemasResponse, error)
	// PublishInsight publishes an insight
	PublishInsight(ctx context.Context, in *PublishInsightRequest, opts ...grpc.CallOption) (*PublishInsightResponse, error)
}

type insightsClient struct {
	cc grpc.ClientConnInterface
}

func NewInsightsClient(cc grpc.ClientConnInterface) InsightsClient {
	return &insightsClient{cc}
}

func (c *insightsClient) CreateInsight(ctx context.Context, in *CreateInsightRequest, opts ...grpc.CallOption) (*CreateInsightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateInsightResponse)
	err := c.cc.Invoke(ctx, Insights_CreateInsight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) ListInsights(ctx context.Context, in *ListInsightsRequest, opts ...grpc.CallOption) (*ListInsightsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListInsightsResponse)
	err := c.cc.Invoke(ctx, Insights_ListInsights_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) ListOrgInsights(ctx context.Context, in *ListOrgInsightsRequest, opts ...grpc.CallOption) (*ListOrgInsightsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOrgInsightsResponse)
	err := c.cc.Invoke(ctx, Insights_ListOrgInsights_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) UpdateInsight(ctx context.Context, in *UpdateInsightRequest, opts ...grpc.CallOption) (*UpdateInsightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateInsightResponse)
	err := c.cc.Invoke(ctx, Insights_UpdateInsight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) DeleteInsight(ctx context.Context, in *DeleteInsightRequest, opts ...grpc.CallOption) (*DeleteInsightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteInsightResponse)
	err := c.cc.Invoke(ctx, Insights_DeleteInsight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) GetInsight(ctx context.Context, in *GetInsightRequest, opts ...grpc.CallOption) (*GetInsightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInsightResponse)
	err := c.cc.Invoke(ctx, Insights_GetInsight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) CreateCommonsInsight(ctx context.Context, in *CreateInsightRequest, opts ...grpc.CallOption) (*CreateInsightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateInsightResponse)
	err := c.cc.Invoke(ctx, Insights_CreateCommonsInsight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) UpdateCommonsInsight(ctx context.Context, in *UpdateInsightRequest, opts ...grpc.CallOption) (*UpdateInsightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateInsightResponse)
	err := c.cc.Invoke(ctx, Insights_UpdateCommonsInsight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) DeleteCommonsInsight(ctx context.Context, in *DeleteInsightRequest, opts ...grpc.CallOption) (*DeleteInsightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteInsightResponse)
	err := c.cc.Invoke(ctx, Insights_DeleteCommonsInsight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) GetVfsSchema(ctx context.Context, in *GetVfsSchemaRequest, opts ...grpc.CallOption) (*GetVfsSchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVfsSchemaResponse)
	err := c.cc.Invoke(ctx, Insights_GetVfsSchema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) ListVfses(ctx context.Context, in *ListVfsesRequest, opts ...grpc.CallOption) (*ListVfsesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVfsesResponse)
	err := c.cc.Invoke(ctx, Insights_ListVfses_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) ListVfsSchemas(ctx context.Context, in *ListVfsSchemasRequest, opts ...grpc.CallOption) (*ListVfsSchemasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVfsSchemasResponse)
	err := c.cc.Invoke(ctx, Insights_ListVfsSchemas_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *insightsClient) PublishInsight(ctx context.Context, in *PublishInsightRequest, opts ...grpc.CallOption) (*PublishInsightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishInsightResponse)
	err := c.cc.Invoke(ctx, Insights_PublishInsight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InsightsServer is the server API for Insights service.
// All implementations must embed UnimplementedInsightsServer
// for forward compatibility
type InsightsServer interface {
	// CreateInsight creates a new insight
	CreateInsight(context.Context, *CreateInsightRequest) (*CreateInsightResponse, error)
	// ListInsights lists insights
	ListInsights(context.Context, *ListInsightsRequest) (*ListInsightsResponse, error)
	// ListOrgInsights lists insights for an org. Used for support app.
	ListOrgInsights(context.Context, *ListOrgInsightsRequest) (*ListOrgInsightsResponse, error)
	// UpdateInsight updates an existing insight
	UpdateInsight(context.Context, *UpdateInsightRequest) (*UpdateInsightResponse, error)
	// DeleteInsight deletes a insight
	DeleteInsight(context.Context, *DeleteInsightRequest) (*DeleteInsightResponse, error)
	// GetInsight gets a insight by id
	GetInsight(context.Context, *GetInsightRequest) (*GetInsightResponse, error)
	// CreateCommonsInsight is deprecated.
	CreateCommonsInsight(context.Context, *CreateInsightRequest) (*CreateInsightResponse, error)
	// UpdateCommonsInsight is deprecated.
	UpdateCommonsInsight(context.Context, *UpdateInsightRequest) (*UpdateInsightResponse, error)
	// DeleteCommonsInsight is deprecated.
	DeleteCommonsInsight(context.Context, *DeleteInsightRequest) (*DeleteInsightResponse, error)
	// GetVfsSchema gets schema for a vfs
	GetVfsSchema(context.Context, *GetVfsSchemaRequest) (*GetVfsSchemaResponse, error)
	// ListVfses lists exported vfs aliases
	ListVfses(context.Context, *ListVfsesRequest) (*ListVfsesResponse, error)
	// ListVfses lists exported vfs aliases
	ListVfsSchemas(context.Context, *ListVfsSchemasRequest) (*ListVfsSchemasResponse, error)
	// PublishInsight publishes an insight
	PublishInsight(context.Context, *PublishInsightRequest) (*PublishInsightResponse, error)
	mustEmbedUnimplementedInsightsServer()
}

// UnimplementedInsightsServer must be embedded to have forward compatible implementations.
type UnimplementedInsightsServer struct {
}

func (UnimplementedInsightsServer) CreateInsight(context.Context, *CreateInsightRequest) (*CreateInsightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInsight not implemented")
}
func (UnimplementedInsightsServer) ListInsights(context.Context, *ListInsightsRequest) (*ListInsightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInsights not implemented")
}
func (UnimplementedInsightsServer) ListOrgInsights(context.Context, *ListOrgInsightsRequest) (*ListOrgInsightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrgInsights not implemented")
}
func (UnimplementedInsightsServer) UpdateInsight(context.Context, *UpdateInsightRequest) (*UpdateInsightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInsight not implemented")
}
func (UnimplementedInsightsServer) DeleteInsight(context.Context, *DeleteInsightRequest) (*DeleteInsightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInsight not implemented")
}
func (UnimplementedInsightsServer) GetInsight(context.Context, *GetInsightRequest) (*GetInsightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInsight not implemented")
}
func (UnimplementedInsightsServer) CreateCommonsInsight(context.Context, *CreateInsightRequest) (*CreateInsightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommonsInsight not implemented")
}
func (UnimplementedInsightsServer) UpdateCommonsInsight(context.Context, *UpdateInsightRequest) (*UpdateInsightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommonsInsight not implemented")
}
func (UnimplementedInsightsServer) DeleteCommonsInsight(context.Context, *DeleteInsightRequest) (*DeleteInsightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommonsInsight not implemented")
}
func (UnimplementedInsightsServer) GetVfsSchema(context.Context, *GetVfsSchemaRequest) (*GetVfsSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVfsSchema not implemented")
}
func (UnimplementedInsightsServer) ListVfses(context.Context, *ListVfsesRequest) (*ListVfsesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVfses not implemented")
}
func (UnimplementedInsightsServer) ListVfsSchemas(context.Context, *ListVfsSchemasRequest) (*ListVfsSchemasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVfsSchemas not implemented")
}
func (UnimplementedInsightsServer) PublishInsight(context.Context, *PublishInsightRequest) (*PublishInsightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishInsight not implemented")
}
func (UnimplementedInsightsServer) mustEmbedUnimplementedInsightsServer() {}

// UnsafeInsightsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InsightsServer will
// result in compilation errors.
type UnsafeInsightsServer interface {
	mustEmbedUnimplementedInsightsServer()
}

func RegisterInsightsServer(s grpc.ServiceRegistrar, srv InsightsServer) {
	s.RegisterService(&Insights_ServiceDesc, srv)
}

func _Insights_CreateInsight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInsightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).CreateInsight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_CreateInsight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).CreateInsight(ctx, req.(*CreateInsightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_ListInsights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInsightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).ListInsights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_ListInsights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).ListInsights(ctx, req.(*ListInsightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_ListOrgInsights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrgInsightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).ListOrgInsights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_ListOrgInsights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).ListOrgInsights(ctx, req.(*ListOrgInsightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_UpdateInsight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInsightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).UpdateInsight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_UpdateInsight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).UpdateInsight(ctx, req.(*UpdateInsightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_DeleteInsight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInsightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).DeleteInsight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_DeleteInsight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).DeleteInsight(ctx, req.(*DeleteInsightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_GetInsight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInsightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).GetInsight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_GetInsight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).GetInsight(ctx, req.(*GetInsightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_CreateCommonsInsight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInsightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).CreateCommonsInsight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_CreateCommonsInsight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).CreateCommonsInsight(ctx, req.(*CreateInsightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_UpdateCommonsInsight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInsightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).UpdateCommonsInsight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_UpdateCommonsInsight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).UpdateCommonsInsight(ctx, req.(*UpdateInsightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_DeleteCommonsInsight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInsightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).DeleteCommonsInsight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_DeleteCommonsInsight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).DeleteCommonsInsight(ctx, req.(*DeleteInsightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_GetVfsSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVfsSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).GetVfsSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_GetVfsSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).GetVfsSchema(ctx, req.(*GetVfsSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_ListVfses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVfsesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).ListVfses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_ListVfses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).ListVfses(ctx, req.(*ListVfsesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_ListVfsSchemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVfsSchemasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).ListVfsSchemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_ListVfsSchemas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).ListVfsSchemas(ctx, req.(*ListVfsSchemasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Insights_PublishInsight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishInsightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsightsServer).PublishInsight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Insights_PublishInsight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsightsServer).PublishInsight(ctx, req.(*PublishInsightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Insights_ServiceDesc is the grpc.ServiceDesc for Insights service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Insights_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.insights.Insights",
	HandlerType: (*InsightsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInsight",
			Handler:    _Insights_CreateInsight_Handler,
		},
		{
			MethodName: "ListInsights",
			Handler:    _Insights_ListInsights_Handler,
		},
		{
			MethodName: "ListOrgInsights",
			Handler:    _Insights_ListOrgInsights_Handler,
		},
		{
			MethodName: "UpdateInsight",
			Handler:    _Insights_UpdateInsight_Handler,
		},
		{
			MethodName: "DeleteInsight",
			Handler:    _Insights_DeleteInsight_Handler,
		},
		{
			MethodName: "GetInsight",
			Handler:    _Insights_GetInsight_Handler,
		},
		{
			MethodName: "CreateCommonsInsight",
			Handler:    _Insights_CreateCommonsInsight_Handler,
		},
		{
			MethodName: "UpdateCommonsInsight",
			Handler:    _Insights_UpdateCommonsInsight_Handler,
		},
		{
			MethodName: "DeleteCommonsInsight",
			Handler:    _Insights_DeleteCommonsInsight_Handler,
		},
		{
			MethodName: "GetVfsSchema",
			Handler:    _Insights_GetVfsSchema_Handler,
		},
		{
			MethodName: "ListVfses",
			Handler:    _Insights_ListVfses_Handler,
		},
		{
			MethodName: "ListVfsSchemas",
			Handler:    _Insights_ListVfsSchemas_Handler,
		},
		{
			MethodName: "PublishInsight",
			Handler:    _Insights_PublishInsight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1alpha1/insights/service.proto",
}
