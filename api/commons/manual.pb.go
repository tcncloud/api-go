// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/commons/manual.proto

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

const file_api_commons_manual_proto_rawDesc = "" +
	"\n" +
	"\x18api/commons/manual.proto\x12\vapi.commons*\xe9\x03\n" +
	"\x15ManualDialGroupStatus\x12\x0f\n" +
	"\vMDG_UNKNOWN\x10\x00\x12\x10\n" +
	"\vMDG_PREPARE\x10\xf0.\x12\x12\n" +
	"\rMDG_SCHEDULED\x10\xd4/\x12\x10\n" +
	"\vMDG_RUNNING\x10\xb80\x12\x12\n" +
	"\rMDG_COMPLETED\x10\x9c1\x12\x1a\n" +
	"\x15MDG_CANCELLED_TIMEOUT\x10\xa61\x12\x17\n" +
	"\x12MDG_CANCELLED_USER\x10\xb01\x12\x18\n" +
	"\x13MDG_CANCELLED_ADMIN\x10\xba1\x12\x19\n" +
	"\x14MDG_SUMMED_COMPLETED\x10\x802\x12!\n" +
	"\x1cMDG_SUMMED_CANCELLED_TIMEOUT\x10\x8a2\x12\x1e\n" +
	"\x19MDG_SUMMED_CANCELLED_USER\x10\x942\x12\x1f\n" +
	"\x1aMDG_SUMMED_CANCELLED_ADMIN\x10\x9e2\x12#\n" +
	"\x1eMDG_ACCOUNTINGEXPORT_COMPLETED\x10\xe42\x12+\n" +
	"&MDG_ACCOUNTINGEXPORT_CANCELLED_TIMEOUT\x10\xee2\x12(\n" +
	"#MDG_ACCOUNTINGEXPORT_CANCELLED_USER\x10\xf82\x12)\n" +
	"$MDG_ACCOUNTINGEXPORT_CANCELLED_ADMIN\x10\x823B\x93\x01\n" +
	"\x0fcom.api.commonsB\vManualProtoP\x01Z&github.com/tcncloud/api-go/api/commons\xa2\x02\x03ACX\xaa\x02\vApi.Commons\xca\x02\vApi\\Commons\xe2\x02\x17Api\\Commons\\GPBMetadata\xea\x02\fApi::Commonsb\x06proto3"

var (
	file_api_commons_manual_proto_rawDescOnce sync.Once
	file_api_commons_manual_proto_rawDescData []byte
)

func file_api_commons_manual_proto_rawDescGZIP() []byte {
	file_api_commons_manual_proto_rawDescOnce.Do(func() {
		file_api_commons_manual_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_manual_proto_rawDesc), len(file_api_commons_manual_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_manual_proto_rawDesc), len(file_api_commons_manual_proto_rawDesc)),
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
	file_api_commons_manual_proto_goTypes = nil
	file_api_commons_manual_proto_depIdxs = nil
}
