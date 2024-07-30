// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/commons/audit/vana_events.proto

package audit

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// VanaFlagEvent represents a flag occurrence in voice analytics.
// It summarizes a set of transcripts flagged over an interval.
type VanaFlagEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. User is not allowed to see flag sid.
	FlagName string `protobuf:"bytes,1,opt,name=flag_name,json=flagName,proto3" json:"flag_name,omitempty"`
	// Required. Used for matching user notification settings.
	FlagSid int64 `protobuf:"varint,2,opt,name=flag_sid,json=flagSid,proto3" json:"flag_sid,omitempty"`
	// Required. Link to voice analytics page that shows reported transcripts.
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// Required. Number of transcripts that were flagged.
	NumTranscripts int64 `protobuf:"varint,4,opt,name=num_transcripts,json=numTranscripts,proto3" json:"num_transcripts,omitempty"`
	// Required. Flag priority.
	Priority int32 `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *VanaFlagEvent) Reset() {
	*x = VanaFlagEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_vana_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VanaFlagEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VanaFlagEvent) ProtoMessage() {}

func (x *VanaFlagEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_vana_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VanaFlagEvent.ProtoReflect.Descriptor instead.
func (*VanaFlagEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_vana_events_proto_rawDescGZIP(), []int{0}
}

func (x *VanaFlagEvent) GetFlagName() string {
	if x != nil {
		return x.FlagName
	}
	return ""
}

func (x *VanaFlagEvent) GetFlagSid() int64 {
	if x != nil {
		return x.FlagSid
	}
	return 0
}

func (x *VanaFlagEvent) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VanaFlagEvent) GetNumTranscripts() int64 {
	if x != nil {
		return x.NumTranscripts
	}
	return 0
}

func (x *VanaFlagEvent) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

// VanaFlagReviewEvent represents a flag that needs review in voice analytics.
// It summarizes a set of transcripts that need review.
type VanaFlagReviewEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. User is not allowed to see flag sid.
	FlagName string `protobuf:"bytes,1,opt,name=flag_name,json=flagName,proto3" json:"flag_name,omitempty"`
	// Required. Used for matching user notification settings.
	FlagSid int64 `protobuf:"varint,2,opt,name=flag_sid,json=flagSid,proto3" json:"flag_sid,omitempty"`
	// Required. Link to voice analytics page that shows reported transcripts.
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// Required. Number of transcripts that were flagged.
	NumTranscripts int64 `protobuf:"varint,4,opt,name=num_transcripts,json=numTranscripts,proto3" json:"num_transcripts,omitempty"`
	// Required. Flag priority.
	Priority int32 `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *VanaFlagReviewEvent) Reset() {
	*x = VanaFlagReviewEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_vana_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VanaFlagReviewEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VanaFlagReviewEvent) ProtoMessage() {}

func (x *VanaFlagReviewEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_vana_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VanaFlagReviewEvent.ProtoReflect.Descriptor instead.
func (*VanaFlagReviewEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_vana_events_proto_rawDescGZIP(), []int{1}
}

func (x *VanaFlagReviewEvent) GetFlagName() string {
	if x != nil {
		return x.FlagName
	}
	return ""
}

func (x *VanaFlagReviewEvent) GetFlagSid() int64 {
	if x != nil {
		return x.FlagSid
	}
	return 0
}

func (x *VanaFlagReviewEvent) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VanaFlagReviewEvent) GetNumTranscripts() int64 {
	if x != nil {
		return x.NumTranscripts
	}
	return 0
}

func (x *VanaFlagReviewEvent) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

// VanaBillingReportEvent represents a billing report in voice analytics.
// It gives a csv to a billing report of transcripts that occured that billing period.
type VanaBillingReportEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Report data start time.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Required. Report data end time.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Required. Report download url.
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *VanaBillingReportEvent) Reset() {
	*x = VanaBillingReportEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_vana_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VanaBillingReportEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VanaBillingReportEvent) ProtoMessage() {}

func (x *VanaBillingReportEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_vana_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VanaBillingReportEvent.ProtoReflect.Descriptor instead.
func (*VanaBillingReportEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_vana_events_proto_rawDescGZIP(), []int{2}
}

func (x *VanaBillingReportEvent) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *VanaBillingReportEvent) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *VanaBillingReportEvent) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// VanaFlagSummaryEvent is the summary of the flagging process.
// It shows what transcripts were flagged.
type VanaFlagSummaryEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Report data start time.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Required. Report data end time.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Required. Flag Summaries
	FlagSummaries []*VanaFlagSummaryEvent_FlagSummary `protobuf:"bytes,3,rep,name=flag_summaries,json=flagSummaries,proto3" json:"flag_summaries,omitempty"`
}

func (x *VanaFlagSummaryEvent) Reset() {
	*x = VanaFlagSummaryEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_vana_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VanaFlagSummaryEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VanaFlagSummaryEvent) ProtoMessage() {}

func (x *VanaFlagSummaryEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_vana_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VanaFlagSummaryEvent.ProtoReflect.Descriptor instead.
func (*VanaFlagSummaryEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_vana_events_proto_rawDescGZIP(), []int{3}
}

func (x *VanaFlagSummaryEvent) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *VanaFlagSummaryEvent) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *VanaFlagSummaryEvent) GetFlagSummaries() []*VanaFlagSummaryEvent_FlagSummary {
	if x != nil {
		return x.FlagSummaries
	}
	return nil
}

// VanaPhraseCorrectionEvent is a user suggestion for what some words in a transcript
// should be changed to when the original translation is off.
type VanaPhraseCorrectionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Start offset of the selected words in the transcript.
	StartOffset *durationpb.Duration `protobuf:"bytes,1,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	// Required. End offset of the selected words in the transcript.
	EndOffset *durationpb.Duration `protobuf:"bytes,2,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`
	// Required. Original text that the transcript was translated to.
	OriginalText string `protobuf:"bytes,3,opt,name=original_text,json=originalText,proto3" json:"original_text,omitempty"`
	// Required. Proposed text of what the transcript text should be.
	ProposedText string `protobuf:"bytes,4,opt,name=proposed_text,json=proposedText,proto3" json:"proposed_text,omitempty"`
	// Required. Url of the call recording for the transcript.
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	// Required. Channel that the words were spoken on.
	Channel uint32 `protobuf:"varint,6,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *VanaPhraseCorrectionEvent) Reset() {
	*x = VanaPhraseCorrectionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_vana_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VanaPhraseCorrectionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VanaPhraseCorrectionEvent) ProtoMessage() {}

func (x *VanaPhraseCorrectionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_vana_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VanaPhraseCorrectionEvent.ProtoReflect.Descriptor instead.
func (*VanaPhraseCorrectionEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_vana_events_proto_rawDescGZIP(), []int{4}
}

func (x *VanaPhraseCorrectionEvent) GetStartOffset() *durationpb.Duration {
	if x != nil {
		return x.StartOffset
	}
	return nil
}

func (x *VanaPhraseCorrectionEvent) GetEndOffset() *durationpb.Duration {
	if x != nil {
		return x.EndOffset
	}
	return nil
}

func (x *VanaPhraseCorrectionEvent) GetOriginalText() string {
	if x != nil {
		return x.OriginalText
	}
	return ""
}

func (x *VanaPhraseCorrectionEvent) GetProposedText() string {
	if x != nil {
		return x.ProposedText
	}
	return ""
}

func (x *VanaPhraseCorrectionEvent) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VanaPhraseCorrectionEvent) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

// VanaCreateTranscriptEvent represents a newly created transcript.
type VanaCreateTranscriptEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TranscriptSid int64 `protobuf:"varint,1,opt,name=transcript_sid,json=transcriptSid,proto3" json:"transcript_sid,omitempty"`
}

func (x *VanaCreateTranscriptEvent) Reset() {
	*x = VanaCreateTranscriptEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_vana_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VanaCreateTranscriptEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VanaCreateTranscriptEvent) ProtoMessage() {}

func (x *VanaCreateTranscriptEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_vana_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VanaCreateTranscriptEvent.ProtoReflect.Descriptor instead.
func (*VanaCreateTranscriptEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_vana_events_proto_rawDescGZIP(), []int{5}
}

func (x *VanaCreateTranscriptEvent) GetTranscriptSid() int64 {
	if x != nil {
		return x.TranscriptSid
	}
	return 0
}

// FlagSummary is an overview of the flags in the transcript.
type VanaFlagSummaryEvent_FlagSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TranscriptSid int64   `protobuf:"varint,1,opt,name=transcript_sid,json=transcriptSid,proto3" json:"transcript_sid,omitempty"`
	FlagSids      []int64 `protobuf:"varint,2,rep,packed,name=flag_sids,json=flagSids,proto3" json:"flag_sids,omitempty"`
}

func (x *VanaFlagSummaryEvent_FlagSummary) Reset() {
	*x = VanaFlagSummaryEvent_FlagSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_audit_vana_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VanaFlagSummaryEvent_FlagSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VanaFlagSummaryEvent_FlagSummary) ProtoMessage() {}

func (x *VanaFlagSummaryEvent_FlagSummary) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_vana_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VanaFlagSummaryEvent_FlagSummary.ProtoReflect.Descriptor instead.
func (*VanaFlagSummaryEvent_FlagSummary) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_vana_events_proto_rawDescGZIP(), []int{3, 0}
}

func (x *VanaFlagSummaryEvent_FlagSummary) GetTranscriptSid() int64 {
	if x != nil {
		return x.TranscriptSid
	}
	return 0
}

func (x *VanaFlagSummaryEvent_FlagSummary) GetFlagSids() []int64 {
	if x != nil {
		return x.FlagSids
	}
	return nil
}

var File_api_commons_audit_vana_events_proto protoreflect.FileDescriptor

var file_api_commons_audit_vana_events_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x0d, 0x56, 0x61,
	0x6e, 0x61, 0x46, 0x6c, 0x61, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x6c, 0x61, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x6c, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x61, 0x67,
	0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x6c, 0x61, 0x67,
	0x53, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x6e, 0x75, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0xa4, 0x01, 0x0a, 0x13, 0x56,
	0x61, 0x6e, 0x61, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0f,
	0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x22, 0x9c, 0x01, 0x0a, 0x16, 0x56, 0x61, 0x6e, 0x61, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0xb7, 0x02, 0x0a, 0x14, 0x56, 0x61, 0x6e, 0x61, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5a, 0x0a, 0x0e, 0x66,
	0x6c, 0x61, 0x67, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x56, 0x61, 0x6e, 0x61, 0x46, 0x6c, 0x61, 0x67,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6c, 0x61,
	0x67, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0d, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x51, 0x0a, 0x0b, 0x46, 0x6c, 0x61, 0x67, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x08, 0x66, 0x6c, 0x61, 0x67, 0x53, 0x69, 0x64, 0x73, 0x22, 0x89, 0x02, 0x0a, 0x19, 0x56,
	0x61, 0x6e, 0x61, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x54, 0x65, 0x78, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x42, 0x0a, 0x19, 0x56, 0x61, 0x6e, 0x61, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x64, 0x42, 0xbc, 0x01, 0x0a, 0x15, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x42, 0x0f, 0x56, 0x61, 0x6e, 0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x41, 0xaa, 0x02, 0x11, 0x41, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0xca,
	0x02, 0x11, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x75,
	0x64, 0x69, 0x74, 0xe2, 0x02, 0x1d, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x3a, 0x3a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_commons_audit_vana_events_proto_rawDescOnce sync.Once
	file_api_commons_audit_vana_events_proto_rawDescData = file_api_commons_audit_vana_events_proto_rawDesc
)

func file_api_commons_audit_vana_events_proto_rawDescGZIP() []byte {
	file_api_commons_audit_vana_events_proto_rawDescOnce.Do(func() {
		file_api_commons_audit_vana_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_audit_vana_events_proto_rawDescData)
	})
	return file_api_commons_audit_vana_events_proto_rawDescData
}

var file_api_commons_audit_vana_events_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_commons_audit_vana_events_proto_goTypes = []any{
	(*VanaFlagEvent)(nil),                    // 0: api.commons.audit.VanaFlagEvent
	(*VanaFlagReviewEvent)(nil),              // 1: api.commons.audit.VanaFlagReviewEvent
	(*VanaBillingReportEvent)(nil),           // 2: api.commons.audit.VanaBillingReportEvent
	(*VanaFlagSummaryEvent)(nil),             // 3: api.commons.audit.VanaFlagSummaryEvent
	(*VanaPhraseCorrectionEvent)(nil),        // 4: api.commons.audit.VanaPhraseCorrectionEvent
	(*VanaCreateTranscriptEvent)(nil),        // 5: api.commons.audit.VanaCreateTranscriptEvent
	(*VanaFlagSummaryEvent_FlagSummary)(nil), // 6: api.commons.audit.VanaFlagSummaryEvent.FlagSummary
	(*timestamppb.Timestamp)(nil),            // 7: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),              // 8: google.protobuf.Duration
}
var file_api_commons_audit_vana_events_proto_depIdxs = []int32{
	7, // 0: api.commons.audit.VanaBillingReportEvent.start_time:type_name -> google.protobuf.Timestamp
	7, // 1: api.commons.audit.VanaBillingReportEvent.end_time:type_name -> google.protobuf.Timestamp
	7, // 2: api.commons.audit.VanaFlagSummaryEvent.start_time:type_name -> google.protobuf.Timestamp
	7, // 3: api.commons.audit.VanaFlagSummaryEvent.end_time:type_name -> google.protobuf.Timestamp
	6, // 4: api.commons.audit.VanaFlagSummaryEvent.flag_summaries:type_name -> api.commons.audit.VanaFlagSummaryEvent.FlagSummary
	8, // 5: api.commons.audit.VanaPhraseCorrectionEvent.start_offset:type_name -> google.protobuf.Duration
	8, // 6: api.commons.audit.VanaPhraseCorrectionEvent.end_offset:type_name -> google.protobuf.Duration
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_commons_audit_vana_events_proto_init() }
func file_api_commons_audit_vana_events_proto_init() {
	if File_api_commons_audit_vana_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commons_audit_vana_events_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*VanaFlagEvent); i {
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
		file_api_commons_audit_vana_events_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*VanaFlagReviewEvent); i {
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
		file_api_commons_audit_vana_events_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*VanaBillingReportEvent); i {
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
		file_api_commons_audit_vana_events_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*VanaFlagSummaryEvent); i {
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
		file_api_commons_audit_vana_events_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*VanaPhraseCorrectionEvent); i {
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
		file_api_commons_audit_vana_events_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*VanaCreateTranscriptEvent); i {
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
		file_api_commons_audit_vana_events_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*VanaFlagSummaryEvent_FlagSummary); i {
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
			RawDescriptor: file_api_commons_audit_vana_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_audit_vana_events_proto_goTypes,
		DependencyIndexes: file_api_commons_audit_vana_events_proto_depIdxs,
		MessageInfos:      file_api_commons_audit_vana_events_proto_msgTypes,
	}.Build()
	File_api_commons_audit_vana_events_proto = out.File
	file_api_commons_audit_vana_events_proto_rawDesc = nil
	file_api_commons_audit_vana_events_proto_goTypes = nil
	file_api_commons_audit_vana_events_proto_depIdxs = nil
}
