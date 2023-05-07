// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/commons/results.proto

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

type CallResult int32

const (
	CallResult_CALL_RESULT_UNKNOWN                                      CallResult = 0    // it's here only b/c protobuf requires it ... do never use it!
	CallResult_CALL_RESULT_PENDING                                      CallResult = 1000 // "CALL_PENDING", "Pending", "PED", "Call is pending result, default call result"),
	CallResult_CALL_RESULT_ANSWERED                                     CallResult = 2000 // "CALL_ANSWERED", "Answered", "ANS", "Call was answered, no other detail available"),
	CallResult_CALL_RESULT_ANSWERED_LINKCALL                            CallResult = 2100 // "CALL_ANSWERED_LINKCALL", "Answered Linkcall", "ANL", "Call was answered and a linkback was attempted"),
	CallResult_CALL_RESULT_ANSWERED_LINKCALL_ABANDONED                  CallResult = 2110 // "CALL_ANSWERED_LINKCALL_ABANDONED", "Answered Linkcall Abandoned", "ALA", "Call was answered and a linkback established, but was abandoned  before agent connect"),
	CallResult_CALL_RESULT_ANSWERED_LINKCALL_AGENT_TALK                 CallResult = 2120 // "CALL_ANSWERED_LINKCALL_AGENT_TALK", "Answered Linkcall Agent Talk", "AGT", "Call was answered and a linkback established to agent"),
	CallResult_CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED                  CallResult = 2130 // "CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED", "Answered Linkcall suspended", "call was suspended"
	CallResult_CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED_TIMEOUT          CallResult = 2140 // "CALL_ANSWERED_LINKCALL_SUSPENDED_TIMEOUT", "Suspend linkcall timeout", "call was timedout"),
	CallResult_CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED_INBOUND_OVERRIDE CallResult = 2150 // "CALL_ANSWERED_LINKCALL_LINKCALL_SUSPENDED_REPLACE_WITH_CALLER", "suspended call reaplaced with caller.)"
	CallResult_CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED_RESUMED          CallResult = 2160 // "CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED_RESUMED", "Answered Linkcall Suspended Resume", "suspended call resumed.)"
	CallResult_CALL_RESULT_ANSWERED_HANGUP                              CallResult = 2200 // "CALL_ANSWERED_HANGUP", "Answered Hangup", "ANH", "Call was answered and the TCN system hung up"),
	CallResult_CALL_RESULT_ANSWERED_VOICEMAIL                           CallResult = 2300 // "CALL_ANSWERED_VOICEMAIL", "Answered VoiceMail", "ANV", "Call was answered and a linkback was attempted, but went to voicemail" ),
	CallResult_CALL_RESULT_MACHINE                                      CallResult = 3000 // "CALL_MACHINE", "Machine", "MAC", "Call was detected as answering machine, no other details available"),
	CallResult_CALL_RESULT_MACHINE_DELIVERED                            CallResult = 3100 // "CALL_MACHINE_DELIVERED", "Machine Delivered", "MAD", "Call was detected as answering machine a message was delivered"),
	CallResult_CALL_RESULT_MACHINE_HANGUP                               CallResult = 3200 // "CALL_MACHINE_HANGUP", "Machine Hangup", "MAH", "Call was detected as answering machine and the TCN system hung up"),
	CallResult_CALL_RESULT_MACHINE_FAILED                               CallResult = 3300 // "CALL_RESULT_MACHINE_FAILED", "Machine Undeliverable", "MAF", "Call was detected as answering machine but remote machine hungup on TCN" ),
	CallResult_CALL_RESULT_FAX                                          CallResult = 4000 // "CALL_FAX", "Fax", "FAX", "Fax was detected, no other details available"),
	CallResult_CALL_RESULT_FAX_DELIVERED                                CallResult = 4100 // "CALL_FAX_DELIVERED", "Fax Delivered", "FAD", "Fax machine detected and a fax was sent"),
	CallResult_CALL_RESULT_BUSY                                         CallResult = 5000 // "CALL_BUSY", "Busy", "BZY", "Call was busy, no other details available"),
	CallResult_CALL_RESULT_NO_ANSWER                                    CallResult = 6000 // "CALL_NOANSWER", "No Answer", "NOA", "Call was not answered, no other details available"),
	CallResult_CALL_RESULT_INVALID                                      CallResult = 7000 // "CALL_INVALID", "Invalid", "INV", "Call was invalid, no other details available"),
	CallResult_CALL_RESULT_INVALID_INCOMPLETE_NUMBER                    CallResult = 7100 // "CALL_INVALID_INCOMPLETE_NUMBER", "Invalid Incomplete Number", "INC", "Call was invalid, incomplete number"),
	CallResult_CALL_RESULT_INVALID_UNKNOWN_PREFIX                       CallResult = 7200 // "CALL_INVALID_UNKNOWN_PREFIX", "Invalid Unknow Prefix", "IUP", "Call was invalid, unknown prefix"),
	CallResult_CALL_RESULT_INVALID_UNKNOWN_PREFIX_NPA                   CallResult = 7210 // "CALL_INVALID_UNKNOWN_PREFIX_NPA", "Invalid Unknown NPA", "INP", "Call was invalid, unknown NPA"),
	CallResult_CALL_RESULT_INVALID_UNKNOWN_PREFIX_NPANXX                CallResult = 7220 // "CALL_INVALID_UNKNOWN_PREFIX_NPANXX", "Invalid Unknown NPA-NXX", "INX", "Call was invalid,unknown NPA-NXX"),
	CallResult_CALL_RESULT_INVALID_PREFIX_NPANXX_NOT_FOUND              CallResult = 7230 // "CALL_INVALID_PREFIX_NPANXX_NOT_FOUND", "Invalid NPA-NXX Not Found", "INF", "Call was invalid, NPA-NXX not found"),
	CallResult_CALL_RESULT_INVALID_NO_ROUTE                             CallResult = 7300 // "CALL_INVALID_NO_ROUTE", "Invalid No Route", "INR", "Call invalid, no route to host"),
	CallResult_CALL_RESULT_INVALID_DISCONNECTED                         CallResult = 7400 // "CALL_INVALID_DISCONNECTED", "Invalid Disconnected", "IDC", "Call invalid, disconnected"),
	CallResult_CALL_RESULT_INVALID_DISCONNECTED_SKIPTRACE               CallResult = 7410 // "CALL_INVALID_DISCONNECTED_SKIPTRACE", "Invalid Disconnected Skiptrace", "IDS", "Call invalid, disconnected, tcn was able to get a number from the recording"),
	CallResult_CALL_RESULT_FAILED                                       CallResult = 8000 // "CALL_FAILED", "Failed", "FAL", "Call failed for unknown reason"),
	CallResult_CALL_RESULT_FAILED_NO_LINES                              CallResult = 8100 // "CALL_FAILED_NO_LINES", "Failed No Lines", "FNL", "Call failed," + " tcn had no available lines to deliver call"),
	CallResult_CALL_RESULT_FAILED_CIRCUITS_BUSY                         CallResult = 8200 // "CALL_FAILED_CIRCUITS_BUSY", "Failed Busy Circuits", "FBC", "Call failed, all circuits busy"),
	CallResult_CALL_RESULT_FAILED_REFUSED                               CallResult = 8300 // "CALL_FAILED_REFUSED", "Failed Refused", "FRE", "Call failed, tcn refused the call"),
	CallResult_CALL_RESULT_FAILED_REFUSED_LEGAL                         CallResult = 8310 // "CALL_FAILED_REFUSED_LEGAL", "Failed Refused Legal", "FRL", "Call failed, tcn refused the call on legal grounds"),
	CallResult_CALL_RESULT_FAILED_REFUSED_TECHNICAL                     CallResult = 8320 // "CALL_FAILED_REFUSED_TECHNICAL", "Failed Refused Technical", "FRT", "Call failed, tcn is unable to dial into requested region for technical reasons"),
	CallResult_CALL_RESULT_FAILED_INTERNAL_ERROR                        CallResult = 8330 // "CALL_RESULT_FAILED_INTERNAL_ERROR", "Failed Internal Technical Problem", "FIE", "Task integrity check failed!!!!!!"), // this could happen only in
	CallResult_CALL_RESULT_CANCELED                                     CallResult = 9000 // "CALL_CANCELED", "Canceled", "CAN", "Call canceled, reason unknown"),
	CallResult_CALL_RESULT_CANCELED_TIMEZONE                            CallResult = 9100 // "CALL_CANCELED_TIMEZONE", "Canceled Timezone", "CAZ", "Call canceled due to timezone restrictions"),
	CallResult_CALL_RESULT_CANCELED_TIMEOUT                             CallResult = 9200 // "CALL_CANCELED_TIMEOUT", "Canceled Timeout", "CAT", "Call canceled due to timeone"),
	CallResult_CALL_RESULT_CANCELED_DNCL                                CallResult = 9300 // "CALL_CANCELED_DNCL", "Canceled DNCL", "CDL", "Call canceled due to Do Not Call List"),
	CallResult_CALL_RESULT_CANCELED_CELLULAR_DNCL                       CallResult = 9310 // "CALL_CANCELED_CELLULAR_DNCL", "Canceled DNC Cellular", "CDC", "Call canceled due to requested cell phone scrubbing"),
	CallResult_CALL_RESULT_CANCELED_DNCL_ZIP_CODE                       CallResult = 9320 // "CALL_CANCELED_DNCL_ZIP_CODE", "Canceled DNC Zip Code", "CDZ", "Call canceled due to requested zip code scrubbing"),
	CallResult_CALL_RESULT_CANCELED_MAX_RETRY                           CallResult = 9400 // "CALL_RESULT_CANCELED_MAX_RETRY", "Canceled Max Retry", "CMR", "Call canceled due to Max Retry was reached"),
	CallResult_CALL_RESULT_CANCELED_INCOMPLETE_NUMBER                   CallResult = 9500 // "CALL_CANCELED_INCOMPLETE_NUMBER", "Cancelled Incomplete Number", "CIN", "Call was cancelled, incomplete number");
)

// Enum value maps for CallResult.
var (
	CallResult_name = map[int32]string{
		0:    "CALL_RESULT_UNKNOWN",
		1000: "CALL_RESULT_PENDING",
		2000: "CALL_RESULT_ANSWERED",
		2100: "CALL_RESULT_ANSWERED_LINKCALL",
		2110: "CALL_RESULT_ANSWERED_LINKCALL_ABANDONED",
		2120: "CALL_RESULT_ANSWERED_LINKCALL_AGENT_TALK",
		2130: "CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED",
		2140: "CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED_TIMEOUT",
		2150: "CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED_INBOUND_OVERRIDE",
		2160: "CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED_RESUMED",
		2200: "CALL_RESULT_ANSWERED_HANGUP",
		2300: "CALL_RESULT_ANSWERED_VOICEMAIL",
		3000: "CALL_RESULT_MACHINE",
		3100: "CALL_RESULT_MACHINE_DELIVERED",
		3200: "CALL_RESULT_MACHINE_HANGUP",
		3300: "CALL_RESULT_MACHINE_FAILED",
		4000: "CALL_RESULT_FAX",
		4100: "CALL_RESULT_FAX_DELIVERED",
		5000: "CALL_RESULT_BUSY",
		6000: "CALL_RESULT_NO_ANSWER",
		7000: "CALL_RESULT_INVALID",
		7100: "CALL_RESULT_INVALID_INCOMPLETE_NUMBER",
		7200: "CALL_RESULT_INVALID_UNKNOWN_PREFIX",
		7210: "CALL_RESULT_INVALID_UNKNOWN_PREFIX_NPA",
		7220: "CALL_RESULT_INVALID_UNKNOWN_PREFIX_NPANXX",
		7230: "CALL_RESULT_INVALID_PREFIX_NPANXX_NOT_FOUND",
		7300: "CALL_RESULT_INVALID_NO_ROUTE",
		7400: "CALL_RESULT_INVALID_DISCONNECTED",
		7410: "CALL_RESULT_INVALID_DISCONNECTED_SKIPTRACE",
		8000: "CALL_RESULT_FAILED",
		8100: "CALL_RESULT_FAILED_NO_LINES",
		8200: "CALL_RESULT_FAILED_CIRCUITS_BUSY",
		8300: "CALL_RESULT_FAILED_REFUSED",
		8310: "CALL_RESULT_FAILED_REFUSED_LEGAL",
		8320: "CALL_RESULT_FAILED_REFUSED_TECHNICAL",
		8330: "CALL_RESULT_FAILED_INTERNAL_ERROR",
		9000: "CALL_RESULT_CANCELED",
		9100: "CALL_RESULT_CANCELED_TIMEZONE",
		9200: "CALL_RESULT_CANCELED_TIMEOUT",
		9300: "CALL_RESULT_CANCELED_DNCL",
		9310: "CALL_RESULT_CANCELED_CELLULAR_DNCL",
		9320: "CALL_RESULT_CANCELED_DNCL_ZIP_CODE",
		9400: "CALL_RESULT_CANCELED_MAX_RETRY",
		9500: "CALL_RESULT_CANCELED_INCOMPLETE_NUMBER",
	}
	CallResult_value = map[string]int32{
		"CALL_RESULT_UNKNOWN":                                      0,
		"CALL_RESULT_PENDING":                                      1000,
		"CALL_RESULT_ANSWERED":                                     2000,
		"CALL_RESULT_ANSWERED_LINKCALL":                            2100,
		"CALL_RESULT_ANSWERED_LINKCALL_ABANDONED":                  2110,
		"CALL_RESULT_ANSWERED_LINKCALL_AGENT_TALK":                 2120,
		"CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED":                  2130,
		"CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED_TIMEOUT":          2140,
		"CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED_INBOUND_OVERRIDE": 2150,
		"CALL_RESULT_ANSWERED_LINKCALL_SUSPENDED_RESUMED":          2160,
		"CALL_RESULT_ANSWERED_HANGUP":                              2200,
		"CALL_RESULT_ANSWERED_VOICEMAIL":                           2300,
		"CALL_RESULT_MACHINE":                                      3000,
		"CALL_RESULT_MACHINE_DELIVERED":                            3100,
		"CALL_RESULT_MACHINE_HANGUP":                               3200,
		"CALL_RESULT_MACHINE_FAILED":                               3300,
		"CALL_RESULT_FAX":                                          4000,
		"CALL_RESULT_FAX_DELIVERED":                                4100,
		"CALL_RESULT_BUSY":                                         5000,
		"CALL_RESULT_NO_ANSWER":                                    6000,
		"CALL_RESULT_INVALID":                                      7000,
		"CALL_RESULT_INVALID_INCOMPLETE_NUMBER":                    7100,
		"CALL_RESULT_INVALID_UNKNOWN_PREFIX":                       7200,
		"CALL_RESULT_INVALID_UNKNOWN_PREFIX_NPA":                   7210,
		"CALL_RESULT_INVALID_UNKNOWN_PREFIX_NPANXX":                7220,
		"CALL_RESULT_INVALID_PREFIX_NPANXX_NOT_FOUND":              7230,
		"CALL_RESULT_INVALID_NO_ROUTE":                             7300,
		"CALL_RESULT_INVALID_DISCONNECTED":                         7400,
		"CALL_RESULT_INVALID_DISCONNECTED_SKIPTRACE":               7410,
		"CALL_RESULT_FAILED":                                       8000,
		"CALL_RESULT_FAILED_NO_LINES":                              8100,
		"CALL_RESULT_FAILED_CIRCUITS_BUSY":                         8200,
		"CALL_RESULT_FAILED_REFUSED":                               8300,
		"CALL_RESULT_FAILED_REFUSED_LEGAL":                         8310,
		"CALL_RESULT_FAILED_REFUSED_TECHNICAL":                     8320,
		"CALL_RESULT_FAILED_INTERNAL_ERROR":                        8330,
		"CALL_RESULT_CANCELED":                                     9000,
		"CALL_RESULT_CANCELED_TIMEZONE":                            9100,
		"CALL_RESULT_CANCELED_TIMEOUT":                             9200,
		"CALL_RESULT_CANCELED_DNCL":                                9300,
		"CALL_RESULT_CANCELED_CELLULAR_DNCL":                       9310,
		"CALL_RESULT_CANCELED_DNCL_ZIP_CODE":                       9320,
		"CALL_RESULT_CANCELED_MAX_RETRY":                           9400,
		"CALL_RESULT_CANCELED_INCOMPLETE_NUMBER":                   9500,
	}
)

func (x CallResult) Enum() *CallResult {
	p := new(CallResult)
	*p = x
	return p
}

func (x CallResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CallResult) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_results_proto_enumTypes[0].Descriptor()
}

func (CallResult) Type() protoreflect.EnumType {
	return &file_api_commons_results_proto_enumTypes[0]
}

func (x CallResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CallResult.Descriptor instead.
func (CallResult) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_results_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_results_proto protoreflect.FileDescriptor

var file_api_commons_results_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0x80, 0x0d, 0x0a, 0x0a, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x18, 0x0a, 0x13, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0xe8, 0x07, 0x12, 0x19, 0x0a, 0x14, 0x43, 0x41,
	0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52,
	0x45, 0x44, 0x10, 0xd0, 0x0f, 0x12, 0x22, 0x0a, 0x1d, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x4c, 0x49,
	0x4e, 0x4b, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0xb4, 0x10, 0x12, 0x2c, 0x0a, 0x27, 0x43, 0x41, 0x4c,
	0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x45,
	0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x41, 0x42, 0x41, 0x4e, 0x44,
	0x4f, 0x4e, 0x45, 0x44, 0x10, 0xbe, 0x10, 0x12, 0x2d, 0x0a, 0x28, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x45, 0x44, 0x5f,
	0x4c, 0x49, 0x4e, 0x4b, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x41, 0x4c, 0x4b, 0x10, 0xc8, 0x10, 0x12, 0x2c, 0x0a, 0x27, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x4c,
	0x49, 0x4e, 0x4b, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45,
	0x44, 0x10, 0xd2, 0x10, 0x12, 0x34, 0x0a, 0x2f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x4c, 0x49, 0x4e,
	0x4b, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0xdc, 0x10, 0x12, 0x3d, 0x0a, 0x38, 0x43, 0x41,
	0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52,
	0x45, 0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x55, 0x53, 0x50,
	0x45, 0x4e, 0x44, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x4f, 0x56,
	0x45, 0x52, 0x52, 0x49, 0x44, 0x45, 0x10, 0xe6, 0x10, 0x12, 0x34, 0x0a, 0x2f, 0x43, 0x41, 0x4c,
	0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x45,
	0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45,
	0x4e, 0x44, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4d, 0x45, 0x44, 0x10, 0xf0, 0x10, 0x12,
	0x20, 0x0a, 0x1b, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x41,
	0x4e, 0x53, 0x57, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x48, 0x41, 0x4e, 0x47, 0x55, 0x50, 0x10, 0x98,
	0x11, 0x12, 0x23, 0x0a, 0x1e, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54,
	0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x4d,
	0x41, 0x49, 0x4c, 0x10, 0xfc, 0x11, 0x12, 0x18, 0x0a, 0x13, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x10, 0xb8, 0x17,
	0x12, 0x22, 0x0a, 0x1d, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45,
	0x44, 0x10, 0x9c, 0x18, 0x12, 0x1f, 0x0a, 0x1a, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x48, 0x41, 0x4e, 0x47,
	0x55, 0x50, 0x10, 0x80, 0x19, 0x12, 0x1f, 0x0a, 0x1a, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0xe4, 0x19, 0x12, 0x14, 0x0a, 0x0f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41, 0x58, 0x10, 0xa0, 0x1f, 0x12, 0x1e, 0x0a, 0x19,
	0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41, 0x58, 0x5f,
	0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x84, 0x20, 0x12, 0x15, 0x0a, 0x10,
	0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x42, 0x55, 0x53, 0x59,
	0x10, 0x88, 0x27, 0x12, 0x1a, 0x0a, 0x15, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x5f, 0x4e, 0x4f, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x10, 0xf0, 0x2e, 0x12,
	0x18, 0x0a, 0x13, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xd8, 0x36, 0x12, 0x2a, 0x0a, 0x25, 0x43, 0x41, 0x4c,
	0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42,
	0x45, 0x52, 0x10, 0xbc, 0x37, 0x12, 0x27, 0x0a, 0x22, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x10, 0xa0, 0x38, 0x12, 0x2b,
	0x0a, 0x26, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x52,
	0x45, 0x46, 0x49, 0x58, 0x5f, 0x4e, 0x50, 0x41, 0x10, 0xaa, 0x38, 0x12, 0x2e, 0x0a, 0x29, 0x43,
	0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49,
	0x58, 0x5f, 0x4e, 0x50, 0x41, 0x4e, 0x58, 0x58, 0x10, 0xb4, 0x38, 0x12, 0x30, 0x0a, 0x2b, 0x43,
	0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x5f, 0x4e, 0x50, 0x41, 0x4e, 0x58, 0x58,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xbe, 0x38, 0x12, 0x21, 0x0a,
	0x1c, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4e, 0x4f, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x10, 0x84, 0x39,
	0x12, 0x25, 0x0a, 0x20, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x10, 0xe8, 0x39, 0x12, 0x2f, 0x0a, 0x2a, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44,
	0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x4b, 0x49, 0x50,
	0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0xf2, 0x39, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x41, 0x4c, 0x4c,
	0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xc0,
	0x3e, 0x12, 0x20, 0x0a, 0x1b, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x4e, 0x4f, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x53,
	0x10, 0xa4, 0x3f, 0x12, 0x25, 0x0a, 0x20, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x43, 0x49, 0x52, 0x43, 0x55, 0x49,
	0x54, 0x53, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x88, 0x40, 0x12, 0x1f, 0x0a, 0x1a, 0x43, 0x41,
	0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x5f, 0x52, 0x45, 0x46, 0x55, 0x53, 0x45, 0x44, 0x10, 0xec, 0x40, 0x12, 0x25, 0x0a, 0x20, 0x43,
	0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x5f, 0x52, 0x45, 0x46, 0x55, 0x53, 0x45, 0x44, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x10,
	0xf6, 0x40, 0x12, 0x29, 0x0a, 0x24, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c,
	0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x46, 0x55, 0x53, 0x45, 0x44,
	0x5f, 0x54, 0x45, 0x43, 0x48, 0x4e, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x80, 0x41, 0x12, 0x26, 0x0a,
	0x21, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x8a, 0x41, 0x12, 0x19, 0x0a, 0x14, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0xa8, 0x46,
	0x12, 0x22, 0x0a, 0x1d, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5a, 0x4f, 0x4e,
	0x45, 0x10, 0x8c, 0x47, 0x12, 0x21, 0x0a, 0x1c, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x4f, 0x55, 0x54, 0x10, 0xf0, 0x47, 0x12, 0x1e, 0x0a, 0x19, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x5f,
	0x44, 0x4e, 0x43, 0x4c, 0x10, 0xd4, 0x48, 0x12, 0x27, 0x0a, 0x22, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x5f,
	0x43, 0x45, 0x4c, 0x4c, 0x55, 0x4c, 0x41, 0x52, 0x5f, 0x44, 0x4e, 0x43, 0x4c, 0x10, 0xde, 0x48,
	0x12, 0x27, 0x0a, 0x22, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x5f, 0x44, 0x4e, 0x43, 0x4c, 0x5f, 0x5a, 0x49,
	0x50, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0xe8, 0x48, 0x12, 0x23, 0x0a, 0x1e, 0x43, 0x41, 0x4c,
	0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45,
	0x44, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x59, 0x10, 0xb8, 0x49, 0x12, 0x2b,
	0x0a, 0x26, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x9c, 0x4a, 0x42, 0x4c, 0x0a, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x63, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_commons_results_proto_rawDescOnce sync.Once
	file_api_commons_results_proto_rawDescData = file_api_commons_results_proto_rawDesc
)

func file_api_commons_results_proto_rawDescGZIP() []byte {
	file_api_commons_results_proto_rawDescOnce.Do(func() {
		file_api_commons_results_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_results_proto_rawDescData)
	})
	return file_api_commons_results_proto_rawDescData
}

var file_api_commons_results_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_results_proto_goTypes = []interface{}{
	(CallResult)(0), // 0: api.commons.CallResult
}
var file_api_commons_results_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_results_proto_init() }
func file_api_commons_results_proto_init() {
	if File_api_commons_results_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_results_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_results_proto_goTypes,
		DependencyIndexes: file_api_commons_results_proto_depIdxs,
		EnumInfos:         file_api_commons_results_proto_enumTypes,
	}.Build()
	File_api_commons_results_proto = out.File
	file_api_commons_results_proto_rawDesc = nil
	file_api_commons_results_proto_goTypes = nil
	file_api_commons_results_proto_depIdxs = nil
}
