// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/v1alpha1/org/businesshours/service.proto

package businesshours

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

var File_api_v1alpha1_org_businesshours_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_org_businesshours_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x1a, 0x17,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbc, 0x0d, 0x0a, 0x14, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc2, 0x01,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x12, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68,
	0x6f, 0x75, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a,
	0x02, 0x08, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0xbe, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x37, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0xba, 0xb8, 0x91, 0x02,
	0x04, 0x0a, 0x02, 0x08, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2f,
	0x67, 0x65, 0x74, 0x12, 0xbe, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x37, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75,
	0x72, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0xba, 0xb8, 0x91,
	0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22,
	0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73,
	0x2f, 0x73, 0x65, 0x74, 0x12, 0xe4, 0x01, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x54, 0x6f, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x12, 0x41, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68,
	0x6f, 0x75, 0x72, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x54, 0x6f, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x54, 0x6f, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0xba, 0xb8, 0x91, 0x02,
	0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a, 0x01, 0x2a, 0x22, 0x2b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2f,
	0x61, 0x64, 0x64, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0xf6, 0x01, 0x0a, 0x1f,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x46, 0x72,
	0x6f, 0x6d, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12,
	0x46, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x46,
	0x72, 0x6f, 0x6d, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x42, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x33, 0x3a, 0x01, 0x2a, 0x22, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x68, 0x6f, 0x75, 0x72, 0x73, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x12, 0xda, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x3e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x48, 0x6f, 0x75, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x48, 0x6f, 0x75, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3e, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2f, 0x3a, 0x01, 0x2a, 0x22, 0x2a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0xca, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x3a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3a, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0xd2,
	0x01, 0x0a, 0x15, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x3c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x42, 0x89, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x68, 0x6f, 0x75, 0x72, 0x73, 0xa2, 0x02, 0x04, 0x41, 0x56, 0x4f, 0x42, 0xaa, 0x02, 0x1e, 0x41,
	0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x2e,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0xca, 0x02, 0x1e,
	0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x4f, 0x72, 0x67,
	0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0xe2, 0x02,
	0x2a, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x4f, 0x72,
	0x67, 0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x4f, 0x72, 0x67,
	0x3a, 0x3a, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_v1alpha1_org_businesshours_service_proto_goTypes = []interface{}{
	(*ListBusinessHoursRequest)(nil),                // 0: api.v1alpha1.org.businesshours.ListBusinessHoursRequest
	(*GetBusinessHoursRequest)(nil),                 // 1: api.v1alpha1.org.businesshours.GetBusinessHoursRequest
	(*SetBusinessHoursRequest)(nil),                 // 2: api.v1alpha1.org.businesshours.SetBusinessHoursRequest
	(*AddIntervalToBusinessHoursRequest)(nil),       // 3: api.v1alpha1.org.businesshours.AddIntervalToBusinessHoursRequest
	(*RemoveIntervalFromBusinessHoursRequest)(nil),  // 4: api.v1alpha1.org.businesshours.RemoveIntervalFromBusinessHoursRequest
	(*UpdateBusinessHoursInfoRequest)(nil),          // 5: api.v1alpha1.org.businesshours.UpdateBusinessHoursInfoRequest
	(*DeleteBusinessHoursRequest)(nil),              // 6: api.v1alpha1.org.businesshours.DeleteBusinessHoursRequest
	(*EvaluateBusinessHoursRequest)(nil),            // 7: api.v1alpha1.org.businesshours.EvaluateBusinessHoursRequest
	(*ListBusinessHoursResponse)(nil),               // 8: api.v1alpha1.org.businesshours.ListBusinessHoursResponse
	(*GetBusinessHoursResponse)(nil),                // 9: api.v1alpha1.org.businesshours.GetBusinessHoursResponse
	(*SetBusinessHoursResponse)(nil),                // 10: api.v1alpha1.org.businesshours.SetBusinessHoursResponse
	(*AddIntervalToBusinessHoursResponse)(nil),      // 11: api.v1alpha1.org.businesshours.AddIntervalToBusinessHoursResponse
	(*RemoveIntervalFromBusinessHoursResponse)(nil), // 12: api.v1alpha1.org.businesshours.RemoveIntervalFromBusinessHoursResponse
	(*UpdateBusinessHoursInfoResponse)(nil),         // 13: api.v1alpha1.org.businesshours.UpdateBusinessHoursInfoResponse
	(*DeleteBusinessHoursResponse)(nil),             // 14: api.v1alpha1.org.businesshours.DeleteBusinessHoursResponse
	(*EvaluateBusinessHoursResponse)(nil),           // 15: api.v1alpha1.org.businesshours.EvaluateBusinessHoursResponse
}
var file_api_v1alpha1_org_businesshours_service_proto_depIdxs = []int32{
	0,  // 0: api.v1alpha1.org.businesshours.BusinessHoursService.ListBusinessHours:input_type -> api.v1alpha1.org.businesshours.ListBusinessHoursRequest
	1,  // 1: api.v1alpha1.org.businesshours.BusinessHoursService.GetBusinessHours:input_type -> api.v1alpha1.org.businesshours.GetBusinessHoursRequest
	2,  // 2: api.v1alpha1.org.businesshours.BusinessHoursService.SetBusinessHours:input_type -> api.v1alpha1.org.businesshours.SetBusinessHoursRequest
	3,  // 3: api.v1alpha1.org.businesshours.BusinessHoursService.AddIntervalToBusinessHours:input_type -> api.v1alpha1.org.businesshours.AddIntervalToBusinessHoursRequest
	4,  // 4: api.v1alpha1.org.businesshours.BusinessHoursService.RemoveIntervalFromBusinessHours:input_type -> api.v1alpha1.org.businesshours.RemoveIntervalFromBusinessHoursRequest
	5,  // 5: api.v1alpha1.org.businesshours.BusinessHoursService.UpdateBusinessHoursInfo:input_type -> api.v1alpha1.org.businesshours.UpdateBusinessHoursInfoRequest
	6,  // 6: api.v1alpha1.org.businesshours.BusinessHoursService.DeleteBusinessHours:input_type -> api.v1alpha1.org.businesshours.DeleteBusinessHoursRequest
	7,  // 7: api.v1alpha1.org.businesshours.BusinessHoursService.EvaluateBusinessHours:input_type -> api.v1alpha1.org.businesshours.EvaluateBusinessHoursRequest
	8,  // 8: api.v1alpha1.org.businesshours.BusinessHoursService.ListBusinessHours:output_type -> api.v1alpha1.org.businesshours.ListBusinessHoursResponse
	9,  // 9: api.v1alpha1.org.businesshours.BusinessHoursService.GetBusinessHours:output_type -> api.v1alpha1.org.businesshours.GetBusinessHoursResponse
	10, // 10: api.v1alpha1.org.businesshours.BusinessHoursService.SetBusinessHours:output_type -> api.v1alpha1.org.businesshours.SetBusinessHoursResponse
	11, // 11: api.v1alpha1.org.businesshours.BusinessHoursService.AddIntervalToBusinessHours:output_type -> api.v1alpha1.org.businesshours.AddIntervalToBusinessHoursResponse
	12, // 12: api.v1alpha1.org.businesshours.BusinessHoursService.RemoveIntervalFromBusinessHours:output_type -> api.v1alpha1.org.businesshours.RemoveIntervalFromBusinessHoursResponse
	13, // 13: api.v1alpha1.org.businesshours.BusinessHoursService.UpdateBusinessHoursInfo:output_type -> api.v1alpha1.org.businesshours.UpdateBusinessHoursInfoResponse
	14, // 14: api.v1alpha1.org.businesshours.BusinessHoursService.DeleteBusinessHours:output_type -> api.v1alpha1.org.businesshours.DeleteBusinessHoursResponse
	15, // 15: api.v1alpha1.org.businesshours.BusinessHoursService.EvaluateBusinessHours:output_type -> api.v1alpha1.org.businesshours.EvaluateBusinessHoursResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_org_businesshours_service_proto_init() }
func file_api_v1alpha1_org_businesshours_service_proto_init() {
	if File_api_v1alpha1_org_businesshours_service_proto != nil {
		return
	}
	file_api_v1alpha1_org_businesshours_entities_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_org_businesshours_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_org_businesshours_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_org_businesshours_service_proto_depIdxs,
	}.Build()
	File_api_v1alpha1_org_businesshours_service_proto = out.File
	file_api_v1alpha1_org_businesshours_service_proto_rawDesc = nil
	file_api_v1alpha1_org_businesshours_service_proto_goTypes = nil
	file_api_v1alpha1_org_businesshours_service_proto_depIdxs = nil
}
