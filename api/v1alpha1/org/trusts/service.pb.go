// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/org/trusts/service.proto

package trusts

import (
	_ "github.com/tcncloud/api-go/annotations"
	org "github.com/tcncloud/api-go/api/v1alpha1/org"
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

var File_api_v1alpha1_org_trusts_service_proto protoreflect.FileDescriptor

const file_api_v1alpha1_org_trusts_service_proto_rawDesc = "" +
	"\n" +
	"%api/v1alpha1/org/trusts/service.proto\x12\x17api.v1alpha1.org.trusts\x1a\x17annotations/authz.proto\x1a*api/v1alpha1/org/agent_profile_group.proto\x1a!api/v1alpha1/org/auth_token.proto\x1a api/v1alpha1/org/huntgroup.proto\x1a\x1dapi/v1alpha1/org/labels.proto\x1a$api/v1alpha1/org/notifications.proto\x1a#api/v1alpha1/org/organization.proto\x1a%api/v1alpha1/org/p3_permissions.proto\x1a\"api/v1alpha1/org/permissions.proto\x1a\"api/v1alpha1/org/preferences.proto\x1a\x1dapi/v1alpha1/org/trusts.proto\x1a\x1bapi/v1alpha1/org/user.proto\x1a\x1cgoogle/api/annotations.proto2\xda\f\n" +
	"\rTrustsService\x12\x95\x01\n" +
	"\vCreateTrust\x12$.api.v1alpha1.org.CreateTrustRequest\x1a%.api.v1alpha1.org.CreateTrustResponse\"9\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xa0\x01\x82\xd3\xe4\x93\x02):\x01*\"$/api/v1alpha1/org/trusts/createtrust\x12\x95\x01\n" +
	"\vAcceptTrust\x12$.api.v1alpha1.org.AcceptTrustRequest\x1a%.api.v1alpha1.org.AcceptTrustResponse\"9\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xa0\x01\x82\xd3\xe4\x93\x02):\x01*\"$/api/v1alpha1/org/trusts/accepttrust\x12\x95\x01\n" +
	"\vRejectTrust\x12$.api.v1alpha1.org.RejectTrustRequest\x1a%.api.v1alpha1.org.RejectTrustResponse\"9\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xa0\x01\x82\xd3\xe4\x93\x02):\x01*\"$/api/v1alpha1/org/trusts/rejecttrust\x12\x89\x01\n" +
	"\bGetTrust\x12!.api.v1alpha1.org.GetTrustRequest\x1a\".api.v1alpha1.org.GetTrustResponse\"6\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xa0\x01\x82\xd3\xe4\x93\x02&:\x01*\"!/api/v1alpha1/org/trusts/gettrust\x12\xb6\x01\n" +
	"\x12ListIncomingTrusts\x12+.api.v1alpha1.org.ListIncomingTrustsRequest\x1a,.api.v1alpha1.org.ListIncomingTrustsResponse\"E\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xa0\x01\x82\xd3\xe4\x93\x025:\x01*\"0/api/v1alpha1/org/trusts/list/listincomingtrusts\x12\xaa\x01\n" +
	"\x0fListGivenTrusts\x12(.api.v1alpha1.org.ListGivenTrustsRequest\x1a).api.v1alpha1.org.ListGivenTrustsResponse\"B\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xa0\x01\x82\xd3\xe4\x93\x022:\x01*\"-/api/v1alpha1/org/trusts/list/listgiventrusts\x12\xbe\x01\n" +
	"\x14ListAssignableTrusts\x12-.api.v1alpha1.org.ListAssignableTrustsRequest\x1a..api.v1alpha1.org.ListAssignableTrustsResponse\"G\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xa0\x01\x82\xd3\xe4\x93\x027:\x01*\"2/api/v1alpha1/org/trusts/list/listassignabletrusts\x12\x95\x01\n" +
	"\vDeleteTrust\x12$.api.v1alpha1.org.DeleteTrustRequest\x1a%.api.v1alpha1.org.DeleteTrustResponse\"9\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xa0\x01\x82\xd3\xe4\x93\x02):\x01*\"$/api/v1alpha1/org/trusts/deletetrust\x12\x95\x01\n" +
	"\vAssignTrust\x12$.api.v1alpha1.org.AssignTrustRequest\x1a%.api.v1alpha1.org.AssignTrustResponse\"9\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xa0\x01\x82\xd3\xe4\x93\x02):\x01*\"$/api/v1alpha1/org/trusts/assigntrust\x12\x9d\x01\n" +
	"\rUnassignTrust\x12&.api.v1alpha1.org.UnassignTrustRequest\x1a'.api.v1alpha1.org.UnassignTrustResponse\";\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xa0\x01\x82\xd3\xe4\x93\x02+:\x01*\"&/api/v1alpha1/org/trusts/unassigntrustB\xdf\x01\n" +
	"\x1bcom.api.v1alpha1.org.trustsB\fServiceProtoP\x01Z2github.com/tcncloud/api-go/api/v1alpha1/org/trusts\xa2\x02\x04AVOT\xaa\x02\x17Api.V1alpha1.Org.Trusts\xca\x02\x17Api\\V1alpha1\\Org\\Trusts\xe2\x02#Api\\V1alpha1\\Org\\Trusts\\GPBMetadata\xea\x02\x1aApi::V1alpha1::Org::Trustsb\x06proto3"

var file_api_v1alpha1_org_trusts_service_proto_goTypes = []any{
	(*org.CreateTrustRequest)(nil),           // 0: api.v1alpha1.org.CreateTrustRequest
	(*org.AcceptTrustRequest)(nil),           // 1: api.v1alpha1.org.AcceptTrustRequest
	(*org.RejectTrustRequest)(nil),           // 2: api.v1alpha1.org.RejectTrustRequest
	(*org.GetTrustRequest)(nil),              // 3: api.v1alpha1.org.GetTrustRequest
	(*org.ListIncomingTrustsRequest)(nil),    // 4: api.v1alpha1.org.ListIncomingTrustsRequest
	(*org.ListGivenTrustsRequest)(nil),       // 5: api.v1alpha1.org.ListGivenTrustsRequest
	(*org.ListAssignableTrustsRequest)(nil),  // 6: api.v1alpha1.org.ListAssignableTrustsRequest
	(*org.DeleteTrustRequest)(nil),           // 7: api.v1alpha1.org.DeleteTrustRequest
	(*org.AssignTrustRequest)(nil),           // 8: api.v1alpha1.org.AssignTrustRequest
	(*org.UnassignTrustRequest)(nil),         // 9: api.v1alpha1.org.UnassignTrustRequest
	(*org.CreateTrustResponse)(nil),          // 10: api.v1alpha1.org.CreateTrustResponse
	(*org.AcceptTrustResponse)(nil),          // 11: api.v1alpha1.org.AcceptTrustResponse
	(*org.RejectTrustResponse)(nil),          // 12: api.v1alpha1.org.RejectTrustResponse
	(*org.GetTrustResponse)(nil),             // 13: api.v1alpha1.org.GetTrustResponse
	(*org.ListIncomingTrustsResponse)(nil),   // 14: api.v1alpha1.org.ListIncomingTrustsResponse
	(*org.ListGivenTrustsResponse)(nil),      // 15: api.v1alpha1.org.ListGivenTrustsResponse
	(*org.ListAssignableTrustsResponse)(nil), // 16: api.v1alpha1.org.ListAssignableTrustsResponse
	(*org.DeleteTrustResponse)(nil),          // 17: api.v1alpha1.org.DeleteTrustResponse
	(*org.AssignTrustResponse)(nil),          // 18: api.v1alpha1.org.AssignTrustResponse
	(*org.UnassignTrustResponse)(nil),        // 19: api.v1alpha1.org.UnassignTrustResponse
}
var file_api_v1alpha1_org_trusts_service_proto_depIdxs = []int32{
	0,  // 0: api.v1alpha1.org.trusts.TrustsService.CreateTrust:input_type -> api.v1alpha1.org.CreateTrustRequest
	1,  // 1: api.v1alpha1.org.trusts.TrustsService.AcceptTrust:input_type -> api.v1alpha1.org.AcceptTrustRequest
	2,  // 2: api.v1alpha1.org.trusts.TrustsService.RejectTrust:input_type -> api.v1alpha1.org.RejectTrustRequest
	3,  // 3: api.v1alpha1.org.trusts.TrustsService.GetTrust:input_type -> api.v1alpha1.org.GetTrustRequest
	4,  // 4: api.v1alpha1.org.trusts.TrustsService.ListIncomingTrusts:input_type -> api.v1alpha1.org.ListIncomingTrustsRequest
	5,  // 5: api.v1alpha1.org.trusts.TrustsService.ListGivenTrusts:input_type -> api.v1alpha1.org.ListGivenTrustsRequest
	6,  // 6: api.v1alpha1.org.trusts.TrustsService.ListAssignableTrusts:input_type -> api.v1alpha1.org.ListAssignableTrustsRequest
	7,  // 7: api.v1alpha1.org.trusts.TrustsService.DeleteTrust:input_type -> api.v1alpha1.org.DeleteTrustRequest
	8,  // 8: api.v1alpha1.org.trusts.TrustsService.AssignTrust:input_type -> api.v1alpha1.org.AssignTrustRequest
	9,  // 9: api.v1alpha1.org.trusts.TrustsService.UnassignTrust:input_type -> api.v1alpha1.org.UnassignTrustRequest
	10, // 10: api.v1alpha1.org.trusts.TrustsService.CreateTrust:output_type -> api.v1alpha1.org.CreateTrustResponse
	11, // 11: api.v1alpha1.org.trusts.TrustsService.AcceptTrust:output_type -> api.v1alpha1.org.AcceptTrustResponse
	12, // 12: api.v1alpha1.org.trusts.TrustsService.RejectTrust:output_type -> api.v1alpha1.org.RejectTrustResponse
	13, // 13: api.v1alpha1.org.trusts.TrustsService.GetTrust:output_type -> api.v1alpha1.org.GetTrustResponse
	14, // 14: api.v1alpha1.org.trusts.TrustsService.ListIncomingTrusts:output_type -> api.v1alpha1.org.ListIncomingTrustsResponse
	15, // 15: api.v1alpha1.org.trusts.TrustsService.ListGivenTrusts:output_type -> api.v1alpha1.org.ListGivenTrustsResponse
	16, // 16: api.v1alpha1.org.trusts.TrustsService.ListAssignableTrusts:output_type -> api.v1alpha1.org.ListAssignableTrustsResponse
	17, // 17: api.v1alpha1.org.trusts.TrustsService.DeleteTrust:output_type -> api.v1alpha1.org.DeleteTrustResponse
	18, // 18: api.v1alpha1.org.trusts.TrustsService.AssignTrust:output_type -> api.v1alpha1.org.AssignTrustResponse
	19, // 19: api.v1alpha1.org.trusts.TrustsService.UnassignTrust:output_type -> api.v1alpha1.org.UnassignTrustResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_org_trusts_service_proto_init() }
func file_api_v1alpha1_org_trusts_service_proto_init() {
	if File_api_v1alpha1_org_trusts_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_org_trusts_service_proto_rawDesc), len(file_api_v1alpha1_org_trusts_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_org_trusts_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_org_trusts_service_proto_depIdxs,
	}.Build()
	File_api_v1alpha1_org_trusts_service_proto = out.File
	file_api_v1alpha1_org_trusts_service_proto_goTypes = nil
	file_api_v1alpha1_org_trusts_service_proto_depIdxs = nil
}
