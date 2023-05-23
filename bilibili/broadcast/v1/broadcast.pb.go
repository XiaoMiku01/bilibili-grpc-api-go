// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: bilibili/broadcast/v1/broadcast.proto

package v1

import (
	rpc "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/rpc"
	any1 "github.com/golang/protobuf/ptypes/any"
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

type Action int32

const (
	Action_UNKNOWN Action = 0 //
	Action_UPDATE  Action = 1 //
	Action_DELETE  Action = 2 //
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "UNKNOWN",
		1: "UPDATE",
		2: "DELETE",
	}
	Action_value = map[string]int32{
		"UNKNOWN": 0,
		"UPDATE":  1,
		"DELETE":  2,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_broadcast_v1_broadcast_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_bilibili_broadcast_v1_broadcast_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_broadcast_proto_rawDescGZIP(), []int{0}
}

// 鉴权请求，通过authorization验证绑定用户mid
type AuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 冷启动id，算法uuid，重新起启会变
	Guid string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	// 连接id，算法uuid，重连会变
	ConnId string `protobuf:"bytes,2,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
	// 最后收到的消息id，用于过虑重连后获取未读的消息
	LastMsgId int64 `protobuf:"varint,3,opt,name=last_msg_id,json=lastMsgId,proto3" json:"last_msg_id,omitempty"`
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReq.ProtoReflect.Descriptor instead.
func (*AuthReq) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_broadcast_proto_rawDescGZIP(), []int{0}
}

func (x *AuthReq) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *AuthReq) GetConnId() string {
	if x != nil {
		return x.ConnId
	}
	return ""
}

func (x *AuthReq) GetLastMsgId() int64 {
	if x != nil {
		return x.LastMsgId
	}
	return 0
}

// 鉴权返回
type AuthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthResp) Reset() {
	*x = AuthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResp) ProtoMessage() {}

func (x *AuthResp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResp.ProtoReflect.Descriptor instead.
func (*AuthResp) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_broadcast_proto_rawDescGZIP(), []int{1}
}

// target_path:
//
//	"/" Service-Name "/" {method name} 参考 gRPC Request Path
type BroadcastFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 请求消息信息
	Options *FrameOption `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	// 业务target_path
	TargetPath string `protobuf:"bytes,2,opt,name=target_path,json=targetPath,proto3" json:"target_path,omitempty"`
	// 业务pb内容
	Body *any1.Any `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *BroadcastFrame) Reset() {
	*x = BroadcastFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastFrame) ProtoMessage() {}

func (x *BroadcastFrame) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastFrame.ProtoReflect.Descriptor instead.
func (*BroadcastFrame) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_broadcast_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastFrame) GetOptions() *FrameOption {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *BroadcastFrame) GetTargetPath() string {
	if x != nil {
		return x.TargetPath
	}
	return ""
}

func (x *BroadcastFrame) GetBody() *any1.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

// message_id:
//
//	client: 本次连接唯一的消息id，可用于回执
//	server: 唯一消息id，可用于上报或者回执
//
// sequence:
//
//	client: 客户端应该每次请求时frame seq++，会返回对应的对称req/resp
//	server: 服务端下行消息，只会返回默认值：0
type FrameOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息id
	MessageId int64 `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// frame序号
	Sequence int64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// 是否进行消息回执(发出MessageAckReq)
	// downstream 上只有服务端设置为true，客户端响应
	// upstream   上只有客户端设置为true，服务端响应
	// 响应帧禁止设置is_ack，协议上禁止循环
	// 通常只有业务帧才可能设置is_ack, 因为协议栈(例如心跳、鉴权)另有响应约定
	IsAck bool `protobuf:"varint,3,opt,name=is_ack,json=isAck,proto3" json:"is_ack,omitempty"`
	// 业务状态码
	Status *rpc.Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// 业务ack来源, 仅downstream时候由服务端填写.
	AckOrigin string `protobuf:"bytes,5,opt,name=ack_origin,json=ackOrigin,proto3" json:"ack_origin,omitempty"`
	Timestamp int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *FrameOption) Reset() {
	*x = FrameOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameOption) ProtoMessage() {}

func (x *FrameOption) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameOption.ProtoReflect.Descriptor instead.
func (*FrameOption) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_broadcast_proto_rawDescGZIP(), []int{3}
}

func (x *FrameOption) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *FrameOption) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *FrameOption) GetIsAck() bool {
	if x != nil {
		return x.IsAck
	}
	return false
}

func (x *FrameOption) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *FrameOption) GetAckOrigin() string {
	if x != nil {
		return x.AckOrigin
	}
	return ""
}

func (x *FrameOption) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// 心跳请求
type HeartbeatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatReq) Reset() {
	*x = HeartbeatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatReq) ProtoMessage() {}

func (x *HeartbeatReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatReq.ProtoReflect.Descriptor instead.
func (*HeartbeatReq) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_broadcast_proto_rawDescGZIP(), []int{4}
}

// 心跳返回
type HeartbeatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatResp) Reset() {
	*x = HeartbeatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResp) ProtoMessage() {}

func (x *HeartbeatResp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResp.ProtoReflect.Descriptor instead.
func (*HeartbeatResp) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_broadcast_proto_rawDescGZIP(), []int{5}
}

// 消息回执
type MessageAckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息id
	AckId int64 `protobuf:"varint,1,opt,name=ack_id,json=ackId,proto3" json:"ack_id,omitempty"`
	// ack来源，由业务指定用于埋点跟踪
	AckOrigin string `protobuf:"bytes,2,opt,name=ack_origin,json=ackOrigin,proto3" json:"ack_origin,omitempty"`
	// 消息对应的target_path，方便业务区分和监控统计
	TargetPath string `protobuf:"bytes,3,opt,name=target_path,json=targetPath,proto3" json:"target_path,omitempty"`
}

func (x *MessageAckReq) Reset() {
	*x = MessageAckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAckReq) ProtoMessage() {}

func (x *MessageAckReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAckReq.ProtoReflect.Descriptor instead.
func (*MessageAckReq) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_broadcast_proto_rawDescGZIP(), []int{6}
}

func (x *MessageAckReq) GetAckId() int64 {
	if x != nil {
		return x.AckId
	}
	return 0
}

func (x *MessageAckReq) GetAckOrigin() string {
	if x != nil {
		return x.AckOrigin
	}
	return ""
}

func (x *MessageAckReq) GetTargetPath() string {
	if x != nil {
		return x.TargetPath
	}
	return ""
}

// target_path
type TargetPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 需要订阅的target_paths
	TargetPaths []string `protobuf:"bytes,1,rep,name=target_paths,json=targetPaths,proto3" json:"target_paths,omitempty"`
}

func (x *TargetPath) Reset() {
	*x = TargetPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetPath) ProtoMessage() {}

func (x *TargetPath) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_broadcast_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetPath.ProtoReflect.Descriptor instead.
func (*TargetPath) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_broadcast_proto_rawDescGZIP(), []int{7}
}

func (x *TargetPath) GetTargetPaths() []string {
	if x != nil {
		return x.TargetPaths
	}
	return nil
}

var File_bilibili_broadcast_v1_broadcast_proto protoreflect.FileDescriptor

var file_bilibili_broadcast_v1_broadcast_proto_rawDesc = []byte{
	0x0a, 0x25, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x19,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x56, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x0a, 0x0a, 0x08, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x99, 0x01, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0xca, 0x01, 0x0a, 0x0b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x69, 0x73, 0x5f, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73,
	0x41, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x6b, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x0e,
	0x0a, 0x0c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x22, 0x0f,
	0x0a, 0x0d, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x66, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x6b, 0x5f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x6b,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x2f, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x73, 0x2a, 0x2d, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x32, 0x8a, 0x03, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1e, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x56,
	0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x23, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x24, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x21, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48,
	0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x21, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x24, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x32, 0x73, 0x0a, 0x0f, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x60, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x25, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x1a, 0x25,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75,
	0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_broadcast_v1_broadcast_proto_rawDescOnce sync.Once
	file_bilibili_broadcast_v1_broadcast_proto_rawDescData = file_bilibili_broadcast_v1_broadcast_proto_rawDesc
)

func file_bilibili_broadcast_v1_broadcast_proto_rawDescGZIP() []byte {
	file_bilibili_broadcast_v1_broadcast_proto_rawDescOnce.Do(func() {
		file_bilibili_broadcast_v1_broadcast_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_broadcast_v1_broadcast_proto_rawDescData)
	})
	return file_bilibili_broadcast_v1_broadcast_proto_rawDescData
}

var file_bilibili_broadcast_v1_broadcast_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bilibili_broadcast_v1_broadcast_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_bilibili_broadcast_v1_broadcast_proto_goTypes = []interface{}{
	(Action)(0),            // 0: bilibili.broadcast.v1.Action
	(*AuthReq)(nil),        // 1: bilibili.broadcast.v1.AuthReq
	(*AuthResp)(nil),       // 2: bilibili.broadcast.v1.AuthResp
	(*BroadcastFrame)(nil), // 3: bilibili.broadcast.v1.BroadcastFrame
	(*FrameOption)(nil),    // 4: bilibili.broadcast.v1.FrameOption
	(*HeartbeatReq)(nil),   // 5: bilibili.broadcast.v1.HeartbeatReq
	(*HeartbeatResp)(nil),  // 6: bilibili.broadcast.v1.HeartbeatResp
	(*MessageAckReq)(nil),  // 7: bilibili.broadcast.v1.MessageAckReq
	(*TargetPath)(nil),     // 8: bilibili.broadcast.v1.TargetPath
	(*any1.Any)(nil),       // 9: google.protobuf.Any
	(*rpc.Status)(nil),     // 10: bilibili.rpc.Status
	(*empty.Empty)(nil),    // 11: google.protobuf.Empty
}
var file_bilibili_broadcast_v1_broadcast_proto_depIdxs = []int32{
	4,  // 0: bilibili.broadcast.v1.BroadcastFrame.options:type_name -> bilibili.broadcast.v1.FrameOption
	9,  // 1: bilibili.broadcast.v1.BroadcastFrame.body:type_name -> google.protobuf.Any
	10, // 2: bilibili.broadcast.v1.FrameOption.status:type_name -> bilibili.rpc.Status
	1,  // 3: bilibili.broadcast.v1.Broadcast.Auth:input_type -> bilibili.broadcast.v1.AuthReq
	5,  // 4: bilibili.broadcast.v1.Broadcast.Heartbeat:input_type -> bilibili.broadcast.v1.HeartbeatReq
	8,  // 5: bilibili.broadcast.v1.Broadcast.Subscribe:input_type -> bilibili.broadcast.v1.TargetPath
	8,  // 6: bilibili.broadcast.v1.Broadcast.Unsubscribe:input_type -> bilibili.broadcast.v1.TargetPath
	7,  // 7: bilibili.broadcast.v1.Broadcast.MessageAck:input_type -> bilibili.broadcast.v1.MessageAckReq
	3,  // 8: bilibili.broadcast.v1.BroadcastTunnel.CreateTunnel:input_type -> bilibili.broadcast.v1.BroadcastFrame
	2,  // 9: bilibili.broadcast.v1.Broadcast.Auth:output_type -> bilibili.broadcast.v1.AuthResp
	6,  // 10: bilibili.broadcast.v1.Broadcast.Heartbeat:output_type -> bilibili.broadcast.v1.HeartbeatResp
	11, // 11: bilibili.broadcast.v1.Broadcast.Subscribe:output_type -> google.protobuf.Empty
	11, // 12: bilibili.broadcast.v1.Broadcast.Unsubscribe:output_type -> google.protobuf.Empty
	11, // 13: bilibili.broadcast.v1.Broadcast.MessageAck:output_type -> google.protobuf.Empty
	3,  // 14: bilibili.broadcast.v1.BroadcastTunnel.CreateTunnel:output_type -> bilibili.broadcast.v1.BroadcastFrame
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_bilibili_broadcast_v1_broadcast_proto_init() }
func file_bilibili_broadcast_v1_broadcast_proto_init() {
	if File_bilibili_broadcast_v1_broadcast_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_broadcast_v1_broadcast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReq); i {
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
		file_bilibili_broadcast_v1_broadcast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResp); i {
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
		file_bilibili_broadcast_v1_broadcast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastFrame); i {
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
		file_bilibili_broadcast_v1_broadcast_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrameOption); i {
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
		file_bilibili_broadcast_v1_broadcast_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatReq); i {
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
		file_bilibili_broadcast_v1_broadcast_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResp); i {
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
		file_bilibili_broadcast_v1_broadcast_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAckReq); i {
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
		file_bilibili_broadcast_v1_broadcast_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetPath); i {
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
			RawDescriptor: file_bilibili_broadcast_v1_broadcast_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_bilibili_broadcast_v1_broadcast_proto_goTypes,
		DependencyIndexes: file_bilibili_broadcast_v1_broadcast_proto_depIdxs,
		EnumInfos:         file_bilibili_broadcast_v1_broadcast_proto_enumTypes,
		MessageInfos:      file_bilibili_broadcast_v1_broadcast_proto_msgTypes,
	}.Build()
	File_bilibili_broadcast_v1_broadcast_proto = out.File
	file_bilibili_broadcast_v1_broadcast_proto_rawDesc = nil
	file_bilibili_broadcast_v1_broadcast_proto_goTypes = nil
	file_bilibili_broadcast_v1_broadcast_proto_depIdxs = nil
}
