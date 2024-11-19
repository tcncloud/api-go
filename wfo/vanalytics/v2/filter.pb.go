// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: wfo/vanalytics/v2/filter.proto

package vanalyticsv2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The filter resource to create.
	Filter *Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *CreateFilterRequest) Reset() {
	*x = CreateFilterRequest{}
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFilterRequest) ProtoMessage() {}

func (x *CreateFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFilterRequest.ProtoReflect.Descriptor instead.
func (*CreateFilterRequest) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_filter_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFilterRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListFiltersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The number of filters to include in a single response. When not
	// provided this defaults to 100.
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The order by which filters will be listed. Follows sql order by
	// syntax. When not provided the order defaults to "name".
	// Supported order by includes:
	//   - (name)
	//   - (name desc)
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// Optional. The next_page_token returned from a previous List request, if any.
	// When provided all other request fields are ignored.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Types that are assignable to Where:
	//
	//	*ListFiltersRequest_Conflict
	//	*ListFiltersRequest_FlagSid
	Where isListFiltersRequest_Where `protobuf_oneof:"where"`
}

func (x *ListFiltersRequest) Reset() {
	*x = ListFiltersRequest{}
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFiltersRequest) ProtoMessage() {}

func (x *ListFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFiltersRequest.ProtoReflect.Descriptor instead.
func (*ListFiltersRequest) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_filter_proto_rawDescGZIP(), []int{1}
}

func (x *ListFiltersRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListFiltersRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListFiltersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (m *ListFiltersRequest) GetWhere() isListFiltersRequest_Where {
	if m != nil {
		return m.Where
	}
	return nil
}

func (x *ListFiltersRequest) GetConflict() *Filter {
	if x, ok := x.GetWhere().(*ListFiltersRequest_Conflict); ok {
		return x.Conflict
	}
	return nil
}

func (x *ListFiltersRequest) GetFlagSid() int64 {
	if x, ok := x.GetWhere().(*ListFiltersRequest_FlagSid); ok {
		return x.FlagSid
	}
	return 0
}

type isListFiltersRequest_Where interface {
	isListFiltersRequest_Where()
}

type ListFiltersRequest_Conflict struct {
	// Optional. The filter by which to find other filters which have a conflicting
	// name field.
	Conflict *Filter `protobuf:"bytes,5,opt,name=conflict,proto3,oneof"`
}

type ListFiltersRequest_FlagSid struct {
	// Optional. Lists filters which are associated with given flag sid.
	FlagSid int64 `protobuf:"varint,6,opt,name=flag_sid,json=flagSid,proto3,oneof"`
}

func (*ListFiltersRequest_Conflict) isListFiltersRequest_Where() {}

func (*ListFiltersRequest_FlagSid) isListFiltersRequest_Where() {}

type ListFiltersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to retrieve the next page of filters, or empty if there are no
	// more filters in the list.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of filters which contains at most one request page_size.
	Filters []*Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListFiltersResponse) Reset() {
	*x = ListFiltersResponse{}
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFiltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFiltersResponse) ProtoMessage() {}

func (x *ListFiltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFiltersResponse.ProtoReflect.Descriptor instead.
func (*ListFiltersResponse) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_filter_proto_rawDescGZIP(), []int{2}
}

func (x *ListFiltersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListFiltersResponse) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type UpdateFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The new filter data.
	Filter *Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. update_mask contains a list of paths to be updated. Possible paths include
	// name and transcript_query. If no update_mask is provided it defaults to
	// using both name and transcript_query.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The ID of the filter to be updated.
	FilterSid int64 `protobuf:"varint,3,opt,name=filter_sid,json=filterSid,proto3" json:"filter_sid,omitempty"`
}

func (x *UpdateFilterRequest) Reset() {
	*x = UpdateFilterRequest{}
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFilterRequest) ProtoMessage() {}

func (x *UpdateFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFilterRequest.ProtoReflect.Descriptor instead.
func (*UpdateFilterRequest) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_filter_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFilterRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *UpdateFilterRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateFilterRequest) GetFilterSid() int64 {
	if x != nil {
		return x.FilterSid
	}
	return 0
}

type DeleteFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The unique id of the filter to be deleted.
	FilterSid int64 `protobuf:"varint,1,opt,name=filter_sid,json=filterSid,proto3" json:"filter_sid,omitempty"`
	// Optional. Return the filter that was deleted.
	Return bool `protobuf:"varint,3,opt,name=return,proto3" json:"return,omitempty"`
}

func (x *DeleteFilterRequest) Reset() {
	*x = DeleteFilterRequest{}
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFilterRequest) ProtoMessage() {}

func (x *DeleteFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFilterRequest.ProtoReflect.Descriptor instead.
func (*DeleteFilterRequest) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_filter_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFilterRequest) GetFilterSid() int64 {
	if x != nil {
		return x.FilterSid
	}
	return 0
}

func (x *DeleteFilterRequest) GetReturn() bool {
	if x != nil {
		return x.Return
	}
	return false
}

type DeleteFilterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The filter that was deleted.
	Filter *Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *DeleteFilterResponse) Reset() {
	*x = DeleteFilterResponse{}
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFilterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFilterResponse) ProtoMessage() {}

func (x *DeleteFilterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFilterResponse.ProtoReflect.Descriptor instead.
func (*DeleteFilterResponse) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_filter_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteFilterResponse) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Where:
	//
	//	*GetFilterRequest_Name
	//	*GetFilterRequest_FilterSid
	Where isGetFilterRequest_Where `protobuf_oneof:"where"`
}

func (x *GetFilterRequest) Reset() {
	*x = GetFilterRequest{}
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilterRequest) ProtoMessage() {}

func (x *GetFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilterRequest.ProtoReflect.Descriptor instead.
func (*GetFilterRequest) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_filter_proto_rawDescGZIP(), []int{6}
}

func (m *GetFilterRequest) GetWhere() isGetFilterRequest_Where {
	if m != nil {
		return m.Where
	}
	return nil
}

func (x *GetFilterRequest) GetName() string {
	if x, ok := x.GetWhere().(*GetFilterRequest_Name); ok {
		return x.Name
	}
	return ""
}

func (x *GetFilterRequest) GetFilterSid() int64 {
	if x, ok := x.GetWhere().(*GetFilterRequest_FilterSid); ok {
		return x.FilterSid
	}
	return 0
}

type isGetFilterRequest_Where interface {
	isGetFilterRequest_Where()
}

type GetFilterRequest_Name struct {
	// The name of the filter.
	Name string `protobuf:"bytes,3,opt,name=name,proto3,oneof"`
}

type GetFilterRequest_FilterSid struct {
	// The unique id of the filter.
	FilterSid int64 `protobuf:"varint,4,opt,name=filter_sid,json=filterSid,proto3,oneof"`
}

func (*GetFilterRequest_Name) isGetFilterRequest_Where() {}

func (*GetFilterRequest_FilterSid) isGetFilterRequest_Where() {}

// A filter resource in the Vanalytics API.
type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The unique id of this filter.
	FilterSid int64 `protobuf:"varint,1,opt,name=filter_sid,json=filterSid,proto3" json:"filter_sid,omitempty"`
	// Required. The name of this filter. Must be non empty and unique across all
	// filters within an organization.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The timestamp when this filter was created. Assigned by the
	// server.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The version of this filter.
	Version int64 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	// The transcript query to be used to filter transcripts.
	TranscriptQuery *TranscriptQuery `protobuf:"bytes,7,opt,name=transcript_query,json=transcriptQuery,proto3" json:"transcript_query,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_filter_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_filter_proto_rawDescGZIP(), []int{7}
}

func (x *Filter) GetFilterSid() int64 {
	if x != nil {
		return x.FilterSid
	}
	return 0
}

func (x *Filter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Filter) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Filter) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Filter) GetTranscriptQuery() *TranscriptQuery {
	if x != nil {
		return x.TranscriptQuery
	}
	return nil
}

var File_wfo_vanalytics_v2_filter_proto protoreflect.FileDescriptor

var file_wfo_vanalytics_v2_filter_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x76, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x22, 0xca, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x37, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x12, 0x1b, 0x0a, 0x08, 0x66,
	0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x07, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x77, 0x68, 0x65, 0x72,
	0x65, 0x22, 0x72, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x33, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x49, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x52, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x69, 0x64,
	0x42, 0x07, 0x0a, 0x05, 0x77, 0x68, 0x65, 0x72, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x06, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x53, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4d,
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0xc5, 0x01,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50,
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
	file_wfo_vanalytics_v2_filter_proto_rawDescOnce sync.Once
	file_wfo_vanalytics_v2_filter_proto_rawDescData = file_wfo_vanalytics_v2_filter_proto_rawDesc
)

func file_wfo_vanalytics_v2_filter_proto_rawDescGZIP() []byte {
	file_wfo_vanalytics_v2_filter_proto_rawDescOnce.Do(func() {
		file_wfo_vanalytics_v2_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_wfo_vanalytics_v2_filter_proto_rawDescData)
	})
	return file_wfo_vanalytics_v2_filter_proto_rawDescData
}

var file_wfo_vanalytics_v2_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_wfo_vanalytics_v2_filter_proto_goTypes = []any{
	(*CreateFilterRequest)(nil),   // 0: wfo.vanalytics.v2.CreateFilterRequest
	(*ListFiltersRequest)(nil),    // 1: wfo.vanalytics.v2.ListFiltersRequest
	(*ListFiltersResponse)(nil),   // 2: wfo.vanalytics.v2.ListFiltersResponse
	(*UpdateFilterRequest)(nil),   // 3: wfo.vanalytics.v2.UpdateFilterRequest
	(*DeleteFilterRequest)(nil),   // 4: wfo.vanalytics.v2.DeleteFilterRequest
	(*DeleteFilterResponse)(nil),  // 5: wfo.vanalytics.v2.DeleteFilterResponse
	(*GetFilterRequest)(nil),      // 6: wfo.vanalytics.v2.GetFilterRequest
	(*Filter)(nil),                // 7: wfo.vanalytics.v2.Filter
	(*fieldmaskpb.FieldMask)(nil), // 8: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*TranscriptQuery)(nil),       // 10: wfo.vanalytics.v2.TranscriptQuery
}
var file_wfo_vanalytics_v2_filter_proto_depIdxs = []int32{
	7,  // 0: wfo.vanalytics.v2.CreateFilterRequest.filter:type_name -> wfo.vanalytics.v2.Filter
	7,  // 1: wfo.vanalytics.v2.ListFiltersRequest.conflict:type_name -> wfo.vanalytics.v2.Filter
	7,  // 2: wfo.vanalytics.v2.ListFiltersResponse.filters:type_name -> wfo.vanalytics.v2.Filter
	7,  // 3: wfo.vanalytics.v2.UpdateFilterRequest.filter:type_name -> wfo.vanalytics.v2.Filter
	8,  // 4: wfo.vanalytics.v2.UpdateFilterRequest.update_mask:type_name -> google.protobuf.FieldMask
	7,  // 5: wfo.vanalytics.v2.DeleteFilterResponse.filter:type_name -> wfo.vanalytics.v2.Filter
	9,  // 6: wfo.vanalytics.v2.Filter.create_time:type_name -> google.protobuf.Timestamp
	10, // 7: wfo.vanalytics.v2.Filter.transcript_query:type_name -> wfo.vanalytics.v2.TranscriptQuery
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_wfo_vanalytics_v2_filter_proto_init() }
func file_wfo_vanalytics_v2_filter_proto_init() {
	if File_wfo_vanalytics_v2_filter_proto != nil {
		return
	}
	file_wfo_vanalytics_v2_transcript_proto_init()
	file_wfo_vanalytics_v2_filter_proto_msgTypes[1].OneofWrappers = []any{
		(*ListFiltersRequest_Conflict)(nil),
		(*ListFiltersRequest_FlagSid)(nil),
	}
	file_wfo_vanalytics_v2_filter_proto_msgTypes[6].OneofWrappers = []any{
		(*GetFilterRequest_Name)(nil),
		(*GetFilterRequest_FilterSid)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wfo_vanalytics_v2_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wfo_vanalytics_v2_filter_proto_goTypes,
		DependencyIndexes: file_wfo_vanalytics_v2_filter_proto_depIdxs,
		MessageInfos:      file_wfo_vanalytics_v2_filter_proto_msgTypes,
	}.Build()
	File_wfo_vanalytics_v2_filter_proto = out.File
	file_wfo_vanalytics_v2_filter_proto_rawDesc = nil
	file_wfo_vanalytics_v2_filter_proto_goTypes = nil
	file_wfo_vanalytics_v2_filter_proto_depIdxs = nil
}
