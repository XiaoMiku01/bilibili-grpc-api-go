// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: bilibili/app/wall/v1/wall.proto

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

// 免流规则信息
type RuleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否支持免流
	Tf bool `protobuf:"varint,1,opt,name=tf,proto3" json:"tf,omitempty"`
	// 操作模式
	// break:无 replace:替换 proxy:代理
	M string `protobuf:"bytes,2,opt,name=m,proto3" json:"m,omitempty"`
	// 操作参数
	A string `protobuf:"bytes,3,opt,name=a,proto3" json:"a,omitempty"`
	// 匹配目标正则
	P       string   `protobuf:"bytes,4,opt,name=p,proto3" json:"p,omitempty"`
	ABackup []string `protobuf:"bytes,5,rep,name=a_backup,json=aBackup,proto3" json:"a_backup,omitempty"`
}

func (x *RuleInfo) Reset() {
	*x = RuleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_wall_v1_wall_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleInfo) ProtoMessage() {}

func (x *RuleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_wall_v1_wall_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleInfo.ProtoReflect.Descriptor instead.
func (*RuleInfo) Descriptor() ([]byte, []int) {
	return file_bilibili_app_wall_v1_wall_proto_rawDescGZIP(), []int{0}
}

func (x *RuleInfo) GetTf() bool {
	if x != nil {
		return x.Tf
	}
	return false
}

func (x *RuleInfo) GetM() string {
	if x != nil {
		return x.M
	}
	return ""
}

func (x *RuleInfo) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *RuleInfo) GetP() string {
	if x != nil {
		return x.P
	}
	return ""
}

func (x *RuleInfo) GetABackup() []string {
	if x != nil {
		return x.ABackup
	}
	return nil
}

// 获取免流规则信息-请求
type RuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RuleRequest) Reset() {
	*x = RuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_wall_v1_wall_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleRequest) ProtoMessage() {}

func (x *RuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_wall_v1_wall_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleRequest.ProtoReflect.Descriptor instead.
func (*RuleRequest) Descriptor() ([]byte, []int) {
	return file_bilibili_app_wall_v1_wall_proto_rawDescGZIP(), []int{1}
}

// 免流规则信息组
type RulesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 免流规则信息
	RulesInfo []*RuleInfo `protobuf:"bytes,1,rep,name=rulesInfo,proto3" json:"rulesInfo,omitempty"`
}

func (x *RulesInfo) Reset() {
	*x = RulesInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_wall_v1_wall_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RulesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RulesInfo) ProtoMessage() {}

func (x *RulesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_wall_v1_wall_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RulesInfo.ProtoReflect.Descriptor instead.
func (*RulesInfo) Descriptor() ([]byte, []int) {
	return file_bilibili_app_wall_v1_wall_proto_rawDescGZIP(), []int{2}
}

func (x *RulesInfo) GetRulesInfo() []*RuleInfo {
	if x != nil {
		return x.RulesInfo
	}
	return nil
}

// 获取免流规则信息-响应
type RulesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 各ISP的免流规则信息组
	// ISP如: cu ct cm
	RulesInfo map[string]*RulesInfo `protobuf:"bytes,1,rep,name=rulesInfo,proto3" json:"rulesInfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HashValue string                `protobuf:"bytes,2,opt,name=hash_value,json=hashValue,proto3" json:"hash_value,omitempty"`
}

func (x *RulesReply) Reset() {
	*x = RulesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_wall_v1_wall_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RulesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RulesReply) ProtoMessage() {}

func (x *RulesReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_wall_v1_wall_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RulesReply.ProtoReflect.Descriptor instead.
func (*RulesReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_wall_v1_wall_proto_rawDescGZIP(), []int{3}
}

func (x *RulesReply) GetRulesInfo() map[string]*RulesInfo {
	if x != nil {
		return x.RulesInfo
	}
	return nil
}

func (x *RulesReply) GetHashValue() string {
	if x != nil {
		return x.HashValue
	}
	return ""
}

var File_bilibili_app_wall_v1_wall_proto protoreflect.FileDescriptor

var file_bilibili_app_wall_v1_wall_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x77,
	0x61, 0x6c, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x5f, 0x0a, 0x08, 0x52, 0x75, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x74, 0x66, 0x12, 0x0c, 0x0a, 0x01, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12,
	0x0c, 0x0a, 0x01, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x70, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x09, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x4d, 0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x61, 0x73, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x5d, 0x0a, 0x0e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x57,
	0x0a, 0x04, 0x57, 0x61, 0x6c, 0x6c, 0x12, 0x4f, 0x0a, 0x08, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x21, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31,
	0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_bilibili_app_wall_v1_wall_proto_rawDescOnce sync.Once
	file_bilibili_app_wall_v1_wall_proto_rawDescData = file_bilibili_app_wall_v1_wall_proto_rawDesc
)

func file_bilibili_app_wall_v1_wall_proto_rawDescGZIP() []byte {
	file_bilibili_app_wall_v1_wall_proto_rawDescOnce.Do(func() {
		file_bilibili_app_wall_v1_wall_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_wall_v1_wall_proto_rawDescData)
	})
	return file_bilibili_app_wall_v1_wall_proto_rawDescData
}

var file_bilibili_app_wall_v1_wall_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bilibili_app_wall_v1_wall_proto_goTypes = []interface{}{
	(*RuleInfo)(nil),    // 0: bilibili.app.wall.v1.RuleInfo
	(*RuleRequest)(nil), // 1: bilibili.app.wall.v1.RuleRequest
	(*RulesInfo)(nil),   // 2: bilibili.app.wall.v1.RulesInfo
	(*RulesReply)(nil),  // 3: bilibili.app.wall.v1.RulesReply
	nil,                 // 4: bilibili.app.wall.v1.RulesReply.RulesInfoEntry
}
var file_bilibili_app_wall_v1_wall_proto_depIdxs = []int32{
	0, // 0: bilibili.app.wall.v1.RulesInfo.rulesInfo:type_name -> bilibili.app.wall.v1.RuleInfo
	4, // 1: bilibili.app.wall.v1.RulesReply.rulesInfo:type_name -> bilibili.app.wall.v1.RulesReply.RulesInfoEntry
	2, // 2: bilibili.app.wall.v1.RulesReply.RulesInfoEntry.value:type_name -> bilibili.app.wall.v1.RulesInfo
	1, // 3: bilibili.app.wall.v1.Wall.RuleInfo:input_type -> bilibili.app.wall.v1.RuleRequest
	3, // 4: bilibili.app.wall.v1.Wall.RuleInfo:output_type -> bilibili.app.wall.v1.RulesReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bilibili_app_wall_v1_wall_proto_init() }
func file_bilibili_app_wall_v1_wall_proto_init() {
	if File_bilibili_app_wall_v1_wall_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_wall_v1_wall_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleInfo); i {
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
		file_bilibili_app_wall_v1_wall_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleRequest); i {
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
		file_bilibili_app_wall_v1_wall_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RulesInfo); i {
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
		file_bilibili_app_wall_v1_wall_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RulesReply); i {
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
			RawDescriptor: file_bilibili_app_wall_v1_wall_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_app_wall_v1_wall_proto_goTypes,
		DependencyIndexes: file_bilibili_app_wall_v1_wall_proto_depIdxs,
		MessageInfos:      file_bilibili_app_wall_v1_wall_proto_msgTypes,
	}.Build()
	File_bilibili_app_wall_v1_wall_proto = out.File
	file_bilibili_app_wall_v1_wall_proto_rawDesc = nil
	file_bilibili_app_wall_v1_wall_proto_goTypes = nil
	file_bilibili_app_wall_v1_wall_proto_depIdxs = nil
}
