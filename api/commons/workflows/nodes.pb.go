// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/commons/workflows/nodes.proto

package workflows

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

// A node is a single step in a flow
type NodeDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// common fields for a node
	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Outputs     []string `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	ErrorNodeId string   `protobuf:"bytes,6,opt,name=error_node_id,json=errorNodeId,proto3" json:"error_node_id,omitempty"`
	// specific node definition
	//
	// Types that are assignable to Definition:
	//
	//	*NodeDefinition_Print
	//	*NodeDefinition_Random
	//	*NodeDefinition_ConsoleInput
	//	*NodeDefinition_Comparator
	//	*NodeDefinition_StoreInput
	//	*NodeDefinition_Chatbot
	//	*NodeDefinition_OmniPrompt
	//	*NodeDefinition_OmniSetSkill
	//	*NodeDefinition_OmniToAgent
	//	*NodeDefinition_OmniError
	//	*NodeDefinition_OmniBotTestStart
	//	*NodeDefinition_OmniBotTestStep
	//	*NodeDefinition_OmniBotTestEnd
	Definition isNodeDefinition_Definition `protobuf_oneof:"definition"`
}

func (x *NodeDefinition) Reset() {
	*x = NodeDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_workflows_nodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeDefinition) ProtoMessage() {}

func (x *NodeDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_workflows_nodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeDefinition.ProtoReflect.Descriptor instead.
func (*NodeDefinition) Descriptor() ([]byte, []int) {
	return file_api_commons_workflows_nodes_proto_rawDescGZIP(), []int{0}
}

func (x *NodeDefinition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NodeDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeDefinition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NodeDefinition) GetOutputs() []string {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *NodeDefinition) GetErrorNodeId() string {
	if x != nil {
		return x.ErrorNodeId
	}
	return ""
}

func (m *NodeDefinition) GetDefinition() isNodeDefinition_Definition {
	if m != nil {
		return m.Definition
	}
	return nil
}

func (x *NodeDefinition) GetPrint() *NodePrint {
	if x, ok := x.GetDefinition().(*NodeDefinition_Print); ok {
		return x.Print
	}
	return nil
}

func (x *NodeDefinition) GetRandom() *NodeRandom {
	if x, ok := x.GetDefinition().(*NodeDefinition_Random); ok {
		return x.Random
	}
	return nil
}

func (x *NodeDefinition) GetConsoleInput() *NodeConsoleInput {
	if x, ok := x.GetDefinition().(*NodeDefinition_ConsoleInput); ok {
		return x.ConsoleInput
	}
	return nil
}

func (x *NodeDefinition) GetComparator() *NodeComparator {
	if x, ok := x.GetDefinition().(*NodeDefinition_Comparator); ok {
		return x.Comparator
	}
	return nil
}

func (x *NodeDefinition) GetStoreInput() *NodeStoreInput {
	if x, ok := x.GetDefinition().(*NodeDefinition_StoreInput); ok {
		return x.StoreInput
	}
	return nil
}

func (x *NodeDefinition) GetChatbot() *NodeChatbot {
	if x, ok := x.GetDefinition().(*NodeDefinition_Chatbot); ok {
		return x.Chatbot
	}
	return nil
}

func (x *NodeDefinition) GetOmniPrompt() *OmniNodePrompt {
	if x, ok := x.GetDefinition().(*NodeDefinition_OmniPrompt); ok {
		return x.OmniPrompt
	}
	return nil
}

func (x *NodeDefinition) GetOmniSetSkill() *OmniNodeSetSkill {
	if x, ok := x.GetDefinition().(*NodeDefinition_OmniSetSkill); ok {
		return x.OmniSetSkill
	}
	return nil
}

func (x *NodeDefinition) GetOmniToAgent() *OmniNodeToAgent {
	if x, ok := x.GetDefinition().(*NodeDefinition_OmniToAgent); ok {
		return x.OmniToAgent
	}
	return nil
}

func (x *NodeDefinition) GetOmniError() *OmniNodeError {
	if x, ok := x.GetDefinition().(*NodeDefinition_OmniError); ok {
		return x.OmniError
	}
	return nil
}

func (x *NodeDefinition) GetOmniBotTestStart() *OmniBotNodeTestStart {
	if x, ok := x.GetDefinition().(*NodeDefinition_OmniBotTestStart); ok {
		return x.OmniBotTestStart
	}
	return nil
}

func (x *NodeDefinition) GetOmniBotTestStep() *OmniBotNodeTestStep {
	if x, ok := x.GetDefinition().(*NodeDefinition_OmniBotTestStep); ok {
		return x.OmniBotTestStep
	}
	return nil
}

func (x *NodeDefinition) GetOmniBotTestEnd() *OmniBotNodeTestEnd {
	if x, ok := x.GetDefinition().(*NodeDefinition_OmniBotTestEnd); ok {
		return x.OmniBotTestEnd
	}
	return nil
}

type isNodeDefinition_Definition interface {
	isNodeDefinition_Definition()
}

type NodeDefinition_Print struct {
	Print *NodePrint `protobuf:"bytes,101,opt,name=print,proto3,oneof"`
}

type NodeDefinition_Random struct {
	Random *NodeRandom `protobuf:"bytes,102,opt,name=random,proto3,oneof"`
}

type NodeDefinition_ConsoleInput struct {
	ConsoleInput *NodeConsoleInput `protobuf:"bytes,103,opt,name=console_input,json=consoleInput,proto3,oneof"`
}

type NodeDefinition_Comparator struct {
	Comparator *NodeComparator `protobuf:"bytes,104,opt,name=comparator,proto3,oneof"`
}

type NodeDefinition_StoreInput struct {
	StoreInput *NodeStoreInput `protobuf:"bytes,105,opt,name=store_input,json=storeInput,proto3,oneof"`
}

type NodeDefinition_Chatbot struct {
	Chatbot *NodeChatbot `protobuf:"bytes,1000,opt,name=chatbot,proto3,oneof"`
}

type NodeDefinition_OmniPrompt struct {
	OmniPrompt *OmniNodePrompt `protobuf:"bytes,201,opt,name=omni_prompt,json=omniPrompt,proto3,oneof"`
}

type NodeDefinition_OmniSetSkill struct {
	OmniSetSkill *OmniNodeSetSkill `protobuf:"bytes,202,opt,name=omni_set_skill,json=omniSetSkill,proto3,oneof"`
}

type NodeDefinition_OmniToAgent struct {
	OmniToAgent *OmniNodeToAgent `protobuf:"bytes,203,opt,name=omni_to_agent,json=omniToAgent,proto3,oneof"`
}

type NodeDefinition_OmniError struct {
	OmniError *OmniNodeError `protobuf:"bytes,204,opt,name=omni_error,json=omniError,proto3,oneof"`
}

type NodeDefinition_OmniBotTestStart struct {
	OmniBotTestStart *OmniBotNodeTestStart `protobuf:"bytes,301,opt,name=omni_bot_test_start,json=omniBotTestStart,proto3,oneof"`
}

type NodeDefinition_OmniBotTestStep struct {
	OmniBotTestStep *OmniBotNodeTestStep `protobuf:"bytes,302,opt,name=omni_bot_test_step,json=omniBotTestStep,proto3,oneof"`
}

type NodeDefinition_OmniBotTestEnd struct {
	OmniBotTestEnd *OmniBotNodeTestEnd `protobuf:"bytes,303,opt,name=omni_bot_test_end,json=omniBotTestEnd,proto3,oneof"`
}

func (*NodeDefinition_Print) isNodeDefinition_Definition() {}

func (*NodeDefinition_Random) isNodeDefinition_Definition() {}

func (*NodeDefinition_ConsoleInput) isNodeDefinition_Definition() {}

func (*NodeDefinition_Comparator) isNodeDefinition_Definition() {}

func (*NodeDefinition_StoreInput) isNodeDefinition_Definition() {}

func (*NodeDefinition_Chatbot) isNodeDefinition_Definition() {}

func (*NodeDefinition_OmniPrompt) isNodeDefinition_Definition() {}

func (*NodeDefinition_OmniSetSkill) isNodeDefinition_Definition() {}

func (*NodeDefinition_OmniToAgent) isNodeDefinition_Definition() {}

func (*NodeDefinition_OmniError) isNodeDefinition_Definition() {}

func (*NodeDefinition_OmniBotTestStart) isNodeDefinition_Definition() {}

func (*NodeDefinition_OmniBotTestStep) isNodeDefinition_Definition() {}

func (*NodeDefinition_OmniBotTestEnd) isNodeDefinition_Definition() {}

var File_api_commons_workflows_nodes_proto protoreflect.FileDescriptor

var file_api_commons_workflows_nodes_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x1a, 0x23, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x5f, 0x62, 0x6f,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x09, 0x0a, 0x0e, 0x4e, 0x6f, 0x64, 0x65,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x38,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x05, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x48, 0x00, 0x52, 0x06, 0x72,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x4e, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x47, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x68, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x48,
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x69, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74,
	0x62, 0x6f, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x48, 0x00,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x12, 0x49, 0x0a, 0x0b, 0x6f, 0x6d, 0x6e,
	0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x4f, 0x6d, 0x6e, 0x69, 0x4e, 0x6f, 0x64, 0x65,
	0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6d, 0x6e, 0x69, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x12, 0x50, 0x0a, 0x0e, 0x6f, 0x6d, 0x6e, 0x69, 0x5f, 0x73, 0x65, 0x74,
	0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0xca, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x4f, 0x6d, 0x6e, 0x69, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65,
	0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x6d, 0x6e, 0x69, 0x53, 0x65,
	0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x4d, 0x0a, 0x0d, 0x6f, 0x6d, 0x6e, 0x69, 0x5f, 0x74,
	0x6f, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0xcb, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x4f, 0x6d, 0x6e, 0x69, 0x4e, 0x6f, 0x64, 0x65, 0x54,
	0x6f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x6d, 0x6e, 0x69, 0x54, 0x6f,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x6f, 0x6d, 0x6e, 0x69, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0xcc, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x2e, 0x4f, 0x6d, 0x6e, 0x69, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x48, 0x00, 0x52, 0x09, 0x6f, 0x6d, 0x6e, 0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5d, 0x0a,
	0x13, 0x6f, 0x6d, 0x6e, 0x69, 0x5f, 0x62, 0x6f, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0xad, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x73, 0x2e, 0x4f, 0x6d, 0x6e, 0x69, 0x42, 0x6f, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x10, 0x6f, 0x6d, 0x6e, 0x69,
	0x42, 0x6f, 0x74, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x5a, 0x0a, 0x12,
	0x6f, 0x6d, 0x6e, 0x69, 0x5f, 0x62, 0x6f, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74,
	0x65, 0x70, 0x18, 0xae, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0x2e, 0x4f, 0x6d, 0x6e, 0x69, 0x42, 0x6f, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x73,
	0x74, 0x53, 0x74, 0x65, 0x70, 0x48, 0x00, 0x52, 0x0f, 0x6f, 0x6d, 0x6e, 0x69, 0x42, 0x6f, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x65, 0x70, 0x12, 0x57, 0x0a, 0x11, 0x6f, 0x6d, 0x6e, 0x69,
	0x5f, 0x62, 0x6f, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0xaf, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x4f, 0x6d, 0x6e,
	0x69, 0x42, 0x6f, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x48,
	0x00, 0x52, 0x0e, 0x6f, 0x6d, 0x6e, 0x69, 0x42, 0x6f, 0x74, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e,
	0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0xcf, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x0a, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0xa2, 0x02, 0x03,
	0x41, 0x43, 0x57, 0xaa, 0x02, 0x15, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0xca, 0x02, 0x15, 0x41, 0x70,
	0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x73, 0xe2, 0x02, 0x21, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x5c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_workflows_nodes_proto_rawDescOnce sync.Once
	file_api_commons_workflows_nodes_proto_rawDescData = file_api_commons_workflows_nodes_proto_rawDesc
)

func file_api_commons_workflows_nodes_proto_rawDescGZIP() []byte {
	file_api_commons_workflows_nodes_proto_rawDescOnce.Do(func() {
		file_api_commons_workflows_nodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_workflows_nodes_proto_rawDescData)
	})
	return file_api_commons_workflows_nodes_proto_rawDescData
}

var file_api_commons_workflows_nodes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_commons_workflows_nodes_proto_goTypes = []interface{}{
	(*NodeDefinition)(nil),       // 0: api.commons.workflows.NodeDefinition
	(*NodePrint)(nil),            // 1: api.commons.workflows.NodePrint
	(*NodeRandom)(nil),           // 2: api.commons.workflows.NodeRandom
	(*NodeConsoleInput)(nil),     // 3: api.commons.workflows.NodeConsoleInput
	(*NodeComparator)(nil),       // 4: api.commons.workflows.NodeComparator
	(*NodeStoreInput)(nil),       // 5: api.commons.workflows.NodeStoreInput
	(*NodeChatbot)(nil),          // 6: api.commons.workflows.NodeChatbot
	(*OmniNodePrompt)(nil),       // 7: api.commons.workflows.OmniNodePrompt
	(*OmniNodeSetSkill)(nil),     // 8: api.commons.workflows.OmniNodeSetSkill
	(*OmniNodeToAgent)(nil),      // 9: api.commons.workflows.OmniNodeToAgent
	(*OmniNodeError)(nil),        // 10: api.commons.workflows.OmniNodeError
	(*OmniBotNodeTestStart)(nil), // 11: api.commons.workflows.OmniBotNodeTestStart
	(*OmniBotNodeTestStep)(nil),  // 12: api.commons.workflows.OmniBotNodeTestStep
	(*OmniBotNodeTestEnd)(nil),   // 13: api.commons.workflows.OmniBotNodeTestEnd
}
var file_api_commons_workflows_nodes_proto_depIdxs = []int32{
	1,  // 0: api.commons.workflows.NodeDefinition.print:type_name -> api.commons.workflows.NodePrint
	2,  // 1: api.commons.workflows.NodeDefinition.random:type_name -> api.commons.workflows.NodeRandom
	3,  // 2: api.commons.workflows.NodeDefinition.console_input:type_name -> api.commons.workflows.NodeConsoleInput
	4,  // 3: api.commons.workflows.NodeDefinition.comparator:type_name -> api.commons.workflows.NodeComparator
	5,  // 4: api.commons.workflows.NodeDefinition.store_input:type_name -> api.commons.workflows.NodeStoreInput
	6,  // 5: api.commons.workflows.NodeDefinition.chatbot:type_name -> api.commons.workflows.NodeChatbot
	7,  // 6: api.commons.workflows.NodeDefinition.omni_prompt:type_name -> api.commons.workflows.OmniNodePrompt
	8,  // 7: api.commons.workflows.NodeDefinition.omni_set_skill:type_name -> api.commons.workflows.OmniNodeSetSkill
	9,  // 8: api.commons.workflows.NodeDefinition.omni_to_agent:type_name -> api.commons.workflows.OmniNodeToAgent
	10, // 9: api.commons.workflows.NodeDefinition.omni_error:type_name -> api.commons.workflows.OmniNodeError
	11, // 10: api.commons.workflows.NodeDefinition.omni_bot_test_start:type_name -> api.commons.workflows.OmniBotNodeTestStart
	12, // 11: api.commons.workflows.NodeDefinition.omni_bot_test_step:type_name -> api.commons.workflows.OmniBotNodeTestStep
	13, // 12: api.commons.workflows.NodeDefinition.omni_bot_test_end:type_name -> api.commons.workflows.OmniBotNodeTestEnd
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_api_commons_workflows_nodes_proto_init() }
func file_api_commons_workflows_nodes_proto_init() {
	if File_api_commons_workflows_nodes_proto != nil {
		return
	}
	file_api_commons_workflows_example_proto_init()
	file_api_commons_workflows_omni_proto_init()
	file_api_commons_workflows_omni_bot_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_commons_workflows_nodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeDefinition); i {
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
	file_api_commons_workflows_nodes_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*NodeDefinition_Print)(nil),
		(*NodeDefinition_Random)(nil),
		(*NodeDefinition_ConsoleInput)(nil),
		(*NodeDefinition_Comparator)(nil),
		(*NodeDefinition_StoreInput)(nil),
		(*NodeDefinition_Chatbot)(nil),
		(*NodeDefinition_OmniPrompt)(nil),
		(*NodeDefinition_OmniSetSkill)(nil),
		(*NodeDefinition_OmniToAgent)(nil),
		(*NodeDefinition_OmniError)(nil),
		(*NodeDefinition_OmniBotTestStart)(nil),
		(*NodeDefinition_OmniBotTestStep)(nil),
		(*NodeDefinition_OmniBotTestEnd)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_workflows_nodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_workflows_nodes_proto_goTypes,
		DependencyIndexes: file_api_commons_workflows_nodes_proto_depIdxs,
		MessageInfos:      file_api_commons_workflows_nodes_proto_msgTypes,
	}.Build()
	File_api_commons_workflows_nodes_proto = out.File
	file_api_commons_workflows_nodes_proto_rawDesc = nil
	file_api_commons_workflows_nodes_proto_goTypes = nil
	file_api_commons_workflows_nodes_proto_depIdxs = nil
}
