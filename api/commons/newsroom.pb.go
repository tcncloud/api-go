// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: api/commons/newsroom.proto

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

// newsroom article statuses -
type ArticleStatus int32

const (
	// initial news article status
	ArticleStatus_STATUS_DRAFT ArticleStatus = 0
	// news article has been published
	ArticleStatus_STATUS_PUBLISHED ArticleStatus = 1
	// new article has been archived/closed
	ArticleStatus_STATUS_ARCHIVED ArticleStatus = 2
)

// Enum value maps for ArticleStatus.
var (
	ArticleStatus_name = map[int32]string{
		0: "STATUS_DRAFT",
		1: "STATUS_PUBLISHED",
		2: "STATUS_ARCHIVED",
	}
	ArticleStatus_value = map[string]int32{
		"STATUS_DRAFT":     0,
		"STATUS_PUBLISHED": 1,
		"STATUS_ARCHIVED":  2,
	}
)

func (x ArticleStatus) Enum() *ArticleStatus {
	p := new(ArticleStatus)
	*p = x
	return p
}

func (x ArticleStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArticleStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_newsroom_proto_enumTypes[0].Descriptor()
}

func (ArticleStatus) Type() protoreflect.EnumType {
	return &file_api_commons_newsroom_proto_enumTypes[0]
}

func (x ArticleStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArticleStatus.Descriptor instead.
func (ArticleStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_newsroom_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_newsroom_proto protoreflect.FileDescriptor

var file_api_commons_newsroom_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x65,
	0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0x4c, 0x0a, 0x0d, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x52, 0x43,
	0x48, 0x49, 0x56, 0x45, 0x44, 0x10, 0x02, 0x42, 0x95, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x0d, 0x4e, 0x65, 0x77,
	0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_newsroom_proto_rawDescOnce sync.Once
	file_api_commons_newsroom_proto_rawDescData = file_api_commons_newsroom_proto_rawDesc
)

func file_api_commons_newsroom_proto_rawDescGZIP() []byte {
	file_api_commons_newsroom_proto_rawDescOnce.Do(func() {
		file_api_commons_newsroom_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_newsroom_proto_rawDescData)
	})
	return file_api_commons_newsroom_proto_rawDescData
}

var file_api_commons_newsroom_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_newsroom_proto_goTypes = []any{
	(ArticleStatus)(0), // 0: api.commons.ArticleStatus
}
var file_api_commons_newsroom_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_newsroom_proto_init() }
func file_api_commons_newsroom_proto_init() {
	if File_api_commons_newsroom_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_newsroom_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_newsroom_proto_goTypes,
		DependencyIndexes: file_api_commons_newsroom_proto_depIdxs,
		EnumInfos:         file_api_commons_newsroom_proto_enumTypes,
	}.Build()
	File_api_commons_newsroom_proto = out.File
	file_api_commons_newsroom_proto_rawDesc = nil
	file_api_commons_newsroom_proto_goTypes = nil
	file_api_commons_newsroom_proto_depIdxs = nil
}
