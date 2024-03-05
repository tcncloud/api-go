// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/commons/audit/delivery_events.proto

package audit

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeliveryFailureEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryDefinitionName int64                  `protobuf:"varint,1,opt,name=delivery_definition_name,json=deliveryDefinitionName,proto3" json:"delivery_definition_name,omitempty"`
	OrgId                  string                 `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	TransactionSid         int64                  `protobuf:"varint,3,opt,name=transaction_sid,json=transactionSid,proto3" json:"transaction_sid,omitempty"`
	AttachmentNames        []string               `protobuf:"bytes,4,rep,name=attachment_names,json=attachmentNames,proto3" json:"attachment_names,omitempty"`
	FailureTime            *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=failure_time,json=failureTime,proto3" json:"failure_time,omitempty"`
	FailureErrorMessage    string                 `protobuf:"bytes,6,opt,name=failure_error_message,json=failureErrorMessage,proto3" json:"failure_error_message,omitempty"`
	Definition             string                 `protobuf:"bytes,7,opt,name=definition,proto3" json:"definition,omitempty"`
	OriginalPayload        string                 `protobuf:"bytes,8,opt,name=original_payload,json=originalPayload,proto3" json:"original_payload,omitempty"`
}

func (x *DeliveryFailureEvent) Reset() {
	*x = DeliveryFailureEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_delivery_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryFailureEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryFailureEvent) ProtoMessage() {}

func (x *DeliveryFailureEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_delivery_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryFailureEvent.ProtoReflect.Descriptor instead.
func (*DeliveryFailureEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_delivery_events_proto_rawDescGZIP(), []int{0}
}

func (x *DeliveryFailureEvent) GetDeliveryDefinitionName() int64 {
	if x != nil {
		return x.DeliveryDefinitionName
	}
	return 0
}

func (x *DeliveryFailureEvent) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *DeliveryFailureEvent) GetTransactionSid() int64 {
	if x != nil {
		return x.TransactionSid
	}
	return 0
}

func (x *DeliveryFailureEvent) GetAttachmentNames() []string {
	if x != nil {
		return x.AttachmentNames
	}
	return nil
}

func (x *DeliveryFailureEvent) GetFailureTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FailureTime
	}
	return nil
}

func (x *DeliveryFailureEvent) GetFailureErrorMessage() string {
	if x != nil {
		return x.FailureErrorMessage
	}
	return ""
}

func (x *DeliveryFailureEvent) GetDefinition() string {
	if x != nil {
		return x.Definition
	}
	return ""
}

func (x *DeliveryFailureEvent) GetOriginalPayload() string {
	if x != nil {
		return x.OriginalPayload
	}
	return ""
}

type DeliverySuccessEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryDefinitionName string                 `protobuf:"bytes,1,opt,name=delivery_definition_name,json=deliveryDefinitionName,proto3" json:"delivery_definition_name,omitempty"`
	OrgId                  string                 `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	TransactionSid         int64                  `protobuf:"varint,3,opt,name=transaction_sid,json=transactionSid,proto3" json:"transaction_sid,omitempty"`
	AttachmentNames        []string               `protobuf:"bytes,4,rep,name=attachment_names,json=attachmentNames,proto3" json:"attachment_names,omitempty"`
	SuccessTime            *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=success_time,json=successTime,proto3" json:"success_time,omitempty"`
	SuccessMessage         string                 `protobuf:"bytes,6,opt,name=success_message,json=successMessage,proto3" json:"success_message,omitempty"`
}

func (x *DeliverySuccessEvent) Reset() {
	*x = DeliverySuccessEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_delivery_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverySuccessEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverySuccessEvent) ProtoMessage() {}

func (x *DeliverySuccessEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_delivery_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverySuccessEvent.ProtoReflect.Descriptor instead.
func (*DeliverySuccessEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_delivery_events_proto_rawDescGZIP(), []int{1}
}

func (x *DeliverySuccessEvent) GetDeliveryDefinitionName() string {
	if x != nil {
		return x.DeliveryDefinitionName
	}
	return ""
}

func (x *DeliverySuccessEvent) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *DeliverySuccessEvent) GetTransactionSid() int64 {
	if x != nil {
		return x.TransactionSid
	}
	return 0
}

func (x *DeliverySuccessEvent) GetAttachmentNames() []string {
	if x != nil {
		return x.AttachmentNames
	}
	return nil
}

func (x *DeliverySuccessEvent) GetSuccessTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SuccessTime
	}
	return nil
}

func (x *DeliverySuccessEvent) GetSuccessMessage() string {
	if x != nil {
		return x.SuccessMessage
	}
	return ""
}

var File_api_commons_audit_delivery_events_proto protoreflect.FileDescriptor

var file_api_commons_audit_delivery_events_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x02,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64,
	0x12, 0x29, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29,
	0x0a, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xa3, 0x02, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0xc0, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0x13, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0xa2, 0x02,
	0x03, 0x41, 0x43, 0x41, 0xaa, 0x02, 0x11, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0xca, 0x02, 0x11, 0x41, 0x70, 0x69, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0xe2, 0x02, 0x1d, 0x41,
	0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_audit_delivery_events_proto_rawDescOnce sync.Once
	file_api_commons_audit_delivery_events_proto_rawDescData = file_api_commons_audit_delivery_events_proto_rawDesc
)

func file_api_commons_audit_delivery_events_proto_rawDescGZIP() []byte {
	file_api_commons_audit_delivery_events_proto_rawDescOnce.Do(func() {
		file_api_commons_audit_delivery_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_audit_delivery_events_proto_rawDescData)
	})
	return file_api_commons_audit_delivery_events_proto_rawDescData
}

var file_api_commons_audit_delivery_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_audit_delivery_events_proto_goTypes = []interface{}{
	(*DeliveryFailureEvent)(nil),  // 0: api.commons.audit.DeliveryFailureEvent
	(*DeliverySuccessEvent)(nil),  // 1: api.commons.audit.DeliverySuccessEvent
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_api_commons_audit_delivery_events_proto_depIdxs = []int32{
	2, // 0: api.commons.audit.DeliveryFailureEvent.failure_time:type_name -> google.protobuf.Timestamp
	2, // 1: api.commons.audit.DeliverySuccessEvent.success_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_commons_audit_delivery_events_proto_init() }
func file_api_commons_audit_delivery_events_proto_init() {
	if File_api_commons_audit_delivery_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commons_audit_delivery_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryFailureEvent); i {
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
		file_api_commons_audit_delivery_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverySuccessEvent); i {
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
			RawDescriptor: file_api_commons_audit_delivery_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_audit_delivery_events_proto_goTypes,
		DependencyIndexes: file_api_commons_audit_delivery_events_proto_depIdxs,
		MessageInfos:      file_api_commons_audit_delivery_events_proto_msgTypes,
	}.Build()
	File_api_commons_audit_delivery_events_proto = out.File
	file_api_commons_audit_delivery_events_proto_rawDesc = nil
	file_api_commons_audit_delivery_events_proto_goTypes = nil
	file_api_commons_audit_delivery_events_proto_depIdxs = nil
}
