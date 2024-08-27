// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: tcnapi/omni/projects/v1/service.proto

package projectsv1

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
	Projects_ListProjects_FullMethodName  = "/tcnapi.omni.projects.v1.Projects/ListProjects"
	Projects_GetProject_FullMethodName    = "/tcnapi.omni.projects.v1.Projects/GetProject"
	Projects_CreateProject_FullMethodName = "/tcnapi.omni.projects.v1.Projects/CreateProject"
	Projects_UpdateProject_FullMethodName = "/tcnapi.omni.projects.v1.Projects/UpdateProject"
	Projects_DeleteProject_FullMethodName = "/tcnapi.omni.projects.v1.Projects/DeleteProject"
)

// ProjectsClient is the client API for Projects service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Projects service is the public api for the omni/projects service.
type ProjectsClient interface {
	// Method to list projects
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	// Method to get a rpoject
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error)
	// Method to create a project
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*Project, error)
	// Method to update a project
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*Project, error)
	// Method to delete a project
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*Project, error)
}

type projectsClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectsClient(cc grpc.ClientConnInterface) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, Projects_ListProjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Project)
	err := c.cc.Invoke(ctx, Projects_GetProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Project)
	err := c.cc.Invoke(ctx, Projects_CreateProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Project)
	err := c.cc.Invoke(ctx, Projects_UpdateProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Project)
	err := c.cc.Invoke(ctx, Projects_DeleteProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectsServer is the server API for Projects service.
// All implementations must embed UnimplementedProjectsServer
// for forward compatibility.
//
// Projects service is the public api for the omni/projects service.
type ProjectsServer interface {
	// Method to list projects
	ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
	// Method to get a rpoject
	GetProject(context.Context, *GetProjectRequest) (*Project, error)
	// Method to create a project
	CreateProject(context.Context, *CreateProjectRequest) (*Project, error)
	// Method to update a project
	UpdateProject(context.Context, *UpdateProjectRequest) (*Project, error)
	// Method to delete a project
	DeleteProject(context.Context, *DeleteProjectRequest) (*Project, error)
	mustEmbedUnimplementedProjectsServer()
}

// UnimplementedProjectsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProjectsServer struct{}

func (UnimplementedProjectsServer) ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (UnimplementedProjectsServer) GetProject(context.Context, *GetProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectsServer) CreateProject(context.Context, *CreateProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectsServer) UpdateProject(context.Context, *UpdateProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedProjectsServer) DeleteProject(context.Context, *DeleteProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedProjectsServer) mustEmbedUnimplementedProjectsServer() {}
func (UnimplementedProjectsServer) testEmbeddedByValue()                  {}

// UnsafeProjectsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectsServer will
// result in compilation errors.
type UnsafeProjectsServer interface {
	mustEmbedUnimplementedProjectsServer()
}

func RegisterProjectsServer(s grpc.ServiceRegistrar, srv ProjectsServer) {
	// If the following call pancis, it indicates UnimplementedProjectsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Projects_ServiceDesc, srv)
}

func _Projects_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Projects_ListProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).ListProjects(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Projects_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Projects_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Projects_UpdateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).UpdateProject(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Projects_DeleteProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Projects_ServiceDesc is the grpc.ServiceDesc for Projects service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Projects_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tcnapi.omni.projects.v1.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProjects",
			Handler:    _Projects_ListProjects_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Projects_GetProject_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _Projects_CreateProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _Projects_UpdateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _Projects_DeleteProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tcnapi/omni/projects/v1/service.proto",
}