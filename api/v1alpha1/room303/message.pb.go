// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/room303/message.proto

package room303

import (
	commons "github.com/tcncloud/api-go/api/commons"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateMessageRequest struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	RoomId  string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Payload string                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// nonce  is set by UI and used to tell if the message coming through the message update stream was created by that client.
	// this is so when the user is logged in on multiple clients (different device, tabs, browser) they can see their updated message on those clients.
	Nonce         string `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateMessageRequest) Reset() {
	*x = CreateMessageRequest{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRequest) ProtoMessage() {}

func (x *CreateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMessageRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *CreateMessageRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *CreateMessageRequest) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type CreateMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *commons.Message       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateMessageResponse) Reset() {
	*x = CreateMessageResponse{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageResponse) ProtoMessage() {}

func (x *CreateMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageResponse.ProtoReflect.Descriptor instead.
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMessageResponse) GetMessage() *commons.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type EditMessageRequest struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	MessageId string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Payload   string                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	RoomId    string                 `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	// nonce  is set by UI and used to tell if the message coming through the message update stream was created by that client.
	// this is so when the user is logged in on multiple clients (different device, tabs, browser) they can see their updated message on those clients.
	Nonce         string `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditMessageRequest) Reset() {
	*x = EditMessageRequest{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditMessageRequest) ProtoMessage() {}

func (x *EditMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditMessageRequest.ProtoReflect.Descriptor instead.
func (*EditMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{2}
}

func (x *EditMessageRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *EditMessageRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *EditMessageRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *EditMessageRequest) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type EditMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *commons.Message       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditMessageResponse) Reset() {
	*x = EditMessageResponse{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditMessageResponse) ProtoMessage() {}

func (x *EditMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditMessageResponse.ProtoReflect.Descriptor instead.
func (*EditMessageResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{3}
}

func (x *EditMessageResponse) GetMessage() *commons.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type GetMessagesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Offset        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMessagesRequest) Reset() {
	*x = GetMessagesRequest{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesRequest) ProtoMessage() {}

func (x *GetMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{4}
}

func (x *GetMessagesRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *GetMessagesRequest) GetOffset() *timestamppb.Timestamp {
	if x != nil {
		return x.Offset
	}
	return nil
}

type GetMessagesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []*commons.Message     `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMessagesResponse) Reset() {
	*x = GetMessagesResponse{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesResponse) ProtoMessage() {}

func (x *GetMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetMessagesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{5}
}

func (x *GetMessagesResponse) GetMessages() []*commons.Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type StreamMessageUpdatesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	MemberId      string                 `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamMessageUpdatesRequest) Reset() {
	*x = StreamMessageUpdatesRequest{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamMessageUpdatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessageUpdatesRequest) ProtoMessage() {}

func (x *StreamMessageUpdatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessageUpdatesRequest.ProtoReflect.Descriptor instead.
func (*StreamMessageUpdatesRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{6}
}

func (x *StreamMessageUpdatesRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *StreamMessageUpdatesRequest) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type StreamMessageUpdatesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *commons.Message       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamMessageUpdatesResponse) Reset() {
	*x = StreamMessageUpdatesResponse{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamMessageUpdatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessageUpdatesResponse) ProtoMessage() {}

func (x *StreamMessageUpdatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessageUpdatesResponse.ProtoReflect.Descriptor instead.
func (*StreamMessageUpdatesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{7}
}

func (x *StreamMessageUpdatesResponse) GetMessage() *commons.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type MarkMessageReadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MessageId     string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	MemberId      string                 `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MarkMessageReadRequest) Reset() {
	*x = MarkMessageReadRequest{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarkMessageReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkMessageReadRequest) ProtoMessage() {}

func (x *MarkMessageReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkMessageReadRequest.ProtoReflect.Descriptor instead.
func (*MarkMessageReadRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{8}
}

func (x *MarkMessageReadRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *MarkMessageReadRequest) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type MarkMessageReadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MarkMessageReadResponse) Reset() {
	*x = MarkMessageReadResponse{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarkMessageReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkMessageReadResponse) ProtoMessage() {}

func (x *MarkMessageReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkMessageReadResponse.ProtoReflect.Descriptor instead.
func (*MarkMessageReadResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{9}
}

type MarkAllMessagesReadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MarkAllMessagesReadRequest) Reset() {
	*x = MarkAllMessagesReadRequest{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarkAllMessagesReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAllMessagesReadRequest) ProtoMessage() {}

func (x *MarkAllMessagesReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAllMessagesReadRequest.ProtoReflect.Descriptor instead.
func (*MarkAllMessagesReadRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{10}
}

func (x *MarkAllMessagesReadRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type MarkAllMessagesReadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RowsUpdated   []*commons.Message     `protobuf:"bytes,1,rep,name=rows_updated,json=rowsUpdated,proto3" json:"rows_updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MarkAllMessagesReadResponse) Reset() {
	*x = MarkAllMessagesReadResponse{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarkAllMessagesReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAllMessagesReadResponse) ProtoMessage() {}

func (x *MarkAllMessagesReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAllMessagesReadResponse.ProtoReflect.Descriptor instead.
func (*MarkAllMessagesReadResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{11}
}

func (x *MarkAllMessagesReadResponse) GetRowsUpdated() []*commons.Message {
	if x != nil {
		return x.RowsUpdated
	}
	return nil
}

type GetUnreadStatsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUnreadStatsRequest) Reset() {
	*x = GetUnreadStatsRequest{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUnreadStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnreadStatsRequest) ProtoMessage() {}

func (x *GetUnreadStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnreadStatsRequest.ProtoReflect.Descriptor instead.
func (*GetUnreadStatsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{12}
}

type GetUnreadStatsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Stats         []*commons.MessageStat `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUnreadStatsResponse) Reset() {
	*x = GetUnreadStatsResponse{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUnreadStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnreadStatsResponse) ProtoMessage() {}

func (x *GetUnreadStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnreadStatsResponse.ProtoReflect.Descriptor instead.
func (*GetUnreadStatsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{13}
}

func (x *GetUnreadStatsResponse) GetStats() []*commons.MessageStat {
	if x != nil {
		return x.Stats
	}
	return nil
}

type DeleteMessageRequest struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	MessageId string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	RoomId    string                 `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	// nonce  is set by UI and used to tell if the message coming through the message update stream was created by that client.
	// this is so when the user is logged in on multiple clients (different device, tabs, browser) they can see their updated message on those clients.
	Nonce         string `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteMessageRequest) Reset() {
	*x = DeleteMessageRequest{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageRequest) ProtoMessage() {}

func (x *DeleteMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageRequest.ProtoReflect.Descriptor instead.
func (*DeleteMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{14}
}

func (x *DeleteMessageRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *DeleteMessageRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *DeleteMessageRequest) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type DeleteMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *commons.Message       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteMessageResponse) Reset() {
	*x = DeleteMessageResponse{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageResponse) ProtoMessage() {}

func (x *DeleteMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageResponse.ProtoReflect.Descriptor instead.
func (*DeleteMessageResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{15}
}

func (x *DeleteMessageResponse) GetMessage() *commons.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type BulkMarkMessageReadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	MessageIds    []string               `protobuf:"bytes,2,rep,name=message_ids,json=messageIds,proto3" json:"message_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BulkMarkMessageReadRequest) Reset() {
	*x = BulkMarkMessageReadRequest{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BulkMarkMessageReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkMarkMessageReadRequest) ProtoMessage() {}

func (x *BulkMarkMessageReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkMarkMessageReadRequest.ProtoReflect.Descriptor instead.
func (*BulkMarkMessageReadRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{16}
}

func (x *BulkMarkMessageReadRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *BulkMarkMessageReadRequest) GetMessageIds() []string {
	if x != nil {
		return x.MessageIds
	}
	return nil
}

type BulkMarkMessageReadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []*commons.Message     `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BulkMarkMessageReadResponse) Reset() {
	*x = BulkMarkMessageReadResponse{}
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BulkMarkMessageReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkMarkMessageReadResponse) ProtoMessage() {}

func (x *BulkMarkMessageReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_message_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkMarkMessageReadResponse.ProtoReflect.Descriptor instead.
func (*BulkMarkMessageReadResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_message_proto_rawDescGZIP(), []int{17}
}

func (x *BulkMarkMessageReadResponse) GetMessages() []*commons.Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_api_v1alpha1_room303_message_proto protoreflect.FileDescriptor

const file_api_v1alpha1_room303_message_proto_rawDesc = "" +
	"\n" +
	"\"api/v1alpha1/room303/message.proto\x12\x14api.v1alpha1.room303\x1a\x19api/commons/room303.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"_\n" +
	"\x14CreateMessageRequest\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x12\x18\n" +
	"\apayload\x18\x02 \x01(\tR\apayload\x12\x14\n" +
	"\x05nonce\x18\x03 \x01(\tR\x05nonce\"G\n" +
	"\x15CreateMessageResponse\x12.\n" +
	"\amessage\x18\x01 \x01(\v2\x14.api.commons.MessageR\amessage\"|\n" +
	"\x12EditMessageRequest\x12\x1d\n" +
	"\n" +
	"message_id\x18\x01 \x01(\tR\tmessageId\x12\x18\n" +
	"\apayload\x18\x02 \x01(\tR\apayload\x12\x17\n" +
	"\aroom_id\x18\x03 \x01(\tR\x06roomId\x12\x14\n" +
	"\x05nonce\x18\x04 \x01(\tR\x05nonce\"E\n" +
	"\x13EditMessageResponse\x12.\n" +
	"\amessage\x18\x01 \x01(\v2\x14.api.commons.MessageR\amessage\"g\n" +
	"\x12GetMessagesRequest\x12\x17\n" +
	"\aroom_id\x18\x02 \x01(\tR\x06roomId\x122\n" +
	"\x06offset\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\x06offsetJ\x04\b\x01\x10\x02\"G\n" +
	"\x13GetMessagesResponse\x120\n" +
	"\bmessages\x18\x01 \x03(\v2\x14.api.commons.MessageR\bmessages\"S\n" +
	"\x1bStreamMessageUpdatesRequest\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x12\x1b\n" +
	"\tmember_id\x18\x02 \x01(\tR\bmemberId\"N\n" +
	"\x1cStreamMessageUpdatesResponse\x12.\n" +
	"\amessage\x18\x01 \x01(\v2\x14.api.commons.MessageR\amessage\"T\n" +
	"\x16MarkMessageReadRequest\x12\x1d\n" +
	"\n" +
	"message_id\x18\x01 \x01(\tR\tmessageId\x12\x1b\n" +
	"\tmember_id\x18\x02 \x01(\tR\bmemberId\"\x19\n" +
	"\x17MarkMessageReadResponse\"5\n" +
	"\x1aMarkAllMessagesReadRequest\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\"V\n" +
	"\x1bMarkAllMessagesReadResponse\x127\n" +
	"\frows_updated\x18\x01 \x03(\v2\x14.api.commons.MessageR\vrowsUpdated\"\x17\n" +
	"\x15GetUnreadStatsRequest\"H\n" +
	"\x16GetUnreadStatsResponse\x12.\n" +
	"\x05stats\x18\x01 \x03(\v2\x18.api.commons.MessageStatR\x05stats\"d\n" +
	"\x14DeleteMessageRequest\x12\x1d\n" +
	"\n" +
	"message_id\x18\x01 \x01(\tR\tmessageId\x12\x17\n" +
	"\aroom_id\x18\x02 \x01(\tR\x06roomId\x12\x14\n" +
	"\x05nonce\x18\x03 \x01(\tR\x05nonce\"G\n" +
	"\x15DeleteMessageResponse\x12.\n" +
	"\amessage\x18\x01 \x01(\v2\x14.api.commons.MessageR\amessage\"V\n" +
	"\x1aBulkMarkMessageReadRequest\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x12\x1f\n" +
	"\vmessage_ids\x18\x02 \x03(\tR\n" +
	"messageIds\"O\n" +
	"\x1bBulkMarkMessageReadResponse\x120\n" +
	"\bmessages\x18\x01 \x03(\v2\x14.api.commons.MessageR\bmessagesB\xcb\x01\n" +
	"\x18com.api.v1alpha1.room303B\fMessageProtoP\x01Z/github.com/tcncloud/api-go/api/v1alpha1/room303\xa2\x02\x03AVR\xaa\x02\x14Api.V1alpha1.Room303\xca\x02\x14Api\\V1alpha1\\Room303\xe2\x02 Api\\V1alpha1\\Room303\\GPBMetadata\xea\x02\x16Api::V1alpha1::Room303b\x06proto3"

var (
	file_api_v1alpha1_room303_message_proto_rawDescOnce sync.Once
	file_api_v1alpha1_room303_message_proto_rawDescData []byte
)

func file_api_v1alpha1_room303_message_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_room303_message_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_room303_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_room303_message_proto_rawDesc), len(file_api_v1alpha1_room303_message_proto_rawDesc)))
	})
	return file_api_v1alpha1_room303_message_proto_rawDescData
}

var file_api_v1alpha1_room303_message_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_api_v1alpha1_room303_message_proto_goTypes = []any{
	(*CreateMessageRequest)(nil),         // 0: api.v1alpha1.room303.CreateMessageRequest
	(*CreateMessageResponse)(nil),        // 1: api.v1alpha1.room303.CreateMessageResponse
	(*EditMessageRequest)(nil),           // 2: api.v1alpha1.room303.EditMessageRequest
	(*EditMessageResponse)(nil),          // 3: api.v1alpha1.room303.EditMessageResponse
	(*GetMessagesRequest)(nil),           // 4: api.v1alpha1.room303.GetMessagesRequest
	(*GetMessagesResponse)(nil),          // 5: api.v1alpha1.room303.GetMessagesResponse
	(*StreamMessageUpdatesRequest)(nil),  // 6: api.v1alpha1.room303.StreamMessageUpdatesRequest
	(*StreamMessageUpdatesResponse)(nil), // 7: api.v1alpha1.room303.StreamMessageUpdatesResponse
	(*MarkMessageReadRequest)(nil),       // 8: api.v1alpha1.room303.MarkMessageReadRequest
	(*MarkMessageReadResponse)(nil),      // 9: api.v1alpha1.room303.MarkMessageReadResponse
	(*MarkAllMessagesReadRequest)(nil),   // 10: api.v1alpha1.room303.MarkAllMessagesReadRequest
	(*MarkAllMessagesReadResponse)(nil),  // 11: api.v1alpha1.room303.MarkAllMessagesReadResponse
	(*GetUnreadStatsRequest)(nil),        // 12: api.v1alpha1.room303.GetUnreadStatsRequest
	(*GetUnreadStatsResponse)(nil),       // 13: api.v1alpha1.room303.GetUnreadStatsResponse
	(*DeleteMessageRequest)(nil),         // 14: api.v1alpha1.room303.DeleteMessageRequest
	(*DeleteMessageResponse)(nil),        // 15: api.v1alpha1.room303.DeleteMessageResponse
	(*BulkMarkMessageReadRequest)(nil),   // 16: api.v1alpha1.room303.BulkMarkMessageReadRequest
	(*BulkMarkMessageReadResponse)(nil),  // 17: api.v1alpha1.room303.BulkMarkMessageReadResponse
	(*commons.Message)(nil),              // 18: api.commons.Message
	(*timestamppb.Timestamp)(nil),        // 19: google.protobuf.Timestamp
	(*commons.MessageStat)(nil),          // 20: api.commons.MessageStat
}
var file_api_v1alpha1_room303_message_proto_depIdxs = []int32{
	18, // 0: api.v1alpha1.room303.CreateMessageResponse.message:type_name -> api.commons.Message
	18, // 1: api.v1alpha1.room303.EditMessageResponse.message:type_name -> api.commons.Message
	19, // 2: api.v1alpha1.room303.GetMessagesRequest.offset:type_name -> google.protobuf.Timestamp
	18, // 3: api.v1alpha1.room303.GetMessagesResponse.messages:type_name -> api.commons.Message
	18, // 4: api.v1alpha1.room303.StreamMessageUpdatesResponse.message:type_name -> api.commons.Message
	18, // 5: api.v1alpha1.room303.MarkAllMessagesReadResponse.rows_updated:type_name -> api.commons.Message
	20, // 6: api.v1alpha1.room303.GetUnreadStatsResponse.stats:type_name -> api.commons.MessageStat
	18, // 7: api.v1alpha1.room303.DeleteMessageResponse.message:type_name -> api.commons.Message
	18, // 8: api.v1alpha1.room303.BulkMarkMessageReadResponse.messages:type_name -> api.commons.Message
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_room303_message_proto_init() }
func file_api_v1alpha1_room303_message_proto_init() {
	if File_api_v1alpha1_room303_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_room303_message_proto_rawDesc), len(file_api_v1alpha1_room303_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_room303_message_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_room303_message_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_room303_message_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_room303_message_proto = out.File
	file_api_v1alpha1_room303_message_proto_goTypes = nil
	file_api_v1alpha1_room303_message_proto_depIdxs = nil
}
