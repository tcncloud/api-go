// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha3/invoice.proto

package entitiesv1alpha3

import (
	decimal "google.golang.org/genproto/googleapis/type/decimal"
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

type Invoice struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	BillingCycle  string                  `protobuf:"bytes,1,opt,name=billing_cycle,json=billingCycle,proto3" json:"billing_cycle,omitempty"`
	CreateTime    *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Rows          []*InvoiceRow           `protobuf:"bytes,3,rep,name=rows,proto3" json:"rows,omitempty"`
	DownloadUrl   *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	mi := &file_services_billing_entities_v1alpha3_invoice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha3_invoice_proto_msgTypes[0]
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
	return file_services_billing_entities_v1alpha3_invoice_proto_rawDescGZIP(), []int{0}
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

func (x *Invoice) GetRows() []*InvoiceRow {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *Invoice) GetDownloadUrl() *wrapperspb.StringValue {
	if x != nil {
		return x.DownloadUrl
	}
	return nil
}

type InvoiceRow struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ProductId     string                 `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Columns       []*InvoiceColumn       `protobuf:"bytes,5,rep,name=columns,proto3" json:"columns,omitempty"`
	RatedAmount   *decimal.Decimal       `protobuf:"bytes,6,opt,name=rated_amount,json=ratedAmount,proto3" json:"rated_amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InvoiceRow) Reset() {
	*x = InvoiceRow{}
	mi := &file_services_billing_entities_v1alpha3_invoice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvoiceRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceRow) ProtoMessage() {}

func (x *InvoiceRow) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha3_invoice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceRow.ProtoReflect.Descriptor instead.
func (*InvoiceRow) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha3_invoice_proto_rawDescGZIP(), []int{1}
}

func (x *InvoiceRow) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *InvoiceRow) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *InvoiceRow) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InvoiceRow) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *InvoiceRow) GetColumns() []*InvoiceColumn {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *InvoiceRow) GetRatedAmount() *decimal.Decimal {
	if x != nil {
		return x.RatedAmount
	}
	return nil
}

type InvoiceColumn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InvoiceColumn) Reset() {
	*x = InvoiceColumn{}
	mi := &file_services_billing_entities_v1alpha3_invoice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvoiceColumn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceColumn) ProtoMessage() {}

func (x *InvoiceColumn) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha3_invoice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceColumn.ProtoReflect.Descriptor instead.
func (*InvoiceColumn) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha3_invoice_proto_rawDescGZIP(), []int{2}
}

func (x *InvoiceColumn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InvoiceColumn) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_services_billing_entities_v1alpha3_invoice_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha3_invoice_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x33, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x79,
	0x63, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x42, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x33, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x6f, 0x77, 0x52, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x55, 0x72, 0x6c, 0x22, 0xa0, 0x02, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x52, 0x6f, 0x77, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x4b, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12,
	0x37, 0x0a, 0x0c, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x0b, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x39, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0xb1, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x42, 0x0c,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x3b, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0xa2, 0x02,
	0x03, 0x53, 0x42, 0x45, 0xaa, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0xca, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0xe2, 0x02,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x25, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_billing_entities_v1alpha3_invoice_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha3_invoice_proto_rawDescData = file_services_billing_entities_v1alpha3_invoice_proto_rawDesc
)

func file_services_billing_entities_v1alpha3_invoice_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha3_invoice_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha3_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_billing_entities_v1alpha3_invoice_proto_rawDescData)
	})
	return file_services_billing_entities_v1alpha3_invoice_proto_rawDescData
}

var file_services_billing_entities_v1alpha3_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_billing_entities_v1alpha3_invoice_proto_goTypes = []any{
	(*Invoice)(nil),                // 0: services.billing.entities.v1alpha3.Invoice
	(*InvoiceRow)(nil),             // 1: services.billing.entities.v1alpha3.InvoiceRow
	(*InvoiceColumn)(nil),          // 2: services.billing.entities.v1alpha3.InvoiceColumn
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*decimal.Decimal)(nil),        // 5: google.type.Decimal
}
var file_services_billing_entities_v1alpha3_invoice_proto_depIdxs = []int32{
	3, // 0: services.billing.entities.v1alpha3.Invoice.create_time:type_name -> google.protobuf.Timestamp
	1, // 1: services.billing.entities.v1alpha3.Invoice.rows:type_name -> services.billing.entities.v1alpha3.InvoiceRow
	4, // 2: services.billing.entities.v1alpha3.Invoice.download_url:type_name -> google.protobuf.StringValue
	3, // 3: services.billing.entities.v1alpha3.InvoiceRow.date:type_name -> google.protobuf.Timestamp
	2, // 4: services.billing.entities.v1alpha3.InvoiceRow.columns:type_name -> services.billing.entities.v1alpha3.InvoiceColumn
	5, // 5: services.billing.entities.v1alpha3.InvoiceRow.rated_amount:type_name -> google.type.Decimal
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha3_invoice_proto_init() }
func file_services_billing_entities_v1alpha3_invoice_proto_init() {
	if File_services_billing_entities_v1alpha3_invoice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_billing_entities_v1alpha3_invoice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha3_invoice_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha3_invoice_proto_depIdxs,
		MessageInfos:      file_services_billing_entities_v1alpha3_invoice_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha3_invoice_proto = out.File
	file_services_billing_entities_v1alpha3_invoice_proto_rawDesc = nil
	file_services_billing_entities_v1alpha3_invoice_proto_goTypes = nil
	file_services_billing_entities_v1alpha3_invoice_proto_depIdxs = nil
}
