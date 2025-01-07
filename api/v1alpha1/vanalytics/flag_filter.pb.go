// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: api/v1alpha1/vanalytics/flag_filter.proto

package vanalytics

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

// CreateFlagFilterRequest is a request for creating a flag filter.
// DEPRECATED
type CreateFlagFilterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The flag filter resource to create.
	FlagFilter    *FlagFilter `protobuf:"bytes,1,opt,name=flag_filter,json=flagFilter,proto3" json:"flag_filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateFlagFilterRequest) Reset() {
	*x = CreateFlagFilterRequest{}
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFlagFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlagFilterRequest) ProtoMessage() {}

func (x *CreateFlagFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlagFilterRequest.ProtoReflect.Descriptor instead.
func (*CreateFlagFilterRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFlagFilterRequest) GetFlagFilter() *FlagFilter {
	if x != nil {
		return x.FlagFilter
	}
	return nil
}

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
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFlagFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlagFiltersRequest) ProtoMessage() {}

func (x *ListFlagFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[1]
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
	return file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescGZIP(), []int{1}
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
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFlagFiltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlagFiltersResponse) ProtoMessage() {}

func (x *ListFlagFiltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[2]
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
	return file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescGZIP(), []int{2}
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

// DeleteFlagFilterRequest is a request for deleting a flag filter.
// DEPRECATED
type DeleteFlagFilterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The unique id of the flag on the flag filter to be deleted.
	FlagSid int64 `protobuf:"varint,1,opt,name=flag_sid,json=flagSid,proto3" json:"flag_sid,omitempty"`
	// Optional. The unique id of the filter on the flag filter to be deleted.
	FilterSid int64 `protobuf:"varint,2,opt,name=filter_sid,json=filterSid,proto3" json:"filter_sid,omitempty"`
	// Optional. If true, then all flag filters with the given flag sid and/or
	// filter sid will be deleted.
	All           bool `protobuf:"varint,3,opt,name=all,proto3" json:"all,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFlagFilterRequest) Reset() {
	*x = DeleteFlagFilterRequest{}
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFlagFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlagFilterRequest) ProtoMessage() {}

func (x *DeleteFlagFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlagFilterRequest.ProtoReflect.Descriptor instead.
func (*DeleteFlagFilterRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteFlagFilterRequest) GetFlagSid() int64 {
	if x != nil {
		return x.FlagSid
	}
	return 0
}

func (x *DeleteFlagFilterRequest) GetFilterSid() int64 {
	if x != nil {
		return x.FilterSid
	}
	return 0
}

func (x *DeleteFlagFilterRequest) GetAll() bool {
	if x != nil {
		return x.All
	}
	return false
}

// DeleteFlagFilterResponse is a response for deleting a flag filter.
// DEPRECATED
type DeleteFlagFilterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFlagFilterResponse) Reset() {
	*x = DeleteFlagFilterResponse{}
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFlagFilterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlagFilterResponse) ProtoMessage() {}

func (x *DeleteFlagFilterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlagFilterResponse.ProtoReflect.Descriptor instead.
func (*DeleteFlagFilterResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescGZIP(), []int{4}
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
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlagFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagFilter) ProtoMessage() {}

func (x *FlagFilter) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes[5]
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
	return file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescGZIP(), []int{5}
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

var File_api_v1alpha1_vanalytics_flag_filter_proto protoreflect.FileDescriptor

var file_api_v1alpha1_vanalytics_flag_filter_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x1a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5f, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0b, 0x66,
	0x6c, 0x61, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x66, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x22, 0xe7, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x66, 0x6c, 0x61, 0x67,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x4d, 0x61, 0x73,
	0x6b, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x17,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x46, 0x0a, 0x0c, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x46, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x66, 0x6c, 0x61, 0x67,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x65, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x22, 0x1a,
	0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xda, 0x01, 0x0a, 0x0a, 0x46,
	0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x6c, 0x61,
	0x67, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x66, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x66,
	0x6c, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x37,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0xe0, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x42, 0x0f, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0xa2, 0x02,
	0x03, 0x41, 0x56, 0x56, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0xca, 0x02,
	0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x56, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a,
	0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescOnce sync.Once
	file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescData = file_api_v1alpha1_vanalytics_flag_filter_proto_rawDesc
)

func file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescData)
	})
	return file_api_v1alpha1_vanalytics_flag_filter_proto_rawDescData
}

var file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1alpha1_vanalytics_flag_filter_proto_goTypes = []any{
	(*CreateFlagFilterRequest)(nil),  // 0: api.v1alpha1.vanalytics.CreateFlagFilterRequest
	(*ListFlagFiltersRequest)(nil),   // 1: api.v1alpha1.vanalytics.ListFlagFiltersRequest
	(*ListFlagFiltersResponse)(nil),  // 2: api.v1alpha1.vanalytics.ListFlagFiltersResponse
	(*DeleteFlagFilterRequest)(nil),  // 3: api.v1alpha1.vanalytics.DeleteFlagFilterRequest
	(*DeleteFlagFilterResponse)(nil), // 4: api.v1alpha1.vanalytics.DeleteFlagFilterResponse
	(*FlagFilter)(nil),               // 5: api.v1alpha1.vanalytics.FlagFilter
	(*fieldmaskpb.FieldMask)(nil),    // 6: google.protobuf.FieldMask
	(*Flag)(nil),                     // 7: api.v1alpha1.vanalytics.Flag
	(*Filter)(nil),                   // 8: api.v1alpha1.vanalytics.Filter
}
var file_api_v1alpha1_vanalytics_flag_filter_proto_depIdxs = []int32{
	5, // 0: api.v1alpha1.vanalytics.CreateFlagFilterRequest.flag_filter:type_name -> api.v1alpha1.vanalytics.FlagFilter
	6, // 1: api.v1alpha1.vanalytics.ListFlagFiltersRequest.flag_mask:type_name -> google.protobuf.FieldMask
	6, // 2: api.v1alpha1.vanalytics.ListFlagFiltersRequest.filter_mask:type_name -> google.protobuf.FieldMask
	5, // 3: api.v1alpha1.vanalytics.ListFlagFiltersResponse.flag_filters:type_name -> api.v1alpha1.vanalytics.FlagFilter
	7, // 4: api.v1alpha1.vanalytics.FlagFilter.flag:type_name -> api.v1alpha1.vanalytics.Flag
	8, // 5: api.v1alpha1.vanalytics.FlagFilter.filter:type_name -> api.v1alpha1.vanalytics.Filter
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_vanalytics_flag_filter_proto_init() }
func file_api_v1alpha1_vanalytics_flag_filter_proto_init() {
	if File_api_v1alpha1_vanalytics_flag_filter_proto != nil {
		return
	}
	file_api_v1alpha1_vanalytics_filter_proto_init()
	file_api_v1alpha1_vanalytics_flag_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_vanalytics_flag_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_vanalytics_flag_filter_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_vanalytics_flag_filter_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_vanalytics_flag_filter_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_vanalytics_flag_filter_proto = out.File
	file_api_v1alpha1_vanalytics_flag_filter_proto_rawDesc = nil
	file_api_v1alpha1_vanalytics_flag_filter_proto_goTypes = nil
	file_api_v1alpha1_vanalytics_flag_filter_proto_depIdxs = nil
}
