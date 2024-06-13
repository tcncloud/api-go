// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/v1alpha1/scorecards/scorecard.proto

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

// CreateScorecardRequest is the request to create a scorecard.
type CreateScorecardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scorecard *commons.Scorecard `protobuf:"bytes,1,opt,name=scorecard,proto3" json:"scorecard,omitempty"`
}

func (x *CreateScorecardRequest) Reset() {
	*x = CreateScorecardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScorecardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScorecardRequest) ProtoMessage() {}

func (x *CreateScorecardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScorecardRequest.ProtoReflect.Descriptor instead.
func (*CreateScorecardRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP(), []int{0}
}

func (x *CreateScorecardRequest) GetScorecard() *commons.Scorecard {
	if x != nil {
		return x.Scorecard
	}
	return nil
}

// CreateScorecardResponse contains the scorecard that was created.
type CreateScorecardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scorecard *commons.Scorecard `protobuf:"bytes,1,opt,name=scorecard,proto3" json:"scorecard,omitempty"`
}

func (x *CreateScorecardResponse) Reset() {
	*x = CreateScorecardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScorecardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScorecardResponse) ProtoMessage() {}

func (x *CreateScorecardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScorecardResponse.ProtoReflect.Descriptor instead.
func (*CreateScorecardResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP(), []int{1}
}

func (x *CreateScorecardResponse) GetScorecard() *commons.Scorecard {
	if x != nil {
		return x.Scorecard
	}
	return nil
}

// ListScorecardsRequest is the request for listing scorecards by certain criteria.
type ListScorecardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorIds       []string                 `protobuf:"bytes,2,rep,name=author_ids,json=authorIds,proto3" json:"author_ids,omitempty"`                                                           // Optional. Results include any of the authors
	CategoryIds     []int64                  `protobuf:"varint,3,rep,packed,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"`                                             // Optional. Results include any of the categories listed.
	States          []commons.ScorecardState `protobuf:"varint,4,rep,packed,name=states,proto3,enum=api.commons.ScorecardState" json:"states,omitempty"`                                          // Optional. Results include scorecard in any of the given states.
	EvaluationTypes []commons.EvaluationType `protobuf:"varint,5,rep,packed,name=evaluation_types,json=evaluationTypes,proto3,enum=api.commons.EvaluationType" json:"evaluation_types,omitempty"` // Optional. Results include scorecard with any of the given types
	CallTypes       []commons.CallType_Enum  `protobuf:"varint,6,rep,packed,name=call_types,json=callTypes,proto3,enum=api.commons.CallType_Enum" json:"call_types,omitempty"`                    // Optional. Results include scorecard with any of the given call types.
}

func (x *ListScorecardsRequest) Reset() {
	*x = ListScorecardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListScorecardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScorecardsRequest) ProtoMessage() {}

func (x *ListScorecardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScorecardsRequest.ProtoReflect.Descriptor instead.
func (*ListScorecardsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP(), []int{2}
}

func (x *ListScorecardsRequest) GetAuthorIds() []string {
	if x != nil {
		return x.AuthorIds
	}
	return nil
}

func (x *ListScorecardsRequest) GetCategoryIds() []int64 {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

func (x *ListScorecardsRequest) GetStates() []commons.ScorecardState {
	if x != nil {
		return x.States
	}
	return nil
}

func (x *ListScorecardsRequest) GetEvaluationTypes() []commons.EvaluationType {
	if x != nil {
		return x.EvaluationTypes
	}
	return nil
}

func (x *ListScorecardsRequest) GetCallTypes() []commons.CallType_Enum {
	if x != nil {
		return x.CallTypes
	}
	return nil
}

// ListScorecardsResponse includes a list of scorecards.
type ListScorecardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scorecards []*commons.Scorecard `protobuf:"bytes,1,rep,name=scorecards,proto3" json:"scorecards,omitempty"`
}

func (x *ListScorecardsResponse) Reset() {
	*x = ListScorecardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListScorecardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScorecardsResponse) ProtoMessage() {}

func (x *ListScorecardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScorecardsResponse.ProtoReflect.Descriptor instead.
func (*ListScorecardsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP(), []int{3}
}

func (x *ListScorecardsResponse) GetScorecards() []*commons.Scorecard {
	if x != nil {
		return x.Scorecards
	}
	return nil
}

// UpdateScorecardRequest is the request for updating a scorecard.
type UpdateScorecardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	Scorecard *commons.Scorecard `protobuf:"bytes,1,opt,name=scorecard,proto3" json:"scorecard,omitempty"`
	// Required. Valid paths: [title, description, pass_score, score_type,
	// evaluation_type, allow_feedback, distribute_weights, category.category_id].
	// To update Sections, use *Section and *ScorecardQuestion methods.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateScorecardRequest) Reset() {
	*x = UpdateScorecardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScorecardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScorecardRequest) ProtoMessage() {}

func (x *UpdateScorecardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScorecardRequest.ProtoReflect.Descriptor instead.
func (*UpdateScorecardRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateScorecardRequest) GetScorecard() *commons.Scorecard {
	if x != nil {
		return x.Scorecard
	}
	return nil
}

func (x *UpdateScorecardRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// UpdateScorecardResponse returns the updated scorecard.
type UpdateScorecardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scorecard *commons.Scorecard `protobuf:"bytes,1,opt,name=scorecard,proto3" json:"scorecard,omitempty"`
}

func (x *UpdateScorecardResponse) Reset() {
	*x = UpdateScorecardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScorecardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScorecardResponse) ProtoMessage() {}

func (x *UpdateScorecardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScorecardResponse.ProtoReflect.Descriptor instead.
func (*UpdateScorecardResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateScorecardResponse) GetScorecard() *commons.Scorecard {
	if x != nil {
		return x.Scorecard
	}
	return nil
}

// DeleteScorecardRequest is the request to delete a scorecard.
type DeleteScorecardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorecardId int64 `protobuf:"varint,2,opt,name=scorecard_id,json=scorecardId,proto3" json:"scorecard_id,omitempty"` // Required.
}

func (x *DeleteScorecardRequest) Reset() {
	*x = DeleteScorecardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScorecardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScorecardRequest) ProtoMessage() {}

func (x *DeleteScorecardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScorecardRequest.ProtoReflect.Descriptor instead.
func (*DeleteScorecardRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteScorecardRequest) GetScorecardId() int64 {
	if x != nil {
		return x.ScorecardId
	}
	return 0
}

// DeleteScorecardResponse contains the deleted scorecard.
type DeleteScorecardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scorecard *commons.Scorecard `protobuf:"bytes,1,opt,name=scorecard,proto3" json:"scorecard,omitempty"`
}

func (x *DeleteScorecardResponse) Reset() {
	*x = DeleteScorecardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScorecardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScorecardResponse) ProtoMessage() {}

func (x *DeleteScorecardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScorecardResponse.ProtoReflect.Descriptor instead.
func (*DeleteScorecardResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteScorecardResponse) GetScorecard() *commons.Scorecard {
	if x != nil {
		return x.Scorecard
	}
	return nil
}

// GetScorecardRequest is the request for getting a scorecard.
type GetScorecardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScorecardId int64 `protobuf:"varint,2,opt,name=scorecard_id,json=scorecardId,proto3" json:"scorecard_id,omitempty"` // Required.
	UseDefault  bool  `protobuf:"varint,3,opt,name=use_default,json=useDefault,proto3" json:"use_default,omitempty"`    // if true, will ignore id and return default
}

func (x *GetScorecardRequest) Reset() {
	*x = GetScorecardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScorecardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScorecardRequest) ProtoMessage() {}

func (x *GetScorecardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScorecardRequest.ProtoReflect.Descriptor instead.
func (*GetScorecardRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP(), []int{8}
}

func (x *GetScorecardRequest) GetScorecardId() int64 {
	if x != nil {
		return x.ScorecardId
	}
	return 0
}

func (x *GetScorecardRequest) GetUseDefault() bool {
	if x != nil {
		return x.UseDefault
	}
	return false
}

// GetScorecardResponse contains a scorecard.
type GetScorecardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scorecard *commons.Scorecard `protobuf:"bytes,1,opt,name=scorecard,proto3" json:"scorecard,omitempty"`
}

func (x *GetScorecardResponse) Reset() {
	*x = GetScorecardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScorecardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScorecardResponse) ProtoMessage() {}

func (x *GetScorecardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScorecardResponse.ProtoReflect.Descriptor instead.
func (*GetScorecardResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP(), []int{9}
}

func (x *GetScorecardResponse) GetScorecard() *commons.Scorecard {
	if x != nil {
		return x.Scorecard
	}
	return nil
}

// ListScorecardsByOrgIdRequest is the request for listing scorecards by certain criteria.
type ListScorecardsByOrgIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgId           string                   `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`                                                                       // Required. Specifies the org in which to list scorecards.
	AuthorIds       []string                 `protobuf:"bytes,2,rep,name=author_ids,json=authorIds,proto3" json:"author_ids,omitempty"`                                                           // Optional. Results include any of the authors
	CategoryIds     []int64                  `protobuf:"varint,3,rep,packed,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"`                                             // Optional. Results include any of the categories listed.
	States          []commons.ScorecardState `protobuf:"varint,4,rep,packed,name=states,proto3,enum=api.commons.ScorecardState" json:"states,omitempty"`                                          // Optional. Results include scorecard in any of the given states.
	EvaluationTypes []commons.EvaluationType `protobuf:"varint,5,rep,packed,name=evaluation_types,json=evaluationTypes,proto3,enum=api.commons.EvaluationType" json:"evaluation_types,omitempty"` // Optional. Results include scorecard with any of the given types
	CallTypes       []commons.CallType_Enum  `protobuf:"varint,6,rep,packed,name=call_types,json=callTypes,proto3,enum=api.commons.CallType_Enum" json:"call_types,omitempty"`                    // Optional. Results include scorecard with any of the given call types.
}

func (x *ListScorecardsByOrgIdRequest) Reset() {
	*x = ListScorecardsByOrgIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListScorecardsByOrgIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScorecardsByOrgIdRequest) ProtoMessage() {}

func (x *ListScorecardsByOrgIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScorecardsByOrgIdRequest.ProtoReflect.Descriptor instead.
func (*ListScorecardsByOrgIdRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP(), []int{10}
}

func (x *ListScorecardsByOrgIdRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *ListScorecardsByOrgIdRequest) GetAuthorIds() []string {
	if x != nil {
		return x.AuthorIds
	}
	return nil
}

func (x *ListScorecardsByOrgIdRequest) GetCategoryIds() []int64 {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

func (x *ListScorecardsByOrgIdRequest) GetStates() []commons.ScorecardState {
	if x != nil {
		return x.States
	}
	return nil
}

func (x *ListScorecardsByOrgIdRequest) GetEvaluationTypes() []commons.EvaluationType {
	if x != nil {
		return x.EvaluationTypes
	}
	return nil
}

func (x *ListScorecardsByOrgIdRequest) GetCallTypes() []commons.CallType_Enum {
	if x != nil {
		return x.CallTypes
	}
	return nil
}

var File_api_v1alpha1_scorecards_scorecard_proto protoreflect.FileDescriptor

var file_api_v1alpha1_scorecards_scorecard_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x63, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x52, 0x09,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x22, 0x4f, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x52,
	0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x22, 0x91, 0x02, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x10, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x52, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x50,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x72, 0x64, 0x52, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x22, 0x8b, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72,
	0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x4f,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x72, 0x64, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x22,
	0x3b, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x17,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x72, 0x64, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x22, 0x59, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x73,
	0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x4c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x52, 0x09, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x22, 0xaf, 0x02, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x73,
	0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x10, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x43, 0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x63,
	0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x42, 0xdf, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x42, 0x0e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
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
	file_api_v1alpha1_scorecards_scorecard_proto_rawDescOnce sync.Once
	file_api_v1alpha1_scorecards_scorecard_proto_rawDescData = file_api_v1alpha1_scorecards_scorecard_proto_rawDesc
)

func file_api_v1alpha1_scorecards_scorecard_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_scorecards_scorecard_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_scorecards_scorecard_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_scorecards_scorecard_proto_rawDescData)
	})
	return file_api_v1alpha1_scorecards_scorecard_proto_rawDescData
}

var file_api_v1alpha1_scorecards_scorecard_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_v1alpha1_scorecards_scorecard_proto_goTypes = []any{
	(*CreateScorecardRequest)(nil),       // 0: api.v1alpha1.scorecards.CreateScorecardRequest
	(*CreateScorecardResponse)(nil),      // 1: api.v1alpha1.scorecards.CreateScorecardResponse
	(*ListScorecardsRequest)(nil),        // 2: api.v1alpha1.scorecards.ListScorecardsRequest
	(*ListScorecardsResponse)(nil),       // 3: api.v1alpha1.scorecards.ListScorecardsResponse
	(*UpdateScorecardRequest)(nil),       // 4: api.v1alpha1.scorecards.UpdateScorecardRequest
	(*UpdateScorecardResponse)(nil),      // 5: api.v1alpha1.scorecards.UpdateScorecardResponse
	(*DeleteScorecardRequest)(nil),       // 6: api.v1alpha1.scorecards.DeleteScorecardRequest
	(*DeleteScorecardResponse)(nil),      // 7: api.v1alpha1.scorecards.DeleteScorecardResponse
	(*GetScorecardRequest)(nil),          // 8: api.v1alpha1.scorecards.GetScorecardRequest
	(*GetScorecardResponse)(nil),         // 9: api.v1alpha1.scorecards.GetScorecardResponse
	(*ListScorecardsByOrgIdRequest)(nil), // 10: api.v1alpha1.scorecards.ListScorecardsByOrgIdRequest
	(*commons.Scorecard)(nil),            // 11: api.commons.Scorecard
	(commons.ScorecardState)(0),          // 12: api.commons.ScorecardState
	(commons.EvaluationType)(0),          // 13: api.commons.EvaluationType
	(commons.CallType_Enum)(0),           // 14: api.commons.CallType.Enum
	(*fieldmaskpb.FieldMask)(nil),        // 15: google.protobuf.FieldMask
}
var file_api_v1alpha1_scorecards_scorecard_proto_depIdxs = []int32{
	11, // 0: api.v1alpha1.scorecards.CreateScorecardRequest.scorecard:type_name -> api.commons.Scorecard
	11, // 1: api.v1alpha1.scorecards.CreateScorecardResponse.scorecard:type_name -> api.commons.Scorecard
	12, // 2: api.v1alpha1.scorecards.ListScorecardsRequest.states:type_name -> api.commons.ScorecardState
	13, // 3: api.v1alpha1.scorecards.ListScorecardsRequest.evaluation_types:type_name -> api.commons.EvaluationType
	14, // 4: api.v1alpha1.scorecards.ListScorecardsRequest.call_types:type_name -> api.commons.CallType.Enum
	11, // 5: api.v1alpha1.scorecards.ListScorecardsResponse.scorecards:type_name -> api.commons.Scorecard
	11, // 6: api.v1alpha1.scorecards.UpdateScorecardRequest.scorecard:type_name -> api.commons.Scorecard
	15, // 7: api.v1alpha1.scorecards.UpdateScorecardRequest.update_mask:type_name -> google.protobuf.FieldMask
	11, // 8: api.v1alpha1.scorecards.UpdateScorecardResponse.scorecard:type_name -> api.commons.Scorecard
	11, // 9: api.v1alpha1.scorecards.DeleteScorecardResponse.scorecard:type_name -> api.commons.Scorecard
	11, // 10: api.v1alpha1.scorecards.GetScorecardResponse.scorecard:type_name -> api.commons.Scorecard
	12, // 11: api.v1alpha1.scorecards.ListScorecardsByOrgIdRequest.states:type_name -> api.commons.ScorecardState
	13, // 12: api.v1alpha1.scorecards.ListScorecardsByOrgIdRequest.evaluation_types:type_name -> api.commons.EvaluationType
	14, // 13: api.v1alpha1.scorecards.ListScorecardsByOrgIdRequest.call_types:type_name -> api.commons.CallType.Enum
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_scorecards_scorecard_proto_init() }
func file_api_v1alpha1_scorecards_scorecard_proto_init() {
	if File_api_v1alpha1_scorecards_scorecard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateScorecardRequest); i {
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
		file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateScorecardResponse); i {
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
		file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListScorecardsRequest); i {
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
		file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListScorecardsResponse); i {
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
		file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateScorecardRequest); i {
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
		file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateScorecardResponse); i {
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
		file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteScorecardRequest); i {
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
		file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteScorecardResponse); i {
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
		file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetScorecardRequest); i {
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
		file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetScorecardResponse); i {
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
		file_api_v1alpha1_scorecards_scorecard_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ListScorecardsByOrgIdRequest); i {
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
			RawDescriptor: file_api_v1alpha1_scorecards_scorecard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_scorecards_scorecard_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_scorecards_scorecard_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_scorecards_scorecard_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_scorecards_scorecard_proto = out.File
	file_api_v1alpha1_scorecards_scorecard_proto_rawDesc = nil
	file_api_v1alpha1_scorecards_scorecard_proto_goTypes = nil
	file_api_v1alpha1_scorecards_scorecard_proto_depIdxs = nil
}
