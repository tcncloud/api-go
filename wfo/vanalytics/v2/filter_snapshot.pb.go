// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: wfo/vanalytics/v2/filter_snapshot.proto

package vanalyticsv2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
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
	// Output only. Assigned by the server. The timestamp when this filter was
	// created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The version of this filter.
	Version int64 `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
	// The transcript query to be used to filter transcripts.
	TranscriptQuery *TranscriptQuery `protobuf:"bytes,8,opt,name=transcript_query,json=transcriptQuery,proto3" json:"transcript_query,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *FilterSnapshot) Reset() {
	*x = FilterSnapshot{}
	mi := &file_wfo_vanalytics_v2_filter_snapshot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilterSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterSnapshot) ProtoMessage() {}

func (x *FilterSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_wfo_vanalytics_v2_filter_snapshot_proto_msgTypes[0]
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

const file_wfo_vanalytics_v2_filter_snapshot_proto_rawDesc = "" +
	"\n" +
	"'wfo/vanalytics/v2/filter_snapshot.proto\x12\x11wfo.vanalytics.v2\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\"wfo/vanalytics/v2/transcript.proto\"\x99\x02\n" +
	"\x0eFilterSnapshot\x12.\n" +
	"\x13filter_snapshot_sid\x18\x01 \x01(\x03R\x11filterSnapshotSid\x12\x1d\n" +
	"\n" +
	"filter_sid\x18\x02 \x01(\x03R\tfilterSid\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12;\n" +
	"\vcreate_time\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x12\x18\n" +
	"\aversion\x18\a \x01(\x03R\aversion\x12M\n" +
	"\x10transcript_query\x18\b \x01(\v2\".wfo.vanalytics.v2.TranscriptQueryR\x0ftranscriptQueryB\xcd\x01\n" +
	"\x15com.wfo.vanalytics.v2B\x13FilterSnapshotProtoP\x01Z9github.com/tcncloud/api-go/wfo/vanalytics/v2;vanalyticsv2\xa2\x02\x03WVX\xaa\x02\x11Wfo.Vanalytics.V2\xca\x02\x11Wfo\\Vanalytics\\V2\xe2\x02\x1dWfo\\Vanalytics\\V2\\GPBMetadata\xea\x02\x13Wfo::Vanalytics::V2b\x06proto3"

var (
	file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescOnce sync.Once
	file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescData []byte
)

func file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescGZIP() []byte {
	file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescOnce.Do(func() {
		file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_wfo_vanalytics_v2_filter_snapshot_proto_rawDesc), len(file_wfo_vanalytics_v2_filter_snapshot_proto_rawDesc)))
	})
	return file_wfo_vanalytics_v2_filter_snapshot_proto_rawDescData
}

var file_wfo_vanalytics_v2_filter_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wfo_vanalytics_v2_filter_snapshot_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_wfo_vanalytics_v2_filter_snapshot_proto_rawDesc), len(file_wfo_vanalytics_v2_filter_snapshot_proto_rawDesc)),
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
	file_wfo_vanalytics_v2_filter_snapshot_proto_goTypes = nil
	file_wfo_vanalytics_v2_filter_snapshot_proto_depIdxs = nil
}
