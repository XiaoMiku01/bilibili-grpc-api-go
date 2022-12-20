// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/app/show/popular/v1/popular.proto

package v1

import (
	v11 "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/app/archive/middleware/v1"
	v1 "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/app/card/v1"
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

// 气泡信息
type Bubble struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文案
	BubbleContent string `protobuf:"bytes,1,opt,name=bubble_content,json=bubbleContent,proto3" json:"bubble_content,omitempty"`
	// 版本
	Version int32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// 起始时间
	Stime int64 `protobuf:"varint,3,opt,name=stime,proto3" json:"stime,omitempty"`
}

func (x *Bubble) Reset() {
	*x = Bubble{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_show_popular_v1_popular_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bubble) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bubble) ProtoMessage() {}

func (x *Bubble) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_show_popular_v1_popular_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bubble.ProtoReflect.Descriptor instead.
func (*Bubble) Descriptor() ([]byte, []int) {
	return file_bilibili_app_show_popular_v1_popular_proto_rawDescGZIP(), []int{0}
}

func (x *Bubble) GetBubbleContent() string {
	if x != nil {
		return x.BubbleContent
	}
	return ""
}

func (x *Bubble) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Bubble) GetStime() int64 {
	if x != nil {
		return x.Stime
	}
	return 0
}

// 配置信息
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 标题
	ItemTitle string `protobuf:"bytes,1,opt,name=item_title,json=itemTitle,proto3" json:"item_title,omitempty"`
	// 底部文案
	BottomText string `protobuf:"bytes,2,opt,name=bottom_text,json=bottomText,proto3" json:"bottom_text,omitempty"`
	// 底部图片url
	BottomTextCover string `protobuf:"bytes,3,opt,name=bottom_text_cover,json=bottomTextCover,proto3" json:"bottom_text_cover,omitempty"`
	// 底部跳转页url
	BottomTextUrl string `protobuf:"bytes,4,opt,name=bottom_text_url,json=bottomTextUrl,proto3" json:"bottom_text_url,omitempty"`
	// 顶部按钮信息列表
	TopItems []*EntranceShow `protobuf:"bytes,5,rep,name=top_items,json=topItems,proto3" json:"top_items,omitempty"`
	// 头图url
	HeadImage string `protobuf:"bytes,6,opt,name=head_image,json=headImage,proto3" json:"head_image,omitempty"`
	// 当前页按钮信息
	PageItems []*EntranceShow `protobuf:"bytes,7,rep,name=page_items,json=pageItems,proto3" json:"page_items,omitempty"`
	Hit int32 `protobuf:"varint,8,opt,name=hit,proto3" json:"hit,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_show_popular_v1_popular_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_show_popular_v1_popular_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_bilibili_app_show_popular_v1_popular_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetItemTitle() string {
	if x != nil {
		return x.ItemTitle
	}
	return ""
}

func (x *Config) GetBottomText() string {
	if x != nil {
		return x.BottomText
	}
	return ""
}

func (x *Config) GetBottomTextCover() string {
	if x != nil {
		return x.BottomTextCover
	}
	return ""
}

func (x *Config) GetBottomTextUrl() string {
	if x != nil {
		return x.BottomTextUrl
	}
	return ""
}

func (x *Config) GetTopItems() []*EntranceShow {
	if x != nil {
		return x.TopItems
	}
	return nil
}

func (x *Config) GetHeadImage() string {
	if x != nil {
		return x.HeadImage
	}
	return ""
}

func (x *Config) GetPageItems() []*EntranceShow {
	if x != nil {
		return x.PageItems
	}
	return nil
}

func (x *Config) GetHit() int32 {
	if x != nil {
		return x.Hit
	}
	return 0
}

// 按钮信息
type EntranceShow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 按钮图标url
	Icon string `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`
	// 按钮名
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// 入口模块id
	ModuleId string `protobuf:"bytes,3,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	// 跳转uri
	Uri string `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`
	// 气泡信息
	Bubble *Bubble `protobuf:"bytes,5,opt,name=bubble,proto3" json:"bubble,omitempty"`
	// 入口id
	EntranceId int64 `protobuf:"varint,6,opt,name=entrance_id,json=entranceId,proto3" json:"entrance_id,omitempty"`
	// 头图url
	TopPhoto string `protobuf:"bytes,7,opt,name=top_photo,json=topPhoto,proto3" json:"top_photo,omitempty"`
	// 入口类型
	EntranceType int32 `protobuf:"varint,8,opt,name=entrance_type,json=entranceType,proto3" json:"entrance_type,omitempty"`
}

func (x *EntranceShow) Reset() {
	*x = EntranceShow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_show_popular_v1_popular_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntranceShow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntranceShow) ProtoMessage() {}

func (x *EntranceShow) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_show_popular_v1_popular_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntranceShow.ProtoReflect.Descriptor instead.
func (*EntranceShow) Descriptor() ([]byte, []int) {
	return file_bilibili_app_show_popular_v1_popular_proto_rawDescGZIP(), []int{2}
}

func (x *EntranceShow) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *EntranceShow) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EntranceShow) GetModuleId() string {
	if x != nil {
		return x.ModuleId
	}
	return ""
}

func (x *EntranceShow) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *EntranceShow) GetBubble() *Bubble {
	if x != nil {
		return x.Bubble
	}
	return nil
}

func (x *EntranceShow) GetEntranceId() int64 {
	if x != nil {
		return x.EntranceId
	}
	return 0
}

func (x *EntranceShow) GetTopPhoto() string {
	if x != nil {
		return x.TopPhoto
	}
	return ""
}

func (x *EntranceShow) GetEntranceType() int32 {
	if x != nil {
		return x.EntranceType
	}
	return 0
}

// 热门列表-响应
type PopularReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 卡片列表
	Items []*v1.Card `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// 配置信息
	Config *Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// 版本
	Ver string `protobuf:"bytes,3,opt,name=ver,proto3" json:"ver,omitempty"`
}

func (x *PopularReply) Reset() {
	*x = PopularReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_show_popular_v1_popular_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopularReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopularReply) ProtoMessage() {}

func (x *PopularReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_show_popular_v1_popular_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopularReply.ProtoReflect.Descriptor instead.
func (*PopularReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_show_popular_v1_popular_proto_rawDescGZIP(), []int{3}
}

func (x *PopularReply) GetItems() []*v1.Card {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PopularReply) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *PopularReply) GetVer() string {
	if x != nil {
		return x.Ver
	}
	return ""
}

// 热门列表-请求
type PopularResultReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 排位索引id，为上此请求末尾项的idx
	Idx int64 `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	// 登录标识
	// 1:未登陆用户第一页 2:登陆用户第一页
	LoginEvent int32 `protobuf:"varint,2,opt,name=login_event,json=loginEvent,proto3" json:"login_event,omitempty"`
	// 清晰度(旧版)
	Qn int32 `protobuf:"varint,3,opt,name=qn,proto3" json:"qn,omitempty"`
	// 视频流版本(旧版)
	Fnver int32 `protobuf:"varint,4,opt,name=fnver,proto3" json:"fnver,omitempty"`
	// 视频流功能(旧版)
	Fnval int32 `protobuf:"varint,5,opt,name=fnval,proto3" json:"fnval,omitempty"`
	// 是否强制使用域名(旧版)
	ForceHost int32 `protobuf:"varint,6,opt,name=force_host,json=forceHost,proto3" json:"force_host,omitempty"`
	// 是否4K(旧版)
	Fourk int32 `protobuf:"varint,7,opt,name=fourk,proto3" json:"fourk,omitempty"`
	// 当前页面spm
	Spmid string `protobuf:"bytes,8,opt,name=spmid,proto3" json:"spmid,omitempty"`
	// 上此请求末尾项的param
	LastParam string `protobuf:"bytes,9,opt,name=last_param,json=lastParam,proto3" json:"last_param,omitempty"`
	// 上此请求的ver
	Ver string `protobuf:"bytes,10,opt,name=ver,proto3" json:"ver,omitempty"`
	// 分品类热门的入口ID
	EntranceId int64 `protobuf:"varint,11,opt,name=entrance_id,json=entranceId,proto3" json:"entrance_id,omitempty"`
	// 热门定位id集合
	LocationIds string `protobuf:"bytes,12,opt,name=location_ids,json=locationIds,proto3" json:"location_ids,omitempty"`
	// 0:tag页 1:中间页
	SourceId int32 `protobuf:"varint,13,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	// 数据埋点上报
	// 0:代表手动刷新 1:代表自动刷新
	Flush int32 `protobuf:"varint,14,opt,name=flush,proto3" json:"flush,omitempty"`
	// 秒开参数
	PlayerArgs *v11.PlayerArgs `protobuf:"bytes,15,opt,name=player_args,json=playerArgs,proto3" json:"player_args,omitempty"`
}

func (x *PopularResultReq) Reset() {
	*x = PopularResultReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_show_popular_v1_popular_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopularResultReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopularResultReq) ProtoMessage() {}

func (x *PopularResultReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_show_popular_v1_popular_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopularResultReq.ProtoReflect.Descriptor instead.
func (*PopularResultReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_show_popular_v1_popular_proto_rawDescGZIP(), []int{4}
}

func (x *PopularResultReq) GetIdx() int64 {
	if x != nil {
		return x.Idx
	}
	return 0
}

func (x *PopularResultReq) GetLoginEvent() int32 {
	if x != nil {
		return x.LoginEvent
	}
	return 0
}

func (x *PopularResultReq) GetQn() int32 {
	if x != nil {
		return x.Qn
	}
	return 0
}

func (x *PopularResultReq) GetFnver() int32 {
	if x != nil {
		return x.Fnver
	}
	return 0
}

func (x *PopularResultReq) GetFnval() int32 {
	if x != nil {
		return x.Fnval
	}
	return 0
}

func (x *PopularResultReq) GetForceHost() int32 {
	if x != nil {
		return x.ForceHost
	}
	return 0
}

func (x *PopularResultReq) GetFourk() int32 {
	if x != nil {
		return x.Fourk
	}
	return 0
}

func (x *PopularResultReq) GetSpmid() string {
	if x != nil {
		return x.Spmid
	}
	return ""
}

func (x *PopularResultReq) GetLastParam() string {
	if x != nil {
		return x.LastParam
	}
	return ""
}

func (x *PopularResultReq) GetVer() string {
	if x != nil {
		return x.Ver
	}
	return ""
}

func (x *PopularResultReq) GetEntranceId() int64 {
	if x != nil {
		return x.EntranceId
	}
	return 0
}

func (x *PopularResultReq) GetLocationIds() string {
	if x != nil {
		return x.LocationIds
	}
	return ""
}

func (x *PopularResultReq) GetSourceId() int32 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *PopularResultReq) GetFlush() int32 {
	if x != nil {
		return x.Flush
	}
	return 0
}

func (x *PopularResultReq) GetPlayerArgs() *v11.PlayerArgs {
	if x != nil {
		return x.PlayerArgs
	}
	return nil
}

var File_bilibili_app_show_popular_v1_popular_proto protoreflect.FileDescriptor

var file_bilibili_app_show_popular_v1_popular_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x68, 0x6f, 0x77, 0x2f, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x63, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x06, 0x42, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x62, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xd1, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x6f,
	0x74, 0x74, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x26, 0x0a,
	0x0f, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x54, 0x65,
	0x78, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x3f, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x08, 0x74, 0x6f,
	0x70, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x65, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x69, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x68, 0x69, 0x74, 0x22, 0x80, 0x02, 0x0a, 0x0c, 0x45,
	0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x12, 0x34, 0x0a, 0x06, 0x62, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x62, 0x62,
	0x6c, 0x65, 0x52, 0x06, 0x62, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x6f, 0x70, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x6f, 0x70, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x88, 0x01,
	0x0a, 0x0c, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x34, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x73, 0x68, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x65, 0x72, 0x22, 0xc5, 0x03, 0x0a, 0x10, 0x50, 0x6f, 0x70,
	0x75, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x78, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x71, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x71, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x6e, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x66, 0x6e, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6e, 0x76, 0x61, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6e, 0x76, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x6f, 0x75, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x72,
	0x6b, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x6d, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x70, 0x6d, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x65, 0x72, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x75,
	0x73, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c, 0x75, 0x73, 0x68, 0x12,
	0x4f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x41, 0x72, 0x67, 0x73, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73,
	0x32, 0x5e, 0x0a, 0x07, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x53, 0x0a, 0x05, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x70, 0x75,
	0x6c, 0x61, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58,
	0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x68, 0x6f, 0x77,
	0x2f, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_show_popular_v1_popular_proto_rawDescOnce sync.Once
	file_bilibili_app_show_popular_v1_popular_proto_rawDescData = file_bilibili_app_show_popular_v1_popular_proto_rawDesc
)

func file_bilibili_app_show_popular_v1_popular_proto_rawDescGZIP() []byte {
	file_bilibili_app_show_popular_v1_popular_proto_rawDescOnce.Do(func() {
		file_bilibili_app_show_popular_v1_popular_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_show_popular_v1_popular_proto_rawDescData)
	})
	return file_bilibili_app_show_popular_v1_popular_proto_rawDescData
}

var file_bilibili_app_show_popular_v1_popular_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bilibili_app_show_popular_v1_popular_proto_goTypes = []interface{}{
	(*Bubble)(nil),           // 0: bilibili.app.show.v1.Bubble
	(*Config)(nil),           // 1: bilibili.app.show.v1.Config
	(*EntranceShow)(nil),     // 2: bilibili.app.show.v1.EntranceShow
	(*PopularReply)(nil),     // 3: bilibili.app.show.v1.PopularReply
	(*PopularResultReq)(nil), // 4: bilibili.app.show.v1.PopularResultReq
	(*v1.Card)(nil),          // 5: bilibili.app.card.v1.Card
	(*v11.PlayerArgs)(nil),   // 6: bilibili.app.archive.middleware.v1.PlayerArgs
}
var file_bilibili_app_show_popular_v1_popular_proto_depIdxs = []int32{
	2, // 0: bilibili.app.show.v1.Config.top_items:type_name -> bilibili.app.show.v1.EntranceShow
	2, // 1: bilibili.app.show.v1.Config.page_items:type_name -> bilibili.app.show.v1.EntranceShow
	0, // 2: bilibili.app.show.v1.EntranceShow.bubble:type_name -> bilibili.app.show.v1.Bubble
	5, // 3: bilibili.app.show.v1.PopularReply.items:type_name -> bilibili.app.card.v1.Card
	1, // 4: bilibili.app.show.v1.PopularReply.config:type_name -> bilibili.app.show.v1.Config
	6, // 5: bilibili.app.show.v1.PopularResultReq.player_args:type_name -> bilibili.app.archive.middleware.v1.PlayerArgs
	4, // 6: bilibili.app.show.v1.Popular.Index:input_type -> bilibili.app.show.v1.PopularResultReq
	3, // 7: bilibili.app.show.v1.Popular.Index:output_type -> bilibili.app.show.v1.PopularReply
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_bilibili_app_show_popular_v1_popular_proto_init() }
func file_bilibili_app_show_popular_v1_popular_proto_init() {
	if File_bilibili_app_show_popular_v1_popular_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_show_popular_v1_popular_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bubble); i {
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
		file_bilibili_app_show_popular_v1_popular_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_bilibili_app_show_popular_v1_popular_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntranceShow); i {
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
		file_bilibili_app_show_popular_v1_popular_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopularReply); i {
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
		file_bilibili_app_show_popular_v1_popular_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopularResultReq); i {
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
			RawDescriptor: file_bilibili_app_show_popular_v1_popular_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_app_show_popular_v1_popular_proto_goTypes,
		DependencyIndexes: file_bilibili_app_show_popular_v1_popular_proto_depIdxs,
		MessageInfos:      file_bilibili_app_show_popular_v1_popular_proto_msgTypes,
	}.Build()
	File_bilibili_app_show_popular_v1_popular_proto = out.File
	file_bilibili_app_show_popular_v1_popular_proto_rawDesc = nil
	file_bilibili_app_show_popular_v1_popular_proto_goTypes = nil
	file_bilibili_app_show_popular_v1_popular_proto_depIdxs = nil
}
