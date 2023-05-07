// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/v1alpha1/scorecards/evaluation_question.proto

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

// CreateEvaluationQuestionRequest is request to create an evaluation question
type CreateEvaluationQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvaluationQuestion *commons.EvaluationQuestion `protobuf:"bytes,1,opt,name=evaluation_question,json=evaluationQuestion,proto3" json:"evaluation_question,omitempty"` // Required - the entity which is requested to be created
}

func (x *CreateEvaluationQuestionRequest) Reset() {
	*x = CreateEvaluationQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEvaluationQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEvaluationQuestionRequest) ProtoMessage() {}

func (x *CreateEvaluationQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEvaluationQuestionRequest.ProtoReflect.Descriptor instead.
func (*CreateEvaluationQuestionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEvaluationQuestionRequest) GetEvaluationQuestion() *commons.EvaluationQuestion {
	if x != nil {
		return x.EvaluationQuestion
	}
	return nil
}

// CreateEvaluationQuestionResponse is response with requested entity
type CreateEvaluationQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvaluationQuestion *commons.EvaluationQuestion `protobuf:"bytes,1,opt,name=evaluation_question,json=evaluationQuestion,proto3" json:"evaluation_question,omitempty"` // Required - the entity which was created
}

func (x *CreateEvaluationQuestionResponse) Reset() {
	*x = CreateEvaluationQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEvaluationQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEvaluationQuestionResponse) ProtoMessage() {}

func (x *CreateEvaluationQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEvaluationQuestionResponse.ProtoReflect.Descriptor instead.
func (*CreateEvaluationQuestionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEvaluationQuestionResponse) GetEvaluationQuestion() *commons.EvaluationQuestion {
	if x != nil {
		return x.EvaluationQuestion
	}
	return nil
}

// UpdateEvaluationQuestionRequest is request to update an existing evaluation question
type UpdateEvaluationQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvaluationQuestion *commons.EvaluationQuestion `protobuf:"bytes,1,opt,name=evaluation_question,json=evaluationQuestion,proto3" json:"evaluation_question,omitempty"` // Required - evaluation question that is to be updated
	UpdateMask         *fieldmaskpb.FieldMask      `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`                         // Required - specification of which fields should be updated
}

func (x *UpdateEvaluationQuestionRequest) Reset() {
	*x = UpdateEvaluationQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEvaluationQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEvaluationQuestionRequest) ProtoMessage() {}

func (x *UpdateEvaluationQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEvaluationQuestionRequest.ProtoReflect.Descriptor instead.
func (*UpdateEvaluationQuestionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateEvaluationQuestionRequest) GetEvaluationQuestion() *commons.EvaluationQuestion {
	if x != nil {
		return x.EvaluationQuestion
	}
	return nil
}

func (x *UpdateEvaluationQuestionRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// UpdateEvaluationQuestionResponse is response with updated entity
type UpdateEvaluationQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvaluationQuestion *commons.EvaluationQuestion `protobuf:"bytes,1,opt,name=evaluation_question,json=evaluationQuestion,proto3" json:"evaluation_question,omitempty"` // updated evaluation question
}

func (x *UpdateEvaluationQuestionResponse) Reset() {
	*x = UpdateEvaluationQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEvaluationQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEvaluationQuestionResponse) ProtoMessage() {}

func (x *UpdateEvaluationQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEvaluationQuestionResponse.ProtoReflect.Descriptor instead.
func (*UpdateEvaluationQuestionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEvaluationQuestionResponse) GetEvaluationQuestion() *commons.EvaluationQuestion {
	if x != nil {
		return x.EvaluationQuestion
	}
	return nil
}

// DeleteEvaluationQuestionRequest is request to delete an evaluation question
type DeleteEvaluationQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvaluationQuestionId int64 `protobuf:"varint,2,opt,name=evaluation_question_id,json=evaluationQuestionId,proto3" json:"evaluation_question_id,omitempty"` // Required - unique id of evaluation question to delete
}

func (x *DeleteEvaluationQuestionRequest) Reset() {
	*x = DeleteEvaluationQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEvaluationQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEvaluationQuestionRequest) ProtoMessage() {}

func (x *DeleteEvaluationQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEvaluationQuestionRequest.ProtoReflect.Descriptor instead.
func (*DeleteEvaluationQuestionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteEvaluationQuestionRequest) GetEvaluationQuestionId() int64 {
	if x != nil {
		return x.EvaluationQuestionId
	}
	return 0
}

// DeleteEvaluationQuestionResponse is response with deleted entity
type DeleteEvaluationQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvaluationQuestion *commons.EvaluationQuestion `protobuf:"bytes,1,opt,name=evaluation_question,json=evaluationQuestion,proto3" json:"evaluation_question,omitempty"` // evaluation question object deleted
}

func (x *DeleteEvaluationQuestionResponse) Reset() {
	*x = DeleteEvaluationQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEvaluationQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEvaluationQuestionResponse) ProtoMessage() {}

func (x *DeleteEvaluationQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEvaluationQuestionResponse.ProtoReflect.Descriptor instead.
func (*DeleteEvaluationQuestionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteEvaluationQuestionResponse) GetEvaluationQuestion() *commons.EvaluationQuestion {
	if x != nil {
		return x.EvaluationQuestion
	}
	return nil
}

var File_api_v1alpha1_scorecards_evaluation_question_proto protoreflect.FileDescriptor

var file_api_v1alpha1_scorecards_evaluation_question_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x1a, 0x1c, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x1f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x50, 0x0a, 0x13, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x74, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x13, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x12, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb0, 0x01, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x13, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x74, 0x0a, 0x20, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50,
	0x0a, 0x13, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x57, 0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x14, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x20, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a,
	0x13, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x65, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0xe8, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x42,
	0x17, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0xa2, 0x02,
	0x03, 0x41, 0x56, 0x53, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0xca, 0x02,
	0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescOnce sync.Once
	file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescData = file_api_v1alpha1_scorecards_evaluation_question_proto_rawDesc
)

func file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescData)
	})
	return file_api_v1alpha1_scorecards_evaluation_question_proto_rawDescData
}

var file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1alpha1_scorecards_evaluation_question_proto_goTypes = []interface{}{
	(*CreateEvaluationQuestionRequest)(nil),  // 0: api.v1alpha1.scorecards.CreateEvaluationQuestionRequest
	(*CreateEvaluationQuestionResponse)(nil), // 1: api.v1alpha1.scorecards.CreateEvaluationQuestionResponse
	(*UpdateEvaluationQuestionRequest)(nil),  // 2: api.v1alpha1.scorecards.UpdateEvaluationQuestionRequest
	(*UpdateEvaluationQuestionResponse)(nil), // 3: api.v1alpha1.scorecards.UpdateEvaluationQuestionResponse
	(*DeleteEvaluationQuestionRequest)(nil),  // 4: api.v1alpha1.scorecards.DeleteEvaluationQuestionRequest
	(*DeleteEvaluationQuestionResponse)(nil), // 5: api.v1alpha1.scorecards.DeleteEvaluationQuestionResponse
	(*commons.EvaluationQuestion)(nil),       // 6: api.commons.EvaluationQuestion
	(*fieldmaskpb.FieldMask)(nil),            // 7: google.protobuf.FieldMask
}
var file_api_v1alpha1_scorecards_evaluation_question_proto_depIdxs = []int32{
	6, // 0: api.v1alpha1.scorecards.CreateEvaluationQuestionRequest.evaluation_question:type_name -> api.commons.EvaluationQuestion
	6, // 1: api.v1alpha1.scorecards.CreateEvaluationQuestionResponse.evaluation_question:type_name -> api.commons.EvaluationQuestion
	6, // 2: api.v1alpha1.scorecards.UpdateEvaluationQuestionRequest.evaluation_question:type_name -> api.commons.EvaluationQuestion
	7, // 3: api.v1alpha1.scorecards.UpdateEvaluationQuestionRequest.update_mask:type_name -> google.protobuf.FieldMask
	6, // 4: api.v1alpha1.scorecards.UpdateEvaluationQuestionResponse.evaluation_question:type_name -> api.commons.EvaluationQuestion
	6, // 5: api.v1alpha1.scorecards.DeleteEvaluationQuestionResponse.evaluation_question:type_name -> api.commons.EvaluationQuestion
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_scorecards_evaluation_question_proto_init() }
func file_api_v1alpha1_scorecards_evaluation_question_proto_init() {
	if File_api_v1alpha1_scorecards_evaluation_question_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEvaluationQuestionRequest); i {
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
		file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEvaluationQuestionResponse); i {
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
		file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEvaluationQuestionRequest); i {
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
		file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEvaluationQuestionResponse); i {
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
		file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEvaluationQuestionRequest); i {
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
		file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEvaluationQuestionResponse); i {
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
			RawDescriptor: file_api_v1alpha1_scorecards_evaluation_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_scorecards_evaluation_question_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_scorecards_evaluation_question_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_scorecards_evaluation_question_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_scorecards_evaluation_question_proto = out.File
	file_api_v1alpha1_scorecards_evaluation_question_proto_rawDesc = nil
	file_api_v1alpha1_scorecards_evaluation_question_proto_goTypes = nil
	file_api_v1alpha1_scorecards_evaluation_question_proto_depIdxs = nil
}
