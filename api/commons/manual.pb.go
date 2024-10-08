// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/commons/manual.proto

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

type ManualDialGroupStatus int32

const (
	ManualDialGroupStatus_MDG_UNKNOWN                            ManualDialGroupStatus = 0
	ManualDialGroupStatus_MDG_PREPARE                            ManualDialGroupStatus = 6000 // "MDG_PREPARE", "Manual Dial Group is being prepared for scheduler"),
	ManualDialGroupStatus_MDG_SCHEDULED                          ManualDialGroupStatus = 6100 // "MDG_SCHEDULED", "Manual Dial Group is waiting for scheduler"),
	ManualDialGroupStatus_MDG_RUNNING                            ManualDialGroupStatus = 6200 // "MDG_RUNNING", "Manual Dial Group is currently executing calls"),
	ManualDialGroupStatus_MDG_COMPLETED                          ManualDialGroupStatus = 6300 // "MDG_COMPLETED", "Manual Dial Group completed normally"),
	ManualDialGroupStatus_MDG_CANCELLED_TIMEOUT                  ManualDialGroupStatus = 6310 // "MDG_CANCELLED_TIMEOUT", "Manual Dial Group was cancelled due to time restrictions"),
	ManualDialGroupStatus_MDG_CANCELLED_USER                     ManualDialGroupStatus = 6320 // "MDG_CANCELLED_USER", "Manual Dial Group was cancelled by login belonging to client"),
	ManualDialGroupStatus_MDG_CANCELLED_ADMIN                    ManualDialGroupStatus = 6330 // "MDG_CANCELLED_ADMIN", "Manual Dial Group was cancelled by some login not belonging to client with permissions"),
	ManualDialGroupStatus_MDG_SUMMED_COMPLETED                   ManualDialGroupStatus = 6400 // "MDG_SUMMED_COMPLETED", "Manual Dial Group completed normally and is summed"),
	ManualDialGroupStatus_MDG_SUMMED_CANCELLED_TIMEOUT           ManualDialGroupStatus = 6410 // "MDG_SUMMED_CANCELLED_TIMEOUT", "Manual Dial Group timedout and is summed"),
	ManualDialGroupStatus_MDG_SUMMED_CANCELLED_USER              ManualDialGroupStatus = 6420 // "MDG_SUMMED_CANCELLED_USER", "Manual Dial Group was cancelled by login belonging to client and is summed"),
	ManualDialGroupStatus_MDG_SUMMED_CANCELLED_ADMIN             ManualDialGroupStatus = 6430 // "MDG_SUMMED_CANCELLED_ADMIN", "Manual Dial Group was cancelled by login not belonging to client and is summed"),
	ManualDialGroupStatus_MDG_ACCOUNTINGEXPORT_COMPLETED         ManualDialGroupStatus = 6500 // "MDG_SUMMED_COMPLETED", "Manual Dial Group completed normally, summed normally, and has been exported into the accounting packages
	ManualDialGroupStatus_MDG_ACCOUNTINGEXPORT_CANCELLED_TIMEOUT ManualDialGroupStatus = 6510 // "MDG_SUMMED_CANCELLED_TIMEOUT", "Manual Dial Group timedout, was summed normally, and has been exported into the accounting packages
	ManualDialGroupStatus_MDG_ACCOUNTINGEXPORT_CANCELLED_USER    ManualDialGroupStatus = 6520 // "MDG_SUMMED_CANCELLED_USER", "Manual Dial Group was cancelled by login belonging to client, was summed normally, and has been exported into the accounting packages
	ManualDialGroupStatus_MDG_ACCOUNTINGEXPORT_CANCELLED_ADMIN   ManualDialGroupStatus = 6530 // "MDG_SUMMED_CANCELLED_ADMIN", "Manual Dial Group was cancelled by login not belonging to client, was summed normally, and has been exported into the accounting packages
)

// Enum value maps for ManualDialGroupStatus.
var (
	ManualDialGroupStatus_name = map[int32]string{
		0:    "MDG_UNKNOWN",
		6000: "MDG_PREPARE",
		6100: "MDG_SCHEDULED",
		6200: "MDG_RUNNING",
		6300: "MDG_COMPLETED",
		6310: "MDG_CANCELLED_TIMEOUT",
		6320: "MDG_CANCELLED_USER",
		6330: "MDG_CANCELLED_ADMIN",
		6400: "MDG_SUMMED_COMPLETED",
		6410: "MDG_SUMMED_CANCELLED_TIMEOUT",
		6420: "MDG_SUMMED_CANCELLED_USER",
		6430: "MDG_SUMMED_CANCELLED_ADMIN",
		6500: "MDG_ACCOUNTINGEXPORT_COMPLETED",
		6510: "MDG_ACCOUNTINGEXPORT_CANCELLED_TIMEOUT",
		6520: "MDG_ACCOUNTINGEXPORT_CANCELLED_USER",
		6530: "MDG_ACCOUNTINGEXPORT_CANCELLED_ADMIN",
	}
	ManualDialGroupStatus_value = map[string]int32{
		"MDG_UNKNOWN":                            0,
		"MDG_PREPARE":                            6000,
		"MDG_SCHEDULED":                          6100,
		"MDG_RUNNING":                            6200,
		"MDG_COMPLETED":                          6300,
		"MDG_CANCELLED_TIMEOUT":                  6310,
		"MDG_CANCELLED_USER":                     6320,
		"MDG_CANCELLED_ADMIN":                    6330,
		"MDG_SUMMED_COMPLETED":                   6400,
		"MDG_SUMMED_CANCELLED_TIMEOUT":           6410,
		"MDG_SUMMED_CANCELLED_USER":              6420,
		"MDG_SUMMED_CANCELLED_ADMIN":             6430,
		"MDG_ACCOUNTINGEXPORT_COMPLETED":         6500,
		"MDG_ACCOUNTINGEXPORT_CANCELLED_TIMEOUT": 6510,
		"MDG_ACCOUNTINGEXPORT_CANCELLED_USER":    6520,
		"MDG_ACCOUNTINGEXPORT_CANCELLED_ADMIN":   6530,
	}
)

func (x ManualDialGroupStatus) Enum() *ManualDialGroupStatus {
	p := new(ManualDialGroupStatus)
	*p = x
	return p
}

func (x ManualDialGroupStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ManualDialGroupStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_manual_proto_enumTypes[0].Descriptor()
}

func (ManualDialGroupStatus) Type() protoreflect.EnumType {
	return &file_api_commons_manual_proto_enumTypes[0]
}

func (x ManualDialGroupStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ManualDialGroupStatus.Descriptor instead.
func (ManualDialGroupStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_manual_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_manual_proto protoreflect.FileDescriptor

var file_api_commons_manual_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x61,
	0x6e, 0x75, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0xe9, 0x03, 0x0a, 0x15, 0x4d, 0x61, 0x6e, 0x75,
	0x61, 0x6c, 0x44, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x44, 0x47, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0b, 0x4d, 0x44, 0x47, 0x5f, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52,
	0x45, 0x10, 0xf0, 0x2e, 0x12, 0x12, 0x0a, 0x0d, 0x4d, 0x44, 0x47, 0x5f, 0x53, 0x43, 0x48, 0x45,
	0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0xd4, 0x2f, 0x12, 0x10, 0x0a, 0x0b, 0x4d, 0x44, 0x47, 0x5f,
	0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0xb8, 0x30, 0x12, 0x12, 0x0a, 0x0d, 0x4d, 0x44,
	0x47, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x9c, 0x31, 0x12, 0x1a,
	0x0a, 0x15, 0x4d, 0x44, 0x47, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0xa6, 0x31, 0x12, 0x17, 0x0a, 0x12, 0x4d, 0x44,
	0x47, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x10, 0xb0, 0x31, 0x12, 0x18, 0x0a, 0x13, 0x4d, 0x44, 0x47, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0xba, 0x31, 0x12, 0x19, 0x0a,
	0x14, 0x4d, 0x44, 0x47, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x80, 0x32, 0x12, 0x21, 0x0a, 0x1c, 0x4d, 0x44, 0x47, 0x5f,
	0x53, 0x55, 0x4d, 0x4d, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x8a, 0x32, 0x12, 0x1e, 0x0a, 0x19, 0x4d,
	0x44, 0x47, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x4c, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x94, 0x32, 0x12, 0x1f, 0x0a, 0x1a, 0x4d,
	0x44, 0x47, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x4c, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x9e, 0x32, 0x12, 0x23, 0x0a, 0x1e,
	0x4d, 0x44, 0x47, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x45, 0x58,
	0x50, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0xe4,
	0x32, 0x12, 0x2b, 0x0a, 0x26, 0x4d, 0x44, 0x47, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x49, 0x4e, 0x47, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x4c, 0x45, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0xee, 0x32, 0x12, 0x28,
	0x0a, 0x23, 0x4d, 0x44, 0x47, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47,
	0x45, 0x58, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0xf8, 0x32, 0x12, 0x29, 0x0a, 0x24, 0x4d, 0x44, 0x47, 0x5f,
	0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x10, 0x82, 0x33, 0x42, 0x93, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x0b, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02,
	0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_commons_manual_proto_rawDescOnce sync.Once
	file_api_commons_manual_proto_rawDescData = file_api_commons_manual_proto_rawDesc
)

func file_api_commons_manual_proto_rawDescGZIP() []byte {
	file_api_commons_manual_proto_rawDescOnce.Do(func() {
		file_api_commons_manual_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_manual_proto_rawDescData)
	})
	return file_api_commons_manual_proto_rawDescData
}

var file_api_commons_manual_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_manual_proto_goTypes = []any{
	(ManualDialGroupStatus)(0), // 0: api.commons.ManualDialGroupStatus
}
var file_api_commons_manual_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_manual_proto_init() }
func file_api_commons_manual_proto_init() {
	if File_api_commons_manual_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_manual_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_manual_proto_goTypes,
		DependencyIndexes: file_api_commons_manual_proto_depIdxs,
		EnumInfos:         file_api_commons_manual_proto_enumTypes,
	}.Build()
	File_api_commons_manual_proto = out.File
	file_api_commons_manual_proto_rawDesc = nil
	file_api_commons_manual_proto_goTypes = nil
	file_api_commons_manual_proto_depIdxs = nil
}
