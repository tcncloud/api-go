// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/org/authconnection/entities.proto

package authconnection

import (
	org "github.com/tcncloud/api-go/api/commons/org"
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

// Request message for the CreateAuthConnection rpc.
type CreateAuthConnectionRequest struct {
	state    protoimpl.MessageState      `protogen:"open.v1"`
	Settings *org.AuthConnectionSettings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	// client secret is the secret required for the provider.
	// Some providers don't allow for non expiring secrets. If this
	// is the case the secret_expiration field on the settings should
	// be set.
	// This field is not part of the ConnectionSettings message
	// because it should never be stored.
	ClientSecret  string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAuthConnectionRequest) Reset() {
	*x = CreateAuthConnectionRequest{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAuthConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthConnectionRequest) ProtoMessage() {}

func (x *CreateAuthConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthConnectionRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthConnectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAuthConnectionRequest) GetSettings() *org.AuthConnectionSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *CreateAuthConnectionRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

// Response message for the CreateAuthConnection rpc.
type CreateAuthConnectionResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the newly created auth connection.
	ConnectionId  string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAuthConnectionResponse) Reset() {
	*x = CreateAuthConnectionResponse{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAuthConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthConnectionResponse) ProtoMessage() {}

func (x *CreateAuthConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthConnectionResponse.ProtoReflect.Descriptor instead.
func (*CreateAuthConnectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAuthConnectionResponse) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

// Request message for the GetAuthConnectionSettings rpc.
type GetAuthConnectionSettingsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthConnectionSettingsRequest) Reset() {
	*x = GetAuthConnectionSettingsRequest{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthConnectionSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthConnectionSettingsRequest) ProtoMessage() {}

func (x *GetAuthConnectionSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthConnectionSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetAuthConnectionSettingsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{2}
}

// Response message for the GetConnectionSettings rpc.
type GetAuthConnectionSettingsResponse struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Settings      *org.AuthConnectionSettings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthConnectionSettingsResponse) Reset() {
	*x = GetAuthConnectionSettingsResponse{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthConnectionSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthConnectionSettingsResponse) ProtoMessage() {}

func (x *GetAuthConnectionSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthConnectionSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetAuthConnectionSettingsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuthConnectionSettingsResponse) GetSettings() *org.AuthConnectionSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

// Request message for the GetAuthConnection rpc.
type GetAuthConnectionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the connetion to retrieve.
	ConnectionId  string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthConnectionRequest) Reset() {
	*x = GetAuthConnectionRequest{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthConnectionRequest) ProtoMessage() {}

func (x *GetAuthConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthConnectionRequest.ProtoReflect.Descriptor instead.
func (*GetAuthConnectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{4}
}

func (x *GetAuthConnectionRequest) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

// Response message for the GetAuthConnection rpc.
type GetAuthConnectionResponse struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Settings      *org.AuthConnectionSettings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthConnectionResponse) Reset() {
	*x = GetAuthConnectionResponse{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthConnectionResponse) ProtoMessage() {}

func (x *GetAuthConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthConnectionResponse.ProtoReflect.Descriptor instead.
func (*GetAuthConnectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{5}
}

func (x *GetAuthConnectionResponse) GetSettings() *org.AuthConnectionSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

// Request message for the DeleteAuthConnection rpc.
type DeleteAuthConnectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ConnectionId  string                 `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAuthConnectionRequest) Reset() {
	*x = DeleteAuthConnectionRequest{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAuthConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthConnectionRequest) ProtoMessage() {}

func (x *DeleteAuthConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthConnectionRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthConnectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAuthConnectionRequest) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

// Response message for the DeleteAuthConnection rpc.
type DeleteAuthConnectionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAuthConnectionResponse) Reset() {
	*x = DeleteAuthConnectionResponse{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAuthConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthConnectionResponse) ProtoMessage() {}

func (x *DeleteAuthConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthConnectionResponse.ProtoReflect.Descriptor instead.
func (*DeleteAuthConnectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{7}
}

// Request message for the UpdateAuthConnectionSecret rpc.
type UpdateAuthConnectionSecretRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the connection that will be updated.
	ConnectionId string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// A secret associated with the provider.
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	// OPTIONAL: The date the secret will expire.
	SecretExpiration *UpdateAuthConnectionSecretRequest_SecretExpiration `protobuf:"bytes,3,opt,name=secret_expiration,json=secretExpiration,proto3" json:"secret_expiration,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *UpdateAuthConnectionSecretRequest) Reset() {
	*x = UpdateAuthConnectionSecretRequest{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAuthConnectionSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthConnectionSecretRequest) ProtoMessage() {}

func (x *UpdateAuthConnectionSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthConnectionSecretRequest.ProtoReflect.Descriptor instead.
func (*UpdateAuthConnectionSecretRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateAuthConnectionSecretRequest) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *UpdateAuthConnectionSecretRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *UpdateAuthConnectionSecretRequest) GetSecretExpiration() *UpdateAuthConnectionSecretRequest_SecretExpiration {
	if x != nil {
		return x.SecretExpiration
	}
	return nil
}

// Response message for the UpdateAuthConnectionSecret rpc.
type UpdateAuthConnectionSecretResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAuthConnectionSecretResponse) Reset() {
	*x = UpdateAuthConnectionSecretResponse{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAuthConnectionSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthConnectionSecretResponse) ProtoMessage() {}

func (x *UpdateAuthConnectionSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthConnectionSecretResponse.ProtoReflect.Descriptor instead.
func (*UpdateAuthConnectionSecretResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{9}
}

// Request message for the UpdateAuthConnectionGroups rpc.
type UpdateAuthConnectionGroupsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The default group. This group is assigned to all users
	// that don't match any custom groups.
	DefaultGroup *org.GroupItem `protobuf:"bytes,1,opt,name=default_group,json=defaultGroup,proto3" json:"default_group,omitempty"`
	// Custom groups. These groups are assigned to users that belong
	// to groups in the sso provider with the same name.
	CustomGroups []*org.GroupItem `protobuf:"bytes,2,rep,name=custom_groups,json=customGroups,proto3" json:"custom_groups,omitempty"`
	// the connection that will be updated.
	ConnectionId  string `protobuf:"bytes,3,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAuthConnectionGroupsRequest) Reset() {
	*x = UpdateAuthConnectionGroupsRequest{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAuthConnectionGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthConnectionGroupsRequest) ProtoMessage() {}

func (x *UpdateAuthConnectionGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthConnectionGroupsRequest.ProtoReflect.Descriptor instead.
func (*UpdateAuthConnectionGroupsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateAuthConnectionGroupsRequest) GetDefaultGroup() *org.GroupItem {
	if x != nil {
		return x.DefaultGroup
	}
	return nil
}

func (x *UpdateAuthConnectionGroupsRequest) GetCustomGroups() []*org.GroupItem {
	if x != nil {
		return x.CustomGroups
	}
	return nil
}

func (x *UpdateAuthConnectionGroupsRequest) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

// Response message for the UpdateAuthConnectionGroups rpc.
type UpdateAuthConnectionGroupsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAuthConnectionGroupsResponse) Reset() {
	*x = UpdateAuthConnectionGroupsResponse{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAuthConnectionGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthConnectionGroupsResponse) ProtoMessage() {}

func (x *UpdateAuthConnectionGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthConnectionGroupsResponse.ProtoReflect.Descriptor instead.
func (*UpdateAuthConnectionGroupsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{11}
}

// Request message for the ListAuthConnectionIds rpc.
type ListAuthConnectionIdsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAuthConnectionIdsRequest) Reset() {
	*x = ListAuthConnectionIdsRequest{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthConnectionIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthConnectionIdsRequest) ProtoMessage() {}

func (x *ListAuthConnectionIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthConnectionIdsRequest.ProtoReflect.Descriptor instead.
func (*ListAuthConnectionIdsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{12}
}

// Response message for the ListAuthConnectionIds rpc.
type ListAuthConnectionIdsResponse struct {
	state         protoimpl.MessageState                      `protogen:"open.v1"`
	Connections   []*ListAuthConnectionIdsResponse_Connection `protobuf:"bytes,1,rep,name=connections,proto3" json:"connections,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAuthConnectionIdsResponse) Reset() {
	*x = ListAuthConnectionIdsResponse{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthConnectionIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthConnectionIdsResponse) ProtoMessage() {}

func (x *ListAuthConnectionIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthConnectionIdsResponse.ProtoReflect.Descriptor instead.
func (*ListAuthConnectionIdsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{13}
}

func (x *ListAuthConnectionIdsResponse) GetConnections() []*ListAuthConnectionIdsResponse_Connection {
	if x != nil {
		return x.Connections
	}
	return nil
}

type UpdateAuthConnectionSecretRequest_SecretExpiration struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAuthConnectionSecretRequest_SecretExpiration) Reset() {
	*x = UpdateAuthConnectionSecretRequest_SecretExpiration{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAuthConnectionSecretRequest_SecretExpiration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthConnectionSecretRequest_SecretExpiration) ProtoMessage() {}

func (x *UpdateAuthConnectionSecretRequest_SecretExpiration) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthConnectionSecretRequest_SecretExpiration.ProtoReflect.Descriptor instead.
func (*UpdateAuthConnectionSecretRequest_SecretExpiration) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{8, 0}
}

func (x *UpdateAuthConnectionSecretRequest_SecretExpiration) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type ListAuthConnectionIdsResponse_Connection struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	AuthConnectionId string                 `protobuf:"bytes,1,opt,name=auth_connection_id,json=authConnectionId,proto3" json:"auth_connection_id,omitempty"`
	Name             string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ListAuthConnectionIdsResponse_Connection) Reset() {
	*x = ListAuthConnectionIdsResponse_Connection{}
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthConnectionIdsResponse_Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthConnectionIdsResponse_Connection) ProtoMessage() {}

func (x *ListAuthConnectionIdsResponse_Connection) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_authconnection_entities_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthConnectionIdsResponse_Connection.ProtoReflect.Descriptor instead.
func (*ListAuthConnectionIdsResponse_Connection) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP(), []int{13, 0}
}

func (x *ListAuthConnectionIdsResponse_Connection) GetAuthConnectionId() string {
	if x != nil {
		return x.AuthConnectionId
	}
	return ""
}

func (x *ListAuthConnectionIdsResponse_Connection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_v1alpha1_org_authconnection_entities_proto protoreflect.FileDescriptor

const file_api_v1alpha1_org_authconnection_entities_proto_rawDesc = "" +
	"\n" +
	".api/v1alpha1/org/authconnection/entities.proto\x12\x1fapi.v1alpha1.org.authconnection\x1a&api/commons/org/auth_connections.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x87\x01\n" +
	"\x1bCreateAuthConnectionRequest\x12C\n" +
	"\bsettings\x18\x01 \x01(\v2'.api.commons.org.AuthConnectionSettingsR\bsettings\x12#\n" +
	"\rclient_secret\x18\x02 \x01(\tR\fclientSecret\"C\n" +
	"\x1cCreateAuthConnectionResponse\x12#\n" +
	"\rconnection_id\x18\x01 \x01(\tR\fconnectionId\"\"\n" +
	" GetAuthConnectionSettingsRequest\"h\n" +
	"!GetAuthConnectionSettingsResponse\x12C\n" +
	"\bsettings\x18\x01 \x01(\v2'.api.commons.org.AuthConnectionSettingsR\bsettings\"?\n" +
	"\x18GetAuthConnectionRequest\x12#\n" +
	"\rconnection_id\x18\x01 \x01(\tR\fconnectionId\"`\n" +
	"\x19GetAuthConnectionResponse\x12C\n" +
	"\bsettings\x18\x01 \x01(\v2'.api.commons.org.AuthConnectionSettingsR\bsettings\"B\n" +
	"\x1bDeleteAuthConnectionRequest\x12#\n" +
	"\rconnection_id\x18\x01 \x01(\tR\fconnectionId\"\x1e\n" +
	"\x1cDeleteAuthConnectionResponse\"\xb4\x02\n" +
	"!UpdateAuthConnectionSecretRequest\x12#\n" +
	"\rconnection_id\x18\x01 \x01(\tR\fconnectionId\x12#\n" +
	"\rclient_secret\x18\x02 \x01(\tR\fclientSecret\x12\x80\x01\n" +
	"\x11secret_expiration\x18\x03 \x01(\v2S.api.v1alpha1.org.authconnection.UpdateAuthConnectionSecretRequest.SecretExpirationR\x10secretExpiration\x1aB\n" +
	"\x10SecretExpiration\x12.\n" +
	"\x04date\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\x04date\"$\n" +
	"\"UpdateAuthConnectionSecretResponse\"\xca\x01\n" +
	"!UpdateAuthConnectionGroupsRequest\x12?\n" +
	"\rdefault_group\x18\x01 \x01(\v2\x1a.api.commons.org.GroupItemR\fdefaultGroup\x12?\n" +
	"\rcustom_groups\x18\x02 \x03(\v2\x1a.api.commons.org.GroupItemR\fcustomGroups\x12#\n" +
	"\rconnection_id\x18\x03 \x01(\tR\fconnectionId\"$\n" +
	"\"UpdateAuthConnectionGroupsResponse\"\x1e\n" +
	"\x1cListAuthConnectionIdsRequest\"\xdc\x01\n" +
	"\x1dListAuthConnectionIdsResponse\x12k\n" +
	"\vconnections\x18\x01 \x03(\v2I.api.v1alpha1.org.authconnection.ListAuthConnectionIdsResponse.ConnectionR\vconnections\x1aN\n" +
	"\n" +
	"Connection\x12,\n" +
	"\x12auth_connection_id\x18\x01 \x01(\tR\x10authConnectionId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04nameB\x90\x02\n" +
	"#com.api.v1alpha1.org.authconnectionB\rEntitiesProtoP\x01Z:github.com/tcncloud/api-go/api/v1alpha1/org/authconnection\xa2\x02\x04AVOA\xaa\x02\x1fApi.V1alpha1.Org.Authconnection\xca\x02\x1fApi\\V1alpha1\\Org\\Authconnection\xe2\x02+Api\\V1alpha1\\Org\\Authconnection\\GPBMetadata\xea\x02\"Api::V1alpha1::Org::Authconnectionb\x06proto3"

var (
	file_api_v1alpha1_org_authconnection_entities_proto_rawDescOnce sync.Once
	file_api_v1alpha1_org_authconnection_entities_proto_rawDescData []byte
)

func file_api_v1alpha1_org_authconnection_entities_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_org_authconnection_entities_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_org_authconnection_entities_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_org_authconnection_entities_proto_rawDesc), len(file_api_v1alpha1_org_authconnection_entities_proto_rawDesc)))
	})
	return file_api_v1alpha1_org_authconnection_entities_proto_rawDescData
}

var file_api_v1alpha1_org_authconnection_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_api_v1alpha1_org_authconnection_entities_proto_goTypes = []any{
	(*CreateAuthConnectionRequest)(nil),                        // 0: api.v1alpha1.org.authconnection.CreateAuthConnectionRequest
	(*CreateAuthConnectionResponse)(nil),                       // 1: api.v1alpha1.org.authconnection.CreateAuthConnectionResponse
	(*GetAuthConnectionSettingsRequest)(nil),                   // 2: api.v1alpha1.org.authconnection.GetAuthConnectionSettingsRequest
	(*GetAuthConnectionSettingsResponse)(nil),                  // 3: api.v1alpha1.org.authconnection.GetAuthConnectionSettingsResponse
	(*GetAuthConnectionRequest)(nil),                           // 4: api.v1alpha1.org.authconnection.GetAuthConnectionRequest
	(*GetAuthConnectionResponse)(nil),                          // 5: api.v1alpha1.org.authconnection.GetAuthConnectionResponse
	(*DeleteAuthConnectionRequest)(nil),                        // 6: api.v1alpha1.org.authconnection.DeleteAuthConnectionRequest
	(*DeleteAuthConnectionResponse)(nil),                       // 7: api.v1alpha1.org.authconnection.DeleteAuthConnectionResponse
	(*UpdateAuthConnectionSecretRequest)(nil),                  // 8: api.v1alpha1.org.authconnection.UpdateAuthConnectionSecretRequest
	(*UpdateAuthConnectionSecretResponse)(nil),                 // 9: api.v1alpha1.org.authconnection.UpdateAuthConnectionSecretResponse
	(*UpdateAuthConnectionGroupsRequest)(nil),                  // 10: api.v1alpha1.org.authconnection.UpdateAuthConnectionGroupsRequest
	(*UpdateAuthConnectionGroupsResponse)(nil),                 // 11: api.v1alpha1.org.authconnection.UpdateAuthConnectionGroupsResponse
	(*ListAuthConnectionIdsRequest)(nil),                       // 12: api.v1alpha1.org.authconnection.ListAuthConnectionIdsRequest
	(*ListAuthConnectionIdsResponse)(nil),                      // 13: api.v1alpha1.org.authconnection.ListAuthConnectionIdsResponse
	(*UpdateAuthConnectionSecretRequest_SecretExpiration)(nil), // 14: api.v1alpha1.org.authconnection.UpdateAuthConnectionSecretRequest.SecretExpiration
	(*ListAuthConnectionIdsResponse_Connection)(nil),           // 15: api.v1alpha1.org.authconnection.ListAuthConnectionIdsResponse.Connection
	(*org.AuthConnectionSettings)(nil),                         // 16: api.commons.org.AuthConnectionSettings
	(*org.GroupItem)(nil),                                      // 17: api.commons.org.GroupItem
	(*timestamppb.Timestamp)(nil),                              // 18: google.protobuf.Timestamp
}
var file_api_v1alpha1_org_authconnection_entities_proto_depIdxs = []int32{
	16, // 0: api.v1alpha1.org.authconnection.CreateAuthConnectionRequest.settings:type_name -> api.commons.org.AuthConnectionSettings
	16, // 1: api.v1alpha1.org.authconnection.GetAuthConnectionSettingsResponse.settings:type_name -> api.commons.org.AuthConnectionSettings
	16, // 2: api.v1alpha1.org.authconnection.GetAuthConnectionResponse.settings:type_name -> api.commons.org.AuthConnectionSettings
	14, // 3: api.v1alpha1.org.authconnection.UpdateAuthConnectionSecretRequest.secret_expiration:type_name -> api.v1alpha1.org.authconnection.UpdateAuthConnectionSecretRequest.SecretExpiration
	17, // 4: api.v1alpha1.org.authconnection.UpdateAuthConnectionGroupsRequest.default_group:type_name -> api.commons.org.GroupItem
	17, // 5: api.v1alpha1.org.authconnection.UpdateAuthConnectionGroupsRequest.custom_groups:type_name -> api.commons.org.GroupItem
	15, // 6: api.v1alpha1.org.authconnection.ListAuthConnectionIdsResponse.connections:type_name -> api.v1alpha1.org.authconnection.ListAuthConnectionIdsResponse.Connection
	18, // 7: api.v1alpha1.org.authconnection.UpdateAuthConnectionSecretRequest.SecretExpiration.date:type_name -> google.protobuf.Timestamp
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_org_authconnection_entities_proto_init() }
func file_api_v1alpha1_org_authconnection_entities_proto_init() {
	if File_api_v1alpha1_org_authconnection_entities_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_org_authconnection_entities_proto_rawDesc), len(file_api_v1alpha1_org_authconnection_entities_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_org_authconnection_entities_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_org_authconnection_entities_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_org_authconnection_entities_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_org_authconnection_entities_proto = out.File
	file_api_v1alpha1_org_authconnection_entities_proto_goTypes = nil
	file_api_v1alpha1_org_authconnection_entities_proto_depIdxs = nil
}
