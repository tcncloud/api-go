// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/commons/labels.proto

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

// EntityType is used to determine what type of Entity
// a label is assigned to.
// The values of these enums MUST match up with the core.EntityType
// iota in the labels app.
type EntityType int32

const (
	EntityType_ENTITY_TYPE_INVALID      EntityType = 0
	EntityType_ENTITY_TYPE_USER         EntityType = 1
	EntityType_ENTITY_TYPE_LMS_PIPELINE EntityType = 2
)

// Enum value maps for EntityType.
var (
	EntityType_name = map[int32]string{
		0: "ENTITY_TYPE_INVALID",
		1: "ENTITY_TYPE_USER",
		2: "ENTITY_TYPE_LMS_PIPELINE",
	}
	EntityType_value = map[string]int32{
		"ENTITY_TYPE_INVALID":      0,
		"ENTITY_TYPE_USER":         1,
		"ENTITY_TYPE_LMS_PIPELINE": 2,
	}
)

func (x EntityType) Enum() *EntityType {
	p := new(EntityType)
	*p = x
	return p
}

func (x EntityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_labels_proto_enumTypes[0].Descriptor()
}

func (EntityType) Type() protoreflect.EnumType {
	return &file_api_commons_labels_proto_enumTypes[0]
}

func (x EntityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityType.Descriptor instead.
func (EntityType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_labels_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_labels_proto protoreflect.FileDescriptor

var file_api_commons_labels_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0x59, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4d, 0x53, 0x5f, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45,
	0x10, 0x02, 0x42, 0x48, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x63, 0x6e, 0x2e, 0x6d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_labels_proto_rawDescOnce sync.Once
	file_api_commons_labels_proto_rawDescData = file_api_commons_labels_proto_rawDesc
)

func file_api_commons_labels_proto_rawDescGZIP() []byte {
	file_api_commons_labels_proto_rawDescOnce.Do(func() {
		file_api_commons_labels_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_labels_proto_rawDescData)
	})
	return file_api_commons_labels_proto_rawDescData
}

var file_api_commons_labels_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_labels_proto_goTypes = []interface{}{
	(EntityType)(0), // 0: api.commons.EntityType
}
var file_api_commons_labels_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_labels_proto_init() }
func file_api_commons_labels_proto_init() {
	if File_api_commons_labels_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_labels_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_labels_proto_goTypes,
		DependencyIndexes: file_api_commons_labels_proto_depIdxs,
		EnumInfos:         file_api_commons_labels_proto_enumTypes,
	}.Build()
	File_api_commons_labels_proto = out.File
	file_api_commons_labels_proto_rawDesc = nil
	file_api_commons_labels_proto_goTypes = nil
	file_api_commons_labels_proto_depIdxs = nil
}
