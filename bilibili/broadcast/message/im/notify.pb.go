// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: bilibili/broadcast/message/im/notify.proto

package im

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

type PLType int32

const (
	PLType_EN_PAYLOAD_NORMAL PLType = 0
	PLType_EN_PAYLOAD_BASE64 PLType = 1
)

// Enum value maps for PLType.
var (
	PLType_name = map[int32]string{
		0: "EN_PAYLOAD_NORMAL",
		1: "EN_PAYLOAD_BASE64",
	}
	PLType_value = map[string]int32{
		"EN_PAYLOAD_NORMAL": 0,
		"EN_PAYLOAD_BASE64": 1,
	}
)

func (x PLType) Enum() *PLType {
	p := new(PLType)
	*p = x
	return p
}

func (x PLType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PLType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_broadcast_message_im_notify_proto_enumTypes[0].Descriptor()
}

func (PLType) Type() protoreflect.EnumType {
	return &file_bilibili_broadcast_message_im_notify_proto_enumTypes[0]
}

func (x PLType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PLType.Descriptor instead.
func (PLType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_im_notify_proto_rawDescGZIP(), []int{0}
}

type CmdId int32

const (
	// 非法cmd
	CmdId_EN_CMD_ID_INVALID CmdId = 0
	// 服务端主动发起
	CmdId_EN_CMD_ID_MSG_NOTIFY CmdId = 1
	CmdId_EN_CMD_ID_KICK_OUT CmdId = 2
)

// Enum value maps for CmdId.
var (
	CmdId_name = map[int32]string{
		0: "EN_CMD_ID_INVALID",
		1: "EN_CMD_ID_MSG_NOTIFY",
		2: "EN_CMD_ID_KICK_OUT",
	}
	CmdId_value = map[string]int32{
		"EN_CMD_ID_INVALID":    0,
		"EN_CMD_ID_MSG_NOTIFY": 1,
		"EN_CMD_ID_KICK_OUT":   2,
	}
)

func (x CmdId) Enum() *CmdId {
	p := new(CmdId)
	*p = x
	return p
}

func (x CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_broadcast_message_im_notify_proto_enumTypes[1].Descriptor()
}

func (CmdId) Type() protoreflect.EnumType {
	return &file_bilibili_broadcast_message_im_notify_proto_enumTypes[1]
}

func (x CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdId.Descriptor instead.
func (CmdId) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_im_notify_proto_rawDescGZIP(), []int{1}
}

type NotifyRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// 命令id
	Cmd uint64 `protobuf:"varint,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	PayloadType PLType `protobuf:"varint,4,opt,name=payload_type,json=payloadType,proto3,enum=bilibili.broadcast.message.im.PLType" json:"payload_type,omitempty"`
}

func (x *NotifyRsp) Reset() {
	*x = NotifyRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_im_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRsp) ProtoMessage() {}

func (x *NotifyRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_im_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRsp.ProtoReflect.Descriptor instead.
func (*NotifyRsp) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_im_notify_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyRsp) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *NotifyRsp) GetCmd() uint64 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *NotifyRsp) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *NotifyRsp) GetPayloadType() PLType {
	if x != nil {
		return x.PayloadType
	}
	return PLType_EN_PAYLOAD_NORMAL
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 发送方uid
	SenderUid uint64 `protobuf:"varint,1,opt,name=sender_uid,json=senderUid,proto3" json:"sender_uid,omitempty"`
	// 接收方类型
	ReceiverType int32 `protobuf:"varint,2,opt,name=receiver_type,json=receiverType,proto3" json:"receiver_type,omitempty"`
	// 接收方id
	ReceiverId uint64 `protobuf:"varint,3,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	// 客户端的序列id 用于服务端去重
	CliMsgId uint64 `protobuf:"varint,4,opt,name=cli_msg_id,json=cliMsgId,proto3" json:"cli_msg_id,omitempty"`
	// 消息类型
	MsgType int32 `protobuf:"varint,5,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	// 消息内容
	Content string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	// 服务端的序列号
	MsgSeqno uint64 `protobuf:"varint,7,opt,name=msg_seqno,json=msgSeqno,proto3" json:"msg_seqno,omitempty"`
	// 消息发送时间（服务端时间）
	Timestamp uint64 `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// at用户列表
	AtUids []uint64 `protobuf:"varint,9,rep,packed,name=at_uids,json=atUids,proto3" json:"at_uids,omitempty"`
	// 多人消息
	RecverIds []uint64 `protobuf:"varint,10,rep,packed,name=recver_ids,json=recverIds,proto3" json:"recver_ids,omitempty"`
	// 消息唯一标示
	MsgKey uint64 `protobuf:"varint,11,opt,name=msg_key,json=msgKey,proto3" json:"msg_key,omitempty"`
	// 消息状态
	MsgStatus uint32 `protobuf:"varint,12,opt,name=msg_status,json=msgStatus,proto3" json:"msg_status,omitempty"`
	// 是否为系统撤销
	SysCancel bool `protobuf:"varint,13,opt,name=sys_cancel,json=sysCancel,proto3" json:"sys_cancel,omitempty"`
	// 是否是多聊消息 目前群通知管理员的部分通知属于该类消息
	IsMultiChat uint32 `protobuf:"varint,14,opt,name=is_multi_chat,json=isMultiChat,proto3" json:"is_multi_chat,omitempty"`
	// 表示撤回的消息的session_seqno 用以后续的比较 实现未读数的正确显示
	WithdrawSeqno uint64 `protobuf:"varint,15,opt,name=withdraw_seqno,json=withdrawSeqno,proto3" json:"withdraw_seqno,omitempty"`
	// 通知码
	NotifyCode string `protobuf:"bytes,16,opt,name=notify_code,json=notifyCode,proto3" json:"notify_code,omitempty"`
	// 消息来源
	MsgSource uint32 `protobuf:"varint,17,opt,name=msg_source,json=msgSource,proto3" json:"msg_source,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_im_notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_im_notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_im_notify_proto_rawDescGZIP(), []int{1}
}

func (x *Msg) GetSenderUid() uint64 {
	if x != nil {
		return x.SenderUid
	}
	return 0
}

func (x *Msg) GetReceiverType() int32 {
	if x != nil {
		return x.ReceiverType
	}
	return 0
}

func (x *Msg) GetReceiverId() uint64 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *Msg) GetCliMsgId() uint64 {
	if x != nil {
		return x.CliMsgId
	}
	return 0
}

func (x *Msg) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *Msg) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Msg) GetMsgSeqno() uint64 {
	if x != nil {
		return x.MsgSeqno
	}
	return 0
}

func (x *Msg) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Msg) GetAtUids() []uint64 {
	if x != nil {
		return x.AtUids
	}
	return nil
}

func (x *Msg) GetRecverIds() []uint64 {
	if x != nil {
		return x.RecverIds
	}
	return nil
}

func (x *Msg) GetMsgKey() uint64 {
	if x != nil {
		return x.MsgKey
	}
	return 0
}

func (x *Msg) GetMsgStatus() uint32 {
	if x != nil {
		return x.MsgStatus
	}
	return 0
}

func (x *Msg) GetSysCancel() bool {
	if x != nil {
		return x.SysCancel
	}
	return false
}

func (x *Msg) GetIsMultiChat() uint32 {
	if x != nil {
		return x.IsMultiChat
	}
	return 0
}

func (x *Msg) GetWithdrawSeqno() uint64 {
	if x != nil {
		return x.WithdrawSeqno
	}
	return 0
}

func (x *Msg) GetNotifyCode() string {
	if x != nil {
		return x.NotifyCode
	}
	return ""
}

func (x *Msg) GetMsgSource() uint32 {
	if x != nil {
		return x.MsgSource
	}
	return 0
}

type NotifyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType uint32 `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	TalkerId uint64 `protobuf:"varint,2,opt,name=talker_id,json=talkerId,proto3" json:"talker_id,omitempty"`
	SessionType uint32 `protobuf:"varint,3,opt,name=session_type,json=sessionType,proto3" json:"session_type,omitempty"`
}

func (x *NotifyInfo) Reset() {
	*x = NotifyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_im_notify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyInfo) ProtoMessage() {}

func (x *NotifyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_im_notify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyInfo.ProtoReflect.Descriptor instead.
func (*NotifyInfo) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_im_notify_proto_rawDescGZIP(), []int{2}
}

func (x *NotifyInfo) GetMsgType() uint32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *NotifyInfo) GetTalkerId() uint64 {
	if x != nil {
		return x.TalkerId
	}
	return 0
}

func (x *NotifyInfo) GetSessionType() uint32 {
	if x != nil {
		return x.SessionType
	}
	return 0
}

type ReqServerNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 最新序列号
	LastestSeqno uint64 `protobuf:"varint,1,opt,name=lastest_seqno,json=lastestSeqno,proto3" json:"lastest_seqno,omitempty"`
	// 即时消息 该类消息主要用于系统通知 当客户端sync msg时 不会sync到此类消息
	InstantMsg *Msg `protobuf:"bytes,2,opt,name=instant_msg,json=instantMsg,proto3" json:"instant_msg,omitempty"`
	NotifyInfo *NotifyInfo `protobuf:"bytes,3,opt,name=notify_info,json=notifyInfo,proto3" json:"notify_info,omitempty"`
}

func (x *ReqServerNotify) Reset() {
	*x = ReqServerNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_im_notify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqServerNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqServerNotify) ProtoMessage() {}

func (x *ReqServerNotify) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_im_notify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqServerNotify.ProtoReflect.Descriptor instead.
func (*ReqServerNotify) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_im_notify_proto_rawDescGZIP(), []int{3}
}

func (x *ReqServerNotify) GetLastestSeqno() uint64 {
	if x != nil {
		return x.LastestSeqno
	}
	return 0
}

func (x *ReqServerNotify) GetInstantMsg() *Msg {
	if x != nil {
		return x.InstantMsg
	}
	return nil
}

func (x *ReqServerNotify) GetNotifyInfo() *NotifyInfo {
	if x != nil {
		return x.NotifyInfo
	}
	return nil
}

var File_bilibili_broadcast_message_im_notify_proto protoreflect.FileDescriptor

var file_bilibili_broadcast_message_im_notify_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x6d, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x6d, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x09, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x48, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x6d, 0x2e, 0x50, 0x4c, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x92,
	0x04, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x63,
	0x6c, 0x69, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x65, 0x71, 0x6e, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x53, 0x65, 0x71, 0x6e, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x74, 0x5f,
	0x75, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x61, 0x74, 0x55, 0x69,
	0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65, 0x63, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6d, 0x73, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x73,
	0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x6d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x79, 0x73,
	0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73,
	0x79, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x69, 0x73, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x68, 0x61, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x73, 0x65, 0x71, 0x6e, 0x6f, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x65,
	0x71, 0x6e, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x73, 0x67, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x67, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x61, 0x6c, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x74, 0x61, 0x6c, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc7, 0x01, 0x0a,
	0x0f, 0x52, 0x65, 0x71, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x6e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x71, 0x6e, 0x6f, 0x12, 0x43, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x6d, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x0a,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x4a, 0x0a, 0x0b, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x6d, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x36, 0x0a, 0x06, 0x50, 0x4c, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x45, 0x4e, 0x5f, 0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x4e,
	0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x4e, 0x5f, 0x50, 0x41,
	0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x36, 0x34, 0x10, 0x01, 0x2a, 0x50,
	0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x4e, 0x5f, 0x43, 0x4d,
	0x44, 0x5f, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x18,
	0x0a, 0x14, 0x45, 0x4e, 0x5f, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x5f, 0x4d, 0x53, 0x47, 0x5f,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x4e, 0x5f, 0x43,
	0x4d, 0x44, 0x5f, 0x49, 0x44, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x02,
	0x32, 0x5b, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x51, 0x0a, 0x0b, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x28, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x69,
	0x6d, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x73, 0x70, 0x30, 0x01, 0x42, 0x4a, 0x5a,
	0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f,
	0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_bilibili_broadcast_message_im_notify_proto_rawDescOnce sync.Once
	file_bilibili_broadcast_message_im_notify_proto_rawDescData = file_bilibili_broadcast_message_im_notify_proto_rawDesc
)

func file_bilibili_broadcast_message_im_notify_proto_rawDescGZIP() []byte {
	file_bilibili_broadcast_message_im_notify_proto_rawDescOnce.Do(func() {
		file_bilibili_broadcast_message_im_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_broadcast_message_im_notify_proto_rawDescData)
	})
	return file_bilibili_broadcast_message_im_notify_proto_rawDescData
}

var file_bilibili_broadcast_message_im_notify_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bilibili_broadcast_message_im_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bilibili_broadcast_message_im_notify_proto_goTypes = []interface{}{
	(PLType)(0),             // 0: bilibili.broadcast.message.im.PLType
	(CmdId)(0),              // 1: bilibili.broadcast.message.im.CmdId
	(*NotifyRsp)(nil),       // 2: bilibili.broadcast.message.im.NotifyRsp
	(*Msg)(nil),             // 3: bilibili.broadcast.message.im.Msg
	(*NotifyInfo)(nil),      // 4: bilibili.broadcast.message.im.NotifyInfo
	(*ReqServerNotify)(nil), // 5: bilibili.broadcast.message.im.ReqServerNotify
	(*empty.Empty)(nil),     // 6: google.protobuf.Empty
}
var file_bilibili_broadcast_message_im_notify_proto_depIdxs = []int32{
	0, // 0: bilibili.broadcast.message.im.NotifyRsp.payload_type:type_name -> bilibili.broadcast.message.im.PLType
	3, // 1: bilibili.broadcast.message.im.ReqServerNotify.instant_msg:type_name -> bilibili.broadcast.message.im.Msg
	4, // 2: bilibili.broadcast.message.im.ReqServerNotify.notify_info:type_name -> bilibili.broadcast.message.im.NotifyInfo
	6, // 3: bilibili.broadcast.message.im.Notify.WatchNotify:input_type -> google.protobuf.Empty
	2, // 4: bilibili.broadcast.message.im.Notify.WatchNotify:output_type -> bilibili.broadcast.message.im.NotifyRsp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bilibili_broadcast_message_im_notify_proto_init() }
func file_bilibili_broadcast_message_im_notify_proto_init() {
	if File_bilibili_broadcast_message_im_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_broadcast_message_im_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyRsp); i {
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
		file_bilibili_broadcast_message_im_notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_bilibili_broadcast_message_im_notify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyInfo); i {
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
		file_bilibili_broadcast_message_im_notify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqServerNotify); i {
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
			RawDescriptor: file_bilibili_broadcast_message_im_notify_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_broadcast_message_im_notify_proto_goTypes,
		DependencyIndexes: file_bilibili_broadcast_message_im_notify_proto_depIdxs,
		EnumInfos:         file_bilibili_broadcast_message_im_notify_proto_enumTypes,
		MessageInfos:      file_bilibili_broadcast_message_im_notify_proto_msgTypes,
	}.Build()
	File_bilibili_broadcast_message_im_notify_proto = out.File
	file_bilibili_broadcast_message_im_notify_proto_rawDesc = nil
	file_bilibili_broadcast_message_im_notify_proto_goTypes = nil
	file_bilibili_broadcast_message_im_notify_proto_depIdxs = nil
}
