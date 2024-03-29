// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/metadata/network/network.proto

package network

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

// 网络类型
type NetworkType int32

const (
	NetworkType_NT_UNKNOWN NetworkType = 0 // 未知
	NetworkType_WIFI       NetworkType = 1 // WIFI
	NetworkType_CELLULAR   NetworkType = 2 // 蜂窝网络
	NetworkType_OFFLINE    NetworkType = 3 // 未连接
	NetworkType_OTHERNET   NetworkType = 4 // 其他网络
	NetworkType_ETHERNET   NetworkType = 5 // 以太网
)

// Enum value maps for NetworkType.
var (
	NetworkType_name = map[int32]string{
		0: "NT_UNKNOWN",
		1: "WIFI",
		2: "CELLULAR",
		3: "OFFLINE",
		4: "OTHERNET",
		5: "ETHERNET",
	}
	NetworkType_value = map[string]int32{
		"NT_UNKNOWN": 0,
		"WIFI":       1,
		"CELLULAR":   2,
		"OFFLINE":    3,
		"OTHERNET":   4,
		"ETHERNET":   5,
	}
)

func (x NetworkType) Enum() *NetworkType {
	p := new(NetworkType)
	*p = x
	return p
}

func (x NetworkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_metadata_network_network_proto_enumTypes[0].Descriptor()
}

func (NetworkType) Type() protoreflect.EnumType {
	return &file_bilibili_metadata_network_network_proto_enumTypes[0]
}

func (x NetworkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkType.Descriptor instead.
func (NetworkType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_metadata_network_network_proto_rawDescGZIP(), []int{0}
}

// 免流类型
type TFType int32

const (
	TFType_TF_UNKNOWN TFType = 0 // 正常计费
	TFType_U_CARD     TFType = 1 // 联通卡
	TFType_U_PKG      TFType = 2 // 联通包
	TFType_C_CARD     TFType = 3 // 移动卡
	TFType_C_PKG      TFType = 4 // 移动包
	TFType_T_CARD     TFType = 5 // 电信卡
	TFType_T_PKG      TFType = 6 // 电信包
)

// Enum value maps for TFType.
var (
	TFType_name = map[int32]string{
		0: "TF_UNKNOWN",
		1: "U_CARD",
		2: "U_PKG",
		3: "C_CARD",
		4: "C_PKG",
		5: "T_CARD",
		6: "T_PKG",
	}
	TFType_value = map[string]int32{
		"TF_UNKNOWN": 0,
		"U_CARD":     1,
		"U_PKG":      2,
		"C_CARD":     3,
		"C_PKG":      4,
		"T_CARD":     5,
		"T_PKG":      6,
	}
)

func (x TFType) Enum() *TFType {
	p := new(TFType)
	*p = x
	return p
}

func (x TFType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TFType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_metadata_network_network_proto_enumTypes[1].Descriptor()
}

func (TFType) Type() protoreflect.EnumType {
	return &file_bilibili_metadata_network_network_proto_enumTypes[1]
}

func (x TFType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TFType.Descriptor instead.
func (TFType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_metadata_network_network_proto_rawDescGZIP(), []int{1}
}

// 网络类型标识
// gRPC头部:x-bili-network-bin
type Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 网络类型
	Type NetworkType `protobuf:"varint,1,opt,name=type,proto3,enum=bilibili.metadata.network.NetworkType" json:"type,omitempty"`
	// 免流类型
	Tf TFType `protobuf:"varint,2,opt,name=tf,proto3,enum=bilibili.metadata.network.TFType" json:"tf,omitempty"`
	// 运营商
	Oid string `protobuf:"bytes,3,opt,name=oid,proto3" json:"oid,omitempty"`
}

func (x *Network) Reset() {
	*x = Network{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_metadata_network_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_metadata_network_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_bilibili_metadata_network_network_proto_rawDescGZIP(), []int{0}
}

func (x *Network) GetType() NetworkType {
	if x != nil {
		return x.Type
	}
	return NetworkType_NT_UNKNOWN
}

func (x *Network) GetTf() TFType {
	if x != nil {
		return x.Tf
	}
	return TFType_TF_UNKNOWN
}

func (x *Network) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

var File_bilibili_metadata_network_network_proto protoreflect.FileDescriptor

var file_bilibili_metadata_network_network_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x22, 0x8a, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x02,
	0x74, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x54, 0x46, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02, 0x74, 0x66, 0x12,
	0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x69,
	0x64, 0x2a, 0x5e, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x57, 0x49, 0x46, 0x49, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x45,
	0x4c, 0x4c, 0x55, 0x4c, 0x41, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x46, 0x46, 0x4c,
	0x49, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x4e, 0x45,
	0x54, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x54, 0x48, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x10,
	0x05, 0x2a, 0x5d, 0x0a, 0x06, 0x54, 0x46, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x54,
	0x46, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55,
	0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x5f, 0x50, 0x4b, 0x47,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x43, 0x5f, 0x50, 0x4b, 0x47, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x5f, 0x43,
	0x41, 0x52, 0x44, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x5f, 0x50, 0x4b, 0x47, 0x10, 0x06,
	0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58,
	0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_metadata_network_network_proto_rawDescOnce sync.Once
	file_bilibili_metadata_network_network_proto_rawDescData = file_bilibili_metadata_network_network_proto_rawDesc
)

func file_bilibili_metadata_network_network_proto_rawDescGZIP() []byte {
	file_bilibili_metadata_network_network_proto_rawDescOnce.Do(func() {
		file_bilibili_metadata_network_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_metadata_network_network_proto_rawDescData)
	})
	return file_bilibili_metadata_network_network_proto_rawDescData
}

var file_bilibili_metadata_network_network_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bilibili_metadata_network_network_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bilibili_metadata_network_network_proto_goTypes = []interface{}{
	(NetworkType)(0), // 0: bilibili.metadata.network.NetworkType
	(TFType)(0),      // 1: bilibili.metadata.network.TFType
	(*Network)(nil),  // 2: bilibili.metadata.network.Network
}
var file_bilibili_metadata_network_network_proto_depIdxs = []int32{
	0, // 0: bilibili.metadata.network.Network.type:type_name -> bilibili.metadata.network.NetworkType
	1, // 1: bilibili.metadata.network.Network.tf:type_name -> bilibili.metadata.network.TFType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bilibili_metadata_network_network_proto_init() }
func file_bilibili_metadata_network_network_proto_init() {
	if File_bilibili_metadata_network_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_metadata_network_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Network); i {
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
			RawDescriptor: file_bilibili_metadata_network_network_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_metadata_network_network_proto_goTypes,
		DependencyIndexes: file_bilibili_metadata_network_network_proto_depIdxs,
		EnumInfos:         file_bilibili_metadata_network_network_proto_enumTypes,
		MessageInfos:      file_bilibili_metadata_network_network_proto_msgTypes,
	}.Build()
	File_bilibili_metadata_network_network_proto = out.File
	file_bilibili_metadata_network_network_proto_rawDesc = nil
	file_bilibili_metadata_network_network_proto_goTypes = nil
	file_bilibili_metadata_network_network_proto_depIdxs = nil
}
