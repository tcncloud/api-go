// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/vanalytics/aclpb/aclpb.proto

package aclpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// AgentCallLog resource.
type AgentCallLog struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// actions are the agent call log actions.
	Actions       []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AgentCallLog) Reset() {
	*x = AgentCallLog{}
	mi := &file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentCallLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCallLog) ProtoMessage() {}

func (x *AgentCallLog) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_msgTypes[0]
	if x != nil {
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
	return file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescGZIP(), []int{0}
}

func (x *AgentCallLog) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

// Action is an agent call log action.
type Action struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Kind:
	//
	//	*Action_CallSkillsInitial
	//	*Action_CallEnded
	Kind          isAction_Kind `protobuf_oneof:"kind"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Action) Reset() {
	*x = Action{}
	mi := &file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescGZIP(), []int{1}
}

func (x *Action) GetKind() isAction_Kind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *Action) GetCallSkillsInitial() *CallSkillsInitial {
	if x != nil {
		if x, ok := x.Kind.(*Action_CallSkillsInitial); ok {
			return x.CallSkillsInitial
		}
	}
	return nil
}

func (x *Action) GetCallEnded() string {
	if x != nil {
		if x, ok := x.Kind.(*Action_CallEnded); ok {
			return x.CallEnded
		}
	}
	return ""
}

type isAction_Kind interface {
	isAction_Kind()
}

type Action_CallSkillsInitial struct {
	// CallSkillsInitial are the initial call skills on a call.
	CallSkillsInitial *CallSkillsInitial `protobuf:"bytes,1,opt,name=call_skills_initial,json=callSkillsInitial,proto3,oneof"`
}

type Action_CallEnded struct {
	// CallEnded is the reason the call ended.
	CallEnded string `protobuf:"bytes,2,opt,name=call_ended,json=callEnded,proto3,oneof"`
}

func (*Action_CallSkillsInitial) isAction_Kind() {}

func (*Action_CallEnded) isAction_Kind() {}

// CallSkillsInitial are the initial call skills on a call.
type CallSkillsInitial struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// need is all the call skills that are needed.
	Need []string `protobuf:"bytes,1,rep,name=need,proto3" json:"need,omitempty"`
	// want is all the call skills that are wanted.
	Want          []string `protobuf:"bytes,2,rep,name=want,proto3" json:"want,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CallSkillsInitial) Reset() {
	*x = CallSkillsInitial{}
	mi := &file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallSkillsInitial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallSkillsInitial) ProtoMessage() {}

func (x *CallSkillsInitial) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallSkillsInitial.ProtoReflect.Descriptor instead.
func (*CallSkillsInitial) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescGZIP(), []int{2}
}

func (x *CallSkillsInitial) GetNeed() []string {
	if x != nil {
		return x.Need
	}
	return nil
}

func (x *CallSkillsInitial) GetWant() []string {
	if x != nil {
		return x.Want
	}
	return nil
}

var File_api_v1alpha1_vanalytics_aclpb_aclpb_proto protoreflect.FileDescriptor

const file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDesc = "" +
	"\n" +
	")api/v1alpha1/vanalytics/aclpb/aclpb.proto\x12\x1dapi.v1alpha1.vanalytics.aclpb\"O\n" +
	"\fAgentCallLog\x12?\n" +
	"\aactions\x18\x01 \x03(\v2%.api.v1alpha1.vanalytics.aclpb.ActionR\aactions\"\x95\x01\n" +
	"\x06Action\x12b\n" +
	"\x13call_skills_initial\x18\x01 \x01(\v20.api.v1alpha1.vanalytics.aclpb.CallSkillsInitialH\x00R\x11callSkillsInitial\x12\x1f\n" +
	"\n" +
	"call_ended\x18\x02 \x01(\tH\x00R\tcallEndedB\x06\n" +
	"\x04kind\";\n" +
	"\x11CallSkillsInitial\x12\x12\n" +
	"\x04need\x18\x01 \x03(\tR\x04need\x12\x12\n" +
	"\x04want\x18\x02 \x03(\tR\x04wantB\x81\x02\n" +
	"!com.api.v1alpha1.vanalytics.aclpbB\n" +
	"AclpbProtoP\x01Z8github.com/tcncloud/api-go/api/v1alpha1/vanalytics/aclpb\xa2\x02\x04AVVA\xaa\x02\x1dApi.V1alpha1.Vanalytics.Aclpb\xca\x02\x1dApi\\V1alpha1\\Vanalytics\\Aclpb\xe2\x02)Api\\V1alpha1\\Vanalytics\\Aclpb\\GPBMetadata\xea\x02 Api::V1alpha1::Vanalytics::Aclpbb\x06proto3"

var (
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescOnce sync.Once
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescData []byte
)

func file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDesc), len(file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDesc)))
	})
	return file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescData
}

var file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_goTypes = []any{
	(*AgentCallLog)(nil),      // 0: api.v1alpha1.vanalytics.aclpb.AgentCallLog
	(*Action)(nil),            // 1: api.v1alpha1.vanalytics.aclpb.Action
	(*CallSkillsInitial)(nil), // 2: api.v1alpha1.vanalytics.aclpb.CallSkillsInitial
}
var file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_depIdxs = []int32{
	1, // 0: api.v1alpha1.vanalytics.aclpb.AgentCallLog.actions:type_name -> api.v1alpha1.vanalytics.aclpb.Action
	2, // 1: api.v1alpha1.vanalytics.aclpb.Action.call_skills_initial:type_name -> api.v1alpha1.vanalytics.aclpb.CallSkillsInitial
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_init() }
func file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_init() {
	if File_api_v1alpha1_vanalytics_aclpb_aclpb_proto != nil {
		return
	}
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_msgTypes[1].OneofWrappers = []any{
		(*Action_CallSkillsInitial)(nil),
		(*Action_CallEnded)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDesc), len(file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_vanalytics_aclpb_aclpb_proto = out.File
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_goTypes = nil
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_depIdxs = nil
}
