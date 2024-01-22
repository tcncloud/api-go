// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: wfo/vanalytics/v2/service.proto

package vanalyticsv2

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
	Vanalytics_SearchTranscripts_FullMethodName = "/wfo.vanalytics.v2.Vanalytics/SearchTranscripts"
)

// VanalyticsClient is the client API for Vanalytics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VanalyticsClient interface {
	// SearchTranscripts searches transcripts by search criteria. The search response
	// contains one page of transcript hits. Traversing the paginated hits is
	// achieved by making use of the given page token.
	SearchTranscripts(ctx context.Context, in *SearchTranscriptsRequest, opts ...grpc.CallOption) (*SearchTranscriptsResponse, error)
}

type vanalyticsClient struct {
	cc grpc.ClientConnInterface
}

func NewVanalyticsClient(cc grpc.ClientConnInterface) VanalyticsClient {
	return &vanalyticsClient{cc}
}

func (c *vanalyticsClient) SearchTranscripts(ctx context.Context, in *SearchTranscriptsRequest, opts ...grpc.CallOption) (*SearchTranscriptsResponse, error) {
	out := new(SearchTranscriptsResponse)
	err := c.cc.Invoke(ctx, Vanalytics_SearchTranscripts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VanalyticsServer is the server API for Vanalytics service.
// All implementations must embed UnimplementedVanalyticsServer
// for forward compatibility
type VanalyticsServer interface {
	// SearchTranscripts searches transcripts by search criteria. The search response
	// contains one page of transcript hits. Traversing the paginated hits is
	// achieved by making use of the given page token.
	SearchTranscripts(context.Context, *SearchTranscriptsRequest) (*SearchTranscriptsResponse, error)
	mustEmbedUnimplementedVanalyticsServer()
}

// UnimplementedVanalyticsServer must be embedded to have forward compatible implementations.
type UnimplementedVanalyticsServer struct {
}

func (UnimplementedVanalyticsServer) SearchTranscripts(context.Context, *SearchTranscriptsRequest) (*SearchTranscriptsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTranscripts not implemented")
}
func (UnimplementedVanalyticsServer) mustEmbedUnimplementedVanalyticsServer() {}

// UnsafeVanalyticsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VanalyticsServer will
// result in compilation errors.
type UnsafeVanalyticsServer interface {
	mustEmbedUnimplementedVanalyticsServer()
}

func RegisterVanalyticsServer(s grpc.ServiceRegistrar, srv VanalyticsServer) {
	s.RegisterService(&Vanalytics_ServiceDesc, srv)
}

func _Vanalytics_SearchTranscripts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTranscriptsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VanalyticsServer).SearchTranscripts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vanalytics_SearchTranscripts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VanalyticsServer).SearchTranscripts(ctx, req.(*SearchTranscriptsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Vanalytics_ServiceDesc is the grpc.ServiceDesc for Vanalytics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Vanalytics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wfo.vanalytics.v2.Vanalytics",
	HandlerType: (*VanalyticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchTranscripts",
			Handler:    _Vanalytics_SearchTranscripts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wfo/vanalytics/v2/service.proto",
}