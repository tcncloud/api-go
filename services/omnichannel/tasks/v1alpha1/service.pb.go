// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/omnichannel/tasks/v1alpha1/service.proto

package tasksv1alpha1

import (
	_ "github.com/tcncloud/api-go/annotations"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_omnichannel_tasks_v1alpha1_service_proto protoreflect.FileDescriptor

const file_services_omnichannel_tasks_v1alpha1_service_proto_rawDesc = "" +
	"\n" +
	"1services/omnichannel/tasks/v1alpha1/service.proto\x12#services.omnichannel.tasks.v1alpha1\x1a\x17annotations/authz.proto\x1a\x1cgoogle/api/annotations.proto\x1a2services/omnichannel/tasks/v1alpha1/entities.proto2\xb2\x03\n" +
	"\fTasksService\x12\xc7\x01\n" +
	"\vCancelTasks\x127.services.omnichannel.tasks.v1alpha1.CancelTasksRequest\x1a8.services.omnichannel.tasks.v1alpha1.CancelTasksResponse\"E\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xb0\t\x82\xd3\xe4\x93\x025:\x01*\"0/services/omnichannel/tasks/v1alpha1/canceltasks\x12\xd7\x01\n" +
	"\x0fBulkCancelTasks\x12;.services.omnichannel.tasks.v1alpha1.BulkCancelTasksRequest\x1a<.services.omnichannel.tasks.v1alpha1.BulkCancelTasksResponse\"I\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xb0\t\x82\xd3\xe4\x93\x029:\x01*\"4/services/omnichannel/tasks/v1alpha1/bulkcanceltasksB\xb4\x02\n" +
	"'com.services.omnichannel.tasks.v1alpha1B\fServiceProtoP\x01ZLgithub.com/tcncloud/api-go/services/omnichannel/tasks/v1alpha1;tasksv1alpha1\xa2\x02\x03SOT\xaa\x02#Services.Omnichannel.Tasks.V1alpha1\xca\x02#Services\\Omnichannel\\Tasks\\V1alpha1\xe2\x02/Services\\Omnichannel\\Tasks\\V1alpha1\\GPBMetadata\xea\x02&Services::Omnichannel::Tasks::V1alpha1b\x06proto3"

var file_services_omnichannel_tasks_v1alpha1_service_proto_goTypes = []any{
	(*CancelTasksRequest)(nil),      // 0: services.omnichannel.tasks.v1alpha1.CancelTasksRequest
	(*BulkCancelTasksRequest)(nil),  // 1: services.omnichannel.tasks.v1alpha1.BulkCancelTasksRequest
	(*CancelTasksResponse)(nil),     // 2: services.omnichannel.tasks.v1alpha1.CancelTasksResponse
	(*BulkCancelTasksResponse)(nil), // 3: services.omnichannel.tasks.v1alpha1.BulkCancelTasksResponse
}
var file_services_omnichannel_tasks_v1alpha1_service_proto_depIdxs = []int32{
	0, // 0: services.omnichannel.tasks.v1alpha1.TasksService.CancelTasks:input_type -> services.omnichannel.tasks.v1alpha1.CancelTasksRequest
	1, // 1: services.omnichannel.tasks.v1alpha1.TasksService.BulkCancelTasks:input_type -> services.omnichannel.tasks.v1alpha1.BulkCancelTasksRequest
	2, // 2: services.omnichannel.tasks.v1alpha1.TasksService.CancelTasks:output_type -> services.omnichannel.tasks.v1alpha1.CancelTasksResponse
	3, // 3: services.omnichannel.tasks.v1alpha1.TasksService.BulkCancelTasks:output_type -> services.omnichannel.tasks.v1alpha1.BulkCancelTasksResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_omnichannel_tasks_v1alpha1_service_proto_init() }
func file_services_omnichannel_tasks_v1alpha1_service_proto_init() {
	if File_services_omnichannel_tasks_v1alpha1_service_proto != nil {
		return
	}
	file_services_omnichannel_tasks_v1alpha1_entities_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_omnichannel_tasks_v1alpha1_service_proto_rawDesc), len(file_services_omnichannel_tasks_v1alpha1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_omnichannel_tasks_v1alpha1_service_proto_goTypes,
		DependencyIndexes: file_services_omnichannel_tasks_v1alpha1_service_proto_depIdxs,
	}.Build()
	File_services_omnichannel_tasks_v1alpha1_service_proto = out.File
	file_services_omnichannel_tasks_v1alpha1_service_proto_goTypes = nil
	file_services_omnichannel_tasks_v1alpha1_service_proto_depIdxs = nil
}
