// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: api/v1alpha1/contactmanager/service.proto

package contactmanager

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

var File_api_v1alpha1_contactmanager_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_contactmanager_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x30, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xeb, 0x0b, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0xca, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4f, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xec, 0x27, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x3f, 0x3a, 0x01, 0x2a, 0x22, 0x3a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0xe2, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x55, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xec, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x45, 0x3a, 0x01, 0x2a, 0x22, 0x40, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xda, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x6e,
	0x63, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x36, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53,
	0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xec, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x43,
	0x3a, 0x01, 0x2a, 0x22, 0x3e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x74, 0x65, 0x6e, 0x63, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0xe6, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4b, 0x59, 0x43, 0x45, 0x6e,
	0x63, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x39, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4b,
	0x59, 0x43, 0x45, 0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x59, 0x43, 0x45, 0x6e, 0x63,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xed, 0x27,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x46, 0x3a, 0x01, 0x2a, 0x22, 0x41, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x6b, 0x79, 0x63, 0x65, 0x6e, 0x63,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0xba, 0x01, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4b, 0x59, 0x43, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x2e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x59, 0x43,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x59, 0x43,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0xba, 0xb8,
	0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xed, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3b, 0x3a, 0x01,
	0x2a, 0x22, 0x36, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x67,
	0x65, 0x74, 0x6b, 0x79, 0x63, 0x6b, 0x65, 0x79, 0x73, 0x12, 0xce, 0x01, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x33, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a,
	0x03, 0x08, 0xec, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x40, 0x3a, 0x01, 0x2a, 0x22, 0x3b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x64, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0xd2, 0x01, 0x0a, 0x10, 0x45,
	0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x64,
	0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0xba, 0xb8,
	0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xec, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x3a, 0x01,
	0x2a, 0x22, 0x3c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x65,
	0x64, 0x69, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0xf5, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x41, 0x56,
	0x43, 0xaa, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0xca,
	0x02, 0x1b, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0xe2, 0x02, 0x27,
	0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_v1alpha1_contactmanager_service_proto_goTypes = []interface{}{
	(*GetContactListRequest)(nil),         // 0: api.v1alpha1.contactmanager.GetContactListRequest
	(*ListContactEntryListRequest)(nil),   // 1: api.v1alpha1.contactmanager.ListContactEntryListRequest
	(*GetEncContactEntryRequest)(nil),     // 2: api.v1alpha1.contactmanager.GetEncContactEntryRequest
	(*GetKYCEncContactEntryRequest)(nil),  // 3: api.v1alpha1.contactmanager.GetKYCEncContactEntryRequest
	(*GetKYCKeysRequest)(nil),             // 4: api.v1alpha1.contactmanager.GetKYCKeysRequest
	(*AddContactEntryRequest)(nil),        // 5: api.v1alpha1.contactmanager.AddContactEntryRequest
	(*EditContactEntryRequest)(nil),       // 6: api.v1alpha1.contactmanager.EditContactEntryRequest
	(*GetContactListResponse)(nil),        // 7: api.v1alpha1.contactmanager.GetContactListResponse
	(*ListContactEntryListResponse)(nil),  // 8: api.v1alpha1.contactmanager.ListContactEntryListResponse
	(*GetEncContactEntryResponse)(nil),    // 9: api.v1alpha1.contactmanager.GetEncContactEntryResponse
	(*GetKYCEncContactEntryResponse)(nil), // 10: api.v1alpha1.contactmanager.GetKYCEncContactEntryResponse
	(*GetKYCKeysResponse)(nil),            // 11: api.v1alpha1.contactmanager.GetKYCKeysResponse
	(*AddContactEntryResponse)(nil),       // 12: api.v1alpha1.contactmanager.AddContactEntryResponse
	(*EditContactEntryResponse)(nil),      // 13: api.v1alpha1.contactmanager.EditContactEntryResponse
}
var file_api_v1alpha1_contactmanager_service_proto_depIdxs = []int32{
	0,  // 0: api.v1alpha1.contactmanager.ContactManager.GetContactList:input_type -> api.v1alpha1.contactmanager.GetContactListRequest
	1,  // 1: api.v1alpha1.contactmanager.ContactManager.ListContactEntryList:input_type -> api.v1alpha1.contactmanager.ListContactEntryListRequest
	2,  // 2: api.v1alpha1.contactmanager.ContactManager.GetEncContactEntry:input_type -> api.v1alpha1.contactmanager.GetEncContactEntryRequest
	3,  // 3: api.v1alpha1.contactmanager.ContactManager.GetKYCEncContactEntry:input_type -> api.v1alpha1.contactmanager.GetKYCEncContactEntryRequest
	4,  // 4: api.v1alpha1.contactmanager.ContactManager.GetKYCKeys:input_type -> api.v1alpha1.contactmanager.GetKYCKeysRequest
	5,  // 5: api.v1alpha1.contactmanager.ContactManager.AddContactEntry:input_type -> api.v1alpha1.contactmanager.AddContactEntryRequest
	6,  // 6: api.v1alpha1.contactmanager.ContactManager.EditContactEntry:input_type -> api.v1alpha1.contactmanager.EditContactEntryRequest
	7,  // 7: api.v1alpha1.contactmanager.ContactManager.GetContactList:output_type -> api.v1alpha1.contactmanager.GetContactListResponse
	8,  // 8: api.v1alpha1.contactmanager.ContactManager.ListContactEntryList:output_type -> api.v1alpha1.contactmanager.ListContactEntryListResponse
	9,  // 9: api.v1alpha1.contactmanager.ContactManager.GetEncContactEntry:output_type -> api.v1alpha1.contactmanager.GetEncContactEntryResponse
	10, // 10: api.v1alpha1.contactmanager.ContactManager.GetKYCEncContactEntry:output_type -> api.v1alpha1.contactmanager.GetKYCEncContactEntryResponse
	11, // 11: api.v1alpha1.contactmanager.ContactManager.GetKYCKeys:output_type -> api.v1alpha1.contactmanager.GetKYCKeysResponse
	12, // 12: api.v1alpha1.contactmanager.ContactManager.AddContactEntry:output_type -> api.v1alpha1.contactmanager.AddContactEntryResponse
	13, // 13: api.v1alpha1.contactmanager.ContactManager.EditContactEntry:output_type -> api.v1alpha1.contactmanager.EditContactEntryResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_contactmanager_service_proto_init() }
func file_api_v1alpha1_contactmanager_service_proto_init() {
	if File_api_v1alpha1_contactmanager_service_proto != nil {
		return
	}
	file_api_v1alpha1_contactmanager_contactmanager_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_contactmanager_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_contactmanager_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_contactmanager_service_proto_depIdxs,
	}.Build()
	File_api_v1alpha1_contactmanager_service_proto = out.File
	file_api_v1alpha1_contactmanager_service_proto_rawDesc = nil
	file_api_v1alpha1_contactmanager_service_proto_goTypes = nil
	file_api_v1alpha1_contactmanager_service_proto_depIdxs = nil
}
