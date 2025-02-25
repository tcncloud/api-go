// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/commons/callmonitor.proto

package commons

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

type HoldQueueMonitorStatus int32

const (
	HoldQueueMonitorStatus_HOLD_QUEUE_MONITOR_STATUS_UNSPECIFIED HoldQueueMonitorStatus = 0
	HoldQueueMonitorStatus_HOLD_QUEUE_MONITOR_STATUS_MONITORING  HoldQueueMonitorStatus = 1
	HoldQueueMonitorStatus_HOLD_QUEUE_MONITOR_STATUS_SUCCESS     HoldQueueMonitorStatus = 2
	HoldQueueMonitorStatus_HOLD_QUEUE_MONITOR_STATUS_FAILED      HoldQueueMonitorStatus = 3
	HoldQueueMonitorStatus_HOLD_QUEUE_MONITOR_STATUS_ENDED       HoldQueueMonitorStatus = 4
)

// Enum value maps for HoldQueueMonitorStatus.
var (
	HoldQueueMonitorStatus_name = map[int32]string{
		0: "HOLD_QUEUE_MONITOR_STATUS_UNSPECIFIED",
		1: "HOLD_QUEUE_MONITOR_STATUS_MONITORING",
		2: "HOLD_QUEUE_MONITOR_STATUS_SUCCESS",
		3: "HOLD_QUEUE_MONITOR_STATUS_FAILED",
		4: "HOLD_QUEUE_MONITOR_STATUS_ENDED",
	}
	HoldQueueMonitorStatus_value = map[string]int32{
		"HOLD_QUEUE_MONITOR_STATUS_UNSPECIFIED": 0,
		"HOLD_QUEUE_MONITOR_STATUS_MONITORING":  1,
		"HOLD_QUEUE_MONITOR_STATUS_SUCCESS":     2,
		"HOLD_QUEUE_MONITOR_STATUS_FAILED":      3,
		"HOLD_QUEUE_MONITOR_STATUS_ENDED":       4,
	}
)

func (x HoldQueueMonitorStatus) Enum() *HoldQueueMonitorStatus {
	p := new(HoldQueueMonitorStatus)
	*p = x
	return p
}

func (x HoldQueueMonitorStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HoldQueueMonitorStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_callmonitor_proto_enumTypes[0].Descriptor()
}

func (HoldQueueMonitorStatus) Type() protoreflect.EnumType {
	return &file_api_commons_callmonitor_proto_enumTypes[0]
}

func (x HoldQueueMonitorStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HoldQueueMonitorStatus.Descriptor instead.
func (HoldQueueMonitorStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_callmonitor_proto_rawDescGZIP(), []int{0}
}

type HoldQueueCallStats struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the call.
	CallId string `protobuf:"bytes,1,opt,name=call_id,json=callId,proto3" json:"call_id,omitempty"`
	// The id of the org.
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// The id of the omni campaign that the call is associated with.
	CampaignId string `protobuf:"bytes,3,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	// The phone number of the caller.
	PhoneNumber string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// The current status of the call.
	Status HoldQueueMonitorStatus `protobuf:"varint,5,opt,name=status,proto3,enum=api.commons.HoldQueueMonitorStatus" json:"status,omitempty"`
	// The total time in milliseconds that the call was monitored. Will be
	// undefined if the call is still being monitored.
	MonitorDurationMillis int64 `protobuf:"varint,7,opt,name=monitor_duration_millis,json=monitorDurationMillis,proto3" json:"monitor_duration_millis,omitempty"`
	// The time that the call was monitored.
	MonitorStartTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=monitor_start_time,json=monitorStartTime,proto3" json:"monitor_start_time,omitempty"`
	// The time that the call was no longer monitored. Will be undefined if
	// the call is still being monitored.
	MonitorEndTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=monitor_end_time,json=monitorEndTime,proto3" json:"monitor_end_time,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *HoldQueueCallStats) Reset() {
	*x = HoldQueueCallStats{}
	mi := &file_api_commons_callmonitor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HoldQueueCallStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HoldQueueCallStats) ProtoMessage() {}

func (x *HoldQueueCallStats) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_callmonitor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HoldQueueCallStats.ProtoReflect.Descriptor instead.
func (*HoldQueueCallStats) Descriptor() ([]byte, []int) {
	return file_api_commons_callmonitor_proto_rawDescGZIP(), []int{0}
}

func (x *HoldQueueCallStats) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

func (x *HoldQueueCallStats) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *HoldQueueCallStats) GetCampaignId() string {
	if x != nil {
		return x.CampaignId
	}
	return ""
}

func (x *HoldQueueCallStats) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *HoldQueueCallStats) GetStatus() HoldQueueMonitorStatus {
	if x != nil {
		return x.Status
	}
	return HoldQueueMonitorStatus_HOLD_QUEUE_MONITOR_STATUS_UNSPECIFIED
}

func (x *HoldQueueCallStats) GetMonitorDurationMillis() int64 {
	if x != nil {
		return x.MonitorDurationMillis
	}
	return 0
}

func (x *HoldQueueCallStats) GetMonitorStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.MonitorStartTime
	}
	return nil
}

func (x *HoldQueueCallStats) GetMonitorEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.MonitorEndTime
	}
	return nil
}

var File_api_commons_callmonitor_proto protoreflect.FileDescriptor

var file_api_commons_callmonitor_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x61,
	0x6c, 0x6c, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03,
	0x0a, 0x12, 0x48, 0x6f, 0x6c, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x48, 0x6f, 0x6c, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x48, 0x0a,
	0x12, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0xdf, 0x01,
	0x0a, 0x16, 0x48, 0x6f, 0x6c, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x25, 0x48, 0x4f, 0x4c, 0x44,
	0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x51, 0x55, 0x45, 0x55,
	0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x25, 0x0a,
	0x21, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x49,
	0x54, 0x4f, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x51, 0x55, 0x45,
	0x55, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x48, 0x4f,
	0x4c, 0x44, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x04, 0x42,
	0x98, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x42, 0x10, 0x43, 0x61, 0x6c, 0x6c, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xa2,
	0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_api_commons_callmonitor_proto_rawDescOnce sync.Once
	file_api_commons_callmonitor_proto_rawDescData []byte
)

func file_api_commons_callmonitor_proto_rawDescGZIP() []byte {
	file_api_commons_callmonitor_proto_rawDescOnce.Do(func() {
		file_api_commons_callmonitor_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_callmonitor_proto_rawDesc), len(file_api_commons_callmonitor_proto_rawDesc)))
	})
	return file_api_commons_callmonitor_proto_rawDescData
}

var file_api_commons_callmonitor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_callmonitor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_commons_callmonitor_proto_goTypes = []any{
	(HoldQueueMonitorStatus)(0),   // 0: api.commons.HoldQueueMonitorStatus
	(*HoldQueueCallStats)(nil),    // 1: api.commons.HoldQueueCallStats
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_api_commons_callmonitor_proto_depIdxs = []int32{
	0, // 0: api.commons.HoldQueueCallStats.status:type_name -> api.commons.HoldQueueMonitorStatus
	2, // 1: api.commons.HoldQueueCallStats.monitor_start_time:type_name -> google.protobuf.Timestamp
	2, // 2: api.commons.HoldQueueCallStats.monitor_end_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_commons_callmonitor_proto_init() }
func file_api_commons_callmonitor_proto_init() {
	if File_api_commons_callmonitor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_callmonitor_proto_rawDesc), len(file_api_commons_callmonitor_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_callmonitor_proto_goTypes,
		DependencyIndexes: file_api_commons_callmonitor_proto_depIdxs,
		EnumInfos:         file_api_commons_callmonitor_proto_enumTypes,
		MessageInfos:      file_api_commons_callmonitor_proto_msgTypes,
	}.Build()
	File_api_commons_callmonitor_proto = out.File
	file_api_commons_callmonitor_proto_goTypes = nil
	file_api_commons_callmonitor_proto_depIdxs = nil
}
