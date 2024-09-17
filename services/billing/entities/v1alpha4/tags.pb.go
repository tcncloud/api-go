// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha4/tags.proto

package entitiesv1alpha4

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

type Category int32

const (
	Category_CATEGORY_UNSPECIFIED                                                 Category = 0
	Category_CATEGORY_COMMUNICATIONS_OMNI_CHAT                                    Category = 100
	Category_CATEGORY_COMMUNICATIONS_OMNI_EMAIL                                   Category = 101
	Category_CATEGORY_COMMUNICATIONS_OMNI_SMS                                     Category = 102
	Category_CATEGORY_COMMUNICATIONS_OMNI_AGENT                                   Category = 103
	Category_CATEGORY_COMMUNICATIONS_OMNI_RESOURCES                               Category = 104
	Category_CATEGORY_DATA_MANAGEMENT_COMPLIANCE_COMPLIANCE                       Category = 200
	Category_CATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_OPTIMIZATION_VOICE_ANALYTICS Category = 300
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0:   "CATEGORY_UNSPECIFIED",
		100: "CATEGORY_COMMUNICATIONS_OMNI_CHAT",
		101: "CATEGORY_COMMUNICATIONS_OMNI_EMAIL",
		102: "CATEGORY_COMMUNICATIONS_OMNI_SMS",
		103: "CATEGORY_COMMUNICATIONS_OMNI_AGENT",
		104: "CATEGORY_COMMUNICATIONS_OMNI_RESOURCES",
		200: "CATEGORY_DATA_MANAGEMENT_COMPLIANCE_COMPLIANCE",
		300: "CATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_OPTIMIZATION_VOICE_ANALYTICS",
	}
	Category_value = map[string]int32{
		"CATEGORY_UNSPECIFIED":                                                 0,
		"CATEGORY_COMMUNICATIONS_OMNI_CHAT":                                    100,
		"CATEGORY_COMMUNICATIONS_OMNI_EMAIL":                                   101,
		"CATEGORY_COMMUNICATIONS_OMNI_SMS":                                     102,
		"CATEGORY_COMMUNICATIONS_OMNI_AGENT":                                   103,
		"CATEGORY_COMMUNICATIONS_OMNI_RESOURCES":                               104,
		"CATEGORY_DATA_MANAGEMENT_COMPLIANCE_COMPLIANCE":                       200,
		"CATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_OPTIMIZATION_VOICE_ANALYTICS": 300,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_services_billing_entities_v1alpha4_tags_proto_enumTypes[0].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_services_billing_entities_v1alpha4_tags_proto_enumTypes[0]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha4_tags_proto_rawDescGZIP(), []int{0}
}

type BillingTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingTagId string                 `protobuf:"bytes,1,opt,name=billing_tag_id,json=billingTagId,proto3" json:"billing_tag_id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DeleteTime   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/tags.proto.
	Category        string   `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	BillingCategory Category `protobuf:"varint,7,opt,name=billing_category,json=billingCategory,proto3,enum=services.billing.entities.v1alpha4.Category" json:"billing_category,omitempty"`
}

func (x *BillingTag) Reset() {
	*x = BillingTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_billing_entities_v1alpha4_tags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingTag) ProtoMessage() {}

func (x *BillingTag) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_tags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingTag.ProtoReflect.Descriptor instead.
func (*BillingTag) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha4_tags_proto_rawDescGZIP(), []int{0}
}

func (x *BillingTag) GetBillingTagId() string {
	if x != nil {
		return x.BillingTagId
	}
	return ""
}

func (x *BillingTag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BillingTag) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BillingTag) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *BillingTag) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/tags.proto.
func (x *BillingTag) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *BillingTag) GetBillingCategory() Category {
	if x != nil {
		return x.BillingCategory
	}
	return Category_CATEGORY_UNSPECIFIED
}

var File_services_billing_entities_v1alpha4_tags_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha4_tags_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x34, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x34, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x02, 0x0a, 0x0a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x54, 0x61, 0x67, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x57, 0x0a, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x34, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2a, 0xed, 0x02,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x55, 0x4e, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f,
	0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x10, 0x64, 0x12, 0x26, 0x0a, 0x22, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x55, 0x4e, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x45, 0x4d, 0x41, 0x49,
	0x4c, 0x10, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f,
	0x43, 0x4f, 0x4d, 0x4d, 0x55, 0x4e, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x4f,
	0x4d, 0x4e, 0x49, 0x5f, 0x53, 0x4d, 0x53, 0x10, 0x66, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x41, 0x54,
	0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x55, 0x4e, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10,
	0x67, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x4f,
	0x4d, 0x4d, 0x55, 0x4e, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x4f, 0x4d, 0x4e,
	0x49, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x53, 0x10, 0x68, 0x12, 0x33, 0x0a,
	0x2e, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4d,
	0x41, 0x4e, 0x41, 0x47, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49,
	0x41, 0x4e, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x10,
	0xc8, 0x01, 0x12, 0x49, 0x0a, 0x44, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x57,
	0x4f, 0x52, 0x4b, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x4e, 0x47, 0x41, 0x47, 0x45, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x4f, 0x50,
	0x54, 0x49, 0x4d, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x4f, 0x49, 0x43, 0x45,
	0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x10, 0xac, 0x02, 0x42, 0xae, 0x02,
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x42, 0x09, 0x54, 0x61, 0x67, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x34, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x34, 0xa2, 0x02, 0x03, 0x53, 0x42, 0x45, 0xaa, 0x02, 0x22, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34,
	0xca, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x34, 0xe2, 0x02, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_billing_entities_v1alpha4_tags_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha4_tags_proto_rawDescData = file_services_billing_entities_v1alpha4_tags_proto_rawDesc
)

func file_services_billing_entities_v1alpha4_tags_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha4_tags_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha4_tags_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_billing_entities_v1alpha4_tags_proto_rawDescData)
	})
	return file_services_billing_entities_v1alpha4_tags_proto_rawDescData
}

var file_services_billing_entities_v1alpha4_tags_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_billing_entities_v1alpha4_tags_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_services_billing_entities_v1alpha4_tags_proto_goTypes = []any{
	(Category)(0),                 // 0: services.billing.entities.v1alpha4.Category
	(*BillingTag)(nil),            // 1: services.billing.entities.v1alpha4.BillingTag
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_services_billing_entities_v1alpha4_tags_proto_depIdxs = []int32{
	2, // 0: services.billing.entities.v1alpha4.BillingTag.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: services.billing.entities.v1alpha4.BillingTag.update_time:type_name -> google.protobuf.Timestamp
	2, // 2: services.billing.entities.v1alpha4.BillingTag.delete_time:type_name -> google.protobuf.Timestamp
	0, // 3: services.billing.entities.v1alpha4.BillingTag.billing_category:type_name -> services.billing.entities.v1alpha4.Category
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha4_tags_proto_init() }
func file_services_billing_entities_v1alpha4_tags_proto_init() {
	if File_services_billing_entities_v1alpha4_tags_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_billing_entities_v1alpha4_tags_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BillingTag); i {
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
			RawDescriptor: file_services_billing_entities_v1alpha4_tags_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha4_tags_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha4_tags_proto_depIdxs,
		EnumInfos:         file_services_billing_entities_v1alpha4_tags_proto_enumTypes,
		MessageInfos:      file_services_billing_entities_v1alpha4_tags_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha4_tags_proto = out.File
	file_services_billing_entities_v1alpha4_tags_proto_rawDesc = nil
	file_services_billing_entities_v1alpha4_tags_proto_goTypes = nil
	file_services_billing_entities_v1alpha4_tags_proto_depIdxs = nil
}
