// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/app/distribution/setting/pegasus.proto

package setting

import (
	v1 "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/app/distribution/v1"
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

type FeedModeValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *v1.Int64Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FeedModeValue) Reset() {
	*x = FeedModeValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedModeValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedModeValue) ProtoMessage() {}

func (x *FeedModeValue) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedModeValue.ProtoReflect.Descriptor instead.
func (*FeedModeValue) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_pegasus_proto_rawDescGZIP(), []int{0}
}

func (x *FeedModeValue) GetValue() *v1.Int64Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type PegasusAutoPlay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Single *v1.Int64Value `protobuf:"bytes,1,opt,name=single,proto3" json:"single,omitempty"`
	Double *v1.Int64Value `protobuf:"bytes,2,opt,name=double,proto3" json:"double,omitempty"`
	SingleAffectedByServerSide *v1.BoolValue `protobuf:"bytes,3,opt,name=single_affected_by_server_side,json=singleAffectedByServerSide,proto3" json:"single_affected_by_server_side,omitempty"`
	DoubleAffectedByServerSide *v1.BoolValue `protobuf:"bytes,4,opt,name=double_affected_by_server_side,json=doubleAffectedByServerSide,proto3" json:"double_affected_by_server_side,omitempty"`
}

func (x *PegasusAutoPlay) Reset() {
	*x = PegasusAutoPlay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PegasusAutoPlay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PegasusAutoPlay) ProtoMessage() {}

func (x *PegasusAutoPlay) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PegasusAutoPlay.ProtoReflect.Descriptor instead.
func (*PegasusAutoPlay) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_pegasus_proto_rawDescGZIP(), []int{1}
}

func (x *PegasusAutoPlay) GetSingle() *v1.Int64Value {
	if x != nil {
		return x.Single
	}
	return nil
}

func (x *PegasusAutoPlay) GetDouble() *v1.Int64Value {
	if x != nil {
		return x.Double
	}
	return nil
}

func (x *PegasusAutoPlay) GetSingleAffectedByServerSide() *v1.BoolValue {
	if x != nil {
		return x.SingleAffectedByServerSide
	}
	return nil
}

func (x *PegasusAutoPlay) GetDoubleAffectedByServerSide() *v1.BoolValue {
	if x != nil {
		return x.DoubleAffectedByServerSide
	}
	return nil
}

type PegasusColumnValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *v1.Int64Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	AffectedByServerSide *v1.BoolValue `protobuf:"bytes,2,opt,name=affected_by_server_side,json=affectedByServerSide,proto3" json:"affected_by_server_side,omitempty"`
}

func (x *PegasusColumnValue) Reset() {
	*x = PegasusColumnValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PegasusColumnValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PegasusColumnValue) ProtoMessage() {}

func (x *PegasusColumnValue) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PegasusColumnValue.ProtoReflect.Descriptor instead.
func (*PegasusColumnValue) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_pegasus_proto_rawDescGZIP(), []int{2}
}

func (x *PegasusColumnValue) GetValue() *v1.Int64Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *PegasusColumnValue) GetAffectedByServerSide() *v1.BoolValue {
	if x != nil {
		return x.AffectedByServerSide
	}
	return nil
}

type PegasusDeviceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column *PegasusColumnValue `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	Mode *FeedModeValue `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
	AutoPlay *PegasusAutoPlay `protobuf:"bytes,3,opt,name=auto_play,json=autoPlay,proto3" json:"auto_play,omitempty"`
}

func (x *PegasusDeviceConfig) Reset() {
	*x = PegasusDeviceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PegasusDeviceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PegasusDeviceConfig) ProtoMessage() {}

func (x *PegasusDeviceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PegasusDeviceConfig.ProtoReflect.Descriptor instead.
func (*PegasusDeviceConfig) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_pegasus_proto_rawDescGZIP(), []int{3}
}

func (x *PegasusDeviceConfig) GetColumn() *PegasusColumnValue {
	if x != nil {
		return x.Column
	}
	return nil
}

func (x *PegasusDeviceConfig) GetMode() *FeedModeValue {
	if x != nil {
		return x.Mode
	}
	return nil
}

func (x *PegasusDeviceConfig) GetAutoPlay() *PegasusAutoPlay {
	if x != nil {
		return x.AutoPlay
	}
	return nil
}

var File_bilibili_app_distribution_setting_pegasus_proto protoreflect.FileDescriptor

var file_bilibili_app_distribution_setting_pegasus_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x67, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x29, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x65, 0x67, 0x61, 0x73, 0x75, 0x73, 0x1a, 0x2f, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a,
	0x0d, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xef,
	0x02, 0x0a, 0x0f, 0x50, 0x65, 0x67, 0x61, 0x73, 0x75, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x50, 0x6c,
	0x61, 0x79, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x73, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x6b, 0x0a, 0x1e, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1a, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x69, 0x64, 0x65, 0x12, 0x6b, 0x0a, 0x1e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x73, 0x69, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x1a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x41, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65,
	0x22, 0xb4, 0x01, 0x0a, 0x12, 0x50, 0x65, 0x67, 0x61, 0x73, 0x75, 0x73, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x5e, 0x0a, 0x17, 0x61, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x69,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x14, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x22, 0x93, 0x02, 0x0a, 0x13, 0x50, 0x65, 0x67, 0x61,
	0x73, 0x75, 0x73, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x55, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x65, 0x67, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x50, 0x65, 0x67, 0x61,
	0x73, 0x75, 0x73, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x4c, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x65, 0x67, 0x61, 0x73, 0x75, 0x73,
	0x2e, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x12, 0x57, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x70, 0x6c, 0x61,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x65, 0x67, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x50, 0x65, 0x67, 0x61, 0x73, 0x75, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x50,
	0x6c, 0x61, 0x79, 0x52, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x42, 0x4e, 0x5a,
	0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f,
	0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_distribution_setting_pegasus_proto_rawDescOnce sync.Once
	file_bilibili_app_distribution_setting_pegasus_proto_rawDescData = file_bilibili_app_distribution_setting_pegasus_proto_rawDesc
)

func file_bilibili_app_distribution_setting_pegasus_proto_rawDescGZIP() []byte {
	file_bilibili_app_distribution_setting_pegasus_proto_rawDescOnce.Do(func() {
		file_bilibili_app_distribution_setting_pegasus_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_distribution_setting_pegasus_proto_rawDescData)
	})
	return file_bilibili_app_distribution_setting_pegasus_proto_rawDescData
}

var file_bilibili_app_distribution_setting_pegasus_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bilibili_app_distribution_setting_pegasus_proto_goTypes = []interface{}{
	(*FeedModeValue)(nil),       // 0: bilibili.app.distribution.setting.pegasus.FeedModeValue
	(*PegasusAutoPlay)(nil),     // 1: bilibili.app.distribution.setting.pegasus.PegasusAutoPlay
	(*PegasusColumnValue)(nil),  // 2: bilibili.app.distribution.setting.pegasus.PegasusColumnValue
	(*PegasusDeviceConfig)(nil), // 3: bilibili.app.distribution.setting.pegasus.PegasusDeviceConfig
	(*v1.Int64Value)(nil),       // 4: bilibili.app.distribution.v1.Int64Value
	(*v1.BoolValue)(nil),        // 5: bilibili.app.distribution.v1.BoolValue
}
var file_bilibili_app_distribution_setting_pegasus_proto_depIdxs = []int32{
	4,  // 0: bilibili.app.distribution.setting.pegasus.FeedModeValue.value:type_name -> bilibili.app.distribution.v1.Int64Value
	4,  // 1: bilibili.app.distribution.setting.pegasus.PegasusAutoPlay.single:type_name -> bilibili.app.distribution.v1.Int64Value
	4,  // 2: bilibili.app.distribution.setting.pegasus.PegasusAutoPlay.double:type_name -> bilibili.app.distribution.v1.Int64Value
	5,  // 3: bilibili.app.distribution.setting.pegasus.PegasusAutoPlay.single_affected_by_server_side:type_name -> bilibili.app.distribution.v1.BoolValue
	5,  // 4: bilibili.app.distribution.setting.pegasus.PegasusAutoPlay.double_affected_by_server_side:type_name -> bilibili.app.distribution.v1.BoolValue
	4,  // 5: bilibili.app.distribution.setting.pegasus.PegasusColumnValue.value:type_name -> bilibili.app.distribution.v1.Int64Value
	5,  // 6: bilibili.app.distribution.setting.pegasus.PegasusColumnValue.affected_by_server_side:type_name -> bilibili.app.distribution.v1.BoolValue
	2,  // 7: bilibili.app.distribution.setting.pegasus.PegasusDeviceConfig.column:type_name -> bilibili.app.distribution.setting.pegasus.PegasusColumnValue
	0,  // 8: bilibili.app.distribution.setting.pegasus.PegasusDeviceConfig.mode:type_name -> bilibili.app.distribution.setting.pegasus.FeedModeValue
	1,  // 9: bilibili.app.distribution.setting.pegasus.PegasusDeviceConfig.auto_play:type_name -> bilibili.app.distribution.setting.pegasus.PegasusAutoPlay
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_bilibili_app_distribution_setting_pegasus_proto_init() }
func file_bilibili_app_distribution_setting_pegasus_proto_init() {
	if File_bilibili_app_distribution_setting_pegasus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedModeValue); i {
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
		file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PegasusAutoPlay); i {
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
		file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PegasusColumnValue); i {
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
		file_bilibili_app_distribution_setting_pegasus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PegasusDeviceConfig); i {
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
			RawDescriptor: file_bilibili_app_distribution_setting_pegasus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_app_distribution_setting_pegasus_proto_goTypes,
		DependencyIndexes: file_bilibili_app_distribution_setting_pegasus_proto_depIdxs,
		MessageInfos:      file_bilibili_app_distribution_setting_pegasus_proto_msgTypes,
	}.Build()
	File_bilibili_app_distribution_setting_pegasus_proto = out.File
	file_bilibili_app_distribution_setting_pegasus_proto_rawDesc = nil
	file_bilibili_app_distribution_setting_pegasus_proto_goTypes = nil
	file_bilibili_app_distribution_setting_pegasus_proto_depIdxs = nil
}