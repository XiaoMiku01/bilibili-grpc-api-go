// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/app/viewunite/ugcanymodel.proto

package viewunite

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

type ViewUgcAny struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ViewUgcAny) Reset() {
	*x = ViewUgcAny{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_viewunite_ugcanymodel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewUgcAny) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewUgcAny) ProtoMessage() {}

func (x *ViewUgcAny) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_viewunite_ugcanymodel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewUgcAny.ProtoReflect.Descriptor instead.
func (*ViewUgcAny) Descriptor() ([]byte, []int) {
	return file_bilibili_app_viewunite_ugcanymodel_proto_rawDescGZIP(), []int{0}
}

var File_bilibili_app_viewunite_ugcanymodel_proto protoreflect.FileDescriptor

var file_bilibili_app_viewunite_ugcanymodel_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76,
	0x69, 0x65, 0x77, 0x75, 0x6e, 0x69, 0x74, 0x65, 0x2f, 0x75, 0x67, 0x63, 0x61, 0x6e, 0x79, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x75, 0x6e, 0x69,
	0x74, 0x65, 0x2e, 0x75, 0x67, 0x63, 0x61, 0x6e, 0x79, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x0c,
	0x0a, 0x0a, 0x56, 0x69, 0x65, 0x77, 0x55, 0x67, 0x63, 0x41, 0x6e, 0x79, 0x42, 0x43, 0x5a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d,
	0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x75, 0x6e, 0x69, 0x74,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_viewunite_ugcanymodel_proto_rawDescOnce sync.Once
	file_bilibili_app_viewunite_ugcanymodel_proto_rawDescData = file_bilibili_app_viewunite_ugcanymodel_proto_rawDesc
)

func file_bilibili_app_viewunite_ugcanymodel_proto_rawDescGZIP() []byte {
	file_bilibili_app_viewunite_ugcanymodel_proto_rawDescOnce.Do(func() {
		file_bilibili_app_viewunite_ugcanymodel_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_viewunite_ugcanymodel_proto_rawDescData)
	})
	return file_bilibili_app_viewunite_ugcanymodel_proto_rawDescData
}

var file_bilibili_app_viewunite_ugcanymodel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bilibili_app_viewunite_ugcanymodel_proto_goTypes = []interface{}{
	(*ViewUgcAny)(nil), // 0: bilibili.app.viewunite.ugcanymodel.ViewUgcAny
}
var file_bilibili_app_viewunite_ugcanymodel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bilibili_app_viewunite_ugcanymodel_proto_init() }
func file_bilibili_app_viewunite_ugcanymodel_proto_init() {
	if File_bilibili_app_viewunite_ugcanymodel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_viewunite_ugcanymodel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewUgcAny); i {
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
			RawDescriptor: file_bilibili_app_viewunite_ugcanymodel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_app_viewunite_ugcanymodel_proto_goTypes,
		DependencyIndexes: file_bilibili_app_viewunite_ugcanymodel_proto_depIdxs,
		MessageInfos:      file_bilibili_app_viewunite_ugcanymodel_proto_msgTypes,
	}.Build()
	File_bilibili_app_viewunite_ugcanymodel_proto = out.File
	file_bilibili_app_viewunite_ugcanymodel_proto_rawDesc = nil
	file_bilibili_app_viewunite_ugcanymodel_proto_goTypes = nil
	file_bilibili_app_viewunite_ugcanymodel_proto_depIdxs = nil
}