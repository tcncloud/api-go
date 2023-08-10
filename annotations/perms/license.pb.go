// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: annotations/perms/license.proto

package perms

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

// Application contains license/permission application name constants.
type Application int32

const (
	Application_APPLICATION_UNSPECIFIED                 Application = 0
	Application_APPLICATION_AGENT                       Application = 1
	Application_APPLICATION_BUSINESS_INTELLIGENCE       Application = 2
	Application_APPLICATION_DELIVERY_SERVICE            Application = 3
	Application_APPLICATION_DEV_TOOLS                   Application = 4
	Application_APPLICATION_INTEGRATIONS                Application = 5
	Application_APPLICATION_LIST_MANAGEMENT_SERVICES    Application = 6
	Application_APPLICATION_NATURAL_LANGUAGE_COMPLIANCE Application = 7
	Application_APPLICATION_OMNI_BOSS                   Application = 8
	Application_APPLICATION_ORGANIZATION                Application = 9
	Application_APPLICATION_ROOM_303                    Application = 10
	Application_APPLICATION_SCORECARDS                  Application = 11
	Application_APPLICATION_SCRIPTS                     Application = 12
	Application_APPLICATION_TICKETS                     Application = 13
	Application_APPLICATION_VOICE_ANALYTICS             Application = 14
	Application_APPLICATION_WORK_FORCE_MANAGEMENT       Application = 15
	Application_APPLICATION_WORKFLOWS                   Application = 16
	Application_APPLICATION_NEWSROOM                    Application = 17
)

// Enum value maps for Application.
var (
	Application_name = map[int32]string{
		0:  "APPLICATION_UNSPECIFIED",
		1:  "APPLICATION_AGENT",
		2:  "APPLICATION_BUSINESS_INTELLIGENCE",
		3:  "APPLICATION_DELIVERY_SERVICE",
		4:  "APPLICATION_DEV_TOOLS",
		5:  "APPLICATION_INTEGRATIONS",
		6:  "APPLICATION_LIST_MANAGEMENT_SERVICES",
		7:  "APPLICATION_NATURAL_LANGUAGE_COMPLIANCE",
		8:  "APPLICATION_OMNI_BOSS",
		9:  "APPLICATION_ORGANIZATION",
		10: "APPLICATION_ROOM_303",
		11: "APPLICATION_SCORECARDS",
		12: "APPLICATION_SCRIPTS",
		13: "APPLICATION_TICKETS",
		14: "APPLICATION_VOICE_ANALYTICS",
		15: "APPLICATION_WORK_FORCE_MANAGEMENT",
		16: "APPLICATION_WORKFLOWS",
		17: "APPLICATION_NEWSROOM",
	}
	Application_value = map[string]int32{
		"APPLICATION_UNSPECIFIED":                 0,
		"APPLICATION_AGENT":                       1,
		"APPLICATION_BUSINESS_INTELLIGENCE":       2,
		"APPLICATION_DELIVERY_SERVICE":            3,
		"APPLICATION_DEV_TOOLS":                   4,
		"APPLICATION_INTEGRATIONS":                5,
		"APPLICATION_LIST_MANAGEMENT_SERVICES":    6,
		"APPLICATION_NATURAL_LANGUAGE_COMPLIANCE": 7,
		"APPLICATION_OMNI_BOSS":                   8,
		"APPLICATION_ORGANIZATION":                9,
		"APPLICATION_ROOM_303":                    10,
		"APPLICATION_SCORECARDS":                  11,
		"APPLICATION_SCRIPTS":                     12,
		"APPLICATION_TICKETS":                     13,
		"APPLICATION_VOICE_ANALYTICS":             14,
		"APPLICATION_WORK_FORCE_MANAGEMENT":       15,
		"APPLICATION_WORKFLOWS":                   16,
		"APPLICATION_NEWSROOM":                    17,
	}
)

func (x Application) Enum() *Application {
	p := new(Application)
	*p = x
	return p
}

func (x Application) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Application) Descriptor() protoreflect.EnumDescriptor {
	return file_annotations_perms_license_proto_enumTypes[0].Descriptor()
}

func (Application) Type() protoreflect.EnumType {
	return &file_annotations_perms_license_proto_enumTypes[0]
}

func (x Application) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Application.Descriptor instead.
func (Application) EnumDescriptor() ([]byte, []int) {
	return file_annotations_perms_license_proto_rawDescGZIP(), []int{0}
}

// Card contains license/permission card label constants.
type Card int32

const (
	Card_CARD_UNSPECIFIED Card = 0
	// Organization cards.
	Card_CARD_ORGANIZATION      Card = 1
	Card_CARD_USERS             Card = 2
	Card_CARD_AGENTS            Card = 3
	Card_CARD_PERMISSION_GROUPS Card = 4
	Card_CARD_LABELS            Card = 5
	Card_CARD_TRUSTS            Card = 6
	Card_CARD_HUNT_GROUPS       Card = 7
	Card_CARD_SOUNDBOARD        Card = 8
	Card_CARD_SUBSCRIPTIONS     Card = 9
	Card_CARD_PBX_MANAGER       Card = 10
)

// Enum value maps for Card.
var (
	Card_name = map[int32]string{
		0:  "CARD_UNSPECIFIED",
		1:  "CARD_ORGANIZATION",
		2:  "CARD_USERS",
		3:  "CARD_AGENTS",
		4:  "CARD_PERMISSION_GROUPS",
		5:  "CARD_LABELS",
		6:  "CARD_TRUSTS",
		7:  "CARD_HUNT_GROUPS",
		8:  "CARD_SOUNDBOARD",
		9:  "CARD_SUBSCRIPTIONS",
		10: "CARD_PBX_MANAGER",
	}
	Card_value = map[string]int32{
		"CARD_UNSPECIFIED":       0,
		"CARD_ORGANIZATION":      1,
		"CARD_USERS":             2,
		"CARD_AGENTS":            3,
		"CARD_PERMISSION_GROUPS": 4,
		"CARD_LABELS":            5,
		"CARD_TRUSTS":            6,
		"CARD_HUNT_GROUPS":       7,
		"CARD_SOUNDBOARD":        8,
		"CARD_SUBSCRIPTIONS":     9,
		"CARD_PBX_MANAGER":       10,
	}
)

func (x Card) Enum() *Card {
	p := new(Card)
	*p = x
	return p
}

func (x Card) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Card) Descriptor() protoreflect.EnumDescriptor {
	return file_annotations_perms_license_proto_enumTypes[1].Descriptor()
}

func (Card) Type() protoreflect.EnumType {
	return &file_annotations_perms_license_proto_enumTypes[1]
}

func (x Card) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Card.Descriptor instead.
func (Card) EnumDescriptor() ([]byte, []int) {
	return file_annotations_perms_license_proto_rawDescGZIP(), []int{1}
}

var File_annotations_perms_license_proto protoreflect.FileDescriptor

var file_annotations_perms_license_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x73, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x73, 0x2a, 0xb8, 0x04, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x50, 0x50, 0x4c,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45, 0x53, 0x53,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x4c, 0x4c, 0x49, 0x47, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12,
	0x20, 0x0a, 0x1c, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10,
	0x03, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x44, 0x45, 0x56, 0x5f, 0x54, 0x4f, 0x4f, 0x4c, 0x53, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18,
	0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x47, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x05, 0x12, 0x28, 0x0a, 0x24, 0x41, 0x50,
	0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x4d,
	0x41, 0x4e, 0x41, 0x47, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x53, 0x10, 0x06, 0x12, 0x2b, 0x0a, 0x27, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x41, 0x4c, 0x5f, 0x4c, 0x41, 0x4e, 0x47,
	0x55, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x10,
	0x07, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4f, 0x4d, 0x4e, 0x49, 0x5f, 0x42, 0x4f, 0x53, 0x53, 0x10, 0x08, 0x12, 0x1c, 0x0a, 0x18,
	0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x52, 0x47, 0x41,
	0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x50,
	0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x33,
	0x30, 0x33, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x43, 0x41, 0x52, 0x44, 0x53, 0x10, 0x0b,
	0x12, 0x17, 0x0a, 0x13, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x53, 0x10, 0x0c, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x50, 0x50,
	0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x53,
	0x10, 0x0d, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43,
	0x53, 0x10, 0x0e, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x4d, 0x41,
	0x4e, 0x41, 0x47, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x0f, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x50,
	0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c,
	0x4f, 0x57, 0x53, 0x10, 0x10, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x45, 0x57, 0x53, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x11, 0x2a,
	0xeb, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x52, 0x44,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15,
	0x0a, 0x11, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x53, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x41, 0x47,
	0x45, 0x4e, 0x54, 0x53, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53,
	0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c,
	0x53, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x52, 0x55, 0x53,
	0x54, 0x53, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x48, 0x55, 0x4e,
	0x54, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x41,
	0x52, 0x44, 0x5f, 0x53, 0x4f, 0x55, 0x4e, 0x44, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x08, 0x12,
	0x16, 0x0a, 0x12, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x52, 0x44, 0x5f,
	0x50, 0x42, 0x58, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x0a, 0x42, 0xb8, 0x01,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x42, 0x0c, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x50, 0x58, 0xaa, 0x02, 0x11, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x73, 0xca,
	0x02, 0x11, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x50, 0x65,
	0x72, 0x6d, 0x73, 0xe2, 0x02, 0x1d, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x5c, 0x50, 0x65, 0x72, 0x6d, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x3a, 0x3a, 0x50, 0x65, 0x72, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_annotations_perms_license_proto_rawDescOnce sync.Once
	file_annotations_perms_license_proto_rawDescData = file_annotations_perms_license_proto_rawDesc
)

func file_annotations_perms_license_proto_rawDescGZIP() []byte {
	file_annotations_perms_license_proto_rawDescOnce.Do(func() {
		file_annotations_perms_license_proto_rawDescData = protoimpl.X.CompressGZIP(file_annotations_perms_license_proto_rawDescData)
	})
	return file_annotations_perms_license_proto_rawDescData
}

var file_annotations_perms_license_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_annotations_perms_license_proto_goTypes = []interface{}{
	(Application)(0), // 0: annotations.perms.Application
	(Card)(0),        // 1: annotations.perms.Card
}
var file_annotations_perms_license_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_annotations_perms_license_proto_init() }
func file_annotations_perms_license_proto_init() {
	if File_annotations_perms_license_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_annotations_perms_license_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_annotations_perms_license_proto_goTypes,
		DependencyIndexes: file_annotations_perms_license_proto_depIdxs,
		EnumInfos:         file_annotations_perms_license_proto_enumTypes,
	}.Build()
	File_annotations_perms_license_proto = out.File
	file_annotations_perms_license_proto_rawDesc = nil
	file_annotations_perms_license_proto_goTypes = nil
	file_annotations_perms_license_proto_depIdxs = nil
}
