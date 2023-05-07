// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/v1alpha1/vanalytics/flag_transcript.proto

package vanalytics

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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The transcript sids to be flagged.
	TranscriptSids []int64 `protobuf:"varint,1,rep,packed,name=transcript_sids,json=transcriptSids,proto3" json:"transcript_sids,omitempty"`
	// Required. The flag for flagging the transcripts.
	Flag *Flag `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *CreateFlagTranscriptRequest) Reset() {
	*x = CreateFlagTranscriptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFlagTranscriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlagTranscriptRequest) ProtoMessage() {}

func (x *CreateFlagTranscriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateFlagTranscriptResponse) Reset() {
	*x = CreateFlagTranscriptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFlagTranscriptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlagTranscriptResponse) ProtoMessage() {}

func (x *CreateFlagTranscriptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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

// SearchFlagTranscriptsRequest is a request for searching flagged transcripts.
type SearchFlagTranscriptsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	OrderBy string `protobuf:"bytes,12,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (x *SearchFlagTranscriptsRequest) Reset() {
	*x = SearchFlagTranscriptsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlagTranscriptsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsRequest) ProtoMessage() {}

func (x *SearchFlagTranscriptsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{2}
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to retrieve the next page of transcripts, or empty if there are no
	// more transcripts in the list.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of hits.
	Hits []*SearchFlagTranscriptsResponse_Hit `protobuf:"bytes,2,rep,name=hits,proto3" json:"hits,omitempty"`
	// Total number of hits in query.
	Total uint64 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *SearchFlagTranscriptsResponse) Reset() {
	*x = SearchFlagTranscriptsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlagTranscriptsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsResponse) ProtoMessage() {}

func (x *SearchFlagTranscriptsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{3}
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The flag sids to filter by.
	Filter []int64 `protobuf:"varint,1,rep,packed,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Minimum number of flags sids within the filter that must be
	// matched on each flagged transcript. All flag sids must be matched when
	// (match <= 0) or (match >= len(filter)).
	Match int32 `protobuf:"varint,2,opt,name=match,proto3" json:"match,omitempty"`
}

func (x *SearchFlagTranscriptsRequest_FlagSid) Reset() {
	*x = SearchFlagTranscriptsRequest_FlagSid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlagTranscriptsRequest_FlagSid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsRequest_FlagSid) ProtoMessage() {}

func (x *SearchFlagTranscriptsRequest_FlagSid) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{2, 0}
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. When true include, when false exclude, transcripts which are
	// flagged with an undefined notify group id.
	IsNull bool `protobuf:"varint,1,opt,name=is_null,json=isNull,proto3" json:"is_null,omitempty"`
	// Optional. The notify group ids to filter by.
	Filter []string `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Minimum number of notify group ids within the filter that must
	// be matched on each flagged transcript. All notify group ids must be
	// matched when (match <= 0) or (match >= len(filter)).
	Match int32 `protobuf:"varint,3,opt,name=match,proto3" json:"match,omitempty"`
}

func (x *SearchFlagTranscriptsRequest_NotifyGroupId) Reset() {
	*x = SearchFlagTranscriptsRequest_NotifyGroupId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlagTranscriptsRequest_NotifyGroupId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsRequest_NotifyGroupId) ProtoMessage() {}

func (x *SearchFlagTranscriptsRequest_NotifyGroupId) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{2, 1}
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. When true include, when false exclude, transcripts which are
	// flagged with an undefined review group id.
	IsNull bool `protobuf:"varint,1,opt,name=is_null,json=isNull,proto3" json:"is_null,omitempty"`
	// Optional. The review group ids to filter by.
	Filter []string `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Minimum number of review group ids within the filter that must
	// be matched on each flagged transcript. All review group ids must be
	// matched when (match <= 0) or (match >= len(filter)).
	Match int32 `protobuf:"varint,3,opt,name=match,proto3" json:"match,omitempty"`
}

func (x *SearchFlagTranscriptsRequest_ReviewGroupId) Reset() {
	*x = SearchFlagTranscriptsRequest_ReviewGroupId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlagTranscriptsRequest_ReviewGroupId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsRequest_ReviewGroupId) ProtoMessage() {}

func (x *SearchFlagTranscriptsRequest_ReviewGroupId) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{2, 2}
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transcript that is flaggged.
	Transcript *Transcript `protobuf:"bytes,1,opt,name=transcript,proto3" json:"transcript,omitempty"`
	// List of flag_snapshot_sids that the transcript is flagged with.
	FlagSnapshotSids []int64 `protobuf:"varint,2,rep,packed,name=flag_snapshot_sids,json=flagSnapshotSids,proto3" json:"flag_snapshot_sids,omitempty"`
	// Map of review statuses. If it exists in the map and has value true it has
	// been reviewed. If it exists in the map and has value false it needs to be
	// reviewed. If it does not exist in the map then it does not require review.
	Reviewed map[int64]bool `protobuf:"bytes,3,rep,name=reviewed,proto3" json:"reviewed,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// List of flag_sids that the transcript is flagged with.
	FlagSids []int64 `protobuf:"varint,4,rep,packed,name=flag_sids,json=flagSids,proto3" json:"flag_sids,omitempty"`
}

func (x *SearchFlagTranscriptsResponse_Hit) Reset() {
	*x = SearchFlagTranscriptsResponse_Hit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlagTranscriptsResponse_Hit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlagTranscriptsResponse_Hit) ProtoMessage() {}

func (x *SearchFlagTranscriptsResponse_Hit) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP(), []int{3, 0}
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

var file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x6c, 0x61, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x64, 0x73, 0x12, 0x31,
	0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x04, 0x66, 0x6c, 0x61,
	0x67, 0x22, 0x1e, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xcb, 0x06, 0x0a, 0x1c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x61, 0x67,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x58,
	0x0a, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x46, 0x6c, 0x61, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x52,
	0x07, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x12, 0x57, 0x0a, 0x12, 0x66, 0x6c, 0x61, 0x67,
	0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x46,
	0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x10, 0x66, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x6b, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x61, 0x67, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52,
	0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x6b,
	0x0a, 0x0f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x61, 0x67, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x0d, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f,
	0x73, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x64, 0x12, 0x2c, 0x0a,
	0x12, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f,
	0x73, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x65, 0x6e, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x1a, 0x37, 0x0a, 0x07, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x1a,
	0x56, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x56, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6e,
	0x75, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4e, 0x75, 0x6c,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x22,
	0xe8, 0x03, 0x0a, 0x1d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x61, 0x67, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4e, 0x0a, 0x04, 0x68, 0x69, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x61, 0x67, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x48, 0x69, 0x74, 0x52, 0x04, 0x68, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x1a,
	0xb8, 0x02, 0x0a, 0x03, 0x48, 0x69, 0x74, 0x12, 0x43, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x2c, 0x0a, 0x12,
	0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x73, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x10, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x69, 0x64, 0x73, 0x12, 0x64, 0x0a, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x61,
	0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x73, 0x1a, 0x3b, 0x0a,
	0x0d, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x39, 0x0a, 0x10, 0x46, 0x6c,
	0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x4e, 0x59, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x4f, 0x44, 0x4f, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x03, 0x42, 0x63, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x63, 0x6e,
	0x2e, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5a,
	0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x3b,
	0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescOnce sync.Once
	file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescData = file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDesc
)

func file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescData)
	})
	return file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDescData
}

var file_api_v1alpha1_vanalytics_flag_transcript_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_v1alpha1_vanalytics_flag_transcript_proto_goTypes = []interface{}{
	(FlagReviewStatus)(0),                              // 0: api.v1alpha1.vanalytics.FlagReviewStatus
	(*CreateFlagTranscriptRequest)(nil),                // 1: api.v1alpha1.vanalytics.CreateFlagTranscriptRequest
	(*CreateFlagTranscriptResponse)(nil),               // 2: api.v1alpha1.vanalytics.CreateFlagTranscriptResponse
	(*SearchFlagTranscriptsRequest)(nil),               // 3: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest
	(*SearchFlagTranscriptsResponse)(nil),              // 4: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse
	(*SearchFlagTranscriptsRequest_FlagSid)(nil),       // 5: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.FlagSid
	(*SearchFlagTranscriptsRequest_NotifyGroupId)(nil), // 6: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.NotifyGroupId
	(*SearchFlagTranscriptsRequest_ReviewGroupId)(nil), // 7: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.ReviewGroupId
	(*SearchFlagTranscriptsResponse_Hit)(nil),          // 8: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit
	nil,                // 9: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit.ReviewedEntry
	(*Flag)(nil),       // 10: api.v1alpha1.vanalytics.Flag
	(*Transcript)(nil), // 11: api.v1alpha1.vanalytics.Transcript
}
var file_api_v1alpha1_vanalytics_flag_transcript_proto_depIdxs = []int32{
	10, // 0: api.v1alpha1.vanalytics.CreateFlagTranscriptRequest.flag:type_name -> api.v1alpha1.vanalytics.Flag
	5,  // 1: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.flag_sid:type_name -> api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.FlagSid
	0,  // 2: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.flag_review_status:type_name -> api.v1alpha1.vanalytics.FlagReviewStatus
	6,  // 3: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.notify_group_id:type_name -> api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.NotifyGroupId
	7,  // 4: api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.review_group_id:type_name -> api.v1alpha1.vanalytics.SearchFlagTranscriptsRequest.ReviewGroupId
	8,  // 5: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.hits:type_name -> api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit
	11, // 6: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit.transcript:type_name -> api.v1alpha1.vanalytics.Transcript
	9,  // 7: api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit.reviewed:type_name -> api.v1alpha1.vanalytics.SearchFlagTranscriptsResponse.Hit.ReviewedEntry
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
	if !protoimpl.UnsafeEnabled {
		file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFlagTranscriptRequest); i {
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
		file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFlagTranscriptResponse); i {
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
		file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlagTranscriptsRequest); i {
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
		file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlagTranscriptsResponse); i {
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
		file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlagTranscriptsRequest_FlagSid); i {
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
		file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlagTranscriptsRequest_NotifyGroupId); i {
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
		file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlagTranscriptsRequest_ReviewGroupId); i {
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
		file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlagTranscriptsResponse_Hit); i {
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
			RawDescriptor: file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_vanalytics_flag_transcript_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_vanalytics_flag_transcript_proto_depIdxs,
		EnumInfos:         file_api_v1alpha1_vanalytics_flag_transcript_proto_enumTypes,
		MessageInfos:      file_api_v1alpha1_vanalytics_flag_transcript_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_vanalytics_flag_transcript_proto = out.File
	file_api_v1alpha1_vanalytics_flag_transcript_proto_rawDesc = nil
	file_api_v1alpha1_vanalytics_flag_transcript_proto_goTypes = nil
	file_api_v1alpha1_vanalytics_flag_transcript_proto_depIdxs = nil
}
