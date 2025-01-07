// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: api/commons/logging.proto

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

type Level int32

const (
	Level_OFF    Level = 0
	Level_FATAL  Level = 1
	Level_PANIC  Level = 2
	Level_DPANIC Level = 3
	Level_ERROR  Level = 4
	Level_WARN   Level = 5
	Level_INFO   Level = 6
	Level_DEBUG  Level = 7
	Level_TRACE  Level = 8
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "OFF",
		1: "FATAL",
		2: "PANIC",
		3: "DPANIC",
		4: "ERROR",
		5: "WARN",
		6: "INFO",
		7: "DEBUG",
		8: "TRACE",
	}
	Level_value = map[string]int32{
		"OFF":    0,
		"FATAL":  1,
		"PANIC":  2,
		"DPANIC": 3,
		"ERROR":  4,
		"WARN":   5,
		"INFO":   6,
		"DEBUG":  7,
		"TRACE":  8,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_logging_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_api_commons_logging_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_logging_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_logging_proto protoreflect.FileDescriptor

var file_api_commons_logging_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0x67, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x46, 0x46, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41,
	0x54, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x4e, 0x49, 0x43, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x44, 0x50, 0x41, 0x4e, 0x49, 0x43, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10,
	0x05, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x44,
	0x45, 0x42, 0x55, 0x47, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10,
	0x08, 0x42, 0x94, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x0c, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03,
	0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2,
	0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_logging_proto_rawDescOnce sync.Once
	file_api_commons_logging_proto_rawDescData = file_api_commons_logging_proto_rawDesc
)

func file_api_commons_logging_proto_rawDescGZIP() []byte {
	file_api_commons_logging_proto_rawDescOnce.Do(func() {
		file_api_commons_logging_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_logging_proto_rawDescData)
	})
	return file_api_commons_logging_proto_rawDescData
}

var file_api_commons_logging_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_logging_proto_goTypes = []any{
	(Level)(0), // 0: api.commons.Level
}
var file_api_commons_logging_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_logging_proto_init() }
func file_api_commons_logging_proto_init() {
	if File_api_commons_logging_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_logging_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_logging_proto_goTypes,
		DependencyIndexes: file_api_commons_logging_proto_depIdxs,
		EnumInfos:         file_api_commons_logging_proto_enumTypes,
	}.Build()
	File_api_commons_logging_proto = out.File
	file_api_commons_logging_proto_rawDesc = nil
	file_api_commons_logging_proto_goTypes = nil
	file_api_commons_logging_proto_depIdxs = nil
}
