// Copyright (c) 2020, TCN Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: api/v1alpha1/ghostnotifier/service.proto

package ghostnotifier

import (
	_ "github.com/tcncloud/api-go/annotations"
	commons "github.com/tcncloud/api-go/api/commons"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListNotificationsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNotificationsReq) Reset() {
	*x = ListNotificationsReq{}
	mi := &file_api_v1alpha1_ghostnotifier_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNotificationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotificationsReq) ProtoMessage() {}

func (x *ListNotificationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_ghostnotifier_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotificationsReq.ProtoReflect.Descriptor instead.
func (*ListNotificationsReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_ghostnotifier_service_proto_rawDescGZIP(), []int{0}
}

var File_api_v1alpha1_ghostnotifier_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_ghostnotifier_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67,
	0x68, 0x6f, 0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x32, 0xce, 0x01, 0x0a, 0x10, 0x47, 0x68, 0x6f, 0x73, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x41, 0x70, 0x69, 0x12, 0xb9, 0x01, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x50, 0xba, 0xb8, 0x91, 0x02, 0x02, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x43, 0x3a, 0x01, 0x2a, 0x22, 0x3e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x2f, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x30, 0x01, 0x42, 0xef, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x67, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x47, 0xaa, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0xca, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5c, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0xe2, 0x02, 0x26, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x47, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1alpha1_ghostnotifier_service_proto_rawDescOnce sync.Once
	file_api_v1alpha1_ghostnotifier_service_proto_rawDescData = file_api_v1alpha1_ghostnotifier_service_proto_rawDesc
)

func file_api_v1alpha1_ghostnotifier_service_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_ghostnotifier_service_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_ghostnotifier_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_ghostnotifier_service_proto_rawDescData)
	})
	return file_api_v1alpha1_ghostnotifier_service_proto_rawDescData
}

var file_api_v1alpha1_ghostnotifier_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1alpha1_ghostnotifier_service_proto_goTypes = []any{
	(*ListNotificationsReq)(nil),      // 0: api.v1alpha1.ghostnotifier.ListNotificationsReq
	(*commons.GhostNotification)(nil), // 1: api.commons.GhostNotification
}
var file_api_v1alpha1_ghostnotifier_service_proto_depIdxs = []int32{
	0, // 0: api.v1alpha1.ghostnotifier.GhostNotifierApi.ListNotifications:input_type -> api.v1alpha1.ghostnotifier.ListNotificationsReq
	1, // 1: api.v1alpha1.ghostnotifier.GhostNotifierApi.ListNotifications:output_type -> api.commons.GhostNotification
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_ghostnotifier_service_proto_init() }
func file_api_v1alpha1_ghostnotifier_service_proto_init() {
	if File_api_v1alpha1_ghostnotifier_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_ghostnotifier_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_ghostnotifier_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_ghostnotifier_service_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_ghostnotifier_service_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_ghostnotifier_service_proto = out.File
	file_api_v1alpha1_ghostnotifier_service_proto_rawDesc = nil
	file_api_v1alpha1_ghostnotifier_service_proto_goTypes = nil
	file_api_v1alpha1_ghostnotifier_service_proto_depIdxs = nil
}
