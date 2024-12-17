// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: api/v0alpha/cfg.proto

package v0alpha

import (
	_ "github.com/tcncloud/api-go/annotations"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetConfigReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RegionId      string                 `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConfigReq) Reset() {
	*x = GetConfigReq{}
	mi := &file_api_v0alpha_cfg_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigReq) ProtoMessage() {}

func (x *GetConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_cfg_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigReq.ProtoReflect.Descriptor instead.
func (*GetConfigReq) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_cfg_proto_rawDescGZIP(), []int{0}
}

func (x *GetConfigReq) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

type WebAgent struct {
	state             protoimpl.MessageState   `protogen:"open.v1"`
	OpenSips          *WebAgent_OpenSips       `protobuf:"bytes,1,opt,name=open_sips,json=openSips,proto3" json:"open_sips,omitempty"`
	EnginePriority    *WebAgent_EnginePriority `protobuf:"bytes,2,opt,name=engine_priority,json=enginePriority,proto3" json:"engine_priority,omitempty"`
	LogLevel          int64                    `protobuf:"varint,3,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	UseStun           int64                    `protobuf:"varint,4,opt,name=use_stun,json=useStun,proto3" json:"use_stun,omitempty"`
	UseFastStun       int64                    `protobuf:"varint,5,opt,name=use_fast_stun,json=useFastStun,proto3" json:"use_fast_stun,omitempty"`
	UseFastIce        int64                    `protobuf:"varint,6,opt,name=use_fast_ice,json=useFastIce,proto3" json:"use_fast_ice,omitempty"`
	IceTimeout        int64                    `protobuf:"varint,7,opt,name=ice_timeout,json=iceTimeout,proto3" json:"ice_timeout,omitempty"`
	SetFinalCodec     int64                    `protobuf:"varint,8,opt,name=set_final_codec,json=setFinalCodec,proto3" json:"set_final_codec,omitempty"`
	UseRport          int64                    `protobuf:"varint,9,opt,name=use_rport,json=useRport,proto3" json:"use_rport,omitempty"`
	Server            *WebAgent_Server         `protobuf:"bytes,10,opt,name=server,proto3" json:"server,omitempty"`
	BaseUrl           string                   `protobuf:"bytes,11,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	StunServerAddress string                   `protobuf:"bytes,12,opt,name=stun_server_address,json=stunServerAddress,proto3" json:"stun_server_address,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *WebAgent) Reset() {
	*x = WebAgent{}
	mi := &file_api_v0alpha_cfg_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebAgent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAgent) ProtoMessage() {}

func (x *WebAgent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_cfg_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAgent.ProtoReflect.Descriptor instead.
func (*WebAgent) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_cfg_proto_rawDescGZIP(), []int{1}
}

func (x *WebAgent) GetOpenSips() *WebAgent_OpenSips {
	if x != nil {
		return x.OpenSips
	}
	return nil
}

func (x *WebAgent) GetEnginePriority() *WebAgent_EnginePriority {
	if x != nil {
		return x.EnginePriority
	}
	return nil
}

func (x *WebAgent) GetLogLevel() int64 {
	if x != nil {
		return x.LogLevel
	}
	return 0
}

func (x *WebAgent) GetUseStun() int64 {
	if x != nil {
		return x.UseStun
	}
	return 0
}

func (x *WebAgent) GetUseFastStun() int64 {
	if x != nil {
		return x.UseFastStun
	}
	return 0
}

func (x *WebAgent) GetUseFastIce() int64 {
	if x != nil {
		return x.UseFastIce
	}
	return 0
}

func (x *WebAgent) GetIceTimeout() int64 {
	if x != nil {
		return x.IceTimeout
	}
	return 0
}

func (x *WebAgent) GetSetFinalCodec() int64 {
	if x != nil {
		return x.SetFinalCodec
	}
	return 0
}

func (x *WebAgent) GetUseRport() int64 {
	if x != nil {
		return x.UseRport
	}
	return 0
}

func (x *WebAgent) GetServer() *WebAgent_Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *WebAgent) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *WebAgent) GetStunServerAddress() string {
	if x != nil {
		return x.StunServerAddress
	}
	return ""
}

type WebAgent_OpenSips struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WebAgent_OpenSips) Reset() {
	*x = WebAgent_OpenSips{}
	mi := &file_api_v0alpha_cfg_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebAgent_OpenSips) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAgent_OpenSips) ProtoMessage() {}

func (x *WebAgent_OpenSips) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_cfg_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAgent_OpenSips.ProtoReflect.Descriptor instead.
func (*WebAgent_OpenSips) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_cfg_proto_rawDescGZIP(), []int{1, 0}
}

func (x *WebAgent_OpenSips) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type WebAgent_EnginePriority struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Java          int64                  `protobuf:"varint,1,opt,name=java,proto3" json:"java,omitempty"`
	Webrtc        int64                  `protobuf:"varint,2,opt,name=webrtc,proto3" json:"webrtc,omitempty"`
	Ns            int64                  `protobuf:"varint,3,opt,name=ns,proto3" json:"ns,omitempty"`
	Flash         int64                  `protobuf:"varint,4,opt,name=flash,proto3" json:"flash,omitempty"`
	App           int64                  `protobuf:"varint,5,opt,name=app,proto3" json:"app,omitempty"`
	P2P           int64                  `protobuf:"varint,6,opt,name=p2p,proto3" json:"p2p,omitempty"`
	AccessNum     int64                  `protobuf:"varint,7,opt,name=access_num,json=accessNum,proto3" json:"access_num,omitempty"`
	Native        int64                  `protobuf:"varint,8,opt,name=native,proto3" json:"native,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WebAgent_EnginePriority) Reset() {
	*x = WebAgent_EnginePriority{}
	mi := &file_api_v0alpha_cfg_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebAgent_EnginePriority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAgent_EnginePriority) ProtoMessage() {}

func (x *WebAgent_EnginePriority) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_cfg_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAgent_EnginePriority.ProtoReflect.Descriptor instead.
func (*WebAgent_EnginePriority) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_cfg_proto_rawDescGZIP(), []int{1, 1}
}

func (x *WebAgent_EnginePriority) GetJava() int64 {
	if x != nil {
		return x.Java
	}
	return 0
}

func (x *WebAgent_EnginePriority) GetWebrtc() int64 {
	if x != nil {
		return x.Webrtc
	}
	return 0
}

func (x *WebAgent_EnginePriority) GetNs() int64 {
	if x != nil {
		return x.Ns
	}
	return 0
}

func (x *WebAgent_EnginePriority) GetFlash() int64 {
	if x != nil {
		return x.Flash
	}
	return 0
}

func (x *WebAgent_EnginePriority) GetApp() int64 {
	if x != nil {
		return x.App
	}
	return 0
}

func (x *WebAgent_EnginePriority) GetP2P() int64 {
	if x != nil {
		return x.P2P
	}
	return 0
}

func (x *WebAgent_EnginePriority) GetAccessNum() int64 {
	if x != nil {
		return x.AccessNum
	}
	return 0
}

func (x *WebAgent_EnginePriority) GetNative() int64 {
	if x != nil {
		return x.Native
	}
	return 0
}

type WebAgent_Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sip           string                 `protobuf:"bytes,1,opt,name=sip,proto3" json:"sip,omitempty"`
	Webrtc        string                 `protobuf:"bytes,2,opt,name=webrtc,proto3" json:"webrtc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WebAgent_Server) Reset() {
	*x = WebAgent_Server{}
	mi := &file_api_v0alpha_cfg_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebAgent_Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAgent_Server) ProtoMessage() {}

func (x *WebAgent_Server) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0alpha_cfg_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAgent_Server.ProtoReflect.Descriptor instead.
func (*WebAgent_Server) Descriptor() ([]byte, []int) {
	return file_api_v0alpha_cfg_proto_rawDescGZIP(), []int{1, 2}
}

func (x *WebAgent_Server) GetSip() string {
	if x != nil {
		return x.Sip
	}
	return ""
}

func (x *WebAgent_Server) GetWebrtc() string {
	if x != nil {
		return x.Webrtc
	}
	return ""
}

var File_api_v0alpha_cfg_proto protoreflect.FileDescriptor

var file_api_v0alpha_cfg_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x66,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x95, 0x06, 0x0a, 0x08, 0x57, 0x65, 0x62,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x69,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x57, 0x65, 0x62, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x4f, 0x70, 0x65, 0x6e, 0x53, 0x69, 0x70, 0x73, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x53, 0x69,
	0x70, 0x73, 0x12, 0x4d, 0x0a, 0x0f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x57, 0x65, 0x62, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x0e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x75, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x75, 0x73, 0x65, 0x53, 0x74, 0x75, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65,
	0x5f, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x75, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x75, 0x73, 0x65, 0x46, 0x61, 0x73, 0x74, 0x53, 0x74, 0x75, 0x6e, 0x12, 0x20, 0x0a,
	0x0c, 0x75, 0x73, 0x65, 0x5f, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x46, 0x61, 0x73, 0x74, 0x49, 0x63, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x26, 0x0a, 0x0f, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x74, 0x46, 0x69,
	0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x5f,
	0x72, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x52, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x57, 0x65, 0x62, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x74, 0x75, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x24, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x69,
	0x70, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0xbd, 0x01, 0x0a,
	0x0e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6a, 0x61, 0x76, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6a,
	0x61, 0x76, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x6c, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x73,
	0x68, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x61, 0x70, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x32, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x70, 0x32, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x1a, 0x32, 0x0a, 0x06,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x62, 0x72,
	0x74, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63,
	0x32, 0x82, 0x01, 0x0a, 0x03, 0x43, 0x66, 0x67, 0x12, 0x7b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57,
	0x65, 0x62, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x57, 0x65, 0x62, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22,
	0x34, 0xba, 0xb8, 0x91, 0x02, 0x02, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01,
	0x2a, 0x22, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x63, 0x66, 0x67, 0x2f, 0x67, 0x65, 0x74, 0x77, 0x65, 0x62, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x90, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x08, 0x43, 0x66, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03,
	0x41, 0x56, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x30, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2,
	0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x56, 0x30, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v0alpha_cfg_proto_rawDescOnce sync.Once
	file_api_v0alpha_cfg_proto_rawDescData = file_api_v0alpha_cfg_proto_rawDesc
)

func file_api_v0alpha_cfg_proto_rawDescGZIP() []byte {
	file_api_v0alpha_cfg_proto_rawDescOnce.Do(func() {
		file_api_v0alpha_cfg_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v0alpha_cfg_proto_rawDescData)
	})
	return file_api_v0alpha_cfg_proto_rawDescData
}

var file_api_v0alpha_cfg_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v0alpha_cfg_proto_goTypes = []any{
	(*GetConfigReq)(nil),            // 0: api.v0alpha.GetConfigReq
	(*WebAgent)(nil),                // 1: api.v0alpha.WebAgent
	(*WebAgent_OpenSips)(nil),       // 2: api.v0alpha.WebAgent.OpenSips
	(*WebAgent_EnginePriority)(nil), // 3: api.v0alpha.WebAgent.EnginePriority
	(*WebAgent_Server)(nil),         // 4: api.v0alpha.WebAgent.Server
}
var file_api_v0alpha_cfg_proto_depIdxs = []int32{
	2, // 0: api.v0alpha.WebAgent.open_sips:type_name -> api.v0alpha.WebAgent.OpenSips
	3, // 1: api.v0alpha.WebAgent.engine_priority:type_name -> api.v0alpha.WebAgent.EnginePriority
	4, // 2: api.v0alpha.WebAgent.server:type_name -> api.v0alpha.WebAgent.Server
	0, // 3: api.v0alpha.Cfg.GetWebAgentConfig:input_type -> api.v0alpha.GetConfigReq
	1, // 4: api.v0alpha.Cfg.GetWebAgentConfig:output_type -> api.v0alpha.WebAgent
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v0alpha_cfg_proto_init() }
func file_api_v0alpha_cfg_proto_init() {
	if File_api_v0alpha_cfg_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v0alpha_cfg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v0alpha_cfg_proto_goTypes,
		DependencyIndexes: file_api_v0alpha_cfg_proto_depIdxs,
		MessageInfos:      file_api_v0alpha_cfg_proto_msgTypes,
	}.Build()
	File_api_v0alpha_cfg_proto = out.File
	file_api_v0alpha_cfg_proto_rawDesc = nil
	file_api_v0alpha_cfg_proto_goTypes = nil
	file_api_v0alpha_cfg_proto_depIdxs = nil
}
