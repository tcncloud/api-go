// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/v1alpha1/org/labels/service.proto

package labels

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
	LabelsService_CreateLabel_FullMethodName         = "/api.v1alpha1.org.labels.LabelsService/CreateLabel"
	LabelsService_GetLabel_FullMethodName            = "/api.v1alpha1.org.labels.LabelsService/GetLabel"
	LabelsService_UpdateLabel_FullMethodName         = "/api.v1alpha1.org.labels.LabelsService/UpdateLabel"
	LabelsService_ListLabels_FullMethodName          = "/api.v1alpha1.org.labels.LabelsService/ListLabels"
	LabelsService_DeleteLabel_FullMethodName         = "/api.v1alpha1.org.labels.LabelsService/DeleteLabel"
	LabelsService_AttachLabel_FullMethodName         = "/api.v1alpha1.org.labels.LabelsService/AttachLabel"
	LabelsService_DetachLabel_FullMethodName         = "/api.v1alpha1.org.labels.LabelsService/DetachLabel"
	LabelsService_GetLabeledEntityMap_FullMethodName = "/api.v1alpha1.org.labels.LabelsService/GetLabeledEntityMap"
	LabelsService_AssignLabels_FullMethodName        = "/api.v1alpha1.org.labels.LabelsService/AssignLabels"
	LabelsService_RevokeLabels_FullMethodName        = "/api.v1alpha1.org.labels.LabelsService/RevokeLabels"
)

// LabelsServiceClient is the client API for LabelsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LabelsServiceClient interface {
	// CreateLabel creates a new label.
	CreateLabel(ctx context.Context, in *CreateLabelRequest, opts ...grpc.CallOption) (*CreateLabelResponse, error)
	// GetLabel gets a single label
	GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*GetLabelResponse, error)
	// UpdateLabel updates a single label
	UpdateLabel(ctx context.Context, in *UpdateLabelRequest, opts ...grpc.CallOption) (*UpdateLabelResponse, error)
	// ListLabels lists all labels for a given organization
	ListLabels(ctx context.Context, in *ListLabelsRequest, opts ...grpc.CallOption) (*ListLabelsResponse, error)
	// DeleteLabel deletes a single label
	DeleteLabel(ctx context.Context, in *DeleteLabelRequest, opts ...grpc.CallOption) (*DeleteLabelResponse, error)
	// AttachLabel attaches a label to a given entity type
	AttachLabel(ctx context.Context, in *AttachLabelRequest, opts ...grpc.CallOption) (*AttachLabelResponse, error)
	// DetachLabel detaches a label from an entity based on an entity type
	DetachLabel(ctx context.Context, in *DetachLabelRequest, opts ...grpc.CallOption) (*DetachLabelResponse, error)
	// GetLabeledEntityMap gives back a map of entity Id to attached labels. The Entity type is specified on the request
	GetLabeledEntityMap(ctx context.Context, in *GetLabeledEntityMapRequest, opts ...grpc.CallOption) (*GetLabeledEntityMapResponse, error)
	// AssignLabels assigns labels to a specific permission group.
	AssignLabels(ctx context.Context, in *AssignLabelsRequest, opts ...grpc.CallOption) (*AssignLabelsResponse, error)
	// RevokeLabels revokes labels from a specific permission group.
	RevokeLabels(ctx context.Context, in *RevokeLabelsRequest, opts ...grpc.CallOption) (*RevokeLabelsResponse, error)
}

type labelsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLabelsServiceClient(cc grpc.ClientConnInterface) LabelsServiceClient {
	return &labelsServiceClient{cc}
}

func (c *labelsServiceClient) CreateLabel(ctx context.Context, in *CreateLabelRequest, opts ...grpc.CallOption) (*CreateLabelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLabelResponse)
	err := c.cc.Invoke(ctx, LabelsService_CreateLabel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelsServiceClient) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*GetLabelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLabelResponse)
	err := c.cc.Invoke(ctx, LabelsService_GetLabel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelsServiceClient) UpdateLabel(ctx context.Context, in *UpdateLabelRequest, opts ...grpc.CallOption) (*UpdateLabelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateLabelResponse)
	err := c.cc.Invoke(ctx, LabelsService_UpdateLabel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelsServiceClient) ListLabels(ctx context.Context, in *ListLabelsRequest, opts ...grpc.CallOption) (*ListLabelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLabelsResponse)
	err := c.cc.Invoke(ctx, LabelsService_ListLabels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelsServiceClient) DeleteLabel(ctx context.Context, in *DeleteLabelRequest, opts ...grpc.CallOption) (*DeleteLabelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteLabelResponse)
	err := c.cc.Invoke(ctx, LabelsService_DeleteLabel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelsServiceClient) AttachLabel(ctx context.Context, in *AttachLabelRequest, opts ...grpc.CallOption) (*AttachLabelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AttachLabelResponse)
	err := c.cc.Invoke(ctx, LabelsService_AttachLabel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelsServiceClient) DetachLabel(ctx context.Context, in *DetachLabelRequest, opts ...grpc.CallOption) (*DetachLabelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetachLabelResponse)
	err := c.cc.Invoke(ctx, LabelsService_DetachLabel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelsServiceClient) GetLabeledEntityMap(ctx context.Context, in *GetLabeledEntityMapRequest, opts ...grpc.CallOption) (*GetLabeledEntityMapResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLabeledEntityMapResponse)
	err := c.cc.Invoke(ctx, LabelsService_GetLabeledEntityMap_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelsServiceClient) AssignLabels(ctx context.Context, in *AssignLabelsRequest, opts ...grpc.CallOption) (*AssignLabelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssignLabelsResponse)
	err := c.cc.Invoke(ctx, LabelsService_AssignLabels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelsServiceClient) RevokeLabels(ctx context.Context, in *RevokeLabelsRequest, opts ...grpc.CallOption) (*RevokeLabelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeLabelsResponse)
	err := c.cc.Invoke(ctx, LabelsService_RevokeLabels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabelsServiceServer is the server API for LabelsService service.
// All implementations must embed UnimplementedLabelsServiceServer
// for forward compatibility.
type LabelsServiceServer interface {
	// CreateLabel creates a new label.
	CreateLabel(context.Context, *CreateLabelRequest) (*CreateLabelResponse, error)
	// GetLabel gets a single label
	GetLabel(context.Context, *GetLabelRequest) (*GetLabelResponse, error)
	// UpdateLabel updates a single label
	UpdateLabel(context.Context, *UpdateLabelRequest) (*UpdateLabelResponse, error)
	// ListLabels lists all labels for a given organization
	ListLabels(context.Context, *ListLabelsRequest) (*ListLabelsResponse, error)
	// DeleteLabel deletes a single label
	DeleteLabel(context.Context, *DeleteLabelRequest) (*DeleteLabelResponse, error)
	// AttachLabel attaches a label to a given entity type
	AttachLabel(context.Context, *AttachLabelRequest) (*AttachLabelResponse, error)
	// DetachLabel detaches a label from an entity based on an entity type
	DetachLabel(context.Context, *DetachLabelRequest) (*DetachLabelResponse, error)
	// GetLabeledEntityMap gives back a map of entity Id to attached labels. The Entity type is specified on the request
	GetLabeledEntityMap(context.Context, *GetLabeledEntityMapRequest) (*GetLabeledEntityMapResponse, error)
	// AssignLabels assigns labels to a specific permission group.
	AssignLabels(context.Context, *AssignLabelsRequest) (*AssignLabelsResponse, error)
	// RevokeLabels revokes labels from a specific permission group.
	RevokeLabels(context.Context, *RevokeLabelsRequest) (*RevokeLabelsResponse, error)
	mustEmbedUnimplementedLabelsServiceServer()
}

// UnimplementedLabelsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLabelsServiceServer struct{}

func (UnimplementedLabelsServiceServer) CreateLabel(context.Context, *CreateLabelRequest) (*CreateLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLabel not implemented")
}
func (UnimplementedLabelsServiceServer) GetLabel(context.Context, *GetLabelRequest) (*GetLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabel not implemented")
}
func (UnimplementedLabelsServiceServer) UpdateLabel(context.Context, *UpdateLabelRequest) (*UpdateLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLabel not implemented")
}
func (UnimplementedLabelsServiceServer) ListLabels(context.Context, *ListLabelsRequest) (*ListLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLabels not implemented")
}
func (UnimplementedLabelsServiceServer) DeleteLabel(context.Context, *DeleteLabelRequest) (*DeleteLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLabel not implemented")
}
func (UnimplementedLabelsServiceServer) AttachLabel(context.Context, *AttachLabelRequest) (*AttachLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachLabel not implemented")
}
func (UnimplementedLabelsServiceServer) DetachLabel(context.Context, *DetachLabelRequest) (*DetachLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachLabel not implemented")
}
func (UnimplementedLabelsServiceServer) GetLabeledEntityMap(context.Context, *GetLabeledEntityMapRequest) (*GetLabeledEntityMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabeledEntityMap not implemented")
}
func (UnimplementedLabelsServiceServer) AssignLabels(context.Context, *AssignLabelsRequest) (*AssignLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignLabels not implemented")
}
func (UnimplementedLabelsServiceServer) RevokeLabels(context.Context, *RevokeLabelsRequest) (*RevokeLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeLabels not implemented")
}
func (UnimplementedLabelsServiceServer) mustEmbedUnimplementedLabelsServiceServer() {}
func (UnimplementedLabelsServiceServer) testEmbeddedByValue()                       {}

// UnsafeLabelsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LabelsServiceServer will
// result in compilation errors.
type UnsafeLabelsServiceServer interface {
	mustEmbedUnimplementedLabelsServiceServer()
}

func RegisterLabelsServiceServer(s grpc.ServiceRegistrar, srv LabelsServiceServer) {
	// If the following call pancis, it indicates UnimplementedLabelsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LabelsService_ServiceDesc, srv)
}

func _LabelsService_CreateLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelsServiceServer).CreateLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelsService_CreateLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelsServiceServer).CreateLabel(ctx, req.(*CreateLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelsService_GetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelsServiceServer).GetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelsService_GetLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelsServiceServer).GetLabel(ctx, req.(*GetLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelsService_UpdateLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelsServiceServer).UpdateLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelsService_UpdateLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelsServiceServer).UpdateLabel(ctx, req.(*UpdateLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelsService_ListLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelsServiceServer).ListLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelsService_ListLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelsServiceServer).ListLabels(ctx, req.(*ListLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelsService_DeleteLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelsServiceServer).DeleteLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelsService_DeleteLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelsServiceServer).DeleteLabel(ctx, req.(*DeleteLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelsService_AttachLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelsServiceServer).AttachLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelsService_AttachLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelsServiceServer).AttachLabel(ctx, req.(*AttachLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelsService_DetachLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetachLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelsServiceServer).DetachLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelsService_DetachLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelsServiceServer).DetachLabel(ctx, req.(*DetachLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelsService_GetLabeledEntityMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabeledEntityMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelsServiceServer).GetLabeledEntityMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelsService_GetLabeledEntityMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelsServiceServer).GetLabeledEntityMap(ctx, req.(*GetLabeledEntityMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelsService_AssignLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelsServiceServer).AssignLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelsService_AssignLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelsServiceServer).AssignLabels(ctx, req.(*AssignLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelsService_RevokeLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelsServiceServer).RevokeLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelsService_RevokeLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelsServiceServer).RevokeLabels(ctx, req.(*RevokeLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LabelsService_ServiceDesc is the grpc.ServiceDesc for LabelsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LabelsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.org.labels.LabelsService",
	HandlerType: (*LabelsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLabel",
			Handler:    _LabelsService_CreateLabel_Handler,
		},
		{
			MethodName: "GetLabel",
			Handler:    _LabelsService_GetLabel_Handler,
		},
		{
			MethodName: "UpdateLabel",
			Handler:    _LabelsService_UpdateLabel_Handler,
		},
		{
			MethodName: "ListLabels",
			Handler:    _LabelsService_ListLabels_Handler,
		},
		{
			MethodName: "DeleteLabel",
			Handler:    _LabelsService_DeleteLabel_Handler,
		},
		{
			MethodName: "AttachLabel",
			Handler:    _LabelsService_AttachLabel_Handler,
		},
		{
			MethodName: "DetachLabel",
			Handler:    _LabelsService_DetachLabel_Handler,
		},
		{
			MethodName: "GetLabeledEntityMap",
			Handler:    _LabelsService_GetLabeledEntityMap_Handler,
		},
		{
			MethodName: "AssignLabels",
			Handler:    _LabelsService_AssignLabels_Handler,
		},
		{
			MethodName: "RevokeLabels",
			Handler:    _LabelsService_RevokeLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1alpha1/org/labels/service.proto",
}
