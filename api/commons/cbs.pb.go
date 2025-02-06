// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/commons/cbs.proto

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

type ScheduledCallbackStatus int32

const (
	ScheduledCallbackStatus_SCS_INVALID  ScheduledCallbackStatus = 0
	ScheduledCallbackStatus_SCS_OPENED   ScheduledCallbackStatus = 1
	ScheduledCallbackStatus_SCS_CANCELED ScheduledCallbackStatus = 2
	ScheduledCallbackStatus_SCS_CLOSED   ScheduledCallbackStatus = 3
	ScheduledCallbackStatus_SCS_READY    ScheduledCallbackStatus = 4
)

// Enum value maps for ScheduledCallbackStatus.
var (
	ScheduledCallbackStatus_name = map[int32]string{
		0: "SCS_INVALID",
		1: "SCS_OPENED",
		2: "SCS_CANCELED",
		3: "SCS_CLOSED",
		4: "SCS_READY",
	}
	ScheduledCallbackStatus_value = map[string]int32{
		"SCS_INVALID":  0,
		"SCS_OPENED":   1,
		"SCS_CANCELED": 2,
		"SCS_CLOSED":   3,
		"SCS_READY":    4,
	}
)

func (x ScheduledCallbackStatus) Enum() *ScheduledCallbackStatus {
	p := new(ScheduledCallbackStatus)
	*p = x
	return p
}

func (x ScheduledCallbackStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScheduledCallbackStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_cbs_proto_enumTypes[0].Descriptor()
}

func (ScheduledCallbackStatus) Type() protoreflect.EnumType {
	return &file_api_commons_cbs_proto_enumTypes[0]
}

func (x ScheduledCallbackStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScheduledCallbackStatus.Descriptor instead.
func (ScheduledCallbackStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_cbs_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_cbs_proto protoreflect.FileDescriptor

var file_api_commons_cbs_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x62,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0x6b, 0x0a, 0x17, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x64, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x43, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x43, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x43, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x43, 0x53, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x43, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10,
	0x04, 0x42, 0x90, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x08, 0x43, 0x62, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
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
	file_api_commons_cbs_proto_rawDescOnce sync.Once
	file_api_commons_cbs_proto_rawDescData []byte
)

func file_api_commons_cbs_proto_rawDescGZIP() []byte {
	file_api_commons_cbs_proto_rawDescOnce.Do(func() {
		file_api_commons_cbs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_cbs_proto_rawDesc), len(file_api_commons_cbs_proto_rawDesc)))
	})
	return file_api_commons_cbs_proto_rawDescData
}

var file_api_commons_cbs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_cbs_proto_goTypes = []any{
	(ScheduledCallbackStatus)(0), // 0: api.commons.ScheduledCallbackStatus
}
var file_api_commons_cbs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_cbs_proto_init() }
func file_api_commons_cbs_proto_init() {
	if File_api_commons_cbs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_cbs_proto_rawDesc), len(file_api_commons_cbs_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_cbs_proto_goTypes,
		DependencyIndexes: file_api_commons_cbs_proto_depIdxs,
		EnumInfos:         file_api_commons_cbs_proto_enumTypes,
	}.Build()
	File_api_commons_cbs_proto = out.File
	file_api_commons_cbs_proto_goTypes = nil
	file_api_commons_cbs_proto_depIdxs = nil
}
