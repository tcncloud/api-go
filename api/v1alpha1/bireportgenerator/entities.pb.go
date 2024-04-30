// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: api/v1alpha1/bireportgenerator/entities.proto

package bireportgenerator

import (
	commons "github.com/tcncloud/api-go/api/commons"
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

// ReportJob that can be scheduled to report dashboard data.
type ReportJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique identifier for the report job
	ReportJobId string `protobuf:"bytes,1,opt,name=report_job_id,json=reportJobId,proto3" json:"report_job_id,omitempty"`
	// name of the report job
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// description of the report job
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// dashboard id to report
	//
	// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
	DashboardId string `protobuf:"bytes,4,opt,name=dashboard_id,json=dashboardId,proto3" json:"dashboard_id,omitempty"`
	// time zone to use for the report schedule
	//
	// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
	TimeZone string `protobuf:"bytes,5,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// time period of data to report with
	TimePeriod commons.TimePeriod `protobuf:"varint,6,opt,name=time_period,json=timePeriod,proto3,enum=api.commons.TimePeriod" json:"time_period,omitempty"`
	// delivery times for the report
	//
	// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
	DeliveryTimes *commons.DeliveryTimes `protobuf:"bytes,7,opt,name=delivery_times,json=deliveryTimes,proto3" json:"delivery_times,omitempty"`
	// days filter to report on
	//
	// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
	DayFilter *commons.DayFilter `protobuf:"bytes,8,opt,name=day_filter,json=dayFilter,proto3" json:"day_filter,omitempty"`
	// months filter to report on
	//
	// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
	Months []commons.Month `protobuf:"varint,9,rep,packed,name=months,proto3,enum=api.commons.Month" json:"months,omitempty"`
	// format options for the report
	FormatOptions *commons.FormatOptions `protobuf:"bytes,10,opt,name=format_options,json=formatOptions,proto3" json:"format_options,omitempty"`
	// delivery options for the report
	//
	// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
	DeliveryOptions *commons.DeliveryOptions `protobuf:"bytes,11,opt,name=delivery_options,json=deliveryOptions,proto3" json:"delivery_options,omitempty"`
	// whether the report job is active
	IsActive bool `protobuf:"varint,12,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// whether to send an empty report if no data is available
	SendEmptyReport bool `protobuf:"varint,13,opt,name=send_empty_report,json=sendEmptyReport,proto3" json:"send_empty_report,omitempty"`
	// dashboard resource id to report
	DashboardResourceId string `protobuf:"bytes,14,opt,name=dashboard_resource_id,json=dashboardResourceId,proto3" json:"dashboard_resource_id,omitempty"`
	// time_zone_wrapper to use for the report schedule
	TimeZoneWrapper *commons.TimeZoneWrapper `protobuf:"bytes,15,opt,name=time_zone_wrapper,json=timeZoneWrapper,proto3" json:"time_zone_wrapper,omitempty"`
	// hide csv footer
	HideCsvFooter bool `protobuf:"varint,16,opt,name=hide_csv_footer,json=hideCsvFooter,proto3" json:"hide_csv_footer,omitempty"`
	// transfer_config_sid to use for the report schedule
	TransferConfigSid int64 `protobuf:"varint,17,opt,name=transfer_config_sid,json=transferConfigSid,proto3" json:"transfer_config_sid,omitempty"`
	// cron expression for the report schedule
	CronExpression *commons.CronExpression `protobuf:"bytes,18,opt,name=cron_expression,json=cronExpression,proto3" json:"cron_expression,omitempty"`
}

func (x *ReportJob) Reset() {
	*x = ReportJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1alpha1_bireportgenerator_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportJob) ProtoMessage() {}

func (x *ReportJob) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_bireportgenerator_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportJob.ProtoReflect.Descriptor instead.
func (*ReportJob) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_bireportgenerator_entities_proto_rawDescGZIP(), []int{0}
}

func (x *ReportJob) GetReportJobId() string {
	if x != nil {
		return x.ReportJobId
	}
	return ""
}

func (x *ReportJob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReportJob) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
func (x *ReportJob) GetDashboardId() string {
	if x != nil {
		return x.DashboardId
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
func (x *ReportJob) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *ReportJob) GetTimePeriod() commons.TimePeriod {
	if x != nil {
		return x.TimePeriod
	}
	return commons.TimePeriod(0)
}

// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
func (x *ReportJob) GetDeliveryTimes() *commons.DeliveryTimes {
	if x != nil {
		return x.DeliveryTimes
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
func (x *ReportJob) GetDayFilter() *commons.DayFilter {
	if x != nil {
		return x.DayFilter
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
func (x *ReportJob) GetMonths() []commons.Month {
	if x != nil {
		return x.Months
	}
	return nil
}

func (x *ReportJob) GetFormatOptions() *commons.FormatOptions {
	if x != nil {
		return x.FormatOptions
	}
	return nil
}

// Deprecated: Marked as deprecated in api/v1alpha1/bireportgenerator/entities.proto.
func (x *ReportJob) GetDeliveryOptions() *commons.DeliveryOptions {
	if x != nil {
		return x.DeliveryOptions
	}
	return nil
}

func (x *ReportJob) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *ReportJob) GetSendEmptyReport() bool {
	if x != nil {
		return x.SendEmptyReport
	}
	return false
}

func (x *ReportJob) GetDashboardResourceId() string {
	if x != nil {
		return x.DashboardResourceId
	}
	return ""
}

func (x *ReportJob) GetTimeZoneWrapper() *commons.TimeZoneWrapper {
	if x != nil {
		return x.TimeZoneWrapper
	}
	return nil
}

func (x *ReportJob) GetHideCsvFooter() bool {
	if x != nil {
		return x.HideCsvFooter
	}
	return false
}

func (x *ReportJob) GetTransferConfigSid() int64 {
	if x != nil {
		return x.TransferConfigSid
	}
	return 0
}

func (x *ReportJob) GetCronExpression() *commons.CronExpression {
	if x != nil {
		return x.CronExpression
	}
	return nil
}

var File_api_v1alpha1_bireportgenerator_entities_proto protoreflect.FileDescriptor

var file_api_v1alpha1_bireportgenerator_entities_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x62,
	0x69, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x62, 0x69,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x1a,
	0x23, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x69, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x07,
	0x0a, 0x09, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x22, 0x0a, 0x0d, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0c, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x0b, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x38, 0x0a,
	0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x45, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x64, 0x61, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x44, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09,
	0x64, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x06, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4b, 0x0a, 0x10,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a,
	0x6f, 0x6e, 0x65, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52,
	0x0f, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x0f, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x63, 0x73, 0x76, 0x5f, 0x66, 0x6f, 0x6f,
	0x74, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x68, 0x69, 0x64, 0x65, 0x43,
	0x73, 0x76, 0x46, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x69, 0x64, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x0f, 0x63, 0x72, 0x6f, 0x6e,
	0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x43, 0x72, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0e,
	0x63, 0x72, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x88,
	0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x62, 0x69, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x62, 0x69, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x42, 0xaa, 0x02, 0x1e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x69, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0xca, 0x02, 0x1e, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x42, 0x69, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0xe2, 0x02, 0x2a, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x42, 0x69, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3a, 0x3a, 0x42, 0x69, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1alpha1_bireportgenerator_entities_proto_rawDescOnce sync.Once
	file_api_v1alpha1_bireportgenerator_entities_proto_rawDescData = file_api_v1alpha1_bireportgenerator_entities_proto_rawDesc
)

func file_api_v1alpha1_bireportgenerator_entities_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_bireportgenerator_entities_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_bireportgenerator_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1alpha1_bireportgenerator_entities_proto_rawDescData)
	})
	return file_api_v1alpha1_bireportgenerator_entities_proto_rawDescData
}

var file_api_v1alpha1_bireportgenerator_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1alpha1_bireportgenerator_entities_proto_goTypes = []interface{}{
	(*ReportJob)(nil),               // 0: api.v1alpha1.bireportgenerator.ReportJob
	(commons.TimePeriod)(0),         // 1: api.commons.TimePeriod
	(*commons.DeliveryTimes)(nil),   // 2: api.commons.DeliveryTimes
	(*commons.DayFilter)(nil),       // 3: api.commons.DayFilter
	(commons.Month)(0),              // 4: api.commons.Month
	(*commons.FormatOptions)(nil),   // 5: api.commons.FormatOptions
	(*commons.DeliveryOptions)(nil), // 6: api.commons.DeliveryOptions
	(*commons.TimeZoneWrapper)(nil), // 7: api.commons.TimeZoneWrapper
	(*commons.CronExpression)(nil),  // 8: api.commons.CronExpression
}
var file_api_v1alpha1_bireportgenerator_entities_proto_depIdxs = []int32{
	1, // 0: api.v1alpha1.bireportgenerator.ReportJob.time_period:type_name -> api.commons.TimePeriod
	2, // 1: api.v1alpha1.bireportgenerator.ReportJob.delivery_times:type_name -> api.commons.DeliveryTimes
	3, // 2: api.v1alpha1.bireportgenerator.ReportJob.day_filter:type_name -> api.commons.DayFilter
	4, // 3: api.v1alpha1.bireportgenerator.ReportJob.months:type_name -> api.commons.Month
	5, // 4: api.v1alpha1.bireportgenerator.ReportJob.format_options:type_name -> api.commons.FormatOptions
	6, // 5: api.v1alpha1.bireportgenerator.ReportJob.delivery_options:type_name -> api.commons.DeliveryOptions
	7, // 6: api.v1alpha1.bireportgenerator.ReportJob.time_zone_wrapper:type_name -> api.commons.TimeZoneWrapper
	8, // 7: api.v1alpha1.bireportgenerator.ReportJob.cron_expression:type_name -> api.commons.CronExpression
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_bireportgenerator_entities_proto_init() }
func file_api_v1alpha1_bireportgenerator_entities_proto_init() {
	if File_api_v1alpha1_bireportgenerator_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1alpha1_bireportgenerator_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportJob); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1alpha1_bireportgenerator_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_bireportgenerator_entities_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_bireportgenerator_entities_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_bireportgenerator_entities_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_bireportgenerator_entities_proto = out.File
	file_api_v1alpha1_bireportgenerator_entities_proto_rawDesc = nil
	file_api_v1alpha1_bireportgenerator_entities_proto_goTypes = nil
	file_api_v1alpha1_bireportgenerator_entities_proto_depIdxs = nil
}
