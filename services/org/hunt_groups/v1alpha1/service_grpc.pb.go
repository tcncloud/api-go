// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: services/org/hunt_groups/v1alpha1/service.proto

package hunt_groupsv1alpha1

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
	HuntGroupsService_ListHuntGroupExileLinks_FullMethodName          = "/services.org.hunt_groups.v1alpha1.HuntGroupsService/ListHuntGroupExileLinks"
	HuntGroupsService_CopyHuntGroupExileLink_FullMethodName           = "/services.org.hunt_groups.v1alpha1.HuntGroupsService/CopyHuntGroupExileLink"
	HuntGroupsService_UpdateHuntGroupExileLinks_FullMethodName        = "/services.org.hunt_groups.v1alpha1.HuntGroupsService/UpdateHuntGroupExileLinks"
	HuntGroupsService_ListHuntGroupAgentTriggers_FullMethodName       = "/services.org.hunt_groups.v1alpha1.HuntGroupsService/ListHuntGroupAgentTriggers"
	HuntGroupsService_CopyHuntGroupAgentTrigger_FullMethodName        = "/services.org.hunt_groups.v1alpha1.HuntGroupsService/CopyHuntGroupAgentTrigger"
	HuntGroupsService_UpdateHuntGroupAgentTriggers_FullMethodName     = "/services.org.hunt_groups.v1alpha1.HuntGroupsService/UpdateHuntGroupAgentTriggers"
	HuntGroupsService_CopyHuntGroupToOrganization_FullMethodName      = "/services.org.hunt_groups.v1alpha1.HuntGroupsService/CopyHuntGroupToOrganization"
	HuntGroupsService_AdminCopyHuntGroupToOrganization_FullMethodName = "/services.org.hunt_groups.v1alpha1.HuntGroupsService/AdminCopyHuntGroupToOrganization"
	HuntGroupsService_AdminListHuntGroups_FullMethodName              = "/services.org.hunt_groups.v1alpha1.HuntGroupsService/AdminListHuntGroups"
)

// HuntGroupsServiceClient is the client API for HuntGroupsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// HuntGroupsService is the service for managing hunt groups and their related entities.
type HuntGroupsServiceClient interface {
	// ListHuntGroupExileLinks returns a list of Exile links for a given hunt group.
	ListHuntGroupExileLinks(ctx context.Context, in *ListHuntGroupExileLinksRequest, opts ...grpc.CallOption) (*ListHuntGroupExileLinksResponse, error)
	// CopyHuntGroupExileLink copies an exile link from one hunt group to another.
	// It will create a new exile link in the destination hunt group with the same
	// settings/parameters as the source exile link.
	CopyHuntGroupExileLink(ctx context.Context, in *CopyHuntGroupExileLinkRequest, opts ...grpc.CallOption) (*CopyHuntGroupExileLinkResponse, error)
	// UpdateHuntGroupExileLinks updates the exile links for a hunt group.
	// It will create any new exile links that do not already exist in the hunt group,
	// update any existing exile links with the new settings/parameters, and
	// delete any exile links that are not in the request.
	UpdateHuntGroupExileLinks(ctx context.Context, in *UpdateHuntGroupExileLinksRequest, opts ...grpc.CallOption) (*UpdateHuntGroupExileLinksResponse, error)
	// ListHuntGroupAgentTriggers returns a list of agent triggers for the given hunt group.
	ListHuntGroupAgentTriggers(ctx context.Context, in *ListHuntGroupAgentTriggersRequest, opts ...grpc.CallOption) (*ListHuntGroupAgentTriggersResponse, error)
	// CopyHuntGroupAgentTrigger copys an agent trigger to the given hunt group in the same org.
	CopyHuntGroupAgentTrigger(ctx context.Context, in *CopyHuntGroupAgentTriggerRequest, opts ...grpc.CallOption) (*CopyHuntGroupAgentTriggerResponse, error)
	// UpdateHuntGroupAgentTriggers updates all agent triggers for the given hunt group.
	UpdateHuntGroupAgentTriggers(ctx context.Context, in *UpdateHuntGroupAgentTriggersRequest, opts ...grpc.CallOption) (*UpdateHuntGroupAgentTriggersResponse, error)
	// CopyHuntGroupToOrganization copies a hunt group to a different organization.
	// The destination organization must be a child of the source organization.
	// This would create a new hunt group in the destination organization with the same
	// settings/parameters and all associated data (skill, call-queue config) as the source hunt group.
	CopyHuntGroupToOrganization(ctx context.Context, in *CopyHuntGroupToOrganizationRequest, opts ...grpc.CallOption) (*CopyHuntGroupToOrganizationResponse, error)
	// AdminCopyHuntGroupToOrganization copies a hunt group to a different organization.
	// This will create a new hunt group in the destination organization with the same
	// settings/parameters and all associated data (skill, call-queue config) as the source hunt group.
	AdminCopyHuntGroupToOrganization(ctx context.Context, in *AdminCopyHuntGroupToOrganizationRequest, opts ...grpc.CallOption) (*AdminCopyHuntGroupToOrganizationResponse, error)
	// AdminListHuntGroups returns a list of hunt groups for the given organization.
	AdminListHuntGroups(ctx context.Context, in *AdminListHuntGroupsRequest, opts ...grpc.CallOption) (*AdminListHuntGroupsResponse, error)
}

type huntGroupsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHuntGroupsServiceClient(cc grpc.ClientConnInterface) HuntGroupsServiceClient {
	return &huntGroupsServiceClient{cc}
}

func (c *huntGroupsServiceClient) ListHuntGroupExileLinks(ctx context.Context, in *ListHuntGroupExileLinksRequest, opts ...grpc.CallOption) (*ListHuntGroupExileLinksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListHuntGroupExileLinksResponse)
	err := c.cc.Invoke(ctx, HuntGroupsService_ListHuntGroupExileLinks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huntGroupsServiceClient) CopyHuntGroupExileLink(ctx context.Context, in *CopyHuntGroupExileLinkRequest, opts ...grpc.CallOption) (*CopyHuntGroupExileLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CopyHuntGroupExileLinkResponse)
	err := c.cc.Invoke(ctx, HuntGroupsService_CopyHuntGroupExileLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huntGroupsServiceClient) UpdateHuntGroupExileLinks(ctx context.Context, in *UpdateHuntGroupExileLinksRequest, opts ...grpc.CallOption) (*UpdateHuntGroupExileLinksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateHuntGroupExileLinksResponse)
	err := c.cc.Invoke(ctx, HuntGroupsService_UpdateHuntGroupExileLinks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huntGroupsServiceClient) ListHuntGroupAgentTriggers(ctx context.Context, in *ListHuntGroupAgentTriggersRequest, opts ...grpc.CallOption) (*ListHuntGroupAgentTriggersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListHuntGroupAgentTriggersResponse)
	err := c.cc.Invoke(ctx, HuntGroupsService_ListHuntGroupAgentTriggers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huntGroupsServiceClient) CopyHuntGroupAgentTrigger(ctx context.Context, in *CopyHuntGroupAgentTriggerRequest, opts ...grpc.CallOption) (*CopyHuntGroupAgentTriggerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CopyHuntGroupAgentTriggerResponse)
	err := c.cc.Invoke(ctx, HuntGroupsService_CopyHuntGroupAgentTrigger_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huntGroupsServiceClient) UpdateHuntGroupAgentTriggers(ctx context.Context, in *UpdateHuntGroupAgentTriggersRequest, opts ...grpc.CallOption) (*UpdateHuntGroupAgentTriggersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateHuntGroupAgentTriggersResponse)
	err := c.cc.Invoke(ctx, HuntGroupsService_UpdateHuntGroupAgentTriggers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huntGroupsServiceClient) CopyHuntGroupToOrganization(ctx context.Context, in *CopyHuntGroupToOrganizationRequest, opts ...grpc.CallOption) (*CopyHuntGroupToOrganizationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CopyHuntGroupToOrganizationResponse)
	err := c.cc.Invoke(ctx, HuntGroupsService_CopyHuntGroupToOrganization_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huntGroupsServiceClient) AdminCopyHuntGroupToOrganization(ctx context.Context, in *AdminCopyHuntGroupToOrganizationRequest, opts ...grpc.CallOption) (*AdminCopyHuntGroupToOrganizationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminCopyHuntGroupToOrganizationResponse)
	err := c.cc.Invoke(ctx, HuntGroupsService_AdminCopyHuntGroupToOrganization_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huntGroupsServiceClient) AdminListHuntGroups(ctx context.Context, in *AdminListHuntGroupsRequest, opts ...grpc.CallOption) (*AdminListHuntGroupsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminListHuntGroupsResponse)
	err := c.cc.Invoke(ctx, HuntGroupsService_AdminListHuntGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HuntGroupsServiceServer is the server API for HuntGroupsService service.
// All implementations must embed UnimplementedHuntGroupsServiceServer
// for forward compatibility.
//
// HuntGroupsService is the service for managing hunt groups and their related entities.
type HuntGroupsServiceServer interface {
	// ListHuntGroupExileLinks returns a list of Exile links for a given hunt group.
	ListHuntGroupExileLinks(context.Context, *ListHuntGroupExileLinksRequest) (*ListHuntGroupExileLinksResponse, error)
	// CopyHuntGroupExileLink copies an exile link from one hunt group to another.
	// It will create a new exile link in the destination hunt group with the same
	// settings/parameters as the source exile link.
	CopyHuntGroupExileLink(context.Context, *CopyHuntGroupExileLinkRequest) (*CopyHuntGroupExileLinkResponse, error)
	// UpdateHuntGroupExileLinks updates the exile links for a hunt group.
	// It will create any new exile links that do not already exist in the hunt group,
	// update any existing exile links with the new settings/parameters, and
	// delete any exile links that are not in the request.
	UpdateHuntGroupExileLinks(context.Context, *UpdateHuntGroupExileLinksRequest) (*UpdateHuntGroupExileLinksResponse, error)
	// ListHuntGroupAgentTriggers returns a list of agent triggers for the given hunt group.
	ListHuntGroupAgentTriggers(context.Context, *ListHuntGroupAgentTriggersRequest) (*ListHuntGroupAgentTriggersResponse, error)
	// CopyHuntGroupAgentTrigger copys an agent trigger to the given hunt group in the same org.
	CopyHuntGroupAgentTrigger(context.Context, *CopyHuntGroupAgentTriggerRequest) (*CopyHuntGroupAgentTriggerResponse, error)
	// UpdateHuntGroupAgentTriggers updates all agent triggers for the given hunt group.
	UpdateHuntGroupAgentTriggers(context.Context, *UpdateHuntGroupAgentTriggersRequest) (*UpdateHuntGroupAgentTriggersResponse, error)
	// CopyHuntGroupToOrganization copies a hunt group to a different organization.
	// The destination organization must be a child of the source organization.
	// This would create a new hunt group in the destination organization with the same
	// settings/parameters and all associated data (skill, call-queue config) as the source hunt group.
	CopyHuntGroupToOrganization(context.Context, *CopyHuntGroupToOrganizationRequest) (*CopyHuntGroupToOrganizationResponse, error)
	// AdminCopyHuntGroupToOrganization copies a hunt group to a different organization.
	// This will create a new hunt group in the destination organization with the same
	// settings/parameters and all associated data (skill, call-queue config) as the source hunt group.
	AdminCopyHuntGroupToOrganization(context.Context, *AdminCopyHuntGroupToOrganizationRequest) (*AdminCopyHuntGroupToOrganizationResponse, error)
	// AdminListHuntGroups returns a list of hunt groups for the given organization.
	AdminListHuntGroups(context.Context, *AdminListHuntGroupsRequest) (*AdminListHuntGroupsResponse, error)
	mustEmbedUnimplementedHuntGroupsServiceServer()
}

// UnimplementedHuntGroupsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHuntGroupsServiceServer struct{}

func (UnimplementedHuntGroupsServiceServer) ListHuntGroupExileLinks(context.Context, *ListHuntGroupExileLinksRequest) (*ListHuntGroupExileLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHuntGroupExileLinks not implemented")
}
func (UnimplementedHuntGroupsServiceServer) CopyHuntGroupExileLink(context.Context, *CopyHuntGroupExileLinkRequest) (*CopyHuntGroupExileLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyHuntGroupExileLink not implemented")
}
func (UnimplementedHuntGroupsServiceServer) UpdateHuntGroupExileLinks(context.Context, *UpdateHuntGroupExileLinksRequest) (*UpdateHuntGroupExileLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHuntGroupExileLinks not implemented")
}
func (UnimplementedHuntGroupsServiceServer) ListHuntGroupAgentTriggers(context.Context, *ListHuntGroupAgentTriggersRequest) (*ListHuntGroupAgentTriggersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHuntGroupAgentTriggers not implemented")
}
func (UnimplementedHuntGroupsServiceServer) CopyHuntGroupAgentTrigger(context.Context, *CopyHuntGroupAgentTriggerRequest) (*CopyHuntGroupAgentTriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyHuntGroupAgentTrigger not implemented")
}
func (UnimplementedHuntGroupsServiceServer) UpdateHuntGroupAgentTriggers(context.Context, *UpdateHuntGroupAgentTriggersRequest) (*UpdateHuntGroupAgentTriggersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHuntGroupAgentTriggers not implemented")
}
func (UnimplementedHuntGroupsServiceServer) CopyHuntGroupToOrganization(context.Context, *CopyHuntGroupToOrganizationRequest) (*CopyHuntGroupToOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyHuntGroupToOrganization not implemented")
}
func (UnimplementedHuntGroupsServiceServer) AdminCopyHuntGroupToOrganization(context.Context, *AdminCopyHuntGroupToOrganizationRequest) (*AdminCopyHuntGroupToOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCopyHuntGroupToOrganization not implemented")
}
func (UnimplementedHuntGroupsServiceServer) AdminListHuntGroups(context.Context, *AdminListHuntGroupsRequest) (*AdminListHuntGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminListHuntGroups not implemented")
}
func (UnimplementedHuntGroupsServiceServer) mustEmbedUnimplementedHuntGroupsServiceServer() {}
func (UnimplementedHuntGroupsServiceServer) testEmbeddedByValue()                           {}

// UnsafeHuntGroupsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HuntGroupsServiceServer will
// result in compilation errors.
type UnsafeHuntGroupsServiceServer interface {
	mustEmbedUnimplementedHuntGroupsServiceServer()
}

func RegisterHuntGroupsServiceServer(s grpc.ServiceRegistrar, srv HuntGroupsServiceServer) {
	// If the following call pancis, it indicates UnimplementedHuntGroupsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HuntGroupsService_ServiceDesc, srv)
}

func _HuntGroupsService_ListHuntGroupExileLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHuntGroupExileLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuntGroupsServiceServer).ListHuntGroupExileLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuntGroupsService_ListHuntGroupExileLinks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuntGroupsServiceServer).ListHuntGroupExileLinks(ctx, req.(*ListHuntGroupExileLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuntGroupsService_CopyHuntGroupExileLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyHuntGroupExileLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuntGroupsServiceServer).CopyHuntGroupExileLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuntGroupsService_CopyHuntGroupExileLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuntGroupsServiceServer).CopyHuntGroupExileLink(ctx, req.(*CopyHuntGroupExileLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuntGroupsService_UpdateHuntGroupExileLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHuntGroupExileLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuntGroupsServiceServer).UpdateHuntGroupExileLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuntGroupsService_UpdateHuntGroupExileLinks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuntGroupsServiceServer).UpdateHuntGroupExileLinks(ctx, req.(*UpdateHuntGroupExileLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuntGroupsService_ListHuntGroupAgentTriggers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHuntGroupAgentTriggersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuntGroupsServiceServer).ListHuntGroupAgentTriggers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuntGroupsService_ListHuntGroupAgentTriggers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuntGroupsServiceServer).ListHuntGroupAgentTriggers(ctx, req.(*ListHuntGroupAgentTriggersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuntGroupsService_CopyHuntGroupAgentTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyHuntGroupAgentTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuntGroupsServiceServer).CopyHuntGroupAgentTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuntGroupsService_CopyHuntGroupAgentTrigger_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuntGroupsServiceServer).CopyHuntGroupAgentTrigger(ctx, req.(*CopyHuntGroupAgentTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuntGroupsService_UpdateHuntGroupAgentTriggers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHuntGroupAgentTriggersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuntGroupsServiceServer).UpdateHuntGroupAgentTriggers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuntGroupsService_UpdateHuntGroupAgentTriggers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuntGroupsServiceServer).UpdateHuntGroupAgentTriggers(ctx, req.(*UpdateHuntGroupAgentTriggersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuntGroupsService_CopyHuntGroupToOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyHuntGroupToOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuntGroupsServiceServer).CopyHuntGroupToOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuntGroupsService_CopyHuntGroupToOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuntGroupsServiceServer).CopyHuntGroupToOrganization(ctx, req.(*CopyHuntGroupToOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuntGroupsService_AdminCopyHuntGroupToOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCopyHuntGroupToOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuntGroupsServiceServer).AdminCopyHuntGroupToOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuntGroupsService_AdminCopyHuntGroupToOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuntGroupsServiceServer).AdminCopyHuntGroupToOrganization(ctx, req.(*AdminCopyHuntGroupToOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuntGroupsService_AdminListHuntGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminListHuntGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuntGroupsServiceServer).AdminListHuntGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuntGroupsService_AdminListHuntGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuntGroupsServiceServer).AdminListHuntGroups(ctx, req.(*AdminListHuntGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HuntGroupsService_ServiceDesc is the grpc.ServiceDesc for HuntGroupsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HuntGroupsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.org.hunt_groups.v1alpha1.HuntGroupsService",
	HandlerType: (*HuntGroupsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListHuntGroupExileLinks",
			Handler:    _HuntGroupsService_ListHuntGroupExileLinks_Handler,
		},
		{
			MethodName: "CopyHuntGroupExileLink",
			Handler:    _HuntGroupsService_CopyHuntGroupExileLink_Handler,
		},
		{
			MethodName: "UpdateHuntGroupExileLinks",
			Handler:    _HuntGroupsService_UpdateHuntGroupExileLinks_Handler,
		},
		{
			MethodName: "ListHuntGroupAgentTriggers",
			Handler:    _HuntGroupsService_ListHuntGroupAgentTriggers_Handler,
		},
		{
			MethodName: "CopyHuntGroupAgentTrigger",
			Handler:    _HuntGroupsService_CopyHuntGroupAgentTrigger_Handler,
		},
		{
			MethodName: "UpdateHuntGroupAgentTriggers",
			Handler:    _HuntGroupsService_UpdateHuntGroupAgentTriggers_Handler,
		},
		{
			MethodName: "CopyHuntGroupToOrganization",
			Handler:    _HuntGroupsService_CopyHuntGroupToOrganization_Handler,
		},
		{
			MethodName: "AdminCopyHuntGroupToOrganization",
			Handler:    _HuntGroupsService_AdminCopyHuntGroupToOrganization_Handler,
		},
		{
			MethodName: "AdminListHuntGroups",
			Handler:    _HuntGroupsService_AdminListHuntGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/org/hunt_groups/v1alpha1/service.proto",
}
