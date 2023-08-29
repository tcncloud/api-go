// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/commons/org/auth_connections.proto

package org

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

// ConnectionType defines the different provider types an auth connection can be.
type ConnectionType int32

const (
	ConnectionType_CONNECTION_TYPE_NONE  ConnectionType = 0
	ConnectionType_CONNECTION_TYPE_OIDC  ConnectionType = 1
	ConnectionType_CONNECTION_TYPE_AZURE ConnectionType = 2
)

// Enum value maps for ConnectionType.
var (
	ConnectionType_name = map[int32]string{
		0: "CONNECTION_TYPE_NONE",
		1: "CONNECTION_TYPE_OIDC",
		2: "CONNECTION_TYPE_AZURE",
	}
	ConnectionType_value = map[string]int32{
		"CONNECTION_TYPE_NONE":  0,
		"CONNECTION_TYPE_OIDC":  1,
		"CONNECTION_TYPE_AZURE": 2,
	}
)

func (x ConnectionType) Enum() *ConnectionType {
	p := new(ConnectionType)
	*p = x
	return p
}

func (x ConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_org_auth_connections_proto_enumTypes[0].Descriptor()
}

func (ConnectionType) Type() protoreflect.EnumType {
	return &file_api_commons_org_auth_connections_proto_enumTypes[0]
}

func (x ConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionType.Descriptor instead.
func (ConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_org_auth_connections_proto_rawDescGZIP(), []int{0}
}

// AuthConnectionSettings is the entity for sso connection information.
type AuthConnectionSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// issuer_url is the url provided by the identity provider
	// used to get authorization tokens.
	IssuerUrl string `protobuf:"bytes,1,opt,name=issuer_url,json=issuerUrl,proto3" json:"issuer_url,omitempty"`
	// tenant_url is the url used to match user's emails so
	// that the connection is used.
	TenantUrl string `protobuf:"bytes,2,opt,name=tenant_url,json=tenantUrl,proto3" json:"tenant_url,omitempty"`
	// client_id is the id of the application created from
	// the auth provider.
	ClientId string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// connection_id is the id of the auth0 connection.
	// This field will be populated automatically during
	// the CreateAuthConnection rpc.
	ConnectionId string `protobuf:"bytes,4,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// secret_expiration is an optional field that is
	// used to know when the connection's secret will
	// expire. Some providers have secret's that don't expire.
	SecretExpiration *AuthConnectionSettings_SecretExpiration `protobuf:"bytes,5,opt,name=secret_expiration,json=secretExpiration,proto3" json:"secret_expiration,omitempty"`
	// default_group defines settings used for all users for this connection
	// regardless of their groups. Group name is ignored for the default. This
	// field can be left nil if no default settings are desired.
	DefaultGroup *GroupItem `protobuf:"bytes,6,opt,name=default_group,json=defaultGroup,proto3" json:"default_group,omitempty"`
	// custom_groups defines the settings that will be used if a user with
	// this connection contains one or more of the matching group names.
	// If a user has multiple matching groups the settings for the highest
	// priority custom group will be used. Priority is determined by the custom
	// groups position in the list.
	CustomGroups []*GroupItem `protobuf:"bytes,7,rep,name=custom_groups,json=customGroups,proto3" json:"custom_groups,omitempty"`
	// org_id is the id of the organization the connection belongs too.
	OrgId string `protobuf:"bytes,8,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// name is the name of the connection.
	Name string `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	// type is what type of provider is used for the connection.
	Type ConnectionType `protobuf:"varint,10,opt,name=type,proto3,enum=api.commons.org.ConnectionType" json:"type,omitempty"`
}

func (x *AuthConnectionSettings) Reset() {
	*x = AuthConnectionSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_org_auth_connections_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthConnectionSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthConnectionSettings) ProtoMessage() {}

func (x *AuthConnectionSettings) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_auth_connections_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthConnectionSettings.ProtoReflect.Descriptor instead.
func (*AuthConnectionSettings) Descriptor() ([]byte, []int) {
	return file_api_commons_org_auth_connections_proto_rawDescGZIP(), []int{0}
}

func (x *AuthConnectionSettings) GetIssuerUrl() string {
	if x != nil {
		return x.IssuerUrl
	}
	return ""
}

func (x *AuthConnectionSettings) GetTenantUrl() string {
	if x != nil {
		return x.TenantUrl
	}
	return ""
}

func (x *AuthConnectionSettings) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AuthConnectionSettings) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *AuthConnectionSettings) GetSecretExpiration() *AuthConnectionSettings_SecretExpiration {
	if x != nil {
		return x.SecretExpiration
	}
	return nil
}

func (x *AuthConnectionSettings) GetDefaultGroup() *GroupItem {
	if x != nil {
		return x.DefaultGroup
	}
	return nil
}

func (x *AuthConnectionSettings) GetCustomGroups() []*GroupItem {
	if x != nil {
		return x.CustomGroups
	}
	return nil
}

func (x *AuthConnectionSettings) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *AuthConnectionSettings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthConnectionSettings) GetType() ConnectionType {
	if x != nil {
		return x.Type
	}
	return ConnectionType_CONNECTION_TYPE_NONE
}

// GroupItem defines settings mapped to a single group.
type GroupItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the group.
	GroupName string `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	// hunt group sid for the group.
	HuntGroupSid int64 `protobuf:"varint,2,opt,name=hunt_group_sid,json=huntGroupSid,proto3" json:"hunt_group_sid,omitempty"`
	// agent profile group for the group.
	AgentProfileGroupId string `protobuf:"bytes,3,opt,name=agent_profile_group_id,json=agentProfileGroupId,proto3" json:"agent_profile_group_id,omitempty"`
	// p3 permission group id for the group.
	P3PermissionGroupId string `protobuf:"bytes,4,opt,name=p3_permission_group_id,json=p3PermissionGroupId,proto3" json:"p3_permission_group_id,omitempty"`
	// list of permission group ids for the group. If a user belongs to multiple
	// groups the list of permission group ids will be appended to the existing
	// permissions the user contains. All other group settings will be overridden
	// by the highest priority group.
	PermissionGroupIds []string `protobuf:"bytes,5,rep,name=permission_group_ids,json=permissionGroupIds,proto3" json:"permission_group_ids,omitempty"`
}

func (x *GroupItem) Reset() {
	*x = GroupItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_org_auth_connections_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupItem) ProtoMessage() {}

func (x *GroupItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_auth_connections_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupItem.ProtoReflect.Descriptor instead.
func (*GroupItem) Descriptor() ([]byte, []int) {
	return file_api_commons_org_auth_connections_proto_rawDescGZIP(), []int{1}
}

func (x *GroupItem) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *GroupItem) GetHuntGroupSid() int64 {
	if x != nil {
		return x.HuntGroupSid
	}
	return 0
}

func (x *GroupItem) GetAgentProfileGroupId() string {
	if x != nil {
		return x.AgentProfileGroupId
	}
	return ""
}

func (x *GroupItem) GetP3PermissionGroupId() string {
	if x != nil {
		return x.P3PermissionGroupId
	}
	return ""
}

func (x *GroupItem) GetPermissionGroupIds() []string {
	if x != nil {
		return x.PermissionGroupIds
	}
	return nil
}

type AuthConnectionSettings_SecretExpiration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *AuthConnectionSettings_SecretExpiration) Reset() {
	*x = AuthConnectionSettings_SecretExpiration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_org_auth_connections_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthConnectionSettings_SecretExpiration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthConnectionSettings_SecretExpiration) ProtoMessage() {}

func (x *AuthConnectionSettings_SecretExpiration) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_auth_connections_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthConnectionSettings_SecretExpiration.ProtoReflect.Descriptor instead.
func (*AuthConnectionSettings_SecretExpiration) Descriptor() ([]byte, []int) {
	return file_api_commons_org_auth_connections_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AuthConnectionSettings_SecretExpiration) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

var File_api_commons_org_auth_connections_proto protoreflect.FileDescriptor

var file_api_commons_org_auth_connections_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x04, 0x0a, 0x16, 0x41,
	0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x65, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0d,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3f, 0x0a,
	0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x15,
	0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x42,
	0x0a, 0x10, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x22, 0xec, 0x01, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x68, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x70, 0x33,
	0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x33, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x14, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x73, 0x2a, 0x5f, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x18, 0x0a,
	0x14, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4f, 0x49, 0x44, 0x43, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x5a, 0x55, 0x52, 0x45,
	0x10, 0x02, 0x42, 0xb5, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x42, 0x14, 0x41, 0x75, 0x74, 0x68,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0xa2, 0x02,
	0x03, 0x41, 0x43, 0x4f, 0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0xca, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0xe2, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_commons_org_auth_connections_proto_rawDescOnce sync.Once
	file_api_commons_org_auth_connections_proto_rawDescData = file_api_commons_org_auth_connections_proto_rawDesc
)

func file_api_commons_org_auth_connections_proto_rawDescGZIP() []byte {
	file_api_commons_org_auth_connections_proto_rawDescOnce.Do(func() {
		file_api_commons_org_auth_connections_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_org_auth_connections_proto_rawDescData)
	})
	return file_api_commons_org_auth_connections_proto_rawDescData
}

var file_api_commons_org_auth_connections_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_org_auth_connections_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_commons_org_auth_connections_proto_goTypes = []interface{}{
	(ConnectionType)(0),                             // 0: api.commons.org.ConnectionType
	(*AuthConnectionSettings)(nil),                  // 1: api.commons.org.AuthConnectionSettings
	(*GroupItem)(nil),                               // 2: api.commons.org.GroupItem
	(*AuthConnectionSettings_SecretExpiration)(nil), // 3: api.commons.org.AuthConnectionSettings.SecretExpiration
	(*timestamppb.Timestamp)(nil),                   // 4: google.protobuf.Timestamp
}
var file_api_commons_org_auth_connections_proto_depIdxs = []int32{
	3, // 0: api.commons.org.AuthConnectionSettings.secret_expiration:type_name -> api.commons.org.AuthConnectionSettings.SecretExpiration
	2, // 1: api.commons.org.AuthConnectionSettings.default_group:type_name -> api.commons.org.GroupItem
	2, // 2: api.commons.org.AuthConnectionSettings.custom_groups:type_name -> api.commons.org.GroupItem
	0, // 3: api.commons.org.AuthConnectionSettings.type:type_name -> api.commons.org.ConnectionType
	4, // 4: api.commons.org.AuthConnectionSettings.SecretExpiration.date:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_commons_org_auth_connections_proto_init() }
func file_api_commons_org_auth_connections_proto_init() {
	if File_api_commons_org_auth_connections_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commons_org_auth_connections_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthConnectionSettings); i {
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
		file_api_commons_org_auth_connections_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupItem); i {
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
		file_api_commons_org_auth_connections_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthConnectionSettings_SecretExpiration); i {
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
			RawDescriptor: file_api_commons_org_auth_connections_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_org_auth_connections_proto_goTypes,
		DependencyIndexes: file_api_commons_org_auth_connections_proto_depIdxs,
		EnumInfos:         file_api_commons_org_auth_connections_proto_enumTypes,
		MessageInfos:      file_api_commons_org_auth_connections_proto_msgTypes,
	}.Build()
	File_api_commons_org_auth_connections_proto = out.File
	file_api_commons_org_auth_connections_proto_rawDesc = nil
	file_api_commons_org_auth_connections_proto_goTypes = nil
	file_api_commons_org_auth_connections_proto_depIdxs = nil
}
