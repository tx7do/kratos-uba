// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: agent/service/v1/i_report.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PostReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportType string `protobuf:"bytes,1,opt,name=reportType,proto3" json:"reportType,omitempty"` // 类型
	AppId      string `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`           // 应用ID
	AppKey     string `protobuf:"bytes,3,opt,name=appKey,proto3" json:"appKey,omitempty"`         // 应用Key
	EventName  string `protobuf:"bytes,4,opt,name=eventName,proto3" json:"eventName,omitempty"`   // 事件名
	Debug      int32  `protobuf:"varint,5,opt,name=debug,proto3" json:"debug,omitempty"`          // 调试
	Content    string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`       // 事件内容
}

func (x *PostReportRequest) Reset() {
	*x = PostReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_service_v1_i_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostReportRequest) ProtoMessage() {}

func (x *PostReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_service_v1_i_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostReportRequest.ProtoReflect.Descriptor instead.
func (*PostReportRequest) Descriptor() ([]byte, []int) {
	return file_agent_service_v1_i_report_proto_rawDescGZIP(), []int{0}
}

func (x *PostReportRequest) GetReportType() string {
	if x != nil {
		return x.ReportType
	}
	return ""
}

func (x *PostReportRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *PostReportRequest) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

func (x *PostReportRequest) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *PostReportRequest) GetDebug() int32 {
	if x != nil {
		return x.Debug
	}
	return 0
}

func (x *PostReportRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PostReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PostReportResponse) Reset() {
	*x = PostReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_service_v1_i_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostReportResponse) ProtoMessage() {}

func (x *PostReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_service_v1_i_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostReportResponse.ProtoReflect.Descriptor instead.
func (*PostReportResponse) Descriptor() ([]byte, []int) {
	return file_agent_service_v1_i_report_proto_rawDescGZIP(), []int{1}
}

func (x *PostReportResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PostReportResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_agent_service_v1_i_report_proto protoreflect.FileDescriptor

var file_agent_service_v1_i_report_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x32, 0x85, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x23, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0xe4, 0x01, 0xba, 0x47, 0x29,
	0x12, 0x27, 0x0a, 0x12, 0xe5, 0xba, 0x94, 0xe7, 0x94, 0xa8, 0xe8, 0xae, 0xa4, 0xe8, 0xaf, 0x81,
	0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0x12, 0x0c, 0xe5, 0xba, 0x94, 0xe7, 0x94, 0xa8, 0xe8, 0xae,
	0xa4, 0xe8, 0xaf, 0x81, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x49, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x30, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x75, 0x62, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x41, 0x53, 0x58, 0xaa, 0x02, 0x10, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_service_v1_i_report_proto_rawDescOnce sync.Once
	file_agent_service_v1_i_report_proto_rawDescData = file_agent_service_v1_i_report_proto_rawDesc
)

func file_agent_service_v1_i_report_proto_rawDescGZIP() []byte {
	file_agent_service_v1_i_report_proto_rawDescOnce.Do(func() {
		file_agent_service_v1_i_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_service_v1_i_report_proto_rawDescData)
	})
	return file_agent_service_v1_i_report_proto_rawDescData
}

var file_agent_service_v1_i_report_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_agent_service_v1_i_report_proto_goTypes = []interface{}{
	(*PostReportRequest)(nil),  // 0: agent.service.v1.PostReportRequest
	(*PostReportResponse)(nil), // 1: agent.service.v1.PostReportResponse
}
var file_agent_service_v1_i_report_proto_depIdxs = []int32{
	0, // 0: agent.service.v1.ReportService.PostReport:input_type -> agent.service.v1.PostReportRequest
	1, // 1: agent.service.v1.ReportService.PostReport:output_type -> agent.service.v1.PostReportResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_agent_service_v1_i_report_proto_init() }
func file_agent_service_v1_i_report_proto_init() {
	if File_agent_service_v1_i_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_service_v1_i_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostReportRequest); i {
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
		file_agent_service_v1_i_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostReportResponse); i {
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
			RawDescriptor: file_agent_service_v1_i_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_service_v1_i_report_proto_goTypes,
		DependencyIndexes: file_agent_service_v1_i_report_proto_depIdxs,
		MessageInfos:      file_agent_service_v1_i_report_proto_msgTypes,
	}.Build()
	File_agent_service_v1_i_report_proto = out.File
	file_agent_service_v1_i_report_proto_rawDesc = nil
	file_agent_service_v1_i_report_proto_goTypes = nil
	file_agent_service_v1_i_report_proto_depIdxs = nil
}
