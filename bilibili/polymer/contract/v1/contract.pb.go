// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/polymer/contract/v1/contract.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type AddContractReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowMessage bool `protobuf:"varint,1,opt,name=allow_message,json=allowMessage,proto3" json:"allow_message,omitempty"`
	AllowReply bool `protobuf:"varint,2,opt,name=allow_reply,json=allowReply,proto3" json:"allow_reply,omitempty"`
	InputText string `protobuf:"bytes,3,opt,name=input_text,json=inputText,proto3" json:"input_text,omitempty"`
	InputTitle string `protobuf:"bytes,4,opt,name=input_title,json=inputTitle,proto3" json:"input_title,omitempty"`
}

func (x *AddContractReply) Reset() {
	*x = AddContractReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddContractReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddContractReply) ProtoMessage() {}

func (x *AddContractReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddContractReply.ProtoReflect.Descriptor instead.
func (*AddContractReply) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_contract_v1_contract_proto_rawDescGZIP(), []int{0}
}

func (x *AddContractReply) GetAllowMessage() bool {
	if x != nil {
		return x.AllowMessage
	}
	return false
}

func (x *AddContractReply) GetAllowReply() bool {
	if x != nil {
		return x.AllowReply
	}
	return false
}

func (x *AddContractReply) GetInputText() string {
	if x != nil {
		return x.InputText
	}
	return ""
}

func (x *AddContractReply) GetInputTitle() string {
	if x != nil {
		return x.InputTitle
	}
	return ""
}

type AddContractReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common *CommonReq `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	Mid int64 `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	UpMid int64 `protobuf:"varint,3,opt,name=up_mid,json=upMid,proto3" json:"up_mid,omitempty"`
	Aid int64 `protobuf:"varint,4,opt,name=aid,proto3" json:"aid,omitempty"`
	Source int32 `protobuf:"varint,5,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *AddContractReq) Reset() {
	*x = AddContractReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddContractReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddContractReq) ProtoMessage() {}

func (x *AddContractReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddContractReq.ProtoReflect.Descriptor instead.
func (*AddContractReq) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_contract_v1_contract_proto_rawDescGZIP(), []int{1}
}

func (x *AddContractReq) GetCommon() *CommonReq {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *AddContractReq) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *AddContractReq) GetUpMid() int64 {
	if x != nil {
		return x.UpMid
	}
	return 0
}

func (x *AddContractReq) GetAid() int64 {
	if x != nil {
		return x.Aid
	}
	return 0
}

func (x *AddContractReq) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

type CommonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform string `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Build int32 `protobuf:"varint,2,opt,name=build,proto3" json:"build,omitempty"`
	Buvid string `protobuf:"bytes,3,opt,name=buvid,proto3" json:"buvid,omitempty"`
	MobiApp string `protobuf:"bytes,4,opt,name=mobi_app,json=mobiApp,proto3" json:"mobi_app,omitempty"`
	Device string `protobuf:"bytes,5,opt,name=device,proto3" json:"device,omitempty"`
	Ip string `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`
	Spmid string `protobuf:"bytes,7,opt,name=spmid,proto3" json:"spmid,omitempty"`
}

func (x *CommonReq) Reset() {
	*x = CommonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonReq) ProtoMessage() {}

func (x *CommonReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonReq.ProtoReflect.Descriptor instead.
func (*CommonReq) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_contract_v1_contract_proto_rawDescGZIP(), []int{2}
}

func (x *CommonReq) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *CommonReq) GetBuild() int32 {
	if x != nil {
		return x.Build
	}
	return 0
}

func (x *CommonReq) GetBuvid() string {
	if x != nil {
		return x.Buvid
	}
	return ""
}

func (x *CommonReq) GetMobiApp() string {
	if x != nil {
		return x.MobiApp
	}
	return ""
}

func (x *CommonReq) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *CommonReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CommonReq) GetSpmid() string {
	if x != nil {
		return x.Spmid
	}
	return ""
}

type ContractCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	SubTitle string `protobuf:"bytes,2,opt,name=sub_title,json=subTitle,proto3" json:"sub_title,omitempty"`
}

func (x *ContractCard) Reset() {
	*x = ContractCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractCard) ProtoMessage() {}

func (x *ContractCard) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractCard.ProtoReflect.Descriptor instead.
func (*ContractCard) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_contract_v1_contract_proto_rawDescGZIP(), []int{3}
}

func (x *ContractCard) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ContractCard) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

type ContractConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFollowDisplay int32 `protobuf:"varint,1,opt,name=is_follow_display,json=isFollowDisplay,proto3" json:"is_follow_display,omitempty"`
	IsTripleDisplay int32 `protobuf:"varint,2,opt,name=is_triple_display,json=isTripleDisplay,proto3" json:"is_triple_display,omitempty"`
	ContractCard *ContractCard `protobuf:"bytes,3,opt,name=contract_card,json=contractCard,proto3" json:"contract_card,omitempty"`
}

func (x *ContractConfigReply) Reset() {
	*x = ContractConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractConfigReply) ProtoMessage() {}

func (x *ContractConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractConfigReply.ProtoReflect.Descriptor instead.
func (*ContractConfigReply) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_contract_v1_contract_proto_rawDescGZIP(), []int{4}
}

func (x *ContractConfigReply) GetIsFollowDisplay() int32 {
	if x != nil {
		return x.IsFollowDisplay
	}
	return 0
}

func (x *ContractConfigReply) GetIsTripleDisplay() int32 {
	if x != nil {
		return x.IsTripleDisplay
	}
	return 0
}

func (x *ContractConfigReply) GetContractCard() *ContractCard {
	if x != nil {
		return x.ContractCard
	}
	return nil
}

type ContractConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common *CommonReq `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	Mid int64 `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	UpMid int64 `protobuf:"varint,3,opt,name=up_mid,json=upMid,proto3" json:"up_mid,omitempty"`
	Aid int64 `protobuf:"varint,4,opt,name=aid,proto3" json:"aid,omitempty"`
	Source int32 `protobuf:"varint,5,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *ContractConfigReq) Reset() {
	*x = ContractConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractConfigReq) ProtoMessage() {}

func (x *ContractConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_polymer_contract_v1_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractConfigReq.ProtoReflect.Descriptor instead.
func (*ContractConfigReq) Descriptor() ([]byte, []int) {
	return file_bilibili_polymer_contract_v1_contract_proto_rawDescGZIP(), []int{5}
}

func (x *ContractConfigReq) GetCommon() *CommonReq {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *ContractConfigReq) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *ContractConfigReq) GetUpMid() int64 {
	if x != nil {
		return x.UpMid
	}
	return 0
}

func (x *ContractConfigReq) GetAid() int64 {
	if x != nil {
		return x.Aid
	}
	return 0
}

func (x *ContractConfigReq) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

var File_bilibili_polymer_contract_v1_contract_proto protoreflect.FileDescriptor

var file_bilibili_polymer_contract_v1_contract_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x6d,
	0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x3f, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x52,
	0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x75, 0x70, 0x5f,
	0x6d, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x70, 0x4d, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x09, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75,
	0x76, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x76, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x62, 0x69, 0x5f, 0x61, 0x70, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x62, 0x69, 0x41, 0x70, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x6d, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x6d, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xbe, 0x01, 0x0a,
	0x13, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x69, 0x73, 0x54,
	0x72, 0x69, 0x70, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x4f, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70,
	0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x72, 0x64, 0x22, 0xa7, 0x01,
	0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x12, 0x3f, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70,
	0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x52, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x75, 0x70, 0x5f, 0x6d, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x70, 0x4d, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x32, 0xc4, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x12, 0x53, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x2c, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70,
	0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x6d, 0x0a, 0x0d, 0x41, 0x64, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x56, 0x32, 0x12, 0x2c, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x2e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x74, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x31, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x49,
	0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61,
	0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x6d, 0x65, 0x72, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_bilibili_polymer_contract_v1_contract_proto_rawDescOnce sync.Once
	file_bilibili_polymer_contract_v1_contract_proto_rawDescData = file_bilibili_polymer_contract_v1_contract_proto_rawDesc
)

func file_bilibili_polymer_contract_v1_contract_proto_rawDescGZIP() []byte {
	file_bilibili_polymer_contract_v1_contract_proto_rawDescOnce.Do(func() {
		file_bilibili_polymer_contract_v1_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_polymer_contract_v1_contract_proto_rawDescData)
	})
	return file_bilibili_polymer_contract_v1_contract_proto_rawDescData
}

var file_bilibili_polymer_contract_v1_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bilibili_polymer_contract_v1_contract_proto_goTypes = []interface{}{
	(*AddContractReply)(nil),    // 0: bilibili.polymer.contract.v1.AddContractReply
	(*AddContractReq)(nil),      // 1: bilibili.polymer.contract.v1.AddContractReq
	(*CommonReq)(nil),           // 2: bilibili.polymer.contract.v1.CommonReq
	(*ContractCard)(nil),        // 3: bilibili.polymer.contract.v1.ContractCard
	(*ContractConfigReply)(nil), // 4: bilibili.polymer.contract.v1.ContractConfigReply
	(*ContractConfigReq)(nil),   // 5: bilibili.polymer.contract.v1.ContractConfigReq
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_bilibili_polymer_contract_v1_contract_proto_depIdxs = []int32{
	2, // 0: bilibili.polymer.contract.v1.AddContractReq.common:type_name -> bilibili.polymer.contract.v1.CommonReq
	3, // 1: bilibili.polymer.contract.v1.ContractConfigReply.contract_card:type_name -> bilibili.polymer.contract.v1.ContractCard
	2, // 2: bilibili.polymer.contract.v1.ContractConfigReq.common:type_name -> bilibili.polymer.contract.v1.CommonReq
	1, // 3: bilibili.polymer.contract.v1.Contract.AddContract:input_type -> bilibili.polymer.contract.v1.AddContractReq
	1, // 4: bilibili.polymer.contract.v1.Contract.AddContractV2:input_type -> bilibili.polymer.contract.v1.AddContractReq
	5, // 5: bilibili.polymer.contract.v1.Contract.ContractConfig:input_type -> bilibili.polymer.contract.v1.ContractConfigReq
	6, // 6: bilibili.polymer.contract.v1.Contract.AddContract:output_type -> google.protobuf.Empty
	0, // 7: bilibili.polymer.contract.v1.Contract.AddContractV2:output_type -> bilibili.polymer.contract.v1.AddContractReply
	4, // 8: bilibili.polymer.contract.v1.Contract.ContractConfig:output_type -> bilibili.polymer.contract.v1.ContractConfigReply
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bilibili_polymer_contract_v1_contract_proto_init() }
func file_bilibili_polymer_contract_v1_contract_proto_init() {
	if File_bilibili_polymer_contract_v1_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_polymer_contract_v1_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddContractReply); i {
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
		file_bilibili_polymer_contract_v1_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddContractReq); i {
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
		file_bilibili_polymer_contract_v1_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonReq); i {
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
		file_bilibili_polymer_contract_v1_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractCard); i {
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
		file_bilibili_polymer_contract_v1_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractConfigReply); i {
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
		file_bilibili_polymer_contract_v1_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractConfigReq); i {
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
			RawDescriptor: file_bilibili_polymer_contract_v1_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_polymer_contract_v1_contract_proto_goTypes,
		DependencyIndexes: file_bilibili_polymer_contract_v1_contract_proto_depIdxs,
		MessageInfos:      file_bilibili_polymer_contract_v1_contract_proto_msgTypes,
	}.Build()
	File_bilibili_polymer_contract_v1_contract_proto = out.File
	file_bilibili_polymer_contract_v1_contract_proto_rawDesc = nil
	file_bilibili_polymer_contract_v1_contract_proto_goTypes = nil
	file_bilibili_polymer_contract_v1_contract_proto_depIdxs = nil
}
