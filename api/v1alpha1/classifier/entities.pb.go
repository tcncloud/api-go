// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/classifier/entities.proto

package classifier

import (
	commons "github.com/tcncloud/api-go/api/commons"
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

type ClassifierEntityTypes struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Types         []commons.ClassifierEntityType `protobuf:"varint,1,rep,packed,name=types,proto3,enum=api.commons.ClassifierEntityType" json:"types,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClassifierEntityTypes) Reset() {
	*x = ClassifierEntityTypes{}
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClassifierEntityTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassifierEntityTypes) ProtoMessage() {}

func (x *ClassifierEntityTypes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassifierEntityTypes.ProtoReflect.Descriptor instead.
func (*ClassifierEntityTypes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{0}
}

func (x *ClassifierEntityTypes) GetTypes() []commons.ClassifierEntityType {
	if x != nil {
		return x.Types
	}
	return nil
}

type FileTemplate struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	FileTemplateId int64                  `protobuf:"varint,1,opt,name=file_template_id,json=fileTemplateId,proto3" json:"file_template_id,omitempty"`
	Filename       string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Fields         []*FileTemplate_Field  `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	// Deprecated: Marked as deprecated in api/v1alpha1/classifier/entities.proto.
	ParseOpts *ParseOpts `protobuf:"bytes,4,opt,name=parse_opts,json=parseOpts,proto3" json:"parse_opts,omitempty"`
	// Deprecated: Marked as deprecated in api/v1alpha1/classifier/entities.proto.
	Constraints   *Constraints `protobuf:"bytes,5,opt,name=constraints,proto3" json:"constraints,omitempty"`
	Foid          int64        `protobuf:"varint,6,opt,name=foid,proto3" json:"foid,omitempty"`
	Opts          *Opts        `protobuf:"bytes,7,opt,name=opts,proto3" json:"opts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileTemplate) Reset() {
	*x = FileTemplate{}
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileTemplate) ProtoMessage() {}

func (x *FileTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[1]
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
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{1}
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

func (x *FileTemplate) GetFields() []*FileTemplate_Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/classifier/entities.proto.
func (x *FileTemplate) GetParseOpts() *ParseOpts {
	if x != nil {
		return x.ParseOpts
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/classifier/entities.proto.
func (x *FileTemplate) GetConstraints() *Constraints {
	if x != nil {
		return x.Constraints
	}
	return nil
}

func (x *FileTemplate) GetFoid() int64 {
	if x != nil {
		return x.Foid
	}
	return 0
}

func (x *FileTemplate) GetOpts() *Opts {
	if x != nil {
		return x.Opts
	}
	return nil
}

type Opts struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// map from field name to date format
	// e.g: "dob": yyyy/mm/dd
	DateFormats map[string]string `protobuf:"bytes,1,rep,name=date_formats,json=dateFormats,proto3" json:"date_formats,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// keys are old field names, values are new field names
	RenameFields  map[string]string `protobuf:"bytes,2,rep,name=rename_fields,json=renameFields,proto3" json:"rename_fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ParseOpts     *ParseOpts        `protobuf:"bytes,3,opt,name=parse_opts,json=parseOpts,proto3" json:"parse_opts,omitempty"`
	Constraints   *Constraints      `protobuf:"bytes,4,opt,name=constraints,proto3" json:"constraints,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Opts) Reset() {
	*x = Opts{}
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Opts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Opts) ProtoMessage() {}

func (x *Opts) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Opts.ProtoReflect.Descriptor instead.
func (*Opts) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{2}
}

func (x *Opts) GetDateFormats() map[string]string {
	if x != nil {
		return x.DateFormats
	}
	return nil
}

func (x *Opts) GetRenameFields() map[string]string {
	if x != nil {
		return x.RenameFields
	}
	return nil
}

func (x *Opts) GetParseOpts() *ParseOpts {
	if x != nil {
		return x.ParseOpts
	}
	return nil
}

func (x *Opts) GetConstraints() *Constraints {
	if x != nil {
		return x.Constraints
	}
	return nil
}

// details on how the data should be
// parsed with respect to the file type
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
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseOpts) ProtoMessage() {}

func (x *ParseOpts) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[3]
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
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{3}
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
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsCsv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsCsv) ProtoMessage() {}

func (x *OptsCsv) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[4]
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
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{4}
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
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsJson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsJson) ProtoMessage() {}

func (x *OptsJson) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[5]
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
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{5}
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
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsJsonL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsJsonL) ProtoMessage() {}

func (x *OptsJsonL) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[6]
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
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{6}
}

type OptsFixed struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// field opts keyed by field name
	Positions     map[string]*OptsFixed_FieldOpts `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	HasHeader     bool                            `protobuf:"varint,2,opt,name=has_header,json=hasHeader,proto3" json:"has_header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OptsFixed) Reset() {
	*x = OptsFixed{}
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsFixed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsFixed) ProtoMessage() {}

func (x *OptsFixed) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[7]
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
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{7}
}

func (x *OptsFixed) GetPositions() map[string]*OptsFixed_FieldOpts {
	if x != nil {
		return x.Positions
	}
	return nil
}

func (x *OptsFixed) GetHasHeader() bool {
	if x != nil {
		return x.HasHeader
	}
	return false
}

type OptsParquet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OptsParquet) Reset() {
	*x = OptsParquet{}
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsParquet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsParquet) ProtoMessage() {}

func (x *OptsParquet) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[8]
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
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{8}
}

// these are options to forbid/allow certain
// types from being guessed by Classifier
type Constraints struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Forbid        map[string]*ClassifierEntityTypes `protobuf:"bytes,1,rep,name=forbid,proto3" json:"forbid,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Allow         map[string]*ClassifierEntityTypes `protobuf:"bytes,2,rep,name=allow,proto3" json:"allow,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Constraints) Reset() {
	*x = Constraints{}
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Constraints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Constraints) ProtoMessage() {}

func (x *Constraints) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Constraints.ProtoReflect.Descriptor instead.
func (*Constraints) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{9}
}

func (x *Constraints) GetForbid() map[string]*ClassifierEntityTypes {
	if x != nil {
		return x.Forbid
	}
	return nil
}

func (x *Constraints) GetAllow() map[string]*ClassifierEntityTypes {
	if x != nil {
		return x.Allow
	}
	return nil
}

type ParseHints struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParseOpts     *ParseOpts             `protobuf:"bytes,1,opt,name=parse_opts,json=parseOpts,proto3" json:"parse_opts,omitempty"`
	Constraints   *Constraints           `protobuf:"bytes,2,opt,name=constraints,proto3" json:"constraints,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParseHints) Reset() {
	*x = ParseHints{}
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseHints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseHints) ProtoMessage() {}

func (x *ParseHints) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseHints.ProtoReflect.Descriptor instead.
func (*ParseHints) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{10}
}

func (x *ParseHints) GetParseOpts() *ParseOpts {
	if x != nil {
		return x.ParseOpts
	}
	return nil
}

func (x *ParseHints) GetConstraints() *Constraints {
	if x != nil {
		return x.Constraints
	}
	return nil
}

type FileTemplate_Field struct {
	state      protoimpl.MessageState       `protogen:"open.v1"`
	SyntaxType string                       `protobuf:"bytes,1,opt,name=syntax_type,json=syntaxType,proto3" json:"syntax_type,omitempty"`
	EntityType commons.ClassifierEntityType `protobuf:"varint,2,opt,name=entity_type,json=entityType,proto3,enum=api.commons.ClassifierEntityType" json:"entity_type,omitempty"`
	Name       string                       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Format     string                       `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
	// example of a value for this field
	RawValue      string `protobuf:"bytes,5,opt,name=raw_value,json=rawValue,proto3" json:"raw_value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileTemplate_Field) Reset() {
	*x = FileTemplate_Field{}
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileTemplate_Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileTemplate_Field) ProtoMessage() {}

func (x *FileTemplate_Field) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileTemplate_Field.ProtoReflect.Descriptor instead.
func (*FileTemplate_Field) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FileTemplate_Field) GetSyntaxType() string {
	if x != nil {
		return x.SyntaxType
	}
	return ""
}

func (x *FileTemplate_Field) GetEntityType() commons.ClassifierEntityType {
	if x != nil {
		return x.EntityType
	}
	return commons.ClassifierEntityType(0)
}

func (x *FileTemplate_Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileTemplate_Field) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *FileTemplate_Field) GetRawValue() string {
	if x != nil {
		return x.RawValue
	}
	return ""
}

type OptsFixed_FieldOpts struct {
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

func (x *OptsFixed_FieldOpts) Reset() {
	*x = OptsFixed_FieldOpts{}
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptsFixed_FieldOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptsFixed_FieldOpts) ProtoMessage() {}

func (x *OptsFixed_FieldOpts) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_classifier_entities_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptsFixed_FieldOpts.ProtoReflect.Descriptor instead.
func (*OptsFixed_FieldOpts) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_classifier_entities_proto_rawDescGZIP(), []int{7, 0}
}

func (x *OptsFixed_FieldOpts) GetStartingPosition() int32 {
	if x != nil {
		return x.StartingPosition
	}
	return 0
}

func (x *OptsFixed_FieldOpts) GetFieldLength() int32 {
	if x != nil {
		return x.FieldLength
	}
	return 0
}

var File_api_v1alpha1_classifier_entities_proto protoreflect.FileDescriptor

const file_api_v1alpha1_classifier_entities_proto_rawDesc = "" +
	"\n" +
	"&api/v1alpha1/classifier/entities.proto\x12\x17api.v1alpha1.classifier\x1a\x1capi/commons/classifier.proto\"P\n" +
	"\x15ClassifierEntityTypes\x127\n" +
	"\x05types\x18\x01 \x03(\x0e2!.api.commons.ClassifierEntityTypeR\x05types\"\xaf\x04\n" +
	"\fFileTemplate\x12,\n" +
	"\x10file_template_id\x18\x01 \x01(\x03B\x020\x01R\x0efileTemplateId\x12\x1a\n" +
	"\bfilename\x18\x02 \x01(\tR\bfilename\x12C\n" +
	"\x06fields\x18\x03 \x03(\v2+.api.v1alpha1.classifier.FileTemplate.FieldR\x06fields\x12E\n" +
	"\n" +
	"parse_opts\x18\x04 \x01(\v2\".api.v1alpha1.classifier.ParseOptsB\x02\x18\x01R\tparseOpts\x12J\n" +
	"\vconstraints\x18\x05 \x01(\v2$.api.v1alpha1.classifier.ConstraintsB\x02\x18\x01R\vconstraints\x12\x12\n" +
	"\x04foid\x18\x06 \x01(\x03R\x04foid\x121\n" +
	"\x04opts\x18\a \x01(\v2\x1d.api.v1alpha1.classifier.OptsR\x04opts\x1a\xb5\x01\n" +
	"\x05Field\x12\x1f\n" +
	"\vsyntax_type\x18\x01 \x01(\tR\n" +
	"syntaxType\x12B\n" +
	"\ventity_type\x18\x02 \x01(\x0e2!.api.commons.ClassifierEntityTypeR\n" +
	"entityType\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x16\n" +
	"\x06format\x18\x04 \x01(\tR\x06format\x12\x1b\n" +
	"\traw_value\x18\x05 \x01(\tR\brawValue\"\xbb\x03\n" +
	"\x04Opts\x12Q\n" +
	"\fdate_formats\x18\x01 \x03(\v2..api.v1alpha1.classifier.Opts.DateFormatsEntryR\vdateFormats\x12T\n" +
	"\rrename_fields\x18\x02 \x03(\v2/.api.v1alpha1.classifier.Opts.RenameFieldsEntryR\frenameFields\x12A\n" +
	"\n" +
	"parse_opts\x18\x03 \x01(\v2\".api.v1alpha1.classifier.ParseOptsR\tparseOpts\x12F\n" +
	"\vconstraints\x18\x04 \x01(\v2$.api.v1alpha1.classifier.ConstraintsR\vconstraints\x1a>\n" +
	"\x10DateFormatsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1a?\n" +
	"\x11RenameFieldsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xbd\x02\n" +
	"\tParseOpts\x124\n" +
	"\x03csv\x18\x01 \x01(\v2 .api.v1alpha1.classifier.OptsCsvH\x00R\x03csv\x127\n" +
	"\x04json\x18\x02 \x01(\v2!.api.v1alpha1.classifier.OptsJsonH\x00R\x04json\x12:\n" +
	"\x05jsonl\x18\x03 \x01(\v2\".api.v1alpha1.classifier.OptsJsonLH\x00R\x05jsonl\x12:\n" +
	"\x05fixed\x18\x04 \x01(\v2\".api.v1alpha1.classifier.OptsFixedH\x00R\x05fixed\x12@\n" +
	"\aparquet\x18\x05 \x01(\v2$.api.v1alpha1.classifier.OptsParquetH\x00R\aparquetB\a\n" +
	"\x05ftype\"{\n" +
	"\aOptsCsv\x12\x1d\n" +
	"\n" +
	"has_header\x18\x01 \x01(\bR\thasHeader\x12\x1b\n" +
	"\tskip_rows\x18\x02 \x01(\x03R\bskipRows\x12\x16\n" +
	"\x06header\x18\x03 \x03(\tR\x06header\x12\x1c\n" +
	"\tseparator\x18\x04 \x01(\tR\tseparator\"-\n" +
	"\bOptsJson\x12!\n" +
	"\frecords_root\x18\x01 \x01(\tR\vrecordsRoot\"\v\n" +
	"\tOptsJsonL\"\xc4\x02\n" +
	"\tOptsFixed\x12O\n" +
	"\tpositions\x18\x01 \x03(\v21.api.v1alpha1.classifier.OptsFixed.PositionsEntryR\tpositions\x12\x1d\n" +
	"\n" +
	"has_header\x18\x02 \x01(\bR\thasHeader\x1a[\n" +
	"\tFieldOpts\x12+\n" +
	"\x11starting_position\x18\r \x01(\x05R\x10startingPosition\x12!\n" +
	"\ffield_length\x18\x0e \x01(\x05R\vfieldLength\x1aj\n" +
	"\x0ePositionsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12B\n" +
	"\x05value\x18\x02 \x01(\v2,.api.v1alpha1.classifier.OptsFixed.FieldOptsR\x05value:\x028\x01\"\r\n" +
	"\vOptsParquet\"\xf3\x02\n" +
	"\vConstraints\x12H\n" +
	"\x06forbid\x18\x01 \x03(\v20.api.v1alpha1.classifier.Constraints.ForbidEntryR\x06forbid\x12E\n" +
	"\x05allow\x18\x02 \x03(\v2/.api.v1alpha1.classifier.Constraints.AllowEntryR\x05allow\x1ai\n" +
	"\vForbidEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12D\n" +
	"\x05value\x18\x02 \x01(\v2..api.v1alpha1.classifier.ClassifierEntityTypesR\x05value:\x028\x01\x1ah\n" +
	"\n" +
	"AllowEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12D\n" +
	"\x05value\x18\x02 \x01(\v2..api.v1alpha1.classifier.ClassifierEntityTypesR\x05value:\x028\x01\"\x97\x01\n" +
	"\n" +
	"ParseHints\x12A\n" +
	"\n" +
	"parse_opts\x18\x01 \x01(\v2\".api.v1alpha1.classifier.ParseOptsR\tparseOpts\x12F\n" +
	"\vconstraints\x18\x02 \x01(\v2$.api.v1alpha1.classifier.ConstraintsR\vconstraintsB\xde\x01\n" +
	"\x1bcom.api.v1alpha1.classifierB\rEntitiesProtoP\x01Z2github.com/tcncloud/api-go/api/v1alpha1/classifier\xa2\x02\x03AVC\xaa\x02\x17Api.V1alpha1.Classifier\xca\x02\x17Api\\V1alpha1\\Classifier\xe2\x02#Api\\V1alpha1\\Classifier\\GPBMetadata\xea\x02\x19Api::V1alpha1::Classifierb\x06proto3"

var (
	file_api_v1alpha1_classifier_entities_proto_rawDescOnce sync.Once
	file_api_v1alpha1_classifier_entities_proto_rawDescData []byte
)

func file_api_v1alpha1_classifier_entities_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_classifier_entities_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_classifier_entities_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_classifier_entities_proto_rawDesc), len(file_api_v1alpha1_classifier_entities_proto_rawDesc)))
	})
	return file_api_v1alpha1_classifier_entities_proto_rawDescData
}

var file_api_v1alpha1_classifier_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_api_v1alpha1_classifier_entities_proto_goTypes = []any{
	(*ClassifierEntityTypes)(nil),     // 0: api.v1alpha1.classifier.ClassifierEntityTypes
	(*FileTemplate)(nil),              // 1: api.v1alpha1.classifier.FileTemplate
	(*Opts)(nil),                      // 2: api.v1alpha1.classifier.Opts
	(*ParseOpts)(nil),                 // 3: api.v1alpha1.classifier.ParseOpts
	(*OptsCsv)(nil),                   // 4: api.v1alpha1.classifier.OptsCsv
	(*OptsJson)(nil),                  // 5: api.v1alpha1.classifier.OptsJson
	(*OptsJsonL)(nil),                 // 6: api.v1alpha1.classifier.OptsJsonL
	(*OptsFixed)(nil),                 // 7: api.v1alpha1.classifier.OptsFixed
	(*OptsParquet)(nil),               // 8: api.v1alpha1.classifier.OptsParquet
	(*Constraints)(nil),               // 9: api.v1alpha1.classifier.Constraints
	(*ParseHints)(nil),                // 10: api.v1alpha1.classifier.ParseHints
	(*FileTemplate_Field)(nil),        // 11: api.v1alpha1.classifier.FileTemplate.Field
	nil,                               // 12: api.v1alpha1.classifier.Opts.DateFormatsEntry
	nil,                               // 13: api.v1alpha1.classifier.Opts.RenameFieldsEntry
	(*OptsFixed_FieldOpts)(nil),       // 14: api.v1alpha1.classifier.OptsFixed.FieldOpts
	nil,                               // 15: api.v1alpha1.classifier.OptsFixed.PositionsEntry
	nil,                               // 16: api.v1alpha1.classifier.Constraints.ForbidEntry
	nil,                               // 17: api.v1alpha1.classifier.Constraints.AllowEntry
	(commons.ClassifierEntityType)(0), // 18: api.commons.ClassifierEntityType
}
var file_api_v1alpha1_classifier_entities_proto_depIdxs = []int32{
	18, // 0: api.v1alpha1.classifier.ClassifierEntityTypes.types:type_name -> api.commons.ClassifierEntityType
	11, // 1: api.v1alpha1.classifier.FileTemplate.fields:type_name -> api.v1alpha1.classifier.FileTemplate.Field
	3,  // 2: api.v1alpha1.classifier.FileTemplate.parse_opts:type_name -> api.v1alpha1.classifier.ParseOpts
	9,  // 3: api.v1alpha1.classifier.FileTemplate.constraints:type_name -> api.v1alpha1.classifier.Constraints
	2,  // 4: api.v1alpha1.classifier.FileTemplate.opts:type_name -> api.v1alpha1.classifier.Opts
	12, // 5: api.v1alpha1.classifier.Opts.date_formats:type_name -> api.v1alpha1.classifier.Opts.DateFormatsEntry
	13, // 6: api.v1alpha1.classifier.Opts.rename_fields:type_name -> api.v1alpha1.classifier.Opts.RenameFieldsEntry
	3,  // 7: api.v1alpha1.classifier.Opts.parse_opts:type_name -> api.v1alpha1.classifier.ParseOpts
	9,  // 8: api.v1alpha1.classifier.Opts.constraints:type_name -> api.v1alpha1.classifier.Constraints
	4,  // 9: api.v1alpha1.classifier.ParseOpts.csv:type_name -> api.v1alpha1.classifier.OptsCsv
	5,  // 10: api.v1alpha1.classifier.ParseOpts.json:type_name -> api.v1alpha1.classifier.OptsJson
	6,  // 11: api.v1alpha1.classifier.ParseOpts.jsonl:type_name -> api.v1alpha1.classifier.OptsJsonL
	7,  // 12: api.v1alpha1.classifier.ParseOpts.fixed:type_name -> api.v1alpha1.classifier.OptsFixed
	8,  // 13: api.v1alpha1.classifier.ParseOpts.parquet:type_name -> api.v1alpha1.classifier.OptsParquet
	15, // 14: api.v1alpha1.classifier.OptsFixed.positions:type_name -> api.v1alpha1.classifier.OptsFixed.PositionsEntry
	16, // 15: api.v1alpha1.classifier.Constraints.forbid:type_name -> api.v1alpha1.classifier.Constraints.ForbidEntry
	17, // 16: api.v1alpha1.classifier.Constraints.allow:type_name -> api.v1alpha1.classifier.Constraints.AllowEntry
	3,  // 17: api.v1alpha1.classifier.ParseHints.parse_opts:type_name -> api.v1alpha1.classifier.ParseOpts
	9,  // 18: api.v1alpha1.classifier.ParseHints.constraints:type_name -> api.v1alpha1.classifier.Constraints
	18, // 19: api.v1alpha1.classifier.FileTemplate.Field.entity_type:type_name -> api.commons.ClassifierEntityType
	14, // 20: api.v1alpha1.classifier.OptsFixed.PositionsEntry.value:type_name -> api.v1alpha1.classifier.OptsFixed.FieldOpts
	0,  // 21: api.v1alpha1.classifier.Constraints.ForbidEntry.value:type_name -> api.v1alpha1.classifier.ClassifierEntityTypes
	0,  // 22: api.v1alpha1.classifier.Constraints.AllowEntry.value:type_name -> api.v1alpha1.classifier.ClassifierEntityTypes
	23, // [23:23] is the sub-list for method output_type
	23, // [23:23] is the sub-list for method input_type
	23, // [23:23] is the sub-list for extension type_name
	23, // [23:23] is the sub-list for extension extendee
	0,  // [0:23] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_classifier_entities_proto_init() }
func file_api_v1alpha1_classifier_entities_proto_init() {
	if File_api_v1alpha1_classifier_entities_proto != nil {
		return
	}
	file_api_v1alpha1_classifier_entities_proto_msgTypes[3].OneofWrappers = []any{
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_classifier_entities_proto_rawDesc), len(file_api_v1alpha1_classifier_entities_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_classifier_entities_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_classifier_entities_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_classifier_entities_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_classifier_entities_proto = out.File
	file_api_v1alpha1_classifier_entities_proto_goTypes = nil
	file_api_v1alpha1_classifier_entities_proto_depIdxs = nil
}
