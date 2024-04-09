// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/commons/classifier.proto

package commons

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

// these are the types that will be used for File Templates
type ClassifierEntityType int32

const (
	ClassifierEntityType_CET_UNKNOWN           ClassifierEntityType = 0
	ClassifierEntityType_CET_CREDIT_CARD       ClassifierEntityType = 1
	ClassifierEntityType_CET_CRYPTO            ClassifierEntityType = 2
	ClassifierEntityType_CET_DATE_TIME         ClassifierEntityType = 3
	ClassifierEntityType_CET_EMAIL_ADDRESS     ClassifierEntityType = 4
	ClassifierEntityType_CET_IBAN_CODE         ClassifierEntityType = 5
	ClassifierEntityType_CET_IP_ADDRESS        ClassifierEntityType = 6
	ClassifierEntityType_CET_NRP               ClassifierEntityType = 7
	ClassifierEntityType_CET_LOCATION          ClassifierEntityType = 8
	ClassifierEntityType_CET_PERSON            ClassifierEntityType = 9
	ClassifierEntityType_CET_PHONE_NUMBER      ClassifierEntityType = 10
	ClassifierEntityType_CET_MEDICAL_LICENSE   ClassifierEntityType = 11
	ClassifierEntityType_CET_URL               ClassifierEntityType = 12
	ClassifierEntityType_CET_US_BANK_NUMBER    ClassifierEntityType = 13
	ClassifierEntityType_CET_US_DRIVER_LICENSE ClassifierEntityType = 14
	ClassifierEntityType_CET_US_ITIN           ClassifierEntityType = 15
	ClassifierEntityType_CET_US_PASSPORT       ClassifierEntityType = 16
	ClassifierEntityType_CET_US_SSN            ClassifierEntityType = 17
	ClassifierEntityType_CET_POSTAL_CODE       ClassifierEntityType = 18
)

// Enum value maps for ClassifierEntityType.
var (
	ClassifierEntityType_name = map[int32]string{
		0:  "CET_UNKNOWN",
		1:  "CET_CREDIT_CARD",
		2:  "CET_CRYPTO",
		3:  "CET_DATE_TIME",
		4:  "CET_EMAIL_ADDRESS",
		5:  "CET_IBAN_CODE",
		6:  "CET_IP_ADDRESS",
		7:  "CET_NRP",
		8:  "CET_LOCATION",
		9:  "CET_PERSON",
		10: "CET_PHONE_NUMBER",
		11: "CET_MEDICAL_LICENSE",
		12: "CET_URL",
		13: "CET_US_BANK_NUMBER",
		14: "CET_US_DRIVER_LICENSE",
		15: "CET_US_ITIN",
		16: "CET_US_PASSPORT",
		17: "CET_US_SSN",
		18: "CET_POSTAL_CODE",
	}
	ClassifierEntityType_value = map[string]int32{
		"CET_UNKNOWN":           0,
		"CET_CREDIT_CARD":       1,
		"CET_CRYPTO":            2,
		"CET_DATE_TIME":         3,
		"CET_EMAIL_ADDRESS":     4,
		"CET_IBAN_CODE":         5,
		"CET_IP_ADDRESS":        6,
		"CET_NRP":               7,
		"CET_LOCATION":          8,
		"CET_PERSON":            9,
		"CET_PHONE_NUMBER":      10,
		"CET_MEDICAL_LICENSE":   11,
		"CET_URL":               12,
		"CET_US_BANK_NUMBER":    13,
		"CET_US_DRIVER_LICENSE": 14,
		"CET_US_ITIN":           15,
		"CET_US_PASSPORT":       16,
		"CET_US_SSN":            17,
		"CET_POSTAL_CODE":       18,
	}
)

func (x ClassifierEntityType) Enum() *ClassifierEntityType {
	p := new(ClassifierEntityType)
	*p = x
	return p
}

func (x ClassifierEntityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClassifierEntityType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_commons_classifier_proto_enumTypes[0].Descriptor()
}

func (ClassifierEntityType) Type() protoreflect.EnumType {
	return &file_api_commons_classifier_proto_enumTypes[0]
}

func (x ClassifierEntityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClassifierEntityType.Descriptor instead.
func (ClassifierEntityType) EnumDescriptor() ([]byte, []int) {
	return file_api_commons_classifier_proto_rawDescGZIP(), []int{0}
}

var File_api_commons_classifier_proto protoreflect.FileDescriptor

var file_api_commons_classifier_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2a, 0x86, 0x03, 0x0a, 0x14,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x45, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x45, 0x54, 0x5f, 0x43, 0x52, 0x45,
	0x44, 0x49, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x45,
	0x54, 0x5f, 0x43, 0x52, 0x59, 0x50, 0x54, 0x4f, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x45,
	0x54, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x15, 0x0a,
	0x11, 0x43, 0x45, 0x54, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x45, 0x54, 0x5f, 0x49, 0x42, 0x41, 0x4e,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x45, 0x54, 0x5f, 0x49,
	0x50, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x45, 0x54, 0x5f, 0x4e, 0x52, 0x50, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x45, 0x54, 0x5f,
	0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x45,
	0x54, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x45,
	0x54, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x0a,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x45, 0x54, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x43, 0x41, 0x4c, 0x5f,
	0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x0b, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x45, 0x54,
	0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0c, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x45, 0x54, 0x5f, 0x55, 0x53,
	0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x0d, 0x12, 0x19,
	0x0a, 0x15, 0x43, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x52, 0x5f,
	0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x0e, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x45, 0x54,
	0x5f, 0x55, 0x53, 0x5f, 0x49, 0x54, 0x49, 0x4e, 0x10, 0x0f, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x45,
	0x54, 0x5f, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x10, 0x12,
	0x0e, 0x0a, 0x0a, 0x43, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x5f, 0x53, 0x53, 0x4e, 0x10, 0x11, 0x12,
	0x13, 0x0a, 0x0f, 0x43, 0x45, 0x54, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x41, 0x4c, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x10, 0x12, 0x42, 0x97, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x42, 0x0f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0c, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commons_classifier_proto_rawDescOnce sync.Once
	file_api_commons_classifier_proto_rawDescData = file_api_commons_classifier_proto_rawDesc
)

func file_api_commons_classifier_proto_rawDescGZIP() []byte {
	file_api_commons_classifier_proto_rawDescOnce.Do(func() {
		file_api_commons_classifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commons_classifier_proto_rawDescData)
	})
	return file_api_commons_classifier_proto_rawDescData
}

var file_api_commons_classifier_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_commons_classifier_proto_goTypes = []interface{}{
	(ClassifierEntityType)(0), // 0: api.commons.ClassifierEntityType
}
var file_api_commons_classifier_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commons_classifier_proto_init() }
func file_api_commons_classifier_proto_init() {
	if File_api_commons_classifier_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_commons_classifier_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_classifier_proto_goTypes,
		DependencyIndexes: file_api_commons_classifier_proto_depIdxs,
		EnumInfos:         file_api_commons_classifier_proto_enumTypes,
	}.Build()
	File_api_commons_classifier_proto = out.File
	file_api_commons_classifier_proto_rawDesc = nil
	file_api_commons_classifier_proto_goTypes = nil
	file_api_commons_classifier_proto_depIdxs = nil
}
