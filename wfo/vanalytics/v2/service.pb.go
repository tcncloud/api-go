// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: wfo/vanalytics/v2/service.proto

package vanalyticsv2

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

var File_wfo_vanalytics_v2_service_proto protoreflect.FileDescriptor

var file_wfo_vanalytics_v2_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x77, 0x66, 0x6f,
	0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x77, 0x66, 0x6f,
	0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x66,
	0x6c, 0x61, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe8, 0x09, 0x0a, 0x0a, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x12, 0xa9, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x12, 0x2b, 0x2e, 0x77, 0x66, 0x6f, 0x2e,
	0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xf4, 0x03,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x77, 0x66, 0x6f, 0x2f,
	0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x12,
	0x87, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x26, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x34, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xf4, 0x03, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x77, 0x66, 0x6f, 0x2f, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x91, 0x01, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e, 0x77, 0x66, 0x6f, 0x2e,
	0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a,
	0x03, 0x08, 0xf4, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f,
	0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x87, 0x01,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x26,
	0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x22, 0x34, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xf4, 0x03, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x95, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0xba, 0xb8, 0x91, 0x02, 0x05,
	0x0a, 0x03, 0x08, 0xf4, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f,
	0x2f, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x7e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x77,
	0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x31, 0xba, 0xb8,
	0x91, 0x02, 0x05, 0x0a, 0x03, 0x08, 0xf4, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01,
	0x2a, 0x22, 0x1c, 0x2f, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x74, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0xc9, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x33, 0x2e,
	0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x34, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a,
	0x03, 0x08, 0xf4, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x3a, 0x01, 0x2a, 0x22, 0x2c, 0x2f,
	0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x66, 0x6c, 0x61, 0x67, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0xa1, 0x01, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x29, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x77, 0x66, 0x6f,
	0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0xba, 0xb8, 0x91, 0x02, 0x05, 0x0a, 0x03, 0x08,
	0xf4, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x77, 0x66,
	0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x66, 0x6c, 0x61, 0x67, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x42,
	0xc6, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x57, 0x56, 0x58, 0xaa, 0x02, 0x11, 0x57, 0x66, 0x6f,
	0x2e, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x56, 0x32, 0xca, 0x02,
	0x11, 0x57, 0x66, 0x6f, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c,
	0x56, 0x32, 0xe2, 0x02, 0x1d, 0x57, 0x66, 0x6f, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x13, 0x57, 0x66, 0x6f, 0x3a, 0x3a, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_wfo_vanalytics_v2_service_proto_goTypes = []interface{}{
	(*SearchTranscriptsRequest)(nil),          // 0: wfo.vanalytics.v2.SearchTranscriptsRequest
	(*CreateFilterRequest)(nil),               // 1: wfo.vanalytics.v2.CreateFilterRequest
	(*ListFiltersRequest)(nil),                // 2: wfo.vanalytics.v2.ListFiltersRequest
	(*UpdateFilterRequest)(nil),               // 3: wfo.vanalytics.v2.UpdateFilterRequest
	(*DeleteFilterRequest)(nil),               // 4: wfo.vanalytics.v2.DeleteFilterRequest
	(*GetFilterRequest)(nil),                  // 5: wfo.vanalytics.v2.GetFilterRequest
	(*ListFlagTranscriptFiltersRequest)(nil),  // 6: wfo.vanalytics.v2.ListFlagTranscriptFiltersRequest
	(*ListFlagFiltersRequest)(nil),            // 7: wfo.vanalytics.v2.ListFlagFiltersRequest
	(*SearchTranscriptsResponse)(nil),         // 8: wfo.vanalytics.v2.SearchTranscriptsResponse
	(*Filter)(nil),                            // 9: wfo.vanalytics.v2.Filter
	(*ListFiltersResponse)(nil),               // 10: wfo.vanalytics.v2.ListFiltersResponse
	(*DeleteFilterResponse)(nil),              // 11: wfo.vanalytics.v2.DeleteFilterResponse
	(*ListFlagTranscriptFiltersResponse)(nil), // 12: wfo.vanalytics.v2.ListFlagTranscriptFiltersResponse
	(*ListFlagFiltersResponse)(nil),           // 13: wfo.vanalytics.v2.ListFlagFiltersResponse
}
var file_wfo_vanalytics_v2_service_proto_depIdxs = []int32{
	0,  // 0: wfo.vanalytics.v2.Vanalytics.SearchTranscripts:input_type -> wfo.vanalytics.v2.SearchTranscriptsRequest
	1,  // 1: wfo.vanalytics.v2.Vanalytics.CreateFilter:input_type -> wfo.vanalytics.v2.CreateFilterRequest
	2,  // 2: wfo.vanalytics.v2.Vanalytics.ListFilters:input_type -> wfo.vanalytics.v2.ListFiltersRequest
	3,  // 3: wfo.vanalytics.v2.Vanalytics.UpdateFilter:input_type -> wfo.vanalytics.v2.UpdateFilterRequest
	4,  // 4: wfo.vanalytics.v2.Vanalytics.DeleteFilter:input_type -> wfo.vanalytics.v2.DeleteFilterRequest
	5,  // 5: wfo.vanalytics.v2.Vanalytics.GetFilter:input_type -> wfo.vanalytics.v2.GetFilterRequest
	6,  // 6: wfo.vanalytics.v2.Vanalytics.ListFlagTranscriptFilters:input_type -> wfo.vanalytics.v2.ListFlagTranscriptFiltersRequest
	7,  // 7: wfo.vanalytics.v2.Vanalytics.ListFlagFilters:input_type -> wfo.vanalytics.v2.ListFlagFiltersRequest
	8,  // 8: wfo.vanalytics.v2.Vanalytics.SearchTranscripts:output_type -> wfo.vanalytics.v2.SearchTranscriptsResponse
	9,  // 9: wfo.vanalytics.v2.Vanalytics.CreateFilter:output_type -> wfo.vanalytics.v2.Filter
	10, // 10: wfo.vanalytics.v2.Vanalytics.ListFilters:output_type -> wfo.vanalytics.v2.ListFiltersResponse
	9,  // 11: wfo.vanalytics.v2.Vanalytics.UpdateFilter:output_type -> wfo.vanalytics.v2.Filter
	11, // 12: wfo.vanalytics.v2.Vanalytics.DeleteFilter:output_type -> wfo.vanalytics.v2.DeleteFilterResponse
	9,  // 13: wfo.vanalytics.v2.Vanalytics.GetFilter:output_type -> wfo.vanalytics.v2.Filter
	12, // 14: wfo.vanalytics.v2.Vanalytics.ListFlagTranscriptFilters:output_type -> wfo.vanalytics.v2.ListFlagTranscriptFiltersResponse
	13, // 15: wfo.vanalytics.v2.Vanalytics.ListFlagFilters:output_type -> wfo.vanalytics.v2.ListFlagFiltersResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_wfo_vanalytics_v2_service_proto_init() }
func file_wfo_vanalytics_v2_service_proto_init() {
	if File_wfo_vanalytics_v2_service_proto != nil {
		return
	}
	file_wfo_vanalytics_v2_filter_proto_init()
	file_wfo_vanalytics_v2_flag_filter_proto_init()
	file_wfo_vanalytics_v2_flag_transcript_filter_proto_init()
	file_wfo_vanalytics_v2_transcript_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wfo_vanalytics_v2_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wfo_vanalytics_v2_service_proto_goTypes,
		DependencyIndexes: file_wfo_vanalytics_v2_service_proto_depIdxs,
	}.Build()
	File_wfo_vanalytics_v2_service_proto = out.File
	file_wfo_vanalytics_v2_service_proto_rawDesc = nil
	file_wfo_vanalytics_v2_service_proto_goTypes = nil
	file_wfo_vanalytics_v2_service_proto_depIdxs = nil
}
