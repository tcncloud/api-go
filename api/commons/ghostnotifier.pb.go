// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/commons/ghostnotifier.proto

package commons

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// a notificaiton used to return a response for an asyncronous request or
// data for or about a backend initiated process that needs to be consumed by the user
type GhostNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a structured id used to identify the given notification
	ReferenceId string `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*GhostNotification_Any
	//	*GhostNotification_Status
	//	*GhostNotification_OmniConversation
	//	*GhostNotification_BackofficeMessage
	//	*GhostNotification_DirectedCallRinging
	//	*GhostNotification_DirectedCallHangup
	//	*GhostNotification_AgentQueuedCallsNotification
	//	*GhostNotification_AuthTokenExpirationNotification
	Payload isGhostNotification_Payload `protobuf_oneof:"payload"`
}

func (x *GhostNotification) Reset() {
	*x = GhostNotification{}
	mi := &file_api_commons_ghostnotifier_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GhostNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GhostNotification) ProtoMessage() {}

func (x *GhostNotification) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_ghostnotifier_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GhostNotification.ProtoReflect.Descriptor instead.
func (*GhostNotification) Descriptor() ([]byte, []int) {
	return file_api_commons_ghostnotifier_proto_rawDescGZIP(), []int{0}
}

func (x *GhostNotification) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (m *GhostNotification) GetPayload() isGhostNotification_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *GhostNotification) GetAny() *anypb.Any {
	if x, ok := x.GetPayload().(*GhostNotification_Any); ok {
		return x.Any
	}
	return nil
}

func (x *GhostNotification) GetStatus() *Status {
	if x, ok := x.GetPayload().(*GhostNotification_Status); ok {
		return x.Status
	}
	return nil
}

func (x *GhostNotification) GetOmniConversation() *OmniConversation {
	if x, ok := x.GetPayload().(*GhostNotification_OmniConversation); ok {
		return x.OmniConversation
	}
	return nil
}

func (x *GhostNotification) GetBackofficeMessage() *AgentBackofficeMessageAlert {
	if x, ok := x.GetPayload().(*GhostNotification_BackofficeMessage); ok {
		return x.BackofficeMessage
	}
	return nil
}

func (x *GhostNotification) GetDirectedCallRinging() *AgentDirectedCallRingingAlert {
	if x, ok := x.GetPayload().(*GhostNotification_DirectedCallRinging); ok {
		return x.DirectedCallRinging
	}
	return nil
}

func (x *GhostNotification) GetDirectedCallHangup() *AgentDirectedCallHangupAlert {
	if x, ok := x.GetPayload().(*GhostNotification_DirectedCallHangup); ok {
		return x.DirectedCallHangup
	}
	return nil
}

func (x *GhostNotification) GetAgentQueuedCallsNotification() *AgentQueuedCallsNotification {
	if x, ok := x.GetPayload().(*GhostNotification_AgentQueuedCallsNotification); ok {
		return x.AgentQueuedCallsNotification
	}
	return nil
}

func (x *GhostNotification) GetAuthTokenExpirationNotification() *AuthTokenExpiration {
	if x, ok := x.GetPayload().(*GhostNotification_AuthTokenExpirationNotification); ok {
		return x.AuthTokenExpirationNotification
	}
	return nil
}

type isGhostNotification_Payload interface {
	isGhostNotification_Payload()
}

type GhostNotification_Any struct {
	// payload used when no other oneof type has been defined
	Any *anypb.Any `protobuf:"bytes,2,opt,name=any,proto3,oneof"`
}

type GhostNotification_Status struct {
	// A way to return a message with a status code and message.
	Status *Status `protobuf:"bytes,3,opt,name=status,proto3,oneof"`
}

type GhostNotification_OmniConversation struct {
	// An omni conversation
	OmniConversation *OmniConversation `protobuf:"bytes,4,opt,name=omni_conversation,json=omniConversation,proto3,oneof"`
}

type GhostNotification_BackofficeMessage struct {
	// an alert to an agent that a message from backoffice came in
	BackofficeMessage *AgentBackofficeMessageAlert `protobuf:"bytes,6,opt,name=backoffice_message,json=backofficeMessage,proto3,oneof"`
}

type GhostNotification_DirectedCallRinging struct {
	// alert that a direct agent call is ringing
	DirectedCallRinging *AgentDirectedCallRingingAlert `protobuf:"bytes,7,opt,name=directed_call_ringing,json=directedCallRinging,proto3,oneof"`
}

type GhostNotification_DirectedCallHangup struct {
	// alert that a direct agent call hungup
	DirectedCallHangup *AgentDirectedCallHangupAlert `protobuf:"bytes,8,opt,name=directed_call_hangup,json=directedCallHangup,proto3,oneof"`
}

type GhostNotification_AgentQueuedCallsNotification struct {
	// notification about the current state of queued call for an agent
	AgentQueuedCallsNotification *AgentQueuedCallsNotification `protobuf:"bytes,9,opt,name=agent_queued_calls_notification,json=agentQueuedCallsNotification,proto3,oneof"`
}

type GhostNotification_AuthTokenExpirationNotification struct {
	// notification that an auth token will soon expire
	AuthTokenExpirationNotification *AuthTokenExpiration `protobuf:"bytes,11,opt,name=auth_token_expiration_notification,json=authTokenExpirationNotification,proto3,oneof"` // ... undetermined other possible payloads
}

func (*GhostNotification_Any) isGhostNotification_Payload() {}

func (*GhostNotification_Status) isGhostNotification_Payload() {}

func (*GhostNotification_OmniConversation) isGhostNotification_Payload() {}

func (*GhostNotification_BackofficeMessage) isGhostNotification_Payload() {}

func (*GhostNotification_DirectedCallRinging) isGhostNotification_Payload() {}

func (*GhostNotification_DirectedCallHangup) isGhostNotification_Payload() {}

func (*GhostNotification_AgentQueuedCallsNotification) isGhostNotification_Payload() {}

func (*GhostNotification_AuthTokenExpirationNotification) isGhostNotification_Payload() {}

// This is based on googles status proto message
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status code
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A developer-facing error message, which should be in English
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_api_commons_ghostnotifier_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_ghostnotifier_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_api_commons_ghostnotifier_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AgentQueuedCallsNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// calls that were queued and have a matching subset of the agent_skills from the request.
	AgentQueueCalls []*AgentQueuedCallsNotification_QueuedCallData `protobuf:"bytes,16,rep,name=agent_queue_calls,json=agentQueueCalls,proto3" json:"agent_queue_calls,omitempty"`
	// calls that were placed on multi-hold by the agent_sid from the request.
	OnHoldCalls []*AgentQueuedCallsNotification_QueuedCallData `protobuf:"bytes,17,rep,name=on_hold_calls,json=onHoldCalls,proto3" json:"on_hold_calls,omitempty"`
	// calls that were placed in the Hold Queue Monitor, are still on hold by the destination (hold music is being played by the other party) and have
	// a matching subset of the agent_skills from the request.
	HqmCalls []*AgentQueuedCallsNotification_QueuedCallData `protobuf:"bytes,18,rep,name=hqm_calls,json=hqmCalls,proto3" json:"hqm_calls,omitempty"`
}

func (x *AgentQueuedCallsNotification) Reset() {
	*x = AgentQueuedCallsNotification{}
	mi := &file_api_commons_ghostnotifier_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentQueuedCallsNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentQueuedCallsNotification) ProtoMessage() {}

func (x *AgentQueuedCallsNotification) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_ghostnotifier_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentQueuedCallsNotification.ProtoReflect.Descriptor instead.
func (*AgentQueuedCallsNotification) Descriptor() ([]byte, []int) {
	return file_api_commons_ghostnotifier_proto_rawDescGZIP(), []int{2}
}

func (x *AgentQueuedCallsNotification) GetAgentQueueCalls() []*AgentQueuedCallsNotification_QueuedCallData {
	if x != nil {
		return x.AgentQueueCalls
	}
	return nil
}

func (x *AgentQueuedCallsNotification) GetOnHoldCalls() []*AgentQueuedCallsNotification_QueuedCallData {
	if x != nil {
		return x.OnHoldCalls
	}
	return nil
}

func (x *AgentQueuedCallsNotification) GetHqmCalls() []*AgentQueuedCallsNotification_QueuedCallData {
	if x != nil {
		return x.HqmCalls
	}
	return nil
}

// AuthTokenExpiration is used to notify the user about an auth token expiration
type AuthTokenExpiration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The token that is expiring
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// The token expiration
	Expiration *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *AuthTokenExpiration) Reset() {
	*x = AuthTokenExpiration{}
	mi := &file_api_commons_ghostnotifier_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthTokenExpiration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthTokenExpiration) ProtoMessage() {}

func (x *AuthTokenExpiration) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_ghostnotifier_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthTokenExpiration.ProtoReflect.Descriptor instead.
func (*AuthTokenExpiration) Descriptor() ([]byte, []int) {
	return file_api_commons_ghostnotifier_proto_rawDescGZIP(), []int{3}
}

func (x *AuthTokenExpiration) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthTokenExpiration) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

type AgentQueuedCallsNotification_QueuedCallData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the found call.
	CallSid int64 `protobuf:"varint,1,opt,name=call_sid,json=callSid,proto3" json:"call_sid,omitempty"`
	// number dialed when the call was placed.
	PhoneNumber string `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// caller_id used when the call was placed.
	CallerId string `protobuf:"bytes,3,opt,name=caller_id,json=callerId,proto3" json:"caller_id,omitempty"`
	// call type of the found call.
	CallType CallType_Enum `protobuf:"varint,4,opt,name=call_type,json=callType,proto3,enum=api.commons.CallType_Enum" json:"call_type,omitempty"`
	// timestamp indicating when the call started.
	StartDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// timestamp indicating when the call was put on hold (only set when call is a hold call).
	HoldDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=hold_date,json=holdDate,proto3" json:"hold_date,omitempty"`
	// formatted skills that the call requires.
	Skills []string `protobuf:"bytes,7,rep,name=skills,proto3" json:"skills,omitempty"`
	// indicates if the call is specific to the agent (multi-hold) or it can be picked up by multiple agents (queued and HQM calls).
	AgentSpecific bool `protobuf:"varint,8,opt,name=agent_specific,json=agentSpecific,proto3" json:"agent_specific,omitempty"`
	// queued notification type of the call.
	QueuedNotificationType QueuedNotificationType `protobuf:"varint,9,opt,name=queued_notification_type,json=queuedNotificationType,proto3,enum=api.commons.QueuedNotificationType" json:"queued_notification_type,omitempty"`
}

func (x *AgentQueuedCallsNotification_QueuedCallData) Reset() {
	*x = AgentQueuedCallsNotification_QueuedCallData{}
	mi := &file_api_commons_ghostnotifier_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentQueuedCallsNotification_QueuedCallData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentQueuedCallsNotification_QueuedCallData) ProtoMessage() {}

func (x *AgentQueuedCallsNotification_QueuedCallData) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_ghostnotifier_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentQueuedCallsNotification_QueuedCallData.ProtoReflect.Descriptor instead.
func (*AgentQueuedCallsNotification_QueuedCallData) Descriptor() ([]byte, []int) {
	return file_api_commons_ghostnotifier_proto_rawDescGZIP(), []int{2, 0}
}

func (x *AgentQueuedCallsNotification_QueuedCallData) GetCallSid() int64 {
	if x != nil {
		return x.CallSid
	}
	return 0
}

func (x *AgentQueuedCallsNotification_QueuedCallData) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *AgentQueuedCallsNotification_QueuedCallData) GetCallerId() string {
	if x != nil {
		return x.CallerId
	}
	return ""
}

func (x *AgentQueuedCallsNotification_QueuedCallData) GetCallType() CallType_Enum {
	if x != nil {
		return x.CallType
	}
	return CallType_INBOUND
}

func (x *AgentQueuedCallsNotification_QueuedCallData) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *AgentQueuedCallsNotification_QueuedCallData) GetHoldDate() *timestamppb.Timestamp {
	if x != nil {
		return x.HoldDate
	}
	return nil
}

func (x *AgentQueuedCallsNotification_QueuedCallData) GetSkills() []string {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *AgentQueuedCallsNotification_QueuedCallData) GetAgentSpecific() bool {
	if x != nil {
		return x.AgentSpecific
	}
	return false
}

func (x *AgentQueuedCallsNotification_QueuedCallData) GetQueuedNotificationType() QueuedNotificationType {
	if x != nil {
		return x.QueuedNotificationType
	}
	return QueuedNotificationType_QueuedNotificationType_GENERAL_INITIAL
}

var File_api_commons_ghostnotifier_proto protoreflect.FileDescriptor

var file_api_commons_ghostnotifier_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x15,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x86, 0x06, 0x0a, 0x11, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x03, 0x61, 0x6e, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x03,
	0x61, 0x6e, 0x79, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x4c, 0x0a, 0x11, 0x6f, 0x6d, 0x6e, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x6d, 0x6e, 0x69,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10,
	0x6f, 0x6d, 0x6e, 0x69, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x59, 0x0a, 0x12, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x48, 0x00, 0x52, 0x11, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x60, 0x0a, 0x15, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x69, 0x6e,
	0x67, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x69, 0x6e, 0x67, 0x69, 0x6e,
	0x67, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x48, 0x00, 0x52, 0x13, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x5d, 0x0a,
	0x14, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x68,
	0x61, 0x6e, 0x67, 0x75, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x61, 0x6e, 0x67, 0x75,
	0x70, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x48, 0x00, 0x52, 0x12, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x61, 0x6e, 0x67, 0x75, 0x70, 0x12, 0x72, 0x0a, 0x1f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6c,
	0x6c, 0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x43,
	0x61, 0x6c, 0x6c, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x1c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x43,
	0x61, 0x6c, 0x6c, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x6f, 0x0a, 0x22, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x1f, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4a, 0x04, 0x08, 0x0a,
	0x10, 0x0b, 0x52, 0x15, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xf2, 0x05, 0x0a, 0x1c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x64, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x64, 0x0a, 0x11, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x43,
	0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x5c, 0x0a, 0x0d, 0x6f, 0x6e, 0x5f, 0x68,
	0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x64, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x6f, 0x6e, 0x48, 0x6f, 0x6c,
	0x64, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x55, 0x0a, 0x09, 0x68, 0x71, 0x6d, 0x5f, 0x63, 0x61,
	0x6c, 0x6c, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x68, 0x71, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x1a, 0xb6, 0x03,
	0x0a, 0x0e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x63,
	0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x37, 0x0a, 0x09, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x68, 0x6f, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x5d, 0x0a, 0x18, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x16,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x67, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x9a, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x42, 0x12, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_ghostnotifier_proto_rawDescOnce sync.Once
	file_api_commons_ghostnotifier_proto_rawDescData = file_api_commons_ghostnotifier_proto_rawDesc
)

func file_api_commons_ghostnotifier_proto_rawDescGZIP() []byte {
	file_api_commons_ghostnotifier_proto_rawDescOnce.Do(func() {
		file_api_commons_ghostnotifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_ghostnotifier_proto_rawDescData)
	})
	return file_api_commons_ghostnotifier_proto_rawDescData
}

var file_api_commons_ghostnotifier_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_commons_ghostnotifier_proto_goTypes = []any{
	(*GhostNotification)(nil),                           // 0: api.commons.GhostNotification
	(*Status)(nil),                                      // 1: api.commons.Status
	(*AgentQueuedCallsNotification)(nil),                // 2: api.commons.AgentQueuedCallsNotification
	(*AuthTokenExpiration)(nil),                         // 3: api.commons.AuthTokenExpiration
	(*AgentQueuedCallsNotification_QueuedCallData)(nil), // 4: api.commons.AgentQueuedCallsNotification.QueuedCallData
	(*anypb.Any)(nil),                                   // 5: google.protobuf.Any
	(*OmniConversation)(nil),                            // 6: api.commons.OmniConversation
	(*AgentBackofficeMessageAlert)(nil),                 // 7: api.commons.AgentBackofficeMessageAlert
	(*AgentDirectedCallRingingAlert)(nil),               // 8: api.commons.AgentDirectedCallRingingAlert
	(*AgentDirectedCallHangupAlert)(nil),                // 9: api.commons.AgentDirectedCallHangupAlert
	(*timestamppb.Timestamp)(nil),                       // 10: google.protobuf.Timestamp
	(CallType_Enum)(0),                                  // 11: api.commons.CallType.Enum
	(QueuedNotificationType)(0),                         // 12: api.commons.QueuedNotificationType
}
var file_api_commons_ghostnotifier_proto_depIdxs = []int32{
	5,  // 0: api.commons.GhostNotification.any:type_name -> google.protobuf.Any
	1,  // 1: api.commons.GhostNotification.status:type_name -> api.commons.Status
	6,  // 2: api.commons.GhostNotification.omni_conversation:type_name -> api.commons.OmniConversation
	7,  // 3: api.commons.GhostNotification.backoffice_message:type_name -> api.commons.AgentBackofficeMessageAlert
	8,  // 4: api.commons.GhostNotification.directed_call_ringing:type_name -> api.commons.AgentDirectedCallRingingAlert
	9,  // 5: api.commons.GhostNotification.directed_call_hangup:type_name -> api.commons.AgentDirectedCallHangupAlert
	2,  // 6: api.commons.GhostNotification.agent_queued_calls_notification:type_name -> api.commons.AgentQueuedCallsNotification
	3,  // 7: api.commons.GhostNotification.auth_token_expiration_notification:type_name -> api.commons.AuthTokenExpiration
	4,  // 8: api.commons.AgentQueuedCallsNotification.agent_queue_calls:type_name -> api.commons.AgentQueuedCallsNotification.QueuedCallData
	4,  // 9: api.commons.AgentQueuedCallsNotification.on_hold_calls:type_name -> api.commons.AgentQueuedCallsNotification.QueuedCallData
	4,  // 10: api.commons.AgentQueuedCallsNotification.hqm_calls:type_name -> api.commons.AgentQueuedCallsNotification.QueuedCallData
	10, // 11: api.commons.AuthTokenExpiration.expiration:type_name -> google.protobuf.Timestamp
	11, // 12: api.commons.AgentQueuedCallsNotification.QueuedCallData.call_type:type_name -> api.commons.CallType.Enum
	10, // 13: api.commons.AgentQueuedCallsNotification.QueuedCallData.start_date:type_name -> google.protobuf.Timestamp
	10, // 14: api.commons.AgentQueuedCallsNotification.QueuedCallData.hold_date:type_name -> google.protobuf.Timestamp
	12, // 15: api.commons.AgentQueuedCallsNotification.QueuedCallData.queued_notification_type:type_name -> api.commons.QueuedNotificationType
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_api_commons_ghostnotifier_proto_init() }
func file_api_commons_ghostnotifier_proto_init() {
	if File_api_commons_ghostnotifier_proto != nil {
		return
	}
	file_api_commons_acd_proto_init()
	file_api_commons_omnichannel_proto_init()
	file_api_commons_ghostnotifier_proto_msgTypes[0].OneofWrappers = []any{
		(*GhostNotification_Any)(nil),
		(*GhostNotification_Status)(nil),
		(*GhostNotification_OmniConversation)(nil),
		(*GhostNotification_BackofficeMessage)(nil),
		(*GhostNotification_DirectedCallRinging)(nil),
		(*GhostNotification_DirectedCallHangup)(nil),
		(*GhostNotification_AgentQueuedCallsNotification)(nil),
		(*GhostNotification_AuthTokenExpirationNotification)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_ghostnotifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_ghostnotifier_proto_goTypes,
		DependencyIndexes: file_api_commons_ghostnotifier_proto_depIdxs,
		MessageInfos:      file_api_commons_ghostnotifier_proto_msgTypes,
	}.Build()
	File_api_commons_ghostnotifier_proto = out.File
	file_api_commons_ghostnotifier_proto_rawDesc = nil
	file_api_commons_ghostnotifier_proto_goTypes = nil
	file_api_commons_ghostnotifier_proto_depIdxs = nil
}
