// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: annotations/perms/tcn.proto

package perms

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

// Tcn defines the enum value annotations which will be used for permissions.
type Tcn struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Work-in-progress tag. Used to delineate permissions that are not
	// yet ready to go live. Previously known as "A la carte" permissions
	// in the legacy customer support licensing tool.
	Wip *bool `protobuf:"varint,1,opt,name=wip" json:"wip,omitempty"`
	// Application which the permission will be grouped under.
	App *Application `protobuf:"varint,2,opt,name=app,enum=annotations.perms.Application" json:"app,omitempty"`
	// The card/sub-app section which it's assigned to.
	// If no card is provided, it will go under the application's default
	// card (under the same name as the app).
	Card *Card `protobuf:"varint,3,opt,name=card,enum=annotations.perms.Card" json:"card,omitempty"`
	// The list of features or effects the permission grants.
	Features      []string `protobuf:"bytes,4,rep,name=features" json:"features,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tcn) Reset() {
	*x = Tcn{}
	mi := &file_annotations_perms_tcn_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tcn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tcn) ProtoMessage() {}

func (x *Tcn) ProtoReflect() protoreflect.Message {
	mi := &file_annotations_perms_tcn_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tcn.ProtoReflect.Descriptor instead.
func (*Tcn) Descriptor() ([]byte, []int) {
	return file_annotations_perms_tcn_proto_rawDescGZIP(), []int{0}
}

func (x *Tcn) GetWip() bool {
	if x != nil && x.Wip != nil {
		return *x.Wip
	}
	return false
}

func (x *Tcn) GetApp() Application {
	if x != nil && x.App != nil {
		return *x.App
	}
	return Application_APPLICATION_UNSPECIFIED
}

func (x *Tcn) GetCard() Card {
	if x != nil && x.Card != nil {
		return *x.Card
	}
	return Card_CARD_UNSPECIFIED
}

func (x *Tcn) GetFeatures() []string {
	if x != nil {
		return x.Features
	}
	return nil
}

var file_annotations_perms_tcn_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*Tcn)(nil),
		Field:         50001,
		Name:          "annotations.perms.options",
		Tag:           "bytes,50001,opt,name=options",
		Filename:      "annotations/perms/tcn.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// Options contain the permission annotation properties.
	//
	// optional annotations.perms.Tcn options = 50001;
	E_Options = &file_annotations_perms_tcn_proto_extTypes[0]
)

var File_annotations_perms_tcn_proto protoreflect.FileDescriptor

var file_annotations_perms_tcn_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x73, 0x2f, 0x74, 0x63, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x73,
	0x1a, 0x1f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x73, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x03, 0x54, 0x63, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x77,
	0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x77, 0x69, 0x70, 0x12, 0x30, 0x0a,
	0x03, 0x61, 0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12,
	0x2b, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x3a, 0x55, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x73, 0x2e, 0x54, 0x63, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0xb4, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x42, 0x08, 0x54, 0x63, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x50, 0x58, 0xaa, 0x02, 0x11, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x73, 0xca, 0x02, 0x11,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x50, 0x65, 0x72, 0x6d,
	0x73, 0xe2, 0x02, 0x1d, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c,
	0x50, 0x65, 0x72, 0x6d, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x12, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a,
	0x3a, 0x50, 0x65, 0x72, 0x6d, 0x73,
})

var (
	file_annotations_perms_tcn_proto_rawDescOnce sync.Once
	file_annotations_perms_tcn_proto_rawDescData []byte
)

func file_annotations_perms_tcn_proto_rawDescGZIP() []byte {
	file_annotations_perms_tcn_proto_rawDescOnce.Do(func() {
		file_annotations_perms_tcn_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_annotations_perms_tcn_proto_rawDesc), len(file_annotations_perms_tcn_proto_rawDesc)))
	})
	return file_annotations_perms_tcn_proto_rawDescData
}

var file_annotations_perms_tcn_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_annotations_perms_tcn_proto_goTypes = []any{
	(*Tcn)(nil),                           // 0: annotations.perms.Tcn
	(Application)(0),                      // 1: annotations.perms.Application
	(Card)(0),                             // 2: annotations.perms.Card
	(*descriptorpb.EnumValueOptions)(nil), // 3: google.protobuf.EnumValueOptions
}
var file_annotations_perms_tcn_proto_depIdxs = []int32{
	1, // 0: annotations.perms.Tcn.app:type_name -> annotations.perms.Application
	2, // 1: annotations.perms.Tcn.card:type_name -> annotations.perms.Card
	3, // 2: annotations.perms.options:extendee -> google.protobuf.EnumValueOptions
	0, // 3: annotations.perms.options:type_name -> annotations.perms.Tcn
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_annotations_perms_tcn_proto_init() }
func file_annotations_perms_tcn_proto_init() {
	if File_annotations_perms_tcn_proto != nil {
		return
	}
	file_annotations_perms_license_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_annotations_perms_tcn_proto_rawDesc), len(file_annotations_perms_tcn_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_annotations_perms_tcn_proto_goTypes,
		DependencyIndexes: file_annotations_perms_tcn_proto_depIdxs,
		MessageInfos:      file_annotations_perms_tcn_proto_msgTypes,
		ExtensionInfos:    file_annotations_perms_tcn_proto_extTypes,
	}.Build()
	File_annotations_perms_tcn_proto = out.File
	file_annotations_perms_tcn_proto_goTypes = nil
	file_annotations_perms_tcn_proto_depIdxs = nil
}
