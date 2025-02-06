// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: services/org/exile_certificate_manager/v1alpha1/exile_configuration.proto

package exile_certificate_managerv1alpha1

import (
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

// CreateExileConfigurationRequest is the request to create a exile configuration
type CreateExileConfigurationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The configuration name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The configuration description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The type of exile configuration
	Type ExileConfigurationType `protobuf:"varint,3,opt,name=type,proto3,enum=services.org.exile_certificate_manager.v1alpha1.ExileConfigurationType" json:"type,omitempty"`
	// The parameters of the exile configuration
	// See services.org.exile_certificate_manager.v1alpha1.ExileConfiguration
	// for the required information.
	Parameters    string `protobuf:"bytes,4,opt,name=parameters,proto3" json:"parameters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateExileConfigurationRequest) Reset() {
	*x = CreateExileConfigurationRequest{}
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateExileConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExileConfigurationRequest) ProtoMessage() {}

func (x *CreateExileConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExileConfigurationRequest.ProtoReflect.Descriptor instead.
func (*CreateExileConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *CreateExileConfigurationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateExileConfigurationRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateExileConfigurationRequest) GetType() ExileConfigurationType {
	if x != nil {
		return x.Type
	}
	return ExileConfigurationType_EXILE_CONFIGURATION_TYPE_UNSPECIFIED
}

func (x *CreateExileConfigurationRequest) GetParameters() string {
	if x != nil {
		return x.Parameters
	}
	return ""
}

// CreateExileConfigurationResponse is the response to creating a exile configuration
type CreateExileConfigurationResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the exile configuration that was created.
	// DEPRECATED: Use exile_configuration instead.
	//
	// Deprecated: Marked as deprecated in services/org/exile_certificate_manager/v1alpha1/exile_configuration.proto.
	ExileConfigurationId string `protobuf:"bytes,1,opt,name=exile_configuration_id,json=exileConfigurationId,proto3" json:"exile_configuration_id,omitempty"`
	// The newly created exile configuration.
	ExileConfiguration *ExileConfiguration `protobuf:"bytes,2,opt,name=exile_configuration,json=exileConfiguration,proto3" json:"exile_configuration,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *CreateExileConfigurationResponse) Reset() {
	*x = CreateExileConfigurationResponse{}
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateExileConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExileConfigurationResponse) ProtoMessage() {}

func (x *CreateExileConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExileConfigurationResponse.ProtoReflect.Descriptor instead.
func (*CreateExileConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in services/org/exile_certificate_manager/v1alpha1/exile_configuration.proto.
func (x *CreateExileConfigurationResponse) GetExileConfigurationId() string {
	if x != nil {
		return x.ExileConfigurationId
	}
	return ""
}

func (x *CreateExileConfigurationResponse) GetExileConfiguration() *ExileConfiguration {
	if x != nil {
		return x.ExileConfiguration
	}
	return nil
}

// UpdateExileConfigurationRequest is the request to update a exile configuration
type UpdateExileConfigurationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The exile configuration to update.
	Configuration *ExileConfiguration `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// The fields to update.
	FieldMask     *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateExileConfigurationRequest) Reset() {
	*x = UpdateExileConfigurationRequest{}
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateExileConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExileConfigurationRequest) ProtoMessage() {}

func (x *UpdateExileConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExileConfigurationRequest.ProtoReflect.Descriptor instead.
func (*UpdateExileConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateExileConfigurationRequest) GetConfiguration() *ExileConfiguration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *UpdateExileConfigurationRequest) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

// UpdateExileConfigurationResponse is the response to updating a exile configuration
type UpdateExileConfigurationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateExileConfigurationResponse) Reset() {
	*x = UpdateExileConfigurationResponse{}
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateExileConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExileConfigurationResponse) ProtoMessage() {}

func (x *UpdateExileConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExileConfigurationResponse.ProtoReflect.Descriptor instead.
func (*UpdateExileConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescGZIP(), []int{3}
}

// DeleteExileConfigurationRequest is the request to delete a exile configuration
type DeleteExileConfigurationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the exile configuration to delete.
	ExileConfigurationId string `protobuf:"bytes,1,opt,name=exile_configuration_id,json=exileConfigurationId,proto3" json:"exile_configuration_id,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *DeleteExileConfigurationRequest) Reset() {
	*x = DeleteExileConfigurationRequest{}
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteExileConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExileConfigurationRequest) ProtoMessage() {}

func (x *DeleteExileConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExileConfigurationRequest.ProtoReflect.Descriptor instead.
func (*DeleteExileConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteExileConfigurationRequest) GetExileConfigurationId() string {
	if x != nil {
		return x.ExileConfigurationId
	}
	return ""
}

// DeleteExileConfigurationResponse is the response to deleting a exile configuration
type DeleteExileConfigurationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteExileConfigurationResponse) Reset() {
	*x = DeleteExileConfigurationResponse{}
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteExileConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExileConfigurationResponse) ProtoMessage() {}

func (x *DeleteExileConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExileConfigurationResponse.ProtoReflect.Descriptor instead.
func (*DeleteExileConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescGZIP(), []int{5}
}

// ListExileConfigurationsRequest is the request to list exile configurations.
type ListExileConfigurationsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The fields on the entity to include in the response.
	FieldMask     *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListExileConfigurationsRequest) Reset() {
	*x = ListExileConfigurationsRequest{}
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListExileConfigurationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExileConfigurationsRequest) ProtoMessage() {}

func (x *ListExileConfigurationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExileConfigurationsRequest.ProtoReflect.Descriptor instead.
func (*ListExileConfigurationsRequest) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescGZIP(), []int{6}
}

func (x *ListExileConfigurationsRequest) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

// ListExileConfigurationsResponse is the response to listing exile configurations.
type ListExileConfigurationsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The list of exile configurations.
	ExileConfigurations []*ExileConfiguration `protobuf:"bytes,1,rep,name=exile_configurations,json=exileConfigurations,proto3" json:"exile_configurations,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ListExileConfigurationsResponse) Reset() {
	*x = ListExileConfigurationsResponse{}
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListExileConfigurationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExileConfigurationsResponse) ProtoMessage() {}

func (x *ListExileConfigurationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExileConfigurationsResponse.ProtoReflect.Descriptor instead.
func (*ListExileConfigurationsResponse) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescGZIP(), []int{7}
}

func (x *ListExileConfigurationsResponse) GetExileConfigurations() []*ExileConfiguration {
	if x != nil {
		return x.ExileConfigurations
	}
	return nil
}

var File_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto protoreflect.FileDescriptor

var file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDesc = string([]byte{
	0x0a, 0x49, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65,
	0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78, 0x69,
	0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4,
	0x01, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x47, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x16, 0x65, 0x78,
	0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x14,
	0x65, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x74, 0x0a, 0x13, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x43, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc7, 0x01, 0x0a, 0x1f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x69,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x22, 0x22, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x65,
	0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65, 0x78, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x22, 0x0a, 0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x22, 0x99, 0x01, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x69, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x14, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x65, 0x78, 0x69, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x93,
	0x03, 0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x17, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65, 0x78,
	0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x53, 0x4f, 0x45, 0xaa, 0x02, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4f, 0x72, 0x67, 0x2e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x39, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x30, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x4f, 0x72,
	0x67, 0x3a, 0x3a, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescOnce sync.Once
	file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescData []byte
)

func file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescGZIP() []byte {
	file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescOnce.Do(func() {
		file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDesc), len(file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDesc)))
	})
	return file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDescData
}

var file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_goTypes = []any{
	(*CreateExileConfigurationRequest)(nil),  // 0: services.org.exile_certificate_manager.v1alpha1.CreateExileConfigurationRequest
	(*CreateExileConfigurationResponse)(nil), // 1: services.org.exile_certificate_manager.v1alpha1.CreateExileConfigurationResponse
	(*UpdateExileConfigurationRequest)(nil),  // 2: services.org.exile_certificate_manager.v1alpha1.UpdateExileConfigurationRequest
	(*UpdateExileConfigurationResponse)(nil), // 3: services.org.exile_certificate_manager.v1alpha1.UpdateExileConfigurationResponse
	(*DeleteExileConfigurationRequest)(nil),  // 4: services.org.exile_certificate_manager.v1alpha1.DeleteExileConfigurationRequest
	(*DeleteExileConfigurationResponse)(nil), // 5: services.org.exile_certificate_manager.v1alpha1.DeleteExileConfigurationResponse
	(*ListExileConfigurationsRequest)(nil),   // 6: services.org.exile_certificate_manager.v1alpha1.ListExileConfigurationsRequest
	(*ListExileConfigurationsResponse)(nil),  // 7: services.org.exile_certificate_manager.v1alpha1.ListExileConfigurationsResponse
	(ExileConfigurationType)(0),              // 8: services.org.exile_certificate_manager.v1alpha1.ExileConfigurationType
	(*ExileConfiguration)(nil),               // 9: services.org.exile_certificate_manager.v1alpha1.ExileConfiguration
	(*fieldmaskpb.FieldMask)(nil),            // 10: google.protobuf.FieldMask
}
var file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_depIdxs = []int32{
	8,  // 0: services.org.exile_certificate_manager.v1alpha1.CreateExileConfigurationRequest.type:type_name -> services.org.exile_certificate_manager.v1alpha1.ExileConfigurationType
	9,  // 1: services.org.exile_certificate_manager.v1alpha1.CreateExileConfigurationResponse.exile_configuration:type_name -> services.org.exile_certificate_manager.v1alpha1.ExileConfiguration
	9,  // 2: services.org.exile_certificate_manager.v1alpha1.UpdateExileConfigurationRequest.configuration:type_name -> services.org.exile_certificate_manager.v1alpha1.ExileConfiguration
	10, // 3: services.org.exile_certificate_manager.v1alpha1.UpdateExileConfigurationRequest.field_mask:type_name -> google.protobuf.FieldMask
	10, // 4: services.org.exile_certificate_manager.v1alpha1.ListExileConfigurationsRequest.field_mask:type_name -> google.protobuf.FieldMask
	9,  // 5: services.org.exile_certificate_manager.v1alpha1.ListExileConfigurationsResponse.exile_configurations:type_name -> services.org.exile_certificate_manager.v1alpha1.ExileConfiguration
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_init() }
func file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_init() {
	if File_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto != nil {
		return
	}
	file_services_org_exile_certificate_manager_v1alpha1_entities_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDesc), len(file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_goTypes,
		DependencyIndexes: file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_depIdxs,
		MessageInfos:      file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_msgTypes,
	}.Build()
	File_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto = out.File
	file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_goTypes = nil
	file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_depIdxs = nil
}
