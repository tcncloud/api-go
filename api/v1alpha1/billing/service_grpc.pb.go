// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v1alpha1/billing/service.proto

package billing

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
	Billing_GetBillingPlan_FullMethodName         = "/api.v1alpha1.billing.Billing/GetBillingPlan"
	Billing_UpdateBillingPlan_FullMethodName      = "/api.v1alpha1.billing.Billing/UpdateBillingPlan"
	Billing_GetInvoice_FullMethodName             = "/api.v1alpha1.billing.Billing/GetInvoice"
	Billing_ExportGeneratedInvoice_FullMethodName = "/api.v1alpha1.billing.Billing/ExportGeneratedInvoice"
)

// BillingClient is the client API for Billing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillingClient interface {
	// GetBillingPlan - returns the billing plan for the provided organization.
	GetBillingPlan(ctx context.Context, in *GetBillingPlanReq, opts ...grpc.CallOption) (*GetBillingPlanRes, error)
	// UpdateBillingPlan - updates the provided billing plan and it's details.
	// If some details are not provided, they will be left as is. However, if
	// deletion is desired, the DeleteBillingDetails method should be used. The
	// billing plan still follows the constraint of only having one billing detail
	// with a specific config type and event type, and so if the request contains
	// more than one billing detail with a config type and event type, the request
	// is malformed and will result in potentially unexpected behavior.
	UpdateBillingPlan(ctx context.Context, in *UpdateBillingPlanReq, opts ...grpc.CallOption) (*UpdateBillingPlanRes, error)
	// GetInvoice - returns the invoice for the organization.
	// If a date is provided, this will return the invoice for the
	// organization that corresponds to the billing cycle that contains
	// the provided date. If no date is provided, this will return the
	// invoice as it currently stands for the current billing cycle.
	GetInvoice(ctx context.Context, in *GetInvoiceReq, opts ...grpc.CallOption) (*GetInvoiceRes, error)
	// ExportGeneratedInvoice - returns the invoice for the organization.
	// If a date is provided, this will return the invoice for the
	// organization that corresponds to the billing cycle that contains
	// the provided date. If no date is provided, this will return the
	// invoice, as it has been last generated, for the current billing cycle.
	// This differs from GetInvoice in that it returns the invoice as
	// it was last generated. It will not take into account new billing
	// events since the last generation.
	ExportGeneratedInvoice(ctx context.Context, in *ExportGeneratedInvoiceReq, opts ...grpc.CallOption) (*ExportGeneratedInvoiceRes, error)
}

type billingClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingClient(cc grpc.ClientConnInterface) BillingClient {
	return &billingClient{cc}
}

func (c *billingClient) GetBillingPlan(ctx context.Context, in *GetBillingPlanReq, opts ...grpc.CallOption) (*GetBillingPlanRes, error) {
	out := new(GetBillingPlanRes)
	err := c.cc.Invoke(ctx, Billing_GetBillingPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) UpdateBillingPlan(ctx context.Context, in *UpdateBillingPlanReq, opts ...grpc.CallOption) (*UpdateBillingPlanRes, error) {
	out := new(UpdateBillingPlanRes)
	err := c.cc.Invoke(ctx, Billing_UpdateBillingPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) GetInvoice(ctx context.Context, in *GetInvoiceReq, opts ...grpc.CallOption) (*GetInvoiceRes, error) {
	out := new(GetInvoiceRes)
	err := c.cc.Invoke(ctx, Billing_GetInvoice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) ExportGeneratedInvoice(ctx context.Context, in *ExportGeneratedInvoiceReq, opts ...grpc.CallOption) (*ExportGeneratedInvoiceRes, error) {
	out := new(ExportGeneratedInvoiceRes)
	err := c.cc.Invoke(ctx, Billing_ExportGeneratedInvoice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingServer is the server API for Billing service.
// All implementations must embed UnimplementedBillingServer
// for forward compatibility
type BillingServer interface {
	// GetBillingPlan - returns the billing plan for the provided organization.
	GetBillingPlan(context.Context, *GetBillingPlanReq) (*GetBillingPlanRes, error)
	// UpdateBillingPlan - updates the provided billing plan and it's details.
	// If some details are not provided, they will be left as is. However, if
	// deletion is desired, the DeleteBillingDetails method should be used. The
	// billing plan still follows the constraint of only having one billing detail
	// with a specific config type and event type, and so if the request contains
	// more than one billing detail with a config type and event type, the request
	// is malformed and will result in potentially unexpected behavior.
	UpdateBillingPlan(context.Context, *UpdateBillingPlanReq) (*UpdateBillingPlanRes, error)
	// GetInvoice - returns the invoice for the organization.
	// If a date is provided, this will return the invoice for the
	// organization that corresponds to the billing cycle that contains
	// the provided date. If no date is provided, this will return the
	// invoice as it currently stands for the current billing cycle.
	GetInvoice(context.Context, *GetInvoiceReq) (*GetInvoiceRes, error)
	// ExportGeneratedInvoice - returns the invoice for the organization.
	// If a date is provided, this will return the invoice for the
	// organization that corresponds to the billing cycle that contains
	// the provided date. If no date is provided, this will return the
	// invoice, as it has been last generated, for the current billing cycle.
	// This differs from GetInvoice in that it returns the invoice as
	// it was last generated. It will not take into account new billing
	// events since the last generation.
	ExportGeneratedInvoice(context.Context, *ExportGeneratedInvoiceReq) (*ExportGeneratedInvoiceRes, error)
	mustEmbedUnimplementedBillingServer()
}

// UnimplementedBillingServer must be embedded to have forward compatible implementations.
type UnimplementedBillingServer struct {
}

func (UnimplementedBillingServer) GetBillingPlan(context.Context, *GetBillingPlanReq) (*GetBillingPlanRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBillingPlan not implemented")
}
func (UnimplementedBillingServer) UpdateBillingPlan(context.Context, *UpdateBillingPlanReq) (*UpdateBillingPlanRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBillingPlan not implemented")
}
func (UnimplementedBillingServer) GetInvoice(context.Context, *GetInvoiceReq) (*GetInvoiceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedBillingServer) ExportGeneratedInvoice(context.Context, *ExportGeneratedInvoiceReq) (*ExportGeneratedInvoiceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportGeneratedInvoice not implemented")
}
func (UnimplementedBillingServer) mustEmbedUnimplementedBillingServer() {}

// UnsafeBillingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillingServer will
// result in compilation errors.
type UnsafeBillingServer interface {
	mustEmbedUnimplementedBillingServer()
}

func RegisterBillingServer(s grpc.ServiceRegistrar, srv BillingServer) {
	s.RegisterService(&Billing_ServiceDesc, srv)
}

func _Billing_GetBillingPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillingPlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).GetBillingPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Billing_GetBillingPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).GetBillingPlan(ctx, req.(*GetBillingPlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_UpdateBillingPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBillingPlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).UpdateBillingPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Billing_UpdateBillingPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).UpdateBillingPlan(ctx, req.(*UpdateBillingPlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Billing_GetInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).GetInvoice(ctx, req.(*GetInvoiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_ExportGeneratedInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportGeneratedInvoiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).ExportGeneratedInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Billing_ExportGeneratedInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).ExportGeneratedInvoice(ctx, req.(*ExportGeneratedInvoiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Billing_ServiceDesc is the grpc.ServiceDesc for Billing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Billing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.billing.Billing",
	HandlerType: (*BillingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBillingPlan",
			Handler:    _Billing_GetBillingPlan_Handler,
		},
		{
			MethodName: "UpdateBillingPlan",
			Handler:    _Billing_UpdateBillingPlan_Handler,
		},
		{
			MethodName: "GetInvoice",
			Handler:    _Billing_GetInvoice_Handler,
		},
		{
			MethodName: "ExportGeneratedInvoice",
			Handler:    _Billing_ExportGeneratedInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1alpha1/billing/service.proto",
}
