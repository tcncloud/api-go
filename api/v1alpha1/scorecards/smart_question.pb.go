// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/v1alpha1/scorecards/smart_question.proto

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

// CreateSmartQuestionRequest is the request to create a new smart question.
type CreateSmartQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartQuestion *commons.SmartQuestion `protobuf:"bytes,1,opt,name=smart_question,json=smartQuestion,proto3" json:"smart_question,omitempty"`
}

func (x *CreateSmartQuestionRequest) Reset() {
	*x = CreateSmartQuestionRequest{}
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSmartQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSmartQuestionRequest) ProtoMessage() {}

func (x *CreateSmartQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSmartQuestionRequest.ProtoReflect.Descriptor instead.
func (*CreateSmartQuestionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_question_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSmartQuestionRequest) GetSmartQuestion() *commons.SmartQuestion {
	if x != nil {
		return x.SmartQuestion
	}
	return nil
}

// CreateSmartQuestionResponse returns the created smart question.
type CreateSmartQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartQuestion *commons.SmartQuestion `protobuf:"bytes,1,opt,name=smart_question,json=smartQuestion,proto3" json:"smart_question,omitempty"`
}

func (x *CreateSmartQuestionResponse) Reset() {
	*x = CreateSmartQuestionResponse{}
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSmartQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSmartQuestionResponse) ProtoMessage() {}

func (x *CreateSmartQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSmartQuestionResponse.ProtoReflect.Descriptor instead.
func (*CreateSmartQuestionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_question_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSmartQuestionResponse) GetSmartQuestion() *commons.SmartQuestion {
	if x != nil {
		return x.SmartQuestion
	}
	return nil
}

// UpdateSmartQuestionRequest is the request to update a smart question.
type UpdateSmartQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartQuestion *commons.SmartQuestion `protobuf:"bytes,1,opt,name=smart_question,json=smartQuestion,proto3" json:"smart_question,omitempty"` // Required.
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`          // Required.
}

func (x *UpdateSmartQuestionRequest) Reset() {
	*x = UpdateSmartQuestionRequest{}
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSmartQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSmartQuestionRequest) ProtoMessage() {}

func (x *UpdateSmartQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSmartQuestionRequest.ProtoReflect.Descriptor instead.
func (*UpdateSmartQuestionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_question_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSmartQuestionRequest) GetSmartQuestion() *commons.SmartQuestion {
	if x != nil {
		return x.SmartQuestion
	}
	return nil
}

func (x *UpdateSmartQuestionRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// UpdateSmartQuestionResponse returns the updated smart question.
type UpdateSmartQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartQuestion *commons.SmartQuestion `protobuf:"bytes,1,opt,name=smart_question,json=smartQuestion,proto3" json:"smart_question,omitempty"`
}

func (x *UpdateSmartQuestionResponse) Reset() {
	*x = UpdateSmartQuestionResponse{}
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSmartQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSmartQuestionResponse) ProtoMessage() {}

func (x *UpdateSmartQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSmartQuestionResponse.ProtoReflect.Descriptor instead.
func (*UpdateSmartQuestionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_question_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSmartQuestionResponse) GetSmartQuestion() *commons.SmartQuestion {
	if x != nil {
		return x.SmartQuestion
	}
	return nil
}

// DeleteSmartQuestionRequest is the request to delete a smart question.
type DeleteSmartQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartQuestionId int64 `protobuf:"varint,2,opt,name=smart_question_id,json=smartQuestionId,proto3" json:"smart_question_id,omitempty"` // Required.
}

func (x *DeleteSmartQuestionRequest) Reset() {
	*x = DeleteSmartQuestionRequest{}
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSmartQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSmartQuestionRequest) ProtoMessage() {}

func (x *DeleteSmartQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSmartQuestionRequest.ProtoReflect.Descriptor instead.
func (*DeleteSmartQuestionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_question_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSmartQuestionRequest) GetSmartQuestionId() int64 {
	if x != nil {
		return x.SmartQuestionId
	}
	return 0
}

// DeleteSmartQuestionResponse returns the deleted smart question.
type DeleteSmartQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartQuestion *commons.SmartQuestion `protobuf:"bytes,1,opt,name=smart_question,json=smartQuestion,proto3" json:"smart_question,omitempty"`
}

func (x *DeleteSmartQuestionResponse) Reset() {
	*x = DeleteSmartQuestionResponse{}
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSmartQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSmartQuestionResponse) ProtoMessage() {}

func (x *DeleteSmartQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_question_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSmartQuestionResponse.ProtoReflect.Descriptor instead.
func (*DeleteSmartQuestionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_question_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSmartQuestionResponse) GetSmartQuestion() *commons.SmartQuestion {
	if x != nil {
		return x.SmartQuestion
	}
	return nil
}

var File_api_v1alpha1_scorecards_smart_question_proto protoreflect.FileDescriptor

var file_api_v1alpha1_scorecards_smart_question_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x6d, 0x61, 0x72,
	0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9c, 0x01, 0x0a, 0x1a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x60, 0x0a, 0x1b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53,
	0x6d, 0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x1a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x6d, 0x61, 0x72, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xe3, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x42, 0x12, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64,
	0x73, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x53, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64,
	0x73, 0xca, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0xe2, 0x02, 0x23, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x72, 0x64, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3a, 0x3a, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_scorecards_smart_question_proto_rawDescOnce sync.Once
	file_api_v1alpha1_scorecards_smart_question_proto_rawDescData = file_api_v1alpha1_scorecards_smart_question_proto_rawDesc
)

func file_api_v1alpha1_scorecards_smart_question_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_scorecards_smart_question_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_scorecards_smart_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_scorecards_smart_question_proto_rawDescData)
	})
	return file_api_v1alpha1_scorecards_smart_question_proto_rawDescData
}

var file_api_v1alpha1_scorecards_smart_question_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1alpha1_scorecards_smart_question_proto_goTypes = []any{
	(*CreateSmartQuestionRequest)(nil),  // 0: api.v1alpha1.scorecards.CreateSmartQuestionRequest
	(*CreateSmartQuestionResponse)(nil), // 1: api.v1alpha1.scorecards.CreateSmartQuestionResponse
	(*UpdateSmartQuestionRequest)(nil),  // 2: api.v1alpha1.scorecards.UpdateSmartQuestionRequest
	(*UpdateSmartQuestionResponse)(nil), // 3: api.v1alpha1.scorecards.UpdateSmartQuestionResponse
	(*DeleteSmartQuestionRequest)(nil),  // 4: api.v1alpha1.scorecards.DeleteSmartQuestionRequest
	(*DeleteSmartQuestionResponse)(nil), // 5: api.v1alpha1.scorecards.DeleteSmartQuestionResponse
	(*commons.SmartQuestion)(nil),       // 6: api.commons.SmartQuestion
	(*fieldmaskpb.FieldMask)(nil),       // 7: google.protobuf.FieldMask
}
var file_api_v1alpha1_scorecards_smart_question_proto_depIdxs = []int32{
	6, // 0: api.v1alpha1.scorecards.CreateSmartQuestionRequest.smart_question:type_name -> api.commons.SmartQuestion
	6, // 1: api.v1alpha1.scorecards.CreateSmartQuestionResponse.smart_question:type_name -> api.commons.SmartQuestion
	6, // 2: api.v1alpha1.scorecards.UpdateSmartQuestionRequest.smart_question:type_name -> api.commons.SmartQuestion
	7, // 3: api.v1alpha1.scorecards.UpdateSmartQuestionRequest.update_mask:type_name -> google.protobuf.FieldMask
	6, // 4: api.v1alpha1.scorecards.UpdateSmartQuestionResponse.smart_question:type_name -> api.commons.SmartQuestion
	6, // 5: api.v1alpha1.scorecards.DeleteSmartQuestionResponse.smart_question:type_name -> api.commons.SmartQuestion
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_scorecards_smart_question_proto_init() }
func file_api_v1alpha1_scorecards_smart_question_proto_init() {
	if File_api_v1alpha1_scorecards_smart_question_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_scorecards_smart_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_scorecards_smart_question_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_scorecards_smart_question_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_scorecards_smart_question_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_scorecards_smart_question_proto = out.File
	file_api_v1alpha1_scorecards_smart_question_proto_rawDesc = nil
	file_api_v1alpha1_scorecards_smart_question_proto_goTypes = nil
	file_api_v1alpha1_scorecards_smart_question_proto_depIdxs = nil
}
