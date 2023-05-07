// Copyright (c) 2019, TCN Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v0alpha/sentinel.proto

package v0alpha

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
	Sentinel_SendEvents_FullMethodName = "/api.v0alpha.Sentinel/SendEvents"
)

// SentinelClient is the client API for Sentinel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SentinelClient interface {
	// Send a json blob of ui events and logs.
	SendEvents(ctx context.Context, in *SendEventsReq, opts ...grpc.CallOption) (*SendEventsRes, error)
}

type sentinelClient struct {
	cc grpc.ClientConnInterface
}

func NewSentinelClient(cc grpc.ClientConnInterface) SentinelClient {
	return &sentinelClient{cc}
}

func (c *sentinelClient) SendEvents(ctx context.Context, in *SendEventsReq, opts ...grpc.CallOption) (*SendEventsRes, error) {
	out := new(SendEventsRes)
	err := c.cc.Invoke(ctx, Sentinel_SendEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SentinelServer is the server API for Sentinel service.
// All implementations must embed UnimplementedSentinelServer
// for forward compatibility
type SentinelServer interface {
	// Send a json blob of ui events and logs.
	SendEvents(context.Context, *SendEventsReq) (*SendEventsRes, error)
	mustEmbedUnimplementedSentinelServer()
}

// UnimplementedSentinelServer must be embedded to have forward compatible implementations.
type UnimplementedSentinelServer struct {
}

func (UnimplementedSentinelServer) SendEvents(context.Context, *SendEventsReq) (*SendEventsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEvents not implemented")
}
func (UnimplementedSentinelServer) mustEmbedUnimplementedSentinelServer() {}

// UnsafeSentinelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SentinelServer will
// result in compilation errors.
type UnsafeSentinelServer interface {
	mustEmbedUnimplementedSentinelServer()
}

func RegisterSentinelServer(s grpc.ServiceRegistrar, srv SentinelServer) {
	s.RegisterService(&Sentinel_ServiceDesc, srv)
}

func _Sentinel_SendEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEventsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentinelServer).SendEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sentinel_SendEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentinelServer).SendEvents(ctx, req.(*SendEventsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Sentinel_ServiceDesc is the grpc.ServiceDesc for Sentinel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sentinel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v0alpha.Sentinel",
	HandlerType: (*SentinelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEvents",
			Handler:    _Sentinel_SendEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v0alpha/sentinel.proto",
}
