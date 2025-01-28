// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: api/v0alpha/sds.proto

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

type GetAgentResponseDataReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CallSid       int64                  `protobuf:"varint,1,opt,name=call_sid,json=callSid,proto3" json:"call_sid,omitempty"`
	CallType      commons.CallType_Enum  `protobuf:"varint,2,opt,name=call_type,json=callType,proto3,enum=api.commons.CallType_Enum" json:"call_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAgentResponseDataReq) Reset() {
	*x = GetAgentResponseDataReq{}
	mi := &file_api_v0alpha_sds_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAgentResponseDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentResponseDataReq) ProtoMessage() {}

func (x *GetAgentResponseDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sds_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentResponseDataReq.ProtoReflect.Descriptor instead.
func (*GetAgentResponseDataReq) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_sds_proto_rawDescGZIP(), []int{0}
}

func (x *GetAgentResponseDataReq) GetCallSid() int64 {
	if x != nil {
		return x.CallSid
	}
	return 0
}

func (x *GetAgentResponseDataReq) GetCallType() commons.CallType_Enum {
	if x != nil {
		return x.CallType
	}
	return commons.CallType_Enum(0)
}

type GetAgentResponseDataRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CallSid       int64                  `protobuf:"varint,1,opt,name=call_sid,json=callSid,proto3" json:"call_sid,omitempty"`
	CallType      commons.CallType_Enum  `protobuf:"varint,2,opt,name=call_type,json=callType,proto3,enum=api.commons.CallType_Enum" json:"call_type,omitempty"`
	Responses     map[string]string      `protobuf:"bytes,3,rep,name=responses,proto3" json:"responses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAgentResponseDataRes) Reset() {
	*x = GetAgentResponseDataRes{}
	mi := &file_api_v0alpha_sds_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAgentResponseDataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentResponseDataRes) ProtoMessage() {}

func (x *GetAgentResponseDataRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sds_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentResponseDataRes.ProtoReflect.Descriptor instead.
func (*GetAgentResponseDataRes) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_sds_proto_rawDescGZIP(), []int{1}
}

func (x *GetAgentResponseDataRes) GetCallSid() int64 {
	if x != nil {
		return x.CallSid
	}
	return 0
}

func (x *GetAgentResponseDataRes) GetCallType() commons.CallType_Enum {
	if x != nil {
		return x.CallType
	}
	return commons.CallType_Enum(0)
}

func (x *GetAgentResponseDataRes) GetResponses() map[string]string {
	if x != nil {
		return x.Responses
	}
	return nil
}

type GetCallReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CallSid       int64                  `protobuf:"varint,2,opt,name=call_sid,json=callSid,proto3" json:"call_sid,omitempty"`
	CallType      commons.CallType_Enum  `protobuf:"varint,3,opt,name=call_type,json=callType,proto3,enum=api.commons.CallType_Enum" json:"call_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCallReq) Reset() {
	*x = GetCallReq{}
	mi := &file_api_v0alpha_sds_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCallReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCallReq) ProtoMessage() {}

func (x *GetCallReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sds_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCallReq.ProtoReflect.Descriptor instead.
func (*GetCallReq) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_sds_proto_rawDescGZIP(), []int{2}
}

func (x *GetCallReq) GetCallSid() int64 {
	if x != nil {
		return x.CallSid
	}
	return 0
}

func (x *GetCallReq) GetCallType() commons.CallType_Enum {
	if x != nil {
		return x.CallType
	}
	return commons.CallType_Enum(0)
}

type UpdateVoicemailBoxReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CallSid       int64                  `protobuf:"varint,2,opt,name=call_sid,json=callSid,proto3" json:"call_sid,omitempty"`
	CallType      commons.CallType_Enum  `protobuf:"varint,3,opt,name=call_type,json=callType,proto3,enum=api.commons.CallType_Enum" json:"call_type,omitempty"`
	PbxExtension  string                 `protobuf:"bytes,4,opt,name=pbx_extension,json=pbxExtension,proto3" json:"pbx_extension,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateVoicemailBoxReq) Reset() {
	*x = UpdateVoicemailBoxReq{}
	mi := &file_api_v0alpha_sds_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateVoicemailBoxReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVoicemailBoxReq) ProtoMessage() {}

func (x *UpdateVoicemailBoxReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sds_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVoicemailBoxReq.ProtoReflect.Descriptor instead.
func (*UpdateVoicemailBoxReq) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_sds_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateVoicemailBoxReq) GetCallSid() int64 {
	if x != nil {
		return x.CallSid
	}
	return 0
}

func (x *UpdateVoicemailBoxReq) GetCallType() commons.CallType_Enum {
	if x != nil {
		return x.CallType
	}
	return commons.CallType_Enum(0)
}

func (x *UpdateVoicemailBoxReq) GetPbxExtension() string {
	if x != nil {
		return x.PbxExtension
	}
	return ""
}

type UpdateVoicemailBoxRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateVoicemailBoxRes) Reset() {
	*x = UpdateVoicemailBoxRes{}
	mi := &file_api_v0alpha_sds_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateVoicemailBoxRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVoicemailBoxRes) ProtoMessage() {}

func (x *UpdateVoicemailBoxRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sds_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVoicemailBoxRes.ProtoReflect.Descriptor instead.
func (*UpdateVoicemailBoxRes) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_sds_proto_rawDescGZIP(), []int{4}
}

type CallObject struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Oid               string                 `protobuf:"bytes,1,opt,name=oid,proto3" json:"oid,omitempty"`
	CallSid           int64                  `protobuf:"varint,2,opt,name=call_sid,json=callSid,proto3" json:"call_sid,omitempty"`
	CallType          commons.CallType_Enum  `protobuf:"varint,3,opt,name=call_type,json=callType,proto3,enum=api.commons.CallType_Enum" json:"call_type,omitempty"`
	Updated           int64                  `protobuf:"varint,10,opt,name=updated,proto3" json:"updated,omitempty"`
	Skills            map[string]bool        `protobuf:"bytes,11,rep,name=skills,proto3" json:"skills,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	RecordingFile     string                 `protobuf:"bytes,12,opt,name=recording_file,json=recordingFile,proto3" json:"recording_file,omitempty"`
	UpdatedDate       *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=updated_date,json=updatedDate,proto3" json:"updated_date,omitempty"`
	SrcNumber         string                 `protobuf:"bytes,14,opt,name=src_number,json=srcNumber,proto3" json:"src_number,omitempty"`
	DstNumber         string                 `protobuf:"bytes,15,opt,name=dst_number,json=dstNumber,proto3" json:"dst_number,omitempty"`
	CallerIdName      string                 `protobuf:"bytes,16,opt,name=caller_id_name,json=callerIdName,proto3" json:"caller_id_name,omitempty"`
	AgentWorker       string                 `protobuf:"bytes,17,opt,name=agent_worker,json=agentWorker,proto3" json:"agent_worker,omitempty"`
	Events            []string               `protobuf:"bytes,18,rep,name=events,proto3" json:"events,omitempty"`
	CallData          string                 `protobuf:"bytes,19,opt,name=call_data,json=callData,proto3" json:"call_data,omitempty"`
	AgentResponseData map[string]string      `protobuf:"bytes,20,rep,name=agent_response_data,json=agentResponseData,proto3" json:"agent_response_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Recorded          bool                   `protobuf:"varint,21,opt,name=recorded,proto3" json:"recorded,omitempty"`
	Connected         bool                   `protobuf:"varint,22,opt,name=connected,proto3" json:"connected,omitempty"`
	Suspended         bool                   `protobuf:"varint,23,opt,name=suspended,proto3" json:"suspended,omitempty"`
	DisconnectReason  string                 `protobuf:"bytes,24,opt,name=disconnect_reason,json=disconnectReason,proto3" json:"disconnect_reason,omitempty"`
	Voicemailed       bool                   `protobuf:"varint,25,opt,name=voicemailed,proto3" json:"voicemailed,omitempty"`
	VoicemailBox      string                 `protobuf:"bytes,26,opt,name=voicemail_box,json=voicemailBox,proto3" json:"voicemail_box,omitempty"`
	Originated        string                 `protobuf:"bytes,27,opt,name=originated,proto3" json:"originated,omitempty"`
	Folder            string                 `protobuf:"bytes,28,opt,name=folder,proto3" json:"folder,omitempty"`
	RtpInfo           string                 `protobuf:"bytes,29,opt,name=rtp_info,json=rtpInfo,proto3" json:"rtp_info,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CallObject) Reset() {
	*x = CallObject{}
	mi := &file_api_v0alpha_sds_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallObject) ProtoMessage() {}

func (x *CallObject) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sds_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallObject.ProtoReflect.Descriptor instead.
func (*CallObject) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_sds_proto_rawDescGZIP(), []int{5}
}

func (x *CallObject) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

func (x *CallObject) GetCallSid() int64 {
	if x != nil {
		return x.CallSid
	}
	return 0
}

func (x *CallObject) GetCallType() commons.CallType_Enum {
	if x != nil {
		return x.CallType
	}
	return commons.CallType_Enum(0)
}

func (x *CallObject) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *CallObject) GetSkills() map[string]bool {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *CallObject) GetRecordingFile() string {
	if x != nil {
		return x.RecordingFile
	}
	return ""
}

func (x *CallObject) GetUpdatedDate() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedDate
	}
	return nil
}

func (x *CallObject) GetSrcNumber() string {
	if x != nil {
		return x.SrcNumber
	}
	return ""
}

func (x *CallObject) GetDstNumber() string {
	if x != nil {
		return x.DstNumber
	}
	return ""
}

func (x *CallObject) GetCallerIdName() string {
	if x != nil {
		return x.CallerIdName
	}
	return ""
}

func (x *CallObject) GetAgentWorker() string {
	if x != nil {
		return x.AgentWorker
	}
	return ""
}

func (x *CallObject) GetEvents() []string {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *CallObject) GetCallData() string {
	if x != nil {
		return x.CallData
	}
	return ""
}

func (x *CallObject) GetAgentResponseData() map[string]string {
	if x != nil {
		return x.AgentResponseData
	}
	return nil
}

func (x *CallObject) GetRecorded() bool {
	if x != nil {
		return x.Recorded
	}
	return false
}

func (x *CallObject) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

func (x *CallObject) GetSuspended() bool {
	if x != nil {
		return x.Suspended
	}
	return false
}

func (x *CallObject) GetDisconnectReason() string {
	if x != nil {
		return x.DisconnectReason
	}
	return ""
}

func (x *CallObject) GetVoicemailed() bool {
	if x != nil {
		return x.Voicemailed
	}
	return false
}

func (x *CallObject) GetVoicemailBox() string {
	if x != nil {
		return x.VoicemailBox
	}
	return ""
}

func (x *CallObject) GetOriginated() string {
	if x != nil {
		return x.Originated
	}
	return ""
}

func (x *CallObject) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

func (x *CallObject) GetRtpInfo() string {
	if x != nil {
		return x.RtpInfo
	}
	return ""
}

type UpdateAgentResponseDataReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CallSid       int64                  `protobuf:"varint,2,opt,name=call_sid,json=callSid,proto3" json:"call_sid,omitempty"`
	CallType      commons.CallType_Enum  `protobuf:"varint,3,opt,name=call_type,json=callType,proto3,enum=api.commons.CallType_Enum" json:"call_type,omitempty"`
	Responses     map[string]string      `protobuf:"bytes,4,rep,name=responses,proto3" json:"responses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAgentResponseDataReq) Reset() {
	*x = UpdateAgentResponseDataReq{}
	mi := &file_api_v0alpha_sds_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAgentResponseDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAgentResponseDataReq) ProtoMessage() {}

func (x *UpdateAgentResponseDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sds_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAgentResponseDataReq.ProtoReflect.Descriptor instead.
func (*UpdateAgentResponseDataReq) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_sds_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAgentResponseDataReq) GetCallSid() int64 {
	if x != nil {
		return x.CallSid
	}
	return 0
}

func (x *UpdateAgentResponseDataReq) GetCallType() commons.CallType_Enum {
	if x != nil {
		return x.CallType
	}
	return commons.CallType_Enum(0)
}

func (x *UpdateAgentResponseDataReq) GetResponses() map[string]string {
	if x != nil {
		return x.Responses
	}
	return nil
}

type UpdateAgentResponseDataRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAgentResponseDataRes) Reset() {
	*x = UpdateAgentResponseDataRes{}
	mi := &file_api_v0alpha_sds_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAgentResponseDataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAgentResponseDataRes) ProtoMessage() {}

func (x *UpdateAgentResponseDataRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_sds_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAgentResponseDataRes.ProtoReflect.Descriptor instead.
func (*UpdateAgentResponseDataRes) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_sds_proto_rawDescGZIP(), []int{7}
}

var File_api_v0alpha_sds_proto protoreflect.FileDescriptor

var file_api_v0alpha_sds_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x22, 0xfe, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x51, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x60, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x09,
	0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x08, 0x63, 0x61, 0x6c,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x62, 0x78, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x62, 0x78, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x52, 0x65,
	0x73, 0x22, 0xeb, 0x07, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x69, 0x64, 0x12, 0x37, 0x0a,
	0x09, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x08, 0x63, 0x61,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x3b, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x72, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x72, 0x63, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x5e, 0x0a, 0x13, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x73,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x75,
	0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x62, 0x6f, 0x78, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x74, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x74, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x39,
	0x0a, 0x0b, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x44, 0x0a, 0x16, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x84, 0x02, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1c, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x32, 0xcc, 0x04, 0x0a, 0x03, 0x53, 0x64, 0x73, 0x12, 0x9b, 0x01, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x22, 0x37, 0xba, 0xb8, 0x91, 0x02, 0x02, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a,
	0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x73, 0x64, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x64, 0x61, 0x74, 0x61, 0x12, 0x67, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x17,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2a, 0xba, 0xb8, 0x91, 0x02, 0x02, 0x18, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x64, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x63,
	0x61, 0x6c, 0x6c, 0x12, 0xa7, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x22, 0x3a, 0xba, 0xb8, 0x91, 0x02, 0x02, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d,
	0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x73, 0x64, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x64, 0x61, 0x74, 0x61, 0x12, 0x93, 0x01,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x42, 0x6f, 0x78, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x22, 0x35, 0xba, 0xb8,
	0x91, 0x02, 0x02, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x64, 0x73,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x62, 0x6f, 0x78, 0x42, 0x90, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x08, 0x53, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x41, 0x56,
	0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca,
	0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x17,
	0x41, 0x70, 0x69, 0x5c, 0x56, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56,
	0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_v0alpha_sds_proto_rawDescOnce sync.Once
	file_api_v0alpha_sds_proto_rawDescData []byte
)

func file_api_v0alpha_sds_proto_rawDescGZIP() []byte {
	file_api_v0alpha_sds_proto_rawDescOnce.Do(func() {
		file_api_v0alpha_sds_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v0alpha_sds_proto_rawDesc), len(file_api_v0alpha_sds_proto_rawDesc)))
	})
	return file_api_v0alpha_sds_proto_rawDescData
}

var file_api_v0alpha_sds_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_v0alpha_sds_proto_goTypes = []any{
	(*GetAgentResponseDataReq)(nil),    // 0: api.v0alpha.GetAgentResponseDataReq
	(*GetAgentResponseDataRes)(nil),    // 1: api.v0alpha.GetAgentResponseDataRes
	(*GetCallReq)(nil),                 // 2: api.v0alpha.GetCallReq
	(*UpdateVoicemailBoxReq)(nil),      // 3: api.v0alpha.UpdateVoicemailBoxReq
	(*UpdateVoicemailBoxRes)(nil),      // 4: api.v0alpha.UpdateVoicemailBoxRes
	(*CallObject)(nil),                 // 5: api.v0alpha.CallObject
	(*UpdateAgentResponseDataReq)(nil), // 6: api.v0alpha.UpdateAgentResponseDataReq
	(*UpdateAgentResponseDataRes)(nil), // 7: api.v0alpha.UpdateAgentResponseDataRes
	nil,                                // 8: api.v0alpha.GetAgentResponseDataRes.ResponsesEntry
	nil,                                // 9: api.v0alpha.CallObject.SkillsEntry
	nil,                                // 10: api.v0alpha.CallObject.AgentResponseDataEntry
	nil,                                // 11: api.v0alpha.UpdateAgentResponseDataReq.ResponsesEntry
	(commons.CallType_Enum)(0),         // 12: api.commons.CallType.Enum
	(*timestamppb.Timestamp)(nil),      // 13: google.protobuf.Timestamp
}
var file_api_v0alpha_sds_proto_depIdxs = []int32{
	12, // 0: api.v0alpha.GetAgentResponseDataReq.call_type:type_name -> api.commons.CallType.Enum
	12, // 1: api.v0alpha.GetAgentResponseDataRes.call_type:type_name -> api.commons.CallType.Enum
	8,  // 2: api.v0alpha.GetAgentResponseDataRes.responses:type_name -> api.v0alpha.GetAgentResponseDataRes.ResponsesEntry
	12, // 3: api.v0alpha.GetCallReq.call_type:type_name -> api.commons.CallType.Enum
	12, // 4: api.v0alpha.UpdateVoicemailBoxReq.call_type:type_name -> api.commons.CallType.Enum
	12, // 5: api.v0alpha.CallObject.call_type:type_name -> api.commons.CallType.Enum
	9,  // 6: api.v0alpha.CallObject.skills:type_name -> api.v0alpha.CallObject.SkillsEntry
	13, // 7: api.v0alpha.CallObject.updated_date:type_name -> google.protobuf.Timestamp
	10, // 8: api.v0alpha.CallObject.agent_response_data:type_name -> api.v0alpha.CallObject.AgentResponseDataEntry
	12, // 9: api.v0alpha.UpdateAgentResponseDataReq.call_type:type_name -> api.commons.CallType.Enum
	11, // 10: api.v0alpha.UpdateAgentResponseDataReq.responses:type_name -> api.v0alpha.UpdateAgentResponseDataReq.ResponsesEntry
	0,  // 11: api.v0alpha.Sds.GetAgentResponseData:input_type -> api.v0alpha.GetAgentResponseDataReq
	2,  // 12: api.v0alpha.Sds.GetCall:input_type -> api.v0alpha.GetCallReq
	6,  // 13: api.v0alpha.Sds.UpdateAgentResponseData:input_type -> api.v0alpha.UpdateAgentResponseDataReq
	3,  // 14: api.v0alpha.Sds.UpdateVoicemailBox:input_type -> api.v0alpha.UpdateVoicemailBoxReq
	1,  // 15: api.v0alpha.Sds.GetAgentResponseData:output_type -> api.v0alpha.GetAgentResponseDataRes
	5,  // 16: api.v0alpha.Sds.GetCall:output_type -> api.v0alpha.CallObject
	7,  // 17: api.v0alpha.Sds.UpdateAgentResponseData:output_type -> api.v0alpha.UpdateAgentResponseDataRes
	4,  // 18: api.v0alpha.Sds.UpdateVoicemailBox:output_type -> api.v0alpha.UpdateVoicemailBoxRes
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_api_v0alpha_sds_proto_init() }
func file_api_v0alpha_sds_proto_init() {
	if File_api_v0alpha_sds_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v0alpha_sds_proto_rawDesc), len(file_api_v0alpha_sds_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v0alpha_sds_proto_goTypes,
		DependencyIndexes: file_api_v0alpha_sds_proto_depIdxs,
		MessageInfos:      file_api_v0alpha_sds_proto_msgTypes,
	}.Build()
	File_api_v0alpha_sds_proto = out.File
	file_api_v0alpha_sds_proto_goTypes = nil
	file_api_v0alpha_sds_proto_depIdxs = nil
}
