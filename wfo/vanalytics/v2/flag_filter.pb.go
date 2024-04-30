// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: wfo/vanalytics/v2/flag_filter.proto

package vanalyticsv2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ListFlagFiltersRequest is a request for listing flag filters.
type ListFlagFiltersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	FlagSids []int64 `protobuf:"varint,6,rep,packed,name=flag_sids,json=flagSids,proto3" json:"flag_sids,omitempty"`
}

func (x *ListFlagFiltersRequest) Reset() {
	*x = ListFlagFiltersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlagFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlagFiltersRequest) ProtoMessage() {}

func (x *ListFlagFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to retrieve the next page. Empty when there are no more pages.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of flag filters.
	FlagFilters []*FlagFilter `protobuf:"bytes,2,rep,name=flag_filters,json=flagFilters,proto3" json:"flag_filters,omitempty"`
}

func (x *ListFlagFiltersResponse) Reset() {
	*x = ListFlagFiltersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlagFiltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlagFiltersResponse) ProtoMessage() {}

func (x *ListFlagFiltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The unique id of this flag filter.
	FlagFilterSid int64 `protobuf:"varint,1,opt,name=flag_filter_sid,json=flagFilterSid,proto3" json:"flag_filter_sid,omitempty"`
	// Required. The unique id of the filter.
	FilterSid int64 `protobuf:"varint,2,opt,name=filter_sid,json=filterSid,proto3" json:"filter_sid,omitempty"`
	// Required. The unique id of the flag.
	FlagSid int64 `protobuf:"varint,3,opt,name=flag_sid,json=flagSid,proto3" json:"flag_sid,omitempty"`
	// Output only. The flag for this flag filter.
	Flag *Flag `protobuf:"bytes,5,opt,name=flag,proto3" json:"flag,omitempty"`
	// Output only. The filter for this flag filter.
	Filter *Filter `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *FlagFilter) Reset() {
	*x = FlagFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagFilter) ProtoMessage() {}

func (x *FlagFilter) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_wfo_vanalytics_v2_flag_filter_proto_rawDesc = []byte{
	0x0a, 0x23, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x77, 0x66, 0x6f, 0x2f,
	0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x77, 0x66, 0x6f, 0x2f,
	0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x6c,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x37, 0x0a, 0x09, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x08,
	0x66, 0x6c, 0x61, 0x67, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69,
	0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69,
	0x64, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x40, 0x0a, 0x0c, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77,
	0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x66, 0x6c, 0x61,
	0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0xce, 0x01, 0x0a, 0x0a, 0x46, 0x6c, 0x61,
	0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x6c, 0x61, 0x67, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x66, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x69, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x66, 0x6c, 0x61,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x6c, 0x61, 0x67,
	0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x31, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0xc9, 0x01, 0x0a, 0x15, 0x63, 0x6f,
	0x6d, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x76, 0x32, 0x42, 0x0f, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x76,
	0x32, 0xa2, 0x02, 0x03, 0x57, 0x56, 0x58, 0xaa, 0x02, 0x11, 0x57, 0x66, 0x6f, 0x2e, 0x56, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x11, 0x57, 0x66,
	0x6f, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c, 0x56, 0x32, 0xe2,
	0x02, 0x1d, 0x57, 0x66, 0x6f, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x13, 0x57, 0x66, 0x6f, 0x3a, 0x3a, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wfo_vanalytics_v2_flag_filter_proto_rawDescOnce sync.Once
	file_wfo_vanalytics_v2_flag_filter_proto_rawDescData = file_wfo_vanalytics_v2_flag_filter_proto_rawDesc
)

func file_wfo_vanalytics_v2_flag_filter_proto_rawDescGZIP() []byte {
	file_wfo_vanalytics_v2_flag_filter_proto_rawDescOnce.Do(func() {
		file_wfo_vanalytics_v2_flag_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_wfo_vanalytics_v2_flag_filter_proto_rawDescData)
	})
	return file_wfo_vanalytics_v2_flag_filter_proto_rawDescData
}

var file_wfo_vanalytics_v2_flag_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wfo_vanalytics_v2_flag_filter_proto_goTypes = []interface{}{
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
	if !protoimpl.UnsafeEnabled {
		file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFlagFiltersRequest); i {
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
		file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFlagFiltersResponse); i {
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
		file_wfo_vanalytics_v2_flag_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagFilter); i {
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
			RawDescriptor: file_wfo_vanalytics_v2_flag_filter_proto_rawDesc,
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
	file_wfo_vanalytics_v2_flag_filter_proto_rawDesc = nil
	file_wfo_vanalytics_v2_flag_filter_proto_goTypes = nil
	file_wfo_vanalytics_v2_flag_filter_proto_depIdxs = nil
}
