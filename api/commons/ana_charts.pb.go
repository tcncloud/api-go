// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/commons/ana_charts.proto

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

type BarChartOrientation int32

const (
	BarChartOrientation_BAR_CHART_ORIENTATION_BAR    BarChartOrientation = 0
	BarChartOrientation_BAR_CHART_ORIENTATION_COLUMN BarChartOrientation = 1
)

// Enum value maps for BarChartOrientation.
var (
	BarChartOrientation_name = map[int32]string{
		0: "BAR_CHART_ORIENTATION_BAR",
		1: "BAR_CHART_ORIENTATION_COLUMN",
	}
	BarChartOrientation_value = map[string]int32{
		"BAR_CHART_ORIENTATION_BAR":    0,
		"BAR_CHART_ORIENTATION_COLUMN": 1,
	}
)

func (x BarChartOrientation) Enum() *BarChartOrientation {
	p := new(BarChartOrientation)
	*p = x
	return p
}

func (x BarChartOrientation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BarChartOrientation) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_ana_charts_proto_enumTypes[0].Descriptor()
}

func (BarChartOrientation) Type() protoreflect.EnumType {
	return &file_api_commons_ana_charts_proto_enumTypes[0]
}

func (x BarChartOrientation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BarChartOrientation.Descriptor instead.
func (BarChartOrientation) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_ana_charts_proto_rawDescGZIP(), []int{0}
}

type ChartOrientation int32

const (
	ChartOrientation_CHART_ORIENTATION_HORIZONTAL ChartOrientation = 0
	ChartOrientation_CHART_ORIENTATION_VERTICAL   ChartOrientation = 1
)

// Enum value maps for ChartOrientation.
var (
	ChartOrientation_name = map[int32]string{
		0: "CHART_ORIENTATION_HORIZONTAL",
		1: "CHART_ORIENTATION_VERTICAL",
	}
	ChartOrientation_value = map[string]int32{
		"CHART_ORIENTATION_HORIZONTAL": 0,
		"CHART_ORIENTATION_VERTICAL":   1,
	}
)

func (x ChartOrientation) Enum() *ChartOrientation {
	p := new(ChartOrientation)
	*p = x
	return p
}

func (x ChartOrientation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChartOrientation) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_ana_charts_proto_enumTypes[1].Descriptor()
}

func (ChartOrientation) Type() protoreflect.EnumType {
	return &file_api_commons_ana_charts_proto_enumTypes[1]
}

func (x ChartOrientation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChartOrientation.Descriptor instead.
func (ChartOrientation) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_ana_charts_proto_rawDescGZIP(), []int{1}
}

type AreaChartChoice int32

const (
	AreaChartChoice_AREA_CHART_CHOICE_AREA       AreaChartChoice = 0
	AreaChartChoice_AREA_CHART_CHOICE_AREASPLINE AreaChartChoice = 1
)

// Enum value maps for AreaChartChoice.
var (
	AreaChartChoice_name = map[int32]string{
		0: "AREA_CHART_CHOICE_AREA",
		1: "AREA_CHART_CHOICE_AREASPLINE",
	}
	AreaChartChoice_value = map[string]int32{
		"AREA_CHART_CHOICE_AREA":       0,
		"AREA_CHART_CHOICE_AREASPLINE": 1,
	}
)

func (x AreaChartChoice) Enum() *AreaChartChoice {
	p := new(AreaChartChoice)
	*p = x
	return p
}

func (x AreaChartChoice) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AreaChartChoice) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_ana_charts_proto_enumTypes[2].Descriptor()
}

func (AreaChartChoice) Type() protoreflect.EnumType {
	return &file_api_commons_ana_charts_proto_enumTypes[2]
}

func (x AreaChartChoice) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AreaChartChoice.Descriptor instead.
func (AreaChartChoice) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_ana_charts_proto_rawDescGZIP(), []int{2}
}

type LineChartStep int32

const (
	LineChartStep_LINE_CHART_STEP_LEFT   LineChartStep = 0
	LineChartStep_LINE_CHART_STEP_CENTER LineChartStep = 1
	LineChartStep_LINE_CHART_STEP_RIGHT  LineChartStep = 2
	LineChartStep_LINE_CHART_STEP_NOLINE LineChartStep = 3
)

// Enum value maps for LineChartStep.
var (
	LineChartStep_name = map[int32]string{
		0: "LINE_CHART_STEP_LEFT",
		1: "LINE_CHART_STEP_CENTER",
		2: "LINE_CHART_STEP_RIGHT",
		3: "LINE_CHART_STEP_NOLINE",
	}
	LineChartStep_value = map[string]int32{
		"LINE_CHART_STEP_LEFT":   0,
		"LINE_CHART_STEP_CENTER": 1,
		"LINE_CHART_STEP_RIGHT":  2,
		"LINE_CHART_STEP_NOLINE": 3,
	}
)

func (x LineChartStep) Enum() *LineChartStep {
	p := new(LineChartStep)
	*p = x
	return p
}

func (x LineChartStep) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LineChartStep) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_ana_charts_proto_enumTypes[3].Descriptor()
}

func (LineChartStep) Type() protoreflect.EnumType {
	return &file_api_commons_ana_charts_proto_enumTypes[3]
}

func (x LineChartStep) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LineChartStep.Descriptor instead.
func (LineChartStep) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_ana_charts_proto_rawDescGZIP(), []int{3}
}

type ChartDisplayLabels int32

const (
	ChartDisplayLabels_CHART_DISPLAY_LABELS_NEVER  ChartDisplayLabels = 0
	ChartDisplayLabels_CHART_DISPLAY_LABELS_ALWAYS ChartDisplayLabels = 1
)

// Enum value maps for ChartDisplayLabels.
var (
	ChartDisplayLabels_name = map[int32]string{
		0: "CHART_DISPLAY_LABELS_NEVER",
		1: "CHART_DISPLAY_LABELS_ALWAYS",
	}
	ChartDisplayLabels_value = map[string]int32{
		"CHART_DISPLAY_LABELS_NEVER":  0,
		"CHART_DISPLAY_LABELS_ALWAYS": 1,
	}
)

func (x ChartDisplayLabels) Enum() *ChartDisplayLabels {
	p := new(ChartDisplayLabels)
	*p = x
	return p
}

func (x ChartDisplayLabels) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChartDisplayLabels) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_ana_charts_proto_enumTypes[4].Descriptor()
}

func (ChartDisplayLabels) Type() protoreflect.EnumType {
	return &file_api_commons_ana_charts_proto_enumTypes[4]
}

func (x ChartDisplayLabels) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChartDisplayLabels.Descriptor instead.
func (ChartDisplayLabels) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_ana_charts_proto_rawDescGZIP(), []int{4}
}

type ThresholdType int32

const (
	ThresholdType_THRESHOLD_TYPE_STATIC     ThresholdType = 0
	ThresholdType_THRESHOLD_TYPE_DATA_POINT ThresholdType = 1
	ThresholdType_THRESHOLD_TYPE_NOT_SET    ThresholdType = 2
)

// Enum value maps for ThresholdType.
var (
	ThresholdType_name = map[int32]string{
		0: "THRESHOLD_TYPE_STATIC",
		1: "THRESHOLD_TYPE_DATA_POINT",
		2: "THRESHOLD_TYPE_NOT_SET",
	}
	ThresholdType_value = map[string]int32{
		"THRESHOLD_TYPE_STATIC":     0,
		"THRESHOLD_TYPE_DATA_POINT": 1,
		"THRESHOLD_TYPE_NOT_SET":    2,
	}
)

func (x ThresholdType) Enum() *ThresholdType {
	p := new(ThresholdType)
	*p = x
	return p
}

func (x ThresholdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ThresholdType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_ana_charts_proto_enumTypes[5].Descriptor()
}

func (ThresholdType) Type() protoreflect.EnumType {
	return &file_api_commons_ana_charts_proto_enumTypes[5]
}

func (x ThresholdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ThresholdType.Descriptor instead.
func (ThresholdType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_ana_charts_proto_rawDescGZIP(), []int{5}
}

type PackedBubbleChoice int32

const (
	PackedBubbleChoice_PACKED_BUBBLE_CHOICE_PACKED PackedBubbleChoice = 0
	PackedBubbleChoice_PACKED_BUBBLE_CHOICE_SPLIT  PackedBubbleChoice = 1
)

// Enum value maps for PackedBubbleChoice.
var (
	PackedBubbleChoice_name = map[int32]string{
		0: "PACKED_BUBBLE_CHOICE_PACKED",
		1: "PACKED_BUBBLE_CHOICE_SPLIT",
	}
	PackedBubbleChoice_value = map[string]int32{
		"PACKED_BUBBLE_CHOICE_PACKED": 0,
		"PACKED_BUBBLE_CHOICE_SPLIT":  1,
	}
)

func (x PackedBubbleChoice) Enum() *PackedBubbleChoice {
	p := new(PackedBubbleChoice)
	*p = x
	return p
}

func (x PackedBubbleChoice) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackedBubbleChoice) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_ana_charts_proto_enumTypes[6].Descriptor()
}

func (PackedBubbleChoice) Type() protoreflect.EnumType {
	return &file_api_commons_ana_charts_proto_enumTypes[6]
}

func (x PackedBubbleChoice) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackedBubbleChoice.Descriptor instead.
func (PackedBubbleChoice) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_ana_charts_proto_rawDescGZIP(), []int{6}
}

type SuffixChoices int32

const (
	SuffixChoices_SUFFIX_CHOICES_NOSUFFIX   SuffixChoices = 0
	SuffixChoices_SUFFIX_CHOICES_DOLLARS    SuffixChoices = 1
	SuffixChoices_SUFFIX_CHOICES_PERCENTAGE SuffixChoices = 2
)

// Enum value maps for SuffixChoices.
var (
	SuffixChoices_name = map[int32]string{
		0: "SUFFIX_CHOICES_NOSUFFIX",
		1: "SUFFIX_CHOICES_DOLLARS",
		2: "SUFFIX_CHOICES_PERCENTAGE",
	}
	SuffixChoices_value = map[string]int32{
		"SUFFIX_CHOICES_NOSUFFIX":   0,
		"SUFFIX_CHOICES_DOLLARS":    1,
		"SUFFIX_CHOICES_PERCENTAGE": 2,
	}
)

func (x SuffixChoices) Enum() *SuffixChoices {
	p := new(SuffixChoices)
	*p = x
	return p
}

func (x SuffixChoices) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SuffixChoices) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_ana_charts_proto_enumTypes[7].Descriptor()
}

func (SuffixChoices) Type() protoreflect.EnumType {
	return &file_api_commons_ana_charts_proto_enumTypes[7]
}

func (x SuffixChoices) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SuffixChoices.Descriptor instead.
func (SuffixChoices) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_ana_charts_proto_rawDescGZIP(), []int{7}
}

var File_api_commons_ana_charts_proto protoreflect.FileDescriptor

var file_api_commons_ana_charts_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x61, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0x56, 0x0a, 0x13, 0x42,
	0x61, 0x72, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x41, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x5f,
	0x4f, 0x52, 0x49, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x41, 0x52, 0x10,
	0x00, 0x12, 0x20, 0x0a, 0x1c, 0x42, 0x41, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x5f, 0x4f,
	0x52, 0x49, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4c, 0x55, 0x4d,
	0x4e, 0x10, 0x01, 0x2a, 0x54, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x72, 0x69, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x48, 0x41, 0x52, 0x54,
	0x5f, 0x4f, 0x52, 0x49, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x48, 0x4f, 0x52,
	0x49, 0x5a, 0x4f, 0x4e, 0x54, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x48, 0x41,
	0x52, 0x54, 0x5f, 0x4f, 0x52, 0x49, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56,
	0x45, 0x52, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x2a, 0x4f, 0x0a, 0x0f, 0x41, 0x72, 0x65,
	0x61, 0x43, 0x68, 0x61, 0x72, 0x74, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x16,
	0x41, 0x52, 0x45, 0x41, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x48, 0x4f, 0x49, 0x43,
	0x45, 0x5f, 0x41, 0x52, 0x45, 0x41, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x52, 0x45, 0x41,
	0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x52,
	0x45, 0x41, 0x53, 0x50, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x2a, 0x7c, 0x0a, 0x0d, 0x4c, 0x69,
	0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x53, 0x74, 0x65, 0x70, 0x12, 0x18, 0x0a, 0x14, 0x4c,
	0x49, 0x4e, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x4c,
	0x45, 0x46, 0x54, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x43, 0x48,
	0x41, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x43, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x10,
	0x01, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x5f,
	0x53, 0x54, 0x45, 0x50, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16,
	0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f,
	0x4e, 0x4f, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x03, 0x2a, 0x55, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1e,
	0x0a, 0x1a, 0x43, 0x48, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f,
	0x4c, 0x41, 0x42, 0x45, 0x4c, 0x53, 0x5f, 0x4e, 0x45, 0x56, 0x45, 0x52, 0x10, 0x00, 0x12, 0x1f,
	0x0a, 0x1b, 0x43, 0x48, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f,
	0x4c, 0x41, 0x42, 0x45, 0x4c, 0x53, 0x5f, 0x41, 0x4c, 0x57, 0x41, 0x59, 0x53, 0x10, 0x01, 0x2a,
	0x65, 0x0a, 0x0d, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x19, 0x0a, 0x15, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x43, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x54,
	0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x48,
	0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x53, 0x45, 0x54, 0x10, 0x02, 0x2a, 0x55, 0x0a, 0x12, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x64,
	0x42, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x1b,
	0x50, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x42, 0x55, 0x42, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x48,
	0x4f, 0x49, 0x43, 0x45, 0x5f, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a,
	0x1a, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x42, 0x55, 0x42, 0x42, 0x4c, 0x45, 0x5f, 0x43,
	0x48, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x50, 0x4c, 0x49, 0x54, 0x10, 0x01, 0x2a, 0x67, 0x0a,
	0x0d, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1b,
	0x0a, 0x17, 0x53, 0x55, 0x46, 0x46, 0x49, 0x58, 0x5f, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x53,
	0x5f, 0x4e, 0x4f, 0x53, 0x55, 0x46, 0x46, 0x49, 0x58, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x53,
	0x55, 0x46, 0x46, 0x49, 0x58, 0x5f, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x53, 0x5f, 0x44, 0x4f,
	0x4c, 0x4c, 0x41, 0x52, 0x53, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x55, 0x46, 0x46, 0x49,
	0x58, 0x5f, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x53, 0x5f, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e,
	0x54, 0x41, 0x47, 0x45, 0x10, 0x02, 0x42, 0x96, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x0e, 0x41, 0x6e, 0x61, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_ana_charts_proto_rawDescOnce sync.Once
	file_api_commons_ana_charts_proto_rawDescData = file_api_commons_ana_charts_proto_rawDesc
)

func file_api_commons_ana_charts_proto_rawDescGZIP() []byte {
	file_api_commons_ana_charts_proto_rawDescOnce.Do(func() {
		file_api_commons_ana_charts_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_ana_charts_proto_rawDescData)
	})
	return file_api_commons_ana_charts_proto_rawDescData
}

var file_api_commons_ana_charts_proto_enumTypes = make([]protoimpl.EnumInfo, 8)
var file_api_commons_ana_charts_proto_goTypes = []interface{}{
	(BarChartOrientation)(0), // 0: api.commons.BarChartOrientation
	(ChartOrientation)(0),    // 1: api.commons.ChartOrientation
	(AreaChartChoice)(0),     // 2: api.commons.AreaChartChoice
	(LineChartStep)(0),       // 3: api.commons.LineChartStep
	(ChartDisplayLabels)(0),  // 4: api.commons.ChartDisplayLabels
	(ThresholdType)(0),       // 5: api.commons.ThresholdType
	(PackedBubbleChoice)(0),  // 6: api.commons.PackedBubbleChoice
	(SuffixChoices)(0),       // 7: api.commons.SuffixChoices
}
var file_api_commons_ana_charts_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_ana_charts_proto_init() }
func file_api_commons_ana_charts_proto_init() {
	if File_api_commons_ana_charts_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_ana_charts_proto_rawDesc,
			NumEnums:      8,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_ana_charts_proto_goTypes,
		DependencyIndexes: file_api_commons_ana_charts_proto_depIdxs,
		EnumInfos:         file_api_commons_ana_charts_proto_enumTypes,
	}.Build()
	File_api_commons_ana_charts_proto = out.File
	file_api_commons_ana_charts_proto_rawDesc = nil
	file_api_commons_ana_charts_proto_goTypes = nil
	file_api_commons_ana_charts_proto_depIdxs = nil
}
