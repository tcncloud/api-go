// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/commons/fts.proto

package commons

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

type AppName int32

const (
	AppName_OUTSIDE      AppName = 0
	AppName_LMS          AppName = 1
	AppName_P3API        AppName = 2
	AppName_ANA          AppName = 3
	AppName_COMPLIANCE   AppName = 4
	AppName_VMDS         AppName = 5
	AppName_EMAIL        AppName = 6
	AppName_CJS          AppName = 7
	AppName_VANALYTICS   AppName = 8
	AppName_INTEGRATIONS AppName = 9
	AppName_OMNICHANNEL  AppName = 10
	AppName_LEARN        AppName = 11
	AppName_BILLING      AppName = 12
	AppName_NEWSROOM     AppName = 13
)

// Enum value maps for AppName.
var (
	AppName_name = map[int32]string{
		0:  "OUTSIDE",
		1:  "LMS",
		2:  "P3API",
		3:  "ANA",
		4:  "COMPLIANCE",
		5:  "VMDS",
		6:  "EMAIL",
		7:  "CJS",
		8:  "VANALYTICS",
		9:  "INTEGRATIONS",
		10: "OMNICHANNEL",
		11: "LEARN",
		12: "BILLING",
		13: "NEWSROOM",
	}
	AppName_value = map[string]int32{
		"OUTSIDE":      0,
		"LMS":          1,
		"P3API":        2,
		"ANA":          3,
		"COMPLIANCE":   4,
		"VMDS":         5,
		"EMAIL":        6,
		"CJS":          7,
		"VANALYTICS":   8,
		"INTEGRATIONS": 9,
		"OMNICHANNEL":  10,
		"LEARN":        11,
		"BILLING":      12,
		"NEWSROOM":     13,
	}
)

func (x AppName) Enum() *AppName {
	p := new(AppName)
	*p = x
	return p
}

func (x AppName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppName) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_fts_proto_enumTypes[0].Descriptor()
}

func (AppName) Type() protoreflect.EnumType {
	return &file_api_commons_fts_proto_enumTypes[0]
}

func (x AppName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppName.Descriptor instead.
func (AppName) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_fts_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_fts_proto protoreflect.FileDescriptor

var file_api_commons_fts_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0xba, 0x01, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x55, 0x54, 0x53, 0x49, 0x44, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x4c, 0x4d, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x33, 0x41, 0x50, 0x49, 0x10,
	0x02, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x41, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x4d,
	0x44, 0x53, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x06, 0x12,
	0x07, 0x0a, 0x03, 0x43, 0x4a, 0x53, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x41, 0x4e, 0x41,
	0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x54, 0x45,
	0x47, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x4d,
	0x4e, 0x49, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x4c,
	0x45, 0x41, 0x52, 0x4e, 0x10, 0x0b, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e,
	0x47, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x45, 0x57, 0x53, 0x52, 0x4f, 0x4f, 0x4d, 0x10,
	0x0d, 0x42, 0x90, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x08, 0x46, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa,
	0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b,
	0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70,
	0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_commons_fts_proto_rawDescOnce sync.Once
	file_api_commons_fts_proto_rawDescData []byte
)

func file_api_commons_fts_proto_rawDescGZIP() []byte {
	file_api_commons_fts_proto_rawDescOnce.Do(func() {
		file_api_commons_fts_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_fts_proto_rawDesc), len(file_api_commons_fts_proto_rawDesc)))
	})
	return file_api_commons_fts_proto_rawDescData
}

var file_api_commons_fts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_fts_proto_goTypes = []any{
	(AppName)(0), // 0: api.commons.AppName
}
var file_api_commons_fts_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_fts_proto_init() }
func file_api_commons_fts_proto_init() {
	if File_api_commons_fts_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_fts_proto_rawDesc), len(file_api_commons_fts_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_fts_proto_goTypes,
		DependencyIndexes: file_api_commons_fts_proto_depIdxs,
		EnumInfos:         file_api_commons_fts_proto_enumTypes,
	}.Build()
	File_api_commons_fts_proto = out.File
	file_api_commons_fts_proto_goTypes = nil
	file_api_commons_fts_proto_depIdxs = nil
}
