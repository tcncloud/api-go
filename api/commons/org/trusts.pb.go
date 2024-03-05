// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/commons/org/trusts.proto

package org

import (
	auth "github.com/tcncloud/api-go/api/commons/auth"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status represents the states that trusts can be in.
type Status int32

const (
	Status_PENDING  Status = 0
	Status_REJECTED Status = 1
	Status_ACCEPTED Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "PENDING",
		1: "REJECTED",
		2: "ACCEPTED",
	}
	Status_value = map[string]int32{
		"PENDING":  0,
		"REJECTED": 1,
		"ACCEPTED": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_org_trusts_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_api_commons_org_trusts_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_org_trusts_proto_rawDescGZIP(), []int{0}
}

// Trust is the entity message for trusts.
type Trust struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the trust.
	TrustId string `protobuf:"bytes,1,opt,name=trust_id,json=trustId,proto3" json:"trust_id,omitempty"`
	// Org ID of the org giving out the trust.
	Grantor string `protobuf:"bytes,2,opt,name=grantor,proto3" json:"grantor,omitempty"`
	// Org ID of the org receiving the trust.
	Grantee string `protobuf:"bytes,3,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// Name of the trust.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the trust.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Permissions the grantor allows users within the
	// grantee org to have when assigned.
	Permissions []auth.Permission `protobuf:"varint,6,rep,packed,name=permissions,proto3,enum=api.commons.auth.Permission" json:"permissions,omitempty"`
	// Labels associated with the trusted permissions.
	Labels []*Label `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty"`
	// Status of the Trust.
	Status  Status `protobuf:"varint,8,opt,name=status,proto3,enum=api.commons.org.Status" json:"status,omitempty"`
	Deleted bool   `protobuf:"varint,9,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Trust) Reset() {
	*x = Trust{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_org_trusts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trust) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trust) ProtoMessage() {}

func (x *Trust) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_trusts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trust.ProtoReflect.Descriptor instead.
func (*Trust) Descriptor() ([]byte, []int) {
	return file_api_commons_org_trusts_proto_rawDescGZIP(), []int{0}
}

func (x *Trust) GetTrustId() string {
	if x != nil {
		return x.TrustId
	}
	return ""
}

func (x *Trust) GetGrantor() string {
	if x != nil {
		return x.Grantor
	}
	return ""
}

func (x *Trust) GetGrantee() string {
	if x != nil {
		return x.Grantee
	}
	return ""
}

func (x *Trust) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Trust) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Trust) GetPermissions() []auth.Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Trust) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Trust) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_PENDING
}

func (x *Trust) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

// TrustGroup contains ALL trusted permission/label groups
// assigned to a user.
type TrustGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Org ID of the org the trusts are from.
	Grantor string `protobuf:"bytes,1,opt,name=grantor,proto3" json:"grantor,omitempty"`
	// All perm/label groups the user has from the grantor org.
	// The length of the labeled_permissions list is equal
	// to the number of assigned trusts the user has from
	// the grantor org.
	LabeledPermissions []*TrustGroup_LabeledPermissions `protobuf:"bytes,2,rep,name=labeled_permissions,json=labeledPermissions,proto3" json:"labeled_permissions,omitempty"`
}

func (x *TrustGroup) Reset() {
	*x = TrustGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_org_trusts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrustGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustGroup) ProtoMessage() {}

func (x *TrustGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_trusts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustGroup.ProtoReflect.Descriptor instead.
func (*TrustGroup) Descriptor() ([]byte, []int) {
	return file_api_commons_org_trusts_proto_rawDescGZIP(), []int{1}
}

func (x *TrustGroup) GetGrantor() string {
	if x != nil {
		return x.Grantor
	}
	return ""
}

func (x *TrustGroup) GetLabeledPermissions() []*TrustGroup_LabeledPermissions {
	if x != nil {
		return x.LabeledPermissions
	}
	return nil
}

// TrustFilter is used to filter trusts during list RPCs.
type TrustFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only return trusts with grantor org id matching this value.
	// Nil will not filter based on grantor.
	Grantor *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=grantor,proto3" json:"grantor,omitempty"`
	// Only return trusts with grantee org id matching this value.
	// Nil will not filter based on grantee.
	Grantee *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// Only return trusts matching ANY of the given statuses.bool
	// Nil will not filter based on status.
	StatusFilter *TrustFilter_StatusFilter `protobuf:"bytes,3,opt,name=status_filter,json=statusFilter,proto3" json:"status_filter,omitempty"`
}

func (x *TrustFilter) Reset() {
	*x = TrustFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_org_trusts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrustFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustFilter) ProtoMessage() {}

func (x *TrustFilter) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_trusts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustFilter.ProtoReflect.Descriptor instead.
func (*TrustFilter) Descriptor() ([]byte, []int) {
	return file_api_commons_org_trusts_proto_rawDescGZIP(), []int{2}
}

func (x *TrustFilter) GetGrantor() *wrapperspb.StringValue {
	if x != nil {
		return x.Grantor
	}
	return nil
}

func (x *TrustFilter) GetGrantee() *wrapperspb.StringValue {
	if x != nil {
		return x.Grantee
	}
	return nil
}

func (x *TrustFilter) GetStatusFilter() *TrustFilter_StatusFilter {
	if x != nil {
		return x.StatusFilter
	}
	return nil
}

type TrustGroup_LabeledPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions []auth.Permission `protobuf:"varint,1,rep,packed,name=permissions,proto3,enum=api.commons.auth.Permission" json:"permissions,omitempty"`
	Labels      []*Label          `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *TrustGroup_LabeledPermissions) Reset() {
	*x = TrustGroup_LabeledPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_org_trusts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrustGroup_LabeledPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustGroup_LabeledPermissions) ProtoMessage() {}

func (x *TrustGroup_LabeledPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_trusts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustGroup_LabeledPermissions.ProtoReflect.Descriptor instead.
func (*TrustGroup_LabeledPermissions) Descriptor() ([]byte, []int) {
	return file_api_commons_org_trusts_proto_rawDescGZIP(), []int{1, 0}
}

func (x *TrustGroup_LabeledPermissions) GetPermissions() []auth.Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *TrustGroup_LabeledPermissions) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

type TrustFilter_StatusFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []Status `protobuf:"varint,1,rep,packed,name=values,proto3,enum=api.commons.org.Status" json:"values,omitempty"`
}

func (x *TrustFilter_StatusFilter) Reset() {
	*x = TrustFilter_StatusFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_org_trusts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrustFilter_StatusFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustFilter_StatusFilter) ProtoMessage() {}

func (x *TrustFilter_StatusFilter) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_trusts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustFilter_StatusFilter.ProtoReflect.Descriptor instead.
func (*TrustFilter_StatusFilter) Descriptor() ([]byte, []int) {
	return file_api_commons_org_trusts_proto_rawDescGZIP(), []int{2, 0}
}

func (x *TrustFilter_StatusFilter) GetValues() []Status {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_api_commons_org_trusts_proto protoreflect.FileDescriptor

var file_api_commons_org_trusts_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x1a,
	0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x02, 0x0a, 0x05,
	0x54, 0x72, 0x75, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x75, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x8e, 0x02, 0x0a, 0x0a, 0x54, 0x72, 0x75, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x12, 0x5f,
	0x0a, 0x13, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x54, 0x72,
	0x75, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x12, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x84, 0x01, 0x0a, 0x12, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x8e, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x75, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x12, 0x36,
	0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x54, 0x72, 0x75, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x3f, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2a, 0x31, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x42, 0xac, 0x01, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f,
	0x72, 0x67, 0x42, 0x0b, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0xa2, 0x02, 0x03,
	0x41, 0x43, 0x4f, 0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x4f, 0x72, 0x67, 0xca, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0xe2, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_commons_org_trusts_proto_rawDescOnce sync.Once
	file_api_commons_org_trusts_proto_rawDescData = file_api_commons_org_trusts_proto_rawDesc
)

func file_api_commons_org_trusts_proto_rawDescGZIP() []byte {
	file_api_commons_org_trusts_proto_rawDescOnce.Do(func() {
		file_api_commons_org_trusts_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_org_trusts_proto_rawDescData)
	})
	return file_api_commons_org_trusts_proto_rawDescData
}

var file_api_commons_org_trusts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_org_trusts_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_commons_org_trusts_proto_goTypes = []interface{}{
	(Status)(0),                           // 0: api.commons.org.Status
	(*Trust)(nil),                         // 1: api.commons.org.Trust
	(*TrustGroup)(nil),                    // 2: api.commons.org.TrustGroup
	(*TrustFilter)(nil),                   // 3: api.commons.org.TrustFilter
	(*TrustGroup_LabeledPermissions)(nil), // 4: api.commons.org.TrustGroup.LabeledPermissions
	(*TrustFilter_StatusFilter)(nil),      // 5: api.commons.org.TrustFilter.StatusFilter
	(auth.Permission)(0),                  // 6: api.commons.auth.Permission
	(*Label)(nil),                         // 7: api.commons.org.Label
	(*wrapperspb.StringValue)(nil),        // 8: google.protobuf.StringValue
}
var file_api_commons_org_trusts_proto_depIdxs = []int32{
	6,  // 0: api.commons.org.Trust.permissions:type_name -> api.commons.auth.Permission
	7,  // 1: api.commons.org.Trust.labels:type_name -> api.commons.org.Label
	0,  // 2: api.commons.org.Trust.status:type_name -> api.commons.org.Status
	4,  // 3: api.commons.org.TrustGroup.labeled_permissions:type_name -> api.commons.org.TrustGroup.LabeledPermissions
	8,  // 4: api.commons.org.TrustFilter.grantor:type_name -> google.protobuf.StringValue
	8,  // 5: api.commons.org.TrustFilter.grantee:type_name -> google.protobuf.StringValue
	5,  // 6: api.commons.org.TrustFilter.status_filter:type_name -> api.commons.org.TrustFilter.StatusFilter
	6,  // 7: api.commons.org.TrustGroup.LabeledPermissions.permissions:type_name -> api.commons.auth.Permission
	7,  // 8: api.commons.org.TrustGroup.LabeledPermissions.labels:type_name -> api.commons.org.Label
	0,  // 9: api.commons.org.TrustFilter.StatusFilter.values:type_name -> api.commons.org.Status
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_commons_org_trusts_proto_init() }
func file_api_commons_org_trusts_proto_init() {
	if File_api_commons_org_trusts_proto != nil {
		return
	}
	file_api_commons_org_labels_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_commons_org_trusts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trust); i {
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
		file_api_commons_org_trusts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrustGroup); i {
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
		file_api_commons_org_trusts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrustFilter); i {
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
		file_api_commons_org_trusts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrustGroup_LabeledPermissions); i {
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
		file_api_commons_org_trusts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrustFilter_StatusFilter); i {
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
			RawDescriptor: file_api_commons_org_trusts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_org_trusts_proto_goTypes,
		DependencyIndexes: file_api_commons_org_trusts_proto_depIdxs,
		EnumInfos:         file_api_commons_org_trusts_proto_enumTypes,
		MessageInfos:      file_api_commons_org_trusts_proto_msgTypes,
	}.Build()
	File_api_commons_org_trusts_proto = out.File
	file_api_commons_org_trusts_proto_rawDesc = nil
	file_api_commons_org_trusts_proto_goTypes = nil
	file_api_commons_org_trusts_proto_depIdxs = nil
}
