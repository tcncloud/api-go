// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/billing/v1alpha1/invoices.proto

package billingv1alpha1

import (
	v1alpha1 "github.com/tcncloud/api-go/services/billing/entities/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type InvoiceFormat int32

const (
	InvoiceFormat_INVOICE_FORMAT_UNSPECIFIED InvoiceFormat = 0
	InvoiceFormat_INVOICE_FORMAT_CSV         InvoiceFormat = 1
)

// Enum value maps for InvoiceFormat.
var (
	InvoiceFormat_name = map[int32]string{
		0: "INVOICE_FORMAT_UNSPECIFIED",
		1: "INVOICE_FORMAT_CSV",
	}
	InvoiceFormat_value = map[string]int32{
		"INVOICE_FORMAT_UNSPECIFIED": 0,
		"INVOICE_FORMAT_CSV":         1,
	}
)

func (x InvoiceFormat) Enum() *InvoiceFormat {
	p := new(InvoiceFormat)
	*p = x
	return p
}

func (x InvoiceFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvoiceFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_services_billing_v1alpha1_invoices_proto_enumTypes[0].Descriptor()
}

func (InvoiceFormat) Type() protoreflect.EnumType {
	return &file_services_billing_v1alpha1_invoices_proto_enumTypes[0]
}

func (x InvoiceFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvoiceFormat.Descriptor instead.
func (InvoiceFormat) EnumDescriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
type CreateInvoiceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	InvoiceId string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	Invoice       *v1alpha1.Invoice `protobuf:"bytes,2,opt,name=invoice,proto3" json:"invoice,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateInvoiceRequest) Reset() {
	*x = CreateInvoiceRequest{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvoiceRequest) ProtoMessage() {}

func (x *CreateInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvoiceRequest.ProtoReflect.Descriptor instead.
func (*CreateInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *CreateInvoiceRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *CreateInvoiceRequest) GetInvoice() *v1alpha1.Invoice {
	if x != nil {
		return x.Invoice
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
type CreateInvoiceResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	InvoiceId     string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateInvoiceResponse) Reset() {
	*x = CreateInvoiceResponse{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvoiceResponse) ProtoMessage() {}

func (x *CreateInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvoiceResponse.ProtoReflect.Descriptor instead.
func (*CreateInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *CreateInvoiceResponse) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
type DeleteInvoiceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	InvoiceId     string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteInvoiceRequest) Reset() {
	*x = DeleteInvoiceRequest{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInvoiceRequest) ProtoMessage() {}

func (x *DeleteInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInvoiceRequest.ProtoReflect.Descriptor instead.
func (*DeleteInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{2}
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *DeleteInvoiceRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
type DeleteInvoiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteInvoiceResponse) Reset() {
	*x = DeleteInvoiceResponse{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInvoiceResponse) ProtoMessage() {}

func (x *DeleteInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInvoiceResponse.ProtoReflect.Descriptor instead.
func (*DeleteInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{3}
}

type ExportInvoiceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	InvoiceId     string                 `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	Format        InvoiceFormat          `protobuf:"varint,2,opt,name=format,proto3,enum=services.billing.v1alpha1.InvoiceFormat" json:"format,omitempty"`
	InvoiceDate   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=invoice_date,json=invoiceDate,proto3" json:"invoice_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportInvoiceRequest) Reset() {
	*x = ExportInvoiceRequest{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportInvoiceRequest) ProtoMessage() {}

func (x *ExportInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportInvoiceRequest.ProtoReflect.Descriptor instead.
func (*ExportInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{4}
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *ExportInvoiceRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *ExportInvoiceRequest) GetFormat() InvoiceFormat {
	if x != nil {
		return x.Format
	}
	return InvoiceFormat_INVOICE_FORMAT_UNSPECIFIED
}

func (x *ExportInvoiceRequest) GetInvoiceDate() *timestamppb.Timestamp {
	if x != nil {
		return x.InvoiceDate
	}
	return nil
}

type ExportInvoiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Invoice       *v1alpha1.Invoice      `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportInvoiceResponse) Reset() {
	*x = ExportInvoiceResponse{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportInvoiceResponse) ProtoMessage() {}

func (x *ExportInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportInvoiceResponse.ProtoReflect.Descriptor instead.
func (*ExportInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{5}
}

func (x *ExportInvoiceResponse) GetInvoice() *v1alpha1.Invoice {
	if x != nil {
		return x.Invoice
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
type GetInvoiceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	InvoiceId     string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInvoiceRequest) Reset() {
	*x = GetInvoiceRequest{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceRequest) ProtoMessage() {}

func (x *GetInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceRequest.ProtoReflect.Descriptor instead.
func (*GetInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{6}
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *GetInvoiceRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
type GetInvoiceResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	Invoice       *v1alpha1.Invoice `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInvoiceResponse) Reset() {
	*x = GetInvoiceResponse{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceResponse) ProtoMessage() {}

func (x *GetInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceResponse.ProtoReflect.Descriptor instead.
func (*GetInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{7}
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *GetInvoiceResponse) GetInvoice() *v1alpha1.Invoice {
	if x != nil {
		return x.Invoice
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
type ListInvoicesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	InvoiceId string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"` // Optional: defaults to no filter.
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	Fields *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=fields,proto3" json:"fields,omitempty"` // Optional: defaults to all fields.
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	Sort []*Sort `protobuf:"bytes,4,rep,name=sort,proto3" json:"sort,omitempty"` // Optional: defaults to no sort.
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	Page          *Page `protobuf:"bytes,5,opt,name=page,proto3" json:"page,omitempty"` // Optional: defaults to no paging.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListInvoicesRequest) Reset() {
	*x = ListInvoicesRequest{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInvoicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvoicesRequest) ProtoMessage() {}

func (x *ListInvoicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvoicesRequest.ProtoReflect.Descriptor instead.
func (*ListInvoicesRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{8}
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *ListInvoicesRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *ListInvoicesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *ListInvoicesRequest) GetFields() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Fields
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *ListInvoicesRequest) GetSort() []*Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *ListInvoicesRequest) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
type ListInvoicesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	Invoices []*v1alpha1.Invoice `protobuf:"bytes,1,rep,name=invoices,proto3" json:"invoices,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	Token         string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"` // Optional: only present if paginating.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListInvoicesResponse) Reset() {
	*x = ListInvoicesResponse{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInvoicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvoicesResponse) ProtoMessage() {}

func (x *ListInvoicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvoicesResponse.ProtoReflect.Descriptor instead.
func (*ListInvoicesResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{9}
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *ListInvoicesResponse) GetInvoices() []*v1alpha1.Invoice {
	if x != nil {
		return x.Invoices
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *ListInvoicesResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
type UpdateInvoiceRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	InvoiceId string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	Invoice *v1alpha1.Invoice `protobuf:"bytes,2,opt,name=invoice,proto3" json:"invoice,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
	UpdateFields  *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_fields,json=updateFields,proto3" json:"update_fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateInvoiceRequest) Reset() {
	*x = UpdateInvoiceRequest{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInvoiceRequest) ProtoMessage() {}

func (x *UpdateInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInvoiceRequest.ProtoReflect.Descriptor instead.
func (*UpdateInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{10}
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *UpdateInvoiceRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *UpdateInvoiceRequest) GetInvoice() *v1alpha1.Invoice {
	if x != nil {
		return x.Invoice
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
func (x *UpdateInvoiceRequest) GetUpdateFields() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateFields
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/v1alpha1/invoices.proto.
type UpdateInvoiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateInvoiceResponse) Reset() {
	*x = UpdateInvoiceResponse{}
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInvoiceResponse) ProtoMessage() {}

func (x *UpdateInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_invoices_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInvoiceResponse.ProtoReflect.Descriptor instead.
func (*UpdateInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_invoices_proto_rawDescGZIP(), []int{11}
}

var File_services_billing_v1alpha1_invoices_proto protoreflect.FileDescriptor

const file_services_billing_v1alpha1_invoices_proto_rawDesc = "" +
	"\n" +
	"(services/billing/v1alpha1/invoices.proto\x12\x19services.billing.v1alpha1\x1a google/protobuf/field_mask.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a0services/billing/entities/v1alpha1/invoice.proto\x1a$services/billing/v1alpha1/core.proto\"\x88\x01\n" +
	"\x14CreateInvoiceRequest\x12!\n" +
	"\n" +
	"invoice_id\x18\x01 \x01(\tB\x02\x18\x01R\tinvoiceId\x12I\n" +
	"\ainvoice\x18\x02 \x01(\v2+.services.billing.entities.v1alpha1.InvoiceB\x02\x18\x01R\ainvoice:\x02\x18\x01\">\n" +
	"\x15CreateInvoiceResponse\x12!\n" +
	"\n" +
	"invoice_id\x18\x01 \x01(\tB\x02\x18\x01R\tinvoiceId:\x02\x18\x01\"=\n" +
	"\x14DeleteInvoiceRequest\x12!\n" +
	"\n" +
	"invoice_id\x18\x01 \x01(\tB\x02\x18\x01R\tinvoiceId:\x02\x18\x01\"\x1b\n" +
	"\x15DeleteInvoiceResponse:\x02\x18\x01\"\xba\x01\n" +
	"\x14ExportInvoiceRequest\x12!\n" +
	"\n" +
	"invoice_id\x18\x01 \x01(\tB\x02\x18\x01R\tinvoiceId\x12@\n" +
	"\x06format\x18\x02 \x01(\x0e2(.services.billing.v1alpha1.InvoiceFormatR\x06format\x12=\n" +
	"\finvoice_date\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\vinvoiceDate\"^\n" +
	"\x15ExportInvoiceResponse\x12E\n" +
	"\ainvoice\x18\x01 \x01(\v2+.services.billing.entities.v1alpha1.InvoiceR\ainvoice\":\n" +
	"\x11GetInvoiceRequest\x12!\n" +
	"\n" +
	"invoice_id\x18\x01 \x01(\tB\x02\x18\x01R\tinvoiceId:\x02\x18\x01\"c\n" +
	"\x12GetInvoiceResponse\x12I\n" +
	"\ainvoice\x18\x01 \x01(\v2+.services.billing.entities.v1alpha1.InvoiceB\x02\x18\x01R\ainvoice:\x02\x18\x01\"\x82\x02\n" +
	"\x13ListInvoicesRequest\x12!\n" +
	"\n" +
	"invoice_id\x18\x01 \x01(\tB\x02\x18\x01R\tinvoiceId\x12\x1a\n" +
	"\x06filter\x18\x02 \x01(\tB\x02\x18\x01R\x06filter\x126\n" +
	"\x06fields\x18\x03 \x01(\v2\x1a.google.protobuf.FieldMaskB\x02\x18\x01R\x06fields\x127\n" +
	"\x04sort\x18\x04 \x03(\v2\x1f.services.billing.v1alpha1.SortB\x02\x18\x01R\x04sort\x127\n" +
	"\x04page\x18\x05 \x01(\v2\x1f.services.billing.v1alpha1.PageB\x02\x18\x01R\x04page:\x02\x18\x01\"\x81\x01\n" +
	"\x14ListInvoicesResponse\x12K\n" +
	"\binvoices\x18\x01 \x03(\v2+.services.billing.entities.v1alpha1.InvoiceB\x02\x18\x01R\binvoices\x12\x18\n" +
	"\x05token\x18\x02 \x01(\tB\x02\x18\x01R\x05token:\x02\x18\x01\"\xcd\x01\n" +
	"\x14UpdateInvoiceRequest\x12!\n" +
	"\n" +
	"invoice_id\x18\x01 \x01(\tB\x02\x18\x01R\tinvoiceId\x12I\n" +
	"\ainvoice\x18\x02 \x01(\v2+.services.billing.entities.v1alpha1.InvoiceB\x02\x18\x01R\ainvoice\x12C\n" +
	"\rupdate_fields\x18\x03 \x01(\v2\x1a.google.protobuf.FieldMaskB\x02\x18\x01R\fupdateFields:\x02\x18\x01\"\x1b\n" +
	"\x15UpdateInvoiceResponse:\x02\x18\x01*G\n" +
	"\rInvoiceFormat\x12\x1e\n" +
	"\x1aINVOICE_FORMAT_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12INVOICE_FORMAT_CSV\x10\x01B\xfa\x01\n" +
	"\x1dcom.services.billing.v1alpha1B\rInvoicesProtoP\x01ZDgithub.com/tcncloud/api-go/services/billing/v1alpha1;billingv1alpha1\xa2\x02\x03SBX\xaa\x02\x19Services.Billing.V1alpha1\xca\x02\x19Services\\Billing\\V1alpha1\xe2\x02%Services\\Billing\\V1alpha1\\GPBMetadata\xea\x02\x1bServices::Billing::V1alpha1b\x06proto3"

var (
	file_services_billing_v1alpha1_invoices_proto_rawDescOnce sync.Once
	file_services_billing_v1alpha1_invoices_proto_rawDescData []byte
)

func file_services_billing_v1alpha1_invoices_proto_rawDescGZIP() []byte {
	file_services_billing_v1alpha1_invoices_proto_rawDescOnce.Do(func() {
		file_services_billing_v1alpha1_invoices_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_v1alpha1_invoices_proto_rawDesc), len(file_services_billing_v1alpha1_invoices_proto_rawDesc)))
	})
	return file_services_billing_v1alpha1_invoices_proto_rawDescData
}

var file_services_billing_v1alpha1_invoices_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_billing_v1alpha1_invoices_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_services_billing_v1alpha1_invoices_proto_goTypes = []any{
	(InvoiceFormat)(0),            // 0: services.billing.v1alpha1.InvoiceFormat
	(*CreateInvoiceRequest)(nil),  // 1: services.billing.v1alpha1.CreateInvoiceRequest
	(*CreateInvoiceResponse)(nil), // 2: services.billing.v1alpha1.CreateInvoiceResponse
	(*DeleteInvoiceRequest)(nil),  // 3: services.billing.v1alpha1.DeleteInvoiceRequest
	(*DeleteInvoiceResponse)(nil), // 4: services.billing.v1alpha1.DeleteInvoiceResponse
	(*ExportInvoiceRequest)(nil),  // 5: services.billing.v1alpha1.ExportInvoiceRequest
	(*ExportInvoiceResponse)(nil), // 6: services.billing.v1alpha1.ExportInvoiceResponse
	(*GetInvoiceRequest)(nil),     // 7: services.billing.v1alpha1.GetInvoiceRequest
	(*GetInvoiceResponse)(nil),    // 8: services.billing.v1alpha1.GetInvoiceResponse
	(*ListInvoicesRequest)(nil),   // 9: services.billing.v1alpha1.ListInvoicesRequest
	(*ListInvoicesResponse)(nil),  // 10: services.billing.v1alpha1.ListInvoicesResponse
	(*UpdateInvoiceRequest)(nil),  // 11: services.billing.v1alpha1.UpdateInvoiceRequest
	(*UpdateInvoiceResponse)(nil), // 12: services.billing.v1alpha1.UpdateInvoiceResponse
	(*v1alpha1.Invoice)(nil),      // 13: services.billing.entities.v1alpha1.Invoice
	(*timestamppb.Timestamp)(nil), // 14: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 15: google.protobuf.FieldMask
	(*Sort)(nil),                  // 16: services.billing.v1alpha1.Sort
	(*Page)(nil),                  // 17: services.billing.v1alpha1.Page
}
var file_services_billing_v1alpha1_invoices_proto_depIdxs = []int32{
	13, // 0: services.billing.v1alpha1.CreateInvoiceRequest.invoice:type_name -> services.billing.entities.v1alpha1.Invoice
	0,  // 1: services.billing.v1alpha1.ExportInvoiceRequest.format:type_name -> services.billing.v1alpha1.InvoiceFormat
	14, // 2: services.billing.v1alpha1.ExportInvoiceRequest.invoice_date:type_name -> google.protobuf.Timestamp
	13, // 3: services.billing.v1alpha1.ExportInvoiceResponse.invoice:type_name -> services.billing.entities.v1alpha1.Invoice
	13, // 4: services.billing.v1alpha1.GetInvoiceResponse.invoice:type_name -> services.billing.entities.v1alpha1.Invoice
	15, // 5: services.billing.v1alpha1.ListInvoicesRequest.fields:type_name -> google.protobuf.FieldMask
	16, // 6: services.billing.v1alpha1.ListInvoicesRequest.sort:type_name -> services.billing.v1alpha1.Sort
	17, // 7: services.billing.v1alpha1.ListInvoicesRequest.page:type_name -> services.billing.v1alpha1.Page
	13, // 8: services.billing.v1alpha1.ListInvoicesResponse.invoices:type_name -> services.billing.entities.v1alpha1.Invoice
	13, // 9: services.billing.v1alpha1.UpdateInvoiceRequest.invoice:type_name -> services.billing.entities.v1alpha1.Invoice
	15, // 10: services.billing.v1alpha1.UpdateInvoiceRequest.update_fields:type_name -> google.protobuf.FieldMask
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_services_billing_v1alpha1_invoices_proto_init() }
func file_services_billing_v1alpha1_invoices_proto_init() {
	if File_services_billing_v1alpha1_invoices_proto != nil {
		return
	}
	file_services_billing_v1alpha1_core_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_v1alpha1_invoices_proto_rawDesc), len(file_services_billing_v1alpha1_invoices_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_v1alpha1_invoices_proto_goTypes,
		DependencyIndexes: file_services_billing_v1alpha1_invoices_proto_depIdxs,
		EnumInfos:         file_services_billing_v1alpha1_invoices_proto_enumTypes,
		MessageInfos:      file_services_billing_v1alpha1_invoices_proto_msgTypes,
	}.Build()
	File_services_billing_v1alpha1_invoices_proto = out.File
	file_services_billing_v1alpha1_invoices_proto_goTypes = nil
	file_services_billing_v1alpha1_invoices_proto_depIdxs = nil
}
