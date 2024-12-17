// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: api/commons/auth/user.proto

package auth

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

// AuthClaims is a proto mapping of the JWT Claims
type AuthClaims struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// AUTH0 user_id
	Auth0UserId string `protobuf:"bytes,1,opt,name=auth0_user_id,json=auth0UserId,proto3" json:"auth0_user_id,omitempty"`
	// ORG user_id
	OrgUserId string `protobuf:"bytes,2,opt,name=org_user_id,json=orgUserId,proto3" json:"org_user_id,omitempty"`
	// ORG ID
	OrgId string `protobuf:"bytes,3,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// API key used in this request for API-based endpoints
	ApiKey string `protobuf:"bytes,4,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// ORG Region ID
	RegionId string `protobuf:"bytes,5,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// ??
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// Impersonation information
	Impersonate string `protobuf:"bytes,7,opt,name=impersonate,proto3" json:"impersonate,omitempty"`
	// P3 Client SID from "client" table
	ClientSid int64 `protobuf:"varint,1000,opt,name=client_sid,json=clientSid,proto3" json:"client_sid,omitempty"`
	// P3 Agent SID from "agent" table
	AgentSid int64 `protobuf:"varint,1001,opt,name=agent_sid,json=agentSid,proto3" json:"agent_sid,omitempty"`
	// Mapped to agent_sid
	LoginSid int64 `protobuf:"varint,1002,opt,name=login_sid,json=loginSid,proto3" json:"login_sid,omitempty"`
	// ActiveOrgId is the org being used in behalf of for the api call
	ActiveOrgId   string `protobuf:"bytes,1100,opt,name=active_org_id,json=activeOrgId,proto3" json:"active_org_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthClaims) Reset() {
	*x = AuthClaims{}
	mi := &file_api_commons_auth_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthClaims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthClaims) ProtoMessage() {}

func (x *AuthClaims) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_auth_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthClaims.ProtoReflect.Descriptor instead.
func (*AuthClaims) Descriptor() ([]byte, []int) {
	return file_api_commons_auth_user_proto_rawDescGZIP(), []int{0}
}

func (x *AuthClaims) GetAuth0UserId() string {
	if x != nil {
		return x.Auth0UserId
	}
	return ""
}

func (x *AuthClaims) GetOrgUserId() string {
	if x != nil {
		return x.OrgUserId
	}
	return ""
}

func (x *AuthClaims) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *AuthClaims) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *AuthClaims) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *AuthClaims) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthClaims) GetImpersonate() string {
	if x != nil {
		return x.Impersonate
	}
	return ""
}

func (x *AuthClaims) GetClientSid() int64 {
	if x != nil {
		return x.ClientSid
	}
	return 0
}

func (x *AuthClaims) GetAgentSid() int64 {
	if x != nil {
		return x.AgentSid
	}
	return 0
}

func (x *AuthClaims) GetLoginSid() int64 {
	if x != nil {
		return x.LoginSid
	}
	return 0
}

func (x *AuthClaims) GetActiveOrgId() string {
	if x != nil {
		return x.ActiveOrgId
	}
	return ""
}

// AuthenticatedUser is a proto serialized version of AuthUser. It can be used to pass all AuthUser information towards the backends
type AuthenticatedUser struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// JWT Claims
	Claims        *AuthClaims `protobuf:"bytes,1,opt,name=claims,proto3" json:"claims,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticatedUser) Reset() {
	*x = AuthenticatedUser{}
	mi := &file_api_commons_auth_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticatedUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticatedUser) ProtoMessage() {}

func (x *AuthenticatedUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_auth_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticatedUser.ProtoReflect.Descriptor instead.
func (*AuthenticatedUser) Descriptor() ([]byte, []int) {
	return file_api_commons_auth_user_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticatedUser) GetClaims() *AuthClaims {
	if x != nil {
		return x.Claims
	}
	return nil
}

var File_api_commons_auth_user_proto protoreflect.FileDescriptor

var file_api_commons_auth_user_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x22,
	0xd4, 0x02, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x22,
	0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x30, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x30, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x6f, 0x72, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x67, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x69, 0x64, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x69, 0x64, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x53, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x69, 0x64,
	0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x69,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x72, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0xcc, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x42, 0xb0, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x41, 0xaa, 0x02, 0x10, 0x41, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0xca, 0x02, 0x10,
	0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x75, 0x74, 0x68,
	0xe2, 0x02, 0x1c, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41,
	0x75, 0x74, 0x68, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x12, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a,
	0x41, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_auth_user_proto_rawDescOnce sync.Once
	file_api_commons_auth_user_proto_rawDescData = file_api_commons_auth_user_proto_rawDesc
)

func file_api_commons_auth_user_proto_rawDescGZIP() []byte {
	file_api_commons_auth_user_proto_rawDescOnce.Do(func() {
		file_api_commons_auth_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_auth_user_proto_rawDescData)
	})
	return file_api_commons_auth_user_proto_rawDescData
}

var file_api_commons_auth_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_auth_user_proto_goTypes = []any{
	(*AuthClaims)(nil),        // 0: api.commons.auth.AuthClaims
	(*AuthenticatedUser)(nil), // 1: api.commons.auth.AuthenticatedUser
}
var file_api_commons_auth_user_proto_depIdxs = []int32{
	0, // 0: api.commons.auth.AuthenticatedUser.claims:type_name -> api.commons.auth.AuthClaims
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_commons_auth_user_proto_init() }
func file_api_commons_auth_user_proto_init() {
	if File_api_commons_auth_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_auth_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_auth_user_proto_goTypes,
		DependencyIndexes: file_api_commons_auth_user_proto_depIdxs,
		MessageInfos:      file_api_commons_auth_user_proto_msgTypes,
	}.Build()
	File_api_commons_auth_user_proto = out.File
	file_api_commons_auth_user_proto_rawDesc = nil
	file_api_commons_auth_user_proto_goTypes = nil
	file_api_commons_auth_user_proto_depIdxs = nil
}
