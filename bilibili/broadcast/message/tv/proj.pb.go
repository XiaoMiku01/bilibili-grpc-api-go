// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/broadcast/message/tv/proj.proto

package tv

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

// 投屏
type ProjReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 投屏命令
	// 1:起播 2:快进 3:快退 4:seek播放进度 5:暂停 6:暂停恢复
	CmdType int64 `protobuf:"varint,1,opt,name=cmd_type,json=cmdType,proto3" json:"cmd_type,omitempty"`
	// 用户id
	Mid int64 `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	// 稿件id
	Aid int64 `protobuf:"varint,3,opt,name=aid,proto3" json:"aid,omitempty"`
	// 视频id
	Cid int64 `protobuf:"varint,4,opt,name=cid,proto3" json:"cid,omitempty"`
	// 视频类型
	// 0:ugc 1:pgc 2:pugv
	VideoType int64 `protobuf:"varint,5,opt,name=video_type,json=videoType,proto3" json:"video_type,omitempty"`
	// 单集id，pgc和pugv需要传
	EpId int64 `protobuf:"varint,6,opt,name=ep_id,json=epId,proto3" json:"ep_id,omitempty"`
	// 剧集id
	SeasonId int64 `protobuf:"varint,7,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	// seek 的位置，cmd位seek时有值，单位秒
	SeekTs int64 `protobuf:"varint,8,opt,name=seek_ts,json=seekTs,proto3" json:"seek_ts,omitempty"`
	// 其他指令对应内容
	Extra string `protobuf:"bytes,9,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *ProjReply) Reset() {
	*x = ProjReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_tv_proj_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjReply) ProtoMessage() {}

func (x *ProjReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_tv_proj_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjReply.ProtoReflect.Descriptor instead.
func (*ProjReply) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_tv_proj_proto_rawDescGZIP(), []int{0}
}

func (x *ProjReply) GetCmdType() int64 {
	if x != nil {
		return x.CmdType
	}
	return 0
}

func (x *ProjReply) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *ProjReply) GetAid() int64 {
	if x != nil {
		return x.Aid
	}
	return 0
}

func (x *ProjReply) GetCid() int64 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *ProjReply) GetVideoType() int64 {
	if x != nil {
		return x.VideoType
	}
	return 0
}

func (x *ProjReply) GetEpId() int64 {
	if x != nil {
		return x.EpId
	}
	return 0
}

func (x *ProjReply) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *ProjReply) GetSeekTs() int64 {
	if x != nil {
		return x.SeekTs
	}
	return 0
}

func (x *ProjReply) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

// 直播状态
type LiveStatusNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 直播状态
	// 1:开播 2:关播 3:截流 4:截流恢复
	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// 文案
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 直播房间号
	Cid int64 `protobuf:"varint,3,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *LiveStatusNotify) Reset() {
	*x = LiveStatusNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_tv_proj_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveStatusNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveStatusNotify) ProtoMessage() {}

func (x *LiveStatusNotify) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_tv_proj_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveStatusNotify.ProtoReflect.Descriptor instead.
func (*LiveStatusNotify) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_tv_proj_proto_rawDescGZIP(), []int{1}
}

func (x *LiveStatusNotify) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LiveStatusNotify) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LiveStatusNotify) GetCid() int64 {
	if x != nil {
		return x.Cid
	}
	return 0
}

type EsportsNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 直播房间号
	Cid int64 `protobuf:"varint,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *EsportsNotify) Reset() {
	*x = EsportsNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_tv_proj_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EsportsNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EsportsNotify) ProtoMessage() {}

func (x *EsportsNotify) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_tv_proj_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EsportsNotify.ProtoReflect.Descriptor instead.
func (*EsportsNotify) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_tv_proj_proto_rawDescGZIP(), []int{2}
}

func (x *EsportsNotify) GetCid() int64 {
	if x != nil {
		return x.Cid
	}
	return 0
}

// 直播插卡
type PublicityNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 插卡id
	PublicityId int64 `protobuf:"varint,1,opt,name=publicity_id,json=publicityId,proto3" json:"publicity_id,omitempty"`
	// 直播房间号
	RoomId int64 `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	// 直播间状态
	// 0:未开播 1:直播中 2:轮播中
	Status int64 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PublicityNotify) Reset() {
	*x = PublicityNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_tv_proj_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicityNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicityNotify) ProtoMessage() {}

func (x *PublicityNotify) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_tv_proj_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicityNotify.ProtoReflect.Descriptor instead.
func (*PublicityNotify) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_tv_proj_proto_rawDescGZIP(), []int{3}
}

func (x *PublicityNotify) GetPublicityId() int64 {
	if x != nil {
		return x.PublicityId
	}
	return 0
}

func (x *PublicityNotify) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *PublicityNotify) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

// 直转点
type LiveSkipNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 直播id
	LiveId int64 `protobuf:"varint,1,opt,name=live_id,json=liveId,proto3" json:"live_id,omitempty"`
}

func (x *LiveSkipNotify) Reset() {
	*x = LiveSkipNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_tv_proj_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveSkipNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveSkipNotify) ProtoMessage() {}

func (x *LiveSkipNotify) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_tv_proj_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveSkipNotify.ProtoReflect.Descriptor instead.
func (*LiveSkipNotify) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_tv_proj_proto_rawDescGZIP(), []int{4}
}

func (x *LiveSkipNotify) GetLiveId() int64 {
	if x != nil {
		return x.LiveId
	}
	return 0
}

var File_bilibili_broadcast_message_tv_proj_proto protoreflect.FileDescriptor

var file_bilibili_broadcast_message_tv_proj_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x76, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x76, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6d, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x61, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x65, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x65, 0x6b, 0x5f, 0x74,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x65, 0x6b, 0x54, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x4e, 0x0a, 0x10, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x63, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0d, 0x45, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x29, 0x0a, 0x0e, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x64, 0x32, 0xa8, 0x03, 0x0a, 0x02, 0x54,
	0x76, 0x12, 0x4a, 0x0a, 0x04, 0x50, 0x72, 0x6f, 0x6a, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x28, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x74,
	0x76, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x12, 0x57, 0x0a,
	0x0a, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x2f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x74, 0x76, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x30, 0x01, 0x12, 0x51, 0x0a, 0x07, 0x45, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2c, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x76, 0x2e, 0x45, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x09, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2e,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x76, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x30, 0x01,
	0x12, 0x53, 0x0a, 0x08, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2d, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x74, 0x76, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x30, 0x01, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x74,
	0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_broadcast_message_tv_proj_proto_rawDescOnce sync.Once
	file_bilibili_broadcast_message_tv_proj_proto_rawDescData = file_bilibili_broadcast_message_tv_proj_proto_rawDesc
)

func file_bilibili_broadcast_message_tv_proj_proto_rawDescGZIP() []byte {
	file_bilibili_broadcast_message_tv_proj_proto_rawDescOnce.Do(func() {
		file_bilibili_broadcast_message_tv_proj_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_broadcast_message_tv_proj_proto_rawDescData)
	})
	return file_bilibili_broadcast_message_tv_proj_proto_rawDescData
}

var file_bilibili_broadcast_message_tv_proj_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bilibili_broadcast_message_tv_proj_proto_goTypes = []interface{}{
	(*ProjReply)(nil),        // 0: bilibili.broadcast.message.tv.ProjReply
	(*LiveStatusNotify)(nil), // 1: bilibili.broadcast.message.tv.LiveStatusNotify
	(*EsportsNotify)(nil),    // 2: bilibili.broadcast.message.tv.EsportsNotify
	(*PublicityNotify)(nil),  // 3: bilibili.broadcast.message.tv.PublicityNotify
	(*LiveSkipNotify)(nil),   // 4: bilibili.broadcast.message.tv.LiveSkipNotify
	(*empty.Empty)(nil),      // 5: google.protobuf.Empty
}
var file_bilibili_broadcast_message_tv_proj_proto_depIdxs = []int32{
	5, // 0: bilibili.broadcast.message.tv.Tv.Proj:input_type -> google.protobuf.Empty
	5, // 1: bilibili.broadcast.message.tv.Tv.LiveStatus:input_type -> google.protobuf.Empty
	5, // 2: bilibili.broadcast.message.tv.Tv.Esports:input_type -> google.protobuf.Empty
	5, // 3: bilibili.broadcast.message.tv.Tv.Publicity:input_type -> google.protobuf.Empty
	5, // 4: bilibili.broadcast.message.tv.Tv.LiveSkip:input_type -> google.protobuf.Empty
	0, // 5: bilibili.broadcast.message.tv.Tv.Proj:output_type -> bilibili.broadcast.message.tv.ProjReply
	1, // 6: bilibili.broadcast.message.tv.Tv.LiveStatus:output_type -> bilibili.broadcast.message.tv.LiveStatusNotify
	2, // 7: bilibili.broadcast.message.tv.Tv.Esports:output_type -> bilibili.broadcast.message.tv.EsportsNotify
	3, // 8: bilibili.broadcast.message.tv.Tv.Publicity:output_type -> bilibili.broadcast.message.tv.PublicityNotify
	4, // 9: bilibili.broadcast.message.tv.Tv.LiveSkip:output_type -> bilibili.broadcast.message.tv.LiveSkipNotify
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bilibili_broadcast_message_tv_proj_proto_init() }
func file_bilibili_broadcast_message_tv_proj_proto_init() {
	if File_bilibili_broadcast_message_tv_proj_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_broadcast_message_tv_proj_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjReply); i {
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
		file_bilibili_broadcast_message_tv_proj_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveStatusNotify); i {
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
		file_bilibili_broadcast_message_tv_proj_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EsportsNotify); i {
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
		file_bilibili_broadcast_message_tv_proj_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicityNotify); i {
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
		file_bilibili_broadcast_message_tv_proj_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveSkipNotify); i {
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
			RawDescriptor: file_bilibili_broadcast_message_tv_proj_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_broadcast_message_tv_proj_proto_goTypes,
		DependencyIndexes: file_bilibili_broadcast_message_tv_proj_proto_depIdxs,
		MessageInfos:      file_bilibili_broadcast_message_tv_proj_proto_msgTypes,
	}.Build()
	File_bilibili_broadcast_message_tv_proj_proto = out.File
	file_bilibili_broadcast_message_tv_proj_proto_rawDesc = nil
	file_bilibili_broadcast_message_tv_proj_proto_goTypes = nil
	file_bilibili_broadcast_message_tv_proj_proto_depIdxs = nil
}
