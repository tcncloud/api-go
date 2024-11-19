// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/commons/org/permissions.proto

package org

import (
	perms "github.com/tcncloud/api-go/annotations/perms"
	commons "github.com/tcncloud/api-go/api/commons"
	auth "github.com/tcncloud/api-go/api/commons/auth"
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

// PermissionGroup defines the permission group entity.
type PermissionGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique id for this permission group.
	PermissionGroupId string `protobuf:"bytes,1,opt,name=permission_group_id,json=permissionGroupId,proto3" json:"permission_group_id,omitempty"`
	// Organization this permission group belongs to.
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Name to identify by.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Brief description of this group of permissions.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// List of permissions in this group.
	Permissions []auth.Permission `protobuf:"varint,5,rep,packed,name=permissions,proto3,enum=api.commons.auth.Permission" json:"permissions,omitempty"`
	// Restricts users from modifying this group.
	ReadOnly bool `protobuf:"varint,6,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	// A list of label ids that are assigned to this group.
	LabelIds []string `protobuf:"bytes,7,rep,name=label_ids,json=labelIds,proto3" json:"label_ids,omitempty"`
}

func (x *PermissionGroup) Reset() {
	*x = PermissionGroup{}
	mi := &file_api_commons_org_permissions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionGroup) ProtoMessage() {}

func (x *PermissionGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_permissions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionGroup.ProtoReflect.Descriptor instead.
func (*PermissionGroup) Descriptor() ([]byte, []int) {
	return file_api_commons_org_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionGroup) GetPermissionGroupId() string {
	if x != nil {
		return x.PermissionGroupId
	}
	return ""
}

func (x *PermissionGroup) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *PermissionGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PermissionGroup) GetPermissions() []auth.Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *PermissionGroup) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

func (x *PermissionGroup) GetLabelIds() []string {
	if x != nil {
		return x.LabelIds
	}
	return nil
}

// P3PermissionGroup entity.
// This was migrated from permissions_group from skunk
// and is requested by backoffice as the source of permissions.
type P3PermissionGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique Id for this entity.
	P3PermissionGroupId string `protobuf:"bytes,1,opt,name=p3_permission_group_id,json=p3PermissionGroupId,proto3" json:"p3_permission_group_id,omitempty"`
	// Owning org.
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Owning org region.
	RegionId string `protobuf:"bytes,3,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// Display name for this group.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// More detailed description for this group.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// List of permissions associated with this group.
	Permissions []commons.Permission `protobuf:"varint,7,rep,packed,name=permissions,proto3,enum=api.commons.Permission" json:"permissions,omitempty"`
}

func (x *P3PermissionGroup) Reset() {
	*x = P3PermissionGroup{}
	mi := &file_api_commons_org_permissions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *P3PermissionGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P3PermissionGroup) ProtoMessage() {}

func (x *P3PermissionGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_permissions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P3PermissionGroup.ProtoReflect.Descriptor instead.
func (*P3PermissionGroup) Descriptor() ([]byte, []int) {
	return file_api_commons_org_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *P3PermissionGroup) GetP3PermissionGroupId() string {
	if x != nil {
		return x.P3PermissionGroupId
	}
	return ""
}

func (x *P3PermissionGroup) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *P3PermissionGroup) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *P3PermissionGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *P3PermissionGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *P3PermissionGroup) GetPermissions() []commons.Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

// License defines the structure of organization license assignments.
type License struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Application containing licenses.
	App perms.Application `protobuf:"varint,1,opt,name=app,proto3,enum=annotations.perms.Application" json:"app,omitempty"`
	// Individual cards for the application.
	Cards []*License_Card `protobuf:"bytes,2,rep,name=cards,proto3" json:"cards,omitempty"`
	// Name of the license (to be displayed in the UI)
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *License) Reset() {
	*x = License{}
	mi := &file_api_commons_org_permissions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *License) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License) ProtoMessage() {}

func (x *License) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_permissions_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License.ProtoReflect.Descriptor instead.
func (*License) Descriptor() ([]byte, []int) {
	return file_api_commons_org_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *License) GetApp() perms.Application {
	if x != nil {
		return x.App
	}
	return perms.Application(0)
}

func (x *License) GetCards() []*License_Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

func (x *License) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type License_Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Card type.
	Type perms.Card `protobuf:"varint,1,opt,name=type,proto3,enum=annotations.perms.Card" json:"type,omitempty"`
	// List of permissions and features.
	Permissions []*License_Permission `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// String version of the card/sub-section.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *License_Card) Reset() {
	*x = License_Card{}
	mi := &file_api_commons_org_permissions_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *License_Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License_Card) ProtoMessage() {}

func (x *License_Card) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_permissions_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License_Card.ProtoReflect.Descriptor instead.
func (*License_Card) Descriptor() ([]byte, []int) {
	return file_api_commons_org_permissions_proto_rawDescGZIP(), []int{2, 0}
}

func (x *License_Card) GetType() perms.Card {
	if x != nil {
		return x.Type
	}
	return perms.Card(0)
}

func (x *License_Card) GetPermissions() []*License_Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *License_Card) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type License_Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permission auth.Permission `protobuf:"varint,1,opt,name=permission,proto3,enum=api.commons.auth.Permission" json:"permission,omitempty"`
	// Whether this permission is enabled in the license.
	Enabled bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// List of features granted by the permission.
	Features []string `protobuf:"bytes,3,rep,name=features,proto3" json:"features,omitempty"`
	// String version of the permission.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *License_Permission) Reset() {
	*x = License_Permission{}
	mi := &file_api_commons_org_permissions_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *License_Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License_Permission) ProtoMessage() {}

func (x *License_Permission) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_permissions_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License_Permission.ProtoReflect.Descriptor instead.
func (*License_Permission) Descriptor() ([]byte, []int) {
	return file_api_commons_org_permissions_proto_rawDescGZIP(), []int{2, 1}
}

func (x *License_Permission) GetPermission() auth.Permission {
	if x != nil {
		return x.Permission
	}
	return auth.Permission(0)
}

func (x *License_Permission) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *License_Permission) GetFeatures() []string {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *License_Permission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_commons_org_permissions_proto protoreflect.FileDescriptor

var file_api_commons_org_permissions_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x6f, 0x72, 0x67, 0x1a, 0x1f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a,
	0x0f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x2e, 0x0a, 0x13, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x22, 0xed, 0x01, 0x0a, 0x11, 0x50, 0x33, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x33, 0x0a,
	0x16, 0x70, 0x33, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70,
	0x33, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xac, 0x03, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x33, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x8e,
	0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a,
	0x94, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c,
	0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xb1, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x42, 0x10,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0xa2, 0x02,
	0x03, 0x41, 0x43, 0x4f, 0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0xca, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0xe2, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_commons_org_permissions_proto_rawDescOnce sync.Once
	file_api_commons_org_permissions_proto_rawDescData = file_api_commons_org_permissions_proto_rawDesc
)

func file_api_commons_org_permissions_proto_rawDescGZIP() []byte {
	file_api_commons_org_permissions_proto_rawDescOnce.Do(func() {
		file_api_commons_org_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_org_permissions_proto_rawDescData)
	})
	return file_api_commons_org_permissions_proto_rawDescData
}

var file_api_commons_org_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_commons_org_permissions_proto_goTypes = []any{
	(*PermissionGroup)(nil),    // 0: api.commons.org.PermissionGroup
	(*P3PermissionGroup)(nil),  // 1: api.commons.org.P3PermissionGroup
	(*License)(nil),            // 2: api.commons.org.License
	(*License_Card)(nil),       // 3: api.commons.org.License.Card
	(*License_Permission)(nil), // 4: api.commons.org.License.Permission
	(auth.Permission)(0),       // 5: api.commons.auth.Permission
	(commons.Permission)(0),    // 6: api.commons.Permission
	(perms.Application)(0),     // 7: annotations.perms.Application
	(perms.Card)(0),            // 8: annotations.perms.Card
}
var file_api_commons_org_permissions_proto_depIdxs = []int32{
	5, // 0: api.commons.org.PermissionGroup.permissions:type_name -> api.commons.auth.Permission
	6, // 1: api.commons.org.P3PermissionGroup.permissions:type_name -> api.commons.Permission
	7, // 2: api.commons.org.License.app:type_name -> annotations.perms.Application
	3, // 3: api.commons.org.License.cards:type_name -> api.commons.org.License.Card
	8, // 4: api.commons.org.License.Card.type:type_name -> annotations.perms.Card
	4, // 5: api.commons.org.License.Card.permissions:type_name -> api.commons.org.License.Permission
	5, // 6: api.commons.org.License.Permission.permission:type_name -> api.commons.auth.Permission
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_commons_org_permissions_proto_init() }
func file_api_commons_org_permissions_proto_init() {
	if File_api_commons_org_permissions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_org_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_org_permissions_proto_goTypes,
		DependencyIndexes: file_api_commons_org_permissions_proto_depIdxs,
		MessageInfos:      file_api_commons_org_permissions_proto_msgTypes,
	}.Build()
	File_api_commons_org_permissions_proto = out.File
	file_api_commons_org_permissions_proto_rawDesc = nil
	file_api_commons_org_permissions_proto_goTypes = nil
	file_api_commons_org_permissions_proto_depIdxs = nil
}
