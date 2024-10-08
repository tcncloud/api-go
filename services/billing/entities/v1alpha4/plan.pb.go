// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha4/plan.proto

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

type BillingPlanType int32

const (
	BillingPlanType_BILLING_PLAN_TYPE_UNSPECIFIED BillingPlanType = 0
	BillingPlanType_BILLING_PLAN_TYPE_DEFAULT     BillingPlanType = 1
	BillingPlanType_BILLING_PLAN_TYPE_ORG         BillingPlanType = 2
)

// Enum value maps for BillingPlanType.
var (
	BillingPlanType_name = map[int32]string{
		0: "BILLING_PLAN_TYPE_UNSPECIFIED",
		1: "BILLING_PLAN_TYPE_DEFAULT",
		2: "BILLING_PLAN_TYPE_ORG",
	}
	BillingPlanType_value = map[string]int32{
		"BILLING_PLAN_TYPE_UNSPECIFIED": 0,
		"BILLING_PLAN_TYPE_DEFAULT":     1,
		"BILLING_PLAN_TYPE_ORG":         2,
	}
)

func (x BillingPlanType) Enum() *BillingPlanType {
	p := new(BillingPlanType)
	*p = x
	return p
}

func (x BillingPlanType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BillingPlanType) Descriptor() protoreflect.EnumDescriptor {
	return file_services_billing_entities_v1alpha4_plan_proto_enumTypes[0].Descriptor()
}

func (BillingPlanType) Type() protoreflect.EnumType {
	return &file_services_billing_entities_v1alpha4_plan_proto_enumTypes[0]
}

func (x BillingPlanType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BillingPlanType.Descriptor instead.
func (BillingPlanType) EnumDescriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha4_plan_proto_rawDescGZIP(), []int{0}
}

type BillingPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingPlanId string                 `protobuf:"bytes,1,opt,name=billing_plan_id,json=billingPlanId,proto3" json:"billing_plan_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type          BillingPlanType        `protobuf:"varint,3,opt,name=type,proto3,enum=services.billing.entities.v1alpha4.BillingPlanType" json:"type,omitempty"`
	IsDraft       bool                   `protobuf:"varint,4,opt,name=is_draft,json=isDraft,proto3" json:"is_draft,omitempty"`
	StartTime     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DeleteTime    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	UserId        string                 `protobuf:"bytes,9,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Optional. if set, this billing plan is for the org, and as
	// a draft can only be applied to that org.
	OrgId string `protobuf:"bytes,10,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
}

func (x *BillingPlan) Reset() {
	*x = BillingPlan{}
	mi := &file_services_billing_entities_v1alpha4_plan_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BillingPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingPlan) ProtoMessage() {}

func (x *BillingPlan) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_plan_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingPlan.ProtoReflect.Descriptor instead.
func (*BillingPlan) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha4_plan_proto_rawDescGZIP(), []int{0}
}

func (x *BillingPlan) GetBillingPlanId() string {
	if x != nil {
		return x.BillingPlanId
	}
	return ""
}

func (x *BillingPlan) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BillingPlan) GetType() BillingPlanType {
	if x != nil {
		return x.Type
	}
	return BillingPlanType_BILLING_PLAN_TYPE_UNSPECIFIED
}

func (x *BillingPlan) GetIsDraft() bool {
	if x != nil {
		return x.IsDraft
	}
	return false
}

func (x *BillingPlan) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *BillingPlan) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BillingPlan) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *BillingPlan) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *BillingPlan) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BillingPlan) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

var File_services_billing_entities_v1alpha4_plan_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha4_plan_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x34, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x34, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x03, 0x0a, 0x0b, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x50, 0x6c, 0x61, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x47, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x34, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x2a, 0x6e, 0x0a, 0x0f, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x42,
	0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d,
	0x0a, 0x19, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x12, 0x19, 0x0a,
	0x15, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4f, 0x52, 0x47, 0x10, 0x02, 0x42, 0xae, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x34, 0x42, 0x09, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x3b,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34,
	0xa2, 0x02, 0x03, 0x53, 0x42, 0x45, 0xaa, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0xca, 0x02, 0x22, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34,
	0xe2, 0x02, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x25, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_services_billing_entities_v1alpha4_plan_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha4_plan_proto_rawDescData = file_services_billing_entities_v1alpha4_plan_proto_rawDesc
)

func file_services_billing_entities_v1alpha4_plan_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha4_plan_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha4_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_billing_entities_v1alpha4_plan_proto_rawDescData)
	})
	return file_services_billing_entities_v1alpha4_plan_proto_rawDescData
}

var file_services_billing_entities_v1alpha4_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_billing_entities_v1alpha4_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_services_billing_entities_v1alpha4_plan_proto_goTypes = []any{
	(BillingPlanType)(0),          // 0: services.billing.entities.v1alpha4.BillingPlanType
	(*BillingPlan)(nil),           // 1: services.billing.entities.v1alpha4.BillingPlan
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_services_billing_entities_v1alpha4_plan_proto_depIdxs = []int32{
	0, // 0: services.billing.entities.v1alpha4.BillingPlan.type:type_name -> services.billing.entities.v1alpha4.BillingPlanType
	2, // 1: services.billing.entities.v1alpha4.BillingPlan.start_time:type_name -> google.protobuf.Timestamp
	2, // 2: services.billing.entities.v1alpha4.BillingPlan.create_time:type_name -> google.protobuf.Timestamp
	2, // 3: services.billing.entities.v1alpha4.BillingPlan.update_time:type_name -> google.protobuf.Timestamp
	2, // 4: services.billing.entities.v1alpha4.BillingPlan.delete_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha4_plan_proto_init() }
func file_services_billing_entities_v1alpha4_plan_proto_init() {
	if File_services_billing_entities_v1alpha4_plan_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_billing_entities_v1alpha4_plan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha4_plan_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha4_plan_proto_depIdxs,
		EnumInfos:         file_services_billing_entities_v1alpha4_plan_proto_enumTypes,
		MessageInfos:      file_services_billing_entities_v1alpha4_plan_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha4_plan_proto = out.File
	file_services_billing_entities_v1alpha4_plan_proto_rawDesc = nil
	file_services_billing_entities_v1alpha4_plan_proto_goTypes = nil
	file_services_billing_entities_v1alpha4_plan_proto_depIdxs = nil
}
