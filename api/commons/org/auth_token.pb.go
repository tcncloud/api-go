// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: api/commons/org/auth_token.proto

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

// AuthToken is an authentication token for a user
type AuthToken struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the raw token value
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// user the token belongs to
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// org the user and token belong to
	OrgId string `protobuf:"bytes,3,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// tokens expiration
	Expiration    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	mi := &file_api_commons_org_auth_token_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_auth_token_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_api_commons_org_auth_token_proto_rawDescGZIP(), []int{0}
}

func (x *AuthToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthToken) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthToken) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *AuthToken) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

// AuthTokenExpiration is used to notify the user about an auth token expiration
//
// Deprecated: Marked as deprecated in api/commons/org/auth_token.proto.
type AuthTokenExpiration struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The token that is expiring
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// The token expiration
	Expiration    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthTokenExpiration) Reset() {
	*x = AuthTokenExpiration{}
	mi := &file_api_commons_org_auth_token_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthTokenExpiration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthTokenExpiration) ProtoMessage() {}

func (x *AuthTokenExpiration) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_auth_token_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthTokenExpiration.ProtoReflect.Descriptor instead.
func (*AuthTokenExpiration) Descriptor() ([]byte, []int) {
	return file_api_commons_org_auth_token_proto_rawDescGZIP(), []int{1}
}

func (x *AuthTokenExpiration) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthTokenExpiration) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

var File_api_commons_org_auth_token_proto protoreflect.FileDescriptor

var file_api_commons_org_auth_token_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x6f, 0x72, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6b, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x02, 0x18,
	0x01, 0x42, 0xaf, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x42, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x4f, 0xaa, 0x02, 0x0f,
	0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0xca,
	0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72,
	0x67, 0xe2, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c,
	0x4f, 0x72, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x11, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a,
	0x4f, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_org_auth_token_proto_rawDescOnce sync.Once
	file_api_commons_org_auth_token_proto_rawDescData = file_api_commons_org_auth_token_proto_rawDesc
)

func file_api_commons_org_auth_token_proto_rawDescGZIP() []byte {
	file_api_commons_org_auth_token_proto_rawDescOnce.Do(func() {
		file_api_commons_org_auth_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_org_auth_token_proto_rawDescData)
	})
	return file_api_commons_org_auth_token_proto_rawDescData
}

var file_api_commons_org_auth_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_org_auth_token_proto_goTypes = []any{
	(*AuthToken)(nil),             // 0: api.commons.org.AuthToken
	(*AuthTokenExpiration)(nil),   // 1: api.commons.org.AuthTokenExpiration
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_api_commons_org_auth_token_proto_depIdxs = []int32{
	2, // 0: api.commons.org.AuthToken.expiration:type_name -> google.protobuf.Timestamp
	2, // 1: api.commons.org.AuthTokenExpiration.expiration:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_commons_org_auth_token_proto_init() }
func file_api_commons_org_auth_token_proto_init() {
	if File_api_commons_org_auth_token_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_org_auth_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_org_auth_token_proto_goTypes,
		DependencyIndexes: file_api_commons_org_auth_token_proto_depIdxs,
		MessageInfos:      file_api_commons_org_auth_token_proto_msgTypes,
	}.Build()
	File_api_commons_org_auth_token_proto = out.File
	file_api_commons_org_auth_token_proto_rawDesc = nil
	file_api_commons_org_auth_token_proto_goTypes = nil
	file_api_commons_org_auth_token_proto_depIdxs = nil
}
