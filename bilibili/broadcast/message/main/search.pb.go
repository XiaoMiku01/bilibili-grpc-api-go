// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/broadcast/message/main/search.proto

package main

import (
	v2 "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/app/dynamic/v2"
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

type Bubble struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paragraphs []*v2.Paragraph `protobuf:"bytes,1,rep,name=paragraphs,proto3" json:"paragraphs,omitempty"`
}

func (x *Bubble) Reset() {
	*x = Bubble{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_main_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bubble) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bubble) ProtoMessage() {}

func (x *Bubble) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_main_search_proto_msgTypes[0]
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
	return file_bilibili_broadcast_message_main_search_proto_rawDescGZIP(), []int{0}
}

func (x *Bubble) GetParagraphs() []*v2.Paragraph {
	if x != nil {
		return x.Paragraphs
	}
	return nil
}

type ChatResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Bubble []*Bubble `protobuf:"bytes,3,rep,name=bubble,proto3" json:"bubble,omitempty"`
	RewriteWord string `protobuf:"bytes,4,opt,name=rewrite_word,json=rewriteWord,proto3" json:"rewrite_word,omitempty"`
	Title string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ChatResult) Reset() {
	*x = ChatResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_message_main_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResult) ProtoMessage() {}

func (x *ChatResult) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_message_main_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResult.ProtoReflect.Descriptor instead.
func (*ChatResult) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_message_main_search_proto_rawDescGZIP(), []int{1}
}

func (x *ChatResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ChatResult) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ChatResult) GetBubble() []*Bubble {
	if x != nil {
		return x.Bubble
	}
	return nil
}

func (x *ChatResult) GetRewriteWord() string {
	if x != nil {
		return x.RewriteWord
	}
	return ""
}

func (x *ChatResult) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_bilibili_broadcast_message_main_search_proto protoreflect.FileDescriptor

var file_bilibili_broadcast_message_main_search_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x06, 0x42, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x12, 0x42, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x73, 0x22, 0xb9, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x06, 0x62, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x52, 0x06, 0x62, 0x75,
	0x62, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x32, 0x61, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x57, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x75, 0x73, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x2b, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x30, 0x01,
	0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58,
	0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_broadcast_message_main_search_proto_rawDescOnce sync.Once
	file_bilibili_broadcast_message_main_search_proto_rawDescData = file_bilibili_broadcast_message_main_search_proto_rawDesc
)

func file_bilibili_broadcast_message_main_search_proto_rawDescGZIP() []byte {
	file_bilibili_broadcast_message_main_search_proto_rawDescOnce.Do(func() {
		file_bilibili_broadcast_message_main_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_broadcast_message_main_search_proto_rawDescData)
	})
	return file_bilibili_broadcast_message_main_search_proto_rawDescData
}

var file_bilibili_broadcast_message_main_search_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bilibili_broadcast_message_main_search_proto_goTypes = []interface{}{
	(*Bubble)(nil),       // 0: bilibili.broadcast.message.main.Bubble
	(*ChatResult)(nil),   // 1: bilibili.broadcast.message.main.ChatResult
	(*v2.Paragraph)(nil), // 2: bilibili.app.dynamic.v2.Paragraph
	(*empty.Empty)(nil),  // 3: google.protobuf.Empty
}
var file_bilibili_broadcast_message_main_search_proto_depIdxs = []int32{
	2, // 0: bilibili.broadcast.message.main.Bubble.paragraphs:type_name -> bilibili.app.dynamic.v2.Paragraph
	0, // 1: bilibili.broadcast.message.main.ChatResult.bubble:type_name -> bilibili.broadcast.message.main.Bubble
	3, // 2: bilibili.broadcast.message.main.Search.ChatResultPush:input_type -> google.protobuf.Empty
	1, // 3: bilibili.broadcast.message.main.Search.ChatResultPush:output_type -> bilibili.broadcast.message.main.ChatResult
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bilibili_broadcast_message_main_search_proto_init() }
func file_bilibili_broadcast_message_main_search_proto_init() {
	if File_bilibili_broadcast_message_main_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_broadcast_message_main_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_bilibili_broadcast_message_main_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResult); i {
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
			RawDescriptor: file_bilibili_broadcast_message_main_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_broadcast_message_main_search_proto_goTypes,
		DependencyIndexes: file_bilibili_broadcast_message_main_search_proto_depIdxs,
		MessageInfos:      file_bilibili_broadcast_message_main_search_proto_msgTypes,
	}.Build()
	File_bilibili_broadcast_message_main_search_proto = out.File
	file_bilibili_broadcast_message_main_search_proto_rawDesc = nil
	file_bilibili_broadcast_message_main_search_proto_goTypes = nil
	file_bilibili_broadcast_message_main_search_proto_depIdxs = nil
}
