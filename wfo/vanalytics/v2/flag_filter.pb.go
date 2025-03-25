// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: wfo/vanalytics/v2/flag_filter.proto

package vanalyticsv2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// ListFlagFiltersRequest is a request for listing flag filters.
type ListFlagFiltersRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The number of flag filters to include in a single response. When not
	// provided this defaults to 100.
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token returned from a previous List request, if any.
	// When provided all other request fields are ignored.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. Field mask for response flag. A missing or empty field mask is
	// interpreted as a field mask containing all possible fields.
	FlagMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=flag_mask,json=flagMask,proto3" json:"flag_mask,omitempty"`
	// Optional. Field mask for response filter. A missing or empty field mask is
	// interpreted as a field mask containing all possible fields.
	FilterMask *fieldmaskpb.FieldMask `protobuf:"bytes,5,opt,name=filter_mask,json=filterMask,proto3" json:"filter_mask,omitempty"`
	// Optional. List of flag sids. Requires response filters to be associated
	// with at least one of the provided flag sids.
	FlagSids      []int64 `protobuf:"varint,6,rep,packed,name=flag_sids,json=flagSids,proto3" json:"flag_sids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFlagFiltersRequest) Reset() {
	*x = ListFlagFiltersRequest{}
	mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFlagFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlagFiltersRequest) ProtoMessage() {}

func (x *ListFlagFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlagFiltersRequest.ProtoReflect.Descriptor instead.
func (*ListFlagFiltersRequest) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_flag_filter_proto_rawDescGZIP(), []int{0}
}

func (x *ListFlagFiltersRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListFlagFiltersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListFlagFiltersRequest) GetFlagMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FlagMask
	}
	return nil
}

func (x *ListFlagFiltersRequest) GetFilterMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FilterMask
	}
	return nil
}

func (x *ListFlagFiltersRequest) GetFlagSids() []int64 {
	if x != nil {
		return x.FlagSids
	}
	return nil
}

// ListFlagFiltersResponse is a response for listing flag filters.
type ListFlagFiltersResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Token to retrieve the next page. Empty when there are no more pages.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of flag filters.
	FlagFilters   []*FlagFilter `protobuf:"bytes,2,rep,name=flag_filters,json=flagFilters,proto3" json:"flag_filters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFlagFiltersResponse) Reset() {
	*x = ListFlagFiltersResponse{}
	mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFlagFiltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlagFiltersResponse) ProtoMessage() {}

func (x *ListFlagFiltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlagFiltersResponse.ProtoReflect.Descriptor instead.
func (*ListFlagFiltersResponse) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_flag_filter_proto_rawDescGZIP(), []int{1}
}

func (x *ListFlagFiltersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListFlagFiltersResponse) GetFlagFilters() []*FlagFilter {
	if x != nil {
		return x.FlagFilters
	}
	return nil
}

// FlagFilter is a resource in the Vanalytics API.
type FlagFilter struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The unique id of this flag filter.
	FlagFilterSid int64 `protobuf:"varint,1,opt,name=flag_filter_sid,json=flagFilterSid,proto3" json:"flag_filter_sid,omitempty"`
	// Required. The unique id of the filter.
	FilterSid int64 `protobuf:"varint,2,opt,name=filter_sid,json=filterSid,proto3" json:"filter_sid,omitempty"`
	// Required. The unique id of the flag.
	FlagSid int64 `protobuf:"varint,3,opt,name=flag_sid,json=flagSid,proto3" json:"flag_sid,omitempty"`
	// Output only. The flag for this flag filter.
	Flag *Flag `protobuf:"bytes,5,opt,name=flag,proto3" json:"flag,omitempty"`
	// Output only. The filter for this flag filter.
	Filter        *Filter `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlagFilter) Reset() {
	*x = FlagFilter{}
	mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlagFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagFilter) ProtoMessage() {}

func (x *FlagFilter) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagFilter.ProtoReflect.Descriptor instead.
func (*FlagFilter) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_flag_filter_proto_rawDescGZIP(), []int{2}
}

func (x *FlagFilter) GetFlagFilterSid() int64 {
	if x != nil {
		return x.FlagFilterSid
	}
	return 0
}

func (x *FlagFilter) GetFilterSid() int64 {
	if x != nil {
		return x.FilterSid
	}
	return 0
}

func (x *FlagFilter) GetFlagSid() int64 {
	if x != nil {
		return x.FlagSid
	}
	return 0
}

func (x *FlagFilter) GetFlag() *Flag {
	if x != nil {
		return x.Flag
	}
	return nil
}

func (x *FlagFilter) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

var File_wfo_vanalytics_v2_flag_filter_proto protoreflect.FileDescriptor

const file_wfo_vanalytics_v2_flag_filter_proto_rawDesc = "" +
	"\n" +
	"#wfo/vanalytics/v2/flag_filter.proto\x12\x11wfo.vanalytics.v2\x1a google/protobuf/field_mask.proto\x1a\x1ewfo/vanalytics/v2/filter.proto\x1a\x1cwfo/vanalytics/v2/flag.proto\"\xe7\x01\n" +
	"\x16ListFlagFiltersRequest\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\rR\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tR\tpageToken\x127\n" +
	"\tflag_mask\x18\x04 \x01(\v2\x1a.google.protobuf.FieldMaskR\bflagMask\x12;\n" +
	"\vfilter_mask\x18\x05 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"filterMask\x12\x1b\n" +
	"\tflag_sids\x18\x06 \x03(\x03R\bflagSids\"\x83\x01\n" +
	"\x17ListFlagFiltersResponse\x12&\n" +
	"\x0fnext_page_token\x18\x01 \x01(\tR\rnextPageToken\x12@\n" +
	"\fflag_filters\x18\x02 \x03(\v2\x1d.wfo.vanalytics.v2.FlagFilterR\vflagFilters\"\xce\x01\n" +
	"\n" +
	"FlagFilter\x12&\n" +
	"\x0fflag_filter_sid\x18\x01 \x01(\x03R\rflagFilterSid\x12\x1d\n" +
	"\n" +
	"filter_sid\x18\x02 \x01(\x03R\tfilterSid\x12\x19\n" +
	"\bflag_sid\x18\x03 \x01(\x03R\aflagSid\x12+\n" +
	"\x04flag\x18\x05 \x01(\v2\x17.wfo.vanalytics.v2.FlagR\x04flag\x121\n" +
	"\x06filter\x18\x06 \x01(\v2\x19.wfo.vanalytics.v2.FilterR\x06filterB\xc9\x01\n" +
	"\x15com.wfo.vanalytics.v2B\x0fFlagFilterProtoP\x01Z9github.com/tcncloud/api-go/wfo/vanalytics/v2;vanalyticsv2\xa2\x02\x03WVX\xaa\x02\x11Wfo.Vanalytics.V2\xca\x02\x11Wfo\\Vanalytics\\V2\xe2\x02\x1dWfo\\Vanalytics\\V2\\GPBMetadata\xea\x02\x13Wfo::Vanalytics::V2b\x06proto3"

var (
	file_wfo_vanalytics_v2_flag_filter_proto_rawDescOnce sync.Once
	file_wfo_vanalytics_v2_flag_filter_proto_rawDescData []byte
)

func file_wfo_vanalytics_v2_flag_filter_proto_rawDescGZIP() []byte {
	file_wfo_vanalytics_v2_flag_filter_proto_rawDescOnce.Do(func() {
		file_wfo_vanalytics_v2_flag_filter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_wfo_vanalytics_v2_flag_filter_proto_rawDesc), len(file_wfo_vanalytics_v2_flag_filter_proto_rawDesc)))
	})
	return file_wfo_vanalytics_v2_flag_filter_proto_rawDescData
}

var file_wfo_vanalytics_v2_flag_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wfo_vanalytics_v2_flag_filter_proto_goTypes = []any{
	(*ListFlagFiltersRequest)(nil),  // 0: wfo.vanalytics.v2.ListFlagFiltersRequest
	(*ListFlagFiltersResponse)(nil), // 1: wfo.vanalytics.v2.ListFlagFiltersResponse
	(*FlagFilter)(nil),              // 2: wfo.vanalytics.v2.FlagFilter
	(*fieldmaskpb.FieldMask)(nil),   // 3: google.protobuf.FieldMask
	(*Flag)(nil),                    // 4: wfo.vanalytics.v2.Flag
	(*Filter)(nil),                  // 5: wfo.vanalytics.v2.Filter
}
var file_wfo_vanalytics_v2_flag_filter_proto_depIdxs = []int32{
	3, // 0: wfo.vanalytics.v2.ListFlagFiltersRequest.flag_mask:type_name -> google.protobuf.FieldMask
	3, // 1: wfo.vanalytics.v2.ListFlagFiltersRequest.filter_mask:type_name -> google.protobuf.FieldMask
	2, // 2: wfo.vanalytics.v2.ListFlagFiltersResponse.flag_filters:type_name -> wfo.vanalytics.v2.FlagFilter
	4, // 3: wfo.vanalytics.v2.FlagFilter.flag:type_name -> wfo.vanalytics.v2.Flag
	5, // 4: wfo.vanalytics.v2.FlagFilter.filter:type_name -> wfo.vanalytics.v2.Filter
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_wfo_vanalytics_v2_flag_filter_proto_init() }
func file_wfo_vanalytics_v2_flag_filter_proto_init() {
	if File_wfo_vanalytics_v2_flag_filter_proto != nil {
		return
	}
	file_wfo_vanalytics_v2_filter_proto_init()
	file_wfo_vanalytics_v2_flag_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_wfo_vanalytics_v2_flag_filter_proto_rawDesc), len(file_wfo_vanalytics_v2_flag_filter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wfo_vanalytics_v2_flag_filter_proto_goTypes,
		DependencyIndexes: file_wfo_vanalytics_v2_flag_filter_proto_depIdxs,
		MessageInfos:      file_wfo_vanalytics_v2_flag_filter_proto_msgTypes,
	}.Build()
	File_wfo_vanalytics_v2_flag_filter_proto = out.File
	file_wfo_vanalytics_v2_flag_filter_proto_goTypes = nil
	file_wfo_vanalytics_v2_flag_filter_proto_depIdxs = nil
}
