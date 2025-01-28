// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: api/commons/audit/lms_events.proto

package audit

import (
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

type LMSPipelineFailureEvent struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ElementId      string                 `protobuf:"bytes,1,opt,name=element_id,json=elementId,proto3" json:"element_id,omitempty"`
	ElementName    string                 `protobuf:"bytes,2,opt,name=element_name,json=elementName,proto3" json:"element_name,omitempty"`
	FileName       string                 `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FailureMessage string                 `protobuf:"bytes,4,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *LMSPipelineFailureEvent) Reset() {
	*x = LMSPipelineFailureEvent{}
	mi := &file_api_commons_audit_lms_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LMSPipelineFailureEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LMSPipelineFailureEvent) ProtoMessage() {}

func (x *LMSPipelineFailureEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_lms_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LMSPipelineFailureEvent.ProtoReflect.Descriptor instead.
func (*LMSPipelineFailureEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_lms_events_proto_rawDescGZIP(), []int{0}
}

func (x *LMSPipelineFailureEvent) GetElementId() string {
	if x != nil {
		return x.ElementId
	}
	return ""
}

func (x *LMSPipelineFailureEvent) GetElementName() string {
	if x != nil {
		return x.ElementName
	}
	return ""
}

func (x *LMSPipelineFailureEvent) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *LMSPipelineFailureEvent) GetFailureMessage() string {
	if x != nil {
		return x.FailureMessage
	}
	return ""
}

type LMSPipelineNoOutputEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ElementId     string                 `protobuf:"bytes,1,opt,name=element_id,json=elementId,proto3" json:"element_id,omitempty"`
	ElementName   string                 `protobuf:"bytes,2,opt,name=element_name,json=elementName,proto3" json:"element_name,omitempty"`
	FileName      string                 `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	EventMessage  string                 `protobuf:"bytes,4,opt,name=event_message,json=eventMessage,proto3" json:"event_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LMSPipelineNoOutputEvent) Reset() {
	*x = LMSPipelineNoOutputEvent{}
	mi := &file_api_commons_audit_lms_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LMSPipelineNoOutputEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LMSPipelineNoOutputEvent) ProtoMessage() {}

func (x *LMSPipelineNoOutputEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_lms_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LMSPipelineNoOutputEvent.ProtoReflect.Descriptor instead.
func (*LMSPipelineNoOutputEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_lms_events_proto_rawDescGZIP(), []int{1}
}

func (x *LMSPipelineNoOutputEvent) GetElementId() string {
	if x != nil {
		return x.ElementId
	}
	return ""
}

func (x *LMSPipelineNoOutputEvent) GetElementName() string {
	if x != nil {
		return x.ElementName
	}
	return ""
}

func (x *LMSPipelineNoOutputEvent) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *LMSPipelineNoOutputEvent) GetEventMessage() string {
	if x != nil {
		return x.EventMessage
	}
	return ""
}

type LMSPipelineSuccessfulEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ElementId     string                 `protobuf:"bytes,1,opt,name=element_id,json=elementId,proto3" json:"element_id,omitempty"`
	ElementName   string                 `protobuf:"bytes,2,opt,name=element_name,json=elementName,proto3" json:"element_name,omitempty"`
	FileName      string                 `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	EventMessage  string                 `protobuf:"bytes,4,opt,name=event_message,json=eventMessage,proto3" json:"event_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LMSPipelineSuccessfulEvent) Reset() {
	*x = LMSPipelineSuccessfulEvent{}
	mi := &file_api_commons_audit_lms_events_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LMSPipelineSuccessfulEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LMSPipelineSuccessfulEvent) ProtoMessage() {}

func (x *LMSPipelineSuccessfulEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_lms_events_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LMSPipelineSuccessfulEvent.ProtoReflect.Descriptor instead.
func (*LMSPipelineSuccessfulEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_lms_events_proto_rawDescGZIP(), []int{2}
}

func (x *LMSPipelineSuccessfulEvent) GetElementId() string {
	if x != nil {
		return x.ElementId
	}
	return ""
}

func (x *LMSPipelineSuccessfulEvent) GetElementName() string {
	if x != nil {
		return x.ElementName
	}
	return ""
}

func (x *LMSPipelineSuccessfulEvent) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *LMSPipelineSuccessfulEvent) GetEventMessage() string {
	if x != nil {
		return x.EventMessage
	}
	return ""
}

var File_api_commons_audit_lms_events_proto protoreflect.FileDescriptor

var file_api_commons_audit_lms_events_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x6c, 0x6d, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x17, 0x4c, 0x4d, 0x53, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x18,
	0x4c, 0x4d, 0x53, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x6f, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa0, 0x01, 0x0a,
	0x1a, 0x4c, 0x4d, 0x53, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0xbb, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0x0e, 0x4c, 0x6d, 0x73, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x41, 0xaa,
	0x02, 0x11, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x74, 0xca, 0x02, 0x11, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0xe2, 0x02, 0x1d, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_commons_audit_lms_events_proto_rawDescOnce sync.Once
	file_api_commons_audit_lms_events_proto_rawDescData []byte
)

func file_api_commons_audit_lms_events_proto_rawDescGZIP() []byte {
	file_api_commons_audit_lms_events_proto_rawDescOnce.Do(func() {
		file_api_commons_audit_lms_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_audit_lms_events_proto_rawDesc), len(file_api_commons_audit_lms_events_proto_rawDesc)))
	})
	return file_api_commons_audit_lms_events_proto_rawDescData
}

var file_api_commons_audit_lms_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_commons_audit_lms_events_proto_goTypes = []any{
	(*LMSPipelineFailureEvent)(nil),    // 0: api.commons.audit.LMSPipelineFailureEvent
	(*LMSPipelineNoOutputEvent)(nil),   // 1: api.commons.audit.LMSPipelineNoOutputEvent
	(*LMSPipelineSuccessfulEvent)(nil), // 2: api.commons.audit.LMSPipelineSuccessfulEvent
}
var file_api_commons_audit_lms_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_audit_lms_events_proto_init() }
func file_api_commons_audit_lms_events_proto_init() {
	if File_api_commons_audit_lms_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_audit_lms_events_proto_rawDesc), len(file_api_commons_audit_lms_events_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_audit_lms_events_proto_goTypes,
		DependencyIndexes: file_api_commons_audit_lms_events_proto_depIdxs,
		MessageInfos:      file_api_commons_audit_lms_events_proto_msgTypes,
	}.Build()
	File_api_commons_audit_lms_events_proto = out.File
	file_api_commons_audit_lms_events_proto_goTypes = nil
	file_api_commons_audit_lms_events_proto_depIdxs = nil
}
