// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: services/org/hunt_groups/v1alpha1/entities.proto

package hunt_groupsv1alpha1

import (
	org "github.com/tcncloud/api-go/api/commons/org"
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

// ExileLink defines a link between a hunt group and the exile service.
type ExileLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The sid or the unique ID of the exile link stored in the database.
	ParameterSid int64 `protobuf:"varint,1,opt,name=parameter_sid,json=parameterSid,proto3" json:"parameter_sid,omitempty"`
	// The sid of the hunt group which the exile link is associated with.
	HuntGroupSid int64 `protobuf:"varint,2,opt,name=hunt_group_sid,json=huntGroupSid,proto3" json:"hunt_group_sid,omitempty"`
	// The display name of the exile link.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the exile link.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// The order of the exile link, used when displaying.
	Order int64 `protobuf:"varint,5,opt,name=order,proto3" json:"order,omitempty"`
	// The parameter data for inbound calls.
	InboundData *ExileLinkData `protobuf:"bytes,6,opt,name=inbound_data,json=inboundData,proto3" json:"inbound_data,omitempty"`
	// The parameter data for outbound calls.
	OutboundData *ExileLinkData `protobuf:"bytes,7,opt,name=outbound_data,json=outboundData,proto3" json:"outbound_data,omitempty"`
}

func (x *ExileLink) Reset() {
	*x = ExileLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExileLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExileLink) ProtoMessage() {}

func (x *ExileLink) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExileLink.ProtoReflect.Descriptor instead.
func (*ExileLink) Descriptor() ([]byte, []int) {
	return file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescGZIP(), []int{0}
}

func (x *ExileLink) GetParameterSid() int64 {
	if x != nil {
		return x.ParameterSid
	}
	return 0
}

func (x *ExileLink) GetHuntGroupSid() int64 {
	if x != nil {
		return x.HuntGroupSid
	}
	return 0
}

func (x *ExileLink) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExileLink) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ExileLink) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *ExileLink) GetInboundData() *ExileLinkData {
	if x != nil {
		return x.InboundData
	}
	return nil
}

func (x *ExileLink) GetOutboundData() *ExileLinkData {
	if x != nil {
		return x.OutboundData
	}
	return nil
}

// ExileLinkData to be passed to exile.
type ExileLinkData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId    *ExileLinkParameter `protobuf:"bytes,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	AlternateId *ExileLinkParameter `protobuf:"bytes,2,opt,name=alternate_id,json=alternateId,proto3" json:"alternate_id,omitempty"`
}

func (x *ExileLinkData) Reset() {
	*x = ExileLinkData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExileLinkData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExileLinkData) ProtoMessage() {}

func (x *ExileLinkData) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExileLinkData.ProtoReflect.Descriptor instead.
func (*ExileLinkData) Descriptor() ([]byte, []int) {
	return file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescGZIP(), []int{1}
}

func (x *ExileLinkData) GetRecordId() *ExileLinkParameter {
	if x != nil {
		return x.RecordId
	}
	return nil
}

func (x *ExileLinkData) GetAlternateId() *ExileLinkParameter {
	if x != nil {
		return x.AlternateId
	}
	return nil
}

// ExileLinkParameter describes the data or data source.
type ExileLinkParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID for the contact field.
	ContactFieldSid int64 `protobuf:"varint,1,opt,name=contact_field_sid,json=contactFieldSid,proto3" json:"contact_field_sid,omitempty"`
	// The helper value of the parameter.
	HelperValue string `protobuf:"bytes,2,opt,name=helper_value,json=helperValue,proto3" json:"helper_value,omitempty"`
	// The source type of the parameter.
	ParameterSourceType org.ParameterSourceType `protobuf:"varint,3,opt,name=parameter_source_type,json=parameterSourceType,proto3,enum=api.commons.org.ParameterSourceType" json:"parameter_source_type,omitempty"`
}

func (x *ExileLinkParameter) Reset() {
	*x = ExileLinkParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExileLinkParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExileLinkParameter) ProtoMessage() {}

func (x *ExileLinkParameter) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExileLinkParameter.ProtoReflect.Descriptor instead.
func (*ExileLinkParameter) Descriptor() ([]byte, []int) {
	return file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescGZIP(), []int{2}
}

func (x *ExileLinkParameter) GetContactFieldSid() int64 {
	if x != nil {
		return x.ContactFieldSid
	}
	return 0
}

func (x *ExileLinkParameter) GetHelperValue() string {
	if x != nil {
		return x.HelperValue
	}
	return ""
}

func (x *ExileLinkParameter) GetParameterSourceType() org.ParameterSourceType {
	if x != nil {
		return x.ParameterSourceType
	}
	return org.ParameterSourceType(0)
}

type ListHuntGroupExileLinksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hunt group sid of the desired exile links
	HuntGroupSid int64 `protobuf:"varint,1,opt,name=hunt_group_sid,json=huntGroupSid,proto3" json:"hunt_group_sid,omitempty"`
}

func (x *ListHuntGroupExileLinksRequest) Reset() {
	*x = ListHuntGroupExileLinksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHuntGroupExileLinksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHuntGroupExileLinksRequest) ProtoMessage() {}

func (x *ListHuntGroupExileLinksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHuntGroupExileLinksRequest.ProtoReflect.Descriptor instead.
func (*ListHuntGroupExileLinksRequest) Descriptor() ([]byte, []int) {
	return file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescGZIP(), []int{3}
}

func (x *ListHuntGroupExileLinksRequest) GetHuntGroupSid() int64 {
	if x != nil {
		return x.HuntGroupSid
	}
	return 0
}

type ListHuntGroupExileLinksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The links for exile
	ExileLinks []*ExileLink `protobuf:"bytes,1,rep,name=exile_links,json=exileLinks,proto3" json:"exile_links,omitempty"`
}

func (x *ListHuntGroupExileLinksResponse) Reset() {
	*x = ListHuntGroupExileLinksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHuntGroupExileLinksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHuntGroupExileLinksResponse) ProtoMessage() {}

func (x *ListHuntGroupExileLinksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHuntGroupExileLinksResponse.ProtoReflect.Descriptor instead.
func (*ListHuntGroupExileLinksResponse) Descriptor() ([]byte, []int) {
	return file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescGZIP(), []int{4}
}

func (x *ListHuntGroupExileLinksResponse) GetExileLinks() []*ExileLink {
	if x != nil {
		return x.ExileLinks
	}
	return nil
}

var File_services_org_hunt_groups_v1alpha1_entities_proto protoreflect.FileDescriptor

var file_services_org_hunt_groups_v1alpha1_entities_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x68,
	0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x68, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x02, 0x0a, 0x09, 0x45, 0x78, 0x69, 0x6c, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x68, 0x75, 0x6e,
	0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x68, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x0c, 0x69,
	0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x0b, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x55, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x6c,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0xbd, 0x01, 0x0a, 0x0d, 0x45, 0x78, 0x69, 0x6c,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x52, 0x0a, 0x09, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x58, 0x0a,
	0x0c, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x61, 0x6c, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x12, 0x45, 0x78, 0x69, 0x6c,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x2a,
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x65,
	0x6c, 0x70, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x58, 0x0a,
	0x15, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x13, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x48,
	0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x68, 0x75, 0x6e,
	0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x68, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x64, 0x22,
	0x70, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x6c,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0a, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x42, 0xab, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0d, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x68, 0x75, 0x6e, 0x74, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x53, 0x4f, 0x48, 0xaa, 0x02, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x72, 0x67, 0x2e, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x23, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x48, 0x75, 0x6e, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescOnce sync.Once
	file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescData = file_services_org_hunt_groups_v1alpha1_entities_proto_rawDesc
)

func file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescGZIP() []byte {
	file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescOnce.Do(func() {
		file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescData)
	})
	return file_services_org_hunt_groups_v1alpha1_entities_proto_rawDescData
}

var file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_org_hunt_groups_v1alpha1_entities_proto_goTypes = []any{
	(*ExileLink)(nil),                       // 0: services.org.hunt_groups.v1alpha1.ExileLink
	(*ExileLinkData)(nil),                   // 1: services.org.hunt_groups.v1alpha1.ExileLinkData
	(*ExileLinkParameter)(nil),              // 2: services.org.hunt_groups.v1alpha1.ExileLinkParameter
	(*ListHuntGroupExileLinksRequest)(nil),  // 3: services.org.hunt_groups.v1alpha1.ListHuntGroupExileLinksRequest
	(*ListHuntGroupExileLinksResponse)(nil), // 4: services.org.hunt_groups.v1alpha1.ListHuntGroupExileLinksResponse
	(org.ParameterSourceType)(0),            // 5: api.commons.org.ParameterSourceType
}
var file_services_org_hunt_groups_v1alpha1_entities_proto_depIdxs = []int32{
	1, // 0: services.org.hunt_groups.v1alpha1.ExileLink.inbound_data:type_name -> services.org.hunt_groups.v1alpha1.ExileLinkData
	1, // 1: services.org.hunt_groups.v1alpha1.ExileLink.outbound_data:type_name -> services.org.hunt_groups.v1alpha1.ExileLinkData
	2, // 2: services.org.hunt_groups.v1alpha1.ExileLinkData.record_id:type_name -> services.org.hunt_groups.v1alpha1.ExileLinkParameter
	2, // 3: services.org.hunt_groups.v1alpha1.ExileLinkData.alternate_id:type_name -> services.org.hunt_groups.v1alpha1.ExileLinkParameter
	5, // 4: services.org.hunt_groups.v1alpha1.ExileLinkParameter.parameter_source_type:type_name -> api.commons.org.ParameterSourceType
	0, // 5: services.org.hunt_groups.v1alpha1.ListHuntGroupExileLinksResponse.exile_links:type_name -> services.org.hunt_groups.v1alpha1.ExileLink
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_services_org_hunt_groups_v1alpha1_entities_proto_init() }
func file_services_org_hunt_groups_v1alpha1_entities_proto_init() {
	if File_services_org_hunt_groups_v1alpha1_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExileLink); i {
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
		file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ExileLinkData); i {
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
		file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ExileLinkParameter); i {
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
		file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListHuntGroupExileLinksRequest); i {
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
		file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListHuntGroupExileLinksResponse); i {
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
			RawDescriptor: file_services_org_hunt_groups_v1alpha1_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_org_hunt_groups_v1alpha1_entities_proto_goTypes,
		DependencyIndexes: file_services_org_hunt_groups_v1alpha1_entities_proto_depIdxs,
		MessageInfos:      file_services_org_hunt_groups_v1alpha1_entities_proto_msgTypes,
	}.Build()
	File_services_org_hunt_groups_v1alpha1_entities_proto = out.File
	file_services_org_hunt_groups_v1alpha1_entities_proto_rawDesc = nil
	file_services_org_hunt_groups_v1alpha1_entities_proto_goTypes = nil
	file_services_org_hunt_groups_v1alpha1_entities_proto_depIdxs = nil
}
