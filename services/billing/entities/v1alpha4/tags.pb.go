// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha4/tags.proto

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
	Category_CATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_OPTIMIZATION_AI_BUNDLE       Category = 400
	Category_CATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_MANAGEMENT_SCHEDULER         Category = 500
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
		400: "CATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_OPTIMIZATION_AI_BUNDLE",
		500: "CATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_MANAGEMENT_SCHEDULER",
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
		"CATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_OPTIMIZATION_AI_BUNDLE":       400,
		"CATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_MANAGEMENT_SCHEDULER":         500,
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
	state        protoimpl.MessageState `protogen:"open.v1"`
	BillingTagId string                 `protobuf:"bytes,1,opt,name=billing_tag_id,json=billingTagId,proto3" json:"billing_tag_id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DeleteTime   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/tags.proto.
	Category        string   `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	BillingCategory Category `protobuf:"varint,7,opt,name=billing_category,json=billingCategory,proto3,enum=services.billing.entities.v1alpha4.Category" json:"billing_category,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *BillingTag) Reset() {
	*x = BillingTag{}
	mi := &file_services_billing_entities_v1alpha4_tags_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BillingTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingTag) ProtoMessage() {}

func (x *BillingTag) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_tags_proto_msgTypes[0]
	if x != nil {
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

const file_services_billing_entities_v1alpha4_tags_proto_rawDesc = "" +
	"\n" +
	"-services/billing/entities/v1alpha4/tags.proto\x12\"services.billing.entities.v1alpha4\x1a\x1fgoogle/protobuf/timestamp.proto\"\xf6\x02\n" +
	"\n" +
	"BillingTag\x12$\n" +
	"\x0ebilling_tag_id\x18\x01 \x01(\tR\fbillingTagId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12;\n" +
	"\vcreate_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x12;\n" +
	"\vupdate_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"updateTime\x12;\n" +
	"\vdelete_time\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"deleteTime\x12\x1e\n" +
	"\bcategory\x18\x06 \x01(\tB\x02\x18\x01R\bcategory\x12W\n" +
	"\x10billing_category\x18\a \x01(\x0e2,.services.billing.entities.v1alpha4.CategoryR\x0fbillingCategory*\xf5\x03\n" +
	"\bCategory\x12\x18\n" +
	"\x14CATEGORY_UNSPECIFIED\x10\x00\x12%\n" +
	"!CATEGORY_COMMUNICATIONS_OMNI_CHAT\x10d\x12&\n" +
	"\"CATEGORY_COMMUNICATIONS_OMNI_EMAIL\x10e\x12$\n" +
	" CATEGORY_COMMUNICATIONS_OMNI_SMS\x10f\x12&\n" +
	"\"CATEGORY_COMMUNICATIONS_OMNI_AGENT\x10g\x12*\n" +
	"&CATEGORY_COMMUNICATIONS_OMNI_RESOURCES\x10h\x123\n" +
	".CATEGORY_DATA_MANAGEMENT_COMPLIANCE_COMPLIANCE\x10\xc8\x01\x12I\n" +
	"DCATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_OPTIMIZATION_VOICE_ANALYTICS\x10\xac\x02\x12C\n" +
	">CATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_OPTIMIZATION_AI_BUNDLE\x10\x90\x03\x12A\n" +
	"<CATEGORY_WORKFORCE_ENGAGEMENT_WORKFORCE_MANAGEMENT_SCHEDULER\x10\xf4\x03B\xae\x02\n" +
	"&com.services.billing.entities.v1alpha4B\tTagsProtoP\x01ZNgithub.com/tcncloud/api-go/services/billing/entities/v1alpha4;entitiesv1alpha4\xa2\x02\x03SBE\xaa\x02\"Services.Billing.Entities.V1alpha4\xca\x02\"Services\\Billing\\Entities\\V1alpha4\xe2\x02.Services\\Billing\\Entities\\V1alpha4\\GPBMetadata\xea\x02%Services::Billing::Entities::V1alpha4b\x06proto3"

var (
	file_services_billing_entities_v1alpha4_tags_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha4_tags_proto_rawDescData []byte
)

func file_services_billing_entities_v1alpha4_tags_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha4_tags_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha4_tags_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha4_tags_proto_rawDesc), len(file_services_billing_entities_v1alpha4_tags_proto_rawDesc)))
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha4_tags_proto_rawDesc), len(file_services_billing_entities_v1alpha4_tags_proto_rawDesc)),
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
	file_services_billing_entities_v1alpha4_tags_proto_goTypes = nil
	file_services_billing_entities_v1alpha4_tags_proto_depIdxs = nil
}
