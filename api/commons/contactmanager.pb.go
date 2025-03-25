// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/commons/contactmanager.proto

package commons

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	state                     protoimpl.MessageState `protogen:"open.v1"`
	ContactManagerEntryId     int64                  `protobuf:"varint,1,opt,name=contact_manager_entry_id,json=contactManagerEntryId,proto3" json:"contact_manager_entry_id,omitempty"`
	ContactManagerEntryListId int64                  `protobuf:"varint,2,opt,name=contact_manager_entry_list_id,json=contactManagerEntryListId,proto3" json:"contact_manager_entry_list_id,omitempty"`
	Key                       string                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                     string                 `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Type                      string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	DateCreated               *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
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
	state                protoimpl.MessageState `protogen:"open.v1"`
	ContactManagerListId int64                  `protobuf:"varint,1,opt,name=contact_manager_list_id,json=contactManagerListId,proto3" json:"contact_manager_list_id,omitempty"`
	OrgId                string                 `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	ProjectId            int64                  `protobuf:"varint,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	FileName             string                 `protobuf:"bytes,4,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Description          string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// List of Columns In The Contact List. UI to render table based on this list
	ListDetails   []string               `protobuf:"bytes,6,rep,name=list_details,json=listDetails,proto3" json:"list_details,omitempty"`
	Ttl           int64                  `protobuf:"varint,7,opt,name=ttl,proto3" json:"ttl,omitempty"`
	DateCreated   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

const file_api_commons_contactmanager_proto_rawDesc = "" +
	"\n" +
	" api/commons/contactmanager.proto\x12\vapi.commons\x1a\x1fgoogle/protobuf/timestamp.proto\"\x97\x02\n" +
	"\x13ContactManagerEntry\x12;\n" +
	"\x18contact_manager_entry_id\x18\x01 \x01(\x03B\x020\x01R\x15contactManagerEntryId\x12D\n" +
	"\x1dcontact_manager_entry_list_id\x18\x02 \x01(\x03B\x020\x01R\x19contactManagerEntryListId\x12\x10\n" +
	"\x03key\x18\x03 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x04 \x01(\tR\x05value\x12\x12\n" +
	"\x04type\x18\x05 \x01(\tR\x04type\x12=\n" +
	"\fdate_created\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\vdateCreated:\x02\x18\x01\"\xc4\x02\n" +
	"\x12ContactManagerList\x129\n" +
	"\x17contact_manager_list_id\x18\x01 \x01(\x03B\x020\x01R\x14contactManagerListId\x12\x15\n" +
	"\x06org_id\x18\x02 \x01(\tR\x05orgId\x12!\n" +
	"\n" +
	"project_id\x18\x03 \x01(\x03B\x020\x01R\tprojectId\x12\x1b\n" +
	"\tfile_name\x18\x04 \x01(\tR\bfileName\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12!\n" +
	"\flist_details\x18\x06 \x03(\tR\vlistDetails\x12\x14\n" +
	"\x03ttl\x18\a \x01(\x03B\x020\x01R\x03ttl\x12=\n" +
	"\fdate_created\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\vdateCreated:\x02\x18\x01\"F\n" +
	"\x16ContactManagerEntryVal\x12\x12\n" +
	"\x04type\x18\x01 \x01(\tR\x04type\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x02\x18\x01*O\n" +
	"\x1aDeDuplicationMergeStrategy\x12\x16\n" +
	"\x12KEEP_EXISTING_LIST\x10\x00\x12\x19\n" +
	"\x15REPLACE_EXISTING_LIST\x10\x01*0\n" +
	"\x11ContactListStatus\x12\x1b\n" +
	"\x17CONTACT_LIST_STATUS_NEW\x10\x00*\xb4\x01\n" +
	"\x12ContactEntryStatus\x12\x1c\n" +
	"\x18CONTACT_ENTRY_STATUS_NEW\x10\x00\x12\x1e\n" +
	"\x1aCONTACT_ENTRY_STATUS_INUSE\x10\x01\x12\x1f\n" +
	"\x1bCONTACT_ENTRY_STATUS_NOTREF\x10\x02\x12\x1d\n" +
	"\x19CONTACT_ENTRY_STATUS_DONE\x10\x03\x12 \n" +
	"\x1cCONTACT_ENTRY_STATUS_EXPIRED\x10\x04B\x9b\x01\n" +
	"\x0fcom.api.commonsB\x13ContactmanagerProtoP\x01Z&github.com/tcncloud/api-go/api/commons\xa2\x02\x03ACX\xaa\x02\vApi.Commons\xca\x02\vApi\\Commons\xe2\x02\x17Api\\Commons\\GPBMetadata\xea\x02\fApi::Commonsb\x06proto3"

var (
	file_api_commons_contactmanager_proto_rawDescOnce sync.Once
	file_api_commons_contactmanager_proto_rawDescData []byte
)

func file_api_commons_contactmanager_proto_rawDescGZIP() []byte {
	file_api_commons_contactmanager_proto_rawDescOnce.Do(func() {
		file_api_commons_contactmanager_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_contactmanager_proto_rawDesc), len(file_api_commons_contactmanager_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_contactmanager_proto_rawDesc), len(file_api_commons_contactmanager_proto_rawDesc)),
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
	file_api_commons_contactmanager_proto_goTypes = nil
	file_api_commons_contactmanager_proto_depIdxs = nil
}
