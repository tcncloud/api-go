// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/lms/entities.proto

package lms

import (
	v0alpha "github.com/tcncloud/api-go/api/v0alpha"
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

type FileTemplateV2 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Template:
	//
	//	*FileTemplateV2_LegacyTemplate
	//	*FileTemplateV2_DockTemplate
	Template      isFileTemplateV2_Template `protobuf_oneof:"template"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileTemplateV2) Reset() {
	*x = FileTemplateV2{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileTemplateV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileTemplateV2) ProtoMessage() {}

func (x *FileTemplateV2) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileTemplateV2.ProtoReflect.Descriptor instead.
func (*FileTemplateV2) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{0}
}

func (x *FileTemplateV2) GetTemplate() isFileTemplateV2_Template {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *FileTemplateV2) GetLegacyTemplate() *v0alpha.FileTemplate {
	if x != nil {
		if x, ok := x.Template.(*FileTemplateV2_LegacyTemplate); ok {
			return x.LegacyTemplate
		}
	}
	return nil
}

func (x *FileTemplateV2) GetDockTemplate() *FileTemplate {
	if x != nil {
		if x, ok := x.Template.(*FileTemplateV2_DockTemplate); ok {
			return x.DockTemplate
		}
	}
	return nil
}

type isFileTemplateV2_Template interface {
	isFileTemplateV2_Template()
}

type FileTemplateV2_LegacyTemplate struct {
	LegacyTemplate *v0alpha.FileTemplate `protobuf:"bytes,1,opt,name=legacy_template,json=legacyTemplate,proto3,oneof"`
}

type FileTemplateV2_DockTemplate struct {
	DockTemplate *FileTemplate `protobuf:"bytes,2,opt,name=dock_template,json=dockTemplate,proto3,oneof"`
}

func (*FileTemplateV2_LegacyTemplate) isFileTemplateV2_Template() {}

func (*FileTemplateV2_DockTemplate) isFileTemplateV2_Template() {}

type FileTemplates struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Templates     []*FileTemplateV2      `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileTemplates) Reset() {
	*x = FileTemplates{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileTemplates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileTemplates) ProtoMessage() {}

func (x *FileTemplates) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileTemplates.ProtoReflect.Descriptor instead.
func (*FileTemplates) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{1}
}

func (x *FileTemplates) GetTemplates() []*FileTemplateV2 {
	if x != nil {
		return x.Templates
	}
	return nil
}

type FileTemplate struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	OrgId          string                 `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	FileTemplateId int64                  `protobuf:"varint,2,opt,name=file_template_id,json=fileTemplateId,proto3" json:"file_template_id,omitempty"`
	Filename       string                 `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Fields         []*Field               `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	ParseOpts      *ParseOpts             `protobuf:"bytes,5,opt,name=parse_opts,json=parseOpts,proto3" json:"parse_opts,omitempty"`
	Foid           int64                  `protobuf:"varint,6,opt,name=foid,proto3" json:"foid,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *FileTemplate) Reset() {
	*x = FileTemplate{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileTemplate) ProtoMessage() {}

func (x *FileTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileTemplate.ProtoReflect.Descriptor instead.
func (*FileTemplate) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{2}
}

func (x *FileTemplate) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *FileTemplate) GetFileTemplateId() int64 {
	if x != nil {
		return x.FileTemplateId
	}
	return 0
}

func (x *FileTemplate) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileTemplate) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *FileTemplate) GetParseOpts() *ParseOpts {
	if x != nil {
		return x.ParseOpts
	}
	return nil
}

func (x *FileTemplate) GetFoid() int64 {
	if x != nil {
		return x.Foid
	}
	return 0
}

type Field struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	SyntaxType string                 `protobuf:"bytes,1,opt,name=syntax_type,json=syntaxType,proto3" json:"syntax_type,omitempty"`
	PresiType  string                 `protobuf:"bytes,2,opt,name=presi_type,json=presiType,proto3" json:"presi_type,omitempty"`
	Name       string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Format     string                 `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
	// example of a value for this field
	RawValue      string `protobuf:"bytes,5,opt,name=raw_value,json=rawValue,proto3" json:"raw_value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Field) Reset() {
	*x = Field{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{3}
}

func (x *Field) GetSyntaxType() string {
	if x != nil {
		return x.SyntaxType
	}
	return ""
}

func (x *Field) GetPresiType() string {
	if x != nil {
		return x.PresiType
	}
	return ""
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *Field) GetRawValue() string {
	if x != nil {
		return x.RawValue
	}
	return ""
}

type ParseOpts struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Ftype:
	//
	//	*ParseOpts_Csv
	//	*ParseOpts_Json
	//	*ParseOpts_Jsonl
	//	*ParseOpts_Fixed
	//	*ParseOpts_Parquet
	Ftype         isParseOpts_Ftype `protobuf_oneof:"ftype"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParseOpts) Reset() {
	*x = ParseOpts{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseOpts) ProtoMessage() {}

func (x *ParseOpts) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseOpts.ProtoReflect.Descriptor instead.
func (*ParseOpts) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{4}
}

func (x *ParseOpts) GetFtype() isParseOpts_Ftype {
	if x != nil {
		return x.Ftype
	}
	return nil
}

func (x *ParseOpts) GetCsv() *OptsCsv {
	if x != nil {
		if x, ok := x.Ftype.(*ParseOpts_Csv); ok {
			return x.Csv
		}
	}
	return nil
}

func (x *ParseOpts) GetJson() *OptsJson {
	if x != nil {
		if x, ok := x.Ftype.(*ParseOpts_Json); ok {
			return x.Json
		}
	}
	return nil
}

func (x *ParseOpts) GetJsonl() *OptsJsonL {
	if x != nil {
		if x, ok := x.Ftype.(*ParseOpts_Jsonl); ok {
			return x.Jsonl
		}
	}
	return nil
}

func (x *ParseOpts) GetFixed() *OptsFixed {
	if x != nil {
		if x, ok := x.Ftype.(*ParseOpts_Fixed); ok {
			return x.Fixed
		}
	}
	return nil
}

func (x *ParseOpts) GetParquet() *OptsParquet {
	if x != nil {
		if x, ok := x.Ftype.(*ParseOpts_Parquet); ok {
			return x.Parquet
		}
	}
	return nil
}

type isParseOpts_Ftype interface {
	isParseOpts_Ftype()
}

type ParseOpts_Csv struct {
	Csv *OptsCsv `protobuf:"bytes,1,opt,name=csv,proto3,oneof"`
}

type ParseOpts_Json struct {
	Json *OptsJson `protobuf:"bytes,2,opt,name=json,proto3,oneof"`
}

type ParseOpts_Jsonl struct {
	Jsonl *OptsJsonL `protobuf:"bytes,3,opt,name=jsonl,proto3,oneof"`
}

type ParseOpts_Fixed struct {
	Fixed *OptsFixed `protobuf:"bytes,4,opt,name=fixed,proto3,oneof"`
}

type ParseOpts_Parquet struct {
	Parquet *OptsParquet `protobuf:"bytes,5,opt,name=parquet,proto3,oneof"`
}

func (*ParseOpts_Csv) isParseOpts_Ftype() {}

func (*ParseOpts_Json) isParseOpts_Ftype() {}

func (*ParseOpts_Jsonl) isParseOpts_Ftype() {}

func (*ParseOpts_Fixed) isParseOpts_Ftype() {}

func (*ParseOpts_Parquet) isParseOpts_Ftype() {}

type OptsCsv struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HasHeader     bool                   `protobuf:"varint,1,opt,name=has_header,json=hasHeader,proto3" json:"has_header,omitempty"`
	SkipRows      int64                  `protobuf:"varint,2,opt,name=skip_rows,json=skipRows,proto3" json:"skip_rows,omitempty"`
	Header        []string               `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty"`
	Separator     string                 `protobuf:"bytes,4,opt,name=separator,proto3" json:"separator,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OptsCsv) Reset() {
	*x = OptsCsv{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsCsv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsCsv) ProtoMessage() {}

func (x *OptsCsv) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptsCsv.ProtoReflect.Descriptor instead.
func (*OptsCsv) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{5}
}

func (x *OptsCsv) GetHasHeader() bool {
	if x != nil {
		return x.HasHeader
	}
	return false
}

func (x *OptsCsv) GetSkipRows() int64 {
	if x != nil {
		return x.SkipRows
	}
	return 0
}

func (x *OptsCsv) GetHeader() []string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OptsCsv) GetSeparator() string {
	if x != nil {
		return x.Separator
	}
	return ""
}

type OptsJson struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecordsRoot   string                 `protobuf:"bytes,1,opt,name=records_root,json=recordsRoot,proto3" json:"records_root,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OptsJson) Reset() {
	*x = OptsJson{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsJson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsJson) ProtoMessage() {}

func (x *OptsJson) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptsJson.ProtoReflect.Descriptor instead.
func (*OptsJson) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{6}
}

func (x *OptsJson) GetRecordsRoot() string {
	if x != nil {
		return x.RecordsRoot
	}
	return ""
}

type OptsJsonL struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OptsJsonL) Reset() {
	*x = OptsJsonL{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsJsonL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsJsonL) ProtoMessage() {}

func (x *OptsJsonL) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptsJsonL.ProtoReflect.Descriptor instead.
func (*OptsJsonL) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{7}
}

type OptsFixed struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HasHeader     bool                   `protobuf:"varint,1,opt,name=has_header,json=hasHeader,proto3" json:"has_header,omitempty"`
	Header        []*OptsFixed_Field     `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OptsFixed) Reset() {
	*x = OptsFixed{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsFixed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsFixed) ProtoMessage() {}

func (x *OptsFixed) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptsFixed.ProtoReflect.Descriptor instead.
func (*OptsFixed) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{8}
}

func (x *OptsFixed) GetHasHeader() bool {
	if x != nil {
		return x.HasHeader
	}
	return false
}

func (x *OptsFixed) GetHeader() []*OptsFixed_Field {
	if x != nil {
		return x.Header
	}
	return nil
}

type OptsParquet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OptsParquet) Reset() {
	*x = OptsParquet{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsParquet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsParquet) ProtoMessage() {}

func (x *OptsParquet) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptsParquet.ProtoReflect.Descriptor instead.
func (*OptsParquet) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{9}
}

type NewTemplate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrgId         string                 `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Filename      string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Data          []byte                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewTemplate) Reset() {
	*x = NewTemplate{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTemplate) ProtoMessage() {}

func (x *NewTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTemplate.ProtoReflect.Descriptor instead.
func (*NewTemplate) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{10}
}

func (x *NewTemplate) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *NewTemplate) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *NewTemplate) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ExistingTemplate struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	FileTemplateId int64                  `protobuf:"varint,1,opt,name=file_template_id,json=fileTemplateId,proto3" json:"file_template_id,omitempty"`
	ParseOpts      *ParseOpts             `protobuf:"bytes,2,opt,name=parse_opts,json=parseOpts,proto3" json:"parse_opts,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ExistingTemplate) Reset() {
	*x = ExistingTemplate{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExistingTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistingTemplate) ProtoMessage() {}

func (x *ExistingTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistingTemplate.ProtoReflect.Descriptor instead.
func (*ExistingTemplate) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{11}
}

func (x *ExistingTemplate) GetFileTemplateId() int64 {
	if x != nil {
		return x.FileTemplateId
	}
	return 0
}

func (x *ExistingTemplate) GetParseOpts() *ParseOpts {
	if x != nil {
		return x.ParseOpts
	}
	return nil
}

type OptsFixed_Field struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// for fixed width files indicates the starting position of the data.
	// if it is -1, starting position is one character after the previous fields
	// starting position + length. if this is the first field and it is -1,
	// starting position is 0
	StartingPosition int32 `protobuf:"varint,13,opt,name=starting_position,json=startingPosition,proto3" json:"starting_position,omitempty"`
	// for fixed width files indicates how many characters to the right of
	// starting_position we will read. this field is required to be greater than
	// 0.
	FieldLength   int32 `protobuf:"varint,14,opt,name=field_length,json=fieldLength,proto3" json:"field_length,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OptsFixed_Field) Reset() {
	*x = OptsFixed_Field{}
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsFixed_Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsFixed_Field) ProtoMessage() {}

func (x *OptsFixed_Field) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_lms_entities_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptsFixed_Field.ProtoReflect.Descriptor instead.
func (*OptsFixed_Field) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_lms_entities_proto_rawDescGZIP(), []int{8, 0}
}

func (x *OptsFixed_Field) GetStartingPosition() int32 {
	if x != nil {
		return x.StartingPosition
	}
	return 0
}

func (x *OptsFixed_Field) GetFieldLength() int32 {
	if x != nil {
		return x.FieldLength
	}
	return 0
}

var File_api_v1alpha1_lms_entities_proto protoreflect.FileDescriptor

const file_api_v1alpha1_lms_entities_proto_rawDesc = "" +
	"\n" +
	"\x1fapi/v1alpha1/lms/entities.proto\x12\x10api.v1alpha1.lms\x1a\x15api/v0alpha/lms.proto\"\xa9\x01\n" +
	"\x0eFileTemplateV2\x12D\n" +
	"\x0flegacy_template\x18\x01 \x01(\v2\x19.api.v0alpha.FileTemplateH\x00R\x0elegacyTemplate\x12E\n" +
	"\rdock_template\x18\x02 \x01(\v2\x1e.api.v1alpha1.lms.FileTemplateH\x00R\fdockTemplateB\n" +
	"\n" +
	"\btemplate\"O\n" +
	"\rFileTemplates\x12>\n" +
	"\ttemplates\x18\x01 \x03(\v2 .api.v1alpha1.lms.FileTemplateV2R\ttemplates\"\xec\x01\n" +
	"\fFileTemplate\x12\x15\n" +
	"\x06org_id\x18\x01 \x01(\tR\x05orgId\x12(\n" +
	"\x10file_template_id\x18\x02 \x01(\x03R\x0efileTemplateId\x12\x1a\n" +
	"\bfilename\x18\x03 \x01(\tR\bfilename\x12/\n" +
	"\x06fields\x18\x04 \x03(\v2\x17.api.v1alpha1.lms.FieldR\x06fields\x12:\n" +
	"\n" +
	"parse_opts\x18\x05 \x01(\v2\x1b.api.v1alpha1.lms.ParseOptsR\tparseOpts\x12\x12\n" +
	"\x04foid\x18\x06 \x01(\x03R\x04foid\"\x90\x01\n" +
	"\x05Field\x12\x1f\n" +
	"\vsyntax_type\x18\x01 \x01(\tR\n" +
	"syntaxType\x12\x1d\n" +
	"\n" +
	"presi_type\x18\x02 \x01(\tR\tpresiType\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x16\n" +
	"\x06format\x18\x04 \x01(\tR\x06format\x12\x1b\n" +
	"\traw_value\x18\x05 \x01(\tR\brawValue\"\x9a\x02\n" +
	"\tParseOpts\x12-\n" +
	"\x03csv\x18\x01 \x01(\v2\x19.api.v1alpha1.lms.OptsCsvH\x00R\x03csv\x120\n" +
	"\x04json\x18\x02 \x01(\v2\x1a.api.v1alpha1.lms.OptsJsonH\x00R\x04json\x123\n" +
	"\x05jsonl\x18\x03 \x01(\v2\x1b.api.v1alpha1.lms.OptsJsonLH\x00R\x05jsonl\x123\n" +
	"\x05fixed\x18\x04 \x01(\v2\x1b.api.v1alpha1.lms.OptsFixedH\x00R\x05fixed\x129\n" +
	"\aparquet\x18\x05 \x01(\v2\x1d.api.v1alpha1.lms.OptsParquetH\x00R\aparquetB\a\n" +
	"\x05ftype\"{\n" +
	"\aOptsCsv\x12\x1d\n" +
	"\n" +
	"has_header\x18\x01 \x01(\bR\thasHeader\x12\x1b\n" +
	"\tskip_rows\x18\x02 \x01(\x03R\bskipRows\x12\x16\n" +
	"\x06header\x18\x03 \x03(\tR\x06header\x12\x1c\n" +
	"\tseparator\x18\x04 \x01(\tR\tseparator\"-\n" +
	"\bOptsJson\x12!\n" +
	"\frecords_root\x18\x01 \x01(\tR\vrecordsRoot\"\v\n" +
	"\tOptsJsonL\"\xbe\x01\n" +
	"\tOptsFixed\x12\x1d\n" +
	"\n" +
	"has_header\x18\x01 \x01(\bR\thasHeader\x129\n" +
	"\x06header\x18\x02 \x03(\v2!.api.v1alpha1.lms.OptsFixed.FieldR\x06header\x1aW\n" +
	"\x05Field\x12+\n" +
	"\x11starting_position\x18\r \x01(\x05R\x10startingPosition\x12!\n" +
	"\ffield_length\x18\x0e \x01(\x05R\vfieldLength\"\r\n" +
	"\vOptsParquet\"T\n" +
	"\vNewTemplate\x12\x15\n" +
	"\x06org_id\x18\x01 \x01(\tR\x05orgId\x12\x1a\n" +
	"\bfilename\x18\x02 \x01(\tR\bfilename\x12\x12\n" +
	"\x04data\x18\x03 \x01(\fR\x04data\"x\n" +
	"\x10ExistingTemplate\x12(\n" +
	"\x10file_template_id\x18\x01 \x01(\x03R\x0efileTemplateId\x12:\n" +
	"\n" +
	"parse_opts\x18\x02 \x01(\v2\x1b.api.v1alpha1.lms.ParseOptsR\tparseOptsB\xb4\x01\n" +
	"\x14com.api.v1alpha1.lmsB\rEntitiesProtoP\x01Z+github.com/tcncloud/api-go/api/v1alpha1/lms\xa2\x02\x03AVL\xaa\x02\x10Api.V1alpha1.Lms\xca\x02\x10Api\\V1alpha1\\Lms\xe2\x02\x1cApi\\V1alpha1\\Lms\\GPBMetadata\xea\x02\x12Api::V1alpha1::Lmsb\x06proto3"

var (
	file_api_v1alpha1_lms_entities_proto_rawDescOnce sync.Once
	file_api_v1alpha1_lms_entities_proto_rawDescData []byte
)

func file_api_v1alpha1_lms_entities_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_lms_entities_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_lms_entities_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_lms_entities_proto_rawDesc), len(file_api_v1alpha1_lms_entities_proto_rawDesc)))
	})
	return file_api_v1alpha1_lms_entities_proto_rawDescData
}

var file_api_v1alpha1_lms_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_v1alpha1_lms_entities_proto_goTypes = []any{
	(*FileTemplateV2)(nil),       // 0: api.v1alpha1.lms.FileTemplateV2
	(*FileTemplates)(nil),        // 1: api.v1alpha1.lms.FileTemplates
	(*FileTemplate)(nil),         // 2: api.v1alpha1.lms.FileTemplate
	(*Field)(nil),                // 3: api.v1alpha1.lms.Field
	(*ParseOpts)(nil),            // 4: api.v1alpha1.lms.ParseOpts
	(*OptsCsv)(nil),              // 5: api.v1alpha1.lms.OptsCsv
	(*OptsJson)(nil),             // 6: api.v1alpha1.lms.OptsJson
	(*OptsJsonL)(nil),            // 7: api.v1alpha1.lms.OptsJsonL
	(*OptsFixed)(nil),            // 8: api.v1alpha1.lms.OptsFixed
	(*OptsParquet)(nil),          // 9: api.v1alpha1.lms.OptsParquet
	(*NewTemplate)(nil),          // 10: api.v1alpha1.lms.NewTemplate
	(*ExistingTemplate)(nil),     // 11: api.v1alpha1.lms.ExistingTemplate
	(*OptsFixed_Field)(nil),      // 12: api.v1alpha1.lms.OptsFixed.Field
	(*v0alpha.FileTemplate)(nil), // 13: api.v0alpha.FileTemplate
}
var file_api_v1alpha1_lms_entities_proto_depIdxs = []int32{
	13, // 0: api.v1alpha1.lms.FileTemplateV2.legacy_template:type_name -> api.v0alpha.FileTemplate
	2,  // 1: api.v1alpha1.lms.FileTemplateV2.dock_template:type_name -> api.v1alpha1.lms.FileTemplate
	0,  // 2: api.v1alpha1.lms.FileTemplates.templates:type_name -> api.v1alpha1.lms.FileTemplateV2
	3,  // 3: api.v1alpha1.lms.FileTemplate.fields:type_name -> api.v1alpha1.lms.Field
	4,  // 4: api.v1alpha1.lms.FileTemplate.parse_opts:type_name -> api.v1alpha1.lms.ParseOpts
	5,  // 5: api.v1alpha1.lms.ParseOpts.csv:type_name -> api.v1alpha1.lms.OptsCsv
	6,  // 6: api.v1alpha1.lms.ParseOpts.json:type_name -> api.v1alpha1.lms.OptsJson
	7,  // 7: api.v1alpha1.lms.ParseOpts.jsonl:type_name -> api.v1alpha1.lms.OptsJsonL
	8,  // 8: api.v1alpha1.lms.ParseOpts.fixed:type_name -> api.v1alpha1.lms.OptsFixed
	9,  // 9: api.v1alpha1.lms.ParseOpts.parquet:type_name -> api.v1alpha1.lms.OptsParquet
	12, // 10: api.v1alpha1.lms.OptsFixed.header:type_name -> api.v1alpha1.lms.OptsFixed.Field
	4,  // 11: api.v1alpha1.lms.ExistingTemplate.parse_opts:type_name -> api.v1alpha1.lms.ParseOpts
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_lms_entities_proto_init() }
func file_api_v1alpha1_lms_entities_proto_init() {
	if File_api_v1alpha1_lms_entities_proto != nil {
		return
	}
	file_api_v1alpha1_lms_entities_proto_msgTypes[0].OneofWrappers = []any{
		(*FileTemplateV2_LegacyTemplate)(nil),
		(*FileTemplateV2_DockTemplate)(nil),
	}
	file_api_v1alpha1_lms_entities_proto_msgTypes[4].OneofWrappers = []any{
		(*ParseOpts_Csv)(nil),
		(*ParseOpts_Json)(nil),
		(*ParseOpts_Jsonl)(nil),
		(*ParseOpts_Fixed)(nil),
		(*ParseOpts_Parquet)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_lms_entities_proto_rawDesc), len(file_api_v1alpha1_lms_entities_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_lms_entities_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_lms_entities_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_lms_entities_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_lms_entities_proto = out.File
	file_api_v1alpha1_lms_entities_proto_goTypes = nil
	file_api_v1alpha1_lms_entities_proto_depIdxs = nil
}
