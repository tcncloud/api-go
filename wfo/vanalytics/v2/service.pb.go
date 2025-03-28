// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: wfo/vanalytics/v2/service.proto

package vanalyticsv2

import (
	_ "github.com/tcncloud/api-go/annotations"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type AuditRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Since         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=since,proto3" json:"since,omitempty"`
	Until         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=until,proto3" json:"until,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuditRequest) Reset() {
	*x = AuditRequest{}
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditRequest) ProtoMessage() {}

func (x *AuditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditRequest.ProtoReflect.Descriptor instead.
func (*AuditRequest) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuditRequest) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

func (x *AuditRequest) GetUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.Until
	}
	return nil
}

type AuditResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// audio time in seconds within billing period
	AudioTime float64 `protobuf:"fixed64,1,opt,name=audio_time,json=audioTime,proto3" json:"audio_time,omitempty"`
	// billed_audio time in seconds within billing period
	BilledAudioTime float64 `protobuf:"fixed64,2,opt,name=billed_audio_time,json=billedAudioTime,proto3" json:"billed_audio_time,omitempty"`
	// last_usage is the date of the last transcription
	// within the billing period
	LastUsage *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_usage,json=lastUsage,proto3" json:"last_usage,omitempty"`
	// billed_transcripts is the number of billed transcripts
	// within the billing period
	BilledTranscripts int64 `protobuf:"varint,4,opt,name=billed_transcripts,json=billedTranscripts,proto3" json:"billed_transcripts,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *AuditResponse) Reset() {
	*x = AuditResponse{}
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditResponse) ProtoMessage() {}

func (x *AuditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditResponse.ProtoReflect.Descriptor instead.
func (*AuditResponse) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_service_proto_rawDescGZIP(), []int{1}
}

func (x *AuditResponse) GetAudioTime() float64 {
	if x != nil {
		return x.AudioTime
	}
	return 0
}

func (x *AuditResponse) GetBilledAudioTime() float64 {
	if x != nil {
		return x.BilledAudioTime
	}
	return 0
}

func (x *AuditResponse) GetLastUsage() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUsage
	}
	return nil
}

func (x *AuditResponse) GetBilledTranscripts() int64 {
	if x != nil {
		return x.BilledTranscripts
	}
	return 0
}

type GetRecordingUrlRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TranscriptSid int64                  `protobuf:"varint,2,opt,name=transcript_sid,json=transcriptSid,proto3" json:"transcript_sid,omitempty"`
	// Optional. Defaults to "wav". Can be "", "wav" or "mp3".
	Kind          string `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecordingUrlRequest) Reset() {
	*x = GetRecordingUrlRequest{}
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecordingUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordingUrlRequest) ProtoMessage() {}

func (x *GetRecordingUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordingUrlRequest.ProtoReflect.Descriptor instead.
func (*GetRecordingUrlRequest) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetRecordingUrlRequest) GetTranscriptSid() int64 {
	if x != nil {
		return x.TranscriptSid
	}
	return 0
}

func (x *GetRecordingUrlRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type GetRecordingUrlResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecordingUrlResponse) Reset() {
	*x = GetRecordingUrlResponse{}
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecordingUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordingUrlResponse) ProtoMessage() {}

func (x *GetRecordingUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordingUrlResponse.ProtoReflect.Descriptor instead.
func (*GetRecordingUrlResponse) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetRecordingUrlResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ListBillingSpanRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. Page size is based on months from now. Defaults to 12.
	PageSize uint32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token returned from a previous List request, if any.
	// When provided all other request fields are ignored.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBillingSpanRequest) Reset() {
	*x = ListBillingSpanRequest{}
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBillingSpanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBillingSpanRequest) ProtoMessage() {}

func (x *ListBillingSpanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBillingSpanRequest.ProtoReflect.Descriptor instead.
func (*ListBillingSpanRequest) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListBillingSpanRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBillingSpanRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListBillingSpanResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Token to retrieve the next page of billing spans, or empty if there are no
	// more billing spans in the list.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of billing spans which contains at most one request page_size.
	Spans         []*BillingSpan `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBillingSpanResponse) Reset() {
	*x = ListBillingSpanResponse{}
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBillingSpanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBillingSpanResponse) ProtoMessage() {}

func (x *ListBillingSpanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBillingSpanResponse.ProtoReflect.Descriptor instead.
func (*ListBillingSpanResponse) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListBillingSpanResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListBillingSpanResponse) GetSpans() []*BillingSpan {
	if x != nil {
		return x.Spans
	}
	return nil
}

type BillingSpan struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Number of calls/transcripts in the billing period.
	Calls int64 `protobuf:"varint,1,opt,name=calls,proto3" json:"calls,omitempty"`
	// Total hours in the billing period. Each call is rounded up to
	// the next 15 seconds and the total for the month is rounded up to
	// the next hour.
	Hours int64 `protobuf:"varint,2,opt,name=hours,proto3" json:"hours,omitempty"`
	// Cost is the total cost for the billing period.
	Cost float64 `protobuf:"fixed64,3,opt,name=cost,proto3" json:"cost,omitempty"`
	// Start time is the start time of the billing span. Transcripts in this span
	// must have create_time >= start_time.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End time is the end time of the billing span. Transcripts in this span must
	// have a create_time < end_time.
	EndTime       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BillingSpan) Reset() {
	*x = BillingSpan{}
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BillingSpan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingSpan) ProtoMessage() {}

func (x *BillingSpan) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingSpan.ProtoReflect.Descriptor instead.
func (*BillingSpan) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_service_proto_rawDescGZIP(), []int{6}
}

func (x *BillingSpan) GetCalls() int64 {
	if x != nil {
		return x.Calls
	}
	return 0
}

func (x *BillingSpan) GetHours() int64 {
	if x != nil {
		return x.Hours
	}
	return 0
}

func (x *BillingSpan) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *BillingSpan) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *BillingSpan) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_wfo_vanalytics_v2_service_proto protoreflect.FileDescriptor

const file_wfo_vanalytics_v2_service_proto_rawDesc = "" +
	"\n" +
	"\x1fwfo/vanalytics/v2/service.proto\x12\x11wfo.vanalytics.v2\x1a\x17annotations/authz.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\"wfo/vanalytics/v2/correction.proto\x1a\x1ewfo/vanalytics/v2/filter.proto\x1a\x1cwfo/vanalytics/v2/flag.proto\x1a#wfo/vanalytics/v2/flag_filter.proto\x1a#wfo/vanalytics/v2/flag_review.proto\x1a%wfo/vanalytics/v2/flag_snapshot.proto\x1a'wfo/vanalytics/v2/flag_transcript.proto\x1a.wfo/vanalytics/v2/flag_transcript_filter.proto\x1a\"wfo/vanalytics/v2/transcript.proto\x1a*wfo/vanalytics/v2/transcript_summary.proto\"r\n" +
	"\fAuditRequest\x120\n" +
	"\x05since\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x05since\x120\n" +
	"\x05until\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\x05until\"\xc4\x01\n" +
	"\rAuditResponse\x12\x1d\n" +
	"\n" +
	"audio_time\x18\x01 \x01(\x01R\taudioTime\x12*\n" +
	"\x11billed_audio_time\x18\x02 \x01(\x01R\x0fbilledAudioTime\x129\n" +
	"\n" +
	"last_usage\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tlastUsage\x12-\n" +
	"\x12billed_transcripts\x18\x04 \x01(\x03R\x11billedTranscripts\"S\n" +
	"\x16GetRecordingUrlRequest\x12%\n" +
	"\x0etranscript_sid\x18\x02 \x01(\x03R\rtranscriptSid\x12\x12\n" +
	"\x04kind\x18\x04 \x01(\tR\x04kind\"+\n" +
	"\x17GetRecordingUrlResponse\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\"T\n" +
	"\x16ListBillingSpanRequest\x12\x1b\n" +
	"\tpage_size\x18\x01 \x01(\rR\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x02 \x01(\tR\tpageToken\"w\n" +
	"\x17ListBillingSpanResponse\x12&\n" +
	"\x0fnext_page_token\x18\x01 \x01(\tR\rnextPageToken\x124\n" +
	"\x05spans\x18\x02 \x03(\v2\x1e.wfo.vanalytics.v2.BillingSpanR\x05spans\"\xbf\x01\n" +
	"\vBillingSpan\x12\x14\n" +
	"\x05calls\x18\x01 \x01(\x03R\x05calls\x12\x14\n" +
	"\x05hours\x18\x02 \x01(\x03R\x05hours\x12\x12\n" +
	"\x04cost\x18\x03 \x01(\x01R\x04cost\x129\n" +
	"\n" +
	"start_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTime\x125\n" +
	"\bend_time\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\aendTime2\xa3'\n" +
	"\n" +
	"Vanalytics\x12\x84\x01\n" +
	"\x05Audit\x12\x1f.wfo.vanalytics.v2.AuditRequest\x1a .wfo.vanalytics.v2.AuditResponse\"8\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02(:\x01*\"#/wfo/vanalytics/v2/vanalytics/audit\x12\xa6\x01\n" +
	"\x0fGetRecordingUrl\x12).wfo.vanalytics.v2.GetRecordingUrlRequest\x1a*.wfo.vanalytics.v2.GetRecordingUrlResponse\"<\xba\xb8\x91\x02\n" +
	"\n" +
	"\x03\b\xf4\x03\n" +
	"\x03\b\xd4\x02\x82\xd3\xe4\x93\x02':\x01*\"\"/wfo/vanalytics/v2/getrecordingurl\x12\xa1\x01\n" +
	"\x0fListBillingSpan\x12).wfo.vanalytics.v2.ListBillingSpanRequest\x1a*.wfo.vanalytics.v2.ListBillingSpanResponse\"7\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02':\x01*\"\"/wfo/vanalytics/v2/listbillingspan\x12\xae\x01\n" +
	"\x11SearchTranscripts\x12+.wfo.vanalytics.v2.SearchTranscriptsRequest\x1a,.wfo.vanalytics.v2.SearchTranscriptsResponse\">\xba\xb8\x91\x02\n" +
	"\n" +
	"\x03\b\xf4\x03\n" +
	"\x03\b\xd4\x02\x82\xd3\xe4\x93\x02):\x01*\"$/wfo/vanalytics/v2/searchtranscripts\x12\xb9\x01\n" +
	"\x15BulkDeleteTranscripts\x12/.wfo.vanalytics.v2.BulkDeleteTranscriptsRequest\x1a0.wfo.vanalytics.v2.BulkDeleteTranscriptsResponse\"=\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xfa\x03\x82\xd3\xe4\x93\x02-:\x01*\"(/wfo/vanalytics/v2/bulkdeletetranscripts\x12\xbd\x01\n" +
	"\x16BulkRestoreTranscripts\x120.wfo.vanalytics.v2.BulkRestoreTranscriptsRequest\x1a1.wfo.vanalytics.v2.BulkRestoreTranscriptsResponse\">\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xfa\x03\x82\xd3\xe4\x93\x02.:\x01*\")/wfo/vanalytics/v2/bulkrestoretranscripts\x12\xc5\x01\n" +
	"\x18ListTranscriptGroupNames\x122.wfo.vanalytics.v2.ListTranscriptGroupNamesRequest\x1a3.wfo.vanalytics.v2.ListTranscriptGroupNamesResponse\"@\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x020:\x01*\"+/wfo/vanalytics/v2/listtranscriptgroupnames\x12\xc1\x01\n" +
	"\x17ListAgentResponseValues\x121.wfo.vanalytics.v2.ListAgentResponseValuesRequest\x1a2.wfo.vanalytics.v2.ListAgentResponseValuesResponse\"?\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02/:\x01*\"*/wfo/vanalytics/v2/listagentresponsevalues\x12\xb5\x01\n" +
	"\x14GetTranscriptSummary\x12..wfo.vanalytics.v2.GetTranscriptSummaryRequest\x1a/.wfo.vanalytics.v2.GetTranscriptSummaryResponse\"<\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02,:\x01*\"'/wfo/vanalytics/v2/gettranscriptsummary\x12\x87\x01\n" +
	"\fCreateFilter\x12&.wfo.vanalytics.v2.CreateFilterRequest\x1a\x19.wfo.vanalytics.v2.Filter\"4\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02$:\x01*\"\x1f/wfo/vanalytics/v2/createfilter\x12\x96\x01\n" +
	"\vListFilters\x12%.wfo.vanalytics.v2.ListFiltersRequest\x1a&.wfo.vanalytics.v2.ListFiltersResponse\"8\xba\xb8\x91\x02\n" +
	"\n" +
	"\x03\b\xf4\x03\n" +
	"\x03\b\xc2\f\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/wfo/vanalytics/v2/listfilters\x12\x87\x01\n" +
	"\fUpdateFilter\x12&.wfo.vanalytics.v2.UpdateFilterRequest\x1a\x19.wfo.vanalytics.v2.Filter\"4\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02$:\x01*\"\x1f/wfo/vanalytics/v2/updatefilter\x12\x95\x01\n" +
	"\fDeleteFilter\x12&.wfo.vanalytics.v2.DeleteFilterRequest\x1a'.wfo.vanalytics.v2.DeleteFilterResponse\"4\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02$:\x01*\"\x1f/wfo/vanalytics/v2/deletefilter\x12~\n" +
	"\tGetFilter\x12#.wfo.vanalytics.v2.GetFilterRequest\x1a\x19.wfo.vanalytics.v2.Filter\"1\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02!:\x01*\"\x1c/wfo/vanalytics/v2/getfilter\x12\xc9\x01\n" +
	"\x19ListFlagTranscriptFilters\x123.wfo.vanalytics.v2.ListFlagTranscriptFiltersRequest\x1a4.wfo.vanalytics.v2.ListFlagTranscriptFiltersResponse\"A\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x021:\x01*\",/wfo/vanalytics/v2/listflagtranscriptfilters\x12\xa1\x01\n" +
	"\x0fListFlagFilters\x12).wfo.vanalytics.v2.ListFlagFiltersRequest\x1a*.wfo.vanalytics.v2.ListFlagFiltersResponse\"7\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02':\x01*\"\"/wfo/vanalytics/v2/listflagfilters\x12v\n" +
	"\aGetFlag\x12!.wfo.vanalytics.v2.GetFlagRequest\x1a\x17.wfo.vanalytics.v2.Flag\"/\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/wfo/vanalytics/v2/getflag\x12\x7f\n" +
	"\n" +
	"CreateFlag\x12$.wfo.vanalytics.v2.CreateFlagRequest\x1a\x17.wfo.vanalytics.v2.Flag\"2\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02\":\x01*\"\x1d/wfo/vanalytics/v2/createflag\x12\x89\x01\n" +
	"\tListFlags\x12#.wfo.vanalytics.v2.ListFlagsRequest\x1a$.wfo.vanalytics.v2.ListFlagsResponse\"1\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02!:\x01*\"\x1c/wfo/vanalytics/v2/listflags\x12\x7f\n" +
	"\n" +
	"UpdateFlag\x12$.wfo.vanalytics.v2.UpdateFlagRequest\x1a\x17.wfo.vanalytics.v2.Flag\"2\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02\":\x01*\"\x1d/wfo/vanalytics/v2/updateflag\x12\x8d\x01\n" +
	"\n" +
	"DeleteFlag\x12$.wfo.vanalytics.v2.DeleteFlagRequest\x1a%.wfo.vanalytics.v2.DeleteFlagResponse\"2\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02\":\x01*\"\x1d/wfo/vanalytics/v2/deleteflag\x12\x97\x01\n" +
	"\x10CreateFlagReview\x12*.wfo.vanalytics.v2.CreateFlagReviewRequest\x1a\x1d.wfo.vanalytics.v2.FlagReview\"8\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02(:\x01*\"#/wfo/vanalytics/v2/createflagreview\x12\xb5\x01\n" +
	"\x14BulkCreateFlagReview\x12..wfo.vanalytics.v2.BulkCreateFlagReviewRequest\x1a/.wfo.vanalytics.v2.BulkCreateFlagReviewResponse\"<\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02,:\x01*\"'/wfo/vanalytics/v2/bulkcreateflagreview\x12\xa1\x01\n" +
	"\x0fListFlagReviews\x12).wfo.vanalytics.v2.ListFlagReviewsRequest\x1a*.wfo.vanalytics.v2.ListFlagReviewsResponse\"7\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02':\x01*\"\"/wfo/vanalytics/v2/listflagreviews\x12\xb5\x01\n" +
	"\x14CreateFlagTranscript\x12..wfo.vanalytics.v2.CreateFlagTranscriptRequest\x1a/.wfo.vanalytics.v2.CreateFlagTranscriptResponse\"<\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02,:\x01*\"'/wfo/vanalytics/v2/createflagtranscript\x12\xa9\x01\n" +
	"\x11ListFlagSnapshots\x12+.wfo.vanalytics.v2.ListFlagSnapshotsRequest\x1a,.wfo.vanalytics.v2.ListFlagSnapshotsResponse\"9\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02):\x01*\"$/wfo/vanalytics/v2/listflagsnapshots\x12\xa5\x01\n" +
	"\x10CreateCorrection\x12*.wfo.vanalytics.v2.CreateCorrectionRequest\x1a+.wfo.vanalytics.v2.CreateCorrectionResponse\"8\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02(:\x01*\"#/wfo/vanalytics/v2/createcorrection\x12\x8e\x01\n" +
	"\rGetCorrection\x12'.wfo.vanalytics.v2.GetCorrectionRequest\x1a\x1d.wfo.vanalytics.v2.Correction\"5\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02%:\x01*\" /wfo/vanalytics/v2/getcorrection\x12\xa5\x01\n" +
	"\x10DeleteCorrection\x12*.wfo.vanalytics.v2.DeleteCorrectionRequest\x1a+.wfo.vanalytics.v2.DeleteCorrectionResponse\"8\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02(:\x01*\"#/wfo/vanalytics/v2/deletecorrection\x12\xa1\x01\n" +
	"\x0fListCorrections\x12).wfo.vanalytics.v2.ListCorrectionsRequest\x1a*.wfo.vanalytics.v2.ListCorrectionsResponse\"7\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02':\x01*\"\"/wfo/vanalytics/v2/listcorrections\x12\xa5\x01\n" +
	"\x10UpdateCorrection\x12*.wfo.vanalytics.v2.UpdateCorrectionRequest\x1a+.wfo.vanalytics.v2.UpdateCorrectionResponse\"8\xba\xb8\x91\x02\x05\n" +
	"\x03\b\xf4\x03\x82\xd3\xe4\x93\x02(:\x01*\"#/wfo/vanalytics/v2/updatecorrectionB\xc6\x01\n" +
	"\x15com.wfo.vanalytics.v2B\fServiceProtoP\x01Z9github.com/tcncloud/api-go/wfo/vanalytics/v2;vanalyticsv2\xa2\x02\x03WVX\xaa\x02\x11Wfo.Vanalytics.V2\xca\x02\x11Wfo\\Vanalytics\\V2\xe2\x02\x1dWfo\\Vanalytics\\V2\\GPBMetadata\xea\x02\x13Wfo::Vanalytics::V2b\x06proto3"

var (
	file_wfo_vanalytics_v2_service_proto_rawDescOnce sync.Once
	file_wfo_vanalytics_v2_service_proto_rawDescData []byte
)

func file_wfo_vanalytics_v2_service_proto_rawDescGZIP() []byte {
	file_wfo_vanalytics_v2_service_proto_rawDescOnce.Do(func() {
		file_wfo_vanalytics_v2_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_wfo_vanalytics_v2_service_proto_rawDesc), len(file_wfo_vanalytics_v2_service_proto_rawDesc)))
	})
	return file_wfo_vanalytics_v2_service_proto_rawDescData
}

var file_wfo_vanalytics_v2_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_wfo_vanalytics_v2_service_proto_goTypes = []any{
	(*AuditRequest)(nil),                      // 0: wfo.vanalytics.v2.AuditRequest
	(*AuditResponse)(nil),                     // 1: wfo.vanalytics.v2.AuditResponse
	(*GetRecordingUrlRequest)(nil),            // 2: wfo.vanalytics.v2.GetRecordingUrlRequest
	(*GetRecordingUrlResponse)(nil),           // 3: wfo.vanalytics.v2.GetRecordingUrlResponse
	(*ListBillingSpanRequest)(nil),            // 4: wfo.vanalytics.v2.ListBillingSpanRequest
	(*ListBillingSpanResponse)(nil),           // 5: wfo.vanalytics.v2.ListBillingSpanResponse
	(*BillingSpan)(nil),                       // 6: wfo.vanalytics.v2.BillingSpan
	(*timestamppb.Timestamp)(nil),             // 7: google.protobuf.Timestamp
	(*SearchTranscriptsRequest)(nil),          // 8: wfo.vanalytics.v2.SearchTranscriptsRequest
	(*BulkDeleteTranscriptsRequest)(nil),      // 9: wfo.vanalytics.v2.BulkDeleteTranscriptsRequest
	(*BulkRestoreTranscriptsRequest)(nil),     // 10: wfo.vanalytics.v2.BulkRestoreTranscriptsRequest
	(*ListTranscriptGroupNamesRequest)(nil),   // 11: wfo.vanalytics.v2.ListTranscriptGroupNamesRequest
	(*ListAgentResponseValuesRequest)(nil),    // 12: wfo.vanalytics.v2.ListAgentResponseValuesRequest
	(*GetTranscriptSummaryRequest)(nil),       // 13: wfo.vanalytics.v2.GetTranscriptSummaryRequest
	(*CreateFilterRequest)(nil),               // 14: wfo.vanalytics.v2.CreateFilterRequest
	(*ListFiltersRequest)(nil),                // 15: wfo.vanalytics.v2.ListFiltersRequest
	(*UpdateFilterRequest)(nil),               // 16: wfo.vanalytics.v2.UpdateFilterRequest
	(*DeleteFilterRequest)(nil),               // 17: wfo.vanalytics.v2.DeleteFilterRequest
	(*GetFilterRequest)(nil),                  // 18: wfo.vanalytics.v2.GetFilterRequest
	(*ListFlagTranscriptFiltersRequest)(nil),  // 19: wfo.vanalytics.v2.ListFlagTranscriptFiltersRequest
	(*ListFlagFiltersRequest)(nil),            // 20: wfo.vanalytics.v2.ListFlagFiltersRequest
	(*GetFlagRequest)(nil),                    // 21: wfo.vanalytics.v2.GetFlagRequest
	(*CreateFlagRequest)(nil),                 // 22: wfo.vanalytics.v2.CreateFlagRequest
	(*ListFlagsRequest)(nil),                  // 23: wfo.vanalytics.v2.ListFlagsRequest
	(*UpdateFlagRequest)(nil),                 // 24: wfo.vanalytics.v2.UpdateFlagRequest
	(*DeleteFlagRequest)(nil),                 // 25: wfo.vanalytics.v2.DeleteFlagRequest
	(*CreateFlagReviewRequest)(nil),           // 26: wfo.vanalytics.v2.CreateFlagReviewRequest
	(*BulkCreateFlagReviewRequest)(nil),       // 27: wfo.vanalytics.v2.BulkCreateFlagReviewRequest
	(*ListFlagReviewsRequest)(nil),            // 28: wfo.vanalytics.v2.ListFlagReviewsRequest
	(*CreateFlagTranscriptRequest)(nil),       // 29: wfo.vanalytics.v2.CreateFlagTranscriptRequest
	(*ListFlagSnapshotsRequest)(nil),          // 30: wfo.vanalytics.v2.ListFlagSnapshotsRequest
	(*CreateCorrectionRequest)(nil),           // 31: wfo.vanalytics.v2.CreateCorrectionRequest
	(*GetCorrectionRequest)(nil),              // 32: wfo.vanalytics.v2.GetCorrectionRequest
	(*DeleteCorrectionRequest)(nil),           // 33: wfo.vanalytics.v2.DeleteCorrectionRequest
	(*ListCorrectionsRequest)(nil),            // 34: wfo.vanalytics.v2.ListCorrectionsRequest
	(*UpdateCorrectionRequest)(nil),           // 35: wfo.vanalytics.v2.UpdateCorrectionRequest
	(*SearchTranscriptsResponse)(nil),         // 36: wfo.vanalytics.v2.SearchTranscriptsResponse
	(*BulkDeleteTranscriptsResponse)(nil),     // 37: wfo.vanalytics.v2.BulkDeleteTranscriptsResponse
	(*BulkRestoreTranscriptsResponse)(nil),    // 38: wfo.vanalytics.v2.BulkRestoreTranscriptsResponse
	(*ListTranscriptGroupNamesResponse)(nil),  // 39: wfo.vanalytics.v2.ListTranscriptGroupNamesResponse
	(*ListAgentResponseValuesResponse)(nil),   // 40: wfo.vanalytics.v2.ListAgentResponseValuesResponse
	(*GetTranscriptSummaryResponse)(nil),      // 41: wfo.vanalytics.v2.GetTranscriptSummaryResponse
	(*Filter)(nil),                            // 42: wfo.vanalytics.v2.Filter
	(*ListFiltersResponse)(nil),               // 43: wfo.vanalytics.v2.ListFiltersResponse
	(*DeleteFilterResponse)(nil),              // 44: wfo.vanalytics.v2.DeleteFilterResponse
	(*ListFlagTranscriptFiltersResponse)(nil), // 45: wfo.vanalytics.v2.ListFlagTranscriptFiltersResponse
	(*ListFlagFiltersResponse)(nil),           // 46: wfo.vanalytics.v2.ListFlagFiltersResponse
	(*Flag)(nil),                              // 47: wfo.vanalytics.v2.Flag
	(*ListFlagsResponse)(nil),                 // 48: wfo.vanalytics.v2.ListFlagsResponse
	(*DeleteFlagResponse)(nil),                // 49: wfo.vanalytics.v2.DeleteFlagResponse
	(*FlagReview)(nil),                        // 50: wfo.vanalytics.v2.FlagReview
	(*BulkCreateFlagReviewResponse)(nil),      // 51: wfo.vanalytics.v2.BulkCreateFlagReviewResponse
	(*ListFlagReviewsResponse)(nil),           // 52: wfo.vanalytics.v2.ListFlagReviewsResponse
	(*CreateFlagTranscriptResponse)(nil),      // 53: wfo.vanalytics.v2.CreateFlagTranscriptResponse
	(*ListFlagSnapshotsResponse)(nil),         // 54: wfo.vanalytics.v2.ListFlagSnapshotsResponse
	(*CreateCorrectionResponse)(nil),          // 55: wfo.vanalytics.v2.CreateCorrectionResponse
	(*Correction)(nil),                        // 56: wfo.vanalytics.v2.Correction
	(*DeleteCorrectionResponse)(nil),          // 57: wfo.vanalytics.v2.DeleteCorrectionResponse
	(*ListCorrectionsResponse)(nil),           // 58: wfo.vanalytics.v2.ListCorrectionsResponse
	(*UpdateCorrectionResponse)(nil),          // 59: wfo.vanalytics.v2.UpdateCorrectionResponse
}
var file_wfo_vanalytics_v2_service_proto_depIdxs = []int32{
	7,  // 0: wfo.vanalytics.v2.AuditRequest.since:type_name -> google.protobuf.Timestamp
	7,  // 1: wfo.vanalytics.v2.AuditRequest.until:type_name -> google.protobuf.Timestamp
	7,  // 2: wfo.vanalytics.v2.AuditResponse.last_usage:type_name -> google.protobuf.Timestamp
	6,  // 3: wfo.vanalytics.v2.ListBillingSpanResponse.spans:type_name -> wfo.vanalytics.v2.BillingSpan
	7,  // 4: wfo.vanalytics.v2.BillingSpan.start_time:type_name -> google.protobuf.Timestamp
	7,  // 5: wfo.vanalytics.v2.BillingSpan.end_time:type_name -> google.protobuf.Timestamp
	0,  // 6: wfo.vanalytics.v2.Vanalytics.Audit:input_type -> wfo.vanalytics.v2.AuditRequest
	2,  // 7: wfo.vanalytics.v2.Vanalytics.GetRecordingUrl:input_type -> wfo.vanalytics.v2.GetRecordingUrlRequest
	4,  // 8: wfo.vanalytics.v2.Vanalytics.ListBillingSpan:input_type -> wfo.vanalytics.v2.ListBillingSpanRequest
	8,  // 9: wfo.vanalytics.v2.Vanalytics.SearchTranscripts:input_type -> wfo.vanalytics.v2.SearchTranscriptsRequest
	9,  // 10: wfo.vanalytics.v2.Vanalytics.BulkDeleteTranscripts:input_type -> wfo.vanalytics.v2.BulkDeleteTranscriptsRequest
	10, // 11: wfo.vanalytics.v2.Vanalytics.BulkRestoreTranscripts:input_type -> wfo.vanalytics.v2.BulkRestoreTranscriptsRequest
	11, // 12: wfo.vanalytics.v2.Vanalytics.ListTranscriptGroupNames:input_type -> wfo.vanalytics.v2.ListTranscriptGroupNamesRequest
	12, // 13: wfo.vanalytics.v2.Vanalytics.ListAgentResponseValues:input_type -> wfo.vanalytics.v2.ListAgentResponseValuesRequest
	13, // 14: wfo.vanalytics.v2.Vanalytics.GetTranscriptSummary:input_type -> wfo.vanalytics.v2.GetTranscriptSummaryRequest
	14, // 15: wfo.vanalytics.v2.Vanalytics.CreateFilter:input_type -> wfo.vanalytics.v2.CreateFilterRequest
	15, // 16: wfo.vanalytics.v2.Vanalytics.ListFilters:input_type -> wfo.vanalytics.v2.ListFiltersRequest
	16, // 17: wfo.vanalytics.v2.Vanalytics.UpdateFilter:input_type -> wfo.vanalytics.v2.UpdateFilterRequest
	17, // 18: wfo.vanalytics.v2.Vanalytics.DeleteFilter:input_type -> wfo.vanalytics.v2.DeleteFilterRequest
	18, // 19: wfo.vanalytics.v2.Vanalytics.GetFilter:input_type -> wfo.vanalytics.v2.GetFilterRequest
	19, // 20: wfo.vanalytics.v2.Vanalytics.ListFlagTranscriptFilters:input_type -> wfo.vanalytics.v2.ListFlagTranscriptFiltersRequest
	20, // 21: wfo.vanalytics.v2.Vanalytics.ListFlagFilters:input_type -> wfo.vanalytics.v2.ListFlagFiltersRequest
	21, // 22: wfo.vanalytics.v2.Vanalytics.GetFlag:input_type -> wfo.vanalytics.v2.GetFlagRequest
	22, // 23: wfo.vanalytics.v2.Vanalytics.CreateFlag:input_type -> wfo.vanalytics.v2.CreateFlagRequest
	23, // 24: wfo.vanalytics.v2.Vanalytics.ListFlags:input_type -> wfo.vanalytics.v2.ListFlagsRequest
	24, // 25: wfo.vanalytics.v2.Vanalytics.UpdateFlag:input_type -> wfo.vanalytics.v2.UpdateFlagRequest
	25, // 26: wfo.vanalytics.v2.Vanalytics.DeleteFlag:input_type -> wfo.vanalytics.v2.DeleteFlagRequest
	26, // 27: wfo.vanalytics.v2.Vanalytics.CreateFlagReview:input_type -> wfo.vanalytics.v2.CreateFlagReviewRequest
	27, // 28: wfo.vanalytics.v2.Vanalytics.BulkCreateFlagReview:input_type -> wfo.vanalytics.v2.BulkCreateFlagReviewRequest
	28, // 29: wfo.vanalytics.v2.Vanalytics.ListFlagReviews:input_type -> wfo.vanalytics.v2.ListFlagReviewsRequest
	29, // 30: wfo.vanalytics.v2.Vanalytics.CreateFlagTranscript:input_type -> wfo.vanalytics.v2.CreateFlagTranscriptRequest
	30, // 31: wfo.vanalytics.v2.Vanalytics.ListFlagSnapshots:input_type -> wfo.vanalytics.v2.ListFlagSnapshotsRequest
	31, // 32: wfo.vanalytics.v2.Vanalytics.CreateCorrection:input_type -> wfo.vanalytics.v2.CreateCorrectionRequest
	32, // 33: wfo.vanalytics.v2.Vanalytics.GetCorrection:input_type -> wfo.vanalytics.v2.GetCorrectionRequest
	33, // 34: wfo.vanalytics.v2.Vanalytics.DeleteCorrection:input_type -> wfo.vanalytics.v2.DeleteCorrectionRequest
	34, // 35: wfo.vanalytics.v2.Vanalytics.ListCorrections:input_type -> wfo.vanalytics.v2.ListCorrectionsRequest
	35, // 36: wfo.vanalytics.v2.Vanalytics.UpdateCorrection:input_type -> wfo.vanalytics.v2.UpdateCorrectionRequest
	1,  // 37: wfo.vanalytics.v2.Vanalytics.Audit:output_type -> wfo.vanalytics.v2.AuditResponse
	3,  // 38: wfo.vanalytics.v2.Vanalytics.GetRecordingUrl:output_type -> wfo.vanalytics.v2.GetRecordingUrlResponse
	5,  // 39: wfo.vanalytics.v2.Vanalytics.ListBillingSpan:output_type -> wfo.vanalytics.v2.ListBillingSpanResponse
	36, // 40: wfo.vanalytics.v2.Vanalytics.SearchTranscripts:output_type -> wfo.vanalytics.v2.SearchTranscriptsResponse
	37, // 41: wfo.vanalytics.v2.Vanalytics.BulkDeleteTranscripts:output_type -> wfo.vanalytics.v2.BulkDeleteTranscriptsResponse
	38, // 42: wfo.vanalytics.v2.Vanalytics.BulkRestoreTranscripts:output_type -> wfo.vanalytics.v2.BulkRestoreTranscriptsResponse
	39, // 43: wfo.vanalytics.v2.Vanalytics.ListTranscriptGroupNames:output_type -> wfo.vanalytics.v2.ListTranscriptGroupNamesResponse
	40, // 44: wfo.vanalytics.v2.Vanalytics.ListAgentResponseValues:output_type -> wfo.vanalytics.v2.ListAgentResponseValuesResponse
	41, // 45: wfo.vanalytics.v2.Vanalytics.GetTranscriptSummary:output_type -> wfo.vanalytics.v2.GetTranscriptSummaryResponse
	42, // 46: wfo.vanalytics.v2.Vanalytics.CreateFilter:output_type -> wfo.vanalytics.v2.Filter
	43, // 47: wfo.vanalytics.v2.Vanalytics.ListFilters:output_type -> wfo.vanalytics.v2.ListFiltersResponse
	42, // 48: wfo.vanalytics.v2.Vanalytics.UpdateFilter:output_type -> wfo.vanalytics.v2.Filter
	44, // 49: wfo.vanalytics.v2.Vanalytics.DeleteFilter:output_type -> wfo.vanalytics.v2.DeleteFilterResponse
	42, // 50: wfo.vanalytics.v2.Vanalytics.GetFilter:output_type -> wfo.vanalytics.v2.Filter
	45, // 51: wfo.vanalytics.v2.Vanalytics.ListFlagTranscriptFilters:output_type -> wfo.vanalytics.v2.ListFlagTranscriptFiltersResponse
	46, // 52: wfo.vanalytics.v2.Vanalytics.ListFlagFilters:output_type -> wfo.vanalytics.v2.ListFlagFiltersResponse
	47, // 53: wfo.vanalytics.v2.Vanalytics.GetFlag:output_type -> wfo.vanalytics.v2.Flag
	47, // 54: wfo.vanalytics.v2.Vanalytics.CreateFlag:output_type -> wfo.vanalytics.v2.Flag
	48, // 55: wfo.vanalytics.v2.Vanalytics.ListFlags:output_type -> wfo.vanalytics.v2.ListFlagsResponse
	47, // 56: wfo.vanalytics.v2.Vanalytics.UpdateFlag:output_type -> wfo.vanalytics.v2.Flag
	49, // 57: wfo.vanalytics.v2.Vanalytics.DeleteFlag:output_type -> wfo.vanalytics.v2.DeleteFlagResponse
	50, // 58: wfo.vanalytics.v2.Vanalytics.CreateFlagReview:output_type -> wfo.vanalytics.v2.FlagReview
	51, // 59: wfo.vanalytics.v2.Vanalytics.BulkCreateFlagReview:output_type -> wfo.vanalytics.v2.BulkCreateFlagReviewResponse
	52, // 60: wfo.vanalytics.v2.Vanalytics.ListFlagReviews:output_type -> wfo.vanalytics.v2.ListFlagReviewsResponse
	53, // 61: wfo.vanalytics.v2.Vanalytics.CreateFlagTranscript:output_type -> wfo.vanalytics.v2.CreateFlagTranscriptResponse
	54, // 62: wfo.vanalytics.v2.Vanalytics.ListFlagSnapshots:output_type -> wfo.vanalytics.v2.ListFlagSnapshotsResponse
	55, // 63: wfo.vanalytics.v2.Vanalytics.CreateCorrection:output_type -> wfo.vanalytics.v2.CreateCorrectionResponse
	56, // 64: wfo.vanalytics.v2.Vanalytics.GetCorrection:output_type -> wfo.vanalytics.v2.Correction
	57, // 65: wfo.vanalytics.v2.Vanalytics.DeleteCorrection:output_type -> wfo.vanalytics.v2.DeleteCorrectionResponse
	58, // 66: wfo.vanalytics.v2.Vanalytics.ListCorrections:output_type -> wfo.vanalytics.v2.ListCorrectionsResponse
	59, // 67: wfo.vanalytics.v2.Vanalytics.UpdateCorrection:output_type -> wfo.vanalytics.v2.UpdateCorrectionResponse
	37, // [37:68] is the sub-list for method output_type
	6,  // [6:37] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_wfo_vanalytics_v2_service_proto_init() }
func file_wfo_vanalytics_v2_service_proto_init() {
	if File_wfo_vanalytics_v2_service_proto != nil {
		return
	}
	file_wfo_vanalytics_v2_correction_proto_init()
	file_wfo_vanalytics_v2_filter_proto_init()
	file_wfo_vanalytics_v2_flag_proto_init()
	file_wfo_vanalytics_v2_flag_filter_proto_init()
	file_wfo_vanalytics_v2_flag_review_proto_init()
	file_wfo_vanalytics_v2_flag_snapshot_proto_init()
	file_wfo_vanalytics_v2_flag_transcript_proto_init()
	file_wfo_vanalytics_v2_flag_transcript_filter_proto_init()
	file_wfo_vanalytics_v2_transcript_proto_init()
	file_wfo_vanalytics_v2_transcript_summary_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_wfo_vanalytics_v2_service_proto_rawDesc), len(file_wfo_vanalytics_v2_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wfo_vanalytics_v2_service_proto_goTypes,
		DependencyIndexes: file_wfo_vanalytics_v2_service_proto_depIdxs,
		MessageInfos:      file_wfo_vanalytics_v2_service_proto_msgTypes,
	}.Build()
	File_wfo_vanalytics_v2_service_proto = out.File
	file_wfo_vanalytics_v2_service_proto_goTypes = nil
	file_wfo_vanalytics_v2_service_proto_depIdxs = nil
}
