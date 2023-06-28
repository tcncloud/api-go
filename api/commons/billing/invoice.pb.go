// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/commons/billing/invoice.proto

package billing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Product - represents a billed product. A product should show up only
// once in a monthly invoice.
type Product int32

const (
	Product_PRODUCT_UNSPECIFIED     Product = 0
	Product_PRODUCT_OTHER           Product = 1
	Product_PRODUCT_AGENT_SEATS     Product = 100
	Product_PRODUCT_EMAILS_SENT     Product = 200
	Product_PRODUCT_EMAILS_RECEIVED Product = 201
	Product_PRODUCT_SMS_SENT        Product = 202
	Product_PRODUCT_SMS_RECEIVED    Product = 203
	Product_PRODUCT_CHAT_SENT       Product = 204
	Product_PRODUCT_CHAT_RECEIVED   Product = 205
	Product_PRODUCT_OMNI            Product = 300
	Product_PRODUCT_VANA            Product = 400
	Product_PRODUCT_COMPLIANCE      Product = 500
)

// Enum value maps for Product.
var (
	Product_name = map[int32]string{
		0:   "PRODUCT_UNSPECIFIED",
		1:   "PRODUCT_OTHER",
		100: "PRODUCT_AGENT_SEATS",
		200: "PRODUCT_EMAILS_SENT",
		201: "PRODUCT_EMAILS_RECEIVED",
		202: "PRODUCT_SMS_SENT",
		203: "PRODUCT_SMS_RECEIVED",
		204: "PRODUCT_CHAT_SENT",
		205: "PRODUCT_CHAT_RECEIVED",
		300: "PRODUCT_OMNI",
		400: "PRODUCT_VANA",
		500: "PRODUCT_COMPLIANCE",
	}
	Product_value = map[string]int32{
		"PRODUCT_UNSPECIFIED":     0,
		"PRODUCT_OTHER":           1,
		"PRODUCT_AGENT_SEATS":     100,
		"PRODUCT_EMAILS_SENT":     200,
		"PRODUCT_EMAILS_RECEIVED": 201,
		"PRODUCT_SMS_SENT":        202,
		"PRODUCT_SMS_RECEIVED":    203,
		"PRODUCT_CHAT_SENT":       204,
		"PRODUCT_CHAT_RECEIVED":   205,
		"PRODUCT_OMNI":            300,
		"PRODUCT_VANA":            400,
		"PRODUCT_COMPLIANCE":      500,
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
	return file_api_commons_billing_invoice_proto_enumTypes[0].Descriptor()
}

func (Product) Type() protoreflect.EnumType {
	return &file_api_commons_billing_invoice_proto_enumTypes[0]
}

func (x Product) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Product.Descriptor instead.
func (Product) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_billing_invoice_proto_rawDescGZIP(), []int{0}
}

// Invoice - a collection of products with their total rated amount for
// a selected month or month-to-date.
type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The items forming the invoice, where a product type
	// should show up only once in this list.
	Items []*InvoiceItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_billing_invoice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_billing_invoice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_api_commons_billing_invoice_proto_rawDescGZIP(), []int{0}
}

func (x *Invoice) GetItems() []*InvoiceItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// InvoiceItem -
type InvoiceItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier for this invoice item
	InvoiceItemSid int64 `protobuf:"varint,1,opt,name=invoice_item_sid,json=invoiceItemSid,proto3" json:"invoice_item_sid,omitempty"`
	// The product this item contains the total for
	Product Product `protobuf:"varint,2,opt,name=product,proto3,enum=api.commons.billing.Product" json:"product,omitempty"`
	// The total amount billed for the product
	Amount float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// time the invoice item was created
	DateCreated *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	// time the invoice item was last modified
	DateModified *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date_modified,json=dateModified,proto3" json:"date_modified,omitempty"`
}

func (x *InvoiceItem) Reset() {
	*x = InvoiceItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_billing_invoice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceItem) ProtoMessage() {}

func (x *InvoiceItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_billing_invoice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceItem.ProtoReflect.Descriptor instead.
func (*InvoiceItem) Descriptor() ([]byte, []int) {
	return file_api_commons_billing_invoice_proto_rawDescGZIP(), []int{1}
}

func (x *InvoiceItem) GetInvoiceItemSid() int64 {
	if x != nil {
		return x.InvoiceItemSid
	}
	return 0
}

func (x *InvoiceItem) GetProduct() Product {
	if x != nil {
		return x.Product
	}
	return Product_PRODUCT_UNSPECIFIED
}

func (x *InvoiceItem) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *InvoiceItem) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

func (x *InvoiceItem) GetDateModified() *timestamppb.Timestamp {
	if x != nil {
		return x.DateModified
	}
	return nil
}

var File_api_commons_billing_invoice_proto protoreflect.FileDescriptor

var file_api_commons_billing_invoice_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x07, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x87, 0x02, 0x0a,
	0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x28, 0x0a, 0x10,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x53, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2a, 0xab, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x50,
	0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x45, 0x41, 0x54, 0x53, 0x10, 0x64, 0x12, 0x18, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x44, 0x55,
	0x43, 0x54, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x53, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x10, 0xc8,
	0x01, 0x12, 0x1c, 0x0a, 0x17, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x45, 0x4d, 0x41,
	0x49, 0x4c, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0xc9, 0x01, 0x12,
	0x15, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x53, 0x4d, 0x53, 0x5f, 0x53,
	0x45, 0x4e, 0x54, 0x10, 0xca, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43,
	0x54, 0x5f, 0x53, 0x4d, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0xcb,
	0x01, 0x12, 0x16, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x54, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x10, 0xcc, 0x01, 0x12, 0x1a, 0x0a, 0x15, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56,
	0x45, 0x44, 0x10, 0xcd, 0x01, 0x12, 0x11, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54,
	0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x10, 0xac, 0x02, 0x12, 0x11, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x44,
	0x55, 0x43, 0x54, 0x5f, 0x56, 0x41, 0x4e, 0x41, 0x10, 0x90, 0x03, 0x12, 0x17, 0x0a, 0x12, 0x50,
	0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43,
	0x45, 0x10, 0xf4, 0x03, 0x42, 0xc5, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x42, 0x0c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0xa2, 0x02, 0x03, 0x41, 0x43, 0x42, 0xaa, 0x02, 0x13, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x13, 0x41,
	0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0xe2, 0x02, 0x1f, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_billing_invoice_proto_rawDescOnce sync.Once
	file_api_commons_billing_invoice_proto_rawDescData = file_api_commons_billing_invoice_proto_rawDesc
)

func file_api_commons_billing_invoice_proto_rawDescGZIP() []byte {
	file_api_commons_billing_invoice_proto_rawDescOnce.Do(func() {
		file_api_commons_billing_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_billing_invoice_proto_rawDescData)
	})
	return file_api_commons_billing_invoice_proto_rawDescData
}

var file_api_commons_billing_invoice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_billing_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_billing_invoice_proto_goTypes = []interface{}{
	(Product)(0),                  // 0: api.commons.billing.Product
	(*Invoice)(nil),               // 1: api.commons.billing.Invoice
	(*InvoiceItem)(nil),           // 2: api.commons.billing.InvoiceItem
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_api_commons_billing_invoice_proto_depIdxs = []int32{
	2, // 0: api.commons.billing.Invoice.items:type_name -> api.commons.billing.InvoiceItem
	0, // 1: api.commons.billing.InvoiceItem.product:type_name -> api.commons.billing.Product
	3, // 2: api.commons.billing.InvoiceItem.date_created:type_name -> google.protobuf.Timestamp
	3, // 3: api.commons.billing.InvoiceItem.date_modified:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_commons_billing_invoice_proto_init() }
func file_api_commons_billing_invoice_proto_init() {
	if File_api_commons_billing_invoice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commons_billing_invoice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
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
		file_api_commons_billing_invoice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvoiceItem); i {
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
			RawDescriptor: file_api_commons_billing_invoice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_billing_invoice_proto_goTypes,
		DependencyIndexes: file_api_commons_billing_invoice_proto_depIdxs,
		EnumInfos:         file_api_commons_billing_invoice_proto_enumTypes,
		MessageInfos:      file_api_commons_billing_invoice_proto_msgTypes,
	}.Build()
	File_api_commons_billing_invoice_proto = out.File
	file_api_commons_billing_invoice_proto_rawDesc = nil
	file_api_commons_billing_invoice_proto_goTypes = nil
	file_api_commons_billing_invoice_proto_depIdxs = nil
}
