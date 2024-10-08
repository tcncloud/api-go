// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: services/billing/v1alpha1/core.proto

package billingv1alpha1

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

type Order int32

const (
	Order_ORDER_UNSPECIFIED Order = 0 // If unspecified, sorts ascending
	Order_ORDER_DESC        Order = 1
)

// Enum value maps for Order.
var (
	Order_name = map[int32]string{
		0: "ORDER_UNSPECIFIED",
		1: "ORDER_DESC",
	}
	Order_value = map[string]int32{
		"ORDER_UNSPECIFIED": 0,
		"ORDER_DESC":        1,
	}
)

func (x Order) Enum() *Order {
	p := new(Order)
	*p = x
	return p
}

func (x Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order) Descriptor() protoreflect.EnumDescriptor {
	return file_services_billing_v1alpha1_core_proto_enumTypes[0].Descriptor()
}

func (Order) Type() protoreflect.EnumType {
	return &file_services_billing_v1alpha1_core_proto_enumTypes[0]
}

func (x Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order.Descriptor instead.
func (Order) EnumDescriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_core_proto_rawDescGZIP(), []int{0}
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"` // Optional: if not specified, will not paginate
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`  // Optional: if not specified, will start from the beginning
}

func (x *Page) Reset() {
	*x = Page{}
	mi := &file_services_billing_v1alpha1_core_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_core_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_core_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Page) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Sort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`                                                   // Required: key to sort by
	Direction Order  `protobuf:"varint,2,opt,name=direction,proto3,enum=services.billing.v1alpha1.Order" json:"direction,omitempty"` // Required: direction to sort by
}

func (x *Sort) Reset() {
	*x = Sort{}
	mi := &file_services_billing_v1alpha1_core_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_core_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_core_proto_rawDescGZIP(), []int{1}
}

func (x *Sort) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Sort) GetDirection() Order {
	if x != nil {
		return x.Direction
	}
	return Order_ORDER_UNSPECIFIED
}

var File_services_billing_v1alpha1_core_proto protoreflect.FileDescriptor

var file_services_billing_v1alpha1_core_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x22, 0x32, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x58, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x3e, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2a,
	0x2e, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x42,
	0xf6, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x42, 0x09, 0x43, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x3b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x42, 0x58, 0xaa, 0x02, 0x19, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x19, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xe2, 0x02, 0x25, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_billing_v1alpha1_core_proto_rawDescOnce sync.Once
	file_services_billing_v1alpha1_core_proto_rawDescData = file_services_billing_v1alpha1_core_proto_rawDesc
)

func file_services_billing_v1alpha1_core_proto_rawDescGZIP() []byte {
	file_services_billing_v1alpha1_core_proto_rawDescOnce.Do(func() {
		file_services_billing_v1alpha1_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_billing_v1alpha1_core_proto_rawDescData)
	})
	return file_services_billing_v1alpha1_core_proto_rawDescData
}

var file_services_billing_v1alpha1_core_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_billing_v1alpha1_core_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_billing_v1alpha1_core_proto_goTypes = []any{
	(Order)(0),   // 0: services.billing.v1alpha1.Order
	(*Page)(nil), // 1: services.billing.v1alpha1.Page
	(*Sort)(nil), // 2: services.billing.v1alpha1.Sort
}
var file_services_billing_v1alpha1_core_proto_depIdxs = []int32{
	0, // 0: services.billing.v1alpha1.Sort.direction:type_name -> services.billing.v1alpha1.Order
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_billing_v1alpha1_core_proto_init() }
func file_services_billing_v1alpha1_core_proto_init() {
	if File_services_billing_v1alpha1_core_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_billing_v1alpha1_core_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_v1alpha1_core_proto_goTypes,
		DependencyIndexes: file_services_billing_v1alpha1_core_proto_depIdxs,
		EnumInfos:         file_services_billing_v1alpha1_core_proto_enumTypes,
		MessageInfos:      file_services_billing_v1alpha1_core_proto_msgTypes,
	}.Build()
	File_services_billing_v1alpha1_core_proto = out.File
	file_services_billing_v1alpha1_core_proto_rawDesc = nil
	file_services_billing_v1alpha1_core_proto_goTypes = nil
	file_services_billing_v1alpha1_core_proto_depIdxs = nil
}
