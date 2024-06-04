// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: services/pbx/v1/service.proto

package pbxv1

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
	PBXService_QueryPbxUsers_FullMethodName         = "/services.pbx.v1.PBXService/QueryPbxUsers"
	PBXService_QueryRingGroups_FullMethodName       = "/services.pbx.v1.PBXService/QueryRingGroups"
	PBXService_UpdatePbxUser_FullMethodName         = "/services.pbx.v1.PBXService/UpdatePbxUser"
	PBXService_UpdateRingGroup_FullMethodName       = "/services.pbx.v1.PBXService/UpdateRingGroup"
	PBXService_CreateRingGroup_FullMethodName       = "/services.pbx.v1.PBXService/CreateRingGroup"
	PBXService_DeleteRingGroup_FullMethodName       = "/services.pbx.v1.PBXService/DeleteRingGroup"
	PBXService_AssignRandomExtension_FullMethodName = "/services.pbx.v1.PBXService/AssignRandomExtension"
)

// PBXServiceClient is the client API for PBXService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PBXServiceClient interface {
	// Queries details of PBX Users based on specified criteria for the authenticated callers ORG
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	QueryPbxUsers(ctx context.Context, in *QueryPbxUsersRequest, opts ...grpc.CallOption) (*QueryPbxUsersResponse, error)
	// Queries details of Ring Groups based on specified criteria for the authenticated callers ORG
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	QueryRingGroups(ctx context.Context, in *QueryRingGroupsRequest, opts ...grpc.CallOption) (*QueryRingGroupsResponse, error)
	// Updates details of a PBX User for the authenticated callers ORG.
	// Allows for updating, activating, and deactivating a user.
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	UpdatePbxUser(ctx context.Context, in *UpdatePbxUserRequest, opts ...grpc.CallOption) (*UpdatePbxUserResponse, error)
	// Updates details of a Ring Group for the authenticated callers ORG. This operation acts as an "upsert".
	//   - If the groupID is in the update mask and the group exists, the group will be updated.
	//   - If the groupID is not in the update mask a group will be created.
	//
	// Allows for creating and updating a ring group.
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is invalid.
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	UpdateRingGroup(ctx context.Context, in *UpdateRingGroupRequest, opts ...grpc.CallOption) (*UpdateRingGroupResponse, error)
	// Creates a ring group for the authenticated caller's ORG.
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is invalid.
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	CreateRingGroup(ctx context.Context, in *CreateRingGroupRequest, opts ...grpc.CallOption) (*CreateRingGroupResponse, error)
	// Deletes a specific Ring Group for the authenticated caller's ORG.
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.InvalidArgument: The groupID is an invalid format.
	//   - grpc.NotFound: The group does not exist or is not in the caller's ORG.
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	DeleteRingGroup(ctx context.Context, in *DeleteRingGroupRequest, opts ...grpc.CallOption) (*DeleteRingGroupResponse, error)
	// Assigns a random extension either to a PBX user or a Ring Group
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	AssignRandomExtension(ctx context.Context, in *AssignRandomExtensionRequest, opts ...grpc.CallOption) (*AssignRandomExtensionResponse, error)
}

type pBXServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPBXServiceClient(cc grpc.ClientConnInterface) PBXServiceClient {
	return &pBXServiceClient{cc}
}

func (c *pBXServiceClient) QueryPbxUsers(ctx context.Context, in *QueryPbxUsersRequest, opts ...grpc.CallOption) (*QueryPbxUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryPbxUsersResponse)
	err := c.cc.Invoke(ctx, PBXService_QueryPbxUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pBXServiceClient) QueryRingGroups(ctx context.Context, in *QueryRingGroupsRequest, opts ...grpc.CallOption) (*QueryRingGroupsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryRingGroupsResponse)
	err := c.cc.Invoke(ctx, PBXService_QueryRingGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pBXServiceClient) UpdatePbxUser(ctx context.Context, in *UpdatePbxUserRequest, opts ...grpc.CallOption) (*UpdatePbxUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePbxUserResponse)
	err := c.cc.Invoke(ctx, PBXService_UpdatePbxUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pBXServiceClient) UpdateRingGroup(ctx context.Context, in *UpdateRingGroupRequest, opts ...grpc.CallOption) (*UpdateRingGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRingGroupResponse)
	err := c.cc.Invoke(ctx, PBXService_UpdateRingGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pBXServiceClient) CreateRingGroup(ctx context.Context, in *CreateRingGroupRequest, opts ...grpc.CallOption) (*CreateRingGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRingGroupResponse)
	err := c.cc.Invoke(ctx, PBXService_CreateRingGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pBXServiceClient) DeleteRingGroup(ctx context.Context, in *DeleteRingGroupRequest, opts ...grpc.CallOption) (*DeleteRingGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRingGroupResponse)
	err := c.cc.Invoke(ctx, PBXService_DeleteRingGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pBXServiceClient) AssignRandomExtension(ctx context.Context, in *AssignRandomExtensionRequest, opts ...grpc.CallOption) (*AssignRandomExtensionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssignRandomExtensionResponse)
	err := c.cc.Invoke(ctx, PBXService_AssignRandomExtension_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PBXServiceServer is the server API for PBXService service.
// All implementations must embed UnimplementedPBXServiceServer
// for forward compatibility
type PBXServiceServer interface {
	// Queries details of PBX Users based on specified criteria for the authenticated callers ORG
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	QueryPbxUsers(context.Context, *QueryPbxUsersRequest) (*QueryPbxUsersResponse, error)
	// Queries details of Ring Groups based on specified criteria for the authenticated callers ORG
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	QueryRingGroups(context.Context, *QueryRingGroupsRequest) (*QueryRingGroupsResponse, error)
	// Updates details of a PBX User for the authenticated callers ORG.
	// Allows for updating, activating, and deactivating a user.
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	UpdatePbxUser(context.Context, *UpdatePbxUserRequest) (*UpdatePbxUserResponse, error)
	// Updates details of a Ring Group for the authenticated callers ORG. This operation acts as an "upsert".
	//   - If the groupID is in the update mask and the group exists, the group will be updated.
	//   - If the groupID is not in the update mask a group will be created.
	//
	// Allows for creating and updating a ring group.
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is invalid.
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	UpdateRingGroup(context.Context, *UpdateRingGroupRequest) (*UpdateRingGroupResponse, error)
	// Creates a ring group for the authenticated caller's ORG.
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is invalid.
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	CreateRingGroup(context.Context, *CreateRingGroupRequest) (*CreateRingGroupResponse, error)
	// Deletes a specific Ring Group for the authenticated caller's ORG.
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.InvalidArgument: The groupID is an invalid format.
	//   - grpc.NotFound: The group does not exist or is not in the caller's ORG.
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	DeleteRingGroup(context.Context, *DeleteRingGroupRequest) (*DeleteRingGroupResponse, error)
	// Assigns a random extension either to a PBX user or a Ring Group
	// Required permissions:
	//
	//	PBX-MANAGER
	//
	// Errors:
	//   - grpc.PermissionDenied: Caller doesn't have the required permissions.
	//   - grpc.Internal: An internal error occurred.
	//   - grpc.Unavailable: The operation is currently unavailable. Likely a transient issue with a downstream service.
	AssignRandomExtension(context.Context, *AssignRandomExtensionRequest) (*AssignRandomExtensionResponse, error)
	mustEmbedUnimplementedPBXServiceServer()
}

// UnimplementedPBXServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPBXServiceServer struct {
}

func (UnimplementedPBXServiceServer) QueryPbxUsers(context.Context, *QueryPbxUsersRequest) (*QueryPbxUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPbxUsers not implemented")
}
func (UnimplementedPBXServiceServer) QueryRingGroups(context.Context, *QueryRingGroupsRequest) (*QueryRingGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRingGroups not implemented")
}
func (UnimplementedPBXServiceServer) UpdatePbxUser(context.Context, *UpdatePbxUserRequest) (*UpdatePbxUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePbxUser not implemented")
}
func (UnimplementedPBXServiceServer) UpdateRingGroup(context.Context, *UpdateRingGroupRequest) (*UpdateRingGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRingGroup not implemented")
}
func (UnimplementedPBXServiceServer) CreateRingGroup(context.Context, *CreateRingGroupRequest) (*CreateRingGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRingGroup not implemented")
}
func (UnimplementedPBXServiceServer) DeleteRingGroup(context.Context, *DeleteRingGroupRequest) (*DeleteRingGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRingGroup not implemented")
}
func (UnimplementedPBXServiceServer) AssignRandomExtension(context.Context, *AssignRandomExtensionRequest) (*AssignRandomExtensionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignRandomExtension not implemented")
}
func (UnimplementedPBXServiceServer) mustEmbedUnimplementedPBXServiceServer() {}

// UnsafePBXServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PBXServiceServer will
// result in compilation errors.
type UnsafePBXServiceServer interface {
	mustEmbedUnimplementedPBXServiceServer()
}

func RegisterPBXServiceServer(s grpc.ServiceRegistrar, srv PBXServiceServer) {
	s.RegisterService(&PBXService_ServiceDesc, srv)
}

func _PBXService_QueryPbxUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPbxUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PBXServiceServer).QueryPbxUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PBXService_QueryPbxUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PBXServiceServer).QueryPbxUsers(ctx, req.(*QueryPbxUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PBXService_QueryRingGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRingGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PBXServiceServer).QueryRingGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PBXService_QueryRingGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PBXServiceServer).QueryRingGroups(ctx, req.(*QueryRingGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PBXService_UpdatePbxUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePbxUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PBXServiceServer).UpdatePbxUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PBXService_UpdatePbxUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PBXServiceServer).UpdatePbxUser(ctx, req.(*UpdatePbxUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PBXService_UpdateRingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PBXServiceServer).UpdateRingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PBXService_UpdateRingGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PBXServiceServer).UpdateRingGroup(ctx, req.(*UpdateRingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PBXService_CreateRingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PBXServiceServer).CreateRingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PBXService_CreateRingGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PBXServiceServer).CreateRingGroup(ctx, req.(*CreateRingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PBXService_DeleteRingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PBXServiceServer).DeleteRingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PBXService_DeleteRingGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PBXServiceServer).DeleteRingGroup(ctx, req.(*DeleteRingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PBXService_AssignRandomExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRandomExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PBXServiceServer).AssignRandomExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PBXService_AssignRandomExtension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PBXServiceServer).AssignRandomExtension(ctx, req.(*AssignRandomExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PBXService_ServiceDesc is the grpc.ServiceDesc for PBXService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PBXService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.pbx.v1.PBXService",
	HandlerType: (*PBXServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryPbxUsers",
			Handler:    _PBXService_QueryPbxUsers_Handler,
		},
		{
			MethodName: "QueryRingGroups",
			Handler:    _PBXService_QueryRingGroups_Handler,
		},
		{
			MethodName: "UpdatePbxUser",
			Handler:    _PBXService_UpdatePbxUser_Handler,
		},
		{
			MethodName: "UpdateRingGroup",
			Handler:    _PBXService_UpdateRingGroup_Handler,
		},
		{
			MethodName: "CreateRingGroup",
			Handler:    _PBXService_CreateRingGroup_Handler,
		},
		{
			MethodName: "DeleteRingGroup",
			Handler:    _PBXService_DeleteRingGroup_Handler,
		},
		{
			MethodName: "AssignRandomExtension",
			Handler:    _PBXService_AssignRandomExtension_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/pbx/v1/service.proto",
}
