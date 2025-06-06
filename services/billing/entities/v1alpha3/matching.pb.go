// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha3/matching.proto

package entitiesv1alpha3

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

type CountryCodePrefix struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The country code to match.
	CountryCode int32 `protobuf:"varint,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// The list of prefixes (if any) to match after successfully
	// matching the country code.
	Prefixes      []string `protobuf:"bytes,2,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CountryCodePrefix) Reset() {
	*x = CountryCodePrefix{}
	mi := &file_services_billing_entities_v1alpha3_matching_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CountryCodePrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountryCodePrefix) ProtoMessage() {}

func (x *CountryCodePrefix) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha3_matching_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountryCodePrefix.ProtoReflect.Descriptor instead.
func (*CountryCodePrefix) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha3_matching_proto_rawDescGZIP(), []int{0}
}

func (x *CountryCodePrefix) GetCountryCode() int32 {
	if x != nil {
		return x.CountryCode
	}
	return 0
}

func (x *CountryCodePrefix) GetPrefixes() []string {
	if x != nil {
		return x.Prefixes
	}
	return nil
}

var File_services_billing_entities_v1alpha3_matching_proto protoreflect.FileDescriptor

const file_services_billing_entities_v1alpha3_matching_proto_rawDesc = "" +
	"\n" +
	"1services/billing/entities/v1alpha3/matching.proto\x12\"services.billing.entities.v1alpha3\"R\n" +
	"\x11CountryCodePrefix\x12!\n" +
	"\fcountry_code\x18\x01 \x01(\x05R\vcountryCode\x12\x1a\n" +
	"\bprefixes\x18\x02 \x03(\tR\bprefixesB\xb2\x02\n" +
	"&com.services.billing.entities.v1alpha3B\rMatchingProtoP\x01ZNgithub.com/tcncloud/api-go/services/billing/entities/v1alpha3;entitiesv1alpha3\xa2\x02\x03SBE\xaa\x02\"Services.Billing.Entities.V1alpha3\xca\x02\"Services\\Billing\\Entities\\V1alpha3\xe2\x02.Services\\Billing\\Entities\\V1alpha3\\GPBMetadata\xea\x02%Services::Billing::Entities::V1alpha3b\x06proto3"

var (
	file_services_billing_entities_v1alpha3_matching_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha3_matching_proto_rawDescData []byte
)

func file_services_billing_entities_v1alpha3_matching_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha3_matching_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha3_matching_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha3_matching_proto_rawDesc), len(file_services_billing_entities_v1alpha3_matching_proto_rawDesc)))
	})
	return file_services_billing_entities_v1alpha3_matching_proto_rawDescData
}

var file_services_billing_entities_v1alpha3_matching_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_services_billing_entities_v1alpha3_matching_proto_goTypes = []any{
	(*CountryCodePrefix)(nil), // 0: services.billing.entities.v1alpha3.CountryCodePrefix
}
var file_services_billing_entities_v1alpha3_matching_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha3_matching_proto_init() }
func file_services_billing_entities_v1alpha3_matching_proto_init() {
	if File_services_billing_entities_v1alpha3_matching_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha3_matching_proto_rawDesc), len(file_services_billing_entities_v1alpha3_matching_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha3_matching_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha3_matching_proto_depIdxs,
		MessageInfos:      file_services_billing_entities_v1alpha3_matching_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha3_matching_proto = out.File
	file_services_billing_entities_v1alpha3_matching_proto_goTypes = nil
	file_services_billing_entities_v1alpha3_matching_proto_depIdxs = nil
}
