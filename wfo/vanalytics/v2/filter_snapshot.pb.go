// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: wfo/vanalytics/v2/filter_snapshot.proto

package vanalyticsv2

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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The unique id of this filter snapshot.
	FilterSnapshotSid int64 `protobuf:"varint,1,opt,name=filter_snapshot_sid,json=filterSnapshotSid,proto3" json:"filter_snapshot_sid,omitempty"`
	// Output only. The unique id of this filter.
	FilterSid int64 `protobuf:"varint,2,opt,name=filter_sid,json=filterSid,proto3" json:"filter_sid,omitempty"`
	// Required. The name of this filter. Must be non empty and unique across all
	// filters within an organization.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Assigned by the server. The timestamp when this filter was
	// created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The version of this filter.
	Version int64 `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
	// The transcript query to be used to filter transcripts.
	TranscriptQuery *TranscriptQuery `protobuf:"bytes,8,opt,name=transcript_query,json=transcriptQuery,proto3" json:"transcript_query,omitempty"`
}

func (x *FilterSnapshot) Reset() {
	*x = FilterSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wfo_vanalytics_v2_filter_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterSnapshot) ProtoMessage() {}

func (x *FilterSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_filter_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescGZIP(), []int{0}
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

func (x *FilterSnapshot) GetTranscriptQuery() *TranscriptQuery {
	if x != nil {
		return x.TranscriptQuery
	}
	return nil
}

var File_wfo_vanalytics_v2_filter_snapshot_proto protoreflect.FileDescriptor

var file_wfo_vanalytics_v2_filter_snapshot_proto_rawDesc = []byte{
	0x0a, 0x27, 0x77, 0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x77, 0x66, 0x6f, 0x2e, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x77,
	0x66, 0x6f, 0x2f, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x99, 0x02, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x53, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x53, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4d,
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0xcd, 0x01,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x66, 0x6f, 0x2e, 0x76, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x13, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x77, 0x66, 0x6f, 0x2f, 0x76,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x57, 0x56, 0x58, 0xaa,
	0x02, 0x11, 0x57, 0x66, 0x6f, 0x2e, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x56, 0x32, 0xca, 0x02, 0x11, 0x57, 0x66, 0x6f, 0x5c, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x1d, 0x57, 0x66, 0x6f, 0x5c, 0x56, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x57, 0x66, 0x6f, 0x3a, 0x3a, 0x56,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescOnce sync.Once
	file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescData = file_wfo_vanalytics_v2_filter_snapshot_proto_rawDesc
)

func file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescGZIP() []byte {
	file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescOnce.Do(func() {
		file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescData)
	})
	return file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescData
}

var file_wfo_vanalytics_v2_filter_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wfo_vanalytics_v2_filter_snapshot_proto_goTypes = []interface{}{
	(*FilterSnapshot)(nil),        // 0: wfo.vanalytics.v2.FilterSnapshot
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*TranscriptQuery)(nil),       // 2: wfo.vanalytics.v2.TranscriptQuery
}
var file_wfo_vanalytics_v2_filter_snapshot_proto_depIdxs = []int32{
	1, // 0: wfo.vanalytics.v2.FilterSnapshot.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: wfo.vanalytics.v2.FilterSnapshot.transcript_query:type_name -> wfo.vanalytics.v2.TranscriptQuery
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wfo_vanalytics_v2_filter_snapshot_proto_init() }
func file_wfo_vanalytics_v2_filter_snapshot_proto_init() {
	if File_wfo_vanalytics_v2_filter_snapshot_proto != nil {
		return
	}
	file_wfo_vanalytics_v2_transcript_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wfo_vanalytics_v2_filter_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterSnapshot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wfo_vanalytics_v2_filter_snapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wfo_vanalytics_v2_filter_snapshot_proto_goTypes,
		DependencyIndexes: file_wfo_vanalytics_v2_filter_snapshot_proto_depIdxs,
		MessageInfos:      file_wfo_vanalytics_v2_filter_snapshot_proto_msgTypes,
	}.Build()
	File_wfo_vanalytics_v2_filter_snapshot_proto = out.File
	file_wfo_vanalytics_v2_filter_snapshot_proto_rawDesc = nil
	file_wfo_vanalytics_v2_filter_snapshot_proto_goTypes = nil
	file_wfo_vanalytics_v2_filter_snapshot_proto_depIdxs = nil
}
