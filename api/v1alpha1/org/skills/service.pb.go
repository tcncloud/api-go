// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: api/v1alpha1/org/skills/service.proto

package skills

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

var File_api_v1alpha1_org_skills_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_org_skills_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x99, 0x16, 0x0a, 0x0d, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xb2, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xba, 0xb8, 0x91,
	0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22,
	0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0xae, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0xba,
	0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01,
	0x2a, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0xb2, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x39, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0xa6, 0x01, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0xba,
	0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01,
	0x2a, 0x22, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0xb2, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x39, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a,
	0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0xd2, 0x01, 0x0a, 0x18, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x6c,
	0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x46, 0x72, 0x6f,
	0x6d, 0x41, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x39, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x6c, 0x6c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0xba, 0xb8,
	0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a,
	0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x66, 0x72, 0x6f, 0x6d, 0x61, 0x6c, 0x6c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12,
	0xb6, 0x01, 0x0a, 0x11, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0xba, 0xb8,
	0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x77, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x3a, 0x01, 0x2a,
	0x22, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0xcd, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x37, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x4f, 0x6e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02,
	0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a, 0x01, 0x2a, 0x22, 0x2b, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x75, 0x73, 0x65, 0x72, 0x6f, 0x6e, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0xb6, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x31,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x77,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f,
	0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0xba, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3b, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x65, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0xa6,
	0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x36, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27,
	0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0xc2, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x12, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0xba,
	0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01,
	0x2a, 0x22, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0xca, 0x01, 0x0a,
	0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x36, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x37, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a,
	0x02, 0x08, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a, 0x01, 0x2a, 0x22, 0x2b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0xa8, 0x01, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x2e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0xba,
	0xb8, 0x91, 0x02, 0x02, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22,
	0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x12, 0xd4, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x39, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0xba, 0xb8, 0x91, 0x02, 0x02,
	0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x3a, 0x01, 0x2a, 0x22, 0x2e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x66, 0x6f, 0x72, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x42, 0xdf, 0x01, 0x0a, 0x1b,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x42, 0x0c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0xa2,
	0x02, 0x04, 0x41, 0x56, 0x4f, 0x53, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0xca, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x4f, 0x72, 0x67, 0x5c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_v1alpha1_org_skills_service_proto_goTypes = []interface{}{
	(*CreateSkillGroupRequest)(nil),           // 0: api.v1alpha1.org.skills.CreateSkillGroupRequest
	(*ListSkillGroupsRequest)(nil),            // 1: api.v1alpha1.org.skills.ListSkillGroupsRequest
	(*UpdateSkillGroupRequest)(nil),           // 2: api.v1alpha1.org.skills.UpdateSkillGroupRequest
	(*GetSkillGroupRequest)(nil),              // 3: api.v1alpha1.org.skills.GetSkillGroupRequest
	(*DeleteSkillGroupRequest)(nil),           // 4: api.v1alpha1.org.skills.DeleteSkillGroupRequest
	(*RemoveSkillFromAllGroupsRequest)(nil),   // 5: api.v1alpha1.org.skills.RemoveSkillFromAllGroupsRequest
	(*AssignSkillGroupsRequest)(nil),          // 6: api.v1alpha1.org.skills.AssignSkillGroupsRequest
	(*UpdateUsersOnSkillGroupRequest)(nil),    // 7: api.v1alpha1.org.skills.UpdateUsersOnSkillGroupRequest
	(*RevokeSkillGroupsRequest)(nil),          // 8: api.v1alpha1.org.skills.RevokeSkillGroupsRequest
	(*GetUserSkillGroupsRequest)(nil),         // 9: api.v1alpha1.org.skills.GetUserSkillGroupsRequest
	(*GetUserSkillsRequest)(nil),              // 10: api.v1alpha1.org.skills.GetUserSkillsRequest
	(*GetSkillGroupMembersRequest)(nil),       // 11: api.v1alpha1.org.skills.GetSkillGroupMembersRequest
	(*ListSkillGroupsMembersRequest)(nil),     // 12: api.v1alpha1.org.skills.ListSkillGroupsMembersRequest
	(*GetAgentSkillsRequest)(nil),             // 13: api.v1alpha1.org.skills.GetAgentSkillsRequest
	(*ListSkillsForCurrentAgentRequest)(nil),  // 14: api.v1alpha1.org.skills.ListSkillsForCurrentAgentRequest
	(*CreateSkillGroupResponse)(nil),          // 15: api.v1alpha1.org.skills.CreateSkillGroupResponse
	(*ListSkillGroupsResponse)(nil),           // 16: api.v1alpha1.org.skills.ListSkillGroupsResponse
	(*UpdateSkillGroupResponse)(nil),          // 17: api.v1alpha1.org.skills.UpdateSkillGroupResponse
	(*GetSkillGroupResponse)(nil),             // 18: api.v1alpha1.org.skills.GetSkillGroupResponse
	(*DeleteSkillGroupResponse)(nil),          // 19: api.v1alpha1.org.skills.DeleteSkillGroupResponse
	(*RemoveSkillFromAllGroupsResponse)(nil),  // 20: api.v1alpha1.org.skills.RemoveSkillFromAllGroupsResponse
	(*AssignSkillGroupsResponse)(nil),         // 21: api.v1alpha1.org.skills.AssignSkillGroupsResponse
	(*UpdateUsersOnSkillGroupResponse)(nil),   // 22: api.v1alpha1.org.skills.UpdateUsersOnSkillGroupResponse
	(*RevokeSkillGroupsResponse)(nil),         // 23: api.v1alpha1.org.skills.RevokeSkillGroupsResponse
	(*GetUserSkillGroupsResponse)(nil),        // 24: api.v1alpha1.org.skills.GetUserSkillGroupsResponse
	(*GetUserSkillsResponse)(nil),             // 25: api.v1alpha1.org.skills.GetUserSkillsResponse
	(*GetSkillGroupMembersResponse)(nil),      // 26: api.v1alpha1.org.skills.GetSkillGroupMembersResponse
	(*ListSkillGroupsMembersResponse)(nil),    // 27: api.v1alpha1.org.skills.ListSkillGroupsMembersResponse
	(*GetAgentSkillsResponse)(nil),            // 28: api.v1alpha1.org.skills.GetAgentSkillsResponse
	(*ListSkillsForCurrentAgentResponse)(nil), // 29: api.v1alpha1.org.skills.ListSkillsForCurrentAgentResponse
}
var file_api_v1alpha1_org_skills_service_proto_depIdxs = []int32{
	0,  // 0: api.v1alpha1.org.skills.SkillsService.CreateSkillGroup:input_type -> api.v1alpha1.org.skills.CreateSkillGroupRequest
	1,  // 1: api.v1alpha1.org.skills.SkillsService.ListSkillGroups:input_type -> api.v1alpha1.org.skills.ListSkillGroupsRequest
	2,  // 2: api.v1alpha1.org.skills.SkillsService.UpdateSkillGroup:input_type -> api.v1alpha1.org.skills.UpdateSkillGroupRequest
	3,  // 3: api.v1alpha1.org.skills.SkillsService.GetSkillGroup:input_type -> api.v1alpha1.org.skills.GetSkillGroupRequest
	4,  // 4: api.v1alpha1.org.skills.SkillsService.DeleteSkillGroup:input_type -> api.v1alpha1.org.skills.DeleteSkillGroupRequest
	5,  // 5: api.v1alpha1.org.skills.SkillsService.RemoveSkillFromAllGroups:input_type -> api.v1alpha1.org.skills.RemoveSkillFromAllGroupsRequest
	6,  // 6: api.v1alpha1.org.skills.SkillsService.AssignSkillGroups:input_type -> api.v1alpha1.org.skills.AssignSkillGroupsRequest
	7,  // 7: api.v1alpha1.org.skills.SkillsService.UpdateUsersOnSkillGroup:input_type -> api.v1alpha1.org.skills.UpdateUsersOnSkillGroupRequest
	8,  // 8: api.v1alpha1.org.skills.SkillsService.RevokeSkillGroups:input_type -> api.v1alpha1.org.skills.RevokeSkillGroupsRequest
	9,  // 9: api.v1alpha1.org.skills.SkillsService.GetUserSkillGroups:input_type -> api.v1alpha1.org.skills.GetUserSkillGroupsRequest
	10, // 10: api.v1alpha1.org.skills.SkillsService.GetUserSkills:input_type -> api.v1alpha1.org.skills.GetUserSkillsRequest
	11, // 11: api.v1alpha1.org.skills.SkillsService.GetSkillGroupMembers:input_type -> api.v1alpha1.org.skills.GetSkillGroupMembersRequest
	12, // 12: api.v1alpha1.org.skills.SkillsService.ListSkillGroupsMembers:input_type -> api.v1alpha1.org.skills.ListSkillGroupsMembersRequest
	13, // 13: api.v1alpha1.org.skills.SkillsService.GetAgentSkills:input_type -> api.v1alpha1.org.skills.GetAgentSkillsRequest
	14, // 14: api.v1alpha1.org.skills.SkillsService.ListSkillsForCurrentAgent:input_type -> api.v1alpha1.org.skills.ListSkillsForCurrentAgentRequest
	15, // 15: api.v1alpha1.org.skills.SkillsService.CreateSkillGroup:output_type -> api.v1alpha1.org.skills.CreateSkillGroupResponse
	16, // 16: api.v1alpha1.org.skills.SkillsService.ListSkillGroups:output_type -> api.v1alpha1.org.skills.ListSkillGroupsResponse
	17, // 17: api.v1alpha1.org.skills.SkillsService.UpdateSkillGroup:output_type -> api.v1alpha1.org.skills.UpdateSkillGroupResponse
	18, // 18: api.v1alpha1.org.skills.SkillsService.GetSkillGroup:output_type -> api.v1alpha1.org.skills.GetSkillGroupResponse
	19, // 19: api.v1alpha1.org.skills.SkillsService.DeleteSkillGroup:output_type -> api.v1alpha1.org.skills.DeleteSkillGroupResponse
	20, // 20: api.v1alpha1.org.skills.SkillsService.RemoveSkillFromAllGroups:output_type -> api.v1alpha1.org.skills.RemoveSkillFromAllGroupsResponse
	21, // 21: api.v1alpha1.org.skills.SkillsService.AssignSkillGroups:output_type -> api.v1alpha1.org.skills.AssignSkillGroupsResponse
	22, // 22: api.v1alpha1.org.skills.SkillsService.UpdateUsersOnSkillGroup:output_type -> api.v1alpha1.org.skills.UpdateUsersOnSkillGroupResponse
	23, // 23: api.v1alpha1.org.skills.SkillsService.RevokeSkillGroups:output_type -> api.v1alpha1.org.skills.RevokeSkillGroupsResponse
	24, // 24: api.v1alpha1.org.skills.SkillsService.GetUserSkillGroups:output_type -> api.v1alpha1.org.skills.GetUserSkillGroupsResponse
	25, // 25: api.v1alpha1.org.skills.SkillsService.GetUserSkills:output_type -> api.v1alpha1.org.skills.GetUserSkillsResponse
	26, // 26: api.v1alpha1.org.skills.SkillsService.GetSkillGroupMembers:output_type -> api.v1alpha1.org.skills.GetSkillGroupMembersResponse
	27, // 27: api.v1alpha1.org.skills.SkillsService.ListSkillGroupsMembers:output_type -> api.v1alpha1.org.skills.ListSkillGroupsMembersResponse
	28, // 28: api.v1alpha1.org.skills.SkillsService.GetAgentSkills:output_type -> api.v1alpha1.org.skills.GetAgentSkillsResponse
	29, // 29: api.v1alpha1.org.skills.SkillsService.ListSkillsForCurrentAgent:output_type -> api.v1alpha1.org.skills.ListSkillsForCurrentAgentResponse
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_org_skills_service_proto_init() }
func file_api_v1alpha1_org_skills_service_proto_init() {
	if File_api_v1alpha1_org_skills_service_proto != nil {
		return
	}
	file_api_v1alpha1_org_skills_entities_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_org_skills_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_org_skills_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_org_skills_service_proto_depIdxs,
	}.Build()
	File_api_v1alpha1_org_skills_service_proto = out.File
	file_api_v1alpha1_org_skills_service_proto_rawDesc = nil
	file_api_v1alpha1_org_skills_service_proto_goTypes = nil
	file_api_v1alpha1_org_skills_service_proto_depIdxs = nil
}
