// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/commons/task_group.proto

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

type TaskGroupStatus int32

const (
	TaskGroupStatus_TG_UNKNOWN                            TaskGroupStatus = 0
	TaskGroupStatus_TG_PREPARE                            TaskGroupStatus = 1000 // "TG_PREPARE", "Task Group is being prepared for scheduler"),
	TaskGroupStatus_TG_SCHEDULED                          TaskGroupStatus = 1100 // "TG_SCHEDULED", "Task Group is waiting for scheduler"),
	TaskGroupStatus_TG_SCHEDULED_LINKING                  TaskGroupStatus = 1110 // "TG_SCHEDULED_LINKING", "Task Group is scheduling a link campaign"),
	TaskGroupStatus_TG_SCHEDULED_PAUSED                   TaskGroupStatus = 1120 // "TG_SCHEDULED_PAUSED", "Task Group will be set to paused when scheduled"),
	TaskGroupStatus_TG_RUNNING                            TaskGroupStatus = 1200 // "TG_RUNNING", "Task Group is currently executing calls"),
	TaskGroupStatus_TG_PAUSED                             TaskGroupStatus = 1210 // "TG_PAUSED", "Task Group has been signaled to pause calls until further notice"),
	TaskGroupStatus_TG_WAITING                            TaskGroupStatus = 1220 // "TG_WAITING", "Task Group has been signaled to wait until the schedule rules it will allow it to run again"),
	TaskGroupStatus_TG_COMPLETED                          TaskGroupStatus = 1300 // "TG_COMPLETED", "Task Group completed normally"),
	TaskGroupStatus_TG_CANCELLED_TIMEOUT                  TaskGroupStatus = 1310 // "TG_CANCELLED_TIMEOUT", "Task Group was cancelled due to time restrictions"),
	TaskGroupStatus_TG_CANCELLED_USER                     TaskGroupStatus = 1320 // "TG_CANCELLED_USER", "Task Group was cancelled by login belonging to client"),
	TaskGroupStatus_TG_CANCELLED_ADMIN                    TaskGroupStatus = 1330 // "TG_CANCELLED_ADMIN", "Task Group was cancelled by some login not belonging to client with permissions"),
	TaskGroupStatus_TG_SUMMED_COMPLETED                   TaskGroupStatus = 1400 // "TG_SUMMED_COMPLETED", "Task Group completed normally and is summed"),
	TaskGroupStatus_TG_SUMMED_CANCELLED_TIMEOUT           TaskGroupStatus = 1410 // "TG_SUMMED_CANCELLED_TIMEOUT", "Task Group timedout and is summed"),
	TaskGroupStatus_TG_SUMMED_CANCELLED_USER              TaskGroupStatus = 1420 // "TG_SUMMED_CANCELLED_USER", "Task Group was cancelled by login belonging to client and is summed"),
	TaskGroupStatus_TG_SUMMED_CANCELLED_ADMIN             TaskGroupStatus = 1430 // "TG_SUMMED_CANCELLED_ADMIN", "Task Group was cancelled by login not belonging to client and is summed")
	TaskGroupStatus_TG_ACCOUNTINGEXPORT_COMPLETED         TaskGroupStatus = 1500 // "TG_SUMMED_COMPLETED", "Task Group completed normally, summed normally, and has been exported into the accounting packages
	TaskGroupStatus_TG_ACCOUNTINGEXPORT_CANCELLED_TIMEOUT TaskGroupStatus = 1510 // "TG_SUMMED_CANCELLED_TIMEOUT", "Task Group timedout, was summed normally, and has been exported into the accounting packages
	TaskGroupStatus_TG_ACCOUNTINGEXPORT_CANCELLED_USER    TaskGroupStatus = 1520 // "TG_SUMMED_CANCELLED_USER", "Task Group was cancelled by login belonging to client, was summed normally, and has been exported into the accounting packages
	TaskGroupStatus_TG_ACCOUNTINGEXPORT_CANCELLED_ADMIN   TaskGroupStatus = 1530 // "TG_SUMMED_CANCELLED_ADMIN", "Task Group was cancelled by login not belonging to client, was summed normally, and has been exported into the accounting packages
)

// Enum value maps for TaskGroupStatus.
var (
	TaskGroupStatus_name = map[int32]string{
		0:    "TG_UNKNOWN",
		1000: "TG_PREPARE",
		1100: "TG_SCHEDULED",
		1110: "TG_SCHEDULED_LINKING",
		1120: "TG_SCHEDULED_PAUSED",
		1200: "TG_RUNNING",
		1210: "TG_PAUSED",
		1220: "TG_WAITING",
		1300: "TG_COMPLETED",
		1310: "TG_CANCELLED_TIMEOUT",
		1320: "TG_CANCELLED_USER",
		1330: "TG_CANCELLED_ADMIN",
		1400: "TG_SUMMED_COMPLETED",
		1410: "TG_SUMMED_CANCELLED_TIMEOUT",
		1420: "TG_SUMMED_CANCELLED_USER",
		1430: "TG_SUMMED_CANCELLED_ADMIN",
		1500: "TG_ACCOUNTINGEXPORT_COMPLETED",
		1510: "TG_ACCOUNTINGEXPORT_CANCELLED_TIMEOUT",
		1520: "TG_ACCOUNTINGEXPORT_CANCELLED_USER",
		1530: "TG_ACCOUNTINGEXPORT_CANCELLED_ADMIN",
	}
	TaskGroupStatus_value = map[string]int32{
		"TG_UNKNOWN":                            0,
		"TG_PREPARE":                            1000,
		"TG_SCHEDULED":                          1100,
		"TG_SCHEDULED_LINKING":                  1110,
		"TG_SCHEDULED_PAUSED":                   1120,
		"TG_RUNNING":                            1200,
		"TG_PAUSED":                             1210,
		"TG_WAITING":                            1220,
		"TG_COMPLETED":                          1300,
		"TG_CANCELLED_TIMEOUT":                  1310,
		"TG_CANCELLED_USER":                     1320,
		"TG_CANCELLED_ADMIN":                    1330,
		"TG_SUMMED_COMPLETED":                   1400,
		"TG_SUMMED_CANCELLED_TIMEOUT":           1410,
		"TG_SUMMED_CANCELLED_USER":              1420,
		"TG_SUMMED_CANCELLED_ADMIN":             1430,
		"TG_ACCOUNTINGEXPORT_COMPLETED":         1500,
		"TG_ACCOUNTINGEXPORT_CANCELLED_TIMEOUT": 1510,
		"TG_ACCOUNTINGEXPORT_CANCELLED_USER":    1520,
		"TG_ACCOUNTINGEXPORT_CANCELLED_ADMIN":   1530,
	}
)

func (x TaskGroupStatus) Enum() *TaskGroupStatus {
	p := new(TaskGroupStatus)
	*p = x
	return p
}

func (x TaskGroupStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskGroupStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_task_group_proto_enumTypes[0].Descriptor()
}

func (TaskGroupStatus) Type() protoreflect.EnumType {
	return &file_api_commons_task_group_proto_enumTypes[0]
}

func (x TaskGroupStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskGroupStatus.Descriptor instead.
func (TaskGroupStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_task_group_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_task_group_proto protoreflect.FileDescriptor

var file_api_commons_task_group_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0xa9, 0x04, 0x0a, 0x0f,
	0x54, 0x61, 0x73, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0e, 0x0a, 0x0a, 0x54, 0x47, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0a, 0x54, 0x47, 0x5f, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52, 0x45, 0x10, 0xe8, 0x07,
	0x12, 0x11, 0x0a, 0x0c, 0x54, 0x47, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44,
	0x10, 0xcc, 0x08, 0x12, 0x19, 0x0a, 0x14, 0x54, 0x47, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55,
	0x4c, 0x45, 0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0xd6, 0x08, 0x12, 0x18,
	0x0a, 0x13, 0x54, 0x47, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x5f, 0x50,
	0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0xe0, 0x08, 0x12, 0x0f, 0x0a, 0x0a, 0x54, 0x47, 0x5f, 0x52,
	0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0xb0, 0x09, 0x12, 0x0e, 0x0a, 0x09, 0x54, 0x47, 0x5f,
	0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0xba, 0x09, 0x12, 0x0f, 0x0a, 0x0a, 0x54, 0x47, 0x5f,
	0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0xc4, 0x09, 0x12, 0x11, 0x0a, 0x0c, 0x54, 0x47,
	0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x94, 0x0a, 0x12, 0x19, 0x0a,
	0x14, 0x54, 0x47, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x9e, 0x0a, 0x12, 0x16, 0x0a, 0x11, 0x54, 0x47, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0xa8, 0x0a,
	0x12, 0x17, 0x0a, 0x12, 0x54, 0x47, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44,
	0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0xb2, 0x0a, 0x12, 0x18, 0x0a, 0x13, 0x54, 0x47, 0x5f,
	0x53, 0x55, 0x4d, 0x4d, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44,
	0x10, 0xf8, 0x0a, 0x12, 0x20, 0x0a, 0x1b, 0x54, 0x47, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x45, 0x44,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0x82, 0x0b, 0x12, 0x1d, 0x0a, 0x18, 0x54, 0x47, 0x5f, 0x53, 0x55, 0x4d, 0x4d,
	0x45, 0x44, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x10, 0x8c, 0x0b, 0x12, 0x1e, 0x0a, 0x19, 0x54, 0x47, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x45,
	0x44, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x4d, 0x49,
	0x4e, 0x10, 0x96, 0x0b, 0x12, 0x22, 0x0a, 0x1d, 0x54, 0x47, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x49, 0x4e, 0x47, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0xdc, 0x0b, 0x12, 0x2a, 0x0a, 0x25, 0x54, 0x47, 0x5f, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0xe6, 0x0b, 0x12, 0x27, 0x0a, 0x22, 0x54, 0x47, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x49, 0x4e, 0x47, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43,
	0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0xf0, 0x0b, 0x12, 0x28, 0x0a,
	0x23, 0x54, 0x47, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x45, 0x58,
	0x50, 0x4f, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x41,
	0x44, 0x4d, 0x49, 0x4e, 0x10, 0xfa, 0x0b, 0x42, 0x96, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x0e, 0x54, 0x61, 0x73,
	0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_task_group_proto_rawDescOnce sync.Once
	file_api_commons_task_group_proto_rawDescData = file_api_commons_task_group_proto_rawDesc
)

func file_api_commons_task_group_proto_rawDescGZIP() []byte {
	file_api_commons_task_group_proto_rawDescOnce.Do(func() {
		file_api_commons_task_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_task_group_proto_rawDescData)
	})
	return file_api_commons_task_group_proto_rawDescData
}

var file_api_commons_task_group_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_task_group_proto_goTypes = []interface{}{
	(TaskGroupStatus)(0), // 0: api.commons.TaskGroupStatus
}
var file_api_commons_task_group_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_task_group_proto_init() }
func file_api_commons_task_group_proto_init() {
	if File_api_commons_task_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_task_group_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_task_group_proto_goTypes,
		DependencyIndexes: file_api_commons_task_group_proto_depIdxs,
		EnumInfos:         file_api_commons_task_group_proto_enumTypes,
	}.Build()
	File_api_commons_task_group_proto = out.File
	file_api_commons_task_group_proto_rawDesc = nil
	file_api_commons_task_group_proto_goTypes = nil
	file_api_commons_task_group_proto_depIdxs = nil
}
