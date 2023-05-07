// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/commons/delivery.proto

package commons

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Encryption) Reset() {
	*x = Encryption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_delivery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Encryption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Encryption) ProtoMessage() {}

func (x *Encryption) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_delivery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_api_commons_delivery_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0xbe, 0x01, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x57, 0x41,
	0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x5f, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x5f, 0x50,
	0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x03,
	0x12, 0x26, 0x0a, 0x22, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x04, 0x42, 0x95, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x0d, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_delivery_proto_rawDescOnce sync.Once
	file_api_commons_delivery_proto_rawDescData = file_api_commons_delivery_proto_rawDesc
)

func file_api_commons_delivery_proto_rawDescGZIP() []byte {
	file_api_commons_delivery_proto_rawDescOnce.Do(func() {
		file_api_commons_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_delivery_proto_rawDescData)
	})
	return file_api_commons_delivery_proto_rawDescData
}

var file_api_commons_delivery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_commons_delivery_proto_goTypes = []interface{}{
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
	if !protoimpl.UnsafeEnabled {
		file_api_commons_delivery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Encryption); i {
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
			RawDescriptor: file_api_commons_delivery_proto_rawDesc,
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
	file_api_commons_delivery_proto_rawDesc = nil
	file_api_commons_delivery_proto_goTypes = nil
	file_api_commons_delivery_proto_depIdxs = nil
}
