// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/polymer/demo/demo.proto

package demo

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

type HelloWorldReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *HelloWorldReq) Reset() {
	*x = HelloWorldReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_demo_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldReq) ProtoMessage() {}

func (x *HelloWorldReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_demo_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldReq.ProtoReflect.Descriptor instead.
func (*HelloWorldReq) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_demo_demo_proto_rawDescGZIP(), []int{0}
}

func (x *HelloWorldReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type HelloWorldResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *HelloWorldResp) Reset() {
	*x = HelloWorldResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_demo_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldResp) ProtoMessage() {}

func (x *HelloWorldResp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_demo_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldResp.ProtoReflect.Descriptor instead.
func (*HelloWorldResp) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_demo_demo_proto_rawDescGZIP(), []int{1}
}

func (x *HelloWorldResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_bilibili_polymer_demo_demo_proto protoreflect.FileDescriptor

var file_bilibili_polymer_demo_demo_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x6d,
	0x65, 0x72, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c,
	0x79, 0x6d, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x22, 0x29, 0x0a, 0x0d, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x0e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b,
	0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_polymer_demo_demo_proto_rawDescOnce sync.Once
	file_bilibili_polymer_demo_demo_proto_rawDescData = file_bilibili_polymer_demo_demo_proto_rawDesc
)

func file_bilibili_polymer_demo_demo_proto_rawDescGZIP() []byte {
	file_bilibili_polymer_demo_demo_proto_rawDescOnce.Do(func() {
		file_bilibili_polymer_demo_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_polymer_demo_demo_proto_rawDescData)
	})
	return file_bilibili_polymer_demo_demo_proto_rawDescData
}

var file_bilibili_polymer_demo_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bilibili_polymer_demo_demo_proto_goTypes = []interface{}{
	(*HelloWorldReq)(nil),  // 0: bilibili.polymer.demo.HelloWorldReq
	(*HelloWorldResp)(nil), // 1: bilibili.polymer.demo.HelloWorldResp
}
var file_bilibili_polymer_demo_demo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bilibili_polymer_demo_demo_proto_init() }
func file_bilibili_polymer_demo_demo_proto_init() {
	if File_bilibili_polymer_demo_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_polymer_demo_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldReq); i {
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
		file_bilibili_polymer_demo_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldResp); i {
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
			RawDescriptor: file_bilibili_polymer_demo_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_polymer_demo_demo_proto_goTypes,
		DependencyIndexes: file_bilibili_polymer_demo_demo_proto_depIdxs,
		MessageInfos:      file_bilibili_polymer_demo_demo_proto_msgTypes,
	}.Build()
	File_bilibili_polymer_demo_demo_proto = out.File
	file_bilibili_polymer_demo_demo_proto_rawDesc = nil
	file_bilibili_polymer_demo_demo_proto_goTypes = nil
	file_bilibili_polymer_demo_demo_proto_depIdxs = nil
}
