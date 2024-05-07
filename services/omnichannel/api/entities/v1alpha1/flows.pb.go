// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: services/omnichannel/api/entities/v1alpha1/flows.proto

package entitiesv1alpha1

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

type FlowFieldName int32

const (
	FlowFieldName_FIELD_NAME_UNSPECIFIED FlowFieldName = 0
	FlowFieldName_FIELD_NAME_USER_INPUT  FlowFieldName = 1
)

// Enum value maps for FlowFieldName.
var (
	FlowFieldName_name = map[int32]string{
		0: "FIELD_NAME_UNSPECIFIED",
		1: "FIELD_NAME_USER_INPUT",
	}
	FlowFieldName_value = map[string]int32{
		"FIELD_NAME_UNSPECIFIED": 0,
		"FIELD_NAME_USER_INPUT":  1,
	}
)

func (x FlowFieldName) Enum() *FlowFieldName {
	p := new(FlowFieldName)
	*p = x
	return p
}

func (x FlowFieldName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlowFieldName) Descriptor() protoreflect.EnumDescriptor {
	return file_services_omnichannel_api_entities_v1alpha1_flows_proto_enumTypes[0].Descriptor()
}

func (FlowFieldName) Type() protoreflect.EnumType {
	return &file_services_omnichannel_api_entities_v1alpha1_flows_proto_enumTypes[0]
}

func (x FlowFieldName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FlowFieldName.Descriptor instead.
func (FlowFieldName) EnumDescriptor() ([]byte, []int) {
	return file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDescGZIP(), []int{0}
}

type FlowPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []*FlowField `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *FlowPayload) Reset() {
	*x = FlowPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_omnichannel_api_entities_v1alpha1_flows_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowPayload) ProtoMessage() {}

func (x *FlowPayload) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_api_entities_v1alpha1_flows_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowPayload.ProtoReflect.Descriptor instead.
func (*FlowPayload) Descriptor() ([]byte, []int) {
	return file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDescGZIP(), []int{0}
}

func (x *FlowPayload) GetFields() []*FlowField {
	if x != nil {
		return x.Fields
	}
	return nil
}

type FlowField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name FlowFieldName `protobuf:"varint,1,opt,name=name,proto3,enum=services.omnichannel.api.entities.v1alpha1.FlowFieldName" json:"name,omitempty"`
	// Types that are assignable to Value:
	//
	//	*FlowField_UserInput
	Value isFlowField_Value `protobuf_oneof:"value"`
}

func (x *FlowField) Reset() {
	*x = FlowField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_omnichannel_api_entities_v1alpha1_flows_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowField) ProtoMessage() {}

func (x *FlowField) ProtoReflect() protoreflect.Message {
	mi := &file_services_omnichannel_api_entities_v1alpha1_flows_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowField.ProtoReflect.Descriptor instead.
func (*FlowField) Descriptor() ([]byte, []int) {
	return file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDescGZIP(), []int{1}
}

func (x *FlowField) GetName() FlowFieldName {
	if x != nil {
		return x.Name
	}
	return FlowFieldName_FIELD_NAME_UNSPECIFIED
}

func (m *FlowField) GetValue() isFlowField_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *FlowField) GetUserInput() string {
	if x, ok := x.GetValue().(*FlowField_UserInput); ok {
		return x.UserInput
	}
	return ""
}

type isFlowField_Value interface {
	isFlowField_Value()
}

type FlowField_UserInput struct {
	UserInput string `protobuf:"bytes,100,opt,name=user_input,json=userInput,proto3,oneof"`
}

func (*FlowField_UserInput) isFlowField_Value() {}

var File_services_omnichannel_api_entities_v1alpha1_flows_proto protoreflect.FileDescriptor

var file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDesc = []byte{
	0x0a, 0x36, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x22, 0x5c, 0x0a, 0x0b, 0x46, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x4d, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f,
	0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x09, 0x46, 0x6c, 0x6f, 0x77, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x4d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x6c, 0x6f, 0x77,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x46, 0x0a, 0x0d, 0x46, 0x6c, 0x6f,
	0x77, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x10,
	0x01, 0x42, 0xe1, 0x02, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x46, 0x6c, 0x6f, 0x77, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x04, 0x53, 0x4f, 0x41,
	0x45, 0xaa, 0x02, 0x2a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x6d, 0x6e,
	0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02,
	0x2a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x36, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a,
	0x3a, 0x4f, 0x6d, 0x6e, 0x69, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDescOnce sync.Once
	file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDescData = file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDesc
)

func file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDescGZIP() []byte {
	file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDescOnce.Do(func() {
		file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDescData)
	})
	return file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDescData
}

var file_services_omnichannel_api_entities_v1alpha1_flows_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_omnichannel_api_entities_v1alpha1_flows_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_omnichannel_api_entities_v1alpha1_flows_proto_goTypes = []interface{}{
	(FlowFieldName)(0),  // 0: services.omnichannel.api.entities.v1alpha1.FlowFieldName
	(*FlowPayload)(nil), // 1: services.omnichannel.api.entities.v1alpha1.FlowPayload
	(*FlowField)(nil),   // 2: services.omnichannel.api.entities.v1alpha1.FlowField
}
var file_services_omnichannel_api_entities_v1alpha1_flows_proto_depIdxs = []int32{
	2, // 0: services.omnichannel.api.entities.v1alpha1.FlowPayload.fields:type_name -> services.omnichannel.api.entities.v1alpha1.FlowField
	0, // 1: services.omnichannel.api.entities.v1alpha1.FlowField.name:type_name -> services.omnichannel.api.entities.v1alpha1.FlowFieldName
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_services_omnichannel_api_entities_v1alpha1_flows_proto_init() }
func file_services_omnichannel_api_entities_v1alpha1_flows_proto_init() {
	if File_services_omnichannel_api_entities_v1alpha1_flows_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_omnichannel_api_entities_v1alpha1_flows_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowPayload); i {
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
		file_services_omnichannel_api_entities_v1alpha1_flows_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowField); i {
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
	file_services_omnichannel_api_entities_v1alpha1_flows_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*FlowField_UserInput)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_omnichannel_api_entities_v1alpha1_flows_proto_goTypes,
		DependencyIndexes: file_services_omnichannel_api_entities_v1alpha1_flows_proto_depIdxs,
		EnumInfos:         file_services_omnichannel_api_entities_v1alpha1_flows_proto_enumTypes,
		MessageInfos:      file_services_omnichannel_api_entities_v1alpha1_flows_proto_msgTypes,
	}.Build()
	File_services_omnichannel_api_entities_v1alpha1_flows_proto = out.File
	file_services_omnichannel_api_entities_v1alpha1_flows_proto_rawDesc = nil
	file_services_omnichannel_api_entities_v1alpha1_flows_proto_goTypes = nil
	file_services_omnichannel_api_entities_v1alpha1_flows_proto_depIdxs = nil
}
