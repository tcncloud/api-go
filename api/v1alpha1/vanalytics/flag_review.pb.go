// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/v1alpha1/vanalytics/flag_review.proto

package vanalytics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// CreateFlagReviewRequest is a request for creating a flag review.
type CreateFlagReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FlagReview resource to create.
	FlagReview *FlagReview `protobuf:"bytes,1,opt,name=flag_review,json=flagReview,proto3" json:"flag_review,omitempty"`
}

func (x *CreateFlagReviewRequest) Reset() {
	*x = CreateFlagReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFlagReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlagReviewRequest) ProtoMessage() {}

func (x *CreateFlagReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlagReviewRequest.ProtoReflect.Descriptor instead.
func (*CreateFlagReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_review_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFlagReviewRequest) GetFlagReview() *FlagReview {
	if x != nil {
		return x.FlagReview
	}
	return nil
}

// BulkCreateFlagReviewRequest is a request for creating many flag reviews.
type BulkCreateFlagReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The ID of the flag for the flag reviews. When not provided
	// all deleted flags will be given a review.
	FlagSid int64 `protobuf:"varint,1,opt,name=flag_sid,json=flagSid,proto3" json:"flag_sid,omitempty"`
	// Optional. Any notes written for the flag reviews.
	Notes string `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *BulkCreateFlagReviewRequest) Reset() {
	*x = BulkCreateFlagReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkCreateFlagReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkCreateFlagReviewRequest) ProtoMessage() {}

func (x *BulkCreateFlagReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkCreateFlagReviewRequest.ProtoReflect.Descriptor instead.
func (*BulkCreateFlagReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_review_proto_rawDescGZIP(), []int{1}
}

func (x *BulkCreateFlagReviewRequest) GetFlagSid() int64 {
	if x != nil {
		return x.FlagSid
	}
	return 0
}

func (x *BulkCreateFlagReviewRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

// BulkCreateFlagReviewResponse is a response for creating many flag reviews.
type BulkCreateFlagReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BulkCreateFlagReviewResponse) Reset() {
	*x = BulkCreateFlagReviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkCreateFlagReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkCreateFlagReviewResponse) ProtoMessage() {}

func (x *BulkCreateFlagReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkCreateFlagReviewResponse.ProtoReflect.Descriptor instead.
func (*BulkCreateFlagReviewResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_review_proto_rawDescGZIP(), []int{2}
}

// ListFlagReviewsRequest is a request for listing reviews on a transcript.
type ListFlagReviewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The number of reviews to include in a single response. When not
	// provided this defaults to 100.
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The order by which reviews will be listed. Follows sql order by
	// syntax. When not provided the order defaults to "flag_review_sid desc".
	// Supported order by includes:
	//   - (flag_review_sid desc)
	//   - (flag_review_sid)
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// Optional. The next_page_token returned from a previous List request, if any.
	// When provided all other request fields are ignored.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. The transcript sid to look for reviews on.
	TranscriptSid int64 `protobuf:"varint,5,opt,name=transcript_sid,json=transcriptSid,proto3" json:"transcript_sid,omitempty"`
}

func (x *ListFlagReviewsRequest) Reset() {
	*x = ListFlagReviewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlagReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlagReviewsRequest) ProtoMessage() {}

func (x *ListFlagReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlagReviewsRequest.ProtoReflect.Descriptor instead.
func (*ListFlagReviewsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_review_proto_rawDescGZIP(), []int{3}
}

func (x *ListFlagReviewsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListFlagReviewsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListFlagReviewsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListFlagReviewsRequest) GetTranscriptSid() int64 {
	if x != nil {
		return x.TranscriptSid
	}
	return 0
}

// ListFlagReviewsResponse is a response for listing reviews on a transcript.
type ListFlagReviewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to retrieve the next page of reviews, or empty if there are no
	// more reviews in the list.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of flag reviews.
	FlagReviews []*FlagReview `protobuf:"bytes,2,rep,name=flag_reviews,json=flagReviews,proto3" json:"flag_reviews,omitempty"`
}

func (x *ListFlagReviewsResponse) Reset() {
	*x = ListFlagReviewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlagReviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlagReviewsResponse) ProtoMessage() {}

func (x *ListFlagReviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlagReviewsResponse.ProtoReflect.Descriptor instead.
func (*ListFlagReviewsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_review_proto_rawDescGZIP(), []int{4}
}

func (x *ListFlagReviewsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListFlagReviewsResponse) GetFlagReviews() []*FlagReview {
	if x != nil {
		return x.FlagReviews
	}
	return nil
}

// FlagReview resource in the Vanalytics API. It represents a flag review
// for transcripts that have been flagged.
type FlagReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The unique id of this flag review.
	FlagReviewSid int64 `protobuf:"varint,1,opt,name=flag_review_sid,json=flagReviewSid,proto3" json:"flag_review_sid,omitempty"`
	// Required. The transcript_sid of the reviewed transcript.
	TranscriptSid int64 `protobuf:"varint,2,opt,name=transcript_sid,json=transcriptSid,proto3" json:"transcript_sid,omitempty"`
	// DEPRECATED. Use flag_snapshot_sid instead.
	//
	// Deprecated: Marked as deprecated in api/v1alpha1/vanalytics/flag_review.proto.
	FlagSid int64 `protobuf:"varint,3,opt,name=flag_sid,json=flagSid,proto3" json:"flag_sid,omitempty"`
	// Output only. The timestamp when this flag review was created. Assigned by the
	// server.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Optional. Any notes written for the flag review.
	Notes string `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	// The ID of the flag snapshot for this flag review.
	FlagSnapshotSid int64 `protobuf:"varint,8,opt,name=flag_snapshot_sid,json=flagSnapshotSid,proto3" json:"flag_snapshot_sid,omitempty"`
}

func (x *FlagReview) Reset() {
	*x = FlagReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagReview) ProtoMessage() {}

func (x *FlagReview) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagReview.ProtoReflect.Descriptor instead.
func (*FlagReview) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_flag_review_proto_rawDescGZIP(), []int{5}
}

func (x *FlagReview) GetFlagReviewSid() int64 {
	if x != nil {
		return x.FlagReviewSid
	}
	return 0
}

func (x *FlagReview) GetTranscriptSid() int64 {
	if x != nil {
		return x.TranscriptSid
	}
	return 0
}

// Deprecated: Marked as deprecated in api/v1alpha1/vanalytics/flag_review.proto.
func (x *FlagReview) GetFlagSid() int64 {
	if x != nil {
		return x.FlagSid
	}
	return 0
}

func (x *FlagReview) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *FlagReview) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *FlagReview) GetFlagSnapshotSid() int64 {
	if x != nil {
		return x.FlagSnapshotSid
	}
	return 0
}

var File_api_v1alpha1_vanalytics_flag_review_proto protoreflect.FileDescriptor

var file_api_v1alpha1_vanalytics_flag_review_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x44, 0x0a, 0x0b, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x0a, 0x66, 0x6c, 0x61, 0x67,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x4e, 0x0a, 0x1b, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x1e, 0x0a, 0x1c, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x64, 0x22,
	0x89, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x46, 0x0a, 0x0c, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x0b,
	0x66, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22, 0xf9, 0x01, 0x0a, 0x0a,
	0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x6c,
	0x61, 0x67, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53,
	0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x08, 0x66, 0x6c, 0x61,
	0x67, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x07, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x66,
	0x6c, 0x61, 0x67, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x73, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x53, 0x69, 0x64, 0x42, 0xe0, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x42, 0x0f, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
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
	file_api_v1alpha1_vanalytics_flag_review_proto_rawDescOnce sync.Once
	file_api_v1alpha1_vanalytics_flag_review_proto_rawDescData = file_api_v1alpha1_vanalytics_flag_review_proto_rawDesc
)

func file_api_v1alpha1_vanalytics_flag_review_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_vanalytics_flag_review_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_vanalytics_flag_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_vanalytics_flag_review_proto_rawDescData)
	})
	return file_api_v1alpha1_vanalytics_flag_review_proto_rawDescData
}

var file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1alpha1_vanalytics_flag_review_proto_goTypes = []any{
	(*CreateFlagReviewRequest)(nil),      // 0: api.v1alpha1.vanalytics.CreateFlagReviewRequest
	(*BulkCreateFlagReviewRequest)(nil),  // 1: api.v1alpha1.vanalytics.BulkCreateFlagReviewRequest
	(*BulkCreateFlagReviewResponse)(nil), // 2: api.v1alpha1.vanalytics.BulkCreateFlagReviewResponse
	(*ListFlagReviewsRequest)(nil),       // 3: api.v1alpha1.vanalytics.ListFlagReviewsRequest
	(*ListFlagReviewsResponse)(nil),      // 4: api.v1alpha1.vanalytics.ListFlagReviewsResponse
	(*FlagReview)(nil),                   // 5: api.v1alpha1.vanalytics.FlagReview
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
}
var file_api_v1alpha1_vanalytics_flag_review_proto_depIdxs = []int32{
	5, // 0: api.v1alpha1.vanalytics.CreateFlagReviewRequest.flag_review:type_name -> api.v1alpha1.vanalytics.FlagReview
	5, // 1: api.v1alpha1.vanalytics.ListFlagReviewsResponse.flag_reviews:type_name -> api.v1alpha1.vanalytics.FlagReview
	6, // 2: api.v1alpha1.vanalytics.FlagReview.create_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_vanalytics_flag_review_proto_init() }
func file_api_v1alpha1_vanalytics_flag_review_proto_init() {
	if File_api_v1alpha1_vanalytics_flag_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFlagReviewRequest); i {
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
		file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BulkCreateFlagReviewRequest); i {
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
		file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BulkCreateFlagReviewResponse); i {
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
		file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListFlagReviewsRequest); i {
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
		file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListFlagReviewsResponse); i {
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
		file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*FlagReview); i {
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
			RawDescriptor: file_api_v1alpha1_vanalytics_flag_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_vanalytics_flag_review_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_vanalytics_flag_review_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_vanalytics_flag_review_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_vanalytics_flag_review_proto = out.File
	file_api_v1alpha1_vanalytics_flag_review_proto_rawDesc = nil
	file_api_v1alpha1_vanalytics_flag_review_proto_goTypes = nil
	file_api_v1alpha1_vanalytics_flag_review_proto_depIdxs = nil
}
