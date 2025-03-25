// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha1/modules.proto

package entitiesv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// BasicConfig represents the basic config for a rating module when
// there is no other data needed
type BasicConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. the amount to rate a single event
	Rate          float64 `protobuf:"fixed64,1,opt,name=rate,proto3" json:"rate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BasicConfig) Reset() {
	*x = BasicConfig{}
	mi := &file_services_billing_entities_v1alpha1_modules_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BasicConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicConfig) ProtoMessage() {}

func (x *BasicConfig) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha1_modules_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicConfig.ProtoReflect.Descriptor instead.
func (*BasicConfig) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha1_modules_proto_rawDescGZIP(), []int{0}
}

func (x *BasicConfig) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

// BasicUnitConfig represents basic config for a rating module that
// rates based on event units
type BasicUnitConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. the size of an individual unit. For example, a unit
	// size of 300 (with bytes as a unit of measurement) and an event
	// of size 400 bytes will get billed as if it were 2 units.
	UnitSize int64 `protobuf:"varint,1,opt,name=unit_size,json=unitSize,proto3" json:"unit_size,omitempty"`
	// Required. the amount to rate each unit
	Rate float64 `protobuf:"fixed64,2,opt,name=rate,proto3" json:"rate,omitempty"`
	// Optional. the minimum number of units to rate; for example, a
	// unit size of 5 seconds and a min_units of 2 would mean that
	// any event less than 10 seconds would be billed as if it were
	// two units (10 seconds).
	MinUnits *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=min_units,json=minUnits,proto3" json:"min_units,omitempty"`
	// Optional. the maximum number of units to rate; for example, a
	// unit size of 5 seconds and a max_units of 2 would mean that
	// any event more than 10 seconds would be billed as if it were
	// two units (10 seconds).
	MaxUnits      *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=max_units,json=maxUnits,proto3" json:"max_units,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BasicUnitConfig) Reset() {
	*x = BasicUnitConfig{}
	mi := &file_services_billing_entities_v1alpha1_modules_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BasicUnitConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicUnitConfig) ProtoMessage() {}

func (x *BasicUnitConfig) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha1_modules_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicUnitConfig.ProtoReflect.Descriptor instead.
func (*BasicUnitConfig) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha1_modules_proto_rawDescGZIP(), []int{1}
}

func (x *BasicUnitConfig) GetUnitSize() int64 {
	if x != nil {
		return x.UnitSize
	}
	return 0
}

func (x *BasicUnitConfig) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *BasicUnitConfig) GetMinUnits() *wrapperspb.Int64Value {
	if x != nil {
		return x.MinUnits
	}
	return nil
}

func (x *BasicUnitConfig) GetMaxUnits() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxUnits
	}
	return nil
}

var File_services_billing_entities_v1alpha1_modules_proto protoreflect.FileDescriptor

const file_services_billing_entities_v1alpha1_modules_proto_rawDesc = "" +
	"\n" +
	"0services/billing/entities/v1alpha1/modules.proto\x12\"services.billing.entities.v1alpha1\x1a\x1egoogle/protobuf/wrappers.proto\"!\n" +
	"\vBasicConfig\x12\x12\n" +
	"\x04rate\x18\x01 \x01(\x01R\x04rate\"\xb6\x01\n" +
	"\x0fBasicUnitConfig\x12\x1b\n" +
	"\tunit_size\x18\x01 \x01(\x03R\bunitSize\x12\x12\n" +
	"\x04rate\x18\x02 \x01(\x01R\x04rate\x128\n" +
	"\tmin_units\x18\x03 \x01(\v2\x1b.google.protobuf.Int64ValueR\bminUnits\x128\n" +
	"\tmax_units\x18\x04 \x01(\v2\x1b.google.protobuf.Int64ValueR\bmaxUnitsB\xb1\x02\n" +
	"&com.services.billing.entities.v1alpha1B\fModulesProtoP\x01ZNgithub.com/tcncloud/api-go/services/billing/entities/v1alpha1;entitiesv1alpha1\xa2\x02\x03SBE\xaa\x02\"Services.Billing.Entities.V1alpha1\xca\x02\"Services\\Billing\\Entities\\V1alpha1\xe2\x02.Services\\Billing\\Entities\\V1alpha1\\GPBMetadata\xea\x02%Services::Billing::Entities::V1alpha1b\x06proto3"

var (
	file_services_billing_entities_v1alpha1_modules_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha1_modules_proto_rawDescData []byte
)

func file_services_billing_entities_v1alpha1_modules_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha1_modules_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha1_modules_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha1_modules_proto_rawDesc), len(file_services_billing_entities_v1alpha1_modules_proto_rawDesc)))
	})
	return file_services_billing_entities_v1alpha1_modules_proto_rawDescData
}

var file_services_billing_entities_v1alpha1_modules_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_billing_entities_v1alpha1_modules_proto_goTypes = []any{
	(*BasicConfig)(nil),           // 0: services.billing.entities.v1alpha1.BasicConfig
	(*BasicUnitConfig)(nil),       // 1: services.billing.entities.v1alpha1.BasicUnitConfig
	(*wrapperspb.Int64Value)(nil), // 2: google.protobuf.Int64Value
}
var file_services_billing_entities_v1alpha1_modules_proto_depIdxs = []int32{
	2, // 0: services.billing.entities.v1alpha1.BasicUnitConfig.min_units:type_name -> google.protobuf.Int64Value
	2, // 1: services.billing.entities.v1alpha1.BasicUnitConfig.max_units:type_name -> google.protobuf.Int64Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha1_modules_proto_init() }
func file_services_billing_entities_v1alpha1_modules_proto_init() {
	if File_services_billing_entities_v1alpha1_modules_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha1_modules_proto_rawDesc), len(file_services_billing_entities_v1alpha1_modules_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha1_modules_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha1_modules_proto_depIdxs,
		MessageInfos:      file_services_billing_entities_v1alpha1_modules_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha1_modules_proto = out.File
	file_services_billing_entities_v1alpha1_modules_proto_goTypes = nil
	file_services_billing_entities_v1alpha1_modules_proto_depIdxs = nil
}
