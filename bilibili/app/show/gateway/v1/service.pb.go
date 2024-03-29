// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/app/show/gateway/v1/service.proto

package v1

import (
	main "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/broadcast/message/main"
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

// 获取Native页进度数据-请求
type GetActProgressReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Native页id
	PageID int64 `protobuf:"varint,1,opt,name=pageID,proto3" json:"pageID,omitempty"`
	// 用户mid
	Mid int64 `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
}

func (x *GetActProgressReq) Reset() {
	*x = GetActProgressReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_show_gateway_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActProgressReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActProgressReq) ProtoMessage() {}

func (x *GetActProgressReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_show_gateway_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActProgressReq.ProtoReflect.Descriptor instead.
func (*GetActProgressReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_show_gateway_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetActProgressReq) GetPageID() int64 {
	if x != nil {
		return x.PageID
	}
	return 0
}

func (x *GetActProgressReq) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

// 获取Native页进度数据-响应
type GetActProgressReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 进度数据
	Event *main.NativePageEvent `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *GetActProgressReply) Reset() {
	*x = GetActProgressReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_show_gateway_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActProgressReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActProgressReply) ProtoMessage() {}

func (x *GetActProgressReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_show_gateway_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActProgressReply.ProtoReflect.Descriptor instead.
func (*GetActProgressReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_show_gateway_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetActProgressReply) GetEvent() *main.NativePageEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_bilibili_app_show_gateway_v1_service_proto protoreflect.FileDescriptor

var file_bilibili_app_show_gateway_v1_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x68, 0x6f, 0x77, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x2c, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0x7f, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x53, 0x68, 0x6f,
	0x77, 0x12, 0x74, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x2f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x31, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31,
	0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_show_gateway_v1_service_proto_rawDescOnce sync.Once
	file_bilibili_app_show_gateway_v1_service_proto_rawDescData = file_bilibili_app_show_gateway_v1_service_proto_rawDesc
)

func file_bilibili_app_show_gateway_v1_service_proto_rawDescGZIP() []byte {
	file_bilibili_app_show_gateway_v1_service_proto_rawDescOnce.Do(func() {
		file_bilibili_app_show_gateway_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_show_gateway_v1_service_proto_rawDescData)
	})
	return file_bilibili_app_show_gateway_v1_service_proto_rawDescData
}

var file_bilibili_app_show_gateway_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bilibili_app_show_gateway_v1_service_proto_goTypes = []interface{}{
	(*GetActProgressReq)(nil),    // 0: bilibili.app.show.gateway.v1.GetActProgressReq
	(*GetActProgressReply)(nil),  // 1: bilibili.app.show.gateway.v1.GetActProgressReply
	(*main.NativePageEvent)(nil), // 2: bilibili.broadcast.message.main.NativePageEvent
}
var file_bilibili_app_show_gateway_v1_service_proto_depIdxs = []int32{
	2, // 0: bilibili.app.show.gateway.v1.GetActProgressReply.event:type_name -> bilibili.broadcast.message.main.NativePageEvent
	0, // 1: bilibili.app.show.gateway.v1.AppShow.GetActProgress:input_type -> bilibili.app.show.gateway.v1.GetActProgressReq
	1, // 2: bilibili.app.show.gateway.v1.AppShow.GetActProgress:output_type -> bilibili.app.show.gateway.v1.GetActProgressReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bilibili_app_show_gateway_v1_service_proto_init() }
func file_bilibili_app_show_gateway_v1_service_proto_init() {
	if File_bilibili_app_show_gateway_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_show_gateway_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActProgressReq); i {
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
		file_bilibili_app_show_gateway_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActProgressReply); i {
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
			RawDescriptor: file_bilibili_app_show_gateway_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_app_show_gateway_v1_service_proto_goTypes,
		DependencyIndexes: file_bilibili_app_show_gateway_v1_service_proto_depIdxs,
		MessageInfos:      file_bilibili_app_show_gateway_v1_service_proto_msgTypes,
	}.Build()
	File_bilibili_app_show_gateway_v1_service_proto = out.File
	file_bilibili_app_show_gateway_v1_service_proto_rawDesc = nil
	file_bilibili_app_show_gateway_v1_service_proto_goTypes = nil
	file_bilibili_app_show_gateway_v1_service_proto_depIdxs = nil
}
