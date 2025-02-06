// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: services/billing/entities/v1alpha4/rates.proto

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

type RateDefinition struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RateDefinitionId string                 `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	SkuId            string                 `protobuf:"bytes,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/rates.proto.
	BillingTag    *BillingTag            `protobuf:"bytes,3,opt,name=billing_tag,json=billingTag,proto3" json:"billing_tag,omitempty"`
	Config        *ProductConfig         `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	IsDraft       bool                   `protobuf:"varint,5,opt,name=is_draft,json=isDraft,proto3" json:"is_draft,omitempty"`
	IsOverwrite   bool                   `protobuf:"varint,6,opt,name=is_overwrite,json=isOverwrite,proto3" json:"is_overwrite,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DeleteTime    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	BillingTagId  string                 `protobuf:"bytes,10,opt,name=billing_tag_id,json=billingTagId,proto3" json:"billing_tag_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RateDefinition) Reset() {
	*x = RateDefinition{}
	mi := &file_services_billing_entities_v1alpha4_rates_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RateDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateDefinition) ProtoMessage() {}

func (x *RateDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_rates_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateDefinition.ProtoReflect.Descriptor instead.
func (*RateDefinition) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha4_rates_proto_rawDescGZIP(), []int{0}
}

func (x *RateDefinition) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

func (x *RateDefinition) GetSkuId() string {
	if x != nil {
		return x.SkuId
	}
	return ""
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/rates.proto.
func (x *RateDefinition) GetBillingTag() *BillingTag {
	if x != nil {
		return x.BillingTag
	}
	return nil
}

func (x *RateDefinition) GetConfig() *ProductConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *RateDefinition) GetIsDraft() bool {
	if x != nil {
		return x.IsDraft
	}
	return false
}

func (x *RateDefinition) GetIsOverwrite() bool {
	if x != nil {
		return x.IsOverwrite
	}
	return false
}

func (x *RateDefinition) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *RateDefinition) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *RateDefinition) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *RateDefinition) GetBillingTagId() string {
	if x != nil {
		return x.BillingTagId
	}
	return ""
}

type MatchingRule struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	MatchingRuleId string                 `protobuf:"bytes,1,opt,name=matching_rule_id,json=matchingRuleId,proto3" json:"matching_rule_id,omitempty"`
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/rates.proto.
	Config     *ProductConfig         `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	RuleConfig *MatchingConfig        `protobuf:"bytes,6,opt,name=rule_config,json=ruleConfig,proto3" json:"rule_config,omitempty"`
	// Types that are valid to be assigned to MatchingConfig:
	//
	//	*MatchingRule_CountryCodePrefix
	MatchingConfig isMatchingRule_MatchingConfig `protobuf_oneof:"matching_config"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MatchingRule) Reset() {
	*x = MatchingRule{}
	mi := &file_services_billing_entities_v1alpha4_rates_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchingRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingRule) ProtoMessage() {}

func (x *MatchingRule) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_entities_v1alpha4_rates_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchingRule.ProtoReflect.Descriptor instead.
func (*MatchingRule) Descriptor() ([]byte, []int) {
	return file_services_billing_entities_v1alpha4_rates_proto_rawDescGZIP(), []int{1}
}

func (x *MatchingRule) GetMatchingRuleId() string {
	if x != nil {
		return x.MatchingRuleId
	}
	return ""
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/rates.proto.
func (x *MatchingRule) GetConfig() *ProductConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *MatchingRule) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *MatchingRule) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *MatchingRule) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *MatchingRule) GetRuleConfig() *MatchingConfig {
	if x != nil {
		return x.RuleConfig
	}
	return nil
}

func (x *MatchingRule) GetMatchingConfig() isMatchingRule_MatchingConfig {
	if x != nil {
		return x.MatchingConfig
	}
	return nil
}

// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/rates.proto.
func (x *MatchingRule) GetCountryCodePrefix() *CountryCodePrefix {
	if x != nil {
		if x, ok := x.MatchingConfig.(*MatchingRule_CountryCodePrefix); ok {
			return x.CountryCodePrefix
		}
	}
	return nil
}

type isMatchingRule_MatchingConfig interface {
	isMatchingRule_MatchingConfig()
}

type MatchingRule_CountryCodePrefix struct {
	// Deprecated: Marked as deprecated in services/billing/entities/v1alpha4/rates.proto.
	CountryCodePrefix *CountryCodePrefix `protobuf:"bytes,100,opt,name=country_code_prefix,json=countryCodePrefix,proto3,oneof"`
}

func (*MatchingRule_CountryCodePrefix) isMatchingRule_MatchingConfig() {}

var File_services_billing_entities_v1alpha4_rates_proto protoreflect.FileDescriptor

var file_services_billing_entities_v1alpha4_rates_proto_rawDesc = string([]byte{
	0x0a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x34, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x34, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2f,
	0x74, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x04, 0x0a, 0x0e, 0x52,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a,
	0x12, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73,
	0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x75,
	0x49, 0x64, 0x12, 0x53, 0x0a, 0x0b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2e, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x67, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x67, 0x12, 0x49, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x73, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x67, 0x49, 0x64, 0x22, 0x93, 0x04,
	0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x4d, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x53,
	0x0a, 0x0b, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x72, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x6b, 0x0a, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x34, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x02, 0x18, 0x01, 0x48, 0x00, 0x52, 0x11, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x42, 0x11, 0x0a, 0x0f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0xaf, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x42, 0x0a,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0x3b, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0xa2, 0x02, 0x03, 0x53,
	0x42, 0x45, 0xaa, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0xca, 0x02, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x34, 0xe2, 0x02, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_services_billing_entities_v1alpha4_rates_proto_rawDescOnce sync.Once
	file_services_billing_entities_v1alpha4_rates_proto_rawDescData []byte
)

func file_services_billing_entities_v1alpha4_rates_proto_rawDescGZIP() []byte {
	file_services_billing_entities_v1alpha4_rates_proto_rawDescOnce.Do(func() {
		file_services_billing_entities_v1alpha4_rates_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha4_rates_proto_rawDesc), len(file_services_billing_entities_v1alpha4_rates_proto_rawDesc)))
	})
	return file_services_billing_entities_v1alpha4_rates_proto_rawDescData
}

var file_services_billing_entities_v1alpha4_rates_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_billing_entities_v1alpha4_rates_proto_goTypes = []any{
	(*RateDefinition)(nil),        // 0: services.billing.entities.v1alpha4.RateDefinition
	(*MatchingRule)(nil),          // 1: services.billing.entities.v1alpha4.MatchingRule
	(*BillingTag)(nil),            // 2: services.billing.entities.v1alpha4.BillingTag
	(*ProductConfig)(nil),         // 3: services.billing.entities.v1alpha4.ProductConfig
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*MatchingConfig)(nil),        // 5: services.billing.entities.v1alpha4.MatchingConfig
	(*CountryCodePrefix)(nil),     // 6: services.billing.entities.v1alpha4.CountryCodePrefix
}
var file_services_billing_entities_v1alpha4_rates_proto_depIdxs = []int32{
	2,  // 0: services.billing.entities.v1alpha4.RateDefinition.billing_tag:type_name -> services.billing.entities.v1alpha4.BillingTag
	3,  // 1: services.billing.entities.v1alpha4.RateDefinition.config:type_name -> services.billing.entities.v1alpha4.ProductConfig
	4,  // 2: services.billing.entities.v1alpha4.RateDefinition.create_time:type_name -> google.protobuf.Timestamp
	4,  // 3: services.billing.entities.v1alpha4.RateDefinition.update_time:type_name -> google.protobuf.Timestamp
	4,  // 4: services.billing.entities.v1alpha4.RateDefinition.delete_time:type_name -> google.protobuf.Timestamp
	3,  // 5: services.billing.entities.v1alpha4.MatchingRule.config:type_name -> services.billing.entities.v1alpha4.ProductConfig
	4,  // 6: services.billing.entities.v1alpha4.MatchingRule.create_time:type_name -> google.protobuf.Timestamp
	4,  // 7: services.billing.entities.v1alpha4.MatchingRule.delete_time:type_name -> google.protobuf.Timestamp
	4,  // 8: services.billing.entities.v1alpha4.MatchingRule.update_time:type_name -> google.protobuf.Timestamp
	5,  // 9: services.billing.entities.v1alpha4.MatchingRule.rule_config:type_name -> services.billing.entities.v1alpha4.MatchingConfig
	6,  // 10: services.billing.entities.v1alpha4.MatchingRule.country_code_prefix:type_name -> services.billing.entities.v1alpha4.CountryCodePrefix
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_services_billing_entities_v1alpha4_rates_proto_init() }
func file_services_billing_entities_v1alpha4_rates_proto_init() {
	if File_services_billing_entities_v1alpha4_rates_proto != nil {
		return
	}
	file_services_billing_entities_v1alpha4_matching_proto_init()
	file_services_billing_entities_v1alpha4_products_proto_init()
	file_services_billing_entities_v1alpha4_tags_proto_init()
	file_services_billing_entities_v1alpha4_rates_proto_msgTypes[1].OneofWrappers = []any{
		(*MatchingRule_CountryCodePrefix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_entities_v1alpha4_rates_proto_rawDesc), len(file_services_billing_entities_v1alpha4_rates_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_entities_v1alpha4_rates_proto_goTypes,
		DependencyIndexes: file_services_billing_entities_v1alpha4_rates_proto_depIdxs,
		MessageInfos:      file_services_billing_entities_v1alpha4_rates_proto_msgTypes,
	}.Build()
	File_services_billing_entities_v1alpha4_rates_proto = out.File
	file_services_billing_entities_v1alpha4_rates_proto_goTypes = nil
	file_services_billing_entities_v1alpha4_rates_proto_depIdxs = nil
}
