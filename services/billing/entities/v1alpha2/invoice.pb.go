// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha2/invoice.proto

package entitiesv1alpha2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Product int32

const (
	Product_PRODUCT_UNSPECIFIED          Product = 0
	Product_PRODUCT_OMNI                 Product = 200
	Product_PRODUCT_OMNI_SEATS           Product = 201
	Product_PRODUCT_OMNI_CHAT_SENT       Product = 202
	Product_PRODUCT_OMNI_CHAT_RECEIVED   Product = 203
	Product_PRODUCT_OMNI_EMAILS_SENT     Product = 204
	Product_PRODUCT_OMNI_EMAILS_RECEIVED Product = 205
	Product_PRODUCT_OMNI_SMS_SENT        Product = 206
	Product_PRODUCT_OMNI_SMS_RECEIVED    Product = 207
	Product_PRODUCT_COMPLIANCE           Product = 300
)

// Enum value maps for Product.
var (
	Product_name = map[int32]string{
		0:   "PRODUCT_UNSPECIFIED",
		200: "PRODUCT_OMNI",
		201: "PRODUCT_OMNI_SEATS",
		202: "PRODUCT_OMNI_CHAT_SENT",
		203: "PRODUCT_OMNI_CHAT_RECEIVED",
		204: "PRODUCT_OMNI_EMAILS_SENT",
		205: "PRODUCT_OMNI_EMAILS_RECEIVED",
		206: "PRODUCT_OMNI_SMS_SENT",
		207: "PRODUCT_OMNI_SMS_RECEIVED",
		300: "PRODUCT_COMPLIANCE",
	}
	Product_value = map[string]int32{
		"PRODUCT_UNSPECIFIED":          0,
		"PRODUCT_OMNI":                 200,
		"PRODUCT_OMNI_SEATS":           201,
		"PRODUCT_OMNI_CHAT_SENT":       202,
		"PRODUCT_OMNI_CHAT_RECEIVED":   203,
		"PRODUCT_OMNI_EMAILS_SENT":     204,
		"PRODUCT_OMNI_EMAILS_RECEIVED": 205,
		"PRODUCT_OMNI_SMS_SENT":        206,
		"PRODUCT_OMNI_SMS_RECEIVED":    207,
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
	return file_services_billing_entities_v1alpha2_invoice_proto_enumTypes[0].Descriptor()
}

func (Product) Type() protoreflect.EnumType {
	return &file_services_billing_entities_v1alpha2_invoice_proto_enumTypes[0]
}

func (x Product) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Product.Descriptor instead.
func (Product) EnumDescriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha2_invoice_proto_rawDescGZIP(), []int{0}
}

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingCycle string                  `protobuf:"bytes,1,opt,name=billing_cycle,json=billingCycle,proto3" json:"billing_cycle,omitempty"`
	CreateTime   *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Items        []*InvoiceItem          `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	DownloadUrl  *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_billing_entities_v1alpha2_invoice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha2_invoice_proto_msgTypes[0]
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
	return file_services_billing_entities_v1alpha2_invoice_proto_rawDescGZIP(), []int{0}
}

func (x *Invoice) GetBillingCycle() string {
	if x != nil {
		return x.BillingCycle
	}
	return ""
}

func (x *Invoice) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Invoice) GetItems() []*InvoiceItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Invoice) GetDownloadUrl() *wrapperspb.StringValue {
	if x != nil {
		return x.DownloadUrl
	}
	return nil
}

// InvoiceItem represents a single line item on an invoice.
type InvoiceItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId    string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Product     Product                `protobuf:"varint,2,opt,name=product,proto3,enum=services.billing.entities.v1alpha2.Product" json:"product,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Date        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Price       float64                `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Columns     []*InvoiceItemColumn   `protobuf:"bytes,6,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *InvoiceItem) Reset() {
	*x = InvoiceItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_billing_entities_v1alpha2_invoice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceItem) ProtoMessage() {}

func (x *InvoiceItem) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha2_invoice_proto_msgTypes[1]
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
	return file_services_billing_entities_v1alpha2_invoice_proto_rawDescGZIP(), []int{1}
}

func (x *InvoiceItem) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *InvoiceItem) GetProduct() Product {
	if x != nil {
		return x.Product
	}
	return Product_PRODUCT_UNSPECIFIED
}

func (x *InvoiceItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InvoiceItem) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *InvoiceItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *InvoiceItem) GetColumns() []*InvoiceItemColumn {
	if x != nil {
		return x.Columns
	}
	return nil
}

// InvoiceItemColumn represents a single column on an invoice item.
type InvoiceItemColumn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InvoiceItemColumn) Reset() {
	*x = InvoiceItemColumn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_billing_entities_v1alpha2_invoice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceItemColumn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceItemColumn) ProtoMessage() {}

func (x *InvoiceItemColumn) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha2_invoice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceItemColumn.ProtoReflect.Descriptor instead.
func (*InvoiceItemColumn) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha2_invoice_proto_rawDescGZIP(), []int{2}
}

func (x *InvoiceItemColumn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InvoiceItemColumn) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_services_billing_entities_v1alpha2_invoice_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha2_invoice_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3f, 0x0a, 0x0c,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x22, 0xaa, 0x02,
	0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0xa3, 0x02, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x10, 0xc8,
	0x01, 0x12, 0x17, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4f, 0x4d, 0x4e,
	0x49, 0x5f, 0x53, 0x45, 0x41, 0x54, 0x53, 0x10, 0xc9, 0x01, 0x12, 0x1b, 0x0a, 0x16, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f,
	0x53, 0x45, 0x4e, 0x54, 0x10, 0xca, 0x01, 0x12, 0x1f, 0x0a, 0x1a, 0x50, 0x52, 0x4f, 0x44, 0x55,
	0x43, 0x54, 0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x52, 0x45, 0x43,
	0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0xcb, 0x01, 0x12, 0x1d, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x44,
	0x55, 0x43, 0x54, 0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x53, 0x5f,
	0x53, 0x45, 0x4e, 0x54, 0x10, 0xcc, 0x01, 0x12, 0x21, 0x0a, 0x1c, 0x50, 0x52, 0x4f, 0x44, 0x55,
	0x43, 0x54, 0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x53, 0x5f, 0x52,
	0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0xcd, 0x01, 0x12, 0x1a, 0x0a, 0x15, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x53, 0x4d, 0x53, 0x5f, 0x53,
	0x45, 0x4e, 0x54, 0x10, 0xce, 0x01, 0x12, 0x1e, 0x0a, 0x19, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43,
	0x54, 0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x53, 0x4d, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49,
	0x56, 0x45, 0x44, 0x10, 0xcf, 0x01, 0x12, 0x17, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43,
	0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x10, 0xac, 0x02, 0x42,
	0xb1, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0c, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68,
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
	file_services_billing_entities_v1alpha2_invoice_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha2_invoice_proto_rawDescData = file_services_billing_entities_v1alpha2_invoice_proto_rawDesc
)

func file_services_billing_entities_v1alpha2_invoice_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha2_invoice_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha2_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_billing_entities_v1alpha2_invoice_proto_rawDescData)
	})
	return file_services_billing_entities_v1alpha2_invoice_proto_rawDescData
}

var file_services_billing_entities_v1alpha2_invoice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_billing_entities_v1alpha2_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_billing_entities_v1alpha2_invoice_proto_goTypes = []interface{}{
	(Product)(0),                   // 0: services.billing.entities.v1alpha2.Product
	(*Invoice)(nil),                // 1: services.billing.entities.v1alpha2.Invoice
	(*InvoiceItem)(nil),            // 2: services.billing.entities.v1alpha2.InvoiceItem
	(*InvoiceItemColumn)(nil),      // 3: services.billing.entities.v1alpha2.InvoiceItemColumn
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
}
var file_services_billing_entities_v1alpha2_invoice_proto_depIdxs = []int32{
	4, // 0: services.billing.entities.v1alpha2.Invoice.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: services.billing.entities.v1alpha2.Invoice.items:type_name -> services.billing.entities.v1alpha2.InvoiceItem
	5, // 2: services.billing.entities.v1alpha2.Invoice.download_url:type_name -> google.protobuf.StringValue
	0, // 3: services.billing.entities.v1alpha2.InvoiceItem.product:type_name -> services.billing.entities.v1alpha2.Product
	4, // 4: services.billing.entities.v1alpha2.InvoiceItem.date:type_name -> google.protobuf.Timestamp
	3, // 5: services.billing.entities.v1alpha2.InvoiceItem.columns:type_name -> services.billing.entities.v1alpha2.InvoiceItemColumn
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha2_invoice_proto_init() }
func file_services_billing_entities_v1alpha2_invoice_proto_init() {
	if File_services_billing_entities_v1alpha2_invoice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_billing_entities_v1alpha2_invoice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_billing_entities_v1alpha2_invoice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_billing_entities_v1alpha2_invoice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvoiceItemColumn); i {
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
			RawDescriptor: file_services_billing_entities_v1alpha2_invoice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha2_invoice_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha2_invoice_proto_depIdxs,
		EnumInfos:         file_services_billing_entities_v1alpha2_invoice_proto_enumTypes,
		MessageInfos:      file_services_billing_entities_v1alpha2_invoice_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha2_invoice_proto = out.File
	file_services_billing_entities_v1alpha2_invoice_proto_rawDesc = nil
	file_services_billing_entities_v1alpha2_invoice_proto_goTypes = nil
	file_services_billing_entities_v1alpha2_invoice_proto_depIdxs = nil
}
