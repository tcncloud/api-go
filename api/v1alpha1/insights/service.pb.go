// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v1alpha1/insights/service.proto

package insights

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

var File_api_v1alpha1_insights_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_insights_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x1a, 0x17, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfa, 0x0d, 0x0a, 0x08, 0x49, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0xae, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x42, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xdc, 0x04, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f,
	0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x69,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0xaa, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x41, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xdc, 0x04, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x31, 0x3a, 0x01, 0x2a, 0x22, 0x2c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x12, 0xae, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x42, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xdc, 0x04, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x69, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x12, 0xae, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x42, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xdc, 0x04, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x69,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0xa2, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0xba, 0xb8, 0x91, 0x02, 0x05,
	0x0a, 0x03, 0x08, 0xdc, 0x04, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a, 0x22, 0x2a,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f,
	0x67, 0x65, 0x74, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0xbc, 0x01, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49,
	0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xdb, 0x04, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39,
	0x3a, 0x01, 0x2a, 0x22, 0x34, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0xbc, 0x01, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x69,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0xba,
	0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xdb, 0x04, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x3a,
	0x01, 0x2a, 0x22, 0x34, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0xbc, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0xba, 0xb8,
	0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xdb, 0x04, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x3a, 0x01,
	0x2a, 0x22, 0x34, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0xaa, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56,
	0x66, 0x73, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x66, 0x73, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x66, 0x73, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x41, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xdc, 0x04, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x31, 0x3a, 0x01, 0x2a, 0x22, 0x2c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x69,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x76, 0x66, 0x73, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x9e, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x66, 0x73,
	0x65, 0x73, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x66, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x66, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xdc,
	0x04, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01, 0x2a, 0x22, 0x29, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x76, 0x66, 0x73, 0x65, 0x73, 0x42, 0xd1, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x49, 0xaa, 0x02, 0x15, 0x41, 0x70,
	0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x73, 0xca, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0xe2, 0x02, 0x21, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x49, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x17, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a,
	0x3a, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_api_v1alpha1_insights_service_proto_goTypes = []interface{}{
	(*CreateInsightRequest)(nil),  // 0: api.v1alpha1.insights.CreateInsightRequest
	(*ListInsightsRequest)(nil),   // 1: api.v1alpha1.insights.ListInsightsRequest
	(*UpdateInsightRequest)(nil),  // 2: api.v1alpha1.insights.UpdateInsightRequest
	(*DeleteInsightRequest)(nil),  // 3: api.v1alpha1.insights.DeleteInsightRequest
	(*GetInsightRequest)(nil),     // 4: api.v1alpha1.insights.GetInsightRequest
	(*GetVfsSchemaRequest)(nil),   // 5: api.v1alpha1.insights.GetVfsSchemaRequest
	(*ListVfsesRequest)(nil),      // 6: api.v1alpha1.insights.ListVfsesRequest
	(*CreateInsightResponse)(nil), // 7: api.v1alpha1.insights.CreateInsightResponse
	(*ListInsightsResponse)(nil),  // 8: api.v1alpha1.insights.ListInsightsResponse
	(*UpdateInsightResponse)(nil), // 9: api.v1alpha1.insights.UpdateInsightResponse
	(*DeleteInsightResponse)(nil), // 10: api.v1alpha1.insights.DeleteInsightResponse
	(*GetInsightResponse)(nil),    // 11: api.v1alpha1.insights.GetInsightResponse
	(*GetVfsSchemaResponse)(nil),  // 12: api.v1alpha1.insights.GetVfsSchemaResponse
	(*ListVfsesResponse)(nil),     // 13: api.v1alpha1.insights.ListVfsesResponse
}
var file_api_v1alpha1_insights_service_proto_depIdxs = []int32{
	0,  // 0: api.v1alpha1.insights.Insights.CreateInsight:input_type -> api.v1alpha1.insights.CreateInsightRequest
	1,  // 1: api.v1alpha1.insights.Insights.ListInsights:input_type -> api.v1alpha1.insights.ListInsightsRequest
	2,  // 2: api.v1alpha1.insights.Insights.UpdateInsight:input_type -> api.v1alpha1.insights.UpdateInsightRequest
	3,  // 3: api.v1alpha1.insights.Insights.DeleteInsight:input_type -> api.v1alpha1.insights.DeleteInsightRequest
	4,  // 4: api.v1alpha1.insights.Insights.GetInsight:input_type -> api.v1alpha1.insights.GetInsightRequest
	0,  // 5: api.v1alpha1.insights.Insights.CreateCommonsInsight:input_type -> api.v1alpha1.insights.CreateInsightRequest
	2,  // 6: api.v1alpha1.insights.Insights.UpdateCommonsInsight:input_type -> api.v1alpha1.insights.UpdateInsightRequest
	3,  // 7: api.v1alpha1.insights.Insights.DeleteCommonsInsight:input_type -> api.v1alpha1.insights.DeleteInsightRequest
	5,  // 8: api.v1alpha1.insights.Insights.GetVfsSchema:input_type -> api.v1alpha1.insights.GetVfsSchemaRequest
	6,  // 9: api.v1alpha1.insights.Insights.ListVfses:input_type -> api.v1alpha1.insights.ListVfsesRequest
	7,  // 10: api.v1alpha1.insights.Insights.CreateInsight:output_type -> api.v1alpha1.insights.CreateInsightResponse
	8,  // 11: api.v1alpha1.insights.Insights.ListInsights:output_type -> api.v1alpha1.insights.ListInsightsResponse
	9,  // 12: api.v1alpha1.insights.Insights.UpdateInsight:output_type -> api.v1alpha1.insights.UpdateInsightResponse
	10, // 13: api.v1alpha1.insights.Insights.DeleteInsight:output_type -> api.v1alpha1.insights.DeleteInsightResponse
	11, // 14: api.v1alpha1.insights.Insights.GetInsight:output_type -> api.v1alpha1.insights.GetInsightResponse
	7,  // 15: api.v1alpha1.insights.Insights.CreateCommonsInsight:output_type -> api.v1alpha1.insights.CreateInsightResponse
	9,  // 16: api.v1alpha1.insights.Insights.UpdateCommonsInsight:output_type -> api.v1alpha1.insights.UpdateInsightResponse
	10, // 17: api.v1alpha1.insights.Insights.DeleteCommonsInsight:output_type -> api.v1alpha1.insights.DeleteInsightResponse
	12, // 18: api.v1alpha1.insights.Insights.GetVfsSchema:output_type -> api.v1alpha1.insights.GetVfsSchemaResponse
	13, // 19: api.v1alpha1.insights.Insights.ListVfses:output_type -> api.v1alpha1.insights.ListVfsesResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_insights_service_proto_init() }
func file_api_v1alpha1_insights_service_proto_init() {
	if File_api_v1alpha1_insights_service_proto != nil {
		return
	}
	file_api_v1alpha1_insights_insight_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_insights_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_insights_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_insights_service_proto_depIdxs,
	}.Build()
	File_api_v1alpha1_insights_service_proto = out.File
	file_api_v1alpha1_insights_service_proto_rawDesc = nil
	file_api_v1alpha1_insights_service_proto_goTypes = nil
	file_api_v1alpha1_insights_service_proto_depIdxs = nil
}
