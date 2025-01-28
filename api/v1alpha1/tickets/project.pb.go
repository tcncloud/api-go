// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: api/v1alpha1/tickets/project.proto

package tickets

import (
	commons "github.com/tcncloud/api-go/api/commons"
	audit "github.com/tcncloud/api-go/api/commons/audit"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EnableProjectReq - Details of project to be enabled in the ticketing system
type EnableProjectReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// project seq id of project to be enabled
	ProjectSid int64 `protobuf:"varint,1,opt,name=project_sid,json=projectSid,proto3" json:"project_sid,omitempty"`
	// Project code to be set
	ProjectCode string `protobuf:"bytes,2,opt,name=project_code,json=projectCode,proto3" json:"project_code,omitempty"`
	// Project title pulled from Omni
	ProjectTitle string `protobuf:"bytes,3,opt,name=project_title,json=projectTitle,proto3" json:"project_title,omitempty"`
	// Is project enabled for ticketing
	IsActive      bool `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnableProjectReq) Reset() {
	*x = EnableProjectReq{}
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnableProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableProjectReq) ProtoMessage() {}

func (x *EnableProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableProjectReq.ProtoReflect.Descriptor instead.
func (*EnableProjectReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_tickets_project_proto_rawDescGZIP(), []int{0}
}

func (x *EnableProjectReq) GetProjectSid() int64 {
	if x != nil {
		return x.ProjectSid
	}
	return 0
}

func (x *EnableProjectReq) GetProjectCode() string {
	if x != nil {
		return x.ProjectCode
	}
	return ""
}

func (x *EnableProjectReq) GetProjectTitle() string {
	if x != nil {
		return x.ProjectTitle
	}
	return ""
}

func (x *EnableProjectReq) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

// EnableProjectRes - Indicates result of the project enable request
type EnableProjectRes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// result of project mapped to ticket was successful
	//
	// Deprecated: Marked as deprecated in api/v1alpha1/tickets/project.proto.
	Success       bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnableProjectRes) Reset() {
	*x = EnableProjectRes{}
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnableProjectRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableProjectRes) ProtoMessage() {}

func (x *EnableProjectRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableProjectRes.ProtoReflect.Descriptor instead.
func (*EnableProjectRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_tickets_project_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in api/v1alpha1/tickets/project.proto.
func (x *EnableProjectRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// ListEnabledProjectsReq - Request to list enabled projects in ticketing system
type ListEnabledProjectsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListEnabledProjectsReq) Reset() {
	*x = ListEnabledProjectsReq{}
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEnabledProjectsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnabledProjectsReq) ProtoMessage() {}

func (x *ListEnabledProjectsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnabledProjectsReq.ProtoReflect.Descriptor instead.
func (*ListEnabledProjectsReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_tickets_project_proto_rawDescGZIP(), []int{2}
}

// ListEnabledProjectsRes - Response wraps list of projects enabled for ticketing system
type ListEnabledProjectsRes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// enabled projects list
	Projects      []*commons.TicketProject `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListEnabledProjectsRes) Reset() {
	*x = ListEnabledProjectsRes{}
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEnabledProjectsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnabledProjectsRes) ProtoMessage() {}

func (x *ListEnabledProjectsRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnabledProjectsRes.ProtoReflect.Descriptor instead.
func (*ListEnabledProjectsRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_tickets_project_proto_rawDescGZIP(), []int{3}
}

func (x *ListEnabledProjectsRes) GetProjects() []*commons.TicketProject {
	if x != nil {
		return x.Projects
	}
	return nil
}

// ListTicketAuditLogReq -
type ListTicketAuditLogReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ticket_sid  scoping filter
	//
	// Deprecated: Marked as deprecated in api/v1alpha1/tickets/project.proto.
	TicketSid int64 `protobuf:"varint,1,opt,name=ticket_sid,json=ticketSid,proto3" json:"ticket_sid,omitempty"`
	// ticket_code
	TicketCode    string `protobuf:"bytes,16,opt,name=ticket_code,json=ticketCode,proto3" json:"ticket_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTicketAuditLogReq) Reset() {
	*x = ListTicketAuditLogReq{}
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTicketAuditLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTicketAuditLogReq) ProtoMessage() {}

func (x *ListTicketAuditLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTicketAuditLogReq.ProtoReflect.Descriptor instead.
func (*ListTicketAuditLogReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_tickets_project_proto_rawDescGZIP(), []int{4}
}

// Deprecated: Marked as deprecated in api/v1alpha1/tickets/project.proto.
func (x *ListTicketAuditLogReq) GetTicketSid() int64 {
	if x != nil {
		return x.TicketSid
	}
	return 0
}

func (x *ListTicketAuditLogReq) GetTicketCode() string {
	if x != nil {
		return x.TicketCode
	}
	return ""
}

// ListTicketAuditLogRes -
type ListTicketAuditLogRes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the retrieved audit events
	Events        []*audit.AuditEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTicketAuditLogRes) Reset() {
	*x = ListTicketAuditLogRes{}
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTicketAuditLogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTicketAuditLogRes) ProtoMessage() {}

func (x *ListTicketAuditLogRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_tickets_project_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTicketAuditLogRes.ProtoReflect.Descriptor instead.
func (*ListTicketAuditLogRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_tickets_project_proto_rawDescGZIP(), []int{5}
}

func (x *ListTicketAuditLogRes) GetEvents() []*audit.AuditEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_api_v1alpha1_tickets_project_proto protoreflect.FileDescriptor

var file_api_v1alpha1_tickets_project_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x10, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02,
	0x30, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x22, 0x30, 0x0a, 0x10, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x22,
	0x50, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x22, 0x5d, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0a, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04,
	0x18, 0x01, 0x30, 0x01, 0x52, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x4e, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x42, 0xcb, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x0c, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0xa2, 0x02,
	0x03, 0x41, 0x56, 0x54, 0xaa, 0x02, 0x14, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0xca, 0x02, 0x14, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0xe2, 0x02, 0x20, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_v1alpha1_tickets_project_proto_rawDescOnce sync.Once
	file_api_v1alpha1_tickets_project_proto_rawDescData []byte
)

func file_api_v1alpha1_tickets_project_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_tickets_project_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_tickets_project_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_tickets_project_proto_rawDesc), len(file_api_v1alpha1_tickets_project_proto_rawDesc)))
	})
	return file_api_v1alpha1_tickets_project_proto_rawDescData
}

var file_api_v1alpha1_tickets_project_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1alpha1_tickets_project_proto_goTypes = []any{
	(*EnableProjectReq)(nil),       // 0: api.v1alpha1.tickets.EnableProjectReq
	(*EnableProjectRes)(nil),       // 1: api.v1alpha1.tickets.EnableProjectRes
	(*ListEnabledProjectsReq)(nil), // 2: api.v1alpha1.tickets.ListEnabledProjectsReq
	(*ListEnabledProjectsRes)(nil), // 3: api.v1alpha1.tickets.ListEnabledProjectsRes
	(*ListTicketAuditLogReq)(nil),  // 4: api.v1alpha1.tickets.ListTicketAuditLogReq
	(*ListTicketAuditLogRes)(nil),  // 5: api.v1alpha1.tickets.ListTicketAuditLogRes
	(*commons.TicketProject)(nil),  // 6: api.commons.TicketProject
	(*audit.AuditEvent)(nil),       // 7: api.commons.audit.AuditEvent
}
var file_api_v1alpha1_tickets_project_proto_depIdxs = []int32{
	6, // 0: api.v1alpha1.tickets.ListEnabledProjectsRes.projects:type_name -> api.commons.TicketProject
	7, // 1: api.v1alpha1.tickets.ListTicketAuditLogRes.events:type_name -> api.commons.audit.AuditEvent
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_tickets_project_proto_init() }
func file_api_v1alpha1_tickets_project_proto_init() {
	if File_api_v1alpha1_tickets_project_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_tickets_project_proto_rawDesc), len(file_api_v1alpha1_tickets_project_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_tickets_project_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_tickets_project_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_tickets_project_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_tickets_project_proto = out.File
	file_api_v1alpha1_tickets_project_proto_goTypes = nil
	file_api_v1alpha1_tickets_project_proto_depIdxs = nil
}
