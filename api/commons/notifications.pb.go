// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/commons/notifications.proto

package commons

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

type NotificationType_Enum int32

const (
	NotificationType_INVALID     NotificationType_Enum = 0
	NotificationType_EMAIL       NotificationType_Enum = 1
	NotificationType_SERVER_PUSH NotificationType_Enum = 2
	NotificationType_IOS         NotificationType_Enum = 3
	NotificationType_ANDROID     NotificationType_Enum = 4
	NotificationType_SMS         NotificationType_Enum = 5
)

// Enum value maps for NotificationType_Enum.
var (
	NotificationType_Enum_name = map[int32]string{
		0: "INVALID",
		1: "EMAIL",
		2: "SERVER_PUSH",
		3: "IOS",
		4: "ANDROID",
		5: "SMS",
	}
	NotificationType_Enum_value = map[string]int32{
		"INVALID":     0,
		"EMAIL":       1,
		"SERVER_PUSH": 2,
		"IOS":         3,
		"ANDROID":     4,
		"SMS":         5,
	}
)

func (x NotificationType_Enum) Enum() *NotificationType_Enum {
	p := new(NotificationType_Enum)
	*p = x
	return p
}

func (x NotificationType_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_notifications_proto_enumTypes[0].Descriptor()
}

func (NotificationType_Enum) Type() protoreflect.EnumType {
	return &file_api_commons_notifications_proto_enumTypes[0]
}

func (x NotificationType_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType_Enum.Descriptor instead.
func (NotificationType_Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_notifications_proto_rawDescGZIP(), []int{1, 0}
}

// FieldValueFilter resents a field mask and value to match on.
// For now we are only supporting string matching.
type FieldValueFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key is a valid field mask
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value is the stringified value at the field mask
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FieldValueFilter) Reset() {
	*x = FieldValueFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldValueFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldValueFilter) ProtoMessage() {}

func (x *FieldValueFilter) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldValueFilter.ProtoReflect.Descriptor instead.
func (*FieldValueFilter) Descriptor() ([]byte, []int) {
	return file_api_commons_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *FieldValueFilter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FieldValueFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type NotificationType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotificationType) Reset() {
	*x = NotificationType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_notifications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationType) ProtoMessage() {}

func (x *NotificationType) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_notifications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationType.ProtoReflect.Descriptor instead.
func (*NotificationType) Descriptor() ([]byte, []int) {
	return file_api_commons_notifications_proto_rawDescGZIP(), []int{1}
}

var File_api_commons_notifications_proto protoreflect.FileDescriptor

var file_api_commons_notifications_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x22, 0x3a,
	0x0a, 0x10, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x62, 0x0a, 0x10, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4e,
	0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x50, 0x55, 0x53, 0x48, 0x10, 0x02, 0x12,
	0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44, 0x52,
	0x4f, 0x49, 0x44, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x4d, 0x53, 0x10, 0x05, 0x42, 0x9a,
	0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x42, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_commons_notifications_proto_rawDescOnce sync.Once
	file_api_commons_notifications_proto_rawDescData = file_api_commons_notifications_proto_rawDesc
)

func file_api_commons_notifications_proto_rawDescGZIP() []byte {
	file_api_commons_notifications_proto_rawDescOnce.Do(func() {
		file_api_commons_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_notifications_proto_rawDescData)
	})
	return file_api_commons_notifications_proto_rawDescData
}

var file_api_commons_notifications_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_notifications_proto_goTypes = []interface{}{
	(NotificationType_Enum)(0), // 0: api.commons.NotificationType.Enum
	(*FieldValueFilter)(nil),   // 1: api.commons.FieldValueFilter
	(*NotificationType)(nil),   // 2: api.commons.NotificationType
}
var file_api_commons_notifications_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_notifications_proto_init() }
func file_api_commons_notifications_proto_init() {
	if File_api_commons_notifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commons_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldValueFilter); i {
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
		file_api_commons_notifications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationType); i {
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
			RawDescriptor: file_api_commons_notifications_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_notifications_proto_goTypes,
		DependencyIndexes: file_api_commons_notifications_proto_depIdxs,
		EnumInfos:         file_api_commons_notifications_proto_enumTypes,
		MessageInfos:      file_api_commons_notifications_proto_msgTypes,
	}.Build()
	File_api_commons_notifications_proto = out.File
	file_api_commons_notifications_proto_rawDesc = nil
	file_api_commons_notifications_proto_goTypes = nil
	file_api_commons_notifications_proto_depIdxs = nil
}
