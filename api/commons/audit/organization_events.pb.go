// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: api/commons/audit/organization_events.proto

package audit

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

// AccessTokensExpiringEvent is the event that is triggered when any users access tokens are expiring.
type AccessTokensExpiringEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the organization.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// The report of all the users access tokens that are expiring.
	Report        string `protobuf:"bytes,2,opt,name=report,proto3" json:"report,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccessTokensExpiringEvent) Reset() {
	*x = AccessTokensExpiringEvent{}
	mi := &file_api_commons_audit_organization_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessTokensExpiringEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessTokensExpiringEvent) ProtoMessage() {}

func (x *AccessTokensExpiringEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_audit_organization_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessTokensExpiringEvent.ProtoReflect.Descriptor instead.
func (*AccessTokensExpiringEvent) Descriptor() ([]byte, []int) {
	return file_api_commons_audit_organization_events_proto_rawDescGZIP(), []int{0}
}

func (x *AccessTokensExpiringEvent) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *AccessTokensExpiringEvent) GetReport() string {
	if x != nil {
		return x.Report
	}
	return ""
}

var File_api_commons_audit_organization_events_proto protoreflect.FileDescriptor

var file_api_commons_audit_organization_events_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x22, 0x4a, 0x0a, 0x19, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a,
	0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0xc4, 0x01, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0x17, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0xa2,
	0x02, 0x03, 0x41, 0x43, 0x41, 0xaa, 0x02, 0x11, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0xca, 0x02, 0x11, 0x41, 0x70, 0x69, 0x5c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0xe2, 0x02, 0x1d,
	0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_audit_organization_events_proto_rawDescOnce sync.Once
	file_api_commons_audit_organization_events_proto_rawDescData = file_api_commons_audit_organization_events_proto_rawDesc
)

func file_api_commons_audit_organization_events_proto_rawDescGZIP() []byte {
	file_api_commons_audit_organization_events_proto_rawDescOnce.Do(func() {
		file_api_commons_audit_organization_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_audit_organization_events_proto_rawDescData)
	})
	return file_api_commons_audit_organization_events_proto_rawDescData
}

var file_api_commons_audit_organization_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_commons_audit_organization_events_proto_goTypes = []any{
	(*AccessTokensExpiringEvent)(nil), // 0: api.commons.audit.AccessTokensExpiringEvent
}
var file_api_commons_audit_organization_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_audit_organization_events_proto_init() }
func file_api_commons_audit_organization_events_proto_init() {
	if File_api_commons_audit_organization_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_audit_organization_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_audit_organization_events_proto_goTypes,
		DependencyIndexes: file_api_commons_audit_organization_events_proto_depIdxs,
		MessageInfos:      file_api_commons_audit_organization_events_proto_msgTypes,
	}.Build()
	File_api_commons_audit_organization_events_proto = out.File
	file_api_commons_audit_organization_events_proto_rawDesc = nil
	file_api_commons_audit_organization_events_proto_goTypes = nil
	file_api_commons_audit_organization_events_proto_depIdxs = nil
}
