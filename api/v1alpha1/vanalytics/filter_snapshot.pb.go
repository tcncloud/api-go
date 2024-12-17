// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: api/v1alpha1/vanalytics/filter_snapshot.proto

package vanalytics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Filter snapshot resource.
type FilterSnapshot struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The unique id of this filter snapshot.
	FilterSnapshotSid int64 `protobuf:"varint,1,opt,name=filter_snapshot_sid,json=filterSnapshotSid,proto3" json:"filter_snapshot_sid,omitempty"`
	// Output only. The unique id of this filter.
	FilterSid int64 `protobuf:"varint,2,opt,name=filter_sid,json=filterSid,proto3" json:"filter_sid,omitempty"`
	// Required. The name of this filter. Must be non empty and unique across all
	// filters within an organization.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The search request to be used to filter transcripts.
	SearchRequest *SearchRequest `protobuf:"bytes,5,opt,name=search_request,json=searchRequest,proto3" json:"search_request,omitempty"`
	// Output only. Assigned by the server. The timestamp when this filter was
	// created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The version of this filter.
	Version       int64 `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FilterSnapshot) Reset() {
	*x = FilterSnapshot{}
	mi := &file_api_v1alpha1_vanalytics_filter_snapshot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilterSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterSnapshot) ProtoMessage() {}

func (x *FilterSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_vanalytics_filter_snapshot_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterSnapshot.ProtoReflect.Descriptor instead.
func (*FilterSnapshot) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *FilterSnapshot) GetFilterSnapshotSid() int64 {
	if x != nil {
		return x.FilterSnapshotSid
	}
	return 0
}

func (x *FilterSnapshot) GetFilterSid() int64 {
	if x != nil {
		return x.FilterSid
	}
	return 0
}

func (x *FilterSnapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FilterSnapshot) GetSearchRequest() *SearchRequest {
	if x != nil {
		return x.SearchRequest
	}
	return nil
}

func (x *FilterSnapshot) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *FilterSnapshot) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_api_v1alpha1_vanalytics_filter_snapshot_proto protoreflect.FileDescriptor

var file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x53, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x53, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0xe4, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x42,
	0x13, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x56,
	0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0xca, 0x02, 0x17, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x56, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDescOnce sync.Once
	file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDescData = file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDesc
)

func file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDescData)
	})
	return file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDescData
}

var file_api_v1alpha1_vanalytics_filter_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1alpha1_vanalytics_filter_snapshot_proto_goTypes = []any{
	(*FilterSnapshot)(nil),        // 0: api.v1alpha1.vanalytics.FilterSnapshot
	(*SearchRequest)(nil),         // 1: api.v1alpha1.vanalytics.SearchRequest
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_api_v1alpha1_vanalytics_filter_snapshot_proto_depIdxs = []int32{
	1, // 0: api.v1alpha1.vanalytics.FilterSnapshot.search_request:type_name -> api.v1alpha1.vanalytics.SearchRequest
	2, // 1: api.v1alpha1.vanalytics.FilterSnapshot.create_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_vanalytics_filter_snapshot_proto_init() }
func file_api_v1alpha1_vanalytics_filter_snapshot_proto_init() {
	if File_api_v1alpha1_vanalytics_filter_snapshot_proto != nil {
		return
	}
	file_api_v1alpha1_vanalytics_transcript_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_vanalytics_filter_snapshot_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_vanalytics_filter_snapshot_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_vanalytics_filter_snapshot_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_vanalytics_filter_snapshot_proto = out.File
	file_api_v1alpha1_vanalytics_filter_snapshot_proto_rawDesc = nil
	file_api_v1alpha1_vanalytics_filter_snapshot_proto_goTypes = nil
	file_api_v1alpha1_vanalytics_filter_snapshot_proto_depIdxs = nil
}
