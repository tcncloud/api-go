// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha1/invoice.proto

package entitiesv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Invoice represents a collection of invoice items
// billed to an org for a given billing cycle.
type Invoice struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
	InvoiceId string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	// the billing cycle this invoice was created for
	BillingCycle string `protobuf:"bytes,2,opt,name=billing_cycle,json=billingCycle,proto3" json:"billing_cycle,omitempty"`
	// the time this invoice was created
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// the items on this invoice
	Items []*InvoiceItem `protobuf:"bytes,6,rep,name=items,proto3" json:"items,omitempty"`
	// the url to download the invoice
	Url *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	// the client this invoice is for
	//
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
	ClientId      string `protobuf:"bytes,8,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	mi := &file_services_billing_entities_v1alpha1_invoice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha1_invoice_proto_msgTypes[0]
	if x != nil {
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
	return file_services_billing_entities_v1alpha1_invoice_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
func (x *Invoice) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
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

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
func (x *Invoice) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
func (x *Invoice) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *Invoice) GetItems() []*InvoiceItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Invoice) GetUrl() *wrapperspb.StringValue {
	if x != nil {
		return x.Url
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
func (x *Invoice) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

// InvoiceItem represents a single line item on an invoice.
type InvoiceItem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
	InvoiceItemId string `protobuf:"bytes,1,opt,name=invoice_item_id,json=invoiceItemId,proto3" json:"invoice_item_id,omitempty"`
	// the product this item represents
	Product Product `protobuf:"varint,2,opt,name=product,proto3,enum=services.billing.entities.v1alpha1.Product" json:"product,omitempty"`
	// the price for this product
	Price float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// description of the item
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// the time this item was made
	Date *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	// other data columns
	Columns []*InvoiceItemColumn `protobuf:"bytes,8,rep,name=columns,proto3" json:"columns,omitempty"`
	// the client this item is for
	ClientId      string `protobuf:"bytes,9,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InvoiceItem) Reset() {
	*x = InvoiceItem{}
	mi := &file_services_billing_entities_v1alpha1_invoice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvoiceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceItem) ProtoMessage() {}

func (x *InvoiceItem) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha1_invoice_proto_msgTypes[1]
	if x != nil {
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
	return file_services_billing_entities_v1alpha1_invoice_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
func (x *InvoiceItem) GetInvoiceItemId() string {
	if x != nil {
		return x.InvoiceItemId
	}
	return ""
}

func (x *InvoiceItem) GetProduct() Product {
	if x != nil {
		return x.Product
	}
	return Product_PRODUCT_UNSPECIFIED
}

func (x *InvoiceItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
func (x *InvoiceItem) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
func (x *InvoiceItem) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
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

func (x *InvoiceItem) GetColumns() []*InvoiceItemColumn {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *InvoiceItem) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

// InvoiceItemColumn represents a single column on an invoice item.
type InvoiceItemColumn struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
	Value         int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	ColumnValue   string `protobuf:"bytes,3,opt,name=column_value,json=columnValue,proto3" json:"column_value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InvoiceItemColumn) Reset() {
	*x = InvoiceItemColumn{}
	mi := &file_services_billing_entities_v1alpha1_invoice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvoiceItemColumn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceItemColumn) ProtoMessage() {}

func (x *InvoiceItemColumn) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha1_invoice_proto_msgTypes[2]
	if x != nil {
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
	return file_services_billing_entities_v1alpha1_invoice_proto_rawDescGZIP(), []int{2}
}

func (x *InvoiceItemColumn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha1/invoice.proto.
func (x *InvoiceItemColumn) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *InvoiceItemColumn) GetColumnValue() string {
	if x != nil {
		return x.ColumnValue
	}
	return ""
}

var File_services_billing_entities_v1alpha1_invoice_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha1_invoice_proto_rawDesc = string([]byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x03, 0x0a, 0x07, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0xd8, 0x03, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x2a, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x45, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3f,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x4f, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x64, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xb1, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x0c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x53, 0x42, 0x45, 0xaa, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x22, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x25, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_services_billing_entities_v1alpha1_invoice_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha1_invoice_proto_rawDescData []byte
)

func file_services_billing_entities_v1alpha1_invoice_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha1_invoice_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha1_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha1_invoice_proto_rawDesc), len(file_services_billing_entities_v1alpha1_invoice_proto_rawDesc)))
	})
	return file_services_billing_entities_v1alpha1_invoice_proto_rawDescData
}

var file_services_billing_entities_v1alpha1_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_billing_entities_v1alpha1_invoice_proto_goTypes = []any{
	(*Invoice)(nil),                // 0: services.billing.entities.v1alpha1.Invoice
	(*InvoiceItem)(nil),            // 1: services.billing.entities.v1alpha1.InvoiceItem
	(*InvoiceItemColumn)(nil),      // 2: services.billing.entities.v1alpha1.InvoiceItemColumn
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(Product)(0),                   // 5: services.billing.entities.v1alpha1.Product
}
var file_services_billing_entities_v1alpha1_invoice_proto_depIdxs = []int32{
	3,  // 0: services.billing.entities.v1alpha1.Invoice.create_time:type_name -> google.protobuf.Timestamp
	3,  // 1: services.billing.entities.v1alpha1.Invoice.update_time:type_name -> google.protobuf.Timestamp
	3,  // 2: services.billing.entities.v1alpha1.Invoice.delete_time:type_name -> google.protobuf.Timestamp
	1,  // 3: services.billing.entities.v1alpha1.Invoice.items:type_name -> services.billing.entities.v1alpha1.InvoiceItem
	4,  // 4: services.billing.entities.v1alpha1.Invoice.url:type_name -> google.protobuf.StringValue
	5,  // 5: services.billing.entities.v1alpha1.InvoiceItem.product:type_name -> services.billing.entities.v1alpha1.Product
	3,  // 6: services.billing.entities.v1alpha1.InvoiceItem.create_time:type_name -> google.protobuf.Timestamp
	3,  // 7: services.billing.entities.v1alpha1.InvoiceItem.update_time:type_name -> google.protobuf.Timestamp
	3,  // 8: services.billing.entities.v1alpha1.InvoiceItem.date:type_name -> google.protobuf.Timestamp
	2,  // 9: services.billing.entities.v1alpha1.InvoiceItem.columns:type_name -> services.billing.entities.v1alpha1.InvoiceItemColumn
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha1_invoice_proto_init() }
func file_services_billing_entities_v1alpha1_invoice_proto_init() {
	if File_services_billing_entities_v1alpha1_invoice_proto != nil {
		return
	}
	file_services_billing_entities_v1alpha1_product_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha1_invoice_proto_rawDesc), len(file_services_billing_entities_v1alpha1_invoice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha1_invoice_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha1_invoice_proto_depIdxs,
		MessageInfos:      file_services_billing_entities_v1alpha1_invoice_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha1_invoice_proto = out.File
	file_services_billing_entities_v1alpha1_invoice_proto_goTypes = nil
	file_services_billing_entities_v1alpha1_invoice_proto_depIdxs = nil
}
