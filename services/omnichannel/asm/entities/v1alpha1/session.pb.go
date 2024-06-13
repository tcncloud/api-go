// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: services/omnichannel/asm/entities/v1alpha1/session.proto

package entitiesv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusState int32

const (
	// unknown
	StatusState_STATUS_STATE_UNKNOWN StatusState = 0
	// no conversations assigned
	StatusState_WAITING StatusState = 1
	// conversations assigned but none open
	StatusState_IDLE StatusState = 2
	// conversation is open
	StatusState_CONVERSATION_OPEN StatusState = 3
)

// Enum value maps for StatusState.
var (
	StatusState_name = map[int32]string{
		0: "STATUS_STATE_UNKNOWN",
		1: "WAITING",
		2: "IDLE",
		3: "CONVERSATION_OPEN",
	}
	StatusState_value = map[string]int32{
		"STATUS_STATE_UNKNOWN": 0,
		"WAITING":              1,
		"IDLE":                 2,
		"CONVERSATION_OPEN":    3,
	}
)

func (x StatusState) Enum() *StatusState {
	p := new(StatusState)
	*p = x
	return p
}

func (x StatusState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusState) Descriptor() protoreflect.EnumDescriptor {
	return file_services_omnichannel_asm_entities_v1alpha1_session_proto_enumTypes[0].Descriptor()
}

func (StatusState) Type() protoreflect.EnumType {
	return &file_services_omnichannel_asm_entities_v1alpha1_session_proto_enumTypes[0]
}

func (x StatusState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusState.Descriptor instead.
func (StatusState) EnumDescriptor() ([]byte, []int) {
	return file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescGZIP(), []int{0}
}

type AsmSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// asm session sid
	AsmSessionSid int64 `protobuf:"varint,1,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	// time the asm session started
	AsmSessionStart *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=asm_session_start,json=asmSessionStart,proto3" json:"asm_session_start,omitempty"`
	// time the asm session ended
	AsmSessionEnd *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=asm_session_end,json=asmSessionEnd,proto3" json:"asm_session_end,omitempty"`
	// voice session is there is one
	VoiceSession *VoiceSession `protobuf:"bytes,6,opt,name=voice_session,json=voiceSession,proto3" json:"voice_session,omitempty"`
}

func (x *AsmSession) Reset() {
	*x = AsmSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsmSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmSession) ProtoMessage() {}

func (x *AsmSession) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmSession.ProtoReflect.Descriptor instead.
func (*AsmSession) Descriptor() ([]byte, []int) {
	return file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescGZIP(), []int{0}
}

func (x *AsmSession) GetAsmSessionSid() int64 {
	if x != nil {
		return x.AsmSessionSid
	}
	return 0
}

func (x *AsmSession) GetAsmSessionStart() *timestamppb.Timestamp {
	if x != nil {
		return x.AsmSessionStart
	}
	return nil
}

func (x *AsmSession) GetAsmSessionEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.AsmSessionEnd
	}
	return nil
}

func (x *AsmSession) GetVoiceSession() *VoiceSession {
	if x != nil {
		return x.VoiceSession
	}
	return nil
}

type VoiceSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// voice session sid
	VoiceSessionSid int64 `protobuf:"varint,1,opt,name=voice_session_sid,json=voiceSessionSid,proto3" json:"voice_session_sid,omitempty"`
	// time the voice session started
	VoiceSessionStart *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=voice_session_start,json=voiceSessionStart,proto3" json:"voice_session_start,omitempty"`
	// time the voice session ended
	VoiceSessionEnd *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=voice_session_end,json=voiceSessionEnd,proto3" json:"voice_session_end,omitempty"`
}

func (x *VoiceSession) Reset() {
	*x = VoiceSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceSession) ProtoMessage() {}

func (x *VoiceSession) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceSession.ProtoReflect.Descriptor instead.
func (*VoiceSession) Descriptor() ([]byte, []int) {
	return file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescGZIP(), []int{1}
}

func (x *VoiceSession) GetVoiceSessionSid() int64 {
	if x != nil {
		return x.VoiceSessionSid
	}
	return 0
}

func (x *VoiceSession) GetVoiceSessionStart() *timestamppb.Timestamp {
	if x != nil {
		return x.VoiceSessionStart
	}
	return nil
}

func (x *VoiceSession) GetVoiceSessionEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.VoiceSessionEnd
	}
	return nil
}

type VoiceRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The pin used to log in via a connected phone
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// The pass used to log in via a connected phone
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// The extention appended
	DialUrl string `protobuf:"bytes,4,opt,name=dial_url,json=dialUrl,proto3" json:"dial_url,omitempty"`
	// pstn phone number that will be used for the agent to dial in
	// if it's an empty string then the voip connection must be used
	PstnPhone string `protobuf:"bytes,5,opt,name=pstn_phone,json=pstnPhone,proto3" json:"pstn_phone,omitempty"`
	// default time zone -
	DefaultTimeZone string `protobuf:"bytes,6,opt,name=default_time_zone,json=defaultTimeZone,proto3" json:"default_time_zone,omitempty"`
	// expiration Timestamp When the registration will expire
	ExpirationTimestamp int64 `protobuf:"varint,7,opt,name=expiration_timestamp,json=expirationTimestamp,proto3" json:"expiration_timestamp,omitempty"`
}

func (x *VoiceRegistration) Reset() {
	*x = VoiceRegistration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceRegistration) ProtoMessage() {}

func (x *VoiceRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceRegistration.ProtoReflect.Descriptor instead.
func (*VoiceRegistration) Descriptor() ([]byte, []int) {
	return file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescGZIP(), []int{2}
}

func (x *VoiceRegistration) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *VoiceRegistration) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *VoiceRegistration) GetDialUrl() string {
	if x != nil {
		return x.DialUrl
	}
	return ""
}

func (x *VoiceRegistration) GetPstnPhone() string {
	if x != nil {
		return x.PstnPhone
	}
	return ""
}

func (x *VoiceRegistration) GetDefaultTimeZone() string {
	if x != nil {
		return x.DefaultTimeZone
	}
	return ""
}

func (x *VoiceRegistration) GetExpirationTimestamp() int64 {
	if x != nil {
		return x.ExpirationTimestamp
	}
	return 0
}

type AsmUserDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// agents user id
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// agents name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// the agents sid
	AgentSid int64 `protobuf:"varint,3,opt,name=agent_sid,json=agentSid,proto3" json:"agent_sid,omitempty"`
	// enum of agents status
	AgentStatus StatusState `protobuf:"varint,4,opt,name=agent_status,json=agentStatus,proto3,enum=services.omnichannel.asm.entities.v1alpha1.StatusState" json:"agent_status,omitempty"`
	// agents profile group name
	AgentProfileGroupName string `protobuf:"bytes,5,opt,name=agent_profile_group_name,json=agentProfileGroupName,proto3" json:"agent_profile_group_name,omitempty"`
	// the agents current conversation
	CurrentConversationSid int64 `protobuf:"varint,6,opt,name=current_conversation_sid,json=currentConversationSid,proto3" json:"current_conversation_sid,omitempty"`
	// time from first customer message to agent response. between all conversations.
	AverageCustomerWaitTimeSeconds int64 `protobuf:"varint,7,opt,name=average_customer_wait_time_seconds,json=averageCustomerWaitTimeSeconds,proto3" json:"average_customer_wait_time_seconds,omitempty"`
	// responste time between all conversations.
	AverageTimeToRespondSeconds int64 `protobuf:"varint,8,opt,name=average_time_to_respond_seconds,json=averageTimeToRespondSeconds,proto3" json:"average_time_to_respond_seconds,omitempty"`
	// last event time
	LastEventTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=last_event_time,json=lastEventTime,proto3" json:"last_event_time,omitempty"`
	// Agents Skills
	Skills map[string]int64 `protobuf:"bytes,10,rep,name=skills,proto3" json:"skills,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Asm Session Sid
	AsmSessionSid *wrapperspb.Int64Value `protobuf:"bytes,11,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	// the list of response events
	Events []*DashboardAgentResponseEvent `protobuf:"bytes,12,rep,name=events,proto3" json:"events,omitempty"`
	// login time
	LoginTime *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=login_time,json=loginTime,proto3" json:"login_time,omitempty"`
}

func (x *AsmUserDetails) Reset() {
	*x = AsmUserDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsmUserDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmUserDetails) ProtoMessage() {}

func (x *AsmUserDetails) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsmUserDetails.ProtoReflect.Descriptor instead.
func (*AsmUserDetails) Descriptor() ([]byte, []int) {
	return file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescGZIP(), []int{3}
}

func (x *AsmUserDetails) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AsmUserDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AsmUserDetails) GetAgentSid() int64 {
	if x != nil {
		return x.AgentSid
	}
	return 0
}

func (x *AsmUserDetails) GetAgentStatus() StatusState {
	if x != nil {
		return x.AgentStatus
	}
	return StatusState_STATUS_STATE_UNKNOWN
}

func (x *AsmUserDetails) GetAgentProfileGroupName() string {
	if x != nil {
		return x.AgentProfileGroupName
	}
	return ""
}

func (x *AsmUserDetails) GetCurrentConversationSid() int64 {
	if x != nil {
		return x.CurrentConversationSid
	}
	return 0
}

func (x *AsmUserDetails) GetAverageCustomerWaitTimeSeconds() int64 {
	if x != nil {
		return x.AverageCustomerWaitTimeSeconds
	}
	return 0
}

func (x *AsmUserDetails) GetAverageTimeToRespondSeconds() int64 {
	if x != nil {
		return x.AverageTimeToRespondSeconds
	}
	return 0
}

func (x *AsmUserDetails) GetLastEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastEventTime
	}
	return nil
}

func (x *AsmUserDetails) GetSkills() map[string]int64 {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *AsmUserDetails) GetAsmSessionSid() *wrapperspb.Int64Value {
	if x != nil {
		return x.AsmSessionSid
	}
	return nil
}

func (x *AsmUserDetails) GetEvents() []*DashboardAgentResponseEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *AsmUserDetails) GetLoginTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LoginTime
	}
	return nil
}

// The response event information for an agent
type DashboardAgentResponseEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the time it took to respond in seconds
	ResponseTimeSeconds int64 `protobuf:"varint,1,opt,name=response_time_seconds,json=responseTimeSeconds,proto3" json:"response_time_seconds,omitempty"`
	// the time of the event
	Time *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// whether or not this was the initial message sent by the agent
	IsInitialAgentMessage bool `protobuf:"varint,3,opt,name=is_initial_agent_message,json=isInitialAgentMessage,proto3" json:"is_initial_agent_message,omitempty"`
}

func (x *DashboardAgentResponseEvent) Reset() {
	*x = DashboardAgentResponseEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DashboardAgentResponseEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardAgentResponseEvent) ProtoMessage() {}

func (x *DashboardAgentResponseEvent) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardAgentResponseEvent.ProtoReflect.Descriptor instead.
func (*DashboardAgentResponseEvent) Descriptor() ([]byte, []int) {
	return file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescGZIP(), []int{4}
}

func (x *DashboardAgentResponseEvent) GetResponseTimeSeconds() int64 {
	if x != nil {
		return x.ResponseTimeSeconds
	}
	return 0
}

func (x *DashboardAgentResponseEvent) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *DashboardAgentResponseEvent) GetIsInitialAgentMessage() bool {
	if x != nil {
		return x.IsInitialAgentMessage
	}
	return false
}

var File_services_omnichannel_asm_entities_v1alpha1_session_proto protoreflect.FileDescriptor

var file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDesc = []byte{
	0x0a, 0x38, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x73, 0x6d, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x0a, 0x41, 0x73, 0x6d, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x73, 0x6d, 0x5f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x61, 0x73, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x12, 0x46,
	0x0a, 0x11, 0x61, 0x73, 0x6d, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x61, 0x73, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x61, 0x73, 0x6d, 0x5f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x61, 0x73, 0x6d,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x12, 0x5d, 0x0a, 0x0d, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e,
	0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xce, 0x01, 0x0a, 0x0c, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x13, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x11, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x46, 0x0a, 0x11, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x22, 0xe4, 0x01, 0x0a, 0x11, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x69, 0x61, 0x6c,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x69, 0x61, 0x6c,
	0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x73, 0x74, 0x6e, 0x5f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x73, 0x74, 0x6e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x31,
	0x0a, 0x14, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0xff, 0x06, 0x0a, 0x0e, 0x41, 0x73, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x12, 0x5a,
	0x0a, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x18, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x16, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69,
	0x64, 0x12, 0x4a, 0x0a, 0x22, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1e, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x57, 0x61,
	0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x44, 0x0a,
	0x1f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x6f,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5e, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61,
	0x73, 0x6d, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x73, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x43, 0x0a, 0x0f, 0x61, 0x73, 0x6d, 0x5f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x61,
	0x73, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x12, 0x5f, 0x0a, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xba, 0x01, 0x0a, 0x1b, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x69, 0x73, 0x5f, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x73, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2a, 0x55, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x44, 0x4c, 0x45, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x03, 0x42, 0xe3, 0x02, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x73, 0x6d, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x04, 0x53, 0x4f, 0x41, 0x45, 0xaa, 0x02, 0x2a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x41, 0x73, 0x6d, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x2a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x5c, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5c, 0x41, 0x73,
	0x6d, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x36, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x4f,
	0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5c, 0x41, 0x73, 0x6d, 0x5c, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x3a, 0x3a, 0x41, 0x73, 0x6d, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescOnce sync.Once
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescData = file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDesc
)

func file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescGZIP() []byte {
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescOnce.Do(func() {
		file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescData)
	})
	return file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescData
}

var file_services_omnichannel_asm_entities_v1alpha1_session_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_omnichannel_asm_entities_v1alpha1_session_proto_goTypes = []any{
	(StatusState)(0),                    // 0: services.omnichannel.asm.entities.v1alpha1.StatusState
	(*AsmSession)(nil),                  // 1: services.omnichannel.asm.entities.v1alpha1.AsmSession
	(*VoiceSession)(nil),                // 2: services.omnichannel.asm.entities.v1alpha1.VoiceSession
	(*VoiceRegistration)(nil),           // 3: services.omnichannel.asm.entities.v1alpha1.VoiceRegistration
	(*AsmUserDetails)(nil),              // 4: services.omnichannel.asm.entities.v1alpha1.AsmUserDetails
	(*DashboardAgentResponseEvent)(nil), // 5: services.omnichannel.asm.entities.v1alpha1.DashboardAgentResponseEvent
	nil,                                 // 6: services.omnichannel.asm.entities.v1alpha1.AsmUserDetails.SkillsEntry
	(*timestamppb.Timestamp)(nil),       // 7: google.protobuf.Timestamp
	(*wrapperspb.Int64Value)(nil),       // 8: google.protobuf.Int64Value
}
var file_services_omnichannel_asm_entities_v1alpha1_session_proto_depIdxs = []int32{
	7,  // 0: services.omnichannel.asm.entities.v1alpha1.AsmSession.asm_session_start:type_name -> google.protobuf.Timestamp
	7,  // 1: services.omnichannel.asm.entities.v1alpha1.AsmSession.asm_session_end:type_name -> google.protobuf.Timestamp
	2,  // 2: services.omnichannel.asm.entities.v1alpha1.AsmSession.voice_session:type_name -> services.omnichannel.asm.entities.v1alpha1.VoiceSession
	7,  // 3: services.omnichannel.asm.entities.v1alpha1.VoiceSession.voice_session_start:type_name -> google.protobuf.Timestamp
	7,  // 4: services.omnichannel.asm.entities.v1alpha1.VoiceSession.voice_session_end:type_name -> google.protobuf.Timestamp
	0,  // 5: services.omnichannel.asm.entities.v1alpha1.AsmUserDetails.agent_status:type_name -> services.omnichannel.asm.entities.v1alpha1.StatusState
	7,  // 6: services.omnichannel.asm.entities.v1alpha1.AsmUserDetails.last_event_time:type_name -> google.protobuf.Timestamp
	6,  // 7: services.omnichannel.asm.entities.v1alpha1.AsmUserDetails.skills:type_name -> services.omnichannel.asm.entities.v1alpha1.AsmUserDetails.SkillsEntry
	8,  // 8: services.omnichannel.asm.entities.v1alpha1.AsmUserDetails.asm_session_sid:type_name -> google.protobuf.Int64Value
	5,  // 9: services.omnichannel.asm.entities.v1alpha1.AsmUserDetails.events:type_name -> services.omnichannel.asm.entities.v1alpha1.DashboardAgentResponseEvent
	7,  // 10: services.omnichannel.asm.entities.v1alpha1.AsmUserDetails.login_time:type_name -> google.protobuf.Timestamp
	7,  // 11: services.omnichannel.asm.entities.v1alpha1.DashboardAgentResponseEvent.time:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_services_omnichannel_asm_entities_v1alpha1_session_proto_init() }
func file_services_omnichannel_asm_entities_v1alpha1_session_proto_init() {
	if File_services_omnichannel_asm_entities_v1alpha1_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AsmSession); i {
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
		file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*VoiceSession); i {
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
		file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*VoiceRegistration); i {
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
		file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AsmUserDetails); i {
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
		file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DashboardAgentResponseEvent); i {
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
			RawDescriptor: file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_omnichannel_asm_entities_v1alpha1_session_proto_goTypes,
		DependencyIndexes: file_services_omnichannel_asm_entities_v1alpha1_session_proto_depIdxs,
		EnumInfos:         file_services_omnichannel_asm_entities_v1alpha1_session_proto_enumTypes,
		MessageInfos:      file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes,
	}.Build()
	File_services_omnichannel_asm_entities_v1alpha1_session_proto = out.File
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDesc = nil
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_goTypes = nil
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_depIdxs = nil
}
