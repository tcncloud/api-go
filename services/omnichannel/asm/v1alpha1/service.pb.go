// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: services/omnichannel/asm/v1alpha1/service.proto

package asmv1alpha1

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

var File_services_omnichannel_asm_v1alpha1_service_proto protoreflect.FileDescriptor

var file_services_omnichannel_asm_v1alpha1_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2f, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x99, 0x0b,
	0x0a, 0x0a, 0x41, 0x73, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc9, 0x01, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x37,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61,
	0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x45, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xac, 0x02, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x35, 0x3a, 0x01, 0x2a, 0x22, 0x30, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x73,
	0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0xbd, 0x01, 0x0a, 0x0a, 0x45, 0x6e, 0x64,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61,
	0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xac, 0x02,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2f, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e,
	0x64, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0xd9, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0xba, 0xb8, 0x91, 0x02, 0x05,
	0x0a, 0x03, 0x08, 0xac, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x3a, 0x01, 0x2a, 0x22, 0x34,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0xc1, 0x01, 0x0a, 0x0b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x12, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x43, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xac, 0x02, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x33, 0x3a, 0x01, 0x2a, 0x22, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f,
	0x61, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0xc5, 0x01, 0x0a, 0x0c, 0x44, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e,
	0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0xba, 0xb8, 0x91, 0x02,
	0x05, 0x0a, 0x03, 0x08, 0xac, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x3a, 0x01, 0x2a, 0x22,
	0x2f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x12, 0xd7, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61,
	0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x73, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xb0, 0x09,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x3a, 0x01, 0x2a, 0x22, 0x2f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2f, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xbc, 0x01, 0x0a, 0x0a, 0x50,
	0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x34, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x34, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xb0,
	0x09, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2f, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70,
	0x75, 0x73, 0x68, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0xa6, 0x02, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x61, 0x73, 0x6d, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x53, 0x4f, 0x41, 0xaa, 0x02, 0x21, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x41, 0x73, 0x6d, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x21, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x5c, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5c, 0x41,
	0x73, 0x6d, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5c, 0x41, 0x73, 0x6d, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x3a, 0x3a, 0x41, 0x73, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_services_omnichannel_asm_v1alpha1_service_proto_goTypes = []interface{}{
	(*CreateSessionRequest)(nil),       // 0: services.omnichannel.asm.v1alpha1.CreateSessionRequest
	(*EndSessionRequest)(nil),          // 1: services.omnichannel.asm.v1alpha1.EndSessionRequest
	(*GetCurrentSessionRequest)(nil),   // 2: services.omnichannel.asm.v1alpha1.GetCurrentSessionRequest
	(*EnableVoiceRequest)(nil),         // 3: services.omnichannel.asm.v1alpha1.EnableVoiceRequest
	(*DisableVoiceRequest)(nil),        // 4: services.omnichannel.asm.v1alpha1.DisableVoiceRequest
	(*ListAsmUserDetailsRequest)(nil),  // 5: services.omnichannel.asm.v1alpha1.ListAsmUserDetailsRequest
	(*PushEventsRequest)(nil),          // 6: services.omnichannel.asm.v1alpha1.PushEventsRequest
	(*CreateSessionResponse)(nil),      // 7: services.omnichannel.asm.v1alpha1.CreateSessionResponse
	(*EndSessionResponse)(nil),         // 8: services.omnichannel.asm.v1alpha1.EndSessionResponse
	(*GetCurrentSessionResponse)(nil),  // 9: services.omnichannel.asm.v1alpha1.GetCurrentSessionResponse
	(*EnableVoiceResponse)(nil),        // 10: services.omnichannel.asm.v1alpha1.EnableVoiceResponse
	(*DisableVoiceResponse)(nil),       // 11: services.omnichannel.asm.v1alpha1.DisableVoiceResponse
	(*ListAsmUserDetailsResponse)(nil), // 12: services.omnichannel.asm.v1alpha1.ListAsmUserDetailsResponse
	(*PushEventResponse)(nil),          // 13: services.omnichannel.asm.v1alpha1.PushEventResponse
}
var file_services_omnichannel_asm_v1alpha1_service_proto_depIdxs = []int32{
	0,  // 0: services.omnichannel.asm.v1alpha1.AsmService.CreateSession:input_type -> services.omnichannel.asm.v1alpha1.CreateSessionRequest
	1,  // 1: services.omnichannel.asm.v1alpha1.AsmService.EndSession:input_type -> services.omnichannel.asm.v1alpha1.EndSessionRequest
	2,  // 2: services.omnichannel.asm.v1alpha1.AsmService.GetCurrentSession:input_type -> services.omnichannel.asm.v1alpha1.GetCurrentSessionRequest
	3,  // 3: services.omnichannel.asm.v1alpha1.AsmService.EnableVoice:input_type -> services.omnichannel.asm.v1alpha1.EnableVoiceRequest
	4,  // 4: services.omnichannel.asm.v1alpha1.AsmService.DisableVoice:input_type -> services.omnichannel.asm.v1alpha1.DisableVoiceRequest
	5,  // 5: services.omnichannel.asm.v1alpha1.AsmService.ListAsmUserDetails:input_type -> services.omnichannel.asm.v1alpha1.ListAsmUserDetailsRequest
	6,  // 6: services.omnichannel.asm.v1alpha1.AsmService.PushEvents:input_type -> services.omnichannel.asm.v1alpha1.PushEventsRequest
	7,  // 7: services.omnichannel.asm.v1alpha1.AsmService.CreateSession:output_type -> services.omnichannel.asm.v1alpha1.CreateSessionResponse
	8,  // 8: services.omnichannel.asm.v1alpha1.AsmService.EndSession:output_type -> services.omnichannel.asm.v1alpha1.EndSessionResponse
	9,  // 9: services.omnichannel.asm.v1alpha1.AsmService.GetCurrentSession:output_type -> services.omnichannel.asm.v1alpha1.GetCurrentSessionResponse
	10, // 10: services.omnichannel.asm.v1alpha1.AsmService.EnableVoice:output_type -> services.omnichannel.asm.v1alpha1.EnableVoiceResponse
	11, // 11: services.omnichannel.asm.v1alpha1.AsmService.DisableVoice:output_type -> services.omnichannel.asm.v1alpha1.DisableVoiceResponse
	12, // 12: services.omnichannel.asm.v1alpha1.AsmService.ListAsmUserDetails:output_type -> services.omnichannel.asm.v1alpha1.ListAsmUserDetailsResponse
	13, // 13: services.omnichannel.asm.v1alpha1.AsmService.PushEvents:output_type -> services.omnichannel.asm.v1alpha1.PushEventResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_omnichannel_asm_v1alpha1_service_proto_init() }
func file_services_omnichannel_asm_v1alpha1_service_proto_init() {
	if File_services_omnichannel_asm_v1alpha1_service_proto != nil {
		return
	}
	file_services_omnichannel_asm_v1alpha1_entities_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_omnichannel_asm_v1alpha1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_omnichannel_asm_v1alpha1_service_proto_goTypes,
		DependencyIndexes: file_services_omnichannel_asm_v1alpha1_service_proto_depIdxs,
	}.Build()
	File_services_omnichannel_asm_v1alpha1_service_proto = out.File
	file_services_omnichannel_asm_v1alpha1_service_proto_rawDesc = nil
	file_services_omnichannel_asm_v1alpha1_service_proto_goTypes = nil
	file_services_omnichannel_asm_v1alpha1_service_proto_depIdxs = nil
}
