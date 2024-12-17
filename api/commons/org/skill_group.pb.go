// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: api/commons/org/skill_group.proto

package org

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

// SkillGroup represents a skill group entity.
type SkillGroup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the skill group
	SkillGroupId string `protobuf:"bytes,1,opt,name=skill_group_id,json=skillGroupId,proto3" json:"skill_group_id,omitempty"`
	// The ID of the org the skill group belongs to.
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// The name of the skill group.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the skill group.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// The skills and proficiencies for the skill group.
	SkillSets     []*SkillSet `protobuf:"bytes,5,rep,name=skill_sets,json=skillSets,proto3" json:"skill_sets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SkillGroup) Reset() {
	*x = SkillGroup{}
	mi := &file_api_commons_org_skill_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SkillGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillGroup) ProtoMessage() {}

func (x *SkillGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_skill_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillGroup.ProtoReflect.Descriptor instead.
func (*SkillGroup) Descriptor() ([]byte, []int) {
	return file_api_commons_org_skill_group_proto_rawDescGZIP(), []int{0}
}

func (x *SkillGroup) GetSkillGroupId() string {
	if x != nil {
		return x.SkillGroupId
	}
	return ""
}

func (x *SkillGroup) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *SkillGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SkillGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SkillGroup) GetSkillSets() []*SkillSet {
	if x != nil {
		return x.SkillSets
	}
	return nil
}

// A set of a skill identifier and the proficiency in that skill
type SkillSet struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The proficiency of the skill as a percent [1,100].
	Proficiency int32 `protobuf:"varint,2,opt,name=proficiency,proto3" json:"proficiency,omitempty"`
	// The sid of the skill
	SkillSid      int64 `protobuf:"varint,3,opt,name=skill_sid,json=skillSid,proto3" json:"skill_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SkillSet) Reset() {
	*x = SkillSet{}
	mi := &file_api_commons_org_skill_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SkillSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillSet) ProtoMessage() {}

func (x *SkillSet) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_skill_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillSet.ProtoReflect.Descriptor instead.
func (*SkillSet) Descriptor() ([]byte, []int) {
	return file_api_commons_org_skill_group_proto_rawDescGZIP(), []int{1}
}

func (x *SkillSet) GetProficiency() int32 {
	if x != nil {
		return x.Proficiency
	}
	return 0
}

func (x *SkillSet) GetSkillSid() int64 {
	if x != nil {
		return x.SkillSid
	}
	return 0
}

var File_api_commons_org_skill_group_proto protoreflect.FileDescriptor

var file_api_commons_org_skill_group_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x6f, 0x72, 0x67, 0x22, 0xb9, 0x01, 0x0a, 0x0a, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f,
	0x73, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x53, 0x65, 0x74, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x74, 0x73,
	0x22, 0x49, 0x0a, 0x08, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x69, 0x64, 0x42, 0xb0, 0x01, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x6f, 0x72, 0x67, 0x42, 0x0f, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f,
	0x72, 0x67, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x4f, 0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0xca, 0x02, 0x0f, 0x41, 0x70, 0x69,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0xe2, 0x02, 0x1b, 0x41,
	0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_org_skill_group_proto_rawDescOnce sync.Once
	file_api_commons_org_skill_group_proto_rawDescData = file_api_commons_org_skill_group_proto_rawDesc
)

func file_api_commons_org_skill_group_proto_rawDescGZIP() []byte {
	file_api_commons_org_skill_group_proto_rawDescOnce.Do(func() {
		file_api_commons_org_skill_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_org_skill_group_proto_rawDescData)
	})
	return file_api_commons_org_skill_group_proto_rawDescData
}

var file_api_commons_org_skill_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_org_skill_group_proto_goTypes = []any{
	(*SkillGroup)(nil), // 0: api.commons.org.SkillGroup
	(*SkillSet)(nil),   // 1: api.commons.org.SkillSet
}
var file_api_commons_org_skill_group_proto_depIdxs = []int32{
	1, // 0: api.commons.org.SkillGroup.skill_sets:type_name -> api.commons.org.SkillSet
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_commons_org_skill_group_proto_init() }
func file_api_commons_org_skill_group_proto_init() {
	if File_api_commons_org_skill_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_org_skill_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_org_skill_group_proto_goTypes,
		DependencyIndexes: file_api_commons_org_skill_group_proto_depIdxs,
		MessageInfos:      file_api_commons_org_skill_group_proto_msgTypes,
	}.Build()
	File_api_commons_org_skill_group_proto = out.File
	file_api_commons_org_skill_group_proto_rawDesc = nil
	file_api_commons_org_skill_group_proto_goTypes = nil
	file_api_commons_org_skill_group_proto_depIdxs = nil
}
