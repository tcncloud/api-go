// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v1alpha1/agenttraining/service.proto

package agenttraining

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

var File_api_v1alpha1_agenttraining_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_agenttraining_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x35, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9e, 0x0e, 0x0a, 0x14, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xf9, 0x01,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x3c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a,
	0x03, 0x08, 0xc4, 0x0c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4f, 0x3a, 0x01, 0x2a, 0x22, 0x4a, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x6f, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0xf9, 0x01, 0x0a, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xc4, 0x0c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4f, 0x3a, 0x01, 0x2a, 0x22, 0x4a, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x8d, 0x02, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x41, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4c,
	0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x64, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xd4, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x54, 0x3a, 0x01, 0x2a, 0x22, 0x4f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x95, 0x02, 0x0a, 0x20, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x43, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x44, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xd4,
	0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x56, 0x3a, 0x01, 0x2a, 0x22, 0x51, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0xf9, 0x01,
	0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x3c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c,
	0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a,
	0x03, 0x08, 0xc4, 0x0c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4f, 0x3a, 0x01, 0x2a, 0x22, 0x4a, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x6f, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0xf9, 0x01, 0x0a, 0x19, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x3c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xc4, 0x0c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4f, 0x3a, 0x01, 0x2a, 0x22, 0x4a, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0xed, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x12, 0x39, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03,
	0x08, 0xc4, 0x0c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4c, 0x3a, 0x01, 0x2a, 0x22, 0x47, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67,
	0x65, 0x74, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x42, 0xef, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0xa2,
	0x02, 0x03, 0x41, 0x56, 0x41, 0xaa, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0xca, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0xe2,
	0x02, 0x26, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_v1alpha1_agenttraining_service_proto_goTypes = []interface{}{
	(*CreateLearningOpportunityRequest)(nil),         // 0: api.v1alpha1.agenttraining.CreateLearningOpportunityRequest
	(*ListLearningOpportunitiesRequest)(nil),         // 1: api.v1alpha1.agenttraining.ListLearningOpportunitiesRequest
	(*ListAgentLearningOpportunitiesRequest)(nil),    // 2: api.v1alpha1.agenttraining.ListAgentLearningOpportunitiesRequest
	(*CompleteAgentLearningOpportunityRequest)(nil),  // 3: api.v1alpha1.agenttraining.CompleteAgentLearningOpportunityRequest
	(*UpdateLearningOpportunityRequest)(nil),         // 4: api.v1alpha1.agenttraining.UpdateLearningOpportunityRequest
	(*DeleteLearningOpportunityRequest)(nil),         // 5: api.v1alpha1.agenttraining.DeleteLearningOpportunityRequest
	(*GetLearningOpportunityRequest)(nil),            // 6: api.v1alpha1.agenttraining.GetLearningOpportunityRequest
	(*CreateLearningOpportunityResponse)(nil),        // 7: api.v1alpha1.agenttraining.CreateLearningOpportunityResponse
	(*ListLearningOpportunitiesResponse)(nil),        // 8: api.v1alpha1.agenttraining.ListLearningOpportunitiesResponse
	(*ListAgentLearningOpportunitiesResponse)(nil),   // 9: api.v1alpha1.agenttraining.ListAgentLearningOpportunitiesResponse
	(*CompleteAgentLearningOpportunityResponse)(nil), // 10: api.v1alpha1.agenttraining.CompleteAgentLearningOpportunityResponse
	(*UpdateLearningOpportunityResponse)(nil),        // 11: api.v1alpha1.agenttraining.UpdateLearningOpportunityResponse
	(*DeleteLearningOpportunityResponse)(nil),        // 12: api.v1alpha1.agenttraining.DeleteLearningOpportunityResponse
	(*GetLearningOpportunityResponse)(nil),           // 13: api.v1alpha1.agenttraining.GetLearningOpportunityResponse
}
var file_api_v1alpha1_agenttraining_service_proto_depIdxs = []int32{
	0,  // 0: api.v1alpha1.agenttraining.AgentTrainingService.CreateLearningOpportunity:input_type -> api.v1alpha1.agenttraining.CreateLearningOpportunityRequest
	1,  // 1: api.v1alpha1.agenttraining.AgentTrainingService.ListLearningOpportunities:input_type -> api.v1alpha1.agenttraining.ListLearningOpportunitiesRequest
	2,  // 2: api.v1alpha1.agenttraining.AgentTrainingService.ListAgentLearningOpportunities:input_type -> api.v1alpha1.agenttraining.ListAgentLearningOpportunitiesRequest
	3,  // 3: api.v1alpha1.agenttraining.AgentTrainingService.CompleteAgentLearningOpportunity:input_type -> api.v1alpha1.agenttraining.CompleteAgentLearningOpportunityRequest
	4,  // 4: api.v1alpha1.agenttraining.AgentTrainingService.UpdateLearningOpportunity:input_type -> api.v1alpha1.agenttraining.UpdateLearningOpportunityRequest
	5,  // 5: api.v1alpha1.agenttraining.AgentTrainingService.DeleteLearningOpportunity:input_type -> api.v1alpha1.agenttraining.DeleteLearningOpportunityRequest
	6,  // 6: api.v1alpha1.agenttraining.AgentTrainingService.GetLearningOpportunity:input_type -> api.v1alpha1.agenttraining.GetLearningOpportunityRequest
	7,  // 7: api.v1alpha1.agenttraining.AgentTrainingService.CreateLearningOpportunity:output_type -> api.v1alpha1.agenttraining.CreateLearningOpportunityResponse
	8,  // 8: api.v1alpha1.agenttraining.AgentTrainingService.ListLearningOpportunities:output_type -> api.v1alpha1.agenttraining.ListLearningOpportunitiesResponse
	9,  // 9: api.v1alpha1.agenttraining.AgentTrainingService.ListAgentLearningOpportunities:output_type -> api.v1alpha1.agenttraining.ListAgentLearningOpportunitiesResponse
	10, // 10: api.v1alpha1.agenttraining.AgentTrainingService.CompleteAgentLearningOpportunity:output_type -> api.v1alpha1.agenttraining.CompleteAgentLearningOpportunityResponse
	11, // 11: api.v1alpha1.agenttraining.AgentTrainingService.UpdateLearningOpportunity:output_type -> api.v1alpha1.agenttraining.UpdateLearningOpportunityResponse
	12, // 12: api.v1alpha1.agenttraining.AgentTrainingService.DeleteLearningOpportunity:output_type -> api.v1alpha1.agenttraining.DeleteLearningOpportunityResponse
	13, // 13: api.v1alpha1.agenttraining.AgentTrainingService.GetLearningOpportunity:output_type -> api.v1alpha1.agenttraining.GetLearningOpportunityResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_agenttraining_service_proto_init() }
func file_api_v1alpha1_agenttraining_service_proto_init() {
	if File_api_v1alpha1_agenttraining_service_proto != nil {
		return
	}
	file_api_v1alpha1_agenttraining_learning_opportunity_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_agenttraining_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_agenttraining_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_agenttraining_service_proto_depIdxs,
	}.Build()
	File_api_v1alpha1_agenttraining_service_proto = out.File
	file_api_v1alpha1_agenttraining_service_proto_rawDesc = nil
	file_api_v1alpha1_agenttraining_service_proto_goTypes = nil
	file_api_v1alpha1_agenttraining_service_proto_depIdxs = nil
}
