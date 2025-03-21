// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha1/product.proto

package entitiesv1alpha1

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

// Product represents a billing product. A product should
// only show up once on an invoice.
type Product int32

const (
	Product_PRODUCT_UNSPECIFIED          Product = 0
	Product_PRODUCT_AGENT_SEATS          Product = 100
	Product_PRODUCT_OMNI                 Product = 200
	Product_PRODUCT_OMNI_CHAT_SENT       Product = 201
	Product_PRODUCT_OMNI_CHAT_RECEIVED   Product = 202
	Product_PRODUCT_OMNI_EMAILS_SENT     Product = 203
	Product_PRODUCT_OMNI_EMAILS_RECEIVED Product = 204
	Product_PRODUCT_OMNI_SMS_SENT        Product = 205
	Product_PRODUCT_OMNI_SMS_RECEIVED    Product = 206
	Product_PRODUCT_COMPLIANCE           Product = 300
)

// Enum value maps for Product.
var (
	Product_name = map[int32]string{
		0:   "PRODUCT_UNSPECIFIED",
		100: "PRODUCT_AGENT_SEATS",
		200: "PRODUCT_OMNI",
		201: "PRODUCT_OMNI_CHAT_SENT",
		202: "PRODUCT_OMNI_CHAT_RECEIVED",
		203: "PRODUCT_OMNI_EMAILS_SENT",
		204: "PRODUCT_OMNI_EMAILS_RECEIVED",
		205: "PRODUCT_OMNI_SMS_SENT",
		206: "PRODUCT_OMNI_SMS_RECEIVED",
		300: "PRODUCT_COMPLIANCE",
	}
	Product_value = map[string]int32{
		"PRODUCT_UNSPECIFIED":          0,
		"PRODUCT_AGENT_SEATS":          100,
		"PRODUCT_OMNI":                 200,
		"PRODUCT_OMNI_CHAT_SENT":       201,
		"PRODUCT_OMNI_CHAT_RECEIVED":   202,
		"PRODUCT_OMNI_EMAILS_SENT":     203,
		"PRODUCT_OMNI_EMAILS_RECEIVED": 204,
		"PRODUCT_OMNI_SMS_SENT":        205,
		"PRODUCT_OMNI_SMS_RECEIVED":    206,
		"PRODUCT_COMPLIANCE":           300,
	}
)

func (x Product) Enum() *Product {
	p := new(Product)
	*p = x
	return p
}

func (x Product) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Product) Descriptor() protoreflect.EnumDescriptor {
	return file_services_billing_entities_v1alpha1_product_proto_enumTypes[0].Descriptor()
}

func (Product) Type() protoreflect.EnumType {
	return &file_services_billing_entities_v1alpha1_product_proto_enumTypes[0]
}

func (x Product) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Product.Descriptor instead.
func (Product) EnumDescriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha1_product_proto_rawDescGZIP(), []int{0}
}

var File_services_billing_entities_v1alpha1_product_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha1_product_proto_rawDesc = string([]byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2a, 0xa3, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50,
	0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x41,
	0x54, 0x53, 0x10, 0x64, 0x12, 0x11, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f,
	0x4f, 0x4d, 0x4e, 0x49, 0x10, 0xc8, 0x01, 0x12, 0x1b, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x44, 0x55,
	0x43, 0x54, 0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x53, 0x45, 0x4e,
	0x54, 0x10, 0xc9, 0x01, 0x12, 0x1f, 0x0a, 0x1a, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f,
	0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56,
	0x45, 0x44, 0x10, 0xca, 0x01, 0x12, 0x1d, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54,
	0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x53, 0x5f, 0x53, 0x45, 0x4e,
	0x54, 0x10, 0xcb, 0x01, 0x12, 0x21, 0x0a, 0x1c, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f,
	0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x45,
	0x49, 0x56, 0x45, 0x44, 0x10, 0xcc, 0x01, 0x12, 0x1a, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x44, 0x55,
	0x43, 0x54, 0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x53, 0x4d, 0x53, 0x5f, 0x53, 0x45, 0x4e, 0x54,
	0x10, 0xcd, 0x01, 0x12, 0x1e, 0x0a, 0x19, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4f,
	0x4d, 0x4e, 0x49, 0x5f, 0x53, 0x4d, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44,
	0x10, 0xce, 0x01, 0x12, 0x17, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x10, 0xac, 0x02, 0x42, 0xb1, 0x02, 0x0a,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x42, 0x45, 0xaa, 0x02, 0x22,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_services_billing_entities_v1alpha1_product_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha1_product_proto_rawDescData []byte
)

func file_services_billing_entities_v1alpha1_product_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha1_product_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha1_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha1_product_proto_rawDesc), len(file_services_billing_entities_v1alpha1_product_proto_rawDesc)))
	})
	return file_services_billing_entities_v1alpha1_product_proto_rawDescData
}

var file_services_billing_entities_v1alpha1_product_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_billing_entities_v1alpha1_product_proto_goTypes = []any{
	(Product)(0), // 0: services.billing.entities.v1alpha1.Product
}
var file_services_billing_entities_v1alpha1_product_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha1_product_proto_init() }
func file_services_billing_entities_v1alpha1_product_proto_init() {
	if File_services_billing_entities_v1alpha1_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha1_product_proto_rawDesc), len(file_services_billing_entities_v1alpha1_product_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha1_product_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha1_product_proto_depIdxs,
		EnumInfos:         file_services_billing_entities_v1alpha1_product_proto_enumTypes,
	}.Build()
	File_services_billing_entities_v1alpha1_product_proto = out.File
	file_services_billing_entities_v1alpha1_product_proto_goTypes = nil
	file_services_billing_entities_v1alpha1_product_proto_depIdxs = nil
}
