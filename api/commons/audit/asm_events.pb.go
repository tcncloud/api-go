// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: api/commons/audit/asm_events.proto

package audit

import (
	commons "github.com/tcncloud/api-go/api/commons"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// AsmAgentLoginEvent - whenever an agent logs into asm
type AsmAgentLoginEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the agent's ID
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// the asm session sid generated when the agent logged in
	AsmSessionSid int64 `protobuf:"varint,2,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AsmAgentLoginEvent) Reset() {
	*x = AsmAgentLoginEvent{}
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmAgentLoginEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmAgentLoginEvent) ProtoMessage() {}

func (x *AsmAgentLoginEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmAgentLoginEvent.ProtoReflect.Descriptor instead.
func (*AsmAgentLoginEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_asm_events_proto_rawDescGZIP(), []int{0}
}

func (x *AsmAgentLoginEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AsmAgentLoginEvent) GetAsmSessionSid() int64 {
	if x != nil {
		return x.AsmSessionSid
	}
	return 0
}

// AsmOpenVoiceEvent - whenever an agent opens voice
type AsmOpenVoiceEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the id of the agent
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// the session sid generated when agent logged into asm
	AsmSessionSid int64 `protobuf:"varint,2,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	// the session sid generated when voice started
	VoiceSessionSid int64 `protobuf:"varint,3,opt,name=voice_session_sid,json=voiceSessionSid,proto3" json:"voice_session_sid,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AsmOpenVoiceEvent) Reset() {
	*x = AsmOpenVoiceEvent{}
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmOpenVoiceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmOpenVoiceEvent) ProtoMessage() {}

func (x *AsmOpenVoiceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmOpenVoiceEvent.ProtoReflect.Descriptor instead.
func (*AsmOpenVoiceEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_asm_events_proto_rawDescGZIP(), []int{1}
}

func (x *AsmOpenVoiceEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AsmOpenVoiceEvent) GetAsmSessionSid() int64 {
	if x != nil {
		return x.AsmSessionSid
	}
	return 0
}

func (x *AsmOpenVoiceEvent) GetVoiceSessionSid() int64 {
	if x != nil {
		return x.VoiceSessionSid
	}
	return 0
}

// AsmOpenOmniAgentEvent - whenever an agent logs into omni agent
type AsmOpenOmniAgentEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the id of the agent
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// the session sid generated when the agent logged into asm
	AsmSessionSid int64 `protobuf:"varint,2,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AsmOpenOmniAgentEvent) Reset() {
	*x = AsmOpenOmniAgentEvent{}
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmOpenOmniAgentEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmOpenOmniAgentEvent) ProtoMessage() {}

func (x *AsmOpenOmniAgentEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmOpenOmniAgentEvent.ProtoReflect.Descriptor instead.
func (*AsmOpenOmniAgentEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_asm_events_proto_rawDescGZIP(), []int{2}
}

func (x *AsmOpenOmniAgentEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AsmOpenOmniAgentEvent) GetAsmSessionSid() int64 {
	if x != nil {
		return x.AsmSessionSid
	}
	return 0
}

// AsmActivateConversationEvent - whenever an agent activates a conversation
type AsmActivateConversationEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the id of the agent
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// the session sid generated when the agent logged into asm
	AsmSessionSid int64 `protobuf:"varint,2,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	// the conversation being activated
	Conversation  *commons.OmniConversation `protobuf:"bytes,3,opt,name=conversation,proto3" json:"conversation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AsmActivateConversationEvent) Reset() {
	*x = AsmActivateConversationEvent{}
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmActivateConversationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmActivateConversationEvent) ProtoMessage() {}

func (x *AsmActivateConversationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmActivateConversationEvent.ProtoReflect.Descriptor instead.
func (*AsmActivateConversationEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_asm_events_proto_rawDescGZIP(), []int{3}
}

func (x *AsmActivateConversationEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AsmActivateConversationEvent) GetAsmSessionSid() int64 {
	if x != nil {
		return x.AsmSessionSid
	}
	return 0
}

func (x *AsmActivateConversationEvent) GetConversation() *commons.OmniConversation {
	if x != nil {
		return x.Conversation
	}
	return nil
}

// AsmDeactivateConversationEvent - whenever agent deactivates a conversation
type AsmDeactivateConversationEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the id of the agent
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// the session sid generated when the agent logged into asm
	AsmSessionSid int64 `protobuf:"varint,2,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	// the conversation being deactivated
	Conversation  *commons.OmniConversation `protobuf:"bytes,3,opt,name=conversation,proto3" json:"conversation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AsmDeactivateConversationEvent) Reset() {
	*x = AsmDeactivateConversationEvent{}
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmDeactivateConversationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmDeactivateConversationEvent) ProtoMessage() {}

func (x *AsmDeactivateConversationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmDeactivateConversationEvent.ProtoReflect.Descriptor instead.
func (*AsmDeactivateConversationEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_asm_events_proto_rawDescGZIP(), []int{4}
}

func (x *AsmDeactivateConversationEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AsmDeactivateConversationEvent) GetAsmSessionSid() int64 {
	if x != nil {
		return x.AsmSessionSid
	}
	return 0
}

func (x *AsmDeactivateConversationEvent) GetConversation() *commons.OmniConversation {
	if x != nil {
		return x.Conversation
	}
	return nil
}

// AsmAgentStateChangedEvent - whenever an agent's state changes
type AsmAgentStateChangedEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the id of the agent
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// the session sid generated when the agent logged into asm
	AsmSessionSid int64 `protobuf:"varint,2,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	// the agent's new status
	NewStatus commons.StatusState `protobuf:"varint,3,opt,name=new_status,json=newStatus,proto3,enum=api.commons.StatusState" json:"new_status,omitempty"`
	// the agent's old status
	OldStatus commons.StatusState `protobuf:"varint,4,opt,name=old_status,json=oldStatus,proto3,enum=api.commons.StatusState" json:"old_status,omitempty"`
	// duration of old status in milliseconds
	OldStatusDurationMilliseconds int64 `protobuf:"varint,5,opt,name=old_status_duration_milliseconds,json=oldStatusDurationMilliseconds,proto3" json:"old_status_duration_milliseconds,omitempty"`
	unknownFields                 protoimpl.UnknownFields
	sizeCache                     protoimpl.SizeCache
}

func (x *AsmAgentStateChangedEvent) Reset() {
	*x = AsmAgentStateChangedEvent{}
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmAgentStateChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmAgentStateChangedEvent) ProtoMessage() {}

func (x *AsmAgentStateChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmAgentStateChangedEvent.ProtoReflect.Descriptor instead.
func (*AsmAgentStateChangedEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_asm_events_proto_rawDescGZIP(), []int{5}
}

func (x *AsmAgentStateChangedEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AsmAgentStateChangedEvent) GetAsmSessionSid() int64 {
	if x != nil {
		return x.AsmSessionSid
	}
	return 0
}

func (x *AsmAgentStateChangedEvent) GetNewStatus() commons.StatusState {
	if x != nil {
		return x.NewStatus
	}
	return commons.StatusState(0)
}

func (x *AsmAgentStateChangedEvent) GetOldStatus() commons.StatusState {
	if x != nil {
		return x.OldStatus
	}
	return commons.StatusState(0)
}

func (x *AsmAgentStateChangedEvent) GetOldStatusDurationMilliseconds() int64 {
	if x != nil {
		return x.OldStatusDurationMilliseconds
	}
	return 0
}

// AsmAgentLogoutEvent - whenever an agent logs out of asm
type AsmAgentLogoutEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the id of the agent
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// the session sid generated when the agent logged into asm
	AsmSessionSid int64 `protobuf:"varint,2,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	// the reason for logging out
	Reason        string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AsmAgentLogoutEvent) Reset() {
	*x = AsmAgentLogoutEvent{}
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmAgentLogoutEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmAgentLogoutEvent) ProtoMessage() {}

func (x *AsmAgentLogoutEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmAgentLogoutEvent.ProtoReflect.Descriptor instead.
func (*AsmAgentLogoutEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_asm_events_proto_rawDescGZIP(), []int{6}
}

func (x *AsmAgentLogoutEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AsmAgentLogoutEvent) GetAsmSessionSid() int64 {
	if x != nil {
		return x.AsmSessionSid
	}
	return 0
}

func (x *AsmAgentLogoutEvent) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

// AsmPauseEvent -
type AsmPauseEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the id of the agent
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// the session sid generated when the agent logged into asm
	AsmSessionSid int64 `protobuf:"varint,2,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AsmPauseEvent) Reset() {
	*x = AsmPauseEvent{}
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmPauseEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmPauseEvent) ProtoMessage() {}

func (x *AsmPauseEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmPauseEvent.ProtoReflect.Descriptor instead.
func (*AsmPauseEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_asm_events_proto_rawDescGZIP(), []int{7}
}

func (x *AsmPauseEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AsmPauseEvent) GetAsmSessionSid() int64 {
	if x != nil {
		return x.AsmSessionSid
	}
	return 0
}

// AsmResumeEvent -
type AsmResumeEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the id of the agent
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// the session sid generated when the agent logged into asm
	AsmSessionSid int64 `protobuf:"varint,2,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AsmResumeEvent) Reset() {
	*x = AsmResumeEvent{}
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmResumeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmResumeEvent) ProtoMessage() {}

func (x *AsmResumeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmResumeEvent.ProtoReflect.Descriptor instead.
func (*AsmResumeEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_asm_events_proto_rawDescGZIP(), []int{8}
}

func (x *AsmResumeEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AsmResumeEvent) GetAsmSessionSid() int64 {
	if x != nil {
		return x.AsmSessionSid
	}
	return 0
}

// AsmConversationPulledEvent -
type AsmConversationPulledEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the conversation pulled
	Conversation  *commons.OmniConversation `protobuf:"bytes,1,opt,name=conversation,proto3" json:"conversation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AsmConversationPulledEvent) Reset() {
	*x = AsmConversationPulledEvent{}
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmConversationPulledEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmConversationPulledEvent) ProtoMessage() {}

func (x *AsmConversationPulledEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_asm_events_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmConversationPulledEvent.ProtoReflect.Descriptor instead.
func (*AsmConversationPulledEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_asm_events_proto_rawDescGZIP(), []int{9}
}

func (x *AsmConversationPulledEvent) GetConversation() *commons.OmniConversation {
	if x != nil {
		return x.Conversation
	}
	return nil
}

var File_api_commons_audit_asm_events_proto protoreflect.FileDescriptor

var file_api_commons_audit_asm_events_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x61, 0x73, 0x6d, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55,
	0x0a, 0x12, 0x41, 0x73, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0f, 0x61, 0x73, 0x6d, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x73, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x69, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x11, 0x41, 0x73, 0x6d, 0x4f, 0x70, 0x65,
	0x6e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x73, 0x6d, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61,
	0x73, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x15, 0x41, 0x73, 0x6d, 0x4f,
	0x70, 0x65, 0x6e, 0x4f, 0x6d, 0x6e, 0x69, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x73,
	0x6d, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x73, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x69, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x1c, 0x41, 0x73, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x61, 0x73, 0x6d, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x73, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x69, 0x64, 0x12, 0x41, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x6d, 0x6e, 0x69, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa4, 0x01, 0x0a, 0x1e, 0x41, 0x73, 0x6d, 0x44,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x73, 0x6d, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x73,
	0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x12, 0x41, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x4f, 0x6d, 0x6e, 0x69, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x97,
	0x02, 0x0a, 0x19, 0x41, 0x73, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x73, 0x6d, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x61, 0x73, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x12, 0x37, 0x0a,
	0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x6e, 0x65, 0x77,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x6f, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x47, 0x0a, 0x20, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1d, 0x6f, 0x6c, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x6e, 0x0a, 0x13, 0x41, 0x73, 0x6d, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x73, 0x6d, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x61, 0x73, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x0d, 0x41, 0x73, 0x6d, 0x50,
	0x61, 0x75, 0x73, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x73, 0x6d, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x73, 0x6d,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x0e, 0x41, 0x73,
	0x6d, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x73, 0x6d, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x61, 0x73, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x22, 0x5f, 0x0a,
	0x1a, 0x41, 0x73, 0x6d, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x75, 0x6c, 0x6c, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x4f, 0x6d, 0x6e, 0x69, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xbb,
	0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0x0e, 0x41, 0x73, 0x6d, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x41, 0xaa, 0x02,
	0x11, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x74, 0xca, 0x02, 0x11, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0xe2, 0x02, 0x1d, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_commons_audit_asm_events_proto_rawDescOnce sync.Once
	file_api_commons_audit_asm_events_proto_rawDescData []byte
)

func file_api_commons_audit_asm_events_proto_rawDescGZIP() []byte {
	file_api_commons_audit_asm_events_proto_rawDescOnce.Do(func() {
		file_api_commons_audit_asm_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_audit_asm_events_proto_rawDesc), len(file_api_commons_audit_asm_events_proto_rawDesc)))
	})
	return file_api_commons_audit_asm_events_proto_rawDescData
}

var file_api_commons_audit_asm_events_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_commons_audit_asm_events_proto_goTypes = []any{
	(*AsmAgentLoginEvent)(nil),             // 0: api.commons.audit.AsmAgentLoginEvent
	(*AsmOpenVoiceEvent)(nil),              // 1: api.commons.audit.AsmOpenVoiceEvent
	(*AsmOpenOmniAgentEvent)(nil),          // 2: api.commons.audit.AsmOpenOmniAgentEvent
	(*AsmActivateConversationEvent)(nil),   // 3: api.commons.audit.AsmActivateConversationEvent
	(*AsmDeactivateConversationEvent)(nil), // 4: api.commons.audit.AsmDeactivateConversationEvent
	(*AsmAgentStateChangedEvent)(nil),      // 5: api.commons.audit.AsmAgentStateChangedEvent
	(*AsmAgentLogoutEvent)(nil),            // 6: api.commons.audit.AsmAgentLogoutEvent
	(*AsmPauseEvent)(nil),                  // 7: api.commons.audit.AsmPauseEvent
	(*AsmResumeEvent)(nil),                 // 8: api.commons.audit.AsmResumeEvent
	(*AsmConversationPulledEvent)(nil),     // 9: api.commons.audit.AsmConversationPulledEvent
	(*commons.OmniConversation)(nil),       // 10: api.commons.OmniConversation
	(commons.StatusState)(0),               // 11: api.commons.StatusState
}
var file_api_commons_audit_asm_events_proto_depIdxs = []int32{
	10, // 0: api.commons.audit.AsmActivateConversationEvent.conversation:type_name -> api.commons.OmniConversation
	10, // 1: api.commons.audit.AsmDeactivateConversationEvent.conversation:type_name -> api.commons.OmniConversation
	11, // 2: api.commons.audit.AsmAgentStateChangedEvent.new_status:type_name -> api.commons.StatusState
	11, // 3: api.commons.audit.AsmAgentStateChangedEvent.old_status:type_name -> api.commons.StatusState
	10, // 4: api.commons.audit.AsmConversationPulledEvent.conversation:type_name -> api.commons.OmniConversation
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_commons_audit_asm_events_proto_init() }
func file_api_commons_audit_asm_events_proto_init() {
	if File_api_commons_audit_asm_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_audit_asm_events_proto_rawDesc), len(file_api_commons_audit_asm_events_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_audit_asm_events_proto_goTypes,
		DependencyIndexes: file_api_commons_audit_asm_events_proto_depIdxs,
		MessageInfos:      file_api_commons_audit_asm_events_proto_msgTypes,
	}.Build()
	File_api_commons_audit_asm_events_proto = out.File
	file_api_commons_audit_asm_events_proto_goTypes = nil
	file_api_commons_audit_asm_events_proto_depIdxs = nil
}
