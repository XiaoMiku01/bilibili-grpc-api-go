// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/app/playeronline/v1/playeronline.proto

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

// 空回复
type NoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoReply) Reset() {
	*x = NoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoReply) ProtoMessage() {}

func (x *NoReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoReply.ProtoReflect.Descriptor instead.
func (*NoReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_playeronline_v1_playeronline_proto_rawDescGZIP(), []int{0}
}

// 获取在线人数-回复
type PlayerOnlineReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalText string `protobuf:"bytes,1,opt,name=total_text,json=totalText,proto3" json:"total_text,omitempty"`
	// 下次轮询间隔时间
	SecNext int64 `protobuf:"varint,2,opt,name=sec_next,json=secNext,proto3" json:"sec_next,omitempty"`
	// 是否底部显示
	BottomShow bool `protobuf:"varint,3,opt,name=bottom_show,json=bottomShow,proto3" json:"bottom_show,omitempty"`
	SdmShow bool `protobuf:"varint,4,opt,name=sdm_show,json=sdmShow,proto3" json:"sdm_show,omitempty"`
	SdmText string `protobuf:"bytes,5,opt,name=sdm_text,json=sdmText,proto3" json:"sdm_text,omitempty"`
	TotalNumber int64 `protobuf:"varint,6,opt,name=total_number,json=totalNumber,proto3" json:"total_number,omitempty"`
	TotalNumberText string `protobuf:"bytes,7,opt,name=total_number_text,json=totalNumberText,proto3" json:"total_number_text,omitempty"`
}

func (x *PlayerOnlineReply) Reset() {
	*x = PlayerOnlineReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerOnlineReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerOnlineReply) ProtoMessage() {}

func (x *PlayerOnlineReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerOnlineReply.ProtoReflect.Descriptor instead.
func (*PlayerOnlineReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_playeronline_v1_playeronline_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerOnlineReply) GetTotalText() string {
	if x != nil {
		return x.TotalText
	}
	return ""
}

func (x *PlayerOnlineReply) GetSecNext() int64 {
	if x != nil {
		return x.SecNext
	}
	return 0
}

func (x *PlayerOnlineReply) GetBottomShow() bool {
	if x != nil {
		return x.BottomShow
	}
	return false
}

func (x *PlayerOnlineReply) GetSdmShow() bool {
	if x != nil {
		return x.SdmShow
	}
	return false
}

func (x *PlayerOnlineReply) GetSdmText() string {
	if x != nil {
		return x.SdmText
	}
	return ""
}

func (x *PlayerOnlineReply) GetTotalNumber() int64 {
	if x != nil {
		return x.TotalNumber
	}
	return 0
}

func (x *PlayerOnlineReply) GetTotalNumberText() string {
	if x != nil {
		return x.TotalNumberText
	}
	return ""
}

// 获取在线人数-请求
type PlayerOnlineReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 稿件 avid
	Aid int64 `protobuf:"varint,1,opt,name=aid,proto3" json:"aid,omitempty"`
	// 视频 cid
	Cid int64 `protobuf:"varint,2,opt,name=cid,proto3" json:"cid,omitempty"`
	// 是否在播放中
	PlayOpen bool `protobuf:"varint,3,opt,name=play_open,json=playOpen,proto3" json:"play_open,omitempty"`
}

func (x *PlayerOnlineReq) Reset() {
	*x = PlayerOnlineReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerOnlineReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerOnlineReq) ProtoMessage() {}

func (x *PlayerOnlineReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerOnlineReq.ProtoReflect.Descriptor instead.
func (*PlayerOnlineReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_playeronline_v1_playeronline_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerOnlineReq) GetAid() int64 {
	if x != nil {
		return x.Aid
	}
	return 0
}

func (x *PlayerOnlineReq) GetCid() int64 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *PlayerOnlineReq) GetPlayOpen() bool {
	if x != nil {
		return x.PlayOpen
	}
	return false
}

type PremiereInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PremiereOverText string `protobuf:"bytes,1,opt,name=premiere_over_text,json=premiereOverText,proto3" json:"premiere_over_text,omitempty"`
	Participant int64 `protobuf:"varint,2,opt,name=participant,proto3" json:"participant,omitempty"`
	Interaction int64 `protobuf:"varint,3,opt,name=interaction,proto3" json:"interaction,omitempty"`
}

func (x *PremiereInfoReply) Reset() {
	*x = PremiereInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PremiereInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PremiereInfoReply) ProtoMessage() {}

func (x *PremiereInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PremiereInfoReply.ProtoReflect.Descriptor instead.
func (*PremiereInfoReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_playeronline_v1_playeronline_proto_rawDescGZIP(), []int{3}
}

func (x *PremiereInfoReply) GetPremiereOverText() string {
	if x != nil {
		return x.PremiereOverText
	}
	return ""
}

func (x *PremiereInfoReply) GetParticipant() int64 {
	if x != nil {
		return x.Participant
	}
	return 0
}

func (x *PremiereInfoReply) GetInteraction() int64 {
	if x != nil {
		return x.Interaction
	}
	return 0
}

type PremiereInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aid int64 `protobuf:"varint,1,opt,name=aid,proto3" json:"aid,omitempty"`
}

func (x *PremiereInfoReq) Reset() {
	*x = PremiereInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PremiereInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PremiereInfoReq) ProtoMessage() {}

func (x *PremiereInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PremiereInfoReq.ProtoReflect.Descriptor instead.
func (*PremiereInfoReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_playeronline_v1_playeronline_proto_rawDescGZIP(), []int{4}
}

func (x *PremiereInfoReq) GetAid() int64 {
	if x != nil {
		return x.Aid
	}
	return 0
}

type ReportWatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aid int64 `protobuf:"varint,1,opt,name=aid,proto3" json:"aid,omitempty"`
	Biz string `protobuf:"bytes,2,opt,name=biz,proto3" json:"biz,omitempty"`
	Buvid string `protobuf:"bytes,3,opt,name=buvid,proto3" json:"buvid,omitempty"`
}

func (x *ReportWatchReq) Reset() {
	*x = ReportWatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportWatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportWatchReq) ProtoMessage() {}

func (x *ReportWatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportWatchReq.ProtoReflect.Descriptor instead.
func (*ReportWatchReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_playeronline_v1_playeronline_proto_rawDescGZIP(), []int{5}
}

func (x *ReportWatchReq) GetAid() int64 {
	if x != nil {
		return x.Aid
	}
	return 0
}

func (x *ReportWatchReq) GetBiz() string {
	if x != nil {
		return x.Biz
	}
	return ""
}

func (x *ReportWatchReq) GetBuvid() string {
	if x != nil {
		return x.Buvid
	}
	return ""
}

var File_bilibili_app_playeronline_v1_playeronline_proto protoreflect.FileDescriptor

var file_bilibili_app_playeronline_v1_playeronline_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x22,
	0x09, 0x0a, 0x07, 0x4e, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xf3, 0x01, 0x0a, 0x11, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f,
	0x74, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x64, 0x6d, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x64, 0x6d, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x64, 0x6d, 0x5f, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x64, 0x6d, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x65, 0x78, 0x74,
	0x22, 0x52, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x61, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6f, 0x70, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x4f, 0x70, 0x65, 0x6e, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x72,
	0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65,
	0x4f, 0x76, 0x65, 0x72, 0x54, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x23, 0x0a, 0x0f,
	0x50, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x69,
	0x64, 0x22, 0x4a, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x61, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x76, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x76, 0x69, 0x64, 0x32, 0xd2, 0x02,
	0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x6e,
	0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2d,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6e,
	0x0a, 0x0c, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x65, 0x6d, 0x69, 0x65, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65,
	0x6d, 0x69, 0x65, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x62,
	0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2c, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_playeronline_v1_playeronline_proto_rawDescOnce sync.Once
	file_bilibili_app_playeronline_v1_playeronline_proto_rawDescData = file_bilibili_app_playeronline_v1_playeronline_proto_rawDesc
)

func file_bilibili_app_playeronline_v1_playeronline_proto_rawDescGZIP() []byte {
	file_bilibili_app_playeronline_v1_playeronline_proto_rawDescOnce.Do(func() {
		file_bilibili_app_playeronline_v1_playeronline_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_playeronline_v1_playeronline_proto_rawDescData)
	})
	return file_bilibili_app_playeronline_v1_playeronline_proto_rawDescData
}

var file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bilibili_app_playeronline_v1_playeronline_proto_goTypes = []interface{}{
	(*NoReply)(nil),           // 0: bilibili.app.playeronline.v1.NoReply
	(*PlayerOnlineReply)(nil), // 1: bilibili.app.playeronline.v1.PlayerOnlineReply
	(*PlayerOnlineReq)(nil),   // 2: bilibili.app.playeronline.v1.PlayerOnlineReq
	(*PremiereInfoReply)(nil), // 3: bilibili.app.playeronline.v1.PremiereInfoReply
	(*PremiereInfoReq)(nil),   // 4: bilibili.app.playeronline.v1.PremiereInfoReq
	(*ReportWatchReq)(nil),    // 5: bilibili.app.playeronline.v1.ReportWatchReq
}
var file_bilibili_app_playeronline_v1_playeronline_proto_depIdxs = []int32{
	2, // 0: bilibili.app.playeronline.v1.PlayerOnline.PlayerOnline:input_type -> bilibili.app.playeronline.v1.PlayerOnlineReq
	4, // 1: bilibili.app.playeronline.v1.PlayerOnline.PremiereInfo:input_type -> bilibili.app.playeronline.v1.PremiereInfoReq
	5, // 2: bilibili.app.playeronline.v1.PlayerOnline.ReportWatch:input_type -> bilibili.app.playeronline.v1.ReportWatchReq
	1, // 3: bilibili.app.playeronline.v1.PlayerOnline.PlayerOnline:output_type -> bilibili.app.playeronline.v1.PlayerOnlineReply
	3, // 4: bilibili.app.playeronline.v1.PlayerOnline.PremiereInfo:output_type -> bilibili.app.playeronline.v1.PremiereInfoReply
	0, // 5: bilibili.app.playeronline.v1.PlayerOnline.ReportWatch:output_type -> bilibili.app.playeronline.v1.NoReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bilibili_app_playeronline_v1_playeronline_proto_init() }
func file_bilibili_app_playeronline_v1_playeronline_proto_init() {
	if File_bilibili_app_playeronline_v1_playeronline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoReply); i {
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
		file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerOnlineReply); i {
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
		file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerOnlineReq); i {
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
		file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PremiereInfoReply); i {
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
		file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PremiereInfoReq); i {
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
		file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportWatchReq); i {
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
			RawDescriptor: file_bilibili_app_playeronline_v1_playeronline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_app_playeronline_v1_playeronline_proto_goTypes,
		DependencyIndexes: file_bilibili_app_playeronline_v1_playeronline_proto_depIdxs,
		MessageInfos:      file_bilibili_app_playeronline_v1_playeronline_proto_msgTypes,
	}.Build()
	File_bilibili_app_playeronline_v1_playeronline_proto = out.File
	file_bilibili_app_playeronline_v1_playeronline_proto_rawDesc = nil
	file_bilibili_app_playeronline_v1_playeronline_proto_goTypes = nil
	file_bilibili_app_playeronline_v1_playeronline_proto_depIdxs = nil
}
