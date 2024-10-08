// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/v1alpha1/soundboard/entities.proto

package soundboard

import (
	commons "github.com/tcncloud/api-go/api/commons"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SoundboardDetails is the core entity which contains metadata for soundboard
// audio files, along with their identifiers.
type SoundboardDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the soundboard, in the snowflake ID format.
	SoundboardId int64 `protobuf:"varint,1,opt,name=soundboard_id,json=soundboardId,proto3" json:"soundboard_id,omitempty"`
	// Name of audio file (does not include path), e.g. 'file.wav'.
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	// Audio file format of soundboard (.wav or .mp3).
	FileType commons.RecordingFileType `protobuf:"varint,3,opt,name=file_type,json=fileType,proto3,enum=api.commons.RecordingFileType" json:"file_type,omitempty"`
	// Soundboard title.
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// Soundboard description.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Timestamp of when the soundboard was created.
	DateCreated *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	// Timestamp of when the soundboard was last updated.
	LastModified *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	// Soundboard audio file size as number of bytes, e.g. '3145728' = 3mb.
	FileSize int64 `protobuf:"varint,8,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	// Length of audio file in seconds, e.g. '192' = 3:12.
	RecordingLength int64 `protobuf:"varint,9,opt,name=recording_length,json=recordingLength,proto3" json:"recording_length,omitempty"`
}

func (x *SoundboardDetails) Reset() {
	*x = SoundboardDetails{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SoundboardDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoundboardDetails) ProtoMessage() {}

func (x *SoundboardDetails) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoundboardDetails.ProtoReflect.Descriptor instead.
func (*SoundboardDetails) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{0}
}

func (x *SoundboardDetails) GetSoundboardId() int64 {
	if x != nil {
		return x.SoundboardId
	}
	return 0
}

func (x *SoundboardDetails) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *SoundboardDetails) GetFileType() commons.RecordingFileType {
	if x != nil {
		return x.FileType
	}
	return commons.RecordingFileType(0)
}

func (x *SoundboardDetails) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SoundboardDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SoundboardDetails) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

func (x *SoundboardDetails) GetLastModified() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModified
	}
	return nil
}

func (x *SoundboardDetails) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *SoundboardDetails) GetRecordingLength() int64 {
	if x != nil {
		return x.RecordingLength
	}
	return 0
}

// Request message for the GetSoundboardFile RPC method.
type GetSoundboardFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the requested soundboard whose file will be returned.
	SoundboardId int64 `protobuf:"varint,1,opt,name=soundboard_id,json=soundboardId,proto3" json:"soundboard_id,omitempty"`
}

func (x *GetSoundboardFileReq) Reset() {
	*x = GetSoundboardFileReq{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSoundboardFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSoundboardFileReq) ProtoMessage() {}

func (x *GetSoundboardFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSoundboardFileReq.ProtoReflect.Descriptor instead.
func (*GetSoundboardFileReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{1}
}

func (x *GetSoundboardFileReq) GetSoundboardId() int64 {
	if x != nil {
		return x.SoundboardId
	}
	return 0
}

// Response message for the GetSoundboardFile RPC method.
type GetSoundboardFileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Segments of the audio file, sent in 3mb chunks.
	SoundboardFile []byte `protobuf:"bytes,1,opt,name=soundboard_file,json=soundboardFile,proto3" json:"soundboard_file,omitempty"`
}

func (x *GetSoundboardFileRes) Reset() {
	*x = GetSoundboardFileRes{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSoundboardFileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSoundboardFileRes) ProtoMessage() {}

func (x *GetSoundboardFileRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSoundboardFileRes.ProtoReflect.Descriptor instead.
func (*GetSoundboardFileRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{2}
}

func (x *GetSoundboardFileRes) GetSoundboardFile() []byte {
	if x != nil {
		return x.SoundboardFile
	}
	return nil
}

// Request message for the GetSoundboard RPC method.
type GetSoundboardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the requested soundboard whose details will be returned.
	SoundboardId int64 `protobuf:"varint,1,opt,name=soundboard_id,json=soundboardId,proto3" json:"soundboard_id,omitempty"`
}

func (x *GetSoundboardReq) Reset() {
	*x = GetSoundboardReq{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSoundboardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSoundboardReq) ProtoMessage() {}

func (x *GetSoundboardReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSoundboardReq.ProtoReflect.Descriptor instead.
func (*GetSoundboardReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{3}
}

func (x *GetSoundboardReq) GetSoundboardId() int64 {
	if x != nil {
		return x.SoundboardId
	}
	return 0
}

// Response message for the GetSoundboard RPC method.
type GetSoundboardRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata of the soundboard.
	Soundboard *SoundboardDetails `protobuf:"bytes,1,opt,name=soundboard,proto3" json:"soundboard,omitempty"`
}

func (x *GetSoundboardRes) Reset() {
	*x = GetSoundboardRes{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSoundboardRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSoundboardRes) ProtoMessage() {}

func (x *GetSoundboardRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSoundboardRes.ProtoReflect.Descriptor instead.
func (*GetSoundboardRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{4}
}

func (x *GetSoundboardRes) GetSoundboard() *SoundboardDetails {
	if x != nil {
		return x.Soundboard
	}
	return nil
}

// Request message for the CreateSoundboard RPC method.
type CreateSoundboardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The metadata of the soundboard.
	Soundboard *SoundboardDetails `protobuf:"bytes,1,opt,name=soundboard,proto3" json:"soundboard,omitempty"`
	// The generated ID received from fts.GetUploadFileUrl.
	FtsId string `protobuf:"bytes,2,opt,name=fts_id,json=ftsId,proto3" json:"fts_id,omitempty"`
}

func (x *CreateSoundboardReq) Reset() {
	*x = CreateSoundboardReq{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSoundboardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSoundboardReq) ProtoMessage() {}

func (x *CreateSoundboardReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSoundboardReq.ProtoReflect.Descriptor instead.
func (*CreateSoundboardReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSoundboardReq) GetSoundboard() *SoundboardDetails {
	if x != nil {
		return x.Soundboard
	}
	return nil
}

func (x *CreateSoundboardReq) GetFtsId() string {
	if x != nil {
		return x.FtsId
	}
	return ""
}

// Response message for the CreateSoundboard RPC method.
type CreateSoundboardRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Generated snowflake ID which will correspond to the soundboard.
	SoundboardId int64 `protobuf:"varint,1,opt,name=soundboard_id,json=soundboardId,proto3" json:"soundboard_id,omitempty"`
	// The generated ID received from fts.GetUploadFileUrl.
	FtsId string `protobuf:"bytes,2,opt,name=fts_id,json=ftsId,proto3" json:"fts_id,omitempty"`
}

func (x *CreateSoundboardRes) Reset() {
	*x = CreateSoundboardRes{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSoundboardRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSoundboardRes) ProtoMessage() {}

func (x *CreateSoundboardRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSoundboardRes.ProtoReflect.Descriptor instead.
func (*CreateSoundboardRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{6}
}

func (x *CreateSoundboardRes) GetSoundboardId() int64 {
	if x != nil {
		return x.SoundboardId
	}
	return 0
}

func (x *CreateSoundboardRes) GetFtsId() string {
	if x != nil {
		return x.FtsId
	}
	return ""
}

// Request message for the ListSoundboards RPC method.
type ListSoundboardsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSoundboardsReq) Reset() {
	*x = ListSoundboardsReq{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSoundboardsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSoundboardsReq) ProtoMessage() {}

func (x *ListSoundboardsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSoundboardsReq.ProtoReflect.Descriptor instead.
func (*ListSoundboardsReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{7}
}

// Response message for the ListSoundboards RPC method.
type ListSoundboardsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each soundboard's metadata belonging to org.
	Soundboards []*SoundboardDetails `protobuf:"bytes,1,rep,name=soundboards,proto3" json:"soundboards,omitempty"`
}

func (x *ListSoundboardsRes) Reset() {
	*x = ListSoundboardsRes{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSoundboardsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSoundboardsRes) ProtoMessage() {}

func (x *ListSoundboardsRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSoundboardsRes.ProtoReflect.Descriptor instead.
func (*ListSoundboardsRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{8}
}

func (x *ListSoundboardsRes) GetSoundboards() []*SoundboardDetails {
	if x != nil {
		return x.Soundboards
	}
	return nil
}

// Request message for the UpdateSoundboard RPC method.
type UpdateSoundboardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Soundboard details to be updated, must contain soundboard and org id.
	Soundboard *SoundboardDetails `protobuf:"bytes,1,opt,name=soundboard,proto3" json:"soundboard,omitempty"`
}

func (x *UpdateSoundboardReq) Reset() {
	*x = UpdateSoundboardReq{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSoundboardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSoundboardReq) ProtoMessage() {}

func (x *UpdateSoundboardReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSoundboardReq.ProtoReflect.Descriptor instead.
func (*UpdateSoundboardReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateSoundboardReq) GetSoundboard() *SoundboardDetails {
	if x != nil {
		return x.Soundboard
	}
	return nil
}

// Response message for the UpdateSoundboard RPC method.
type UpdateSoundboardRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the soundboard which was updated.
	SoundboardId int64 `protobuf:"varint,1,opt,name=soundboard_id,json=soundboardId,proto3" json:"soundboard_id,omitempty"`
}

func (x *UpdateSoundboardRes) Reset() {
	*x = UpdateSoundboardRes{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSoundboardRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSoundboardRes) ProtoMessage() {}

func (x *UpdateSoundboardRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSoundboardRes.ProtoReflect.Descriptor instead.
func (*UpdateSoundboardRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateSoundboardRes) GetSoundboardId() int64 {
	if x != nil {
		return x.SoundboardId
	}
	return 0
}

// Request message for the DeleteSoundboard RPC method.
type DeleteSoundboardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the soundboard whose file and details will be deleted.
	SoundboardId int64 `protobuf:"varint,1,opt,name=soundboard_id,json=soundboardId,proto3" json:"soundboard_id,omitempty"`
}

func (x *DeleteSoundboardReq) Reset() {
	*x = DeleteSoundboardReq{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSoundboardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSoundboardReq) ProtoMessage() {}

func (x *DeleteSoundboardReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSoundboardReq.ProtoReflect.Descriptor instead.
func (*DeleteSoundboardReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteSoundboardReq) GetSoundboardId() int64 {
	if x != nil {
		return x.SoundboardId
	}
	return 0
}

// Response message for the DeleteSoundboard RPC method.
type DeleteSoundboardRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSoundboardRes) Reset() {
	*x = DeleteSoundboardRes{}
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSoundboardRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSoundboardRes) ProtoMessage() {}

func (x *DeleteSoundboardRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_soundboard_entities_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSoundboardRes.ProtoReflect.Descriptor instead.
func (*DeleteSoundboardRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP(), []int{12}
}

var File_api_v1alpha1_soundboard_entities_proto protoreflect.FileDescriptor

var file_api_v1alpha1_soundboard_entities_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x03, 0x0a, 0x11, 0x53, 0x6f,
	0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x27, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x6e,
	0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x22, 0x3f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x0d, 0x73, 0x6f,
	0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x22, 0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x6e, 0x64,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x6e,
	0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x22, 0x5e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x22, 0x78, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x6e, 0x64,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x4a, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x6e,
	0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x6f, 0x75, 0x6e,
	0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x66, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x74, 0x73, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x12, 0x27, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0c, 0x73,
	0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x66,
	0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x74, 0x73,
	0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x22, 0x62, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x12, 0x4c,
	0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x6f,
	0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x0b, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x22, 0x61, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x4a, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22,
	0x3e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30,
	0x01, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22,
	0x3e, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30,
	0x01, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22,
	0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x42, 0xde, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x73, 0x6f, 0x75, 0x6e,
	0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0xa2, 0x02, 0x03, 0x41, 0x56,
	0x53, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0xca, 0x02, 0x17, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x53, 0x6f, 0x75, 0x6e, 0x64,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x5c, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x53, 0x6f, 0x75,
	0x6e, 0x64, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_soundboard_entities_proto_rawDescOnce sync.Once
	file_api_v1alpha1_soundboard_entities_proto_rawDescData = file_api_v1alpha1_soundboard_entities_proto_rawDesc
)

func file_api_v1alpha1_soundboard_entities_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_soundboard_entities_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_soundboard_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_soundboard_entities_proto_rawDescData)
	})
	return file_api_v1alpha1_soundboard_entities_proto_rawDescData
}

var file_api_v1alpha1_soundboard_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_v1alpha1_soundboard_entities_proto_goTypes = []any{
	(*SoundboardDetails)(nil),      // 0: api.v1alpha1.soundboard.SoundboardDetails
	(*GetSoundboardFileReq)(nil),   // 1: api.v1alpha1.soundboard.GetSoundboardFileReq
	(*GetSoundboardFileRes)(nil),   // 2: api.v1alpha1.soundboard.GetSoundboardFileRes
	(*GetSoundboardReq)(nil),       // 3: api.v1alpha1.soundboard.GetSoundboardReq
	(*GetSoundboardRes)(nil),       // 4: api.v1alpha1.soundboard.GetSoundboardRes
	(*CreateSoundboardReq)(nil),    // 5: api.v1alpha1.soundboard.CreateSoundboardReq
	(*CreateSoundboardRes)(nil),    // 6: api.v1alpha1.soundboard.CreateSoundboardRes
	(*ListSoundboardsReq)(nil),     // 7: api.v1alpha1.soundboard.ListSoundboardsReq
	(*ListSoundboardsRes)(nil),     // 8: api.v1alpha1.soundboard.ListSoundboardsRes
	(*UpdateSoundboardReq)(nil),    // 9: api.v1alpha1.soundboard.UpdateSoundboardReq
	(*UpdateSoundboardRes)(nil),    // 10: api.v1alpha1.soundboard.UpdateSoundboardRes
	(*DeleteSoundboardReq)(nil),    // 11: api.v1alpha1.soundboard.DeleteSoundboardReq
	(*DeleteSoundboardRes)(nil),    // 12: api.v1alpha1.soundboard.DeleteSoundboardRes
	(commons.RecordingFileType)(0), // 13: api.commons.RecordingFileType
	(*timestamppb.Timestamp)(nil),  // 14: google.protobuf.Timestamp
}
var file_api_v1alpha1_soundboard_entities_proto_depIdxs = []int32{
	13, // 0: api.v1alpha1.soundboard.SoundboardDetails.file_type:type_name -> api.commons.RecordingFileType
	14, // 1: api.v1alpha1.soundboard.SoundboardDetails.date_created:type_name -> google.protobuf.Timestamp
	14, // 2: api.v1alpha1.soundboard.SoundboardDetails.last_modified:type_name -> google.protobuf.Timestamp
	0,  // 3: api.v1alpha1.soundboard.GetSoundboardRes.soundboard:type_name -> api.v1alpha1.soundboard.SoundboardDetails
	0,  // 4: api.v1alpha1.soundboard.CreateSoundboardReq.soundboard:type_name -> api.v1alpha1.soundboard.SoundboardDetails
	0,  // 5: api.v1alpha1.soundboard.ListSoundboardsRes.soundboards:type_name -> api.v1alpha1.soundboard.SoundboardDetails
	0,  // 6: api.v1alpha1.soundboard.UpdateSoundboardReq.soundboard:type_name -> api.v1alpha1.soundboard.SoundboardDetails
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_soundboard_entities_proto_init() }
func file_api_v1alpha1_soundboard_entities_proto_init() {
	if File_api_v1alpha1_soundboard_entities_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_soundboard_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_soundboard_entities_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_soundboard_entities_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_soundboard_entities_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_soundboard_entities_proto = out.File
	file_api_v1alpha1_soundboard_entities_proto_rawDesc = nil
	file_api_v1alpha1_soundboard_entities_proto_goTypes = nil
	file_api_v1alpha1_soundboard_entities_proto_depIdxs = nil
}
