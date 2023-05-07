// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/v1alpha1/scorecards/evaluation.proto

package scorecards

import (
	commons "github.com/tcncloud/api-go/api/commons"
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

// CreateEvaluationRequest is request to create a new evaluation
type CreateEvaluationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evaluation *commons.Evaluation `protobuf:"bytes,1,opt,name=evaluation,proto3" json:"evaluation,omitempty"` // entity to create
}

func (x *CreateEvaluationRequest) Reset() {
	*x = CreateEvaluationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEvaluationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEvaluationRequest) ProtoMessage() {}

func (x *CreateEvaluationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEvaluationRequest.ProtoReflect.Descriptor instead.
func (*CreateEvaluationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEvaluationRequest) GetEvaluation() *commons.Evaluation {
	if x != nil {
		return x.Evaluation
	}
	return nil
}

// CreateEvaluationResponse is response with created entity
type CreateEvaluationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evaluation *commons.Evaluation `protobuf:"bytes,1,opt,name=evaluation,proto3" json:"evaluation,omitempty"` // entity created
}

func (x *CreateEvaluationResponse) Reset() {
	*x = CreateEvaluationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEvaluationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEvaluationResponse) ProtoMessage() {}

func (x *CreateEvaluationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEvaluationResponse.ProtoReflect.Descriptor instead.
func (*CreateEvaluationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEvaluationResponse) GetEvaluation() *commons.Evaluation {
	if x != nil {
		return x.Evaluation
	}
	return nil
}

// DeleteEvaluationRequest is request to delete an evaluation
type DeleteEvaluationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvaluationId int64 `protobuf:"varint,2,opt,name=evaluation_id,json=evaluationId,proto3" json:"evaluation_id,omitempty"` // Required - unique id of evaluation to get
}

func (x *DeleteEvaluationRequest) Reset() {
	*x = DeleteEvaluationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEvaluationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEvaluationRequest) ProtoMessage() {}

func (x *DeleteEvaluationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEvaluationRequest.ProtoReflect.Descriptor instead.
func (*DeleteEvaluationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteEvaluationRequest) GetEvaluationId() int64 {
	if x != nil {
		return x.EvaluationId
	}
	return 0
}

// DeleteEvaluationReponse is response with requested entity
type DeleteEvaluationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evaluation *commons.Evaluation `protobuf:"bytes,1,opt,name=evaluation,proto3" json:"evaluation,omitempty"` // entity deleted
}

func (x *DeleteEvaluationResponse) Reset() {
	*x = DeleteEvaluationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEvaluationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEvaluationResponse) ProtoMessage() {}

func (x *DeleteEvaluationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEvaluationResponse.ProtoReflect.Descriptor instead.
func (*DeleteEvaluationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteEvaluationResponse) GetEvaluation() *commons.Evaluation {
	if x != nil {
		return x.Evaluation
	}
	return nil
}

// GetEvaluationRequest is request to get an evaluation
type GetEvaluationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvaluationId int64 `protobuf:"varint,2,opt,name=evaluation_id,json=evaluationId,proto3" json:"evaluation_id,omitempty"` // Required - unique id of evaluation to get
}

func (x *GetEvaluationRequest) Reset() {
	*x = GetEvaluationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEvaluationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEvaluationRequest) ProtoMessage() {}

func (x *GetEvaluationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEvaluationRequest.ProtoReflect.Descriptor instead.
func (*GetEvaluationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_proto_rawDescGZIP(), []int{4}
}

func (x *GetEvaluationRequest) GetEvaluationId() int64 {
	if x != nil {
		return x.EvaluationId
	}
	return 0
}

// GetEvaluationReponse is response with requested entity
type GetEvaluationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evaluation *commons.Evaluation `protobuf:"bytes,1,opt,name=evaluation,proto3" json:"evaluation,omitempty"`
}

func (x *GetEvaluationResponse) Reset() {
	*x = GetEvaluationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEvaluationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEvaluationResponse) ProtoMessage() {}

func (x *GetEvaluationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEvaluationResponse.ProtoReflect.Descriptor instead.
func (*GetEvaluationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_proto_rawDescGZIP(), []int{5}
}

func (x *GetEvaluationResponse) GetEvaluation() *commons.Evaluation {
	if x != nil {
		return x.Evaluation
	}
	return nil
}

// ScoreEvaluationRequest is request to score an evaluation
type ScoreEvaluationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvaluationId int64 `protobuf:"varint,3,opt,name=evaluation_id,json=evaluationId,proto3" json:"evaluation_id,omitempty"` // Required - unique id of evaluation to score
}

func (x *ScoreEvaluationRequest) Reset() {
	*x = ScoreEvaluationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreEvaluationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreEvaluationRequest) ProtoMessage() {}

func (x *ScoreEvaluationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreEvaluationRequest.ProtoReflect.Descriptor instead.
func (*ScoreEvaluationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_proto_rawDescGZIP(), []int{6}
}

func (x *ScoreEvaluationRequest) GetEvaluationId() int64 {
	if x != nil {
		return x.EvaluationId
	}
	return 0
}

// ScoreEvaluationResponse returns the evaluation that was scored
type ScoreEvaluationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evaluation *commons.Evaluation `protobuf:"bytes,1,opt,name=evaluation,proto3" json:"evaluation,omitempty"`
}

func (x *ScoreEvaluationResponse) Reset() {
	*x = ScoreEvaluationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreEvaluationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreEvaluationResponse) ProtoMessage() {}

func (x *ScoreEvaluationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreEvaluationResponse.ProtoReflect.Descriptor instead.
func (*ScoreEvaluationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_proto_rawDescGZIP(), []int{7}
}

func (x *ScoreEvaluationResponse) GetEvaluation() *commons.Evaluation {
	if x != nil {
		return x.Evaluation
	}
	return nil
}

// ListEvaluationsRequest is request to get list of evaluations
type ListEvaluationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorerId     []string            `protobuf:"bytes,2,rep,name=scorer_id,json=scorerId,proto3" json:"scorer_id,omitempty"`                     // optional, list by scorer_id
	CompletedAt  *commons.TimeFilter `protobuf:"bytes,3,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`            // optional, filter completed_at by a specific range
	CategoryIds  []int64             `protobuf:"varint,4,rep,packed,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"`    // optional, list by category_ids
	AgentUserIds []string            `protobuf:"bytes,5,rep,name=agent_user_ids,json=agentUserIds,proto3" json:"agent_user_ids,omitempty"`       // Optional. List by agent user ids.
	ScorecardIds []int64             `protobuf:"varint,6,rep,packed,name=scorecard_ids,json=scorecardIds,proto3" json:"scorecard_ids,omitempty"` // Optional. List by scorecard_ids
}

func (x *ListEvaluationsRequest) Reset() {
	*x = ListEvaluationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEvaluationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEvaluationsRequest) ProtoMessage() {}

func (x *ListEvaluationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEvaluationsRequest.ProtoReflect.Descriptor instead.
func (*ListEvaluationsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_proto_rawDescGZIP(), []int{8}
}

func (x *ListEvaluationsRequest) GetScorerId() []string {
	if x != nil {
		return x.ScorerId
	}
	return nil
}

func (x *ListEvaluationsRequest) GetCompletedAt() *commons.TimeFilter {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

func (x *ListEvaluationsRequest) GetCategoryIds() []int64 {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

func (x *ListEvaluationsRequest) GetAgentUserIds() []string {
	if x != nil {
		return x.AgentUserIds
	}
	return nil
}

func (x *ListEvaluationsRequest) GetScorecardIds() []int64 {
	if x != nil {
		return x.ScorecardIds
	}
	return nil
}

// ListEvaluationsResponse returns a list of evaluations
type ListEvaluationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evaluations []*commons.Evaluation `protobuf:"bytes,1,rep,name=evaluations,proto3" json:"evaluations,omitempty"`
}

func (x *ListEvaluationsResponse) Reset() {
	*x = ListEvaluationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEvaluationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEvaluationsResponse) ProtoMessage() {}

func (x *ListEvaluationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEvaluationsResponse.ProtoReflect.Descriptor instead.
func (*ListEvaluationsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_evaluation_proto_rawDescGZIP(), []int{9}
}

func (x *ListEvaluationsResponse) GetEvaluations() []*commons.Evaluation {
	if x != nil {
		return x.Evaluations
	}
	return nil
}

var File_api_v1alpha1_scorecards_evaluation_proto protoreflect.FileDescriptor

var file_api_v1alpha1_scorecards_evaluation_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x52, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x17, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x18, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d,
	0x0a, 0x16, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x52, 0x0a,
	0x17, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xdf, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x73, 0x22, 0x54, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x0b, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x63, 0x0a, 0x26, 0x63, 0x6f, 0x6d,
	0x2e, 0x74, 0x63, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x3b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_scorecards_evaluation_proto_rawDescOnce sync.Once
	file_api_v1alpha1_scorecards_evaluation_proto_rawDescData = file_api_v1alpha1_scorecards_evaluation_proto_rawDesc
)

func file_api_v1alpha1_scorecards_evaluation_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_scorecards_evaluation_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_scorecards_evaluation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_scorecards_evaluation_proto_rawDescData)
	})
	return file_api_v1alpha1_scorecards_evaluation_proto_rawDescData
}

var file_api_v1alpha1_scorecards_evaluation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v1alpha1_scorecards_evaluation_proto_goTypes = []interface{}{
	(*CreateEvaluationRequest)(nil),  // 0: api.v1alpha1.scorecards.CreateEvaluationRequest
	(*CreateEvaluationResponse)(nil), // 1: api.v1alpha1.scorecards.CreateEvaluationResponse
	(*DeleteEvaluationRequest)(nil),  // 2: api.v1alpha1.scorecards.DeleteEvaluationRequest
	(*DeleteEvaluationResponse)(nil), // 3: api.v1alpha1.scorecards.DeleteEvaluationResponse
	(*GetEvaluationRequest)(nil),     // 4: api.v1alpha1.scorecards.GetEvaluationRequest
	(*GetEvaluationResponse)(nil),    // 5: api.v1alpha1.scorecards.GetEvaluationResponse
	(*ScoreEvaluationRequest)(nil),   // 6: api.v1alpha1.scorecards.ScoreEvaluationRequest
	(*ScoreEvaluationResponse)(nil),  // 7: api.v1alpha1.scorecards.ScoreEvaluationResponse
	(*ListEvaluationsRequest)(nil),   // 8: api.v1alpha1.scorecards.ListEvaluationsRequest
	(*ListEvaluationsResponse)(nil),  // 9: api.v1alpha1.scorecards.ListEvaluationsResponse
	(*commons.Evaluation)(nil),       // 10: api.commons.Evaluation
	(*commons.TimeFilter)(nil),       // 11: api.commons.TimeFilter
}
var file_api_v1alpha1_scorecards_evaluation_proto_depIdxs = []int32{
	10, // 0: api.v1alpha1.scorecards.CreateEvaluationRequest.evaluation:type_name -> api.commons.Evaluation
	10, // 1: api.v1alpha1.scorecards.CreateEvaluationResponse.evaluation:type_name -> api.commons.Evaluation
	10, // 2: api.v1alpha1.scorecards.DeleteEvaluationResponse.evaluation:type_name -> api.commons.Evaluation
	10, // 3: api.v1alpha1.scorecards.GetEvaluationResponse.evaluation:type_name -> api.commons.Evaluation
	10, // 4: api.v1alpha1.scorecards.ScoreEvaluationResponse.evaluation:type_name -> api.commons.Evaluation
	11, // 5: api.v1alpha1.scorecards.ListEvaluationsRequest.completed_at:type_name -> api.commons.TimeFilter
	10, // 6: api.v1alpha1.scorecards.ListEvaluationsResponse.evaluations:type_name -> api.commons.Evaluation
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_scorecards_evaluation_proto_init() }
func file_api_v1alpha1_scorecards_evaluation_proto_init() {
	if File_api_v1alpha1_scorecards_evaluation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEvaluationRequest); i {
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
		file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEvaluationResponse); i {
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
		file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEvaluationRequest); i {
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
		file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEvaluationResponse); i {
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
		file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEvaluationRequest); i {
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
		file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEvaluationResponse); i {
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
		file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreEvaluationRequest); i {
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
		file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreEvaluationResponse); i {
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
		file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEvaluationsRequest); i {
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
		file_api_v1alpha1_scorecards_evaluation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEvaluationsResponse); i {
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
			RawDescriptor: file_api_v1alpha1_scorecards_evaluation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_scorecards_evaluation_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_scorecards_evaluation_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_scorecards_evaluation_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_scorecards_evaluation_proto = out.File
	file_api_v1alpha1_scorecards_evaluation_proto_rawDesc = nil
	file_api_v1alpha1_scorecards_evaluation_proto_goTypes = nil
	file_api_v1alpha1_scorecards_evaluation_proto_depIdxs = nil
}
