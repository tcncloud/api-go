// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/v1alpha1/newsroom/service.proto

package newsroom

import (
	_ "github.com/tcncloud/api-go/annotations"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_v1alpha1_newsroom_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_newsroom_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e,
	0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x1a, 0x17, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x89, 0x12, 0x0a, 0x0b, 0x4e, 0x65,
	0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x41, 0x50, 0x49, 0x12, 0xb5, 0x01, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e,
	0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65,
	0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3d, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xe8, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x6e, 0x65, 0x77, 0x73, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0xb1, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03,
	0x08, 0xe7, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73,
	0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x6e, 0x65, 0x77, 0x73, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0xb9, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77,
	0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x30, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73,
	0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65,
	0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3e, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xe7, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2e, 0x3a, 0x01, 0x2a, 0x22, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x67,
	0x65, 0x74, 0x6e, 0x65, 0x77, 0x73, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x62, 0x79, 0x69,
	0x64, 0x12, 0xb5, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0xba, 0xb8, 0x91, 0x02,
	0x05, 0x0a, 0x03, 0x08, 0xe8, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22,
	0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e,
	0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6e, 0x65,
	0x77, 0x73, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0xc9, 0x01, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x12, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x42, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xe9, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0xc5, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12,
	0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e,
	0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0xba, 0xb8, 0x91, 0x02,
	0x05, 0x0a, 0x03, 0x08, 0xe7, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x3a, 0x01, 0x2a, 0x22,
	0x2c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e,
	0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0xcd, 0x01,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x35, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f,
	0x6d, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a,
	0x03, 0x08, 0xe7, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x3a, 0x01, 0x2a, 0x22, 0x2e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77,
	0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x62, 0x79, 0x69, 0x64, 0x12, 0xa1, 0x01,
	0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x2a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65,
	0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03,
	0x08, 0xe7, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73,
	0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0xa9, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x46, 0x6f, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x4e, 0x65, 0x77, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x77, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3a, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xe7, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x67,
	0x65, 0x74, 0x6e, 0x65, 0x77, 0x73, 0x66, 0x6f, 0x72, 0x75, 0x73, 0x65, 0x72, 0x12, 0xc5, 0x01,
	0x0a, 0x15, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73,
	0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x41, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xe8, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x31, 0x3a, 0x01, 0x2a, 0x22, 0x2c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x73, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0xd1, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x12, 0x36, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x46, 0x6f, 0x72,
	0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x44, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xe7, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x34, 0x3a, 0x01, 0x2a, 0x22, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x66, 0x6f, 0x72, 0x6e, 0x65,
	0x77, 0x73, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0xc9, 0x01, 0x0a, 0x16, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x42, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xe8, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x6e, 0x65, 0x77, 0x73, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0xd1, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x77,
	0x73, 0x72, 0x6f, 0x6f, 0x6d, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x4e, 0xaa, 0x02, 0x15, 0x41, 0x70,
	0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0xca, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x4e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0xe2, 0x02, 0x21, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x4e, 0x65, 0x77, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x17, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a,
	0x3a, 0x4e, 0x65, 0x77, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_api_v1alpha1_newsroom_service_proto_goTypes = []any{
	(*CreateNewsArticleRequest)(nil),         // 0: api.v1alpha1.newsroom.CreateNewsArticleRequest
	(*ListNewsArticlesRequest)(nil),          // 1: api.v1alpha1.newsroom.ListNewsArticlesRequest
	(*GetNewsArticleByIdRequest)(nil),        // 2: api.v1alpha1.newsroom.GetNewsArticleByIdRequest
	(*UpdateNewsArticleRequest)(nil),         // 3: api.v1alpha1.newsroom.UpdateNewsArticleRequest
	(*CreatePublishedArticleRequest)(nil),    // 4: api.v1alpha1.newsroom.CreatePublishedArticleRequest
	(*ListPublishedArticlesRequest)(nil),     // 5: api.v1alpha1.newsroom.ListPublishedArticlesRequest
	(*GetPublishedArticleByIdRequest)(nil),   // 6: api.v1alpha1.newsroom.GetPublishedArticleByIdRequest
	(*UserActivityRequest)(nil),              // 7: api.v1alpha1.newsroom.UserActivityRequest
	(*GetNewsForUserRequest)(nil),            // 8: api.v1alpha1.newsroom.GetNewsForUserRequest
	(*StoreNewsArticleImageRequest)(nil),     // 9: api.v1alpha1.newsroom.StoreNewsArticleImageRequest
	(*ListImagesForNewsArticleRequest)(nil),  // 10: api.v1alpha1.newsroom.ListImagesForNewsArticleRequest
	(*UploadNewsArticleImageRequest)(nil),    // 11: api.v1alpha1.newsroom.UploadNewsArticleImageRequest
	(*CreateNewsArticleResponse)(nil),        // 12: api.v1alpha1.newsroom.CreateNewsArticleResponse
	(*ListNewsArticlesResponse)(nil),         // 13: api.v1alpha1.newsroom.ListNewsArticlesResponse
	(*GetNewsArticleByIdResponse)(nil),       // 14: api.v1alpha1.newsroom.GetNewsArticleByIdResponse
	(*UpdateNewsArticleResponse)(nil),        // 15: api.v1alpha1.newsroom.UpdateNewsArticleResponse
	(*CreatePublishedArticleResponse)(nil),   // 16: api.v1alpha1.newsroom.CreatePublishedArticleResponse
	(*ListPublishedArticlesResponse)(nil),    // 17: api.v1alpha1.newsroom.ListPublishedArticlesResponse
	(*GetPublishedArticleByIdResponse)(nil),  // 18: api.v1alpha1.newsroom.GetPublishedArticleByIdResponse
	(*UserActivityResponse)(nil),             // 19: api.v1alpha1.newsroom.UserActivityResponse
	(*GetNewsForUserResponse)(nil),           // 20: api.v1alpha1.newsroom.GetNewsForUserResponse
	(*StoreNewsArticleImageResponse)(nil),    // 21: api.v1alpha1.newsroom.StoreNewsArticleImageResponse
	(*ListImagesForNewsArticleResponse)(nil), // 22: api.v1alpha1.newsroom.ListImagesForNewsArticleResponse
	(*UploadNewsArticleImageResponse)(nil),   // 23: api.v1alpha1.newsroom.UploadNewsArticleImageResponse
}
var file_api_v1alpha1_newsroom_service_proto_depIdxs = []int32{
	0,  // 0: api.v1alpha1.newsroom.NewsroomAPI.CreateNewsArticle:input_type -> api.v1alpha1.newsroom.CreateNewsArticleRequest
	1,  // 1: api.v1alpha1.newsroom.NewsroomAPI.ListNewsArticles:input_type -> api.v1alpha1.newsroom.ListNewsArticlesRequest
	2,  // 2: api.v1alpha1.newsroom.NewsroomAPI.GetNewsArticleById:input_type -> api.v1alpha1.newsroom.GetNewsArticleByIdRequest
	3,  // 3: api.v1alpha1.newsroom.NewsroomAPI.UpdateNewsArticle:input_type -> api.v1alpha1.newsroom.UpdateNewsArticleRequest
	4,  // 4: api.v1alpha1.newsroom.NewsroomAPI.CreatePublishedArticle:input_type -> api.v1alpha1.newsroom.CreatePublishedArticleRequest
	5,  // 5: api.v1alpha1.newsroom.NewsroomAPI.ListPublishedArticles:input_type -> api.v1alpha1.newsroom.ListPublishedArticlesRequest
	6,  // 6: api.v1alpha1.newsroom.NewsroomAPI.GetPublishedArticleById:input_type -> api.v1alpha1.newsroom.GetPublishedArticleByIdRequest
	7,  // 7: api.v1alpha1.newsroom.NewsroomAPI.UserActivity:input_type -> api.v1alpha1.newsroom.UserActivityRequest
	8,  // 8: api.v1alpha1.newsroom.NewsroomAPI.GetNewsForUser:input_type -> api.v1alpha1.newsroom.GetNewsForUserRequest
	9,  // 9: api.v1alpha1.newsroom.NewsroomAPI.StoreNewsArticleImage:input_type -> api.v1alpha1.newsroom.StoreNewsArticleImageRequest
	10, // 10: api.v1alpha1.newsroom.NewsroomAPI.ListImagesForNewsArticle:input_type -> api.v1alpha1.newsroom.ListImagesForNewsArticleRequest
	11, // 11: api.v1alpha1.newsroom.NewsroomAPI.UploadNewsArticleImage:input_type -> api.v1alpha1.newsroom.UploadNewsArticleImageRequest
	12, // 12: api.v1alpha1.newsroom.NewsroomAPI.CreateNewsArticle:output_type -> api.v1alpha1.newsroom.CreateNewsArticleResponse
	13, // 13: api.v1alpha1.newsroom.NewsroomAPI.ListNewsArticles:output_type -> api.v1alpha1.newsroom.ListNewsArticlesResponse
	14, // 14: api.v1alpha1.newsroom.NewsroomAPI.GetNewsArticleById:output_type -> api.v1alpha1.newsroom.GetNewsArticleByIdResponse
	15, // 15: api.v1alpha1.newsroom.NewsroomAPI.UpdateNewsArticle:output_type -> api.v1alpha1.newsroom.UpdateNewsArticleResponse
	16, // 16: api.v1alpha1.newsroom.NewsroomAPI.CreatePublishedArticle:output_type -> api.v1alpha1.newsroom.CreatePublishedArticleResponse
	17, // 17: api.v1alpha1.newsroom.NewsroomAPI.ListPublishedArticles:output_type -> api.v1alpha1.newsroom.ListPublishedArticlesResponse
	18, // 18: api.v1alpha1.newsroom.NewsroomAPI.GetPublishedArticleById:output_type -> api.v1alpha1.newsroom.GetPublishedArticleByIdResponse
	19, // 19: api.v1alpha1.newsroom.NewsroomAPI.UserActivity:output_type -> api.v1alpha1.newsroom.UserActivityResponse
	20, // 20: api.v1alpha1.newsroom.NewsroomAPI.GetNewsForUser:output_type -> api.v1alpha1.newsroom.GetNewsForUserResponse
	21, // 21: api.v1alpha1.newsroom.NewsroomAPI.StoreNewsArticleImage:output_type -> api.v1alpha1.newsroom.StoreNewsArticleImageResponse
	22, // 22: api.v1alpha1.newsroom.NewsroomAPI.ListImagesForNewsArticle:output_type -> api.v1alpha1.newsroom.ListImagesForNewsArticleResponse
	23, // 23: api.v1alpha1.newsroom.NewsroomAPI.UploadNewsArticleImage:output_type -> api.v1alpha1.newsroom.UploadNewsArticleImageResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_newsroom_service_proto_init() }
func file_api_v1alpha1_newsroom_service_proto_init() {
	if File_api_v1alpha1_newsroom_service_proto != nil {
		return
	}
	file_api_v1alpha1_newsroom_entities_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_newsroom_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_newsroom_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_newsroom_service_proto_depIdxs,
	}.Build()
	File_api_v1alpha1_newsroom_service_proto = out.File
	file_api_v1alpha1_newsroom_service_proto_rawDesc = nil
	file_api_v1alpha1_newsroom_service_proto_goTypes = nil
	file_api_v1alpha1_newsroom_service_proto_depIdxs = nil
}
