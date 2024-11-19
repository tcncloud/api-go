// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/v1alpha1/explorer/entities.proto

package explorer

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

// ExportFormat is an enum for the format of a report.
type ExportFormat int32

const (
	ExportFormat_REPORT_FORMAT_UNSPECIFIED ExportFormat = 0
	ExportFormat_REPORT_FORMAT_CSV         ExportFormat = 1
	ExportFormat_REPORT_FORMAT_PARQUET     ExportFormat = 2
)

// Enum value maps for ExportFormat.
var (
	ExportFormat_name = map[int32]string{
		0: "REPORT_FORMAT_UNSPECIFIED",
		1: "REPORT_FORMAT_CSV",
		2: "REPORT_FORMAT_PARQUET",
	}
	ExportFormat_value = map[string]int32{
		"REPORT_FORMAT_UNSPECIFIED": 0,
		"REPORT_FORMAT_CSV":         1,
		"REPORT_FORMAT_PARQUET":     2,
	}
)

func (x ExportFormat) Enum() *ExportFormat {
	p := new(ExportFormat)
	*p = x
	return p
}

func (x ExportFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExportFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1alpha1_explorer_entities_proto_enumTypes[0].Descriptor()
}

func (ExportFormat) Type() protoreflect.EnumType {
	return &file_api_v1alpha1_explorer_entities_proto_enumTypes[0]
}

func (x ExportFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExportFormat.Descriptor instead.
func (ExportFormat) EnumDescriptor() ([]byte, []int) {
	return file_api_v1alpha1_explorer_entities_proto_rawDescGZIP(), []int{0}
}

// SchemaType is an enum for the type of a schema field.
type SchemaType int32

const (
	SchemaType_SCHEMA_TYPE_UNSPECIFIED  SchemaType = 0
	SchemaType_SCHEMA_TYPE_INT          SchemaType = 2
	SchemaType_SCHEMA_TYPE_FLOAT        SchemaType = 3
	SchemaType_SCHEMA_TYPE_STRING       SchemaType = 5
	SchemaType_SCHEMA_TYPE_BOOL         SchemaType = 6
	SchemaType_SCHEMA_TYPE_TIMESTAMP    SchemaType = 7
	SchemaType_SCHEMA_TYPE_INT_ARRAY    SchemaType = 8
	SchemaType_SCHEMA_TYPE_FLOAT_ARRAY  SchemaType = 9
	SchemaType_SCHEMA_TYPE_STRING_ARRAY SchemaType = 10
	SchemaType_SCHEMA_TYPE_BOOL_ARRAY   SchemaType = 11
	SchemaType_SCHEMA_TYPE_MAP          SchemaType = 12
)

// Enum value maps for SchemaType.
var (
	SchemaType_name = map[int32]string{
		0:  "SCHEMA_TYPE_UNSPECIFIED",
		2:  "SCHEMA_TYPE_INT",
		3:  "SCHEMA_TYPE_FLOAT",
		5:  "SCHEMA_TYPE_STRING",
		6:  "SCHEMA_TYPE_BOOL",
		7:  "SCHEMA_TYPE_TIMESTAMP",
		8:  "SCHEMA_TYPE_INT_ARRAY",
		9:  "SCHEMA_TYPE_FLOAT_ARRAY",
		10: "SCHEMA_TYPE_STRING_ARRAY",
		11: "SCHEMA_TYPE_BOOL_ARRAY",
		12: "SCHEMA_TYPE_MAP",
	}
	SchemaType_value = map[string]int32{
		"SCHEMA_TYPE_UNSPECIFIED":  0,
		"SCHEMA_TYPE_INT":          2,
		"SCHEMA_TYPE_FLOAT":        3,
		"SCHEMA_TYPE_STRING":       5,
		"SCHEMA_TYPE_BOOL":         6,
		"SCHEMA_TYPE_TIMESTAMP":    7,
		"SCHEMA_TYPE_INT_ARRAY":    8,
		"SCHEMA_TYPE_FLOAT_ARRAY":  9,
		"SCHEMA_TYPE_STRING_ARRAY": 10,
		"SCHEMA_TYPE_BOOL_ARRAY":   11,
		"SCHEMA_TYPE_MAP":          12,
	}
)

func (x SchemaType) Enum() *SchemaType {
	p := new(SchemaType)
	*p = x
	return p
}

func (x SchemaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SchemaType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1alpha1_explorer_entities_proto_enumTypes[1].Descriptor()
}

func (SchemaType) Type() protoreflect.EnumType {
	return &file_api_v1alpha1_explorer_entities_proto_enumTypes[1]
}

func (x SchemaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SchemaType.Descriptor instead.
func (SchemaType) EnumDescriptor() ([]byte, []int) {
	return file_api_v1alpha1_explorer_entities_proto_rawDescGZIP(), []int{1}
}

// DatasourceType is an enum for the type of a datasource.
type DatasourceType int32

const (
	DatasourceType_DATASOURCE_TYPE_UNSPECIFIED DatasourceType = 0
	DatasourceType_DATASOURCE_TYPE_VFS         DatasourceType = 1
	DatasourceType_DATASOURCE_TYPE_CLICKHOUSE  DatasourceType = 2
)

// Enum value maps for DatasourceType.
var (
	DatasourceType_name = map[int32]string{
		0: "DATASOURCE_TYPE_UNSPECIFIED",
		1: "DATASOURCE_TYPE_VFS",
		2: "DATASOURCE_TYPE_CLICKHOUSE",
	}
	DatasourceType_value = map[string]int32{
		"DATASOURCE_TYPE_UNSPECIFIED": 0,
		"DATASOURCE_TYPE_VFS":         1,
		"DATASOURCE_TYPE_CLICKHOUSE":  2,
	}
)

func (x DatasourceType) Enum() *DatasourceType {
	p := new(DatasourceType)
	*p = x
	return p
}

func (x DatasourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatasourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1alpha1_explorer_entities_proto_enumTypes[2].Descriptor()
}

func (DatasourceType) Type() protoreflect.EnumType {
	return &file_api_v1alpha1_explorer_entities_proto_enumTypes[2]
}

func (x DatasourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DatasourceType.Descriptor instead.
func (DatasourceType) EnumDescriptor() ([]byte, []int) {
	return file_api_v1alpha1_explorer_entities_proto_rawDescGZIP(), []int{2}
}

// SchemaField is a field in a schema.
type SchemaField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ColumnType          SchemaType `protobuf:"varint,2,opt,name=column_type,json=columnType,proto3,enum=api.v1alpha1.explorer.SchemaType" json:"column_type,omitempty"`
	IsPrimaryKey        bool       `protobuf:"varint,3,opt,name=is_primary_key,json=isPrimaryKey,proto3" json:"is_primary_key,omitempty"`
	IsLowCardinality    bool       `protobuf:"varint,4,opt,name=is_low_cardinality,json=isLowCardinality,proto3" json:"is_low_cardinality,omitempty"`
	ColumnDescription   string     `protobuf:"bytes,5,opt,name=column_description,json=columnDescription,proto3" json:"column_description,omitempty"`
	IsTimeFilter        bool       `protobuf:"varint,6,opt,name=is_time_filter,json=isTimeFilter,proto3" json:"is_time_filter,omitempty"`
	IsDefaultTimeFilter bool       `protobuf:"varint,7,opt,name=is_default_time_filter,json=isDefaultTimeFilter,proto3" json:"is_default_time_filter,omitempty"`
}

func (x *SchemaField) Reset() {
	*x = SchemaField{}
	mi := &file_api_v1alpha1_explorer_entities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SchemaField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaField) ProtoMessage() {}

func (x *SchemaField) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_explorer_entities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaField.ProtoReflect.Descriptor instead.
func (*SchemaField) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_explorer_entities_proto_rawDescGZIP(), []int{0}
}

func (x *SchemaField) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SchemaField) GetColumnType() SchemaType {
	if x != nil {
		return x.ColumnType
	}
	return SchemaType_SCHEMA_TYPE_UNSPECIFIED
}

func (x *SchemaField) GetIsPrimaryKey() bool {
	if x != nil {
		return x.IsPrimaryKey
	}
	return false
}

func (x *SchemaField) GetIsLowCardinality() bool {
	if x != nil {
		return x.IsLowCardinality
	}
	return false
}

func (x *SchemaField) GetColumnDescription() string {
	if x != nil {
		return x.ColumnDescription
	}
	return ""
}

func (x *SchemaField) GetIsTimeFilter() bool {
	if x != nil {
		return x.IsTimeFilter
	}
	return false
}

func (x *SchemaField) GetIsDefaultTimeFilter() bool {
	if x != nil {
		return x.IsDefaultTimeFilter
	}
	return false
}

// Schema is a schema for a datasource.
type Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DatasourceType   DatasourceType `protobuf:"varint,2,opt,name=datasource_type,json=datasourceType,proto3,enum=api.v1alpha1.explorer.DatasourceType" json:"datasource_type,omitempty"`
	Fields           []*SchemaField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	TableDescription string         `protobuf:"bytes,4,opt,name=table_description,json=tableDescription,proto3" json:"table_description,omitempty"`
	Category         string         `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	SubCategory      string         `protobuf:"bytes,6,opt,name=sub_category,json=subCategory,proto3" json:"sub_category,omitempty"`
}

func (x *Schema) Reset() {
	*x = Schema{}
	mi := &file_api_v1alpha1_explorer_entities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_explorer_entities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_explorer_entities_proto_rawDescGZIP(), []int{1}
}

func (x *Schema) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Schema) GetDatasourceType() DatasourceType {
	if x != nil {
		return x.DatasourceType
	}
	return DatasourceType_DATASOURCE_TYPE_UNSPECIFIED
}

func (x *Schema) GetFields() []*SchemaField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Schema) GetTableDescription() string {
	if x != nil {
		return x.TableDescription
	}
	return ""
}

func (x *Schema) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Schema) GetSubCategory() string {
	if x != nil {
		return x.SubCategory
	}
	return ""
}

// Parameter is a parameter for a query.
type Parameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters map[string]*Parameters_Parameter `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Parameters) Reset() {
	*x = Parameters{}
	mi := &file_api_v1alpha1_explorer_entities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Parameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameters) ProtoMessage() {}

func (x *Parameters) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_explorer_entities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameters.ProtoReflect.Descriptor instead.
func (*Parameters) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_explorer_entities_proto_rawDescGZIP(), []int{2}
}

func (x *Parameters) GetParameters() map[string]*Parameters_Parameter {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type Parameters_Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	DataType string `protobuf:"bytes,3,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
}

func (x *Parameters_Parameter) Reset() {
	*x = Parameters_Parameter{}
	mi := &file_api_v1alpha1_explorer_entities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Parameters_Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameters_Parameter) ProtoMessage() {}

func (x *Parameters_Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_explorer_entities_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameters_Parameter.ProtoReflect.Descriptor instead.
func (*Parameters_Parameter) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_explorer_entities_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Parameters_Parameter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Parameters_Parameter) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

var File_api_v1alpha1_explorer_entities_proto protoreflect.FileDescriptor

var file_api_v1alpha1_explorer_entities_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x22, 0xc3, 0x02,
	0x0a, 0x0b, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x42, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x69,
	0x73, 0x5f, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x4c, 0x6f, 0x77, 0x43, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x69, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x33,
	0x0a, 0x16, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13,
	0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x94, 0x02, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2b,
	0x0a, 0x11, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x8b, 0x02, 0x0a, 0x0a, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x51, 0x0a, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x65, 0x78, 0x70,
	0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3e, 0x0a, 0x09,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x6a, 0x0a, 0x0f,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x41, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x5f, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x50, 0x4f,
	0x52, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x50, 0x4f, 0x52,
	0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x43, 0x53, 0x56, 0x10, 0x01, 0x12, 0x19,
	0x0a, 0x15, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f,
	0x50, 0x41, 0x52, 0x51, 0x55, 0x45, 0x54, 0x10, 0x02, 0x2a, 0xa5, 0x02, 0x0a, 0x0a, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x43, 0x48, 0x45,
	0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x43,
	0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10,
	0x03, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x43, 0x48,
	0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x06, 0x12,
	0x19, 0x0a, 0x15, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x07, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x43,
	0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x5f, 0x41, 0x52,
	0x52, 0x41, 0x59, 0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59,
	0x10, 0x09, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x0a,
	0x12, 0x1a, 0x0a, 0x16, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x42, 0x4f, 0x4f, 0x4c, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x0b, 0x12, 0x13, 0x0a, 0x0f,
	0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x50, 0x10,
	0x0c, 0x2a, 0x6a, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x41, 0x54, 0x41, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x41, 0x54, 0x41, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x46, 0x53, 0x10, 0x01, 0x12, 0x1e, 0x0a,
	0x1a, 0x44, 0x41, 0x54, 0x41, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x48, 0x4f, 0x55, 0x53, 0x45, 0x10, 0x02, 0x42, 0xd2, 0x01,
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x42, 0x0d, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0xa2, 0x02,
	0x03, 0x41, 0x56, 0x45, 0xaa, 0x02, 0x15, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0xca, 0x02, 0x15, 0x41,
	0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x45, 0x78, 0x70, 0x6c,
	0x6f, 0x72, 0x65, 0x72, 0xe2, 0x02, 0x21, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5c, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1alpha1_explorer_entities_proto_rawDescOnce sync.Once
	file_api_v1alpha1_explorer_entities_proto_rawDescData = file_api_v1alpha1_explorer_entities_proto_rawDesc
)

func file_api_v1alpha1_explorer_entities_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_explorer_entities_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_explorer_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_explorer_entities_proto_rawDescData)
	})
	return file_api_v1alpha1_explorer_entities_proto_rawDescData
}

var file_api_v1alpha1_explorer_entities_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_v1alpha1_explorer_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1alpha1_explorer_entities_proto_goTypes = []any{
	(ExportFormat)(0),            // 0: api.v1alpha1.explorer.ExportFormat
	(SchemaType)(0),              // 1: api.v1alpha1.explorer.SchemaType
	(DatasourceType)(0),          // 2: api.v1alpha1.explorer.DatasourceType
	(*SchemaField)(nil),          // 3: api.v1alpha1.explorer.SchemaField
	(*Schema)(nil),               // 4: api.v1alpha1.explorer.Schema
	(*Parameters)(nil),           // 5: api.v1alpha1.explorer.Parameters
	(*Parameters_Parameter)(nil), // 6: api.v1alpha1.explorer.Parameters.Parameter
	nil,                          // 7: api.v1alpha1.explorer.Parameters.ParametersEntry
}
var file_api_v1alpha1_explorer_entities_proto_depIdxs = []int32{
	1, // 0: api.v1alpha1.explorer.SchemaField.column_type:type_name -> api.v1alpha1.explorer.SchemaType
	2, // 1: api.v1alpha1.explorer.Schema.datasource_type:type_name -> api.v1alpha1.explorer.DatasourceType
	3, // 2: api.v1alpha1.explorer.Schema.fields:type_name -> api.v1alpha1.explorer.SchemaField
	7, // 3: api.v1alpha1.explorer.Parameters.parameters:type_name -> api.v1alpha1.explorer.Parameters.ParametersEntry
	6, // 4: api.v1alpha1.explorer.Parameters.ParametersEntry.value:type_name -> api.v1alpha1.explorer.Parameters.Parameter
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_explorer_entities_proto_init() }
func file_api_v1alpha1_explorer_entities_proto_init() {
	if File_api_v1alpha1_explorer_entities_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_explorer_entities_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_explorer_entities_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_explorer_entities_proto_depIdxs,
		EnumInfos:         file_api_v1alpha1_explorer_entities_proto_enumTypes,
		MessageInfos:      file_api_v1alpha1_explorer_entities_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_explorer_entities_proto = out.File
	file_api_v1alpha1_explorer_entities_proto_rawDesc = nil
	file_api_v1alpha1_explorer_entities_proto_goTypes = nil
	file_api_v1alpha1_explorer_entities_proto_depIdxs = nil
}
