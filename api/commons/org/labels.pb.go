// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: api/commons/org/labels.proto

package org

import (
	commons "github.com/tcncloud/api-go/api/commons"
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

// Label is an entity message.
type Label struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// labels are owned by orgs. this
	// could be nasty in regards to
	// org trusts and groups.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// simply the name of the label.
	// 'Team A', 'Medical', etc
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// the description of the label.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// Hex color code of the label.
	Color string `protobuf:"bytes,7,opt,name=color,proto3" json:"color,omitempty"`
	// Id of the label.
	LabelId string `protobuf:"bytes,8,opt,name=label_id,json=labelId,proto3" json:"label_id,omitempty"`
	// whether or not the label is deleted.
	Deleted       bool `protobuf:"varint,9,opt,name=deleted,proto3" json:"deleted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Label) Reset() {
	*x = Label{}
	mi := &file_api_commons_org_labels_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_labels_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_api_commons_org_labels_proto_rawDescGZIP(), []int{0}
}

func (x *Label) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Label) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Label) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Label) GetLabelId() string {
	if x != nil {
		return x.LabelId
	}
	return ""
}

func (x *Label) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

// Entity message for a label assignment.
type LabelAssignment struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the label.
	LabelId string `protobuf:"bytes,1,opt,name=label_id,json=labelId,proto3" json:"label_id,omitempty"`
	// type of entity being labelled.
	Type commons.EntityType `protobuf:"varint,2,opt,name=type,proto3,enum=api.commons.EntityType" json:"type,omitempty"`
	// id of the entity being labelled.
	EntityId string `protobuf:"bytes,3,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// the id of the organization the label belongs too
	OrgId string `protobuf:"bytes,4,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// The label associated with the label_id.
	// This will only be populated during the GetLabelAssignments RPC
	// if the request field PopulateLabelInfo is set to true.
	Label         *Label `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelAssignment) Reset() {
	*x = LabelAssignment{}
	mi := &file_api_commons_org_labels_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelAssignment) ProtoMessage() {}

func (x *LabelAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_labels_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelAssignment.ProtoReflect.Descriptor instead.
func (*LabelAssignment) Descriptor() ([]byte, []int) {
	return file_api_commons_org_labels_proto_rawDescGZIP(), []int{1}
}

func (x *LabelAssignment) GetLabelId() string {
	if x != nil {
		return x.LabelId
	}
	return ""
}

func (x *LabelAssignment) GetType() commons.EntityType {
	if x != nil {
		return x.Type
	}
	return commons.EntityType(0)
}

func (x *LabelAssignment) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *LabelAssignment) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *LabelAssignment) GetLabel() *Label {
	if x != nil {
		return x.Label
	}
	return nil
}

var File_api_commons_org_labels_proto protoreflect.FileDescriptor

var file_api_commons_org_labels_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x1a,
	0x18, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x05, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4a, 0x04, 0x08, 0x04, 0x10,
	0x05, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x22, 0xbb, 0x01, 0x0a, 0x0f, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0xac, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x42, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x4f, 0xaa, 0x02,
	0x0f, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x72, 0x67,
	0xca, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f,
	0x72, 0x67, 0xe2, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x11, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a,
	0x3a, 0x4f, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_commons_org_labels_proto_rawDescOnce sync.Once
	file_api_commons_org_labels_proto_rawDescData []byte
)

func file_api_commons_org_labels_proto_rawDescGZIP() []byte {
	file_api_commons_org_labels_proto_rawDescOnce.Do(func() {
		file_api_commons_org_labels_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_org_labels_proto_rawDesc), len(file_api_commons_org_labels_proto_rawDesc)))
	})
	return file_api_commons_org_labels_proto_rawDescData
}

var file_api_commons_org_labels_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_commons_org_labels_proto_goTypes = []any{
	(*Label)(nil),           // 0: api.commons.org.Label
	(*LabelAssignment)(nil), // 1: api.commons.org.LabelAssignment
	(commons.EntityType)(0), // 2: api.commons.EntityType
}
var file_api_commons_org_labels_proto_depIdxs = []int32{
	2, // 0: api.commons.org.LabelAssignment.type:type_name -> api.commons.EntityType
	0, // 1: api.commons.org.LabelAssignment.label:type_name -> api.commons.org.Label
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_commons_org_labels_proto_init() }
func file_api_commons_org_labels_proto_init() {
	if File_api_commons_org_labels_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_org_labels_proto_rawDesc), len(file_api_commons_org_labels_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_org_labels_proto_goTypes,
		DependencyIndexes: file_api_commons_org_labels_proto_depIdxs,
		MessageInfos:      file_api_commons_org_labels_proto_msgTypes,
	}.Build()
	File_api_commons_org_labels_proto = out.File
	file_api_commons_org_labels_proto_goTypes = nil
	file_api_commons_org_labels_proto_depIdxs = nil
}
