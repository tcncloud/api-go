// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/v1alpha1/vanalytics/correction.proto

package vanalytics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// UpdateCorrectionRequest is a request for updating a correction.
type UpdateCorrectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The correction resource to update.
	Correction *Correction `protobuf:"bytes,1,opt,name=correction,proto3" json:"correction,omitempty"`
	// The fields to update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateCorrectionRequest) Reset() {
	*x = UpdateCorrectionRequest{}
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCorrectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCorrectionRequest) ProtoMessage() {}

func (x *UpdateCorrectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCorrectionRequest.ProtoReflect.Descriptor instead.
func (*UpdateCorrectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_correction_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateCorrectionRequest) GetCorrection() *Correction {
	if x != nil {
		return x.Correction
	}
	return nil
}

func (x *UpdateCorrectionRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// UpdateCorrectionResponse is a response for updating a correction.
type UpdateCorrectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The correction resource that was updated.
	Correction *Correction `protobuf:"bytes,1,opt,name=correction,proto3" json:"correction,omitempty"`
}

func (x *UpdateCorrectionResponse) Reset() {
	*x = UpdateCorrectionResponse{}
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCorrectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCorrectionResponse) ProtoMessage() {}

func (x *UpdateCorrectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCorrectionResponse.ProtoReflect.Descriptor instead.
func (*UpdateCorrectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_correction_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCorrectionResponse) GetCorrection() *Correction {
	if x != nil {
		return x.Correction
	}
	return nil
}

// CreateCorrectionRequest is a request for creating a correction.
type CreateCorrectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The correction resource to create.
	Correction *Correction `protobuf:"bytes,1,opt,name=correction,proto3" json:"correction,omitempty"`
}

func (x *CreateCorrectionRequest) Reset() {
	*x = CreateCorrectionRequest{}
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCorrectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCorrectionRequest) ProtoMessage() {}

func (x *CreateCorrectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCorrectionRequest.ProtoReflect.Descriptor instead.
func (*CreateCorrectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_correction_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCorrectionRequest) GetCorrection() *Correction {
	if x != nil {
		return x.Correction
	}
	return nil
}

// CreateCorrectionResponse is a response for creating a correction.
type CreateCorrectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The correction resource that was created.
	Correction *Correction `protobuf:"bytes,1,opt,name=correction,proto3" json:"correction,omitempty"`
}

func (x *CreateCorrectionResponse) Reset() {
	*x = CreateCorrectionResponse{}
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCorrectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCorrectionResponse) ProtoMessage() {}

func (x *CreateCorrectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCorrectionResponse.ProtoReflect.Descriptor instead.
func (*CreateCorrectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_correction_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCorrectionResponse) GetCorrection() *Correction {
	if x != nil {
		return x.Correction
	}
	return nil
}

type GetCorrectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The unique id of the correction.
	CorrectionSid int64 `protobuf:"varint,2,opt,name=correction_sid,json=correctionSid,proto3" json:"correction_sid,omitempty"`
}

func (x *GetCorrectionRequest) Reset() {
	*x = GetCorrectionRequest{}
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCorrectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCorrectionRequest) ProtoMessage() {}

func (x *GetCorrectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCorrectionRequest.ProtoReflect.Descriptor instead.
func (*GetCorrectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_correction_proto_rawDescGZIP(), []int{4}
}

func (x *GetCorrectionRequest) GetCorrectionSid() int64 {
	if x != nil {
		return x.CorrectionSid
	}
	return 0
}

// ListCorrectionsRequest is a request for listing corrections.
type ListCorrectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The transcript sid of the transcript that the corrections are on.
	TranscriptSid int64 `protobuf:"varint,2,opt,name=transcript_sid,json=transcriptSid,proto3" json:"transcript_sid,omitempty"`
}

func (x *ListCorrectionsRequest) Reset() {
	*x = ListCorrectionsRequest{}
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCorrectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCorrectionsRequest) ProtoMessage() {}

func (x *ListCorrectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCorrectionsRequest.ProtoReflect.Descriptor instead.
func (*ListCorrectionsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_correction_proto_rawDescGZIP(), []int{5}
}

func (x *ListCorrectionsRequest) GetTranscriptSid() int64 {
	if x != nil {
		return x.TranscriptSid
	}
	return 0
}

// ListCorrectionsResponse is a response for listing corrections.
type ListCorrectionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of corrections.
	Corrections []*Correction `protobuf:"bytes,1,rep,name=corrections,proto3" json:"corrections,omitempty"`
}

func (x *ListCorrectionsResponse) Reset() {
	*x = ListCorrectionsResponse{}
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCorrectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCorrectionsResponse) ProtoMessage() {}

func (x *ListCorrectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCorrectionsResponse.ProtoReflect.Descriptor instead.
func (*ListCorrectionsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_correction_proto_rawDescGZIP(), []int{6}
}

func (x *ListCorrectionsResponse) GetCorrections() []*Correction {
	if x != nil {
		return x.Corrections
	}
	return nil
}

type DeleteCorrectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The unique id of the correction to be deleted.
	CorrectionSid int64 `protobuf:"varint,1,opt,name=correction_sid,json=correctionSid,proto3" json:"correction_sid,omitempty"`
	// Optional. Return the correction that was deleted.
	Return bool `protobuf:"varint,3,opt,name=return,proto3" json:"return,omitempty"`
}

func (x *DeleteCorrectionRequest) Reset() {
	*x = DeleteCorrectionRequest{}
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCorrectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCorrectionRequest) ProtoMessage() {}

func (x *DeleteCorrectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCorrectionRequest.ProtoReflect.Descriptor instead.
func (*DeleteCorrectionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_correction_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCorrectionRequest) GetCorrectionSid() int64 {
	if x != nil {
		return x.CorrectionSid
	}
	return 0
}

func (x *DeleteCorrectionRequest) GetReturn() bool {
	if x != nil {
		return x.Return
	}
	return false
}

type DeleteCorrectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The correction that was deleted.
	Correction *Correction `protobuf:"bytes,1,opt,name=correction,proto3" json:"correction,omitempty"`
}

func (x *DeleteCorrectionResponse) Reset() {
	*x = DeleteCorrectionResponse{}
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCorrectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCorrectionResponse) ProtoMessage() {}

func (x *DeleteCorrectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCorrectionResponse.ProtoReflect.Descriptor instead.
func (*DeleteCorrectionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_correction_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCorrectionResponse) GetCorrection() *Correction {
	if x != nil {
		return x.Correction
	}
	return nil
}

// Correction is a resource in the Vanalytics API.
type Correction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The unique id of this correction.
	CorrectionSid int64 `protobuf:"varint,1,opt,name=correction_sid,json=correctionSid,proto3" json:"correction_sid,omitempty"`
	// Required. The transcript sid of the transcript that the correction is on.
	TranscriptSid int64 `protobuf:"varint,3,opt,name=transcript_sid,json=transcriptSid,proto3" json:"transcript_sid,omitempty"`
	// Required. Start offset of the selected words in the transcript.
	StartOffset *durationpb.Duration `protobuf:"bytes,4,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	// Required. End offset of the selected words in the transcript.
	EndOffset *durationpb.Duration `protobuf:"bytes,5,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`
	// Required. Proposed text of what the transcript selected text should be.
	ProposedText string `protobuf:"bytes,6,opt,name=proposed_text,json=proposedText,proto3" json:"proposed_text,omitempty"`
	// Required. Channel that the words were spoken on.
	Channel uint32 `protobuf:"varint,7,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *Correction) Reset() {
	*x = Correction{}
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Correction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Correction) ProtoMessage() {}

func (x *Correction) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_correction_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Correction.ProtoReflect.Descriptor instead.
func (*Correction) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_correction_proto_rawDescGZIP(), []int{9}
}

func (x *Correction) GetCorrectionSid() int64 {
	if x != nil {
		return x.CorrectionSid
	}
	return 0
}

func (x *Correction) GetTranscriptSid() int64 {
	if x != nil {
		return x.TranscriptSid
	}
	return 0
}

func (x *Correction) GetStartOffset() *durationpb.Duration {
	if x != nil {
		return x.StartOffset
	}
	return nil
}

func (x *Correction) GetEndOffset() *durationpb.Duration {
	if x != nil {
		return x.EndOffset
	}
	return nil
}

func (x *Correction) GetProposedText() string {
	if x != nil {
		return x.ProposedText
	}
	return ""
}

func (x *Correction) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

var File_api_v1alpha1_vanalytics_correction_proto protoreflect.FileDescriptor

var file_api_v1alpha1_vanalytics_correction_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x73, 0x6b, 0x22, 0x5f, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x72,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x43, 0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5e, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x43, 0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5f, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43,
	0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x72, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x73, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x53, 0x69, 0x64, 0x22, 0x60, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x72,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x45, 0x0a, 0x0b, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x58, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x22, 0x5f, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a,
	0x0a, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x72, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x91, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x64, 0x12,
	0x3c, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x38, 0x0a,
	0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x6e,
	0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0xe0, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x42, 0x0f, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0xa2, 0x02, 0x03,
	0x41, 0x56, 0x56, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0xca, 0x02, 0x17,
	0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x56, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x56,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1alpha1_vanalytics_correction_proto_rawDescOnce sync.Once
	file_api_v1alpha1_vanalytics_correction_proto_rawDescData = file_api_v1alpha1_vanalytics_correction_proto_rawDesc
)

func file_api_v1alpha1_vanalytics_correction_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_vanalytics_correction_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_vanalytics_correction_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_vanalytics_correction_proto_rawDescData)
	})
	return file_api_v1alpha1_vanalytics_correction_proto_rawDescData
}

var file_api_v1alpha1_vanalytics_correction_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v1alpha1_vanalytics_correction_proto_goTypes = []any{
	(*UpdateCorrectionRequest)(nil),  // 0: api.v1alpha1.vanalytics.UpdateCorrectionRequest
	(*UpdateCorrectionResponse)(nil), // 1: api.v1alpha1.vanalytics.UpdateCorrectionResponse
	(*CreateCorrectionRequest)(nil),  // 2: api.v1alpha1.vanalytics.CreateCorrectionRequest
	(*CreateCorrectionResponse)(nil), // 3: api.v1alpha1.vanalytics.CreateCorrectionResponse
	(*GetCorrectionRequest)(nil),     // 4: api.v1alpha1.vanalytics.GetCorrectionRequest
	(*ListCorrectionsRequest)(nil),   // 5: api.v1alpha1.vanalytics.ListCorrectionsRequest
	(*ListCorrectionsResponse)(nil),  // 6: api.v1alpha1.vanalytics.ListCorrectionsResponse
	(*DeleteCorrectionRequest)(nil),  // 7: api.v1alpha1.vanalytics.DeleteCorrectionRequest
	(*DeleteCorrectionResponse)(nil), // 8: api.v1alpha1.vanalytics.DeleteCorrectionResponse
	(*Correction)(nil),               // 9: api.v1alpha1.vanalytics.Correction
	(*fieldmaskpb.FieldMask)(nil),    // 10: google.protobuf.FieldMask
	(*durationpb.Duration)(nil),      // 11: google.protobuf.Duration
}
var file_api_v1alpha1_vanalytics_correction_proto_depIdxs = []int32{
	9,  // 0: api.v1alpha1.vanalytics.UpdateCorrectionRequest.correction:type_name -> api.v1alpha1.vanalytics.Correction
	10, // 1: api.v1alpha1.vanalytics.UpdateCorrectionRequest.update_mask:type_name -> google.protobuf.FieldMask
	9,  // 2: api.v1alpha1.vanalytics.UpdateCorrectionResponse.correction:type_name -> api.v1alpha1.vanalytics.Correction
	9,  // 3: api.v1alpha1.vanalytics.CreateCorrectionRequest.correction:type_name -> api.v1alpha1.vanalytics.Correction
	9,  // 4: api.v1alpha1.vanalytics.CreateCorrectionResponse.correction:type_name -> api.v1alpha1.vanalytics.Correction
	9,  // 5: api.v1alpha1.vanalytics.ListCorrectionsResponse.corrections:type_name -> api.v1alpha1.vanalytics.Correction
	9,  // 6: api.v1alpha1.vanalytics.DeleteCorrectionResponse.correction:type_name -> api.v1alpha1.vanalytics.Correction
	11, // 7: api.v1alpha1.vanalytics.Correction.start_offset:type_name -> google.protobuf.Duration
	11, // 8: api.v1alpha1.vanalytics.Correction.end_offset:type_name -> google.protobuf.Duration
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_vanalytics_correction_proto_init() }
func file_api_v1alpha1_vanalytics_correction_proto_init() {
	if File_api_v1alpha1_vanalytics_correction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_vanalytics_correction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_vanalytics_correction_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_vanalytics_correction_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_vanalytics_correction_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_vanalytics_correction_proto = out.File
	file_api_v1alpha1_vanalytics_correction_proto_rawDesc = nil
	file_api_v1alpha1_vanalytics_correction_proto_goTypes = nil
	file_api_v1alpha1_vanalytics_correction_proto_depIdxs = nil
}
