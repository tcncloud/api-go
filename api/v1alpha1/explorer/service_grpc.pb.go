// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/v1alpha1/explorer/service.proto

package explorer

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
	ExplorerService_ListDatasourceSchemas_FullMethodName = "/api.v1alpha1.explorer.ExplorerService/ListDatasourceSchemas"
	ExplorerService_Query_FullMethodName                 = "/api.v1alpha1.explorer.ExplorerService/Query"
	ExplorerService_GetSupportQuery_FullMethodName       = "/api.v1alpha1.explorer.ExplorerService/GetSupportQuery"
	ExplorerService_GetQueryExplain_FullMethodName       = "/api.v1alpha1.explorer.ExplorerService/GetQueryExplain"
	ExplorerService_GetWeeksOfData_FullMethodName        = "/api.v1alpha1.explorer.ExplorerService/GetWeeksOfData"
)

// ExplorerServiceClient is the client API for ExplorerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ExplorerService is the service for the explorer API.
type ExplorerServiceClient interface {
	// ListDatasourceSchemas lists all accessible datasources and their schemas.
	ListDatasourceSchemas(ctx context.Context, in *ListDatasourceSchemasRequest, opts ...grpc.CallOption) (*ListDatasourceSchemasResponse, error)
	// Query queries a datasource.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	GetSupportQuery(ctx context.Context, in *SupportQueryRequest, opts ...grpc.CallOption) (*SupportQueryResponse, error)
	GetQueryExplain(ctx context.Context, in *QueryExplainRequest, opts ...grpc.CallOption) (*QueryExplainResponse, error)
	// GetWeeksOfData returns the number of weeks of data an org is limited to and the cutoff date.
	GetWeeksOfData(ctx context.Context, in *GetWeeksOfDataRequest, opts ...grpc.CallOption) (*GetWeeksOfDataResponse, error)
}

type explorerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExplorerServiceClient(cc grpc.ClientConnInterface) ExplorerServiceClient {
	return &explorerServiceClient{cc}
}

func (c *explorerServiceClient) ListDatasourceSchemas(ctx context.Context, in *ListDatasourceSchemasRequest, opts ...grpc.CallOption) (*ListDatasourceSchemasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDatasourceSchemasResponse)
	err := c.cc.Invoke(ctx, ExplorerService_ListDatasourceSchemas_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, ExplorerService_Query_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) GetSupportQuery(ctx context.Context, in *SupportQueryRequest, opts ...grpc.CallOption) (*SupportQueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SupportQueryResponse)
	err := c.cc.Invoke(ctx, ExplorerService_GetSupportQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) GetQueryExplain(ctx context.Context, in *QueryExplainRequest, opts ...grpc.CallOption) (*QueryExplainResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryExplainResponse)
	err := c.cc.Invoke(ctx, ExplorerService_GetQueryExplain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) GetWeeksOfData(ctx context.Context, in *GetWeeksOfDataRequest, opts ...grpc.CallOption) (*GetWeeksOfDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWeeksOfDataResponse)
	err := c.cc.Invoke(ctx, ExplorerService_GetWeeksOfData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExplorerServiceServer is the server API for ExplorerService service.
// All implementations must embed UnimplementedExplorerServiceServer
// for forward compatibility.
//
// ExplorerService is the service for the explorer API.
type ExplorerServiceServer interface {
	// ListDatasourceSchemas lists all accessible datasources and their schemas.
	ListDatasourceSchemas(context.Context, *ListDatasourceSchemasRequest) (*ListDatasourceSchemasResponse, error)
	// Query queries a datasource.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	GetSupportQuery(context.Context, *SupportQueryRequest) (*SupportQueryResponse, error)
	GetQueryExplain(context.Context, *QueryExplainRequest) (*QueryExplainResponse, error)
	// GetWeeksOfData returns the number of weeks of data an org is limited to and the cutoff date.
	GetWeeksOfData(context.Context, *GetWeeksOfDataRequest) (*GetWeeksOfDataResponse, error)
	mustEmbedUnimplementedExplorerServiceServer()
}

// UnimplementedExplorerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExplorerServiceServer struct{}

func (UnimplementedExplorerServiceServer) ListDatasourceSchemas(context.Context, *ListDatasourceSchemasRequest) (*ListDatasourceSchemasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatasourceSchemas not implemented")
}
func (UnimplementedExplorerServiceServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedExplorerServiceServer) GetSupportQuery(context.Context, *SupportQueryRequest) (*SupportQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportQuery not implemented")
}
func (UnimplementedExplorerServiceServer) GetQueryExplain(context.Context, *QueryExplainRequest) (*QueryExplainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryExplain not implemented")
}
func (UnimplementedExplorerServiceServer) GetWeeksOfData(context.Context, *GetWeeksOfDataRequest) (*GetWeeksOfDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeeksOfData not implemented")
}
func (UnimplementedExplorerServiceServer) mustEmbedUnimplementedExplorerServiceServer() {}
func (UnimplementedExplorerServiceServer) testEmbeddedByValue()                         {}

// UnsafeExplorerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExplorerServiceServer will
// result in compilation errors.
type UnsafeExplorerServiceServer interface {
	mustEmbedUnimplementedExplorerServiceServer()
}

func RegisterExplorerServiceServer(s grpc.ServiceRegistrar, srv ExplorerServiceServer) {
	// If the following call pancis, it indicates UnimplementedExplorerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExplorerService_ServiceDesc, srv)
}

func _ExplorerService_ListDatasourceSchemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatasourceSchemasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).ListDatasourceSchemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExplorerService_ListDatasourceSchemas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).ListDatasourceSchemas(ctx, req.(*ListDatasourceSchemasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExplorerService_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_GetSupportQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetSupportQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExplorerService_GetSupportQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetSupportQuery(ctx, req.(*SupportQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_GetQueryExplain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryExplainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetQueryExplain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExplorerService_GetQueryExplain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetQueryExplain(ctx, req.(*QueryExplainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_GetWeeksOfData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeeksOfDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetWeeksOfData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExplorerService_GetWeeksOfData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetWeeksOfData(ctx, req.(*GetWeeksOfDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExplorerService_ServiceDesc is the grpc.ServiceDesc for ExplorerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExplorerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.explorer.ExplorerService",
	HandlerType: (*ExplorerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDatasourceSchemas",
			Handler:    _ExplorerService_ListDatasourceSchemas_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _ExplorerService_Query_Handler,
		},
		{
			MethodName: "GetSupportQuery",
			Handler:    _ExplorerService_GetSupportQuery_Handler,
		},
		{
			MethodName: "GetQueryExplain",
			Handler:    _ExplorerService_GetQueryExplain_Handler,
		},
		{
			MethodName: "GetWeeksOfData",
			Handler:    _ExplorerService_GetWeeksOfData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1alpha1/explorer/service.proto",
}
