// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/commons/vanalytics.proto

package commons

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Interval enumerates common dynamic time periods.
type Interval int32

const (
	Interval_TODAY         Interval = 0
	Interval_YESTERDAY     Interval = 1
	Interval_THIS_WEEK     Interval = 2
	Interval_LAST_7_DAYS   Interval = 3
	Interval_LAST_WEEK     Interval = 4
	Interval_LAST_14_DAYS  Interval = 5
	Interval_THIS_MONTH    Interval = 6
	Interval_LAST_30_DAYS  Interval = 7
	Interval_LAST_60_DAYS  Interval = 8
	Interval_LAST_90_DAYS  Interval = 9
	Interval_LAST_180_DAYS Interval = 10
)

// Enum value maps for Interval.
var (
	Interval_name = map[int32]string{
		0:  "TODAY",
		1:  "YESTERDAY",
		2:  "THIS_WEEK",
		3:  "LAST_7_DAYS",
		4:  "LAST_WEEK",
		5:  "LAST_14_DAYS",
		6:  "THIS_MONTH",
		7:  "LAST_30_DAYS",
		8:  "LAST_60_DAYS",
		9:  "LAST_90_DAYS",
		10: "LAST_180_DAYS",
	}
	Interval_value = map[string]int32{
		"TODAY":         0,
		"YESTERDAY":     1,
		"THIS_WEEK":     2,
		"LAST_7_DAYS":   3,
		"LAST_WEEK":     4,
		"LAST_14_DAYS":  5,
		"THIS_MONTH":    6,
		"LAST_30_DAYS":  7,
		"LAST_60_DAYS":  8,
		"LAST_90_DAYS":  9,
		"LAST_180_DAYS": 10,
	}
)

func (x Interval) Enum() *Interval {
	p := new(Interval)
	*p = x
	return p
}

func (x Interval) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Interval) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_vanalytics_proto_enumTypes[0].Descriptor()
}

func (Interval) Type() protoreflect.EnumType {
	return &file_api_commons_vanalytics_proto_enumTypes[0]
}

func (x Interval) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Interval.Descriptor instead.
func (Interval) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_vanalytics_proto_rawDescGZIP(), []int{0}
}

// TranscriptSummaryStatus defines a set of possible statuses
// in which transcript summaries may exist.
type TranscriptSummaryStatus int32

const (
	// Initial - Summary is currently
	// queued and awaiting generation.
	TranscriptSummaryStatus_TRANSCRIPT_SUMMARY_STATUS_QUEUED TranscriptSummaryStatus = 0
	// Summary could not be generated.
	TranscriptSummaryStatus_TRANSCRIPT_SUMMARY_STATUS_QUEUED_ERRORED TranscriptSummaryStatus = -1
	// Summary has been generated.
	// Waiting to be added to transcript.
	TranscriptSummaryStatus_TRANSCRIPT_SUMMARY_STATUS_SUMMARIZED TranscriptSummaryStatus = 1
	// Summary could not be added to transcript.
	TranscriptSummaryStatus_TRANSCRIPT_SUMMARY_STATUS_SUMMARIZED_ERRORED TranscriptSummaryStatus = -2
	// Summary is generated and visible.
	TranscriptSummaryStatus_TRANSCRIPT_SUMMARY_STATUS_VISIBLE TranscriptSummaryStatus = 2
)

// Enum value maps for TranscriptSummaryStatus.
var (
	TranscriptSummaryStatus_name = map[int32]string{
		0:  "TRANSCRIPT_SUMMARY_STATUS_QUEUED",
		-1: "TRANSCRIPT_SUMMARY_STATUS_QUEUED_ERRORED",
		1:  "TRANSCRIPT_SUMMARY_STATUS_SUMMARIZED",
		-2: "TRANSCRIPT_SUMMARY_STATUS_SUMMARIZED_ERRORED",
		2:  "TRANSCRIPT_SUMMARY_STATUS_VISIBLE",
	}
	TranscriptSummaryStatus_value = map[string]int32{
		"TRANSCRIPT_SUMMARY_STATUS_QUEUED":             0,
		"TRANSCRIPT_SUMMARY_STATUS_QUEUED_ERRORED":     -1,
		"TRANSCRIPT_SUMMARY_STATUS_SUMMARIZED":         1,
		"TRANSCRIPT_SUMMARY_STATUS_SUMMARIZED_ERRORED": -2,
		"TRANSCRIPT_SUMMARY_STATUS_VISIBLE":            2,
	}
)

func (x TranscriptSummaryStatus) Enum() *TranscriptSummaryStatus {
	p := new(TranscriptSummaryStatus)
	*p = x
	return p
}

func (x TranscriptSummaryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TranscriptSummaryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_vanalytics_proto_enumTypes[1].Descriptor()
}

func (TranscriptSummaryStatus) Type() protoreflect.EnumType {
	return &file_api_commons_vanalytics_proto_enumTypes[1]
}

func (x TranscriptSummaryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TranscriptSummaryStatus.Descriptor instead.
func (TranscriptSummaryStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_vanalytics_proto_rawDescGZIP(), []int{1}
}

var File_api_commons_vanalytics_proto protoreflect.FileDescriptor

var file_api_commons_vanalytics_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0xbe, 0x01, 0x0a, 0x08,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x4f, 0x44, 0x41,
	0x59, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x59, 0x45, 0x53, 0x54, 0x45, 0x52, 0x44, 0x41, 0x59,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x48, 0x49, 0x53, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x37, 0x5f, 0x44, 0x41, 0x59, 0x53,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x10,
	0x04, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x31, 0x34, 0x5f, 0x44, 0x41, 0x59,
	0x53, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x48, 0x49, 0x53, 0x5f, 0x4d, 0x4f, 0x4e, 0x54,
	0x48, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x33, 0x30, 0x5f, 0x44,
	0x41, 0x59, 0x53, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x36, 0x30,
	0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x41, 0x53, 0x54, 0x5f,
	0x39, 0x30, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x41, 0x53,
	0x54, 0x5f, 0x31, 0x38, 0x30, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x0a, 0x2a, 0x82, 0x02, 0x0a,
	0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x00, 0x12, 0x35,
	0x0a, 0x28, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x5f, 0x53, 0x55, 0x4d,
	0x4d, 0x41, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x51, 0x55, 0x45, 0x55,
	0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x45, 0x44, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x43, 0x52,
	0x49, 0x50, 0x54, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x39, 0x0a, 0x2c, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x5f, 0x53, 0x55,
	0x4d, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x4d,
	0x4d, 0x41, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x45, 0x44, 0x10,
	0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x59,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10,
	0x02, 0x42, 0x97, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x0f, 0x56, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_commons_vanalytics_proto_rawDescOnce sync.Once
	file_api_commons_vanalytics_proto_rawDescData = file_api_commons_vanalytics_proto_rawDesc
)

func file_api_commons_vanalytics_proto_rawDescGZIP() []byte {
	file_api_commons_vanalytics_proto_rawDescOnce.Do(func() {
		file_api_commons_vanalytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_vanalytics_proto_rawDescData)
	})
	return file_api_commons_vanalytics_proto_rawDescData
}

var file_api_commons_vanalytics_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_commons_vanalytics_proto_goTypes = []any{
	(Interval)(0),                // 0: api.commons.Interval
	(TranscriptSummaryStatus)(0), // 1: api.commons.TranscriptSummaryStatus
}
var file_api_commons_vanalytics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_vanalytics_proto_init() }
func file_api_commons_vanalytics_proto_init() {
	if File_api_commons_vanalytics_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_vanalytics_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_vanalytics_proto_goTypes,
		DependencyIndexes: file_api_commons_vanalytics_proto_depIdxs,
		EnumInfos:         file_api_commons_vanalytics_proto_enumTypes,
	}.Build()
	File_api_commons_vanalytics_proto = out.File
	file_api_commons_vanalytics_proto_rawDesc = nil
	file_api_commons_vanalytics_proto_goTypes = nil
	file_api_commons_vanalytics_proto_depIdxs = nil
}
