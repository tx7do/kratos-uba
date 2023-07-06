// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: report/service/v1/kafka.proto

package v1

import (
	_ "github.com/google/gnostic/openapiv3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "kratos-bi/gen/api/go/common/pagination"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReportType int32

const (
	ReportType_ReportTypeUnknown ReportType = 0
	ReportType_ReportTypeUser    ReportType = 1
	ReportType_ReportTypeEvent   ReportType = 2
)

// Enum value maps for ReportType.
var (
	ReportType_name = map[int32]string{
		0: "ReportTypeUnknown",
		1: "ReportTypeUser",
		2: "ReportTypeEvent",
	}
	ReportType_value = map[string]int32{
		"ReportTypeUnknown": 0,
		"ReportTypeUser":    1,
		"ReportTypeEvent":   2,
	}
)

func (x ReportType) Enum() *ReportType {
	p := new(ReportType)
	*p = x
	return p
}

func (x ReportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportType) Descriptor() protoreflect.EnumDescriptor {
	return file_report_service_v1_kafka_proto_enumTypes[0].Descriptor()
}

func (ReportType) Type() protoreflect.EnumType {
	return &file_report_service_v1_kafka_proto_enumTypes[0]
}

func (x ReportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportType.Descriptor instead.
func (ReportType) EnumDescriptor() ([]byte, []int) {
	return file_report_service_v1_kafka_proto_rawDescGZIP(), []int{0}
}

type UserReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId           string `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	DistinctId      string `protobuf:"bytes,2,opt,name=distinctId,proto3" json:"distinctId,omitempty"`
	TableId         string `protobuf:"bytes,3,opt,name=tableId,proto3" json:"tableId,omitempty"`
	Ip              string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	ReportType      int32  `protobuf:"varint,5,opt,name=reportType,proto3" json:"reportType,omitempty"`
	Debug           string `protobuf:"bytes,6,opt,name=debug,proto3" json:"debug,omitempty"`
	RequestData     []byte `protobuf:"bytes,7,opt,name=requestData,proto3" json:"requestData,omitempty"`
	ReportTime      string `protobuf:"bytes,8,opt,name=reportTime,proto3" json:"reportTime,omitempty"`
	ConsumptionTime string `protobuf:"bytes,9,opt,name=consumptionTime,proto3" json:"consumptionTime,omitempty"`
	EventName       string `protobuf:"bytes,10,opt,name=eventName,proto3" json:"eventName,omitempty"`
	Offset          int64  `protobuf:"varint,11,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *UserReport) Reset() {
	*x = UserReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_v1_kafka_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReport) ProtoMessage() {}

func (x *UserReport) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_v1_kafka_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReport.ProtoReflect.Descriptor instead.
func (*UserReport) Descriptor() ([]byte, []int) {
	return file_report_service_v1_kafka_proto_rawDescGZIP(), []int{0}
}

func (x *UserReport) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *UserReport) GetDistinctId() string {
	if x != nil {
		return x.DistinctId
	}
	return ""
}

func (x *UserReport) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

func (x *UserReport) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *UserReport) GetReportType() int32 {
	if x != nil {
		return x.ReportType
	}
	return 0
}

func (x *UserReport) GetDebug() string {
	if x != nil {
		return x.Debug
	}
	return ""
}

func (x *UserReport) GetRequestData() []byte {
	if x != nil {
		return x.RequestData
	}
	return nil
}

func (x *UserReport) GetReportTime() string {
	if x != nil {
		return x.ReportTime
	}
	return ""
}

func (x *UserReport) GetConsumptionTime() string {
	if x != nil {
		return x.ConsumptionTime
	}
	return ""
}

func (x *UserReport) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *UserReport) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type EventReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId    string            `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	DistinctId   string            `protobuf:"bytes,2,opt,name=distinctId,proto3" json:"distinctId,omitempty"`
	UserId       string            `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	PartEvent    string            `protobuf:"bytes,4,opt,name=partEvent,proto3" json:"partEvent,omitempty"`
	PartDate     string            `protobuf:"bytes,5,opt,name=partDate,proto3" json:"partDate,omitempty"`
	MpPlatform   string            `protobuf:"bytes,6,opt,name=mpPlatform,proto3" json:"mpPlatform,omitempty"`
	LibVersion   string            `protobuf:"bytes,7,opt,name=libVersion,proto3" json:"libVersion,omitempty"`
	Os           string            `protobuf:"bytes,8,opt,name=os,proto3" json:"os,omitempty"`
	OsVersion    string            `protobuf:"bytes,20,opt,name=osVersion,proto3" json:"osVersion,omitempty"`
	ScreenWidth  string            `protobuf:"bytes,9,opt,name=screenWidth,proto3" json:"screenWidth,omitempty"`    // 屏幕宽
	ScreenHeight string            `protobuf:"bytes,11,opt,name=screenHeight,proto3" json:"screenHeight,omitempty"` // 屏幕高
	DeviceId     string            `protobuf:"bytes,12,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	NetworkType  string            `protobuf:"bytes,13,opt,name=networkType,proto3" json:"networkType,omitempty"`
	DeviceModel  string            `protobuf:"bytes,14,opt,name=deviceModel,proto3" json:"deviceModel,omitempty"`
	CountryCode  string            `protobuf:"bytes,10,opt,name=countryCode,proto3" json:"countryCode,omitempty"` // 国家
	Province     string            `protobuf:"bytes,16,opt,name=province,proto3" json:"province,omitempty"`       // 省
	City         string            `protobuf:"bytes,15,opt,name=city,proto3" json:"city,omitempty"`               // 城市
	Lib          string            `protobuf:"bytes,17,opt,name=lib,proto3" json:"lib,omitempty"`
	Scene        string            `protobuf:"bytes,18,opt,name=scene,proto3" json:"scene,omitempty"`
	Manufacturer string            `protobuf:"bytes,19,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Ip           string            `protobuf:"bytes,22,opt,name=ip,proto3" json:"ip,omitempty"`
	TableId      string            `protobuf:"bytes,23,opt,name=tableId,proto3" json:"tableId,omitempty"`
	Properties   map[string]string `protobuf:"bytes,24,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EventReport) Reset() {
	*x = EventReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_v1_kafka_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventReport) ProtoMessage() {}

func (x *EventReport) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_v1_kafka_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventReport.ProtoReflect.Descriptor instead.
func (*EventReport) Descriptor() ([]byte, []int) {
	return file_report_service_v1_kafka_proto_rawDescGZIP(), []int{1}
}

func (x *EventReport) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *EventReport) GetDistinctId() string {
	if x != nil {
		return x.DistinctId
	}
	return ""
}

func (x *EventReport) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EventReport) GetPartEvent() string {
	if x != nil {
		return x.PartEvent
	}
	return ""
}

func (x *EventReport) GetPartDate() string {
	if x != nil {
		return x.PartDate
	}
	return ""
}

func (x *EventReport) GetMpPlatform() string {
	if x != nil {
		return x.MpPlatform
	}
	return ""
}

func (x *EventReport) GetLibVersion() string {
	if x != nil {
		return x.LibVersion
	}
	return ""
}

func (x *EventReport) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *EventReport) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *EventReport) GetScreenWidth() string {
	if x != nil {
		return x.ScreenWidth
	}
	return ""
}

func (x *EventReport) GetScreenHeight() string {
	if x != nil {
		return x.ScreenHeight
	}
	return ""
}

func (x *EventReport) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *EventReport) GetNetworkType() string {
	if x != nil {
		return x.NetworkType
	}
	return ""
}

func (x *EventReport) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *EventReport) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *EventReport) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *EventReport) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *EventReport) GetLib() string {
	if x != nil {
		return x.Lib
	}
	return ""
}

func (x *EventReport) GetScene() string {
	if x != nil {
		return x.Scene
	}
	return ""
}

func (x *EventReport) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *EventReport) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *EventReport) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

func (x *EventReport) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

type AcceptStatusReportData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataName      *string `protobuf:"bytes,1,opt,name=dataName,proto3,oneof" json:"dataName,omitempty"`
	ReportType    *string `protobuf:"bytes,2,opt,name=reportType,proto3,oneof" json:"reportType,omitempty"`
	ReportData    *string `protobuf:"bytes,3,opt,name=reportData,proto3,oneof" json:"reportData,omitempty"`
	ErrorReason   *string `protobuf:"bytes,4,opt,name=errorReason,proto3,oneof" json:"errorReason,omitempty"`
	ErrorHandling *string `protobuf:"bytes,5,opt,name=errorHandling,proto3,oneof" json:"errorHandling,omitempty"`
	Status        *int32  `protobuf:"varint,6,opt,name=status,proto3,oneof" json:"status,omitempty"`
	PartEvent     *string `protobuf:"bytes,7,opt,name=partEvent,proto3,oneof" json:"partEvent,omitempty"`
	PartDate      *string `protobuf:"bytes,8,opt,name=partDate,proto3,oneof" json:"partDate,omitempty"`
}

func (x *AcceptStatusReportData) Reset() {
	*x = AcceptStatusReportData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_v1_kafka_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptStatusReportData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptStatusReportData) ProtoMessage() {}

func (x *AcceptStatusReportData) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_v1_kafka_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptStatusReportData.ProtoReflect.Descriptor instead.
func (*AcceptStatusReportData) Descriptor() ([]byte, []int) {
	return file_report_service_v1_kafka_proto_rawDescGZIP(), []int{2}
}

func (x *AcceptStatusReportData) GetDataName() string {
	if x != nil && x.DataName != nil {
		return *x.DataName
	}
	return ""
}

func (x *AcceptStatusReportData) GetReportType() string {
	if x != nil && x.ReportType != nil {
		return *x.ReportType
	}
	return ""
}

func (x *AcceptStatusReportData) GetReportData() string {
	if x != nil && x.ReportData != nil {
		return *x.ReportData
	}
	return ""
}

func (x *AcceptStatusReportData) GetErrorReason() string {
	if x != nil && x.ErrorReason != nil {
		return *x.ErrorReason
	}
	return ""
}

func (x *AcceptStatusReportData) GetErrorHandling() string {
	if x != nil && x.ErrorHandling != nil {
		return *x.ErrorHandling
	}
	return ""
}

func (x *AcceptStatusReportData) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *AcceptStatusReportData) GetPartEvent() string {
	if x != nil && x.PartEvent != nil {
		return *x.PartEvent
	}
	return ""
}

func (x *AcceptStatusReportData) GetPartDate() string {
	if x != nil && x.PartDate != nil {
		return *x.PartDate
	}
	return ""
}

type RealTimeWarehousingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventName  *string `protobuf:"bytes,1,opt,name=eventName,json=event_name,proto3,oneof" json:"eventName,omitempty"`
	ReportData *string `protobuf:"bytes,2,opt,name=reportData,json=report_data,proto3,oneof" json:"reportData,omitempty"`
	CreateTime *string `protobuf:"bytes,3,opt,name=createTime,json=create_time,proto3,oneof" json:"createTime,omitempty"`
}

func (x *RealTimeWarehousingData) Reset() {
	*x = RealTimeWarehousingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_v1_kafka_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RealTimeWarehousingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RealTimeWarehousingData) ProtoMessage() {}

func (x *RealTimeWarehousingData) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_v1_kafka_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RealTimeWarehousingData.ProtoReflect.Descriptor instead.
func (*RealTimeWarehousingData) Descriptor() ([]byte, []int) {
	return file_report_service_v1_kafka_proto_rawDescGZIP(), []int{3}
}

func (x *RealTimeWarehousingData) GetEventName() string {
	if x != nil && x.EventName != nil {
		return *x.EventName
	}
	return ""
}

func (x *RealTimeWarehousingData) GetReportData() string {
	if x != nil && x.ReportData != nil {
		return *x.ReportData
	}
	return ""
}

func (x *RealTimeWarehousingData) GetCreateTime() string {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return ""
}

type KafkaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId           *int64  `protobuf:"varint,1,opt,name=appId,proto3,oneof" json:"appId,omitempty"`
	DistinctId      *int64  `protobuf:"varint,2,opt,name=distinctId,proto3,oneof" json:"distinctId,omitempty"`
	TableId         *int64  `protobuf:"varint,3,opt,name=tableId,proto3,oneof" json:"tableId,omitempty"`
	Ip              *string `protobuf:"bytes,4,opt,name=ip,proto3,oneof" json:"ip,omitempty"`
	ReportType      *int32  `protobuf:"varint,5,opt,name=reportType,proto3,oneof" json:"reportType,omitempty"`
	Debug           *string `protobuf:"bytes,6,opt,name=debug,proto3,oneof" json:"debug,omitempty"`
	ReqData         *string `protobuf:"bytes,7,opt,name=reqData,proto3,oneof" json:"reqData,omitempty"`
	ReportTime      *string `protobuf:"bytes,8,opt,name=reportTime,proto3,oneof" json:"reportTime,omitempty"`
	ConsumptionTime *string `protobuf:"bytes,9,opt,name=consumptionTime,proto3,oneof" json:"consumptionTime,omitempty"`
	EventName       *string `protobuf:"bytes,10,opt,name=eventName,proto3,oneof" json:"eventName,omitempty"`
}

func (x *KafkaData) Reset() {
	*x = KafkaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_v1_kafka_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaData) ProtoMessage() {}

func (x *KafkaData) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_v1_kafka_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaData.ProtoReflect.Descriptor instead.
func (*KafkaData) Descriptor() ([]byte, []int) {
	return file_report_service_v1_kafka_proto_rawDescGZIP(), []int{4}
}

func (x *KafkaData) GetAppId() int64 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *KafkaData) GetDistinctId() int64 {
	if x != nil && x.DistinctId != nil {
		return *x.DistinctId
	}
	return 0
}

func (x *KafkaData) GetTableId() int64 {
	if x != nil && x.TableId != nil {
		return *x.TableId
	}
	return 0
}

func (x *KafkaData) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

func (x *KafkaData) GetReportType() int32 {
	if x != nil && x.ReportType != nil {
		return *x.ReportType
	}
	return 0
}

func (x *KafkaData) GetDebug() string {
	if x != nil && x.Debug != nil {
		return *x.Debug
	}
	return ""
}

func (x *KafkaData) GetReqData() string {
	if x != nil && x.ReqData != nil {
		return *x.ReqData
	}
	return ""
}

func (x *KafkaData) GetReportTime() string {
	if x != nil && x.ReportTime != nil {
		return *x.ReportTime
	}
	return ""
}

func (x *KafkaData) GetConsumptionTime() string {
	if x != nil && x.ConsumptionTime != nil {
		return *x.ConsumptionTime
	}
	return ""
}

func (x *KafkaData) GetEventName() string {
	if x != nil && x.EventName != nil {
		return *x.EventName
	}
	return ""
}

var File_report_service_v1_kafka_proto protoreflect.FileDescriptor

var file_report_service_v1_kafka_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x02, 0x0a, 0x0a, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x20,
	0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x88, 0x06, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x70, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x70, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x69, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x57, 0x69, 0x64, 0x74, 0x68,
	0x12, 0x22, 0x0a, 0x0c, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x69, 0x62, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x69, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa9, 0x03, 0x0a, 0x16,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01,
	0x01, 0x12, 0x25, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x04, 0x52, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x21, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x08, 0x70, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x61, 0x6d,
	0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e,
	0x67, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x61, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xd7, 0x03, 0x0a, 0x09, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0a,
	0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02,
	0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x02, 0x69, 0x70, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x06, 0x52, 0x07, 0x72, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x49, 0x64,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x69, 0x70, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x72, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x4c, 0x0a, 0x0a, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x42, 0x2b, 0x5a, 0x29, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2d, 0x62, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_service_v1_kafka_proto_rawDescOnce sync.Once
	file_report_service_v1_kafka_proto_rawDescData = file_report_service_v1_kafka_proto_rawDesc
)

func file_report_service_v1_kafka_proto_rawDescGZIP() []byte {
	file_report_service_v1_kafka_proto_rawDescOnce.Do(func() {
		file_report_service_v1_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_service_v1_kafka_proto_rawDescData)
	})
	return file_report_service_v1_kafka_proto_rawDescData
}

var file_report_service_v1_kafka_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_report_service_v1_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_report_service_v1_kafka_proto_goTypes = []interface{}{
	(ReportType)(0),                 // 0: report.service.v1.ReportType
	(*UserReport)(nil),              // 1: report.service.v1.UserReport
	(*EventReport)(nil),             // 2: report.service.v1.EventReport
	(*AcceptStatusReportData)(nil),  // 3: report.service.v1.AcceptStatusReportData
	(*RealTimeWarehousingData)(nil), // 4: report.service.v1.RealTimeWarehousingData
	(*KafkaData)(nil),               // 5: report.service.v1.KafkaData
	nil,                             // 6: report.service.v1.EventReport.PropertiesEntry
}
var file_report_service_v1_kafka_proto_depIdxs = []int32{
	6, // 0: report.service.v1.EventReport.properties:type_name -> report.service.v1.EventReport.PropertiesEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_report_service_v1_kafka_proto_init() }
func file_report_service_v1_kafka_proto_init() {
	if File_report_service_v1_kafka_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_report_service_v1_kafka_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReport); i {
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
		file_report_service_v1_kafka_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventReport); i {
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
		file_report_service_v1_kafka_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptStatusReportData); i {
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
		file_report_service_v1_kafka_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RealTimeWarehousingData); i {
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
		file_report_service_v1_kafka_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaData); i {
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
	file_report_service_v1_kafka_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_report_service_v1_kafka_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_report_service_v1_kafka_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_report_service_v1_kafka_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_report_service_v1_kafka_proto_goTypes,
		DependencyIndexes: file_report_service_v1_kafka_proto_depIdxs,
		EnumInfos:         file_report_service_v1_kafka_proto_enumTypes,
		MessageInfos:      file_report_service_v1_kafka_proto_msgTypes,
	}.Build()
	File_report_service_v1_kafka_proto = out.File
	file_report_service_v1_kafka_proto_rawDesc = nil
	file_report_service_v1_kafka_proto_goTypes = nil
	file_report_service_v1_kafka_proto_depIdxs = nil
}
