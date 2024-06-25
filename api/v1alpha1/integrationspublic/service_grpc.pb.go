// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: api/v1alpha1/integrationspublic/service.proto

package integrationspublic

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
	IntegrationsPublic_GetLinkData_FullMethodName        = "/api.v1alpha1.integrationspublic.IntegrationsPublic/GetLinkData"
	IntegrationsPublic_SubmitVerification_FullMethodName = "/api.v1alpha1.integrationspublic.IntegrationsPublic/SubmitVerification"
	IntegrationsPublic_SessionKeepAlive_FullMethodName   = "/api.v1alpha1.integrationspublic.IntegrationsPublic/SessionKeepAlive"
	IntegrationsPublic_GetInvoice_FullMethodName         = "/api.v1alpha1.integrationspublic.IntegrationsPublic/GetInvoice"
	IntegrationsPublic_SubmitPayment_FullMethodName      = "/api.v1alpha1.integrationspublic.IntegrationsPublic/SubmitPayment"
	IntegrationsPublic_GetReceipt_FullMethodName         = "/api.v1alpha1.integrationspublic.IntegrationsPublic/GetReceipt"
)

// IntegrationsPublicClient is the client API for IntegrationsPublic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntegrationsPublicClient interface {
	GetLinkData(ctx context.Context, in *GetLinkDataReq, opts ...grpc.CallOption) (*GetLinkDataRes, error)
	SubmitVerification(ctx context.Context, in *SubmitVerificationReq, opts ...grpc.CallOption) (*SubmitVerificationRes, error)
	SessionKeepAlive(ctx context.Context, in *SessionKeepAliveReq, opts ...grpc.CallOption) (*SessionKeepAliveRes, error)
	GetInvoice(ctx context.Context, in *GetInvoiceReq, opts ...grpc.CallOption) (*GetInvoiceRes, error)
	SubmitPayment(ctx context.Context, in *SubmitPaymentReq, opts ...grpc.CallOption) (*SubmitPaymentRes, error)
	GetReceipt(ctx context.Context, in *GetReceiptReq, opts ...grpc.CallOption) (*GetReceiptRes, error)
}

type integrationsPublicClient struct {
	cc grpc.ClientConnInterface
}

func NewIntegrationsPublicClient(cc grpc.ClientConnInterface) IntegrationsPublicClient {
	return &integrationsPublicClient{cc}
}

func (c *integrationsPublicClient) GetLinkData(ctx context.Context, in *GetLinkDataReq, opts ...grpc.CallOption) (*GetLinkDataRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLinkDataRes)
	err := c.cc.Invoke(ctx, IntegrationsPublic_GetLinkData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationsPublicClient) SubmitVerification(ctx context.Context, in *SubmitVerificationReq, opts ...grpc.CallOption) (*SubmitVerificationRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitVerificationRes)
	err := c.cc.Invoke(ctx, IntegrationsPublic_SubmitVerification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationsPublicClient) SessionKeepAlive(ctx context.Context, in *SessionKeepAliveReq, opts ...grpc.CallOption) (*SessionKeepAliveRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionKeepAliveRes)
	err := c.cc.Invoke(ctx, IntegrationsPublic_SessionKeepAlive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationsPublicClient) GetInvoice(ctx context.Context, in *GetInvoiceReq, opts ...grpc.CallOption) (*GetInvoiceRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInvoiceRes)
	err := c.cc.Invoke(ctx, IntegrationsPublic_GetInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationsPublicClient) SubmitPayment(ctx context.Context, in *SubmitPaymentReq, opts ...grpc.CallOption) (*SubmitPaymentRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitPaymentRes)
	err := c.cc.Invoke(ctx, IntegrationsPublic_SubmitPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationsPublicClient) GetReceipt(ctx context.Context, in *GetReceiptReq, opts ...grpc.CallOption) (*GetReceiptRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReceiptRes)
	err := c.cc.Invoke(ctx, IntegrationsPublic_GetReceipt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntegrationsPublicServer is the server API for IntegrationsPublic service.
// All implementations must embed UnimplementedIntegrationsPublicServer
// for forward compatibility
type IntegrationsPublicServer interface {
	GetLinkData(context.Context, *GetLinkDataReq) (*GetLinkDataRes, error)
	SubmitVerification(context.Context, *SubmitVerificationReq) (*SubmitVerificationRes, error)
	SessionKeepAlive(context.Context, *SessionKeepAliveReq) (*SessionKeepAliveRes, error)
	GetInvoice(context.Context, *GetInvoiceReq) (*GetInvoiceRes, error)
	SubmitPayment(context.Context, *SubmitPaymentReq) (*SubmitPaymentRes, error)
	GetReceipt(context.Context, *GetReceiptReq) (*GetReceiptRes, error)
	mustEmbedUnimplementedIntegrationsPublicServer()
}

// UnimplementedIntegrationsPublicServer must be embedded to have forward compatible implementations.
type UnimplementedIntegrationsPublicServer struct {
}

func (UnimplementedIntegrationsPublicServer) GetLinkData(context.Context, *GetLinkDataReq) (*GetLinkDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkData not implemented")
}
func (UnimplementedIntegrationsPublicServer) SubmitVerification(context.Context, *SubmitVerificationReq) (*SubmitVerificationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitVerification not implemented")
}
func (UnimplementedIntegrationsPublicServer) SessionKeepAlive(context.Context, *SessionKeepAliveReq) (*SessionKeepAliveRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionKeepAlive not implemented")
}
func (UnimplementedIntegrationsPublicServer) GetInvoice(context.Context, *GetInvoiceReq) (*GetInvoiceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedIntegrationsPublicServer) SubmitPayment(context.Context, *SubmitPaymentReq) (*SubmitPaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitPayment not implemented")
}
func (UnimplementedIntegrationsPublicServer) GetReceipt(context.Context, *GetReceiptReq) (*GetReceiptRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedIntegrationsPublicServer) mustEmbedUnimplementedIntegrationsPublicServer() {}

// UnsafeIntegrationsPublicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntegrationsPublicServer will
// result in compilation errors.
type UnsafeIntegrationsPublicServer interface {
	mustEmbedUnimplementedIntegrationsPublicServer()
}

func RegisterIntegrationsPublicServer(s grpc.ServiceRegistrar, srv IntegrationsPublicServer) {
	s.RegisterService(&IntegrationsPublic_ServiceDesc, srv)
}

func _IntegrationsPublic_GetLinkData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsPublicServer).GetLinkData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsPublic_GetLinkData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsPublicServer).GetLinkData(ctx, req.(*GetLinkDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationsPublic_SubmitVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitVerificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsPublicServer).SubmitVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsPublic_SubmitVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsPublicServer).SubmitVerification(ctx, req.(*SubmitVerificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationsPublic_SessionKeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionKeepAliveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsPublicServer).SessionKeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsPublic_SessionKeepAlive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsPublicServer).SessionKeepAlive(ctx, req.(*SessionKeepAliveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationsPublic_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsPublicServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsPublic_GetInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsPublicServer).GetInvoice(ctx, req.(*GetInvoiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationsPublic_SubmitPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitPaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsPublicServer).SubmitPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsPublic_SubmitPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsPublicServer).SubmitPayment(ctx, req.(*SubmitPaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationsPublic_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReceiptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsPublicServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsPublic_GetReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsPublicServer).GetReceipt(ctx, req.(*GetReceiptReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IntegrationsPublic_ServiceDesc is the grpc.ServiceDesc for IntegrationsPublic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntegrationsPublic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.integrationspublic.IntegrationsPublic",
	HandlerType: (*IntegrationsPublicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLinkData",
			Handler:    _IntegrationsPublic_GetLinkData_Handler,
		},
		{
			MethodName: "SubmitVerification",
			Handler:    _IntegrationsPublic_SubmitVerification_Handler,
		},
		{
			MethodName: "SessionKeepAlive",
			Handler:    _IntegrationsPublic_SessionKeepAlive_Handler,
		},
		{
			MethodName: "GetInvoice",
			Handler:    _IntegrationsPublic_GetInvoice_Handler,
		},
		{
			MethodName: "SubmitPayment",
			Handler:    _IntegrationsPublic_SubmitPayment_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _IntegrationsPublic_GetReceipt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1alpha1/integrationspublic/service.proto",
}
