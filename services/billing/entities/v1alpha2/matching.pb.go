// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha2/matching.proto

package entitiesv1alpha2

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

// MatchingRule represents a rule for matching an event to a
// rate definition.
type MatchingRule int32

const (
	MatchingRule_MATCHING_RULE_UNSPECIFIED    MatchingRule = 0
	MatchingRule_MATCHING_RULE_COUNTRY_PREFIX MatchingRule = 1
)

// Enum value maps for MatchingRule.
var (
	MatchingRule_name = map[int32]string{
		0: "MATCHING_RULE_UNSPECIFIED",
		1: "MATCHING_RULE_COUNTRY_PREFIX",
	}
	MatchingRule_value = map[string]int32{
		"MATCHING_RULE_UNSPECIFIED":    0,
		"MATCHING_RULE_COUNTRY_PREFIX": 1,
	}
)

func (x MatchingRule) Enum() *MatchingRule {
	p := new(MatchingRule)
	*p = x
	return p
}

func (x MatchingRule) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchingRule) Descriptor() protoreflect.EnumDescriptor {
	return file_services_billing_entities_v1alpha2_matching_proto_enumTypes[0].Descriptor()
}

func (MatchingRule) Type() protoreflect.EnumType {
	return &file_services_billing_entities_v1alpha2_matching_proto_enumTypes[0]
}

func (x MatchingRule) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchingRule.Descriptor instead.
func (MatchingRule) EnumDescriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha2_matching_proto_rawDescGZIP(), []int{0}
}

// MatchingConfig represents the configuration for matching
// an event to a rate definition.
type MatchingConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the name of the matching configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// the configuration data.
	//
	// Types that are valid to be assigned to Config:
	//
	//	*MatchingConfig_CountryPrefix
	Config        isMatchingConfig_Config `protobuf_oneof:"config"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchingConfig) Reset() {
	*x = MatchingConfig{}
	mi := &file_services_billing_entities_v1alpha2_matching_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingConfig) ProtoMessage() {}

func (x *MatchingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha2_matching_proto_msgTypes[0]
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
	return file_services_billing_entities_v1alpha2_matching_proto_rawDescGZIP(), []int{0}
}

func (x *MatchingConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MatchingConfig) GetConfig() isMatchingConfig_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *MatchingConfig) GetCountryPrefix() *MatchingConfigCountryPrefix {
	if x != nil {
		if x, ok := x.Config.(*MatchingConfig_CountryPrefix); ok {
			return x.CountryPrefix
		}
	}
	return nil
}

type isMatchingConfig_Config interface {
	isMatchingConfig_Config()
}

type MatchingConfig_CountryPrefix struct {
	CountryPrefix *MatchingConfigCountryPrefix `protobuf:"bytes,100,opt,name=country_prefix,json=countryPrefix,proto3,oneof"`
}

func (*MatchingConfig_CountryPrefix) isMatchingConfig_Config() {}

// MatchingConfigCountryPrefix represents the configuration for matching
// an event to a rate definition using the country code and a set of
// prefixes. The longest prefix match is used.
type MatchingConfigCountryPrefix struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The country code to match.
	CountryCode int32 `protobuf:"varint,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// The list of prefixes (if any) to match after successfully matching
	// the country code.
	Prefixes      []string `protobuf:"bytes,2,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchingConfigCountryPrefix) Reset() {
	*x = MatchingConfigCountryPrefix{}
	mi := &file_services_billing_entities_v1alpha2_matching_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchingConfigCountryPrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingConfigCountryPrefix) ProtoMessage() {}

func (x *MatchingConfigCountryPrefix) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha2_matching_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchingConfigCountryPrefix.ProtoReflect.Descriptor instead.
func (*MatchingConfigCountryPrefix) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha2_matching_proto_rawDescGZIP(), []int{1}
}

func (x *MatchingConfigCountryPrefix) GetCountryCode() int32 {
	if x != nil {
		return x.CountryCode
	}
	return 0
}

func (x *MatchingConfigCountryPrefix) GetPrefixes() []string {
	if x != nil {
		return x.Prefixes
	}
	return nil
}

var File_services_billing_entities_v1alpha2_matching_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha2_matching_proto_rawDesc = string([]byte{
	0x0a, 0x31, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x22, 0x98, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x68,
	0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x5c, 0x0a, 0x1b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73,
	0x2a, 0x4f, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65,
	0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x55, 0x4c,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x20, 0x0a, 0x1c, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x55, 0x4c, 0x45,
	0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x10,
	0x01, 0x42, 0xb2, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0d, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x3b, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2, 0x02, 0x03,
	0x53, 0x42, 0x45, 0xaa, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xe2, 0x02, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x25, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_services_billing_entities_v1alpha2_matching_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha2_matching_proto_rawDescData []byte
)

func file_services_billing_entities_v1alpha2_matching_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha2_matching_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha2_matching_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha2_matching_proto_rawDesc), len(file_services_billing_entities_v1alpha2_matching_proto_rawDesc)))
	})
	return file_services_billing_entities_v1alpha2_matching_proto_rawDescData
}

var file_services_billing_entities_v1alpha2_matching_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_billing_entities_v1alpha2_matching_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_billing_entities_v1alpha2_matching_proto_goTypes = []any{
	(MatchingRule)(0),                   // 0: services.billing.entities.v1alpha2.MatchingRule
	(*MatchingConfig)(nil),              // 1: services.billing.entities.v1alpha2.MatchingConfig
	(*MatchingConfigCountryPrefix)(nil), // 2: services.billing.entities.v1alpha2.MatchingConfigCountryPrefix
}
var file_services_billing_entities_v1alpha2_matching_proto_depIdxs = []int32{
	2, // 0: services.billing.entities.v1alpha2.MatchingConfig.country_prefix:type_name -> services.billing.entities.v1alpha2.MatchingConfigCountryPrefix
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha2_matching_proto_init() }
func file_services_billing_entities_v1alpha2_matching_proto_init() {
	if File_services_billing_entities_v1alpha2_matching_proto != nil {
		return
	}
	file_services_billing_entities_v1alpha2_matching_proto_msgTypes[0].OneofWrappers = []any{
		(*MatchingConfig_CountryPrefix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha2_matching_proto_rawDesc), len(file_services_billing_entities_v1alpha2_matching_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha2_matching_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha2_matching_proto_depIdxs,
		EnumInfos:         file_services_billing_entities_v1alpha2_matching_proto_enumTypes,
		MessageInfos:      file_services_billing_entities_v1alpha2_matching_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha2_matching_proto = out.File
	file_services_billing_entities_v1alpha2_matching_proto_goTypes = nil
	file_services_billing_entities_v1alpha2_matching_proto_depIdxs = nil
}
