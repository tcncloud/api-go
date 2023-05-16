// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v1alpha1/idp/service.proto

package idp

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const ()

// IdentityProviderClient is the client API for IdentityProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityProviderClient interface {
}

type identityProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityProviderClient(cc grpc.ClientConnInterface) IdentityProviderClient {
	return &identityProviderClient{cc}
}

// IdentityProviderServer is the server API for IdentityProvider service.
// All implementations must embed UnimplementedIdentityProviderServer
// for forward compatibility
type IdentityProviderServer interface {
	mustEmbedUnimplementedIdentityProviderServer()
}

// UnimplementedIdentityProviderServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityProviderServer struct {
}

func (UnimplementedIdentityProviderServer) mustEmbedUnimplementedIdentityProviderServer() {}

// UnsafeIdentityProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityProviderServer will
// result in compilation errors.
type UnsafeIdentityProviderServer interface {
	mustEmbedUnimplementedIdentityProviderServer()
}

func RegisterIdentityProviderServer(s grpc.ServiceRegistrar, srv IdentityProviderServer) {
	s.RegisterService(&IdentityProvider_ServiceDesc, srv)
}

// IdentityProvider_ServiceDesc is the grpc.ServiceDesc for IdentityProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.idp.IdentityProvider",
	HandlerType: (*IdentityProviderServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "api/v1alpha1/idp/service.proto",
}
