// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/room303/member.proto

package room303

import (
	commons "github.com/tcncloud/api-go/api/commons"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type AddRoomMemberRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Admin         bool                   `protobuf:"varint,3,opt,name=admin,proto3" json:"admin,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRoomMemberRequest) Reset() {
	*x = AddRoomMemberRequest{}
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRoomMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoomMemberRequest) ProtoMessage() {}

func (x *AddRoomMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoomMemberRequest.ProtoReflect.Descriptor instead.
func (*AddRoomMemberRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_member_proto_rawDescGZIP(), []int{0}
}

func (x *AddRoomMemberRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *AddRoomMemberRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddRoomMemberRequest) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

type RemoveRoomMemberRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoomId        string                 `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRoomMemberRequest) Reset() {
	*x = RemoveRoomMemberRequest{}
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRoomMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoomMemberRequest) ProtoMessage() {}

func (x *RemoveRoomMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoomMemberRequest.ProtoReflect.Descriptor instead.
func (*RemoveRoomMemberRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_member_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveRoomMemberRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemoveRoomMemberRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type RemoveRoomMemberResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRoomMemberResponse) Reset() {
	*x = RemoveRoomMemberResponse{}
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRoomMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoomMemberResponse) ProtoMessage() {}

func (x *RemoveRoomMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoomMemberResponse.ProtoReflect.Descriptor instead.
func (*RemoveRoomMemberResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_member_proto_rawDescGZIP(), []int{2}
}

type ListRoomMembersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRoomMembersRequest) Reset() {
	*x = ListRoomMembersRequest{}
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRoomMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomMembersRequest) ProtoMessage() {}

func (x *ListRoomMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomMembersRequest.ProtoReflect.Descriptor instead.
func (*ListRoomMembersRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_member_proto_rawDescGZIP(), []int{3}
}

func (x *ListRoomMembersRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type ListRoomMembersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Members       []*commons.Member      `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRoomMembersResponse) Reset() {
	*x = ListRoomMembersResponse{}
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRoomMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomMembersResponse) ProtoMessage() {}

func (x *ListRoomMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomMembersResponse.ProtoReflect.Descriptor instead.
func (*ListRoomMembersResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_member_proto_rawDescGZIP(), []int{4}
}

func (x *ListRoomMembersResponse) GetMembers() []*commons.Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type SetAdminForRoomMemberRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetAdminForRoomMemberRequest) Reset() {
	*x = SetAdminForRoomMemberRequest{}
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAdminForRoomMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAdminForRoomMemberRequest) ProtoMessage() {}

func (x *SetAdminForRoomMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAdminForRoomMemberRequest.ProtoReflect.Descriptor instead.
func (*SetAdminForRoomMemberRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_member_proto_rawDescGZIP(), []int{5}
}

func (x *SetAdminForRoomMemberRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *SetAdminForRoomMemberRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SetAdminForRoomMemberResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetAdminForRoomMemberResponse) Reset() {
	*x = SetAdminForRoomMemberResponse{}
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAdminForRoomMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAdminForRoomMemberResponse) ProtoMessage() {}

func (x *SetAdminForRoomMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAdminForRoomMemberResponse.ProtoReflect.Descriptor instead.
func (*SetAdminForRoomMemberResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_member_proto_rawDescGZIP(), []int{6}
}

type JoinRoomRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JoinRoomRequest) Reset() {
	*x = JoinRoomRequest{}
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRequest) ProtoMessage() {}

func (x *JoinRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomRequest.ProtoReflect.Descriptor instead.
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_member_proto_rawDescGZIP(), []int{7}
}

func (x *JoinRoomRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type GetRoomMemberRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoomMemberRequest) Reset() {
	*x = GetRoomMemberRequest{}
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoomMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomMemberRequest) ProtoMessage() {}

func (x *GetRoomMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_room303_member_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomMemberRequest.ProtoReflect.Descriptor instead.
func (*GetRoomMemberRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_room303_member_proto_rawDescGZIP(), []int{8}
}

func (x *GetRoomMemberRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *GetRoomMemberRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_api_v1alpha1_room303_member_proto protoreflect.FileDescriptor

const file_api_v1alpha1_room303_member_proto_rawDesc = "" +
	"\n" +
	"!api/v1alpha1/room303/member.proto\x12\x14api.v1alpha1.room303\x1a\x19api/commons/room303.proto\"^\n" +
	"\x14AddRoomMemberRequest\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x14\n" +
	"\x05admin\x18\x03 \x01(\bR\x05admin\"K\n" +
	"\x17RemoveRoomMemberRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x17\n" +
	"\aroom_id\x18\x02 \x01(\tR\x06roomId\"\x1a\n" +
	"\x18RemoveRoomMemberResponse\"1\n" +
	"\x16ListRoomMembersRequest\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\"H\n" +
	"\x17ListRoomMembersResponse\x12-\n" +
	"\amembers\x18\x01 \x03(\v2\x13.api.commons.MemberR\amembers\"P\n" +
	"\x1cSetAdminForRoomMemberRequest\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\"\x1f\n" +
	"\x1dSetAdminForRoomMemberResponse\"*\n" +
	"\x0fJoinRoomRequest\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\"H\n" +
	"\x14GetRoomMemberRequest\x12\x17\n" +
	"\aroom_id\x18\x02 \x01(\tR\x06roomId\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\tR\x06userIdB\xca\x01\n" +
	"\x18com.api.v1alpha1.room303B\vMemberProtoP\x01Z/github.com/tcncloud/api-go/api/v1alpha1/room303\xa2\x02\x03AVR\xaa\x02\x14Api.V1alpha1.Room303\xca\x02\x14Api\\V1alpha1\\Room303\xe2\x02 Api\\V1alpha1\\Room303\\GPBMetadata\xea\x02\x16Api::V1alpha1::Room303b\x06proto3"

var (
	file_api_v1alpha1_room303_member_proto_rawDescOnce sync.Once
	file_api_v1alpha1_room303_member_proto_rawDescData []byte
)

func file_api_v1alpha1_room303_member_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_room303_member_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_room303_member_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_room303_member_proto_rawDesc), len(file_api_v1alpha1_room303_member_proto_rawDesc)))
	})
	return file_api_v1alpha1_room303_member_proto_rawDescData
}

var file_api_v1alpha1_room303_member_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_v1alpha1_room303_member_proto_goTypes = []any{
	(*AddRoomMemberRequest)(nil),          // 0: api.v1alpha1.room303.AddRoomMemberRequest
	(*RemoveRoomMemberRequest)(nil),       // 1: api.v1alpha1.room303.RemoveRoomMemberRequest
	(*RemoveRoomMemberResponse)(nil),      // 2: api.v1alpha1.room303.RemoveRoomMemberResponse
	(*ListRoomMembersRequest)(nil),        // 3: api.v1alpha1.room303.ListRoomMembersRequest
	(*ListRoomMembersResponse)(nil),       // 4: api.v1alpha1.room303.ListRoomMembersResponse
	(*SetAdminForRoomMemberRequest)(nil),  // 5: api.v1alpha1.room303.SetAdminForRoomMemberRequest
	(*SetAdminForRoomMemberResponse)(nil), // 6: api.v1alpha1.room303.SetAdminForRoomMemberResponse
	(*JoinRoomRequest)(nil),               // 7: api.v1alpha1.room303.JoinRoomRequest
	(*GetRoomMemberRequest)(nil),          // 8: api.v1alpha1.room303.GetRoomMemberRequest
	(*commons.Member)(nil),                // 9: api.commons.Member
}
var file_api_v1alpha1_room303_member_proto_depIdxs = []int32{
	9, // 0: api.v1alpha1.room303.ListRoomMembersResponse.members:type_name -> api.commons.Member
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_room303_member_proto_init() }
func file_api_v1alpha1_room303_member_proto_init() {
	if File_api_v1alpha1_room303_member_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_room303_member_proto_rawDesc), len(file_api_v1alpha1_room303_member_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_room303_member_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_room303_member_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_room303_member_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_room303_member_proto = out.File
	file_api_v1alpha1_room303_member_proto_goTypes = nil
	file_api_v1alpha1_room303_member_proto_depIdxs = nil
}
