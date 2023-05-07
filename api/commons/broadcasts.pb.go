// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/commons/broadcasts.proto

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

type BroadcastType int32

const (
	BroadcastType_TYPE_UNKNOWN                     BroadcastType = 0
	BroadcastType_TYPE_OUTBOUND                    BroadcastType = 1000
	BroadcastType_TYPE_OUTBOUND_PREVIEW_ONLY       BroadcastType = 1100
	BroadcastType_TYPE_OUTBOUND_MAC_ONLY           BroadcastType = 1200
	BroadcastType_TYPE_OUTBOUND_RINGLESS_VOICEMAIL BroadcastType = 1300
	BroadcastType_TYPE_INBOUND                     BroadcastType = 2000
	BroadcastType_TYPE_MANUAL                      BroadcastType = 3000
	BroadcastType_TYPE_SMS                         BroadcastType = 4000
	BroadcastType_TYPE_EMAIL                       BroadcastType = 5000
	BroadcastType_TYPE_OUTBOUND_INBOUND            BroadcastType = 7000
	BroadcastType_TYPE_OUTBOUND_MANUAL             BroadcastType = 8000
	BroadcastType_TYPE_INBOUND_MANUAL              BroadcastType = 9000
	BroadcastType_TYPE_OUTBOUND_INBOUND_MANUAL     BroadcastType = 10000
)

// Enum value maps for BroadcastType.
var (
	BroadcastType_name = map[int32]string{
		0:     "TYPE_UNKNOWN",
		1000:  "TYPE_OUTBOUND",
		1100:  "TYPE_OUTBOUND_PREVIEW_ONLY",
		1200:  "TYPE_OUTBOUND_MAC_ONLY",
		1300:  "TYPE_OUTBOUND_RINGLESS_VOICEMAIL",
		2000:  "TYPE_INBOUND",
		3000:  "TYPE_MANUAL",
		4000:  "TYPE_SMS",
		5000:  "TYPE_EMAIL",
		7000:  "TYPE_OUTBOUND_INBOUND",
		8000:  "TYPE_OUTBOUND_MANUAL",
		9000:  "TYPE_INBOUND_MANUAL",
		10000: "TYPE_OUTBOUND_INBOUND_MANUAL",
	}
	BroadcastType_value = map[string]int32{
		"TYPE_UNKNOWN":                     0,
		"TYPE_OUTBOUND":                    1000,
		"TYPE_OUTBOUND_PREVIEW_ONLY":       1100,
		"TYPE_OUTBOUND_MAC_ONLY":           1200,
		"TYPE_OUTBOUND_RINGLESS_VOICEMAIL": 1300,
		"TYPE_INBOUND":                     2000,
		"TYPE_MANUAL":                      3000,
		"TYPE_SMS":                         4000,
		"TYPE_EMAIL":                       5000,
		"TYPE_OUTBOUND_INBOUND":            7000,
		"TYPE_OUTBOUND_MANUAL":             8000,
		"TYPE_INBOUND_MANUAL":              9000,
		"TYPE_OUTBOUND_INBOUND_MANUAL":     10000,
	}
)

func (x BroadcastType) Enum() *BroadcastType {
	p := new(BroadcastType)
	*p = x
	return p
}

func (x BroadcastType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BroadcastType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_broadcasts_proto_enumTypes[0].Descriptor()
}

func (BroadcastType) Type() protoreflect.EnumType {
	return &file_api_commons_broadcasts_proto_enumTypes[0]
}

func (x BroadcastType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BroadcastType.Descriptor instead.
func (BroadcastType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_broadcasts_proto_rawDescGZIP(), []int{0}
}

type TemplateType_Enum int32

const (
	TemplateType_UNKNOWN            TemplateType_Enum = 0
	TemplateType_STANDARD           TemplateType_Enum = 1
	TemplateType_LAYERED            TemplateType_Enum = 2
	TemplateType_INBOUND            TemplateType_Enum = 3
	TemplateType_MAC_ONLY           TemplateType_Enum = 4
	TemplateType_PREVIEW_ONLY       TemplateType_Enum = 5
	TemplateType_RINGLESS_VOICEMAIL TemplateType_Enum = 6
)

// Enum value maps for TemplateType_Enum.
var (
	TemplateType_Enum_name = map[int32]string{
		0: "UNKNOWN",
		1: "STANDARD",
		2: "LAYERED",
		3: "INBOUND",
		4: "MAC_ONLY",
		5: "PREVIEW_ONLY",
		6: "RINGLESS_VOICEMAIL",
	}
	TemplateType_Enum_value = map[string]int32{
		"UNKNOWN":            0,
		"STANDARD":           1,
		"LAYERED":            2,
		"INBOUND":            3,
		"MAC_ONLY":           4,
		"PREVIEW_ONLY":       5,
		"RINGLESS_VOICEMAIL": 6,
	}
)

func (x TemplateType_Enum) Enum() *TemplateType_Enum {
	p := new(TemplateType_Enum)
	*p = x
	return p
}

func (x TemplateType_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TemplateType_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_broadcasts_proto_enumTypes[1].Descriptor()
}

func (TemplateType_Enum) Type() protoreflect.EnumType {
	return &file_api_commons_broadcasts_proto_enumTypes[1]
}

func (x TemplateType_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TemplateType_Enum.Descriptor instead.
func (TemplateType_Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_broadcasts_proto_rawDescGZIP(), []int{0, 0}
}

// Enum to represent the different template types that broadcast templates can be.
type TemplateType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TemplateType) Reset() {
	*x = TemplateType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_broadcasts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateType) ProtoMessage() {}

func (x *TemplateType) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_broadcasts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateType.ProtoReflect.Descriptor instead.
func (*TemplateType) Descriptor() ([]byte, []int) {
	return file_api_commons_broadcasts_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_broadcasts_proto protoreflect.FileDescriptor

var file_api_commons_broadcasts_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x0c,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x73, 0x0a, 0x04,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x41, 0x43,
	0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x45, 0x56, 0x49,
	0x45, 0x57, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x49, 0x4e,
	0x47, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10,
	0x06, 0x2a, 0xd3, 0x02, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x55,
	0x54, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xe8, 0x07, 0x12, 0x1f, 0x0a, 0x1a, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4f, 0x55, 0x54, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49,
	0x45, 0x57, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0xcc, 0x08, 0x12, 0x1b, 0x0a, 0x16, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x4d, 0x41, 0x43, 0x5f,
	0x4f, 0x4e, 0x4c, 0x59, 0x10, 0xb0, 0x09, 0x12, 0x25, 0x0a, 0x20, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4f, 0x55, 0x54, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x52, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x53,
	0x53, 0x5f, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x94, 0x0a, 0x12, 0x11,
	0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xd0,
	0x0f, 0x12, 0x10, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c,
	0x10, 0xb8, 0x17, 0x12, 0x0d, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4d, 0x53, 0x10,
	0xa0, 0x1f, 0x12, 0x0f, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c,
	0x10, 0x88, 0x27, 0x12, 0x1a, 0x0a, 0x15, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x42,
	0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xd8, 0x36, 0x12,
	0x19, 0x0a, 0x14, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x42, 0x4f, 0x55, 0x4e, 0x44,
	0x5f, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x10, 0xc0, 0x3e, 0x12, 0x18, 0x0a, 0x13, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x4d, 0x41, 0x4e, 0x55, 0x41,
	0x4c, 0x10, 0xa8, 0x46, 0x12, 0x21, 0x0a, 0x1c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x55, 0x54,
	0x42, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x4d, 0x41,
	0x4e, 0x55, 0x41, 0x4c, 0x10, 0x90, 0x4e, 0x42, 0x97, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x0f, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_broadcasts_proto_rawDescOnce sync.Once
	file_api_commons_broadcasts_proto_rawDescData = file_api_commons_broadcasts_proto_rawDesc
)

func file_api_commons_broadcasts_proto_rawDescGZIP() []byte {
	file_api_commons_broadcasts_proto_rawDescOnce.Do(func() {
		file_api_commons_broadcasts_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_broadcasts_proto_rawDescData)
	})
	return file_api_commons_broadcasts_proto_rawDescData
}

var file_api_commons_broadcasts_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_commons_broadcasts_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_commons_broadcasts_proto_goTypes = []interface{}{
	(BroadcastType)(0),     // 0: api.commons.BroadcastType
	(TemplateType_Enum)(0), // 1: api.commons.TemplateType.Enum
	(*TemplateType)(nil),   // 2: api.commons.TemplateType
}
var file_api_commons_broadcasts_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_broadcasts_proto_init() }
func file_api_commons_broadcasts_proto_init() {
	if File_api_commons_broadcasts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commons_broadcasts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateType); i {
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
			RawDescriptor: file_api_commons_broadcasts_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_broadcasts_proto_goTypes,
		DependencyIndexes: file_api_commons_broadcasts_proto_depIdxs,
		EnumInfos:         file_api_commons_broadcasts_proto_enumTypes,
		MessageInfos:      file_api_commons_broadcasts_proto_msgTypes,
	}.Build()
	File_api_commons_broadcasts_proto = out.File
	file_api_commons_broadcasts_proto_rawDesc = nil
	file_api_commons_broadcasts_proto_goTypes = nil
	file_api_commons_broadcasts_proto_depIdxs = nil
}
