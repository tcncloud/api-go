// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: services/omnichannel/asm/entities/v1alpha1/session.proto

package entitiesv1alpha1

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
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x0a, 0x41, 0x73, 0x6d, 0x53,
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
	0x70, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x73, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x42, 0xe3, 0x02, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x73, 0x6d, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02,
	0x04, 0x53, 0x4f, 0x41, 0x45, 0xaa, 0x02, 0x2a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x41, 0x73, 0x6d,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x2a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x4f, 0x6d,
	0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5c, 0x41, 0x73, 0x6d, 0x5c, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2,
	0x02, 0x36, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x4f, 0x6d, 0x6e, 0x69, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5c, 0x41, 0x73, 0x6d, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3a, 0x3a, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x3a, 0x3a, 0x41, 0x73, 0x6d, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_omnichannel_asm_entities_v1alpha1_session_proto_goTypes = []interface{}{
	(*AsmSession)(nil),            // 0: services.omnichannel.asm.entities.v1alpha1.AsmSession
	(*VoiceSession)(nil),          // 1: services.omnichannel.asm.entities.v1alpha1.VoiceSession
	(*VoiceRegistration)(nil),     // 2: services.omnichannel.asm.entities.v1alpha1.VoiceRegistration
	(*AsmUserDetails)(nil),        // 3: services.omnichannel.asm.entities.v1alpha1.AsmUserDetails
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_services_omnichannel_asm_entities_v1alpha1_session_proto_depIdxs = []int32{
	4, // 0: services.omnichannel.asm.entities.v1alpha1.AsmSession.asm_session_start:type_name -> google.protobuf.Timestamp
	4, // 1: services.omnichannel.asm.entities.v1alpha1.AsmSession.asm_session_end:type_name -> google.protobuf.Timestamp
	1, // 2: services.omnichannel.asm.entities.v1alpha1.AsmSession.voice_session:type_name -> services.omnichannel.asm.entities.v1alpha1.VoiceSession
	4, // 3: services.omnichannel.asm.entities.v1alpha1.VoiceSession.voice_session_start:type_name -> google.protobuf.Timestamp
	4, // 4: services.omnichannel.asm.entities.v1alpha1.VoiceSession.voice_session_end:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_services_omnichannel_asm_entities_v1alpha1_session_proto_init() }
func file_services_omnichannel_asm_entities_v1alpha1_session_proto_init() {
	if File_services_omnichannel_asm_entities_v1alpha1_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_omnichannel_asm_entities_v1alpha1_session_proto_goTypes,
		DependencyIndexes: file_services_omnichannel_asm_entities_v1alpha1_session_proto_depIdxs,
		MessageInfos:      file_services_omnichannel_asm_entities_v1alpha1_session_proto_msgTypes,
	}.Build()
	File_services_omnichannel_asm_entities_v1alpha1_session_proto = out.File
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_rawDesc = nil
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_goTypes = nil
	file_services_omnichannel_asm_entities_v1alpha1_session_proto_depIdxs = nil
}