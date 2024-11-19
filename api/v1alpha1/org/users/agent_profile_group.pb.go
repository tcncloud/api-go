// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/v1alpha1/org/users/agent_profile_group.proto

package users

import (
	org "github.com/tcncloud/api-go/api/commons/org"
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

// Request message for the GetAgentProfileGroup rpc.
type GetAgentProfileGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the agent profile group.
	AgentProfileGroupId string `protobuf:"bytes,1,opt,name=agent_profile_group_id,json=agentProfileGroupId,proto3" json:"agent_profile_group_id,omitempty"`
}

func (x *GetAgentProfileGroupRequest) Reset() {
	*x = GetAgentProfileGroupRequest{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAgentProfileGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentProfileGroupRequest) ProtoMessage() {}

func (x *GetAgentProfileGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentProfileGroupRequest.ProtoReflect.Descriptor instead.
func (*GetAgentProfileGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{0}
}

func (x *GetAgentProfileGroupRequest) GetAgentProfileGroupId() string {
	if x != nil {
		return x.AgentProfileGroupId
	}
	return ""
}

// Response message for the GetAgentProfileGroup rpc.
type GetAgentProfileGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Agent response group corresponding to the ID.
	AgentProfileGroup *org.AgentProfileGroup `protobuf:"bytes,1,opt,name=agent_profile_group,json=agentProfileGroup,proto3" json:"agent_profile_group,omitempty"`
}

func (x *GetAgentProfileGroupResponse) Reset() {
	*x = GetAgentProfileGroupResponse{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAgentProfileGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentProfileGroupResponse) ProtoMessage() {}

func (x *GetAgentProfileGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentProfileGroupResponse.ProtoReflect.Descriptor instead.
func (*GetAgentProfileGroupResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{1}
}

func (x *GetAgentProfileGroupResponse) GetAgentProfileGroup() *org.AgentProfileGroup {
	if x != nil {
		return x.AgentProfileGroup
	}
	return nil
}

// Request message for the UpdateAgentProfileGroup rpc.
type UpdateAgentProfileGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Agent response group to update
	AgentProfileGroup *org.AgentProfileGroup `protobuf:"bytes,1,opt,name=agent_profile_group,json=agentProfileGroup,proto3" json:"agent_profile_group,omitempty"`
}

func (x *UpdateAgentProfileGroupRequest) Reset() {
	*x = UpdateAgentProfileGroupRequest{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAgentProfileGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAgentProfileGroupRequest) ProtoMessage() {}

func (x *UpdateAgentProfileGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAgentProfileGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateAgentProfileGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAgentProfileGroupRequest) GetAgentProfileGroup() *org.AgentProfileGroup {
	if x != nil {
		return x.AgentProfileGroup
	}
	return nil
}

// Response message for the UpdateAgentProfileGroup rpc.
type UpdateAgentProfileGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAgentProfileGroupResponse) Reset() {
	*x = UpdateAgentProfileGroupResponse{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAgentProfileGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAgentProfileGroupResponse) ProtoMessage() {}

func (x *UpdateAgentProfileGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAgentProfileGroupResponse.ProtoReflect.Descriptor instead.
func (*UpdateAgentProfileGroupResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{3}
}

// Request message for the ListAgentProfileGroups rpc.
type ListAgentProfileGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAgentProfileGroupsRequest) Reset() {
	*x = ListAgentProfileGroupsRequest{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAgentProfileGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAgentProfileGroupsRequest) ProtoMessage() {}

func (x *ListAgentProfileGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAgentProfileGroupsRequest.ProtoReflect.Descriptor instead.
func (*ListAgentProfileGroupsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{4}
}

// Response message for the ListAgentProfileGroups rpc.
type ListAgentProfileGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of agent profile groups belonging to the org ID.
	AgentProfileGroups []*org.AgentProfileGroup `protobuf:"bytes,1,rep,name=agent_profile_groups,json=agentProfileGroups,proto3" json:"agent_profile_groups,omitempty"`
}

func (x *ListAgentProfileGroupsResponse) Reset() {
	*x = ListAgentProfileGroupsResponse{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAgentProfileGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAgentProfileGroupsResponse) ProtoMessage() {}

func (x *ListAgentProfileGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAgentProfileGroupsResponse.ProtoReflect.Descriptor instead.
func (*ListAgentProfileGroupsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{5}
}

func (x *ListAgentProfileGroupsResponse) GetAgentProfileGroups() []*org.AgentProfileGroup {
	if x != nil {
		return x.AgentProfileGroups
	}
	return nil
}

// Request message for the CreateAgentProfileGroup rpc.
type CreateAgentProfileGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Agent profile group to create.
	AgentProfileGroup *org.AgentProfileGroup `protobuf:"bytes,1,opt,name=agent_profile_group,json=agentProfileGroup,proto3" json:"agent_profile_group,omitempty"`
}

func (x *CreateAgentProfileGroupRequest) Reset() {
	*x = CreateAgentProfileGroupRequest{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAgentProfileGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAgentProfileGroupRequest) ProtoMessage() {}

func (x *CreateAgentProfileGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAgentProfileGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateAgentProfileGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{6}
}

func (x *CreateAgentProfileGroupRequest) GetAgentProfileGroup() *org.AgentProfileGroup {
	if x != nil {
		return x.AgentProfileGroup
	}
	return nil
}

// Response message for the CreateAgentProfileGroup rpc.
type CreateAgentProfileGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Created agent profile group.
	AgentProfileGroupId string `protobuf:"bytes,1,opt,name=agent_profile_group_id,json=agentProfileGroupId,proto3" json:"agent_profile_group_id,omitempty"`
}

func (x *CreateAgentProfileGroupResponse) Reset() {
	*x = CreateAgentProfileGroupResponse{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAgentProfileGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAgentProfileGroupResponse) ProtoMessage() {}

func (x *CreateAgentProfileGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAgentProfileGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateAgentProfileGroupResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{7}
}

func (x *CreateAgentProfileGroupResponse) GetAgentProfileGroupId() string {
	if x != nil {
		return x.AgentProfileGroupId
	}
	return ""
}

// Request message for the DeleteAgentProfileGroup rpc.
type DeleteAgentProfileGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Agent profile group ID to delete.
	AgentProfileGroupId string `protobuf:"bytes,1,opt,name=agent_profile_group_id,json=agentProfileGroupId,proto3" json:"agent_profile_group_id,omitempty"`
}

func (x *DeleteAgentProfileGroupRequest) Reset() {
	*x = DeleteAgentProfileGroupRequest{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAgentProfileGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAgentProfileGroupRequest) ProtoMessage() {}

func (x *DeleteAgentProfileGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAgentProfileGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteAgentProfileGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAgentProfileGroupRequest) GetAgentProfileGroupId() string {
	if x != nil {
		return x.AgentProfileGroupId
	}
	return ""
}

// Response message for the DeleteAgentProfileGroup rpc.
type DeleteAgentProfileGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAgentProfileGroupResponse) Reset() {
	*x = DeleteAgentProfileGroupResponse{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAgentProfileGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAgentProfileGroupResponse) ProtoMessage() {}

func (x *DeleteAgentProfileGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAgentProfileGroupResponse.ProtoReflect.Descriptor instead.
func (*DeleteAgentProfileGroupResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{9}
}

// Request message for the AssignAgentProfileGroups rpc.
type AssignAgentProfileGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Agent profile group ID to assign to provided users.
	AgentProfileGroupId string `protobuf:"bytes,1,opt,name=agent_profile_group_id,json=agentProfileGroupId,proto3" json:"agent_profile_group_id,omitempty"`
	// List of user IDs to assign to the agent profile group
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *AssignAgentProfileGroupsRequest) Reset() {
	*x = AssignAgentProfileGroupsRequest{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignAgentProfileGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignAgentProfileGroupsRequest) ProtoMessage() {}

func (x *AssignAgentProfileGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignAgentProfileGroupsRequest.ProtoReflect.Descriptor instead.
func (*AssignAgentProfileGroupsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{10}
}

func (x *AssignAgentProfileGroupsRequest) GetAgentProfileGroupId() string {
	if x != nil {
		return x.AgentProfileGroupId
	}
	return ""
}

func (x *AssignAgentProfileGroupsRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// Response message for the AssignAgentProfileGroups rpc.
type AssignAgentProfileGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssignAgentProfileGroupsResponse) Reset() {
	*x = AssignAgentProfileGroupsResponse{}
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignAgentProfileGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignAgentProfileGroupsResponse) ProtoMessage() {}

func (x *AssignAgentProfileGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignAgentProfileGroupsResponse.ProtoReflect.Descriptor instead.
func (*AssignAgentProfileGroupsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP(), []int{11}
}

var File_api_v1alpha1_org_users_agent_profile_group_proto protoreflect.FileDescriptor

var file_api_v1alpha1_org_users_agent_profile_group_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x29, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x13, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x11, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x74, 0x0a,
	0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x52, 0x0a, 0x13, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x11, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x22, 0x21, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x76, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x14, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x12, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22,
	0x74, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x52, 0x0a, 0x13, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x11, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x56, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x55, 0x0a,
	0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x33, 0x0a, 0x16, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x71, 0x0a, 0x1f, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x22, 0x0a, 0x20, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xe3,
	0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x16, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0xa2, 0x02, 0x04, 0x41, 0x56, 0x4f,
	0x55, 0xaa, 0x02, 0x16, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4f, 0x72, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0xca, 0x02, 0x16, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x55, 0x73,
	0x65, 0x72, 0x73, 0xe2, 0x02, 0x22, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescOnce sync.Once
	file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescData = file_api_v1alpha1_org_users_agent_profile_group_proto_rawDesc
)

func file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescData)
	})
	return file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescData
}

var file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_v1alpha1_org_users_agent_profile_group_proto_goTypes = []any{
	(*GetAgentProfileGroupRequest)(nil),      // 0: api.v1alpha1.org.users.GetAgentProfileGroupRequest
	(*GetAgentProfileGroupResponse)(nil),     // 1: api.v1alpha1.org.users.GetAgentProfileGroupResponse
	(*UpdateAgentProfileGroupRequest)(nil),   // 2: api.v1alpha1.org.users.UpdateAgentProfileGroupRequest
	(*UpdateAgentProfileGroupResponse)(nil),  // 3: api.v1alpha1.org.users.UpdateAgentProfileGroupResponse
	(*ListAgentProfileGroupsRequest)(nil),    // 4: api.v1alpha1.org.users.ListAgentProfileGroupsRequest
	(*ListAgentProfileGroupsResponse)(nil),   // 5: api.v1alpha1.org.users.ListAgentProfileGroupsResponse
	(*CreateAgentProfileGroupRequest)(nil),   // 6: api.v1alpha1.org.users.CreateAgentProfileGroupRequest
	(*CreateAgentProfileGroupResponse)(nil),  // 7: api.v1alpha1.org.users.CreateAgentProfileGroupResponse
	(*DeleteAgentProfileGroupRequest)(nil),   // 8: api.v1alpha1.org.users.DeleteAgentProfileGroupRequest
	(*DeleteAgentProfileGroupResponse)(nil),  // 9: api.v1alpha1.org.users.DeleteAgentProfileGroupResponse
	(*AssignAgentProfileGroupsRequest)(nil),  // 10: api.v1alpha1.org.users.AssignAgentProfileGroupsRequest
	(*AssignAgentProfileGroupsResponse)(nil), // 11: api.v1alpha1.org.users.AssignAgentProfileGroupsResponse
	(*org.AgentProfileGroup)(nil),            // 12: api.commons.org.AgentProfileGroup
}
var file_api_v1alpha1_org_users_agent_profile_group_proto_depIdxs = []int32{
	12, // 0: api.v1alpha1.org.users.GetAgentProfileGroupResponse.agent_profile_group:type_name -> api.commons.org.AgentProfileGroup
	12, // 1: api.v1alpha1.org.users.UpdateAgentProfileGroupRequest.agent_profile_group:type_name -> api.commons.org.AgentProfileGroup
	12, // 2: api.v1alpha1.org.users.ListAgentProfileGroupsResponse.agent_profile_groups:type_name -> api.commons.org.AgentProfileGroup
	12, // 3: api.v1alpha1.org.users.CreateAgentProfileGroupRequest.agent_profile_group:type_name -> api.commons.org.AgentProfileGroup
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_org_users_agent_profile_group_proto_init() }
func file_api_v1alpha1_org_users_agent_profile_group_proto_init() {
	if File_api_v1alpha1_org_users_agent_profile_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_org_users_agent_profile_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_org_users_agent_profile_group_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_org_users_agent_profile_group_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_org_users_agent_profile_group_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_org_users_agent_profile_group_proto = out.File
	file_api_v1alpha1_org_users_agent_profile_group_proto_rawDesc = nil
	file_api_v1alpha1_org_users_agent_profile_group_proto_goTypes = nil
	file_api_v1alpha1_org_users_agent_profile_group_proto_depIdxs = nil
}
