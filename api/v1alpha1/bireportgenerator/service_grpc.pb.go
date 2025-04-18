// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/v1alpha1/bireportgenerator/service.proto

package bireportgenerator

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
	BIReportGeneratorService_CreateReportJob_FullMethodName      = "/api.v1alpha1.bireportgenerator.BIReportGeneratorService/CreateReportJob"
	BIReportGeneratorService_ListReportJobs_FullMethodName       = "/api.v1alpha1.bireportgenerator.BIReportGeneratorService/ListReportJobs"
	BIReportGeneratorService_UpdateReportJob_FullMethodName      = "/api.v1alpha1.bireportgenerator.BIReportGeneratorService/UpdateReportJob"
	BIReportGeneratorService_DeleteReportJob_FullMethodName      = "/api.v1alpha1.bireportgenerator.BIReportGeneratorService/DeleteReportJob"
	BIReportGeneratorService_GetReportJob_FullMethodName         = "/api.v1alpha1.bireportgenerator.BIReportGeneratorService/GetReportJob"
	BIReportGeneratorService_GenerateReport_FullMethodName       = "/api.v1alpha1.bireportgenerator.BIReportGeneratorService/GenerateReport"
	BIReportGeneratorService_ListReportLogsStream_FullMethodName = "/api.v1alpha1.bireportgenerator.BIReportGeneratorService/ListReportLogsStream"
)

// BIReportGeneratorServiceClient is the client API for BIReportGeneratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// BIReportGeneratorService is the service for generating BI reports.
type BIReportGeneratorServiceClient interface {
	// CreateReportJob creates a new report job.
	CreateReportJob(ctx context.Context, in *CreateReportJobRequest, opts ...grpc.CallOption) (*CreateReportJobResponse, error)
	// ListReportJobs lists report jobs.
	ListReportJobs(ctx context.Context, in *ListReportJobsRequest, opts ...grpc.CallOption) (*ListReportJobsResponse, error)
	// UpdateReportJob updates a report job.
	UpdateReportJob(ctx context.Context, in *UpdateReportJobRequest, opts ...grpc.CallOption) (*UpdateReportJobResponse, error)
	// DeleteReportJob deletes a report job.
	DeleteReportJob(ctx context.Context, in *DeleteReportJobRequest, opts ...grpc.CallOption) (*DeleteReportJobResponse, error)
	// GetReportJob gets a report job.
	GetReportJob(ctx context.Context, in *GetReportJobRequest, opts ...grpc.CallOption) (*GetReportJobResponse, error)
	GenerateReport(ctx context.Context, in *GenerateReportRequest, opts ...grpc.CallOption) (*GenerateReportResponse, error)
	// ListReportLogsStream lists report logs with streaming
	ListReportLogsStream(ctx context.Context, in *ListReportLogsStreamRequest, opts ...grpc.CallOption) (BIReportGeneratorService_ListReportLogsStreamClient, error)
}

type bIReportGeneratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBIReportGeneratorServiceClient(cc grpc.ClientConnInterface) BIReportGeneratorServiceClient {
	return &bIReportGeneratorServiceClient{cc}
}

func (c *bIReportGeneratorServiceClient) CreateReportJob(ctx context.Context, in *CreateReportJobRequest, opts ...grpc.CallOption) (*CreateReportJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateReportJobResponse)
	err := c.cc.Invoke(ctx, BIReportGeneratorService_CreateReportJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIReportGeneratorServiceClient) ListReportJobs(ctx context.Context, in *ListReportJobsRequest, opts ...grpc.CallOption) (*ListReportJobsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListReportJobsResponse)
	err := c.cc.Invoke(ctx, BIReportGeneratorService_ListReportJobs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIReportGeneratorServiceClient) UpdateReportJob(ctx context.Context, in *UpdateReportJobRequest, opts ...grpc.CallOption) (*UpdateReportJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateReportJobResponse)
	err := c.cc.Invoke(ctx, BIReportGeneratorService_UpdateReportJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIReportGeneratorServiceClient) DeleteReportJob(ctx context.Context, in *DeleteReportJobRequest, opts ...grpc.CallOption) (*DeleteReportJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteReportJobResponse)
	err := c.cc.Invoke(ctx, BIReportGeneratorService_DeleteReportJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIReportGeneratorServiceClient) GetReportJob(ctx context.Context, in *GetReportJobRequest, opts ...grpc.CallOption) (*GetReportJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReportJobResponse)
	err := c.cc.Invoke(ctx, BIReportGeneratorService_GetReportJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIReportGeneratorServiceClient) GenerateReport(ctx context.Context, in *GenerateReportRequest, opts ...grpc.CallOption) (*GenerateReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateReportResponse)
	err := c.cc.Invoke(ctx, BIReportGeneratorService_GenerateReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIReportGeneratorServiceClient) ListReportLogsStream(ctx context.Context, in *ListReportLogsStreamRequest, opts ...grpc.CallOption) (BIReportGeneratorService_ListReportLogsStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BIReportGeneratorService_ServiceDesc.Streams[0], BIReportGeneratorService_ListReportLogsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &bIReportGeneratorServiceListReportLogsStreamClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BIReportGeneratorService_ListReportLogsStreamClient interface {
	Recv() (*ListReportLogsStreamResponse, error)
	grpc.ClientStream
}

type bIReportGeneratorServiceListReportLogsStreamClient struct {
	grpc.ClientStream
}

func (x *bIReportGeneratorServiceListReportLogsStreamClient) Recv() (*ListReportLogsStreamResponse, error) {
	m := new(ListReportLogsStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BIReportGeneratorServiceServer is the server API for BIReportGeneratorService service.
// All implementations must embed UnimplementedBIReportGeneratorServiceServer
// for forward compatibility.
//
// BIReportGeneratorService is the service for generating BI reports.
type BIReportGeneratorServiceServer interface {
	// CreateReportJob creates a new report job.
	CreateReportJob(context.Context, *CreateReportJobRequest) (*CreateReportJobResponse, error)
	// ListReportJobs lists report jobs.
	ListReportJobs(context.Context, *ListReportJobsRequest) (*ListReportJobsResponse, error)
	// UpdateReportJob updates a report job.
	UpdateReportJob(context.Context, *UpdateReportJobRequest) (*UpdateReportJobResponse, error)
	// DeleteReportJob deletes a report job.
	DeleteReportJob(context.Context, *DeleteReportJobRequest) (*DeleteReportJobResponse, error)
	// GetReportJob gets a report job.
	GetReportJob(context.Context, *GetReportJobRequest) (*GetReportJobResponse, error)
	GenerateReport(context.Context, *GenerateReportRequest) (*GenerateReportResponse, error)
	// ListReportLogsStream lists report logs with streaming
	ListReportLogsStream(*ListReportLogsStreamRequest, BIReportGeneratorService_ListReportLogsStreamServer) error
	mustEmbedUnimplementedBIReportGeneratorServiceServer()
}

// UnimplementedBIReportGeneratorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBIReportGeneratorServiceServer struct{}

func (UnimplementedBIReportGeneratorServiceServer) CreateReportJob(context.Context, *CreateReportJobRequest) (*CreateReportJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReportJob not implemented")
}
func (UnimplementedBIReportGeneratorServiceServer) ListReportJobs(context.Context, *ListReportJobsRequest) (*ListReportJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReportJobs not implemented")
}
func (UnimplementedBIReportGeneratorServiceServer) UpdateReportJob(context.Context, *UpdateReportJobRequest) (*UpdateReportJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReportJob not implemented")
}
func (UnimplementedBIReportGeneratorServiceServer) DeleteReportJob(context.Context, *DeleteReportJobRequest) (*DeleteReportJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReportJob not implemented")
}
func (UnimplementedBIReportGeneratorServiceServer) GetReportJob(context.Context, *GetReportJobRequest) (*GetReportJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReportJob not implemented")
}
func (UnimplementedBIReportGeneratorServiceServer) GenerateReport(context.Context, *GenerateReportRequest) (*GenerateReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateReport not implemented")
}
func (UnimplementedBIReportGeneratorServiceServer) ListReportLogsStream(*ListReportLogsStreamRequest, BIReportGeneratorService_ListReportLogsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListReportLogsStream not implemented")
}
func (UnimplementedBIReportGeneratorServiceServer) mustEmbedUnimplementedBIReportGeneratorServiceServer() {
}
func (UnimplementedBIReportGeneratorServiceServer) testEmbeddedByValue() {}

// UnsafeBIReportGeneratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BIReportGeneratorServiceServer will
// result in compilation errors.
type UnsafeBIReportGeneratorServiceServer interface {
	mustEmbedUnimplementedBIReportGeneratorServiceServer()
}

func RegisterBIReportGeneratorServiceServer(s grpc.ServiceRegistrar, srv BIReportGeneratorServiceServer) {
	// If the following call pancis, it indicates UnimplementedBIReportGeneratorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BIReportGeneratorService_ServiceDesc, srv)
}

func _BIReportGeneratorService_CreateReportJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReportJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIReportGeneratorServiceServer).CreateReportJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BIReportGeneratorService_CreateReportJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIReportGeneratorServiceServer).CreateReportJob(ctx, req.(*CreateReportJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIReportGeneratorService_ListReportJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReportJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIReportGeneratorServiceServer).ListReportJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BIReportGeneratorService_ListReportJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIReportGeneratorServiceServer).ListReportJobs(ctx, req.(*ListReportJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIReportGeneratorService_UpdateReportJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReportJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIReportGeneratorServiceServer).UpdateReportJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BIReportGeneratorService_UpdateReportJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIReportGeneratorServiceServer).UpdateReportJob(ctx, req.(*UpdateReportJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIReportGeneratorService_DeleteReportJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReportJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIReportGeneratorServiceServer).DeleteReportJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BIReportGeneratorService_DeleteReportJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIReportGeneratorServiceServer).DeleteReportJob(ctx, req.(*DeleteReportJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIReportGeneratorService_GetReportJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIReportGeneratorServiceServer).GetReportJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BIReportGeneratorService_GetReportJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIReportGeneratorServiceServer).GetReportJob(ctx, req.(*GetReportJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIReportGeneratorService_GenerateReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIReportGeneratorServiceServer).GenerateReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BIReportGeneratorService_GenerateReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIReportGeneratorServiceServer).GenerateReport(ctx, req.(*GenerateReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIReportGeneratorService_ListReportLogsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListReportLogsStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BIReportGeneratorServiceServer).ListReportLogsStream(m, &bIReportGeneratorServiceListReportLogsStreamServer{ServerStream: stream})
}

type BIReportGeneratorService_ListReportLogsStreamServer interface {
	Send(*ListReportLogsStreamResponse) error
	grpc.ServerStream
}

type bIReportGeneratorServiceListReportLogsStreamServer struct {
	grpc.ServerStream
}

func (x *bIReportGeneratorServiceListReportLogsStreamServer) Send(m *ListReportLogsStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// BIReportGeneratorService_ServiceDesc is the grpc.ServiceDesc for BIReportGeneratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BIReportGeneratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.bireportgenerator.BIReportGeneratorService",
	HandlerType: (*BIReportGeneratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReportJob",
			Handler:    _BIReportGeneratorService_CreateReportJob_Handler,
		},
		{
			MethodName: "ListReportJobs",
			Handler:    _BIReportGeneratorService_ListReportJobs_Handler,
		},
		{
			MethodName: "UpdateReportJob",
			Handler:    _BIReportGeneratorService_UpdateReportJob_Handler,
		},
		{
			MethodName: "DeleteReportJob",
			Handler:    _BIReportGeneratorService_DeleteReportJob_Handler,
		},
		{
			MethodName: "GetReportJob",
			Handler:    _BIReportGeneratorService_GetReportJob_Handler,
		},
		{
			MethodName: "GenerateReport",
			Handler:    _BIReportGeneratorService_GenerateReport_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListReportLogsStream",
			Handler:       _BIReportGeneratorService_ListReportLogsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1alpha1/bireportgenerator/service.proto",
}
