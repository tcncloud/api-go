// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/commons/user.proto

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

// Facilitate user filtering on archived state
type UserArchivedStateFilter int32

const (
	// Do not filter (get all users)
	UserArchivedStateFilter_USER_ARCHIVED_STATE_FILTER_ALL UserArchivedStateFilter = 0
	// Filter to only archived user
	UserArchivedStateFilter_USER_ARCHIVED_STATE_FILTER_ARCHIVED UserArchivedStateFilter = 1
	// Filter to only unarchived (active) user
	UserArchivedStateFilter_USER_ARCHIVED_STATE_FILTER_UNARCHIVED UserArchivedStateFilter = 2
)

// Enum value maps for UserArchivedStateFilter.
var (
	UserArchivedStateFilter_name = map[int32]string{
		0: "USER_ARCHIVED_STATE_FILTER_ALL",
		1: "USER_ARCHIVED_STATE_FILTER_ARCHIVED",
		2: "USER_ARCHIVED_STATE_FILTER_UNARCHIVED",
	}
	UserArchivedStateFilter_value = map[string]int32{
		"USER_ARCHIVED_STATE_FILTER_ALL":        0,
		"USER_ARCHIVED_STATE_FILTER_ARCHIVED":   1,
		"USER_ARCHIVED_STATE_FILTER_UNARCHIVED": 2,
	}
)

func (x UserArchivedStateFilter) Enum() *UserArchivedStateFilter {
	p := new(UserArchivedStateFilter)
	*p = x
	return p
}

func (x UserArchivedStateFilter) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserArchivedStateFilter) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_user_proto_enumTypes[0].Descriptor()
}

func (UserArchivedStateFilter) Type() protoreflect.EnumType {
	return &file_api_commons_user_proto_enumTypes[0]
}

func (x UserArchivedStateFilter) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserArchivedStateFilter.Descriptor instead.
func (UserArchivedStateFilter) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_user_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_user_proto protoreflect.FileDescriptor

var file_api_commons_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0x91, 0x01, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56,
	0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f,
	0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x52,
	0x43, 0x48, 0x49, 0x56, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4c,
	0x54, 0x45, 0x52, 0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x29,
	0x0a, 0x25, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x44, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x41,
	0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x44, 0x10, 0x02, 0x42, 0x91, 0x01, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_user_proto_rawDescOnce sync.Once
	file_api_commons_user_proto_rawDescData = file_api_commons_user_proto_rawDesc
)

func file_api_commons_user_proto_rawDescGZIP() []byte {
	file_api_commons_user_proto_rawDescOnce.Do(func() {
		file_api_commons_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_user_proto_rawDescData)
	})
	return file_api_commons_user_proto_rawDescData
}

var file_api_commons_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_user_proto_goTypes = []interface{}{
	(UserArchivedStateFilter)(0), // 0: api.commons.UserArchivedStateFilter
}
var file_api_commons_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_user_proto_init() }
func file_api_commons_user_proto_init() {
	if File_api_commons_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_user_proto_goTypes,
		DependencyIndexes: file_api_commons_user_proto_depIdxs,
		EnumInfos:         file_api_commons_user_proto_enumTypes,
	}.Build()
	File_api_commons_user_proto = out.File
	file_api_commons_user_proto_rawDesc = nil
	file_api_commons_user_proto_goTypes = nil
	file_api_commons_user_proto_depIdxs = nil
}
