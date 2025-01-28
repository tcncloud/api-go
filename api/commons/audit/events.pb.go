// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: api/commons/audit/events.proto

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

type DummyEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DummyEvent) Reset() {
	*x = DummyEvent{}
	mi := &file_api_commons_audit_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DummyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyEvent) ProtoMessage() {}

func (x *DummyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyEvent.ProtoReflect.Descriptor instead.
func (*DummyEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_events_proto_rawDescGZIP(), []int{0}
}

func (x *DummyEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_commons_audit_events_proto protoreflect.FileDescriptor

var file_api_commons_audit_events_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x22, 0x26, 0x0a, 0x0a, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0xb8, 0x01, 0x0a, 0x15,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x41, 0xaa, 0x02, 0x11, 0x41, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0xca, 0x02, 0x11, 0x41,
	0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74,
	0xe2, 0x02, 0x1d, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x13, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a,
	0x3a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_commons_audit_events_proto_rawDescOnce sync.Once
	file_api_commons_audit_events_proto_rawDescData []byte
)

func file_api_commons_audit_events_proto_rawDescGZIP() []byte {
	file_api_commons_audit_events_proto_rawDescOnce.Do(func() {
		file_api_commons_audit_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_audit_events_proto_rawDesc), len(file_api_commons_audit_events_proto_rawDesc)))
	})
	return file_api_commons_audit_events_proto_rawDescData
}

var file_api_commons_audit_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_commons_audit_events_proto_goTypes = []any{
	(*DummyEvent)(nil), // 0: api.commons.audit.DummyEvent
}
var file_api_commons_audit_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_audit_events_proto_init() }
func file_api_commons_audit_events_proto_init() {
	if File_api_commons_audit_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_audit_events_proto_rawDesc), len(file_api_commons_audit_events_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_audit_events_proto_goTypes,
		DependencyIndexes: file_api_commons_audit_events_proto_depIdxs,
		MessageInfos:      file_api_commons_audit_events_proto_msgTypes,
	}.Build()
	File_api_commons_audit_events_proto = out.File
	file_api_commons_audit_events_proto_goTypes = nil
	file_api_commons_audit_events_proto_depIdxs = nil
}
