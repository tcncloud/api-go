// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: api/v1alpha1/vanalytics/filter.proto

package vanalytics

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
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFilterRequest) ProtoMessage() {}

func (x *CreateFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_filter_proto_rawDescGZIP(), []int{0}
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
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFiltersRequest) ProtoMessage() {}

func (x *ListFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_filter_proto_rawDescGZIP(), []int{1}
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
	// Optional. The filter by which to find other filters which have conflicting
	// name or search request field(s).
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
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFiltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFiltersResponse) ProtoMessage() {}

func (x *ListFiltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_filter_proto_rawDescGZIP(), []int{2}
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
	// name and search_request. If no update_mask is provided it defaults to
	// using both name and search_request.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The ID of the filter to be updated.
	FilterSid int64 `protobuf:"varint,3,opt,name=filter_sid,json=filterSid,proto3" json:"filter_sid,omitempty"`
}

func (x *UpdateFilterRequest) Reset() {
	*x = UpdateFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFilterRequest) ProtoMessage() {}

func (x *UpdateFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_filter_proto_rawDescGZIP(), []int{3}
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
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFilterRequest) ProtoMessage() {}

func (x *DeleteFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_filter_proto_rawDescGZIP(), []int{4}
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
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFilterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFilterResponse) ProtoMessage() {}

func (x *DeleteFilterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_filter_proto_rawDescGZIP(), []int{5}
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
	//	*GetFilterRequest_SearchRequest
	//	*GetFilterRequest_Name
	//	*GetFilterRequest_FilterSid
	Where isGetFilterRequest_Where `protobuf_oneof:"where"`
}

func (x *GetFilterRequest) Reset() {
	*x = GetFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilterRequest) ProtoMessage() {}

func (x *GetFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_filter_proto_rawDescGZIP(), []int{6}
}

func (m *GetFilterRequest) GetWhere() isGetFilterRequest_Where {
	if m != nil {
		return m.Where
	}
	return nil
}

func (x *GetFilterRequest) GetSearchRequest() *SearchRequest {
	if x, ok := x.GetWhere().(*GetFilterRequest_SearchRequest); ok {
		return x.SearchRequest
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

type GetFilterRequest_SearchRequest struct {
	// The search request to be used to filter transcripts.
	SearchRequest *SearchRequest `protobuf:"bytes,2,opt,name=search_request,json=searchRequest,proto3,oneof"`
}

type GetFilterRequest_Name struct {
	// The name of the filter.
	Name string `protobuf:"bytes,3,opt,name=name,proto3,oneof"`
}

type GetFilterRequest_FilterSid struct {
	// The unique id of the filter.
	FilterSid int64 `protobuf:"varint,4,opt,name=filter_sid,json=filterSid,proto3,oneof"`
}

func (*GetFilterRequest_SearchRequest) isGetFilterRequest_Where() {}

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
	// The search request to be used to filter transcripts.
	SearchRequest *SearchRequest `protobuf:"bytes,4,opt,name=search_request,json=searchRequest,proto3" json:"search_request,omitempty"`
	// Output only. The timestamp when this filter was created. Assigned by the
	// server.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The version of this filter.
	Version int64 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_filter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_filter_proto_rawDescGZIP(), []int{7}
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

func (x *Filter) GetSearchRequest() *SearchRequest {
	if x != nil {
		return x.SearchRequest
	}
	return nil
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

var File_api_v1alpha1_vanalytics_filter_proto protoreflect.FileDescriptor

var file_api_v1alpha1_vanalytics_filter_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x1a,
	0x28, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xd0, 0x01, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3d, 0x0a, 0x08, 0x63, 0x6f,
	0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x08, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x6c, 0x61,
	0x67, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x07, 0x66,
	0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x77, 0x68, 0x65, 0x72, 0x65, 0x22,
	0x78, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x39,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x53, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x22, 0x4f, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xa3, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53,
	0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x77, 0x68, 0x65, 0x72, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x06,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x53, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0xdc, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x42,
	0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x56, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0xca, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0xe2, 0x02, 0x23, 0x41,
	0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x56, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3a, 0x3a, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_vanalytics_filter_proto_rawDescOnce sync.Once
	file_api_v1alpha1_vanalytics_filter_proto_rawDescData = file_api_v1alpha1_vanalytics_filter_proto_rawDesc
)

func file_api_v1alpha1_vanalytics_filter_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_vanalytics_filter_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_vanalytics_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_vanalytics_filter_proto_rawDescData)
	})
	return file_api_v1alpha1_vanalytics_filter_proto_rawDescData
}

var file_api_v1alpha1_vanalytics_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_v1alpha1_vanalytics_filter_proto_goTypes = []interface{}{
	(*CreateFilterRequest)(nil),   // 0: api.v1alpha1.vanalytics.CreateFilterRequest
	(*ListFiltersRequest)(nil),    // 1: api.v1alpha1.vanalytics.ListFiltersRequest
	(*ListFiltersResponse)(nil),   // 2: api.v1alpha1.vanalytics.ListFiltersResponse
	(*UpdateFilterRequest)(nil),   // 3: api.v1alpha1.vanalytics.UpdateFilterRequest
	(*DeleteFilterRequest)(nil),   // 4: api.v1alpha1.vanalytics.DeleteFilterRequest
	(*DeleteFilterResponse)(nil),  // 5: api.v1alpha1.vanalytics.DeleteFilterResponse
	(*GetFilterRequest)(nil),      // 6: api.v1alpha1.vanalytics.GetFilterRequest
	(*Filter)(nil),                // 7: api.v1alpha1.vanalytics.Filter
	(*fieldmaskpb.FieldMask)(nil), // 8: google.protobuf.FieldMask
	(*SearchRequest)(nil),         // 9: api.v1alpha1.vanalytics.SearchRequest
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_api_v1alpha1_vanalytics_filter_proto_depIdxs = []int32{
	7,  // 0: api.v1alpha1.vanalytics.CreateFilterRequest.filter:type_name -> api.v1alpha1.vanalytics.Filter
	7,  // 1: api.v1alpha1.vanalytics.ListFiltersRequest.conflict:type_name -> api.v1alpha1.vanalytics.Filter
	7,  // 2: api.v1alpha1.vanalytics.ListFiltersResponse.filters:type_name -> api.v1alpha1.vanalytics.Filter
	7,  // 3: api.v1alpha1.vanalytics.UpdateFilterRequest.filter:type_name -> api.v1alpha1.vanalytics.Filter
	8,  // 4: api.v1alpha1.vanalytics.UpdateFilterRequest.update_mask:type_name -> google.protobuf.FieldMask
	7,  // 5: api.v1alpha1.vanalytics.DeleteFilterResponse.filter:type_name -> api.v1alpha1.vanalytics.Filter
	9,  // 6: api.v1alpha1.vanalytics.GetFilterRequest.search_request:type_name -> api.v1alpha1.vanalytics.SearchRequest
	9,  // 7: api.v1alpha1.vanalytics.Filter.search_request:type_name -> api.v1alpha1.vanalytics.SearchRequest
	10, // 8: api.v1alpha1.vanalytics.Filter.create_time:type_name -> google.protobuf.Timestamp
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_vanalytics_filter_proto_init() }
func file_api_v1alpha1_vanalytics_filter_proto_init() {
	if File_api_v1alpha1_vanalytics_filter_proto != nil {
		return
	}
	file_api_v1alpha1_vanalytics_transcript_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1alpha1_vanalytics_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFilterRequest); i {
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
		file_api_v1alpha1_vanalytics_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFiltersRequest); i {
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
		file_api_v1alpha1_vanalytics_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFiltersResponse); i {
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
		file_api_v1alpha1_vanalytics_filter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFilterRequest); i {
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
		file_api_v1alpha1_vanalytics_filter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFilterRequest); i {
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
		file_api_v1alpha1_vanalytics_filter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFilterResponse); i {
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
		file_api_v1alpha1_vanalytics_filter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilterRequest); i {
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
		file_api_v1alpha1_vanalytics_filter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
	file_api_v1alpha1_vanalytics_filter_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ListFiltersRequest_Conflict)(nil),
		(*ListFiltersRequest_FlagSid)(nil),
	}
	file_api_v1alpha1_vanalytics_filter_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*GetFilterRequest_SearchRequest)(nil),
		(*GetFilterRequest_Name)(nil),
		(*GetFilterRequest_FilterSid)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_vanalytics_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_vanalytics_filter_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_vanalytics_filter_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_vanalytics_filter_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_vanalytics_filter_proto = out.File
	file_api_v1alpha1_vanalytics_filter_proto_rawDesc = nil
	file_api_v1alpha1_vanalytics_filter_proto_goTypes = nil
	file_api_v1alpha1_vanalytics_filter_proto_depIdxs = nil
}
