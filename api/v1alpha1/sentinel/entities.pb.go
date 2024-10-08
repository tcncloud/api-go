// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/v1alpha1/sentinel/entities.proto

package sentinel

import (
	commons "github.com/tcncloud/api-go/api/commons"
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

type SentinelEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*SentinelEvent_LogEvent
	Event isSentinelEvent_Event `protobuf_oneof:"event"`
}

func (x *SentinelEvent) Reset() {
	*x = SentinelEvent{}
	mi := &file_api_v1alpha1_sentinel_entities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SentinelEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentinelEvent) ProtoMessage() {}

func (x *SentinelEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_sentinel_entities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SentinelEvent.ProtoReflect.Descriptor instead.
func (*SentinelEvent) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_sentinel_entities_proto_rawDescGZIP(), []int{0}
}

func (m *SentinelEvent) GetEvent() isSentinelEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *SentinelEvent) GetLogEvent() *LogEvent {
	if x, ok := x.GetEvent().(*SentinelEvent_LogEvent); ok {
		return x.LogEvent
	}
	return nil
}

type isSentinelEvent_Event interface {
	isSentinelEvent_Event()
}

type SentinelEvent_LogEvent struct {
	LogEvent *LogEvent `protobuf:"bytes,1,opt,name=log_event,json=logEvent,proto3,oneof"`
}

func (*SentinelEvent_LogEvent) isSentinelEvent_Event() {}

type LogEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId string `protobuf:"bytes,3,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// the browsers session id
	SessionId string `protobuf:"bytes,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Message   string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	// the browsers current uri
	Location   string                 `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	StackTrace string                 `protobuf:"bytes,7,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	Timestamp  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Metadata   map[string]string      `protobuf:"bytes,9,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Severity   commons.Level          `protobuf:"varint,10,opt,name=severity,proto3,enum=api.commons.Level" json:"severity,omitempty"`
}

func (x *LogEvent) Reset() {
	*x = LogEvent{}
	mi := &file_api_v1alpha1_sentinel_entities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEvent) ProtoMessage() {}

func (x *LogEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_sentinel_entities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEvent.ProtoReflect.Descriptor instead.
func (*LogEvent) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_sentinel_entities_proto_rawDescGZIP(), []int{1}
}

func (x *LogEvent) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *LogEvent) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *LogEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogEvent) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *LogEvent) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

func (x *LogEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogEvent) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *LogEvent) GetSeverity() commons.Level {
	if x != nil {
		return x.Severity
	}
	return commons.Level(0)
}

type SendEventsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*SentinelEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *SendEventsReq) Reset() {
	*x = SendEventsReq{}
	mi := &file_api_v1alpha1_sentinel_entities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendEventsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEventsReq) ProtoMessage() {}

func (x *SendEventsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_sentinel_entities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEventsReq.ProtoReflect.Descriptor instead.
func (*SendEventsReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_sentinel_entities_proto_rawDescGZIP(), []int{2}
}

func (x *SendEventsReq) GetEvents() []*SentinelEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type SendEventsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendEventsRes) Reset() {
	*x = SendEventsRes{}
	mi := &file_api_v1alpha1_sentinel_entities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendEventsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEventsRes) ProtoMessage() {}

func (x *SendEventsRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_sentinel_entities_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEventsRes.ProtoReflect.Descriptor instead.
func (*SendEventsRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_sentinel_entities_proto_rawDescGZIP(), []int{3}
}

var File_api_v1alpha1_sentinel_entities_proto protoreflect.FileDescriptor

var file_api_v1alpha1_sentinel_entities_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6c, 0x1a, 0x19, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0d, 0x53, 0x65, 0x6e,
	0x74, 0x69, 0x6e, 0x65, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x09, 0x6c, 0x6f,
	0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x65, 0x6e,
	0x74, 0x69, 0x6e, 0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x08, 0x6c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x8d, 0x03, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x49, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x65, 0x6e,
	0x74, 0x69, 0x6e, 0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x4d, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x3c, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x6e,
	0x74, 0x69, 0x6e, 0x65, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x42, 0xd2, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65,
	0x6c, 0x42, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x6e, 0x74,
	0x69, 0x6e, 0x65, 0x6c, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x53, 0xaa, 0x02, 0x15, 0x41, 0x70, 0x69,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6e,
	0x65, 0x6c, 0xca, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6c, 0xe2, 0x02, 0x21, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6e,
	0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x17, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a,
	0x53, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_sentinel_entities_proto_rawDescOnce sync.Once
	file_api_v1alpha1_sentinel_entities_proto_rawDescData = file_api_v1alpha1_sentinel_entities_proto_rawDesc
)

func file_api_v1alpha1_sentinel_entities_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_sentinel_entities_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_sentinel_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_sentinel_entities_proto_rawDescData)
	})
	return file_api_v1alpha1_sentinel_entities_proto_rawDescData
}

var file_api_v1alpha1_sentinel_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1alpha1_sentinel_entities_proto_goTypes = []any{
	(*SentinelEvent)(nil),         // 0: api.v1alpha1.sentinel.SentinelEvent
	(*LogEvent)(nil),              // 1: api.v1alpha1.sentinel.LogEvent
	(*SendEventsReq)(nil),         // 2: api.v1alpha1.sentinel.SendEventsReq
	(*SendEventsRes)(nil),         // 3: api.v1alpha1.sentinel.SendEventsRes
	nil,                           // 4: api.v1alpha1.sentinel.LogEvent.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(commons.Level)(0),            // 6: api.commons.Level
}
var file_api_v1alpha1_sentinel_entities_proto_depIdxs = []int32{
	1, // 0: api.v1alpha1.sentinel.SentinelEvent.log_event:type_name -> api.v1alpha1.sentinel.LogEvent
	5, // 1: api.v1alpha1.sentinel.LogEvent.timestamp:type_name -> google.protobuf.Timestamp
	4, // 2: api.v1alpha1.sentinel.LogEvent.metadata:type_name -> api.v1alpha1.sentinel.LogEvent.MetadataEntry
	6, // 3: api.v1alpha1.sentinel.LogEvent.severity:type_name -> api.commons.Level
	0, // 4: api.v1alpha1.sentinel.SendEventsReq.events:type_name -> api.v1alpha1.sentinel.SentinelEvent
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_sentinel_entities_proto_init() }
func file_api_v1alpha1_sentinel_entities_proto_init() {
	if File_api_v1alpha1_sentinel_entities_proto != nil {
		return
	}
	file_api_v1alpha1_sentinel_entities_proto_msgTypes[0].OneofWrappers = []any{
		(*SentinelEvent_LogEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_sentinel_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_sentinel_entities_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_sentinel_entities_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_sentinel_entities_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_sentinel_entities_proto = out.File
	file_api_v1alpha1_sentinel_entities_proto_rawDesc = nil
	file_api_v1alpha1_sentinel_entities_proto_goTypes = nil
	file_api_v1alpha1_sentinel_entities_proto_depIdxs = nil
}
