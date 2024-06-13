// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: services/omnichannel/oauth/v1alpha1/entities.proto

package oauthv1alpha1

import (
	commons "github.com/tcncloud/api-go/api/commons"
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

type GetConnectedInboxOAuthURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthenticationType   commons.ConnectedInboxAuthenticationType `protobuf:"varint,1,opt,name=authentication_type,json=authenticationType,proto3,enum=api.commons.ConnectedInboxAuthenticationType" json:"authentication_type,omitempty"`
	ReturningRedirectUri string                                   `protobuf:"bytes,2,opt,name=returning_redirect_uri,json=returningRedirectUri,proto3" json:"returning_redirect_uri,omitempty"`
}

func (x *GetConnectedInboxOAuthURLRequest) Reset() {
	*x = GetConnectedInboxOAuthURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_omnichannel_oauth_v1alpha1_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectedInboxOAuthURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectedInboxOAuthURLRequest) ProtoMessage() {}

func (x *GetConnectedInboxOAuthURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_oauth_v1alpha1_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectedInboxOAuthURLRequest.ProtoReflect.Descriptor instead.
func (*GetConnectedInboxOAuthURLRequest) Descriptor() ([]byte, []int) {
	return file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDescGZIP(), []int{0}
}

func (x *GetConnectedInboxOAuthURLRequest) GetAuthenticationType() commons.ConnectedInboxAuthenticationType {
	if x != nil {
		return x.AuthenticationType
	}
	return commons.ConnectedInboxAuthenticationType(0)
}

func (x *GetConnectedInboxOAuthURLRequest) GetReturningRedirectUri() string {
	if x != nil {
		return x.ReturningRedirectUri
	}
	return ""
}

type GetConnectedInboxOAuthURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OauthUrl string `protobuf:"bytes,1,opt,name=oauth_url,json=oauthUrl,proto3" json:"oauth_url,omitempty"`
}

func (x *GetConnectedInboxOAuthURLResponse) Reset() {
	*x = GetConnectedInboxOAuthURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_omnichannel_oauth_v1alpha1_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectedInboxOAuthURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectedInboxOAuthURLResponse) ProtoMessage() {}

func (x *GetConnectedInboxOAuthURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_oauth_v1alpha1_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectedInboxOAuthURLResponse.ProtoReflect.Descriptor instead.
func (*GetConnectedInboxOAuthURLResponse) Descriptor() ([]byte, []int) {
	return file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDescGZIP(), []int{1}
}

func (x *GetConnectedInboxOAuthURLResponse) GetOauthUrl() string {
	if x != nil {
		return x.OauthUrl
	}
	return ""
}

var File_services_omnichannel_oauth_v1alpha1_entities_proto protoreflect.FileDescriptor

var file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f,
	0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5e, 0x0a,
	0x13, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a,
	0x16, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x69, 0x22, 0x40, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x55, 0x72, 0x6c, 0x42, 0xb5, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x42, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x53, 0x4f, 0x4f, 0xaa, 0x02, 0x23, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x23, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5c, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xe2, 0x02, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x4f, 0x6d,
	0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5c, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x26, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a,
	0x3a, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x3a, 0x3a, 0x4f, 0x61,
	0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDescOnce sync.Once
	file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDescData = file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDesc
)

func file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDescGZIP() []byte {
	file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDescOnce.Do(func() {
		file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDescData)
	})
	return file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDescData
}

var file_services_omnichannel_oauth_v1alpha1_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_omnichannel_oauth_v1alpha1_entities_proto_goTypes = []any{
	(*GetConnectedInboxOAuthURLRequest)(nil),      // 0: services.omnichannel.oauth.v1alpha1.GetConnectedInboxOAuthURLRequest
	(*GetConnectedInboxOAuthURLResponse)(nil),     // 1: services.omnichannel.oauth.v1alpha1.GetConnectedInboxOAuthURLResponse
	(commons.ConnectedInboxAuthenticationType)(0), // 2: api.commons.ConnectedInboxAuthenticationType
}
var file_services_omnichannel_oauth_v1alpha1_entities_proto_depIdxs = []int32{
	2, // 0: services.omnichannel.oauth.v1alpha1.GetConnectedInboxOAuthURLRequest.authentication_type:type_name -> api.commons.ConnectedInboxAuthenticationType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_omnichannel_oauth_v1alpha1_entities_proto_init() }
func file_services_omnichannel_oauth_v1alpha1_entities_proto_init() {
	if File_services_omnichannel_oauth_v1alpha1_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_omnichannel_oauth_v1alpha1_entities_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetConnectedInboxOAuthURLRequest); i {
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
		file_services_omnichannel_oauth_v1alpha1_entities_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetConnectedInboxOAuthURLResponse); i {
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
			RawDescriptor: file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_omnichannel_oauth_v1alpha1_entities_proto_goTypes,
		DependencyIndexes: file_services_omnichannel_oauth_v1alpha1_entities_proto_depIdxs,
		MessageInfos:      file_services_omnichannel_oauth_v1alpha1_entities_proto_msgTypes,
	}.Build()
	File_services_omnichannel_oauth_v1alpha1_entities_proto = out.File
	file_services_omnichannel_oauth_v1alpha1_entities_proto_rawDesc = nil
	file_services_omnichannel_oauth_v1alpha1_entities_proto_goTypes = nil
	file_services_omnichannel_oauth_v1alpha1_entities_proto_depIdxs = nil
}
