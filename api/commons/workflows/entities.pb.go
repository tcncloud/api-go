// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/commons/workflows/entities.proto

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

// The overarching object that defines a flow
type FlowDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	EntryNodeId string            `protobuf:"bytes,4,opt,name=entry_node_id,json=entryNodeId,proto3" json:"entry_node_id,omitempty"`
	OrgId       string            `protobuf:"bytes,5,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Nodes       []*NodeDefinition `protobuf:"bytes,10,rep,name=nodes,proto3" json:"nodes,omitempty"`
	ErrorNodeId string            `protobuf:"bytes,11,opt,name=error_node_id,json=errorNodeId,proto3" json:"error_node_id,omitempty"`
}

func (x *FlowDefinition) Reset() {
	*x = FlowDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_workflows_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowDefinition) ProtoMessage() {}

func (x *FlowDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_workflows_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowDefinition.ProtoReflect.Descriptor instead.
func (*FlowDefinition) Descriptor() ([]byte, []int) {
	return file_api_commons_workflows_entities_proto_rawDescGZIP(), []int{0}
}

func (x *FlowDefinition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FlowDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlowDefinition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FlowDefinition) GetEntryNodeId() string {
	if x != nil {
		return x.EntryNodeId
	}
	return ""
}

func (x *FlowDefinition) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *FlowDefinition) GetNodes() []*NodeDefinition {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *FlowDefinition) GetErrorNodeId() string {
	if x != nil {
		return x.ErrorNodeId
	}
	return ""
}

type FlowState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envelope *Envelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
	State    *State    `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Payload  *Payload  `protobuf:"bytes,100,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *FlowState) Reset() {
	*x = FlowState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_workflows_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowState) ProtoMessage() {}

func (x *FlowState) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_workflows_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowState.ProtoReflect.Descriptor instead.
func (*FlowState) Descriptor() ([]byte, []int) {
	return file_api_commons_workflows_entities_proto_rawDescGZIP(), []int{1}
}

func (x *FlowState) GetEnvelope() *Envelope {
	if x != nil {
		return x.Envelope
	}
	return nil
}

func (x *FlowState) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *FlowState) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_workflows_entities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_workflows_entities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_api_commons_workflows_entities_proto_rawDescGZIP(), []int{2}
}

func (x *Envelope) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Envelope) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlowDone    bool   `protobuf:"varint,1,opt,name=flow_done,json=flowDone,proto3" json:"flow_done,omitempty"`
	FlowId      string `protobuf:"bytes,2,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	CurrentNode string `protobuf:"bytes,3,opt,name=current_node,json=currentNode,proto3" json:"current_node,omitempty"`
	Error       string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_workflows_entities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_workflows_entities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_api_commons_workflows_entities_proto_rawDescGZIP(), []int{3}
}

func (x *State) GetFlowDone() bool {
	if x != nil {
		return x.FlowDone
	}
	return false
}

func (x *State) GetFlowId() string {
	if x != nil {
		return x.FlowId
	}
	return ""
}

func (x *State) GetCurrentNode() string {
	if x != nil {
		return x.CurrentNode
	}
	return ""
}

func (x *State) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_workflows_entities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_workflows_entities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_api_commons_workflows_entities_proto_rawDescGZIP(), []int{4}
}

func (x *Payload) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Payload) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*Message_Text
	Payload isMessage_Payload `protobuf_oneof:"payload"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_workflows_entities_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_workflows_entities_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_commons_workflows_entities_proto_rawDescGZIP(), []int{5}
}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Message) GetText() string {
	if x, ok := x.GetPayload().(*Message_Text); ok {
		return x.Text
	}
	return ""
}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

func (*Message_Text) isMessage_Payload() {}

var File_api_commons_workflows_entities_proto protoreflect.FileDescriptor

var file_api_commons_workflows_entities_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x1a, 0x21, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf2, 0x01, 0x0a, 0x0e, 0x46, 0x6c, 0x6f, 0x77, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x67, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x12, 0x22, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0xb6, 0x01, 0x0a, 0x09, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x31,
	0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x22, 0x76, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66,
	0x6c, 0x6f, 0x77, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x33, 0x0a, 0x07, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2a,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42,
	0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0xd2, 0x01, 0x0a, 0x19, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43,
	0x57, 0xaa, 0x02, 0x15, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0xca, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0xe2, 0x02, 0x21, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_workflows_entities_proto_rawDescOnce sync.Once
	file_api_commons_workflows_entities_proto_rawDescData = file_api_commons_workflows_entities_proto_rawDesc
)

func file_api_commons_workflows_entities_proto_rawDescGZIP() []byte {
	file_api_commons_workflows_entities_proto_rawDescOnce.Do(func() {
		file_api_commons_workflows_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_workflows_entities_proto_rawDescData)
	})
	return file_api_commons_workflows_entities_proto_rawDescData
}

var file_api_commons_workflows_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_commons_workflows_entities_proto_goTypes = []interface{}{
	(*FlowDefinition)(nil), // 0: api.commons.workflows.FlowDefinition
	(*FlowState)(nil),      // 1: api.commons.workflows.FlowState
	(*Envelope)(nil),       // 2: api.commons.workflows.Envelope
	(*State)(nil),          // 3: api.commons.workflows.State
	(*Payload)(nil),        // 4: api.commons.workflows.Payload
	(*Message)(nil),        // 5: api.commons.workflows.Message
	(*NodeDefinition)(nil), // 6: api.commons.workflows.NodeDefinition
}
var file_api_commons_workflows_entities_proto_depIdxs = []int32{
	6, // 0: api.commons.workflows.FlowDefinition.nodes:type_name -> api.commons.workflows.NodeDefinition
	2, // 1: api.commons.workflows.FlowState.envelope:type_name -> api.commons.workflows.Envelope
	3, // 2: api.commons.workflows.FlowState.state:type_name -> api.commons.workflows.State
	4, // 3: api.commons.workflows.FlowState.payload:type_name -> api.commons.workflows.Payload
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_commons_workflows_entities_proto_init() }
func file_api_commons_workflows_entities_proto_init() {
	if File_api_commons_workflows_entities_proto != nil {
		return
	}
	file_api_commons_workflows_nodes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_commons_workflows_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowDefinition); i {
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
		file_api_commons_workflows_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowState); i {
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
		file_api_commons_workflows_entities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
		file_api_commons_workflows_entities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_api_commons_workflows_entities_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_api_commons_workflows_entities_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
	file_api_commons_workflows_entities_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Message_Text)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_workflows_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_workflows_entities_proto_goTypes,
		DependencyIndexes: file_api_commons_workflows_entities_proto_depIdxs,
		MessageInfos:      file_api_commons_workflows_entities_proto_msgTypes,
	}.Build()
	File_api_commons_workflows_entities_proto = out.File
	file_api_commons_workflows_entities_proto_rawDesc = nil
	file_api_commons_workflows_entities_proto_goTypes = nil
	file_api_commons_workflows_entities_proto_depIdxs = nil
}