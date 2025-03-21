// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/commons/call.proto

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

type CallStatus int32

const (
	CallStatus_CALL_UNKNOWN   CallStatus = 0
	CallStatus_CALL_SCHEDULED CallStatus = 3100 // "CALL_SCHEDULED", "Call is currently waiting for scheduler"),
	CallStatus_CALL_RUNNING   CallStatus = 3200 // "CALL_RUNNING", "Call is currently being executed"),
	CallStatus_CALL_COMPLETED CallStatus = 3300 // "CALL_COMPLTED", "Call has been executed"),
)

// Enum value maps for CallStatus.
var (
	CallStatus_name = map[int32]string{
		0:    "CALL_UNKNOWN",
		3100: "CALL_SCHEDULED",
		3200: "CALL_RUNNING",
		3300: "CALL_COMPLETED",
	}
	CallStatus_value = map[string]int32{
		"CALL_UNKNOWN":   0,
		"CALL_SCHEDULED": 3100,
		"CALL_RUNNING":   3200,
		"CALL_COMPLETED": 3300,
	}
)

func (x CallStatus) Enum() *CallStatus {
	p := new(CallStatus)
	*p = x
	return p
}

func (x CallStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CallStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_call_proto_enumTypes[0].Descriptor()
}

func (CallStatus) Type() protoreflect.EnumType {
	return &file_api_commons_call_proto_enumTypes[0]
}

func (x CallStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CallStatus.Descriptor instead.
func (CallStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_call_proto_rawDescGZIP(), []int{0}
}

type SimpleCallData struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	TaskSid           int64                  `protobuf:"varint,1,opt,name=task_sid,json=taskSid,proto3" json:"task_sid,omitempty"`
	CallSid           int64                  `protobuf:"varint,2,opt,name=call_sid,json=callSid,proto3" json:"call_sid,omitempty"`
	TaskGroupSid      int64                  `protobuf:"varint,3,opt,name=task_group_sid,json=taskGroupSid,proto3" json:"task_group_sid,omitempty"`
	ClientSid         int64                  `protobuf:"varint,4,opt,name=client_sid,json=clientSid,proto3" json:"client_sid,omitempty"`
	CountrySid        int64                  `protobuf:"varint,5,opt,name=country_sid,json=countrySid,proto3" json:"country_sid,omitempty"`
	AgentSid          int64                  `protobuf:"varint,6,opt,name=agent_sid,json=agentSid,proto3" json:"agent_sid,omitempty"`
	StartTime         int64                  `protobuf:"varint,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	CallerId          string                 `protobuf:"bytes,8,opt,name=caller_id,json=callerId,proto3" json:"caller_id,omitempty"`
	PhoneNumber       string                 `protobuf:"bytes,9,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	CountryCode       string                 `protobuf:"bytes,10,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	DeliveryDuration  int32                  `protobuf:"varint,11,opt,name=delivery_duration,json=deliveryDuration,proto3" json:"delivery_duration,omitempty"`
	LinkCallDuration  int32                  `protobuf:"varint,12,opt,name=link_call_duration,json=linkCallDuration,proto3" json:"link_call_duration,omitempty"`
	Result            CallResult             `protobuf:"varint,13,opt,name=result,proto3,enum=api.commons.CallResult" json:"result,omitempty"`
	SipCode           int32                  `protobuf:"varint,14,opt,name=sip_code,json=sipCode,proto3" json:"sip_code,omitempty"`
	DoRecord          bool                   `protobuf:"varint,15,opt,name=do_record,json=doRecord,proto3" json:"do_record,omitempty"`
	RecordingFileName string                 `protobuf:"bytes,16,opt,name=recording_file_name,json=recordingFileName,proto3" json:"recording_file_name,omitempty"`
	// WARNING: don't delete these fields randomly, the frontend checks the inverse
	// so reserved fields will trigger a false positive for the scrub. Most likely you
	// will need to hardcode the value to True in matrix-api (like custom calling rules)
	IsDialValidationOk          bool              `protobuf:"varint,17,opt,name=is_dial_validation_ok,json=isDialValidationOk,proto3" json:"is_dial_validation_ok,omitempty"`
	IsTimeZoneScrubOk           bool              `protobuf:"varint,18,opt,name=is_time_zone_scrub_ok,json=isTimeZoneScrubOk,proto3" json:"is_time_zone_scrub_ok,omitempty"`
	IsCellPhoneScrubOk          bool              `protobuf:"varint,19,opt,name=is_cell_phone_scrub_ok,json=isCellPhoneScrubOk,proto3" json:"is_cell_phone_scrub_ok,omitempty"`
	IsCustomCallingRulesScrubOk bool              `protobuf:"varint,20,opt,name=is_custom_calling_rules_scrub_ok,json=isCustomCallingRulesScrubOk,proto3" json:"is_custom_calling_rules_scrub_ok,omitempty"`
	IsDnclScrubOk               bool              `protobuf:"varint,21,opt,name=is_dncl_scrub_ok,json=isDnclScrubOk,proto3" json:"is_dncl_scrub_ok,omitempty"`
	UseGlobalTimeZoneScrub      bool              `protobuf:"varint,22,opt,name=use_global_time_zone_scrub,json=useGlobalTimeZoneScrub,proto3" json:"use_global_time_zone_scrub,omitempty"`
	DoCellPhoneScrub            bool              `protobuf:"varint,23,opt,name=do_cell_phone_scrub,json=doCellPhoneScrub,proto3" json:"do_cell_phone_scrub,omitempty"`
	DoDnclScrub                 bool              `protobuf:"varint,25,opt,name=do_dncl_scrub,json=doDnclScrub,proto3" json:"do_dncl_scrub,omitempty"`
	CallDataType                string            `protobuf:"bytes,26,opt,name=call_data_type,json=callDataType,proto3" json:"call_data_type,omitempty"`
	CallerIdCountryCode         string            `protobuf:"bytes,28,opt,name=caller_id_country_code,json=callerIdCountryCode,proto3" json:"caller_id_country_code,omitempty"`
	CallerIdCountrySid          int64             `protobuf:"varint,29,opt,name=caller_id_country_sid,json=callerIdCountrySid,proto3" json:"caller_id_country_sid,omitempty"`
	ZipCode                     string            `protobuf:"bytes,30,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	IsPreviewByRecord           bool              `protobuf:"varint,31,opt,name=is_preview_by_record,json=isPreviewByRecord,proto3" json:"is_preview_by_record,omitempty"`
	RuleSetName                 string            `protobuf:"bytes,32,opt,name=rule_set_name,json=ruleSetName,proto3" json:"rule_set_name,omitempty"`
	IsNaturalComplianceOk       bool              `protobuf:"varint,33,opt,name=is_natural_compliance_ok,json=isNaturalComplianceOk,proto3" json:"is_natural_compliance_ok,omitempty"`
	SimpleMetaData              []*SimpleKeyValue `protobuf:"bytes,34,rep,name=simple_meta_data,json=simpleMetaData,proto3" json:"simple_meta_data,omitempty"`
	SimpleResultMetaData        []*SimpleKeyValue `protobuf:"bytes,35,rep,name=simple_result_meta_data,json=simpleResultMetaData,proto3" json:"simple_result_meta_data,omitempty"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *SimpleCallData) Reset() {
	*x = SimpleCallData{}
	mi := &file_api_commons_call_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimpleCallData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleCallData) ProtoMessage() {}

func (x *SimpleCallData) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_call_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleCallData.ProtoReflect.Descriptor instead.
func (*SimpleCallData) Descriptor() ([]byte, []int) {
	return file_api_commons_call_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleCallData) GetTaskSid() int64 {
	if x != nil {
		return x.TaskSid
	}
	return 0
}

func (x *SimpleCallData) GetCallSid() int64 {
	if x != nil {
		return x.CallSid
	}
	return 0
}

func (x *SimpleCallData) GetTaskGroupSid() int64 {
	if x != nil {
		return x.TaskGroupSid
	}
	return 0
}

func (x *SimpleCallData) GetClientSid() int64 {
	if x != nil {
		return x.ClientSid
	}
	return 0
}

func (x *SimpleCallData) GetCountrySid() int64 {
	if x != nil {
		return x.CountrySid
	}
	return 0
}

func (x *SimpleCallData) GetAgentSid() int64 {
	if x != nil {
		return x.AgentSid
	}
	return 0
}

func (x *SimpleCallData) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *SimpleCallData) GetCallerId() string {
	if x != nil {
		return x.CallerId
	}
	return ""
}

func (x *SimpleCallData) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *SimpleCallData) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *SimpleCallData) GetDeliveryDuration() int32 {
	if x != nil {
		return x.DeliveryDuration
	}
	return 0
}

func (x *SimpleCallData) GetLinkCallDuration() int32 {
	if x != nil {
		return x.LinkCallDuration
	}
	return 0
}

func (x *SimpleCallData) GetResult() CallResult {
	if x != nil {
		return x.Result
	}
	return CallResult_CALL_RESULT_UNKNOWN
}

func (x *SimpleCallData) GetSipCode() int32 {
	if x != nil {
		return x.SipCode
	}
	return 0
}

func (x *SimpleCallData) GetDoRecord() bool {
	if x != nil {
		return x.DoRecord
	}
	return false
}

func (x *SimpleCallData) GetRecordingFileName() string {
	if x != nil {
		return x.RecordingFileName
	}
	return ""
}

func (x *SimpleCallData) GetIsDialValidationOk() bool {
	if x != nil {
		return x.IsDialValidationOk
	}
	return false
}

func (x *SimpleCallData) GetIsTimeZoneScrubOk() bool {
	if x != nil {
		return x.IsTimeZoneScrubOk
	}
	return false
}

func (x *SimpleCallData) GetIsCellPhoneScrubOk() bool {
	if x != nil {
		return x.IsCellPhoneScrubOk
	}
	return false
}

func (x *SimpleCallData) GetIsCustomCallingRulesScrubOk() bool {
	if x != nil {
		return x.IsCustomCallingRulesScrubOk
	}
	return false
}

func (x *SimpleCallData) GetIsDnclScrubOk() bool {
	if x != nil {
		return x.IsDnclScrubOk
	}
	return false
}

func (x *SimpleCallData) GetUseGlobalTimeZoneScrub() bool {
	if x != nil {
		return x.UseGlobalTimeZoneScrub
	}
	return false
}

func (x *SimpleCallData) GetDoCellPhoneScrub() bool {
	if x != nil {
		return x.DoCellPhoneScrub
	}
	return false
}

func (x *SimpleCallData) GetDoDnclScrub() bool {
	if x != nil {
		return x.DoDnclScrub
	}
	return false
}

func (x *SimpleCallData) GetCallDataType() string {
	if x != nil {
		return x.CallDataType
	}
	return ""
}

func (x *SimpleCallData) GetCallerIdCountryCode() string {
	if x != nil {
		return x.CallerIdCountryCode
	}
	return ""
}

func (x *SimpleCallData) GetCallerIdCountrySid() int64 {
	if x != nil {
		return x.CallerIdCountrySid
	}
	return 0
}

func (x *SimpleCallData) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *SimpleCallData) GetIsPreviewByRecord() bool {
	if x != nil {
		return x.IsPreviewByRecord
	}
	return false
}

func (x *SimpleCallData) GetRuleSetName() string {
	if x != nil {
		return x.RuleSetName
	}
	return ""
}

func (x *SimpleCallData) GetIsNaturalComplianceOk() bool {
	if x != nil {
		return x.IsNaturalComplianceOk
	}
	return false
}

func (x *SimpleCallData) GetSimpleMetaData() []*SimpleKeyValue {
	if x != nil {
		return x.SimpleMetaData
	}
	return nil
}

func (x *SimpleCallData) GetSimpleResultMetaData() []*SimpleKeyValue {
	if x != nil {
		return x.SimpleResultMetaData
	}
	return nil
}

type SimpleKeyValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SimpleKeyValue) Reset() {
	*x = SimpleKeyValue{}
	mi := &file_api_commons_call_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimpleKeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleKeyValue) ProtoMessage() {}

func (x *SimpleKeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_call_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleKeyValue.ProtoReflect.Descriptor instead.
func (*SimpleKeyValue) Descriptor() ([]byte, []int) {
	return file_api_commons_call_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleKeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SimpleKeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SimpleRecordData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskSid       int64                  `protobuf:"varint,1,opt,name=task_sid,json=taskSid,proto3" json:"task_sid,omitempty"`
	TaskGroupSid  int64                  `protobuf:"varint,2,opt,name=task_group_sid,json=taskGroupSid,proto3" json:"task_group_sid,omitempty"`
	AgentSid      int64                  `protobuf:"varint,3,opt,name=agent_sid,json=agentSid,proto3" json:"agent_sid,omitempty"`
	StartTime     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	StopTime      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=stop_time,json=stopTime,proto3" json:"stop_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SimpleRecordData) Reset() {
	*x = SimpleRecordData{}
	mi := &file_api_commons_call_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimpleRecordData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleRecordData) ProtoMessage() {}

func (x *SimpleRecordData) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_call_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleRecordData.ProtoReflect.Descriptor instead.
func (*SimpleRecordData) Descriptor() ([]byte, []int) {
	return file_api_commons_call_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleRecordData) GetTaskSid() int64 {
	if x != nil {
		return x.TaskSid
	}
	return 0
}

func (x *SimpleRecordData) GetTaskGroupSid() int64 {
	if x != nil {
		return x.TaskGroupSid
	}
	return 0
}

func (x *SimpleRecordData) GetAgentSid() int64 {
	if x != nil {
		return x.AgentSid
	}
	return 0
}

func (x *SimpleRecordData) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *SimpleRecordData) GetStopTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StopTime
	}
	return nil
}

var File_api_commons_call_proto protoreflect.FileDescriptor

var file_api_commons_call_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x61,
	0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb5, 0x0b, 0x0a, 0x0e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x12,
	0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x61,
	0x6c, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73,
	0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x6f, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x6b, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x44, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x6b, 0x12, 0x30, 0x0a, 0x15, 0x69, 0x73, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x5f, 0x6f, 0x6b, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e,
	0x65, 0x53, 0x63, 0x72, 0x75, 0x62, 0x4f, 0x6b, 0x12, 0x32, 0x0a, 0x16, 0x69, 0x73, 0x5f, 0x63,
	0x65, 0x6c, 0x6c, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x5f,
	0x6f, 0x6b, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x43, 0x65, 0x6c, 0x6c,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x53, 0x63, 0x72, 0x75, 0x62, 0x4f, 0x6b, 0x12, 0x45, 0x0a, 0x20,
	0x69, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x5f, 0x6f, 0x6b,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x69, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x43, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53, 0x63, 0x72, 0x75,
	0x62, 0x4f, 0x6b, 0x12, 0x27, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x64, 0x6e, 0x63, 0x6c, 0x5f, 0x73,
	0x63, 0x72, 0x75, 0x62, 0x5f, 0x6f, 0x6b, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69,
	0x73, 0x44, 0x6e, 0x63, 0x6c, 0x53, 0x63, 0x72, 0x75, 0x62, 0x4f, 0x6b, 0x12, 0x3a, 0x0a, 0x1a,
	0x75, 0x73, 0x65, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x16, 0x75, 0x73, 0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x5a,
	0x6f, 0x6e, 0x65, 0x53, 0x63, 0x72, 0x75, 0x62, 0x12, 0x2d, 0x0a, 0x13, 0x64, 0x6f, 0x5f, 0x63,
	0x65, 0x6c, 0x6c, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x6f, 0x43, 0x65, 0x6c, 0x6c, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x53, 0x63, 0x72, 0x75, 0x62, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x6f, 0x5f, 0x64, 0x6e,
	0x63, 0x6c, 0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x64, 0x6f, 0x44, 0x6e, 0x63, 0x6c, 0x53, 0x63, 0x72, 0x75, 0x62, 0x12, 0x24, 0x0a, 0x0e, 0x63,
	0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x33, 0x0a, 0x16, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x64, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x62, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x42, 0x79, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x69, 0x73, 0x5f,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x6f, 0x6b, 0x18, 0x21, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x73, 0x4e,
	0x61, 0x74, 0x75, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x4f, 0x6b, 0x12, 0x45, 0x0a, 0x10, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x22, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x52, 0x0a, 0x17, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x23, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x4a, 0x04, 0x08,
	0x18, 0x10, 0x19, 0x4a, 0x04, 0x08, 0x1b, 0x10, 0x1c, 0x22, 0x38, 0x0a, 0x0e, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xe4, 0x01, 0x0a, 0x10, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x53, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x61, 0x73,
	0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x5b, 0x0a, 0x0a, 0x43, 0x61,
	0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x41, 0x4c, 0x4c,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x41,
	0x4c, 0x4c, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x9c, 0x18, 0x12,
	0x11, 0x0a, 0x0c, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x80, 0x19, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x54, 0x45, 0x44, 0x10, 0xe4, 0x19, 0x42, 0x91, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x09, 0x43, 0x61, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_api_commons_call_proto_rawDescOnce sync.Once
	file_api_commons_call_proto_rawDescData []byte
)

func file_api_commons_call_proto_rawDescGZIP() []byte {
	file_api_commons_call_proto_rawDescOnce.Do(func() {
		file_api_commons_call_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_call_proto_rawDesc), len(file_api_commons_call_proto_rawDesc)))
	})
	return file_api_commons_call_proto_rawDescData
}

var file_api_commons_call_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_call_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_commons_call_proto_goTypes = []any{
	(CallStatus)(0),               // 0: api.commons.CallStatus
	(*SimpleCallData)(nil),        // 1: api.commons.SimpleCallData
	(*SimpleKeyValue)(nil),        // 2: api.commons.SimpleKeyValue
	(*SimpleRecordData)(nil),      // 3: api.commons.SimpleRecordData
	(CallResult)(0),               // 4: api.commons.CallResult
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_api_commons_call_proto_depIdxs = []int32{
	4, // 0: api.commons.SimpleCallData.result:type_name -> api.commons.CallResult
	2, // 1: api.commons.SimpleCallData.simple_meta_data:type_name -> api.commons.SimpleKeyValue
	2, // 2: api.commons.SimpleCallData.simple_result_meta_data:type_name -> api.commons.SimpleKeyValue
	5, // 3: api.commons.SimpleRecordData.start_time:type_name -> google.protobuf.Timestamp
	5, // 4: api.commons.SimpleRecordData.stop_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_commons_call_proto_init() }
func file_api_commons_call_proto_init() {
	if File_api_commons_call_proto != nil {
		return
	}
	file_api_commons_results_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_call_proto_rawDesc), len(file_api_commons_call_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_call_proto_goTypes,
		DependencyIndexes: file_api_commons_call_proto_depIdxs,
		EnumInfos:         file_api_commons_call_proto_enumTypes,
		MessageInfos:      file_api_commons_call_proto_msgTypes,
	}.Build()
	File_api_commons_call_proto = out.File
	file_api_commons_call_proto_goTypes = nil
	file_api_commons_call_proto_depIdxs = nil
}
