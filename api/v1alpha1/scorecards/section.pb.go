// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/v1alpha1/scorecards/section.proto

package scorecards

import (
	commons "github.com/tcncloud/api-go/api/commons"
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

// CreateSectionRequest is the request to create a new section.
type CreateSectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section *commons.Section `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *CreateSectionRequest) Reset() {
	*x = CreateSectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSectionRequest) ProtoMessage() {}

func (x *CreateSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSectionRequest.ProtoReflect.Descriptor instead.
func (*CreateSectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_section_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSectionRequest) GetSection() *commons.Section {
	if x != nil {
		return x.Section
	}
	return nil
}

// CreateSectionResponse returns the created section.
type CreateSectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section *commons.Section `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *CreateSectionResponse) Reset() {
	*x = CreateSectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSectionResponse) ProtoMessage() {}

func (x *CreateSectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSectionResponse.ProtoReflect.Descriptor instead.
func (*CreateSectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_section_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSectionResponse) GetSection() *commons.Section {
	if x != nil {
		return x.Section
	}
	return nil
}

// ListSectionsRequest is request to list sections by scorecard id.
type ListSectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorecardId int64 `protobuf:"varint,2,opt,name=scorecard_id,json=scorecardId,proto3" json:"scorecard_id,omitempty"` // Required.
}

func (x *ListSectionsRequest) Reset() {
	*x = ListSectionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSectionsRequest) ProtoMessage() {}

func (x *ListSectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSectionsRequest.ProtoReflect.Descriptor instead.
func (*ListSectionsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_section_proto_rawDescGZIP(), []int{2}
}

func (x *ListSectionsRequest) GetScorecardId() int64 {
	if x != nil {
		return x.ScorecardId
	}
	return 0
}

// ListSectionsResponse returns a list of sections.
type ListSectionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sections []*commons.Section `protobuf:"bytes,1,rep,name=sections,proto3" json:"sections,omitempty"`
}

func (x *ListSectionsResponse) Reset() {
	*x = ListSectionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSectionsResponse) ProtoMessage() {}

func (x *ListSectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSectionsResponse.ProtoReflect.Descriptor instead.
func (*ListSectionsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_section_proto_rawDescGZIP(), []int{3}
}

func (x *ListSectionsResponse) GetSections() []*commons.Section {
	if x != nil {
		return x.Sections
	}
	return nil
}

// UpdateSectionRequest is the request to update a section.
type UpdateSectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	Section *commons.Section `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
	// Required. Valid paths: [title, description, weight].
	// To update Questions, use *ScorecardQuestion methods.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateSectionRequest) Reset() {
	*x = UpdateSectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSectionRequest) ProtoMessage() {}

func (x *UpdateSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSectionRequest.ProtoReflect.Descriptor instead.
func (*UpdateSectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_section_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSectionRequest) GetSection() *commons.Section {
	if x != nil {
		return x.Section
	}
	return nil
}

func (x *UpdateSectionRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// UpdateSectionResponse returns the updated section.
type UpdateSectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section *commons.Section `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"` // Updated entity
}

func (x *UpdateSectionResponse) Reset() {
	*x = UpdateSectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSectionResponse) ProtoMessage() {}

func (x *UpdateSectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSectionResponse.ProtoReflect.Descriptor instead.
func (*UpdateSectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_section_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSectionResponse) GetSection() *commons.Section {
	if x != nil {
		return x.Section
	}
	return nil
}

// DeleteSectionRequest is the request to delete a section.
type DeleteSectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SectionId int64 `protobuf:"varint,2,opt,name=section_id,json=sectionId,proto3" json:"section_id,omitempty"` // Required.
}

func (x *DeleteSectionRequest) Reset() {
	*x = DeleteSectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSectionRequest) ProtoMessage() {}

func (x *DeleteSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSectionRequest.ProtoReflect.Descriptor instead.
func (*DeleteSectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_section_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSectionRequest) GetSectionId() int64 {
	if x != nil {
		return x.SectionId
	}
	return 0
}

// DeleteSectionResponse returns the deleted section.
type DeleteSectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section *commons.Section `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"` // Deleted entity
}

func (x *DeleteSectionResponse) Reset() {
	*x = DeleteSectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSectionResponse) ProtoMessage() {}

func (x *DeleteSectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSectionResponse.ProtoReflect.Descriptor instead.
func (*DeleteSectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_section_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSectionResponse) GetSection() *commons.Section {
	if x != nil {
		return x.Section
	}
	return nil
}

// GetSectionRequest is the request to get a section.
type GetSectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SectionId int64 `protobuf:"varint,2,opt,name=section_id,json=sectionId,proto3" json:"section_id,omitempty"` // Unique id of section.
}

func (x *GetSectionRequest) Reset() {
	*x = GetSectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSectionRequest) ProtoMessage() {}

func (x *GetSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSectionRequest.ProtoReflect.Descriptor instead.
func (*GetSectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_section_proto_rawDescGZIP(), []int{8}
}

func (x *GetSectionRequest) GetSectionId() int64 {
	if x != nil {
		return x.SectionId
	}
	return 0
}

// GetSectionResponse is the response to getting a section.
type GetSectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section *commons.Section `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"` // returned section
}

func (x *GetSectionResponse) Reset() {
	*x = GetSectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSectionResponse) ProtoMessage() {}

func (x *GetSectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_section_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSectionResponse.ProtoReflect.Descriptor instead.
func (*GetSectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_section_proto_rawDescGZIP(), []int{9}
}

func (x *GetSectionResponse) GetSection() *commons.Section {
	if x != nil {
		return x.Section
	}
	return nil
}

var File_api_v1alpha1_scorecards_section_proto protoreflect.FileDescriptor

var file_api_v1alpha1_scorecards_section_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x46, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x38, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x47, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x15, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xdd,
	0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x42, 0x0c,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72,
	0x64, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x53, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72,
	0x64, 0x73, 0xca, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0xe2, 0x02, 0x23, 0x41,
	0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3a, 0x3a, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_scorecards_section_proto_rawDescOnce sync.Once
	file_api_v1alpha1_scorecards_section_proto_rawDescData = file_api_v1alpha1_scorecards_section_proto_rawDesc
)

func file_api_v1alpha1_scorecards_section_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_scorecards_section_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_scorecards_section_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_scorecards_section_proto_rawDescData)
	})
	return file_api_v1alpha1_scorecards_section_proto_rawDescData
}

var file_api_v1alpha1_scorecards_section_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v1alpha1_scorecards_section_proto_goTypes = []interface{}{
	(*CreateSectionRequest)(nil),  // 0: api.v1alpha1.scorecards.CreateSectionRequest
	(*CreateSectionResponse)(nil), // 1: api.v1alpha1.scorecards.CreateSectionResponse
	(*ListSectionsRequest)(nil),   // 2: api.v1alpha1.scorecards.ListSectionsRequest
	(*ListSectionsResponse)(nil),  // 3: api.v1alpha1.scorecards.ListSectionsResponse
	(*UpdateSectionRequest)(nil),  // 4: api.v1alpha1.scorecards.UpdateSectionRequest
	(*UpdateSectionResponse)(nil), // 5: api.v1alpha1.scorecards.UpdateSectionResponse
	(*DeleteSectionRequest)(nil),  // 6: api.v1alpha1.scorecards.DeleteSectionRequest
	(*DeleteSectionResponse)(nil), // 7: api.v1alpha1.scorecards.DeleteSectionResponse
	(*GetSectionRequest)(nil),     // 8: api.v1alpha1.scorecards.GetSectionRequest
	(*GetSectionResponse)(nil),    // 9: api.v1alpha1.scorecards.GetSectionResponse
	(*commons.Section)(nil),       // 10: api.commons.Section
	(*fieldmaskpb.FieldMask)(nil), // 11: google.protobuf.FieldMask
}
var file_api_v1alpha1_scorecards_section_proto_depIdxs = []int32{
	10, // 0: api.v1alpha1.scorecards.CreateSectionRequest.section:type_name -> api.commons.Section
	10, // 1: api.v1alpha1.scorecards.CreateSectionResponse.section:type_name -> api.commons.Section
	10, // 2: api.v1alpha1.scorecards.ListSectionsResponse.sections:type_name -> api.commons.Section
	10, // 3: api.v1alpha1.scorecards.UpdateSectionRequest.section:type_name -> api.commons.Section
	11, // 4: api.v1alpha1.scorecards.UpdateSectionRequest.update_mask:type_name -> google.protobuf.FieldMask
	10, // 5: api.v1alpha1.scorecards.UpdateSectionResponse.section:type_name -> api.commons.Section
	10, // 6: api.v1alpha1.scorecards.DeleteSectionResponse.section:type_name -> api.commons.Section
	10, // 7: api.v1alpha1.scorecards.GetSectionResponse.section:type_name -> api.commons.Section
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_scorecards_section_proto_init() }
func file_api_v1alpha1_scorecards_section_proto_init() {
	if File_api_v1alpha1_scorecards_section_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1alpha1_scorecards_section_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSectionRequest); i {
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
		file_api_v1alpha1_scorecards_section_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSectionResponse); i {
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
		file_api_v1alpha1_scorecards_section_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSectionsRequest); i {
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
		file_api_v1alpha1_scorecards_section_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSectionsResponse); i {
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
		file_api_v1alpha1_scorecards_section_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSectionRequest); i {
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
		file_api_v1alpha1_scorecards_section_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSectionResponse); i {
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
		file_api_v1alpha1_scorecards_section_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSectionRequest); i {
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
		file_api_v1alpha1_scorecards_section_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSectionResponse); i {
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
		file_api_v1alpha1_scorecards_section_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSectionRequest); i {
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
		file_api_v1alpha1_scorecards_section_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSectionResponse); i {
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
			RawDescriptor: file_api_v1alpha1_scorecards_section_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_scorecards_section_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_scorecards_section_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_scorecards_section_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_scorecards_section_proto = out.File
	file_api_v1alpha1_scorecards_section_proto_rawDesc = nil
	file_api_v1alpha1_scorecards_section_proto_goTypes = nil
	file_api_v1alpha1_scorecards_section_proto_depIdxs = nil
}
