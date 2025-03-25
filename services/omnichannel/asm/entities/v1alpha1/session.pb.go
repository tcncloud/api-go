// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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
	unsafe "unsafe"
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// asm session sid
	AsmSessionSid int64 `protobuf:"varint,1,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	// time the asm session started
	AsmSessionStart *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=asm_session_start,json=asmSessionStart,proto3" json:"asm_session_start,omitempty"`
	// time the asm session ended
	AsmSessionEnd *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=asm_session_end,json=asmSessionEnd,proto3" json:"asm_session_end,omitempty"`
	// voice session is there is one
	VoiceSession  *VoiceSession `protobuf:"bytes,6,opt,name=voice_session,json=voiceSession,proto3" json:"voice_session,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AsmSession) Reset() {
	*x = AsmSession{}
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmSession) ProtoMessage() {}

func (x *AsmSession) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[0]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// voice session sid
	VoiceSessionSid int64 `protobuf:"varint,1,opt,name=voice_session_sid,json=voiceSessionSid,proto3" json:"voice_session_sid,omitempty"`
	// time the voice session started
	VoiceSessionStart *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=voice_session_start,json=voiceSessionStart,proto3" json:"voice_session_start,omitempty"`
	// time the voice session ended
	VoiceSessionEnd *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=voice_session_end,json=voiceSessionEnd,proto3" json:"voice_session_end,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *VoiceSession) Reset() {
	*x = VoiceSession{}
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoiceSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceSession) ProtoMessage() {}

func (x *VoiceSession) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[1]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
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
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *VoiceRegistration) Reset() {
	*x = VoiceRegistration{}
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoiceRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceRegistration) ProtoMessage() {}

func (x *VoiceRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[2]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
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
	Skills map[string]int64 `protobuf:"bytes,10,rep,name=skills,proto3" json:"skills,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Asm Session Sid
	AsmSessionSid *wrapperspb.Int64Value `protobuf:"bytes,11,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	// the list of response events
	Events []*DashboardAgentResponseEvent `protobuf:"bytes,12,rep,name=events,proto3" json:"events,omitempty"`
	// login time
	LoginTime     *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=login_time,json=loginTime,proto3" json:"login_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AsmUserDetails) Reset() {
	*x = AsmUserDetails{}
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsmUserDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsmUserDetails) ProtoMessage() {}

func (x *AsmUserDetails) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[3]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// the time it took to respond in seconds
	ResponseTimeSeconds int64 `protobuf:"varint,1,opt,name=response_time_seconds,json=responseTimeSeconds,proto3" json:"response_time_seconds,omitempty"`
	// the time of the event
	Time *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// whether or not this was the initial message sent by the agent
	IsInitialAgentMessage bool `protobuf:"varint,3,opt,name=is_initial_agent_message,json=isInitialAgentMessage,proto3" json:"is_initial_agent_message,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *DashboardAgentResponseEvent) Reset() {
	*x = DashboardAgentResponseEvent{}
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DashboardAgentResponseEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardAgentResponseEvent) ProtoMessage() {}

func (x *DashboardAgentResponseEvent) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[4]
	if x != nil {
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

const file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDesc = "" +
	"\n" +
	"8services/omnichannel/asm/entities/v1alpha1/session.proto\x12*services.omnichannel.asm.entities.v1alpha1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\x9f\x02\n" +
	"\n" +
	"AsmSession\x12&\n" +
	"\x0fasm_session_sid\x18\x01 \x01(\x03R\rasmSessionSid\x12F\n" +
	"\x11asm_session_start\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\x0fasmSessionStart\x12B\n" +
	"\x0fasm_session_end\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\rasmSessionEnd\x12]\n" +
	"\rvoice_session\x18\x06 \x01(\v28.services.omnichannel.asm.entities.v1alpha1.VoiceSessionR\fvoiceSession\"\xce\x01\n" +
	"\fVoiceSession\x12*\n" +
	"\x11voice_session_sid\x18\x01 \x01(\x03R\x0fvoiceSessionSid\x12J\n" +
	"\x13voice_session_start\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x11voiceSessionStart\x12F\n" +
	"\x11voice_session_end\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\x0fvoiceSessionEnd\"\xe4\x01\n" +
	"\x11VoiceRegistration\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12\x19\n" +
	"\bdial_url\x18\x04 \x01(\tR\adialUrl\x12\x1d\n" +
	"\n" +
	"pstn_phone\x18\x05 \x01(\tR\tpstnPhone\x12*\n" +
	"\x11default_time_zone\x18\x06 \x01(\tR\x0fdefaultTimeZone\x121\n" +
	"\x14expiration_timestamp\x18\a \x01(\x03R\x13expirationTimestamp\"\xff\x06\n" +
	"\x0eAsmUserDetails\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1b\n" +
	"\tagent_sid\x18\x03 \x01(\x03R\bagentSid\x12Z\n" +
	"\fagent_status\x18\x04 \x01(\x0e27.services.omnichannel.asm.entities.v1alpha1.StatusStateR\vagentStatus\x127\n" +
	"\x18agent_profile_group_name\x18\x05 \x01(\tR\x15agentProfileGroupName\x12<\n" +
	"\x18current_conversation_sid\x18\x06 \x01(\x03B\x020\x01R\x16currentConversationSid\x12J\n" +
	"\"average_customer_wait_time_seconds\x18\a \x01(\x03R\x1eaverageCustomerWaitTimeSeconds\x12D\n" +
	"\x1faverage_time_to_respond_seconds\x18\b \x01(\x03R\x1baverageTimeToRespondSeconds\x12B\n" +
	"\x0flast_event_time\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\rlastEventTime\x12^\n" +
	"\x06skills\x18\n" +
	" \x03(\v2F.services.omnichannel.asm.entities.v1alpha1.AsmUserDetails.SkillsEntryR\x06skills\x12C\n" +
	"\x0fasm_session_sid\x18\v \x01(\v2\x1b.google.protobuf.Int64ValueR\rasmSessionSid\x12_\n" +
	"\x06events\x18\f \x03(\v2G.services.omnichannel.asm.entities.v1alpha1.DashboardAgentResponseEventR\x06events\x129\n" +
	"\n" +
	"login_time\x18\r \x01(\v2\x1a.google.protobuf.TimestampR\tloginTime\x1a9\n" +
	"\vSkillsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x03R\x05value:\x028\x01\"\xba\x01\n" +
	"\x1bDashboardAgentResponseEvent\x122\n" +
	"\x15response_time_seconds\x18\x01 \x01(\x03R\x13responseTimeSeconds\x12.\n" +
	"\x04time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x04time\x127\n" +
	"\x18is_initial_agent_message\x18\x03 \x01(\bR\x15isInitialAgentMessage*U\n" +
	"\vStatusState\x12\x18\n" +
	"\x14STATUS_STATE_UNKNOWN\x10\x00\x12\v\n" +
	"\aWAITING\x10\x01\x12\b\n" +
	"\x04IDLE\x10\x02\x12\x15\n" +
	"\x11CONVERSATION_OPEN\x10\x03B\xe3\x02\n" +
	".com.services.omnichannel.asm.entities.v1alpha1B\fSessionProtoP\x01ZVgithub.com/tcncloud/api-go/services/omnichannel/asm/entities/v1alpha1;entitiesv1alpha1\xa2\x02\x04SOAE\xaa\x02*Services.Omnichannel.Asm.Entities.V1alpha1\xca\x02*Services\\Omnichannel\\Asm\\Entities\\V1alpha1\xe2\x026Services\\Omnichannel\\Asm\\Entities\\V1alpha1\\GPBMetadata\xea\x02.Services::Omnichannel::Asm::Entities::V1alpha1b\x06proto3"

var (
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescOnce sync.Once
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescData []byte
)

func file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescGZIP() []byte {
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescOnce.Do(func() {
		file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDesc), len(file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDesc)))
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDesc), len(file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDesc)),
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
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_goTypes = nil
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_depIdxs = nil
}
