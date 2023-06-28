// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/commons/org/agent_profile_group.proto

package org

import (
	commons "github.com/tcncloud/api-go/api/commons"
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

type AgentProfileGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrgId          string           `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Name           string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	PriorityGroups []*PriorityGroup `protobuf:"bytes,4,rep,name=priority_groups,json=priorityGroups,proto3" json:"priority_groups,omitempty"`
	// last time agent profile group was updated
	LastUpdated  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	DefaultGroup bool                   `protobuf:"varint,10,opt,name=default_group,json=defaultGroup,proto3" json:"default_group,omitempty"`
}

func (x *AgentProfileGroup) Reset() {
	*x = AgentProfileGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_org_agent_profile_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentProfileGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentProfileGroup) ProtoMessage() {}

func (x *AgentProfileGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_agent_profile_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentProfileGroup.ProtoReflect.Descriptor instead.
func (*AgentProfileGroup) Descriptor() ([]byte, []int) {
	return file_api_commons_org_agent_profile_group_proto_rawDescGZIP(), []int{0}
}

func (x *AgentProfileGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AgentProfileGroup) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *AgentProfileGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AgentProfileGroup) GetPriorityGroups() []*PriorityGroup {
	if x != nil {
		return x.PriorityGroups
	}
	return nil
}

func (x *AgentProfileGroup) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *AgentProfileGroup) GetDefaultGroup() bool {
	if x != nil {
		return x.DefaultGroup
	}
	return false
}

// PriorityGroup is a type and threshold in array for AgentProfileGroup
type PriorityGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// threshold for the priority group
	Threshold int32 `protobuf:"varint,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	// channel type for the priority group
	ChannelType commons.ChannelType `protobuf:"varint,2,opt,name=channel_type,json=channelType,proto3,enum=api.commons.ChannelType" json:"channel_type,omitempty"`
}

func (x *PriorityGroup) Reset() {
	*x = PriorityGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_org_agent_profile_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriorityGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriorityGroup) ProtoMessage() {}

func (x *PriorityGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_agent_profile_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriorityGroup.ProtoReflect.Descriptor instead.
func (*PriorityGroup) Descriptor() ([]byte, []int) {
	return file_api_commons_org_agent_profile_group_proto_rawDescGZIP(), []int{1}
}

func (x *PriorityGroup) GetThreshold() int32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *PriorityGroup) GetChannelType() commons.ChannelType {
	if x != nil {
		return x.ChannelType
	}
	return commons.ChannelType(0)
}

var File_api_commons_org_agent_profile_group_proto protoreflect.FileDescriptor

var file_api_commons_org_agent_profile_group_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x1a, 0x1d, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x01, 0x0a,
	0x11, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a,
	0x0f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x6a, 0x0a, 0x0d, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x3b, 0x0a, 0x0c, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x42, 0xb7, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x42, 0x16,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x6f, 0x72, 0x67, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x4f, 0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0xca, 0x02, 0x0f, 0x41,
	0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0xe2, 0x02,
	0x1b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x4f, 0x72, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_org_agent_profile_group_proto_rawDescOnce sync.Once
	file_api_commons_org_agent_profile_group_proto_rawDescData = file_api_commons_org_agent_profile_group_proto_rawDesc
)

func file_api_commons_org_agent_profile_group_proto_rawDescGZIP() []byte {
	file_api_commons_org_agent_profile_group_proto_rawDescOnce.Do(func() {
		file_api_commons_org_agent_profile_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_org_agent_profile_group_proto_rawDescData)
	})
	return file_api_commons_org_agent_profile_group_proto_rawDescData
}

var file_api_commons_org_agent_profile_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_org_agent_profile_group_proto_goTypes = []interface{}{
	(*AgentProfileGroup)(nil),     // 0: api.commons.org.AgentProfileGroup
	(*PriorityGroup)(nil),         // 1: api.commons.org.PriorityGroup
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(commons.ChannelType)(0),      // 3: api.commons.ChannelType
}
var file_api_commons_org_agent_profile_group_proto_depIdxs = []int32{
	1, // 0: api.commons.org.AgentProfileGroup.priority_groups:type_name -> api.commons.org.PriorityGroup
	2, // 1: api.commons.org.AgentProfileGroup.last_updated:type_name -> google.protobuf.Timestamp
	3, // 2: api.commons.org.PriorityGroup.channel_type:type_name -> api.commons.ChannelType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_commons_org_agent_profile_group_proto_init() }
func file_api_commons_org_agent_profile_group_proto_init() {
	if File_api_commons_org_agent_profile_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commons_org_agent_profile_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentProfileGroup); i {
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
		file_api_commons_org_agent_profile_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriorityGroup); i {
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
			RawDescriptor: file_api_commons_org_agent_profile_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_org_agent_profile_group_proto_goTypes,
		DependencyIndexes: file_api_commons_org_agent_profile_group_proto_depIdxs,
		MessageInfos:      file_api_commons_org_agent_profile_group_proto_msgTypes,
	}.Build()
	File_api_commons_org_agent_profile_group_proto = out.File
	file_api_commons_org_agent_profile_group_proto_rawDesc = nil
	file_api_commons_org_agent_profile_group_proto_goTypes = nil
	file_api_commons_org_agent_profile_group_proto_depIdxs = nil
}
