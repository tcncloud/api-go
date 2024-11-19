// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/v1alpha1/scorecards/scorecard_question.proto

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

// GetScorecardQuestionRequest is the request to get a scorecard question.
type GetScorecardQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorecardQuestionId int64 `protobuf:"varint,2,opt,name=scorecard_question_id,json=scorecardQuestionId,proto3" json:"scorecard_question_id,omitempty"` // Required - unique id of scorecard question to get.
}

func (x *GetScorecardQuestionRequest) Reset() {
	*x = GetScorecardQuestionRequest{}
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetScorecardQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScorecardQuestionRequest) ProtoMessage() {}

func (x *GetScorecardQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScorecardQuestionRequest.ProtoReflect.Descriptor instead.
func (*GetScorecardQuestionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescGZIP(), []int{0}
}

func (x *GetScorecardQuestionRequest) GetScorecardQuestionId() int64 {
	if x != nil {
		return x.ScorecardQuestionId
	}
	return 0
}

// GetScorecardQuestionResponse returns the scorecard question requested.
type GetScorecardQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorecardQuestion *commons.ScorecardQuestion `protobuf:"bytes,1,opt,name=scorecard_question,json=scorecardQuestion,proto3" json:"scorecard_question,omitempty"` // Requested scorecard question resource.
}

func (x *GetScorecardQuestionResponse) Reset() {
	*x = GetScorecardQuestionResponse{}
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetScorecardQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScorecardQuestionResponse) ProtoMessage() {}

func (x *GetScorecardQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScorecardQuestionResponse.ProtoReflect.Descriptor instead.
func (*GetScorecardQuestionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescGZIP(), []int{1}
}

func (x *GetScorecardQuestionResponse) GetScorecardQuestion() *commons.ScorecardQuestion {
	if x != nil {
		return x.ScorecardQuestion
	}
	return nil
}

// CreateScorecardQuestionRequest is the request to create a new scorecard question.
type CreateScorecardQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorecardQuestion *commons.ScorecardQuestion `protobuf:"bytes,1,opt,name=scorecard_question,json=scorecardQuestion,proto3" json:"scorecard_question,omitempty"`
}

func (x *CreateScorecardQuestionRequest) Reset() {
	*x = CreateScorecardQuestionRequest{}
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateScorecardQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScorecardQuestionRequest) ProtoMessage() {}

func (x *CreateScorecardQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScorecardQuestionRequest.ProtoReflect.Descriptor instead.
func (*CreateScorecardQuestionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescGZIP(), []int{2}
}

func (x *CreateScorecardQuestionRequest) GetScorecardQuestion() *commons.ScorecardQuestion {
	if x != nil {
		return x.ScorecardQuestion
	}
	return nil
}

// CreateScorecardQuestionResponse returns the created scorecard.
type CreateScorecardQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorecardQuestion *commons.ScorecardQuestion `protobuf:"bytes,1,opt,name=scorecard_question,json=scorecardQuestion,proto3" json:"scorecard_question,omitempty"`
}

func (x *CreateScorecardQuestionResponse) Reset() {
	*x = CreateScorecardQuestionResponse{}
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateScorecardQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScorecardQuestionResponse) ProtoMessage() {}

func (x *CreateScorecardQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScorecardQuestionResponse.ProtoReflect.Descriptor instead.
func (*CreateScorecardQuestionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescGZIP(), []int{3}
}

func (x *CreateScorecardQuestionResponse) GetScorecardQuestion() *commons.ScorecardQuestion {
	if x != nil {
		return x.ScorecardQuestion
	}
	return nil
}

// UpdateScorecardQuestionRequest is the request to update a scorecard question.
type UpdateScorecardQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorecardQuestion *commons.ScorecardQuestion `protobuf:"bytes,1,opt,name=scorecard_question,json=scorecardQuestion,proto3" json:"scorecard_question,omitempty"` // Required.
	UpdateMask        *fieldmaskpb.FieldMask     `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`                      // Required.
}

func (x *UpdateScorecardQuestionRequest) Reset() {
	*x = UpdateScorecardQuestionRequest{}
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateScorecardQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScorecardQuestionRequest) ProtoMessage() {}

func (x *UpdateScorecardQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScorecardQuestionRequest.ProtoReflect.Descriptor instead.
func (*UpdateScorecardQuestionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateScorecardQuestionRequest) GetScorecardQuestion() *commons.ScorecardQuestion {
	if x != nil {
		return x.ScorecardQuestion
	}
	return nil
}

func (x *UpdateScorecardQuestionRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// UpdateScorecardQuestionResponse returns the updated scorecard question.
type UpdateScorecardQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorecardQuestion *commons.ScorecardQuestion `protobuf:"bytes,1,opt,name=scorecard_question,json=scorecardQuestion,proto3" json:"scorecard_question,omitempty"`
}

func (x *UpdateScorecardQuestionResponse) Reset() {
	*x = UpdateScorecardQuestionResponse{}
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateScorecardQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScorecardQuestionResponse) ProtoMessage() {}

func (x *UpdateScorecardQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScorecardQuestionResponse.ProtoReflect.Descriptor instead.
func (*UpdateScorecardQuestionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateScorecardQuestionResponse) GetScorecardQuestion() *commons.ScorecardQuestion {
	if x != nil {
		return x.ScorecardQuestion
	}
	return nil
}

// DeleteScorecardQuestionRequest is the request to delete a scorecard question.
type DeleteScorecardQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorecardQuestionId int64 `protobuf:"varint,2,opt,name=scorecard_question_id,json=scorecardQuestionId,proto3" json:"scorecard_question_id,omitempty"` // Required.
}

func (x *DeleteScorecardQuestionRequest) Reset() {
	*x = DeleteScorecardQuestionRequest{}
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteScorecardQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScorecardQuestionRequest) ProtoMessage() {}

func (x *DeleteScorecardQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScorecardQuestionRequest.ProtoReflect.Descriptor instead.
func (*DeleteScorecardQuestionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteScorecardQuestionRequest) GetScorecardQuestionId() int64 {
	if x != nil {
		return x.ScorecardQuestionId
	}
	return 0
}

// DeleteScorecardQuestionResponse returns the deleted scorecard question.
type DeleteScorecardQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorecardQuestion *commons.ScorecardQuestion `protobuf:"bytes,1,opt,name=scorecard_question,json=scorecardQuestion,proto3" json:"scorecard_question,omitempty"`
}

func (x *DeleteScorecardQuestionResponse) Reset() {
	*x = DeleteScorecardQuestionResponse{}
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteScorecardQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScorecardQuestionResponse) ProtoMessage() {}

func (x *DeleteScorecardQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScorecardQuestionResponse.ProtoReflect.Descriptor instead.
func (*DeleteScorecardQuestionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteScorecardQuestionResponse) GetScorecardQuestion() *commons.ScorecardQuestion {
	if x != nil {
		return x.ScorecardQuestion
	}
	return nil
}

var File_api_v1alpha1_scorecards_scorecard_question_proto protoreflect.FileDescriptor

var file_api_v1alpha1_scorecards_scorecard_question_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x1a, 0x1c, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x6d,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x12, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6f, 0x0a,
	0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4d, 0x0a, 0x12, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x70,
	0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72,
	0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4d, 0x0a, 0x12, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xac, 0x01, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x12, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x11, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22,
	0x70, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x54, 0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x13, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72,
	0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xe7, 0x01, 0x0a, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x42, 0x16, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x53, 0xaa, 0x02, 0x17,
	0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0xca, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64,
	0x73, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescOnce sync.Once
	file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescData = file_api_v1alpha1_scorecards_scorecard_question_proto_rawDesc
)

func file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescData)
	})
	return file_api_v1alpha1_scorecards_scorecard_question_proto_rawDescData
}

var file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_v1alpha1_scorecards_scorecard_question_proto_goTypes = []any{
	(*GetScorecardQuestionRequest)(nil),     // 0: api.v1alpha1.scorecards.GetScorecardQuestionRequest
	(*GetScorecardQuestionResponse)(nil),    // 1: api.v1alpha1.scorecards.GetScorecardQuestionResponse
	(*CreateScorecardQuestionRequest)(nil),  // 2: api.v1alpha1.scorecards.CreateScorecardQuestionRequest
	(*CreateScorecardQuestionResponse)(nil), // 3: api.v1alpha1.scorecards.CreateScorecardQuestionResponse
	(*UpdateScorecardQuestionRequest)(nil),  // 4: api.v1alpha1.scorecards.UpdateScorecardQuestionRequest
	(*UpdateScorecardQuestionResponse)(nil), // 5: api.v1alpha1.scorecards.UpdateScorecardQuestionResponse
	(*DeleteScorecardQuestionRequest)(nil),  // 6: api.v1alpha1.scorecards.DeleteScorecardQuestionRequest
	(*DeleteScorecardQuestionResponse)(nil), // 7: api.v1alpha1.scorecards.DeleteScorecardQuestionResponse
	(*commons.ScorecardQuestion)(nil),       // 8: api.commons.ScorecardQuestion
	(*fieldmaskpb.FieldMask)(nil),           // 9: google.protobuf.FieldMask
}
var file_api_v1alpha1_scorecards_scorecard_question_proto_depIdxs = []int32{
	8, // 0: api.v1alpha1.scorecards.GetScorecardQuestionResponse.scorecard_question:type_name -> api.commons.ScorecardQuestion
	8, // 1: api.v1alpha1.scorecards.CreateScorecardQuestionRequest.scorecard_question:type_name -> api.commons.ScorecardQuestion
	8, // 2: api.v1alpha1.scorecards.CreateScorecardQuestionResponse.scorecard_question:type_name -> api.commons.ScorecardQuestion
	8, // 3: api.v1alpha1.scorecards.UpdateScorecardQuestionRequest.scorecard_question:type_name -> api.commons.ScorecardQuestion
	9, // 4: api.v1alpha1.scorecards.UpdateScorecardQuestionRequest.update_mask:type_name -> google.protobuf.FieldMask
	8, // 5: api.v1alpha1.scorecards.UpdateScorecardQuestionResponse.scorecard_question:type_name -> api.commons.ScorecardQuestion
	8, // 6: api.v1alpha1.scorecards.DeleteScorecardQuestionResponse.scorecard_question:type_name -> api.commons.ScorecardQuestion
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_scorecards_scorecard_question_proto_init() }
func file_api_v1alpha1_scorecards_scorecard_question_proto_init() {
	if File_api_v1alpha1_scorecards_scorecard_question_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_scorecards_scorecard_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_scorecards_scorecard_question_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_scorecards_scorecard_question_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_scorecards_scorecard_question_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_scorecards_scorecard_question_proto = out.File
	file_api_v1alpha1_scorecards_scorecard_question_proto_rawDesc = nil
	file_api_v1alpha1_scorecards_scorecard_question_proto_goTypes = nil
	file_api_v1alpha1_scorecards_scorecard_question_proto_depIdxs = nil
}
