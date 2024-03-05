// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/commons/task.proto

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

type TaskStatus int32

const (
	TaskStatus_TASK_UNKNOWN           TaskStatus = 0
	TaskStatus_TASK_SCHEDULED         TaskStatus = 2100 // "TASK_SCHEDULED", "Task is waiting for scheduler"),
	TaskStatus_TASK_WAITING           TaskStatus = 2110 // "TASK_WAITING", "Task was running and is waiting to send next call"),
	TaskStatus_TASK_PREPARING         TaskStatus = 2120 // "TASK_PREPARING", " Task is currently beeing prepared by the scheduler"),
	TaskStatus_TASK_RUNNING           TaskStatus = 2200 // "TASK_RUNNING", "Task is currently executing calls"),
	TaskStatus_TASK_COMPLETED         TaskStatus = 2300 // "TASK_COMPLETED", "Task was completed normally"),
	TaskStatus_TASK_CANCELLED_TIMEOUT TaskStatus = 2310 // "TASK_CANCELLED_TIMEOUT", "Task was cancelled due to time restrictions"),
	TaskStatus_TASK_CANCELLED_USER    TaskStatus = 2320 // "TASK_CANCELLED_USER", "Task was cancelled by login belonging to client"),
	TaskStatus_TASK_CANCELLED_ADMIN   TaskStatus = 2330 // "TASK_CANCELLED_ADMIN", "Task was cancelled by some login not belonging to client with permissions"),
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0:    "TASK_UNKNOWN",
		2100: "TASK_SCHEDULED",
		2110: "TASK_WAITING",
		2120: "TASK_PREPARING",
		2200: "TASK_RUNNING",
		2300: "TASK_COMPLETED",
		2310: "TASK_CANCELLED_TIMEOUT",
		2320: "TASK_CANCELLED_USER",
		2330: "TASK_CANCELLED_ADMIN",
	}
	TaskStatus_value = map[string]int32{
		"TASK_UNKNOWN":           0,
		"TASK_SCHEDULED":         2100,
		"TASK_WAITING":           2110,
		"TASK_PREPARING":         2120,
		"TASK_RUNNING":           2200,
		"TASK_COMPLETED":         2300,
		"TASK_CANCELLED_TIMEOUT": 2310,
		"TASK_CANCELLED_USER":    2320,
		"TASK_CANCELLED_ADMIN":   2330,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_task_proto_enumTypes[0].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_api_commons_task_proto_enumTypes[0]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_task_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_task_proto protoreflect.FileDescriptor

var file_api_commons_task_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0xd5, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0e, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53,
	0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0xb4, 0x10, 0x12, 0x11, 0x0a, 0x0c, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0xbe, 0x10, 0x12, 0x13,
	0x0a, 0x0e, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52, 0x49, 0x4e, 0x47,
	0x10, 0xc8, 0x10, 0x12, 0x11, 0x0a, 0x0c, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x55, 0x4e, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x98, 0x11, 0x12, 0x13, 0x0a, 0x0e, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0xfc, 0x11, 0x12, 0x1b, 0x0a, 0x16, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x86, 0x12, 0x12, 0x18, 0x0a, 0x13, 0x54, 0x41, 0x53, 0x4b,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10,
	0x90, 0x12, 0x12, 0x19, 0x0a, 0x14, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x9a, 0x12, 0x42, 0x91, 0x01,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x42, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_task_proto_rawDescOnce sync.Once
	file_api_commons_task_proto_rawDescData = file_api_commons_task_proto_rawDesc
)

func file_api_commons_task_proto_rawDescGZIP() []byte {
	file_api_commons_task_proto_rawDescOnce.Do(func() {
		file_api_commons_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_task_proto_rawDescData)
	})
	return file_api_commons_task_proto_rawDescData
}

var file_api_commons_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_task_proto_goTypes = []interface{}{
	(TaskStatus)(0), // 0: api.commons.TaskStatus
}
var file_api_commons_task_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_task_proto_init() }
func file_api_commons_task_proto_init() {
	if File_api_commons_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_task_proto_goTypes,
		DependencyIndexes: file_api_commons_task_proto_depIdxs,
		EnumInfos:         file_api_commons_task_proto_enumTypes,
	}.Build()
	File_api_commons_task_proto = out.File
	file_api_commons_task_proto_rawDesc = nil
	file_api_commons_task_proto_goTypes = nil
	file_api_commons_task_proto_depIdxs = nil
}
