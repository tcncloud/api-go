// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/org/users/agent_profile_group.proto

package users

import (
	org "github.com/tcncloud/api-go/api/commons/org"
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

// Request message for the GetAgentProfileGroup rpc.
type GetAgentProfileGroupRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the agent profile group.
	AgentProfileGroupId string `protobuf:"bytes,1,opt,name=agent_profile_group_id,json=agentProfileGroupId,proto3" json:"agent_profile_group_id,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Agent response group corresponding to the ID.
	AgentProfileGroup *org.AgentProfileGroup `protobuf:"bytes,1,opt,name=agent_profile_group,json=agentProfileGroup,proto3" json:"agent_profile_group,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Agent response group to update
	AgentProfileGroup *org.AgentProfileGroup `protobuf:"bytes,1,opt,name=agent_profile_group,json=agentProfileGroup,proto3" json:"agent_profile_group,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of agent profile groups belonging to the org ID.
	AgentProfileGroups []*org.AgentProfileGroup `protobuf:"bytes,1,rep,name=agent_profile_groups,json=agentProfileGroups,proto3" json:"agent_profile_groups,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Agent profile group to create.
	AgentProfileGroup *org.AgentProfileGroup `protobuf:"bytes,1,opt,name=agent_profile_group,json=agentProfileGroup,proto3" json:"agent_profile_group,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Created agent profile group.
	AgentProfileGroupId string `protobuf:"bytes,1,opt,name=agent_profile_group_id,json=agentProfileGroupId,proto3" json:"agent_profile_group_id,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Agent profile group ID to delete.
	AgentProfileGroupId string `protobuf:"bytes,1,opt,name=agent_profile_group_id,json=agentProfileGroupId,proto3" json:"agent_profile_group_id,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Agent profile group ID to assign to provided users.
	AgentProfileGroupId string `protobuf:"bytes,1,opt,name=agent_profile_group_id,json=agentProfileGroupId,proto3" json:"agent_profile_group_id,omitempty"`
	// List of user IDs to assign to the agent profile group
	UserIds       []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

const file_api_v1alpha1_org_users_agent_profile_group_proto_rawDesc = "" +
	"\n" +
	"0api/v1alpha1/org/users/agent_profile_group.proto\x12\x16api.v1alpha1.org.users\x1a)api/commons/org/agent_profile_group.proto\"R\n" +
	"\x1bGetAgentProfileGroupRequest\x123\n" +
	"\x16agent_profile_group_id\x18\x01 \x01(\tR\x13agentProfileGroupId\"r\n" +
	"\x1cGetAgentProfileGroupResponse\x12R\n" +
	"\x13agent_profile_group\x18\x01 \x01(\v2\".api.commons.org.AgentProfileGroupR\x11agentProfileGroup\"t\n" +
	"\x1eUpdateAgentProfileGroupRequest\x12R\n" +
	"\x13agent_profile_group\x18\x01 \x01(\v2\".api.commons.org.AgentProfileGroupR\x11agentProfileGroup\"!\n" +
	"\x1fUpdateAgentProfileGroupResponse\"\x1f\n" +
	"\x1dListAgentProfileGroupsRequest\"v\n" +
	"\x1eListAgentProfileGroupsResponse\x12T\n" +
	"\x14agent_profile_groups\x18\x01 \x03(\v2\".api.commons.org.AgentProfileGroupR\x12agentProfileGroups\"t\n" +
	"\x1eCreateAgentProfileGroupRequest\x12R\n" +
	"\x13agent_profile_group\x18\x01 \x01(\v2\".api.commons.org.AgentProfileGroupR\x11agentProfileGroup\"V\n" +
	"\x1fCreateAgentProfileGroupResponse\x123\n" +
	"\x16agent_profile_group_id\x18\x01 \x01(\tR\x13agentProfileGroupId\"U\n" +
	"\x1eDeleteAgentProfileGroupRequest\x123\n" +
	"\x16agent_profile_group_id\x18\x01 \x01(\tR\x13agentProfileGroupId\"!\n" +
	"\x1fDeleteAgentProfileGroupResponse\"q\n" +
	"\x1fAssignAgentProfileGroupsRequest\x123\n" +
	"\x16agent_profile_group_id\x18\x01 \x01(\tR\x13agentProfileGroupId\x12\x19\n" +
	"\buser_ids\x18\x02 \x03(\tR\auserIds\"\"\n" +
	" AssignAgentProfileGroupsResponseB\xe3\x01\n" +
	"\x1acom.api.v1alpha1.org.usersB\x16AgentProfileGroupProtoP\x01Z1github.com/tcncloud/api-go/api/v1alpha1/org/users\xa2\x02\x04AVOU\xaa\x02\x16Api.V1alpha1.Org.Users\xca\x02\x16Api\\V1alpha1\\Org\\Users\xe2\x02\"Api\\V1alpha1\\Org\\Users\\GPBMetadata\xea\x02\x19Api::V1alpha1::Org::Usersb\x06proto3"

var (
	file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescOnce sync.Once
	file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescData []byte
)

func file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_org_users_agent_profile_group_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_org_users_agent_profile_group_proto_rawDesc), len(file_api_v1alpha1_org_users_agent_profile_group_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_org_users_agent_profile_group_proto_rawDesc), len(file_api_v1alpha1_org_users_agent_profile_group_proto_rawDesc)),
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
	file_api_v1alpha1_org_users_agent_profile_group_proto_goTypes = nil
	file_api_v1alpha1_org_users_agent_profile_group_proto_depIdxs = nil
}
