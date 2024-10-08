// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/commons/contactmanager.proto

package commons

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

type DeDuplicationMergeStrategy int32

const (
	DeDuplicationMergeStrategy_KEEP_EXISTING_LIST    DeDuplicationMergeStrategy = 0
	DeDuplicationMergeStrategy_REPLACE_EXISTING_LIST DeDuplicationMergeStrategy = 1
)

// Enum value maps for DeDuplicationMergeStrategy.
var (
	DeDuplicationMergeStrategy_name = map[int32]string{
		0: "KEEP_EXISTING_LIST",
		1: "REPLACE_EXISTING_LIST",
	}
	DeDuplicationMergeStrategy_value = map[string]int32{
		"KEEP_EXISTING_LIST":    0,
		"REPLACE_EXISTING_LIST": 1,
	}
)

func (x DeDuplicationMergeStrategy) Enum() *DeDuplicationMergeStrategy {
	p := new(DeDuplicationMergeStrategy)
	*p = x
	return p
}

func (x DeDuplicationMergeStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeDuplicationMergeStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_contactmanager_proto_enumTypes[0].Descriptor()
}

func (DeDuplicationMergeStrategy) Type() protoreflect.EnumType {
	return &file_api_commons_contactmanager_proto_enumTypes[0]
}

func (x DeDuplicationMergeStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeDuplicationMergeStrategy.Descriptor instead.
func (DeDuplicationMergeStrategy) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_contactmanager_proto_rawDescGZIP(), []int{0}
}

type ContactListStatus int32

const (
	// NEW
	ContactListStatus_CONTACT_LIST_STATUS_NEW ContactListStatus = 0
)

// Enum value maps for ContactListStatus.
var (
	ContactListStatus_name = map[int32]string{
		0: "CONTACT_LIST_STATUS_NEW",
	}
	ContactListStatus_value = map[string]int32{
		"CONTACT_LIST_STATUS_NEW": 0,
	}
)

func (x ContactListStatus) Enum() *ContactListStatus {
	p := new(ContactListStatus)
	*p = x
	return p
}

func (x ContactListStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContactListStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_contactmanager_proto_enumTypes[1].Descriptor()
}

func (ContactListStatus) Type() protoreflect.EnumType {
	return &file_api_commons_contactmanager_proto_enumTypes[1]
}

func (x ContactListStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContactListStatus.Descriptor instead.
func (ContactListStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_contactmanager_proto_rawDescGZIP(), []int{1}
}

type ContactEntryStatus int32

const (
	ContactEntryStatus_CONTACT_ENTRY_STATUS_NEW     ContactEntryStatus = 0 //NEW
	ContactEntryStatus_CONTACT_ENTRY_STATUS_INUSE   ContactEntryStatus = 1 //IN USE
	ContactEntryStatus_CONTACT_ENTRY_STATUS_NOTREF  ContactEntryStatus = 2 //NOT REF
	ContactEntryStatus_CONTACT_ENTRY_STATUS_DONE    ContactEntryStatus = 3 //DONE
	ContactEntryStatus_CONTACT_ENTRY_STATUS_EXPIRED ContactEntryStatus = 4 //EXPIRED
)

// Enum value maps for ContactEntryStatus.
var (
	ContactEntryStatus_name = map[int32]string{
		0: "CONTACT_ENTRY_STATUS_NEW",
		1: "CONTACT_ENTRY_STATUS_INUSE",
		2: "CONTACT_ENTRY_STATUS_NOTREF",
		3: "CONTACT_ENTRY_STATUS_DONE",
		4: "CONTACT_ENTRY_STATUS_EXPIRED",
	}
	ContactEntryStatus_value = map[string]int32{
		"CONTACT_ENTRY_STATUS_NEW":     0,
		"CONTACT_ENTRY_STATUS_INUSE":   1,
		"CONTACT_ENTRY_STATUS_NOTREF":  2,
		"CONTACT_ENTRY_STATUS_DONE":    3,
		"CONTACT_ENTRY_STATUS_EXPIRED": 4,
	}
)

func (x ContactEntryStatus) Enum() *ContactEntryStatus {
	p := new(ContactEntryStatus)
	*p = x
	return p
}

func (x ContactEntryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContactEntryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_contactmanager_proto_enumTypes[2].Descriptor()
}

func (ContactEntryStatus) Type() protoreflect.EnumType {
	return &file_api_commons_contactmanager_proto_enumTypes[2]
}

func (x ContactEntryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContactEntryStatus.Descriptor instead.
func (ContactEntryStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_contactmanager_proto_rawDescGZIP(), []int{2}
}

// Deprecated: Marked as deprecated in api/commons/contactmanager.proto.
type ContactManagerEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactManagerEntryId     int64                  `protobuf:"varint,1,opt,name=contact_manager_entry_id,json=contactManagerEntryId,proto3" json:"contact_manager_entry_id,omitempty"`
	ContactManagerEntryListId int64                  `protobuf:"varint,2,opt,name=contact_manager_entry_list_id,json=contactManagerEntryListId,proto3" json:"contact_manager_entry_list_id,omitempty"`
	Key                       string                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                     string                 `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Type                      string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	DateCreated               *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
}

func (x *ContactManagerEntry) Reset() {
	*x = ContactManagerEntry{}
	mi := &file_api_commons_contactmanager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContactManagerEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactManagerEntry) ProtoMessage() {}

func (x *ContactManagerEntry) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_contactmanager_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactManagerEntry.ProtoReflect.Descriptor instead.
func (*ContactManagerEntry) Descriptor() ([]byte, []int) {
	return file_api_commons_contactmanager_proto_rawDescGZIP(), []int{0}
}

func (x *ContactManagerEntry) GetContactManagerEntryId() int64 {
	if x != nil {
		return x.ContactManagerEntryId
	}
	return 0
}

func (x *ContactManagerEntry) GetContactManagerEntryListId() int64 {
	if x != nil {
		return x.ContactManagerEntryListId
	}
	return 0
}

func (x *ContactManagerEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ContactManagerEntry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ContactManagerEntry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ContactManagerEntry) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

// Deprecated: Marked as deprecated in api/commons/contactmanager.proto.
type ContactManagerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactManagerListId int64  `protobuf:"varint,1,opt,name=contact_manager_list_id,json=contactManagerListId,proto3" json:"contact_manager_list_id,omitempty"`
	OrgId                string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	ProjectId            int64  `protobuf:"varint,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	FileName             string `protobuf:"bytes,4,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Description          string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// List of Columns In The Contact List. UI to render table based on this list
	ListDetails []string               `protobuf:"bytes,6,rep,name=list_details,json=listDetails,proto3" json:"list_details,omitempty"`
	Ttl         int64                  `protobuf:"varint,7,opt,name=ttl,proto3" json:"ttl,omitempty"`
	DateCreated *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
}

func (x *ContactManagerList) Reset() {
	*x = ContactManagerList{}
	mi := &file_api_commons_contactmanager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContactManagerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactManagerList) ProtoMessage() {}

func (x *ContactManagerList) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_contactmanager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactManagerList.ProtoReflect.Descriptor instead.
func (*ContactManagerList) Descriptor() ([]byte, []int) {
	return file_api_commons_contactmanager_proto_rawDescGZIP(), []int{1}
}

func (x *ContactManagerList) GetContactManagerListId() int64 {
	if x != nil {
		return x.ContactManagerListId
	}
	return 0
}

func (x *ContactManagerList) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *ContactManagerList) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *ContactManagerList) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ContactManagerList) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ContactManagerList) GetListDetails() []string {
	if x != nil {
		return x.ListDetails
	}
	return nil
}

func (x *ContactManagerList) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *ContactManagerList) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

// Deprecated: Marked as deprecated in api/commons/contactmanager.proto.
type ContactManagerEntryVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ContactManagerEntryVal) Reset() {
	*x = ContactManagerEntryVal{}
	mi := &file_api_commons_contactmanager_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContactManagerEntryVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactManagerEntryVal) ProtoMessage() {}

func (x *ContactManagerEntryVal) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_contactmanager_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactManagerEntryVal.ProtoReflect.Descriptor instead.
func (*ContactManagerEntryVal) Descriptor() ([]byte, []int) {
	return file_api_commons_contactmanager_proto_rawDescGZIP(), []int{2}
}

func (x *ContactManagerEntryVal) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ContactManagerEntryVal) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_api_commons_contactmanager_proto protoreflect.FileDescriptor

var file_api_commons_contactmanager_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x97, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x15,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x1d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x22, 0xc4, 0x02, 0x0a, 0x12, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x39, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x3d,
	0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x02, 0x18,
	0x01, 0x22, 0x46, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x18, 0x01, 0x2a, 0x4f, 0x0a, 0x1a, 0x44, 0x65, 0x44,
	0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x16, 0x0a, 0x12, 0x4b, 0x45, 0x45, 0x50, 0x5f,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x00, 0x12,
	0x19, 0x0a, 0x15, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54,
	0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x01, 0x2a, 0x30, 0x0a, 0x11, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x00, 0x2a, 0xb4, 0x01, 0x0a,
	0x12, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x45,
	0x4e, 0x54, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x10,
	0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x45, 0x4e, 0x54,
	0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x55, 0x53, 0x45, 0x10,
	0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x45, 0x4e, 0x54,
	0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x52, 0x45, 0x46,
	0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x45, 0x4e,
	0x54, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10,
	0x03, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x45, 0x4e, 0x54,
	0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45,
	0x44, 0x10, 0x04, 0x42, 0x9b, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_contactmanager_proto_rawDescOnce sync.Once
	file_api_commons_contactmanager_proto_rawDescData = file_api_commons_contactmanager_proto_rawDesc
)

func file_api_commons_contactmanager_proto_rawDescGZIP() []byte {
	file_api_commons_contactmanager_proto_rawDescOnce.Do(func() {
		file_api_commons_contactmanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_contactmanager_proto_rawDescData)
	})
	return file_api_commons_contactmanager_proto_rawDescData
}

var file_api_commons_contactmanager_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_commons_contactmanager_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_commons_contactmanager_proto_goTypes = []any{
	(DeDuplicationMergeStrategy)(0), // 0: api.commons.DeDuplicationMergeStrategy
	(ContactListStatus)(0),          // 1: api.commons.ContactListStatus
	(ContactEntryStatus)(0),         // 2: api.commons.ContactEntryStatus
	(*ContactManagerEntry)(nil),     // 3: api.commons.ContactManagerEntry
	(*ContactManagerList)(nil),      // 4: api.commons.ContactManagerList
	(*ContactManagerEntryVal)(nil),  // 5: api.commons.ContactManagerEntryVal
	(*timestamppb.Timestamp)(nil),   // 6: google.protobuf.Timestamp
}
var file_api_commons_contactmanager_proto_depIdxs = []int32{
	6, // 0: api.commons.ContactManagerEntry.date_created:type_name -> google.protobuf.Timestamp
	6, // 1: api.commons.ContactManagerList.date_created:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_commons_contactmanager_proto_init() }
func file_api_commons_contactmanager_proto_init() {
	if File_api_commons_contactmanager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_contactmanager_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_contactmanager_proto_goTypes,
		DependencyIndexes: file_api_commons_contactmanager_proto_depIdxs,
		EnumInfos:         file_api_commons_contactmanager_proto_enumTypes,
		MessageInfos:      file_api_commons_contactmanager_proto_msgTypes,
	}.Build()
	File_api_commons_contactmanager_proto = out.File
	file_api_commons_contactmanager_proto_rawDesc = nil
	file_api_commons_contactmanager_proto_goTypes = nil
	file_api_commons_contactmanager_proto_depIdxs = nil
}
