// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/billing/entities.proto

package billing

import (
	_ "github.com/tcncloud/api-go/annotations"
	billing "github.com/tcncloud/api-go/api/commons/billing"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
type InvoiceFormat int32

const (
	InvoiceFormat_INVOICE_FORMAT_UNSPECIFIED InvoiceFormat = 0
	InvoiceFormat_INVOICE_FORMAT_PROTO       InvoiceFormat = 1
	InvoiceFormat_INVOICE_FORMAT_CSV         InvoiceFormat = 2
)

// Enum value maps for InvoiceFormat.
var (
	InvoiceFormat_name = map[int32]string{
		0: "INVOICE_FORMAT_UNSPECIFIED",
		1: "INVOICE_FORMAT_PROTO",
		2: "INVOICE_FORMAT_CSV",
	}
	InvoiceFormat_value = map[string]int32{
		"INVOICE_FORMAT_UNSPECIFIED": 0,
		"INVOICE_FORMAT_PROTO":       1,
		"INVOICE_FORMAT_CSV":         2,
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
	return file_api_v1alpha1_billing_entities_proto_enumTypes[0].Descriptor()
}

func (InvoiceFormat) Type() protoreflect.EnumType {
	return &file_api_v1alpha1_billing_entities_proto_enumTypes[0]
}

func (x InvoiceFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvoiceFormat.Descriptor instead.
func (InvoiceFormat) EnumDescriptor() ([]byte, []int) {
	return file_api_v1alpha1_billing_entities_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
type GetBillingPlanReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	OrgId         string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBillingPlanReq) Reset() {
	*x = GetBillingPlanReq{}
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBillingPlanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBillingPlanReq) ProtoMessage() {}

func (x *GetBillingPlanReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBillingPlanReq.ProtoReflect.Descriptor instead.
func (*GetBillingPlanReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_billing_entities_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetBillingPlanReq) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
type GetBillingPlanRes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	BillingPlan   *billing.Plan `protobuf:"bytes,1,opt,name=billing_plan,json=billingPlan,proto3" json:"billing_plan,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBillingPlanRes) Reset() {
	*x = GetBillingPlanRes{}
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBillingPlanRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBillingPlanRes) ProtoMessage() {}

func (x *GetBillingPlanRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBillingPlanRes.ProtoReflect.Descriptor instead.
func (*GetBillingPlanRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_billing_entities_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetBillingPlanRes) GetBillingPlan() *billing.Plan {
	if x != nil {
		return x.BillingPlan
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
type UpdateBillingPlanReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	BillingDetails []*billing.Detail `protobuf:"bytes,1,rep,name=billing_details,json=billingDetails,proto3" json:"billing_details,omitempty"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	OrgId         string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBillingPlanReq) Reset() {
	*x = UpdateBillingPlanReq{}
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBillingPlanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBillingPlanReq) ProtoMessage() {}

func (x *UpdateBillingPlanReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBillingPlanReq.ProtoReflect.Descriptor instead.
func (*UpdateBillingPlanReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_billing_entities_proto_rawDescGZIP(), []int{2}
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *UpdateBillingPlanReq) GetBillingDetails() []*billing.Detail {
	if x != nil {
		return x.BillingDetails
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *UpdateBillingPlanReq) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
type UpdateBillingPlanRes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	BillingPlan   *billing.Plan `protobuf:"bytes,1,opt,name=billing_plan,json=billingPlan,proto3" json:"billing_plan,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBillingPlanRes) Reset() {
	*x = UpdateBillingPlanRes{}
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBillingPlanRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBillingPlanRes) ProtoMessage() {}

func (x *UpdateBillingPlanRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBillingPlanRes.ProtoReflect.Descriptor instead.
func (*UpdateBillingPlanRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_billing_entities_proto_rawDescGZIP(), []int{3}
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *UpdateBillingPlanRes) GetBillingPlan() *billing.Plan {
	if x != nil {
		return x.BillingPlan
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
type GetInvoiceReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	InvoiceDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=invoice_date,json=invoiceDate,proto3" json:"invoice_date,omitempty"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	Format InvoiceFormat `protobuf:"varint,3,opt,name=format,proto3,enum=api.v1alpha1.billing.InvoiceFormat" json:"format,omitempty"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	InvoiceFormat billing.InvoiceFormat `protobuf:"varint,4,opt,name=invoice_format,json=invoiceFormat,proto3,enum=api.commons.billing.InvoiceFormat" json:"invoice_format,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInvoiceReq) Reset() {
	*x = GetInvoiceReq{}
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInvoiceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceReq) ProtoMessage() {}

func (x *GetInvoiceReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceReq.ProtoReflect.Descriptor instead.
func (*GetInvoiceReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_billing_entities_proto_rawDescGZIP(), []int{4}
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetInvoiceReq) GetInvoiceDate() *timestamppb.Timestamp {
	if x != nil {
		return x.InvoiceDate
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetInvoiceReq) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetInvoiceReq) GetFormat() InvoiceFormat {
	if x != nil {
		return x.Format
	}
	return InvoiceFormat_INVOICE_FORMAT_UNSPECIFIED
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetInvoiceReq) GetInvoiceFormat() billing.InvoiceFormat {
	if x != nil {
		return x.InvoiceFormat
	}
	return billing.InvoiceFormat(0)
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
type GetInvoiceRes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	Invoice *billing.Invoice `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"`
	// Types that are valid to be assigned to Format:
	//
	//	*GetInvoiceRes_Proto
	//	*GetInvoiceRes_CsvUrl
	Format isGetInvoiceRes_Format `protobuf_oneof:"format"`
	// Types that are valid to be assigned to InvoiceData:
	//
	//	*GetInvoiceRes_InvoiceProto
	//	*GetInvoiceRes_InvoiceCsvUrl
	InvoiceData isGetInvoiceRes_InvoiceData `protobuf_oneof:"invoice_data"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	BillingCycle  string `protobuf:"bytes,4,opt,name=billing_cycle,json=billingCycle,proto3" json:"billing_cycle,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInvoiceRes) Reset() {
	*x = GetInvoiceRes{}
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInvoiceRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceRes) ProtoMessage() {}

func (x *GetInvoiceRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceRes.ProtoReflect.Descriptor instead.
func (*GetInvoiceRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_billing_entities_proto_rawDescGZIP(), []int{5}
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetInvoiceRes) GetInvoice() *billing.Invoice {
	if x != nil {
		return x.Invoice
	}
	return nil
}

func (x *GetInvoiceRes) GetFormat() isGetInvoiceRes_Format {
	if x != nil {
		return x.Format
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetInvoiceRes) GetProto() *billing.Invoice {
	if x != nil {
		if x, ok := x.Format.(*GetInvoiceRes_Proto); ok {
			return x.Proto
		}
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetInvoiceRes) GetCsvUrl() string {
	if x != nil {
		if x, ok := x.Format.(*GetInvoiceRes_CsvUrl); ok {
			return x.CsvUrl
		}
	}
	return ""
}

func (x *GetInvoiceRes) GetInvoiceData() isGetInvoiceRes_InvoiceData {
	if x != nil {
		return x.InvoiceData
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetInvoiceRes) GetInvoiceProto() *billing.Invoice {
	if x != nil {
		if x, ok := x.InvoiceData.(*GetInvoiceRes_InvoiceProto); ok {
			return x.InvoiceProto
		}
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetInvoiceRes) GetInvoiceCsvUrl() string {
	if x != nil {
		if x, ok := x.InvoiceData.(*GetInvoiceRes_InvoiceCsvUrl); ok {
			return x.InvoiceCsvUrl
		}
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *GetInvoiceRes) GetBillingCycle() string {
	if x != nil {
		return x.BillingCycle
	}
	return ""
}

type isGetInvoiceRes_Format interface {
	isGetInvoiceRes_Format()
}

type GetInvoiceRes_Proto struct {
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	Proto *billing.Invoice `protobuf:"bytes,2,opt,name=proto,proto3,oneof"`
}

type GetInvoiceRes_CsvUrl struct {
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	CsvUrl string `protobuf:"bytes,3,opt,name=csv_url,json=csvUrl,proto3,oneof"`
}

func (*GetInvoiceRes_Proto) isGetInvoiceRes_Format() {}

func (*GetInvoiceRes_CsvUrl) isGetInvoiceRes_Format() {}

type isGetInvoiceRes_InvoiceData interface {
	isGetInvoiceRes_InvoiceData()
}

type GetInvoiceRes_InvoiceProto struct {
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	InvoiceProto *billing.Invoice `protobuf:"bytes,100,opt,name=invoice_proto,json=invoiceProto,proto3,oneof"`
}

type GetInvoiceRes_InvoiceCsvUrl struct {
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	InvoiceCsvUrl string `protobuf:"bytes,101,opt,name=invoice_csv_url,json=invoiceCsvUrl,proto3,oneof"`
}

func (*GetInvoiceRes_InvoiceProto) isGetInvoiceRes_InvoiceData() {}

func (*GetInvoiceRes_InvoiceCsvUrl) isGetInvoiceRes_InvoiceData() {}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
type ExportGeneratedInvoiceReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	InvoiceDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=invoice_date,json=invoiceDate,proto3" json:"invoice_date,omitempty"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	Format InvoiceFormat `protobuf:"varint,3,opt,name=format,proto3,enum=api.v1alpha1.billing.InvoiceFormat" json:"format,omitempty"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	InvoiceFormat billing.InvoiceFormat `protobuf:"varint,4,opt,name=invoice_format,json=invoiceFormat,proto3,enum=api.commons.billing.InvoiceFormat" json:"invoice_format,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportGeneratedInvoiceReq) Reset() {
	*x = ExportGeneratedInvoiceReq{}
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportGeneratedInvoiceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportGeneratedInvoiceReq) ProtoMessage() {}

func (x *ExportGeneratedInvoiceReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportGeneratedInvoiceReq.ProtoReflect.Descriptor instead.
func (*ExportGeneratedInvoiceReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_billing_entities_proto_rawDescGZIP(), []int{6}
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *ExportGeneratedInvoiceReq) GetInvoiceDate() *timestamppb.Timestamp {
	if x != nil {
		return x.InvoiceDate
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *ExportGeneratedInvoiceReq) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *ExportGeneratedInvoiceReq) GetFormat() InvoiceFormat {
	if x != nil {
		return x.Format
	}
	return InvoiceFormat_INVOICE_FORMAT_UNSPECIFIED
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *ExportGeneratedInvoiceReq) GetInvoiceFormat() billing.InvoiceFormat {
	if x != nil {
		return x.InvoiceFormat
	}
	return billing.InvoiceFormat(0)
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
type ExportGeneratedInvoiceRes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Format:
	//
	//	*ExportGeneratedInvoiceRes_Proto
	//	*ExportGeneratedInvoiceRes_CsvUrl
	Format isExportGeneratedInvoiceRes_Format `protobuf_oneof:"format"`
	// Types that are valid to be assigned to InvoiceData:
	//
	//	*ExportGeneratedInvoiceRes_InvoiceProto
	//	*ExportGeneratedInvoiceRes_InvoiceCsvUrl
	InvoiceData isExportGeneratedInvoiceRes_InvoiceData `protobuf_oneof:"invoice_data"`
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	BillingCycle  string `protobuf:"bytes,4,opt,name=billing_cycle,json=billingCycle,proto3" json:"billing_cycle,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportGeneratedInvoiceRes) Reset() {
	*x = ExportGeneratedInvoiceRes{}
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportGeneratedInvoiceRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportGeneratedInvoiceRes) ProtoMessage() {}

func (x *ExportGeneratedInvoiceRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_billing_entities_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportGeneratedInvoiceRes.ProtoReflect.Descriptor instead.
func (*ExportGeneratedInvoiceRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_billing_entities_proto_rawDescGZIP(), []int{7}
}

func (x *ExportGeneratedInvoiceRes) GetFormat() isExportGeneratedInvoiceRes_Format {
	if x != nil {
		return x.Format
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *ExportGeneratedInvoiceRes) GetProto() *billing.Invoice {
	if x != nil {
		if x, ok := x.Format.(*ExportGeneratedInvoiceRes_Proto); ok {
			return x.Proto
		}
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *ExportGeneratedInvoiceRes) GetCsvUrl() string {
	if x != nil {
		if x, ok := x.Format.(*ExportGeneratedInvoiceRes_CsvUrl); ok {
			return x.CsvUrl
		}
	}
	return ""
}

func (x *ExportGeneratedInvoiceRes) GetInvoiceData() isExportGeneratedInvoiceRes_InvoiceData {
	if x != nil {
		return x.InvoiceData
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *ExportGeneratedInvoiceRes) GetInvoiceProto() *billing.Invoice {
	if x != nil {
		if x, ok := x.InvoiceData.(*ExportGeneratedInvoiceRes_InvoiceProto); ok {
			return x.InvoiceProto
		}
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *ExportGeneratedInvoiceRes) GetInvoiceCsvUrl() string {
	if x != nil {
		if x, ok := x.InvoiceData.(*ExportGeneratedInvoiceRes_InvoiceCsvUrl); ok {
			return x.InvoiceCsvUrl
		}
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
func (x *ExportGeneratedInvoiceRes) GetBillingCycle() string {
	if x != nil {
		return x.BillingCycle
	}
	return ""
}

type isExportGeneratedInvoiceRes_Format interface {
	isExportGeneratedInvoiceRes_Format()
}

type ExportGeneratedInvoiceRes_Proto struct {
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	Proto *billing.Invoice `protobuf:"bytes,1,opt,name=proto,proto3,oneof"`
}

type ExportGeneratedInvoiceRes_CsvUrl struct {
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	CsvUrl string `protobuf:"bytes,2,opt,name=csv_url,json=csvUrl,proto3,oneof"`
}

func (*ExportGeneratedInvoiceRes_Proto) isExportGeneratedInvoiceRes_Format() {}

func (*ExportGeneratedInvoiceRes_CsvUrl) isExportGeneratedInvoiceRes_Format() {}

type isExportGeneratedInvoiceRes_InvoiceData interface {
	isExportGeneratedInvoiceRes_InvoiceData()
}

type ExportGeneratedInvoiceRes_InvoiceProto struct {
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	InvoiceProto *billing.Invoice `protobuf:"bytes,100,opt,name=invoice_proto,json=invoiceProto,proto3,oneof"`
}

type ExportGeneratedInvoiceRes_InvoiceCsvUrl struct {
	// Deprecated: Marked as deprecated in api/v1alpha1/billing/entities.proto.
	InvoiceCsvUrl string `protobuf:"bytes,101,opt,name=invoice_csv_url,json=invoiceCsvUrl,proto3,oneof"`
}

func (*ExportGeneratedInvoiceRes_InvoiceProto) isExportGeneratedInvoiceRes_InvoiceData() {}

func (*ExportGeneratedInvoiceRes_InvoiceCsvUrl) isExportGeneratedInvoiceRes_InvoiceData() {}

var File_api_v1alpha1_billing_entities_proto protoreflect.FileDescriptor

const file_api_v1alpha1_billing_entities_proto_rawDesc = "" +
	"\n" +
	"#api/v1alpha1/billing/entities.proto\x12\x14api.v1alpha1.billing\x1a\x17annotations/authz.proto\x1a api/commons/billing/detail.proto\x1a!api/commons/billing/invoice.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"2\n" +
	"\x11GetBillingPlanReq\x12\x19\n" +
	"\x06org_id\x18\x01 \x01(\tB\x02\x18\x01R\x05orgId:\x02\x18\x01\"Y\n" +
	"\x11GetBillingPlanRes\x12@\n" +
	"\fbilling_plan\x18\x01 \x01(\v2\x19.api.commons.billing.PlanB\x02\x18\x01R\vbillingPlan:\x02\x18\x01\"\x7f\n" +
	"\x14UpdateBillingPlanReq\x12H\n" +
	"\x0fbilling_details\x18\x01 \x03(\v2\x1b.api.commons.billing.DetailB\x02\x18\x01R\x0ebillingDetails\x12\x19\n" +
	"\x06org_id\x18\x02 \x01(\tB\x02\x18\x01R\x05orgId:\x02\x18\x01\"\\\n" +
	"\x14UpdateBillingPlanRes\x12@\n" +
	"\fbilling_plan\x18\x01 \x01(\v2\x19.api.commons.billing.PlanB\x02\x18\x01R\vbillingPlan:\x02\x18\x01\"\x81\x02\n" +
	"\rGetInvoiceReq\x12A\n" +
	"\finvoice_date\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampB\x02\x18\x01R\vinvoiceDate\x12\x19\n" +
	"\x06org_id\x18\x02 \x01(\tB\x02\x18\x01R\x05orgId\x12?\n" +
	"\x06format\x18\x03 \x01(\x0e2#.api.v1alpha1.billing.InvoiceFormatB\x02\x18\x01R\x06format\x12M\n" +
	"\x0einvoice_format\x18\x04 \x01(\x0e2\".api.commons.billing.InvoiceFormatB\x02\x18\x01R\rinvoiceFormat:\x02\x18\x01\"\xe2\x02\n" +
	"\rGetInvoiceRes\x12:\n" +
	"\ainvoice\x18\x01 \x01(\v2\x1c.api.commons.billing.InvoiceB\x02\x18\x01R\ainvoice\x128\n" +
	"\x05proto\x18\x02 \x01(\v2\x1c.api.commons.billing.InvoiceB\x02\x18\x01H\x00R\x05proto\x12\x1d\n" +
	"\acsv_url\x18\x03 \x01(\tB\x02\x18\x01H\x00R\x06csvUrl\x12G\n" +
	"\rinvoice_proto\x18d \x01(\v2\x1c.api.commons.billing.InvoiceB\x02\x18\x01H\x01R\finvoiceProto\x12,\n" +
	"\x0finvoice_csv_url\x18e \x01(\tB\x02\x18\x01H\x01R\rinvoiceCsvUrl\x12'\n" +
	"\rbilling_cycle\x18\x04 \x01(\tB\x02\x18\x01R\fbillingCycle:\x02\x18\x01B\b\n" +
	"\x06formatB\x0e\n" +
	"\finvoice_data\"\x8d\x02\n" +
	"\x19ExportGeneratedInvoiceReq\x12A\n" +
	"\finvoice_date\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampB\x02\x18\x01R\vinvoiceDate\x12\x19\n" +
	"\x06org_id\x18\x02 \x01(\tB\x02\x18\x01R\x05orgId\x12?\n" +
	"\x06format\x18\x03 \x01(\x0e2#.api.v1alpha1.billing.InvoiceFormatB\x02\x18\x01R\x06format\x12M\n" +
	"\x0einvoice_format\x18\x04 \x01(\x0e2\".api.commons.billing.InvoiceFormatB\x02\x18\x01R\rinvoiceFormat:\x02\x18\x01\"\xb2\x02\n" +
	"\x19ExportGeneratedInvoiceRes\x128\n" +
	"\x05proto\x18\x01 \x01(\v2\x1c.api.commons.billing.InvoiceB\x02\x18\x01H\x00R\x05proto\x12\x1d\n" +
	"\acsv_url\x18\x02 \x01(\tB\x02\x18\x01H\x00R\x06csvUrl\x12G\n" +
	"\rinvoice_proto\x18d \x01(\v2\x1c.api.commons.billing.InvoiceB\x02\x18\x01H\x01R\finvoiceProto\x12,\n" +
	"\x0finvoice_csv_url\x18e \x01(\tB\x02\x18\x01H\x01R\rinvoiceCsvUrl\x12'\n" +
	"\rbilling_cycle\x18\x04 \x01(\tB\x02\x18\x01R\fbillingCycle:\x02\x18\x01B\b\n" +
	"\x06formatB\x0e\n" +
	"\finvoice_data*e\n" +
	"\rInvoiceFormat\x12\x1e\n" +
	"\x1aINVOICE_FORMAT_UNSPECIFIED\x10\x00\x12\x18\n" +
	"\x14INVOICE_FORMAT_PROTO\x10\x01\x12\x16\n" +
	"\x12INVOICE_FORMAT_CSV\x10\x02\x1a\x02\x18\x01B\xcc\x01\n" +
	"\x18com.api.v1alpha1.billingB\rEntitiesProtoP\x01Z/github.com/tcncloud/api-go/api/v1alpha1/billing\xa2\x02\x03AVB\xaa\x02\x14Api.V1alpha1.Billing\xca\x02\x14Api\\V1alpha1\\Billing\xe2\x02 Api\\V1alpha1\\Billing\\GPBMetadata\xea\x02\x16Api::V1alpha1::Billingb\x06proto3"

var (
	file_api_v1alpha1_billing_entities_proto_rawDescOnce sync.Once
	file_api_v1alpha1_billing_entities_proto_rawDescData []byte
)

func file_api_v1alpha1_billing_entities_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_billing_entities_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_billing_entities_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_billing_entities_proto_rawDesc), len(file_api_v1alpha1_billing_entities_proto_rawDesc)))
	})
	return file_api_v1alpha1_billing_entities_proto_rawDescData
}

var file_api_v1alpha1_billing_entities_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1alpha1_billing_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_v1alpha1_billing_entities_proto_goTypes = []any{
	(InvoiceFormat)(0),                // 0: api.v1alpha1.billing.InvoiceFormat
	(*GetBillingPlanReq)(nil),         // 1: api.v1alpha1.billing.GetBillingPlanReq
	(*GetBillingPlanRes)(nil),         // 2: api.v1alpha1.billing.GetBillingPlanRes
	(*UpdateBillingPlanReq)(nil),      // 3: api.v1alpha1.billing.UpdateBillingPlanReq
	(*UpdateBillingPlanRes)(nil),      // 4: api.v1alpha1.billing.UpdateBillingPlanRes
	(*GetInvoiceReq)(nil),             // 5: api.v1alpha1.billing.GetInvoiceReq
	(*GetInvoiceRes)(nil),             // 6: api.v1alpha1.billing.GetInvoiceRes
	(*ExportGeneratedInvoiceReq)(nil), // 7: api.v1alpha1.billing.ExportGeneratedInvoiceReq
	(*ExportGeneratedInvoiceRes)(nil), // 8: api.v1alpha1.billing.ExportGeneratedInvoiceRes
	(*billing.Plan)(nil),              // 9: api.commons.billing.Plan
	(*billing.Detail)(nil),            // 10: api.commons.billing.Detail
	(*timestamppb.Timestamp)(nil),     // 11: google.protobuf.Timestamp
	(billing.InvoiceFormat)(0),        // 12: api.commons.billing.InvoiceFormat
	(*billing.Invoice)(nil),           // 13: api.commons.billing.Invoice
}
var file_api_v1alpha1_billing_entities_proto_depIdxs = []int32{
	9,  // 0: api.v1alpha1.billing.GetBillingPlanRes.billing_plan:type_name -> api.commons.billing.Plan
	10, // 1: api.v1alpha1.billing.UpdateBillingPlanReq.billing_details:type_name -> api.commons.billing.Detail
	9,  // 2: api.v1alpha1.billing.UpdateBillingPlanRes.billing_plan:type_name -> api.commons.billing.Plan
	11, // 3: api.v1alpha1.billing.GetInvoiceReq.invoice_date:type_name -> google.protobuf.Timestamp
	0,  // 4: api.v1alpha1.billing.GetInvoiceReq.format:type_name -> api.v1alpha1.billing.InvoiceFormat
	12, // 5: api.v1alpha1.billing.GetInvoiceReq.invoice_format:type_name -> api.commons.billing.InvoiceFormat
	13, // 6: api.v1alpha1.billing.GetInvoiceRes.invoice:type_name -> api.commons.billing.Invoice
	13, // 7: api.v1alpha1.billing.GetInvoiceRes.proto:type_name -> api.commons.billing.Invoice
	13, // 8: api.v1alpha1.billing.GetInvoiceRes.invoice_proto:type_name -> api.commons.billing.Invoice
	11, // 9: api.v1alpha1.billing.ExportGeneratedInvoiceReq.invoice_date:type_name -> google.protobuf.Timestamp
	0,  // 10: api.v1alpha1.billing.ExportGeneratedInvoiceReq.format:type_name -> api.v1alpha1.billing.InvoiceFormat
	12, // 11: api.v1alpha1.billing.ExportGeneratedInvoiceReq.invoice_format:type_name -> api.commons.billing.InvoiceFormat
	13, // 12: api.v1alpha1.billing.ExportGeneratedInvoiceRes.proto:type_name -> api.commons.billing.Invoice
	13, // 13: api.v1alpha1.billing.ExportGeneratedInvoiceRes.invoice_proto:type_name -> api.commons.billing.Invoice
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_billing_entities_proto_init() }
func file_api_v1alpha1_billing_entities_proto_init() {
	if File_api_v1alpha1_billing_entities_proto != nil {
		return
	}
	file_api_v1alpha1_billing_entities_proto_msgTypes[5].OneofWrappers = []any{
		(*GetInvoiceRes_Proto)(nil),
		(*GetInvoiceRes_CsvUrl)(nil),
		(*GetInvoiceRes_InvoiceProto)(nil),
		(*GetInvoiceRes_InvoiceCsvUrl)(nil),
	}
	file_api_v1alpha1_billing_entities_proto_msgTypes[7].OneofWrappers = []any{
		(*ExportGeneratedInvoiceRes_Proto)(nil),
		(*ExportGeneratedInvoiceRes_CsvUrl)(nil),
		(*ExportGeneratedInvoiceRes_InvoiceProto)(nil),
		(*ExportGeneratedInvoiceRes_InvoiceCsvUrl)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_billing_entities_proto_rawDesc), len(file_api_v1alpha1_billing_entities_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_billing_entities_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_billing_entities_proto_depIdxs,
		EnumInfos:         file_api_v1alpha1_billing_entities_proto_enumTypes,
		MessageInfos:      file_api_v1alpha1_billing_entities_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_billing_entities_proto = out.File
	file_api_v1alpha1_billing_entities_proto_goTypes = nil
	file_api_v1alpha1_billing_entities_proto_depIdxs = nil
}
