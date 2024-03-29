// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/pgc/gateway/player/v3/playurl.proto

package v3

import (
	playershared "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/playershared"
	any1 "github.com/golang/protobuf/ptypes/any"
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

// 播放页信息-请求
type PlayViewReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 视频信息
	Vod *playershared.VideoVod `protobuf:"bytes,1,opt,name=vod,proto3" json:"vod,omitempty"`
	// 当前页spm
	Spmid string `protobuf:"bytes,2,opt,name=spmid,proto3" json:"spmid,omitempty"`
	// 上一页spm
	FromSpmid string `protobuf:"bytes,3,opt,name=from_spmid,json=fromSpmid,proto3" json:"from_spmid,omitempty"`
	// 青少年模式
	TeenagersMode int32 `protobuf:"varint,4,opt,name=teenagers_mode,json=teenagersMode,proto3" json:"teenagers_mode,omitempty"`
	ExtraContent map[string]string `protobuf:"bytes,5,rep,name=extra_content,json=extraContent,proto3" json:"extra_content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PlayViewReq) Reset() {
	*x = PlayViewReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_pgc_gateway_player_v3_playurl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayViewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayViewReq) ProtoMessage() {}

func (x *PlayViewReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_pgc_gateway_player_v3_playurl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayViewReq.ProtoReflect.Descriptor instead.
func (*PlayViewReq) Descriptor() ([]byte, []int) {
	return file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDescGZIP(), []int{0}
}

func (x *PlayViewReq) GetVod() *playershared.VideoVod {
	if x != nil {
		return x.Vod
	}
	return nil
}

func (x *PlayViewReq) GetSpmid() string {
	if x != nil {
		return x.Spmid
	}
	return ""
}

func (x *PlayViewReq) GetFromSpmid() string {
	if x != nil {
		return x.FromSpmid
	}
	return ""
}

func (x *PlayViewReq) GetTeenagersMode() int32 {
	if x != nil {
		return x.TeenagersMode
	}
	return 0
}

func (x *PlayViewReq) GetExtraContent() map[string]string {
	if x != nil {
		return x.ExtraContent
	}
	return nil
}

// 播放页信息-响应
type PlayViewReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VodInfo     *playershared.VodInfo     `protobuf:"bytes,1,opt,name=vod_info,json=vodInfo,proto3" json:"vod_info,omitempty"`
	PlayArcConf *playershared.PlayArcConf `protobuf:"bytes,2,opt,name=play_arc_conf,json=playArcConf,proto3" json:"play_arc_conf,omitempty"`
	Supplement  *any1.Any                 `protobuf:"bytes,3,opt,name=supplement,proto3" json:"supplement,omitempty"`
	PlayArc     *playershared.PlayArc     `protobuf:"bytes,4,opt,name=play_arc,json=playArc,proto3" json:"play_arc,omitempty"`
	QnTrialInfo *playershared.QnTrialInfo `protobuf:"bytes,5,opt,name=qn_trial_info,json=qnTrialInfo,proto3" json:"qn_trial_info,omitempty"`
	Event       *playershared.Event       `protobuf:"bytes,6,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *PlayViewReply) Reset() {
	*x = PlayViewReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_pgc_gateway_player_v3_playurl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayViewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayViewReply) ProtoMessage() {}

func (x *PlayViewReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_pgc_gateway_player_v3_playurl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayViewReply.ProtoReflect.Descriptor instead.
func (*PlayViewReply) Descriptor() ([]byte, []int) {
	return file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDescGZIP(), []int{1}
}

func (x *PlayViewReply) GetVodInfo() *playershared.VodInfo {
	if x != nil {
		return x.VodInfo
	}
	return nil
}

func (x *PlayViewReply) GetPlayArcConf() *playershared.PlayArcConf {
	if x != nil {
		return x.PlayArcConf
	}
	return nil
}

func (x *PlayViewReply) GetSupplement() *any1.Any {
	if x != nil {
		return x.Supplement
	}
	return nil
}

func (x *PlayViewReply) GetPlayArc() *playershared.PlayArc {
	if x != nil {
		return x.PlayArc
	}
	return nil
}

func (x *PlayViewReply) GetQnTrialInfo() *playershared.QnTrialInfo {
	if x != nil {
		return x.QnTrialInfo
	}
	return nil
}

func (x *PlayViewReply) GetEvent() *playershared.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_bilibili_pgc_gateway_player_v3_playurl_proto protoreflect.FileDescriptor

var file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x70, 0x67, 0x63, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x33,
	0x2f, 0x70, 0x6c, 0x61, 0x79, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x67, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x28,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x56, 0x69, 0x65, 0x77,
	0x52, 0x65, 0x71, 0x12, 0x31, 0x0a, 0x03, 0x76, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x56, 0x6f,
	0x64, 0x52, 0x03, 0x76, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x6d, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x6d, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x70, 0x6d, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x70, 0x6d, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x65, 0x65, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x65, 0x65, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x62, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x67, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x56,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x3f, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x72, 0x61, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xff, 0x02, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79,
	0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x08, 0x76, 0x6f, 0x64,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x56, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x76, 0x6f, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x46, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x61, 0x72, 0x63,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x72, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x52,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x72, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x34, 0x0a, 0x0a,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x39, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x61, 0x72, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x41, 0x72, 0x63, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x72, 0x63, 0x12, 0x46, 0x0a,
	0x0d, 0x71, 0x6e, 0x5f, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x51, 0x6e, 0x54,
	0x72, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x71, 0x6e, 0x54, 0x72, 0x69, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75,
	0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2f, 0x70, 0x67, 0x63, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDescOnce sync.Once
	file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDescData = file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDesc
)

func file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDescGZIP() []byte {
	file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDescOnce.Do(func() {
		file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDescData)
	})
	return file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDescData
}

var file_bilibili_pgc_gateway_player_v3_playurl_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bilibili_pgc_gateway_player_v3_playurl_proto_goTypes = []interface{}{
	(*PlayViewReq)(nil),              // 0: bilibili.pgc.gateway.player.v3.PlayViewReq
	(*PlayViewReply)(nil),            // 1: bilibili.pgc.gateway.player.v3.PlayViewReply
	nil,                              // 2: bilibili.pgc.gateway.player.v3.PlayViewReq.ExtraContentEntry
	(*playershared.VideoVod)(nil),    // 3: bilibili.playershared.VideoVod
	(*playershared.VodInfo)(nil),     // 4: bilibili.playershared.VodInfo
	(*playershared.PlayArcConf)(nil), // 5: bilibili.playershared.PlayArcConf
	(*any1.Any)(nil),                 // 6: google.protobuf.Any
	(*playershared.PlayArc)(nil),     // 7: bilibili.playershared.PlayArc
	(*playershared.QnTrialInfo)(nil), // 8: bilibili.playershared.QnTrialInfo
	(*playershared.Event)(nil),       // 9: bilibili.playershared.Event
}
var file_bilibili_pgc_gateway_player_v3_playurl_proto_depIdxs = []int32{
	3, // 0: bilibili.pgc.gateway.player.v3.PlayViewReq.vod:type_name -> bilibili.playershared.VideoVod
	2, // 1: bilibili.pgc.gateway.player.v3.PlayViewReq.extra_content:type_name -> bilibili.pgc.gateway.player.v3.PlayViewReq.ExtraContentEntry
	4, // 2: bilibili.pgc.gateway.player.v3.PlayViewReply.vod_info:type_name -> bilibili.playershared.VodInfo
	5, // 3: bilibili.pgc.gateway.player.v3.PlayViewReply.play_arc_conf:type_name -> bilibili.playershared.PlayArcConf
	6, // 4: bilibili.pgc.gateway.player.v3.PlayViewReply.supplement:type_name -> google.protobuf.Any
	7, // 5: bilibili.pgc.gateway.player.v3.PlayViewReply.play_arc:type_name -> bilibili.playershared.PlayArc
	8, // 6: bilibili.pgc.gateway.player.v3.PlayViewReply.qn_trial_info:type_name -> bilibili.playershared.QnTrialInfo
	9, // 7: bilibili.pgc.gateway.player.v3.PlayViewReply.event:type_name -> bilibili.playershared.Event
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_bilibili_pgc_gateway_player_v3_playurl_proto_init() }
func file_bilibili_pgc_gateway_player_v3_playurl_proto_init() {
	if File_bilibili_pgc_gateway_player_v3_playurl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_pgc_gateway_player_v3_playurl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayViewReq); i {
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
		file_bilibili_pgc_gateway_player_v3_playurl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayViewReply); i {
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
			RawDescriptor: file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_pgc_gateway_player_v3_playurl_proto_goTypes,
		DependencyIndexes: file_bilibili_pgc_gateway_player_v3_playurl_proto_depIdxs,
		MessageInfos:      file_bilibili_pgc_gateway_player_v3_playurl_proto_msgTypes,
	}.Build()
	File_bilibili_pgc_gateway_player_v3_playurl_proto = out.File
	file_bilibili_pgc_gateway_player_v3_playurl_proto_rawDesc = nil
	file_bilibili_pgc_gateway_player_v3_playurl_proto_goTypes = nil
	file_bilibili_pgc_gateway_player_v3_playurl_proto_depIdxs = nil
}
