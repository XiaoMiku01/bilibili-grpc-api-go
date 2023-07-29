// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/app/distribution/setting/experimental.proto

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

type DynamicSelect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fold *v1.BoolValue `protobuf:"bytes,1,opt,name=fold,proto3" json:"fold,omitempty"`
}

func (x *DynamicSelect) Reset() {
	*x = DynamicSelect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_experimental_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicSelect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicSelect) ProtoMessage() {}

func (x *DynamicSelect) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_experimental_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicSelect.ProtoReflect.Descriptor instead.
func (*DynamicSelect) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_experimental_proto_rawDescGZIP(), []int{0}
}

func (x *DynamicSelect) GetFold() *v1.BoolValue {
	if x != nil {
		return x.Fold
	}
	return nil
}

type Exp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *v1.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Bucket *v1.Int32Value `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
}

func (x *Exp) Reset() {
	*x = Exp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_experimental_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exp) ProtoMessage() {}

func (x *Exp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_experimental_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exp.ProtoReflect.Descriptor instead.
func (*Exp) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_experimental_proto_rawDescGZIP(), []int{1}
}

func (x *Exp) GetId() *v1.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Exp) GetBucket() *v1.Int32Value {
	if x != nil {
		return x.Bucket
	}
	return nil
}

type ExperimentalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag *v1.StringValue `protobuf:"bytes,1,opt,name=flag,proto3" json:"flag,omitempty"`
	Exps []*Exp `protobuf:"bytes,2,rep,name=exps,proto3" json:"exps,omitempty"`
}

func (x *ExperimentalConfig) Reset() {
	*x = ExperimentalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_experimental_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentalConfig) ProtoMessage() {}

func (x *ExperimentalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_experimental_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentalConfig.ProtoReflect.Descriptor instead.
func (*ExperimentalConfig) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_experimental_proto_rawDescGZIP(), []int{2}
}

func (x *ExperimentalConfig) GetFlag() *v1.StringValue {
	if x != nil {
		return x.Flag
	}
	return nil
}

func (x *ExperimentalConfig) GetExps() []*Exp {
	if x != nil {
		return x.Exps
	}
	return nil
}

type MultipleTusConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopLeft *TopLeft `protobuf:"bytes,1,opt,name=top_left,json=topLeft,proto3" json:"top_left,omitempty"`
	DynamicSelect *DynamicSelect `protobuf:"bytes,2,opt,name=dynamic_select,json=dynamicSelect,proto3" json:"dynamic_select,omitempty"`
}

func (x *MultipleTusConfig) Reset() {
	*x = MultipleTusConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_experimental_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleTusConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleTusConfig) ProtoMessage() {}

func (x *MultipleTusConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_experimental_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleTusConfig.ProtoReflect.Descriptor instead.
func (*MultipleTusConfig) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_experimental_proto_rawDescGZIP(), []int{3}
}

func (x *MultipleTusConfig) GetTopLeft() *TopLeft {
	if x != nil {
		return x.TopLeft
	}
	return nil
}

func (x *MultipleTusConfig) GetDynamicSelect() *DynamicSelect {
	if x != nil {
		return x.DynamicSelect
	}
	return nil
}

// APP首页头像跳转信息
type TopLeft struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url *v1.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	StoryForegroundImage *v1.StringValue `protobuf:"bytes,2,opt,name=story_foreground_image,json=storyForegroundImage,proto3" json:"story_foreground_image,omitempty"`
	StoryBackgroundImage *v1.StringValue `protobuf:"bytes,3,opt,name=story_background_image,json=storyBackgroundImage,proto3" json:"story_background_image,omitempty"`
	ListenForegroundImage *v1.StringValue `protobuf:"bytes,4,opt,name=listen_foreground_image,json=listenForegroundImage,proto3" json:"listen_foreground_image,omitempty"`
	ListenBackgroundImage *v1.StringValue `protobuf:"bytes,5,opt,name=listen_background_image,json=listenBackgroundImage,proto3" json:"listen_background_image,omitempty"`
	IosStoryForegroundImage *v1.StringValue `protobuf:"bytes,6,opt,name=ios_story_foreground_image,json=iosStoryForegroundImage,proto3" json:"ios_story_foreground_image,omitempty"`
	IosStoryBackgroundImage *v1.StringValue `protobuf:"bytes,7,opt,name=ios_story_background_image,json=iosStoryBackgroundImage,proto3" json:"ios_story_background_image,omitempty"`
	IosListenForegroundImage *v1.StringValue `protobuf:"bytes,8,opt,name=ios_listen_foreground_image,json=iosListenForegroundImage,proto3" json:"ios_listen_foreground_image,omitempty"`
	IosListenBackgroundImage *v1.StringValue `protobuf:"bytes,9,opt,name=ios_listen_background_image,json=iosListenBackgroundImage,proto3" json:"ios_listen_background_image,omitempty"`
	Goto *v1.StringValue `protobuf:"bytes,10,opt,name=goto,proto3" json:"goto,omitempty"`
	UrlV2 *v1.StringValue `protobuf:"bytes,11,opt,name=url_v2,json=urlV2,proto3" json:"url_v2,omitempty"`
	GotoV2 *v1.Int64Value `protobuf:"bytes,12,opt,name=goto_v2,json=gotoV2,proto3" json:"goto_v2,omitempty"`
	Badge *v1.StringValue `protobuf:"bytes,13,opt,name=badge,proto3" json:"badge,omitempty"`
}

func (x *TopLeft) Reset() {
	*x = TopLeft{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_experimental_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopLeft) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopLeft) ProtoMessage() {}

func (x *TopLeft) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_experimental_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopLeft.ProtoReflect.Descriptor instead.
func (*TopLeft) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_experimental_proto_rawDescGZIP(), []int{4}
}

func (x *TopLeft) GetUrl() *v1.StringValue {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *TopLeft) GetStoryForegroundImage() *v1.StringValue {
	if x != nil {
		return x.StoryForegroundImage
	}
	return nil
}

func (x *TopLeft) GetStoryBackgroundImage() *v1.StringValue {
	if x != nil {
		return x.StoryBackgroundImage
	}
	return nil
}

func (x *TopLeft) GetListenForegroundImage() *v1.StringValue {
	if x != nil {
		return x.ListenForegroundImage
	}
	return nil
}

func (x *TopLeft) GetListenBackgroundImage() *v1.StringValue {
	if x != nil {
		return x.ListenBackgroundImage
	}
	return nil
}

func (x *TopLeft) GetIosStoryForegroundImage() *v1.StringValue {
	if x != nil {
		return x.IosStoryForegroundImage
	}
	return nil
}

func (x *TopLeft) GetIosStoryBackgroundImage() *v1.StringValue {
	if x != nil {
		return x.IosStoryBackgroundImage
	}
	return nil
}

func (x *TopLeft) GetIosListenForegroundImage() *v1.StringValue {
	if x != nil {
		return x.IosListenForegroundImage
	}
	return nil
}

func (x *TopLeft) GetIosListenBackgroundImage() *v1.StringValue {
	if x != nil {
		return x.IosListenBackgroundImage
	}
	return nil
}

func (x *TopLeft) GetGoto() *v1.StringValue {
	if x != nil {
		return x.Goto
	}
	return nil
}

func (x *TopLeft) GetUrlV2() *v1.StringValue {
	if x != nil {
		return x.UrlV2
	}
	return nil
}

func (x *TopLeft) GetGotoV2() *v1.Int64Value {
	if x != nil {
		return x.GotoV2
	}
	return nil
}

func (x *TopLeft) GetBadge() *v1.StringValue {
	if x != nil {
		return x.Badge
	}
	return nil
}

var File_bilibili_app_distribution_setting_experimental_proto protoreflect.FileDescriptor

var file_bilibili_app_distribution_setting_experimental_proto_rawDesc = []byte{
	0x0a, 0x34, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x1a, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0d, 0x44, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x66, 0x6f, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x66, 0x6f, 0x6c, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x03, 0x45, 0x78, 0x70, 0x12, 0x38, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x9c, 0x01, 0x0a, 0x12, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3d, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12,
	0x47, 0x0a, 0x04, 0x65, 0x78, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x45,
	0x78, 0x70, 0x52, 0x04, 0x65, 0x78, 0x70, 0x73, 0x22, 0xcd, 0x01, 0x0a, 0x11, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x54, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x52,
	0x0a, 0x08, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x2e, 0x54, 0x6f, 0x70, 0x4c, 0x65, 0x66, 0x74, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x4c, 0x65,
	0x66, 0x74, 0x12, 0x64, 0x0a, 0x0e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x44, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x0d, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x22, 0xf7, 0x08, 0x0a, 0x07, 0x54, 0x6f, 0x70,
	0x4c, 0x65, 0x66, 0x74, 0x12, 0x3b, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x5f, 0x0a, 0x16, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x46, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x5f, 0x0a, 0x16, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x61, 0x0a, 0x17, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x66, 0x6f,
	0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x15, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x61, 0x0a, 0x17, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x15, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x66, 0x0a, 0x1a, 0x69, 0x6f, 0x73,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x17, 0x69, 0x6f, 0x73, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x46, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x66, 0x0a, 0x1a, 0x69, 0x6f, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x17, 0x69, 0x6f, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x68, 0x0a, 0x1b, 0x69, 0x6f, 0x73,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x69, 0x6f, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x68, 0x0a, 0x1b, 0x69, 0x6f, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x18, 0x69, 0x6f, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x42, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a,
	0x04, 0x67, 0x6f, 0x74, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x67, 0x6f, 0x74, 0x6f, 0x12, 0x40, 0x0a, 0x06,
	0x75, 0x72, 0x6c, 0x5f, 0x76, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x75, 0x72, 0x6c, 0x56, 0x32, 0x12, 0x41,
	0x0a, 0x07, 0x67, 0x6f, 0x74, 0x6f, 0x5f, 0x76, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x67, 0x6f, 0x74, 0x6f, 0x56,
	0x32, 0x12, 0x3f, 0x0a, 0x05, 0x62, 0x61, 0x64, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x62, 0x61, 0x64,
	0x67, 0x65, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_distribution_setting_experimental_proto_rawDescOnce sync.Once
	file_bilibili_app_distribution_setting_experimental_proto_rawDescData = file_bilibili_app_distribution_setting_experimental_proto_rawDesc
)

func file_bilibili_app_distribution_setting_experimental_proto_rawDescGZIP() []byte {
	file_bilibili_app_distribution_setting_experimental_proto_rawDescOnce.Do(func() {
		file_bilibili_app_distribution_setting_experimental_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_distribution_setting_experimental_proto_rawDescData)
	})
	return file_bilibili_app_distribution_setting_experimental_proto_rawDescData
}

var file_bilibili_app_distribution_setting_experimental_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bilibili_app_distribution_setting_experimental_proto_goTypes = []interface{}{
	(*DynamicSelect)(nil),      // 0: bilibili.app.distribution.setting.experimental.DynamicSelect
	(*Exp)(nil),                // 1: bilibili.app.distribution.setting.experimental.Exp
	(*ExperimentalConfig)(nil), // 2: bilibili.app.distribution.setting.experimental.ExperimentalConfig
	(*MultipleTusConfig)(nil),  // 3: bilibili.app.distribution.setting.experimental.MultipleTusConfig
	(*TopLeft)(nil),            // 4: bilibili.app.distribution.setting.experimental.TopLeft
	(*v1.BoolValue)(nil),       // 5: bilibili.app.distribution.v1.BoolValue
	(*v1.Int64Value)(nil),      // 6: bilibili.app.distribution.v1.Int64Value
	(*v1.Int32Value)(nil),      // 7: bilibili.app.distribution.v1.Int32Value
	(*v1.StringValue)(nil),     // 8: bilibili.app.distribution.v1.StringValue
}
var file_bilibili_app_distribution_setting_experimental_proto_depIdxs = []int32{
	5,  // 0: bilibili.app.distribution.setting.experimental.DynamicSelect.fold:type_name -> bilibili.app.distribution.v1.BoolValue
	6,  // 1: bilibili.app.distribution.setting.experimental.Exp.id:type_name -> bilibili.app.distribution.v1.Int64Value
	7,  // 2: bilibili.app.distribution.setting.experimental.Exp.bucket:type_name -> bilibili.app.distribution.v1.Int32Value
	8,  // 3: bilibili.app.distribution.setting.experimental.ExperimentalConfig.flag:type_name -> bilibili.app.distribution.v1.StringValue
	1,  // 4: bilibili.app.distribution.setting.experimental.ExperimentalConfig.exps:type_name -> bilibili.app.distribution.setting.experimental.Exp
	4,  // 5: bilibili.app.distribution.setting.experimental.MultipleTusConfig.top_left:type_name -> bilibili.app.distribution.setting.experimental.TopLeft
	0,  // 6: bilibili.app.distribution.setting.experimental.MultipleTusConfig.dynamic_select:type_name -> bilibili.app.distribution.setting.experimental.DynamicSelect
	8,  // 7: bilibili.app.distribution.setting.experimental.TopLeft.url:type_name -> bilibili.app.distribution.v1.StringValue
	8,  // 8: bilibili.app.distribution.setting.experimental.TopLeft.story_foreground_image:type_name -> bilibili.app.distribution.v1.StringValue
	8,  // 9: bilibili.app.distribution.setting.experimental.TopLeft.story_background_image:type_name -> bilibili.app.distribution.v1.StringValue
	8,  // 10: bilibili.app.distribution.setting.experimental.TopLeft.listen_foreground_image:type_name -> bilibili.app.distribution.v1.StringValue
	8,  // 11: bilibili.app.distribution.setting.experimental.TopLeft.listen_background_image:type_name -> bilibili.app.distribution.v1.StringValue
	8,  // 12: bilibili.app.distribution.setting.experimental.TopLeft.ios_story_foreground_image:type_name -> bilibili.app.distribution.v1.StringValue
	8,  // 13: bilibili.app.distribution.setting.experimental.TopLeft.ios_story_background_image:type_name -> bilibili.app.distribution.v1.StringValue
	8,  // 14: bilibili.app.distribution.setting.experimental.TopLeft.ios_listen_foreground_image:type_name -> bilibili.app.distribution.v1.StringValue
	8,  // 15: bilibili.app.distribution.setting.experimental.TopLeft.ios_listen_background_image:type_name -> bilibili.app.distribution.v1.StringValue
	8,  // 16: bilibili.app.distribution.setting.experimental.TopLeft.goto:type_name -> bilibili.app.distribution.v1.StringValue
	8,  // 17: bilibili.app.distribution.setting.experimental.TopLeft.url_v2:type_name -> bilibili.app.distribution.v1.StringValue
	6,  // 18: bilibili.app.distribution.setting.experimental.TopLeft.goto_v2:type_name -> bilibili.app.distribution.v1.Int64Value
	8,  // 19: bilibili.app.distribution.setting.experimental.TopLeft.badge:type_name -> bilibili.app.distribution.v1.StringValue
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_bilibili_app_distribution_setting_experimental_proto_init() }
func file_bilibili_app_distribution_setting_experimental_proto_init() {
	if File_bilibili_app_distribution_setting_experimental_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_distribution_setting_experimental_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicSelect); i {
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
		file_bilibili_app_distribution_setting_experimental_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exp); i {
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
		file_bilibili_app_distribution_setting_experimental_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentalConfig); i {
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
		file_bilibili_app_distribution_setting_experimental_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleTusConfig); i {
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
		file_bilibili_app_distribution_setting_experimental_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopLeft); i {
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
			RawDescriptor: file_bilibili_app_distribution_setting_experimental_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_app_distribution_setting_experimental_proto_goTypes,
		DependencyIndexes: file_bilibili_app_distribution_setting_experimental_proto_depIdxs,
		MessageInfos:      file_bilibili_app_distribution_setting_experimental_proto_msgTypes,
	}.Build()
	File_bilibili_app_distribution_setting_experimental_proto = out.File
	file_bilibili_app_distribution_setting_experimental_proto_rawDesc = nil
	file_bilibili_app_distribution_setting_experimental_proto_goTypes = nil
	file_bilibili_app_distribution_setting_experimental_proto_depIdxs = nil
}