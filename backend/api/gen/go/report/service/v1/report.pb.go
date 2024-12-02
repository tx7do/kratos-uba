// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: report/service/v1/report.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 报告
type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 报告ID
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_v1_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_v1_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_report_service_v1_report_proto_rawDescGZIP(), []int{0}
}

func (x *Report) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 获取报告列表 - 答复
type ListReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Report `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int32     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListReportResponse) Reset() {
	*x = ListReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_v1_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReportResponse) ProtoMessage() {}

func (x *ListReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_v1_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReportResponse.ProtoReflect.Descriptor instead.
func (*ListReportResponse) Descriptor() ([]byte, []int) {
	return file_report_service_v1_report_proto_rawDescGZIP(), []int{1}
}

func (x *ListReportResponse) GetItems() []*Report {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListReportResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 获取报告数据 - 请求
type GetReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetReportRequest) Reset() {
	*x = GetReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_v1_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportRequest) ProtoMessage() {}

func (x *GetReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_v1_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportRequest.ProtoReflect.Descriptor instead.
func (*GetReportRequest) Descriptor() ([]byte, []int) {
	return file_report_service_v1_report_proto_rawDescGZIP(), []int{2}
}

func (x *GetReportRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 创建报告 - 请求
type CreateReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *Report `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	OperatorId uint32  `protobuf:"varint,2,opt,name=operatorId,proto3" json:"operatorId,omitempty"`
}

func (x *CreateReportRequest) Reset() {
	*x = CreateReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_v1_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReportRequest) ProtoMessage() {}

func (x *CreateReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_v1_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReportRequest.ProtoReflect.Descriptor instead.
func (*CreateReportRequest) Descriptor() ([]byte, []int) {
	return file_report_service_v1_report_proto_rawDescGZIP(), []int{3}
}

func (x *CreateReportRequest) GetUser() *Report {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreateReportRequest) GetOperatorId() uint32 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

// 更新报告 - 请求
type UpdateReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User       *Report `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	OperatorId uint32  `protobuf:"varint,3,opt,name=operatorId,proto3" json:"operatorId,omitempty"`
}

func (x *UpdateReportRequest) Reset() {
	*x = UpdateReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_v1_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReportRequest) ProtoMessage() {}

func (x *UpdateReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_v1_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReportRequest.ProtoReflect.Descriptor instead.
func (*UpdateReportRequest) Descriptor() ([]byte, []int) {
	return file_report_service_v1_report_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateReportRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateReportRequest) GetUser() *Report {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UpdateReportRequest) GetOperatorId() uint32 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

// 删除报告 - 请求
type DeleteReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OperatorId uint32 `protobuf:"varint,2,opt,name=operatorId,proto3" json:"operatorId,omitempty"`
}

func (x *DeleteReportRequest) Reset() {
	*x = DeleteReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_v1_report_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReportRequest) ProtoMessage() {}

func (x *DeleteReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_v1_report_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReportRequest.ProtoReflect.Descriptor instead.
func (*DeleteReportRequest) Descriptor() ([]byte, []int) {
	return file_report_service_v1_report_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteReportRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteReportRequest) GetOperatorId() uint32 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

var File_report_service_v1_report_proto protoreflect.FileDescriptor

var file_report_service_v1_report_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0e, 0xba, 0x47,
	0x0b, 0x92, 0x02, 0x08, 0xe6, 0x8a, 0xa5, 0xe5, 0x91, 0x8a, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x5b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x22, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x64, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x45, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x32, 0xac, 0x03, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x00, 0x12, 0x50, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0xbd, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2d, 0x75, 0x62, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x52, 0x53, 0x58, 0xaa, 0x02, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_service_v1_report_proto_rawDescOnce sync.Once
	file_report_service_v1_report_proto_rawDescData = file_report_service_v1_report_proto_rawDesc
)

func file_report_service_v1_report_proto_rawDescGZIP() []byte {
	file_report_service_v1_report_proto_rawDescOnce.Do(func() {
		file_report_service_v1_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_service_v1_report_proto_rawDescData)
	})
	return file_report_service_v1_report_proto_rawDescData
}

var file_report_service_v1_report_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_report_service_v1_report_proto_goTypes = []interface{}{
	(*Report)(nil),              // 0: report.service.v1.Report
	(*ListReportResponse)(nil),  // 1: report.service.v1.ListReportResponse
	(*GetReportRequest)(nil),    // 2: report.service.v1.GetReportRequest
	(*CreateReportRequest)(nil), // 3: report.service.v1.CreateReportRequest
	(*UpdateReportRequest)(nil), // 4: report.service.v1.UpdateReportRequest
	(*DeleteReportRequest)(nil), // 5: report.service.v1.DeleteReportRequest
	(*v1.PagingRequest)(nil),    // 6: pagination.PagingRequest
	(*emptypb.Empty)(nil),       // 7: google.protobuf.Empty
}
var file_report_service_v1_report_proto_depIdxs = []int32{
	0, // 0: report.service.v1.ListReportResponse.items:type_name -> report.service.v1.Report
	0, // 1: report.service.v1.CreateReportRequest.user:type_name -> report.service.v1.Report
	0, // 2: report.service.v1.UpdateReportRequest.user:type_name -> report.service.v1.Report
	6, // 3: report.service.v1.ReportService.ListReport:input_type -> pagination.PagingRequest
	2, // 4: report.service.v1.ReportService.GetReport:input_type -> report.service.v1.GetReportRequest
	3, // 5: report.service.v1.ReportService.CreateReport:input_type -> report.service.v1.CreateReportRequest
	4, // 6: report.service.v1.ReportService.UpdateReport:input_type -> report.service.v1.UpdateReportRequest
	5, // 7: report.service.v1.ReportService.DeleteReport:input_type -> report.service.v1.DeleteReportRequest
	1, // 8: report.service.v1.ReportService.ListReport:output_type -> report.service.v1.ListReportResponse
	0, // 9: report.service.v1.ReportService.GetReport:output_type -> report.service.v1.Report
	0, // 10: report.service.v1.ReportService.CreateReport:output_type -> report.service.v1.Report
	0, // 11: report.service.v1.ReportService.UpdateReport:output_type -> report.service.v1.Report
	7, // 12: report.service.v1.ReportService.DeleteReport:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_report_service_v1_report_proto_init() }
func file_report_service_v1_report_proto_init() {
	if File_report_service_v1_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_report_service_v1_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_report_service_v1_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReportResponse); i {
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
		file_report_service_v1_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportRequest); i {
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
		file_report_service_v1_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReportRequest); i {
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
		file_report_service_v1_report_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReportRequest); i {
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
		file_report_service_v1_report_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReportRequest); i {
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
			RawDescriptor: file_report_service_v1_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_report_service_v1_report_proto_goTypes,
		DependencyIndexes: file_report_service_v1_report_proto_depIdxs,
		MessageInfos:      file_report_service_v1_report_proto_msgTypes,
	}.Build()
	File_report_service_v1_report_proto = out.File
	file_report_service_v1_report_proto_rawDesc = nil
	file_report_service_v1_report_proto_goTypes = nil
	file_report_service_v1_report_proto_depIdxs = nil
}