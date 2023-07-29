// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/metadata/metadata.proto

package metadata

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

// 请求元数据
// gRPC头部:x-bili-metadata-bin
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 登录 access_key
	AccessKey string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// 包类型, 如 `android`
	MobiApp string `protobuf:"bytes,2,opt,name=mobi_app,json=mobiApp,proto3" json:"mobi_app,omitempty"`
	// 运行设备, 留空即可
	Device string `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	// 构建id, 如 `7380300`
	Build int32 `protobuf:"varint,4,opt,name=build,proto3" json:"build,omitempty"`
	// APP分发渠道, 如 `master`
	Channel string `protobuf:"bytes,5,opt,name=channel,proto3" json:"channel,omitempty"`
	// 设备唯一标识
	Buvid string `protobuf:"bytes,6,opt,name=buvid,proto3" json:"buvid,omitempty"`
	// 平台类型, 如 `android`
	Platform string `protobuf:"bytes,7,opt,name=platform,proto3" json:"platform,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_metadata_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_metadata_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_bilibili_metadata_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Metadata) GetMobiApp() string {
	if x != nil {
		return x.MobiApp
	}
	return ""
}

func (x *Metadata) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *Metadata) GetBuild() int32 {
	if x != nil {
		return x.Build
	}
	return 0
}

func (x *Metadata) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Metadata) GetBuvid() string {
	if x != nil {
		return x.Buvid
	}
	return ""
}

func (x *Metadata) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

var File_bilibili_metadata_metadata_proto protoreflect.FileDescriptor

var file_bilibili_metadata_metadata_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbe, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x62, 0x69, 0x5f, 0x61, 0x70, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x62, 0x69, 0x41, 0x70, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x76, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x76, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_metadata_metadata_proto_rawDescOnce sync.Once
	file_bilibili_metadata_metadata_proto_rawDescData = file_bilibili_metadata_metadata_proto_rawDesc
)

func file_bilibili_metadata_metadata_proto_rawDescGZIP() []byte {
	file_bilibili_metadata_metadata_proto_rawDescOnce.Do(func() {
		file_bilibili_metadata_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_metadata_metadata_proto_rawDescData)
	})
	return file_bilibili_metadata_metadata_proto_rawDescData
}

var file_bilibili_metadata_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bilibili_metadata_metadata_proto_goTypes = []interface{}{
	(*Metadata)(nil), // 0: bilibili.metadata.Metadata
}
var file_bilibili_metadata_metadata_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bilibili_metadata_metadata_proto_init() }
func file_bilibili_metadata_metadata_proto_init() {
	if File_bilibili_metadata_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_metadata_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
			RawDescriptor: file_bilibili_metadata_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_metadata_metadata_proto_goTypes,
		DependencyIndexes: file_bilibili_metadata_metadata_proto_depIdxs,
		MessageInfos:      file_bilibili_metadata_metadata_proto_msgTypes,
	}.Build()
	File_bilibili_metadata_metadata_proto = out.File
	file_bilibili_metadata_metadata_proto_rawDesc = nil
	file_bilibili_metadata_metadata_proto_goTypes = nil
	file_bilibili_metadata_metadata_proto_depIdxs = nil
}
