// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/v1alpha1/scorecards/smart_evaluation.proto

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

// CreateSmartEvaluationRequest represents a request to create a smart evaluation.
type CreateSmartEvaluationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartEvaluation *commons.SmartEvaluation `protobuf:"bytes,1,opt,name=smart_evaluation,json=smartEvaluation,proto3" json:"smart_evaluation,omitempty"`
}

func (x *CreateSmartEvaluationRequest) Reset() {
	*x = CreateSmartEvaluationRequest{}
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSmartEvaluationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSmartEvaluationRequest) ProtoMessage() {}

func (x *CreateSmartEvaluationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSmartEvaluationRequest.ProtoReflect.Descriptor instead.
func (*CreateSmartEvaluationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSmartEvaluationRequest) GetSmartEvaluation() *commons.SmartEvaluation {
	if x != nil {
		return x.SmartEvaluation
	}
	return nil
}

// CreateSmartEvaluationResponse represents a response to create a smart evaluation.
type CreateSmartEvaluationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartEvaluation *commons.SmartEvaluation `protobuf:"bytes,1,opt,name=smart_evaluation,json=smartEvaluation,proto3" json:"smart_evaluation,omitempty"`
}

func (x *CreateSmartEvaluationResponse) Reset() {
	*x = CreateSmartEvaluationResponse{}
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSmartEvaluationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSmartEvaluationResponse) ProtoMessage() {}

func (x *CreateSmartEvaluationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSmartEvaluationResponse.ProtoReflect.Descriptor instead.
func (*CreateSmartEvaluationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSmartEvaluationResponse) GetSmartEvaluation() *commons.SmartEvaluation {
	if x != nil {
		return x.SmartEvaluation
	}
	return nil
}

// ListSmartEvaluationsRequest represents a request to list smart evaluations.
type ListSmartEvaluationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The order by which smart evaluations will be listed. Follows sql
	// order by syntax. Defaults to "complete_time desc, smart_evaluation_id desc".
	OrderBy string `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// Optional - number of smart evaluations included in response. Defaults to 500.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Required. Fields to return.
	//
	// Example selecting score, section points, and question answer:
	//
	//	{
	//	  paths: [
	//	    "score",
	//	    "smart_evaluation_section.points",
	//	    "smart_evaluation_section.smart_evaluation_question.answer"
	//	  ]
	//	}
	ReturnFields *fieldmaskpb.FieldMask `protobuf:"bytes,5,opt,name=return_fields,json=returnFields,proto3" json:"return_fields,omitempty"`
	// Optional. The standard list filter as described in https://google.aip.dev/160.
	// Multiple comparisons can be provided when separated with a logical AND
	// operator. Supported fields, operators and functions are listed below.
	//
	// +----------------------------+-----------+-----------------+-----------+
	// |                      field |      type |       operators | functions |
	// +----------------------------+-----------+-----------------+-----------+
	// |               scorecard_id |   integer |               = |       any |
	// |                category_id |   integer |               = |       any |
	// |              agent_user_id |    string |               = |       any |
	// |             transcript_sid |   integer | =, >=, <=, >, < |           |
	// |              complete_time | timestamp | =, >=, <=, >, < |           |
	// +----------------------------+-----------+-----------------+-----------+
	//
	// Examples:
	// transcript_sid >= 1
	//
	// scorecard_id = 0 AND
	// agent_user_id = any("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001") AND
	// complete_time >= '2012-04-21T11:30:00-04:00'
	Filter string `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListSmartEvaluationsRequest) Reset() {
	*x = ListSmartEvaluationsRequest{}
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSmartEvaluationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSmartEvaluationsRequest) ProtoMessage() {}

func (x *ListSmartEvaluationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSmartEvaluationsRequest.ProtoReflect.Descriptor instead.
func (*ListSmartEvaluationsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescGZIP(), []int{2}
}

func (x *ListSmartEvaluationsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListSmartEvaluationsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSmartEvaluationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListSmartEvaluationsRequest) GetReturnFields() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReturnFields
	}
	return nil
}

func (x *ListSmartEvaluationsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// ListSmartEvaluationsResponse represents a response to list smart evaluations.
type ListSmartEvaluationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartEvaluations []*commons.SmartEvaluation `protobuf:"bytes,1,rep,name=smart_evaluations,json=smartEvaluations,proto3" json:"smart_evaluations,omitempty"`
	// Token to retrieve the next page, or empty if there are no
	// more entries in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListSmartEvaluationsResponse) Reset() {
	*x = ListSmartEvaluationsResponse{}
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSmartEvaluationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSmartEvaluationsResponse) ProtoMessage() {}

func (x *ListSmartEvaluationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSmartEvaluationsResponse.ProtoReflect.Descriptor instead.
func (*ListSmartEvaluationsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescGZIP(), []int{3}
}

func (x *ListSmartEvaluationsResponse) GetSmartEvaluations() []*commons.SmartEvaluation {
	if x != nil {
		return x.SmartEvaluations
	}
	return nil
}

func (x *ListSmartEvaluationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// UpdateSmartEvaluationRequest represents a request to update a smart evaluation.
type UpdateSmartEvaluationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartEvaluation *commons.SmartEvaluation `protobuf:"bytes,1,opt,name=smart_evaluation,json=smartEvaluation,proto3" json:"smart_evaluation,omitempty"` // Required - the smart evaluation  to update.
	UpdateMask      *fieldmaskpb.FieldMask   `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`                // Required - specification of which fields should be updated.
}

func (x *UpdateSmartEvaluationRequest) Reset() {
	*x = UpdateSmartEvaluationRequest{}
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSmartEvaluationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSmartEvaluationRequest) ProtoMessage() {}

func (x *UpdateSmartEvaluationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSmartEvaluationRequest.ProtoReflect.Descriptor instead.
func (*UpdateSmartEvaluationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSmartEvaluationRequest) GetSmartEvaluation() *commons.SmartEvaluation {
	if x != nil {
		return x.SmartEvaluation
	}
	return nil
}

func (x *UpdateSmartEvaluationRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// UpdateSmartEvaluationResponse represents a response to update a smart evaluation.
type UpdateSmartEvaluationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartEvaluation *commons.SmartEvaluation `protobuf:"bytes,1,opt,name=smart_evaluation,json=smartEvaluation,proto3" json:"smart_evaluation,omitempty"`
}

func (x *UpdateSmartEvaluationResponse) Reset() {
	*x = UpdateSmartEvaluationResponse{}
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSmartEvaluationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSmartEvaluationResponse) ProtoMessage() {}

func (x *UpdateSmartEvaluationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSmartEvaluationResponse.ProtoReflect.Descriptor instead.
func (*UpdateSmartEvaluationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSmartEvaluationResponse) GetSmartEvaluation() *commons.SmartEvaluation {
	if x != nil {
		return x.SmartEvaluation
	}
	return nil
}

// DeleteSmartEvaluationRequest represents a request to delete a smart evaluation.
type DeleteSmartEvaluationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartEvaluationId int64 `protobuf:"varint,2,opt,name=smart_evaluation_id,json=smartEvaluationId,proto3" json:"smart_evaluation_id,omitempty"`
}

func (x *DeleteSmartEvaluationRequest) Reset() {
	*x = DeleteSmartEvaluationRequest{}
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSmartEvaluationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSmartEvaluationRequest) ProtoMessage() {}

func (x *DeleteSmartEvaluationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSmartEvaluationRequest.ProtoReflect.Descriptor instead.
func (*DeleteSmartEvaluationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSmartEvaluationRequest) GetSmartEvaluationId() int64 {
	if x != nil {
		return x.SmartEvaluationId
	}
	return 0
}

// DeleteSmartEvaluationResponse represents a response to delete a smart evaluation.
type DeleteSmartEvaluationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSmartEvaluationResponse) Reset() {
	*x = DeleteSmartEvaluationResponse{}
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSmartEvaluationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSmartEvaluationResponse) ProtoMessage() {}

func (x *DeleteSmartEvaluationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSmartEvaluationResponse.ProtoReflect.Descriptor instead.
func (*DeleteSmartEvaluationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescGZIP(), []int{7}
}

// GetSmartEvaluationRequest represents a request to get a smart evaluation.
type GetSmartEvaluationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartEvaluationId int64 `protobuf:"varint,2,opt,name=smart_evaluation_id,json=smartEvaluationId,proto3" json:"smart_evaluation_id,omitempty"`
}

func (x *GetSmartEvaluationRequest) Reset() {
	*x = GetSmartEvaluationRequest{}
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSmartEvaluationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSmartEvaluationRequest) ProtoMessage() {}

func (x *GetSmartEvaluationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSmartEvaluationRequest.ProtoReflect.Descriptor instead.
func (*GetSmartEvaluationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescGZIP(), []int{8}
}

func (x *GetSmartEvaluationRequest) GetSmartEvaluationId() int64 {
	if x != nil {
		return x.SmartEvaluationId
	}
	return 0
}

// GetSmartEvaluationResponse represents a response to get a smart evaluation.
type GetSmartEvaluationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartEvaluation *commons.SmartEvaluation `protobuf:"bytes,1,opt,name=smart_evaluation,json=smartEvaluation,proto3" json:"smart_evaluation,omitempty"`
}

func (x *GetSmartEvaluationResponse) Reset() {
	*x = GetSmartEvaluationResponse{}
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSmartEvaluationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSmartEvaluationResponse) ProtoMessage() {}

func (x *GetSmartEvaluationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSmartEvaluationResponse.ProtoReflect.Descriptor instead.
func (*GetSmartEvaluationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescGZIP(), []int{9}
}

func (x *GetSmartEvaluationResponse) GetSmartEvaluation() *commons.SmartEvaluation {
	if x != nil {
		return x.SmartEvaluation
	}
	return nil
}

var File_api_v1alpha1_scorecards_smart_evaluation_proto protoreflect.FileDescriptor

var file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x1c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x10, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x68, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72,
	0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x6d, 0x61, 0x72,
	0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xcd, 0x01, 0x0a,
	0x1b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x91, 0x01, 0x0a,
	0x1c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x11, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0xa4, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x47, 0x0a, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x68, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x4e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x4b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x13, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x65, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a,
	0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xe5, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x72, 0x64, 0x73, 0x42, 0x14, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x45, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32,
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
	file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescOnce sync.Once
	file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescData = file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDesc
)

func file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescData)
	})
	return file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDescData
}

var file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v1alpha1_scorecards_smart_evaluation_proto_goTypes = []any{
	(*CreateSmartEvaluationRequest)(nil),  // 0: api.v1alpha1.scorecards.CreateSmartEvaluationRequest
	(*CreateSmartEvaluationResponse)(nil), // 1: api.v1alpha1.scorecards.CreateSmartEvaluationResponse
	(*ListSmartEvaluationsRequest)(nil),   // 2: api.v1alpha1.scorecards.ListSmartEvaluationsRequest
	(*ListSmartEvaluationsResponse)(nil),  // 3: api.v1alpha1.scorecards.ListSmartEvaluationsResponse
	(*UpdateSmartEvaluationRequest)(nil),  // 4: api.v1alpha1.scorecards.UpdateSmartEvaluationRequest
	(*UpdateSmartEvaluationResponse)(nil), // 5: api.v1alpha1.scorecards.UpdateSmartEvaluationResponse
	(*DeleteSmartEvaluationRequest)(nil),  // 6: api.v1alpha1.scorecards.DeleteSmartEvaluationRequest
	(*DeleteSmartEvaluationResponse)(nil), // 7: api.v1alpha1.scorecards.DeleteSmartEvaluationResponse
	(*GetSmartEvaluationRequest)(nil),     // 8: api.v1alpha1.scorecards.GetSmartEvaluationRequest
	(*GetSmartEvaluationResponse)(nil),    // 9: api.v1alpha1.scorecards.GetSmartEvaluationResponse
	(*commons.SmartEvaluation)(nil),       // 10: api.commons.SmartEvaluation
	(*fieldmaskpb.FieldMask)(nil),         // 11: google.protobuf.FieldMask
}
var file_api_v1alpha1_scorecards_smart_evaluation_proto_depIdxs = []int32{
	10, // 0: api.v1alpha1.scorecards.CreateSmartEvaluationRequest.smart_evaluation:type_name -> api.commons.SmartEvaluation
	10, // 1: api.v1alpha1.scorecards.CreateSmartEvaluationResponse.smart_evaluation:type_name -> api.commons.SmartEvaluation
	11, // 2: api.v1alpha1.scorecards.ListSmartEvaluationsRequest.return_fields:type_name -> google.protobuf.FieldMask
	10, // 3: api.v1alpha1.scorecards.ListSmartEvaluationsResponse.smart_evaluations:type_name -> api.commons.SmartEvaluation
	10, // 4: api.v1alpha1.scorecards.UpdateSmartEvaluationRequest.smart_evaluation:type_name -> api.commons.SmartEvaluation
	11, // 5: api.v1alpha1.scorecards.UpdateSmartEvaluationRequest.update_mask:type_name -> google.protobuf.FieldMask
	10, // 6: api.v1alpha1.scorecards.UpdateSmartEvaluationResponse.smart_evaluation:type_name -> api.commons.SmartEvaluation
	10, // 7: api.v1alpha1.scorecards.GetSmartEvaluationResponse.smart_evaluation:type_name -> api.commons.SmartEvaluation
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_scorecards_smart_evaluation_proto_init() }
func file_api_v1alpha1_scorecards_smart_evaluation_proto_init() {
	if File_api_v1alpha1_scorecards_smart_evaluation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_scorecards_smart_evaluation_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_scorecards_smart_evaluation_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_scorecards_smart_evaluation_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_scorecards_smart_evaluation_proto = out.File
	file_api_v1alpha1_scorecards_smart_evaluation_proto_rawDesc = nil
	file_api_v1alpha1_scorecards_smart_evaluation_proto_goTypes = nil
	file_api_v1alpha1_scorecards_smart_evaluation_proto_depIdxs = nil
}
