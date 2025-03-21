// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/commons/inbound.proto

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

type InboundGroupStatus int32

const (
	InboundGroupStatus_IBG_UNKNOWN                            InboundGroupStatus = 0    // never use this
	InboundGroupStatus_IBG_PREPARE                            InboundGroupStatus = 5000 // "IBG_PREPARE", "Inbound Group is being prepared for scheduler"),
	InboundGroupStatus_IBG_SCHEDULED                          InboundGroupStatus = 5100 // "IBG_SCHEDULED", "Inbound Group is waiting for scheduler"),
	InboundGroupStatus_IBG_RUNNING                            InboundGroupStatus = 5200 // "IBG_RUNNING", "Inbound Group is currently executing calls"),
	InboundGroupStatus_IBG_PAUSED                             InboundGroupStatus = 5210 // "IBG_PAUSED", "Inbound Group has been signaled to pause calls until further notice"),
	InboundGroupStatus_IBG_COMPLETED                          InboundGroupStatus = 5300 // "IBG_COMPLETED", "Inbound Group completed normally"),
	InboundGroupStatus_IBG_CANCELLED_TIMEOUT                  InboundGroupStatus = 5310 // "IBG_CANCELLED_TIMEOUT", "Inbound Group was cancelled due to time restrictions"),
	InboundGroupStatus_IBG_CANCELLED_USER                     InboundGroupStatus = 5320 // "IBG_CANCELLED_USER", "Inbound Group was cancelled by login belonging to client"),
	InboundGroupStatus_IBG_CANCELLED_ADMIN                    InboundGroupStatus = 5330 // "IBG_CANCELLED_ADMIN", "Inbound Group was cancelled by some login not belonging to client with permissions"),
	InboundGroupStatus_IBG_SUMMED_COMPLETED                   InboundGroupStatus = 5400 // "IBG_SUMMED_COMPLETED", "Inbound Group completed normally and is summed"),
	InboundGroupStatus_IBG_SUMMED_CANCELLED_TIMEOUT           InboundGroupStatus = 5410 // "IBG_SUMMED_CANCELLED_TIMEOUT", "Inbound Group timedout and is summed"),
	InboundGroupStatus_IBG_SUMMED_CANCELLED_USER              InboundGroupStatus = 5420 // "IBG_SUMMED_CANCELLED_USER", "Inbound Group was cancelled by login belonging to client and is summed"),
	InboundGroupStatus_IBG_SUMMED_CANCELLED_ADMIN             InboundGroupStatus = 5430 // "IBG_SUMMED_CANCELLED_ADMIN", "Inbound Group was cancelled by login not belonging to client and is summed"),
	InboundGroupStatus_IBG_ACCOUNTINGEXPORT_COMPLETED         InboundGroupStatus = 5500 // "IBG_SUMMED_COMPLETED", "Inbound Group completed normally, summed normally, and has been exported into the accounting packages
	InboundGroupStatus_IBG_ACCOUNTINGEXPORT_CANCELLED_TIMEOUT InboundGroupStatus = 5510 // "IBG_SUMMED_CANCELLED_TIMEOUT", "Inbound Group timedout, was summed normally, and has been exported into the accounting packages
	InboundGroupStatus_IBG_ACCOUNTINGEXPORT_CANCELLED_USER    InboundGroupStatus = 5520 // "IBG_SUMMED_CANCELLED_USER", "Inbound Group was cancelled by login belonging to client, was summed normally, and has been exported into the accounting packages
	InboundGroupStatus_IBG_ACCOUNTINGEXPORT_CANCELLED_ADMIN   InboundGroupStatus = 5530 // "IBG_SUMMED_CANCELLED_ADMIN", "Inbound Group was cancelled by login not belonging to client, was summed normally, and has been exported into the accounting packages
)

// Enum value maps for InboundGroupStatus.
var (
	InboundGroupStatus_name = map[int32]string{
		0:    "IBG_UNKNOWN",
		5000: "IBG_PREPARE",
		5100: "IBG_SCHEDULED",
		5200: "IBG_RUNNING",
		5210: "IBG_PAUSED",
		5300: "IBG_COMPLETED",
		5310: "IBG_CANCELLED_TIMEOUT",
		5320: "IBG_CANCELLED_USER",
		5330: "IBG_CANCELLED_ADMIN",
		5400: "IBG_SUMMED_COMPLETED",
		5410: "IBG_SUMMED_CANCELLED_TIMEOUT",
		5420: "IBG_SUMMED_CANCELLED_USER",
		5430: "IBG_SUMMED_CANCELLED_ADMIN",
		5500: "IBG_ACCOUNTINGEXPORT_COMPLETED",
		5510: "IBG_ACCOUNTINGEXPORT_CANCELLED_TIMEOUT",
		5520: "IBG_ACCOUNTINGEXPORT_CANCELLED_USER",
		5530: "IBG_ACCOUNTINGEXPORT_CANCELLED_ADMIN",
	}
	InboundGroupStatus_value = map[string]int32{
		"IBG_UNKNOWN":                            0,
		"IBG_PREPARE":                            5000,
		"IBG_SCHEDULED":                          5100,
		"IBG_RUNNING":                            5200,
		"IBG_PAUSED":                             5210,
		"IBG_COMPLETED":                          5300,
		"IBG_CANCELLED_TIMEOUT":                  5310,
		"IBG_CANCELLED_USER":                     5320,
		"IBG_CANCELLED_ADMIN":                    5330,
		"IBG_SUMMED_COMPLETED":                   5400,
		"IBG_SUMMED_CANCELLED_TIMEOUT":           5410,
		"IBG_SUMMED_CANCELLED_USER":              5420,
		"IBG_SUMMED_CANCELLED_ADMIN":             5430,
		"IBG_ACCOUNTINGEXPORT_COMPLETED":         5500,
		"IBG_ACCOUNTINGEXPORT_CANCELLED_TIMEOUT": 5510,
		"IBG_ACCOUNTINGEXPORT_CANCELLED_USER":    5520,
		"IBG_ACCOUNTINGEXPORT_CANCELLED_ADMIN":   5530,
	}
)

func (x InboundGroupStatus) Enum() *InboundGroupStatus {
	p := new(InboundGroupStatus)
	*p = x
	return p
}

func (x InboundGroupStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InboundGroupStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_inbound_proto_enumTypes[0].Descriptor()
}

func (InboundGroupStatus) Type() protoreflect.EnumType {
	return &file_api_commons_inbound_proto_enumTypes[0]
}

func (x InboundGroupStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InboundGroupStatus.Descriptor instead.
func (InboundGroupStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_inbound_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_inbound_proto protoreflect.FileDescriptor

var file_api_commons_inbound_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x6e,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0xf7, 0x03, 0x0a, 0x12, 0x49, 0x6e, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0f, 0x0a, 0x0b, 0x49, 0x42, 0x47, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0b, 0x49, 0x42, 0x47, 0x5f, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52, 0x45, 0x10,
	0x88, 0x27, 0x12, 0x12, 0x0a, 0x0d, 0x49, 0x42, 0x47, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55,
	0x4c, 0x45, 0x44, 0x10, 0xec, 0x27, 0x12, 0x10, 0x0a, 0x0b, 0x49, 0x42, 0x47, 0x5f, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0xd0, 0x28, 0x12, 0x0f, 0x0a, 0x0a, 0x49, 0x42, 0x47, 0x5f,
	0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0xda, 0x28, 0x12, 0x12, 0x0a, 0x0d, 0x49, 0x42, 0x47,
	0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0xb4, 0x29, 0x12, 0x1a, 0x0a,
	0x15, 0x49, 0x42, 0x47, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0xbe, 0x29, 0x12, 0x17, 0x0a, 0x12, 0x49, 0x42, 0x47,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10,
	0xc8, 0x29, 0x12, 0x18, 0x0a, 0x13, 0x49, 0x42, 0x47, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x4c, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0xd2, 0x29, 0x12, 0x19, 0x0a, 0x14,
	0x49, 0x42, 0x47, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x54, 0x45, 0x44, 0x10, 0x98, 0x2a, 0x12, 0x21, 0x0a, 0x1c, 0x49, 0x42, 0x47, 0x5f, 0x53,
	0x55, 0x4d, 0x4d, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0xa2, 0x2a, 0x12, 0x1e, 0x0a, 0x19, 0x49, 0x42,
	0x47, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c,
	0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0xac, 0x2a, 0x12, 0x1f, 0x0a, 0x1a, 0x49, 0x42,
	0x47, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c,
	0x45, 0x44, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0xb6, 0x2a, 0x12, 0x23, 0x0a, 0x1e, 0x49,
	0x42, 0x47, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x45, 0x58, 0x50,
	0x4f, 0x52, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0xfc, 0x2a,
	0x12, 0x2b, 0x0a, 0x26, 0x49, 0x42, 0x47, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49,
	0x4e, 0x47, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c,
	0x45, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x86, 0x2b, 0x12, 0x28, 0x0a,
	0x23, 0x49, 0x42, 0x47, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x45,
	0x58, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x10, 0x90, 0x2b, 0x12, 0x29, 0x0a, 0x24, 0x49, 0x42, 0x47, 0x5f, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10,
	0x9a, 0x2b, 0x42, 0x94, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x0c, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02,
	0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_api_commons_inbound_proto_rawDescOnce sync.Once
	file_api_commons_inbound_proto_rawDescData []byte
)

func file_api_commons_inbound_proto_rawDescGZIP() []byte {
	file_api_commons_inbound_proto_rawDescOnce.Do(func() {
		file_api_commons_inbound_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_inbound_proto_rawDesc), len(file_api_commons_inbound_proto_rawDesc)))
	})
	return file_api_commons_inbound_proto_rawDescData
}

var file_api_commons_inbound_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_inbound_proto_goTypes = []any{
	(InboundGroupStatus)(0), // 0: api.commons.InboundGroupStatus
}
var file_api_commons_inbound_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_inbound_proto_init() }
func file_api_commons_inbound_proto_init() {
	if File_api_commons_inbound_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_inbound_proto_rawDesc), len(file_api_commons_inbound_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_inbound_proto_goTypes,
		DependencyIndexes: file_api_commons_inbound_proto_depIdxs,
		EnumInfos:         file_api_commons_inbound_proto_enumTypes,
	}.Build()
	File_api_commons_inbound_proto = out.File
	file_api_commons_inbound_proto_goTypes = nil
	file_api_commons_inbound_proto_depIdxs = nil
}
