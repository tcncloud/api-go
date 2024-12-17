// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: api/v1alpha1/lms/api.proto

package lms

import (
	_ "github.com/tcncloud/api-go/annotations"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeleteFileTemplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileTemplate  *FileTemplateV2        `protobuf:"bytes,1,opt,name=file_template,json=fileTemplate,proto3" json:"file_template,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFileTemplateRequest) Reset() {
	*x = DeleteFileTemplateRequest{}
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFileTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileTemplateRequest) ProtoMessage() {}

func (x *DeleteFileTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileTemplateRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_api_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteFileTemplateRequest) GetFileTemplate() *FileTemplateV2 {
	if x != nil {
		return x.FileTemplate
	}
	return nil
}

type DeleteFileTemplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileTemplate  *FileTemplateV2        `protobuf:"bytes,1,opt,name=file_template,json=fileTemplate,proto3" json:"file_template,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFileTemplateResponse) Reset() {
	*x = DeleteFileTemplateResponse{}
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFileTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileTemplateResponse) ProtoMessage() {}

func (x *DeleteFileTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileTemplateResponse.ProtoReflect.Descriptor instead.
func (*DeleteFileTemplateResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_api_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteFileTemplateResponse) GetFileTemplate() *FileTemplateV2 {
	if x != nil {
		return x.FileTemplate
	}
	return nil
}

type GetFileTemplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileTemplate  *FileTemplateV2        `protobuf:"bytes,1,opt,name=file_template,json=fileTemplate,proto3" json:"file_template,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFileTemplateRequest) Reset() {
	*x = GetFileTemplateRequest{}
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFileTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileTemplateRequest) ProtoMessage() {}

func (x *GetFileTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileTemplateRequest.ProtoReflect.Descriptor instead.
func (*GetFileTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetFileTemplateRequest) GetFileTemplate() *FileTemplateV2 {
	if x != nil {
		return x.FileTemplate
	}
	return nil
}

type GetFileTemplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileTemplate  *FileTemplateV2        `protobuf:"bytes,1,opt,name=file_template,json=fileTemplate,proto3" json:"file_template,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFileTemplateResponse) Reset() {
	*x = GetFileTemplateResponse{}
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFileTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileTemplateResponse) ProtoMessage() {}

func (x *GetFileTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileTemplateResponse.ProtoReflect.Descriptor instead.
func (*GetFileTemplateResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetFileTemplateResponse) GetFileTemplate() *FileTemplateV2 {
	if x != nil {
		return x.FileTemplate
	}
	return nil
}

type ListFileTemplatesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         *FileTemplateV2        `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	LastId        string                 `protobuf:"bytes,2,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
	PageSize      int32                  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFileTemplatesRequest) Reset() {
	*x = ListFileTemplatesRequest{}
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFileTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileTemplatesRequest) ProtoMessage() {}

func (x *ListFileTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileTemplatesRequest.ProtoReflect.Descriptor instead.
func (*ListFileTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_api_proto_rawDescGZIP(), []int{4}
}

func (x *ListFileTemplatesRequest) GetValue() *FileTemplateV2 {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ListFileTemplatesRequest) GetLastId() string {
	if x != nil {
		return x.LastId
	}
	return ""
}

func (x *ListFileTemplatesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListFileTemplatesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileTemplates []*FileTemplateV2      `protobuf:"bytes,1,rep,name=file_templates,json=fileTemplates,proto3" json:"file_templates,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFileTemplatesResponse) Reset() {
	*x = ListFileTemplatesResponse{}
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFileTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileTemplatesResponse) ProtoMessage() {}

func (x *ListFileTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileTemplatesResponse.ProtoReflect.Descriptor instead.
func (*ListFileTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListFileTemplatesResponse) GetFileTemplates() []*FileTemplateV2 {
	if x != nil {
		return x.FileTemplates
	}
	return nil
}

type ParseFileTemplateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to RetrieveBy:
	//
	//	*ParseFileTemplateRequest_NewTemplate
	//	*ParseFileTemplateRequest_ExistingTemplate
	RetrieveBy    isParseFileTemplateRequest_RetrieveBy `protobuf_oneof:"retrieve_by"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParseFileTemplateRequest) Reset() {
	*x = ParseFileTemplateRequest{}
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseFileTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseFileTemplateRequest) ProtoMessage() {}

func (x *ParseFileTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseFileTemplateRequest.ProtoReflect.Descriptor instead.
func (*ParseFileTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_api_proto_rawDescGZIP(), []int{6}
}

func (x *ParseFileTemplateRequest) GetRetrieveBy() isParseFileTemplateRequest_RetrieveBy {
	if x != nil {
		return x.RetrieveBy
	}
	return nil
}

func (x *ParseFileTemplateRequest) GetNewTemplate() *NewTemplate {
	if x != nil {
		if x, ok := x.RetrieveBy.(*ParseFileTemplateRequest_NewTemplate); ok {
			return x.NewTemplate
		}
	}
	return nil
}

func (x *ParseFileTemplateRequest) GetExistingTemplate() *ExistingTemplate {
	if x != nil {
		if x, ok := x.RetrieveBy.(*ParseFileTemplateRequest_ExistingTemplate); ok {
			return x.ExistingTemplate
		}
	}
	return nil
}

type isParseFileTemplateRequest_RetrieveBy interface {
	isParseFileTemplateRequest_RetrieveBy()
}

type ParseFileTemplateRequest_NewTemplate struct {
	NewTemplate *NewTemplate `protobuf:"bytes,1,opt,name=new_template,json=newTemplate,proto3,oneof"`
}

type ParseFileTemplateRequest_ExistingTemplate struct {
	ExistingTemplate *ExistingTemplate `protobuf:"bytes,2,opt,name=existing_template,json=existingTemplate,proto3,oneof"`
}

func (*ParseFileTemplateRequest_NewTemplate) isParseFileTemplateRequest_RetrieveBy() {}

func (*ParseFileTemplateRequest_ExistingTemplate) isParseFileTemplateRequest_RetrieveBy() {}

type ParseFileTemplateResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParseFileTemplateResult) Reset() {
	*x = ParseFileTemplateResult{}
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseFileTemplateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseFileTemplateResult) ProtoMessage() {}

func (x *ParseFileTemplateResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseFileTemplateResult.ProtoReflect.Descriptor instead.
func (*ParseFileTemplateResult) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_api_proto_rawDescGZIP(), []int{7}
}

type UpdateFileTemplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileTemplate  *FileTemplateV2        `protobuf:"bytes,1,opt,name=file_template,json=fileTemplate,proto3" json:"file_template,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFileTemplateRequest) Reset() {
	*x = UpdateFileTemplateRequest{}
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFileTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileTemplateRequest) ProtoMessage() {}

func (x *UpdateFileTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileTemplateRequest.ProtoReflect.Descriptor instead.
func (*UpdateFileTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_api_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateFileTemplateRequest) GetFileTemplate() *FileTemplateV2 {
	if x != nil {
		return x.FileTemplate
	}
	return nil
}

type UpdateFileTemplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileTemplate  *FileTemplateV2        `protobuf:"bytes,1,opt,name=file_template,json=fileTemplate,proto3" json:"file_template,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFileTemplateResponse) Reset() {
	*x = UpdateFileTemplateResponse{}
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFileTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileTemplateResponse) ProtoMessage() {}

func (x *UpdateFileTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_api_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileTemplateResponse.ProtoReflect.Descriptor instead.
func (*UpdateFileTemplateResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_api_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateFileTemplateResponse) GetFileTemplate() *FileTemplateV2 {
	if x != nil {
		return x.FileTemplate
	}
	return nil
}

var File_api_v1alpha1_lms_api_proto protoreflect.FileDescriptor

var file_api_v1alpha1_lms_api_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c,
	0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x1a, 0x17,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x32, 0x52, 0x0c, 0x66, 0x69,
	0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x63, 0x0a, 0x1a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c,
	0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56,
	0x32, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22,
	0x5f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0d, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x56, 0x32, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x22, 0x60, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x56, 0x32, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d,
	0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x32,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x64, 0x0a,
	0x19, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x56, 0x32, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x18, 0x50, 0x61, 0x72, 0x73, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x42, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x51, 0x0a, 0x11, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c,
	0x6d, 0x73, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x10, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x5f, 0x62, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x50, 0x61, 0x72, 0x73, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x62, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45,
	0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x32, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x63, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x32, 0x52, 0x0c, 0x66, 0x69,
	0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x32, 0xd0, 0x06, 0x0a, 0x03, 0x4c,
	0x4d, 0x53, 0x12, 0xaa, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xe9, 0x07,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x6d, 0x73, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x9e, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a,
	0x03, 0x08, 0xe8, 0x07, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x6d, 0x73,
	0x2f, 0x67, 0x65, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0xa6, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x38, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xe8, 0x07, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6c, 0x6d, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x66, 0x69, 0x6c, 0x65,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0xa4, 0x01, 0x0a, 0x11, 0x50, 0x61,
	0x72, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c,
	0x6d, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x38, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08,
	0xe9, 0x07, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x6d, 0x73, 0x2f, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0xaa, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x39, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xe9, 0x07, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x6d, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0xaf, 0x01,
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x6c, 0x6d, 0x73, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x6d, 0x73, 0xa2,
	0x02, 0x03, 0x41, 0x56, 0x4c, 0xaa, 0x02, 0x10, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6d, 0x73, 0xca, 0x02, 0x10, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x4c, 0x6d, 0x73, 0xe2, 0x02, 0x1c, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x4c, 0x6d, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x4c, 0x6d, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_lms_api_proto_rawDescOnce sync.Once
	file_api_v1alpha1_lms_api_proto_rawDescData = file_api_v1alpha1_lms_api_proto_rawDesc
)

func file_api_v1alpha1_lms_api_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_lms_api_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_lms_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_lms_api_proto_rawDescData)
	})
	return file_api_v1alpha1_lms_api_proto_rawDescData
}

var file_api_v1alpha1_lms_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v1alpha1_lms_api_proto_goTypes = []any{
	(*DeleteFileTemplateRequest)(nil),  // 0: api.v1alpha1.lms.DeleteFileTemplateRequest
	(*DeleteFileTemplateResponse)(nil), // 1: api.v1alpha1.lms.DeleteFileTemplateResponse
	(*GetFileTemplateRequest)(nil),     // 2: api.v1alpha1.lms.GetFileTemplateRequest
	(*GetFileTemplateResponse)(nil),    // 3: api.v1alpha1.lms.GetFileTemplateResponse
	(*ListFileTemplatesRequest)(nil),   // 4: api.v1alpha1.lms.ListFileTemplatesRequest
	(*ListFileTemplatesResponse)(nil),  // 5: api.v1alpha1.lms.ListFileTemplatesResponse
	(*ParseFileTemplateRequest)(nil),   // 6: api.v1alpha1.lms.ParseFileTemplateRequest
	(*ParseFileTemplateResult)(nil),    // 7: api.v1alpha1.lms.ParseFileTemplateResult
	(*UpdateFileTemplateRequest)(nil),  // 8: api.v1alpha1.lms.UpdateFileTemplateRequest
	(*UpdateFileTemplateResponse)(nil), // 9: api.v1alpha1.lms.UpdateFileTemplateResponse
	(*FileTemplateV2)(nil),             // 10: api.v1alpha1.lms.FileTemplateV2
	(*NewTemplate)(nil),                // 11: api.v1alpha1.lms.NewTemplate
	(*ExistingTemplate)(nil),           // 12: api.v1alpha1.lms.ExistingTemplate
}
var file_api_v1alpha1_lms_api_proto_depIdxs = []int32{
	10, // 0: api.v1alpha1.lms.DeleteFileTemplateRequest.file_template:type_name -> api.v1alpha1.lms.FileTemplateV2
	10, // 1: api.v1alpha1.lms.DeleteFileTemplateResponse.file_template:type_name -> api.v1alpha1.lms.FileTemplateV2
	10, // 2: api.v1alpha1.lms.GetFileTemplateRequest.file_template:type_name -> api.v1alpha1.lms.FileTemplateV2
	10, // 3: api.v1alpha1.lms.GetFileTemplateResponse.file_template:type_name -> api.v1alpha1.lms.FileTemplateV2
	10, // 4: api.v1alpha1.lms.ListFileTemplatesRequest.value:type_name -> api.v1alpha1.lms.FileTemplateV2
	10, // 5: api.v1alpha1.lms.ListFileTemplatesResponse.file_templates:type_name -> api.v1alpha1.lms.FileTemplateV2
	11, // 6: api.v1alpha1.lms.ParseFileTemplateRequest.new_template:type_name -> api.v1alpha1.lms.NewTemplate
	12, // 7: api.v1alpha1.lms.ParseFileTemplateRequest.existing_template:type_name -> api.v1alpha1.lms.ExistingTemplate
	10, // 8: api.v1alpha1.lms.UpdateFileTemplateRequest.file_template:type_name -> api.v1alpha1.lms.FileTemplateV2
	10, // 9: api.v1alpha1.lms.UpdateFileTemplateResponse.file_template:type_name -> api.v1alpha1.lms.FileTemplateV2
	0,  // 10: api.v1alpha1.lms.LMS.DeleteFileTemplate:input_type -> api.v1alpha1.lms.DeleteFileTemplateRequest
	2,  // 11: api.v1alpha1.lms.LMS.GetFileTemplate:input_type -> api.v1alpha1.lms.GetFileTemplateRequest
	4,  // 12: api.v1alpha1.lms.LMS.ListFileTemplates:input_type -> api.v1alpha1.lms.ListFileTemplatesRequest
	6,  // 13: api.v1alpha1.lms.LMS.ParseFileTemplate:input_type -> api.v1alpha1.lms.ParseFileTemplateRequest
	8,  // 14: api.v1alpha1.lms.LMS.UpdateFileTemplate:input_type -> api.v1alpha1.lms.UpdateFileTemplateRequest
	1,  // 15: api.v1alpha1.lms.LMS.DeleteFileTemplate:output_type -> api.v1alpha1.lms.DeleteFileTemplateResponse
	3,  // 16: api.v1alpha1.lms.LMS.GetFileTemplate:output_type -> api.v1alpha1.lms.GetFileTemplateResponse
	5,  // 17: api.v1alpha1.lms.LMS.ListFileTemplates:output_type -> api.v1alpha1.lms.ListFileTemplatesResponse
	7,  // 18: api.v1alpha1.lms.LMS.ParseFileTemplate:output_type -> api.v1alpha1.lms.ParseFileTemplateResult
	9,  // 19: api.v1alpha1.lms.LMS.UpdateFileTemplate:output_type -> api.v1alpha1.lms.UpdateFileTemplateResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_lms_api_proto_init() }
func file_api_v1alpha1_lms_api_proto_init() {
	if File_api_v1alpha1_lms_api_proto != nil {
		return
	}
	file_api_v1alpha1_lms_entities_proto_init()
	file_api_v1alpha1_lms_api_proto_msgTypes[6].OneofWrappers = []any{
		(*ParseFileTemplateRequest_NewTemplate)(nil),
		(*ParseFileTemplateRequest_ExistingTemplate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_lms_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_lms_api_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_lms_api_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_lms_api_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_lms_api_proto = out.File
	file_api_v1alpha1_lms_api_proto_rawDesc = nil
	file_api_v1alpha1_lms_api_proto_goTypes = nil
	file_api_v1alpha1_lms_api_proto_depIdxs = nil
}
