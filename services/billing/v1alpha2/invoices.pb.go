// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: services/billing/v1alpha2/invoices.proto

package billingv1alpha2

import (
	v1alpha2 "github.com/tcncloud/api-go/services/billing/entities/v1alpha2"
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
	return file_services_billing_v1alpha2_invoices_proto_enumTypes[0].Descriptor()
}

func (InvoiceFormat) Type() protoreflect.EnumType {
	return &file_services_billing_v1alpha2_invoices_proto_enumTypes[0]
}

func (x InvoiceFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvoiceFormat.Descriptor instead.
func (InvoiceFormat) EnumDescriptor() ([]byte, []int) {
	return file_services_billing_v1alpha2_invoices_proto_rawDescGZIP(), []int{0}
}

type ExportInvoiceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Format        InvoiceFormat          `protobuf:"varint,1,opt,name=format,proto3,enum=services.billing.v1alpha2.InvoiceFormat" json:"format,omitempty"`
	InvoiceDate   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=invoice_date,json=invoiceDate,proto3" json:"invoice_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportInvoiceRequest) Reset() {
	*x = ExportInvoiceRequest{}
	mi := &file_services_billing_v1alpha2_invoices_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportInvoiceRequest) ProtoMessage() {}

func (x *ExportInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha2_invoices_proto_msgTypes[0]
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
	return file_services_billing_v1alpha2_invoices_proto_rawDescGZIP(), []int{0}
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
	Invoice       *v1alpha2.Invoice      `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportInvoiceResponse) Reset() {
	*x = ExportInvoiceResponse{}
	mi := &file_services_billing_v1alpha2_invoices_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportInvoiceResponse) ProtoMessage() {}

func (x *ExportInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha2_invoices_proto_msgTypes[1]
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
	return file_services_billing_v1alpha2_invoices_proto_rawDescGZIP(), []int{1}
}

func (x *ExportInvoiceResponse) GetInvoice() *v1alpha2.Invoice {
	if x != nil {
		return x.Invoice
	}
	return nil
}

var File_services_billing_v1alpha2_invoices_proto protoreflect.FileDescriptor

var file_services_billing_v1alpha2_invoices_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x14, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x22, 0x5e, 0x0a, 0x15, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x2a, 0x47, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x46,
	0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x46,
	0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x43, 0x53, 0x56, 0x10, 0x01, 0x42, 0xfa, 0x01, 0x0a, 0x1d,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0d, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x3b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0xa2, 0x02, 0x03, 0x53, 0x42, 0x58, 0xaa, 0x02, 0x19, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x19, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0xe2, 0x02, 0x25, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_services_billing_v1alpha2_invoices_proto_rawDescOnce sync.Once
	file_services_billing_v1alpha2_invoices_proto_rawDescData []byte
)

func file_services_billing_v1alpha2_invoices_proto_rawDescGZIP() []byte {
	file_services_billing_v1alpha2_invoices_proto_rawDescOnce.Do(func() {
		file_services_billing_v1alpha2_invoices_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_v1alpha2_invoices_proto_rawDesc), len(file_services_billing_v1alpha2_invoices_proto_rawDesc)))
	})
	return file_services_billing_v1alpha2_invoices_proto_rawDescData
}

var file_services_billing_v1alpha2_invoices_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_billing_v1alpha2_invoices_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_billing_v1alpha2_invoices_proto_goTypes = []any{
	(InvoiceFormat)(0),            // 0: services.billing.v1alpha2.InvoiceFormat
	(*ExportInvoiceRequest)(nil),  // 1: services.billing.v1alpha2.ExportInvoiceRequest
	(*ExportInvoiceResponse)(nil), // 2: services.billing.v1alpha2.ExportInvoiceResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*v1alpha2.Invoice)(nil),      // 4: services.billing.entities.v1alpha2.Invoice
}
var file_services_billing_v1alpha2_invoices_proto_depIdxs = []int32{
	0, // 0: services.billing.v1alpha2.ExportInvoiceRequest.format:type_name -> services.billing.v1alpha2.InvoiceFormat
	3, // 1: services.billing.v1alpha2.ExportInvoiceRequest.invoice_date:type_name -> google.protobuf.Timestamp
	4, // 2: services.billing.v1alpha2.ExportInvoiceResponse.invoice:type_name -> services.billing.entities.v1alpha2.Invoice
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_services_billing_v1alpha2_invoices_proto_init() }
func file_services_billing_v1alpha2_invoices_proto_init() {
	if File_services_billing_v1alpha2_invoices_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_v1alpha2_invoices_proto_rawDesc), len(file_services_billing_v1alpha2_invoices_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_v1alpha2_invoices_proto_goTypes,
		DependencyIndexes: file_services_billing_v1alpha2_invoices_proto_depIdxs,
		EnumInfos:         file_services_billing_v1alpha2_invoices_proto_enumTypes,
		MessageInfos:      file_services_billing_v1alpha2_invoices_proto_msgTypes,
	}.Build()
	File_services_billing_v1alpha2_invoices_proto = out.File
	file_services_billing_v1alpha2_invoices_proto_goTypes = nil
	file_services_billing_v1alpha2_invoices_proto_depIdxs = nil
}
