// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: api/commons/billing/modules/modules.proto

package modules

import (
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

// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
type BasicConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
	Rate float64 `protobuf:"fixed64,1,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *BasicConfig) Reset() {
	*x = BasicConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_billing_modules_modules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicConfig) ProtoMessage() {}

func (x *BasicConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_billing_modules_modules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_commons_billing_modules_modules_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
func (x *BasicConfig) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
type BasicAmountConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
	Rate float64 `protobuf:"fixed64,2,opt,name=rate,proto3" json:"rate,omitempty"`
	// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
	MinIncrement *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=min_increment,json=minIncrement,proto3" json:"min_increment,omitempty"`
	// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
	MaxIncrement *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=max_increment,json=maxIncrement,proto3" json:"max_increment,omitempty"`
}

func (x *BasicAmountConfig) Reset() {
	*x = BasicAmountConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_billing_modules_modules_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicAmountConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicAmountConfig) ProtoMessage() {}

func (x *BasicAmountConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_billing_modules_modules_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicAmountConfig.ProtoReflect.Descriptor instead.
func (*BasicAmountConfig) Descriptor() ([]byte, []int) {
	return file_api_commons_billing_modules_modules_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
func (x *BasicAmountConfig) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
func (x *BasicAmountConfig) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
func (x *BasicAmountConfig) GetMinIncrement() *wrapperspb.Int64Value {
	if x != nil {
		return x.MinIncrement
	}
	return nil
}

// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
func (x *BasicAmountConfig) GetMaxIncrement() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxIncrement
	}
	return nil
}

var File_api_commons_billing_modules_modules_proto protoreflect.FileDescriptor

var file_api_commons_billing_modules_modules_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x3a,
	0x02, 0x18, 0x01, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x44, 0x0a,
	0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0c, 0x6d, 0x61, 0x78,
	0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x02, 0x18, 0x01, 0x42, 0xf7, 0x01,
	0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x42, 0x0c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0xa2, 0x02, 0x04, 0x41, 0x43, 0x42, 0x4d,
	0xaa, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0xca, 0x02,
	0x1b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0xe2, 0x02, 0x27, 0x41,
	0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_billing_modules_modules_proto_rawDescOnce sync.Once
	file_api_commons_billing_modules_modules_proto_rawDescData = file_api_commons_billing_modules_modules_proto_rawDesc
)

func file_api_commons_billing_modules_modules_proto_rawDescGZIP() []byte {
	file_api_commons_billing_modules_modules_proto_rawDescOnce.Do(func() {
		file_api_commons_billing_modules_modules_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_billing_modules_modules_proto_rawDescData)
	})
	return file_api_commons_billing_modules_modules_proto_rawDescData
}

var file_api_commons_billing_modules_modules_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_billing_modules_modules_proto_goTypes = []interface{}{
	(*BasicConfig)(nil),           // 0: api.commons.billing.modules.BasicConfig
	(*BasicAmountConfig)(nil),     // 1: api.commons.billing.modules.BasicAmountConfig
	(*wrapperspb.Int64Value)(nil), // 2: google.protobuf.Int64Value
}
var file_api_commons_billing_modules_modules_proto_depIdxs = []int32{
	2, // 0: api.commons.billing.modules.BasicAmountConfig.min_increment:type_name -> google.protobuf.Int64Value
	2, // 1: api.commons.billing.modules.BasicAmountConfig.max_increment:type_name -> google.protobuf.Int64Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_commons_billing_modules_modules_proto_init() }
func file_api_commons_billing_modules_modules_proto_init() {
	if File_api_commons_billing_modules_modules_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commons_billing_modules_modules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_commons_billing_modules_modules_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicAmountConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_billing_modules_modules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_billing_modules_modules_proto_goTypes,
		DependencyIndexes: file_api_commons_billing_modules_modules_proto_depIdxs,
		MessageInfos:      file_api_commons_billing_modules_modules_proto_msgTypes,
	}.Build()
	File_api_commons_billing_modules_modules_proto = out.File
	file_api_commons_billing_modules_modules_proto_rawDesc = nil
	file_api_commons_billing_modules_modules_proto_goTypes = nil
	file_api_commons_billing_modules_modules_proto_depIdxs = nil
}
