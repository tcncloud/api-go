// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v1alpha1/srec/service.proto

package srec

import (
	_ "github.com/tcncloud/api-go/annotations"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// ListScreenRecordingsRequest is a request for listing screen recordings.
type ListScreenRecordingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The number of screen recordings to include in a single response. When not
	// provided this defaults to 100.
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token returned from a previous List request, if any.
	// When provided all other request fields are ignored.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListScreenRecordingsRequest) Reset() {
	*x = ListScreenRecordingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_srec_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListScreenRecordingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScreenRecordingsRequest) ProtoMessage() {}

func (x *ListScreenRecordingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_srec_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScreenRecordingsRequest.ProtoReflect.Descriptor instead.
func (*ListScreenRecordingsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_srec_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListScreenRecordingsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListScreenRecordingsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// ListScreenRecordingsResponse is a response for listing screen recordings.
type ListScreenRecordingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to retrieve the next page of screen recordings, or empty if there are no
	// more screen recordings in the list.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of screen recordings which contains at most one request page_size.
	Recordings []*ScreenRecording `protobuf:"bytes,2,rep,name=recordings,proto3" json:"recordings,omitempty"`
}

func (x *ListScreenRecordingsResponse) Reset() {
	*x = ListScreenRecordingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_srec_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListScreenRecordingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScreenRecordingsResponse) ProtoMessage() {}

func (x *ListScreenRecordingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_srec_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScreenRecordingsResponse.ProtoReflect.Descriptor instead.
func (*ListScreenRecordingsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_srec_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListScreenRecordingsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListScreenRecordingsResponse) GetRecordings() []*ScreenRecording {
	if x != nil {
		return x.Recordings
	}
	return nil
}

// ScreenRecording is a resource in the SREC API.
type ScreenRecording struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The unique id of this screen recording.
	SessionId int64 `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Required. The agent first name.
	AgentFirstName string `protobuf:"bytes,3,opt,name=agent_first_name,json=agentFirstName,proto3" json:"agent_first_name,omitempty"`
	// Required. The agent last name.
	AgentLastName string `protobuf:"bytes,4,opt,name=agent_last_name,json=agentLastName,proto3" json:"agent_last_name,omitempty"`
	// Required. Start time is the start time of the screen recording.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Required. Audio time in milliseconds.
	AudioTime uint32 `protobuf:"varint,6,opt,name=audio_time,json=audioTime,proto3" json:"audio_time,omitempty"`
	// Required. Audio bytes of the screen recording.
	AudioBytes int64 `protobuf:"varint,7,opt,name=audio_bytes,json=audioBytes,proto3" json:"audio_bytes,omitempty"`
}

func (x *ScreenRecording) Reset() {
	*x = ScreenRecording{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_srec_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreenRecording) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreenRecording) ProtoMessage() {}

func (x *ScreenRecording) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_srec_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreenRecording.ProtoReflect.Descriptor instead.
func (*ScreenRecording) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_srec_service_proto_rawDescGZIP(), []int{2}
}

func (x *ScreenRecording) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *ScreenRecording) GetAgentFirstName() string {
	if x != nil {
		return x.AgentFirstName
	}
	return ""
}

func (x *ScreenRecording) GetAgentLastName() string {
	if x != nil {
		return x.AgentLastName
	}
	return ""
}

func (x *ScreenRecording) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *ScreenRecording) GetAudioTime() uint32 {
	if x != nil {
		return x.AudioTime
	}
	return 0
}

func (x *ScreenRecording) GetAudioBytes() int64 {
	if x != nil {
		return x.AudioBytes
	}
	return 0
}

// GetScreenRecordingURLRequest is a request for getting a screen recording url.
type GetScreenRecordingURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The session id of this screen recording. Must be non empty.
	SessionId int64 `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *GetScreenRecordingURLRequest) Reset() {
	*x = GetScreenRecordingURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_srec_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScreenRecordingURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScreenRecordingURLRequest) ProtoMessage() {}

func (x *GetScreenRecordingURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_srec_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScreenRecordingURLRequest.ProtoReflect.Descriptor instead.
func (*GetScreenRecordingURLRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_srec_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetScreenRecordingURLRequest) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

// GetScreenRecordingURLResponse is a response for getting a screen recording url.
type GetScreenRecordingURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The url for the screen recording.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetScreenRecordingURLResponse) Reset() {
	*x = GetScreenRecordingURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_srec_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScreenRecordingURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScreenRecordingURLResponse) ProtoMessage() {}

func (x *GetScreenRecordingURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_srec_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScreenRecordingURLResponse.ProtoReflect.Descriptor instead.
func (*GetScreenRecordingURLResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_srec_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetScreenRecordingURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// DeleteScreenRecordingRequest is a request for deleting a screen recording.
type DeleteScreenRecordingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The unique id of the screen recording to be deleted.
	SessionId int64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *DeleteScreenRecordingRequest) Reset() {
	*x = DeleteScreenRecordingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_srec_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScreenRecordingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScreenRecordingRequest) ProtoMessage() {}

func (x *DeleteScreenRecordingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_srec_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScreenRecordingRequest.ProtoReflect.Descriptor instead.
func (*DeleteScreenRecordingRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_srec_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteScreenRecordingRequest) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

// DeleteScreenRecordingResponse is a response for deleting a screen recording.
type DeleteScreenRecordingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteScreenRecordingResponse) Reset() {
	*x = DeleteScreenRecordingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_srec_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScreenRecordingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScreenRecordingResponse) ProtoMessage() {}

func (x *DeleteScreenRecordingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_srec_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScreenRecordingResponse.ProtoReflect.Descriptor instead.
func (*DeleteScreenRecordingResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_srec_service_proto_rawDescGZIP(), []int{6}
}

var File_api_v1alpha1_srec_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_srec_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x72, 0x65, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x73, 0x72, 0x65, 0x63, 0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x1b,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x42, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x73, 0x72, 0x65, 0x63, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x22, 0xfd, 0x01, 0x0a, 0x0f, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x3d, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc5, 0x04, 0x0a, 0x04, 0x53, 0x72, 0x65, 0x63, 0x12,
	0xba, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x72, 0x65, 0x63, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x72, 0x65, 0x63, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0xba, 0xb8, 0x91, 0x02, 0x05,
	0x0a, 0x03, 0x08, 0xf9, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x3a, 0x01, 0x2a, 0x22, 0x2c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x72,
	0x65, 0x63, 0x2f, 0x73, 0x72, 0x65, 0x63, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0xbe, 0x01, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x12, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x72, 0x65, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x72, 0x65, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0xba, 0xb8, 0x91, 0x02, 0x05,
	0x0a, 0x03, 0x08, 0xf9, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x72,
	0x65, 0x63, 0x2f, 0x73, 0x72, 0x65, 0x63, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x75, 0x72, 0x6c, 0x12, 0xbe, 0x01,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x72, 0x65, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x72, 0x65, 0x63, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0xba, 0xb8, 0x91, 0x02,
	0x05, 0x0a, 0x03, 0x08, 0xfb, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22,
	0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x72, 0x65, 0x63, 0x2f, 0x73, 0x72, 0x65, 0x63, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x42, 0xb9,
	0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x73, 0x72, 0x65, 0x63, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x73, 0x72, 0x65, 0x63, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x53, 0xaa, 0x02, 0x11, 0x41,
	0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x72, 0x65, 0x63,
	0xca, 0x02, 0x11, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x53, 0x72, 0x65, 0x63, 0xe2, 0x02, 0x1d, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5c, 0x53, 0x72, 0x65, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x53, 0x72, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_v1alpha1_srec_service_proto_rawDescOnce sync.Once
	file_api_v1alpha1_srec_service_proto_rawDescData = file_api_v1alpha1_srec_service_proto_rawDesc
)

func file_api_v1alpha1_srec_service_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_srec_service_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_srec_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_srec_service_proto_rawDescData)
	})
	return file_api_v1alpha1_srec_service_proto_rawDescData
}

var file_api_v1alpha1_srec_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_v1alpha1_srec_service_proto_goTypes = []interface{}{
	(*ListScreenRecordingsRequest)(nil),   // 0: api.v1alpha1.srec.ListScreenRecordingsRequest
	(*ListScreenRecordingsResponse)(nil),  // 1: api.v1alpha1.srec.ListScreenRecordingsResponse
	(*ScreenRecording)(nil),               // 2: api.v1alpha1.srec.ScreenRecording
	(*GetScreenRecordingURLRequest)(nil),  // 3: api.v1alpha1.srec.GetScreenRecordingURLRequest
	(*GetScreenRecordingURLResponse)(nil), // 4: api.v1alpha1.srec.GetScreenRecordingURLResponse
	(*DeleteScreenRecordingRequest)(nil),  // 5: api.v1alpha1.srec.DeleteScreenRecordingRequest
	(*DeleteScreenRecordingResponse)(nil), // 6: api.v1alpha1.srec.DeleteScreenRecordingResponse
	(*timestamppb.Timestamp)(nil),         // 7: google.protobuf.Timestamp
}
var file_api_v1alpha1_srec_service_proto_depIdxs = []int32{
	2, // 0: api.v1alpha1.srec.ListScreenRecordingsResponse.recordings:type_name -> api.v1alpha1.srec.ScreenRecording
	7, // 1: api.v1alpha1.srec.ScreenRecording.start_time:type_name -> google.protobuf.Timestamp
	0, // 2: api.v1alpha1.srec.Srec.ListScreenRecordings:input_type -> api.v1alpha1.srec.ListScreenRecordingsRequest
	3, // 3: api.v1alpha1.srec.Srec.GetScreenRecordingURL:input_type -> api.v1alpha1.srec.GetScreenRecordingURLRequest
	5, // 4: api.v1alpha1.srec.Srec.DeleteScreenRecording:input_type -> api.v1alpha1.srec.DeleteScreenRecordingRequest
	1, // 5: api.v1alpha1.srec.Srec.ListScreenRecordings:output_type -> api.v1alpha1.srec.ListScreenRecordingsResponse
	4, // 6: api.v1alpha1.srec.Srec.GetScreenRecordingURL:output_type -> api.v1alpha1.srec.GetScreenRecordingURLResponse
	6, // 7: api.v1alpha1.srec.Srec.DeleteScreenRecording:output_type -> api.v1alpha1.srec.DeleteScreenRecordingResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_srec_service_proto_init() }
func file_api_v1alpha1_srec_service_proto_init() {
	if File_api_v1alpha1_srec_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1alpha1_srec_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListScreenRecordingsRequest); i {
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
		file_api_v1alpha1_srec_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListScreenRecordingsResponse); i {
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
		file_api_v1alpha1_srec_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreenRecording); i {
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
		file_api_v1alpha1_srec_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScreenRecordingURLRequest); i {
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
		file_api_v1alpha1_srec_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScreenRecordingURLResponse); i {
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
		file_api_v1alpha1_srec_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteScreenRecordingRequest); i {
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
		file_api_v1alpha1_srec_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteScreenRecordingResponse); i {
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
			RawDescriptor: file_api_v1alpha1_srec_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_srec_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_srec_service_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_srec_service_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_srec_service_proto = out.File
	file_api_v1alpha1_srec_service_proto_rawDesc = nil
	file_api_v1alpha1_srec_service_proto_goTypes = nil
	file_api_v1alpha1_srec_service_proto_depIdxs = nil
}
