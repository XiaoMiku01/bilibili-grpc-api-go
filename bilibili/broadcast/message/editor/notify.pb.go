// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/broadcast/message/editor/notify.proto

package editor

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

type Notify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息唯一标示
	MsgId int64 `protobuf:"varint,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	// 消息类型
	MsgType int32 `protobuf:"varint,2,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	// 接收方uid
	ReceiverUid int64 `protobuf:"varint,3,opt,name=receiver_uid,json=receiverUid,proto3" json:"receiver_uid,omitempty"`
	// 接收方类型
	ReceiverType int32 `protobuf:"varint,4,opt,name=receiver_type,json=receiverType,proto3" json:"receiver_type,omitempty"`
	// 故事的版本
	StoryVersion int64 `protobuf:"varint,5,opt,name=story_version,json=storyVersion,proto3" json:"story_version,omitempty"`
	// 操作结果的hash值
	OpHash int64 `protobuf:"varint,6,opt,name=op_hash,json=opHash,proto3" json:"op_hash,omitempty"`
	// 操作产生用户的uid
	OpSender int64 `protobuf:"varint,7,opt,name=op_sender,json=opSender,proto3" json:"op_sender,omitempty"`
	// patch内容
	OpContent string `protobuf:"bytes,8,opt,name=op_content,json=opContent,proto3" json:"op_content,omitempty"`
}

func (x *Notify) Reset() {
	*x = Notify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_editor_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notify) ProtoMessage() {}

func (x *Notify) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_editor_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notify.ProtoReflect.Descriptor instead.
func (*Notify) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_editor_notify_proto_rawDescGZIP(), []int{0}
}

func (x *Notify) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *Notify) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *Notify) GetReceiverUid() int64 {
	if x != nil {
		return x.ReceiverUid
	}
	return 0
}

func (x *Notify) GetReceiverType() int32 {
	if x != nil {
		return x.ReceiverType
	}
	return 0
}

func (x *Notify) GetStoryVersion() int64 {
	if x != nil {
		return x.StoryVersion
	}
	return 0
}

func (x *Notify) GetOpHash() int64 {
	if x != nil {
		return x.OpHash
	}
	return 0
}

func (x *Notify) GetOpSender() int64 {
	if x != nil {
		return x.OpSender
	}
	return 0
}

func (x *Notify) GetOpContent() string {
	if x != nil {
		return x.OpContent
	}
	return ""
}

var File_bilibili_broadcast_message_editor_notify_proto protoreflect.FileDescriptor

var file_bilibili_broadcast_message_editor_notify_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x65, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x65, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfc, 0x01, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x6d,
	0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x73, 0x67,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x55, 0x69, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x70, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x70, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32,
	0x69, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x56, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x30, 0x01, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b,
	0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_bilibili_broadcast_message_editor_notify_proto_rawDescOnce sync.Once
	file_bilibili_broadcast_message_editor_notify_proto_rawDescData = file_bilibili_broadcast_message_editor_notify_proto_rawDesc
)

func file_bilibili_broadcast_message_editor_notify_proto_rawDescGZIP() []byte {
	file_bilibili_broadcast_message_editor_notify_proto_rawDescOnce.Do(func() {
		file_bilibili_broadcast_message_editor_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_broadcast_message_editor_notify_proto_rawDescData)
	})
	return file_bilibili_broadcast_message_editor_notify_proto_rawDescData
}

var file_bilibili_broadcast_message_editor_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bilibili_broadcast_message_editor_notify_proto_goTypes = []interface{}{
	(*Notify)(nil),      // 0: bilibili.broadcast.message.editor.Notify
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
}
var file_bilibili_broadcast_message_editor_notify_proto_depIdxs = []int32{
	1, // 0: bilibili.broadcast.message.editor.OperationNotify.OperationNotify:input_type -> google.protobuf.Empty
	0, // 1: bilibili.broadcast.message.editor.OperationNotify.OperationNotify:output_type -> bilibili.broadcast.message.editor.Notify
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bilibili_broadcast_message_editor_notify_proto_init() }
func file_bilibili_broadcast_message_editor_notify_proto_init() {
	if File_bilibili_broadcast_message_editor_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_broadcast_message_editor_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notify); i {
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
			RawDescriptor: file_bilibili_broadcast_message_editor_notify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_broadcast_message_editor_notify_proto_goTypes,
		DependencyIndexes: file_bilibili_broadcast_message_editor_notify_proto_depIdxs,
		MessageInfos:      file_bilibili_broadcast_message_editor_notify_proto_msgTypes,
	}.Build()
	File_bilibili_broadcast_message_editor_notify_proto = out.File
	file_bilibili_broadcast_message_editor_notify_proto_rawDesc = nil
	file_bilibili_broadcast_message_editor_notify_proto_goTypes = nil
	file_bilibili_broadcast_message_editor_notify_proto_depIdxs = nil
}
