// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/pgc/service/premiere/v1/premiere.proto

package v1

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

// 获取首播状态-请求
type PremiereStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 剧集epid
	EpId int64 `protobuf:"varint,1,opt,name=ep_id,json=epId,proto3" json:"ep_id,omitempty"`
}

func (x *PremiereStatusReq) Reset() {
	*x = PremiereStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_pgc_service_premiere_v1_premiere_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PremiereStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PremiereStatusReq) ProtoMessage() {}

func (x *PremiereStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_pgc_service_premiere_v1_premiere_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PremiereStatusReq.ProtoReflect.Descriptor instead.
func (*PremiereStatusReq) Descriptor() ([]byte, []int) {
	return file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDescGZIP(), []int{0}
}

func (x *PremiereStatusReq) GetEpId() int64 {
	if x != nil {
		return x.EpId
	}
	return 0
}

// 获取首播状态-响应
type PremiereStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务端播放进度 单位ms 用户实际播放进度：progress - delay_time
	Progress int64 `protobuf:"varint,1,opt,name=progress,proto3" json:"progress,omitempty"`
	// 起播时间戳 单位ms
	StartTime int64 `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// 延迟播放时长 单位ms
	DelayTime int64 `protobuf:"varint,3,opt,name=delay_time,json=delayTime,proto3" json:"delay_time,omitempty"`
	// 首播在线人数
	OnlineCount int64 `protobuf:"varint,4,opt,name=online_count,json=onlineCount,proto3" json:"online_count,omitempty"`
	// 首播状态
	// 1:预热 2:首播中 3:紧急停播 4:已结束
	Status int32 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	// 首播结束后跳转类型
	// 1:下架 2:转点播
	AfterPremiereType int32 `protobuf:"varint,6,opt,name=after_premiere_type,json=afterPremiereType,proto3" json:"after_premiere_type,omitempty"`
}

func (x *PremiereStatusReply) Reset() {
	*x = PremiereStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_pgc_service_premiere_v1_premiere_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PremiereStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PremiereStatusReply) ProtoMessage() {}

func (x *PremiereStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_pgc_service_premiere_v1_premiere_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PremiereStatusReply.ProtoReflect.Descriptor instead.
func (*PremiereStatusReply) Descriptor() ([]byte, []int) {
	return file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDescGZIP(), []int{1}
}

func (x *PremiereStatusReply) GetProgress() int64 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *PremiereStatusReply) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *PremiereStatusReply) GetDelayTime() int64 {
	if x != nil {
		return x.DelayTime
	}
	return 0
}

func (x *PremiereStatusReply) GetOnlineCount() int64 {
	if x != nil {
		return x.OnlineCount
	}
	return 0
}

func (x *PremiereStatusReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PremiereStatusReply) GetAfterPremiereType() int32 {
	if x != nil {
		return x.AfterPremiereType
	}
	return 0
}

var File_bilibili_pgc_service_premiere_v1_premiere_proto protoreflect.FileDescriptor

var file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x70, 0x67, 0x63, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x67, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0x28, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x05, 0x65, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x65, 0x70, 0x49, 0x64, 0x22, 0xda, 0x01,
	0x0a, 0x13, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x61, 0x66, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x32, 0x80, 0x01, 0x0a, 0x08, 0x50,
	0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x12, 0x74, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x33, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x67, 0x63,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x35, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x70, 0x67, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x65,
	0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x4d, 0x5a,
	0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f,
	0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2f, 0x70, 0x67, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDescOnce sync.Once
	file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDescData = file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDesc
)

func file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDescGZIP() []byte {
	file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDescOnce.Do(func() {
		file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDescData)
	})
	return file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDescData
}

var file_bilibili_pgc_service_premiere_v1_premiere_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bilibili_pgc_service_premiere_v1_premiere_proto_goTypes = []interface{}{
	(*PremiereStatusReq)(nil),   // 0: bilibili.pgc.service.premiere.v1.PremiereStatusReq
	(*PremiereStatusReply)(nil), // 1: bilibili.pgc.service.premiere.v1.PremiereStatusReply
}
var file_bilibili_pgc_service_premiere_v1_premiere_proto_depIdxs = []int32{
	0, // 0: bilibili.pgc.service.premiere.v1.Premiere.Status:input_type -> bilibili.pgc.service.premiere.v1.PremiereStatusReq
	1, // 1: bilibili.pgc.service.premiere.v1.Premiere.Status:output_type -> bilibili.pgc.service.premiere.v1.PremiereStatusReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bilibili_pgc_service_premiere_v1_premiere_proto_init() }
func file_bilibili_pgc_service_premiere_v1_premiere_proto_init() {
	if File_bilibili_pgc_service_premiere_v1_premiere_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_pgc_service_premiere_v1_premiere_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PremiereStatusReq); i {
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
		file_bilibili_pgc_service_premiere_v1_premiere_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PremiereStatusReply); i {
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
			RawDescriptor: file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_pgc_service_premiere_v1_premiere_proto_goTypes,
		DependencyIndexes: file_bilibili_pgc_service_premiere_v1_premiere_proto_depIdxs,
		MessageInfos:      file_bilibili_pgc_service_premiere_v1_premiere_proto_msgTypes,
	}.Build()
	File_bilibili_pgc_service_premiere_v1_premiere_proto = out.File
	file_bilibili_pgc_service_premiere_v1_premiere_proto_rawDesc = nil
	file_bilibili_pgc_service_premiere_v1_premiere_proto_goTypes = nil
	file_bilibili_pgc_service_premiere_v1_premiere_proto_depIdxs = nil
}
