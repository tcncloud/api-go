// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/commons/communication.proto

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

type EmailType_Enum int32

const (
	EmailType_OUTBOUND EmailType_Enum = 0
)

// Enum value maps for EmailType_Enum.
var (
	EmailType_Enum_name = map[int32]string{
		0: "OUTBOUND",
	}
	EmailType_Enum_value = map[string]int32{
		"OUTBOUND": 0,
	}
)

func (x EmailType_Enum) Enum() *EmailType_Enum {
	p := new(EmailType_Enum)
	*p = x
	return p
}

func (x EmailType_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmailType_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_communication_proto_enumTypes[0].Descriptor()
}

func (EmailType_Enum) Type() protoreflect.EnumType {
	return &file_api_commons_communication_proto_enumTypes[0]
}

func (x EmailType_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmailType_Enum.Descriptor instead.
func (EmailType_Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_communication_proto_rawDescGZIP(), []int{1, 0}
}

type SmsType_Enum int32

const (
	SmsType_OUTBOUND SmsType_Enum = 0
)

// Enum value maps for SmsType_Enum.
var (
	SmsType_Enum_name = map[int32]string{
		0: "OUTBOUND",
	}
	SmsType_Enum_value = map[string]int32{
		"OUTBOUND": 0,
	}
)

func (x SmsType_Enum) Enum() *SmsType_Enum {
	p := new(SmsType_Enum)
	*p = x
	return p
}

func (x SmsType_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SmsType_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_communication_proto_enumTypes[1].Descriptor()
}

func (SmsType_Enum) Type() protoreflect.EnumType {
	return &file_api_commons_communication_proto_enumTypes[1]
}

func (x SmsType_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SmsType_Enum.Descriptor instead.
func (SmsType_Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_communication_proto_rawDescGZIP(), []int{2, 0}
}

type WhatsAppType_Enum int32

const (
	WhatsAppType_OUTBOUND WhatsAppType_Enum = 0
)

// Enum value maps for WhatsAppType_Enum.
var (
	WhatsAppType_Enum_name = map[int32]string{
		0: "OUTBOUND",
	}
	WhatsAppType_Enum_value = map[string]int32{
		"OUTBOUND": 0,
	}
)

func (x WhatsAppType_Enum) Enum() *WhatsAppType_Enum {
	p := new(WhatsAppType_Enum)
	*p = x
	return p
}

func (x WhatsAppType_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WhatsAppType_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_communication_proto_enumTypes[2].Descriptor()
}

func (WhatsAppType_Enum) Type() protoreflect.EnumType {
	return &file_api_commons_communication_proto_enumTypes[2]
}

func (x WhatsAppType_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WhatsAppType_Enum.Descriptor instead.
func (WhatsAppType_Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_communication_proto_rawDescGZIP(), []int{3, 0}
}

// type of communication
type CommType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*CommType_CallType
	//	*CommType_EmailType
	//	*CommType_SmsType
	//	*CommType_WhatsappType
	Type isCommType_Type `protobuf_oneof:"type"`
}

func (x *CommType) Reset() {
	*x = CommType{}
	mi := &file_api_commons_communication_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommType) ProtoMessage() {}

func (x *CommType) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_communication_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommType.ProtoReflect.Descriptor instead.
func (*CommType) Descriptor() ([]byte, []int) {
	return file_api_commons_communication_proto_rawDescGZIP(), []int{0}
}

func (m *CommType) GetType() isCommType_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *CommType) GetCallType() CallType_Enum {
	if x, ok := x.GetType().(*CommType_CallType); ok {
		return x.CallType
	}
	return CallType_INBOUND
}

func (x *CommType) GetEmailType() EmailType_Enum {
	if x, ok := x.GetType().(*CommType_EmailType); ok {
		return x.EmailType
	}
	return EmailType_OUTBOUND
}

func (x *CommType) GetSmsType() SmsType_Enum {
	if x, ok := x.GetType().(*CommType_SmsType); ok {
		return x.SmsType
	}
	return SmsType_OUTBOUND
}

func (x *CommType) GetWhatsappType() WhatsAppType_Enum {
	if x, ok := x.GetType().(*CommType_WhatsappType); ok {
		return x.WhatsappType
	}
	return WhatsAppType_OUTBOUND
}

type isCommType_Type interface {
	isCommType_Type()
}

type CommType_CallType struct {
	CallType CallType_Enum `protobuf:"varint,1,opt,name=call_type,json=callType,proto3,enum=api.commons.CallType_Enum,oneof"`
}

type CommType_EmailType struct {
	EmailType EmailType_Enum `protobuf:"varint,2,opt,name=email_type,json=emailType,proto3,enum=api.commons.EmailType_Enum,oneof"`
}

type CommType_SmsType struct {
	SmsType SmsType_Enum `protobuf:"varint,3,opt,name=sms_type,json=smsType,proto3,enum=api.commons.SmsType_Enum,oneof"`
}

type CommType_WhatsappType struct {
	WhatsappType WhatsAppType_Enum `protobuf:"varint,4,opt,name=whatsapp_type,json=whatsappType,proto3,enum=api.commons.WhatsAppType_Enum,oneof"`
}

func (*CommType_CallType) isCommType_Type() {}

func (*CommType_EmailType) isCommType_Type() {}

func (*CommType_SmsType) isCommType_Type() {}

func (*CommType_WhatsappType) isCommType_Type() {}

type EmailType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmailType) Reset() {
	*x = EmailType{}
	mi := &file_api_commons_communication_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailType) ProtoMessage() {}

func (x *EmailType) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_communication_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailType.ProtoReflect.Descriptor instead.
func (*EmailType) Descriptor() ([]byte, []int) {
	return file_api_commons_communication_proto_rawDescGZIP(), []int{1}
}

type SmsType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SmsType) Reset() {
	*x = SmsType{}
	mi := &file_api_commons_communication_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmsType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsType) ProtoMessage() {}

func (x *SmsType) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_communication_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsType.ProtoReflect.Descriptor instead.
func (*SmsType) Descriptor() ([]byte, []int) {
	return file_api_commons_communication_proto_rawDescGZIP(), []int{2}
}

type WhatsAppType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WhatsAppType) Reset() {
	*x = WhatsAppType{}
	mi := &file_api_commons_communication_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WhatsAppType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhatsAppType) ProtoMessage() {}

func (x *WhatsAppType) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_communication_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhatsAppType.ProtoReflect.Descriptor instead.
func (*WhatsAppType) Descriptor() ([]byte, []int) {
	return file_api_commons_communication_proto_rawDescGZIP(), []int{3}
}

var File_api_commons_communication_proto protoreflect.FileDescriptor

var file_api_commons_communication_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x15,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x48, 0x00, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a,
	0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00,
	0x52, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x73,
	0x6d, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00, 0x52, 0x07, 0x73, 0x6d, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x57, 0x68, 0x61, 0x74, 0x73, 0x41, 0x70,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00, 0x52, 0x0c, 0x77, 0x68,
	0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x14, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x55, 0x54, 0x42, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x00, 0x22, 0x1f, 0x0a, 0x07, 0x53, 0x6d, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x14, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x55, 0x54, 0x42,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x22, 0x24, 0x0a, 0x0c, 0x57, 0x68, 0x61, 0x74, 0x73, 0x41,
	0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0x14, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0c,
	0x0a, 0x08, 0x4f, 0x55, 0x54, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x42, 0x9a, 0x01, 0x0a,
	0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x42, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02,
	0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_commons_communication_proto_rawDescOnce sync.Once
	file_api_commons_communication_proto_rawDescData = file_api_commons_communication_proto_rawDesc
)

func file_api_commons_communication_proto_rawDescGZIP() []byte {
	file_api_commons_communication_proto_rawDescOnce.Do(func() {
		file_api_commons_communication_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_communication_proto_rawDescData)
	})
	return file_api_commons_communication_proto_rawDescData
}

var file_api_commons_communication_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_commons_communication_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_commons_communication_proto_goTypes = []any{
	(EmailType_Enum)(0),    // 0: api.commons.EmailType.Enum
	(SmsType_Enum)(0),      // 1: api.commons.SmsType.Enum
	(WhatsAppType_Enum)(0), // 2: api.commons.WhatsAppType.Enum
	(*CommType)(nil),       // 3: api.commons.CommType
	(*EmailType)(nil),      // 4: api.commons.EmailType
	(*SmsType)(nil),        // 5: api.commons.SmsType
	(*WhatsAppType)(nil),   // 6: api.commons.WhatsAppType
	(CallType_Enum)(0),     // 7: api.commons.CallType.Enum
}
var file_api_commons_communication_proto_depIdxs = []int32{
	7, // 0: api.commons.CommType.call_type:type_name -> api.commons.CallType.Enum
	0, // 1: api.commons.CommType.email_type:type_name -> api.commons.EmailType.Enum
	1, // 2: api.commons.CommType.sms_type:type_name -> api.commons.SmsType.Enum
	2, // 3: api.commons.CommType.whatsapp_type:type_name -> api.commons.WhatsAppType.Enum
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_commons_communication_proto_init() }
func file_api_commons_communication_proto_init() {
	if File_api_commons_communication_proto != nil {
		return
	}
	file_api_commons_acd_proto_init()
	file_api_commons_communication_proto_msgTypes[0].OneofWrappers = []any{
		(*CommType_CallType)(nil),
		(*CommType_EmailType)(nil),
		(*CommType_SmsType)(nil),
		(*CommType_WhatsappType)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_communication_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_communication_proto_goTypes,
		DependencyIndexes: file_api_commons_communication_proto_depIdxs,
		EnumInfos:         file_api_commons_communication_proto_enumTypes,
		MessageInfos:      file_api_commons_communication_proto_msgTypes,
	}.Build()
	File_api_commons_communication_proto = out.File
	file_api_commons_communication_proto_rawDesc = nil
	file_api_commons_communication_proto_goTypes = nil
	file_api_commons_communication_proto_depIdxs = nil
}
