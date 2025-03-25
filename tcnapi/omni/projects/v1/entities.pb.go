// (-- api-linter: core::0133::request-id-field=disabled
//     aip.dev/not-precedent: We want all ids to get generated
//     by the system and not by the user. --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: tcnapi/omni/projects/v1/entities.proto

package projectsv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// request used to get a list of projects
type ListProjectsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the parent of the listed projects
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// returned page
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// used to specify the page token
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// used to specify the filter
	Filter        string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProjectsRequest) Reset() {
	*x = ListProjectsRequest{}
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsRequest) ProtoMessage() {}

func (x *ListProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsRequest.ProtoReflect.Descriptor instead.
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return file_tcnapi_omni_projects_v1_entities_proto_rawDescGZIP(), []int{0}
}

func (x *ListProjectsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListProjectsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListProjectsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListProjectsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// response used listing all projects
type ListProjectsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the list of projects
	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	// the next page token
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProjectsResponse) Reset() {
	*x = ListProjectsResponse{}
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsResponse) ProtoMessage() {}

func (x *ListProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsResponse.ProtoReflect.Descriptor instead.
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return file_tcnapi_omni_projects_v1_entities_proto_rawDescGZIP(), []int{1}
}

func (x *ListProjectsResponse) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *ListProjectsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// get single project based on name and mask used to getting particular data
type GetProjectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the name of the project
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProjectRequest) Reset() {
	*x = GetProjectRequest{}
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectRequest) ProtoMessage() {}

func (x *GetProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectRequest.ProtoReflect.Descriptor instead.
func (*GetProjectRequest) Descriptor() ([]byte, []int) {
	return file_tcnapi_omni_projects_v1_entities_proto_rawDescGZIP(), []int{2}
}

func (x *GetProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// request used to create a project
type CreateProjectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the parent of the project
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// the project to be created
	Project       *Project `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProjectRequest) Reset() {
	*x = CreateProjectRequest{}
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRequest) ProtoMessage() {}

func (x *CreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_tcnapi_omni_projects_v1_entities_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProjectRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateProjectRequest) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

// request used to update a project
type UpdateProjectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the project to be updated
	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// the update mask
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProjectRequest) Reset() {
	*x = UpdateProjectRequest{}
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectRequest) ProtoMessage() {}

func (x *UpdateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectRequest.ProtoReflect.Descriptor instead.
func (*UpdateProjectRequest) Descriptor() ([]byte, []int) {
	return file_tcnapi_omni_projects_v1_entities_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProjectRequest) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *UpdateProjectRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// request used to delete a project
type DeleteProjectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the name of the project
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProjectRequest) Reset() {
	*x = DeleteProjectRequest{}
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProjectRequest) ProtoMessage() {}

func (x *DeleteProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tcnapi_omni_projects_v1_entities_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) {
	return file_tcnapi_omni_projects_v1_entities_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_tcnapi_omni_projects_v1_entities_proto protoreflect.FileDescriptor

const file_tcnapi_omni_projects_v1_entities_proto_rawDesc = "" +
	"\n" +
	"&tcnapi/omni/projects/v1/entities.proto\x12\x17tcnapi.omni.projects.v1\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a google/protobuf/field_mask.proto\x1a&tcnapi/omni/projects/v1/projects.proto\"\xb9\x01\n" +
	"\x13ListProjectsRequest\x12?\n" +
	"\x06parent\x18\x01 \x01(\tB'\xe0A\x02\xfaA!\x12\x1ftcnapi.omni.projects.v1/ProjectR\x06parent\x12 \n" +
	"\tpage_size\x18\x02 \x01(\x05B\x03\xe0A\x01R\bpageSize\x12\"\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tB\x03\xe0A\x01R\tpageToken\x12\x1b\n" +
	"\x06filter\x18\x04 \x01(\tB\x03\xe0A\x01R\x06filter\"|\n" +
	"\x14ListProjectsResponse\x12<\n" +
	"\bprojects\x18\x01 \x03(\v2 .tcnapi.omni.projects.v1.ProjectR\bprojects\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"P\n" +
	"\x11GetProjectRequest\x12;\n" +
	"\x04name\x18\x01 \x01(\tB'\xe0A\x02\xfaA!\n" +
	"\x1ftcnapi.omni.projects.v1/ProjectR\x04name\"\x98\x01\n" +
	"\x14CreateProjectRequest\x12?\n" +
	"\x06parent\x18\x01 \x01(\tB'\xe0A\x02\xfaA!\x12\x1ftcnapi.omni.projects.v1/ProjectR\x06parent\x12?\n" +
	"\aproject\x18\x03 \x01(\v2 .tcnapi.omni.projects.v1.ProjectB\x03\xe0A\x02R\aproject\"\x99\x01\n" +
	"\x14UpdateProjectRequest\x12?\n" +
	"\aproject\x18\x01 \x01(\v2 .tcnapi.omni.projects.v1.ProjectB\x03\xe0A\x02R\aproject\x12@\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskB\x03\xe0A\x01R\n" +
	"updateMask\"S\n" +
	"\x14DeleteProjectRequest\x12;\n" +
	"\x04name\x18\x01 \x01(\tB'\xe0A\x02\xfaA!\n" +
	"\x1ftcnapi.omni.projects.v1/ProjectR\x04nameB\xea\x01\n" +
	"\x1bcom.tcnapi.omni.projects.v1B\rEntitiesProtoP\x01Z=github.com/tcncloud/api-go/tcnapi/omni/projects/v1;projectsv1\xa2\x02\x03TOP\xaa\x02\x17Tcnapi.Omni.Projects.V1\xca\x02\x17Tcnapi\\Omni\\Projects\\V1\xe2\x02#Tcnapi\\Omni\\Projects\\V1\\GPBMetadata\xea\x02\x1aTcnapi::Omni::Projects::V1b\x06proto3"

var (
	file_tcnapi_omni_projects_v1_entities_proto_rawDescOnce sync.Once
	file_tcnapi_omni_projects_v1_entities_proto_rawDescData []byte
)

func file_tcnapi_omni_projects_v1_entities_proto_rawDescGZIP() []byte {
	file_tcnapi_omni_projects_v1_entities_proto_rawDescOnce.Do(func() {
		file_tcnapi_omni_projects_v1_entities_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tcnapi_omni_projects_v1_entities_proto_rawDesc), len(file_tcnapi_omni_projects_v1_entities_proto_rawDesc)))
	})
	return file_tcnapi_omni_projects_v1_entities_proto_rawDescData
}

var file_tcnapi_omni_projects_v1_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tcnapi_omni_projects_v1_entities_proto_goTypes = []any{
	(*ListProjectsRequest)(nil),   // 0: tcnapi.omni.projects.v1.ListProjectsRequest
	(*ListProjectsResponse)(nil),  // 1: tcnapi.omni.projects.v1.ListProjectsResponse
	(*GetProjectRequest)(nil),     // 2: tcnapi.omni.projects.v1.GetProjectRequest
	(*CreateProjectRequest)(nil),  // 3: tcnapi.omni.projects.v1.CreateProjectRequest
	(*UpdateProjectRequest)(nil),  // 4: tcnapi.omni.projects.v1.UpdateProjectRequest
	(*DeleteProjectRequest)(nil),  // 5: tcnapi.omni.projects.v1.DeleteProjectRequest
	(*Project)(nil),               // 6: tcnapi.omni.projects.v1.Project
	(*fieldmaskpb.FieldMask)(nil), // 7: google.protobuf.FieldMask
}
var file_tcnapi_omni_projects_v1_entities_proto_depIdxs = []int32{
	6, // 0: tcnapi.omni.projects.v1.ListProjectsResponse.projects:type_name -> tcnapi.omni.projects.v1.Project
	6, // 1: tcnapi.omni.projects.v1.CreateProjectRequest.project:type_name -> tcnapi.omni.projects.v1.Project
	6, // 2: tcnapi.omni.projects.v1.UpdateProjectRequest.project:type_name -> tcnapi.omni.projects.v1.Project
	7, // 3: tcnapi.omni.projects.v1.UpdateProjectRequest.update_mask:type_name -> google.protobuf.FieldMask
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tcnapi_omni_projects_v1_entities_proto_init() }
func file_tcnapi_omni_projects_v1_entities_proto_init() {
	if File_tcnapi_omni_projects_v1_entities_proto != nil {
		return
	}
	file_tcnapi_omni_projects_v1_projects_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tcnapi_omni_projects_v1_entities_proto_rawDesc), len(file_tcnapi_omni_projects_v1_entities_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tcnapi_omni_projects_v1_entities_proto_goTypes,
		DependencyIndexes: file_tcnapi_omni_projects_v1_entities_proto_depIdxs,
		MessageInfos:      file_tcnapi_omni_projects_v1_entities_proto_msgTypes,
	}.Build()
	File_tcnapi_omni_projects_v1_entities_proto = out.File
	file_tcnapi_omni_projects_v1_entities_proto_goTypes = nil
	file_tcnapi_omni_projects_v1_entities_proto_depIdxs = nil
}
