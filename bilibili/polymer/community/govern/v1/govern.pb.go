// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/polymer/community/govern/v1/govern.proto

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

type AntiHarassmentLimit int32

const (
	AntiHarassmentLimit_DefaultLimit  AntiHarassmentLimit = 0 //
	AntiHarassmentLimit_FollowLimit   AntiHarassmentLimit = 1 //
	AntiHarassmentLimit_ReFollowLimit AntiHarassmentLimit = 2 //
	AntiHarassmentLimit_TwoWayFollow  AntiHarassmentLimit = 3 //
	AntiHarassmentLimit_AllLimit      AntiHarassmentLimit = 4 //
)

// Enum value maps for AntiHarassmentLimit.
var (
	AntiHarassmentLimit_name = map[int32]string{
		0: "DefaultLimit",
		1: "FollowLimit",
		2: "ReFollowLimit",
		3: "TwoWayFollow",
		4: "AllLimit",
	}
	AntiHarassmentLimit_value = map[string]int32{
		"DefaultLimit":  0,
		"FollowLimit":   1,
		"ReFollowLimit": 2,
		"TwoWayFollow":  3,
		"AllLimit":      4,
	}
)

func (x AntiHarassmentLimit) Enum() *AntiHarassmentLimit {
	p := new(AntiHarassmentLimit)
	*p = x
	return p
}

func (x AntiHarassmentLimit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AntiHarassmentLimit) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_polymer_community_govern_v1_govern_proto_enumTypes[0].Descriptor()
}

func (AntiHarassmentLimit) Type() protoreflect.EnumType {
	return &file_bilibili_polymer_community_govern_v1_govern_proto_enumTypes[0]
}

func (x AntiHarassmentLimit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AntiHarassmentLimit.Descriptor instead.
func (AntiHarassmentLimit) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_polymer_community_govern_v1_govern_proto_rawDescGZIP(), []int{0}
}

type BizType int32

const (
	BizType_InvalidBizType BizType = 0 //
	BizType_Im             BizType = 1 //
	BizType_Dm             BizType = 2 //
	BizType_Reply          BizType = 3 //
	BizType_ReplyMe        BizType = 4 //
	BizType_LikeMe         BizType = 5 //
	BizType_AtMe           BizType = 6 //
)

// Enum value maps for BizType.
var (
	BizType_name = map[int32]string{
		0: "InvalidBizType",
		1: "Im",
		2: "Dm",
		3: "Reply",
		4: "ReplyMe",
		5: "LikeMe",
		6: "AtMe",
	}
	BizType_value = map[string]int32{
		"InvalidBizType": 0,
		"Im":             1,
		"Dm":             2,
		"Reply":          3,
		"ReplyMe":        4,
		"LikeMe":         5,
		"AtMe":           6,
	}
)

func (x BizType) Enum() *BizType {
	p := new(BizType)
	*p = x
	return p
}

func (x BizType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BizType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_polymer_community_govern_v1_govern_proto_enumTypes[1].Descriptor()
}

func (BizType) Type() protoreflect.EnumType {
	return &file_bilibili_polymer_community_govern_v1_govern_proto_enumTypes[1]
}

func (x BizType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BizType.Descriptor instead.
func (BizType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_polymer_community_govern_v1_govern_proto_rawDescGZIP(), []int{1}
}

type AntiHarassmentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit                 int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	FollowTimeLimitSecond int32 `protobuf:"varint,2,opt,name=follow_time_limit_second,json=followTimeLimitSecond,proto3" json:"follow_time_limit_second,omitempty"`
	ExpireTime            int64 `protobuf:"varint,3,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *AntiHarassmentInfo) Reset() {
	*x = AntiHarassmentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AntiHarassmentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AntiHarassmentInfo) ProtoMessage() {}

func (x *AntiHarassmentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AntiHarassmentInfo.ProtoReflect.Descriptor instead.
func (*AntiHarassmentInfo) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_community_govern_v1_govern_proto_rawDescGZIP(), []int{0}
}

func (x *AntiHarassmentInfo) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *AntiHarassmentInfo) GetFollowTimeLimitSecond() int32 {
	if x != nil {
		return x.FollowTimeLimitSecond
	}
	return 0
}

func (x *AntiHarassmentInfo) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

type AntiHarassmentSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mid                 int64               `protobuf:"varint,1,opt,name=mid,proto3" json:"mid,omitempty"`
	AutoLimit           bool                `protobuf:"varint,2,opt,name=auto_limit,json=autoLimit,proto3" json:"auto_limit,omitempty"`
	Im                  *AntiHarassmentInfo `protobuf:"bytes,3,opt,name=im,proto3" json:"im,omitempty"`
	Reply               *AntiHarassmentInfo `protobuf:"bytes,4,opt,name=reply,proto3" json:"reply,omitempty"`
	Dm                  *AntiHarassmentInfo `protobuf:"bytes,5,opt,name=dm,proto3" json:"dm,omitempty"`
	ReplyMe             *AntiHarassmentInfo `protobuf:"bytes,6,opt,name=reply_me,json=replyMe,proto3" json:"reply_me,omitempty"`
	LikeMe              *AntiHarassmentInfo `protobuf:"bytes,7,opt,name=like_me,json=likeMe,proto3" json:"like_me,omitempty"`
	AtMe                *AntiHarassmentInfo `protobuf:"bytes,8,opt,name=at_me,json=atMe,proto3" json:"at_me,omitempty"`
	AutoLimitExpireTime int64               `protobuf:"varint,9,opt,name=auto_limit_expire_time,json=autoLimitExpireTime,proto3" json:"auto_limit_expire_time,omitempty"`
}

func (x *AntiHarassmentSetting) Reset() {
	*x = AntiHarassmentSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AntiHarassmentSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AntiHarassmentSetting) ProtoMessage() {}

func (x *AntiHarassmentSetting) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AntiHarassmentSetting.ProtoReflect.Descriptor instead.
func (*AntiHarassmentSetting) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_community_govern_v1_govern_proto_rawDescGZIP(), []int{1}
}

func (x *AntiHarassmentSetting) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *AntiHarassmentSetting) GetAutoLimit() bool {
	if x != nil {
		return x.AutoLimit
	}
	return false
}

func (x *AntiHarassmentSetting) GetIm() *AntiHarassmentInfo {
	if x != nil {
		return x.Im
	}
	return nil
}

func (x *AntiHarassmentSetting) GetReply() *AntiHarassmentInfo {
	if x != nil {
		return x.Reply
	}
	return nil
}

func (x *AntiHarassmentSetting) GetDm() *AntiHarassmentInfo {
	if x != nil {
		return x.Dm
	}
	return nil
}

func (x *AntiHarassmentSetting) GetReplyMe() *AntiHarassmentInfo {
	if x != nil {
		return x.ReplyMe
	}
	return nil
}

func (x *AntiHarassmentSetting) GetLikeMe() *AntiHarassmentInfo {
	if x != nil {
		return x.LikeMe
	}
	return nil
}

func (x *AntiHarassmentSetting) GetAtMe() *AntiHarassmentInfo {
	if x != nil {
		return x.AtMe
	}
	return nil
}

func (x *AntiHarassmentSetting) GetAutoLimitExpireTime() int64 {
	if x != nil {
		return x.AutoLimitExpireTime
	}
	return 0
}

type LoadAntiHarassmentSettingsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizType int32 `protobuf:"varint,1,opt,name=biz_type,json=bizType,proto3" json:"biz_type,omitempty"`
	RecvMid int64 `protobuf:"varint,2,opt,name=recv_mid,json=recvMid,proto3" json:"recv_mid,omitempty"`
	SendMid int64 `protobuf:"varint,3,opt,name=send_mid,json=sendMid,proto3" json:"send_mid,omitempty"`
}

func (x *LoadAntiHarassmentSettingsReq) Reset() {
	*x = LoadAntiHarassmentSettingsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadAntiHarassmentSettingsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadAntiHarassmentSettingsReq) ProtoMessage() {}

func (x *LoadAntiHarassmentSettingsReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadAntiHarassmentSettingsReq.ProtoReflect.Descriptor instead.
func (*LoadAntiHarassmentSettingsReq) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_community_govern_v1_govern_proto_rawDescGZIP(), []int{2}
}

func (x *LoadAntiHarassmentSettingsReq) GetBizType() int32 {
	if x != nil {
		return x.BizType
	}
	return 0
}

func (x *LoadAntiHarassmentSettingsReq) GetRecvMid() int64 {
	if x != nil {
		return x.RecvMid
	}
	return 0
}

func (x *LoadAntiHarassmentSettingsReq) GetSendMid() int64 {
	if x != nil {
		return x.SendMid
	}
	return 0
}

type LoadAntiHarassmentSettingsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AntiHarassmentRet     bool                   `protobuf:"varint,1,opt,name=anti_harassment_ret,json=antiHarassmentRet,proto3" json:"anti_harassment_ret,omitempty"`
	AntiHarassmentSetting *AntiHarassmentSetting `protobuf:"bytes,2,opt,name=anti_harassment_setting,json=antiHarassmentSetting,proto3" json:"anti_harassment_setting,omitempty"`
	ShowWindow            int32                  `protobuf:"varint,3,opt,name=show_window,json=showWindow,proto3" json:"show_window,omitempty"`
}

func (x *LoadAntiHarassmentSettingsRsp) Reset() {
	*x = LoadAntiHarassmentSettingsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadAntiHarassmentSettingsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadAntiHarassmentSettingsRsp) ProtoMessage() {}

func (x *LoadAntiHarassmentSettingsRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadAntiHarassmentSettingsRsp.ProtoReflect.Descriptor instead.
func (*LoadAntiHarassmentSettingsRsp) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_community_govern_v1_govern_proto_rawDescGZIP(), []int{3}
}

func (x *LoadAntiHarassmentSettingsRsp) GetAntiHarassmentRet() bool {
	if x != nil {
		return x.AntiHarassmentRet
	}
	return false
}

func (x *LoadAntiHarassmentSettingsRsp) GetAntiHarassmentSetting() *AntiHarassmentSetting {
	if x != nil {
		return x.AntiHarassmentSetting
	}
	return nil
}

func (x *LoadAntiHarassmentSettingsRsp) GetShowWindow() int32 {
	if x != nil {
		return x.ShowWindow
	}
	return 0
}

type StoreAntiHarassmentSettingsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizType               int32                  `protobuf:"varint,1,opt,name=biz_type,json=bizType,proto3" json:"biz_type,omitempty"`
	Mid                   int64                  `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	AntiHarassmentSetting *AntiHarassmentSetting `protobuf:"bytes,3,opt,name=anti_harassment_setting,json=antiHarassmentSetting,proto3" json:"anti_harassment_setting,omitempty"`
}

func (x *StoreAntiHarassmentSettingsReq) Reset() {
	*x = StoreAntiHarassmentSettingsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreAntiHarassmentSettingsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreAntiHarassmentSettingsReq) ProtoMessage() {}

func (x *StoreAntiHarassmentSettingsReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreAntiHarassmentSettingsReq.ProtoReflect.Descriptor instead.
func (*StoreAntiHarassmentSettingsReq) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_community_govern_v1_govern_proto_rawDescGZIP(), []int{4}
}

func (x *StoreAntiHarassmentSettingsReq) GetBizType() int32 {
	if x != nil {
		return x.BizType
	}
	return 0
}

func (x *StoreAntiHarassmentSettingsReq) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *StoreAntiHarassmentSettingsReq) GetAntiHarassmentSetting() *AntiHarassmentSetting {
	if x != nil {
		return x.AntiHarassmentSetting
	}
	return nil
}

type StoreAntiHarassmentSettingsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreAntiHarassmentSettingsRsp) Reset() {
	*x = StoreAntiHarassmentSettingsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreAntiHarassmentSettingsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreAntiHarassmentSettingsRsp) ProtoMessage() {}

func (x *StoreAntiHarassmentSettingsRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreAntiHarassmentSettingsRsp.ProtoReflect.Descriptor instead.
func (*StoreAntiHarassmentSettingsRsp) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_community_govern_v1_govern_proto_rawDescGZIP(), []int{5}
}

var File_bilibili_polymer_community_govern_v1_govern_proto protoreflect.FileDescriptor

var file_bilibili_polymer_community_govern_v1_govern_proto_rawDesc = []byte{
	0x0a, 0x31, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x6d,
	0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x76,
	0x65, 0x72, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f,
	0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e,
	0x2e, 0x76, 0x31, 0x22, 0x84, 0x01, 0x0a, 0x12, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61,
	0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x37, 0x0a, 0x18, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x15, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb4, 0x04, 0x0a, 0x15, 0x41,
	0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x75, 0x74, 0x6f,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x42, 0x0a, 0x02, 0x69, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c,
	0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x02, 0x69, 0x6d, 0x12, 0x48, 0x0a, 0x05, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61,
	0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x42, 0x0a, 0x02, 0x64, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x02, 0x64, 0x6d, 0x12, 0x4d, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x5f, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x74, 0x69, 0x48,
	0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x12, 0x4b, 0x0a, 0x07, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67,
	0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72,
	0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6c, 0x69, 0x6b,
	0x65, 0x4d, 0x65, 0x12, 0x47, 0x0a, 0x05, 0x61, 0x74, 0x5f, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f,
	0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x61, 0x74, 0x4d, 0x65, 0x12, 0x33, 0x0a, 0x16,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x61, 0x75,
	0x74, 0x6f, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x70, 0x0a, 0x1d, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72,
	0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x7a, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x69, 0x7a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x72, 0x65, 0x63, 0x76, 0x5f, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x72, 0x65, 0x63, 0x76, 0x4d, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x6d, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x64,
	0x4d, 0x69, 0x64, 0x22, 0xdf, 0x01, 0x0a, 0x1d, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x6e, 0x74, 0x69,
	0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x6e, 0x74, 0x69, 0x5f, 0x68, 0x61,
	0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x61, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x74, 0x12, 0x6d, 0x0a, 0x17, 0x61, 0x6e, 0x74, 0x69, 0x5f, 0x68, 0x61,
	0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f,
	0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61,
	0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x15, 0x61,
	0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x22, 0xbc, 0x01, 0x0a, 0x1e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41,
	0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x7a, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x69, 0x7a, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x6d, 0x0a, 0x17, 0x61, 0x6e, 0x74, 0x69, 0x5f, 0x68, 0x61,
	0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f,
	0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61,
	0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x15, 0x61,
	0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x22, 0x20, 0x0a, 0x1e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x6e, 0x74,
	0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x73, 0x70, 0x2a, 0x6b, 0x0a, 0x13, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61,
	0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x10, 0x0a,
	0x0c, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x77, 0x6f, 0x57, 0x61, 0x79, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x10, 0x04, 0x2a, 0x5b, 0x0a, 0x07, 0x42, 0x69, 0x7a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x42, 0x69, 0x7a, 0x54, 0x79, 0x70, 0x65,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x6d, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x44, 0x6d,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x69,
	0x6b, 0x65, 0x4d, 0x65, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x74, 0x4d, 0x65, 0x10, 0x06,
	0x32, 0xd4, 0x02, 0x0a, 0x15, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x1b, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3e, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x3e, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x73, 0x70, 0x12, 0x9a, 0x01, 0x0a, 0x1a, 0x4c,
	0x6f, 0x61, 0x64, 0x41, 0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3d, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x41,
	0x6e, 0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x3d, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x6e,
	0x74, 0x69, 0x48, 0x61, 0x72, 0x61, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x73, 0x70, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31,
	0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x70,
	0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2f, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_bilibili_polymer_community_govern_v1_govern_proto_rawDescOnce sync.Once
	file_bilibili_polymer_community_govern_v1_govern_proto_rawDescData = file_bilibili_polymer_community_govern_v1_govern_proto_rawDesc
)

func file_bilibili_polymer_community_govern_v1_govern_proto_rawDescGZIP() []byte {
	file_bilibili_polymer_community_govern_v1_govern_proto_rawDescOnce.Do(func() {
		file_bilibili_polymer_community_govern_v1_govern_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_polymer_community_govern_v1_govern_proto_rawDescData)
	})
	return file_bilibili_polymer_community_govern_v1_govern_proto_rawDescData
}

var file_bilibili_polymer_community_govern_v1_govern_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bilibili_polymer_community_govern_v1_govern_proto_goTypes = []interface{}{
	(AntiHarassmentLimit)(0),               // 0: bilibili.polymer.app.govern.v1.AntiHarassmentLimit
	(BizType)(0),                           // 1: bilibili.polymer.app.govern.v1.BizType
	(*AntiHarassmentInfo)(nil),             // 2: bilibili.polymer.app.govern.v1.AntiHarassmentInfo
	(*AntiHarassmentSetting)(nil),          // 3: bilibili.polymer.app.govern.v1.AntiHarassmentSetting
	(*LoadAntiHarassmentSettingsReq)(nil),  // 4: bilibili.polymer.app.govern.v1.LoadAntiHarassmentSettingsReq
	(*LoadAntiHarassmentSettingsRsp)(nil),  // 5: bilibili.polymer.app.govern.v1.LoadAntiHarassmentSettingsRsp
	(*StoreAntiHarassmentSettingsReq)(nil), // 6: bilibili.polymer.app.govern.v1.StoreAntiHarassmentSettingsReq
	(*StoreAntiHarassmentSettingsRsp)(nil), // 7: bilibili.polymer.app.govern.v1.StoreAntiHarassmentSettingsRsp
}
var file_bilibili_polymer_community_govern_v1_govern_proto_depIdxs = []int32{
	2,  // 0: bilibili.polymer.app.govern.v1.AntiHarassmentSetting.im:type_name -> bilibili.polymer.app.govern.v1.AntiHarassmentInfo
	2,  // 1: bilibili.polymer.app.govern.v1.AntiHarassmentSetting.reply:type_name -> bilibili.polymer.app.govern.v1.AntiHarassmentInfo
	2,  // 2: bilibili.polymer.app.govern.v1.AntiHarassmentSetting.dm:type_name -> bilibili.polymer.app.govern.v1.AntiHarassmentInfo
	2,  // 3: bilibili.polymer.app.govern.v1.AntiHarassmentSetting.reply_me:type_name -> bilibili.polymer.app.govern.v1.AntiHarassmentInfo
	2,  // 4: bilibili.polymer.app.govern.v1.AntiHarassmentSetting.like_me:type_name -> bilibili.polymer.app.govern.v1.AntiHarassmentInfo
	2,  // 5: bilibili.polymer.app.govern.v1.AntiHarassmentSetting.at_me:type_name -> bilibili.polymer.app.govern.v1.AntiHarassmentInfo
	3,  // 6: bilibili.polymer.app.govern.v1.LoadAntiHarassmentSettingsRsp.anti_harassment_setting:type_name -> bilibili.polymer.app.govern.v1.AntiHarassmentSetting
	3,  // 7: bilibili.polymer.app.govern.v1.StoreAntiHarassmentSettingsReq.anti_harassment_setting:type_name -> bilibili.polymer.app.govern.v1.AntiHarassmentSetting
	6,  // 8: bilibili.polymer.app.govern.v1.AntiHarassmentService.StoreAntiHarassmentSettings:input_type -> bilibili.polymer.app.govern.v1.StoreAntiHarassmentSettingsReq
	4,  // 9: bilibili.polymer.app.govern.v1.AntiHarassmentService.LoadAntiHarassmentSettings:input_type -> bilibili.polymer.app.govern.v1.LoadAntiHarassmentSettingsReq
	7,  // 10: bilibili.polymer.app.govern.v1.AntiHarassmentService.StoreAntiHarassmentSettings:output_type -> bilibili.polymer.app.govern.v1.StoreAntiHarassmentSettingsRsp
	5,  // 11: bilibili.polymer.app.govern.v1.AntiHarassmentService.LoadAntiHarassmentSettings:output_type -> bilibili.polymer.app.govern.v1.LoadAntiHarassmentSettingsRsp
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_bilibili_polymer_community_govern_v1_govern_proto_init() }
func file_bilibili_polymer_community_govern_v1_govern_proto_init() {
	if File_bilibili_polymer_community_govern_v1_govern_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AntiHarassmentInfo); i {
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
		file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AntiHarassmentSetting); i {
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
		file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadAntiHarassmentSettingsReq); i {
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
		file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadAntiHarassmentSettingsRsp); i {
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
		file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreAntiHarassmentSettingsReq); i {
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
		file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreAntiHarassmentSettingsRsp); i {
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
			RawDescriptor: file_bilibili_polymer_community_govern_v1_govern_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_polymer_community_govern_v1_govern_proto_goTypes,
		DependencyIndexes: file_bilibili_polymer_community_govern_v1_govern_proto_depIdxs,
		EnumInfos:         file_bilibili_polymer_community_govern_v1_govern_proto_enumTypes,
		MessageInfos:      file_bilibili_polymer_community_govern_v1_govern_proto_msgTypes,
	}.Build()
	File_bilibili_polymer_community_govern_v1_govern_proto = out.File
	file_bilibili_polymer_community_govern_v1_govern_proto_rawDesc = nil
	file_bilibili_polymer_community_govern_v1_govern_proto_goTypes = nil
	file_bilibili_polymer_community_govern_v1_govern_proto_depIdxs = nil
}