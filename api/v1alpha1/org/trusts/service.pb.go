// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/v1alpha1/org/trusts/service.proto

package trusts

import (
	_ "github.com/tcncloud/api-go/annotations"
	org "github.com/tcncloud/api-go/api/v1alpha1/org"
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

var File_api_v1alpha1_org_trusts_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_org_trusts_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73,
	0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x68, 0x75, 0x6e, 0x74, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x33, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xda, 0x0c,
	0x0a, 0x0d, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x95, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x75, 0x73, 0x74, 0x12,
	0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xba, 0xb8,
	0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xa0, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01,
	0x2a, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x12, 0x95, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xa0, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x74, 0x72, 0x75, 0x73, 0x74, 0x12,
	0x95, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x12,
	0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xba, 0xb8,
	0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xa0, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01,
	0x2a, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x74, 0x72, 0x75, 0x73, 0x74, 0x12, 0x89, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x75, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x75, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0xba, 0xb8, 0x91,
	0x02, 0x05, 0x0a, 0x03, 0x08, 0xa0, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a,
	0x22, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x6f, 0x72, 0x67, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x12, 0xb6, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xa0,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x3a, 0x01, 0x2a, 0x22, 0x30, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x12, 0xaa, 0x01, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73,
	0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x69, 0x76, 0x65, 0x6e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xa0,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x67, 0x69,
	0x76, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x12, 0xbe, 0x01, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x75, 0x73,
	0x74, 0x73, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x47, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xa0, 0x01, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x37, 0x3a, 0x01, 0x2a, 0x22, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x12, 0x95, 0x01, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x75, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03,
	0x08, 0xa0, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x12, 0x95, 0x01, 0x0a, 0x0b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x72, 0x75, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x39, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xa0, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x2f, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x12, 0x9d, 0x01, 0x0a, 0x0d, 0x55,
	0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0xba,
	0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xa0, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x3a,
	0x01, 0x2a, 0x22, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x2f, 0x75, 0x6e, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x42, 0xdf, 0x01, 0x0a, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0xa2, 0x02,
	0x04, 0x41, 0x56, 0x4f, 0x54, 0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0xca,
	0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x4f,
	0x72, 0x67, 0x5c, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a,
	0x3a, 0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x54, 0x72, 0x75, 0x73, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_api_v1alpha1_org_trusts_service_proto_goTypes = []any{
	(*org.CreateTrustRequest)(nil),           // 0: api.v1alpha1.org.CreateTrustRequest
	(*org.AcceptTrustRequest)(nil),           // 1: api.v1alpha1.org.AcceptTrustRequest
	(*org.RejectTrustRequest)(nil),           // 2: api.v1alpha1.org.RejectTrustRequest
	(*org.GetTrustRequest)(nil),              // 3: api.v1alpha1.org.GetTrustRequest
	(*org.ListIncomingTrustsRequest)(nil),    // 4: api.v1alpha1.org.ListIncomingTrustsRequest
	(*org.ListGivenTrustsRequest)(nil),       // 5: api.v1alpha1.org.ListGivenTrustsRequest
	(*org.ListAssignableTrustsRequest)(nil),  // 6: api.v1alpha1.org.ListAssignableTrustsRequest
	(*org.DeleteTrustRequest)(nil),           // 7: api.v1alpha1.org.DeleteTrustRequest
	(*org.AssignTrustRequest)(nil),           // 8: api.v1alpha1.org.AssignTrustRequest
	(*org.UnassignTrustRequest)(nil),         // 9: api.v1alpha1.org.UnassignTrustRequest
	(*org.CreateTrustResponse)(nil),          // 10: api.v1alpha1.org.CreateTrustResponse
	(*org.AcceptTrustResponse)(nil),          // 11: api.v1alpha1.org.AcceptTrustResponse
	(*org.RejectTrustResponse)(nil),          // 12: api.v1alpha1.org.RejectTrustResponse
	(*org.GetTrustResponse)(nil),             // 13: api.v1alpha1.org.GetTrustResponse
	(*org.ListIncomingTrustsResponse)(nil),   // 14: api.v1alpha1.org.ListIncomingTrustsResponse
	(*org.ListGivenTrustsResponse)(nil),      // 15: api.v1alpha1.org.ListGivenTrustsResponse
	(*org.ListAssignableTrustsResponse)(nil), // 16: api.v1alpha1.org.ListAssignableTrustsResponse
	(*org.DeleteTrustResponse)(nil),          // 17: api.v1alpha1.org.DeleteTrustResponse
	(*org.AssignTrustResponse)(nil),          // 18: api.v1alpha1.org.AssignTrustResponse
	(*org.UnassignTrustResponse)(nil),        // 19: api.v1alpha1.org.UnassignTrustResponse
}
var file_api_v1alpha1_org_trusts_service_proto_depIdxs = []int32{
	0,  // 0: api.v1alpha1.org.trusts.TrustsService.CreateTrust:input_type -> api.v1alpha1.org.CreateTrustRequest
	1,  // 1: api.v1alpha1.org.trusts.TrustsService.AcceptTrust:input_type -> api.v1alpha1.org.AcceptTrustRequest
	2,  // 2: api.v1alpha1.org.trusts.TrustsService.RejectTrust:input_type -> api.v1alpha1.org.RejectTrustRequest
	3,  // 3: api.v1alpha1.org.trusts.TrustsService.GetTrust:input_type -> api.v1alpha1.org.GetTrustRequest
	4,  // 4: api.v1alpha1.org.trusts.TrustsService.ListIncomingTrusts:input_type -> api.v1alpha1.org.ListIncomingTrustsRequest
	5,  // 5: api.v1alpha1.org.trusts.TrustsService.ListGivenTrusts:input_type -> api.v1alpha1.org.ListGivenTrustsRequest
	6,  // 6: api.v1alpha1.org.trusts.TrustsService.ListAssignableTrusts:input_type -> api.v1alpha1.org.ListAssignableTrustsRequest
	7,  // 7: api.v1alpha1.org.trusts.TrustsService.DeleteTrust:input_type -> api.v1alpha1.org.DeleteTrustRequest
	8,  // 8: api.v1alpha1.org.trusts.TrustsService.AssignTrust:input_type -> api.v1alpha1.org.AssignTrustRequest
	9,  // 9: api.v1alpha1.org.trusts.TrustsService.UnassignTrust:input_type -> api.v1alpha1.org.UnassignTrustRequest
	10, // 10: api.v1alpha1.org.trusts.TrustsService.CreateTrust:output_type -> api.v1alpha1.org.CreateTrustResponse
	11, // 11: api.v1alpha1.org.trusts.TrustsService.AcceptTrust:output_type -> api.v1alpha1.org.AcceptTrustResponse
	12, // 12: api.v1alpha1.org.trusts.TrustsService.RejectTrust:output_type -> api.v1alpha1.org.RejectTrustResponse
	13, // 13: api.v1alpha1.org.trusts.TrustsService.GetTrust:output_type -> api.v1alpha1.org.GetTrustResponse
	14, // 14: api.v1alpha1.org.trusts.TrustsService.ListIncomingTrusts:output_type -> api.v1alpha1.org.ListIncomingTrustsResponse
	15, // 15: api.v1alpha1.org.trusts.TrustsService.ListGivenTrusts:output_type -> api.v1alpha1.org.ListGivenTrustsResponse
	16, // 16: api.v1alpha1.org.trusts.TrustsService.ListAssignableTrusts:output_type -> api.v1alpha1.org.ListAssignableTrustsResponse
	17, // 17: api.v1alpha1.org.trusts.TrustsService.DeleteTrust:output_type -> api.v1alpha1.org.DeleteTrustResponse
	18, // 18: api.v1alpha1.org.trusts.TrustsService.AssignTrust:output_type -> api.v1alpha1.org.AssignTrustResponse
	19, // 19: api.v1alpha1.org.trusts.TrustsService.UnassignTrust:output_type -> api.v1alpha1.org.UnassignTrustResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_org_trusts_service_proto_init() }
func file_api_v1alpha1_org_trusts_service_proto_init() {
	if File_api_v1alpha1_org_trusts_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_org_trusts_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_org_trusts_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_org_trusts_service_proto_depIdxs,
	}.Build()
	File_api_v1alpha1_org_trusts_service_proto = out.File
	file_api_v1alpha1_org_trusts_service_proto_rawDesc = nil
	file_api_v1alpha1_org_trusts_service_proto_goTypes = nil
	file_api_v1alpha1_org_trusts_service_proto_depIdxs = nil
}
