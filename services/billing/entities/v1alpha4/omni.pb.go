// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha4/omni.proto

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

type OmniSmsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/omni.proto.
	Name     string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Prefixes *CountryCodePrefix `protobuf:"bytes,2,opt,name=prefixes,proto3" json:"prefixes,omitempty"`
	Config   *BasicConfig       `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *OmniSmsConfig) Reset() {
	*x = OmniSmsConfig{}
	mi := &file_services_billing_entities_v1alpha4_omni_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OmniSmsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OmniSmsConfig) ProtoMessage() {}

func (x *OmniSmsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_omni_proto_msgTypes[0]
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
	return file_services_billing_entities_v1alpha4_omni_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/omni.proto.
func (x *OmniSmsConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/omni.proto.
	Name     string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Prefixes *CountryCodePrefix `protobuf:"bytes,2,opt,name=prefixes,proto3" json:"prefixes,omitempty"`
	Config   *BasicUnitConfig   `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *OmniSmsUnitConfig) Reset() {
	*x = OmniSmsUnitConfig{}
	mi := &file_services_billing_entities_v1alpha4_omni_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OmniSmsUnitConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OmniSmsUnitConfig) ProtoMessage() {}

func (x *OmniSmsUnitConfig) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_omni_proto_msgTypes[1]
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
	return file_services_billing_entities_v1alpha4_omni_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/omni.proto.
func (x *OmniSmsUnitConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
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

var File_services_billing_entities_v1alpha4_omni_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha4_omni_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x34, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x34, 0x1a, 0x31, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x0d, 0x4f, 0x6d, 0x6e,
	0x69, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x08, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xcb,
	0x01, 0x0a, 0x11, 0x4f, 0x6d, 0x6e, 0x69, 0x53, 0x6d, 0x73, 0x55, 0x6e, 0x69, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x08,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x34, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x08, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12,
	0x4b, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x34, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x55, 0x6e, 0x69, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xae, 0x02, 0x0a,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x42, 0x09, 0x4f, 0x6d, 0x6e, 0x69, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x34, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x34, 0xa2, 0x02, 0x03, 0x53, 0x42, 0x45, 0xaa, 0x02, 0x22, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0xca,
	0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x34, 0xe2, 0x02, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_billing_entities_v1alpha4_omni_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha4_omni_proto_rawDescData = file_services_billing_entities_v1alpha4_omni_proto_rawDesc
)

func file_services_billing_entities_v1alpha4_omni_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha4_omni_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha4_omni_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_billing_entities_v1alpha4_omni_proto_rawDescData)
	})
	return file_services_billing_entities_v1alpha4_omni_proto_rawDescData
}

var file_services_billing_entities_v1alpha4_omni_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_billing_entities_v1alpha4_omni_proto_goTypes = []any{
	(*OmniSmsConfig)(nil),     // 0: services.billing.entities.v1alpha4.OmniSmsConfig
	(*OmniSmsUnitConfig)(nil), // 1: services.billing.entities.v1alpha4.OmniSmsUnitConfig
	(*CountryCodePrefix)(nil), // 2: services.billing.entities.v1alpha4.CountryCodePrefix
	(*BasicConfig)(nil),       // 3: services.billing.entities.v1alpha4.BasicConfig
	(*BasicUnitConfig)(nil),   // 4: services.billing.entities.v1alpha4.BasicUnitConfig
}
var file_services_billing_entities_v1alpha4_omni_proto_depIdxs = []int32{
	2, // 0: services.billing.entities.v1alpha4.OmniSmsConfig.prefixes:type_name -> services.billing.entities.v1alpha4.CountryCodePrefix
	3, // 1: services.billing.entities.v1alpha4.OmniSmsConfig.config:type_name -> services.billing.entities.v1alpha4.BasicConfig
	2, // 2: services.billing.entities.v1alpha4.OmniSmsUnitConfig.prefixes:type_name -> services.billing.entities.v1alpha4.CountryCodePrefix
	4, // 3: services.billing.entities.v1alpha4.OmniSmsUnitConfig.config:type_name -> services.billing.entities.v1alpha4.BasicUnitConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha4_omni_proto_init() }
func file_services_billing_entities_v1alpha4_omni_proto_init() {
	if File_services_billing_entities_v1alpha4_omni_proto != nil {
		return
	}
	file_services_billing_entities_v1alpha4_matching_proto_init()
	file_services_billing_entities_v1alpha4_modules_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_billing_entities_v1alpha4_omni_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha4_omni_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha4_omni_proto_depIdxs,
		MessageInfos:      file_services_billing_entities_v1alpha4_omni_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha4_omni_proto = out.File
	file_services_billing_entities_v1alpha4_omni_proto_rawDesc = nil
	file_services_billing_entities_v1alpha4_omni_proto_goTypes = nil
	file_services_billing_entities_v1alpha4_omni_proto_depIdxs = nil
}
