// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: services/org/exile_certificate_manager/v1alpha1/exile_certificate.proto

package exile_certificate_managerv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CreateExileCertificateRequest is the request for creating a exile certificate.
type CreateExileCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the exile certificate to be created.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the exile certificate to be created.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateExileCertificateRequest) Reset() {
	*x = CreateExileCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExileCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExileCertificateRequest) ProtoMessage() {}

func (x *CreateExileCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExileCertificateRequest.ProtoReflect.Descriptor instead.
func (*CreateExileCertificateRequest) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *CreateExileCertificateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateExileCertificateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// CreateExileCertificateResponse is the response for creating a exile certificate.
type CreateExileCertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The base-64 encoded certificate that was created.
	EncodedExileCertificate string `protobuf:"bytes,1,opt,name=encoded_exile_certificate,json=encodedExileCertificate,proto3" json:"encoded_exile_certificate,omitempty"`
}

func (x *CreateExileCertificateResponse) Reset() {
	*x = CreateExileCertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExileCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExileCertificateResponse) ProtoMessage() {}

func (x *CreateExileCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExileCertificateResponse.ProtoReflect.Descriptor instead.
func (*CreateExileCertificateResponse) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescGZIP(), []int{1}
}

func (x *CreateExileCertificateResponse) GetEncodedExileCertificate() string {
	if x != nil {
		return x.EncodedExileCertificate
	}
	return ""
}

// RevokeExileCertificateRequest is the request for revoking a exile certificate.
type RevokeExileCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the exile certificate to delete.
	ExileCertificateId string `protobuf:"bytes,1,opt,name=exile_certificate_id,json=exileCertificateId,proto3" json:"exile_certificate_id,omitempty"`
}

func (x *RevokeExileCertificateRequest) Reset() {
	*x = RevokeExileCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeExileCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeExileCertificateRequest) ProtoMessage() {}

func (x *RevokeExileCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeExileCertificateRequest.ProtoReflect.Descriptor instead.
func (*RevokeExileCertificateRequest) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescGZIP(), []int{2}
}

func (x *RevokeExileCertificateRequest) GetExileCertificateId() string {
	if x != nil {
		return x.ExileCertificateId
	}
	return ""
}

// RevokeExileCertificateResponse is the response for revoking a exile certificate.
type RevokeExileCertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RevokeExileCertificateResponse) Reset() {
	*x = RevokeExileCertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeExileCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeExileCertificateResponse) ProtoMessage() {}

func (x *RevokeExileCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeExileCertificateResponse.ProtoReflect.Descriptor instead.
func (*RevokeExileCertificateResponse) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescGZIP(), []int{3}
}

// AssignExileConfigurationRequest is the request for assigning a exile configuration
type AssignExileConfigurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the exile certificate to update.
	ExileCertificateId string `protobuf:"bytes,1,opt,name=exile_certificate_id,json=exileCertificateId,proto3" json:"exile_certificate_id,omitempty"`
	// The id of the exile configuration to assign.
	ExileConfigurationId string `protobuf:"bytes,2,opt,name=exile_configuration_id,json=exileConfigurationId,proto3" json:"exile_configuration_id,omitempty"`
}

func (x *AssignExileConfigurationRequest) Reset() {
	*x = AssignExileConfigurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignExileConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignExileConfigurationRequest) ProtoMessage() {}

func (x *AssignExileConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignExileConfigurationRequest.ProtoReflect.Descriptor instead.
func (*AssignExileConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescGZIP(), []int{4}
}

func (x *AssignExileConfigurationRequest) GetExileCertificateId() string {
	if x != nil {
		return x.ExileCertificateId
	}
	return ""
}

func (x *AssignExileConfigurationRequest) GetExileConfigurationId() string {
	if x != nil {
		return x.ExileConfigurationId
	}
	return ""
}

// AssignExileConfigurationResponse is the response for assigning a exile configuration
type AssignExileConfigurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssignExileConfigurationResponse) Reset() {
	*x = AssignExileConfigurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignExileConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignExileConfigurationResponse) ProtoMessage() {}

func (x *AssignExileConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignExileConfigurationResponse.ProtoReflect.Descriptor instead.
func (*AssignExileConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescGZIP(), []int{5}
}

// UnassignExileConfigurationRequest is the request for unassigning a exile configuration
type UnassignExileConfigurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the exile certificate to update.
	ExileCertificateId string `protobuf:"bytes,1,opt,name=exile_certificate_id,json=exileCertificateId,proto3" json:"exile_certificate_id,omitempty"`
}

func (x *UnassignExileConfigurationRequest) Reset() {
	*x = UnassignExileConfigurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnassignExileConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnassignExileConfigurationRequest) ProtoMessage() {}

func (x *UnassignExileConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnassignExileConfigurationRequest.ProtoReflect.Descriptor instead.
func (*UnassignExileConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescGZIP(), []int{6}
}

func (x *UnassignExileConfigurationRequest) GetExileCertificateId() string {
	if x != nil {
		return x.ExileCertificateId
	}
	return ""
}

// UnassignExileConfigurationResponse is the response for unassigning a exile configuration
type UnassignExileConfigurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnassignExileConfigurationResponse) Reset() {
	*x = UnassignExileConfigurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnassignExileConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnassignExileConfigurationResponse) ProtoMessage() {}

func (x *UnassignExileConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnassignExileConfigurationResponse.ProtoReflect.Descriptor instead.
func (*UnassignExileConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescGZIP(), []int{7}
}

// ListExileCertificatesRequest is the request for listing exile certificate.
type ListExileCertificatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The fields on the entity to include in the response.
	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
}

func (x *ListExileCertificatesRequest) Reset() {
	*x = ListExileCertificatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExileCertificatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExileCertificatesRequest) ProtoMessage() {}

func (x *ListExileCertificatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExileCertificatesRequest.ProtoReflect.Descriptor instead.
func (*ListExileCertificatesRequest) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescGZIP(), []int{8}
}

func (x *ListExileCertificatesRequest) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

// ListExileCertificateResponses is the response for listing exile certificate.
type ListExileCertificatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of exile certificate.
	ExileCertificates []*ExileCertificate `protobuf:"bytes,1,rep,name=exile_certificates,json=exileCertificates,proto3" json:"exile_certificates,omitempty"`
}

func (x *ListExileCertificatesResponse) Reset() {
	*x = ListExileCertificatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExileCertificatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExileCertificatesResponse) ProtoMessage() {}

func (x *ListExileCertificatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExileCertificatesResponse.ProtoReflect.Descriptor instead.
func (*ListExileCertificatesResponse) Descriptor() ([]byte, []int) {
	return file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescGZIP(), []int{9}
}

func (x *ListExileCertificatesResponse) GetExileCertificates() []*ExileCertificate {
	if x != nil {
		return x.ExileCertificates
	}
	return nil
}

var File_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto protoreflect.FileDescriptor

var file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDesc = []byte{
	0x0a, 0x47, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65,
	0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x1d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69,
	0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x19, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64,
	0x5f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x22, 0x51, 0x0a, 0x1d, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x1e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x78,
	0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x1f, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x78,
	0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x16,
	0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65, 0x78,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x22, 0x0a, 0x20, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x0a, 0x21, 0x55, 0x6e, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x65,
	0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65, 0x78, 0x69, 0x6c, 0x65,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x24, 0x0a,
	0x22, 0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x59, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x69, 0x6c, 0x65,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x91,
	0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x70, 0x0a, 0x12, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c,
	0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45,
	0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x11, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x73, 0x42, 0x91, 0x03, 0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x15, 0x45, 0x78, 0x69, 0x6c,
	0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78, 0x69,
	0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x53, 0x4f, 0x45, 0xaa, 0x02, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x2e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x39, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x30, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a,
	0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescOnce sync.Once
	file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescData = file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDesc
)

func file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescGZIP() []byte {
	file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescOnce.Do(func() {
		file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescData)
	})
	return file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDescData
}

var file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_goTypes = []any{
	(*CreateExileCertificateRequest)(nil),      // 0: services.org.exile_certificate_manager.v1alpha1.CreateExileCertificateRequest
	(*CreateExileCertificateResponse)(nil),     // 1: services.org.exile_certificate_manager.v1alpha1.CreateExileCertificateResponse
	(*RevokeExileCertificateRequest)(nil),      // 2: services.org.exile_certificate_manager.v1alpha1.RevokeExileCertificateRequest
	(*RevokeExileCertificateResponse)(nil),     // 3: services.org.exile_certificate_manager.v1alpha1.RevokeExileCertificateResponse
	(*AssignExileConfigurationRequest)(nil),    // 4: services.org.exile_certificate_manager.v1alpha1.AssignExileConfigurationRequest
	(*AssignExileConfigurationResponse)(nil),   // 5: services.org.exile_certificate_manager.v1alpha1.AssignExileConfigurationResponse
	(*UnassignExileConfigurationRequest)(nil),  // 6: services.org.exile_certificate_manager.v1alpha1.UnassignExileConfigurationRequest
	(*UnassignExileConfigurationResponse)(nil), // 7: services.org.exile_certificate_manager.v1alpha1.UnassignExileConfigurationResponse
	(*ListExileCertificatesRequest)(nil),       // 8: services.org.exile_certificate_manager.v1alpha1.ListExileCertificatesRequest
	(*ListExileCertificatesResponse)(nil),      // 9: services.org.exile_certificate_manager.v1alpha1.ListExileCertificatesResponse
	(*fieldmaskpb.FieldMask)(nil),              // 10: google.protobuf.FieldMask
	(*ExileCertificate)(nil),                   // 11: services.org.exile_certificate_manager.v1alpha1.ExileCertificate
}
var file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_depIdxs = []int32{
	10, // 0: services.org.exile_certificate_manager.v1alpha1.ListExileCertificatesRequest.field_mask:type_name -> google.protobuf.FieldMask
	11, // 1: services.org.exile_certificate_manager.v1alpha1.ListExileCertificatesResponse.exile_certificates:type_name -> services.org.exile_certificate_manager.v1alpha1.ExileCertificate
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_init() }
func file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_init() {
	if File_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto != nil {
		return
	}
	file_services_org_exile_certificate_manager_v1alpha1_entities_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateExileCertificateRequest); i {
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
		file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateExileCertificateResponse); i {
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
		file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RevokeExileCertificateRequest); i {
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
		file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RevokeExileCertificateResponse); i {
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
		file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AssignExileConfigurationRequest); i {
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
		file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AssignExileConfigurationResponse); i {
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
		file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UnassignExileConfigurationRequest); i {
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
		file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UnassignExileConfigurationResponse); i {
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
		file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ListExileCertificatesRequest); i {
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
		file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ListExileCertificatesResponse); i {
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
			RawDescriptor: file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_goTypes,
		DependencyIndexes: file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_depIdxs,
		MessageInfos:      file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_msgTypes,
	}.Build()
	File_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto = out.File
	file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_rawDesc = nil
	file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_goTypes = nil
	file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_depIdxs = nil
}
