// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/app/resource/privacy/v1/api.proto

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

// 隐私开关状态
type PrivacyConfigState int32

const (
	PrivacyConfigState_close PrivacyConfigState = 0 // 关闭
	PrivacyConfigState_open  PrivacyConfigState = 1 // 打开
)

// Enum value maps for PrivacyConfigState.
var (
	PrivacyConfigState_name = map[int32]string{
		0: "close",
		1: "open",
	}
	PrivacyConfigState_value = map[string]int32{
		"close": 0,
		"open":  1,
	}
)

func (x PrivacyConfigState) Enum() *PrivacyConfigState {
	p := new(PrivacyConfigState)
	*p = x
	return p
}

func (x PrivacyConfigState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrivacyConfigState) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_app_resource_privacy_v1_api_proto_enumTypes[0].Descriptor()
}

func (PrivacyConfigState) Type() protoreflect.EnumType {
	return &file_bilibili_app_resource_privacy_v1_api_proto_enumTypes[0]
}

func (x PrivacyConfigState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrivacyConfigState.Descriptor instead.
func (PrivacyConfigState) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_app_resource_privacy_v1_api_proto_rawDescGZIP(), []int{0}
}

// 隐私开关类型
type PrivacyConfigType int32

const (
	PrivacyConfigType_none         PrivacyConfigType = 0 //
	PrivacyConfigType_dynamic_city PrivacyConfigType = 1 // 动态同城
)

// Enum value maps for PrivacyConfigType.
var (
	PrivacyConfigType_name = map[int32]string{
		0: "none",
		1: "dynamic_city",
	}
	PrivacyConfigType_value = map[string]int32{
		"none":         0,
		"dynamic_city": 1,
	}
)

func (x PrivacyConfigType) Enum() *PrivacyConfigType {
	p := new(PrivacyConfigType)
	*p = x
	return p
}

func (x PrivacyConfigType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrivacyConfigType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_app_resource_privacy_v1_api_proto_enumTypes[1].Descriptor()
}

func (PrivacyConfigType) Type() protoreflect.EnumType {
	return &file_bilibili_app_resource_privacy_v1_api_proto_enumTypes[1]
}

func (x PrivacyConfigType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrivacyConfigType.Descriptor instead.
func (PrivacyConfigType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_app_resource_privacy_v1_api_proto_rawDescGZIP(), []int{1}
}

// 空请求
type NoArgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoArgRequest) Reset() {
	*x = NoArgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoArgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoArgRequest) ProtoMessage() {}

func (x *NoArgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoArgRequest.ProtoReflect.Descriptor instead.
func (*NoArgRequest) Descriptor() ([]byte, []int) {
	return file_bilibili_app_resource_privacy_v1_api_proto_rawDescGZIP(), []int{0}
}

// 空响应
type NoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoReply) Reset() {
	*x = NoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoReply) ProtoMessage() {}

func (x *NoReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoReply.ProtoReflect.Descriptor instead.
func (*NoReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_resource_privacy_v1_api_proto_rawDescGZIP(), []int{1}
}

// 隐私设置
type PrivacyConfigItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 隐私开关类型
	PrivacyConfigType PrivacyConfigType `protobuf:"varint,1,opt,name=privacy_config_type,json=privacyConfigType,proto3,enum=bilibili.app.resource.privacy.v1.PrivacyConfigType" json:"privacy_config_type,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// 隐私开关状态
	State PrivacyConfigState `protobuf:"varint,3,opt,name=state,proto3,enum=bilibili.app.resource.privacy.v1.PrivacyConfigState" json:"state,omitempty"`
	SubTitle string `protobuf:"bytes,4,opt,name=sub_title,json=subTitle,proto3" json:"sub_title,omitempty"`
	SubTitleUri string `protobuf:"bytes,5,opt,name=sub_title_uri,json=subTitleUri,proto3" json:"sub_title_uri,omitempty"`
}

func (x *PrivacyConfigItem) Reset() {
	*x = PrivacyConfigItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivacyConfigItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivacyConfigItem) ProtoMessage() {}

func (x *PrivacyConfigItem) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivacyConfigItem.ProtoReflect.Descriptor instead.
func (*PrivacyConfigItem) Descriptor() ([]byte, []int) {
	return file_bilibili_app_resource_privacy_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *PrivacyConfigItem) GetPrivacyConfigType() PrivacyConfigType {
	if x != nil {
		return x.PrivacyConfigType
	}
	return PrivacyConfigType_none
}

func (x *PrivacyConfigItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PrivacyConfigItem) GetState() PrivacyConfigState {
	if x != nil {
		return x.State
	}
	return PrivacyConfigState_close
}

func (x *PrivacyConfigItem) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *PrivacyConfigItem) GetSubTitleUri() string {
	if x != nil {
		return x.SubTitleUri
	}
	return ""
}

// 获取隐私设置-响应
type PrivacyConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 隐私设置
	PrivacyConfigItem *PrivacyConfigItem `protobuf:"bytes,1,opt,name=privacy_config_item,json=privacyConfigItem,proto3" json:"privacy_config_item,omitempty"`
}

func (x *PrivacyConfigReply) Reset() {
	*x = PrivacyConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivacyConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivacyConfigReply) ProtoMessage() {}

func (x *PrivacyConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivacyConfigReply.ProtoReflect.Descriptor instead.
func (*PrivacyConfigReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_resource_privacy_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *PrivacyConfigReply) GetPrivacyConfigItem() *PrivacyConfigItem {
	if x != nil {
		return x.PrivacyConfigItem
	}
	return nil
}

// 修改隐私设置-请求
type SetPrivacyConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 隐私开关类型
	PrivacyConfigType PrivacyConfigType `protobuf:"varint,1,opt,name=privacy_config_type,json=privacyConfigType,proto3,enum=bilibili.app.resource.privacy.v1.PrivacyConfigType" json:"privacy_config_type,omitempty"`
	// 隐私开关状态
	State PrivacyConfigState `protobuf:"varint,2,opt,name=state,proto3,enum=bilibili.app.resource.privacy.v1.PrivacyConfigState" json:"state,omitempty"`
}

func (x *SetPrivacyConfigRequest) Reset() {
	*x = SetPrivacyConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPrivacyConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPrivacyConfigRequest) ProtoMessage() {}

func (x *SetPrivacyConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPrivacyConfigRequest.ProtoReflect.Descriptor instead.
func (*SetPrivacyConfigRequest) Descriptor() ([]byte, []int) {
	return file_bilibili_app_resource_privacy_v1_api_proto_rawDescGZIP(), []int{4}
}

func (x *SetPrivacyConfigRequest) GetPrivacyConfigType() PrivacyConfigType {
	if x != nil {
		return x.PrivacyConfigType
	}
	return PrivacyConfigType_none
}

func (x *SetPrivacyConfigRequest) GetState() PrivacyConfigState {
	if x != nil {
		return x.State
	}
	return PrivacyConfigState_close
}

var File_bilibili_app_resource_privacy_v1_api_proto protoreflect.FileDescriptor

var file_bilibili_app_resource_privacy_v1_api_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x0e,
	0x0a, 0x0c, 0x4e, 0x6f, 0x41, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x09,
	0x0a, 0x07, 0x4e, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x9b, 0x02, 0x0a, 0x11, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x63, 0x0a, 0x13, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x11, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x22, 0x79, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x63, 0x0a,
	0x13, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x11, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74,
	0x65, 0x6d, 0x22, 0xca, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x63,
	0x0a, 0x13, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x11, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x34, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a,
	0x29, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x10, 0x01, 0x2a, 0x2f, 0x0a, 0x11, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x10, 0x01, 0x32, 0xfa, 0x01, 0x0a, 0x07,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x75, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x41, 0x72,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x78,
	0x0a, 0x10, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x39, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30,
	0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_resource_privacy_v1_api_proto_rawDescOnce sync.Once
	file_bilibili_app_resource_privacy_v1_api_proto_rawDescData = file_bilibili_app_resource_privacy_v1_api_proto_rawDesc
)

func file_bilibili_app_resource_privacy_v1_api_proto_rawDescGZIP() []byte {
	file_bilibili_app_resource_privacy_v1_api_proto_rawDescOnce.Do(func() {
		file_bilibili_app_resource_privacy_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_resource_privacy_v1_api_proto_rawDescData)
	})
	return file_bilibili_app_resource_privacy_v1_api_proto_rawDescData
}

var file_bilibili_app_resource_privacy_v1_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bilibili_app_resource_privacy_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bilibili_app_resource_privacy_v1_api_proto_goTypes = []interface{}{
	(PrivacyConfigState)(0),         // 0: bilibili.app.resource.privacy.v1.PrivacyConfigState
	(PrivacyConfigType)(0),          // 1: bilibili.app.resource.privacy.v1.PrivacyConfigType
	(*NoArgRequest)(nil),            // 2: bilibili.app.resource.privacy.v1.NoArgRequest
	(*NoReply)(nil),                 // 3: bilibili.app.resource.privacy.v1.NoReply
	(*PrivacyConfigItem)(nil),       // 4: bilibili.app.resource.privacy.v1.PrivacyConfigItem
	(*PrivacyConfigReply)(nil),      // 5: bilibili.app.resource.privacy.v1.PrivacyConfigReply
	(*SetPrivacyConfigRequest)(nil), // 6: bilibili.app.resource.privacy.v1.SetPrivacyConfigRequest
}
var file_bilibili_app_resource_privacy_v1_api_proto_depIdxs = []int32{
	1, // 0: bilibili.app.resource.privacy.v1.PrivacyConfigItem.privacy_config_type:type_name -> bilibili.app.resource.privacy.v1.PrivacyConfigType
	0, // 1: bilibili.app.resource.privacy.v1.PrivacyConfigItem.state:type_name -> bilibili.app.resource.privacy.v1.PrivacyConfigState
	4, // 2: bilibili.app.resource.privacy.v1.PrivacyConfigReply.privacy_config_item:type_name -> bilibili.app.resource.privacy.v1.PrivacyConfigItem
	1, // 3: bilibili.app.resource.privacy.v1.SetPrivacyConfigRequest.privacy_config_type:type_name -> bilibili.app.resource.privacy.v1.PrivacyConfigType
	0, // 4: bilibili.app.resource.privacy.v1.SetPrivacyConfigRequest.state:type_name -> bilibili.app.resource.privacy.v1.PrivacyConfigState
	2, // 5: bilibili.app.resource.privacy.v1.Privacy.PrivacyConfig:input_type -> bilibili.app.resource.privacy.v1.NoArgRequest
	6, // 6: bilibili.app.resource.privacy.v1.Privacy.SetPrivacyConfig:input_type -> bilibili.app.resource.privacy.v1.SetPrivacyConfigRequest
	5, // 7: bilibili.app.resource.privacy.v1.Privacy.PrivacyConfig:output_type -> bilibili.app.resource.privacy.v1.PrivacyConfigReply
	3, // 8: bilibili.app.resource.privacy.v1.Privacy.SetPrivacyConfig:output_type -> bilibili.app.resource.privacy.v1.NoReply
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_bilibili_app_resource_privacy_v1_api_proto_init() }
func file_bilibili_app_resource_privacy_v1_api_proto_init() {
	if File_bilibili_app_resource_privacy_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoArgRequest); i {
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
		file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoReply); i {
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
		file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivacyConfigItem); i {
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
		file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivacyConfigReply); i {
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
		file_bilibili_app_resource_privacy_v1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPrivacyConfigRequest); i {
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
			RawDescriptor: file_bilibili_app_resource_privacy_v1_api_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_app_resource_privacy_v1_api_proto_goTypes,
		DependencyIndexes: file_bilibili_app_resource_privacy_v1_api_proto_depIdxs,
		EnumInfos:         file_bilibili_app_resource_privacy_v1_api_proto_enumTypes,
		MessageInfos:      file_bilibili_app_resource_privacy_v1_api_proto_msgTypes,
	}.Build()
	File_bilibili_app_resource_privacy_v1_api_proto = out.File
	file_bilibili_app_resource_privacy_v1_api_proto_rawDesc = nil
	file_bilibili_app_resource_privacy_v1_api_proto_goTypes = nil
	file_bilibili_app_resource_privacy_v1_api_proto_depIdxs = nil
}
