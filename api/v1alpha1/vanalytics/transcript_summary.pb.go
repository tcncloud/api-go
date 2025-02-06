// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/v1alpha1/vanalytics/transcript_summary.proto

package vanalytics

import (
	commons "github.com/tcncloud/api-go/api/commons"
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

// GetTranscriptSummaryRequest is a request for getting a transcript summary.
type GetTranscriptSummaryRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required - transcript to get summary of.
	TranscriptSid int64 `protobuf:"varint,2,opt,name=transcript_sid,json=transcriptSid,proto3" json:"transcript_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTranscriptSummaryRequest) Reset() {
	*x = GetTranscriptSummaryRequest{}
	mi := &file_api_v1alpha1_vanalytics_transcript_summary_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTranscriptSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTranscriptSummaryRequest) ProtoMessage() {}

func (x *GetTranscriptSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_transcript_summary_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTranscriptSummaryRequest.ProtoReflect.Descriptor instead.
func (*GetTranscriptSummaryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDescGZIP(), []int{0}
}

func (x *GetTranscriptSummaryRequest) GetTranscriptSid() int64 {
	if x != nil {
		return x.TranscriptSid
	}
	return 0
}

// GetTranscriptSummaryResponse is a response for getting a transcript summary.
type GetTranscriptSummaryResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	TranscriptSummary *TranscriptSummary     `protobuf:"bytes,1,opt,name=transcript_summary,json=transcriptSummary,proto3" json:"transcript_summary,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetTranscriptSummaryResponse) Reset() {
	*x = GetTranscriptSummaryResponse{}
	mi := &file_api_v1alpha1_vanalytics_transcript_summary_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTranscriptSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTranscriptSummaryResponse) ProtoMessage() {}

func (x *GetTranscriptSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_transcript_summary_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTranscriptSummaryResponse.ProtoReflect.Descriptor instead.
func (*GetTranscriptSummaryResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDescGZIP(), []int{1}
}

func (x *GetTranscriptSummaryResponse) GetTranscriptSummary() *TranscriptSummary {
	if x != nil {
		return x.TranscriptSummary
	}
	return nil
}

// TranscriptSummary is an AI generated summary based on a parent transcript.
type TranscriptSummary struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Summary content provided in multiple bullet points.
	BulletPoints []string `protobuf:"bytes,1,rep,name=bullet_points,json=bulletPoints,proto3" json:"bullet_points,omitempty"`
	// Status of the summary.
	Status        commons.TranscriptSummaryStatus `protobuf:"varint,2,opt,name=status,proto3,enum=api.commons.TranscriptSummaryStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TranscriptSummary) Reset() {
	*x = TranscriptSummary{}
	mi := &file_api_v1alpha1_vanalytics_transcript_summary_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TranscriptSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscriptSummary) ProtoMessage() {}

func (x *TranscriptSummary) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_transcript_summary_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscriptSummary.ProtoReflect.Descriptor instead.
func (*TranscriptSummary) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDescGZIP(), []int{2}
}

func (x *TranscriptSummary) GetBulletPoints() []string {
	if x != nil {
		return x.BulletPoints
	}
	return nil
}

func (x *TranscriptSummary) GetStatus() commons.TranscriptSummaryStatus {
	if x != nil {
		return x.Status
	}
	return commons.TranscriptSummaryStatus(0)
}

var File_api_v1alpha1_vanalytics_transcript_summary_proto protoreflect.FileDescriptor

var file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDesc = string([]byte{
	0x0a, 0x30, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1c, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x64, 0x22,
	0x79, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x59, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x76, 0x0a, 0x11, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0xe7, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x42, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0xa2, 0x02, 0x03, 0x41, 0x56, 0x56, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0xca, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3a, 0x3a, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDescOnce sync.Once
	file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDescData []byte
)

func file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDesc), len(file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDesc)))
	})
	return file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDescData
}

var file_api_v1alpha1_vanalytics_transcript_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1alpha1_vanalytics_transcript_summary_proto_goTypes = []any{
	(*GetTranscriptSummaryRequest)(nil),  // 0: api.v1alpha1.vanalytics.GetTranscriptSummaryRequest
	(*GetTranscriptSummaryResponse)(nil), // 1: api.v1alpha1.vanalytics.GetTranscriptSummaryResponse
	(*TranscriptSummary)(nil),            // 2: api.v1alpha1.vanalytics.TranscriptSummary
	(commons.TranscriptSummaryStatus)(0), // 3: api.commons.TranscriptSummaryStatus
}
var file_api_v1alpha1_vanalytics_transcript_summary_proto_depIdxs = []int32{
	2, // 0: api.v1alpha1.vanalytics.GetTranscriptSummaryResponse.transcript_summary:type_name -> api.v1alpha1.vanalytics.TranscriptSummary
	3, // 1: api.v1alpha1.vanalytics.TranscriptSummary.status:type_name -> api.commons.TranscriptSummaryStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_vanalytics_transcript_summary_proto_init() }
func file_api_v1alpha1_vanalytics_transcript_summary_proto_init() {
	if File_api_v1alpha1_vanalytics_transcript_summary_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDesc), len(file_api_v1alpha1_vanalytics_transcript_summary_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_vanalytics_transcript_summary_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_vanalytics_transcript_summary_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_vanalytics_transcript_summary_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_vanalytics_transcript_summary_proto = out.File
	file_api_v1alpha1_vanalytics_transcript_summary_proto_goTypes = nil
	file_api_v1alpha1_vanalytics_transcript_summary_proto_depIdxs = nil
}
