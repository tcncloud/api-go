// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha4/plan.proto

package entitiesv1alpha4

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
	state         protoimpl.MessageState `protogen:"open.v1"`
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
	OrgId         string `protobuf:"bytes,10,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

const file_services_billing_entities_v1alpha4_plan_proto_rawDesc = "" +
	"\n" +
	"-services/billing/entities/v1alpha4/plan.proto\x12\"services.billing.entities.v1alpha4\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd1\x03\n" +
	"\vBillingPlan\x12&\n" +
	"\x0fbilling_plan_id\x18\x01 \x01(\tR\rbillingPlanId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12G\n" +
	"\x04type\x18\x03 \x01(\x0e23.services.billing.entities.v1alpha4.BillingPlanTypeR\x04type\x12\x19\n" +
	"\bis_draft\x18\x04 \x01(\bR\aisDraft\x129\n" +
	"\n" +
	"start_time\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTime\x12;\n" +
	"\vcreate_time\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x12;\n" +
	"\vupdate_time\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"updateTime\x12;\n" +
	"\vdelete_time\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"deleteTime\x12\x17\n" +
	"\auser_id\x18\t \x01(\tR\x06userId\x12\x15\n" +
	"\x06org_id\x18\n" +
	" \x01(\tR\x05orgId*n\n" +
	"\x0fBillingPlanType\x12!\n" +
	"\x1dBILLING_PLAN_TYPE_UNSPECIFIED\x10\x00\x12\x1d\n" +
	"\x19BILLING_PLAN_TYPE_DEFAULT\x10\x01\x12\x19\n" +
	"\x15BILLING_PLAN_TYPE_ORG\x10\x02B\xae\x02\n" +
	"&com.services.billing.entities.v1alpha4B\tPlanProtoP\x01ZNgithub.com/tcncloud/api-go/services/billing/entities/v1alpha4;entitiesv1alpha4\xa2\x02\x03SBE\xaa\x02\"Services.Billing.Entities.V1alpha4\xca\x02\"Services\\Billing\\Entities\\V1alpha4\xe2\x02.Services\\Billing\\Entities\\V1alpha4\\GPBMetadata\xea\x02%Services::Billing::Entities::V1alpha4b\x06proto3"

var (
	file_services_billing_entities_v1alpha4_plan_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha4_plan_proto_rawDescData []byte
)

func file_services_billing_entities_v1alpha4_plan_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha4_plan_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha4_plan_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha4_plan_proto_rawDesc), len(file_services_billing_entities_v1alpha4_plan_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha4_plan_proto_rawDesc), len(file_services_billing_entities_v1alpha4_plan_proto_rawDesc)),
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
	file_services_billing_entities_v1alpha4_plan_proto_goTypes = nil
	file_services_billing_entities_v1alpha4_plan_proto_depIdxs = nil
}
