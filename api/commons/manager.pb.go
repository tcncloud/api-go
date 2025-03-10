// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/commons/manager.proto

package commons

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type AgentInfo int32

const (
	AgentInfo_AGENT_INFO_ACTIVE_AGENTS        AgentInfo = 0
	AgentInfo_AGENT_INFO_IN_CONFERENCE_AGENTS AgentInfo = 1
	AgentInfo_AGENT_INFO_MANUAL_AGENTS        AgentInfo = 2
	AgentInfo_AGENT_INFO_PAUSED_AGENTS        AgentInfo = 3
	AgentInfo_AGENT_INFO_PREVIEW_AGENTS       AgentInfo = 4
	AgentInfo_AGENT_INFO_READY_AGENTS         AgentInfo = 5
	AgentInfo_AGENT_INFO_TOTAL_AGENTS         AgentInfo = 6
	AgentInfo_AGENT_INFO_TRANSFER_AGENTS      AgentInfo = 7
	AgentInfo_AGENT_INFO_WRAP_UP_AGENTS       AgentInfo = 8
)

// Enum value maps for AgentInfo.
var (
	AgentInfo_name = map[int32]string{
		0: "AGENT_INFO_ACTIVE_AGENTS",
		1: "AGENT_INFO_IN_CONFERENCE_AGENTS",
		2: "AGENT_INFO_MANUAL_AGENTS",
		3: "AGENT_INFO_PAUSED_AGENTS",
		4: "AGENT_INFO_PREVIEW_AGENTS",
		5: "AGENT_INFO_READY_AGENTS",
		6: "AGENT_INFO_TOTAL_AGENTS",
		7: "AGENT_INFO_TRANSFER_AGENTS",
		8: "AGENT_INFO_WRAP_UP_AGENTS",
	}
	AgentInfo_value = map[string]int32{
		"AGENT_INFO_ACTIVE_AGENTS":        0,
		"AGENT_INFO_IN_CONFERENCE_AGENTS": 1,
		"AGENT_INFO_MANUAL_AGENTS":        2,
		"AGENT_INFO_PAUSED_AGENTS":        3,
		"AGENT_INFO_PREVIEW_AGENTS":       4,
		"AGENT_INFO_READY_AGENTS":         5,
		"AGENT_INFO_TOTAL_AGENTS":         6,
		"AGENT_INFO_TRANSFER_AGENTS":      7,
		"AGENT_INFO_WRAP_UP_AGENTS":       8,
	}
)

func (x AgentInfo) Enum() *AgentInfo {
	p := new(AgentInfo)
	*p = x
	return p
}

func (x AgentInfo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentInfo) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_manager_proto_enumTypes[0].Descriptor()
}

func (AgentInfo) Type() protoreflect.EnumType {
	return &file_api_commons_manager_proto_enumTypes[0]
}

func (x AgentInfo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentInfo.Descriptor instead.
func (AgentInfo) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_manager_proto_rawDescGZIP(), []int{0}
}

type SkillStats int32

const (
	SkillStats_SKILL_STATS_AGENT_PEERED_CALLS       SkillStats = 0
	SkillStats_SKILL_STATS_AVERAGE_LENGTH           SkillStats = 1
	SkillStats_SKILL_STATS_CALL_COUNT               SkillStats = 2
	SkillStats_SKILL_STATS_CALL_SKILL_MAP           SkillStats = 3
	SkillStats_SKILL_STATS_CALL_TYPE                SkillStats = 4
	SkillStats_SKILL_STATS_CALLS_LIST               SkillStats = 5
	SkillStats_SKILL_STATS_DAILY_BY_SKILLS_KEY      SkillStats = 6
	SkillStats_SKILL_STATS_LONGEST_IN_QUEUE         SkillStats = 7
	SkillStats_SKILL_STATS_MATCHING_AGENTS_MD       SkillStats = 8  // manual dial
	SkillStats_SKILL_STATS_MATCHING_AGENTS_PC       SkillStats = 9  // preview call
	SkillStats_SKILL_STATS_MATCHING_AGENTS_LI       SkillStats = 10 // logged in
	SkillStats_SKILL_STATS_MATCHING_AGENTS_OC       SkillStats = 11 // on call
	SkillStats_SKILL_STATS_MATCHING_AGENTS_P        SkillStats = 12 // paused
	SkillStats_SKILL_STATS_MATCHING_AGENTS_W        SkillStats = 13 // waiting
	SkillStats_SKILL_STATS_MATCHING_AGENTS_WU       SkillStats = 14 // wrap up
	SkillStats_SKILL_STATS_MATCHING_AGENTS_TC       SkillStats = 15 // transfer call
	SkillStats_SKILL_STATS_MAXIMUM_LENGTH           SkillStats = 16
	SkillStats_SKILL_STATS_MINIMUM_LENGTH           SkillStats = 17
	SkillStats_SKILL_STATS_PBX_EXTENSION            SkillStats = 18
	SkillStats_SKILL_STATS_QUEUED_NOTIFICATION_TYPE SkillStats = 19
	SkillStats_SKILL_STATS_SKILL_SET                SkillStats = 20
	SkillStats_SKILL_STATS_TOTAL_LENGTH_FOR_AVERAGE SkillStats = 21
)

// Enum value maps for SkillStats.
var (
	SkillStats_name = map[int32]string{
		0:  "SKILL_STATS_AGENT_PEERED_CALLS",
		1:  "SKILL_STATS_AVERAGE_LENGTH",
		2:  "SKILL_STATS_CALL_COUNT",
		3:  "SKILL_STATS_CALL_SKILL_MAP",
		4:  "SKILL_STATS_CALL_TYPE",
		5:  "SKILL_STATS_CALLS_LIST",
		6:  "SKILL_STATS_DAILY_BY_SKILLS_KEY",
		7:  "SKILL_STATS_LONGEST_IN_QUEUE",
		8:  "SKILL_STATS_MATCHING_AGENTS_MD",
		9:  "SKILL_STATS_MATCHING_AGENTS_PC",
		10: "SKILL_STATS_MATCHING_AGENTS_LI",
		11: "SKILL_STATS_MATCHING_AGENTS_OC",
		12: "SKILL_STATS_MATCHING_AGENTS_P",
		13: "SKILL_STATS_MATCHING_AGENTS_W",
		14: "SKILL_STATS_MATCHING_AGENTS_WU",
		15: "SKILL_STATS_MATCHING_AGENTS_TC",
		16: "SKILL_STATS_MAXIMUM_LENGTH",
		17: "SKILL_STATS_MINIMUM_LENGTH",
		18: "SKILL_STATS_PBX_EXTENSION",
		19: "SKILL_STATS_QUEUED_NOTIFICATION_TYPE",
		20: "SKILL_STATS_SKILL_SET",
		21: "SKILL_STATS_TOTAL_LENGTH_FOR_AVERAGE",
	}
	SkillStats_value = map[string]int32{
		"SKILL_STATS_AGENT_PEERED_CALLS":       0,
		"SKILL_STATS_AVERAGE_LENGTH":           1,
		"SKILL_STATS_CALL_COUNT":               2,
		"SKILL_STATS_CALL_SKILL_MAP":           3,
		"SKILL_STATS_CALL_TYPE":                4,
		"SKILL_STATS_CALLS_LIST":               5,
		"SKILL_STATS_DAILY_BY_SKILLS_KEY":      6,
		"SKILL_STATS_LONGEST_IN_QUEUE":         7,
		"SKILL_STATS_MATCHING_AGENTS_MD":       8,
		"SKILL_STATS_MATCHING_AGENTS_PC":       9,
		"SKILL_STATS_MATCHING_AGENTS_LI":       10,
		"SKILL_STATS_MATCHING_AGENTS_OC":       11,
		"SKILL_STATS_MATCHING_AGENTS_P":        12,
		"SKILL_STATS_MATCHING_AGENTS_W":        13,
		"SKILL_STATS_MATCHING_AGENTS_WU":       14,
		"SKILL_STATS_MATCHING_AGENTS_TC":       15,
		"SKILL_STATS_MAXIMUM_LENGTH":           16,
		"SKILL_STATS_MINIMUM_LENGTH":           17,
		"SKILL_STATS_PBX_EXTENSION":            18,
		"SKILL_STATS_QUEUED_NOTIFICATION_TYPE": 19,
		"SKILL_STATS_SKILL_SET":                20,
		"SKILL_STATS_TOTAL_LENGTH_FOR_AVERAGE": 21,
	}
)

func (x SkillStats) Enum() *SkillStats {
	p := new(SkillStats)
	*p = x
	return p
}

func (x SkillStats) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SkillStats) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_manager_proto_enumTypes[1].Descriptor()
}

func (SkillStats) Type() protoreflect.EnumType {
	return &file_api_commons_manager_proto_enumTypes[1]
}

func (x SkillStats) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SkillStats.Descriptor instead.
func (SkillStats) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_manager_proto_rawDescGZIP(), []int{1}
}

type SkillQueues int32

const (
	SkillQueues_SKILL_QUEUES_ACD_QUEUE          SkillQueues = 0
	SkillQueues_SKILL_QUEUES_MULTI_HOLD         SkillQueues = 1
	SkillQueues_SKILL_QUEUES_SIMPLE_HOLD        SkillQueues = 2
	SkillQueues_SKILL_QUEUES_TRANSFER           SkillQueues = 3
	SkillQueues_SKILL_QUEUES_SUSPENDED_CALLBACK SkillQueues = 4
	SkillQueues_SKILL_QUEUES_GRAND_TOTALS       SkillQueues = 5
)

// Enum value maps for SkillQueues.
var (
	SkillQueues_name = map[int32]string{
		0: "SKILL_QUEUES_ACD_QUEUE",
		1: "SKILL_QUEUES_MULTI_HOLD",
		2: "SKILL_QUEUES_SIMPLE_HOLD",
		3: "SKILL_QUEUES_TRANSFER",
		4: "SKILL_QUEUES_SUSPENDED_CALLBACK",
		5: "SKILL_QUEUES_GRAND_TOTALS",
	}
	SkillQueues_value = map[string]int32{
		"SKILL_QUEUES_ACD_QUEUE":          0,
		"SKILL_QUEUES_MULTI_HOLD":         1,
		"SKILL_QUEUES_SIMPLE_HOLD":        2,
		"SKILL_QUEUES_TRANSFER":           3,
		"SKILL_QUEUES_SUSPENDED_CALLBACK": 4,
		"SKILL_QUEUES_GRAND_TOTALS":       5,
	}
)

func (x SkillQueues) Enum() *SkillQueues {
	p := new(SkillQueues)
	*p = x
	return p
}

func (x SkillQueues) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SkillQueues) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_manager_proto_enumTypes[2].Descriptor()
}

func (SkillQueues) Type() protoreflect.EnumType {
	return &file_api_commons_manager_proto_enumTypes[2]
}

func (x SkillQueues) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SkillQueues.Descriptor instead.
func (SkillQueues) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_manager_proto_rawDescGZIP(), []int{2}
}

type AgentStats int32

const (
	AgentStats_AGENT_STATS_AGENT_NAME         AgentStats = 0
	AgentStats_AGENT_STATS_AGENT_STATUS       AgentStats = 1
	AgentStats_AGENT_STATS_DURATION_IN_STATUS AgentStats = 2
	AgentStats_AGENT_STATS_LOGIN_DATE_TIME    AgentStats = 3
	AgentStats_AGENT_STATS_LOGIN_DURATION     AgentStats = 4
	AgentStats_AGENT_STATS_AGENT_SKILLS       AgentStats = 5
	AgentStats_AGENT_STATS_AGENT_SID          AgentStats = 6
	AgentStats_AGENT_STATS_SESSION_ID         AgentStats = 7
	AgentStats_AGENT_STATS_HUNT_GROUP_NAME    AgentStats = 8
	AgentStats_AGENT_STATS_CALL_COUNT         AgentStats = 9
	AgentStats_AGENT_STATS_PAUSE_CODE         AgentStats = 10
	AgentStats_AGENT_STATS_RECORDING_STATUS   AgentStats = 11
	AgentStats_AGENT_STATS_MULTI_HOLD_COUNT   AgentStats = 12
	AgentStats_AGENT_STATS_SIMPLE_HOLD_COUNT  AgentStats = 13
	AgentStats_AGENT_STATS_TOTAL_HOLD_COUNT   AgentStats = 14
)

// Enum value maps for AgentStats.
var (
	AgentStats_name = map[int32]string{
		0:  "AGENT_STATS_AGENT_NAME",
		1:  "AGENT_STATS_AGENT_STATUS",
		2:  "AGENT_STATS_DURATION_IN_STATUS",
		3:  "AGENT_STATS_LOGIN_DATE_TIME",
		4:  "AGENT_STATS_LOGIN_DURATION",
		5:  "AGENT_STATS_AGENT_SKILLS",
		6:  "AGENT_STATS_AGENT_SID",
		7:  "AGENT_STATS_SESSION_ID",
		8:  "AGENT_STATS_HUNT_GROUP_NAME",
		9:  "AGENT_STATS_CALL_COUNT",
		10: "AGENT_STATS_PAUSE_CODE",
		11: "AGENT_STATS_RECORDING_STATUS",
		12: "AGENT_STATS_MULTI_HOLD_COUNT",
		13: "AGENT_STATS_SIMPLE_HOLD_COUNT",
		14: "AGENT_STATS_TOTAL_HOLD_COUNT",
	}
	AgentStats_value = map[string]int32{
		"AGENT_STATS_AGENT_NAME":         0,
		"AGENT_STATS_AGENT_STATUS":       1,
		"AGENT_STATS_DURATION_IN_STATUS": 2,
		"AGENT_STATS_LOGIN_DATE_TIME":    3,
		"AGENT_STATS_LOGIN_DURATION":     4,
		"AGENT_STATS_AGENT_SKILLS":       5,
		"AGENT_STATS_AGENT_SID":          6,
		"AGENT_STATS_SESSION_ID":         7,
		"AGENT_STATS_HUNT_GROUP_NAME":    8,
		"AGENT_STATS_CALL_COUNT":         9,
		"AGENT_STATS_PAUSE_CODE":         10,
		"AGENT_STATS_RECORDING_STATUS":   11,
		"AGENT_STATS_MULTI_HOLD_COUNT":   12,
		"AGENT_STATS_SIMPLE_HOLD_COUNT":  13,
		"AGENT_STATS_TOTAL_HOLD_COUNT":   14,
	}
)

func (x AgentStats) Enum() *AgentStats {
	p := new(AgentStats)
	*p = x
	return p
}

func (x AgentStats) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentStats) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_manager_proto_enumTypes[3].Descriptor()
}

func (AgentStats) Type() protoreflect.EnumType {
	return &file_api_commons_manager_proto_enumTypes[3]
}

func (x AgentStats) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentStats.Descriptor instead.
func (AgentStats) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_manager_proto_rawDescGZIP(), []int{3}
}

type ManagerBargeInMode int32

const (
	ManagerBargeInMode_MONITOR         ManagerBargeInMode = 0
	ManagerBargeInMode_FULL_CONFERENCE ManagerBargeInMode = 1
	ManagerBargeInMode_AGENT_WHISPER   ManagerBargeInMode = 2
)

// Enum value maps for ManagerBargeInMode.
var (
	ManagerBargeInMode_name = map[int32]string{
		0: "MONITOR",
		1: "FULL_CONFERENCE",
		2: "AGENT_WHISPER",
	}
	ManagerBargeInMode_value = map[string]int32{
		"MONITOR":         0,
		"FULL_CONFERENCE": 1,
		"AGENT_WHISPER":   2,
	}
)

func (x ManagerBargeInMode) Enum() *ManagerBargeInMode {
	p := new(ManagerBargeInMode)
	*p = x
	return p
}

func (x ManagerBargeInMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ManagerBargeInMode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_manager_proto_enumTypes[4].Descriptor()
}

func (ManagerBargeInMode) Type() protoreflect.EnumType {
	return &file_api_commons_manager_proto_enumTypes[4]
}

func (x ManagerBargeInMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ManagerBargeInMode.Descriptor instead.
func (ManagerBargeInMode) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_manager_proto_rawDescGZIP(), []int{4}
}

var File_api_commons_manager_proto protoreflect.FileDescriptor

var file_api_commons_manager_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0xa2, 0x02, 0x0a, 0x09, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f,
	0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x41, 0x47, 0x45, 0x4e,
	0x54, 0x53, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e,
	0x46, 0x4f, 0x5f, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45,
	0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x47, 0x45,
	0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x5f, 0x41,
	0x47, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x47, 0x45, 0x4e, 0x54,
	0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x5f, 0x41, 0x47, 0x45,
	0x4e, 0x54, 0x53, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x49,
	0x4e, 0x46, 0x4f, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x41, 0x47, 0x45, 0x4e,
	0x54, 0x53, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e,
	0x46, 0x4f, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x10,
	0x05, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f,
	0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x06, 0x12, 0x1e,
	0x0a, 0x1a, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x07, 0x12, 0x1d,
	0x0a, 0x19, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x57, 0x52, 0x41,
	0x50, 0x5f, 0x55, 0x50, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x08, 0x2a, 0xf6, 0x05,
	0x0a, 0x0a, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x1e,
	0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x41, 0x47, 0x45, 0x4e,
	0x54, 0x5f, 0x50, 0x45, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x53, 0x10, 0x00,
	0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f,
	0x41, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f,
	0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a,
	0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x43, 0x41, 0x4c, 0x4c,
	0x5f, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x4d, 0x41, 0x50, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15,
	0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x43, 0x41, 0x4c, 0x4c,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x4b, 0x49, 0x4c, 0x4c,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x53, 0x5f, 0x4c, 0x49, 0x53,
	0x54, 0x10, 0x05, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x53, 0x5f, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x5f, 0x42, 0x59, 0x5f, 0x53, 0x4b, 0x49, 0x4c,
	0x4c, 0x53, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x06, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x4b, 0x49, 0x4c,
	0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x45, 0x53, 0x54, 0x5f,
	0x49, 0x4e, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x10, 0x07, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x4b,
	0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x49,
	0x4e, 0x47, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x4d, 0x44, 0x10, 0x08, 0x12, 0x22,
	0x0a, 0x1e, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x50, 0x43,
	0x10, 0x09, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x53, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54,
	0x53, 0x5f, 0x4c, 0x49, 0x10, 0x0a, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x41,
	0x47, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x4f, 0x43, 0x10, 0x0b, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x4b,
	0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x49,
	0x4e, 0x47, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x50, 0x10, 0x0c, 0x12, 0x21, 0x0a,
	0x1d, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x4d, 0x41, 0x54,
	0x43, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x57, 0x10, 0x0d,
	0x12, 0x22, 0x0a, 0x1e, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x5f,
	0x57, 0x55, 0x10, 0x0e, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x53, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x47, 0x45,
	0x4e, 0x54, 0x53, 0x5f, 0x54, 0x43, 0x10, 0x0f, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x4b, 0x49, 0x4c,
	0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x4d, 0x41, 0x58, 0x49, 0x4d, 0x55, 0x4d, 0x5f,
	0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x10, 0x10, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x4b, 0x49, 0x4c,
	0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x4d, 0x55, 0x4d, 0x5f,
	0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x10, 0x11, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x4b, 0x49, 0x4c,
	0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x50, 0x42, 0x58, 0x5f, 0x45, 0x58, 0x54, 0x45,
	0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x12, 0x12, 0x28, 0x0a, 0x24, 0x53, 0x4b, 0x49, 0x4c, 0x4c,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x5f, 0x4e, 0x4f,
	0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10,
	0x13, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53,
	0x5f, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x14, 0x12, 0x28, 0x0a, 0x24,
	0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x54, 0x4f, 0x54, 0x41,
	0x4c, 0x5f, 0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41, 0x56, 0x45,
	0x52, 0x41, 0x47, 0x45, 0x10, 0x15, 0x2a, 0xc3, 0x01, 0x0a, 0x0b, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f,
	0x51, 0x55, 0x45, 0x55, 0x45, 0x53, 0x5f, 0x41, 0x43, 0x44, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45,
	0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x51, 0x55, 0x45, 0x55,
	0x45, 0x53, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x01, 0x12,
	0x1c, 0x0a, 0x18, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x53, 0x5f,
	0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x02, 0x12, 0x19, 0x0a,
	0x15, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x53, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x4b, 0x49, 0x4c,
	0x4c, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x53, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44,
	0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x04, 0x12, 0x1d, 0x0a,
	0x19, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x53, 0x5f, 0x47, 0x52,
	0x41, 0x4e, 0x44, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x53, 0x10, 0x05, 0x2a, 0xe2, 0x03, 0x0a,
	0x0a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x41,
	0x47, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x47, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x53, 0x5f, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x47, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x47,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f,
	0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x47,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x53, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x47, 0x45, 0x4e,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x49,
	0x44, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x53, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x07, 0x12,
	0x1f, 0x0a, 0x1b, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x48,
	0x55, 0x4e, 0x54, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x08,
	0x12, 0x1a, 0x0a, 0x16, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f,
	0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x09, 0x12, 0x1a, 0x0a, 0x16,
	0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x50, 0x41, 0x55, 0x53,
	0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x0a, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x47, 0x45, 0x4e,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x0b, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x47,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f,
	0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x0c, 0x12, 0x21, 0x0a, 0x1d,
	0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x53, 0x49, 0x4d, 0x50,
	0x4c, 0x45, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x0d, 0x12,
	0x20, 0x0a, 0x1c, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x54,
	0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10,
	0x0e, 0x2a, 0x49, 0x0a, 0x12, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42, 0x61, 0x72, 0x67,
	0x65, 0x49, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x49, 0x54,
	0x4f, 0x52, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x4e,
	0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x47, 0x45,
	0x4e, 0x54, 0x5f, 0x57, 0x48, 0x49, 0x53, 0x50, 0x45, 0x52, 0x10, 0x02, 0x42, 0x94, 0x01, 0x0a,
	0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x42, 0x0c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02,
	0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41,
	0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_commons_manager_proto_rawDescOnce sync.Once
	file_api_commons_manager_proto_rawDescData []byte
)

func file_api_commons_manager_proto_rawDescGZIP() []byte {
	file_api_commons_manager_proto_rawDescOnce.Do(func() {
		file_api_commons_manager_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_manager_proto_rawDesc), len(file_api_commons_manager_proto_rawDesc)))
	})
	return file_api_commons_manager_proto_rawDescData
}

var file_api_commons_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_api_commons_manager_proto_goTypes = []any{
	(AgentInfo)(0),          // 0: api.commons.AgentInfo
	(SkillStats)(0),         // 1: api.commons.SkillStats
	(SkillQueues)(0),        // 2: api.commons.SkillQueues
	(AgentStats)(0),         // 3: api.commons.AgentStats
	(ManagerBargeInMode)(0), // 4: api.commons.ManagerBargeInMode
}
var file_api_commons_manager_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_manager_proto_init() }
func file_api_commons_manager_proto_init() {
	if File_api_commons_manager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_manager_proto_rawDesc), len(file_api_commons_manager_proto_rawDesc)),
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_manager_proto_goTypes,
		DependencyIndexes: file_api_commons_manager_proto_depIdxs,
		EnumInfos:         file_api_commons_manager_proto_enumTypes,
	}.Build()
	File_api_commons_manager_proto = out.File
	file_api_commons_manager_proto_goTypes = nil
	file_api_commons_manager_proto_depIdxs = nil
}
