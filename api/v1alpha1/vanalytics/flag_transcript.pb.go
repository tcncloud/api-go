// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/vanalytics/flag_transcript.proto

package vanalytics

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

// FlagReviewStatus is an enum for the review status of a flagged transcript.
type FlagReviewStatus int32

const (
	FlagReviewStatus_ANY  FlagReviewStatus = 0
	FlagReviewStatus_TODO FlagReviewStatus = 1
	FlagReviewStatus_DONE FlagReviewStatus = 2
	FlagReviewStatus_NONE FlagReviewStatus = 3
)

// Enum value maps for FlagReviewStatus.
var (
	FlagReviewStatus_name = map[int32]string{
		0: "ANY",
		1: "TODO",
		2: "DONE",
		3: "NONE",
	}
	FlagReviewStatus_value = map[string]int32{
		"ANY":  0,
		"TODO": 1,
		"DONE": 2,
		"NONE": 3,
	}
)

func (x FlagReviewStatus) Enum() *FlagReviewStatus {
	p := new(FlagReviewStatus)
	*p = x
	return p
}

func (x FlagReviewStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlagReviewStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_enumTypes[0].Descriptor()
}

func (FlagReviewStatus) Type() protoreflect.EnumType {
	return &file_api_v1alpha1_vanalytics_flag_transcript_proto_enumTypes[0]
}

func (x FlagReviewStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FlagReviewStatus.Descriptor instead.
func (FlagReviewStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{0}
}

// CreateFlagTranscriptRequest is a request for creating a flag transcript.
type CreateFlagTranscriptRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The transcript sids to be flagged.
	TranscriptSids []int64 `protobuf:"varint,1,rep,packed,name=transcript_sids,json=transcriptSids,proto3" json:"transcript_sids,omitempty"`
	// Required. The flag for flagging the transcripts.
	Flag          *Flag `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateFlagTranscriptRequest) Reset() {
	*x = CreateFlagTranscriptRequest{}
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFlagTranscriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlagTranscriptRequest) ProtoMessage() {}

func (x *CreateFlagTranscriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlagTranscriptRequest.ProtoReflect.Descriptor instead.
func (*CreateFlagTranscriptRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFlagTranscriptRequest) GetTranscriptSids() []int64 {
	if x != nil {
		return x.TranscriptSids
	}
	return nil
}

func (x *CreateFlagTranscriptRequest) GetFlag() *Flag {
	if x != nil {
		return x.Flag
	}
	return nil
}

// CreateFlagTranscriptResponse is a response for creating a flag transcript.
type CreateFlagTranscriptResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateFlagTranscriptResponse) Reset() {
	*x = CreateFlagTranscriptResponse{}
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFlagTranscriptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlagTranscriptResponse) ProtoMessage() {}

func (x *CreateFlagTranscriptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlagTranscriptResponse.ProtoReflect.Descriptor instead.
func (*CreateFlagTranscriptResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{1}
}

// DeleteFlagTranscriptRequest is a request for deleting a flag transcript.
type DeleteFlagTranscriptRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The org id of the flag transcripts to be deleted.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Deprecated. Use string filter instead.
	//
	// Deprecated: Marked as deprecated in api/v1alpha1/vanalytics/flag_transcript.proto.
	TranscriptSids []int64 `protobuf:"varint,2,rep,packed,name=transcript_sids,json=transcriptSids,proto3" json:"transcript_sids,omitempty"`
	// Optional. The standard list filter as described in https://google.aip.dev/160.
	// Multiple comparisons can be provided when separated with a logical AND
	// operator. Supported fields, operators and functions are listed below.
	//
	// +----------------------------+-----------+--------------+-----------+
	// |                      field |      type |    operators | functions |
	// +----------------------------+-----------+--------------+-----------+
	// |             transcript_sid |   integer |            = |       any |
	// |                 start_time | timestamp | >=, <=, >, < | timestamp |
	// |         flag_summary.count |   integer | >=, <=, >, < |           |
	// | flag_summary.review_status |      enum |            = |       any |
	// +----------------------------+-----------+--------------+-----------+
	//
	// Examples:
	// transcript_sid = any(1, 2, 3)
	//
	// flag_summary.count > 0 AND
	// flag_summary.review_status = any(done, none) AND
	// start_time >= timestamp('2012-04-21T11:30:00-04:00')
	Filter        string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFlagTranscriptRequest) Reset() {
	*x = DeleteFlagTranscriptRequest{}
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFlagTranscriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlagTranscriptRequest) ProtoMessage() {}

func (x *DeleteFlagTranscriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlagTranscriptRequest.ProtoReflect.Descriptor instead.
func (*DeleteFlagTranscriptRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteFlagTranscriptRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1alpha1/vanalytics/flag_transcript.proto.
func (x *DeleteFlagTranscriptRequest) GetTranscriptSids() []int64 {
	if x != nil {
		return x.TranscriptSids
	}
	return nil
}

func (x *DeleteFlagTranscriptRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// DeleteFlagTranscriptResponse is a response for deleting a flag transcript.
type DeleteFlagTranscriptResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFlagTranscriptResponse) Reset() {
	*x = DeleteFlagTranscriptResponse{}
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFlagTranscriptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlagTranscriptResponse) ProtoMessage() {}

func (x *DeleteFlagTranscriptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlagTranscriptResponse.ProtoReflect.Descriptor instead.
func (*DeleteFlagTranscriptResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{3}
}

// SearchFlagTranscriptsRequest is a request for searching flagged transcripts.
type SearchFlagTranscriptsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The number of hits to include in a single response. When not
	// provided this defaults to 100.
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token returned from a previous List request, if any.
	// When provided all other request fields are ignored.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. Filter by flag sid.
	FlagSid *SearchFlagTranscriptsRequest_FlagSid `protobuf:"bytes,5,opt,name=flag_sid,json=flagSid,proto3" json:"flag_sid,omitempty"`
	// Optional. Filters transcripts by review status.
	FlagReviewStatus FlagReviewStatus `protobuf:"varint,6,opt,name=flag_review_status,json=flagReviewStatus,proto3,enum=api.v1alpha1.vanalytics.FlagReviewStatus" json:"flag_review_status,omitempty"`
	// Optional. Filter by notify group id.
	NotifyGroupId *SearchFlagTranscriptsRequest_NotifyGroupId `protobuf:"bytes,7,opt,name=notify_group_id,json=notifyGroupId,proto3" json:"notify_group_id,omitempty"`
	// Optional. Filter by review group id.
	ReviewGroupId *SearchFlagTranscriptsRequest_ReviewGroupId `protobuf:"bytes,8,opt,name=review_group_id,json=reviewGroupId,proto3" json:"review_group_id,omitempty"`
	// Optional. Filter where transript sid >= start transcript sid.
	StartTranscriptSid int64 `protobuf:"varint,10,opt,name=start_transcript_sid,json=startTranscriptSid,proto3" json:"start_transcript_sid,omitempty"`
	// Optional. Filter where transript sid <= end transcript sid.
	EndTranscriptSid int64 `protobuf:"varint,11,opt,name=end_transcript_sid,json=endTranscriptSid,proto3" json:"end_transcript_sid,omitempty"`
	// Optional. The order by which flag transcripts will be listed. Follows sql order by
	// syntax. When not provided the order defaults to "transcript_sid".
	// Supported order by includes:
	//   - (transcript_sid)
	//   - (transcript_sid desc)
	//   - (create_time, transcript_sid)
	//   - (create_time desc, transcript_sid desc)
	//   - (audio_time, transcript_sid)
	//   - (audio_time desc, transcript_sid desc)
	OrderBy       string `protobuf:"bytes,12,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchFlagTranscriptsRequest) Reset() {
	*x = SearchFlagTranscriptsRequest{}
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchFlagTranscriptsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsRequest) ProtoMessage() {}

func (x *SearchFlagTranscriptsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlagTranscriptsRequest.ProtoReflect.Descriptor instead.
func (*SearchFlagTranscriptsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{4}
}

func (x *SearchFlagTranscriptsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchFlagTranscriptsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *SearchFlagTranscriptsRequest) GetFlagSid() *SearchFlagTranscriptsRequest_FlagSid {
	if x != nil {
		return x.FlagSid
	}
	return nil
}

func (x *SearchFlagTranscriptsRequest) GetFlagReviewStatus() FlagReviewStatus {
	if x != nil {
		return x.FlagReviewStatus
	}
	return FlagReviewStatus_ANY
}

func (x *SearchFlagTranscriptsRequest) GetNotifyGroupId() *SearchFlagTranscriptsRequest_NotifyGroupId {
	if x != nil {
		return x.NotifyGroupId
	}
	return nil
}

func (x *SearchFlagTranscriptsRequest) GetReviewGroupId() *SearchFlagTranscriptsRequest_ReviewGroupId {
	if x != nil {
		return x.ReviewGroupId
	}
	return nil
}

func (x *SearchFlagTranscriptsRequest) GetStartTranscriptSid() int64 {
	if x != nil {
		return x.StartTranscriptSid
	}
	return 0
}

func (x *SearchFlagTranscriptsRequest) GetEndTranscriptSid() int64 {
	if x != nil {
		return x.EndTranscriptSid
	}
	return 0
}

func (x *SearchFlagTranscriptsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

// SearchFlagTranscriptsResponse is a response for searching flagged transcripts.
type SearchFlagTranscriptsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Token to retrieve the next page of transcripts, or empty if there are no
	// more transcripts in the list.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of hits.
	Hits []*SearchFlagTranscriptsResponse_Hit `protobuf:"bytes,2,rep,name=hits,proto3" json:"hits,omitempty"`
	// Total number of hits in query.
	Total         uint64 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchFlagTranscriptsResponse) Reset() {
	*x = SearchFlagTranscriptsResponse{}
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchFlagTranscriptsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsResponse) ProtoMessage() {}

func (x *SearchFlagTranscriptsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlagTranscriptsResponse.ProtoReflect.Descriptor instead.
func (*SearchFlagTranscriptsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{5}
}

func (x *SearchFlagTranscriptsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *SearchFlagTranscriptsResponse) GetHits() []*SearchFlagTranscriptsResponse_Hit {
	if x != nil {
		return x.Hits
	}
	return nil
}

func (x *SearchFlagTranscriptsResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// FlagSid filtering.
type SearchFlagTranscriptsRequest_FlagSid struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The flag sids to filter by.
	Filter []int64 `protobuf:"varint,1,rep,packed,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Minimum number of flags sids within the filter that must be
	// matched on each flagged transcript. All flag sids must be matched when
	// (match <= 0) or (match >= len(filter)).
	Match         int32 `protobuf:"varint,2,opt,name=match,proto3" json:"match,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchFlagTranscriptsRequest_FlagSid) Reset() {
	*x = SearchFlagTranscriptsRequest_FlagSid{}
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchFlagTranscriptsRequest_FlagSid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsRequest_FlagSid) ProtoMessage() {}

func (x *SearchFlagTranscriptsRequest_FlagSid) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlagTranscriptsRequest_FlagSid.ProtoReflect.Descriptor instead.
func (*SearchFlagTranscriptsRequest_FlagSid) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{4, 0}
}

func (x *SearchFlagTranscriptsRequest_FlagSid) GetFilter() []int64 {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *SearchFlagTranscriptsRequest_FlagSid) GetMatch() int32 {
	if x != nil {
		return x.Match
	}
	return 0
}

// NotifyGroupId filtering.
type SearchFlagTranscriptsRequest_NotifyGroupId struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. When true include, when false exclude, transcripts which are
	// flagged with an undefined notify group id.
	IsNull bool `protobuf:"varint,1,opt,name=is_null,json=isNull,proto3" json:"is_null,omitempty"`
	// Optional. The notify group ids to filter by.
	Filter []string `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Minimum number of notify group ids within the filter that must
	// be matched on each flagged transcript. All notify group ids must be
	// matched when (match <= 0) or (match >= len(filter)).
	Match         int32 `protobuf:"varint,3,opt,name=match,proto3" json:"match,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchFlagTranscriptsRequest_NotifyGroupId) Reset() {
	*x = SearchFlagTranscriptsRequest_NotifyGroupId{}
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchFlagTranscriptsRequest_NotifyGroupId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsRequest_NotifyGroupId) ProtoMessage() {}

func (x *SearchFlagTranscriptsRequest_NotifyGroupId) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlagTranscriptsRequest_NotifyGroupId.ProtoReflect.Descriptor instead.
func (*SearchFlagTranscriptsRequest_NotifyGroupId) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{4, 1}
}

func (x *SearchFlagTranscriptsRequest_NotifyGroupId) GetIsNull() bool {
	if x != nil {
		return x.IsNull
	}
	return false
}

func (x *SearchFlagTranscriptsRequest_NotifyGroupId) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *SearchFlagTranscriptsRequest_NotifyGroupId) GetMatch() int32 {
	if x != nil {
		return x.Match
	}
	return 0
}

// ReviewGroupId filtering.
type SearchFlagTranscriptsRequest_ReviewGroupId struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. When true include, when false exclude, transcripts which are
	// flagged with an undefined review group id.
	IsNull bool `protobuf:"varint,1,opt,name=is_null,json=isNull,proto3" json:"is_null,omitempty"`
	// Optional. The review group ids to filter by.
	Filter []string `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Minimum number of review group ids within the filter that must
	// be matched on each flagged transcript. All review group ids must be
	// matched when (match <= 0) or (match >= len(filter)).
	Match         int32 `protobuf:"varint,3,opt,name=match,proto3" json:"match,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchFlagTranscriptsRequest_ReviewGroupId) Reset() {
	*x = SearchFlagTranscriptsRequest_ReviewGroupId{}
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchFlagTranscriptsRequest_ReviewGroupId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsRequest_ReviewGroupId) ProtoMessage() {}

func (x *SearchFlagTranscriptsRequest_ReviewGroupId) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlagTranscriptsRequest_ReviewGroupId.ProtoReflect.Descriptor instead.
func (*SearchFlagTranscriptsRequest_ReviewGroupId) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{4, 2}
}

func (x *SearchFlagTranscriptsRequest_ReviewGroupId) GetIsNull() bool {
	if x != nil {
		return x.IsNull
	}
	return false
}

func (x *SearchFlagTranscriptsRequest_ReviewGroupId) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *SearchFlagTranscriptsRequest_ReviewGroupId) GetMatch() int32 {
	if x != nil {
		return x.Match
	}
	return 0
}

// Hits for searching flag transcripts.
type SearchFlagTranscriptsResponse_Hit struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Transcript that is flaggged.
	Transcript *Transcript `protobuf:"bytes,1,opt,name=transcript,proto3" json:"transcript,omitempty"`
	// List of flag_snapshot_sids that the transcript is flagged with.
	FlagSnapshotSids []int64 `protobuf:"varint,2,rep,packed,name=flag_snapshot_sids,json=flagSnapshotSids,proto3" json:"flag_snapshot_sids,omitempty"`
	// Map of review statuses. If it exists in the map and has value true it has
	// been reviewed. If it exists in the map and has value false it needs to be
	// reviewed. If it does not exist in the map then it does not require review.
	Reviewed map[int64]bool `protobuf:"bytes,3,rep,name=reviewed,proto3" json:"reviewed,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// List of flag_sids that the transcript is flagged with.
	FlagSids      []int64 `protobuf:"varint,4,rep,packed,name=flag_sids,json=flagSids,proto3" json:"flag_sids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchFlagTranscriptsResponse_Hit) Reset() {
	*x = SearchFlagTranscriptsResponse_Hit{}
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchFlagTranscriptsResponse_Hit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsResponse_Hit) ProtoMessage() {}

func (x *SearchFlagTranscriptsResponse_Hit) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlagTranscriptsResponse_Hit.ProtoReflect.Descriptor instead.
func (*SearchFlagTranscriptsResponse_Hit) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{5, 0}
}

func (x *SearchFlagTranscriptsResponse_Hit) GetTranscript() *Transcript {
	if x != nil {
		return x.Transcript
	}
	return nil
}

func (x *SearchFlagTranscriptsResponse_Hit) GetFlagSnapshotSids() []int64 {
	if x != nil {
		return x.FlagSnapshotSids
	}
	return nil
}

func (x *SearchFlagTranscriptsResponse_Hit) GetReviewed() map[int64]bool {
	if x != nil {
		return x.Reviewed
	}
	return nil
}

func (x *SearchFlagTranscriptsResponse_Hit) GetFlagSids() []int64 {
	if x != nil {
		return x.FlagSids
	}
	return nil
}

var File_api_v1alpha1_vanalytics_flag_transcript_proto protoreflect.FileDescriptor

const file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDesc = "" +
	"\n" +
	"-api/v1alpha1/vanalytics/flag_transcript.proto\x12\x17api.v1alpha1.vanalytics\x1a\"api/v1alpha1/vanalytics/flag.proto\x1a(api/v1alpha1/vanalytics/transcript.proto\"y\n" +
	"\x1bCreateFlagTranscriptRequest\x12'\n" +
	"\x0ftranscript_sids\x18\x01 \x03(\x03R\x0etranscriptSids\x121\n" +
	"\x04flag\x18\x02 \x01(\v2\x1d.api.v1alpha1.vanalytics.FlagR\x04flag\"\x1e\n" +
	"\x1cCreateFlagTranscriptResponse\"y\n" +
	"\x1bDeleteFlagTranscriptRequest\x12\x15\n" +
	"\x06org_id\x18\x01 \x01(\tR\x05orgId\x12+\n" +
	"\x0ftranscript_sids\x18\x02 \x03(\x03B\x02\x18\x01R\x0etranscriptSids\x12\x16\n" +
	"\x06filter\x18\x03 \x01(\tR\x06filter\"\x1e\n" +
	"\x1cDeleteFlagTranscriptResponse\"\xcb\x06\n" +
	"\x1cSearchFlagTranscriptsRequest\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\rR\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tR\tpageToken\x12X\n" +
	"\bflag_sid\x18\x05 \x01(\v2=.api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.FlagSidR\aflagSid\x12W\n" +
	"\x12flag_review_status\x18\x06 \x01(\x0e2).api.v1alpha1.vanalytics.FlagReviewStatusR\x10flagReviewStatus\x12k\n" +
	"\x0fnotify_group_id\x18\a \x01(\v2C.api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.NotifyGroupIdR\rnotifyGroupId\x12k\n" +
	"\x0freview_group_id\x18\b \x01(\v2C.api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.ReviewGroupIdR\rreviewGroupId\x120\n" +
	"\x14start_transcript_sid\x18\n" +
	" \x01(\x03R\x12startTranscriptSid\x12,\n" +
	"\x12end_transcript_sid\x18\v \x01(\x03R\x10endTranscriptSid\x12\x19\n" +
	"\border_by\x18\f \x01(\tR\aorderBy\x1a7\n" +
	"\aFlagSid\x12\x16\n" +
	"\x06filter\x18\x01 \x03(\x03R\x06filter\x12\x14\n" +
	"\x05match\x18\x02 \x01(\x05R\x05match\x1aV\n" +
	"\rNotifyGroupId\x12\x17\n" +
	"\ais_null\x18\x01 \x01(\bR\x06isNull\x12\x16\n" +
	"\x06filter\x18\x02 \x03(\tR\x06filter\x12\x14\n" +
	"\x05match\x18\x03 \x01(\x05R\x05match\x1aV\n" +
	"\rReviewGroupId\x12\x17\n" +
	"\ais_null\x18\x01 \x01(\bR\x06isNull\x12\x16\n" +
	"\x06filter\x18\x02 \x03(\tR\x06filter\x12\x14\n" +
	"\x05match\x18\x03 \x01(\x05R\x05match\"\xe8\x03\n" +
	"\x1dSearchFlagTranscriptsResponse\x12&\n" +
	"\x0fnext_page_token\x18\x01 \x01(\tR\rnextPageToken\x12N\n" +
	"\x04hits\x18\x02 \x03(\v2:.api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.HitR\x04hits\x12\x14\n" +
	"\x05total\x18\x03 \x01(\x04R\x05total\x1a\xb8\x02\n" +
	"\x03Hit\x12C\n" +
	"\n" +
	"transcript\x18\x01 \x01(\v2#.api.v1alpha1.vanalytics.TranscriptR\n" +
	"transcript\x12,\n" +
	"\x12flag_snapshot_sids\x18\x02 \x03(\x03R\x10flagSnapshotSids\x12d\n" +
	"\breviewed\x18\x03 \x03(\v2H.api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit.ReviewedEntryR\breviewed\x12\x1b\n" +
	"\tflag_sids\x18\x04 \x03(\x03R\bflagSids\x1a;\n" +
	"\rReviewedEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x03R\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\bR\x05value:\x028\x01*9\n" +
	"\x10FlagReviewStatus\x12\a\n" +
	"\x03ANY\x10\x00\x12\b\n" +
	"\x04TODO\x10\x01\x12\b\n" +
	"\x04DONE\x10\x02\x12\b\n" +
	"\x04NONE\x10\x03B\xe4\x01\n" +
	"\x1bcom.api.v1alpha1.vanalyticsB\x13FlagTranscriptProtoP\x01Z2github.com/tcncloud/api-go/api/v1alpha1/vanalytics\xa2\x02\x03AVV\xaa\x02\x17Api.V1alpha1.Vanalytics\xca\x02\x17Api\\V1alpha1\\Vanalytics\xe2\x02#Api\\V1alpha1\\Vanalytics\\GPBMetadata\xea\x02\x19Api::V1alpha1::Vanalyticsb\x06proto3"

var (
	file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescOnce sync.Once
	file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescData []byte
)

func file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDesc), len(file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDesc)))
	})
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescData
}

var file_api_v1alpha1_vanalytics_flag_transcript_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_v1alpha1_vanalytics_flag_transcript_proto_goTypes = []any{
	(FlagReviewStatus)(0),                              // 0: api.v1alpha1.vanalytics.FlagReviewStatus
	(*CreateFlagTranscriptRequest)(nil),                // 1: api.v1alpha1.vanalytics.CreateFlagTranscriptRequest
	(*CreateFlagTranscriptResponse)(nil),               // 2: api.v1alpha1.vanalytics.CreateFlagTranscriptResponse
	(*DeleteFlagTranscriptRequest)(nil),                // 3: api.v1alpha1.vanalytics.DeleteFlagTranscriptRequest
	(*DeleteFlagTranscriptResponse)(nil),               // 4: api.v1alpha1.vanalytics.DeleteFlagTranscriptResponse
	(*SearchFlagTranscriptsRequest)(nil),               // 5: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest
	(*SearchFlagTranscriptsResponse)(nil),              // 6: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse
	(*SearchFlagTranscriptsRequest_FlagSid)(nil),       // 7: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.FlagSid
	(*SearchFlagTranscriptsRequest_NotifyGroupId)(nil), // 8: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.NotifyGroupId
	(*SearchFlagTranscriptsRequest_ReviewGroupId)(nil), // 9: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.ReviewGroupId
	(*SearchFlagTranscriptsResponse_Hit)(nil),          // 10: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit
	nil,                // 11: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit.ReviewedEntry
	(*Flag)(nil),       // 12: api.v1alpha1.vanalytics.Flag
	(*Transcript)(nil), // 13: api.v1alpha1.vanalytics.Transcript
}
var file_api_v1alpha1_vanalytics_flag_transcript_proto_depIdxs = []int32{
	12, // 0: api.v1alpha1.vanalytics.CreateFlagTranscriptRequest.flag:type_name -> api.v1alpha1.vanalytics.Flag
	7,  // 1: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.flag_sid:type_name -> api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.FlagSid
	0,  // 2: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.flag_review_status:type_name -> api.v1alpha1.vanalytics.FlagReviewStatus
	8,  // 3: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.notify_group_id:type_name -> api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.NotifyGroupId
	9,  // 4: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.review_group_id:type_name -> api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.ReviewGroupId
	10, // 5: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.hits:type_name -> api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit
	13, // 6: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit.transcript:type_name -> api.v1alpha1.vanalytics.Transcript
	11, // 7: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit.reviewed:type_name -> api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit.ReviewedEntry
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_vanalytics_flag_transcript_proto_init() }
func file_api_v1alpha1_vanalytics_flag_transcript_proto_init() {
	if File_api_v1alpha1_vanalytics_flag_transcript_proto != nil {
		return
	}
	file_api_v1alpha1_vanalytics_flag_proto_init()
	file_api_v1alpha1_vanalytics_transcript_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDesc), len(file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_vanalytics_flag_transcript_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_vanalytics_flag_transcript_proto_depIdxs,
		EnumInfos:         file_api_v1alpha1_vanalytics_flag_transcript_proto_enumTypes,
		MessageInfos:      file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_vanalytics_flag_transcript_proto = out.File
	file_api_v1alpha1_vanalytics_flag_transcript_proto_goTypes = nil
	file_api_v1alpha1_vanalytics_flag_transcript_proto_depIdxs = nil
}
