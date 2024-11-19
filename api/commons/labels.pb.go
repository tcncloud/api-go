// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
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

// LabeledEntity is an enum that represents the different types of entities that can be labeled.
type LabeledEntity int32

const (
	// LABELED_ENTITY_UNSPECIFIED is an unspecified entity type.
	LabeledEntity_LABELED_ENTITY_UNSPECIFIED LabeledEntity = 0
	// LABELED_ENTITY_SKILL_GROUP is a skill group entity type.
	LabeledEntity_LABELED_ENTITY_SKILL_GROUP LabeledEntity = 1
	// LABELED_ENTITY_USER is a user entity type.
	LabeledEntity_LABELED_ENTITY_USER LabeledEntity = 2
	// LABELED_ENTITY_TICKET is a ticket entity type.
	LabeledEntity_LABELED_ENTITY_TICKET LabeledEntity = 3
)

// Enum value maps for LabeledEntity.
var (
	LabeledEntity_name = map[int32]string{
		0: "LABELED_ENTITY_UNSPECIFIED",
		1: "LABELED_ENTITY_SKILL_GROUP",
		2: "LABELED_ENTITY_USER",
		3: "LABELED_ENTITY_TICKET",
	}
	LabeledEntity_value = map[string]int32{
		"LABELED_ENTITY_UNSPECIFIED": 0,
		"LABELED_ENTITY_SKILL_GROUP": 1,
		"LABELED_ENTITY_USER":        2,
		"LABELED_ENTITY_TICKET":      3,
	}
)

func (x LabeledEntity) Enum() *LabeledEntity {
	p := new(LabeledEntity)
	*p = x
	return p
}

func (x LabeledEntity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LabeledEntity) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_labels_proto_enumTypes[1].Descriptor()
}

func (LabeledEntity) Type() protoreflect.EnumType {
	return &file_api_commons_labels_proto_enumTypes[1]
}

func (x LabeledEntity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LabeledEntity.Descriptor instead.
func (LabeledEntity) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_labels_proto_rawDescGZIP(), []int{1}
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
	0x10, 0x02, 0x2a, 0x83, 0x01, 0x0a, 0x0d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x1a, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x45, 0x44, 0x5f,
	0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x45, 0x44, 0x5f,
	0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x45, 0x44, 0x5f,
	0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x02, 0x12, 0x19, 0x0a,
	0x15, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x45, 0x44, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f,
	0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x03, 0x42, 0x93, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_api_commons_labels_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_commons_labels_proto_goTypes = []any{
	(EntityType)(0),    // 0: api.commons.EntityType
	(LabeledEntity)(0), // 1: api.commons.LabeledEntity
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
			NumEnums:      2,
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
