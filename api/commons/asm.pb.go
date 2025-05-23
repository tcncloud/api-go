// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/commons/asm.proto

package commons

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

type AsmSubsessionType int32

const (
	AsmSubsessionType_VOICE AsmSubsessionType = 0
	AsmSubsessionType_OMNI  AsmSubsessionType = 1
)

// Enum value maps for AsmSubsessionType.
var (
	AsmSubsessionType_name = map[int32]string{
		0: "VOICE",
		1: "OMNI",
	}
	AsmSubsessionType_value = map[string]int32{
		"VOICE": 0,
		"OMNI":  1,
	}
)

func (x AsmSubsessionType) Enum() *AsmSubsessionType {
	p := new(AsmSubsessionType)
	*p = x
	return p
}

func (x AsmSubsessionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AsmSubsessionType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_asm_proto_enumTypes[0].Descriptor()
}

func (AsmSubsessionType) Type() protoreflect.EnumType {
	return &file_api_commons_asm_proto_enumTypes[0]
}

func (x AsmSubsessionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AsmSubsessionType.Descriptor instead.
func (AsmSubsessionType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_asm_proto_rawDescGZIP(), []int{0}
}

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
	return file_api_commons_asm_proto_enumTypes[1].Descriptor()
}

func (StatusState) Type() protoreflect.EnumType {
	return &file_api_commons_asm_proto_enumTypes[1]
}

func (x StatusState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusState.Descriptor instead.
func (StatusState) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_asm_proto_rawDescGZIP(), []int{1}
}

type DashboardAgentInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// agents user id
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// agents name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// agents user name
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// agents hunt group name
	HuntGroupName string `protobuf:"bytes,4,opt,name=hunt_group_name,json=huntGroupName,proto3" json:"hunt_group_name,omitempty"`
	// partner agent id
	PartnerAgentId string `protobuf:"bytes,6,opt,name=partner_agent_id,json=partnerAgentId,proto3" json:"partner_agent_id,omitempty"`
	// the agents hunt group sid
	HuntGroupSid int64 `protobuf:"varint,9,opt,name=hunt_group_sid,json=huntGroupSid,proto3" json:"hunt_group_sid,omitempty"`
	// the agents sid
	AgentSid int64 `protobuf:"varint,10,opt,name=agent_sid,json=agentSid,proto3" json:"agent_sid,omitempty"`
	// the user caller id
	UserCallerId string `protobuf:"bytes,12,opt,name=user_caller_id,json=userCallerId,proto3" json:"user_caller_id,omitempty"`
	// agents first name
	FirstName string `protobuf:"bytes,13,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// agents last name
	LastName string `protobuf:"bytes,14,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// the date of creation
	Created *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=created,proto3" json:"created,omitempty"`
	// the date of last update
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	// agents profile group id
	AgentProfileGroupId string `protobuf:"bytes,17,opt,name=agent_profile_group_id,json=agentProfileGroupId,proto3" json:"agent_profile_group_id,omitempty"`
	// agents profile group name
	AgentProfileGroupName string `protobuf:"bytes,18,opt,name=agent_profile_group_name,json=agentProfileGroupName,proto3" json:"agent_profile_group_name,omitempty"` // TODO add loginTime
	// enum of agents status
	AgentStatus StatusState `protobuf:"varint,19,opt,name=agent_status,json=agentStatus,proto3,enum=api.commons.StatusState" json:"agent_status,omitempty"`
	// the agents current conversation
	CurrentConversationSid int64 `protobuf:"varint,20,opt,name=current_conversation_sid,json=currentConversationSid,proto3" json:"current_conversation_sid,omitempty"`
	// time from first customer message to agent response. between all conversations.
	AverageCustomerWaitTimeSeconds int64 `protobuf:"varint,22,opt,name=average_customer_wait_time_seconds,json=averageCustomerWaitTimeSeconds,proto3" json:"average_customer_wait_time_seconds,omitempty"`
	// responste time between all conversations.
	AverageTimeToRespondSeconds int64 `protobuf:"varint,23,opt,name=average_time_to_respond_seconds,json=averageTimeToRespondSeconds,proto3" json:"average_time_to_respond_seconds,omitempty"`
	// time from first customer message to conversation close. between all conversations.
	AverageConversationDurationSeconds int64 `protobuf:"varint,24,opt,name=average_conversation_duration_seconds,json=averageConversationDurationSeconds,proto3" json:"average_conversation_duration_seconds,omitempty"`
	// login time
	LoginTime *timestamppb.Timestamp `protobuf:"bytes,25,opt,name=login_time,json=loginTime,proto3" json:"login_time,omitempty"`
	// last event time
	LastEventTime *timestamppb.Timestamp `protobuf:"bytes,26,opt,name=last_event_time,json=lastEventTime,proto3" json:"last_event_time,omitempty"`
	// the list of response events
	Events []*DashboardAgentInfo_DashboardAgentResponseEvent `protobuf:"bytes,27,rep,name=events,proto3" json:"events,omitempty"`
	// Agents Skills
	Skills map[string]int64 `protobuf:"bytes,28,rep,name=skills,proto3" json:"skills,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Asm Session Sid
	AsmSessionSid *wrapperspb.Int64Value `protobuf:"bytes,29,opt,name=asm_session_sid,json=asmSessionSid,proto3" json:"asm_session_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DashboardAgentInfo) Reset() {
	*x = DashboardAgentInfo{}
	mi := &file_api_commons_asm_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DashboardAgentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardAgentInfo) ProtoMessage() {}

func (x *DashboardAgentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_asm_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardAgentInfo.ProtoReflect.Descriptor instead.
func (*DashboardAgentInfo) Descriptor() ([]byte, []int) {
	return file_api_commons_asm_proto_rawDescGZIP(), []int{0}
}

func (x *DashboardAgentInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DashboardAgentInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DashboardAgentInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *DashboardAgentInfo) GetHuntGroupName() string {
	if x != nil {
		return x.HuntGroupName
	}
	return ""
}

func (x *DashboardAgentInfo) GetPartnerAgentId() string {
	if x != nil {
		return x.PartnerAgentId
	}
	return ""
}

func (x *DashboardAgentInfo) GetHuntGroupSid() int64 {
	if x != nil {
		return x.HuntGroupSid
	}
	return 0
}

func (x *DashboardAgentInfo) GetAgentSid() int64 {
	if x != nil {
		return x.AgentSid
	}
	return 0
}

func (x *DashboardAgentInfo) GetUserCallerId() string {
	if x != nil {
		return x.UserCallerId
	}
	return ""
}

func (x *DashboardAgentInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *DashboardAgentInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *DashboardAgentInfo) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *DashboardAgentInfo) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *DashboardAgentInfo) GetAgentProfileGroupId() string {
	if x != nil {
		return x.AgentProfileGroupId
	}
	return ""
}

func (x *DashboardAgentInfo) GetAgentProfileGroupName() string {
	if x != nil {
		return x.AgentProfileGroupName
	}
	return ""
}

func (x *DashboardAgentInfo) GetAgentStatus() StatusState {
	if x != nil {
		return x.AgentStatus
	}
	return StatusState_STATUS_STATE_UNKNOWN
}

func (x *DashboardAgentInfo) GetCurrentConversationSid() int64 {
	if x != nil {
		return x.CurrentConversationSid
	}
	return 0
}

func (x *DashboardAgentInfo) GetAverageCustomerWaitTimeSeconds() int64 {
	if x != nil {
		return x.AverageCustomerWaitTimeSeconds
	}
	return 0
}

func (x *DashboardAgentInfo) GetAverageTimeToRespondSeconds() int64 {
	if x != nil {
		return x.AverageTimeToRespondSeconds
	}
	return 0
}

func (x *DashboardAgentInfo) GetAverageConversationDurationSeconds() int64 {
	if x != nil {
		return x.AverageConversationDurationSeconds
	}
	return 0
}

func (x *DashboardAgentInfo) GetLoginTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LoginTime
	}
	return nil
}

func (x *DashboardAgentInfo) GetLastEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastEventTime
	}
	return nil
}

func (x *DashboardAgentInfo) GetEvents() []*DashboardAgentInfo_DashboardAgentResponseEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *DashboardAgentInfo) GetSkills() map[string]int64 {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *DashboardAgentInfo) GetAsmSessionSid() *wrapperspb.Int64Value {
	if x != nil {
		return x.AsmSessionSid
	}
	return nil
}

type StreamAgentStateRes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to AgentState:
	//
	//	*StreamAgentStateRes_State
	//	*StreamAgentStateRes_HeartBeat
	//	*StreamAgentStateRes_CallQueueAdd
	//	*StreamAgentStateRes_CallQueueRemove
	AgentState    isStreamAgentStateRes_AgentState `protobuf_oneof:"agent_state"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamAgentStateRes) Reset() {
	*x = StreamAgentStateRes{}
	mi := &file_api_commons_asm_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamAgentStateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAgentStateRes) ProtoMessage() {}

func (x *StreamAgentStateRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_asm_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAgentStateRes.ProtoReflect.Descriptor instead.
func (*StreamAgentStateRes) Descriptor() ([]byte, []int) {
	return file_api_commons_asm_proto_rawDescGZIP(), []int{1}
}

func (x *StreamAgentStateRes) GetAgentState() isStreamAgentStateRes_AgentState {
	if x != nil {
		return x.AgentState
	}
	return nil
}

func (x *StreamAgentStateRes) GetState() *AgentState {
	if x != nil {
		if x, ok := x.AgentState.(*StreamAgentStateRes_State); ok {
			return x.State
		}
	}
	return nil
}

func (x *StreamAgentStateRes) GetHeartBeat() *KeepAlive {
	if x != nil {
		if x, ok := x.AgentState.(*StreamAgentStateRes_HeartBeat); ok {
			return x.HeartBeat
		}
	}
	return nil
}

func (x *StreamAgentStateRes) GetCallQueueAdd() *QueueCallAdd {
	if x != nil {
		if x, ok := x.AgentState.(*StreamAgentStateRes_CallQueueAdd); ok {
			return x.CallQueueAdd
		}
	}
	return nil
}

func (x *StreamAgentStateRes) GetCallQueueRemove() *QueueCallRemove {
	if x != nil {
		if x, ok := x.AgentState.(*StreamAgentStateRes_CallQueueRemove); ok {
			return x.CallQueueRemove
		}
	}
	return nil
}

type isStreamAgentStateRes_AgentState interface {
	isStreamAgentStateRes_AgentState()
}

type StreamAgentStateRes_State struct {
	State *AgentState `protobuf:"bytes,1,opt,name=state,proto3,oneof"`
}

type StreamAgentStateRes_HeartBeat struct {
	HeartBeat *KeepAlive `protobuf:"bytes,2,opt,name=heart_beat,json=heartBeat,proto3,oneof"`
}

type StreamAgentStateRes_CallQueueAdd struct {
	CallQueueAdd *QueueCallAdd `protobuf:"bytes,3,opt,name=call_queue_add,json=callQueueAdd,proto3,oneof"`
}

type StreamAgentStateRes_CallQueueRemove struct {
	CallQueueRemove *QueueCallRemove `protobuf:"bytes,4,opt,name=call_queue_remove,json=callQueueRemove,proto3,oneof"`
}

func (*StreamAgentStateRes_State) isStreamAgentStateRes_AgentState() {}

func (*StreamAgentStateRes_HeartBeat) isStreamAgentStateRes_AgentState() {}

func (*StreamAgentStateRes_CallQueueAdd) isStreamAgentStateRes_AgentState() {}

func (*StreamAgentStateRes_CallQueueRemove) isStreamAgentStateRes_AgentState() {}

type ManagerStreamAgentStateRes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to ManagerAgentState:
	//
	//	*ManagerStreamAgentStateRes_State
	//	*ManagerStreamAgentStateRes_HeartBeat
	ManagerAgentState isManagerStreamAgentStateRes_ManagerAgentState `protobuf_oneof:"manager_agent_state"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ManagerStreamAgentStateRes) Reset() {
	*x = ManagerStreamAgentStateRes{}
	mi := &file_api_commons_asm_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManagerStreamAgentStateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerStreamAgentStateRes) ProtoMessage() {}

func (x *ManagerStreamAgentStateRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_asm_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerStreamAgentStateRes.ProtoReflect.Descriptor instead.
func (*ManagerStreamAgentStateRes) Descriptor() ([]byte, []int) {
	return file_api_commons_asm_proto_rawDescGZIP(), []int{2}
}

func (x *ManagerStreamAgentStateRes) GetManagerAgentState() isManagerStreamAgentStateRes_ManagerAgentState {
	if x != nil {
		return x.ManagerAgentState
	}
	return nil
}

func (x *ManagerStreamAgentStateRes) GetState() *AgentState {
	if x != nil {
		if x, ok := x.ManagerAgentState.(*ManagerStreamAgentStateRes_State); ok {
			return x.State
		}
	}
	return nil
}

func (x *ManagerStreamAgentStateRes) GetHeartBeat() *KeepAlive {
	if x != nil {
		if x, ok := x.ManagerAgentState.(*ManagerStreamAgentStateRes_HeartBeat); ok {
			return x.HeartBeat
		}
	}
	return nil
}

type isManagerStreamAgentStateRes_ManagerAgentState interface {
	isManagerStreamAgentStateRes_ManagerAgentState()
}

type ManagerStreamAgentStateRes_State struct {
	State *AgentState `protobuf:"bytes,1,opt,name=state,proto3,oneof"`
}

type ManagerStreamAgentStateRes_HeartBeat struct {
	HeartBeat *KeepAlive `protobuf:"bytes,2,opt,name=heart_beat,json=heartBeat,proto3,oneof"`
}

func (*ManagerStreamAgentStateRes_State) isManagerStreamAgentStateRes_ManagerAgentState() {}

func (*ManagerStreamAgentStateRes_HeartBeat) isManagerStreamAgentStateRes_ManagerAgentState() {}

type KeepAlive struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KeepAlive) Reset() {
	*x = KeepAlive{}
	mi := &file_api_commons_asm_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeepAlive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepAlive) ProtoMessage() {}

func (x *KeepAlive) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_asm_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepAlive.ProtoReflect.Descriptor instead.
func (*KeepAlive) Descriptor() ([]byte, []int) {
	return file_api_commons_asm_proto_rawDescGZIP(), []int{3}
}

type QueueCallAdd struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// number dialed when the call was placed.
	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// caller_id used when the call was placed.
	CallerId string `protobuf:"bytes,2,opt,name=caller_id,json=callerId,proto3" json:"caller_id,omitempty"`
	// timestamp indicating when the call started.
	StartDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// timestamp indicating when the call was put on hold (only set when call is a hold call).
	HoldDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=hold_date,json=holdDate,proto3" json:"hold_date,omitempty"`
	// formatted skills that the call requires.
	FormattedSkills map[string]string `protobuf:"bytes,5,rep,name=formatted_skills,json=formattedSkills,proto3" json:"formatted_skills,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// indicates if the call is specific to the agent (multi-hold) or it can be picked up by multiple agents (queued and HQM calls).
	AgentSpecific bool `protobuf:"varint,6,opt,name=agent_specific,json=agentSpecific,proto3" json:"agent_specific,omitempty"`
	// queued notification type of the call.
	QueuedNotificationType QueuedNotificationType `protobuf:"varint,7,opt,name=queued_notification_type,json=queuedNotificationType,proto3,enum=api.commons.QueuedNotificationType" json:"queued_notification_type,omitempty"`
	// callersid
	CallerSid *CallerSid `protobuf:"bytes,8,opt,name=caller_sid,json=callerSid,proto3" json:"caller_sid,omitempty"`
	// skills
	Skills        map[string]bool `protobuf:"bytes,9,rep,name=skills,proto3" json:"skills,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueueCallAdd) Reset() {
	*x = QueueCallAdd{}
	mi := &file_api_commons_asm_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueueCallAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueCallAdd) ProtoMessage() {}

func (x *QueueCallAdd) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_asm_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueCallAdd.ProtoReflect.Descriptor instead.
func (*QueueCallAdd) Descriptor() ([]byte, []int) {
	return file_api_commons_asm_proto_rawDescGZIP(), []int{4}
}

func (x *QueueCallAdd) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *QueueCallAdd) GetCallerId() string {
	if x != nil {
		return x.CallerId
	}
	return ""
}

func (x *QueueCallAdd) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *QueueCallAdd) GetHoldDate() *timestamppb.Timestamp {
	if x != nil {
		return x.HoldDate
	}
	return nil
}

func (x *QueueCallAdd) GetFormattedSkills() map[string]string {
	if x != nil {
		return x.FormattedSkills
	}
	return nil
}

func (x *QueueCallAdd) GetAgentSpecific() bool {
	if x != nil {
		return x.AgentSpecific
	}
	return false
}

func (x *QueueCallAdd) GetQueuedNotificationType() QueuedNotificationType {
	if x != nil {
		return x.QueuedNotificationType
	}
	return QueuedNotificationType_QueuedNotificationType_GENERAL_INITIAL
}

func (x *QueueCallAdd) GetCallerSid() *CallerSid {
	if x != nil {
		return x.CallerSid
	}
	return nil
}

func (x *QueueCallAdd) GetSkills() map[string]bool {
	if x != nil {
		return x.Skills
	}
	return nil
}

type QueueCallRemove struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CallerSid     *CallerSid             `protobuf:"bytes,1,opt,name=caller_sid,json=callerSid,proto3" json:"caller_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueueCallRemove) Reset() {
	*x = QueueCallRemove{}
	mi := &file_api_commons_asm_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueueCallRemove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueCallRemove) ProtoMessage() {}

func (x *QueueCallRemove) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_asm_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueCallRemove.ProtoReflect.Descriptor instead.
func (*QueueCallRemove) Descriptor() ([]byte, []int) {
	return file_api_commons_asm_proto_rawDescGZIP(), []int{5}
}

func (x *QueueCallRemove) GetCallerSid() *CallerSid {
	if x != nil {
		return x.CallerSid
	}
	return nil
}

// The response event information for an agent
type DashboardAgentInfo_DashboardAgentResponseEvent struct {
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

func (x *DashboardAgentInfo_DashboardAgentResponseEvent) Reset() {
	*x = DashboardAgentInfo_DashboardAgentResponseEvent{}
	mi := &file_api_commons_asm_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DashboardAgentInfo_DashboardAgentResponseEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardAgentInfo_DashboardAgentResponseEvent) ProtoMessage() {}

func (x *DashboardAgentInfo_DashboardAgentResponseEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_asm_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardAgentInfo_DashboardAgentResponseEvent.ProtoReflect.Descriptor instead.
func (*DashboardAgentInfo_DashboardAgentResponseEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_asm_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DashboardAgentInfo_DashboardAgentResponseEvent) GetResponseTimeSeconds() int64 {
	if x != nil {
		return x.ResponseTimeSeconds
	}
	return 0
}

func (x *DashboardAgentInfo_DashboardAgentResponseEvent) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *DashboardAgentInfo_DashboardAgentResponseEvent) GetIsInitialAgentMessage() bool {
	if x != nil {
		return x.IsInitialAgentMessage
	}
	return false
}

var File_api_commons_asm_proto protoreflect.FileDescriptor

const file_api_commons_asm_proto_rawDesc = "" +
	"\n" +
	"\x15api/commons/asm.proto\x12\vapi.commons\x1a\x15api/commons/acd.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\xee\v\n" +
	"\x12DashboardAgentInfo\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1b\n" +
	"\tuser_name\x18\x03 \x01(\tR\buserName\x12&\n" +
	"\x0fhunt_group_name\x18\x04 \x01(\tR\rhuntGroupName\x12(\n" +
	"\x10partner_agent_id\x18\x06 \x01(\tR\x0epartnerAgentId\x12$\n" +
	"\x0ehunt_group_sid\x18\t \x01(\x03R\fhuntGroupSid\x12\x1b\n" +
	"\tagent_sid\x18\n" +
	" \x01(\x03R\bagentSid\x12$\n" +
	"\x0euser_caller_id\x18\f \x01(\tR\fuserCallerId\x12\x1d\n" +
	"\n" +
	"first_name\x18\r \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x0e \x01(\tR\blastName\x124\n" +
	"\acreated\x18\x0f \x01(\v2\x1a.google.protobuf.TimestampR\acreated\x12=\n" +
	"\flast_updated\x18\x10 \x01(\v2\x1a.google.protobuf.TimestampR\vlastUpdated\x123\n" +
	"\x16agent_profile_group_id\x18\x11 \x01(\tR\x13agentProfileGroupId\x127\n" +
	"\x18agent_profile_group_name\x18\x12 \x01(\tR\x15agentProfileGroupName\x12;\n" +
	"\fagent_status\x18\x13 \x01(\x0e2\x18.api.commons.StatusStateR\vagentStatus\x12<\n" +
	"\x18current_conversation_sid\x18\x14 \x01(\x03B\x020\x01R\x16currentConversationSid\x12J\n" +
	"\"average_customer_wait_time_seconds\x18\x16 \x01(\x03R\x1eaverageCustomerWaitTimeSeconds\x12D\n" +
	"\x1faverage_time_to_respond_seconds\x18\x17 \x01(\x03R\x1baverageTimeToRespondSeconds\x12Q\n" +
	"%average_conversation_duration_seconds\x18\x18 \x01(\x03R\"averageConversationDurationSeconds\x129\n" +
	"\n" +
	"login_time\x18\x19 \x01(\v2\x1a.google.protobuf.TimestampR\tloginTime\x12B\n" +
	"\x0flast_event_time\x18\x1a \x01(\v2\x1a.google.protobuf.TimestampR\rlastEventTime\x12S\n" +
	"\x06events\x18\x1b \x03(\v2;.api.commons.DashboardAgentInfo.DashboardAgentResponseEventR\x06events\x12C\n" +
	"\x06skills\x18\x1c \x03(\v2+.api.commons.DashboardAgentInfo.SkillsEntryR\x06skills\x12C\n" +
	"\x0fasm_session_sid\x18\x1d \x01(\v2\x1b.google.protobuf.Int64ValueR\rasmSessionSid\x1a\xba\x01\n" +
	"\x1bDashboardAgentResponseEvent\x122\n" +
	"\x15response_time_seconds\x18\x01 \x01(\x03R\x13responseTimeSeconds\x12.\n" +
	"\x04time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x04time\x127\n" +
	"\x18is_initial_agent_message\x18\x03 \x01(\bR\x15isInitialAgentMessage\x1a9\n" +
	"\vSkillsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x03R\x05value:\x028\x01\"\x9d\x02\n" +
	"\x13StreamAgentStateRes\x12/\n" +
	"\x05state\x18\x01 \x01(\v2\x17.api.commons.AgentStateH\x00R\x05state\x127\n" +
	"\n" +
	"heart_beat\x18\x02 \x01(\v2\x16.api.commons.KeepAliveH\x00R\theartBeat\x12A\n" +
	"\x0ecall_queue_add\x18\x03 \x01(\v2\x19.api.commons.QueueCallAddH\x00R\fcallQueueAdd\x12J\n" +
	"\x11call_queue_remove\x18\x04 \x01(\v2\x1c.api.commons.QueueCallRemoveH\x00R\x0fcallQueueRemoveB\r\n" +
	"\vagent_state\"\x9d\x01\n" +
	"\x1aManagerStreamAgentStateRes\x12/\n" +
	"\x05state\x18\x01 \x01(\v2\x17.api.commons.AgentStateH\x00R\x05state\x127\n" +
	"\n" +
	"heart_beat\x18\x02 \x01(\v2\x16.api.commons.KeepAliveH\x00R\theartBeatB\x15\n" +
	"\x13manager_agent_state\"\v\n" +
	"\tKeepAlive\"\x98\x05\n" +
	"\fQueueCallAdd\x12!\n" +
	"\fphone_number\x18\x01 \x01(\tR\vphoneNumber\x12\x1b\n" +
	"\tcaller_id\x18\x02 \x01(\tR\bcallerId\x129\n" +
	"\n" +
	"start_date\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tstartDate\x127\n" +
	"\thold_date\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\bholdDate\x12Y\n" +
	"\x10formatted_skills\x18\x05 \x03(\v2..api.commons.QueueCallAdd.FormattedSkillsEntryR\x0fformattedSkills\x12%\n" +
	"\x0eagent_specific\x18\x06 \x01(\bR\ragentSpecific\x12]\n" +
	"\x18queued_notification_type\x18\a \x01(\x0e2#.api.commons.QueuedNotificationTypeR\x16queuedNotificationType\x125\n" +
	"\n" +
	"caller_sid\x18\b \x01(\v2\x16.api.commons.CallerSidR\tcallerSid\x12=\n" +
	"\x06skills\x18\t \x03(\v2%.api.commons.QueueCallAdd.SkillsEntryR\x06skills\x1aB\n" +
	"\x14FormattedSkillsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1a9\n" +
	"\vSkillsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\bR\x05value:\x028\x01\"H\n" +
	"\x0fQueueCallRemove\x125\n" +
	"\n" +
	"caller_sid\x18\x01 \x01(\v2\x16.api.commons.CallerSidR\tcallerSid*(\n" +
	"\x11AsmSubsessionType\x12\t\n" +
	"\x05VOICE\x10\x00\x12\b\n" +
	"\x04OMNI\x10\x01*U\n" +
	"\vStatusState\x12\x18\n" +
	"\x14STATUS_STATE_UNKNOWN\x10\x00\x12\v\n" +
	"\aWAITING\x10\x01\x12\b\n" +
	"\x04IDLE\x10\x02\x12\x15\n" +
	"\x11CONVERSATION_OPEN\x10\x03B\x90\x01\n" +
	"\x0fcom.api.commonsB\bAsmProtoP\x01Z&github.com/tcncloud/api-go/api/commons\xa2\x02\x03ACX\xaa\x02\vApi.Commons\xca\x02\vApi\\Commons\xe2\x02\x17Api\\Commons\\GPBMetadata\xea\x02\fApi::Commonsb\x06proto3"

var (
	file_api_commons_asm_proto_rawDescOnce sync.Once
	file_api_commons_asm_proto_rawDescData []byte
)

func file_api_commons_asm_proto_rawDescGZIP() []byte {
	file_api_commons_asm_proto_rawDescOnce.Do(func() {
		file_api_commons_asm_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_asm_proto_rawDesc), len(file_api_commons_asm_proto_rawDesc)))
	})
	return file_api_commons_asm_proto_rawDescData
}

var file_api_commons_asm_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_commons_asm_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_commons_asm_proto_goTypes = []any{
	(AsmSubsessionType)(0),                                 // 0: api.commons.AsmSubsessionType
	(StatusState)(0),                                       // 1: api.commons.StatusState
	(*DashboardAgentInfo)(nil),                             // 2: api.commons.DashboardAgentInfo
	(*StreamAgentStateRes)(nil),                            // 3: api.commons.StreamAgentStateRes
	(*ManagerStreamAgentStateRes)(nil),                     // 4: api.commons.ManagerStreamAgentStateRes
	(*KeepAlive)(nil),                                      // 5: api.commons.KeepAlive
	(*QueueCallAdd)(nil),                                   // 6: api.commons.QueueCallAdd
	(*QueueCallRemove)(nil),                                // 7: api.commons.QueueCallRemove
	(*DashboardAgentInfo_DashboardAgentResponseEvent)(nil), // 8: api.commons.DashboardAgentInfo.DashboardAgentResponseEvent
	nil,                           // 9: api.commons.DashboardAgentInfo.SkillsEntry
	nil,                           // 10: api.commons.QueueCallAdd.FormattedSkillsEntry
	nil,                           // 11: api.commons.QueueCallAdd.SkillsEntry
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
	(*wrapperspb.Int64Value)(nil), // 13: google.protobuf.Int64Value
	(*AgentState)(nil),            // 14: api.commons.AgentState
	(QueuedNotificationType)(0),   // 15: api.commons.QueuedNotificationType
	(*CallerSid)(nil),             // 16: api.commons.CallerSid
}
var file_api_commons_asm_proto_depIdxs = []int32{
	12, // 0: api.commons.DashboardAgentInfo.created:type_name -> google.protobuf.Timestamp
	12, // 1: api.commons.DashboardAgentInfo.last_updated:type_name -> google.protobuf.Timestamp
	1,  // 2: api.commons.DashboardAgentInfo.agent_status:type_name -> api.commons.StatusState
	12, // 3: api.commons.DashboardAgentInfo.login_time:type_name -> google.protobuf.Timestamp
	12, // 4: api.commons.DashboardAgentInfo.last_event_time:type_name -> google.protobuf.Timestamp
	8,  // 5: api.commons.DashboardAgentInfo.events:type_name -> api.commons.DashboardAgentInfo.DashboardAgentResponseEvent
	9,  // 6: api.commons.DashboardAgentInfo.skills:type_name -> api.commons.DashboardAgentInfo.SkillsEntry
	13, // 7: api.commons.DashboardAgentInfo.asm_session_sid:type_name -> google.protobuf.Int64Value
	14, // 8: api.commons.StreamAgentStateRes.state:type_name -> api.commons.AgentState
	5,  // 9: api.commons.StreamAgentStateRes.heart_beat:type_name -> api.commons.KeepAlive
	6,  // 10: api.commons.StreamAgentStateRes.call_queue_add:type_name -> api.commons.QueueCallAdd
	7,  // 11: api.commons.StreamAgentStateRes.call_queue_remove:type_name -> api.commons.QueueCallRemove
	14, // 12: api.commons.ManagerStreamAgentStateRes.state:type_name -> api.commons.AgentState
	5,  // 13: api.commons.ManagerStreamAgentStateRes.heart_beat:type_name -> api.commons.KeepAlive
	12, // 14: api.commons.QueueCallAdd.start_date:type_name -> google.protobuf.Timestamp
	12, // 15: api.commons.QueueCallAdd.hold_date:type_name -> google.protobuf.Timestamp
	10, // 16: api.commons.QueueCallAdd.formatted_skills:type_name -> api.commons.QueueCallAdd.FormattedSkillsEntry
	15, // 17: api.commons.QueueCallAdd.queued_notification_type:type_name -> api.commons.QueuedNotificationType
	16, // 18: api.commons.QueueCallAdd.caller_sid:type_name -> api.commons.CallerSid
	11, // 19: api.commons.QueueCallAdd.skills:type_name -> api.commons.QueueCallAdd.SkillsEntry
	16, // 20: api.commons.QueueCallRemove.caller_sid:type_name -> api.commons.CallerSid
	12, // 21: api.commons.DashboardAgentInfo.DashboardAgentResponseEvent.time:type_name -> google.protobuf.Timestamp
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_api_commons_asm_proto_init() }
func file_api_commons_asm_proto_init() {
	if File_api_commons_asm_proto != nil {
		return
	}
	file_api_commons_acd_proto_init()
	file_api_commons_asm_proto_msgTypes[1].OneofWrappers = []any{
		(*StreamAgentStateRes_State)(nil),
		(*StreamAgentStateRes_HeartBeat)(nil),
		(*StreamAgentStateRes_CallQueueAdd)(nil),
		(*StreamAgentStateRes_CallQueueRemove)(nil),
	}
	file_api_commons_asm_proto_msgTypes[2].OneofWrappers = []any{
		(*ManagerStreamAgentStateRes_State)(nil),
		(*ManagerStreamAgentStateRes_HeartBeat)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_asm_proto_rawDesc), len(file_api_commons_asm_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_asm_proto_goTypes,
		DependencyIndexes: file_api_commons_asm_proto_depIdxs,
		EnumInfos:         file_api_commons_asm_proto_enumTypes,
		MessageInfos:      file_api_commons_asm_proto_msgTypes,
	}.Build()
	File_api_commons_asm_proto = out.File
	file_api_commons_asm_proto_goTypes = nil
	file_api_commons_asm_proto_depIdxs = nil
}
