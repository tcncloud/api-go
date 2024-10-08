// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/v1alpha1/org/users/auth_token.proto

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

// Request message for the CreateAuthToken rpc.
type CreateAuthTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAuthTokenRequest) Reset() {
	*x = CreateAuthTokenRequest{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthTokenRequest) ProtoMessage() {}

func (x *CreateAuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{0}
}

// Response message for the CreateAuthToken rpc.
type CreateAuthTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Newly generated auth token.
	AuthToken *org.AuthToken `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
}

func (x *CreateAuthTokenResponse) Reset() {
	*x = CreateAuthTokenResponse{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAuthTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthTokenResponse) ProtoMessage() {}

func (x *CreateAuthTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateAuthTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAuthTokenResponse) GetAuthToken() *org.AuthToken {
	if x != nil {
		return x.AuthToken
	}
	return nil
}

// Request message for the CreateAuthTokenByUserId rpc.
type CreateAuthTokenByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User id creating new token for.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateAuthTokenByUserIdRequest) Reset() {
	*x = CreateAuthTokenByUserIdRequest{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAuthTokenByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthTokenByUserIdRequest) ProtoMessage() {}

func (x *CreateAuthTokenByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthTokenByUserIdRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthTokenByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAuthTokenByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Response message for the CreateAuthTokenByUserId rpc.
type CreateAuthTokenByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Newly generated auth token.
	AuthToken *org.AuthToken `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
}

func (x *CreateAuthTokenByUserIdResponse) Reset() {
	*x = CreateAuthTokenByUserIdResponse{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAuthTokenByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthTokenByUserIdResponse) ProtoMessage() {}

func (x *CreateAuthTokenByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthTokenByUserIdResponse.ProtoReflect.Descriptor instead.
func (*CreateAuthTokenByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAuthTokenByUserIdResponse) GetAuthToken() *org.AuthToken {
	if x != nil {
		return x.AuthToken
	}
	return nil
}

// Request message for the ListAuthTokensRequest rpc.
type ListAuthTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAuthTokensRequest) Reset() {
	*x = ListAuthTokensRequest{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthTokensRequest) ProtoMessage() {}

func (x *ListAuthTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthTokensRequest.ProtoReflect.Descriptor instead.
func (*ListAuthTokensRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{4}
}

// Response message for the ListAuthTokensResponse rpc.
type ListAuthTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of current users' auth tokens.
	AuthTokens []*org.AuthToken `protobuf:"bytes,1,rep,name=auth_tokens,json=authTokens,proto3" json:"auth_tokens,omitempty"`
}

func (x *ListAuthTokensResponse) Reset() {
	*x = ListAuthTokensResponse{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthTokensResponse) ProtoMessage() {}

func (x *ListAuthTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthTokensResponse.ProtoReflect.Descriptor instead.
func (*ListAuthTokensResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{5}
}

func (x *ListAuthTokensResponse) GetAuthTokens() []*org.AuthToken {
	if x != nil {
		return x.AuthTokens
	}
	return nil
}

// Request message for the ListAuthTokensByUserIdRequest rpc.
type ListAuthTokensByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User id to get list of auth tokens.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListAuthTokensByUserIdRequest) Reset() {
	*x = ListAuthTokensByUserIdRequest{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthTokensByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthTokensByUserIdRequest) ProtoMessage() {}

func (x *ListAuthTokensByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthTokensByUserIdRequest.ProtoReflect.Descriptor instead.
func (*ListAuthTokensByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{6}
}

func (x *ListAuthTokensByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Response message for the ListAuthTokensByUserIdResponse rpc.
type ListAuthTokensByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of requested users auth tokens.
	AuthTokens []*org.AuthToken `protobuf:"bytes,1,rep,name=auth_tokens,json=authTokens,proto3" json:"auth_tokens,omitempty"`
}

func (x *ListAuthTokensByUserIdResponse) Reset() {
	*x = ListAuthTokensByUserIdResponse{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthTokensByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthTokensByUserIdResponse) ProtoMessage() {}

func (x *ListAuthTokensByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthTokensByUserIdResponse.ProtoReflect.Descriptor instead.
func (*ListAuthTokensByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{7}
}

func (x *ListAuthTokensByUserIdResponse) GetAuthTokens() []*org.AuthToken {
	if x != nil {
		return x.AuthTokens
	}
	return nil
}

// Request message for the SetAuthTokenExpirationRequest rpc.
type SetAuthTokenExpirationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to reset expiration for that belongs to the current user.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *SetAuthTokenExpirationRequest) Reset() {
	*x = SetAuthTokenExpirationRequest{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAuthTokenExpirationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAuthTokenExpirationRequest) ProtoMessage() {}

func (x *SetAuthTokenExpirationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAuthTokenExpirationRequest.ProtoReflect.Descriptor instead.
func (*SetAuthTokenExpirationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{8}
}

func (x *SetAuthTokenExpirationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Response message for the SetAuthTokenExpirationResponse rpc.
type SetAuthTokenExpirationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetAuthTokenExpirationResponse) Reset() {
	*x = SetAuthTokenExpirationResponse{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAuthTokenExpirationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAuthTokenExpirationResponse) ProtoMessage() {}

func (x *SetAuthTokenExpirationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAuthTokenExpirationResponse.ProtoReflect.Descriptor instead.
func (*SetAuthTokenExpirationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{9}
}

// Request message for the SetAuthTokenExpirationByUserIdRequest rpc.
type SetAuthTokenExpirationByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to reset expiration for that belongs to the given user.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// User token belongs to.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SetAuthTokenExpirationByUserIdRequest) Reset() {
	*x = SetAuthTokenExpirationByUserIdRequest{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAuthTokenExpirationByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAuthTokenExpirationByUserIdRequest) ProtoMessage() {}

func (x *SetAuthTokenExpirationByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAuthTokenExpirationByUserIdRequest.ProtoReflect.Descriptor instead.
func (*SetAuthTokenExpirationByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{10}
}

func (x *SetAuthTokenExpirationByUserIdRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SetAuthTokenExpirationByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Response message for the SetAuthTokenExpirationByUserIdResponse rpc.
type SetAuthTokenExpirationByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetAuthTokenExpirationByUserIdResponse) Reset() {
	*x = SetAuthTokenExpirationByUserIdResponse{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAuthTokenExpirationByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAuthTokenExpirationByUserIdResponse) ProtoMessage() {}

func (x *SetAuthTokenExpirationByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAuthTokenExpirationByUserIdResponse.ProtoReflect.Descriptor instead.
func (*SetAuthTokenExpirationByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{11}
}

// Request message for the DeleteAuthTokenRequest rpc.
type DeleteAuthTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to delete for current user.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DeleteAuthTokenRequest) Reset() {
	*x = DeleteAuthTokenRequest{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthTokenRequest) ProtoMessage() {}

func (x *DeleteAuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthTokenRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteAuthTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Response message for the DeleteAuthTokenResponse rpc.
type DeleteAuthTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAuthTokenResponse) Reset() {
	*x = DeleteAuthTokenResponse{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAuthTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthTokenResponse) ProtoMessage() {}

func (x *DeleteAuthTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthTokenResponse.ProtoReflect.Descriptor instead.
func (*DeleteAuthTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{13}
}

// Request message for the DeleteAuthTokenByUserIdRequest rpc.
type DeleteAuthTokenByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token to delete.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// User token belongs to.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteAuthTokenByUserIdRequest) Reset() {
	*x = DeleteAuthTokenByUserIdRequest{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAuthTokenByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthTokenByUserIdRequest) ProtoMessage() {}

func (x *DeleteAuthTokenByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthTokenByUserIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthTokenByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{14}
}

func (x *DeleteAuthTokenByUserIdRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteAuthTokenByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Response message for the DeleteAuthTokenByUserIdResponse rpc.
type DeleteAuthTokenByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAuthTokenByUserIdResponse) Reset() {
	*x = DeleteAuthTokenByUserIdResponse{}
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAuthTokenByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthTokenByUserIdResponse) ProtoMessage() {}

func (x *DeleteAuthTokenByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_users_auth_token_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthTokenByUserIdResponse.ProtoReflect.Descriptor instead.
func (*DeleteAuthTokenByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP(), []int{15}
}

var File_api_v1alpha1_org_users_auth_token_proto protoreflect.FileDescriptor

var file_api_v1alpha1_org_users_auth_token_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x1a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x54, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x39, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5c,
	0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x17, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x55, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x38, 0x0a, 0x1d,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x35, 0x0a, 0x1d, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x20, 0x0a, 0x1e,
	0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56,
	0x0a, 0x25, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x26, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2e, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x0a, 0x1e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x1f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0xdb, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x0e,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0xa2, 0x02, 0x04, 0x41, 0x56, 0x4f, 0x55, 0xaa, 0x02, 0x16, 0x41, 0x70, 0x69,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x73, 0xca, 0x02, 0x16, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x73, 0xe2, 0x02, 0x22, 0x41,
	0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x4f, 0x72, 0x67, 0x5c,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_org_users_auth_token_proto_rawDescOnce sync.Once
	file_api_v1alpha1_org_users_auth_token_proto_rawDescData = file_api_v1alpha1_org_users_auth_token_proto_rawDesc
)

func file_api_v1alpha1_org_users_auth_token_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_org_users_auth_token_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_org_users_auth_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_org_users_auth_token_proto_rawDescData)
	})
	return file_api_v1alpha1_org_users_auth_token_proto_rawDescData
}

var file_api_v1alpha1_org_users_auth_token_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_api_v1alpha1_org_users_auth_token_proto_goTypes = []any{
	(*CreateAuthTokenRequest)(nil),                 // 0: api.v1alpha1.org.users.CreateAuthTokenRequest
	(*CreateAuthTokenResponse)(nil),                // 1: api.v1alpha1.org.users.CreateAuthTokenResponse
	(*CreateAuthTokenByUserIdRequest)(nil),         // 2: api.v1alpha1.org.users.CreateAuthTokenByUserIdRequest
	(*CreateAuthTokenByUserIdResponse)(nil),        // 3: api.v1alpha1.org.users.CreateAuthTokenByUserIdResponse
	(*ListAuthTokensRequest)(nil),                  // 4: api.v1alpha1.org.users.ListAuthTokensRequest
	(*ListAuthTokensResponse)(nil),                 // 5: api.v1alpha1.org.users.ListAuthTokensResponse
	(*ListAuthTokensByUserIdRequest)(nil),          // 6: api.v1alpha1.org.users.ListAuthTokensByUserIdRequest
	(*ListAuthTokensByUserIdResponse)(nil),         // 7: api.v1alpha1.org.users.ListAuthTokensByUserIdResponse
	(*SetAuthTokenExpirationRequest)(nil),          // 8: api.v1alpha1.org.users.SetAuthTokenExpirationRequest
	(*SetAuthTokenExpirationResponse)(nil),         // 9: api.v1alpha1.org.users.SetAuthTokenExpirationResponse
	(*SetAuthTokenExpirationByUserIdRequest)(nil),  // 10: api.v1alpha1.org.users.SetAuthTokenExpirationByUserIdRequest
	(*SetAuthTokenExpirationByUserIdResponse)(nil), // 11: api.v1alpha1.org.users.SetAuthTokenExpirationByUserIdResponse
	(*DeleteAuthTokenRequest)(nil),                 // 12: api.v1alpha1.org.users.DeleteAuthTokenRequest
	(*DeleteAuthTokenResponse)(nil),                // 13: api.v1alpha1.org.users.DeleteAuthTokenResponse
	(*DeleteAuthTokenByUserIdRequest)(nil),         // 14: api.v1alpha1.org.users.DeleteAuthTokenByUserIdRequest
	(*DeleteAuthTokenByUserIdResponse)(nil),        // 15: api.v1alpha1.org.users.DeleteAuthTokenByUserIdResponse
	(*org.AuthToken)(nil),                          // 16: api.commons.org.AuthToken
}
var file_api_v1alpha1_org_users_auth_token_proto_depIdxs = []int32{
	16, // 0: api.v1alpha1.org.users.CreateAuthTokenResponse.auth_token:type_name -> api.commons.org.AuthToken
	16, // 1: api.v1alpha1.org.users.CreateAuthTokenByUserIdResponse.auth_token:type_name -> api.commons.org.AuthToken
	16, // 2: api.v1alpha1.org.users.ListAuthTokensResponse.auth_tokens:type_name -> api.commons.org.AuthToken
	16, // 3: api.v1alpha1.org.users.ListAuthTokensByUserIdResponse.auth_tokens:type_name -> api.commons.org.AuthToken
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_org_users_auth_token_proto_init() }
func file_api_v1alpha1_org_users_auth_token_proto_init() {
	if File_api_v1alpha1_org_users_auth_token_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_org_users_auth_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_org_users_auth_token_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_org_users_auth_token_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_org_users_auth_token_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_org_users_auth_token_proto = out.File
	file_api_v1alpha1_org_users_auth_token_proto_rawDesc = nil
	file_api_v1alpha1_org_users_auth_token_proto_goTypes = nil
	file_api_v1alpha1_org_users_auth_token_proto_depIdxs = nil
}
