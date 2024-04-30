// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: api/commons/org_preferences.proto

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

type BroadcastTemplateOrdering int32

const (
	BroadcastTemplateOrdering_BY_NAME_ASC             BroadcastTemplateOrdering = 0
	BroadcastTemplateOrdering_BY_NAME_DESC            BroadcastTemplateOrdering = 1
	BroadcastTemplateOrdering_BY_TEMPLATE_NUMBER_ASC  BroadcastTemplateOrdering = 2
	BroadcastTemplateOrdering_BY_TEMPLATE_NUMBER_DESC BroadcastTemplateOrdering = 3
	BroadcastTemplateOrdering_BY_MODIFIED_DATE_ASC    BroadcastTemplateOrdering = 4
	BroadcastTemplateOrdering_BY_MODIFIED_DATE_DESC   BroadcastTemplateOrdering = 5
)

// Enum value maps for BroadcastTemplateOrdering.
var (
	BroadcastTemplateOrdering_name = map[int32]string{
		0: "BY_NAME_ASC",
		1: "BY_NAME_DESC",
		2: "BY_TEMPLATE_NUMBER_ASC",
		3: "BY_TEMPLATE_NUMBER_DESC",
		4: "BY_MODIFIED_DATE_ASC",
		5: "BY_MODIFIED_DATE_DESC",
	}
	BroadcastTemplateOrdering_value = map[string]int32{
		"BY_NAME_ASC":             0,
		"BY_NAME_DESC":            1,
		"BY_TEMPLATE_NUMBER_ASC":  2,
		"BY_TEMPLATE_NUMBER_DESC": 3,
		"BY_MODIFIED_DATE_ASC":    4,
		"BY_MODIFIED_DATE_DESC":   5,
	}
)

func (x BroadcastTemplateOrdering) Enum() *BroadcastTemplateOrdering {
	p := new(BroadcastTemplateOrdering)
	*p = x
	return p
}

func (x BroadcastTemplateOrdering) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BroadcastTemplateOrdering) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_org_preferences_proto_enumTypes[0].Descriptor()
}

func (BroadcastTemplateOrdering) Type() protoreflect.EnumType {
	return &file_api_commons_org_preferences_proto_enumTypes[0]
}

func (x BroadcastTemplateOrdering) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BroadcastTemplateOrdering.Descriptor instead.
func (BroadcastTemplateOrdering) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_org_preferences_proto_rawDescGZIP(), []int{0}
}

type ScheduleByTimeZoneScope int32

const (
	ScheduleByTimeZoneScope_BOTH       ScheduleByTimeZoneScope = 0
	ScheduleByTimeZoneScope_STOP_DATE  ScheduleByTimeZoneScope = 1
	ScheduleByTimeZoneScope_START_DATE ScheduleByTimeZoneScope = 2
)

// Enum value maps for ScheduleByTimeZoneScope.
var (
	ScheduleByTimeZoneScope_name = map[int32]string{
		0: "BOTH",
		1: "STOP_DATE",
		2: "START_DATE",
	}
	ScheduleByTimeZoneScope_value = map[string]int32{
		"BOTH":       0,
		"STOP_DATE":  1,
		"START_DATE": 2,
	}
)

func (x ScheduleByTimeZoneScope) Enum() *ScheduleByTimeZoneScope {
	p := new(ScheduleByTimeZoneScope)
	*p = x
	return p
}

func (x ScheduleByTimeZoneScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScheduleByTimeZoneScope) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_org_preferences_proto_enumTypes[1].Descriptor()
}

func (ScheduleByTimeZoneScope) Type() protoreflect.EnumType {
	return &file_api_commons_org_preferences_proto_enumTypes[1]
}

func (x ScheduleByTimeZoneScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScheduleByTimeZoneScope.Descriptor instead.
func (ScheduleByTimeZoneScope) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_org_preferences_proto_rawDescGZIP(), []int{1}
}

type AnsweringMachineDetection int32

const (
	AnsweringMachineDetection_OPTIMIZE_MACHINE_DETECTION           AnsweringMachineDetection = 0
	AnsweringMachineDetection_OPTIMIZE_MACHINE_DETECTION_SLOW_LIVE AnsweringMachineDetection = 1
	AnsweringMachineDetection_OPTIMIZE_MACHINE_DELIVERY            AnsweringMachineDetection = 2
	AnsweringMachineDetection_BALANCED_DETECTION_AND_DELIVERY      AnsweringMachineDetection = 3
)

// Enum value maps for AnsweringMachineDetection.
var (
	AnsweringMachineDetection_name = map[int32]string{
		0: "OPTIMIZE_MACHINE_DETECTION",
		1: "OPTIMIZE_MACHINE_DETECTION_SLOW_LIVE",
		2: "OPTIMIZE_MACHINE_DELIVERY",
		3: "BALANCED_DETECTION_AND_DELIVERY",
	}
	AnsweringMachineDetection_value = map[string]int32{
		"OPTIMIZE_MACHINE_DETECTION":           0,
		"OPTIMIZE_MACHINE_DETECTION_SLOW_LIVE": 1,
		"OPTIMIZE_MACHINE_DELIVERY":            2,
		"BALANCED_DETECTION_AND_DELIVERY":      3,
	}
)

func (x AnsweringMachineDetection) Enum() *AnsweringMachineDetection {
	p := new(AnsweringMachineDetection)
	*p = x
	return p
}

func (x AnsweringMachineDetection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnsweringMachineDetection) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_org_preferences_proto_enumTypes[2].Descriptor()
}

func (AnsweringMachineDetection) Type() protoreflect.EnumType {
	return &file_api_commons_org_preferences_proto_enumTypes[2]
}

func (x AnsweringMachineDetection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnsweringMachineDetection.Descriptor instead.
func (AnsweringMachineDetection) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_org_preferences_proto_rawDescGZIP(), []int{2}
}

type DialListPenetrationStrategy int32

const (
	DialListPenetrationStrategy_DEPTH_FIRST   DialListPenetrationStrategy = 0
	DialListPenetrationStrategy_BREADTH_FIRST DialListPenetrationStrategy = 1
)

// Enum value maps for DialListPenetrationStrategy.
var (
	DialListPenetrationStrategy_name = map[int32]string{
		0: "DEPTH_FIRST",
		1: "BREADTH_FIRST",
	}
	DialListPenetrationStrategy_value = map[string]int32{
		"DEPTH_FIRST":   0,
		"BREADTH_FIRST": 1,
	}
)

func (x DialListPenetrationStrategy) Enum() *DialListPenetrationStrategy {
	p := new(DialListPenetrationStrategy)
	*p = x
	return p
}

func (x DialListPenetrationStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DialListPenetrationStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_org_preferences_proto_enumTypes[3].Descriptor()
}

func (DialListPenetrationStrategy) Type() protoreflect.EnumType {
	return &file_api_commons_org_preferences_proto_enumTypes[3]
}

func (x DialListPenetrationStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DialListPenetrationStrategy.Descriptor instead.
func (DialListPenetrationStrategy) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_org_preferences_proto_rawDescGZIP(), []int{3}
}

type StandardReportFilter int32

const (
	StandardReportFilter_NO_PREFERENCE                            StandardReportFilter = 0  // WHERE_NO_PREFERENCE
	StandardReportFilter_FILTER_BY_ANSWERED_CALLS                 StandardReportFilter = 1  // WHERE_ANSWERED_CALLS
	StandardReportFilter_FILTER_BY_ANSWERED_HANGUPS               StandardReportFilter = 2  // WHERE_ANSWERED_LINKCALL_HANGUP_CALLS
	StandardReportFilter_FILTER_BY_ANSWERED_LINKCALL              StandardReportFilter = 3  // WHERE_ANSWERED_LINKCALL_CALLS
	StandardReportFilter_FILTER_BY_ANY_KEY_PRESSED                StandardReportFilter = 4  // WHERE_ANY_KEY
	StandardReportFilter_FILTER_BY_BUSY                           StandardReportFilter = 5  // WHERE_BUSY_CALLS
	StandardReportFilter_FILTER_BY_CANCELED_CALLS                 StandardReportFilter = 6  // WHERE_CANCELED_CALLS
	StandardReportFilter_FILTER_BY_CONFIRM_KEYS_3_THROUGH_6       StandardReportFilter = 7  // WHERE_DIGIT_RESPONSE_CALLS
	StandardReportFilter_FILTER_BY_CONNECTED_CALLS                StandardReportFilter = 8  // WHERE_CONNECTED_CALLS
	StandardReportFilter_FILTER_BY_DNCL_CANCELED                  StandardReportFilter = 9  // WHERE_CANCELED_DNCL_CALLS
	StandardReportFilter_FILTER_BY_FAILED_CALLS                   StandardReportFilter = 10 // WHERE_FAILED_CALLS
	StandardReportFilter_FILTER_BY_INVALID_CALLS                  StandardReportFilter = 11 // WHERE_INVALID_CALLS
	StandardReportFilter_FILTER_BY_LINKCALL_ABANDONED             StandardReportFilter = 12 // WHERE_ANSWERED_LINKCALL_ABANDONNED_CALLS
	StandardReportFilter_FILTER_BY_MACHINE_CALLS                  StandardReportFilter = 13 // WHERE_MACHINE_CALLS
	StandardReportFilter_FILTER_BY_MACHINE_DELIVERY               StandardReportFilter = 14 // WHERE_MACHINE_DELIVERED_CALLS
	StandardReportFilter_FILTER_BY_MACHINE_HANGUP                 StandardReportFilter = 15 // WHERE_MACHINE_HANGUP_CALLS
	StandardReportFilter_FILTER_BY_MACHINE_AND_ANY_KEY            StandardReportFilter = 16 // WHERE_MACHINE_PLUS_ANY_KEY
	StandardReportFilter_FILTER_BY_NO_ANSWER                      StandardReportFilter = 17 // WHERE_NOANSWER_CALLS
	StandardReportFilter_FILTER_BY_NO_KEYS_PRESSED                StandardReportFilter = 18 // WHERE_NO_KEYS_PRESSED
	StandardReportFilter_FILTER_BY_NOT_CANCELED_CALLS             StandardReportFilter = 19 // WHERE_NOT_CANCELED_CALLS
	StandardReportFilter_FILTER_BY_UNCONNECTED_CALLS              StandardReportFilter = 20 // WHERE_UNCONNECTED_CALLS
	StandardReportFilter_FILTER_BY_UNCONNECTED_EXCLUDE_INVALID    StandardReportFilter = 21 // WHERE_UNCONNECTED_EXCLUDE_INVALID
	StandardReportFilter_FILTER_BY_0_KEY                          StandardReportFilter = 22 // WHERE_DIGIT_0_KEY
	StandardReportFilter_FILTER_BY_1_KEY                          StandardReportFilter = 23 // WHERE_DIGIT_1_KEY
	StandardReportFilter_FILTER_BY_2_KEY                          StandardReportFilter = 24 // WHERE_DIGIT_2_KEY
	StandardReportFilter_FILTER_BY_3_KEY                          StandardReportFilter = 25 // WHERE_DIGIT_3_KEY
	StandardReportFilter_FILTER_BY_4_KEY                          StandardReportFilter = 26 // WHERE_DIGIT_4_KEY
	StandardReportFilter_FILTER_BY_5_KEY                          StandardReportFilter = 27 // WHERE_DIGIT_5_KEY
	StandardReportFilter_FILTER_BY_6_KEY                          StandardReportFilter = 28 // WHERE_DIGIT_6_KEY
	StandardReportFilter_FILTER_BY_7_KEY                          StandardReportFilter = 29 // WHERE_DIGIT_7_KEY
	StandardReportFilter_FILTER_BY_8_KEY                          StandardReportFilter = 30 // WHERE_DIGIT_8_KEY
	StandardReportFilter_FILTER_BY_9_KEY                          StandardReportFilter = 31 // WHERE_DIGIT_9_KEY
	StandardReportFilter_FILTER_BY_STAR_KEY                       StandardReportFilter = 32 // WHERE_DIGIT_STAR_KEY
	StandardReportFilter_FILTER_BY_POUND_KEY                      StandardReportFilter = 33 // WHERE_DIGIT_POUND_KEY
	StandardReportFilter_FILTER_BY_MACHINE_HANGUP_AND_UNCONNECTED StandardReportFilter = 34 // WHERE_MACHINE_HANGUP_AND_UNCONNECTED
	StandardReportFilter_FILTER_BY_EXCLUDING_CANCELED_AND_INVALID StandardReportFilter = 35 // WHERE_ALL_EXCLUDING_INVALID_AND_CANCELED
)

// Enum value maps for StandardReportFilter.
var (
	StandardReportFilter_name = map[int32]string{
		0:  "NO_PREFERENCE",
		1:  "FILTER_BY_ANSWERED_CALLS",
		2:  "FILTER_BY_ANSWERED_HANGUPS",
		3:  "FILTER_BY_ANSWERED_LINKCALL",
		4:  "FILTER_BY_ANY_KEY_PRESSED",
		5:  "FILTER_BY_BUSY",
		6:  "FILTER_BY_CANCELED_CALLS",
		7:  "FILTER_BY_CONFIRM_KEYS_3_THROUGH_6",
		8:  "FILTER_BY_CONNECTED_CALLS",
		9:  "FILTER_BY_DNCL_CANCELED",
		10: "FILTER_BY_FAILED_CALLS",
		11: "FILTER_BY_INVALID_CALLS",
		12: "FILTER_BY_LINKCALL_ABANDONED",
		13: "FILTER_BY_MACHINE_CALLS",
		14: "FILTER_BY_MACHINE_DELIVERY",
		15: "FILTER_BY_MACHINE_HANGUP",
		16: "FILTER_BY_MACHINE_AND_ANY_KEY",
		17: "FILTER_BY_NO_ANSWER",
		18: "FILTER_BY_NO_KEYS_PRESSED",
		19: "FILTER_BY_NOT_CANCELED_CALLS",
		20: "FILTER_BY_UNCONNECTED_CALLS",
		21: "FILTER_BY_UNCONNECTED_EXCLUDE_INVALID",
		22: "FILTER_BY_0_KEY",
		23: "FILTER_BY_1_KEY",
		24: "FILTER_BY_2_KEY",
		25: "FILTER_BY_3_KEY",
		26: "FILTER_BY_4_KEY",
		27: "FILTER_BY_5_KEY",
		28: "FILTER_BY_6_KEY",
		29: "FILTER_BY_7_KEY",
		30: "FILTER_BY_8_KEY",
		31: "FILTER_BY_9_KEY",
		32: "FILTER_BY_STAR_KEY",
		33: "FILTER_BY_POUND_KEY",
		34: "FILTER_BY_MACHINE_HANGUP_AND_UNCONNECTED",
		35: "FILTER_BY_EXCLUDING_CANCELED_AND_INVALID",
	}
	StandardReportFilter_value = map[string]int32{
		"NO_PREFERENCE":                            0,
		"FILTER_BY_ANSWERED_CALLS":                 1,
		"FILTER_BY_ANSWERED_HANGUPS":               2,
		"FILTER_BY_ANSWERED_LINKCALL":              3,
		"FILTER_BY_ANY_KEY_PRESSED":                4,
		"FILTER_BY_BUSY":                           5,
		"FILTER_BY_CANCELED_CALLS":                 6,
		"FILTER_BY_CONFIRM_KEYS_3_THROUGH_6":       7,
		"FILTER_BY_CONNECTED_CALLS":                8,
		"FILTER_BY_DNCL_CANCELED":                  9,
		"FILTER_BY_FAILED_CALLS":                   10,
		"FILTER_BY_INVALID_CALLS":                  11,
		"FILTER_BY_LINKCALL_ABANDONED":             12,
		"FILTER_BY_MACHINE_CALLS":                  13,
		"FILTER_BY_MACHINE_DELIVERY":               14,
		"FILTER_BY_MACHINE_HANGUP":                 15,
		"FILTER_BY_MACHINE_AND_ANY_KEY":            16,
		"FILTER_BY_NO_ANSWER":                      17,
		"FILTER_BY_NO_KEYS_PRESSED":                18,
		"FILTER_BY_NOT_CANCELED_CALLS":             19,
		"FILTER_BY_UNCONNECTED_CALLS":              20,
		"FILTER_BY_UNCONNECTED_EXCLUDE_INVALID":    21,
		"FILTER_BY_0_KEY":                          22,
		"FILTER_BY_1_KEY":                          23,
		"FILTER_BY_2_KEY":                          24,
		"FILTER_BY_3_KEY":                          25,
		"FILTER_BY_4_KEY":                          26,
		"FILTER_BY_5_KEY":                          27,
		"FILTER_BY_6_KEY":                          28,
		"FILTER_BY_7_KEY":                          29,
		"FILTER_BY_8_KEY":                          30,
		"FILTER_BY_9_KEY":                          31,
		"FILTER_BY_STAR_KEY":                       32,
		"FILTER_BY_POUND_KEY":                      33,
		"FILTER_BY_MACHINE_HANGUP_AND_UNCONNECTED": 34,
		"FILTER_BY_EXCLUDING_CANCELED_AND_INVALID": 35,
	}
)

func (x StandardReportFilter) Enum() *StandardReportFilter {
	p := new(StandardReportFilter)
	*p = x
	return p
}

func (x StandardReportFilter) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StandardReportFilter) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_org_preferences_proto_enumTypes[4].Descriptor()
}

func (StandardReportFilter) Type() protoreflect.EnumType {
	return &file_api_commons_org_preferences_proto_enumTypes[4]
}

func (x StandardReportFilter) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StandardReportFilter.Descriptor instead.
func (StandardReportFilter) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_org_preferences_proto_rawDescGZIP(), []int{4}
}

// LocalePreferences represents the organization's locale settings.
type LocalePreferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Display language in Operator for users of the organization.
	// Must be a valid language tag supported in Operator.
	// IETF BCP 47 - https://en.wikipedia.org/wiki/IETF_language_tag
	LanguageTag string `protobuf:"bytes,1,opt,name=language_tag,json=languageTag,proto3" json:"language_tag,omitempty"`
	// The direction of the script in Operator used in the organization.
	// By default, it is set to false to indicate left-to-right.
	UseScriptDirectionRightToLeft bool `protobuf:"varint,2,opt,name=use_script_direction_right_to_left,json=useScriptDirectionRightToLeft,proto3" json:"use_script_direction_right_to_left,omitempty"`
	// The default currency used in the organization.
	// Must be a valid currency code supported in Operator.
	// ISO 4217 - https://en.wikipedia.org/wiki/ISO_4217.
	DefaultCurrency string `protobuf:"bytes,3,opt,name=default_currency,json=defaultCurrency,proto3" json:"default_currency,omitempty"`
}

func (x *LocalePreferences) Reset() {
	*x = LocalePreferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commons_org_preferences_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalePreferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalePreferences) ProtoMessage() {}

func (x *LocalePreferences) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_preferences_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalePreferences.ProtoReflect.Descriptor instead.
func (*LocalePreferences) Descriptor() ([]byte, []int) {
	return file_api_commons_org_preferences_proto_rawDescGZIP(), []int{0}
}

func (x *LocalePreferences) GetLanguageTag() string {
	if x != nil {
		return x.LanguageTag
	}
	return ""
}

func (x *LocalePreferences) GetUseScriptDirectionRightToLeft() bool {
	if x != nil {
		return x.UseScriptDirectionRightToLeft
	}
	return false
}

func (x *LocalePreferences) GetDefaultCurrency() string {
	if x != nil {
		return x.DefaultCurrency
	}
	return ""
}

var File_api_commons_org_preferences_proto protoreflect.FileDescriptor

var file_api_commons_org_preferences_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x22, 0xac, 0x01, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x61, 0x67, 0x12, 0x49, 0x0a, 0x22, 0x75, 0x73, 0x65,
	0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x75, 0x73, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x69, 0x67, 0x68, 0x74, 0x54, 0x6f,
	0x4c, 0x65, 0x66, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2a,
	0xac, 0x01, 0x0a, 0x19, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x0a,
	0x0b, 0x42, 0x59, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x42, 0x59, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x42, 0x59, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x5f,
	0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17,
	0x42, 0x59, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42,
	0x45, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x59, 0x5f,
	0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x53,
	0x43, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x59, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x05, 0x2a, 0x42,
	0x0a, 0x17, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x54,
	0x48, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x44, 0x41, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x45,
	0x10, 0x02, 0x2a, 0xa9, 0x01, 0x0a, 0x19, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x4d, 0x41, 0x43,
	0x48, 0x49, 0x4e, 0x45, 0x5f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00,
	0x12, 0x28, 0x0a, 0x24, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x4d, 0x41, 0x43,
	0x48, 0x49, 0x4e, 0x45, 0x5f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x4c, 0x4f, 0x57, 0x5f, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x50,
	0x54, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x42, 0x41, 0x4c,
	0x41, 0x4e, 0x43, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x41, 0x4e, 0x44, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x10, 0x03, 0x2a, 0x41,
	0x0a, 0x1b, 0x44, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x6e, 0x65, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0f, 0x0a,
	0x0b, 0x44, 0x45, 0x50, 0x54, 0x48, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x42, 0x52, 0x45, 0x41, 0x44, 0x54, 0x48, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10,
	0x01, 0x2a, 0x9b, 0x08, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f,
	0x5f, 0x50, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x00, 0x12, 0x1c, 0x0a,
	0x18, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45,
	0x52, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x53, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x45,
	0x44, 0x5f, 0x48, 0x41, 0x4e, 0x47, 0x55, 0x50, 0x53, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x45,
	0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19,
	0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x41, 0x4e, 0x59, 0x5f, 0x4b, 0x45,
	0x59, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x53, 0x45, 0x44, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x05, 0x12,
	0x1c, 0x0a, 0x18, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x53, 0x10, 0x06, 0x12, 0x26, 0x0a,
	0x22, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x52, 0x4d, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x5f, 0x33, 0x5f, 0x54, 0x48, 0x52, 0x4f, 0x55, 0x47,
	0x48, 0x5f, 0x36, 0x10, 0x07, 0x12, 0x1d, 0x0a, 0x19, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f,
	0x42, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c,
	0x4c, 0x53, 0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42,
	0x59, 0x5f, 0x44, 0x4e, 0x43, 0x4c, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10,
	0x09, 0x12, 0x1a, 0x0a, 0x16, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x53, 0x10, 0x0a, 0x12, 0x1b, 0x0a,
	0x17, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x53, 0x10, 0x0b, 0x12, 0x20, 0x0a, 0x1c, 0x46, 0x49,
	0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x43, 0x41, 0x4c, 0x4c,
	0x5f, 0x41, 0x42, 0x41, 0x4e, 0x44, 0x4f, 0x4e, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x1b, 0x0a, 0x17,
	0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e,
	0x45, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x53, 0x10, 0x0d, 0x12, 0x1e, 0x0a, 0x1a, 0x46, 0x49, 0x4c,
	0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x10, 0x0e, 0x12, 0x1c, 0x0a, 0x18, 0x46, 0x49, 0x4c,
	0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x48,
	0x41, 0x4e, 0x47, 0x55, 0x50, 0x10, 0x0f, 0x12, 0x21, 0x0a, 0x1d, 0x46, 0x49, 0x4c, 0x54, 0x45,
	0x52, 0x5f, 0x42, 0x59, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x41, 0x4e, 0x44,
	0x5f, 0x41, 0x4e, 0x59, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x10, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49,
	0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x4e, 0x4f, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45,
	0x52, 0x10, 0x11, 0x12, 0x1d, 0x0a, 0x19, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59,
	0x5f, 0x4e, 0x4f, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x53, 0x45, 0x44,
	0x10, 0x12, 0x12, 0x20, 0x0a, 0x1c, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c,
	0x4c, 0x53, 0x10, 0x13, 0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42,
	0x59, 0x5f, 0x55, 0x4e, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x43, 0x41,
	0x4c, 0x4c, 0x53, 0x10, 0x14, 0x12, 0x29, 0x0a, 0x25, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f,
	0x42, 0x59, 0x5f, 0x55, 0x4e, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x45,
	0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x15,
	0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x30, 0x5f,
	0x4b, 0x45, 0x59, 0x10, 0x16, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f,
	0x42, 0x59, 0x5f, 0x31, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x17, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49,
	0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x32, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x18, 0x12,
	0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x33, 0x5f, 0x4b,
	0x45, 0x59, 0x10, 0x19, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42,
	0x59, 0x5f, 0x34, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x1a, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c,
	0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x35, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x1b, 0x12, 0x13,
	0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x36, 0x5f, 0x4b, 0x45,
	0x59, 0x10, 0x1c, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59,
	0x5f, 0x37, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x1d, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x54,
	0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x38, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x1e, 0x12, 0x13, 0x0a,
	0x0f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x39, 0x5f, 0x4b, 0x45, 0x59,
	0x10, 0x1f, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f,
	0x53, 0x54, 0x41, 0x52, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x20, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49,
	0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x50, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x4b, 0x45,
	0x59, 0x10, 0x21, 0x12, 0x2c, 0x0a, 0x28, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59,
	0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x48, 0x41, 0x4e, 0x47, 0x55, 0x50, 0x5f,
	0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10,
	0x22, 0x12, 0x2c, 0x0a, 0x28, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x45,
	0x58, 0x43, 0x4c, 0x55, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45,
	0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x23, 0x42,
	0x9b, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x42, 0x13, 0x4f, 0x72, 0x67, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_org_preferences_proto_rawDescOnce sync.Once
	file_api_commons_org_preferences_proto_rawDescData = file_api_commons_org_preferences_proto_rawDesc
)

func file_api_commons_org_preferences_proto_rawDescGZIP() []byte {
	file_api_commons_org_preferences_proto_rawDescOnce.Do(func() {
		file_api_commons_org_preferences_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_org_preferences_proto_rawDescData)
	})
	return file_api_commons_org_preferences_proto_rawDescData
}

var file_api_commons_org_preferences_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_api_commons_org_preferences_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_commons_org_preferences_proto_goTypes = []interface{}{
	(BroadcastTemplateOrdering)(0),   // 0: api.commons.BroadcastTemplateOrdering
	(ScheduleByTimeZoneScope)(0),     // 1: api.commons.ScheduleByTimeZoneScope
	(AnsweringMachineDetection)(0),   // 2: api.commons.AnsweringMachineDetection
	(DialListPenetrationStrategy)(0), // 3: api.commons.DialListPenetrationStrategy
	(StandardReportFilter)(0),        // 4: api.commons.StandardReportFilter
	(*LocalePreferences)(nil),        // 5: api.commons.LocalePreferences
}
var file_api_commons_org_preferences_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_org_preferences_proto_init() }
func file_api_commons_org_preferences_proto_init() {
	if File_api_commons_org_preferences_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commons_org_preferences_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalePreferences); i {
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
			RawDescriptor: file_api_commons_org_preferences_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_org_preferences_proto_goTypes,
		DependencyIndexes: file_api_commons_org_preferences_proto_depIdxs,
		EnumInfos:         file_api_commons_org_preferences_proto_enumTypes,
		MessageInfos:      file_api_commons_org_preferences_proto_msgTypes,
	}.Build()
	File_api_commons_org_preferences_proto = out.File
	file_api_commons_org_preferences_proto_rawDesc = nil
	file_api_commons_org_preferences_proto_goTypes = nil
	file_api_commons_org_preferences_proto_depIdxs = nil
}
