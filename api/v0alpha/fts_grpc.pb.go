// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: api/v0alpha/fts.proto

package v0alpha

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
	FTS_GetUploadFileUrl_FullMethodName = "/api.v0alpha.FTS/GetUploadFileUrl"
)

// FTSClient is the client API for FTS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FTSClient interface {
	GetUploadFileUrl(ctx context.Context, in *GetUploadFileUrlReq, opts ...grpc.CallOption) (*GetUploadFileUrlRes, error)
}

type fTSClient struct {
	cc grpc.ClientConnInterface
}

func NewFTSClient(cc grpc.ClientConnInterface) FTSClient {
	return &fTSClient{cc}
}

func (c *fTSClient) GetUploadFileUrl(ctx context.Context, in *GetUploadFileUrlReq, opts ...grpc.CallOption) (*GetUploadFileUrlRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUploadFileUrlRes)
	err := c.cc.Invoke(ctx, FTS_GetUploadFileUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FTSServer is the server API for FTS service.
// All implementations must embed UnimplementedFTSServer
// for forward compatibility
type FTSServer interface {
	GetUploadFileUrl(context.Context, *GetUploadFileUrlReq) (*GetUploadFileUrlRes, error)
	mustEmbedUnimplementedFTSServer()
}

// UnimplementedFTSServer must be embedded to have forward compatible implementations.
type UnimplementedFTSServer struct {
}

func (UnimplementedFTSServer) GetUploadFileUrl(context.Context, *GetUploadFileUrlReq) (*GetUploadFileUrlRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUploadFileUrl not implemented")
}
func (UnimplementedFTSServer) mustEmbedUnimplementedFTSServer() {}

// UnsafeFTSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FTSServer will
// result in compilation errors.
type UnsafeFTSServer interface {
	mustEmbedUnimplementedFTSServer()
}

func RegisterFTSServer(s grpc.ServiceRegistrar, srv FTSServer) {
	s.RegisterService(&FTS_ServiceDesc, srv)
}

func _FTS_GetUploadFileUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUploadFileUrlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FTSServer).GetUploadFileUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FTS_GetUploadFileUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FTSServer).GetUploadFileUrl(ctx, req.(*GetUploadFileUrlReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FTS_ServiceDesc is the grpc.ServiceDesc for FTS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FTS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v0alpha.FTS",
	HandlerType: (*FTSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUploadFileUrl",
			Handler:    _FTS_GetUploadFileUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v0alpha/fts.proto",
}
