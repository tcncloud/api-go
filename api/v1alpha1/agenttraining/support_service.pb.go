// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: api/v1alpha1/agenttraining/support_service.proto

package agenttraining

import (
	_ "github.com/tcncloud/api-go/annotations"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_v1alpha1_agenttraining_support_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_agenttraining_support_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x17,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x2f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcd, 0x04, 0x0a,
	0x1b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x95, 0x02, 0x0a,
	0x20, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x49,
	0x64, 0x12, 0x43, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xc8,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x5d, 0x3a, 0x01, 0x2a, 0x22, 0x58, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x79, 0x6f,
	0x72, 0x67, 0x69, 0x64, 0x12, 0x95, 0x02, 0x0a, 0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x43, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x42, 0x79, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0xba,
	0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xc8, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x5d, 0x3a,
	0x01, 0x2a, 0x22, 0x58, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x62, 0x79, 0x6f, 0x72, 0x67, 0x69, 0x64, 0x42, 0xf6, 0x01, 0x0a,
	0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x42,
	0x13, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0xa2, 0x02, 0x03,
	0x41, 0x56, 0x41, 0xaa, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0xca, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0xe2, 0x02, 0x26,
	0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_v1alpha1_agenttraining_support_service_proto_goTypes = []interface{}{
	(*ListLearningOpportunitiesByOrgIdRequest)(nil), // 0: api.v1alpha1.agenttraining.ListLearningOpportunitiesByOrgIdRequest
	(*DeleteLearningOpportunityByOrgIdRequest)(nil), // 1: api.v1alpha1.agenttraining.DeleteLearningOpportunityByOrgIdRequest
	(*ListLearningOpportunitiesResponse)(nil),       // 2: api.v1alpha1.agenttraining.ListLearningOpportunitiesResponse
	(*DeleteLearningOpportunityResponse)(nil),       // 3: api.v1alpha1.agenttraining.DeleteLearningOpportunityResponse
}
var file_api_v1alpha1_agenttraining_support_service_proto_depIdxs = []int32{
	0, // 0: api.v1alpha1.agenttraining.AgentTrainingSupportService.ListLearningOpportunitiesByOrgId:input_type -> api.v1alpha1.agenttraining.ListLearningOpportunitiesByOrgIdRequest
	1, // 1: api.v1alpha1.agenttraining.AgentTrainingSupportService.DeleteLearningOpportunityByOrgId:input_type -> api.v1alpha1.agenttraining.DeleteLearningOpportunityByOrgIdRequest
	2, // 2: api.v1alpha1.agenttraining.AgentTrainingSupportService.ListLearningOpportunitiesByOrgId:output_type -> api.v1alpha1.agenttraining.ListLearningOpportunitiesResponse
	3, // 3: api.v1alpha1.agenttraining.AgentTrainingSupportService.DeleteLearningOpportunityByOrgId:output_type -> api.v1alpha1.agenttraining.DeleteLearningOpportunityResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_agenttraining_support_service_proto_init() }
func file_api_v1alpha1_agenttraining_support_service_proto_init() {
	if File_api_v1alpha1_agenttraining_support_service_proto != nil {
		return
	}
	file_api_v1alpha1_agenttraining_learning_opportunity_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_agenttraining_support_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_agenttraining_support_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_agenttraining_support_service_proto_depIdxs,
	}.Build()
	File_api_v1alpha1_agenttraining_support_service_proto = out.File
	file_api_v1alpha1_agenttraining_support_service_proto_rawDesc = nil
	file_api_v1alpha1_agenttraining_support_service_proto_goTypes = nil
	file_api_v1alpha1_agenttraining_support_service_proto_depIdxs = nil
}
