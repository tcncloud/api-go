// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: services/org/hunt_groups/v1alpha1/service.proto

package hunt_groupsv1alpha1

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

var File_services_org_hunt_groups_v1alpha1_service_proto protoreflect.FileDescriptor

var file_services_org_hunt_groups_v1alpha1_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x68,
	0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x88, 0x0e,
	0x0a, 0x11, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xf0, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x6e, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12,
	0x41, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68,
	0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x42, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x6e, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08,
	0xec, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3e, 0x3a, 0x01, 0x2a, 0x22, 0x39, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x68, 0x75, 0x6e, 0x74, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x68, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x78, 0x69, 0x6c,
	0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0xec, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x70, 0x79, 0x48,
	0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x40, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x48, 0x75, 0x6e, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08,
	0xed, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x3a, 0x01, 0x2a, 0x22, 0x38, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x68, 0x75, 0x6e, 0x74, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63,
	0x6f, 0x70, 0x79, 0x68, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x78, 0x69, 0x6c,
	0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0xf8, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x12, 0x43, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x75,
	0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x69, 0x6c,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50,
	0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xed, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x40,
	0x3a, 0x01, 0x2a, 0x22, 0x3b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x68, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x68, 0x75, 0x6e,
	0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x12, 0xfc, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12,
	0x44, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68,
	0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75,
	0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0xba, 0xb8,
	0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xec, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x3a, 0x01,
	0x2a, 0x22, 0x3c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67,
	0x2f, 0x68, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x68, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12,
	0xf8, 0x01, 0x0a, 0x19, 0x43, 0x6f, 0x70, 0x79, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x43, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e,
	0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x44, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x48, 0x75, 0x6e, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a,
	0x03, 0x08, 0xed, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x40, 0x3a, 0x01, 0x2a, 0x22, 0x3b, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x68, 0x75, 0x6e,
	0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x63, 0x6f, 0x70, 0x79, 0x68, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x84, 0x02, 0x0a, 0x1c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12, 0x46, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x75,
	0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0xba, 0xb8,
	0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xed, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x43, 0x3a, 0x01,
	0x2a, 0x22, 0x3e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67,
	0x2f, 0x68, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x68, 0x75, 0x6e, 0x74, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x12, 0x94, 0x02, 0x0a, 0x20, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x70, 0x79, 0x48,
	0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x6f, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x43, 0x6f, 0x70, 0x79, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x6f, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x4b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x70, 0x79,
	0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x6f, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x57, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xc8, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x47, 0x3a, 0x01, 0x2a, 0x22, 0x42, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x6f, 0x72, 0x67, 0x2f, 0x68, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x63, 0x6f, 0x70,
	0x79, 0x68, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x74, 0x6f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xaa, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x75,
	0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x68, 0x75, 0x6e, 0x74,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x4f, 0x48, 0xaa, 0x02, 0x20, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x2e, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x20,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x48, 0x75, 0x6e,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x2c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x4f, 0x72, 0x67, 0x5c,
	0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x23, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x3a,
	0x3a, 0x48, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_services_org_hunt_groups_v1alpha1_service_proto_goTypes = []any{
	(*ListHuntGroupExileLinksRequest)(nil),           // 0: services.org.hunt_groups.v1alpha1.ListHuntGroupExileLinksRequest
	(*CopyHuntGroupExileLinkRequest)(nil),            // 1: services.org.hunt_groups.v1alpha1.CopyHuntGroupExileLinkRequest
	(*UpdateHuntGroupExileLinksRequest)(nil),         // 2: services.org.hunt_groups.v1alpha1.UpdateHuntGroupExileLinksRequest
	(*ListHuntGroupAgentTriggersRequest)(nil),        // 3: services.org.hunt_groups.v1alpha1.ListHuntGroupAgentTriggersRequest
	(*CopyHuntGroupAgentTriggerRequest)(nil),         // 4: services.org.hunt_groups.v1alpha1.CopyHuntGroupAgentTriggerRequest
	(*UpdateHuntGroupAgentTriggersRequest)(nil),      // 5: services.org.hunt_groups.v1alpha1.UpdateHuntGroupAgentTriggersRequest
	(*AdminCopyHuntGroupToOrganizationRequest)(nil),  // 6: services.org.hunt_groups.v1alpha1.AdminCopyHuntGroupToOrganizationRequest
	(*ListHuntGroupExileLinksResponse)(nil),          // 7: services.org.hunt_groups.v1alpha1.ListHuntGroupExileLinksResponse
	(*CopyHuntGroupExileLinkResponse)(nil),           // 8: services.org.hunt_groups.v1alpha1.CopyHuntGroupExileLinkResponse
	(*UpdateHuntGroupExileLinksResponse)(nil),        // 9: services.org.hunt_groups.v1alpha1.UpdateHuntGroupExileLinksResponse
	(*ListHuntGroupAgentTriggersResponse)(nil),       // 10: services.org.hunt_groups.v1alpha1.ListHuntGroupAgentTriggersResponse
	(*CopyHuntGroupAgentTriggerResponse)(nil),        // 11: services.org.hunt_groups.v1alpha1.CopyHuntGroupAgentTriggerResponse
	(*UpdateHuntGroupAgentTriggersResponse)(nil),     // 12: services.org.hunt_groups.v1alpha1.UpdateHuntGroupAgentTriggersResponse
	(*AdminCopyHuntGroupToOrganizationResponse)(nil), // 13: services.org.hunt_groups.v1alpha1.AdminCopyHuntGroupToOrganizationResponse
}
var file_services_org_hunt_groups_v1alpha1_service_proto_depIdxs = []int32{
	0,  // 0: services.org.hunt_groups.v1alpha1.HuntGroupsService.ListHuntGroupExileLinks:input_type -> services.org.hunt_groups.v1alpha1.ListHuntGroupExileLinksRequest
	1,  // 1: services.org.hunt_groups.v1alpha1.HuntGroupsService.CopyHuntGroupExileLink:input_type -> services.org.hunt_groups.v1alpha1.CopyHuntGroupExileLinkRequest
	2,  // 2: services.org.hunt_groups.v1alpha1.HuntGroupsService.UpdateHuntGroupExileLinks:input_type -> services.org.hunt_groups.v1alpha1.UpdateHuntGroupExileLinksRequest
	3,  // 3: services.org.hunt_groups.v1alpha1.HuntGroupsService.ListHuntGroupAgentTriggers:input_type -> services.org.hunt_groups.v1alpha1.ListHuntGroupAgentTriggersRequest
	4,  // 4: services.org.hunt_groups.v1alpha1.HuntGroupsService.CopyHuntGroupAgentTrigger:input_type -> services.org.hunt_groups.v1alpha1.CopyHuntGroupAgentTriggerRequest
	5,  // 5: services.org.hunt_groups.v1alpha1.HuntGroupsService.UpdateHuntGroupAgentTriggers:input_type -> services.org.hunt_groups.v1alpha1.UpdateHuntGroupAgentTriggersRequest
	6,  // 6: services.org.hunt_groups.v1alpha1.HuntGroupsService.AdminCopyHuntGroupToOrganization:input_type -> services.org.hunt_groups.v1alpha1.AdminCopyHuntGroupToOrganizationRequest
	7,  // 7: services.org.hunt_groups.v1alpha1.HuntGroupsService.ListHuntGroupExileLinks:output_type -> services.org.hunt_groups.v1alpha1.ListHuntGroupExileLinksResponse
	8,  // 8: services.org.hunt_groups.v1alpha1.HuntGroupsService.CopyHuntGroupExileLink:output_type -> services.org.hunt_groups.v1alpha1.CopyHuntGroupExileLinkResponse
	9,  // 9: services.org.hunt_groups.v1alpha1.HuntGroupsService.UpdateHuntGroupExileLinks:output_type -> services.org.hunt_groups.v1alpha1.UpdateHuntGroupExileLinksResponse
	10, // 10: services.org.hunt_groups.v1alpha1.HuntGroupsService.ListHuntGroupAgentTriggers:output_type -> services.org.hunt_groups.v1alpha1.ListHuntGroupAgentTriggersResponse
	11, // 11: services.org.hunt_groups.v1alpha1.HuntGroupsService.CopyHuntGroupAgentTrigger:output_type -> services.org.hunt_groups.v1alpha1.CopyHuntGroupAgentTriggerResponse
	12, // 12: services.org.hunt_groups.v1alpha1.HuntGroupsService.UpdateHuntGroupAgentTriggers:output_type -> services.org.hunt_groups.v1alpha1.UpdateHuntGroupAgentTriggersResponse
	13, // 13: services.org.hunt_groups.v1alpha1.HuntGroupsService.AdminCopyHuntGroupToOrganization:output_type -> services.org.hunt_groups.v1alpha1.AdminCopyHuntGroupToOrganizationResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_org_hunt_groups_v1alpha1_service_proto_init() }
func file_services_org_hunt_groups_v1alpha1_service_proto_init() {
	if File_services_org_hunt_groups_v1alpha1_service_proto != nil {
		return
	}
	file_services_org_hunt_groups_v1alpha1_entities_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_org_hunt_groups_v1alpha1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_org_hunt_groups_v1alpha1_service_proto_goTypes,
		DependencyIndexes: file_services_org_hunt_groups_v1alpha1_service_proto_depIdxs,
	}.Build()
	File_services_org_hunt_groups_v1alpha1_service_proto = out.File
	file_services_org_hunt_groups_v1alpha1_service_proto_rawDesc = nil
	file_services_org_hunt_groups_v1alpha1_service_proto_goTypes = nil
	file_services_org_hunt_groups_v1alpha1_service_proto_depIdxs = nil
}
