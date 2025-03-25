// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha3/omni.proto

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

type OmniSmsConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Prefixes      *CountryCodePrefix     `protobuf:"bytes,1,opt,name=prefixes,proto3" json:"prefixes,omitempty"`
	Config        *BasicConfig           `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OmniSmsConfig) Reset() {
	*x = OmniSmsConfig{}
	mi := &file_services_billing_entities_v1alpha3_omni_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OmniSmsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OmniSmsConfig) ProtoMessage() {}

func (x *OmniSmsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha3_omni_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OmniSmsConfig.ProtoReflect.Descriptor instead.
func (*OmniSmsConfig) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha3_omni_proto_rawDescGZIP(), []int{0}
}

func (x *OmniSmsConfig) GetPrefixes() *CountryCodePrefix {
	if x != nil {
		return x.Prefixes
	}
	return nil
}

func (x *OmniSmsConfig) GetConfig() *BasicConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type OmniSmsUnitConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Prefixes      *CountryCodePrefix     `protobuf:"bytes,1,opt,name=prefixes,proto3" json:"prefixes,omitempty"`
	Config        *BasicUnitConfig       `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OmniSmsUnitConfig) Reset() {
	*x = OmniSmsUnitConfig{}
	mi := &file_services_billing_entities_v1alpha3_omni_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OmniSmsUnitConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OmniSmsUnitConfig) ProtoMessage() {}

func (x *OmniSmsUnitConfig) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha3_omni_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OmniSmsUnitConfig.ProtoReflect.Descriptor instead.
func (*OmniSmsUnitConfig) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha3_omni_proto_rawDescGZIP(), []int{1}
}

func (x *OmniSmsUnitConfig) GetPrefixes() *CountryCodePrefix {
	if x != nil {
		return x.Prefixes
	}
	return nil
}

func (x *OmniSmsUnitConfig) GetConfig() *BasicUnitConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_services_billing_entities_v1alpha3_omni_proto protoreflect.FileDescriptor

const file_services_billing_entities_v1alpha3_omni_proto_rawDesc = "" +
	"\n" +
	"-services/billing/entities/v1alpha3/omni.proto\x12\"services.billing.entities.v1alpha3\x1a1services/billing/entities/v1alpha3/matching.proto\x1a0services/billing/entities/v1alpha3/modules.proto\"\xab\x01\n" +
	"\rOmniSmsConfig\x12Q\n" +
	"\bprefixes\x18\x01 \x01(\v25.services.billing.entities.v1alpha3.CountryCodePrefixR\bprefixes\x12G\n" +
	"\x06config\x18\x02 \x01(\v2/.services.billing.entities.v1alpha3.BasicConfigR\x06config\"\xb3\x01\n" +
	"\x11OmniSmsUnitConfig\x12Q\n" +
	"\bprefixes\x18\x01 \x01(\v25.services.billing.entities.v1alpha3.CountryCodePrefixR\bprefixes\x12K\n" +
	"\x06config\x18\x02 \x01(\v23.services.billing.entities.v1alpha3.BasicUnitConfigR\x06configB\xae\x02\n" +
	"&com.services.billing.entities.v1alpha3B\tOmniProtoP\x01ZNgithub.com/tcncloud/api-go/services/billing/entities/v1alpha3;entitiesv1alpha3\xa2\x02\x03SBE\xaa\x02\"Services.Billing.Entities.V1alpha3\xca\x02\"Services\\Billing\\Entities\\V1alpha3\xe2\x02.Services\\Billing\\Entities\\V1alpha3\\GPBMetadata\xea\x02%Services::Billing::Entities::V1alpha3b\x06proto3"

var (
	file_services_billing_entities_v1alpha3_omni_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha3_omni_proto_rawDescData []byte
)

func file_services_billing_entities_v1alpha3_omni_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha3_omni_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha3_omni_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha3_omni_proto_rawDesc), len(file_services_billing_entities_v1alpha3_omni_proto_rawDesc)))
	})
	return file_services_billing_entities_v1alpha3_omni_proto_rawDescData
}

var file_services_billing_entities_v1alpha3_omni_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_billing_entities_v1alpha3_omni_proto_goTypes = []any{
	(*OmniSmsConfig)(nil),     // 0: services.billing.entities.v1alpha3.OmniSmsConfig
	(*OmniSmsUnitConfig)(nil), // 1: services.billing.entities.v1alpha3.OmniSmsUnitConfig
	(*CountryCodePrefix)(nil), // 2: services.billing.entities.v1alpha3.CountryCodePrefix
	(*BasicConfig)(nil),       // 3: services.billing.entities.v1alpha3.BasicConfig
	(*BasicUnitConfig)(nil),   // 4: services.billing.entities.v1alpha3.BasicUnitConfig
}
var file_services_billing_entities_v1alpha3_omni_proto_depIdxs = []int32{
	2, // 0: services.billing.entities.v1alpha3.OmniSmsConfig.prefixes:type_name -> services.billing.entities.v1alpha3.CountryCodePrefix
	3, // 1: services.billing.entities.v1alpha3.OmniSmsConfig.config:type_name -> services.billing.entities.v1alpha3.BasicConfig
	2, // 2: services.billing.entities.v1alpha3.OmniSmsUnitConfig.prefixes:type_name -> services.billing.entities.v1alpha3.CountryCodePrefix
	4, // 3: services.billing.entities.v1alpha3.OmniSmsUnitConfig.config:type_name -> services.billing.entities.v1alpha3.BasicUnitConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha3_omni_proto_init() }
func file_services_billing_entities_v1alpha3_omni_proto_init() {
	if File_services_billing_entities_v1alpha3_omni_proto != nil {
		return
	}
	file_services_billing_entities_v1alpha3_matching_proto_init()
	file_services_billing_entities_v1alpha3_modules_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha3_omni_proto_rawDesc), len(file_services_billing_entities_v1alpha3_omni_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha3_omni_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha3_omni_proto_depIdxs,
		MessageInfos:      file_services_billing_entities_v1alpha3_omni_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha3_omni_proto = out.File
	file_services_billing_entities_v1alpha3_omni_proto_goTypes = nil
	file_services_billing_entities_v1alpha3_omni_proto_depIdxs = nil
}
