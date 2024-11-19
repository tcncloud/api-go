// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha4/matching.proto

package entitiesv1alpha4

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

// MatchingConfig -
type MatchingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Config:
	//
	//	*MatchingConfig_CountryCodePrefix
	Config isMatchingConfig_Config `protobuf_oneof:"config"`
}

func (x *MatchingConfig) Reset() {
	*x = MatchingConfig{}
	mi := &file_services_billing_entities_v1alpha4_matching_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingConfig) ProtoMessage() {}

func (x *MatchingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_matching_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchingConfig.ProtoReflect.Descriptor instead.
func (*MatchingConfig) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha4_matching_proto_rawDescGZIP(), []int{0}
}

func (m *MatchingConfig) GetConfig() isMatchingConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *MatchingConfig) GetCountryCodePrefix() *CountryCodePrefix {
	if x, ok := x.GetConfig().(*MatchingConfig_CountryCodePrefix); ok {
		return x.CountryCodePrefix
	}
	return nil
}

type isMatchingConfig_Config interface {
	isMatchingConfig_Config()
}

type MatchingConfig_CountryCodePrefix struct {
	CountryCodePrefix *CountryCodePrefix `protobuf:"bytes,1,opt,name=country_code_prefix,json=countryCodePrefix,proto3,oneof"`
}

func (*MatchingConfig_CountryCodePrefix) isMatchingConfig_Config() {}

type CountryCodePrefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The country code to match.
	CountryCode int32 `protobuf:"varint,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// The list of prefixes (if any) to match after successfully
	// matching the country code.
	Prefixes []string `protobuf:"bytes,2,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	// The id of the matching rule.
	MatchingRuleId string `protobuf:"bytes,3,opt,name=matching_rule_id,json=matchingRuleId,proto3" json:"matching_rule_id,omitempty"`
	Name           string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CountryCodePrefix) Reset() {
	*x = CountryCodePrefix{}
	mi := &file_services_billing_entities_v1alpha4_matching_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CountryCodePrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountryCodePrefix) ProtoMessage() {}

func (x *CountryCodePrefix) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_matching_proto_msgTypes[1]
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
	return file_services_billing_entities_v1alpha4_matching_proto_rawDescGZIP(), []int{1}
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

func (x *CountryCodePrefix) GetMatchingRuleId() string {
	if x != nil {
		return x.MatchingRuleId
	}
	return ""
}

func (x *CountryCodePrefix) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_services_billing_entities_v1alpha4_matching_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha4_matching_proto_rawDesc = []byte{
	0x0a, 0x31, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x34, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x22, 0x83, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x67, 0x0a, 0x13, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x48, 0x00,
	0x52, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x90, 0x01,
	0x0a, 0x11, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0xb2, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x42, 0x0d, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x3b, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0xa2, 0x02, 0x03, 0x53,
	0x42, 0x45, 0xaa, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0xca, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0xe2, 0x02, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_billing_entities_v1alpha4_matching_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha4_matching_proto_rawDescData = file_services_billing_entities_v1alpha4_matching_proto_rawDesc
)

func file_services_billing_entities_v1alpha4_matching_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha4_matching_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha4_matching_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_billing_entities_v1alpha4_matching_proto_rawDescData)
	})
	return file_services_billing_entities_v1alpha4_matching_proto_rawDescData
}

var file_services_billing_entities_v1alpha4_matching_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_billing_entities_v1alpha4_matching_proto_goTypes = []any{
	(*MatchingConfig)(nil),    // 0: services.billing.entities.v1alpha4.MatchingConfig
	(*CountryCodePrefix)(nil), // 1: services.billing.entities.v1alpha4.CountryCodePrefix
}
var file_services_billing_entities_v1alpha4_matching_proto_depIdxs = []int32{
	1, // 0: services.billing.entities.v1alpha4.MatchingConfig.country_code_prefix:type_name -> services.billing.entities.v1alpha4.CountryCodePrefix
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha4_matching_proto_init() }
func file_services_billing_entities_v1alpha4_matching_proto_init() {
	if File_services_billing_entities_v1alpha4_matching_proto != nil {
		return
	}
	file_services_billing_entities_v1alpha4_matching_proto_msgTypes[0].OneofWrappers = []any{
		(*MatchingConfig_CountryCodePrefix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_billing_entities_v1alpha4_matching_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha4_matching_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha4_matching_proto_depIdxs,
		MessageInfos:      file_services_billing_entities_v1alpha4_matching_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha4_matching_proto = out.File
	file_services_billing_entities_v1alpha4_matching_proto_rawDesc = nil
	file_services_billing_entities_v1alpha4_matching_proto_goTypes = nil
	file_services_billing_entities_v1alpha4_matching_proto_depIdxs = nil
}
