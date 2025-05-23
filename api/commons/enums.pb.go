// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/commons/enums.proto

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

type Month int32

const (
	Month_MONTH_JANUARY   Month = 0
	Month_MONTH_FEBRUARY  Month = 1
	Month_MONTH_MARCH     Month = 2
	Month_MONTH_APRIL     Month = 3
	Month_MONTH_MAY       Month = 4
	Month_MONTH_JUNE      Month = 5
	Month_MONTH_JULY      Month = 6
	Month_MONTH_AUGUST    Month = 7
	Month_MONTH_SEPTEMBER Month = 8
	Month_MONTH_OCTOBER   Month = 9
	Month_MONTH_NOVEMBER  Month = 10
	Month_MONTH_DECEMBER  Month = 11
)

// Enum value maps for Month.
var (
	Month_name = map[int32]string{
		0:  "MONTH_JANUARY",
		1:  "MONTH_FEBRUARY",
		2:  "MONTH_MARCH",
		3:  "MONTH_APRIL",
		4:  "MONTH_MAY",
		5:  "MONTH_JUNE",
		6:  "MONTH_JULY",
		7:  "MONTH_AUGUST",
		8:  "MONTH_SEPTEMBER",
		9:  "MONTH_OCTOBER",
		10: "MONTH_NOVEMBER",
		11: "MONTH_DECEMBER",
	}
	Month_value = map[string]int32{
		"MONTH_JANUARY":   0,
		"MONTH_FEBRUARY":  1,
		"MONTH_MARCH":     2,
		"MONTH_APRIL":     3,
		"MONTH_MAY":       4,
		"MONTH_JUNE":      5,
		"MONTH_JULY":      6,
		"MONTH_AUGUST":    7,
		"MONTH_SEPTEMBER": 8,
		"MONTH_OCTOBER":   9,
		"MONTH_NOVEMBER":  10,
		"MONTH_DECEMBER":  11,
	}
)

func (x Month) Enum() *Month {
	p := new(Month)
	*p = x
	return p
}

func (x Month) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Month) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_enums_proto_enumTypes[0].Descriptor()
}

func (Month) Type() protoreflect.EnumType {
	return &file_api_commons_enums_proto_enumTypes[0]
}

func (x Month) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Month.Descriptor instead.
func (Month) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_enums_proto_rawDescGZIP(), []int{0}
}

type Weekday_Enum int32

const (
	Weekday_SUNDAY    Weekday_Enum = 0
	Weekday_MONDAY    Weekday_Enum = 1
	Weekday_TUESDAY   Weekday_Enum = 2
	Weekday_WEDNESDAY Weekday_Enum = 3
	Weekday_THURSDAY  Weekday_Enum = 4
	Weekday_FRIDAY    Weekday_Enum = 5
	Weekday_SATURDAY  Weekday_Enum = 6
)

// Enum value maps for Weekday_Enum.
var (
	Weekday_Enum_name = map[int32]string{
		0: "SUNDAY",
		1: "MONDAY",
		2: "TUESDAY",
		3: "WEDNESDAY",
		4: "THURSDAY",
		5: "FRIDAY",
		6: "SATURDAY",
	}
	Weekday_Enum_value = map[string]int32{
		"SUNDAY":    0,
		"MONDAY":    1,
		"TUESDAY":   2,
		"WEDNESDAY": 3,
		"THURSDAY":  4,
		"FRIDAY":    5,
		"SATURDAY":  6,
	}
)

func (x Weekday_Enum) Enum() *Weekday_Enum {
	p := new(Weekday_Enum)
	*p = x
	return p
}

func (x Weekday_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Weekday_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_enums_proto_enumTypes[1].Descriptor()
}

func (Weekday_Enum) Type() protoreflect.EnumType {
	return &file_api_commons_enums_proto_enumTypes[1]
}

func (x Weekday_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Weekday_Enum.Descriptor instead.
func (Weekday_Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_enums_proto_rawDescGZIP(), []int{0, 0}
}

type CronRequestType_Enum int32

const (
	CronRequestType_INVALID CronRequestType_Enum = 0
	CronRequestType_SFTP    CronRequestType_Enum = 1
)

// Enum value maps for CronRequestType_Enum.
var (
	CronRequestType_Enum_name = map[int32]string{
		0: "INVALID",
		1: "SFTP",
	}
	CronRequestType_Enum_value = map[string]int32{
		"INVALID": 0,
		"SFTP":    1,
	}
)

func (x CronRequestType_Enum) Enum() *CronRequestType_Enum {
	p := new(CronRequestType_Enum)
	*p = x
	return p
}

func (x CronRequestType_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CronRequestType_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_enums_proto_enumTypes[2].Descriptor()
}

func (CronRequestType_Enum) Type() protoreflect.EnumType {
	return &file_api_commons_enums_proto_enumTypes[2]
}

func (x CronRequestType_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CronRequestType_Enum.Descriptor instead.
func (CronRequestType_Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_enums_proto_rawDescGZIP(), []int{1, 0}
}

type Weekday struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Weekday) Reset() {
	*x = Weekday{}
	mi := &file_api_commons_enums_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Weekday) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weekday) ProtoMessage() {}

func (x *Weekday) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_enums_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weekday.ProtoReflect.Descriptor instead.
func (*Weekday) Descriptor() ([]byte, []int) {
	return file_api_commons_enums_proto_rawDescGZIP(), []int{0}
}

type CronRequestType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CronRequestType) Reset() {
	*x = CronRequestType{}
	mi := &file_api_commons_enums_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CronRequestType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronRequestType) ProtoMessage() {}

func (x *CronRequestType) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_enums_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronRequestType.ProtoReflect.Descriptor instead.
func (*CronRequestType) Descriptor() ([]byte, []int) {
	return file_api_commons_enums_proto_rawDescGZIP(), []int{1}
}

var File_api_commons_enums_proto protoreflect.FileDescriptor

const file_api_commons_enums_proto_rawDesc = "" +
	"\n" +
	"\x17api/commons/enums.proto\x12\vapi.commons\"m\n" +
	"\aWeekday\"b\n" +
	"\x04Enum\x12\n" +
	"\n" +
	"\x06SUNDAY\x10\x00\x12\n" +
	"\n" +
	"\x06MONDAY\x10\x01\x12\v\n" +
	"\aTUESDAY\x10\x02\x12\r\n" +
	"\tWEDNESDAY\x10\x03\x12\f\n" +
	"\bTHURSDAY\x10\x04\x12\n" +
	"\n" +
	"\x06FRIDAY\x10\x05\x12\f\n" +
	"\bSATURDAY\x10\x06\"0\n" +
	"\x0fCronRequestType\"\x1d\n" +
	"\x04Enum\x12\v\n" +
	"\aINVALID\x10\x00\x12\b\n" +
	"\x04SFTP\x10\x01*\xe1\x01\n" +
	"\x05Month\x12\x11\n" +
	"\rMONTH_JANUARY\x10\x00\x12\x12\n" +
	"\x0eMONTH_FEBRUARY\x10\x01\x12\x0f\n" +
	"\vMONTH_MARCH\x10\x02\x12\x0f\n" +
	"\vMONTH_APRIL\x10\x03\x12\r\n" +
	"\tMONTH_MAY\x10\x04\x12\x0e\n" +
	"\n" +
	"MONTH_JUNE\x10\x05\x12\x0e\n" +
	"\n" +
	"MONTH_JULY\x10\x06\x12\x10\n" +
	"\fMONTH_AUGUST\x10\a\x12\x13\n" +
	"\x0fMONTH_SEPTEMBER\x10\b\x12\x11\n" +
	"\rMONTH_OCTOBER\x10\t\x12\x12\n" +
	"\x0eMONTH_NOVEMBER\x10\n" +
	"\x12\x12\n" +
	"\x0eMONTH_DECEMBER\x10\vB\x92\x01\n" +
	"\x0fcom.api.commonsB\n" +
	"EnumsProtoP\x01Z&github.com/tcncloud/api-go/api/commons\xa2\x02\x03ACX\xaa\x02\vApi.Commons\xca\x02\vApi\\Commons\xe2\x02\x17Api\\Commons\\GPBMetadata\xea\x02\fApi::Commonsb\x06proto3"

var (
	file_api_commons_enums_proto_rawDescOnce sync.Once
	file_api_commons_enums_proto_rawDescData []byte
)

func file_api_commons_enums_proto_rawDescGZIP() []byte {
	file_api_commons_enums_proto_rawDescOnce.Do(func() {
		file_api_commons_enums_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_enums_proto_rawDesc), len(file_api_commons_enums_proto_rawDesc)))
	})
	return file_api_commons_enums_proto_rawDescData
}

var file_api_commons_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_commons_enums_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_enums_proto_goTypes = []any{
	(Month)(0),                // 0: api.commons.Month
	(Weekday_Enum)(0),         // 1: api.commons.Weekday.Enum
	(CronRequestType_Enum)(0), // 2: api.commons.CronRequestType.Enum
	(*Weekday)(nil),           // 3: api.commons.Weekday
	(*CronRequestType)(nil),   // 4: api.commons.CronRequestType
}
var file_api_commons_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_enums_proto_init() }
func file_api_commons_enums_proto_init() {
	if File_api_commons_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_enums_proto_rawDesc), len(file_api_commons_enums_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_enums_proto_goTypes,
		DependencyIndexes: file_api_commons_enums_proto_depIdxs,
		EnumInfos:         file_api_commons_enums_proto_enumTypes,
		MessageInfos:      file_api_commons_enums_proto_msgTypes,
	}.Build()
	File_api_commons_enums_proto = out.File
	file_api_commons_enums_proto_goTypes = nil
	file_api_commons_enums_proto_depIdxs = nil
}
