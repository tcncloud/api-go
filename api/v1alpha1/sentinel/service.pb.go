// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/sentinel/service.proto

package sentinel

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

var File_api_v1alpha1_sentinel_service_proto protoreflect.FileDescriptor

const file_api_v1alpha1_sentinel_service_proto_rawDesc = "" +
	"\n" +
	"#api/v1alpha1/sentinel/service.proto\x12\x15api.v1alpha1.sentinel\x1a\x17annotations/authz.proto\x1a$api/v1alpha1/sentinel/entities.proto\x1a\x1cgoogle/api/annotations.proto2\x9a\x01\n" +
	"\bSentinel\x12\x8d\x01\n" +
	"\n" +
	"SendEvents\x12$.api.v1alpha1.sentinel.SendEventsReq\x1a$.api.v1alpha1.sentinel.SendEventsRes\"3\xba\xb8\x91\x02\x02\x18\x01\x82\xd3\xe4\x93\x02&:\x01*\"!/api/v1alpha1/sentinel/sendeventsB\xd1\x01\n" +
	"\x19com.api.v1alpha1.sentinelB\fServiceProtoP\x01Z0github.com/tcncloud/api-go/api/v1alpha1/sentinel\xa2\x02\x03AVS\xaa\x02\x15Api.V1alpha1.Sentinel\xca\x02\x15Api\\V1alpha1\\Sentinel\xe2\x02!Api\\V1alpha1\\Sentinel\\GPBMetadata\xea\x02\x17Api::V1alpha1::Sentinelb\x06proto3"

var file_api_v1alpha1_sentinel_service_proto_goTypes = []any{
	(*SendEventsReq)(nil), // 0: api.v1alpha1.sentinel.SendEventsReq
	(*SendEventsRes)(nil), // 1: api.v1alpha1.sentinel.SendEventsRes
}
var file_api_v1alpha1_sentinel_service_proto_depIdxs = []int32{
	0, // 0: api.v1alpha1.sentinel.Sentinel.SendEvents:input_type -> api.v1alpha1.sentinel.SendEventsReq
	1, // 1: api.v1alpha1.sentinel.Sentinel.SendEvents:output_type -> api.v1alpha1.sentinel.SendEventsRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_sentinel_service_proto_init() }
func file_api_v1alpha1_sentinel_service_proto_init() {
	if File_api_v1alpha1_sentinel_service_proto != nil {
		return
	}
	file_api_v1alpha1_sentinel_entities_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_sentinel_service_proto_rawDesc), len(file_api_v1alpha1_sentinel_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_sentinel_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_sentinel_service_proto_depIdxs,
	}.Build()
	File_api_v1alpha1_sentinel_service_proto = out.File
	file_api_v1alpha1_sentinel_service_proto_goTypes = nil
	file_api_v1alpha1_sentinel_service_proto_depIdxs = nil
}
