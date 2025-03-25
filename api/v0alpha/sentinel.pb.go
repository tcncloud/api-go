// Copyright (c) 2019, TCN Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v0alpha/sentinel.proto

package v0alpha

import (
	_ "github.com/tcncloud/api-go/annotations"
	commons "github.com/tcncloud/api-go/api/commons"
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

// sentinel event
type SentinelEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Event:
	//
	//	*SentinelEvent_LogEvent
	Event         isSentinelEvent_Event `protobuf_oneof:"event"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SentinelEvent) Reset() {
	*x = SentinelEvent{}
	mi := &file_api_v0alpha_sentinel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SentinelEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentinelEvent) ProtoMessage() {}

func (x *SentinelEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sentinel_proto_msgTypes[0]
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
	return file_api_v0alpha_sentinel_proto_rawDescGZIP(), []int{0}
}

func (x *SentinelEvent) GetEvent() isSentinelEvent_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *SentinelEvent) GetLogEvent() *LogEvent {
	if x != nil {
		if x, ok := x.Event.(*SentinelEvent_LogEvent); ok {
			return x.LogEvent
		}
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

// log events details to be sent
type LogEvent struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	TraceId string                 `protobuf:"bytes,3,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// the browsers session id
	SessionId string `protobuf:"bytes,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Message   string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	// the browsers current uri
	Location      string                 `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	StackTrace    string                 `protobuf:"bytes,7,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Metadata      map[string]string      `protobuf:"bytes,9,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Severity      commons.Level          `protobuf:"varint,10,opt,name=severity,proto3,enum=api.commons.Level" json:"severity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogEvent) Reset() {
	*x = LogEvent{}
	mi := &file_api_v0alpha_sentinel_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEvent) ProtoMessage() {}

func (x *LogEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sentinel_proto_msgTypes[1]
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
	return file_api_v0alpha_sentinel_proto_rawDescGZIP(), []int{1}
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

// request to sentinel send events
type SendEventsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Events        []*SentinelEvent       `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendEventsReq) Reset() {
	*x = SendEventsReq{}
	mi := &file_api_v0alpha_sentinel_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendEventsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEventsReq) ProtoMessage() {}

func (x *SendEventsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sentinel_proto_msgTypes[2]
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
	return file_api_v0alpha_sentinel_proto_rawDescGZIP(), []int{2}
}

func (x *SendEventsReq) GetEvents() []*SentinelEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

// response to sentinel send events
type SendEventsRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendEventsRes) Reset() {
	*x = SendEventsRes{}
	mi := &file_api_v0alpha_sentinel_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendEventsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEventsRes) ProtoMessage() {}

func (x *SendEventsRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sentinel_proto_msgTypes[3]
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
	return file_api_v0alpha_sentinel_proto_rawDescGZIP(), []int{3}
}

var File_api_v0alpha_sentinel_proto protoreflect.FileDescriptor

const file_api_v0alpha_sentinel_proto_rawDesc = "" +
	"\n" +
	"\x1aapi/v0alpha/sentinel.proto\x12\vapi.v0alpha\x1a\x17annotations/authz.proto\x1a\x19api/commons/logging.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"N\n" +
	"\rSentinelEvent\x124\n" +
	"\tlog_event\x18\x01 \x01(\v2\x15.api.v0alpha.LogEventH\x00R\blogEventB\a\n" +
	"\x05event\"\x83\x03\n" +
	"\bLogEvent\x12\x19\n" +
	"\btrace_id\x18\x03 \x01(\tR\atraceId\x12\x1d\n" +
	"\n" +
	"session_id\x18\x04 \x01(\tR\tsessionId\x12\x18\n" +
	"\amessage\x18\x05 \x01(\tR\amessage\x12\x1a\n" +
	"\blocation\x18\x06 \x01(\tR\blocation\x12\x1f\n" +
	"\vstack_trace\x18\a \x01(\tR\n" +
	"stackTrace\x128\n" +
	"\ttimestamp\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12?\n" +
	"\bmetadata\x18\t \x03(\v2#.api.v0alpha.LogEvent.MetadataEntryR\bmetadata\x12.\n" +
	"\bseverity\x18\n" +
	" \x01(\x0e2\x12.api.commons.LevelR\bseverity\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"C\n" +
	"\rSendEventsReq\x122\n" +
	"\x06events\x18\x01 \x03(\v2\x1a.api.v0alpha.SentinelEventR\x06events\"\x0f\n" +
	"\rSendEventsRes2\x84\x01\n" +
	"\bSentinel\x12x\n" +
	"\n" +
	"SendEvents\x12\x1a.api.v0alpha.SendEventsReq\x1a\x1a.api.v0alpha.SendEventsRes\"2\xba\xb8\x91\x02\x02\x18\x01\x82\xd3\xe4\x93\x02%:\x01*\" /api/v0alpha/sentinel/sendeventsB\x95\x01\n" +
	"\x0fcom.api.v0alphaB\rSentinelProtoP\x01Z&github.com/tcncloud/api-go/api/v0alpha\xa2\x02\x03AVX\xaa\x02\vApi.V0alpha\xca\x02\vApi\\V0alpha\xe2\x02\x17Api\\V0alpha\\GPBMetadata\xea\x02\fApi::V0alphab\x06proto3"

var (
	file_api_v0alpha_sentinel_proto_rawDescOnce sync.Once
	file_api_v0alpha_sentinel_proto_rawDescData []byte
)

func file_api_v0alpha_sentinel_proto_rawDescGZIP() []byte {
	file_api_v0alpha_sentinel_proto_rawDescOnce.Do(func() {
		file_api_v0alpha_sentinel_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v0alpha_sentinel_proto_rawDesc), len(file_api_v0alpha_sentinel_proto_rawDesc)))
	})
	return file_api_v0alpha_sentinel_proto_rawDescData
}

var file_api_v0alpha_sentinel_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v0alpha_sentinel_proto_goTypes = []any{
	(*SentinelEvent)(nil),         // 0: api.v0alpha.SentinelEvent
	(*LogEvent)(nil),              // 1: api.v0alpha.LogEvent
	(*SendEventsReq)(nil),         // 2: api.v0alpha.SendEventsReq
	(*SendEventsRes)(nil),         // 3: api.v0alpha.SendEventsRes
	nil,                           // 4: api.v0alpha.LogEvent.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(commons.Level)(0),            // 6: api.commons.Level
}
var file_api_v0alpha_sentinel_proto_depIdxs = []int32{
	1, // 0: api.v0alpha.SentinelEvent.log_event:type_name -> api.v0alpha.LogEvent
	5, // 1: api.v0alpha.LogEvent.timestamp:type_name -> google.protobuf.Timestamp
	4, // 2: api.v0alpha.LogEvent.metadata:type_name -> api.v0alpha.LogEvent.MetadataEntry
	6, // 3: api.v0alpha.LogEvent.severity:type_name -> api.commons.Level
	0, // 4: api.v0alpha.SendEventsReq.events:type_name -> api.v0alpha.SentinelEvent
	2, // 5: api.v0alpha.Sentinel.SendEvents:input_type -> api.v0alpha.SendEventsReq
	3, // 6: api.v0alpha.Sentinel.SendEvents:output_type -> api.v0alpha.SendEventsRes
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_v0alpha_sentinel_proto_init() }
func file_api_v0alpha_sentinel_proto_init() {
	if File_api_v0alpha_sentinel_proto != nil {
		return
	}
	file_api_v0alpha_sentinel_proto_msgTypes[0].OneofWrappers = []any{
		(*SentinelEvent_LogEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v0alpha_sentinel_proto_rawDesc), len(file_api_v0alpha_sentinel_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v0alpha_sentinel_proto_goTypes,
		DependencyIndexes: file_api_v0alpha_sentinel_proto_depIdxs,
		MessageInfos:      file_api_v0alpha_sentinel_proto_msgTypes,
	}.Build()
	File_api_v0alpha_sentinel_proto = out.File
	file_api_v0alpha_sentinel_proto_goTypes = nil
	file_api_v0alpha_sentinel_proto_depIdxs = nil
}
