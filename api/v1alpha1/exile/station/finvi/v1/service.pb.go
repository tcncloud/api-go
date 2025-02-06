// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/v1alpha1/exile/station/finvi/v1/service.proto

package finviv1

import (
	_ "github.com/tcncloud/api-go/annotations"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type PopAccountReq struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	RecordId string                 `protobuf:"bytes,10,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	UserId   string                 `protobuf:"bytes,11,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// optional
	AlternateId string `protobuf:"bytes,12,opt,name=alternate_id,json=alternateId,proto3" json:"alternate_id,omitempty"`
	// (call_sid, call_type) is primary key
	CallSid       int64  `protobuf:"varint,13,opt,name=call_sid,json=callSid,proto3" json:"call_sid,omitempty"`
	CallType      string `protobuf:"bytes,14,opt,name=call_type,json=callType,proto3" json:"call_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PopAccountReq) Reset() {
	*x = PopAccountReq{}
	mi := &file_api_v1alpha1_exile_station_finvi_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PopAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopAccountReq) ProtoMessage() {}

func (x *PopAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_exile_station_finvi_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopAccountReq.ProtoReflect.Descriptor instead.
func (*PopAccountReq) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *PopAccountReq) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *PopAccountReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PopAccountReq) GetAlternateId() string {
	if x != nil {
		return x.AlternateId
	}
	return ""
}

func (x *PopAccountReq) GetCallSid() int64 {
	if x != nil {
		return x.CallSid
	}
	return 0
}

func (x *PopAccountReq) GetCallType() string {
	if x != nil {
		return x.CallType
	}
	return ""
}

type PopAccountRes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	JobId int64                  `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// if true only job_id and async will be populated.
	// all other data on the response will be sent as jsonl to the async handler
	Async         bool `protobuf:"varint,2,opt,name=async,proto3" json:"async,omitempty"` // TBD
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PopAccountRes) Reset() {
	*x = PopAccountRes{}
	mi := &file_api_v1alpha1_exile_station_finvi_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PopAccountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopAccountRes) ProtoMessage() {}

func (x *PopAccountRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_exile_station_finvi_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopAccountRes.ProtoReflect.Descriptor instead.
func (*PopAccountRes) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *PopAccountRes) GetJobId() int64 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (x *PopAccountRes) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

var File_api_v1alpha1_exile_station_finvi_v1_service_proto protoreflect.FileDescriptor

var file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDesc = string([]byte{
	0x0a, 0x31, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65,
	0x78, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x6e,
	0x76, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x23, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x66, 0x69, 0x6e, 0x76, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xac, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x61,
	0x6c, 0x6c, 0x53, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x3c,
	0x0a, 0x0d, 0x50, 0x6f, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x32, 0xe6, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x46, 0x69, 0x6e, 0x76, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xce, 0x01, 0x0a, 0x0a, 0x50, 0x6f, 0x70, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x66, 0x69, 0x6e, 0x76, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x70, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6e, 0x76, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x58, 0xba, 0xb8, 0x91,
	0x02, 0x05, 0x0a, 0x03, 0x08, 0xac, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x48, 0x3a, 0x01, 0x2a,
	0x22, 0x43, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x65, 0x78, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x69,
	0x6e, 0x76, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x66, 0x69,
	0x6e, 0x76, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6f, 0x70, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xb2, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x65, 0x78, 0x69, 0x6c, 0x65,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6e, 0x76, 0x69, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x78, 0x69, 0x6c, 0x65,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x6e, 0x76, 0x69, 0x2f, 0x76,
	0x31, 0x3b, 0x66, 0x69, 0x6e, 0x76, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x41, 0x56, 0x45, 0x53,
	0x46, 0xaa, 0x02, 0x23, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46,
	0x69, 0x6e, 0x76, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x5c, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x46, 0x69, 0x6e, 0x76, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2f,
	0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x45, 0x78, 0x69,
	0x6c, 0x65, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x46, 0x69, 0x6e, 0x76, 0x69,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x28, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a,
	0x3a, 0x45, 0x78, 0x69, 0x6c, 0x65, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x3a, 0x46, 0x69, 0x6e, 0x76, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDescOnce sync.Once
	file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDescData []byte
)

func file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDesc), len(file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDesc)))
	})
	return file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDescData
}

var file_api_v1alpha1_exile_station_finvi_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1alpha1_exile_station_finvi_v1_service_proto_goTypes = []any{
	(*PopAccountReq)(nil), // 0: api.v1alpha1.exile.station.finvi.v1.PopAccountReq
	(*PopAccountRes)(nil), // 1: api.v1alpha1.exile.station.finvi.v1.PopAccountRes
}
var file_api_v1alpha1_exile_station_finvi_v1_service_proto_depIdxs = []int32{
	0, // 0: api.v1alpha1.exile.station.finvi.v1.GenericFinviService.PopAccount:input_type -> api.v1alpha1.exile.station.finvi.v1.PopAccountReq
	1, // 1: api.v1alpha1.exile.station.finvi.v1.GenericFinviService.PopAccount:output_type -> api.v1alpha1.exile.station.finvi.v1.PopAccountRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_exile_station_finvi_v1_service_proto_init() }
func file_api_v1alpha1_exile_station_finvi_v1_service_proto_init() {
	if File_api_v1alpha1_exile_station_finvi_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDesc), len(file_api_v1alpha1_exile_station_finvi_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1alpha1_exile_station_finvi_v1_service_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_exile_station_finvi_v1_service_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_exile_station_finvi_v1_service_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_exile_station_finvi_v1_service_proto = out.File
	file_api_v1alpha1_exile_station_finvi_v1_service_proto_goTypes = nil
	file_api_v1alpha1_exile_station_finvi_v1_service_proto_depIdxs = nil
}
