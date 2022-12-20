// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/broadcast/v1/room.proto

package v1

import (
	rpc "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/rpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RoomErrorEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RoomErrorEvent) Reset() {
	*x = RoomErrorEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomErrorEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomErrorEvent) ProtoMessage() {}

func (x *RoomErrorEvent) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomErrorEvent.ProtoReflect.Descriptor instead.
func (*RoomErrorEvent) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_room_proto_rawDescGZIP(), []int{0}
}

func (x *RoomErrorEvent) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type RoomJoinEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoomJoinEvent) Reset() {
	*x = RoomJoinEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomJoinEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomJoinEvent) ProtoMessage() {}

func (x *RoomJoinEvent) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomJoinEvent.ProtoReflect.Descriptor instead.
func (*RoomJoinEvent) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_room_proto_rawDescGZIP(), []int{1}
}

type RoomLeaveEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoomLeaveEvent) Reset() {
	*x = RoomLeaveEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomLeaveEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomLeaveEvent) ProtoMessage() {}

func (x *RoomLeaveEvent) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomLeaveEvent.ProtoReflect.Descriptor instead.
func (*RoomLeaveEvent) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_room_proto_rawDescGZIP(), []int{2}
}

type RoomMessageEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetPath string `protobuf:"bytes,1,opt,name=target_path,json=targetPath,proto3" json:"target_path,omitempty"`
	Body *anypb.Any `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *RoomMessageEvent) Reset() {
	*x = RoomMessageEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomMessageEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomMessageEvent) ProtoMessage() {}

func (x *RoomMessageEvent) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomMessageEvent.ProtoReflect.Descriptor instead.
func (*RoomMessageEvent) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_room_proto_rawDescGZIP(), []int{3}
}

func (x *RoomMessageEvent) GetTargetPath() string {
	if x != nil {
		return x.TargetPath
	}
	return ""
}

func (x *RoomMessageEvent) GetBody() *anypb.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

type RoomOnlineEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Online int32 `protobuf:"varint,1,opt,name=online,proto3" json:"online,omitempty"`
	AllOnline int32 `protobuf:"varint,2,opt,name=all_online,json=allOnline,proto3" json:"all_online,omitempty"`
}

func (x *RoomOnlineEvent) Reset() {
	*x = RoomOnlineEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomOnlineEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomOnlineEvent) ProtoMessage() {}

func (x *RoomOnlineEvent) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomOnlineEvent.ProtoReflect.Descriptor instead.
func (*RoomOnlineEvent) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_room_proto_rawDescGZIP(), []int{4}
}

func (x *RoomOnlineEvent) GetOnline() int32 {
	if x != nil {
		return x.Online
	}
	return 0
}

func (x *RoomOnlineEvent) GetAllOnline() int32 {
	if x != nil {
		return x.AllOnline
	}
	return 0
}

type RoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// {type}://{room_id}
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Event:
	//
	//	*RoomReq_Join
	//	*RoomReq_Leave
	//	*RoomReq_Online
	//	*RoomReq_Msg
	Event isRoomReq_Event `protobuf_oneof:"event"`
}

func (x *RoomReq) Reset() {
	*x = RoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomReq) ProtoMessage() {}

func (x *RoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomReq.ProtoReflect.Descriptor instead.
func (*RoomReq) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_room_proto_rawDescGZIP(), []int{5}
}

func (x *RoomReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *RoomReq) GetEvent() isRoomReq_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *RoomReq) GetJoin() *RoomJoinEvent {
	if x, ok := x.GetEvent().(*RoomReq_Join); ok {
		return x.Join
	}
	return nil
}

func (x *RoomReq) GetLeave() *RoomLeaveEvent {
	if x, ok := x.GetEvent().(*RoomReq_Leave); ok {
		return x.Leave
	}
	return nil
}

func (x *RoomReq) GetOnline() *RoomOnlineEvent {
	if x, ok := x.GetEvent().(*RoomReq_Online); ok {
		return x.Online
	}
	return nil
}

func (x *RoomReq) GetMsg() *RoomMessageEvent {
	if x, ok := x.GetEvent().(*RoomReq_Msg); ok {
		return x.Msg
	}
	return nil
}

type isRoomReq_Event interface {
	isRoomReq_Event()
}

type RoomReq_Join struct {
	Join *RoomJoinEvent `protobuf:"bytes,2,opt,name=join,proto3,oneof"`
}

type RoomReq_Leave struct {
	Leave *RoomLeaveEvent `protobuf:"bytes,3,opt,name=leave,proto3,oneof"`
}

type RoomReq_Online struct {
	Online *RoomOnlineEvent `protobuf:"bytes,4,opt,name=online,proto3,oneof"`
}

type RoomReq_Msg struct {
	Msg *RoomMessageEvent `protobuf:"bytes,5,opt,name=msg,proto3,oneof"`
}

func (*RoomReq_Join) isRoomReq_Event() {}

func (*RoomReq_Leave) isRoomReq_Event() {}

func (*RoomReq_Online) isRoomReq_Event() {}

func (*RoomReq_Msg) isRoomReq_Event() {}

type RoomResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// {type}://{room_id}
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Event:
	//
	//	*RoomResp_Join
	//	*RoomResp_Leave
	//	*RoomResp_Online
	//	*RoomResp_Msg
	//	*RoomResp_Err
	Event isRoomResp_Event `protobuf_oneof:"event"`
}

func (x *RoomResp) Reset() {
	*x = RoomResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomResp) ProtoMessage() {}

func (x *RoomResp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_room_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomResp.ProtoReflect.Descriptor instead.
func (*RoomResp) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_room_proto_rawDescGZIP(), []int{6}
}

func (x *RoomResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *RoomResp) GetEvent() isRoomResp_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *RoomResp) GetJoin() *RoomJoinEvent {
	if x, ok := x.GetEvent().(*RoomResp_Join); ok {
		return x.Join
	}
	return nil
}

func (x *RoomResp) GetLeave() *RoomLeaveEvent {
	if x, ok := x.GetEvent().(*RoomResp_Leave); ok {
		return x.Leave
	}
	return nil
}

func (x *RoomResp) GetOnline() *RoomOnlineEvent {
	if x, ok := x.GetEvent().(*RoomResp_Online); ok {
		return x.Online
	}
	return nil
}

func (x *RoomResp) GetMsg() *RoomMessageEvent {
	if x, ok := x.GetEvent().(*RoomResp_Msg); ok {
		return x.Msg
	}
	return nil
}

func (x *RoomResp) GetErr() *RoomErrorEvent {
	if x, ok := x.GetEvent().(*RoomResp_Err); ok {
		return x.Err
	}
	return nil
}

type isRoomResp_Event interface {
	isRoomResp_Event()
}

type RoomResp_Join struct {
	Join *RoomJoinEvent `protobuf:"bytes,2,opt,name=join,proto3,oneof"`
}

type RoomResp_Leave struct {
	Leave *RoomLeaveEvent `protobuf:"bytes,3,opt,name=leave,proto3,oneof"`
}

type RoomResp_Online struct {
	Online *RoomOnlineEvent `protobuf:"bytes,4,opt,name=online,proto3,oneof"`
}

type RoomResp_Msg struct {
	Msg *RoomMessageEvent `protobuf:"bytes,5,opt,name=msg,proto3,oneof"`
}

type RoomResp_Err struct {
	Err *RoomErrorEvent `protobuf:"bytes,6,opt,name=err,proto3,oneof"`
}

func (*RoomResp_Join) isRoomResp_Event() {}

func (*RoomResp_Leave) isRoomResp_Event() {}

func (*RoomResp_Online) isRoomResp_Event() {}

func (*RoomResp_Msg) isRoomResp_Event() {}

func (*RoomResp_Err) isRoomResp_Event() {}

var File_bilibili_broadcast_v1_room_proto protoreflect.FileDescriptor

var file_bilibili_broadcast_v1_room_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3e, 0x0a, 0x0e, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x0f, 0x0a, 0x0d, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x10, 0x0a, 0x0e, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x22, 0x5d, 0x0a, 0x10, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x48, 0x0a, 0x0f, 0x52, 0x6f, 0x6f, 0x6d, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x6c, 0x6c, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x61, 0x6c, 0x6c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x9c, 0x02, 0x0a, 0x07,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x04, 0x6a,
	0x6f, 0x69, 0x6e, 0x12, 0x3d, 0x0a, 0x05, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x65, 0x61,
	0x76, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3b, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xd8, 0x02, 0x0a, 0x08, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x04, 0x6a,
	0x6f, 0x69, 0x6e, 0x12, 0x3d, 0x0a, 0x05, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x65, 0x61,
	0x76, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3b, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x39, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x03, 0x65, 0x72, 0x72, 0x42, 0x07, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0x5d, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x4c, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x1e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x1a,
	0x1f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_broadcast_v1_room_proto_rawDescOnce sync.Once
	file_bilibili_broadcast_v1_room_proto_rawDescData = file_bilibili_broadcast_v1_room_proto_rawDesc
)

func file_bilibili_broadcast_v1_room_proto_rawDescGZIP() []byte {
	file_bilibili_broadcast_v1_room_proto_rawDescOnce.Do(func() {
		file_bilibili_broadcast_v1_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_broadcast_v1_room_proto_rawDescData)
	})
	return file_bilibili_broadcast_v1_room_proto_rawDescData
}

var file_bilibili_broadcast_v1_room_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bilibili_broadcast_v1_room_proto_goTypes = []interface{}{
	(*RoomErrorEvent)(nil),   // 0: bilibili.broadcast.v1.RoomErrorEvent
	(*RoomJoinEvent)(nil),    // 1: bilibili.broadcast.v1.RoomJoinEvent
	(*RoomLeaveEvent)(nil),   // 2: bilibili.broadcast.v1.RoomLeaveEvent
	(*RoomMessageEvent)(nil), // 3: bilibili.broadcast.v1.RoomMessageEvent
	(*RoomOnlineEvent)(nil),  // 4: bilibili.broadcast.v1.RoomOnlineEvent
	(*RoomReq)(nil),          // 5: bilibili.broadcast.v1.RoomReq
	(*RoomResp)(nil),         // 6: bilibili.broadcast.v1.RoomResp
	(*rpc.Status)(nil),       // 7: bilibili.rpc.Status
	(*anypb.Any)(nil),        // 8: google.protobuf.Any
}
var file_bilibili_broadcast_v1_room_proto_depIdxs = []int32{
	7,  // 0: bilibili.broadcast.v1.RoomErrorEvent.status:type_name -> bilibili.rpc.Status
	8,  // 1: bilibili.broadcast.v1.RoomMessageEvent.body:type_name -> google.protobuf.Any
	1,  // 2: bilibili.broadcast.v1.RoomReq.join:type_name -> bilibili.broadcast.v1.RoomJoinEvent
	2,  // 3: bilibili.broadcast.v1.RoomReq.leave:type_name -> bilibili.broadcast.v1.RoomLeaveEvent
	4,  // 4: bilibili.broadcast.v1.RoomReq.online:type_name -> bilibili.broadcast.v1.RoomOnlineEvent
	3,  // 5: bilibili.broadcast.v1.RoomReq.msg:type_name -> bilibili.broadcast.v1.RoomMessageEvent
	1,  // 6: bilibili.broadcast.v1.RoomResp.join:type_name -> bilibili.broadcast.v1.RoomJoinEvent
	2,  // 7: bilibili.broadcast.v1.RoomResp.leave:type_name -> bilibili.broadcast.v1.RoomLeaveEvent
	4,  // 8: bilibili.broadcast.v1.RoomResp.online:type_name -> bilibili.broadcast.v1.RoomOnlineEvent
	3,  // 9: bilibili.broadcast.v1.RoomResp.msg:type_name -> bilibili.broadcast.v1.RoomMessageEvent
	0,  // 10: bilibili.broadcast.v1.RoomResp.err:type_name -> bilibili.broadcast.v1.RoomErrorEvent
	5,  // 11: bilibili.broadcast.v1.BroadcastRoom.Enter:input_type -> bilibili.broadcast.v1.RoomReq
	6,  // 12: bilibili.broadcast.v1.BroadcastRoom.Enter:output_type -> bilibili.broadcast.v1.RoomResp
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_bilibili_broadcast_v1_room_proto_init() }
func file_bilibili_broadcast_v1_room_proto_init() {
	if File_bilibili_broadcast_v1_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_broadcast_v1_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomErrorEvent); i {
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
		file_bilibili_broadcast_v1_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomJoinEvent); i {
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
		file_bilibili_broadcast_v1_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomLeaveEvent); i {
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
		file_bilibili_broadcast_v1_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomMessageEvent); i {
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
		file_bilibili_broadcast_v1_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomOnlineEvent); i {
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
		file_bilibili_broadcast_v1_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomReq); i {
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
		file_bilibili_broadcast_v1_room_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomResp); i {
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
	file_bilibili_broadcast_v1_room_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*RoomReq_Join)(nil),
		(*RoomReq_Leave)(nil),
		(*RoomReq_Online)(nil),
		(*RoomReq_Msg)(nil),
	}
	file_bilibili_broadcast_v1_room_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*RoomResp_Join)(nil),
		(*RoomResp_Leave)(nil),
		(*RoomResp_Online)(nil),
		(*RoomResp_Msg)(nil),
		(*RoomResp_Err)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bilibili_broadcast_v1_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_broadcast_v1_room_proto_goTypes,
		DependencyIndexes: file_bilibili_broadcast_v1_room_proto_depIdxs,
		MessageInfos:      file_bilibili_broadcast_v1_room_proto_msgTypes,
	}.Build()
	File_bilibili_broadcast_v1_room_proto = out.File
	file_bilibili_broadcast_v1_room_proto_rawDesc = nil
	file_bilibili_broadcast_v1_room_proto_goTypes = nil
	file_bilibili_broadcast_v1_room_proto_depIdxs = nil
}
