// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha1/plan.proto

package entitiesv1alpha1

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

// BillingPlanStatus represents the status of a billing plan
type BillingPlanStatus int32

const (
	BillingPlanStatus_BILLING_PLAN_STATUS_UNSPECIFIED BillingPlanStatus = 0
	// A billing plan is initially created in "creating" status
	// and is transitioned to "created" explicitly by committing
	// the billing plan changes.
	BillingPlanStatus_BILLING_PLAN_STATUS_CREATING BillingPlanStatus = 100
	// A billing plan is in "created" status when it is ready to
	// be used.
	BillingPlanStatus_BILLING_PLAN_STATUS_CREATED BillingPlanStatus = 200
)

// Enum value maps for BillingPlanStatus.
var (
	BillingPlanStatus_name = map[int32]string{
		0:   "BILLING_PLAN_STATUS_UNSPECIFIED",
		100: "BILLING_PLAN_STATUS_CREATING",
		200: "BILLING_PLAN_STATUS_CREATED",
	}
	BillingPlanStatus_value = map[string]int32{
		"BILLING_PLAN_STATUS_UNSPECIFIED": 0,
		"BILLING_PLAN_STATUS_CREATING":    100,
		"BILLING_PLAN_STATUS_CREATED":     200,
	}
)

func (x BillingPlanStatus) Enum() *BillingPlanStatus {
	p := new(BillingPlanStatus)
	*p = x
	return p
}

func (x BillingPlanStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BillingPlanStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_services_billing_entities_v1alpha1_plan_proto_enumTypes[0].Descriptor()
}

func (BillingPlanStatus) Type() protoreflect.EnumType {
	return &file_services_billing_entities_v1alpha1_plan_proto_enumTypes[0]
}

func (x BillingPlanStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BillingPlanStatus.Descriptor instead.
func (BillingPlanStatus) EnumDescriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha1_plan_proto_rawDescGZIP(), []int{0}
}

// BillingPlan represents a collection of rate definitions
type BillingPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the billing plan identifier
	BillingPlanId string `protobuf:"bytes,1,opt,name=billing_plan_id,json=billingPlanId,proto3" json:"billing_plan_id,omitempty"`
	// time the billing plan was created
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// time the billing plan was last updated
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// the time from which this billing plan will take effect
	StartTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// the time this billing plan ended; can be null
	EndTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// the time this billing plan was deleted; can be null
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// the billing plan rate definitions
	RateDefinitionIds []string `protobuf:"bytes,7,rep,name=rate_definition_ids,json=rateDefinitionIds,proto3" json:"rate_definition_ids,omitempty"`
	// the billing plan status
	Status BillingPlanStatus `protobuf:"varint,8,opt,name=status,proto3,enum=services.billing.entities.v1alpha1.BillingPlanStatus" json:"status,omitempty"`
}

func (x *BillingPlan) Reset() {
	*x = BillingPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_billing_entities_v1alpha1_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingPlan) ProtoMessage() {}

func (x *BillingPlan) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha1_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_services_billing_entities_v1alpha1_plan_proto_rawDescGZIP(), []int{0}
}

func (x *BillingPlan) GetBillingPlanId() string {
	if x != nil {
		return x.BillingPlanId
	}
	return ""
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

func (x *BillingPlan) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *BillingPlan) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *BillingPlan) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *BillingPlan) GetRateDefinitionIds() []string {
	if x != nil {
		return x.RateDefinitionIds
	}
	return nil
}

func (x *BillingPlan) GetStatus() BillingPlanStatus {
	if x != nil {
		return x.Status
	}
	return BillingPlanStatus_BILLING_PLAN_STATUS_UNSPECIFIED
}

var File_services_billing_entities_v1alpha1_plan_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha1_plan_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x03, 0x0a, 0x0b, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x50, 0x6c, 0x61, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x11, 0x72, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x4d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2a, 0x7c, 0x0a, 0x11, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50,
	0x6c, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x1f, 0x42, 0x49, 0x4c,
	0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20,
	0x0a, 0x1c, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x64,
	0x12, 0x20, 0x0a, 0x1b, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4c, 0x41, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10,
	0xc8, 0x01, 0x42, 0xae, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x09, 0x50,
	0x6c, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x42, 0x45,
	0xaa, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a,
	0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_billing_entities_v1alpha1_plan_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha1_plan_proto_rawDescData = file_services_billing_entities_v1alpha1_plan_proto_rawDesc
)

func file_services_billing_entities_v1alpha1_plan_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha1_plan_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha1_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_billing_entities_v1alpha1_plan_proto_rawDescData)
	})
	return file_services_billing_entities_v1alpha1_plan_proto_rawDescData
}

var file_services_billing_entities_v1alpha1_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_billing_entities_v1alpha1_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_services_billing_entities_v1alpha1_plan_proto_goTypes = []interface{}{
	(BillingPlanStatus)(0),        // 0: services.billing.entities.v1alpha1.BillingPlanStatus
	(*BillingPlan)(nil),           // 1: services.billing.entities.v1alpha1.BillingPlan
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_services_billing_entities_v1alpha1_plan_proto_depIdxs = []int32{
	2, // 0: services.billing.entities.v1alpha1.BillingPlan.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: services.billing.entities.v1alpha1.BillingPlan.update_time:type_name -> google.protobuf.Timestamp
	2, // 2: services.billing.entities.v1alpha1.BillingPlan.start_time:type_name -> google.protobuf.Timestamp
	2, // 3: services.billing.entities.v1alpha1.BillingPlan.end_time:type_name -> google.protobuf.Timestamp
	2, // 4: services.billing.entities.v1alpha1.BillingPlan.delete_time:type_name -> google.protobuf.Timestamp
	0, // 5: services.billing.entities.v1alpha1.BillingPlan.status:type_name -> services.billing.entities.v1alpha1.BillingPlanStatus
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha1_plan_proto_init() }
func file_services_billing_entities_v1alpha1_plan_proto_init() {
	if File_services_billing_entities_v1alpha1_plan_proto != nil {
		return
	}
	file_services_billing_entities_v1alpha1_rates_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_billing_entities_v1alpha1_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingPlan); i {
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
			RawDescriptor: file_services_billing_entities_v1alpha1_plan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha1_plan_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha1_plan_proto_depIdxs,
		EnumInfos:         file_services_billing_entities_v1alpha1_plan_proto_enumTypes,
		MessageInfos:      file_services_billing_entities_v1alpha1_plan_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha1_plan_proto = out.File
	file_services_billing_entities_v1alpha1_plan_proto_rawDesc = nil
	file_services_billing_entities_v1alpha1_plan_proto_goTypes = nil
	file_services_billing_entities_v1alpha1_plan_proto_depIdxs = nil
}
