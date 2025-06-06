// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/commons/chat.proto

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

type ChatMessageSenderType int32

const (
	ChatMessageSenderType_CHAT_MESSAGE_SENDER_TYPE_AGENT    ChatMessageSenderType = 0
	ChatMessageSenderType_CHAT_MESSAGE_SENDER_TYPE_CUSTOMER ChatMessageSenderType = 1
	ChatMessageSenderType_CHAT_MESSAGE_SENDER_TYPE_MANAGER  ChatMessageSenderType = 2
)

// Enum value maps for ChatMessageSenderType.
var (
	ChatMessageSenderType_name = map[int32]string{
		0: "CHAT_MESSAGE_SENDER_TYPE_AGENT",
		1: "CHAT_MESSAGE_SENDER_TYPE_CUSTOMER",
		2: "CHAT_MESSAGE_SENDER_TYPE_MANAGER",
	}
	ChatMessageSenderType_value = map[string]int32{
		"CHAT_MESSAGE_SENDER_TYPE_AGENT":    0,
		"CHAT_MESSAGE_SENDER_TYPE_CUSTOMER": 1,
		"CHAT_MESSAGE_SENDER_TYPE_MANAGER":  2,
	}
)

func (x ChatMessageSenderType) Enum() *ChatMessageSenderType {
	p := new(ChatMessageSenderType)
	*p = x
	return p
}

func (x ChatMessageSenderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatMessageSenderType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_chat_proto_enumTypes[0].Descriptor()
}

func (ChatMessageSenderType) Type() protoreflect.EnumType {
	return &file_api_commons_chat_proto_enumTypes[0]
}

func (x ChatMessageSenderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatMessageSenderType.Descriptor instead.
func (ChatMessageSenderType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_chat_proto_rawDescGZIP(), []int{0}
}

// The status for a chat campaign
type ChatCampaignStatus int32

const (
	ChatCampaignStatus_CHAT_CAMPAIGN_UNKNOWN  ChatCampaignStatus = 0
	ChatCampaignStatus_CHAT_CAMPAIGN_RUNNING  ChatCampaignStatus = 15010 //Chat campaign started running
	ChatCampaignStatus_CHAT_CAMPAIGN_STOPPED  ChatCampaignStatus = 15050 //Chat campaign has been executed
	ChatCampaignStatus_CHAT_CAMPAIGN_ARCHIVED ChatCampaignStatus = 15060 //Chat campaign has been archived
	ChatCampaignStatus_CHAT_CAMPAIGN_PAUSED   ChatCampaignStatus = 15070 //Chat campaign has been paused
)

// Enum value maps for ChatCampaignStatus.
var (
	ChatCampaignStatus_name = map[int32]string{
		0:     "CHAT_CAMPAIGN_UNKNOWN",
		15010: "CHAT_CAMPAIGN_RUNNING",
		15050: "CHAT_CAMPAIGN_STOPPED",
		15060: "CHAT_CAMPAIGN_ARCHIVED",
		15070: "CHAT_CAMPAIGN_PAUSED",
	}
	ChatCampaignStatus_value = map[string]int32{
		"CHAT_CAMPAIGN_UNKNOWN":  0,
		"CHAT_CAMPAIGN_RUNNING":  15010,
		"CHAT_CAMPAIGN_STOPPED":  15050,
		"CHAT_CAMPAIGN_ARCHIVED": 15060,
		"CHAT_CAMPAIGN_PAUSED":   15070,
	}
)

func (x ChatCampaignStatus) Enum() *ChatCampaignStatus {
	p := new(ChatCampaignStatus)
	*p = x
	return p
}

func (x ChatCampaignStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatCampaignStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_chat_proto_enumTypes[1].Descriptor()
}

func (ChatCampaignStatus) Type() protoreflect.EnumType {
	return &file_api_commons_chat_proto_enumTypes[1]
}

func (x ChatCampaignStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatCampaignStatus.Descriptor instead.
func (ChatCampaignStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_chat_proto_rawDescGZIP(), []int{1}
}

// The type of events for chat campaigns
type ChatCampaignEvent int32

const (
	ChatCampaignEvent_CHAT_CAMPAIGN_EVENT_UNKNOWN   ChatCampaignEvent = 0
	ChatCampaignEvent_CHAT_CAMPAIGN_EVENT_SCHEDULED ChatCampaignEvent = 1 // when the chat campaign is scheduled
	ChatCampaignEvent_CHAT_CAMPAIGN_EVENT_STOPPED   ChatCampaignEvent = 2 // when the manager stops the chat campaign
	ChatCampaignEvent_CHAT_CAMPAIGN_EVENT_ARCHIVED  ChatCampaignEvent = 3 // when the manager archives the chat campaign
	ChatCampaignEvent_CHAT_CAMPAIGN_EVENT_PAUSED    ChatCampaignEvent = 4 // when the chat campaign paused by manager
)

// Enum value maps for ChatCampaignEvent.
var (
	ChatCampaignEvent_name = map[int32]string{
		0: "CHAT_CAMPAIGN_EVENT_UNKNOWN",
		1: "CHAT_CAMPAIGN_EVENT_SCHEDULED",
		2: "CHAT_CAMPAIGN_EVENT_STOPPED",
		3: "CHAT_CAMPAIGN_EVENT_ARCHIVED",
		4: "CHAT_CAMPAIGN_EVENT_PAUSED",
	}
	ChatCampaignEvent_value = map[string]int32{
		"CHAT_CAMPAIGN_EVENT_UNKNOWN":   0,
		"CHAT_CAMPAIGN_EVENT_SCHEDULED": 1,
		"CHAT_CAMPAIGN_EVENT_STOPPED":   2,
		"CHAT_CAMPAIGN_EVENT_ARCHIVED":  3,
		"CHAT_CAMPAIGN_EVENT_PAUSED":    4,
	}
)

func (x ChatCampaignEvent) Enum() *ChatCampaignEvent {
	p := new(ChatCampaignEvent)
	*p = x
	return p
}

func (x ChatCampaignEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatCampaignEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_chat_proto_enumTypes[2].Descriptor()
}

func (ChatCampaignEvent) Type() protoreflect.EnumType {
	return &file_api_commons_chat_proto_enumTypes[2]
}

func (x ChatCampaignEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatCampaignEvent.Descriptor instead.
func (ChatCampaignEvent) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_chat_proto_rawDescGZIP(), []int{2}
}

type ChatMessageType int32

const (
	ChatMessageType_CHAT_REPLY_FROM_CUSTOMER ChatMessageType = 0 // Received chat message reply from customer
	ChatMessageType_CHAT_REPLY_FROM_AGENT    ChatMessageType = 1 // Its a chat message reply by an agent
)

// Enum value maps for ChatMessageType.
var (
	ChatMessageType_name = map[int32]string{
		0: "CHAT_REPLY_FROM_CUSTOMER",
		1: "CHAT_REPLY_FROM_AGENT",
	}
	ChatMessageType_value = map[string]int32{
		"CHAT_REPLY_FROM_CUSTOMER": 0,
		"CHAT_REPLY_FROM_AGENT":    1,
	}
)

func (x ChatMessageType) Enum() *ChatMessageType {
	p := new(ChatMessageType)
	*p = x
	return p
}

func (x ChatMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_chat_proto_enumTypes[3].Descriptor()
}

func (ChatMessageType) Type() protoreflect.EnumType {
	return &file_api_commons_chat_proto_enumTypes[3]
}

func (x ChatMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatMessageType.Descriptor instead.
func (ChatMessageType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_chat_proto_rawDescGZIP(), []int{3}
}

type ChatColorProperties struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	PrimaryColor       string                 `protobuf:"bytes,1,opt,name=primary_color,json=primaryColor,proto3" json:"primary_color,omitempty"`
	HeaderTextColor    string                 `protobuf:"bytes,2,opt,name=header_text_color,json=headerTextColor,proto3" json:"header_text_color,omitempty"`
	ParagraphTextColor string                 `protobuf:"bytes,3,opt,name=paragraph_text_color,json=paragraphTextColor,proto3" json:"paragraph_text_color,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ChatColorProperties) Reset() {
	*x = ChatColorProperties{}
	mi := &file_api_commons_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatColorProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatColorProperties) ProtoMessage() {}

func (x *ChatColorProperties) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatColorProperties.ProtoReflect.Descriptor instead.
func (*ChatColorProperties) Descriptor() ([]byte, []int) {
	return file_api_commons_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatColorProperties) GetPrimaryColor() string {
	if x != nil {
		return x.PrimaryColor
	}
	return ""
}

func (x *ChatColorProperties) GetHeaderTextColor() string {
	if x != nil {
		return x.HeaderTextColor
	}
	return ""
}

func (x *ChatColorProperties) GetParagraphTextColor() string {
	if x != nil {
		return x.ParagraphTextColor
	}
	return ""
}

// The header data for the chat widget associated to a given chat campaign
type ChatHeader struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the main text to display
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// the text to display beneath the header
	Subheader     string `protobuf:"bytes,2,opt,name=subheader,proto3" json:"subheader,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatHeader) Reset() {
	*x = ChatHeader{}
	mi := &file_api_commons_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatHeader) ProtoMessage() {}

func (x *ChatHeader) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatHeader.ProtoReflect.Descriptor instead.
func (*ChatHeader) Descriptor() ([]byte, []int) {
	return file_api_commons_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ChatHeader) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *ChatHeader) GetSubheader() string {
	if x != nil {
		return x.Subheader
	}
	return ""
}

type HoursOfOperation struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Monday        []*HoursOfOperationRange `protobuf:"bytes,1,rep,name=monday,proto3" json:"monday,omitempty"`
	Tuesday       []*HoursOfOperationRange `protobuf:"bytes,2,rep,name=tuesday,proto3" json:"tuesday,omitempty"`
	Wednesday     []*HoursOfOperationRange `protobuf:"bytes,3,rep,name=wednesday,proto3" json:"wednesday,omitempty"`
	Thursday      []*HoursOfOperationRange `protobuf:"bytes,4,rep,name=thursday,proto3" json:"thursday,omitempty"`
	Friday        []*HoursOfOperationRange `protobuf:"bytes,5,rep,name=friday,proto3" json:"friday,omitempty"`
	Saturday      []*HoursOfOperationRange `protobuf:"bytes,6,rep,name=saturday,proto3" json:"saturday,omitempty"`
	Sunday        []*HoursOfOperationRange `protobuf:"bytes,7,rep,name=sunday,proto3" json:"sunday,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HoursOfOperation) Reset() {
	*x = HoursOfOperation{}
	mi := &file_api_commons_chat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HoursOfOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HoursOfOperation) ProtoMessage() {}

func (x *HoursOfOperation) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_chat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HoursOfOperation.ProtoReflect.Descriptor instead.
func (*HoursOfOperation) Descriptor() ([]byte, []int) {
	return file_api_commons_chat_proto_rawDescGZIP(), []int{2}
}

func (x *HoursOfOperation) GetMonday() []*HoursOfOperationRange {
	if x != nil {
		return x.Monday
	}
	return nil
}

func (x *HoursOfOperation) GetTuesday() []*HoursOfOperationRange {
	if x != nil {
		return x.Tuesday
	}
	return nil
}

func (x *HoursOfOperation) GetWednesday() []*HoursOfOperationRange {
	if x != nil {
		return x.Wednesday
	}
	return nil
}

func (x *HoursOfOperation) GetThursday() []*HoursOfOperationRange {
	if x != nil {
		return x.Thursday
	}
	return nil
}

func (x *HoursOfOperation) GetFriday() []*HoursOfOperationRange {
	if x != nil {
		return x.Friday
	}
	return nil
}

func (x *HoursOfOperation) GetSaturday() []*HoursOfOperationRange {
	if x != nil {
		return x.Saturday
	}
	return nil
}

func (x *HoursOfOperation) GetSunday() []*HoursOfOperationRange {
	if x != nil {
		return x.Sunday
	}
	return nil
}

type HoursOfOperationRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StartHour     int64                  `protobuf:"varint,1,opt,name=start_hour,json=startHour,proto3" json:"start_hour,omitempty"`
	StartMinute   int64                  `protobuf:"varint,2,opt,name=start_minute,json=startMinute,proto3" json:"start_minute,omitempty"`
	EndHour       int64                  `protobuf:"varint,3,opt,name=end_hour,json=endHour,proto3" json:"end_hour,omitempty"`
	EndMinute     int64                  `protobuf:"varint,4,opt,name=end_minute,json=endMinute,proto3" json:"end_minute,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HoursOfOperationRange) Reset() {
	*x = HoursOfOperationRange{}
	mi := &file_api_commons_chat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HoursOfOperationRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HoursOfOperationRange) ProtoMessage() {}

func (x *HoursOfOperationRange) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_chat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HoursOfOperationRange.ProtoReflect.Descriptor instead.
func (*HoursOfOperationRange) Descriptor() ([]byte, []int) {
	return file_api_commons_chat_proto_rawDescGZIP(), []int{3}
}

func (x *HoursOfOperationRange) GetStartHour() int64 {
	if x != nil {
		return x.StartHour
	}
	return 0
}

func (x *HoursOfOperationRange) GetStartMinute() int64 {
	if x != nil {
		return x.StartMinute
	}
	return 0
}

func (x *HoursOfOperationRange) GetEndHour() int64 {
	if x != nil {
		return x.EndHour
	}
	return 0
}

func (x *HoursOfOperationRange) GetEndMinute() int64 {
	if x != nil {
		return x.EndMinute
	}
	return 0
}

var File_api_commons_chat_proto protoreflect.FileDescriptor

const file_api_commons_chat_proto_rawDesc = "" +
	"\n" +
	"\x16api/commons/chat.proto\x12\vapi.commons\"\x98\x01\n" +
	"\x13ChatColorProperties\x12#\n" +
	"\rprimary_color\x18\x01 \x01(\tR\fprimaryColor\x12*\n" +
	"\x11header_text_color\x18\x02 \x01(\tR\x0fheaderTextColor\x120\n" +
	"\x14paragraph_text_color\x18\x03 \x01(\tR\x12paragraphTextColor\"B\n" +
	"\n" +
	"ChatHeader\x12\x16\n" +
	"\x06header\x18\x01 \x01(\tR\x06header\x12\x1c\n" +
	"\tsubheader\x18\x02 \x01(\tR\tsubheader\"\xc6\x03\n" +
	"\x10HoursOfOperation\x12:\n" +
	"\x06monday\x18\x01 \x03(\v2\".api.commons.HoursOfOperationRangeR\x06monday\x12<\n" +
	"\atuesday\x18\x02 \x03(\v2\".api.commons.HoursOfOperationRangeR\atuesday\x12@\n" +
	"\twednesday\x18\x03 \x03(\v2\".api.commons.HoursOfOperationRangeR\twednesday\x12>\n" +
	"\bthursday\x18\x04 \x03(\v2\".api.commons.HoursOfOperationRangeR\bthursday\x12:\n" +
	"\x06friday\x18\x05 \x03(\v2\".api.commons.HoursOfOperationRangeR\x06friday\x12>\n" +
	"\bsaturday\x18\x06 \x03(\v2\".api.commons.HoursOfOperationRangeR\bsaturday\x12:\n" +
	"\x06sunday\x18\a \x03(\v2\".api.commons.HoursOfOperationRangeR\x06sunday\"\x93\x01\n" +
	"\x15HoursOfOperationRange\x12\x1d\n" +
	"\n" +
	"start_hour\x18\x01 \x01(\x03R\tstartHour\x12!\n" +
	"\fstart_minute\x18\x02 \x01(\x03R\vstartMinute\x12\x19\n" +
	"\bend_hour\x18\x03 \x01(\x03R\aendHour\x12\x1d\n" +
	"\n" +
	"end_minute\x18\x04 \x01(\x03R\tendMinute*\x88\x01\n" +
	"\x15ChatMessageSenderType\x12\"\n" +
	"\x1eCHAT_MESSAGE_SENDER_TYPE_AGENT\x10\x00\x12%\n" +
	"!CHAT_MESSAGE_SENDER_TYPE_CUSTOMER\x10\x01\x12$\n" +
	" CHAT_MESSAGE_SENDER_TYPE_MANAGER\x10\x02*\x9f\x01\n" +
	"\x12ChatCampaignStatus\x12\x19\n" +
	"\x15CHAT_CAMPAIGN_UNKNOWN\x10\x00\x12\x1a\n" +
	"\x15CHAT_CAMPAIGN_RUNNING\x10\xa2u\x12\x1a\n" +
	"\x15CHAT_CAMPAIGN_STOPPED\x10\xcau\x12\x1b\n" +
	"\x16CHAT_CAMPAIGN_ARCHIVED\x10\xd4u\x12\x19\n" +
	"\x14CHAT_CAMPAIGN_PAUSED\x10\xdeu*\xba\x01\n" +
	"\x11ChatCampaignEvent\x12\x1f\n" +
	"\x1bCHAT_CAMPAIGN_EVENT_UNKNOWN\x10\x00\x12!\n" +
	"\x1dCHAT_CAMPAIGN_EVENT_SCHEDULED\x10\x01\x12\x1f\n" +
	"\x1bCHAT_CAMPAIGN_EVENT_STOPPED\x10\x02\x12 \n" +
	"\x1cCHAT_CAMPAIGN_EVENT_ARCHIVED\x10\x03\x12\x1e\n" +
	"\x1aCHAT_CAMPAIGN_EVENT_PAUSED\x10\x04*J\n" +
	"\x0fChatMessageType\x12\x1c\n" +
	"\x18CHAT_REPLY_FROM_CUSTOMER\x10\x00\x12\x19\n" +
	"\x15CHAT_REPLY_FROM_AGENT\x10\x01B\x91\x01\n" +
	"\x0fcom.api.commonsB\tChatProtoP\x01Z&github.com/tcncloud/api-go/api/commons\xa2\x02\x03ACX\xaa\x02\vApi.Commons\xca\x02\vApi\\Commons\xe2\x02\x17Api\\Commons\\GPBMetadata\xea\x02\fApi::Commonsb\x06proto3"

var (
	file_api_commons_chat_proto_rawDescOnce sync.Once
	file_api_commons_chat_proto_rawDescData []byte
)

func file_api_commons_chat_proto_rawDescGZIP() []byte {
	file_api_commons_chat_proto_rawDescOnce.Do(func() {
		file_api_commons_chat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_chat_proto_rawDesc), len(file_api_commons_chat_proto_rawDesc)))
	})
	return file_api_commons_chat_proto_rawDescData
}

var file_api_commons_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_api_commons_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_commons_chat_proto_goTypes = []any{
	(ChatMessageSenderType)(0),    // 0: api.commons.ChatMessageSenderType
	(ChatCampaignStatus)(0),       // 1: api.commons.ChatCampaignStatus
	(ChatCampaignEvent)(0),        // 2: api.commons.ChatCampaignEvent
	(ChatMessageType)(0),          // 3: api.commons.ChatMessageType
	(*ChatColorProperties)(nil),   // 4: api.commons.ChatColorProperties
	(*ChatHeader)(nil),            // 5: api.commons.ChatHeader
	(*HoursOfOperation)(nil),      // 6: api.commons.HoursOfOperation
	(*HoursOfOperationRange)(nil), // 7: api.commons.HoursOfOperationRange
}
var file_api_commons_chat_proto_depIdxs = []int32{
	7, // 0: api.commons.HoursOfOperation.monday:type_name -> api.commons.HoursOfOperationRange
	7, // 1: api.commons.HoursOfOperation.tuesday:type_name -> api.commons.HoursOfOperationRange
	7, // 2: api.commons.HoursOfOperation.wednesday:type_name -> api.commons.HoursOfOperationRange
	7, // 3: api.commons.HoursOfOperation.thursday:type_name -> api.commons.HoursOfOperationRange
	7, // 4: api.commons.HoursOfOperation.friday:type_name -> api.commons.HoursOfOperationRange
	7, // 5: api.commons.HoursOfOperation.saturday:type_name -> api.commons.HoursOfOperationRange
	7, // 6: api.commons.HoursOfOperation.sunday:type_name -> api.commons.HoursOfOperationRange
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_commons_chat_proto_init() }
func file_api_commons_chat_proto_init() {
	if File_api_commons_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_chat_proto_rawDesc), len(file_api_commons_chat_proto_rawDesc)),
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_chat_proto_goTypes,
		DependencyIndexes: file_api_commons_chat_proto_depIdxs,
		EnumInfos:         file_api_commons_chat_proto_enumTypes,
		MessageInfos:      file_api_commons_chat_proto_msgTypes,
	}.Build()
	File_api_commons_chat_proto = out.File
	file_api_commons_chat_proto_goTypes = nil
	file_api_commons_chat_proto_depIdxs = nil
}
