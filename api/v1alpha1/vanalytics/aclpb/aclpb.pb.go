// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/v1alpha1/vanalytics/aclpb/aclpb.proto

package aclpb

import (
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
	Actions []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*Action_CallSkillsInitial
	//	*Action_CallEnded
	Kind isAction_Kind `protobuf_oneof:"kind"`
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

func (m *Action) GetKind() isAction_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Action) GetCallSkillsInitial() *CallSkillsInitial {
	if x, ok := x.GetKind().(*Action_CallSkillsInitial); ok {
		return x.CallSkillsInitial
	}
	return nil
}

func (x *Action) GetCallEnded() string {
	if x, ok := x.GetKind().(*Action_CallEnded); ok {
		return x.CallEnded
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// need is all the call skills that are needed.
	Need []string `protobuf:"bytes,1,rep,name=need,proto3" json:"need,omitempty"`
	// want is all the call skills that are wanted.
	Want []string `protobuf:"bytes,2,rep,name=want,proto3" json:"want,omitempty"`
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

var file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x63, 0x6c, 0x70, 0x62, 0x2f,
	0x61, 0x63, 0x6c, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x63, 0x6c, 0x70, 0x62, 0x22, 0x4f, 0x0a, 0x0c, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x12, 0x3f, 0x0a, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x63, 0x6c, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x13, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x63,
	0x6c, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x11, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x65, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x77, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x77, 0x61, 0x6e, 0x74,
	0x42, 0x81, 0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x61, 0x63, 0x6c, 0x70, 0x62, 0x42, 0x0a, 0x41, 0x63, 0x6c, 0x70, 0x62, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x63, 0x6c, 0x70, 0x62, 0xa2, 0x02,
	0x04, 0x41, 0x56, 0x56, 0x41, 0xaa, 0x02, 0x1d, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x41, 0x63, 0x6c, 0x70, 0x62, 0xca, 0x02, 0x1d, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c,
	0x41, 0x63, 0x6c, 0x70, 0x62, 0xe2, 0x02, 0x29, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c,
	0x41, 0x63, 0x6c, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x20, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3a, 0x3a, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x3a, 0x3a, 0x41,
	0x63, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescOnce sync.Once
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescData = file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDesc
)

func file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDescData)
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
			RawDescriptor: file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDesc,
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
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_rawDesc = nil
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_goTypes = nil
	file_api_v1alpha1_vanalytics_aclpb_aclpb_proto_depIdxs = nil
}
