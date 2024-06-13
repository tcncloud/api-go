// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: wfo/vanalytics/v2/agent_call_log.proto

package vanalyticsv2

import (
	commons "github.com/tcncloud/api-go/api/commons"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AgentCallLog resource.
type AgentCallLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// actions are the agent call log actions.
	Actions []*AgentCallLog_Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *AgentCallLog) Reset() {
	*x = AgentCallLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCallLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCallLog) ProtoMessage() {}

func (x *AgentCallLog) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCallLog.ProtoReflect.Descriptor instead.
func (*AgentCallLog) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_agent_call_log_proto_rawDescGZIP(), []int{0}
}

func (x *AgentCallLog) GetActions() []*AgentCallLog_Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

// Query constraints on agent call log.
type AgentCallLogQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Query constraints on call skills initial.
	CallSkillsInitial *AgentCallLogQuery_CallSkillsInitial `protobuf:"bytes,1,opt,name=call_skills_initial,json=callSkillsInitial,proto3" json:"call_skills_initial,omitempty"`
	// Query constraints on the reason the call ended.
	CallEnded *AgentCallLogQuery_CallEnded `protobuf:"bytes,2,opt,name=call_ended,json=callEnded,proto3" json:"call_ended,omitempty"`
}

func (x *AgentCallLogQuery) Reset() {
	*x = AgentCallLogQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCallLogQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCallLogQuery) ProtoMessage() {}

func (x *AgentCallLogQuery) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCallLogQuery.ProtoReflect.Descriptor instead.
func (*AgentCallLogQuery) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_agent_call_log_proto_rawDescGZIP(), []int{1}
}

func (x *AgentCallLogQuery) GetCallSkillsInitial() *AgentCallLogQuery_CallSkillsInitial {
	if x != nil {
		return x.CallSkillsInitial
	}
	return nil
}

func (x *AgentCallLogQuery) GetCallEnded() *AgentCallLogQuery_CallEnded {
	if x != nil {
		return x.CallEnded
	}
	return nil
}

// Action is an agent call log action.
type AgentCallLog_Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*AgentCallLog_Action_CallSkillsInitial
	//	*AgentCallLog_Action_CallEnded
	Kind isAgentCallLog_Action_Kind `protobuf_oneof:"kind"`
}

func (x *AgentCallLog_Action) Reset() {
	*x = AgentCallLog_Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCallLog_Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCallLog_Action) ProtoMessage() {}

func (x *AgentCallLog_Action) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCallLog_Action.ProtoReflect.Descriptor instead.
func (*AgentCallLog_Action) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_agent_call_log_proto_rawDescGZIP(), []int{0, 0}
}

func (m *AgentCallLog_Action) GetKind() isAgentCallLog_Action_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *AgentCallLog_Action) GetCallSkillsInitial() *AgentCallLog_CallSkillsInitial {
	if x, ok := x.GetKind().(*AgentCallLog_Action_CallSkillsInitial); ok {
		return x.CallSkillsInitial
	}
	return nil
}

func (x *AgentCallLog_Action) GetCallEnded() string {
	if x, ok := x.GetKind().(*AgentCallLog_Action_CallEnded); ok {
		return x.CallEnded
	}
	return ""
}

type isAgentCallLog_Action_Kind interface {
	isAgentCallLog_Action_Kind()
}

type AgentCallLog_Action_CallSkillsInitial struct {
	// CallSkillsInitial are the initial call skills on a call.
	CallSkillsInitial *AgentCallLog_CallSkillsInitial `protobuf:"bytes,1,opt,name=call_skills_initial,json=callSkillsInitial,proto3,oneof"`
}

type AgentCallLog_Action_CallEnded struct {
	// CallEnded is the reason the call ended.
	CallEnded string `protobuf:"bytes,2,opt,name=call_ended,json=callEnded,proto3,oneof"`
}

func (*AgentCallLog_Action_CallSkillsInitial) isAgentCallLog_Action_Kind() {}

func (*AgentCallLog_Action_CallEnded) isAgentCallLog_Action_Kind() {}

// CallSkillsInitial are the initial call skills on a call.
type AgentCallLog_CallSkillsInitial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// need is all the call skills that are needed.
	Need []string `protobuf:"bytes,1,rep,name=need,proto3" json:"need,omitempty"`
	// want is all the call skills that are wanted.
	Want []string `protobuf:"bytes,2,rep,name=want,proto3" json:"want,omitempty"`
}

func (x *AgentCallLog_CallSkillsInitial) Reset() {
	*x = AgentCallLog_CallSkillsInitial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCallLog_CallSkillsInitial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCallLog_CallSkillsInitial) ProtoMessage() {}

func (x *AgentCallLog_CallSkillsInitial) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCallLog_CallSkillsInitial.ProtoReflect.Descriptor instead.
func (*AgentCallLog_CallSkillsInitial) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_agent_call_log_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AgentCallLog_CallSkillsInitial) GetNeed() []string {
	if x != nil {
		return x.Need
	}
	return nil
}

func (x *AgentCallLog_CallSkillsInitial) GetWant() []string {
	if x != nil {
		return x.Want
	}
	return nil
}

// Query constraints on call skills initial.
type AgentCallLogQuery_CallSkillsInitial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Query constraints on call skills initial need.
	Need *AgentCallLogQuery_CallSkillsInitial_Need `protobuf:"bytes,1,opt,name=need,proto3" json:"need,omitempty"`
	// Query constraints on call skills initial want.
	Want *AgentCallLogQuery_CallSkillsInitial_Want `protobuf:"bytes,2,opt,name=want,proto3" json:"want,omitempty"`
}

func (x *AgentCallLogQuery_CallSkillsInitial) Reset() {
	*x = AgentCallLogQuery_CallSkillsInitial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCallLogQuery_CallSkillsInitial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCallLogQuery_CallSkillsInitial) ProtoMessage() {}

func (x *AgentCallLogQuery_CallSkillsInitial) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCallLogQuery_CallSkillsInitial.ProtoReflect.Descriptor instead.
func (*AgentCallLogQuery_CallSkillsInitial) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_agent_call_log_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AgentCallLogQuery_CallSkillsInitial) GetNeed() *AgentCallLogQuery_CallSkillsInitial_Need {
	if x != nil {
		return x.Need
	}
	return nil
}

func (x *AgentCallLogQuery_CallSkillsInitial) GetWant() *AgentCallLogQuery_CallSkillsInitial_Want {
	if x != nil {
		return x.Want
	}
	return nil
}

// Query constraints on the reason the call ended.
type AgentCallLogQuery_CallEnded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Query constraints on the reason the call ended.
	Reasons []commons.AgentCallLogCallEnded `protobuf:"varint,1,rep,packed,name=reasons,proto3,enum=api.commons.AgentCallLogCallEnded" json:"reasons,omitempty"`
}

func (x *AgentCallLogQuery_CallEnded) Reset() {
	*x = AgentCallLogQuery_CallEnded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCallLogQuery_CallEnded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCallLogQuery_CallEnded) ProtoMessage() {}

func (x *AgentCallLogQuery_CallEnded) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCallLogQuery_CallEnded.ProtoReflect.Descriptor instead.
func (*AgentCallLogQuery_CallEnded) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_agent_call_log_proto_rawDescGZIP(), []int{1, 1}
}

func (x *AgentCallLogQuery_CallEnded) GetReasons() []commons.AgentCallLogCallEnded {
	if x != nil {
		return x.Reasons
	}
	return nil
}

// Query constraints on call skills initial need.
type AgentCallLogQuery_CallSkillsInitial_Need struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requires all call specific transcript hits to have a needed agent
	// call log in the list.
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// Requires all values to match.
	All bool `protobuf:"varint,2,opt,name=all,proto3" json:"all,omitempty"`
}

func (x *AgentCallLogQuery_CallSkillsInitial_Need) Reset() {
	*x = AgentCallLogQuery_CallSkillsInitial_Need{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCallLogQuery_CallSkillsInitial_Need) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCallLogQuery_CallSkillsInitial_Need) ProtoMessage() {}

func (x *AgentCallLogQuery_CallSkillsInitial_Need) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCallLogQuery_CallSkillsInitial_Need.ProtoReflect.Descriptor instead.
func (*AgentCallLogQuery_CallSkillsInitial_Need) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_agent_call_log_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *AgentCallLogQuery_CallSkillsInitial_Need) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *AgentCallLogQuery_CallSkillsInitial_Need) GetAll() bool {
	if x != nil {
		return x.All
	}
	return false
}

// Query constraints on call skills initial want.
type AgentCallLogQuery_CallSkillsInitial_Want struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requires all call specific transcript hits to have a wanted agent
	// call log in the list.
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// Requires all values to match.
	All bool `protobuf:"varint,2,opt,name=all,proto3" json:"all,omitempty"`
}

func (x *AgentCallLogQuery_CallSkillsInitial_Want) Reset() {
	*x = AgentCallLogQuery_CallSkillsInitial_Want{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCallLogQuery_CallSkillsInitial_Want) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCallLogQuery_CallSkillsInitial_Want) ProtoMessage() {}

func (x *AgentCallLogQuery_CallSkillsInitial_Want) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCallLogQuery_CallSkillsInitial_Want.ProtoReflect.Descriptor instead.
func (*AgentCallLogQuery_CallSkillsInitial_Want) Descriptor() ([]byte, []int) {
	return file_wfo_vanalytics_v2_agent_call_log_proto_rawDescGZIP(), []int{1, 0, 1}
}

func (x *AgentCallLogQuery_CallSkillsInitial_Want) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *AgentCallLogQuery_CallSkillsInitial_Want) GetAll() bool {
	if x != nil {
		return x.All
	}
	return false
}

var File_wfo_vanalytics_v2_agent_call_log_proto protoreflect.FileDescriptor

var file_wfo_vanalytics_v2_agent_call_log_proto_rawDesc = []byte{
	0x0a, 0x26, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x15, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c,
	0x4c, 0x6f, 0x67, 0x12, 0x40, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x61,
	0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x96, 0x01, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x63, 0x0a, 0x13, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x5f,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x48, 0x00, 0x52, 0x11, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x63, 0x61, 0x6c,
	0x6c, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x1a, 0x3b,
	0x0a, 0x11, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x65, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x61, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x77, 0x61, 0x6e, 0x74, 0x22, 0xb1, 0x04, 0x0a, 0x11,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x66, 0x0a, 0x13, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x11, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x4d, 0x0a, 0x0a, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x52, 0x09, 0x63,
	0x61, 0x6c, 0x6c, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x1a, 0x99, 0x02, 0x0a, 0x11, 0x43, 0x61, 0x6c,
	0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x4f,
	0x0a, 0x04, 0x6e, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x77,
	0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x2e, 0x4e, 0x65, 0x65, 0x64, 0x52, 0x04, 0x6e, 0x65, 0x65, 0x64, 0x12,
	0x4f, 0x0a, 0x04, 0x77, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x57, 0x61, 0x6e, 0x74, 0x52, 0x04, 0x77, 0x61, 0x6e, 0x74,
	0x1a, 0x30, 0x0a, 0x04, 0x4e, 0x65, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61,
	0x6c, 0x6c, 0x1a, 0x30, 0x0a, 0x04, 0x57, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x61, 0x6c, 0x6c, 0x1a, 0x49, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x6e, 0x64, 0x65,
	0x64, 0x12, 0x3c, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x43, 0x61, 0x6c,
	0x6c, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x42,
	0xcb, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x11, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x77, 0x66, 0x6f, 0x2f, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x57, 0x56, 0x58, 0xaa,
	0x02, 0x11, 0x57, 0x66, 0x6f, 0x2e, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x56, 0x32, 0xca, 0x02, 0x11, 0x57, 0x66, 0x6f, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x1d, 0x57, 0x66, 0x6f, 0x5c, 0x56, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x57, 0x66, 0x6f, 0x3a, 0x3a, 0x56,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wfo_vanalytics_v2_agent_call_log_proto_rawDescOnce sync.Once
	file_wfo_vanalytics_v2_agent_call_log_proto_rawDescData = file_wfo_vanalytics_v2_agent_call_log_proto_rawDesc
)

func file_wfo_vanalytics_v2_agent_call_log_proto_rawDescGZIP() []byte {
	file_wfo_vanalytics_v2_agent_call_log_proto_rawDescOnce.Do(func() {
		file_wfo_vanalytics_v2_agent_call_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_wfo_vanalytics_v2_agent_call_log_proto_rawDescData)
	})
	return file_wfo_vanalytics_v2_agent_call_log_proto_rawDescData
}

var file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_wfo_vanalytics_v2_agent_call_log_proto_goTypes = []any{
	(*AgentCallLog)(nil),                             // 0: wfo.vanalytics.v2.AgentCallLog
	(*AgentCallLogQuery)(nil),                        // 1: wfo.vanalytics.v2.AgentCallLogQuery
	(*AgentCallLog_Action)(nil),                      // 2: wfo.vanalytics.v2.AgentCallLog.Action
	(*AgentCallLog_CallSkillsInitial)(nil),           // 3: wfo.vanalytics.v2.AgentCallLog.CallSkillsInitial
	(*AgentCallLogQuery_CallSkillsInitial)(nil),      // 4: wfo.vanalytics.v2.AgentCallLogQuery.CallSkillsInitial
	(*AgentCallLogQuery_CallEnded)(nil),              // 5: wfo.vanalytics.v2.AgentCallLogQuery.CallEnded
	(*AgentCallLogQuery_CallSkillsInitial_Need)(nil), // 6: wfo.vanalytics.v2.AgentCallLogQuery.CallSkillsInitial.Need
	(*AgentCallLogQuery_CallSkillsInitial_Want)(nil), // 7: wfo.vanalytics.v2.AgentCallLogQuery.CallSkillsInitial.Want
	(commons.AgentCallLogCallEnded)(0),               // 8: api.commons.AgentCallLogCallEnded
}
var file_wfo_vanalytics_v2_agent_call_log_proto_depIdxs = []int32{
	2, // 0: wfo.vanalytics.v2.AgentCallLog.actions:type_name -> wfo.vanalytics.v2.AgentCallLog.Action
	4, // 1: wfo.vanalytics.v2.AgentCallLogQuery.call_skills_initial:type_name -> wfo.vanalytics.v2.AgentCallLogQuery.CallSkillsInitial
	5, // 2: wfo.vanalytics.v2.AgentCallLogQuery.call_ended:type_name -> wfo.vanalytics.v2.AgentCallLogQuery.CallEnded
	3, // 3: wfo.vanalytics.v2.AgentCallLog.Action.call_skills_initial:type_name -> wfo.vanalytics.v2.AgentCallLog.CallSkillsInitial
	6, // 4: wfo.vanalytics.v2.AgentCallLogQuery.CallSkillsInitial.need:type_name -> wfo.vanalytics.v2.AgentCallLogQuery.CallSkillsInitial.Need
	7, // 5: wfo.vanalytics.v2.AgentCallLogQuery.CallSkillsInitial.want:type_name -> wfo.vanalytics.v2.AgentCallLogQuery.CallSkillsInitial.Want
	8, // 6: wfo.vanalytics.v2.AgentCallLogQuery.CallEnded.reasons:type_name -> api.commons.AgentCallLogCallEnded
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_wfo_vanalytics_v2_agent_call_log_proto_init() }
func file_wfo_vanalytics_v2_agent_call_log_proto_init() {
	if File_wfo_vanalytics_v2_agent_call_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AgentCallLog); i {
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
		file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AgentCallLogQuery); i {
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
		file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AgentCallLog_Action); i {
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
		file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AgentCallLog_CallSkillsInitial); i {
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
		file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AgentCallLogQuery_CallSkillsInitial); i {
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
		file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AgentCallLogQuery_CallEnded); i {
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
		file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*AgentCallLogQuery_CallSkillsInitial_Need); i {
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
		file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*AgentCallLogQuery_CallSkillsInitial_Want); i {
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
	file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes[2].OneofWrappers = []any{
		(*AgentCallLog_Action_CallSkillsInitial)(nil),
		(*AgentCallLog_Action_CallEnded)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wfo_vanalytics_v2_agent_call_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wfo_vanalytics_v2_agent_call_log_proto_goTypes,
		DependencyIndexes: file_wfo_vanalytics_v2_agent_call_log_proto_depIdxs,
		MessageInfos:      file_wfo_vanalytics_v2_agent_call_log_proto_msgTypes,
	}.Build()
	File_wfo_vanalytics_v2_agent_call_log_proto = out.File
	file_wfo_vanalytics_v2_agent_call_log_proto_rawDesc = nil
	file_wfo_vanalytics_v2_agent_call_log_proto_goTypes = nil
	file_wfo_vanalytics_v2_agent_call_log_proto_depIdxs = nil
}
