// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v1alpha1/org/skills/service.proto

package skills

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SkillsService_CreateSkillGroup_FullMethodName     = "/api.v1alpha1.org.skills.SkillsService/CreateSkillGroup"
	SkillsService_ListSkillGroups_FullMethodName      = "/api.v1alpha1.org.skills.SkillsService/ListSkillGroups"
	SkillsService_UpdateSkillGroup_FullMethodName     = "/api.v1alpha1.org.skills.SkillsService/UpdateSkillGroup"
	SkillsService_GetSkillGroup_FullMethodName        = "/api.v1alpha1.org.skills.SkillsService/GetSkillGroup"
	SkillsService_DeleteSkillGroup_FullMethodName     = "/api.v1alpha1.org.skills.SkillsService/DeleteSkillGroup"
	SkillsService_AssignSkillGroups_FullMethodName    = "/api.v1alpha1.org.skills.SkillsService/AssignSkillGroups"
	SkillsService_RevokeSkillGroups_FullMethodName    = "/api.v1alpha1.org.skills.SkillsService/RevokeSkillGroups"
	SkillsService_GetUserSkillGroups_FullMethodName   = "/api.v1alpha1.org.skills.SkillsService/GetUserSkillGroups"
	SkillsService_GetUserSkills_FullMethodName        = "/api.v1alpha1.org.skills.SkillsService/GetUserSkills"
	SkillsService_GetSkillGroupMembers_FullMethodName = "/api.v1alpha1.org.skills.SkillsService/GetSkillGroupMembers"
)

// SkillsServiceClient is the client API for SkillsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SkillsServiceClient interface {
	// CreateSkillGroup creates a new skill group.
	CreateSkillGroup(ctx context.Context, in *CreateSkillGroupRequest, opts ...grpc.CallOption) (*CreateSkillGroupResponse, error)
	// ListSkillGroups lists the skill groups belonging to an organization.
	ListSkillGroups(ctx context.Context, in *ListSkillGroupsRequest, opts ...grpc.CallOption) (*ListSkillGroupsResponse, error)
	// UpdateSkillGroup updates a single skill group.
	UpdateSkillGroup(ctx context.Context, in *UpdateSkillGroupRequest, opts ...grpc.CallOption) (*UpdateSkillGroupResponse, error)
	// GetSkillGroup gets a single skill group.
	GetSkillGroup(ctx context.Context, in *GetSkillGroupRequest, opts ...grpc.CallOption) (*GetSkillGroupResponse, error)
	// DeleteSkillGroup deletes a skill group.
	DeleteSkillGroup(ctx context.Context, in *DeleteSkillGroupRequest, opts ...grpc.CallOption) (*DeleteSkillGroupResponse, error)
	// AssignSkillGroups assigns a user to the given skill groups.
	AssignSkillGroups(ctx context.Context, in *AssignSkillGroupsRequest, opts ...grpc.CallOption) (*AssignSkillGroupsResponse, error)
	// RevokeSkillGroups revokes the given skill groups from a user.
	RevokeSkillGroups(ctx context.Context, in *RevokeSkillGroupsRequest, opts ...grpc.CallOption) (*RevokeSkillGroupsResponse, error)
	// GetUserSkillGroups gets the skill groups assigned to a user.
	GetUserSkillGroups(ctx context.Context, in *GetUserSkillGroupsRequest, opts ...grpc.CallOption) (*GetUserSkillGroupsResponse, error)
	// GetUserSkills gets a user's skill proficiencies.
	GetUserSkills(ctx context.Context, in *GetUserSkillsRequest, opts ...grpc.CallOption) (*GetUserSkillsResponse, error)
	// GetSkillGroupMembers gets the members of a skill group.
	GetSkillGroupMembers(ctx context.Context, in *GetSkillGroupMembersRequest, opts ...grpc.CallOption) (*GetSkillGroupMembersResponse, error)
}

type skillsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSkillsServiceClient(cc grpc.ClientConnInterface) SkillsServiceClient {
	return &skillsServiceClient{cc}
}

func (c *skillsServiceClient) CreateSkillGroup(ctx context.Context, in *CreateSkillGroupRequest, opts ...grpc.CallOption) (*CreateSkillGroupResponse, error) {
	out := new(CreateSkillGroupResponse)
	err := c.cc.Invoke(ctx, SkillsService_CreateSkillGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillsServiceClient) ListSkillGroups(ctx context.Context, in *ListSkillGroupsRequest, opts ...grpc.CallOption) (*ListSkillGroupsResponse, error) {
	out := new(ListSkillGroupsResponse)
	err := c.cc.Invoke(ctx, SkillsService_ListSkillGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillsServiceClient) UpdateSkillGroup(ctx context.Context, in *UpdateSkillGroupRequest, opts ...grpc.CallOption) (*UpdateSkillGroupResponse, error) {
	out := new(UpdateSkillGroupResponse)
	err := c.cc.Invoke(ctx, SkillsService_UpdateSkillGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillsServiceClient) GetSkillGroup(ctx context.Context, in *GetSkillGroupRequest, opts ...grpc.CallOption) (*GetSkillGroupResponse, error) {
	out := new(GetSkillGroupResponse)
	err := c.cc.Invoke(ctx, SkillsService_GetSkillGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillsServiceClient) DeleteSkillGroup(ctx context.Context, in *DeleteSkillGroupRequest, opts ...grpc.CallOption) (*DeleteSkillGroupResponse, error) {
	out := new(DeleteSkillGroupResponse)
	err := c.cc.Invoke(ctx, SkillsService_DeleteSkillGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillsServiceClient) AssignSkillGroups(ctx context.Context, in *AssignSkillGroupsRequest, opts ...grpc.CallOption) (*AssignSkillGroupsResponse, error) {
	out := new(AssignSkillGroupsResponse)
	err := c.cc.Invoke(ctx, SkillsService_AssignSkillGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillsServiceClient) RevokeSkillGroups(ctx context.Context, in *RevokeSkillGroupsRequest, opts ...grpc.CallOption) (*RevokeSkillGroupsResponse, error) {
	out := new(RevokeSkillGroupsResponse)
	err := c.cc.Invoke(ctx, SkillsService_RevokeSkillGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillsServiceClient) GetUserSkillGroups(ctx context.Context, in *GetUserSkillGroupsRequest, opts ...grpc.CallOption) (*GetUserSkillGroupsResponse, error) {
	out := new(GetUserSkillGroupsResponse)
	err := c.cc.Invoke(ctx, SkillsService_GetUserSkillGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillsServiceClient) GetUserSkills(ctx context.Context, in *GetUserSkillsRequest, opts ...grpc.CallOption) (*GetUserSkillsResponse, error) {
	out := new(GetUserSkillsResponse)
	err := c.cc.Invoke(ctx, SkillsService_GetUserSkills_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillsServiceClient) GetSkillGroupMembers(ctx context.Context, in *GetSkillGroupMembersRequest, opts ...grpc.CallOption) (*GetSkillGroupMembersResponse, error) {
	out := new(GetSkillGroupMembersResponse)
	err := c.cc.Invoke(ctx, SkillsService_GetSkillGroupMembers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkillsServiceServer is the server API for SkillsService service.
// All implementations must embed UnimplementedSkillsServiceServer
// for forward compatibility
type SkillsServiceServer interface {
	// CreateSkillGroup creates a new skill group.
	CreateSkillGroup(context.Context, *CreateSkillGroupRequest) (*CreateSkillGroupResponse, error)
	// ListSkillGroups lists the skill groups belonging to an organization.
	ListSkillGroups(context.Context, *ListSkillGroupsRequest) (*ListSkillGroupsResponse, error)
	// UpdateSkillGroup updates a single skill group.
	UpdateSkillGroup(context.Context, *UpdateSkillGroupRequest) (*UpdateSkillGroupResponse, error)
	// GetSkillGroup gets a single skill group.
	GetSkillGroup(context.Context, *GetSkillGroupRequest) (*GetSkillGroupResponse, error)
	// DeleteSkillGroup deletes a skill group.
	DeleteSkillGroup(context.Context, *DeleteSkillGroupRequest) (*DeleteSkillGroupResponse, error)
	// AssignSkillGroups assigns a user to the given skill groups.
	AssignSkillGroups(context.Context, *AssignSkillGroupsRequest) (*AssignSkillGroupsResponse, error)
	// RevokeSkillGroups revokes the given skill groups from a user.
	RevokeSkillGroups(context.Context, *RevokeSkillGroupsRequest) (*RevokeSkillGroupsResponse, error)
	// GetUserSkillGroups gets the skill groups assigned to a user.
	GetUserSkillGroups(context.Context, *GetUserSkillGroupsRequest) (*GetUserSkillGroupsResponse, error)
	// GetUserSkills gets a user's skill proficiencies.
	GetUserSkills(context.Context, *GetUserSkillsRequest) (*GetUserSkillsResponse, error)
	// GetSkillGroupMembers gets the members of a skill group.
	GetSkillGroupMembers(context.Context, *GetSkillGroupMembersRequest) (*GetSkillGroupMembersResponse, error)
	mustEmbedUnimplementedSkillsServiceServer()
}

// UnimplementedSkillsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSkillsServiceServer struct {
}

func (UnimplementedSkillsServiceServer) CreateSkillGroup(context.Context, *CreateSkillGroupRequest) (*CreateSkillGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSkillGroup not implemented")
}
func (UnimplementedSkillsServiceServer) ListSkillGroups(context.Context, *ListSkillGroupsRequest) (*ListSkillGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSkillGroups not implemented")
}
func (UnimplementedSkillsServiceServer) UpdateSkillGroup(context.Context, *UpdateSkillGroupRequest) (*UpdateSkillGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSkillGroup not implemented")
}
func (UnimplementedSkillsServiceServer) GetSkillGroup(context.Context, *GetSkillGroupRequest) (*GetSkillGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkillGroup not implemented")
}
func (UnimplementedSkillsServiceServer) DeleteSkillGroup(context.Context, *DeleteSkillGroupRequest) (*DeleteSkillGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSkillGroup not implemented")
}
func (UnimplementedSkillsServiceServer) AssignSkillGroups(context.Context, *AssignSkillGroupsRequest) (*AssignSkillGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignSkillGroups not implemented")
}
func (UnimplementedSkillsServiceServer) RevokeSkillGroups(context.Context, *RevokeSkillGroupsRequest) (*RevokeSkillGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeSkillGroups not implemented")
}
func (UnimplementedSkillsServiceServer) GetUserSkillGroups(context.Context, *GetUserSkillGroupsRequest) (*GetUserSkillGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSkillGroups not implemented")
}
func (UnimplementedSkillsServiceServer) GetUserSkills(context.Context, *GetUserSkillsRequest) (*GetUserSkillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSkills not implemented")
}
func (UnimplementedSkillsServiceServer) GetSkillGroupMembers(context.Context, *GetSkillGroupMembersRequest) (*GetSkillGroupMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkillGroupMembers not implemented")
}
func (UnimplementedSkillsServiceServer) mustEmbedUnimplementedSkillsServiceServer() {}

// UnsafeSkillsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SkillsServiceServer will
// result in compilation errors.
type UnsafeSkillsServiceServer interface {
	mustEmbedUnimplementedSkillsServiceServer()
}

func RegisterSkillsServiceServer(s grpc.ServiceRegistrar, srv SkillsServiceServer) {
	s.RegisterService(&SkillsService_ServiceDesc, srv)
}

func _SkillsService_CreateSkillGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSkillGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillsServiceServer).CreateSkillGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillsService_CreateSkillGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillsServiceServer).CreateSkillGroup(ctx, req.(*CreateSkillGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillsService_ListSkillGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSkillGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillsServiceServer).ListSkillGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillsService_ListSkillGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillsServiceServer).ListSkillGroups(ctx, req.(*ListSkillGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillsService_UpdateSkillGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSkillGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillsServiceServer).UpdateSkillGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillsService_UpdateSkillGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillsServiceServer).UpdateSkillGroup(ctx, req.(*UpdateSkillGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillsService_GetSkillGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkillGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillsServiceServer).GetSkillGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillsService_GetSkillGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillsServiceServer).GetSkillGroup(ctx, req.(*GetSkillGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillsService_DeleteSkillGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSkillGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillsServiceServer).DeleteSkillGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillsService_DeleteSkillGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillsServiceServer).DeleteSkillGroup(ctx, req.(*DeleteSkillGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillsService_AssignSkillGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignSkillGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillsServiceServer).AssignSkillGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillsService_AssignSkillGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillsServiceServer).AssignSkillGroups(ctx, req.(*AssignSkillGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillsService_RevokeSkillGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeSkillGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillsServiceServer).RevokeSkillGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillsService_RevokeSkillGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillsServiceServer).RevokeSkillGroups(ctx, req.(*RevokeSkillGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillsService_GetUserSkillGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSkillGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillsServiceServer).GetUserSkillGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillsService_GetUserSkillGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillsServiceServer).GetUserSkillGroups(ctx, req.(*GetUserSkillGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillsService_GetUserSkills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSkillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillsServiceServer).GetUserSkills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillsService_GetUserSkills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillsServiceServer).GetUserSkills(ctx, req.(*GetUserSkillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillsService_GetSkillGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkillGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillsServiceServer).GetSkillGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillsService_GetSkillGroupMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillsServiceServer).GetSkillGroupMembers(ctx, req.(*GetSkillGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SkillsService_ServiceDesc is the grpc.ServiceDesc for SkillsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SkillsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.org.skills.SkillsService",
	HandlerType: (*SkillsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSkillGroup",
			Handler:    _SkillsService_CreateSkillGroup_Handler,
		},
		{
			MethodName: "ListSkillGroups",
			Handler:    _SkillsService_ListSkillGroups_Handler,
		},
		{
			MethodName: "UpdateSkillGroup",
			Handler:    _SkillsService_UpdateSkillGroup_Handler,
		},
		{
			MethodName: "GetSkillGroup",
			Handler:    _SkillsService_GetSkillGroup_Handler,
		},
		{
			MethodName: "DeleteSkillGroup",
			Handler:    _SkillsService_DeleteSkillGroup_Handler,
		},
		{
			MethodName: "AssignSkillGroups",
			Handler:    _SkillsService_AssignSkillGroups_Handler,
		},
		{
			MethodName: "RevokeSkillGroups",
			Handler:    _SkillsService_RevokeSkillGroups_Handler,
		},
		{
			MethodName: "GetUserSkillGroups",
			Handler:    _SkillsService_GetUserSkillGroups_Handler,
		},
		{
			MethodName: "GetUserSkills",
			Handler:    _SkillsService_GetUserSkills_Handler,
		},
		{
			MethodName: "GetSkillGroupMembers",
			Handler:    _SkillsService_GetSkillGroupMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1alpha1/org/skills/service.proto",
}