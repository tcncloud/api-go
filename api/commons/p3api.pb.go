// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/commons/p3api.proto

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

// States for interrupting peering
type InterruptedPeeringStatus int32

const (
	InterruptedPeeringStatus_InterruptPeeringStatus_MANUAL  InterruptedPeeringStatus = 0
	InterruptedPeeringStatus_InterruptPeeringStatus_PREVIEW InterruptedPeeringStatus = 1
)

// Enum value maps for InterruptedPeeringStatus.
var (
	InterruptedPeeringStatus_name = map[int32]string{
		0: "InterruptPeeringStatus_MANUAL",
		1: "InterruptPeeringStatus_PREVIEW",
	}
	InterruptedPeeringStatus_value = map[string]int32{
		"InterruptPeeringStatus_MANUAL":  0,
		"InterruptPeeringStatus_PREVIEW": 1,
	}
)

func (x InterruptedPeeringStatus) Enum() *InterruptedPeeringStatus {
	p := new(InterruptedPeeringStatus)
	*p = x
	return p
}

func (x InterruptedPeeringStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InterruptedPeeringStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_p3api_proto_enumTypes[0].Descriptor()
}

func (InterruptedPeeringStatus) Type() protoreflect.EnumType {
	return &file_api_commons_p3api_proto_enumTypes[0]
}

func (x InterruptedPeeringStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InterruptedPeeringStatus.Descriptor instead.
func (InterruptedPeeringStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{0}
}

type RecalculateBillingMonth int32

const (
	RecalculateBillingMonth_RECALCULATE_BILLING_MONTH_CURRENT  RecalculateBillingMonth = 0
	RecalculateBillingMonth_RECALCULATE_BILLING_MONTH_PREVIOUS RecalculateBillingMonth = 1
)

// Enum value maps for RecalculateBillingMonth.
var (
	RecalculateBillingMonth_name = map[int32]string{
		0: "RECALCULATE_BILLING_MONTH_CURRENT",
		1: "RECALCULATE_BILLING_MONTH_PREVIOUS",
	}
	RecalculateBillingMonth_value = map[string]int32{
		"RECALCULATE_BILLING_MONTH_CURRENT":  0,
		"RECALCULATE_BILLING_MONTH_PREVIOUS": 1,
	}
)

func (x RecalculateBillingMonth) Enum() *RecalculateBillingMonth {
	p := new(RecalculateBillingMonth)
	*p = x
	return p
}

func (x RecalculateBillingMonth) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecalculateBillingMonth) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_p3api_proto_enumTypes[1].Descriptor()
}

func (RecalculateBillingMonth) Type() protoreflect.EnumType {
	return &file_api_commons_p3api_proto_enumTypes[1]
}

func (x RecalculateBillingMonth) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecalculateBillingMonth.Descriptor instead.
func (RecalculateBillingMonth) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{1}
}

type RecalculateBillingType int32

const (
	RecalculateBillingType_RECALCULATE_BILLING_TYPE_OUTBOUND_CALLS    RecalculateBillingType = 0
	RecalculateBillingType_RECALCULATE_BILLING_TYPE_INBOUND_CALLS     RecalculateBillingType = 1
	RecalculateBillingType_RECALCULATE_BILLING_TYPE_AGENTS            RecalculateBillingType = 2
	RecalculateBillingType_RECALCULATE_BILLING_TYPE_SMS               RecalculateBillingType = 3
	RecalculateBillingType_RECALCULATE_BILLING_TYPE_EMAIL             RecalculateBillingType = 4
	RecalculateBillingType_RECALCULATE_BILLING_TYPE_MANUAL_DIAL_CALLS RecalculateBillingType = 6
)

// Enum value maps for RecalculateBillingType.
var (
	RecalculateBillingType_name = map[int32]string{
		0: "RECALCULATE_BILLING_TYPE_OUTBOUND_CALLS",
		1: "RECALCULATE_BILLING_TYPE_INBOUND_CALLS",
		2: "RECALCULATE_BILLING_TYPE_AGENTS",
		3: "RECALCULATE_BILLING_TYPE_SMS",
		4: "RECALCULATE_BILLING_TYPE_EMAIL",
		6: "RECALCULATE_BILLING_TYPE_MANUAL_DIAL_CALLS",
	}
	RecalculateBillingType_value = map[string]int32{
		"RECALCULATE_BILLING_TYPE_OUTBOUND_CALLS":    0,
		"RECALCULATE_BILLING_TYPE_INBOUND_CALLS":     1,
		"RECALCULATE_BILLING_TYPE_AGENTS":            2,
		"RECALCULATE_BILLING_TYPE_SMS":               3,
		"RECALCULATE_BILLING_TYPE_EMAIL":             4,
		"RECALCULATE_BILLING_TYPE_MANUAL_DIAL_CALLS": 6,
	}
)

func (x RecalculateBillingType) Enum() *RecalculateBillingType {
	p := new(RecalculateBillingType)
	*p = x
	return p
}

func (x RecalculateBillingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecalculateBillingType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_p3api_proto_enumTypes[2].Descriptor()
}

func (RecalculateBillingType) Type() protoreflect.EnumType {
	return &file_api_commons_p3api_proto_enumTypes[2]
}

func (x RecalculateBillingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecalculateBillingType.Descriptor instead.
func (RecalculateBillingType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{2}
}

type PhoneBookSIPType int32

const (
	PhoneBookSIPType_PHONE_BOOK_SIP_TYPE_OUTBOUND PhoneBookSIPType = 0
	PhoneBookSIPType_PHONE_BOOK_SIP_TYPE_TRANSFER PhoneBookSIPType = 1
)

// Enum value maps for PhoneBookSIPType.
var (
	PhoneBookSIPType_name = map[int32]string{
		0: "PHONE_BOOK_SIP_TYPE_OUTBOUND",
		1: "PHONE_BOOK_SIP_TYPE_TRANSFER",
	}
	PhoneBookSIPType_value = map[string]int32{
		"PHONE_BOOK_SIP_TYPE_OUTBOUND": 0,
		"PHONE_BOOK_SIP_TYPE_TRANSFER": 1,
	}
)

func (x PhoneBookSIPType) Enum() *PhoneBookSIPType {
	p := new(PhoneBookSIPType)
	*p = x
	return p
}

func (x PhoneBookSIPType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhoneBookSIPType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_p3api_proto_enumTypes[3].Descriptor()
}

func (PhoneBookSIPType) Type() protoreflect.EnumType {
	return &file_api_commons_p3api_proto_enumTypes[3]
}

func (x PhoneBookSIPType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhoneBookSIPType.Descriptor instead.
func (PhoneBookSIPType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{3}
}

type PhoneBookPhoneNumberType int32

const (
	PhoneBookPhoneNumberType_PHONE_BOOK_PHONE_NUMBER_TYPE_CALLER_ID PhoneBookPhoneNumberType = 0
	PhoneBookPhoneNumberType_PHONE_BOOK_PHONE_NUMBER_TYPE_OUTBOUND  PhoneBookPhoneNumberType = 1
	PhoneBookPhoneNumberType_PHONE_BOOK_PHONE_NUMBER_TYPE_TRANSFER  PhoneBookPhoneNumberType = 2
)

// Enum value maps for PhoneBookPhoneNumberType.
var (
	PhoneBookPhoneNumberType_name = map[int32]string{
		0: "PHONE_BOOK_PHONE_NUMBER_TYPE_CALLER_ID",
		1: "PHONE_BOOK_PHONE_NUMBER_TYPE_OUTBOUND",
		2: "PHONE_BOOK_PHONE_NUMBER_TYPE_TRANSFER",
	}
	PhoneBookPhoneNumberType_value = map[string]int32{
		"PHONE_BOOK_PHONE_NUMBER_TYPE_CALLER_ID": 0,
		"PHONE_BOOK_PHONE_NUMBER_TYPE_OUTBOUND":  1,
		"PHONE_BOOK_PHONE_NUMBER_TYPE_TRANSFER":  2,
	}
)

func (x PhoneBookPhoneNumberType) Enum() *PhoneBookPhoneNumberType {
	p := new(PhoneBookPhoneNumberType)
	*p = x
	return p
}

func (x PhoneBookPhoneNumberType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhoneBookPhoneNumberType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_p3api_proto_enumTypes[4].Descriptor()
}

func (PhoneBookPhoneNumberType) Type() protoreflect.EnumType {
	return &file_api_commons_p3api_proto_enumTypes[4]
}

func (x PhoneBookPhoneNumberType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhoneBookPhoneNumberType.Descriptor instead.
func (PhoneBookPhoneNumberType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{4}
}

type CallHistorySearchType_Enum int32

const (
	CallHistorySearchType_STANDARD CallHistorySearchType_Enum = 0
	CallHistorySearchType_BY_ID    CallHistorySearchType_Enum = 1
	CallHistorySearchType_BY_AGENT CallHistorySearchType_Enum = 2
)

// Enum value maps for CallHistorySearchType_Enum.
var (
	CallHistorySearchType_Enum_name = map[int32]string{
		0: "STANDARD",
		1: "BY_ID",
		2: "BY_AGENT",
	}
	CallHistorySearchType_Enum_value = map[string]int32{
		"STANDARD": 0,
		"BY_ID":    1,
		"BY_AGENT": 2,
	}
)

func (x CallHistorySearchType_Enum) Enum() *CallHistorySearchType_Enum {
	p := new(CallHistorySearchType_Enum)
	*p = x
	return p
}

func (x CallHistorySearchType_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CallHistorySearchType_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_p3api_proto_enumTypes[5].Descriptor()
}

func (CallHistorySearchType_Enum) Type() protoreflect.EnumType {
	return &file_api_commons_p3api_proto_enumTypes[5]
}

func (x CallHistorySearchType_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CallHistorySearchType_Enum.Descriptor instead.
func (CallHistorySearchType_Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{4, 0}
}

type CallHistorySearchScope_Enum int32

const (
	CallHistorySearchScope_ALL                    CallHistorySearchScope_Enum = 0
	CallHistorySearchScope_TWENTY_FOUR_HOURS      CallHistorySearchScope_Enum = 1
	CallHistorySearchScope_TWO_DAYS               CallHistorySearchScope_Enum = 2
	CallHistorySearchScope_THREE_DAYS             CallHistorySearchScope_Enum = 3
	CallHistorySearchScope_FIVE_DAYS              CallHistorySearchScope_Enum = 4
	CallHistorySearchScope_SEVEN_DAYS             CallHistorySearchScope_Enum = 5
	CallHistorySearchScope_THIRTY_DAYS            CallHistorySearchScope_Enum = 6
	CallHistorySearchScope_SIXTY_DAYS             CallHistorySearchScope_Enum = 7
	CallHistorySearchScope_NINETY_DAYS            CallHistorySearchScope_Enum = 8
	CallHistorySearchScope_ONEHUNDRED_TWENTY_DAYS CallHistorySearchScope_Enum = 9
	CallHistorySearchScope_ONEHUNDRED_FIFTY_DAYS  CallHistorySearchScope_Enum = 10
	CallHistorySearchScope_ONEHUNDRED_EIGHTY_DAYS CallHistorySearchScope_Enum = 11
)

// Enum value maps for CallHistorySearchScope_Enum.
var (
	CallHistorySearchScope_Enum_name = map[int32]string{
		0:  "ALL",
		1:  "TWENTY_FOUR_HOURS",
		2:  "TWO_DAYS",
		3:  "THREE_DAYS",
		4:  "FIVE_DAYS",
		5:  "SEVEN_DAYS",
		6:  "THIRTY_DAYS",
		7:  "SIXTY_DAYS",
		8:  "NINETY_DAYS",
		9:  "ONEHUNDRED_TWENTY_DAYS",
		10: "ONEHUNDRED_FIFTY_DAYS",
		11: "ONEHUNDRED_EIGHTY_DAYS",
	}
	CallHistorySearchScope_Enum_value = map[string]int32{
		"ALL":                    0,
		"TWENTY_FOUR_HOURS":      1,
		"TWO_DAYS":               2,
		"THREE_DAYS":             3,
		"FIVE_DAYS":              4,
		"SEVEN_DAYS":             5,
		"THIRTY_DAYS":            6,
		"SIXTY_DAYS":             7,
		"NINETY_DAYS":            8,
		"ONEHUNDRED_TWENTY_DAYS": 9,
		"ONEHUNDRED_FIFTY_DAYS":  10,
		"ONEHUNDRED_EIGHTY_DAYS": 11,
	}
)

func (x CallHistorySearchScope_Enum) Enum() *CallHistorySearchScope_Enum {
	p := new(CallHistorySearchScope_Enum)
	*p = x
	return p
}

func (x CallHistorySearchScope_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CallHistorySearchScope_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_p3api_proto_enumTypes[6].Descriptor()
}

func (CallHistorySearchScope_Enum) Type() protoreflect.EnumType {
	return &file_api_commons_p3api_proto_enumTypes[6]
}

func (x CallHistorySearchScope_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CallHistorySearchScope_Enum.Descriptor instead.
func (CallHistorySearchScope_Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{5, 0}
}

type ListPhoneBookOrderBy_Enum int32

const (
	ListPhoneBookOrderBy_PHONE_BOOK_SID      ListPhoneBookOrderBy_Enum = 0
	ListPhoneBookOrderBy_ENTRY_TYPE          ListPhoneBookOrderBy_Enum = 1
	ListPhoneBookOrderBy_ENTRY_NAME          ListPhoneBookOrderBy_Enum = 2
	ListPhoneBookOrderBy_CLIENT_SID          ListPhoneBookOrderBy_Enum = 3
	ListPhoneBookOrderBy_HUNT_GROUP_SID      ListPhoneBookOrderBy_Enum = 4
	ListPhoneBookOrderBy_PHONE_NUMBER        ListPhoneBookOrderBy_Enum = 5
	ListPhoneBookOrderBy_PHONE_NUMBER_TYPE   ListPhoneBookOrderBy_Enum = 6
	ListPhoneBookOrderBy_PHONE_NUMBER_HIDDEN ListPhoneBookOrderBy_Enum = 7
)

// Enum value maps for ListPhoneBookOrderBy_Enum.
var (
	ListPhoneBookOrderBy_Enum_name = map[int32]string{
		0: "PHONE_BOOK_SID",
		1: "ENTRY_TYPE",
		2: "ENTRY_NAME",
		3: "CLIENT_SID",
		4: "HUNT_GROUP_SID",
		5: "PHONE_NUMBER",
		6: "PHONE_NUMBER_TYPE",
		7: "PHONE_NUMBER_HIDDEN",
	}
	ListPhoneBookOrderBy_Enum_value = map[string]int32{
		"PHONE_BOOK_SID":      0,
		"ENTRY_TYPE":          1,
		"ENTRY_NAME":          2,
		"CLIENT_SID":          3,
		"HUNT_GROUP_SID":      4,
		"PHONE_NUMBER":        5,
		"PHONE_NUMBER_TYPE":   6,
		"PHONE_NUMBER_HIDDEN": 7,
	}
)

func (x ListPhoneBookOrderBy_Enum) Enum() *ListPhoneBookOrderBy_Enum {
	p := new(ListPhoneBookOrderBy_Enum)
	*p = x
	return p
}

func (x ListPhoneBookOrderBy_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListPhoneBookOrderBy_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_p3api_proto_enumTypes[7].Descriptor()
}

func (ListPhoneBookOrderBy_Enum) Type() protoreflect.EnumType {
	return &file_api_commons_p3api_proto_enumTypes[7]
}

func (x ListPhoneBookOrderBy_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListPhoneBookOrderBy_Enum.Descriptor instead.
func (ListPhoneBookOrderBy_Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{6, 0}
}

type ClientInfoDataRow struct {
	state                      protoimpl.MessageState `protogen:"open.v1"`
	FieldLabel                 string                 `protobuf:"bytes,1,opt,name=field_label,json=fieldLabel,proto3" json:"field_label,omitempty"`
	FieldValue                 string                 `protobuf:"bytes,2,opt,name=field_value,json=fieldValue,proto3" json:"field_value,omitempty"`
	IsPhone                    bool                   `protobuf:"varint,3,opt,name=is_phone,json=isPhone,proto3" json:"is_phone,omitempty"`
	DialedNumber               bool                   `protobuf:"varint,4,opt,name=dialed_number,json=dialedNumber,proto3" json:"dialed_number,omitempty"`
	ContactFieldDescriptionSid int64                  `protobuf:"varint,5,opt,name=contact_field_description_sid,json=contactFieldDescriptionSid,proto3" json:"contact_field_description_sid,omitempty"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *ClientInfoDataRow) Reset() {
	*x = ClientInfoDataRow{}
	mi := &file_api_commons_p3api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientInfoDataRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfoDataRow) ProtoMessage() {}

func (x *ClientInfoDataRow) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_p3api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfoDataRow.ProtoReflect.Descriptor instead.
func (*ClientInfoDataRow) Descriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{0}
}

func (x *ClientInfoDataRow) GetFieldLabel() string {
	if x != nil {
		return x.FieldLabel
	}
	return ""
}

func (x *ClientInfoDataRow) GetFieldValue() string {
	if x != nil {
		return x.FieldValue
	}
	return ""
}

func (x *ClientInfoDataRow) GetIsPhone() bool {
	if x != nil {
		return x.IsPhone
	}
	return false
}

func (x *ClientInfoDataRow) GetDialedNumber() bool {
	if x != nil {
		return x.DialedNumber
	}
	return false
}

func (x *ClientInfoDataRow) GetContactFieldDescriptionSid() int64 {
	if x != nil {
		return x.ContactFieldDescriptionSid
	}
	return 0
}

type RGBColor struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Red           int32                  `protobuf:"varint,1,opt,name=red,proto3" json:"red,omitempty"`
	Green         int32                  `protobuf:"varint,2,opt,name=green,proto3" json:"green,omitempty"`
	Blue          int32                  `protobuf:"varint,3,opt,name=blue,proto3" json:"blue,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RGBColor) Reset() {
	*x = RGBColor{}
	mi := &file_api_commons_p3api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RGBColor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RGBColor) ProtoMessage() {}

func (x *RGBColor) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_p3api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RGBColor.ProtoReflect.Descriptor instead.
func (*RGBColor) Descriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{1}
}

func (x *RGBColor) GetRed() int32 {
	if x != nil {
		return x.Red
	}
	return 0
}

func (x *RGBColor) GetGreen() int32 {
	if x != nil {
		return x.Green
	}
	return 0
}

func (x *RGBColor) GetBlue() int32 {
	if x != nil {
		return x.Blue
	}
	return 0
}

// Dialed Number Field Settings
type DialedNumberFieldSettings struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Should the dialed number be displayed?
	ShouldDisplay bool `protobuf:"varint,1,opt,name=should_display,json=shouldDisplay,proto3" json:"should_display,omitempty"`
	// Color of the text of the dialed number
	Color *RGBColor `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	// Color of the background of the dialed number
	BgColor *RGBColor `protobuf:"bytes,3,opt,name=bg_color,json=bgColor,proto3" json:"bg_color,omitempty"`
	// Shows a copy button in the row that copies the field value to your clipboard
	AllowAgentCopy bool `protobuf:"varint,4,opt,name=allow_agent_copy,json=allowAgentCopy,proto3" json:"allow_agent_copy,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DialedNumberFieldSettings) Reset() {
	*x = DialedNumberFieldSettings{}
	mi := &file_api_commons_p3api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DialedNumberFieldSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DialedNumberFieldSettings) ProtoMessage() {}

func (x *DialedNumberFieldSettings) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_p3api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DialedNumberFieldSettings.ProtoReflect.Descriptor instead.
func (*DialedNumberFieldSettings) Descriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{2}
}

func (x *DialedNumberFieldSettings) GetShouldDisplay() bool {
	if x != nil {
		return x.ShouldDisplay
	}
	return false
}

func (x *DialedNumberFieldSettings) GetColor() *RGBColor {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *DialedNumberFieldSettings) GetBgColor() *RGBColor {
	if x != nil {
		return x.BgColor
	}
	return nil
}

func (x *DialedNumberFieldSettings) GetAllowAgentCopy() bool {
	if x != nil {
		return x.AllowAgentCopy
	}
	return false
}

// ClientInfoDisplayTemplateRow
type ClientInfoDisplayTemplateRow struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Field label for the row
	FieldLabel string `protobuf:"bytes,1,opt,name=field_label,json=fieldLabel,proto3" json:"field_label,omitempty"`
	// Color of the text in the row
	Color *RGBColor `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	// Backgorund color in the row
	BgColor *RGBColor `protobuf:"bytes,3,opt,name=bg_color,json=bgColor,proto3" json:"bg_color,omitempty"`
	// Sid that relates the row to a contact field description
	ContactFieldDescriptionSid int64 `protobuf:"varint,4,opt,name=contact_field_description_sid,json=contactFieldDescriptionSid,proto3" json:"contact_field_description_sid,omitempty"`
	// Shows a copy button in the row that copies the field value to your clipboard
	AllowAgentCopy bool `protobuf:"varint,5,opt,name=allow_agent_copy,json=allowAgentCopy,proto3" json:"allow_agent_copy,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ClientInfoDisplayTemplateRow) Reset() {
	*x = ClientInfoDisplayTemplateRow{}
	mi := &file_api_commons_p3api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientInfoDisplayTemplateRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfoDisplayTemplateRow) ProtoMessage() {}

func (x *ClientInfoDisplayTemplateRow) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_p3api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfoDisplayTemplateRow.ProtoReflect.Descriptor instead.
func (*ClientInfoDisplayTemplateRow) Descriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{3}
}

func (x *ClientInfoDisplayTemplateRow) GetFieldLabel() string {
	if x != nil {
		return x.FieldLabel
	}
	return ""
}

func (x *ClientInfoDisplayTemplateRow) GetColor() *RGBColor {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *ClientInfoDisplayTemplateRow) GetBgColor() *RGBColor {
	if x != nil {
		return x.BgColor
	}
	return nil
}

func (x *ClientInfoDisplayTemplateRow) GetContactFieldDescriptionSid() int64 {
	if x != nil {
		return x.ContactFieldDescriptionSid
	}
	return 0
}

func (x *ClientInfoDisplayTemplateRow) GetAllowAgentCopy() bool {
	if x != nil {
		return x.AllowAgentCopy
	}
	return false
}

type CallHistorySearchType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CallHistorySearchType) Reset() {
	*x = CallHistorySearchType{}
	mi := &file_api_commons_p3api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallHistorySearchType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallHistorySearchType) ProtoMessage() {}

func (x *CallHistorySearchType) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_p3api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallHistorySearchType.ProtoReflect.Descriptor instead.
func (*CallHistorySearchType) Descriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{4}
}

type CallHistorySearchScope struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CallHistorySearchScope) Reset() {
	*x = CallHistorySearchScope{}
	mi := &file_api_commons_p3api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallHistorySearchScope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallHistorySearchScope) ProtoMessage() {}

func (x *CallHistorySearchScope) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_p3api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallHistorySearchScope.ProtoReflect.Descriptor instead.
func (*CallHistorySearchScope) Descriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{5}
}

type ListPhoneBookOrderBy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPhoneBookOrderBy) Reset() {
	*x = ListPhoneBookOrderBy{}
	mi := &file_api_commons_p3api_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPhoneBookOrderBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPhoneBookOrderBy) ProtoMessage() {}

func (x *ListPhoneBookOrderBy) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_p3api_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPhoneBookOrderBy.ProtoReflect.Descriptor instead.
func (*ListPhoneBookOrderBy) Descriptor() ([]byte, []int) {
	return file_api_commons_p3api_proto_rawDescGZIP(), []int{6}
}

var File_api_commons_p3api_proto protoreflect.FileDescriptor

const file_api_commons_p3api_proto_rawDesc = "" +
	"\n" +
	"\x17api/commons/p3api.proto\x12\vapi.commons\"\xd8\x01\n" +
	"\x11ClientInfoDataRow\x12\x1f\n" +
	"\vfield_label\x18\x01 \x01(\tR\n" +
	"fieldLabel\x12\x1f\n" +
	"\vfield_value\x18\x02 \x01(\tR\n" +
	"fieldValue\x12\x19\n" +
	"\bis_phone\x18\x03 \x01(\bR\aisPhone\x12#\n" +
	"\rdialed_number\x18\x04 \x01(\bR\fdialedNumber\x12A\n" +
	"\x1dcontact_field_description_sid\x18\x05 \x01(\x03R\x1acontactFieldDescriptionSid\"F\n" +
	"\bRGBColor\x12\x10\n" +
	"\x03red\x18\x01 \x01(\x05R\x03red\x12\x14\n" +
	"\x05green\x18\x02 \x01(\x05R\x05green\x12\x12\n" +
	"\x04blue\x18\x03 \x01(\x05R\x04blue\"\xcb\x01\n" +
	"\x19DialedNumberFieldSettings\x12%\n" +
	"\x0eshould_display\x18\x01 \x01(\bR\rshouldDisplay\x12+\n" +
	"\x05color\x18\x02 \x01(\v2\x15.api.commons.RGBColorR\x05color\x120\n" +
	"\bbg_color\x18\x03 \x01(\v2\x15.api.commons.RGBColorR\abgColor\x12(\n" +
	"\x10allow_agent_copy\x18\x04 \x01(\bR\x0eallowAgentCopy\"\x8b\x02\n" +
	"\x1cClientInfoDisplayTemplateRow\x12\x1f\n" +
	"\vfield_label\x18\x01 \x01(\tR\n" +
	"fieldLabel\x12+\n" +
	"\x05color\x18\x02 \x01(\v2\x15.api.commons.RGBColorR\x05color\x120\n" +
	"\bbg_color\x18\x03 \x01(\v2\x15.api.commons.RGBColorR\abgColor\x12A\n" +
	"\x1dcontact_field_description_sid\x18\x04 \x01(\x03R\x1acontactFieldDescriptionSid\x12(\n" +
	"\x10allow_agent_copy\x18\x05 \x01(\bR\x0eallowAgentCopy\"F\n" +
	"\x15CallHistorySearchType\"-\n" +
	"\x04Enum\x12\f\n" +
	"\bSTANDARD\x10\x00\x12\t\n" +
	"\x05BY_ID\x10\x01\x12\f\n" +
	"\bBY_AGENT\x10\x02\"\x83\x02\n" +
	"\x16CallHistorySearchScope\"\xe8\x01\n" +
	"\x04Enum\x12\a\n" +
	"\x03ALL\x10\x00\x12\x15\n" +
	"\x11TWENTY_FOUR_HOURS\x10\x01\x12\f\n" +
	"\bTWO_DAYS\x10\x02\x12\x0e\n" +
	"\n" +
	"THREE_DAYS\x10\x03\x12\r\n" +
	"\tFIVE_DAYS\x10\x04\x12\x0e\n" +
	"\n" +
	"SEVEN_DAYS\x10\x05\x12\x0f\n" +
	"\vTHIRTY_DAYS\x10\x06\x12\x0e\n" +
	"\n" +
	"SIXTY_DAYS\x10\a\x12\x0f\n" +
	"\vNINETY_DAYS\x10\b\x12\x1a\n" +
	"\x16ONEHUNDRED_TWENTY_DAYS\x10\t\x12\x19\n" +
	"\x15ONEHUNDRED_FIFTY_DAYS\x10\n" +
	"\x12\x1a\n" +
	"\x16ONEHUNDRED_EIGHTY_DAYS\x10\v\"\xb9\x01\n" +
	"\x14ListPhoneBookOrderBy\"\xa0\x01\n" +
	"\x04Enum\x12\x12\n" +
	"\x0ePHONE_BOOK_SID\x10\x00\x12\x0e\n" +
	"\n" +
	"ENTRY_TYPE\x10\x01\x12\x0e\n" +
	"\n" +
	"ENTRY_NAME\x10\x02\x12\x0e\n" +
	"\n" +
	"CLIENT_SID\x10\x03\x12\x12\n" +
	"\x0eHUNT_GROUP_SID\x10\x04\x12\x10\n" +
	"\fPHONE_NUMBER\x10\x05\x12\x15\n" +
	"\x11PHONE_NUMBER_TYPE\x10\x06\x12\x17\n" +
	"\x13PHONE_NUMBER_HIDDEN\x10\a*a\n" +
	"\x18InterruptedPeeringStatus\x12!\n" +
	"\x1dInterruptPeeringStatus_MANUAL\x10\x00\x12\"\n" +
	"\x1eInterruptPeeringStatus_PREVIEW\x10\x01*h\n" +
	"\x17RecalculateBillingMonth\x12%\n" +
	"!RECALCULATE_BILLING_MONTH_CURRENT\x10\x00\x12&\n" +
	"\"RECALCULATE_BILLING_MONTH_PREVIOUS\x10\x01*\xb9\x02\n" +
	"\x16RecalculateBillingType\x12+\n" +
	"'RECALCULATE_BILLING_TYPE_OUTBOUND_CALLS\x10\x00\x12*\n" +
	"&RECALCULATE_BILLING_TYPE_INBOUND_CALLS\x10\x01\x12#\n" +
	"\x1fRECALCULATE_BILLING_TYPE_AGENTS\x10\x02\x12 \n" +
	"\x1cRECALCULATE_BILLING_TYPE_SMS\x10\x03\x12\"\n" +
	"\x1eRECALCULATE_BILLING_TYPE_EMAIL\x10\x04\x12.\n" +
	"*RECALCULATE_BILLING_TYPE_MANUAL_DIAL_CALLS\x10\x06\"\x04\b\x05\x10\x05*%RECALCULATE_BILLING_TYPE_VOCAL_DIRECT*V\n" +
	"\x10PhoneBookSIPType\x12 \n" +
	"\x1cPHONE_BOOK_SIP_TYPE_OUTBOUND\x10\x00\x12 \n" +
	"\x1cPHONE_BOOK_SIP_TYPE_TRANSFER\x10\x01*\x9c\x01\n" +
	"\x18PhoneBookPhoneNumberType\x12*\n" +
	"&PHONE_BOOK_PHONE_NUMBER_TYPE_CALLER_ID\x10\x00\x12)\n" +
	"%PHONE_BOOK_PHONE_NUMBER_TYPE_OUTBOUND\x10\x01\x12)\n" +
	"%PHONE_BOOK_PHONE_NUMBER_TYPE_TRANSFER\x10\x02B\x92\x01\n" +
	"\x0fcom.api.commonsB\n" +
	"P3apiProtoP\x01Z&github.com/tcncloud/api-go/api/commons\xa2\x02\x03ACX\xaa\x02\vApi.Commons\xca\x02\vApi\\Commons\xe2\x02\x17Api\\Commons\\GPBMetadata\xea\x02\fApi::Commonsb\x06proto3"

var (
	file_api_commons_p3api_proto_rawDescOnce sync.Once
	file_api_commons_p3api_proto_rawDescData []byte
)

func file_api_commons_p3api_proto_rawDescGZIP() []byte {
	file_api_commons_p3api_proto_rawDescOnce.Do(func() {
		file_api_commons_p3api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_p3api_proto_rawDesc), len(file_api_commons_p3api_proto_rawDesc)))
	})
	return file_api_commons_p3api_proto_rawDescData
}

var file_api_commons_p3api_proto_enumTypes = make([]protoimpl.EnumInfo, 8)
var file_api_commons_p3api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_commons_p3api_proto_goTypes = []any{
	(InterruptedPeeringStatus)(0),        // 0: api.commons.InterruptedPeeringStatus
	(RecalculateBillingMonth)(0),         // 1: api.commons.RecalculateBillingMonth
	(RecalculateBillingType)(0),          // 2: api.commons.RecalculateBillingType
	(PhoneBookSIPType)(0),                // 3: api.commons.PhoneBookSIPType
	(PhoneBookPhoneNumberType)(0),        // 4: api.commons.PhoneBookPhoneNumberType
	(CallHistorySearchType_Enum)(0),      // 5: api.commons.CallHistorySearchType.Enum
	(CallHistorySearchScope_Enum)(0),     // 6: api.commons.CallHistorySearchScope.Enum
	(ListPhoneBookOrderBy_Enum)(0),       // 7: api.commons.ListPhoneBookOrderBy.Enum
	(*ClientInfoDataRow)(nil),            // 8: api.commons.ClientInfoDataRow
	(*RGBColor)(nil),                     // 9: api.commons.RGBColor
	(*DialedNumberFieldSettings)(nil),    // 10: api.commons.DialedNumberFieldSettings
	(*ClientInfoDisplayTemplateRow)(nil), // 11: api.commons.ClientInfoDisplayTemplateRow
	(*CallHistorySearchType)(nil),        // 12: api.commons.CallHistorySearchType
	(*CallHistorySearchScope)(nil),       // 13: api.commons.CallHistorySearchScope
	(*ListPhoneBookOrderBy)(nil),         // 14: api.commons.ListPhoneBookOrderBy
}
var file_api_commons_p3api_proto_depIdxs = []int32{
	9, // 0: api.commons.DialedNumberFieldSettings.color:type_name -> api.commons.RGBColor
	9, // 1: api.commons.DialedNumberFieldSettings.bg_color:type_name -> api.commons.RGBColor
	9, // 2: api.commons.ClientInfoDisplayTemplateRow.color:type_name -> api.commons.RGBColor
	9, // 3: api.commons.ClientInfoDisplayTemplateRow.bg_color:type_name -> api.commons.RGBColor
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_commons_p3api_proto_init() }
func file_api_commons_p3api_proto_init() {
	if File_api_commons_p3api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_p3api_proto_rawDesc), len(file_api_commons_p3api_proto_rawDesc)),
			NumEnums:      8,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_p3api_proto_goTypes,
		DependencyIndexes: file_api_commons_p3api_proto_depIdxs,
		EnumInfos:         file_api_commons_p3api_proto_enumTypes,
		MessageInfos:      file_api_commons_p3api_proto_msgTypes,
	}.Build()
	File_api_commons_p3api_proto = out.File
	file_api_commons_p3api_proto_goTypes = nil
	file_api_commons_p3api_proto_depIdxs = nil
}
