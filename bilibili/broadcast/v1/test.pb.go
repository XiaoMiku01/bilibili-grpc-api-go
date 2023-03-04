// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: bilibili/broadcast/v1/test.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *AddParams) Reset() {
	*x = AddParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddParams) ProtoMessage() {}

func (x *AddParams) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddParams.ProtoReflect.Descriptor instead.
func (*AddParams) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_test_proto_rawDescGZIP(), []int{0}
}

func (x *AddParams) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *AddParams) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type AddResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	R int32 `protobuf:"varint,1,opt,name=r,proto3" json:"r,omitempty"`
}

func (x *AddResult) Reset() {
	*x = AddResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResult) ProtoMessage() {}

func (x *AddResult) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResult.ProtoReflect.Descriptor instead.
func (*AddResult) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_test_proto_rawDescGZIP(), []int{1}
}

func (x *AddResult) GetR() int32 {
	if x != nil {
		return x.R
	}
	return 0
}

type TestResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务id
	Taskid int64 `protobuf:"varint,1,opt,name=taskid,proto3" json:"taskid,omitempty"`
	// 时间戳
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 消息
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// 扩展
	Extra *anypb.Any `protobuf:"bytes,4,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *TestResp) Reset() {
	*x = TestResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v1_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResp) ProtoMessage() {}

func (x *TestResp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v1_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResp.ProtoReflect.Descriptor instead.
func (*TestResp) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v1_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestResp) GetTaskid() int64 {
	if x != nil {
		return x.Taskid
	}
	return 0
}

func (x *TestResp) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TestResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TestResp) GetExtra() *anypb.Any {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_bilibili_broadcast_v1_test_proto protoreflect.FileDescriptor

var file_bilibili_broadcast_v1_test_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x27, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0c,
	0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x22, 0x19, 0x0a, 0x09, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x32, 0x53,
	0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x54,
	0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x30, 0x01, 0x32, 0x49, 0x0a, 0x05, 0x54, 0x65, 0x73, 0x74, 0x32, 0x12, 0x40, 0x0a, 0x04,
	0x54, 0x65, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x42,
	0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61,
	0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_broadcast_v1_test_proto_rawDescOnce sync.Once
	file_bilibili_broadcast_v1_test_proto_rawDescData = file_bilibili_broadcast_v1_test_proto_rawDesc
)

func file_bilibili_broadcast_v1_test_proto_rawDescGZIP() []byte {
	file_bilibili_broadcast_v1_test_proto_rawDescOnce.Do(func() {
		file_bilibili_broadcast_v1_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_broadcast_v1_test_proto_rawDescData)
	})
	return file_bilibili_broadcast_v1_test_proto_rawDescData
}

var file_bilibili_broadcast_v1_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bilibili_broadcast_v1_test_proto_goTypes = []interface{}{
	(*AddParams)(nil),     // 0: bilibili.broadcast.v1.AddParams
	(*AddResult)(nil),     // 1: bilibili.broadcast.v1.AddResult
	(*TestResp)(nil),      // 2: bilibili.broadcast.v1.TestResp
	(*anypb.Any)(nil),     // 3: google.protobuf.Any
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_bilibili_broadcast_v1_test_proto_depIdxs = []int32{
	3, // 0: bilibili.broadcast.v1.TestResp.extra:type_name -> google.protobuf.Any
	4, // 1: bilibili.broadcast.v1.Test.WatchTestEvent:input_type -> google.protobuf.Empty
	0, // 2: bilibili.broadcast.v1.Test2.Test:input_type -> bilibili.broadcast.v1.AddParams
	2, // 3: bilibili.broadcast.v1.Test.WatchTestEvent:output_type -> bilibili.broadcast.v1.TestResp
	4, // 4: bilibili.broadcast.v1.Test2.Test:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bilibili_broadcast_v1_test_proto_init() }
func file_bilibili_broadcast_v1_test_proto_init() {
	if File_bilibili_broadcast_v1_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_broadcast_v1_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddParams); i {
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
		file_bilibili_broadcast_v1_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResult); i {
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
		file_bilibili_broadcast_v1_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResp); i {
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
			RawDescriptor: file_bilibili_broadcast_v1_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_bilibili_broadcast_v1_test_proto_goTypes,
		DependencyIndexes: file_bilibili_broadcast_v1_test_proto_depIdxs,
		MessageInfos:      file_bilibili_broadcast_v1_test_proto_msgTypes,
	}.Build()
	File_bilibili_broadcast_v1_test_proto = out.File
	file_bilibili_broadcast_v1_test_proto_rawDesc = nil
	file_bilibili_broadcast_v1_test_proto_goTypes = nil
	file_bilibili_broadcast_v1_test_proto_depIdxs = nil
}
