// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha4/invoice.proto

package entitiesv1alpha4

import (
	_ "google.golang.org/genproto/googleapis/type/decimal"
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
	mi := &file_services_billing_entities_v1alpha4_invoice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_invoice_proto_msgTypes[0]
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
	return file_services_billing_entities_v1alpha4_invoice_proto_rawDescGZIP(), []int{0}
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
	Columns       []*InvoiceColumn       `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InvoiceRow) Reset() {
	*x = InvoiceRow{}
	mi := &file_services_billing_entities_v1alpha4_invoice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvoiceRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceRow) ProtoMessage() {}

func (x *InvoiceRow) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_invoice_proto_msgTypes[1]
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
	return file_services_billing_entities_v1alpha4_invoice_proto_rawDescGZIP(), []int{1}
}

func (x *InvoiceRow) GetColumns() []*InvoiceColumn {
	if x != nil {
		return x.Columns
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
	mi := &file_services_billing_entities_v1alpha4_invoice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvoiceColumn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceColumn) ProtoMessage() {}

func (x *InvoiceColumn) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_invoice_proto_msgTypes[2]
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
	return file_services_billing_entities_v1alpha4_invoice_proto_rawDescGZIP(), []int{2}
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

var File_services_billing_entities_v1alpha4_invoice_proto protoreflect.FileDescriptor

const file_services_billing_entities_v1alpha4_invoice_proto_rawDesc = "" +
	"\n" +
	"0services/billing/entities/v1alpha4/invoice.proto\x12\"services.billing.entities.v1alpha4\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x19google/type/decimal.proto\"\xf0\x01\n" +
	"\aInvoice\x12#\n" +
	"\rbilling_cycle\x18\x01 \x01(\tR\fbillingCycle\x12;\n" +
	"\vcreate_time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x12B\n" +
	"\x04rows\x18\x03 \x03(\v2..services.billing.entities.v1alpha4.InvoiceRowR\x04rows\x12?\n" +
	"\fdownload_url\x18\x04 \x01(\v2\x1c.google.protobuf.StringValueR\vdownloadUrl\"Y\n" +
	"\n" +
	"InvoiceRow\x12K\n" +
	"\acolumns\x18\x01 \x03(\v21.services.billing.entities.v1alpha4.InvoiceColumnR\acolumns\"9\n" +
	"\rInvoiceColumn\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05valueB\xb1\x02\n" +
	"&com.services.billing.entities.v1alpha4B\fInvoiceProtoP\x01ZNgithub.com/tcncloud/api-go/services/billing/entities/v1alpha4;entitiesv1alpha4\xa2\x02\x03SBE\xaa\x02\"Services.Billing.Entities.V1alpha4\xca\x02\"Services\\Billing\\Entities\\V1alpha4\xe2\x02.Services\\Billing\\Entities\\V1alpha4\\GPBMetadata\xea\x02%Services::Billing::Entities::V1alpha4b\x06proto3"

var (
	file_services_billing_entities_v1alpha4_invoice_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha4_invoice_proto_rawDescData []byte
)

func file_services_billing_entities_v1alpha4_invoice_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha4_invoice_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha4_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha4_invoice_proto_rawDesc), len(file_services_billing_entities_v1alpha4_invoice_proto_rawDesc)))
	})
	return file_services_billing_entities_v1alpha4_invoice_proto_rawDescData
}

var file_services_billing_entities_v1alpha4_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_billing_entities_v1alpha4_invoice_proto_goTypes = []any{
	(*Invoice)(nil),                // 0: services.billing.entities.v1alpha4.Invoice
	(*InvoiceRow)(nil),             // 1: services.billing.entities.v1alpha4.InvoiceRow
	(*InvoiceColumn)(nil),          // 2: services.billing.entities.v1alpha4.InvoiceColumn
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
}
var file_services_billing_entities_v1alpha4_invoice_proto_depIdxs = []int32{
	3, // 0: services.billing.entities.v1alpha4.Invoice.create_time:type_name -> google.protobuf.Timestamp
	1, // 1: services.billing.entities.v1alpha4.Invoice.rows:type_name -> services.billing.entities.v1alpha4.InvoiceRow
	4, // 2: services.billing.entities.v1alpha4.Invoice.download_url:type_name -> google.protobuf.StringValue
	2, // 3: services.billing.entities.v1alpha4.InvoiceRow.columns:type_name -> services.billing.entities.v1alpha4.InvoiceColumn
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha4_invoice_proto_init() }
func file_services_billing_entities_v1alpha4_invoice_proto_init() {
	if File_services_billing_entities_v1alpha4_invoice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha4_invoice_proto_rawDesc), len(file_services_billing_entities_v1alpha4_invoice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha4_invoice_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha4_invoice_proto_depIdxs,
		MessageInfos:      file_services_billing_entities_v1alpha4_invoice_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha4_invoice_proto = out.File
	file_services_billing_entities_v1alpha4_invoice_proto_goTypes = nil
	file_services_billing_entities_v1alpha4_invoice_proto_depIdxs = nil
}
