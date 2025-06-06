// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/billing/v1alpha1/rates.proto

package billingv1alpha1

import (
	v1alpha1 "github.com/tcncloud/api-go/services/billing/entities/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

type CreateDefaultRateDefinitionRequest struct {
	state            protoimpl.MessageState   `protogen:"open.v1"`
	RateDefinitionId string                   `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	RateDefinition   *v1alpha1.RateDefinition `protobuf:"bytes,2,opt,name=rate_definition,json=rateDefinition,proto3" json:"rate_definition,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateDefaultRateDefinitionRequest) Reset() {
	*x = CreateDefaultRateDefinitionRequest{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDefaultRateDefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDefaultRateDefinitionRequest) ProtoMessage() {}

func (x *CreateDefaultRateDefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDefaultRateDefinitionRequest.ProtoReflect.Descriptor instead.
func (*CreateDefaultRateDefinitionRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDefaultRateDefinitionRequest) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

func (x *CreateDefaultRateDefinitionRequest) GetRateDefinition() *v1alpha1.RateDefinition {
	if x != nil {
		return x.RateDefinition
	}
	return nil
}

type CreateDefaultRateDefinitionResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RateDefinitionId string                 `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateDefaultRateDefinitionResponse) Reset() {
	*x = CreateDefaultRateDefinitionResponse{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDefaultRateDefinitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDefaultRateDefinitionResponse) ProtoMessage() {}

func (x *CreateDefaultRateDefinitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDefaultRateDefinitionResponse.ProtoReflect.Descriptor instead.
func (*CreateDefaultRateDefinitionResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDefaultRateDefinitionResponse) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

type CreateRateDefinitionRequest struct {
	state            protoimpl.MessageState   `protogen:"open.v1"`
	RateDefinitionId string                   `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	RateDefinition   *v1alpha1.RateDefinition `protobuf:"bytes,2,opt,name=rate_definition,json=rateDefinition,proto3" json:"rate_definition,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateRateDefinitionRequest) Reset() {
	*x = CreateRateDefinitionRequest{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRateDefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRateDefinitionRequest) ProtoMessage() {}

func (x *CreateRateDefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRateDefinitionRequest.ProtoReflect.Descriptor instead.
func (*CreateRateDefinitionRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRateDefinitionRequest) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

func (x *CreateRateDefinitionRequest) GetRateDefinition() *v1alpha1.RateDefinition {
	if x != nil {
		return x.RateDefinition
	}
	return nil
}

type CreateRateDefinitionResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RateDefinitionId string                 `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateRateDefinitionResponse) Reset() {
	*x = CreateRateDefinitionResponse{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRateDefinitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRateDefinitionResponse) ProtoMessage() {}

func (x *CreateRateDefinitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRateDefinitionResponse.ProtoReflect.Descriptor instead.
func (*CreateRateDefinitionResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRateDefinitionResponse) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

type DeleteDefaultRateDefinitionRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RateDefinitionId string                 `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *DeleteDefaultRateDefinitionRequest) Reset() {
	*x = DeleteDefaultRateDefinitionRequest{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDefaultRateDefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDefaultRateDefinitionRequest) ProtoMessage() {}

func (x *DeleteDefaultRateDefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDefaultRateDefinitionRequest.ProtoReflect.Descriptor instead.
func (*DeleteDefaultRateDefinitionRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDefaultRateDefinitionRequest) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

type DeleteDefaultRateDefinitionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDefaultRateDefinitionResponse) Reset() {
	*x = DeleteDefaultRateDefinitionResponse{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDefaultRateDefinitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDefaultRateDefinitionResponse) ProtoMessage() {}

func (x *DeleteDefaultRateDefinitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDefaultRateDefinitionResponse.ProtoReflect.Descriptor instead.
func (*DeleteDefaultRateDefinitionResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{5}
}

type DeleteRateDefinitionRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RateDefinitionId string                 `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *DeleteRateDefinitionRequest) Reset() {
	*x = DeleteRateDefinitionRequest{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRateDefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRateDefinitionRequest) ProtoMessage() {}

func (x *DeleteRateDefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRateDefinitionRequest.ProtoReflect.Descriptor instead.
func (*DeleteRateDefinitionRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRateDefinitionRequest) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

type DeleteRateDefinitionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRateDefinitionResponse) Reset() {
	*x = DeleteRateDefinitionResponse{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRateDefinitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRateDefinitionResponse) ProtoMessage() {}

func (x *DeleteRateDefinitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRateDefinitionResponse.ProtoReflect.Descriptor instead.
func (*DeleteRateDefinitionResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{7}
}

type GetRateDefinitionRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RateDefinitionId string                 `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetRateDefinitionRequest) Reset() {
	*x = GetRateDefinitionRequest{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRateDefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRateDefinitionRequest) ProtoMessage() {}

func (x *GetRateDefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRateDefinitionRequest.ProtoReflect.Descriptor instead.
func (*GetRateDefinitionRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{8}
}

func (x *GetRateDefinitionRequest) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

type GetRateDefinitionResponse struct {
	state          protoimpl.MessageState   `protogen:"open.v1"`
	RateDefinition *v1alpha1.RateDefinition `protobuf:"bytes,1,opt,name=rate_definition,json=rateDefinition,proto3" json:"rate_definition,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetRateDefinitionResponse) Reset() {
	*x = GetRateDefinitionResponse{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRateDefinitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRateDefinitionResponse) ProtoMessage() {}

func (x *GetRateDefinitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRateDefinitionResponse.ProtoReflect.Descriptor instead.
func (*GetRateDefinitionResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{9}
}

func (x *GetRateDefinitionResponse) GetRateDefinition() *v1alpha1.RateDefinition {
	if x != nil {
		return x.RateDefinition
	}
	return nil
}

type ListRateDefinitionsRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RateDefinitionId string                 `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	Filter           string                 `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"` // Optional: defaults to no filter
	Fields           *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=fields,proto3" json:"fields,omitempty"` // Optional: defaults to all fields.
	Sort             []*Sort                `protobuf:"bytes,4,rep,name=sort,proto3" json:"sort,omitempty"`     // Optional: defaults to no sort.
	Page             *Page                  `protobuf:"bytes,5,opt,name=page,proto3" json:"page,omitempty"`     // Optional: defaults to no paging.
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ListRateDefinitionsRequest) Reset() {
	*x = ListRateDefinitionsRequest{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRateDefinitionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRateDefinitionsRequest) ProtoMessage() {}

func (x *ListRateDefinitionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRateDefinitionsRequest.ProtoReflect.Descriptor instead.
func (*ListRateDefinitionsRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{10}
}

func (x *ListRateDefinitionsRequest) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

func (x *ListRateDefinitionsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListRateDefinitionsRequest) GetFields() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ListRateDefinitionsRequest) GetSort() []*Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *ListRateDefinitionsRequest) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListRateDefinitionsResponse struct {
	state           protoimpl.MessageState     `protogen:"open.v1"`
	RateDefinitions []*v1alpha1.RateDefinition `protobuf:"bytes,1,rep,name=rate_definitions,json=rateDefinitions,proto3" json:"rate_definitions,omitempty"`
	Token           string                     `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"` // Optional: only present if paginating.
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListRateDefinitionsResponse) Reset() {
	*x = ListRateDefinitionsResponse{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRateDefinitionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRateDefinitionsResponse) ProtoMessage() {}

func (x *ListRateDefinitionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRateDefinitionsResponse.ProtoReflect.Descriptor instead.
func (*ListRateDefinitionsResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{11}
}

func (x *ListRateDefinitionsResponse) GetRateDefinitions() []*v1alpha1.RateDefinition {
	if x != nil {
		return x.RateDefinitions
	}
	return nil
}

func (x *ListRateDefinitionsResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateDefaultRateDefinitionRequest struct {
	state            protoimpl.MessageState   `protogen:"open.v1"`
	RateDefinitionId string                   `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	RateDefinition   *v1alpha1.RateDefinition `protobuf:"bytes,2,opt,name=rate_definition,json=rateDefinition,proto3" json:"rate_definition,omitempty"`
	UpdateFields     *fieldmaskpb.FieldMask   `protobuf:"bytes,3,opt,name=update_fields,json=updateFields,proto3" json:"update_fields,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *UpdateDefaultRateDefinitionRequest) Reset() {
	*x = UpdateDefaultRateDefinitionRequest{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDefaultRateDefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDefaultRateDefinitionRequest) ProtoMessage() {}

func (x *UpdateDefaultRateDefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDefaultRateDefinitionRequest.ProtoReflect.Descriptor instead.
func (*UpdateDefaultRateDefinitionRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateDefaultRateDefinitionRequest) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

func (x *UpdateDefaultRateDefinitionRequest) GetRateDefinition() *v1alpha1.RateDefinition {
	if x != nil {
		return x.RateDefinition
	}
	return nil
}

func (x *UpdateDefaultRateDefinitionRequest) GetUpdateFields() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateFields
	}
	return nil
}

type UpdateDefaultRateDefinitionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDefaultRateDefinitionResponse) Reset() {
	*x = UpdateDefaultRateDefinitionResponse{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDefaultRateDefinitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDefaultRateDefinitionResponse) ProtoMessage() {}

func (x *UpdateDefaultRateDefinitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDefaultRateDefinitionResponse.ProtoReflect.Descriptor instead.
func (*UpdateDefaultRateDefinitionResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{13}
}

type UpdateRateDefinitionRequest struct {
	state            protoimpl.MessageState   `protogen:"open.v1"`
	RateDefinitionId string                   `protobuf:"bytes,1,opt,name=rate_definition_id,json=rateDefinitionId,proto3" json:"rate_definition_id,omitempty"`
	RateDefinition   *v1alpha1.RateDefinition `protobuf:"bytes,2,opt,name=rate_definition,json=rateDefinition,proto3" json:"rate_definition,omitempty"`
	UpdateFields     *fieldmaskpb.FieldMask   `protobuf:"bytes,3,opt,name=update_fields,json=updateFields,proto3" json:"update_fields,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *UpdateRateDefinitionRequest) Reset() {
	*x = UpdateRateDefinitionRequest{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRateDefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRateDefinitionRequest) ProtoMessage() {}

func (x *UpdateRateDefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRateDefinitionRequest.ProtoReflect.Descriptor instead.
func (*UpdateRateDefinitionRequest) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{14}
}

func (x *UpdateRateDefinitionRequest) GetRateDefinitionId() string {
	if x != nil {
		return x.RateDefinitionId
	}
	return ""
}

func (x *UpdateRateDefinitionRequest) GetRateDefinition() *v1alpha1.RateDefinition {
	if x != nil {
		return x.RateDefinition
	}
	return nil
}

func (x *UpdateRateDefinitionRequest) GetUpdateFields() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateFields
	}
	return nil
}

type UpdateRateDefinitionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRateDefinitionResponse) Reset() {
	*x = UpdateRateDefinitionResponse{}
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRateDefinitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRateDefinitionResponse) ProtoMessage() {}

func (x *UpdateRateDefinitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_billing_v1alpha1_rates_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRateDefinitionResponse.ProtoReflect.Descriptor instead.
func (*UpdateRateDefinitionResponse) Descriptor() ([]byte, []int) {
	return file_services_billing_v1alpha1_rates_proto_rawDescGZIP(), []int{15}
}

var File_services_billing_v1alpha1_rates_proto protoreflect.FileDescriptor

const file_services_billing_v1alpha1_rates_proto_rawDesc = "" +
	"\n" +
	"%services/billing/v1alpha1/rates.proto\x12\x19services.billing.v1alpha1\x1a google/protobuf/field_mask.proto\x1a.services/billing/entities/v1alpha1/rates.proto\x1a$services/billing/v1alpha1/core.proto\"\xaf\x01\n" +
	"\"CreateDefaultRateDefinitionRequest\x12,\n" +
	"\x12rate_definition_id\x18\x01 \x01(\tR\x10rateDefinitionId\x12[\n" +
	"\x0frate_definition\x18\x02 \x01(\v22.services.billing.entities.v1alpha1.RateDefinitionR\x0erateDefinition\"S\n" +
	"#CreateDefaultRateDefinitionResponse\x12,\n" +
	"\x12rate_definition_id\x18\x01 \x01(\tR\x10rateDefinitionId\"\xa8\x01\n" +
	"\x1bCreateRateDefinitionRequest\x12,\n" +
	"\x12rate_definition_id\x18\x01 \x01(\tR\x10rateDefinitionId\x12[\n" +
	"\x0frate_definition\x18\x02 \x01(\v22.services.billing.entities.v1alpha1.RateDefinitionR\x0erateDefinition\"L\n" +
	"\x1cCreateRateDefinitionResponse\x12,\n" +
	"\x12rate_definition_id\x18\x01 \x01(\tR\x10rateDefinitionId\"R\n" +
	"\"DeleteDefaultRateDefinitionRequest\x12,\n" +
	"\x12rate_definition_id\x18\x01 \x01(\tR\x10rateDefinitionId\"%\n" +
	"#DeleteDefaultRateDefinitionResponse\"K\n" +
	"\x1bDeleteRateDefinitionRequest\x12,\n" +
	"\x12rate_definition_id\x18\x01 \x01(\tR\x10rateDefinitionId\"\x1e\n" +
	"\x1cDeleteRateDefinitionResponse\"H\n" +
	"\x18GetRateDefinitionRequest\x12,\n" +
	"\x12rate_definition_id\x18\x01 \x01(\tR\x10rateDefinitionId\"x\n" +
	"\x19GetRateDefinitionResponse\x12[\n" +
	"\x0frate_definition\x18\x01 \x01(\v22.services.billing.entities.v1alpha1.RateDefinitionR\x0erateDefinition\"\x80\x02\n" +
	"\x1aListRateDefinitionsRequest\x12,\n" +
	"\x12rate_definition_id\x18\x01 \x01(\tR\x10rateDefinitionId\x12\x16\n" +
	"\x06filter\x18\x02 \x01(\tR\x06filter\x122\n" +
	"\x06fields\x18\x03 \x01(\v2\x1a.google.protobuf.FieldMaskR\x06fields\x123\n" +
	"\x04sort\x18\x04 \x03(\v2\x1f.services.billing.v1alpha1.SortR\x04sort\x123\n" +
	"\x04page\x18\x05 \x01(\v2\x1f.services.billing.v1alpha1.PageR\x04page\"\x92\x01\n" +
	"\x1bListRateDefinitionsResponse\x12]\n" +
	"\x10rate_definitions\x18\x01 \x03(\v22.services.billing.entities.v1alpha1.RateDefinitionR\x0frateDefinitions\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\"\xf0\x01\n" +
	"\"UpdateDefaultRateDefinitionRequest\x12,\n" +
	"\x12rate_definition_id\x18\x01 \x01(\tR\x10rateDefinitionId\x12[\n" +
	"\x0frate_definition\x18\x02 \x01(\v22.services.billing.entities.v1alpha1.RateDefinitionR\x0erateDefinition\x12?\n" +
	"\rupdate_fields\x18\x03 \x01(\v2\x1a.google.protobuf.FieldMaskR\fupdateFields\"%\n" +
	"#UpdateDefaultRateDefinitionResponse\"\xe9\x01\n" +
	"\x1bUpdateRateDefinitionRequest\x12,\n" +
	"\x12rate_definition_id\x18\x01 \x01(\tR\x10rateDefinitionId\x12[\n" +
	"\x0frate_definition\x18\x02 \x01(\v22.services.billing.entities.v1alpha1.RateDefinitionR\x0erateDefinition\x12?\n" +
	"\rupdate_fields\x18\x03 \x01(\v2\x1a.google.protobuf.FieldMaskR\fupdateFields\"\x1e\n" +
	"\x1cUpdateRateDefinitionResponseB\xf7\x01\n" +
	"\x1dcom.services.billing.v1alpha1B\n" +
	"RatesProtoP\x01ZDgithub.com/tcncloud/api-go/services/billing/v1alpha1;billingv1alpha1\xa2\x02\x03SBX\xaa\x02\x19Services.Billing.V1alpha1\xca\x02\x19Services\\Billing\\V1alpha1\xe2\x02%Services\\Billing\\V1alpha1\\GPBMetadata\xea\x02\x1bServices::Billing::V1alpha1b\x06proto3"

var (
	file_services_billing_v1alpha1_rates_proto_rawDescOnce sync.Once
	file_services_billing_v1alpha1_rates_proto_rawDescData []byte
)

func file_services_billing_v1alpha1_rates_proto_rawDescGZIP() []byte {
	file_services_billing_v1alpha1_rates_proto_rawDescOnce.Do(func() {
		file_services_billing_v1alpha1_rates_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_billing_v1alpha1_rates_proto_rawDesc), len(file_services_billing_v1alpha1_rates_proto_rawDesc)))
	})
	return file_services_billing_v1alpha1_rates_proto_rawDescData
}

var file_services_billing_v1alpha1_rates_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_services_billing_v1alpha1_rates_proto_goTypes = []any{
	(*CreateDefaultRateDefinitionRequest)(nil),  // 0: services.billing.v1alpha1.CreateDefaultRateDefinitionRequest
	(*CreateDefaultRateDefinitionResponse)(nil), // 1: services.billing.v1alpha1.CreateDefaultRateDefinitionResponse
	(*CreateRateDefinitionRequest)(nil),         // 2: services.billing.v1alpha1.CreateRateDefinitionRequest
	(*CreateRateDefinitionResponse)(nil),        // 3: services.billing.v1alpha1.CreateRateDefinitionResponse
	(*DeleteDefaultRateDefinitionRequest)(nil),  // 4: services.billing.v1alpha1.DeleteDefaultRateDefinitionRequest
	(*DeleteDefaultRateDefinitionResponse)(nil), // 5: services.billing.v1alpha1.DeleteDefaultRateDefinitionResponse
	(*DeleteRateDefinitionRequest)(nil),         // 6: services.billing.v1alpha1.DeleteRateDefinitionRequest
	(*DeleteRateDefinitionResponse)(nil),        // 7: services.billing.v1alpha1.DeleteRateDefinitionResponse
	(*GetRateDefinitionRequest)(nil),            // 8: services.billing.v1alpha1.GetRateDefinitionRequest
	(*GetRateDefinitionResponse)(nil),           // 9: services.billing.v1alpha1.GetRateDefinitionResponse
	(*ListRateDefinitionsRequest)(nil),          // 10: services.billing.v1alpha1.ListRateDefinitionsRequest
	(*ListRateDefinitionsResponse)(nil),         // 11: services.billing.v1alpha1.ListRateDefinitionsResponse
	(*UpdateDefaultRateDefinitionRequest)(nil),  // 12: services.billing.v1alpha1.UpdateDefaultRateDefinitionRequest
	(*UpdateDefaultRateDefinitionResponse)(nil), // 13: services.billing.v1alpha1.UpdateDefaultRateDefinitionResponse
	(*UpdateRateDefinitionRequest)(nil),         // 14: services.billing.v1alpha1.UpdateRateDefinitionRequest
	(*UpdateRateDefinitionResponse)(nil),        // 15: services.billing.v1alpha1.UpdateRateDefinitionResponse
	(*v1alpha1.RateDefinition)(nil),             // 16: services.billing.entities.v1alpha1.RateDefinition
	(*fieldmaskpb.FieldMask)(nil),               // 17: google.protobuf.FieldMask
	(*Sort)(nil),                                // 18: services.billing.v1alpha1.Sort
	(*Page)(nil),                                // 19: services.billing.v1alpha1.Page
}
var file_services_billing_v1alpha1_rates_proto_depIdxs = []int32{
	16, // 0: services.billing.v1alpha1.CreateDefaultRateDefinitionRequest.rate_definition:type_name -> services.billing.entities.v1alpha1.RateDefinition
	16, // 1: services.billing.v1alpha1.CreateRateDefinitionRequest.rate_definition:type_name -> services.billing.entities.v1alpha1.RateDefinition
	16, // 2: services.billing.v1alpha1.GetRateDefinitionResponse.rate_definition:type_name -> services.billing.entities.v1alpha1.RateDefinition
	17, // 3: services.billing.v1alpha1.ListRateDefinitionsRequest.fields:type_name -> google.protobuf.FieldMask
	18, // 4: services.billing.v1alpha1.ListRateDefinitionsRequest.sort:type_name -> services.billing.v1alpha1.Sort
	19, // 5: services.billing.v1alpha1.ListRateDefinitionsRequest.page:type_name -> services.billing.v1alpha1.Page
	16, // 6: services.billing.v1alpha1.ListRateDefinitionsResponse.rate_definitions:type_name -> services.billing.entities.v1alpha1.RateDefinition
	16, // 7: services.billing.v1alpha1.UpdateDefaultRateDefinitionRequest.rate_definition:type_name -> services.billing.entities.v1alpha1.RateDefinition
	17, // 8: services.billing.v1alpha1.UpdateDefaultRateDefinitionRequest.update_fields:type_name -> google.protobuf.FieldMask
	16, // 9: services.billing.v1alpha1.UpdateRateDefinitionRequest.rate_definition:type_name -> services.billing.entities.v1alpha1.RateDefinition
	17, // 10: services.billing.v1alpha1.UpdateRateDefinitionRequest.update_fields:type_name -> google.protobuf.FieldMask
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_services_billing_v1alpha1_rates_proto_init() }
func file_services_billing_v1alpha1_rates_proto_init() {
	if File_services_billing_v1alpha1_rates_proto != nil {
		return
	}
	file_services_billing_v1alpha1_core_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_billing_v1alpha1_rates_proto_rawDesc), len(file_services_billing_v1alpha1_rates_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_billing_v1alpha1_rates_proto_goTypes,
		DependencyIndexes: file_services_billing_v1alpha1_rates_proto_depIdxs,
		MessageInfos:      file_services_billing_v1alpha1_rates_proto_msgTypes,
	}.Build()
	File_services_billing_v1alpha1_rates_proto = out.File
	file_services_billing_v1alpha1_rates_proto_goTypes = nil
	file_services_billing_v1alpha1_rates_proto_depIdxs = nil
}
