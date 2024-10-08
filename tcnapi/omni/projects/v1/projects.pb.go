// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: tcnapi/omni/projects/v1/projects.proto

package projectsv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This enum represents the state of a project
type Project_State int32

const (
	// Project state is Unknown
	Project_STATE_UNSPECIFIED Project_State = 0
	// Project state is Open
	Project_STATE_OPEN Project_State = 17000
	// Project state is Closed
	Project_STATE_CLOSED Project_State = 17010
)

// Enum value maps for Project_State.
var (
	Project_State_name = map[int32]string{
		0:     "STATE_UNSPECIFIED",
		17000: "STATE_OPEN",
		17010: "STATE_CLOSED",
	}
	Project_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_OPEN":        17000,
		"STATE_CLOSED":      17010,
	}
)

func (x Project_State) Enum() *Project_State {
	p := new(Project_State)
	*p = x
	return p
}

func (x Project_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Project_State) Descriptor() protoreflect.EnumDescriptor {
	return file_tcnapi_omni_projects_v1_projects_proto_enumTypes[0].Descriptor()
}

func (Project_State) Type() protoreflect.EnumType {
	return &file_tcnapi_omni_projects_v1_projects_proto_enumTypes[0]
}

func (x Project_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Project_State.Descriptor instead.
func (Project_State) EnumDescriptor() ([]byte, []int) {
	return file_tcnapi_omni_projects_v1_projects_proto_rawDescGZIP(), []int{0, 0}
}

// Project -
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project identifier
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// project title
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// project description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// project state
	State Project_State `protobuf:"varint,4,opt,name=state,proto3,enum=tcnapi.omni.projects.v1.Project_State" json:"state,omitempty"`
	// creation date
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	mi := &file_tcnapi_omni_projects_v1_projects_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_tcnapi_omni_projects_v1_projects_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_tcnapi_omni_projects_v1_projects_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Project) GetState() Project_State {
	if x != nil {
		return x.State
	}
	return Project_STATE_UNSPECIFIED
}

func (x *Project) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

var File_tcnapi_omni_projects_v1_projects_proto protoreflect.FileDescriptor

var file_tcnapi_omni_projects_v1_projects_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x63, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x63, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98,
	0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x25,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x74, 0x63, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x6d,
	0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0a, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0xe8, 0x84, 0x01, 0x12, 0x12, 0x0a, 0x0c,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0xf2, 0x84, 0x01,
	0x3a, 0x67, 0xea, 0x41, 0x64, 0x0a, 0x1f, 0x74, 0x63, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x6d,
	0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2f, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6f, 0x72, 0x67, 0x73, 0x2f, 0x7b, 0x6f,
	0x72, 0x67, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x32, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0xea, 0x01, 0x0a, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x63, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x63, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x6d,
	0x6e, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x4f, 0x50, 0xaa,
	0x02, 0x17, 0x54, 0x63, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x6d, 0x6e, 0x69, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x54, 0x63, 0x6e, 0x61,
	0x70, 0x69, 0x5c, 0x4f, 0x6d, 0x6e, 0x69, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x54, 0x63, 0x6e, 0x61, 0x70, 0x69, 0x5c, 0x4f, 0x6d, 0x6e,
	0x69, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x54, 0x63, 0x6e, 0x61,
	0x70, 0x69, 0x3a, 0x3a, 0x4f, 0x6d, 0x6e, 0x69, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tcnapi_omni_projects_v1_projects_proto_rawDescOnce sync.Once
	file_tcnapi_omni_projects_v1_projects_proto_rawDescData = file_tcnapi_omni_projects_v1_projects_proto_rawDesc
)

func file_tcnapi_omni_projects_v1_projects_proto_rawDescGZIP() []byte {
	file_tcnapi_omni_projects_v1_projects_proto_rawDescOnce.Do(func() {
		file_tcnapi_omni_projects_v1_projects_proto_rawDescData = protoimpl.X.CompressGZIP(file_tcnapi_omni_projects_v1_projects_proto_rawDescData)
	})
	return file_tcnapi_omni_projects_v1_projects_proto_rawDescData
}

var file_tcnapi_omni_projects_v1_projects_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tcnapi_omni_projects_v1_projects_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tcnapi_omni_projects_v1_projects_proto_goTypes = []any{
	(Project_State)(0),            // 0: tcnapi.omni.projects.v1.Project.State
	(*Project)(nil),               // 1: tcnapi.omni.projects.v1.Project
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_tcnapi_omni_projects_v1_projects_proto_depIdxs = []int32{
	0, // 0: tcnapi.omni.projects.v1.Project.state:type_name -> tcnapi.omni.projects.v1.Project.State
	2, // 1: tcnapi.omni.projects.v1.Project.create_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tcnapi_omni_projects_v1_projects_proto_init() }
func file_tcnapi_omni_projects_v1_projects_proto_init() {
	if File_tcnapi_omni_projects_v1_projects_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tcnapi_omni_projects_v1_projects_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tcnapi_omni_projects_v1_projects_proto_goTypes,
		DependencyIndexes: file_tcnapi_omni_projects_v1_projects_proto_depIdxs,
		EnumInfos:         file_tcnapi_omni_projects_v1_projects_proto_enumTypes,
		MessageInfos:      file_tcnapi_omni_projects_v1_projects_proto_msgTypes,
	}.Build()
	File_tcnapi_omni_projects_v1_projects_proto = out.File
	file_tcnapi_omni_projects_v1_projects_proto_rawDesc = nil
	file_tcnapi_omni_projects_v1_projects_proto_goTypes = nil
	file_tcnapi_omni_projects_v1_projects_proto_depIdxs = nil
}
