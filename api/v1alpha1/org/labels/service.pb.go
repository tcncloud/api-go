// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: api/v1alpha1/org/labels/service.proto

package labels

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

var File_api_v1alpha1_org_labels_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_org_labels_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f,
	0x72, 0x67, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xa0, 0x09, 0x0a, 0x0d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xa3, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xba, 0xb8,
	0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0x96, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01,
	0x2a, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x94, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0xba, 0xb8, 0x91, 0x02, 0x02,
	0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0xa3,
	0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xba, 0xb8, 0x91, 0x02, 0x05,
	0x0a, 0x03, 0x08, 0x96, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x9c, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0xba, 0xb8,
	0x91, 0x02, 0x02, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0xa3, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39,
	0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0x96, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29,
	0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0xa3, 0x01, 0x0a, 0x0b, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0x97, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0xc0, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x12, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65,
	0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3e, 0xba, 0xb8, 0x91, 0x02, 0x02, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x31, 0x3a, 0x01, 0x2a, 0x22, 0x2c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x67,
	0x65, 0x74, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x6d,
	0x61, 0x70, 0x42, 0xdf, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0xa2, 0x02, 0x04, 0x41, 0x56, 0x4f, 0x4c, 0xaa, 0x02, 0x17,
	0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x72, 0x67,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0xca, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_v1alpha1_org_labels_service_proto_goTypes = []interface{}{
	(*CreateLabelRequest)(nil),          // 0: api.v1alpha1.org.labels.CreateLabelRequest
	(*GetLabelRequest)(nil),             // 1: api.v1alpha1.org.labels.GetLabelRequest
	(*UpdateLabelRequest)(nil),          // 2: api.v1alpha1.org.labels.UpdateLabelRequest
	(*ListLabelsRequest)(nil),           // 3: api.v1alpha1.org.labels.ListLabelsRequest
	(*DeleteLabelRequest)(nil),          // 4: api.v1alpha1.org.labels.DeleteLabelRequest
	(*AttachLabelRequest)(nil),          // 5: api.v1alpha1.org.labels.AttachLabelRequest
	(*GetLabeledEntityMapRequest)(nil),  // 6: api.v1alpha1.org.labels.GetLabeledEntityMapRequest
	(*CreateLabelResponse)(nil),         // 7: api.v1alpha1.org.labels.CreateLabelResponse
	(*GetLabelResponse)(nil),            // 8: api.v1alpha1.org.labels.GetLabelResponse
	(*UpdateLabelResponse)(nil),         // 9: api.v1alpha1.org.labels.UpdateLabelResponse
	(*ListLabelsResponse)(nil),          // 10: api.v1alpha1.org.labels.ListLabelsResponse
	(*DeleteLabelResponse)(nil),         // 11: api.v1alpha1.org.labels.DeleteLabelResponse
	(*AttachLabelResponse)(nil),         // 12: api.v1alpha1.org.labels.AttachLabelResponse
	(*GetLabeledEntityMapResponse)(nil), // 13: api.v1alpha1.org.labels.GetLabeledEntityMapResponse
}
var file_api_v1alpha1_org_labels_service_proto_depIdxs = []int32{
	0,  // 0: api.v1alpha1.org.labels.LabelsService.CreateLabel:input_type -> api.v1alpha1.org.labels.CreateLabelRequest
	1,  // 1: api.v1alpha1.org.labels.LabelsService.GetLabel:input_type -> api.v1alpha1.org.labels.GetLabelRequest
	2,  // 2: api.v1alpha1.org.labels.LabelsService.UpdateLabel:input_type -> api.v1alpha1.org.labels.UpdateLabelRequest
	3,  // 3: api.v1alpha1.org.labels.LabelsService.ListLabels:input_type -> api.v1alpha1.org.labels.ListLabelsRequest
	4,  // 4: api.v1alpha1.org.labels.LabelsService.DeleteLabel:input_type -> api.v1alpha1.org.labels.DeleteLabelRequest
	5,  // 5: api.v1alpha1.org.labels.LabelsService.AttachLabel:input_type -> api.v1alpha1.org.labels.AttachLabelRequest
	6,  // 6: api.v1alpha1.org.labels.LabelsService.GetLabeledEntityMap:input_type -> api.v1alpha1.org.labels.GetLabeledEntityMapRequest
	7,  // 7: api.v1alpha1.org.labels.LabelsService.CreateLabel:output_type -> api.v1alpha1.org.labels.CreateLabelResponse
	8,  // 8: api.v1alpha1.org.labels.LabelsService.GetLabel:output_type -> api.v1alpha1.org.labels.GetLabelResponse
	9,  // 9: api.v1alpha1.org.labels.LabelsService.UpdateLabel:output_type -> api.v1alpha1.org.labels.UpdateLabelResponse
	10, // 10: api.v1alpha1.org.labels.LabelsService.ListLabels:output_type -> api.v1alpha1.org.labels.ListLabelsResponse
	11, // 11: api.v1alpha1.org.labels.LabelsService.DeleteLabel:output_type -> api.v1alpha1.org.labels.DeleteLabelResponse
	12, // 12: api.v1alpha1.org.labels.LabelsService.AttachLabel:output_type -> api.v1alpha1.org.labels.AttachLabelResponse
	13, // 13: api.v1alpha1.org.labels.LabelsService.GetLabeledEntityMap:output_type -> api.v1alpha1.org.labels.GetLabeledEntityMapResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_org_labels_service_proto_init() }
func file_api_v1alpha1_org_labels_service_proto_init() {
	if File_api_v1alpha1_org_labels_service_proto != nil {
		return
	}
	file_api_v1alpha1_org_labels_entities_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_org_labels_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_org_labels_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_org_labels_service_proto_depIdxs,
	}.Build()
	File_api_v1alpha1_org_labels_service_proto = out.File
	file_api_v1alpha1_org_labels_service_proto_rawDesc = nil
	file_api_v1alpha1_org_labels_service_proto_goTypes = nil
	file_api_v1alpha1_org_labels_service_proto_depIdxs = nil
}