// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: wfo/vanalytics/v2/flag_snapshot.proto

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

// ListFlagSnapshotsRequest is a request for listing flag snapshots.
type ListFlagSnapshotsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The number of snapshots to include in a single response. When not
	// provided this defaults to 100.
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The order by which snapshots will be listed. Follows sql order by
	// syntax. When not provided the order defaults to "flag_snapshot_sid desc".
	// Supported order by includes:
	//   - (flag_snapshot_sid desc)
	//   - (flag_snapshot_sid)
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// Optional. The next_page_token returned from a previous List request, if any.
	// When provided all other request fields are ignored.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. List of flag_snapshots to filter on.
	FlagSnapshotSids []int64 `protobuf:"varint,5,rep,packed,name=flag_snapshot_sids,json=flagSnapshotSids,proto3" json:"flag_snapshot_sids,omitempty"`
	// Optional. mask contains a list of fields to be returned. Possible paths include
	// flag_snapshot_sid, flag_sid, name, review_group_id, notify_group_id, create_time,
	// version, and priority. If no mask is provided it defaults to using all paths.
	Mask *fieldmaskpb.FieldMask `protobuf:"bytes,6,opt,name=mask,proto3" json:"mask,omitempty"`
	// Optional. Transcript sid to filter on.
	TranscriptSid int64 `protobuf:"varint,7,opt,name=transcript_sid,json=transcriptSid,proto3" json:"transcript_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFlagSnapshotsRequest) Reset() {
	*x = ListFlagSnapshotsRequest{}
	mi := &file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFlagSnapshotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlagSnapshotsRequest) ProtoMessage() {}

func (x *ListFlagSnapshotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlagSnapshotsRequest.ProtoReflect.Descriptor instead.
func (*ListFlagSnapshotsRequest) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *ListFlagSnapshotsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListFlagSnapshotsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListFlagSnapshotsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListFlagSnapshotsRequest) GetFlagSnapshotSids() []int64 {
	if x != nil {
		return x.FlagSnapshotSids
	}
	return nil
}

func (x *ListFlagSnapshotsRequest) GetMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Mask
	}
	return nil
}

func (x *ListFlagSnapshotsRequest) GetTranscriptSid() int64 {
	if x != nil {
		return x.TranscriptSid
	}
	return 0
}

// ListFlagSnapshotsResponse is a response for listing flag snapshots.
type ListFlagSnapshotsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Token to retrieve the next page of snapshots, or empty if there are no
	// more snapshots in the list.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of flag snapshots.
	FlagSnapshots []*FlagSnapshot `protobuf:"bytes,2,rep,name=flag_snapshots,json=flagSnapshots,proto3" json:"flag_snapshots,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFlagSnapshotsResponse) Reset() {
	*x = ListFlagSnapshotsResponse{}
	mi := &file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFlagSnapshotsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlagSnapshotsResponse) ProtoMessage() {}

func (x *ListFlagSnapshotsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlagSnapshotsResponse.ProtoReflect.Descriptor instead.
func (*ListFlagSnapshotsResponse) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *ListFlagSnapshotsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListFlagSnapshotsResponse) GetFlagSnapshots() []*FlagSnapshot {
	if x != nil {
		return x.FlagSnapshots
	}
	return nil
}

// FlagSnapshot is a resource in the Vanalytics API.
type FlagSnapshot struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The unique id of this flag.
	FlagSnapshotSid int64 `protobuf:"varint,1,opt,name=flag_snapshot_sid,json=flagSnapshotSid,proto3" json:"flag_snapshot_sid,omitempty"`
	// Output only. The unique id of this flag.
	FlagSid int64 `protobuf:"varint,2,opt,name=flag_sid,json=flagSid,proto3" json:"flag_sid,omitempty"`
	// Required. The name of this flag. Must be non empty and unique across all
	// flags within an organization.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The org permission group id which users must have in order to
	// to review flagged transcripts. When not provided flagged transcripts will
	// not require review.
	ReviewGroupId string `protobuf:"bytes,5,opt,name=review_group_id,json=reviewGroupId,proto3" json:"review_group_id,omitempty"`
	// Optional. The notify group id for this flag.
	NotifyGroupId string `protobuf:"bytes,6,opt,name=notify_group_id,json=notifyGroupId,proto3" json:"notify_group_id,omitempty"`
	// Optional. The priority of the flag. Defaults to 0.
	Priority int32 `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	// Output only. The version of this flag.
	Version int64 `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
	// Output only. The timestamp when this flag snapshot was created. Assigned by the
	// server.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Determines whether this flag must be reviewed.
	MustReview bool `protobuf:"varint,10,opt,name=must_review,json=mustReview,proto3" json:"must_review,omitempty"`
	// Output only. Determines whether this flag must be notified.
	MustNotify bool `protobuf:"varint,11,opt,name=must_notify,json=mustNotify,proto3" json:"must_notify,omitempty"`
	// Required. Boolean expression of filters which a transcript must match
	// for this flag to be applied.
	BoolExpr *FlagSnapshot_BoolExpr `protobuf:"bytes,12,opt,name=bool_expr,json=boolExpr,proto3" json:"bool_expr,omitempty"`
	// Optional. Specifies dncl lists to update
	// if a transcript is flagged.
	DnclList      []*DnclList `protobuf:"bytes,13,rep,name=dncl_list,json=dnclList,proto3" json:"dncl_list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlagSnapshot) Reset() {
	*x = FlagSnapshot{}
	mi := &file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlagSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagSnapshot) ProtoMessage() {}

func (x *FlagSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagSnapshot.ProtoReflect.Descriptor instead.
func (*FlagSnapshot) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescGZIP(), []int{2}
}

func (x *FlagSnapshot) GetFlagSnapshotSid() int64 {
	if x != nil {
		return x.FlagSnapshotSid
	}
	return 0
}

func (x *FlagSnapshot) GetFlagSid() int64 {
	if x != nil {
		return x.FlagSid
	}
	return 0
}

func (x *FlagSnapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlagSnapshot) GetReviewGroupId() string {
	if x != nil {
		return x.ReviewGroupId
	}
	return ""
}

func (x *FlagSnapshot) GetNotifyGroupId() string {
	if x != nil {
		return x.NotifyGroupId
	}
	return ""
}

func (x *FlagSnapshot) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *FlagSnapshot) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *FlagSnapshot) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *FlagSnapshot) GetMustReview() bool {
	if x != nil {
		return x.MustReview
	}
	return false
}

func (x *FlagSnapshot) GetMustNotify() bool {
	if x != nil {
		return x.MustNotify
	}
	return false
}

func (x *FlagSnapshot) GetBoolExpr() *FlagSnapshot_BoolExpr {
	if x != nil {
		return x.BoolExpr
	}
	return nil
}

func (x *FlagSnapshot) GetDnclList() []*DnclList {
	if x != nil {
		return x.DnclList
	}
	return nil
}

// BoolExpr defines a boolean expression of filters.
type FlagSnapshot_BoolExpr struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. Boolean and operator.
	And []*FlagSnapshot_BoolExpr `protobuf:"bytes,1,rep,name=and,proto3" json:"and,omitempty"`
	// Optional. Boolean or operator.
	Or []*FlagSnapshot_BoolExpr `protobuf:"bytes,2,rep,name=or,proto3" json:"or,omitempty"`
	// Optional. Filter to match.
	Filter *FlagSnapshot_BoolExpr_Filter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Boolean not operator.
	Not           *FlagSnapshot_BoolExpr `protobuf:"bytes,4,opt,name=not,proto3" json:"not,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlagSnapshot_BoolExpr) Reset() {
	*x = FlagSnapshot_BoolExpr{}
	mi := &file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlagSnapshot_BoolExpr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagSnapshot_BoolExpr) ProtoMessage() {}

func (x *FlagSnapshot_BoolExpr) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagSnapshot_BoolExpr.ProtoReflect.Descriptor instead.
func (*FlagSnapshot_BoolExpr) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescGZIP(), []int{2, 0}
}

func (x *FlagSnapshot_BoolExpr) GetAnd() []*FlagSnapshot_BoolExpr {
	if x != nil {
		return x.And
	}
	return nil
}

func (x *FlagSnapshot_BoolExpr) GetOr() []*FlagSnapshot_BoolExpr {
	if x != nil {
		return x.Or
	}
	return nil
}

func (x *FlagSnapshot_BoolExpr) GetFilter() *FlagSnapshot_BoolExpr_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *FlagSnapshot_BoolExpr) GetNot() *FlagSnapshot_BoolExpr {
	if x != nil {
		return x.Not
	}
	return nil
}

// Filter defines a filter.
type FlagSnapshot_BoolExpr_Filter struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. Filter sid.
	FilterSid     int64 `protobuf:"varint,1,opt,name=filter_sid,json=filterSid,proto3" json:"filter_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlagSnapshot_BoolExpr_Filter) Reset() {
	*x = FlagSnapshot_BoolExpr_Filter{}
	mi := &file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlagSnapshot_BoolExpr_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagSnapshot_BoolExpr_Filter) ProtoMessage() {}

func (x *FlagSnapshot_BoolExpr_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagSnapshot_BoolExpr_Filter.ProtoReflect.Descriptor instead.
func (*FlagSnapshot_BoolExpr_Filter) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *FlagSnapshot_BoolExpr_Filter) GetFilterSid() int64 {
	if x != nil {
		return x.FilterSid
	}
	return 0
}

var File_wfo_vanalytics_v2_flag_snapshot_proto protoreflect.FileDescriptor

var file_wfo_vanalytics_v2_flag_snapshot_proto_rawDesc = []byte{
	0x0a, 0x25, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x77,
	0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x64, 0x6e, 0x63, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf6, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x10, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x69,
	0x64, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x6d, 0x61,
	0x73, 0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x5f, 0x73, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x46, 0x0a, 0x0e, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x6c, 0x61, 0x67,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x0d, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x22, 0xa0, 0x06, 0x0a, 0x0c, 0x46, 0x6c, 0x61, 0x67,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x6c, 0x61, 0x67,
	0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x53, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6d, 0x75, 0x73,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x75, 0x73, 0x74, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6d, 0x75,
	0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x45, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6c,
	0x5f, 0x65, 0x78, 0x70, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x77, 0x66,
	0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x46, 0x6c, 0x61, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x45, 0x78, 0x70, 0x72, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6c, 0x45, 0x78, 0x70, 0x72, 0x12,
	0x38, 0x0a, 0x09, 0x64, 0x6e, 0x63, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x6e, 0x63, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x08, 0x64, 0x6e, 0x63, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0xae, 0x02, 0x0a, 0x08, 0x42, 0x6f,
	0x6f, 0x6c, 0x45, 0x78, 0x70, 0x72, 0x12, 0x3a, 0x0a, 0x03, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x45, 0x78, 0x70, 0x72, 0x52, 0x03, 0x61,
	0x6e, 0x64, 0x12, 0x38, 0x0a, 0x02, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x45, 0x78, 0x70, 0x72, 0x52, 0x02, 0x6f, 0x72, 0x12, 0x47, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x77,
	0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x45, 0x78, 0x70, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x03, 0x6e, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x45, 0x78, 0x70, 0x72, 0x52, 0x03, 0x6e, 0x6f,
	0x74, 0x1a, 0x27, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x69, 0x64, 0x42, 0xcb, 0x01, 0x0a, 0x15, 0x63,
	0x6f, 0x6d, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x32, 0x42, 0x11, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x57, 0x56, 0x58, 0xaa, 0x02, 0x11, 0x57, 0x66, 0x6f,
	0x2e, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x56, 0x32, 0xca, 0x02,
	0x11, 0x57, 0x66, 0x6f, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c,
	0x56, 0x32, 0xe2, 0x02, 0x1d, 0x57, 0x66, 0x6f, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x13, 0x57, 0x66, 0x6f, 0x3a, 0x3a, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescOnce sync.Once
	file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescData = file_wfo_vanalytics_v2_flag_snapshot_proto_rawDesc
)

func file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescGZIP() []byte {
	file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescOnce.Do(func() {
		file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescData)
	})
	return file_wfo_vanalytics_v2_flag_snapshot_proto_rawDescData
}

var file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wfo_vanalytics_v2_flag_snapshot_proto_goTypes = []any{
	(*ListFlagSnapshotsRequest)(nil),     // 0: wfo.vanalytics.v2.ListFlagSnapshotsRequest
	(*ListFlagSnapshotsResponse)(nil),    // 1: wfo.vanalytics.v2.ListFlagSnapshotsResponse
	(*FlagSnapshot)(nil),                 // 2: wfo.vanalytics.v2.FlagSnapshot
	(*FlagSnapshot_BoolExpr)(nil),        // 3: wfo.vanalytics.v2.FlagSnapshot.BoolExpr
	(*FlagSnapshot_BoolExpr_Filter)(nil), // 4: wfo.vanalytics.v2.FlagSnapshot.BoolExpr.Filter
	(*fieldmaskpb.FieldMask)(nil),        // 5: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
	(*DnclList)(nil),                     // 7: wfo.vanalytics.v2.DnclList
}
var file_wfo_vanalytics_v2_flag_snapshot_proto_depIdxs = []int32{
	5, // 0: wfo.vanalytics.v2.ListFlagSnapshotsRequest.mask:type_name -> google.protobuf.FieldMask
	2, // 1: wfo.vanalytics.v2.ListFlagSnapshotsResponse.flag_snapshots:type_name -> wfo.vanalytics.v2.FlagSnapshot
	6, // 2: wfo.vanalytics.v2.FlagSnapshot.create_time:type_name -> google.protobuf.Timestamp
	3, // 3: wfo.vanalytics.v2.FlagSnapshot.bool_expr:type_name -> wfo.vanalytics.v2.FlagSnapshot.BoolExpr
	7, // 4: wfo.vanalytics.v2.FlagSnapshot.dncl_list:type_name -> wfo.vanalytics.v2.DnclList
	3, // 5: wfo.vanalytics.v2.FlagSnapshot.BoolExpr.and:type_name -> wfo.vanalytics.v2.FlagSnapshot.BoolExpr
	3, // 6: wfo.vanalytics.v2.FlagSnapshot.BoolExpr.or:type_name -> wfo.vanalytics.v2.FlagSnapshot.BoolExpr
	4, // 7: wfo.vanalytics.v2.FlagSnapshot.BoolExpr.filter:type_name -> wfo.vanalytics.v2.FlagSnapshot.BoolExpr.Filter
	3, // 8: wfo.vanalytics.v2.FlagSnapshot.BoolExpr.not:type_name -> wfo.vanalytics.v2.FlagSnapshot.BoolExpr
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_wfo_vanalytics_v2_flag_snapshot_proto_init() }
func file_wfo_vanalytics_v2_flag_snapshot_proto_init() {
	if File_wfo_vanalytics_v2_flag_snapshot_proto != nil {
		return
	}
	file_wfo_vanalytics_v2_dncl_list_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wfo_vanalytics_v2_flag_snapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wfo_vanalytics_v2_flag_snapshot_proto_goTypes,
		DependencyIndexes: file_wfo_vanalytics_v2_flag_snapshot_proto_depIdxs,
		MessageInfos:      file_wfo_vanalytics_v2_flag_snapshot_proto_msgTypes,
	}.Build()
	File_wfo_vanalytics_v2_flag_snapshot_proto = out.File
	file_wfo_vanalytics_v2_flag_snapshot_proto_rawDesc = nil
	file_wfo_vanalytics_v2_flag_snapshot_proto_goTypes = nil
	file_wfo_vanalytics_v2_flag_snapshot_proto_depIdxs = nil
}
