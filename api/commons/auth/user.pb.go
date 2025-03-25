// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/commons/auth/user.proto

package auth

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

const file_api_commons_auth_user_proto_rawDesc = "" +
	"\n" +
	"\x1bapi/commons/auth/user.proto\x12\x10api.commons.auth\"\xd4\x02\n" +
	"\n" +
	"AuthClaims\x12\"\n" +
	"\rauth0_user_id\x18\x01 \x01(\tR\vauth0UserId\x12\x1e\n" +
	"\vorg_user_id\x18\x02 \x01(\tR\torgUserId\x12\x15\n" +
	"\x06org_id\x18\x03 \x01(\tR\x05orgId\x12\x17\n" +
	"\aapi_key\x18\x04 \x01(\tR\x06apiKey\x12\x1b\n" +
	"\tregion_id\x18\x05 \x01(\tR\bregionId\x12\x12\n" +
	"\x04name\x18\x06 \x01(\tR\x04name\x12 \n" +
	"\vimpersonate\x18\a \x01(\tR\vimpersonate\x12\x1e\n" +
	"\n" +
	"client_sid\x18\xe8\a \x01(\x03R\tclientSid\x12\x1c\n" +
	"\tagent_sid\x18\xe9\a \x01(\x03R\bagentSid\x12\x1c\n" +
	"\tlogin_sid\x18\xea\a \x01(\x03R\bloginSid\x12#\n" +
	"\ractive_org_id\x18\xcc\b \x01(\tR\vactiveOrgId\"I\n" +
	"\x11AuthenticatedUser\x124\n" +
	"\x06claims\x18\x01 \x01(\v2\x1c.api.commons.auth.AuthClaimsR\x06claimsB\xb0\x01\n" +
	"\x14com.api.commons.authB\tUserProtoP\x01Z+github.com/tcncloud/api-go/api/commons/auth\xa2\x02\x03ACA\xaa\x02\x10Api.Commons.Auth\xca\x02\x10Api\\Commons\\Auth\xe2\x02\x1cApi\\Commons\\Auth\\GPBMetadata\xea\x02\x12Api::Commons::Authb\x06proto3"

var (
	file_api_commons_auth_user_proto_rawDescOnce sync.Once
	file_api_commons_auth_user_proto_rawDescData []byte
)

func file_api_commons_auth_user_proto_rawDescGZIP() []byte {
	file_api_commons_auth_user_proto_rawDescOnce.Do(func() {
		file_api_commons_auth_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_auth_user_proto_rawDesc), len(file_api_commons_auth_user_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_auth_user_proto_rawDesc), len(file_api_commons_auth_user_proto_rawDesc)),
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
	file_api_commons_auth_user_proto_goTypes = nil
	file_api_commons_auth_user_proto_depIdxs = nil
}
