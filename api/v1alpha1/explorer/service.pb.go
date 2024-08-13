// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/v1alpha1/explorer/service.proto

package explorer

import (
	_ "github.com/tcncloud/api-go/annotations"
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

// ListDatasourceSchemasRequest is the request to list datasource schemas.
type ListDatasourceSchemasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// datasource_names is a table name filter for the datasources to list.
	DatasourceNames []string `protobuf:"bytes,1,rep,name=datasource_names,json=datasourceNames,proto3" json:"datasource_names,omitempty"`
	// datasource_type is the type of the datasource to list.
	// If not specified, all datasources will be listed.
	DatasourceType DatasourceType `protobuf:"varint,2,opt,name=datasource_type,json=datasourceType,proto3,enum=api.v1alpha1.explorer.DatasourceType" json:"datasource_type,omitempty"`
}

func (x *ListDatasourceSchemasRequest) Reset() {
	*x = ListDatasourceSchemasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_explorer_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasourceSchemasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasourceSchemasRequest) ProtoMessage() {}

func (x *ListDatasourceSchemasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_explorer_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasourceSchemasRequest.ProtoReflect.Descriptor instead.
func (*ListDatasourceSchemasRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_explorer_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListDatasourceSchemasRequest) GetDatasourceNames() []string {
	if x != nil {
		return x.DatasourceNames
	}
	return nil
}

func (x *ListDatasourceSchemasRequest) GetDatasourceType() DatasourceType {
	if x != nil {
		return x.DatasourceType
	}
	return DatasourceType_DATASOURCE_TYPE_UNSPECIFIED
}

// ListDatasourceSchemasResponse contains datasources and their schemas.
type ListDatasourceSchemasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of schemas
	Schemas []*Schema `protobuf:"bytes,1,rep,name=schemas,proto3" json:"schemas,omitempty"`
}

func (x *ListDatasourceSchemasResponse) Reset() {
	*x = ListDatasourceSchemasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_explorer_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatasourceSchemasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatasourceSchemasResponse) ProtoMessage() {}

func (x *ListDatasourceSchemasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_explorer_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatasourceSchemasResponse.ProtoReflect.Descriptor instead.
func (*ListDatasourceSchemasResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_explorer_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListDatasourceSchemasResponse) GetSchemas() []*Schema {
	if x != nil {
		return x.Schemas
	}
	return nil
}

// QueryRequest is the request to query a datasource.
type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// datasource_name is the name of the datasource to query.
	DatasourceName string `protobuf:"bytes,1,opt,name=datasource_name,json=datasourceName,proto3" json:"datasource_name,omitempty"`
	// datasource_type is the type of the datasource to query.
	DatasourceType DatasourceType `protobuf:"varint,2,opt,name=datasource_type,json=datasourceType,proto3,enum=api.v1alpha1.explorer.DatasourceType" json:"datasource_type,omitempty"`
	// query is the query to execute.
	//
	// Types that are assignable to Query:
	//
	//	*QueryRequest_Pipeline
	//	*QueryRequest_Prql
	Query isQueryRequest_Query `protobuf_oneof:"query"`
	// org_id for ownership of the data
	OrgIds []string `protobuf:"bytes,5,rep,name=org_ids,json=orgIds,proto3" json:"org_ids,omitempty"`
	// start_time is the start time of the query.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// end_time is the end time of the query.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// timezone is IESG timezone name
	// this is how the timezone is represented in the query
	Timezone string `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	// pipeline_parameters is the parameters for the pipeline.
	PipelineParameters *Parameters `protobuf:"bytes,9,opt,name=pipeline_parameters,json=pipelineParameters,proto3" json:"pipeline_parameters,omitempty"`
	// ui_trace_id is the trace id of the query.
	UiTraceId string `protobuf:"bytes,10,opt,name=ui_trace_id,json=uiTraceId,proto3" json:"ui_trace_id,omitempty"`
	// comment is the comment for the query.
	Comment string `protobuf:"bytes,11,opt,name=comment,proto3" json:"comment,omitempty"`
	// format is the format of the result.
	Format ExportFormat `protobuf:"varint,12,opt,name=format,proto3,enum=api.v1alpha1.explorer.ExportFormat" json:"format,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_explorer_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_explorer_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_explorer_service_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRequest) GetDatasourceName() string {
	if x != nil {
		return x.DatasourceName
	}
	return ""
}

func (x *QueryRequest) GetDatasourceType() DatasourceType {
	if x != nil {
		return x.DatasourceType
	}
	return DatasourceType_DATASOURCE_TYPE_UNSPECIFIED
}

func (m *QueryRequest) GetQuery() isQueryRequest_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *QueryRequest) GetPipeline() string {
	if x, ok := x.GetQuery().(*QueryRequest_Pipeline); ok {
		return x.Pipeline
	}
	return ""
}

func (x *QueryRequest) GetPrql() string {
	if x, ok := x.GetQuery().(*QueryRequest_Prql); ok {
		return x.Prql
	}
	return ""
}

func (x *QueryRequest) GetOrgIds() []string {
	if x != nil {
		return x.OrgIds
	}
	return nil
}

func (x *QueryRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *QueryRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *QueryRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *QueryRequest) GetPipelineParameters() *Parameters {
	if x != nil {
		return x.PipelineParameters
	}
	return nil
}

func (x *QueryRequest) GetUiTraceId() string {
	if x != nil {
		return x.UiTraceId
	}
	return ""
}

func (x *QueryRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *QueryRequest) GetFormat() ExportFormat {
	if x != nil {
		return x.Format
	}
	return ExportFormat_REPORT_FORMAT_UNSPECIFIED
}

type isQueryRequest_Query interface {
	isQueryRequest_Query()
}

type QueryRequest_Pipeline struct {
	// pipeline to be compiled to prql
	Pipeline string `protobuf:"bytes,3,opt,name=pipeline,proto3,oneof"`
}

type QueryRequest_Prql struct {
	// prql query to execute
	Prql string `protobuf:"bytes,4,opt,name=prql,proto3,oneof"`
}

func (*QueryRequest_Pipeline) isQueryRequest_Query() {}

func (*QueryRequest_Prql) isQueryRequest_Query() {}

// QueryResponse contains the result of a datasource query.
type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// result_url is the URL to the result of the query.
	ResultUrl string `protobuf:"bytes,1,opt,name=result_url,json=resultUrl,proto3" json:"result_url,omitempty"`
	// result_size_bytes is the size of the result in bytes.
	ResultSizeBytes int64 `protobuf:"varint,2,opt,name=result_size_bytes,json=resultSizeBytes,proto3" json:"result_size_bytes,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_explorer_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_explorer_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_explorer_service_proto_rawDescGZIP(), []int{3}
}

func (x *QueryResponse) GetResultUrl() string {
	if x != nil {
		return x.ResultUrl
	}
	return ""
}

func (x *QueryResponse) GetResultSizeBytes() int64 {
	if x != nil {
		return x.ResultSizeBytes
	}
	return 0
}

var File_api_v1alpha1_explorer_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_explorer_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x1a, 0x17, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x1c, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x65,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x58, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x22, 0xb6, 0x04, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x70, 0x72, 0x71, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x72, 0x71, 0x6c, 0x12, 0x17,
	0x0a, 0x07, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x52, 0x0a, 0x13, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x12, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x75, 0x69, 0x5f,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x69, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x42, 0x07, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x5a, 0x0a, 0x0d, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x69, 0x7a, 0x65,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x32, 0xed, 0x02, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xcb, 0x01, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x12, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47,
	0xba, 0xb8, 0x91, 0x02, 0x02, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3a, 0x3a, 0x01, 0x2a,
	0x22, 0x35, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65,
	0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x8b, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0xba, 0xb8,
	0x91, 0x02, 0x02, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0xd1, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x72, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x78, 0x70,
	0x6c, 0x6f, 0x72, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x45, 0xaa, 0x02, 0x15, 0x41, 0x70,
	0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x72, 0xca, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0xe2, 0x02, 0x21, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x45, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x17, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a,
	0x3a, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1alpha1_explorer_service_proto_rawDescOnce sync.Once
	file_api_v1alpha1_explorer_service_proto_rawDescData = file_api_v1alpha1_explorer_service_proto_rawDesc
)

func file_api_v1alpha1_explorer_service_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_explorer_service_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_explorer_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_explorer_service_proto_rawDescData)
	})
	return file_api_v1alpha1_explorer_service_proto_rawDescData
}

var file_api_v1alpha1_explorer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1alpha1_explorer_service_proto_goTypes = []any{
	(*ListDatasourceSchemasRequest)(nil),  // 0: api.v1alpha1.explorer.ListDatasourceSchemasRequest
	(*ListDatasourceSchemasResponse)(nil), // 1: api.v1alpha1.explorer.ListDatasourceSchemasResponse
	(*QueryRequest)(nil),                  // 2: api.v1alpha1.explorer.QueryRequest
	(*QueryResponse)(nil),                 // 3: api.v1alpha1.explorer.QueryResponse
	(DatasourceType)(0),                   // 4: api.v1alpha1.explorer.DatasourceType
	(*Schema)(nil),                        // 5: api.v1alpha1.explorer.Schema
	(*timestamppb.Timestamp)(nil),         // 6: google.protobuf.Timestamp
	(*Parameters)(nil),                    // 7: api.v1alpha1.explorer.Parameters
	(ExportFormat)(0),                     // 8: api.v1alpha1.explorer.ExportFormat
}
var file_api_v1alpha1_explorer_service_proto_depIdxs = []int32{
	4, // 0: api.v1alpha1.explorer.ListDatasourceSchemasRequest.datasource_type:type_name -> api.v1alpha1.explorer.DatasourceType
	5, // 1: api.v1alpha1.explorer.ListDatasourceSchemasResponse.schemas:type_name -> api.v1alpha1.explorer.Schema
	4, // 2: api.v1alpha1.explorer.QueryRequest.datasource_type:type_name -> api.v1alpha1.explorer.DatasourceType
	6, // 3: api.v1alpha1.explorer.QueryRequest.start_time:type_name -> google.protobuf.Timestamp
	6, // 4: api.v1alpha1.explorer.QueryRequest.end_time:type_name -> google.protobuf.Timestamp
	7, // 5: api.v1alpha1.explorer.QueryRequest.pipeline_parameters:type_name -> api.v1alpha1.explorer.Parameters
	8, // 6: api.v1alpha1.explorer.QueryRequest.format:type_name -> api.v1alpha1.explorer.ExportFormat
	0, // 7: api.v1alpha1.explorer.ExplorerService.ListDatasourceSchemas:input_type -> api.v1alpha1.explorer.ListDatasourceSchemasRequest
	2, // 8: api.v1alpha1.explorer.ExplorerService.Query:input_type -> api.v1alpha1.explorer.QueryRequest
	1, // 9: api.v1alpha1.explorer.ExplorerService.ListDatasourceSchemas:output_type -> api.v1alpha1.explorer.ListDatasourceSchemasResponse
	3, // 10: api.v1alpha1.explorer.ExplorerService.Query:output_type -> api.v1alpha1.explorer.QueryResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_explorer_service_proto_init() }
func file_api_v1alpha1_explorer_service_proto_init() {
	if File_api_v1alpha1_explorer_service_proto != nil {
		return
	}
	file_api_v1alpha1_explorer_entities_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1alpha1_explorer_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListDatasourceSchemasRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1alpha1_explorer_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListDatasourceSchemasResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1alpha1_explorer_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*QueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1alpha1_explorer_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*QueryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_v1alpha1_explorer_service_proto_msgTypes[2].OneofWrappers = []any{
		(*QueryRequest_Pipeline)(nil),
		(*QueryRequest_Prql)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_explorer_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_explorer_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_explorer_service_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_explorer_service_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_explorer_service_proto = out.File
	file_api_v1alpha1_explorer_service_proto_rawDesc = nil
	file_api_v1alpha1_explorer_service_proto_goTypes = nil
	file_api_v1alpha1_explorer_service_proto_depIdxs = nil
}
