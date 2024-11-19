// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/commons/insights.proto

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

type InsightPermissionType int32

const (
	InsightPermissionType_INSIGHT_PERMISSION_TYPE_COMMON_LIBRARY InsightPermissionType = 0
	InsightPermissionType_INSIGHT_PERMISSION_TYPE_ORG            InsightPermissionType = 1
)

// Enum value maps for InsightPermissionType.
var (
	InsightPermissionType_name = map[int32]string{
		0: "INSIGHT_PERMISSION_TYPE_COMMON_LIBRARY",
		1: "INSIGHT_PERMISSION_TYPE_ORG",
	}
	InsightPermissionType_value = map[string]int32{
		"INSIGHT_PERMISSION_TYPE_COMMON_LIBRARY": 0,
		"INSIGHT_PERMISSION_TYPE_ORG":            1,
	}
)

func (x InsightPermissionType) Enum() *InsightPermissionType {
	p := new(InsightPermissionType)
	*p = x
	return p
}

func (x InsightPermissionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsightPermissionType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_insights_proto_enumTypes[0].Descriptor()
}

func (InsightPermissionType) Type() protoreflect.EnumType {
	return &file_api_commons_insights_proto_enumTypes[0]
}

func (x InsightPermissionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsightPermissionType.Descriptor instead.
func (InsightPermissionType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_insights_proto_rawDescGZIP(), []int{0}
}

type InsightType int32

const (
	InsightType_INSIGHT_TYPE_TABLE_VIEW InsightType = 0
)

// Enum value maps for InsightType.
var (
	InsightType_name = map[int32]string{
		0: "INSIGHT_TYPE_TABLE_VIEW",
	}
	InsightType_value = map[string]int32{
		"INSIGHT_TYPE_TABLE_VIEW": 0,
	}
)

func (x InsightType) Enum() *InsightType {
	p := new(InsightType)
	*p = x
	return p
}

func (x InsightType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsightType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_insights_proto_enumTypes[1].Descriptor()
}

func (InsightType) Type() protoreflect.EnumType {
	return &file_api_commons_insights_proto_enumTypes[1]
}

func (x InsightType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsightType.Descriptor instead.
func (InsightType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_insights_proto_rawDescGZIP(), []int{1}
}

// InsightVfsSchemaType is enum for column types in a vfs
type InsightVfsSchemaType int32

const (
	InsightVfsSchemaType_INSIGHT_VFS_SCHEMA_TYPE_STRING   InsightVfsSchemaType = 0 // data type for string
	InsightVfsSchemaType_INSIGHT_VFS_SCHEMA_TYPE_INT64    InsightVfsSchemaType = 1 // data type for int
	InsightVfsSchemaType_INSIGHT_VFS_SCHEMA_TYPE_FLOAT64  InsightVfsSchemaType = 2 // data type for float
	InsightVfsSchemaType_INSIGHT_VFS_SCHEMA_TYPE_BOOLEAN  InsightVfsSchemaType = 3 // data type for bool
	InsightVfsSchemaType_INSIGHT_VFS_SCHEMA_TYPE_DATETIME InsightVfsSchemaType = 4 // data type for datetime
)

// Enum value maps for InsightVfsSchemaType.
var (
	InsightVfsSchemaType_name = map[int32]string{
		0: "INSIGHT_VFS_SCHEMA_TYPE_STRING",
		1: "INSIGHT_VFS_SCHEMA_TYPE_INT64",
		2: "INSIGHT_VFS_SCHEMA_TYPE_FLOAT64",
		3: "INSIGHT_VFS_SCHEMA_TYPE_BOOLEAN",
		4: "INSIGHT_VFS_SCHEMA_TYPE_DATETIME",
	}
	InsightVfsSchemaType_value = map[string]int32{
		"INSIGHT_VFS_SCHEMA_TYPE_STRING":   0,
		"INSIGHT_VFS_SCHEMA_TYPE_INT64":    1,
		"INSIGHT_VFS_SCHEMA_TYPE_FLOAT64":  2,
		"INSIGHT_VFS_SCHEMA_TYPE_BOOLEAN":  3,
		"INSIGHT_VFS_SCHEMA_TYPE_DATETIME": 4,
	}
)

func (x InsightVfsSchemaType) Enum() *InsightVfsSchemaType {
	p := new(InsightVfsSchemaType)
	*p = x
	return p
}

func (x InsightVfsSchemaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsightVfsSchemaType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_insights_proto_enumTypes[2].Descriptor()
}

func (InsightVfsSchemaType) Type() protoreflect.EnumType {
	return &file_api_commons_insights_proto_enumTypes[2]
}

func (x InsightVfsSchemaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsightVfsSchemaType.Descriptor instead.
func (InsightVfsSchemaType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_insights_proto_rawDescGZIP(), []int{2}
}

var File_api_commons_insights_proto protoreflect.FileDescriptor

var file_api_commons_insights_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0x64, 0x0a, 0x15, 0x49, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2a, 0x0a, 0x26, 0x49, 0x4e, 0x53, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f,
	0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x42, 0x52, 0x41, 0x52, 0x59, 0x10, 0x00, 0x12, 0x1f,
	0x0a, 0x1b, 0x49, 0x4e, 0x53, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52, 0x47, 0x10, 0x01, 0x2a,
	0x2a, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x49, 0x4e, 0x53, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x41, 0x42, 0x4c, 0x45, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x10, 0x00, 0x2a, 0xcd, 0x01, 0x0a, 0x14,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x56, 0x66, 0x73, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x4e, 0x53, 0x49, 0x47, 0x48, 0x54, 0x5f,
	0x56, 0x46, 0x53, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x4e, 0x53, 0x49,
	0x47, 0x48, 0x54, 0x5f, 0x56, 0x46, 0x53, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x49,
	0x4e, 0x53, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x56, 0x46, 0x53, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d,
	0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x02,
	0x12, 0x23, 0x0a, 0x1f, 0x49, 0x4e, 0x53, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x56, 0x46, 0x53, 0x5f,
	0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c,
	0x45, 0x41, 0x4e, 0x10, 0x03, 0x12, 0x24, 0x0a, 0x20, 0x49, 0x4e, 0x53, 0x49, 0x47, 0x48, 0x54,
	0x5f, 0x56, 0x46, 0x53, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x04, 0x42, 0x95, 0x01, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42,
	0x0d, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02,
	0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41,
	0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_insights_proto_rawDescOnce sync.Once
	file_api_commons_insights_proto_rawDescData = file_api_commons_insights_proto_rawDesc
)

func file_api_commons_insights_proto_rawDescGZIP() []byte {
	file_api_commons_insights_proto_rawDescOnce.Do(func() {
		file_api_commons_insights_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_insights_proto_rawDescData)
	})
	return file_api_commons_insights_proto_rawDescData
}

var file_api_commons_insights_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_commons_insights_proto_goTypes = []any{
	(InsightPermissionType)(0), // 0: api.commons.InsightPermissionType
	(InsightType)(0),           // 1: api.commons.InsightType
	(InsightVfsSchemaType)(0),  // 2: api.commons.InsightVfsSchemaType
}
var file_api_commons_insights_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_insights_proto_init() }
func file_api_commons_insights_proto_init() {
	if File_api_commons_insights_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_insights_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_insights_proto_goTypes,
		DependencyIndexes: file_api_commons_insights_proto_depIdxs,
		EnumInfos:         file_api_commons_insights_proto_enumTypes,
	}.Build()
	File_api_commons_insights_proto = out.File
	file_api_commons_insights_proto_rawDesc = nil
	file_api_commons_insights_proto_goTypes = nil
	file_api_commons_insights_proto_depIdxs = nil
}
