// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: services/org/exile_certificate_manager/v1alpha1/service.proto

package exile_certificate_managerv1alpha1

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

var File_services_org_exile_certificate_manager_v1alpha1_service_proto protoreflect.FileDescriptor

var file_services_org_exile_certificate_manager_v1alpha1_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65,
	0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78,
	0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x47, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x49, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65,
	0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9e, 0x14, 0x0a, 0x1e,
	0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x94,
	0x02, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x4e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0xba, 0xb8, 0x91, 0x02,
	0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4a, 0x3a, 0x01, 0x2a, 0x22, 0x45,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78,
	0x69, 0x6c, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x94, 0x02, 0x0a, 0x16, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x4e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x4f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x59, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x4a, 0x3a, 0x01, 0x2a, 0x22, 0x45, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x65, 0x78, 0x69, 0x6c,
	0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x9c, 0x02, 0x0a,
	0x18, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x51, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b,
	0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4c, 0x3a,
	0x01, 0x2a, 0x22, 0x47, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xa4, 0x02, 0x0a, 0x1a,
	0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x53,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78,
	0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x5d, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x4e, 0x3a, 0x01, 0x2a, 0x22, 0x49, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x75, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x8f, 0x02, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x69, 0x6c, 0x65,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x4d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c,
	0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0xba, 0xb8, 0x91,
	0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x48, 0x3a, 0x01, 0x2a, 0x22,
	0x43, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65,
	0x78, 0x69, 0x6c, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x9c, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x50, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x51, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08,
	0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4c, 0x3a, 0x01, 0x2a, 0x22, 0x47, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x9c, 0x02, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x50, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x51, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4c, 0x3a, 0x01, 0x2a, 0x22, 0x47, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x9c, 0x02, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x50, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65,
	0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x51, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x4c, 0x3a, 0x01, 0x2a, 0x22, 0x47, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x65,
	0x78, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x98, 0x02, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78, 0x69,
	0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x50,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x78,
	0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x5a, 0xba, 0xb8, 0x91, 0x02, 0x04, 0x0a, 0x02, 0x08, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x4b, 0x3a, 0x01, 0x2a, 0x22, 0x46, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x88, 0x03, 0x0a,
	0x33, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x78,
	0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x4f, 0x45, 0xaa, 0x02, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x2e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x39, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x30, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a,
	0x3a, 0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_services_org_exile_certificate_manager_v1alpha1_service_proto_goTypes = []any{
	(*CreateExileCertificateRequest)(nil),      // 0: services.org.exile_certificate_manager.v1alpha1.CreateExileCertificateRequest
	(*RevokeExileCertificateRequest)(nil),      // 1: services.org.exile_certificate_manager.v1alpha1.RevokeExileCertificateRequest
	(*AssignExileConfigurationRequest)(nil),    // 2: services.org.exile_certificate_manager.v1alpha1.AssignExileConfigurationRequest
	(*UnassignExileConfigurationRequest)(nil),  // 3: services.org.exile_certificate_manager.v1alpha1.UnassignExileConfigurationRequest
	(*ListExileCertificatesRequest)(nil),       // 4: services.org.exile_certificate_manager.v1alpha1.ListExileCertificatesRequest
	(*CreateExileConfigurationRequest)(nil),    // 5: services.org.exile_certificate_manager.v1alpha1.CreateExileConfigurationRequest
	(*UpdateExileConfigurationRequest)(nil),    // 6: services.org.exile_certificate_manager.v1alpha1.UpdateExileConfigurationRequest
	(*DeleteExileConfigurationRequest)(nil),    // 7: services.org.exile_certificate_manager.v1alpha1.DeleteExileConfigurationRequest
	(*ListExileConfigurationsRequest)(nil),     // 8: services.org.exile_certificate_manager.v1alpha1.ListExileConfigurationsRequest
	(*CreateExileCertificateResponse)(nil),     // 9: services.org.exile_certificate_manager.v1alpha1.CreateExileCertificateResponse
	(*RevokeExileCertificateResponse)(nil),     // 10: services.org.exile_certificate_manager.v1alpha1.RevokeExileCertificateResponse
	(*AssignExileConfigurationResponse)(nil),   // 11: services.org.exile_certificate_manager.v1alpha1.AssignExileConfigurationResponse
	(*UnassignExileConfigurationResponse)(nil), // 12: services.org.exile_certificate_manager.v1alpha1.UnassignExileConfigurationResponse
	(*ListExileCertificatesResponse)(nil),      // 13: services.org.exile_certificate_manager.v1alpha1.ListExileCertificatesResponse
	(*CreateExileConfigurationResponse)(nil),   // 14: services.org.exile_certificate_manager.v1alpha1.CreateExileConfigurationResponse
	(*UpdateExileConfigurationResponse)(nil),   // 15: services.org.exile_certificate_manager.v1alpha1.UpdateExileConfigurationResponse
	(*DeleteExileConfigurationResponse)(nil),   // 16: services.org.exile_certificate_manager.v1alpha1.DeleteExileConfigurationResponse
	(*ListExileConfigurationsResponse)(nil),    // 17: services.org.exile_certificate_manager.v1alpha1.ListExileConfigurationsResponse
}
var file_services_org_exile_certificate_manager_v1alpha1_service_proto_depIdxs = []int32{
	0,  // 0: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.CreateExileCertificate:input_type -> services.org.exile_certificate_manager.v1alpha1.CreateExileCertificateRequest
	1,  // 1: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.RevokeExileCertificate:input_type -> services.org.exile_certificate_manager.v1alpha1.RevokeExileCertificateRequest
	2,  // 2: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.AssignExileConfiguration:input_type -> services.org.exile_certificate_manager.v1alpha1.AssignExileConfigurationRequest
	3,  // 3: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.UnassignExileConfiguration:input_type -> services.org.exile_certificate_manager.v1alpha1.UnassignExileConfigurationRequest
	4,  // 4: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.ListExileCertificates:input_type -> services.org.exile_certificate_manager.v1alpha1.ListExileCertificatesRequest
	5,  // 5: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.CreateExileConfiguration:input_type -> services.org.exile_certificate_manager.v1alpha1.CreateExileConfigurationRequest
	6,  // 6: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.UpdateExileConfiguration:input_type -> services.org.exile_certificate_manager.v1alpha1.UpdateExileConfigurationRequest
	7,  // 7: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.DeleteExileConfiguration:input_type -> services.org.exile_certificate_manager.v1alpha1.DeleteExileConfigurationRequest
	8,  // 8: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.ListExileConfigurations:input_type -> services.org.exile_certificate_manager.v1alpha1.ListExileConfigurationsRequest
	9,  // 9: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.CreateExileCertificate:output_type -> services.org.exile_certificate_manager.v1alpha1.CreateExileCertificateResponse
	10, // 10: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.RevokeExileCertificate:output_type -> services.org.exile_certificate_manager.v1alpha1.RevokeExileCertificateResponse
	11, // 11: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.AssignExileConfiguration:output_type -> services.org.exile_certificate_manager.v1alpha1.AssignExileConfigurationResponse
	12, // 12: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.UnassignExileConfiguration:output_type -> services.org.exile_certificate_manager.v1alpha1.UnassignExileConfigurationResponse
	13, // 13: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.ListExileCertificates:output_type -> services.org.exile_certificate_manager.v1alpha1.ListExileCertificatesResponse
	14, // 14: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.CreateExileConfiguration:output_type -> services.org.exile_certificate_manager.v1alpha1.CreateExileConfigurationResponse
	15, // 15: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.UpdateExileConfiguration:output_type -> services.org.exile_certificate_manager.v1alpha1.UpdateExileConfigurationResponse
	16, // 16: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.DeleteExileConfiguration:output_type -> services.org.exile_certificate_manager.v1alpha1.DeleteExileConfigurationResponse
	17, // 17: services.org.exile_certificate_manager.v1alpha1.ExileCertificateManagerService.ListExileConfigurations:output_type -> services.org.exile_certificate_manager.v1alpha1.ListExileConfigurationsResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_org_exile_certificate_manager_v1alpha1_service_proto_init() }
func file_services_org_exile_certificate_manager_v1alpha1_service_proto_init() {
	if File_services_org_exile_certificate_manager_v1alpha1_service_proto != nil {
		return
	}
	file_services_org_exile_certificate_manager_v1alpha1_exile_certificate_proto_init()
	file_services_org_exile_certificate_manager_v1alpha1_exile_configuration_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_org_exile_certificate_manager_v1alpha1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_org_exile_certificate_manager_v1alpha1_service_proto_goTypes,
		DependencyIndexes: file_services_org_exile_certificate_manager_v1alpha1_service_proto_depIdxs,
	}.Build()
	File_services_org_exile_certificate_manager_v1alpha1_service_proto = out.File
	file_services_org_exile_certificate_manager_v1alpha1_service_proto_rawDesc = nil
	file_services_org_exile_certificate_manager_v1alpha1_service_proto_goTypes = nil
	file_services_org_exile_certificate_manager_v1alpha1_service_proto_depIdxs = nil
}
