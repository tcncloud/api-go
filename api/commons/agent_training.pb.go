// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/commons/agent_training.proto

package commons

import (
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

// LearningOpportunityStatus describes the status of the learning opportunity.
type LearningOpportunityStatus int32

const (
	LearningOpportunityStatus_STATUS_OPEN      LearningOpportunityStatus = 0 // Default status - Ready to be completed by the agent.
	LearningOpportunityStatus_STATUS_COMPLETED LearningOpportunityStatus = 1 // Agent has completed the learning opportunity.
)

// Enum value maps for LearningOpportunityStatus.
var (
	LearningOpportunityStatus_name = map[int32]string{
		0: "STATUS_OPEN",
		1: "STATUS_COMPLETED",
	}
	LearningOpportunityStatus_value = map[string]int32{
		"STATUS_OPEN":      0,
		"STATUS_COMPLETED": 1,
	}
)

func (x LearningOpportunityStatus) Enum() *LearningOpportunityStatus {
	p := new(LearningOpportunityStatus)
	*p = x
	return p
}

func (x LearningOpportunityStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LearningOpportunityStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_agent_training_proto_enumTypes[0].Descriptor()
}

func (LearningOpportunityStatus) Type() protoreflect.EnumType {
	return &file_api_commons_agent_training_proto_enumTypes[0]
}

func (x LearningOpportunityStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LearningOpportunityStatus.Descriptor instead.
func (LearningOpportunityStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_agent_training_proto_rawDescGZIP(), []int{0}
}

// LearningOpportunityOrigin describes the possible origins
// that a learning opportunity can be created from.
type LearningOpportunityOrigin int32

const (
	LearningOpportunityOrigin_UNDEFINED            LearningOpportunityOrigin = 0 // Default value - Undefined origin.
	LearningOpportunityOrigin_QUALITY_EVALUATION   LearningOpportunityOrigin = 1 // Created from Quality Evaluation.
	LearningOpportunityOrigin_AUTO_EVALUATION      LearningOpportunityOrigin = 2 // Created from Auto Evaluation.
	LearningOpportunityOrigin_FLAG_EVALUATION      LearningOpportunityOrigin = 3 // Created from Flag Evaluation.
	LearningOpportunityOrigin_CONVERSATION         LearningOpportunityOrigin = 4 // Created from Conversation (voice analytics recordings).
	LearningOpportunityOrigin_FLAGGED_CONVERSATION LearningOpportunityOrigin = 5 // Created from Flagged Conversation (voice analytics flagged recordings).
)

// Enum value maps for LearningOpportunityOrigin.
var (
	LearningOpportunityOrigin_name = map[int32]string{
		0: "UNDEFINED",
		1: "QUALITY_EVALUATION",
		2: "AUTO_EVALUATION",
		3: "FLAG_EVALUATION",
		4: "CONVERSATION",
		5: "FLAGGED_CONVERSATION",
	}
	LearningOpportunityOrigin_value = map[string]int32{
		"UNDEFINED":            0,
		"QUALITY_EVALUATION":   1,
		"AUTO_EVALUATION":      2,
		"FLAG_EVALUATION":      3,
		"CONVERSATION":         4,
		"FLAGGED_CONVERSATION": 5,
	}
)

func (x LearningOpportunityOrigin) Enum() *LearningOpportunityOrigin {
	p := new(LearningOpportunityOrigin)
	*p = x
	return p
}

func (x LearningOpportunityOrigin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LearningOpportunityOrigin) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_agent_training_proto_enumTypes[1].Descriptor()
}

func (LearningOpportunityOrigin) Type() protoreflect.EnumType {
	return &file_api_commons_agent_training_proto_enumTypes[1]
}

func (x LearningOpportunityOrigin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LearningOpportunityOrigin.Descriptor instead.
func (LearningOpportunityOrigin) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_agent_training_proto_rawDescGZIP(), []int{1}
}

// LearningOpportunity represents a single learning opportunity entity.
type LearningOpportunity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LearningOpportunityId int64                     `protobuf:"varint,2,opt,name=learning_opportunity_id,json=learningOpportunityId,proto3" json:"learning_opportunity_id,omitempty"` // Unique id of the learning opportunity.
	CallSid               int64                     `protobuf:"varint,3,opt,name=call_sid,json=callSid,proto3" json:"call_sid,omitempty"`                                             // Call sid of the learning opportunity's related call.
	CallType              CallType_Enum             `protobuf:"varint,4,opt,name=call_type,json=callType,proto3,enum=api.commons.CallType_Enum" json:"call_type,omitempty"`           // Call type of the learning opportunity's related call.
	TranscriptSid         int64                     `protobuf:"varint,5,opt,name=transcript_sid,json=transcriptSid,proto3" json:"transcript_sid,omitempty"`                           // (optional) Transcript sid of the learning opportunity's related transcript.
	AgentUserId           string                    `protobuf:"bytes,6,opt,name=agent_user_id,json=agentUserId,proto3" json:"agent_user_id,omitempty"`                                // Agent user id of the agent receiving the learning opportunity.
	StartOffset           int32                     `protobuf:"varint,7,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`                                 // Start offset (in milliseconds) of the learning opportunity.
	EndOffset             int32                     `protobuf:"varint,8,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`                                       // End offset (in milliseconds) of the learning opportunity.
	Description           string                    `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`                                                     // (optional) Description text.
	CreatedAt             *timestamppb.Timestamp    `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                                       // Time the learning opportunity was created at.
	Title                 string                    `protobuf:"bytes,12,opt,name=title,proto3" json:"title,omitempty"`                                                                // Title of the learning opportunity.
	Status                LearningOpportunityStatus `protobuf:"varint,13,opt,name=status,proto3,enum=api.commons.LearningOpportunityStatus" json:"status,omitempty"`                  // Current status (ex: OPEN, COMPLETED).
	Origin                LearningOpportunityOrigin `protobuf:"varint,14,opt,name=origin,proto3,enum=api.commons.LearningOpportunityOrigin" json:"origin,omitempty"`                  // Origin (ie. opportunity created from).
	CreatorUserId         string                    `protobuf:"bytes,15,opt,name=creator_user_id,json=creatorUserId,proto3" json:"creator_user_id,omitempty"`                         // User id for the creator of the learning opportunity.
}

func (x *LearningOpportunity) Reset() {
	*x = LearningOpportunity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_agent_training_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LearningOpportunity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LearningOpportunity) ProtoMessage() {}

func (x *LearningOpportunity) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_agent_training_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LearningOpportunity.ProtoReflect.Descriptor instead.
func (*LearningOpportunity) Descriptor() ([]byte, []int) {
	return file_api_commons_agent_training_proto_rawDescGZIP(), []int{0}
}

func (x *LearningOpportunity) GetLearningOpportunityId() int64 {
	if x != nil {
		return x.LearningOpportunityId
	}
	return 0
}

func (x *LearningOpportunity) GetCallSid() int64 {
	if x != nil {
		return x.CallSid
	}
	return 0
}

func (x *LearningOpportunity) GetCallType() CallType_Enum {
	if x != nil {
		return x.CallType
	}
	return CallType_INBOUND
}

func (x *LearningOpportunity) GetTranscriptSid() int64 {
	if x != nil {
		return x.TranscriptSid
	}
	return 0
}

func (x *LearningOpportunity) GetAgentUserId() string {
	if x != nil {
		return x.AgentUserId
	}
	return ""
}

func (x *LearningOpportunity) GetStartOffset() int32 {
	if x != nil {
		return x.StartOffset
	}
	return 0
}

func (x *LearningOpportunity) GetEndOffset() int32 {
	if x != nil {
		return x.EndOffset
	}
	return 0
}

func (x *LearningOpportunity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LearningOpportunity) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *LearningOpportunity) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *LearningOpportunity) GetStatus() LearningOpportunityStatus {
	if x != nil {
		return x.Status
	}
	return LearningOpportunityStatus_STATUS_OPEN
}

func (x *LearningOpportunity) GetOrigin() LearningOpportunityOrigin {
	if x != nil {
		return x.Origin
	}
	return LearningOpportunityOrigin_UNDEFINED
}

func (x *LearningOpportunity) GetCreatorUserId() string {
	if x != nil {
		return x.CreatorUserId
	}
	return ""
}

// CallIdentifier is used to completely identify calls.
type CallIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid  int64         `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`                                  // call sid
	Type CallType_Enum `protobuf:"varint,2,opt,name=type,proto3,enum=api.commons.CallType_Enum" json:"type,omitempty"` // call type
}

func (x *CallIdentifier) Reset() {
	*x = CallIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_agent_training_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallIdentifier) ProtoMessage() {}

func (x *CallIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_agent_training_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallIdentifier.ProtoReflect.Descriptor instead.
func (*CallIdentifier) Descriptor() ([]byte, []int) {
	return file_api_commons_agent_training_proto_rawDescGZIP(), []int{1}
}

func (x *CallIdentifier) GetSid() int64 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *CallIdentifier) GetType() CallType_Enum {
	if x != nil {
		return x.Type
	}
	return CallType_INBOUND
}

var File_api_commons_agent_training_proto protoreflect.FileDescriptor

var file_api_commons_agent_training_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a,
	0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x04, 0x0a, 0x13, 0x4c, 0x65, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12,
	0x36, 0x0a, 0x17, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x15, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x53,
	0x69, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53,
	0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65,
	0x6e, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x0e, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x42, 0x0a, 0x19, 0x4c, 0x65, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f,
	0x50, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x01, 0x2a, 0x98, 0x01, 0x0a, 0x19,
	0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44,
	0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x51, 0x55, 0x41, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x45, 0x56, 0x41, 0x4c, 0x55, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x45, 0x56, 0x41, 0x4c, 0x55, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x45, 0x56,
	0x41, 0x4c, 0x55, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14,
	0x46, 0x4c, 0x41, 0x47, 0x47, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x42, 0x9a, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x12, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02,
	0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41,
	0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_agent_training_proto_rawDescOnce sync.Once
	file_api_commons_agent_training_proto_rawDescData = file_api_commons_agent_training_proto_rawDesc
)

func file_api_commons_agent_training_proto_rawDescGZIP() []byte {
	file_api_commons_agent_training_proto_rawDescOnce.Do(func() {
		file_api_commons_agent_training_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_agent_training_proto_rawDescData)
	})
	return file_api_commons_agent_training_proto_rawDescData
}

var file_api_commons_agent_training_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_commons_agent_training_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_agent_training_proto_goTypes = []interface{}{
	(LearningOpportunityStatus)(0), // 0: api.commons.LearningOpportunityStatus
	(LearningOpportunityOrigin)(0), // 1: api.commons.LearningOpportunityOrigin
	(*LearningOpportunity)(nil),    // 2: api.commons.LearningOpportunity
	(*CallIdentifier)(nil),         // 3: api.commons.CallIdentifier
	(CallType_Enum)(0),             // 4: api.commons.CallType.Enum
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
}
var file_api_commons_agent_training_proto_depIdxs = []int32{
	4, // 0: api.commons.LearningOpportunity.call_type:type_name -> api.commons.CallType.Enum
	5, // 1: api.commons.LearningOpportunity.created_at:type_name -> google.protobuf.Timestamp
	0, // 2: api.commons.LearningOpportunity.status:type_name -> api.commons.LearningOpportunityStatus
	1, // 3: api.commons.LearningOpportunity.origin:type_name -> api.commons.LearningOpportunityOrigin
	4, // 4: api.commons.CallIdentifier.type:type_name -> api.commons.CallType.Enum
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_commons_agent_training_proto_init() }
func file_api_commons_agent_training_proto_init() {
	if File_api_commons_agent_training_proto != nil {
		return
	}
	file_api_commons_acd_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_commons_agent_training_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LearningOpportunity); i {
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
		file_api_commons_agent_training_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallIdentifier); i {
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
			RawDescriptor: file_api_commons_agent_training_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_agent_training_proto_goTypes,
		DependencyIndexes: file_api_commons_agent_training_proto_depIdxs,
		EnumInfos:         file_api_commons_agent_training_proto_enumTypes,
		MessageInfos:      file_api_commons_agent_training_proto_msgTypes,
	}.Build()
	File_api_commons_agent_training_proto = out.File
	file_api_commons_agent_training_proto_rawDesc = nil
	file_api_commons_agent_training_proto_goTypes = nil
	file_api_commons_agent_training_proto_depIdxs = nil
}
