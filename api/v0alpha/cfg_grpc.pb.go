// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: api/v0alpha/cfg.proto

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
	Cfg_GetWebAgentConfig_FullMethodName = "/api.v0alpha.Cfg/GetWebAgentConfig"
)

// CfgClient is the client API for Cfg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CfgClient interface {
	GetWebAgentConfig(ctx context.Context, in *GetConfigReq, opts ...grpc.CallOption) (*WebAgent, error)
}

type cfgClient struct {
	cc grpc.ClientConnInterface
}

func NewCfgClient(cc grpc.ClientConnInterface) CfgClient {
	return &cfgClient{cc}
}

func (c *cfgClient) GetWebAgentConfig(ctx context.Context, in *GetConfigReq, opts ...grpc.CallOption) (*WebAgent, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebAgent)
	err := c.cc.Invoke(ctx, Cfg_GetWebAgentConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CfgServer is the server API for Cfg service.
// All implementations must embed UnimplementedCfgServer
// for forward compatibility
type CfgServer interface {
	GetWebAgentConfig(context.Context, *GetConfigReq) (*WebAgent, error)
	mustEmbedUnimplementedCfgServer()
}

// UnimplementedCfgServer must be embedded to have forward compatible implementations.
type UnimplementedCfgServer struct {
}

func (UnimplementedCfgServer) GetWebAgentConfig(context.Context, *GetConfigReq) (*WebAgent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebAgentConfig not implemented")
}
func (UnimplementedCfgServer) mustEmbedUnimplementedCfgServer() {}

// UnsafeCfgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CfgServer will
// result in compilation errors.
type UnsafeCfgServer interface {
	mustEmbedUnimplementedCfgServer()
}

func RegisterCfgServer(s grpc.ServiceRegistrar, srv CfgServer) {
	s.RegisterService(&Cfg_ServiceDesc, srv)
}

func _Cfg_GetWebAgentConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CfgServer).GetWebAgentConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cfg_GetWebAgentConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CfgServer).GetWebAgentConfig(ctx, req.(*GetConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Cfg_ServiceDesc is the grpc.ServiceDesc for Cfg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cfg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v0alpha.Cfg",
	HandlerType: (*CfgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWebAgentConfig",
			Handler:    _Cfg_GetWebAgentConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v0alpha/cfg.proto",
}
