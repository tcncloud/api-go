// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/commons/audit/billing_events.proto

package audit

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BillingCommitBillingPlanEvent is fired when a billing plan is committed
type BillingCommitBillingPlanEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingPlanId string `protobuf:"bytes,1,opt,name=billing_plan_id,json=billingPlanId,proto3" json:"billing_plan_id,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BillingCommitBillingPlanEvent) Reset() {
	*x = BillingCommitBillingPlanEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_billing_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingCommitBillingPlanEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingCommitBillingPlanEvent) ProtoMessage() {}

func (x *BillingCommitBillingPlanEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_billing_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingCommitBillingPlanEvent.ProtoReflect.Descriptor instead.
func (*BillingCommitBillingPlanEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_billing_events_proto_rawDescGZIP(), []int{0}
}

func (x *BillingCommitBillingPlanEvent) GetBillingPlanId() string {
	if x != nil {
		return x.BillingPlanId
	}
	return ""
}

func (x *BillingCommitBillingPlanEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// BillingCreateBillingPlanEvent is fired when a billing plan is created
type BillingCreateBillingPlanEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingPlan string `protobuf:"bytes,1,opt,name=billing_plan,json=billingPlan,proto3" json:"billing_plan,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BillingCreateBillingPlanEvent) Reset() {
	*x = BillingCreateBillingPlanEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_billing_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingCreateBillingPlanEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingCreateBillingPlanEvent) ProtoMessage() {}

func (x *BillingCreateBillingPlanEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_billing_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingCreateBillingPlanEvent.ProtoReflect.Descriptor instead.
func (*BillingCreateBillingPlanEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_billing_events_proto_rawDescGZIP(), []int{1}
}

func (x *BillingCreateBillingPlanEvent) GetBillingPlan() string {
	if x != nil {
		return x.BillingPlan
	}
	return ""
}

func (x *BillingCreateBillingPlanEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// BillingCreateInvoiceEvent is fired when an invoice is created
//
// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
type BillingCreateInvoiceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
	Invoice string `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"`
	// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BillingCreateInvoiceEvent) Reset() {
	*x = BillingCreateInvoiceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_billing_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingCreateInvoiceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingCreateInvoiceEvent) ProtoMessage() {}

func (x *BillingCreateInvoiceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_billing_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingCreateInvoiceEvent.ProtoReflect.Descriptor instead.
func (*BillingCreateInvoiceEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_billing_events_proto_rawDescGZIP(), []int{2}
}

// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
func (x *BillingCreateInvoiceEvent) GetInvoice() string {
	if x != nil {
		return x.Invoice
	}
	return ""
}

// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
func (x *BillingCreateInvoiceEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// BillingCreateRateDefinitionEvent is fired when a rate definition is created
type BillingCreateRateDefinitionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RateDefinition string `protobuf:"bytes,1,opt,name=rate_definition,json=rateDefinition,proto3" json:"rate_definition,omitempty"`
	UserId         string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BillingCreateRateDefinitionEvent) Reset() {
	*x = BillingCreateRateDefinitionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_billing_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingCreateRateDefinitionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingCreateRateDefinitionEvent) ProtoMessage() {}

func (x *BillingCreateRateDefinitionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_billing_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingCreateRateDefinitionEvent.ProtoReflect.Descriptor instead.
func (*BillingCreateRateDefinitionEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_billing_events_proto_rawDescGZIP(), []int{3}
}

func (x *BillingCreateRateDefinitionEvent) GetRateDefinition() string {
	if x != nil {
		return x.RateDefinition
	}
	return ""
}

func (x *BillingCreateRateDefinitionEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// BillingDeleteBillingPlanEvent is fired when a billing plan is deleted
type BillingDeleteBillingPlanEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingPlanId string `protobuf:"bytes,1,opt,name=billing_plan_id,json=billingPlanId,proto3" json:"billing_plan_id,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BillingDeleteBillingPlanEvent) Reset() {
	*x = BillingDeleteBillingPlanEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_billing_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingDeleteBillingPlanEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingDeleteBillingPlanEvent) ProtoMessage() {}

func (x *BillingDeleteBillingPlanEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_billing_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingDeleteBillingPlanEvent.ProtoReflect.Descriptor instead.
func (*BillingDeleteBillingPlanEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_billing_events_proto_rawDescGZIP(), []int{4}
}

func (x *BillingDeleteBillingPlanEvent) GetBillingPlanId() string {
	if x != nil {
		return x.BillingPlanId
	}
	return ""
}

func (x *BillingDeleteBillingPlanEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// BillingDeleteInvoiceEvent is fired when an invoice is deleted
//
// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
type BillingDeleteInvoiceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
	InvoiceId string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BillingDeleteInvoiceEvent) Reset() {
	*x = BillingDeleteInvoiceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_billing_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingDeleteInvoiceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingDeleteInvoiceEvent) ProtoMessage() {}

func (x *BillingDeleteInvoiceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_billing_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingDeleteInvoiceEvent.ProtoReflect.Descriptor instead.
func (*BillingDeleteInvoiceEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_billing_events_proto_rawDescGZIP(), []int{5}
}

// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
func (x *BillingDeleteInvoiceEvent) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
func (x *BillingDeleteInvoiceEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// BillingDeleteRateDefinitionEvent is fired when a rate definition is deleted
type BillingDeleteRateDefinitionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RateDefinitionId string `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	UserId           string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BillingDeleteRateDefinitionEvent) Reset() {
	*x = BillingDeleteRateDefinitionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_billing_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingDeleteRateDefinitionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingDeleteRateDefinitionEvent) ProtoMessage() {}

func (x *BillingDeleteRateDefinitionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_billing_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingDeleteRateDefinitionEvent.ProtoReflect.Descriptor instead.
func (*BillingDeleteRateDefinitionEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_billing_events_proto_rawDescGZIP(), []int{6}
}

func (x *BillingDeleteRateDefinitionEvent) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

func (x *BillingDeleteRateDefinitionEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// BillingExportInvoiceEvent is fired when an invoice is exported
type BillingExportInvoiceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvoiceId string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BillingExportInvoiceEvent) Reset() {
	*x = BillingExportInvoiceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_billing_events_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingExportInvoiceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingExportInvoiceEvent) ProtoMessage() {}

func (x *BillingExportInvoiceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_billing_events_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingExportInvoiceEvent.ProtoReflect.Descriptor instead.
func (*BillingExportInvoiceEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_billing_events_proto_rawDescGZIP(), []int{7}
}

func (x *BillingExportInvoiceEvent) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *BillingExportInvoiceEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// BillingUpdateBillingPlanEvent is fired when a billing plan is updated
type BillingUpdateBillingPlanEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingPlan string `protobuf:"bytes,1,opt,name=billing_plan,json=billingPlan,proto3" json:"billing_plan,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BillingUpdateBillingPlanEvent) Reset() {
	*x = BillingUpdateBillingPlanEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_billing_events_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingUpdateBillingPlanEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingUpdateBillingPlanEvent) ProtoMessage() {}

func (x *BillingUpdateBillingPlanEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_billing_events_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingUpdateBillingPlanEvent.ProtoReflect.Descriptor instead.
func (*BillingUpdateBillingPlanEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_billing_events_proto_rawDescGZIP(), []int{8}
}

func (x *BillingUpdateBillingPlanEvent) GetBillingPlan() string {
	if x != nil {
		return x.BillingPlan
	}
	return ""
}

func (x *BillingUpdateBillingPlanEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// BillingUpdateInvoiceEvent is fired when an invoice is updated
//
// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
type BillingUpdateInvoiceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
	Invoice string `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"`
	// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BillingUpdateInvoiceEvent) Reset() {
	*x = BillingUpdateInvoiceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_billing_events_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingUpdateInvoiceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingUpdateInvoiceEvent) ProtoMessage() {}

func (x *BillingUpdateInvoiceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_billing_events_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingUpdateInvoiceEvent.ProtoReflect.Descriptor instead.
func (*BillingUpdateInvoiceEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_billing_events_proto_rawDescGZIP(), []int{9}
}

// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
func (x *BillingUpdateInvoiceEvent) GetInvoice() string {
	if x != nil {
		return x.Invoice
	}
	return ""
}

// Deprecated: Marked as deprecated in api/commons/audit/billing_events.proto.
func (x *BillingUpdateInvoiceEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// BillingUpdateRateDefinitionEvent is fired when a rate definition is updated
type BillingUpdateRateDefinitionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RateDefinition string `protobuf:"bytes,1,opt,name=rate_definition,json=rateDefinition,proto3" json:"rate_definition,omitempty"`
	UserId         string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BillingUpdateRateDefinitionEvent) Reset() {
	*x = BillingUpdateRateDefinitionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_billing_events_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingUpdateRateDefinitionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingUpdateRateDefinitionEvent) ProtoMessage() {}

func (x *BillingUpdateRateDefinitionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_billing_events_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingUpdateRateDefinitionEvent.ProtoReflect.Descriptor instead.
func (*BillingUpdateRateDefinitionEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_billing_events_proto_rawDescGZIP(), []int{10}
}

func (x *BillingUpdateRateDefinitionEvent) GetRateDefinition() string {
	if x != nil {
		return x.RateDefinition
	}
	return ""
}

func (x *BillingUpdateRateDefinitionEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_api_commons_audit_billing_events_proto protoreflect.FileDescriptor

var file_api_commons_audit_billing_events_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x22, 0x60, 0x0a, 0x1d, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6c,
	0x61, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5b, 0x0a,
	0x1d, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x19, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x22, 0x64, 0x0a, 0x20, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x1d,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a,
	0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50,
	0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5f,
	0x0a, 0x19, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0a, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x22,
	0x69, 0x0a, 0x20, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x72, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x19, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x5b, 0x0a, 0x1d, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50,
	0x6c, 0x61, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x19,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x07, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x22, 0x64, 0x0a, 0x20, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0xbf,
	0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0x12, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0xa2, 0x02, 0x03, 0x41,
	0x43, 0x41, 0xaa, 0x02, 0x11, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0xca, 0x02, 0x11, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0xe2, 0x02, 0x1d, 0x41, 0x70, 0x69,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_audit_billing_events_proto_rawDescOnce sync.Once
	file_api_commons_audit_billing_events_proto_rawDescData = file_api_commons_audit_billing_events_proto_rawDesc
)

func file_api_commons_audit_billing_events_proto_rawDescGZIP() []byte {
	file_api_commons_audit_billing_events_proto_rawDescOnce.Do(func() {
		file_api_commons_audit_billing_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_audit_billing_events_proto_rawDescData)
	})
	return file_api_commons_audit_billing_events_proto_rawDescData
}

var file_api_commons_audit_billing_events_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_commons_audit_billing_events_proto_goTypes = []interface{}{
	(*BillingCommitBillingPlanEvent)(nil),    // 0: api.commons.audit.BillingCommitBillingPlanEvent
	(*BillingCreateBillingPlanEvent)(nil),    // 1: api.commons.audit.BillingCreateBillingPlanEvent
	(*BillingCreateInvoiceEvent)(nil),        // 2: api.commons.audit.BillingCreateInvoiceEvent
	(*BillingCreateRateDefinitionEvent)(nil), // 3: api.commons.audit.BillingCreateRateDefinitionEvent
	(*BillingDeleteBillingPlanEvent)(nil),    // 4: api.commons.audit.BillingDeleteBillingPlanEvent
	(*BillingDeleteInvoiceEvent)(nil),        // 5: api.commons.audit.BillingDeleteInvoiceEvent
	(*BillingDeleteRateDefinitionEvent)(nil), // 6: api.commons.audit.BillingDeleteRateDefinitionEvent
	(*BillingExportInvoiceEvent)(nil),        // 7: api.commons.audit.BillingExportInvoiceEvent
	(*BillingUpdateBillingPlanEvent)(nil),    // 8: api.commons.audit.BillingUpdateBillingPlanEvent
	(*BillingUpdateInvoiceEvent)(nil),        // 9: api.commons.audit.BillingUpdateInvoiceEvent
	(*BillingUpdateRateDefinitionEvent)(nil), // 10: api.commons.audit.BillingUpdateRateDefinitionEvent
}
var file_api_commons_audit_billing_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_audit_billing_events_proto_init() }
func file_api_commons_audit_billing_events_proto_init() {
	if File_api_commons_audit_billing_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commons_audit_billing_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingCommitBillingPlanEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_commons_audit_billing_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingCreateBillingPlanEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_commons_audit_billing_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingCreateInvoiceEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_commons_audit_billing_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingCreateRateDefinitionEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_commons_audit_billing_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingDeleteBillingPlanEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_commons_audit_billing_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingDeleteInvoiceEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_commons_audit_billing_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingDeleteRateDefinitionEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_commons_audit_billing_events_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingExportInvoiceEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_commons_audit_billing_events_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingUpdateBillingPlanEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_commons_audit_billing_events_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingUpdateInvoiceEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_commons_audit_billing_events_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingUpdateRateDefinitionEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_audit_billing_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_audit_billing_events_proto_goTypes,
		DependencyIndexes: file_api_commons_audit_billing_events_proto_depIdxs,
		MessageInfos:      file_api_commons_audit_billing_events_proto_msgTypes,
	}.Build()
	File_api_commons_audit_billing_events_proto = out.File
	file_api_commons_audit_billing_events_proto_rawDesc = nil
	file_api_commons_audit_billing_events_proto_goTypes = nil
	file_api_commons_audit_billing_events_proto_depIdxs = nil
}
