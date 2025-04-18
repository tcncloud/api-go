// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/vanalytics/expr.proto

package vanalytics

import (
	commons "github.com/tcncloud/api-go/api/commons"
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

type Uint32Expr struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Where:
	//
	//	*Uint32Expr_Gt
	//	*Uint32Expr_Gte
	//	*Uint32Expr_Lt
	//	*Uint32Expr_Lte
	//	*Uint32Expr_Eq
	//	*Uint32Expr_NotEq
	//	*Uint32Expr_Range
	Where         isUint32Expr_Where `protobuf_oneof:"where"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Uint32Expr) Reset() {
	*x = Uint32Expr{}
	mi := &file_api_v1alpha1_vanalytics_expr_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Uint32Expr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint32Expr) ProtoMessage() {}

func (x *Uint32Expr) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_expr_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uint32Expr.ProtoReflect.Descriptor instead.
func (*Uint32Expr) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_expr_proto_rawDescGZIP(), []int{0}
}

func (x *Uint32Expr) GetWhere() isUint32Expr_Where {
	if x != nil {
		return x.Where
	}
	return nil
}

func (x *Uint32Expr) GetGt() uint32 {
	if x != nil {
		if x, ok := x.Where.(*Uint32Expr_Gt); ok {
			return x.Gt
		}
	}
	return 0
}

func (x *Uint32Expr) GetGte() uint32 {
	if x != nil {
		if x, ok := x.Where.(*Uint32Expr_Gte); ok {
			return x.Gte
		}
	}
	return 0
}

func (x *Uint32Expr) GetLt() uint32 {
	if x != nil {
		if x, ok := x.Where.(*Uint32Expr_Lt); ok {
			return x.Lt
		}
	}
	return 0
}

func (x *Uint32Expr) GetLte() uint32 {
	if x != nil {
		if x, ok := x.Where.(*Uint32Expr_Lte); ok {
			return x.Lte
		}
	}
	return 0
}

func (x *Uint32Expr) GetEq() uint32 {
	if x != nil {
		if x, ok := x.Where.(*Uint32Expr_Eq); ok {
			return x.Eq
		}
	}
	return 0
}

func (x *Uint32Expr) GetNotEq() uint32 {
	if x != nil {
		if x, ok := x.Where.(*Uint32Expr_NotEq); ok {
			return x.NotEq
		}
	}
	return 0
}

func (x *Uint32Expr) GetRange() *Uint32Range {
	if x != nil {
		if x, ok := x.Where.(*Uint32Expr_Range); ok {
			return x.Range
		}
	}
	return nil
}

type isUint32Expr_Where interface {
	isUint32Expr_Where()
}

type Uint32Expr_Gt struct {
	Gt uint32 `protobuf:"varint,1,opt,name=gt,proto3,oneof"`
}

type Uint32Expr_Gte struct {
	Gte uint32 `protobuf:"varint,2,opt,name=gte,proto3,oneof"`
}

type Uint32Expr_Lt struct {
	Lt uint32 `protobuf:"varint,3,opt,name=lt,proto3,oneof"`
}

type Uint32Expr_Lte struct {
	Lte uint32 `protobuf:"varint,4,opt,name=lte,proto3,oneof"`
}

type Uint32Expr_Eq struct {
	Eq uint32 `protobuf:"varint,5,opt,name=eq,proto3,oneof"`
}

type Uint32Expr_NotEq struct {
	NotEq uint32 `protobuf:"varint,6,opt,name=not_eq,json=notEq,proto3,oneof"`
}

type Uint32Expr_Range struct {
	Range *Uint32Range `protobuf:"bytes,7,opt,name=range,proto3,oneof"`
}

func (*Uint32Expr_Gt) isUint32Expr_Where() {}

func (*Uint32Expr_Gte) isUint32Expr_Where() {}

func (*Uint32Expr_Lt) isUint32Expr_Where() {}

func (*Uint32Expr_Lte) isUint32Expr_Where() {}

func (*Uint32Expr_Eq) isUint32Expr_Where() {}

func (*Uint32Expr_NotEq) isUint32Expr_Where() {}

func (*Uint32Expr_Range) isUint32Expr_Where() {}

type Uint32Range struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          uint32                 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To            uint32                 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	IncludeFrom   bool                   `protobuf:"varint,3,opt,name=include_from,json=includeFrom,proto3" json:"include_from,omitempty"`
	IncludeTo     bool                   `protobuf:"varint,4,opt,name=include_to,json=includeTo,proto3" json:"include_to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Uint32Range) Reset() {
	*x = Uint32Range{}
	mi := &file_api_v1alpha1_vanalytics_expr_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Uint32Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint32Range) ProtoMessage() {}

func (x *Uint32Range) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_expr_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uint32Range.ProtoReflect.Descriptor instead.
func (*Uint32Range) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_expr_proto_rawDescGZIP(), []int{1}
}

func (x *Uint32Range) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *Uint32Range) GetTo() uint32 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *Uint32Range) GetIncludeFrom() bool {
	if x != nil {
		return x.IncludeFrom
	}
	return false
}

func (x *Uint32Range) GetIncludeTo() bool {
	if x != nil {
		return x.IncludeTo
	}
	return false
}

type TimestampExpr struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Where:
	//
	//	*TimestampExpr_Range
	//	*TimestampExpr_Moment
	Where         isTimestampExpr_Where `protobuf_oneof:"where"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimestampExpr) Reset() {
	*x = TimestampExpr{}
	mi := &file_api_v1alpha1_vanalytics_expr_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimestampExpr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampExpr) ProtoMessage() {}

func (x *TimestampExpr) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_expr_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampExpr.ProtoReflect.Descriptor instead.
func (*TimestampExpr) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_expr_proto_rawDescGZIP(), []int{2}
}

func (x *TimestampExpr) GetWhere() isTimestampExpr_Where {
	if x != nil {
		return x.Where
	}
	return nil
}

func (x *TimestampExpr) GetRange() *TimestampRange {
	if x != nil {
		if x, ok := x.Where.(*TimestampExpr_Range); ok {
			return x.Range
		}
	}
	return nil
}

func (x *TimestampExpr) GetMoment() *Moment {
	if x != nil {
		if x, ok := x.Where.(*TimestampExpr_Moment); ok {
			return x.Moment
		}
	}
	return nil
}

type isTimestampExpr_Where interface {
	isTimestampExpr_Where()
}

type TimestampExpr_Range struct {
	// Optional. Static time period to match.
	Range *TimestampRange `protobuf:"bytes,1,opt,name=range,proto3,oneof"`
}

type TimestampExpr_Moment struct {
	// Optional. Dynamic time period to match.
	Moment *Moment `protobuf:"bytes,2,opt,name=moment,proto3,oneof"`
}

func (*TimestampExpr_Range) isTimestampExpr_Where() {}

func (*TimestampExpr_Moment) isTimestampExpr_Where() {}

type Moment struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. Timezone of the client.
	TimeZone string `protobuf:"bytes,1,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// Optional. Dynamic time period to match which
	// defaults to today.
	Interval      commons.Interval `protobuf:"varint,2,opt,name=interval,proto3,enum=api.commons.Interval" json:"interval,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Moment) Reset() {
	*x = Moment{}
	mi := &file_api_v1alpha1_vanalytics_expr_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Moment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Moment) ProtoMessage() {}

func (x *Moment) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_expr_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Moment.ProtoReflect.Descriptor instead.
func (*Moment) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_expr_proto_rawDescGZIP(), []int{3}
}

func (x *Moment) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *Moment) GetInterval() commons.Interval {
	if x != nil {
		return x.Interval
	}
	return commons.Interval(0)
}

type TimestampRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	IncludeFrom   bool                   `protobuf:"varint,3,opt,name=include_from,json=includeFrom,proto3" json:"include_from,omitempty"`
	IncludeTo     bool                   `protobuf:"varint,4,opt,name=include_to,json=includeTo,proto3" json:"include_to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimestampRange) Reset() {
	*x = TimestampRange{}
	mi := &file_api_v1alpha1_vanalytics_expr_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimestampRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampRange) ProtoMessage() {}

func (x *TimestampRange) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_expr_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampRange.ProtoReflect.Descriptor instead.
func (*TimestampRange) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_expr_proto_rawDescGZIP(), []int{4}
}

func (x *TimestampRange) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *TimestampRange) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *TimestampRange) GetIncludeFrom() bool {
	if x != nil {
		return x.IncludeFrom
	}
	return false
}

func (x *TimestampRange) GetIncludeTo() bool {
	if x != nil {
		return x.IncludeTo
	}
	return false
}

var File_api_v1alpha1_vanalytics_expr_proto protoreflect.FileDescriptor

const file_api_v1alpha1_vanalytics_expr_proto_rawDesc = "" +
	"\n" +
	"\"api/v1alpha1/vanalytics/expr.proto\x12\x17api.v1alpha1.vanalytics\x1a\x1capi/commons/vanalytics.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xca\x01\n" +
	"\n" +
	"Uint32Expr\x12\x10\n" +
	"\x02gt\x18\x01 \x01(\rH\x00R\x02gt\x12\x12\n" +
	"\x03gte\x18\x02 \x01(\rH\x00R\x03gte\x12\x10\n" +
	"\x02lt\x18\x03 \x01(\rH\x00R\x02lt\x12\x12\n" +
	"\x03lte\x18\x04 \x01(\rH\x00R\x03lte\x12\x10\n" +
	"\x02eq\x18\x05 \x01(\rH\x00R\x02eq\x12\x17\n" +
	"\x06not_eq\x18\x06 \x01(\rH\x00R\x05notEq\x12<\n" +
	"\x05range\x18\a \x01(\v2$.api.v1alpha1.vanalytics.Uint32RangeH\x00R\x05rangeB\a\n" +
	"\x05where\"s\n" +
	"\vUint32Range\x12\x12\n" +
	"\x04from\x18\x01 \x01(\rR\x04from\x12\x0e\n" +
	"\x02to\x18\x02 \x01(\rR\x02to\x12!\n" +
	"\finclude_from\x18\x03 \x01(\bR\vincludeFrom\x12\x1d\n" +
	"\n" +
	"include_to\x18\x04 \x01(\bR\tincludeTo\"\x94\x01\n" +
	"\rTimestampExpr\x12?\n" +
	"\x05range\x18\x01 \x01(\v2'.api.v1alpha1.vanalytics.TimestampRangeH\x00R\x05range\x129\n" +
	"\x06moment\x18\x02 \x01(\v2\x1f.api.v1alpha1.vanalytics.MomentH\x00R\x06momentB\a\n" +
	"\x05where\"X\n" +
	"\x06Moment\x12\x1b\n" +
	"\ttime_zone\x18\x01 \x01(\tR\btimeZone\x121\n" +
	"\binterval\x18\x02 \x01(\x0e2\x15.api.commons.IntervalR\binterval\"\xae\x01\n" +
	"\x0eTimestampRange\x12.\n" +
	"\x04from\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\x04from\x12*\n" +
	"\x02to\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x02to\x12!\n" +
	"\finclude_from\x18\x03 \x01(\bR\vincludeFrom\x12\x1d\n" +
	"\n" +
	"include_to\x18\x04 \x01(\bR\tincludeToB\xda\x01\n" +
	"\x1bcom.api.v1alpha1.vanalyticsB\tExprProtoP\x01Z2github.com/tcncloud/api-go/api/v1alpha1/vanalytics\xa2\x02\x03AVV\xaa\x02\x17Api.V1alpha1.Vanalytics\xca\x02\x17Api\\V1alpha1\\Vanalytics\xe2\x02#Api\\V1alpha1\\Vanalytics\\GPBMetadata\xea\x02\x19Api::V1alpha1::Vanalyticsb\x06proto3"

var (
	file_api_v1alpha1_vanalytics_expr_proto_rawDescOnce sync.Once
	file_api_v1alpha1_vanalytics_expr_proto_rawDescData []byte
)

func file_api_v1alpha1_vanalytics_expr_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_vanalytics_expr_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_vanalytics_expr_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_vanalytics_expr_proto_rawDesc), len(file_api_v1alpha1_vanalytics_expr_proto_rawDesc)))
	})
	return file_api_v1alpha1_vanalytics_expr_proto_rawDescData
}

var file_api_v1alpha1_vanalytics_expr_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1alpha1_vanalytics_expr_proto_goTypes = []any{
	(*Uint32Expr)(nil),            // 0: api.v1alpha1.vanalytics.Uint32Expr
	(*Uint32Range)(nil),           // 1: api.v1alpha1.vanalytics.Uint32Range
	(*TimestampExpr)(nil),         // 2: api.v1alpha1.vanalytics.TimestampExpr
	(*Moment)(nil),                // 3: api.v1alpha1.vanalytics.Moment
	(*TimestampRange)(nil),        // 4: api.v1alpha1.vanalytics.TimestampRange
	(commons.Interval)(0),         // 5: api.commons.Interval
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_api_v1alpha1_vanalytics_expr_proto_depIdxs = []int32{
	1, // 0: api.v1alpha1.vanalytics.Uint32Expr.range:type_name -> api.v1alpha1.vanalytics.Uint32Range
	4, // 1: api.v1alpha1.vanalytics.TimestampExpr.range:type_name -> api.v1alpha1.vanalytics.TimestampRange
	3, // 2: api.v1alpha1.vanalytics.TimestampExpr.moment:type_name -> api.v1alpha1.vanalytics.Moment
	5, // 3: api.v1alpha1.vanalytics.Moment.interval:type_name -> api.commons.Interval
	6, // 4: api.v1alpha1.vanalytics.TimestampRange.from:type_name -> google.protobuf.Timestamp
	6, // 5: api.v1alpha1.vanalytics.TimestampRange.to:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_vanalytics_expr_proto_init() }
func file_api_v1alpha1_vanalytics_expr_proto_init() {
	if File_api_v1alpha1_vanalytics_expr_proto != nil {
		return
	}
	file_api_v1alpha1_vanalytics_expr_proto_msgTypes[0].OneofWrappers = []any{
		(*Uint32Expr_Gt)(nil),
		(*Uint32Expr_Gte)(nil),
		(*Uint32Expr_Lt)(nil),
		(*Uint32Expr_Lte)(nil),
		(*Uint32Expr_Eq)(nil),
		(*Uint32Expr_NotEq)(nil),
		(*Uint32Expr_Range)(nil),
	}
	file_api_v1alpha1_vanalytics_expr_proto_msgTypes[2].OneofWrappers = []any{
		(*TimestampExpr_Range)(nil),
		(*TimestampExpr_Moment)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_vanalytics_expr_proto_rawDesc), len(file_api_v1alpha1_vanalytics_expr_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_vanalytics_expr_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_vanalytics_expr_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_vanalytics_expr_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_vanalytics_expr_proto = out.File
	file_api_v1alpha1_vanalytics_expr_proto_goTypes = nil
	file_api_v1alpha1_vanalytics_expr_proto_depIdxs = nil
}
