// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha2/modules.proto

package entitiesv1alpha2

import (
	decimal "google.golang.org/genproto/googleapis/type/decimal"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha2/modules.proto.
	Rate float64 `protobuf:"fixed64,1,opt,name=rate,proto3" json:"rate,omitempty"`
	// Required. the amount to rate a single event
	RateDecimal *decimal.Decimal `protobuf:"bytes,2,opt,name=rate_decimal,json=rateDecimal,proto3" json:"rate_decimal,omitempty"`
}

func (x *BasicConfig) Reset() {
	*x = BasicConfig{}
	mi := &file_services_billing_entities_v1alpha2_modules_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BasicConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicConfig) ProtoMessage() {}

func (x *BasicConfig) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha2_modules_proto_msgTypes[0]
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
	return file_services_billing_entities_v1alpha2_modules_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha2/modules.proto.
func (x *BasicConfig) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *BasicConfig) GetRateDecimal() *decimal.Decimal {
	if x != nil {
		return x.RateDecimal
	}
	return nil
}

// BasicUnitConfig represents basic config for a rating module that
// rates based on event units
type BasicUnitConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. the size of an individual unit. For example, a unit
	// size of 300 (with bytes as a unit of measurement) and an event
	// of size 400 bytes will get billed as if it were 2 units.
	UnitSize int64 `protobuf:"varint,1,opt,name=unit_size,json=unitSize,proto3" json:"unit_size,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha2/modules.proto.
	Rate float64 `protobuf:"fixed64,2,opt,name=rate,proto3" json:"rate,omitempty"`
	// Optional. the minimum number of units to rate; for example, a
	// unit size of 5 seconds and a min_units of 2 would mean that
	// any event less than 10 seconds would be billed as if it were
	// two units (10 seconds). This is per event.
	MinUnits *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=min_units,json=minUnits,proto3" json:"min_units,omitempty"`
	// Optional. the maximum number of units to rate; for example, a
	// unit size of 5 seconds and a max_units of 2 would mean that
	// any event more than 10 seconds would be billed as if it were
	// two units (10 seconds). This is per event.
	MaxUnits *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=max_units,json=maxUnits,proto3" json:"max_units,omitempty"`
	// Optional. the minimum number of units to rate per cycle
	MinUnitsPerCycle *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=min_units_per_cycle,json=minUnitsPerCycle,proto3" json:"min_units_per_cycle,omitempty"`
	// Optional. the maximum number of units to rate per cycle
	MaxUnitsPerCycle *wrapperspb.Int64Value `protobuf:"bytes,6,opt,name=max_units_per_cycle,json=maxUnitsPerCycle,proto3" json:"max_units_per_cycle,omitempty"`
	// Required. the amount to rate each unit
	RateDecimal *decimal.Decimal `protobuf:"bytes,7,opt,name=rate_decimal,json=rateDecimal,proto3" json:"rate_decimal,omitempty"`
}

func (x *BasicUnitConfig) Reset() {
	*x = BasicUnitConfig{}
	mi := &file_services_billing_entities_v1alpha2_modules_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BasicUnitConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicUnitConfig) ProtoMessage() {}

func (x *BasicUnitConfig) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha2_modules_proto_msgTypes[1]
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
	return file_services_billing_entities_v1alpha2_modules_proto_rawDescGZIP(), []int{1}
}

func (x *BasicUnitConfig) GetUnitSize() int64 {
	if x != nil {
		return x.UnitSize
	}
	return 0
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha2/modules.proto.
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

func (x *BasicUnitConfig) GetMinUnitsPerCycle() *wrapperspb.Int64Value {
	if x != nil {
		return x.MinUnitsPerCycle
	}
	return nil
}

func (x *BasicUnitConfig) GetMaxUnitsPerCycle() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxUnitsPerCycle
	}
	return nil
}

func (x *BasicUnitConfig) GetRateDecimal() *decimal.Decimal {
	if x != nil {
		return x.RateDecimal
	}
	return nil
}

var File_services_billing_entities_v1alpha2_modules_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha2_modules_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5e, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x16, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x63,
	0x69, 0x6d, 0x61, 0x6c, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61,
	0x6c, 0x22, 0x8b, 0x03, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x55, 0x6e, 0x69, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x16, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x6d, 0x69,
	0x6e, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x6e, 0x69, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x4a,
	0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x55, 0x6e, 0x69,
	0x74, 0x73, 0x50, 0x65, 0x72, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x13, 0x6d, 0x61,
	0x78, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x50, 0x65,
	0x72, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x42,
	0xb1, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0c, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2, 0x02, 0x03, 0x53, 0x42, 0x45,
	0xaa, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xe2, 0x02, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a,
	0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_billing_entities_v1alpha2_modules_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha2_modules_proto_rawDescData = file_services_billing_entities_v1alpha2_modules_proto_rawDesc
)

func file_services_billing_entities_v1alpha2_modules_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha2_modules_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha2_modules_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_billing_entities_v1alpha2_modules_proto_rawDescData)
	})
	return file_services_billing_entities_v1alpha2_modules_proto_rawDescData
}

var file_services_billing_entities_v1alpha2_modules_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_billing_entities_v1alpha2_modules_proto_goTypes = []any{
	(*BasicConfig)(nil),           // 0: services.billing.entities.v1alpha2.BasicConfig
	(*BasicUnitConfig)(nil),       // 1: services.billing.entities.v1alpha2.BasicUnitConfig
	(*decimal.Decimal)(nil),       // 2: google.type.Decimal
	(*wrapperspb.Int64Value)(nil), // 3: google.protobuf.Int64Value
}
var file_services_billing_entities_v1alpha2_modules_proto_depIdxs = []int32{
	2, // 0: services.billing.entities.v1alpha2.BasicConfig.rate_decimal:type_name -> google.type.Decimal
	3, // 1: services.billing.entities.v1alpha2.BasicUnitConfig.min_units:type_name -> google.protobuf.Int64Value
	3, // 2: services.billing.entities.v1alpha2.BasicUnitConfig.max_units:type_name -> google.protobuf.Int64Value
	3, // 3: services.billing.entities.v1alpha2.BasicUnitConfig.min_units_per_cycle:type_name -> google.protobuf.Int64Value
	3, // 4: services.billing.entities.v1alpha2.BasicUnitConfig.max_units_per_cycle:type_name -> google.protobuf.Int64Value
	2, // 5: services.billing.entities.v1alpha2.BasicUnitConfig.rate_decimal:type_name -> google.type.Decimal
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha2_modules_proto_init() }
func file_services_billing_entities_v1alpha2_modules_proto_init() {
	if File_services_billing_entities_v1alpha2_modules_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_billing_entities_v1alpha2_modules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha2_modules_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha2_modules_proto_depIdxs,
		MessageInfos:      file_services_billing_entities_v1alpha2_modules_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha2_modules_proto = out.File
	file_services_billing_entities_v1alpha2_modules_proto_rawDesc = nil
	file_services_billing_entities_v1alpha2_modules_proto_goTypes = nil
	file_services_billing_entities_v1alpha2_modules_proto_depIdxs = nil
}
