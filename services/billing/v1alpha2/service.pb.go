// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/billing/v1alpha2/service.proto

package billingv1alpha2

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

var File_services_billing_v1alpha2_service_proto protoreflect.FileDescriptor

const file_services_billing_v1alpha2_service_proto_rawDesc = "" +
	"\n" +
	"'services/billing/v1alpha2/service.proto\x12\x19services.billing.v1alpha2\x1a\x17annotations/authz.proto\x1a\x1cgoogle/api/annotations.proto\x1a(services/billing/v1alpha2/invoices.proto\x1a%services/billing/v1alpha2/rates.proto2\xbb\x19\n" +
	"\x0eBillingService\x12\xec\x01\n" +
	"\x1bCreateDefaultRateDefinition\x12=.services.billing.v1alpha2.CreateDefaultRateDefinitionRequest\x1a>.services.billing.v1alpha2.CreateDefaultRateDefinitionResponse\"N\xba\xb8\x91\x02\b\n" +
	"\x06\b\xc8\x01\b\xf1\x01\x82\xd3\xe4\x93\x02;:\x01*\"6/services/billing/v1alpha2/createdefaultratedefinition\x12\xf0\x01\n" +
	"\x1cCreateDefaultRateDefinitions\x12>.services.billing.v1alpha2.CreateDefaultRateDefinitionsRequest\x1a?.services.billing.v1alpha2.CreateDefaultRateDefinitionsResponse\"O\xba\xb8\x91\x02\b\n" +
	"\x06\b\xc8\x01\b\xf1\x01\x82\xd3\xe4\x93\x02<:\x01*\"7/services/billing/v1alpha2/createdefaultratedefinitions\x12\xcd\x01\n" +
	"\x14CreateRateDefinition\x126.services.billing.v1alpha2.CreateRateDefinitionRequest\x1a7.services.billing.v1alpha2.CreateRateDefinitionResponse\"D\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xc8\x01\x82\xd3\xe4\x93\x024:\x01*\"//services/billing/v1alpha2/createratedefinition\x12\xd1\x01\n" +
	"\x15CreateRateDefinitions\x127.services.billing.v1alpha2.CreateRateDefinitionsRequest\x1a8.services.billing.v1alpha2.CreateRateDefinitionsResponse\"E\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xc8\x01\x82\xd3\xe4\x93\x025:\x01*\"0/services/billing/v1alpha2/createratedefinitions\x12\xec\x01\n" +
	"\x1bDeleteDefaultRateDefinition\x12=.services.billing.v1alpha2.DeleteDefaultRateDefinitionRequest\x1a>.services.billing.v1alpha2.DeleteDefaultRateDefinitionResponse\"N\xba\xb8\x91\x02\b\n" +
	"\x06\b\xc8\x01\b\xf1\x01\x82\xd3\xe4\x93\x02;:\x01*\"6/services/billing/v1alpha2/deletedefaultratedefinition\x12\xf0\x01\n" +
	"\x1cDeleteDefaultRateDefinitions\x12>.services.billing.v1alpha2.DeleteDefaultRateDefinitionsRequest\x1a?.services.billing.v1alpha2.DeleteDefaultRateDefinitionsResponse\"O\xba\xb8\x91\x02\b\n" +
	"\x06\b\xc8\x01\b\xf1\x01\x82\xd3\xe4\x93\x02<:\x01*\"7/services/billing/v1alpha2/deletedefaultratedefinitions\x12\xcd\x01\n" +
	"\x14DeleteRateDefinition\x126.services.billing.v1alpha2.DeleteRateDefinitionRequest\x1a7.services.billing.v1alpha2.DeleteRateDefinitionResponse\"D\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xc8\x01\x82\xd3\xe4\x93\x024:\x01*\"//services/billing/v1alpha2/deleteratedefinition\x12\xd1\x01\n" +
	"\x15DeleteRateDefinitions\x127.services.billing.v1alpha2.DeleteRateDefinitionsRequest\x1a8.services.billing.v1alpha2.DeleteRateDefinitionsResponse\"E\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xc8\x01\x82\xd3\xe4\x93\x025:\x01*\"0/services/billing/v1alpha2/deleteratedefinitions\x12\xb1\x01\n" +
	"\rExportInvoice\x12/.services.billing.v1alpha2.ExportInvoiceRequest\x1a0.services.billing.v1alpha2.ExportInvoiceResponse\"=\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf1\x01\x82\xd3\xe4\x93\x02-:\x01*\"(/services/billing/v1alpha2/exportinvoice\x12\xc1\x01\n" +
	"\x11GetRateDefinition\x123.services.billing.v1alpha2.GetRateDefinitionRequest\x1a4.services.billing.v1alpha2.GetRateDefinitionResponse\"A\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xc8\x01\x82\xd3\xe4\x93\x021:\x01*\",/services/billing/v1alpha2/getratedefinition\x12\xb5\x01\n" +
	"\x0eGetRateHistory\x120.services.billing.v1alpha2.GetRateHistoryRequest\x1a1.services.billing.v1alpha2.GetRateHistoryResponse\">\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xc8\x01\x82\xd3\xe4\x93\x02.:\x01*\")/services/billing/v1alpha2/getratehistory\x12\xe1\x01\n" +
	"\x19ListActiveRateDefinitions\x12;.services.billing.v1alpha2.ListActiveRateDefinitionsRequest\x1a<.services.billing.v1alpha2.ListActiveRateDefinitionsResponse\"I\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xc8\x01\x82\xd3\xe4\x93\x029:\x01*\"4/services/billing/v1alpha2/listactiveratedefinitions\x12\xc9\x01\n" +
	"\x13ListRateDefinitions\x125.services.billing.v1alpha2.ListRateDefinitionsRequest\x1a6.services.billing.v1alpha2.ListRateDefinitionsResponse\"C\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xc8\x01\x82\xd3\xe4\x93\x023:\x01*\"./services/billing/v1alpha2/listratedefinitions\x12\xec\x01\n" +
	"\x1bUpdateDefaultRateDefinition\x12=.services.billing.v1alpha2.UpdateDefaultRateDefinitionRequest\x1a>.services.billing.v1alpha2.UpdateDefaultRateDefinitionResponse\"N\xba\xb8\x91\x02\b\n" +
	"\x06\b\xc8\x01\b\xf1\x01\x82\xd3\xe4\x93\x02;:\x01*\"6/services/billing/v1alpha2/updatedefaultratedefinition\x12\xcd\x01\n" +
	"\x14UpdateRateDefinition\x126.services.billing.v1alpha2.UpdateRateDefinitionRequest\x1a7.services.billing.v1alpha2.UpdateRateDefinitionResponse\"D\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xc8\x01\x82\xd3\xe4\x93\x024:\x01*\"//services/billing/v1alpha2/updateratedefinitionB\xf9\x01\n" +
	"\x1dcom.services.billing.v1alpha2B\fServiceProtoP\x01ZDgithub.com/tcncloud/api-go/services/billing/v1alpha2;billingv1alpha2\xa2\x02\x03SBX\xaa\x02\x19Services.Billing.V1alpha2\xca\x02\x19Services\\Billing\\V1alpha2\xe2\x02%Services\\Billing\\V1alpha2\\GPBMetadata\xea\x02\x1bServices::Billing::V1alpha2b\x06proto3"

var file_services_billing_v1alpha2_service_proto_goTypes = []any{
	(*CreateDefaultRateDefinitionRequest)(nil),   // 0: services.billing.v1alpha2.CreateDefaultRateDefinitionRequest
	(*CreateDefaultRateDefinitionsRequest)(nil),  // 1: services.billing.v1alpha2.CreateDefaultRateDefinitionsRequest
	(*CreateRateDefinitionRequest)(nil),          // 2: services.billing.v1alpha2.CreateRateDefinitionRequest
	(*CreateRateDefinitionsRequest)(nil),         // 3: services.billing.v1alpha2.CreateRateDefinitionsRequest
	(*DeleteDefaultRateDefinitionRequest)(nil),   // 4: services.billing.v1alpha2.DeleteDefaultRateDefinitionRequest
	(*DeleteDefaultRateDefinitionsRequest)(nil),  // 5: services.billing.v1alpha2.DeleteDefaultRateDefinitionsRequest
	(*DeleteRateDefinitionRequest)(nil),          // 6: services.billing.v1alpha2.DeleteRateDefinitionRequest
	(*DeleteRateDefinitionsRequest)(nil),         // 7: services.billing.v1alpha2.DeleteRateDefinitionsRequest
	(*ExportInvoiceRequest)(nil),                 // 8: services.billing.v1alpha2.ExportInvoiceRequest
	(*GetRateDefinitionRequest)(nil),             // 9: services.billing.v1alpha2.GetRateDefinitionRequest
	(*GetRateHistoryRequest)(nil),                // 10: services.billing.v1alpha2.GetRateHistoryRequest
	(*ListActiveRateDefinitionsRequest)(nil),     // 11: services.billing.v1alpha2.ListActiveRateDefinitionsRequest
	(*ListRateDefinitionsRequest)(nil),           // 12: services.billing.v1alpha2.ListRateDefinitionsRequest
	(*UpdateDefaultRateDefinitionRequest)(nil),   // 13: services.billing.v1alpha2.UpdateDefaultRateDefinitionRequest
	(*UpdateRateDefinitionRequest)(nil),          // 14: services.billing.v1alpha2.UpdateRateDefinitionRequest
	(*CreateDefaultRateDefinitionResponse)(nil),  // 15: services.billing.v1alpha2.CreateDefaultRateDefinitionResponse
	(*CreateDefaultRateDefinitionsResponse)(nil), // 16: services.billing.v1alpha2.CreateDefaultRateDefinitionsResponse
	(*CreateRateDefinitionResponse)(nil),         // 17: services.billing.v1alpha2.CreateRateDefinitionResponse
	(*CreateRateDefinitionsResponse)(nil),        // 18: services.billing.v1alpha2.CreateRateDefinitionsResponse
	(*DeleteDefaultRateDefinitionResponse)(nil),  // 19: services.billing.v1alpha2.DeleteDefaultRateDefinitionResponse
	(*DeleteDefaultRateDefinitionsResponse)(nil), // 20: services.billing.v1alpha2.DeleteDefaultRateDefinitionsResponse
	(*DeleteRateDefinitionResponse)(nil),         // 21: services.billing.v1alpha2.DeleteRateDefinitionResponse
	(*DeleteRateDefinitionsResponse)(nil),        // 22: services.billing.v1alpha2.DeleteRateDefinitionsResponse
	(*ExportInvoiceResponse)(nil),                // 23: services.billing.v1alpha2.ExportInvoiceResponse
	(*GetRateDefinitionResponse)(nil),            // 24: services.billing.v1alpha2.GetRateDefinitionResponse
	(*GetRateHistoryResponse)(nil),               // 25: services.billing.v1alpha2.GetRateHistoryResponse
	(*ListActiveRateDefinitionsResponse)(nil),    // 26: services.billing.v1alpha2.ListActiveRateDefinitionsResponse
	(*ListRateDefinitionsResponse)(nil),          // 27: services.billing.v1alpha2.ListRateDefinitionsResponse
	(*UpdateDefaultRateDefinitionResponse)(nil),  // 28: services.billing.v1alpha2.UpdateDefaultRateDefinitionResponse
	(*UpdateRateDefinitionResponse)(nil),         // 29: services.billing.v1alpha2.UpdateRateDefinitionResponse
}
var file_services_billing_v1alpha2_service_proto_depIdxs = []int32{
	0,  // 0: services.billing.v1alpha2.BillingService.CreateDefaultRateDefinition:input_type -> services.billing.v1alpha2.CreateDefaultRateDefinitionRequest
	1,  // 1: services.billing.v1alpha2.BillingService.CreateDefaultRateDefinitions:input_type -> services.billing.v1alpha2.CreateDefaultRateDefinitionsRequest
	2,  // 2: services.billing.v1alpha2.BillingService.CreateRateDefinition:input_type -> services.billing.v1alpha2.CreateRateDefinitionRequest
	3,  // 3: services.billing.v1alpha2.BillingService.CreateRateDefinitions:input_type -> services.billing.v1alpha2.CreateRateDefinitionsRequest
	4,  // 4: services.billing.v1alpha2.BillingService.DeleteDefaultRateDefinition:input_type -> services.billing.v1alpha2.DeleteDefaultRateDefinitionRequest
	5,  // 5: services.billing.v1alpha2.BillingService.DeleteDefaultRateDefinitions:input_type -> services.billing.v1alpha2.DeleteDefaultRateDefinitionsRequest
	6,  // 6: services.billing.v1alpha2.BillingService.DeleteRateDefinition:input_type -> services.billing.v1alpha2.DeleteRateDefinitionRequest
	7,  // 7: services.billing.v1alpha2.BillingService.DeleteRateDefinitions:input_type -> services.billing.v1alpha2.DeleteRateDefinitionsRequest
	8,  // 8: services.billing.v1alpha2.BillingService.ExportInvoice:input_type -> services.billing.v1alpha2.ExportInvoiceRequest
	9,  // 9: services.billing.v1alpha2.BillingService.GetRateDefinition:input_type -> services.billing.v1alpha2.GetRateDefinitionRequest
	10, // 10: services.billing.v1alpha2.BillingService.GetRateHistory:input_type -> services.billing.v1alpha2.GetRateHistoryRequest
	11, // 11: services.billing.v1alpha2.BillingService.ListActiveRateDefinitions:input_type -> services.billing.v1alpha2.ListActiveRateDefinitionsRequest
	12, // 12: services.billing.v1alpha2.BillingService.ListRateDefinitions:input_type -> services.billing.v1alpha2.ListRateDefinitionsRequest
	13, // 13: services.billing.v1alpha2.BillingService.UpdateDefaultRateDefinition:input_type -> services.billing.v1alpha2.UpdateDefaultRateDefinitionRequest
	14, // 14: services.billing.v1alpha2.BillingService.UpdateRateDefinition:input_type -> services.billing.v1alpha2.UpdateRateDefinitionRequest
	15, // 15: services.billing.v1alpha2.BillingService.CreateDefaultRateDefinition:output_type -> services.billing.v1alpha2.CreateDefaultRateDefinitionResponse
	16, // 16: services.billing.v1alpha2.BillingService.CreateDefaultRateDefinitions:output_type -> services.billing.v1alpha2.CreateDefaultRateDefinitionsResponse
	17, // 17: services.billing.v1alpha2.BillingService.CreateRateDefinition:output_type -> services.billing.v1alpha2.CreateRateDefinitionResponse
	18, // 18: services.billing.v1alpha2.BillingService.CreateRateDefinitions:output_type -> services.billing.v1alpha2.CreateRateDefinitionsResponse
	19, // 19: services.billing.v1alpha2.BillingService.DeleteDefaultRateDefinition:output_type -> services.billing.v1alpha2.DeleteDefaultRateDefinitionResponse
	20, // 20: services.billing.v1alpha2.BillingService.DeleteDefaultRateDefinitions:output_type -> services.billing.v1alpha2.DeleteDefaultRateDefinitionsResponse
	21, // 21: services.billing.v1alpha2.BillingService.DeleteRateDefinition:output_type -> services.billing.v1alpha2.DeleteRateDefinitionResponse
	22, // 22: services.billing.v1alpha2.BillingService.DeleteRateDefinitions:output_type -> services.billing.v1alpha2.DeleteRateDefinitionsResponse
	23, // 23: services.billing.v1alpha2.BillingService.ExportInvoice:output_type -> services.billing.v1alpha2.ExportInvoiceResponse
	24, // 24: services.billing.v1alpha2.BillingService.GetRateDefinition:output_type -> services.billing.v1alpha2.GetRateDefinitionResponse
	25, // 25: services.billing.v1alpha2.BillingService.GetRateHistory:output_type -> services.billing.v1alpha2.GetRateHistoryResponse
	26, // 26: services.billing.v1alpha2.BillingService.ListActiveRateDefinitions:output_type -> services.billing.v1alpha2.ListActiveRateDefinitionsResponse
	27, // 27: services.billing.v1alpha2.BillingService.ListRateDefinitions:output_type -> services.billing.v1alpha2.ListRateDefinitionsResponse
	28, // 28: services.billing.v1alpha2.BillingService.UpdateDefaultRateDefinition:output_type -> services.billing.v1alpha2.UpdateDefaultRateDefinitionResponse
	29, // 29: services.billing.v1alpha2.BillingService.UpdateRateDefinition:output_type -> services.billing.v1alpha2.UpdateRateDefinitionResponse
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_billing_v1alpha2_service_proto_init() }
func file_services_billing_v1alpha2_service_proto_init() {
	if File_services_billing_v1alpha2_service_proto != nil {
		return
	}
	file_services_billing_v1alpha2_invoices_proto_init()
	file_services_billing_v1alpha2_rates_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_v1alpha2_service_proto_rawDesc), len(file_services_billing_v1alpha2_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_billing_v1alpha2_service_proto_goTypes,
		DependencyIndexes: file_services_billing_v1alpha2_service_proto_depIdxs,
	}.Build()
	File_services_billing_v1alpha2_service_proto = out.File
	file_services_billing_v1alpha2_service_proto_goTypes = nil
	file_services_billing_v1alpha2_service_proto_depIdxs = nil
}
