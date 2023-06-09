// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v1alpha1/tickets/service.proto

package tickets

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
	Tickets_CreateTicket_FullMethodName        = "/api.v1alpha1.tickets.Tickets/CreateTicket"
	Tickets_EditTicket_FullMethodName          = "/api.v1alpha1.tickets.Tickets/EditTicket"
	Tickets_ListTickets_FullMethodName         = "/api.v1alpha1.tickets.Tickets/ListTickets"
	Tickets_AssignTicket_FullMethodName        = "/api.v1alpha1.tickets.Tickets/AssignTicket"
	Tickets_CloseTicket_FullMethodName         = "/api.v1alpha1.tickets.Tickets/CloseTicket"
	Tickets_ViewTicket_FullMethodName          = "/api.v1alpha1.tickets.Tickets/ViewTicket"
	Tickets_CreateComment_FullMethodName       = "/api.v1alpha1.tickets.Tickets/CreateComment"
	Tickets_EnableProject_FullMethodName       = "/api.v1alpha1.tickets.Tickets/EnableProject"
	Tickets_ListEnabledProjects_FullMethodName = "/api.v1alpha1.tickets.Tickets/ListEnabledProjects"
	Tickets_CreateSLA_FullMethodName           = "/api.v1alpha1.tickets.Tickets/CreateSLA"
	Tickets_ListSLA_FullMethodName             = "/api.v1alpha1.tickets.Tickets/ListSLA"
	Tickets_UpdateSLA_FullMethodName           = "/api.v1alpha1.tickets.Tickets/UpdateSLA"
	Tickets_ListSLACondition_FullMethodName    = "/api.v1alpha1.tickets.Tickets/ListSLACondition"
	Tickets_ReplyComment_FullMethodName        = "/api.v1alpha1.tickets.Tickets/ReplyComment"
	Tickets_ListTicketAuditLog_FullMethodName  = "/api.v1alpha1.tickets.Tickets/ListTicketAuditLog"
	Tickets_AssignSelf_FullMethodName          = "/api.v1alpha1.tickets.Tickets/AssignSelf"
)

// TicketsClient is the client API for Tickets service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketsClient interface {
	// Public Method to create a ticket.
	CreateTicket(ctx context.Context, in *CreateTicketReq, opts ...grpc.CallOption) (*CreateTicketRes, error)
	// Public Method to create a ticket.
	EditTicket(ctx context.Context, in *EditTicketReq, opts ...grpc.CallOption) (*EditTicketRes, error)
	// Public method to list tickets
	ListTickets(ctx context.Context, in *ListTicketsReq, opts ...grpc.CallOption) (*ListTicketsRes, error)
	// Public method to assign ticket
	AssignTicket(ctx context.Context, in *AssignTicketReq, opts ...grpc.CallOption) (*AssignTicketRes, error)
	// Closes the ticket
	CloseTicket(ctx context.Context, in *CloseTicketReq, opts ...grpc.CallOption) (*CloseTicketRes, error)
	// Public method to view ticket
	ViewTicket(ctx context.Context, in *ViewTicketReq, opts ...grpc.CallOption) (*ViewTicketRes, error)
	// Public Method to create a Comment.
	CreateComment(ctx context.Context, in *CreateCommentReq, opts ...grpc.CallOption) (*CreateCommentRes, error)
	// Public method to Enable Project for Ticketing system
	EnableProject(ctx context.Context, in *EnableProjectReq, opts ...grpc.CallOption) (*EnableProjectRes, error)
	// Public method to List projects enabled for Ticketing system
	ListEnabledProjects(ctx context.Context, in *ListEnabledProjectsReq, opts ...grpc.CallOption) (*ListEnabledProjectsRes, error)
	// Public Method to create a sla.
	CreateSLA(ctx context.Context, in *CreateSlaReq, opts ...grpc.CallOption) (*CreateSlaRes, error)
	// Public method to list sla
	ListSLA(ctx context.Context, in *ListSlaReq, opts ...grpc.CallOption) (*ListSlaRes, error)
	// Public method to update sla
	UpdateSLA(ctx context.Context, in *UpdateSlaReq, opts ...grpc.CallOption) (*UpdateSlaRes, error)
	// Public method to list sla_condition
	ListSLACondition(ctx context.Context, in *ListSlaConditionReq, opts ...grpc.CallOption) (*ListSlaConditionRes, error)
	// Public method to list sla_condition
	ReplyComment(ctx context.Context, in *ReplyCommentReq, opts ...grpc.CallOption) (*ReplyCommentRes, error)
	// Public method to List audit log for Ticketing system
	ListTicketAuditLog(ctx context.Context, in *ListTicketAuditLogReq, opts ...grpc.CallOption) (*ListTicketAuditLogRes, error)
	// Public method to assign a ticket
	AssignSelf(ctx context.Context, in *CreateSelfAssignReq, opts ...grpc.CallOption) (*CreateSelfAssignRes, error)
}

type ticketsClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketsClient(cc grpc.ClientConnInterface) TicketsClient {
	return &ticketsClient{cc}
}

func (c *ticketsClient) CreateTicket(ctx context.Context, in *CreateTicketReq, opts ...grpc.CallOption) (*CreateTicketRes, error) {
	out := new(CreateTicketRes)
	err := c.cc.Invoke(ctx, Tickets_CreateTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) EditTicket(ctx context.Context, in *EditTicketReq, opts ...grpc.CallOption) (*EditTicketRes, error) {
	out := new(EditTicketRes)
	err := c.cc.Invoke(ctx, Tickets_EditTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) ListTickets(ctx context.Context, in *ListTicketsReq, opts ...grpc.CallOption) (*ListTicketsRes, error) {
	out := new(ListTicketsRes)
	err := c.cc.Invoke(ctx, Tickets_ListTickets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) AssignTicket(ctx context.Context, in *AssignTicketReq, opts ...grpc.CallOption) (*AssignTicketRes, error) {
	out := new(AssignTicketRes)
	err := c.cc.Invoke(ctx, Tickets_AssignTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) CloseTicket(ctx context.Context, in *CloseTicketReq, opts ...grpc.CallOption) (*CloseTicketRes, error) {
	out := new(CloseTicketRes)
	err := c.cc.Invoke(ctx, Tickets_CloseTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) ViewTicket(ctx context.Context, in *ViewTicketReq, opts ...grpc.CallOption) (*ViewTicketRes, error) {
	out := new(ViewTicketRes)
	err := c.cc.Invoke(ctx, Tickets_ViewTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) CreateComment(ctx context.Context, in *CreateCommentReq, opts ...grpc.CallOption) (*CreateCommentRes, error) {
	out := new(CreateCommentRes)
	err := c.cc.Invoke(ctx, Tickets_CreateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) EnableProject(ctx context.Context, in *EnableProjectReq, opts ...grpc.CallOption) (*EnableProjectRes, error) {
	out := new(EnableProjectRes)
	err := c.cc.Invoke(ctx, Tickets_EnableProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) ListEnabledProjects(ctx context.Context, in *ListEnabledProjectsReq, opts ...grpc.CallOption) (*ListEnabledProjectsRes, error) {
	out := new(ListEnabledProjectsRes)
	err := c.cc.Invoke(ctx, Tickets_ListEnabledProjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) CreateSLA(ctx context.Context, in *CreateSlaReq, opts ...grpc.CallOption) (*CreateSlaRes, error) {
	out := new(CreateSlaRes)
	err := c.cc.Invoke(ctx, Tickets_CreateSLA_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) ListSLA(ctx context.Context, in *ListSlaReq, opts ...grpc.CallOption) (*ListSlaRes, error) {
	out := new(ListSlaRes)
	err := c.cc.Invoke(ctx, Tickets_ListSLA_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) UpdateSLA(ctx context.Context, in *UpdateSlaReq, opts ...grpc.CallOption) (*UpdateSlaRes, error) {
	out := new(UpdateSlaRes)
	err := c.cc.Invoke(ctx, Tickets_UpdateSLA_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) ListSLACondition(ctx context.Context, in *ListSlaConditionReq, opts ...grpc.CallOption) (*ListSlaConditionRes, error) {
	out := new(ListSlaConditionRes)
	err := c.cc.Invoke(ctx, Tickets_ListSLACondition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) ReplyComment(ctx context.Context, in *ReplyCommentReq, opts ...grpc.CallOption) (*ReplyCommentRes, error) {
	out := new(ReplyCommentRes)
	err := c.cc.Invoke(ctx, Tickets_ReplyComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) ListTicketAuditLog(ctx context.Context, in *ListTicketAuditLogReq, opts ...grpc.CallOption) (*ListTicketAuditLogRes, error) {
	out := new(ListTicketAuditLogRes)
	err := c.cc.Invoke(ctx, Tickets_ListTicketAuditLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) AssignSelf(ctx context.Context, in *CreateSelfAssignReq, opts ...grpc.CallOption) (*CreateSelfAssignRes, error) {
	out := new(CreateSelfAssignRes)
	err := c.cc.Invoke(ctx, Tickets_AssignSelf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketsServer is the server API for Tickets service.
// All implementations must embed UnimplementedTicketsServer
// for forward compatibility
type TicketsServer interface {
	// Public Method to create a ticket.
	CreateTicket(context.Context, *CreateTicketReq) (*CreateTicketRes, error)
	// Public Method to create a ticket.
	EditTicket(context.Context, *EditTicketReq) (*EditTicketRes, error)
	// Public method to list tickets
	ListTickets(context.Context, *ListTicketsReq) (*ListTicketsRes, error)
	// Public method to assign ticket
	AssignTicket(context.Context, *AssignTicketReq) (*AssignTicketRes, error)
	// Closes the ticket
	CloseTicket(context.Context, *CloseTicketReq) (*CloseTicketRes, error)
	// Public method to view ticket
	ViewTicket(context.Context, *ViewTicketReq) (*ViewTicketRes, error)
	// Public Method to create a Comment.
	CreateComment(context.Context, *CreateCommentReq) (*CreateCommentRes, error)
	// Public method to Enable Project for Ticketing system
	EnableProject(context.Context, *EnableProjectReq) (*EnableProjectRes, error)
	// Public method to List projects enabled for Ticketing system
	ListEnabledProjects(context.Context, *ListEnabledProjectsReq) (*ListEnabledProjectsRes, error)
	// Public Method to create a sla.
	CreateSLA(context.Context, *CreateSlaReq) (*CreateSlaRes, error)
	// Public method to list sla
	ListSLA(context.Context, *ListSlaReq) (*ListSlaRes, error)
	// Public method to update sla
	UpdateSLA(context.Context, *UpdateSlaReq) (*UpdateSlaRes, error)
	// Public method to list sla_condition
	ListSLACondition(context.Context, *ListSlaConditionReq) (*ListSlaConditionRes, error)
	// Public method to list sla_condition
	ReplyComment(context.Context, *ReplyCommentReq) (*ReplyCommentRes, error)
	// Public method to List audit log for Ticketing system
	ListTicketAuditLog(context.Context, *ListTicketAuditLogReq) (*ListTicketAuditLogRes, error)
	// Public method to assign a ticket
	AssignSelf(context.Context, *CreateSelfAssignReq) (*CreateSelfAssignRes, error)
	mustEmbedUnimplementedTicketsServer()
}

// UnimplementedTicketsServer must be embedded to have forward compatible implementations.
type UnimplementedTicketsServer struct {
}

func (UnimplementedTicketsServer) CreateTicket(context.Context, *CreateTicketReq) (*CreateTicketRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedTicketsServer) EditTicket(context.Context, *EditTicketReq) (*EditTicketRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditTicket not implemented")
}
func (UnimplementedTicketsServer) ListTickets(context.Context, *ListTicketsReq) (*ListTicketsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTickets not implemented")
}
func (UnimplementedTicketsServer) AssignTicket(context.Context, *AssignTicketReq) (*AssignTicketRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignTicket not implemented")
}
func (UnimplementedTicketsServer) CloseTicket(context.Context, *CloseTicketReq) (*CloseTicketRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTicket not implemented")
}
func (UnimplementedTicketsServer) ViewTicket(context.Context, *ViewTicketReq) (*ViewTicketRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewTicket not implemented")
}
func (UnimplementedTicketsServer) CreateComment(context.Context, *CreateCommentReq) (*CreateCommentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedTicketsServer) EnableProject(context.Context, *EnableProjectReq) (*EnableProjectRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableProject not implemented")
}
func (UnimplementedTicketsServer) ListEnabledProjects(context.Context, *ListEnabledProjectsReq) (*ListEnabledProjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledProjects not implemented")
}
func (UnimplementedTicketsServer) CreateSLA(context.Context, *CreateSlaReq) (*CreateSlaRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSLA not implemented")
}
func (UnimplementedTicketsServer) ListSLA(context.Context, *ListSlaReq) (*ListSlaRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSLA not implemented")
}
func (UnimplementedTicketsServer) UpdateSLA(context.Context, *UpdateSlaReq) (*UpdateSlaRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSLA not implemented")
}
func (UnimplementedTicketsServer) ListSLACondition(context.Context, *ListSlaConditionReq) (*ListSlaConditionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSLACondition not implemented")
}
func (UnimplementedTicketsServer) ReplyComment(context.Context, *ReplyCommentReq) (*ReplyCommentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplyComment not implemented")
}
func (UnimplementedTicketsServer) ListTicketAuditLog(context.Context, *ListTicketAuditLogReq) (*ListTicketAuditLogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTicketAuditLog not implemented")
}
func (UnimplementedTicketsServer) AssignSelf(context.Context, *CreateSelfAssignReq) (*CreateSelfAssignRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignSelf not implemented")
}
func (UnimplementedTicketsServer) mustEmbedUnimplementedTicketsServer() {}

// UnsafeTicketsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketsServer will
// result in compilation errors.
type UnsafeTicketsServer interface {
	mustEmbedUnimplementedTicketsServer()
}

func RegisterTicketsServer(s grpc.ServiceRegistrar, srv TicketsServer) {
	s.RegisterService(&Tickets_ServiceDesc, srv)
}

func _Tickets_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_CreateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).CreateTicket(ctx, req.(*CreateTicketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_EditTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditTicketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).EditTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_EditTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).EditTicket(ctx, req.(*EditTicketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_ListTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTicketsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).ListTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_ListTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).ListTickets(ctx, req.(*ListTicketsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_AssignTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignTicketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).AssignTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_AssignTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).AssignTicket(ctx, req.(*AssignTicketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_CloseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseTicketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).CloseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_CloseTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).CloseTicket(ctx, req.(*CloseTicketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_ViewTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewTicketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).ViewTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_ViewTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).ViewTicket(ctx, req.(*ViewTicketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).CreateComment(ctx, req.(*CreateCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_EnableProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).EnableProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_EnableProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).EnableProject(ctx, req.(*EnableProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_ListEnabledProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledProjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).ListEnabledProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_ListEnabledProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).ListEnabledProjects(ctx, req.(*ListEnabledProjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_CreateSLA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSlaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).CreateSLA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_CreateSLA_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).CreateSLA(ctx, req.(*CreateSlaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_ListSLA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSlaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).ListSLA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_ListSLA_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).ListSLA(ctx, req.(*ListSlaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_UpdateSLA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSlaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).UpdateSLA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_UpdateSLA_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).UpdateSLA(ctx, req.(*UpdateSlaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_ListSLACondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSlaConditionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).ListSLACondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_ListSLACondition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).ListSLACondition(ctx, req.(*ListSlaConditionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_ReplyComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).ReplyComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_ReplyComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).ReplyComment(ctx, req.(*ReplyCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_ListTicketAuditLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTicketAuditLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).ListTicketAuditLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_ListTicketAuditLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).ListTicketAuditLog(ctx, req.(*ListTicketAuditLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_AssignSelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSelfAssignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).AssignSelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tickets_AssignSelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).AssignSelf(ctx, req.(*CreateSelfAssignReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Tickets_ServiceDesc is the grpc.ServiceDesc for Tickets service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tickets_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.tickets.Tickets",
	HandlerType: (*TicketsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTicket",
			Handler:    _Tickets_CreateTicket_Handler,
		},
		{
			MethodName: "EditTicket",
			Handler:    _Tickets_EditTicket_Handler,
		},
		{
			MethodName: "ListTickets",
			Handler:    _Tickets_ListTickets_Handler,
		},
		{
			MethodName: "AssignTicket",
			Handler:    _Tickets_AssignTicket_Handler,
		},
		{
			MethodName: "CloseTicket",
			Handler:    _Tickets_CloseTicket_Handler,
		},
		{
			MethodName: "ViewTicket",
			Handler:    _Tickets_ViewTicket_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _Tickets_CreateComment_Handler,
		},
		{
			MethodName: "EnableProject",
			Handler:    _Tickets_EnableProject_Handler,
		},
		{
			MethodName: "ListEnabledProjects",
			Handler:    _Tickets_ListEnabledProjects_Handler,
		},
		{
			MethodName: "CreateSLA",
			Handler:    _Tickets_CreateSLA_Handler,
		},
		{
			MethodName: "ListSLA",
			Handler:    _Tickets_ListSLA_Handler,
		},
		{
			MethodName: "UpdateSLA",
			Handler:    _Tickets_UpdateSLA_Handler,
		},
		{
			MethodName: "ListSLACondition",
			Handler:    _Tickets_ListSLACondition_Handler,
		},
		{
			MethodName: "ReplyComment",
			Handler:    _Tickets_ReplyComment_Handler,
		},
		{
			MethodName: "ListTicketAuditLog",
			Handler:    _Tickets_ListTicketAuditLog_Handler,
		},
		{
			MethodName: "AssignSelf",
			Handler:    _Tickets_AssignSelf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1alpha1/tickets/service.proto",
}
