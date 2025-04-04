// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/commons/billing/modules/modules.proto

package modules

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

// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
type BasicConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
	Rate          float64 `protobuf:"fixed64,1,opt,name=rate,proto3" json:"rate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BasicConfig) Reset() {
	*x = BasicConfig{}
	mi := &file_api_commons_billing_modules_modules_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BasicConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicConfig) ProtoMessage() {}

func (x *BasicConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_billing_modules_modules_proto_msgTypes[0]
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
	Rate float64 `protobuf:"fixed64,2,opt,name=rate,proto3" json:"rate,omitempty"`
	// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
	MinIncrement *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=min_increment,json=minIncrement,proto3" json:"min_increment,omitempty"`
	// Deprecated: Marked as deprecated in api/commons/billing/modules/modules.proto.
	MaxIncrement  *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=max_increment,json=maxIncrement,proto3" json:"max_increment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BasicAmountConfig) Reset() {
	*x = BasicAmountConfig{}
	mi := &file_api_commons_billing_modules_modules_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BasicAmountConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicAmountConfig) ProtoMessage() {}

func (x *BasicAmountConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_billing_modules_modules_proto_msgTypes[1]
	if x != nil {
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

const file_api_commons_billing_modules_modules_proto_rawDesc = "" +
	"\n" +
	")api/commons/billing/modules/modules.proto\x12\x1bapi.commons.billing.modules\x1a\x1egoogle/protobuf/wrappers.proto\")\n" +
	"\vBasicConfig\x12\x16\n" +
	"\x04rate\x18\x01 \x01(\x01B\x02\x18\x01R\x04rate:\x02\x18\x01\"\xd7\x01\n" +
	"\x11BasicAmountConfig\x12\x1a\n" +
	"\x06amount\x18\x01 \x01(\x03B\x02\x18\x01R\x06amount\x12\x16\n" +
	"\x04rate\x18\x02 \x01(\x01B\x02\x18\x01R\x04rate\x12D\n" +
	"\rmin_increment\x18\x03 \x01(\v2\x1b.google.protobuf.Int64ValueB\x02\x18\x01R\fminIncrement\x12D\n" +
	"\rmax_increment\x18\x04 \x01(\v2\x1b.google.protobuf.Int64ValueB\x02\x18\x01R\fmaxIncrement:\x02\x18\x01B\xf7\x01\n" +
	"\x1fcom.api.commons.billing.modulesB\fModulesProtoP\x01Z6github.com/tcncloud/api-go/api/commons/billing/modules\xa2\x02\x04ACBM\xaa\x02\x1bApi.Commons.Billing.Modules\xca\x02\x1bApi\\Commons\\Billing\\Modules\xe2\x02'Api\\Commons\\Billing\\Modules\\GPBMetadata\xea\x02\x1eApi::Commons::Billing::Modulesb\x06proto3"

var (
	file_api_commons_billing_modules_modules_proto_rawDescOnce sync.Once
	file_api_commons_billing_modules_modules_proto_rawDescData []byte
)

func file_api_commons_billing_modules_modules_proto_rawDescGZIP() []byte {
	file_api_commons_billing_modules_modules_proto_rawDescOnce.Do(func() {
		file_api_commons_billing_modules_modules_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_billing_modules_modules_proto_rawDesc), len(file_api_commons_billing_modules_modules_proto_rawDesc)))
	})
	return file_api_commons_billing_modules_modules_proto_rawDescData
}

var file_api_commons_billing_modules_modules_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_billing_modules_modules_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_billing_modules_modules_proto_rawDesc), len(file_api_commons_billing_modules_modules_proto_rawDesc)),
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
	file_api_commons_billing_modules_modules_proto_goTypes = nil
	file_api_commons_billing_modules_modules_proto_depIdxs = nil
}
