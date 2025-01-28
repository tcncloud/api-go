// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: api/commons/services.proto

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

type CloudServices int32

const (
	CloudServices_ORG_SERVICE CloudServices = 0
)

// Enum value maps for CloudServices.
var (
	CloudServices_name = map[int32]string{
		0: "ORG_SERVICE",
	}
	CloudServices_value = map[string]int32{
		"ORG_SERVICE": 0,
	}
)

func (x CloudServices) Enum() *CloudServices {
	p := new(CloudServices)
	*p = x
	return p
}

func (x CloudServices) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudServices) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_services_proto_enumTypes[0].Descriptor()
}

func (CloudServices) Type() protoreflect.EnumType {
	return &file_api_commons_services_proto_enumTypes[0]
}

func (x CloudServices) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CloudServices.Descriptor instead.
func (CloudServices) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_services_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_services_proto protoreflect.FileDescriptor

var file_api_commons_services_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0x20, 0x0a, 0x0d, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x52,
	0x47, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x00, 0x42, 0x95, 0x01, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42,
	0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02,
	0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41,
	0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_commons_services_proto_rawDescOnce sync.Once
	file_api_commons_services_proto_rawDescData []byte
)

func file_api_commons_services_proto_rawDescGZIP() []byte {
	file_api_commons_services_proto_rawDescOnce.Do(func() {
		file_api_commons_services_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_services_proto_rawDesc), len(file_api_commons_services_proto_rawDesc)))
	})
	return file_api_commons_services_proto_rawDescData
}

var file_api_commons_services_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_services_proto_goTypes = []any{
	(CloudServices)(0), // 0: api.commons.CloudServices
}
var file_api_commons_services_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_services_proto_init() }
func file_api_commons_services_proto_init() {
	if File_api_commons_services_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_services_proto_rawDesc), len(file_api_commons_services_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_services_proto_goTypes,
		DependencyIndexes: file_api_commons_services_proto_depIdxs,
		EnumInfos:         file_api_commons_services_proto_enumTypes,
	}.Build()
	File_api_commons_services_proto = out.File
	file_api_commons_services_proto_goTypes = nil
	file_api_commons_services_proto_depIdxs = nil
}
