// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: services/omnichannel/tasks/v1alpha1/service.proto

package tasksv1alpha1

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
	TasksService_CancelTasks_FullMethodName = "/services.omnichannel.tasks.v1alpha1.TasksService/CancelTasks"
)

// TasksServiceClient is the client API for TasksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TasksServiceClient interface {
	CancelTasks(ctx context.Context, in *CancelTasksRequest, opts ...grpc.CallOption) (*CancelTasksResponse, error)
}

type tasksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTasksServiceClient(cc grpc.ClientConnInterface) TasksServiceClient {
	return &tasksServiceClient{cc}
}

func (c *tasksServiceClient) CancelTasks(ctx context.Context, in *CancelTasksRequest, opts ...grpc.CallOption) (*CancelTasksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelTasksResponse)
	err := c.cc.Invoke(ctx, TasksService_CancelTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServiceServer is the server API for TasksService service.
// All implementations must embed UnimplementedTasksServiceServer
// for forward compatibility
type TasksServiceServer interface {
	CancelTasks(context.Context, *CancelTasksRequest) (*CancelTasksResponse, error)
	mustEmbedUnimplementedTasksServiceServer()
}

// UnimplementedTasksServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTasksServiceServer struct {
}

func (UnimplementedTasksServiceServer) CancelTasks(context.Context, *CancelTasksRequest) (*CancelTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTasks not implemented")
}
func (UnimplementedTasksServiceServer) mustEmbedUnimplementedTasksServiceServer() {}

// UnsafeTasksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TasksServiceServer will
// result in compilation errors.
type UnsafeTasksServiceServer interface {
	mustEmbedUnimplementedTasksServiceServer()
}

func RegisterTasksServiceServer(s grpc.ServiceRegistrar, srv TasksServiceServer) {
	s.RegisterService(&TasksService_ServiceDesc, srv)
}

func _TasksService_CancelTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).CancelTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TasksService_CancelTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).CancelTasks(ctx, req.(*CancelTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TasksService_ServiceDesc is the grpc.ServiceDesc for TasksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TasksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.omnichannel.tasks.v1alpha1.TasksService",
	HandlerType: (*TasksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelTasks",
			Handler:    _TasksService_CancelTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/omnichannel/tasks/v1alpha1/service.proto",
}