// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: api/commons/org/organization.proto

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

// Organization contains the basic properties for an organization.
type Organization struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// An organization's globally unique identifier.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Map of regionId to clientSid. Shows an org's enabled regions.
	EnabledRegions map[string]int64 `protobuf:"bytes,2,rep,name=enabled_regions,json=enabledRegions,proto3" json:"enabled_regions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Default region org was first enbabled for.
	RegionId string `protobuf:"bytes,3,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// Billing_prefix + clientSid. Used for integrations and billing.
	BillingId string `protobuf:"bytes,4,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
	// First enabled region's clientSid; will be deprecated for
	// enabled_regions map.
	ClientSid int64 `protobuf:"varint,5,opt,name=client_sid,json=clientSid,proto3" json:"client_sid,omitempty"`
	// The organization name.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// Date of organization creation.
	AddDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=add_date,json=addDate,proto3" json:"add_date,omitempty"`
	// Whether account is manual only.
	IsManualOnlyAccount bool `protobuf:"varint,8,opt,name=is_manual_only_account,json=isManualOnlyAccount,proto3" json:"is_manual_only_account,omitempty"`
	// Backoffice UI theme/skin.
	BackofficeTheme commons.ClientSkin `protobuf:"varint,9,opt,name=backoffice_theme,json=backofficeTheme,proto3,enum=api.commons.ClientSkin" json:"backoffice_theme,omitempty"`
	// Archived specifies this organization as no longer active.
	Archived bool `protobuf:"varint,10,opt,name=archived,proto3" json:"archived,omitempty"`
	// Salesforce ID used for integrations.
	CrmId string `protobuf:"bytes,11,opt,name=crm_id,json=crmId,proto3" json:"crm_id,omitempty"`
	// Organization's time zone.
	TimeZone commons.TimeZone `protobuf:"varint,12,opt,name=time_zone,json=timeZone,proto3,enum=api.commons.TimeZone" json:"time_zone,omitempty"`
	// Organization callbacks service ID.
	CallbacksServiceId string `protobuf:"bytes,13,opt,name=callbacks_service_id,json=callbacksServiceId,proto3" json:"callbacks_service_id,omitempty"`
	// Organization that is the parent.
	P3OwnerId     string `protobuf:"bytes,14,opt,name=p3_owner_id,json=p3OwnerId,proto3" json:"p3_owner_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Organization) Reset() {
	*x = Organization{}
	mi := &file_api_commons_org_organization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_organization_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_api_commons_org_organization_proto_rawDescGZIP(), []int{0}
}

func (x *Organization) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Organization) GetEnabledRegions() map[string]int64 {
	if x != nil {
		return x.EnabledRegions
	}
	return nil
}

func (x *Organization) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *Organization) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

func (x *Organization) GetClientSid() int64 {
	if x != nil {
		return x.ClientSid
	}
	return 0
}

func (x *Organization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Organization) GetAddDate() *timestamppb.Timestamp {
	if x != nil {
		return x.AddDate
	}
	return nil
}

func (x *Organization) GetIsManualOnlyAccount() bool {
	if x != nil {
		return x.IsManualOnlyAccount
	}
	return false
}

func (x *Organization) GetBackofficeTheme() commons.ClientSkin {
	if x != nil {
		return x.BackofficeTheme
	}
	return commons.ClientSkin(0)
}

func (x *Organization) GetArchived() bool {
	if x != nil {
		return x.Archived
	}
	return false
}

func (x *Organization) GetCrmId() string {
	if x != nil {
		return x.CrmId
	}
	return ""
}

func (x *Organization) GetTimeZone() commons.TimeZone {
	if x != nil {
		return x.TimeZone
	}
	return commons.TimeZone(0)
}

func (x *Organization) GetCallbacksServiceId() string {
	if x != nil {
		return x.CallbacksServiceId
	}
	return ""
}

func (x *Organization) GetP3OwnerId() string {
	if x != nil {
		return x.P3OwnerId
	}
	return ""
}

type OrganizationDetails struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Organization      *Organization          `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	LastScheduledDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_scheduled_date,json=lastScheduledDate,proto3" json:"last_scheduled_date,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *OrganizationDetails) Reset() {
	*x = OrganizationDetails{}
	mi := &file_api_commons_org_organization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationDetails) ProtoMessage() {}

func (x *OrganizationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_organization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationDetails.ProtoReflect.Descriptor instead.
func (*OrganizationDetails) Descriptor() ([]byte, []int) {
	return file_api_commons_org_organization_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationDetails) GetOrganization() *Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

func (x *OrganizationDetails) GetLastScheduledDate() *timestamppb.Timestamp {
	if x != nil {
		return x.LastScheduledDate
	}
	return nil
}

var File_api_commons_org_organization_proto protoreflect.FileDescriptor

var file_api_commons_org_organization_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x6f, 0x72, 0x67, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x05,
	0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15,
	0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x5a, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x69, 0x73, 0x5f, 0x6d, 0x61,
	0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x69, 0x73, 0x4d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x4f, 0x6e, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x10,
	0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x68, 0x65, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x6b, 0x69, 0x6e, 0x52,
	0x0f, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x54, 0x68, 0x65, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x63, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x72,
	0x6d, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x70, 0x33, 0x5f,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x33, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x41, 0x0a, 0x13, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa4, 0x01, 0x0a,
	0x13, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x42, 0xb2, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x42, 0x11, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0xa2, 0x02, 0x03, 0x41,
	0x43, 0x4f, 0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x4f, 0x72, 0x67, 0xca, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0xe2, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_org_organization_proto_rawDescOnce sync.Once
	file_api_commons_org_organization_proto_rawDescData = file_api_commons_org_organization_proto_rawDesc
)

func file_api_commons_org_organization_proto_rawDescGZIP() []byte {
	file_api_commons_org_organization_proto_rawDescOnce.Do(func() {
		file_api_commons_org_organization_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_org_organization_proto_rawDescData)
	})
	return file_api_commons_org_organization_proto_rawDescData
}

var file_api_commons_org_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_commons_org_organization_proto_goTypes = []any{
	(*Organization)(nil),          // 0: api.commons.org.Organization
	(*OrganizationDetails)(nil),   // 1: api.commons.org.OrganizationDetails
	nil,                           // 2: api.commons.org.Organization.EnabledRegionsEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(commons.ClientSkin)(0),       // 4: api.commons.ClientSkin
	(commons.TimeZone)(0),         // 5: api.commons.TimeZone
}
var file_api_commons_org_organization_proto_depIdxs = []int32{
	2, // 0: api.commons.org.Organization.enabled_regions:type_name -> api.commons.org.Organization.EnabledRegionsEntry
	3, // 1: api.commons.org.Organization.add_date:type_name -> google.protobuf.Timestamp
	4, // 2: api.commons.org.Organization.backoffice_theme:type_name -> api.commons.ClientSkin
	5, // 3: api.commons.org.Organization.time_zone:type_name -> api.commons.TimeZone
	0, // 4: api.commons.org.OrganizationDetails.organization:type_name -> api.commons.org.Organization
	3, // 5: api.commons.org.OrganizationDetails.last_scheduled_date:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_commons_org_organization_proto_init() }
func file_api_commons_org_organization_proto_init() {
	if File_api_commons_org_organization_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_org_organization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_org_organization_proto_goTypes,
		DependencyIndexes: file_api_commons_org_organization_proto_depIdxs,
		MessageInfos:      file_api_commons_org_organization_proto_msgTypes,
	}.Build()
	File_api_commons_org_organization_proto = out.File
	file_api_commons_org_organization_proto_rawDesc = nil
	file_api_commons_org_organization_proto_goTypes = nil
	file_api_commons_org_organization_proto_depIdxs = nil
}
