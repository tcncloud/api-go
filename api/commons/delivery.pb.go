// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/commons/delivery.proto

package commons

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

type TransferStatus int32

const (
	TransferStatus_TRANSFER_STATUS_WAITING              TransferStatus = 0
	TransferStatus_TRANSFER_STATUS_RUNNING              TransferStatus = 1
	TransferStatus_TRANSFER_STATUS_DONE_SUCCESS         TransferStatus = 2
	TransferStatus_TRANSFER_STATUS_DONE_PARTIAL_FAILURE TransferStatus = 3
	TransferStatus_TRANSFER_STATUS_DONE_TOTAL_FAILURE   TransferStatus = 4
)

// Enum value maps for TransferStatus.
var (
	TransferStatus_name = map[int32]string{
		0: "TRANSFER_STATUS_WAITING",
		1: "TRANSFER_STATUS_RUNNING",
		2: "TRANSFER_STATUS_DONE_SUCCESS",
		3: "TRANSFER_STATUS_DONE_PARTIAL_FAILURE",
		4: "TRANSFER_STATUS_DONE_TOTAL_FAILURE",
	}
	TransferStatus_value = map[string]int32{
		"TRANSFER_STATUS_WAITING":              0,
		"TRANSFER_STATUS_RUNNING":              1,
		"TRANSFER_STATUS_DONE_SUCCESS":         2,
		"TRANSFER_STATUS_DONE_PARTIAL_FAILURE": 3,
		"TRANSFER_STATUS_DONE_TOTAL_FAILURE":   4,
	}
)

func (x TransferStatus) Enum() *TransferStatus {
	p := new(TransferStatus)
	*p = x
	return p
}

func (x TransferStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransferStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_delivery_proto_enumTypes[0].Descriptor()
}

func (TransferStatus) Type() protoreflect.EnumType {
	return &file_api_commons_delivery_proto_enumTypes[0]
}

func (x TransferStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransferStatus.Descriptor instead.
func (TransferStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_delivery_proto_rawDescGZIP(), []int{0}
}

// TODO
type Encryption struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Encryption) Reset() {
	*x = Encryption{}
	mi := &file_api_commons_delivery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Encryption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Encryption) ProtoMessage() {}

func (x *Encryption) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_delivery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Encryption.ProtoReflect.Descriptor instead.
func (*Encryption) Descriptor() ([]byte, []int) {
	return file_api_commons_delivery_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_delivery_proto protoreflect.FileDescriptor

const file_api_commons_delivery_proto_rawDesc = "" +
	"\n" +
	"\x1aapi/commons/delivery.proto\x12\vapi.commons\"\f\n" +
	"\n" +
	"Encryption*\xbe\x01\n" +
	"\x0eTransferStatus\x12\x1b\n" +
	"\x17TRANSFER_STATUS_WAITING\x10\x00\x12\x1b\n" +
	"\x17TRANSFER_STATUS_RUNNING\x10\x01\x12 \n" +
	"\x1cTRANSFER_STATUS_DONE_SUCCESS\x10\x02\x12(\n" +
	"$TRANSFER_STATUS_DONE_PARTIAL_FAILURE\x10\x03\x12&\n" +
	"\"TRANSFER_STATUS_DONE_TOTAL_FAILURE\x10\x04B\x95\x01\n" +
	"\x0fcom.api.commonsB\rDeliveryProtoP\x01Z&github.com/tcncloud/api-go/api/commons\xa2\x02\x03ACX\xaa\x02\vApi.Commons\xca\x02\vApi\\Commons\xe2\x02\x17Api\\Commons\\GPBMetadata\xea\x02\fApi::Commonsb\x06proto3"

var (
	file_api_commons_delivery_proto_rawDescOnce sync.Once
	file_api_commons_delivery_proto_rawDescData []byte
)

func file_api_commons_delivery_proto_rawDescGZIP() []byte {
	file_api_commons_delivery_proto_rawDescOnce.Do(func() {
		file_api_commons_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_delivery_proto_rawDesc), len(file_api_commons_delivery_proto_rawDesc)))
	})
	return file_api_commons_delivery_proto_rawDescData
}

var file_api_commons_delivery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_commons_delivery_proto_goTypes = []any{
	(TransferStatus)(0), // 0: api.commons.TransferStatus
	(*Encryption)(nil),  // 1: api.commons.Encryption
}
var file_api_commons_delivery_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_delivery_proto_init() }
func file_api_commons_delivery_proto_init() {
	if File_api_commons_delivery_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_delivery_proto_rawDesc), len(file_api_commons_delivery_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_delivery_proto_goTypes,
		DependencyIndexes: file_api_commons_delivery_proto_depIdxs,
		EnumInfos:         file_api_commons_delivery_proto_enumTypes,
		MessageInfos:      file_api_commons_delivery_proto_msgTypes,
	}.Build()
	File_api_commons_delivery_proto = out.File
	file_api_commons_delivery_proto_goTypes = nil
	file_api_commons_delivery_proto_depIdxs = nil
}
