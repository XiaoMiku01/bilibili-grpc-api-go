// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: bilibili/main/common/arch/doll/v1/doll.proto

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

type ErrorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error int32 `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Delay int64 `protobuf:"varint,3,opt,name=delay,proto3" json:"delay,omitempty"`
}

func (x *ErrorRequest) Reset() {
	*x = ErrorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorRequest) ProtoMessage() {}

func (x *ErrorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorRequest.ProtoReflect.Descriptor instead.
func (*ErrorRequest) Descriptor() ([]byte, []int) {
	return file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorRequest) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *ErrorRequest) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ErrorRequest) GetDelay() int64 {
	if x != nil {
		return x.Delay
	}
	return 0
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Time int64 `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ErrorResponse) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescGZIP(), []int{2}
}

func (x *PingRequest) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Time int64 `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescGZIP(), []int{3}
}

func (x *PingResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PingResponse) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type SayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SayRequest) Reset() {
	*x = SayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayRequest) ProtoMessage() {}

func (x *SayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayRequest.ProtoReflect.Descriptor instead.
func (*SayRequest) Descriptor() ([]byte, []int) {
	return file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescGZIP(), []int{4}
}

func (x *SayRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Time int64 `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *SayResponse) Reset() {
	*x = SayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayResponse) ProtoMessage() {}

func (x *SayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayResponse.ProtoReflect.Descriptor instead.
func (*SayResponse) Descriptor() ([]byte, []int) {
	return file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescGZIP(), []int{5}
}

func (x *SayResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SayResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SayResponse) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_bilibili_main_common_arch_doll_v1_doll_proto protoreflect.FileDescriptor

var file_bilibili_main_common_arch_doll_v1_doll_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x64, 0x6f, 0x6c, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x64, 0x6f, 0x6c, 0x6c, 0x2e, 0x76,
	0x31, 0x22, 0x4e, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x22, 0x37, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x36, 0x0a,
	0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x53, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x4f, 0x0a,
	0x0b, 0x53, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0xc1,
	0x02, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x67, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x2e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x64, 0x6f, 0x6c, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x64, 0x6f, 0x6c, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x64, 0x0a, 0x03, 0x53, 0x61, 0x79, 0x12, 0x2d, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x64, 0x6f, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x64, 0x6f, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x2f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x64, 0x6f, 0x6c, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x64, 0x6f, 0x6c,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x64, 0x6f, 0x6c, 0x6c, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescOnce sync.Once
	file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescData = file_bilibili_main_common_arch_doll_v1_doll_proto_rawDesc
)

func file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescGZIP() []byte {
	file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescOnce.Do(func() {
		file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescData)
	})
	return file_bilibili_main_common_arch_doll_v1_doll_proto_rawDescData
}

var file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bilibili_main_common_arch_doll_v1_doll_proto_goTypes = []interface{}{
	(*ErrorRequest)(nil),  // 0: bilibili.main.common.arch.doll.v1.ErrorRequest
	(*ErrorResponse)(nil), // 1: bilibili.main.common.arch.doll.v1.ErrorResponse
	(*PingRequest)(nil),   // 2: bilibili.main.common.arch.doll.v1.PingRequest
	(*PingResponse)(nil),  // 3: bilibili.main.common.arch.doll.v1.PingResponse
	(*SayRequest)(nil),    // 4: bilibili.main.common.arch.doll.v1.SayRequest
	(*SayResponse)(nil),   // 5: bilibili.main.common.arch.doll.v1.SayResponse
}
var file_bilibili_main_common_arch_doll_v1_doll_proto_depIdxs = []int32{
	2, // 0: bilibili.main.common.arch.doll.v1.Echo.Ping:input_type -> bilibili.main.common.arch.doll.v1.PingRequest
	4, // 1: bilibili.main.common.arch.doll.v1.Echo.Say:input_type -> bilibili.main.common.arch.doll.v1.SayRequest
	0, // 2: bilibili.main.common.arch.doll.v1.Echo.Error:input_type -> bilibili.main.common.arch.doll.v1.ErrorRequest
	3, // 3: bilibili.main.common.arch.doll.v1.Echo.Ping:output_type -> bilibili.main.common.arch.doll.v1.PingResponse
	5, // 4: bilibili.main.common.arch.doll.v1.Echo.Say:output_type -> bilibili.main.common.arch.doll.v1.SayResponse
	1, // 5: bilibili.main.common.arch.doll.v1.Echo.Error:output_type -> bilibili.main.common.arch.doll.v1.ErrorResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bilibili_main_common_arch_doll_v1_doll_proto_init() }
func file_bilibili_main_common_arch_doll_v1_doll_proto_init() {
	if File_bilibili_main_common_arch_doll_v1_doll_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorRequest); i {
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
		file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
		file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayRequest); i {
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
		file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayResponse); i {
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
			RawDescriptor: file_bilibili_main_common_arch_doll_v1_doll_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_main_common_arch_doll_v1_doll_proto_goTypes,
		DependencyIndexes: file_bilibili_main_common_arch_doll_v1_doll_proto_depIdxs,
		MessageInfos:      file_bilibili_main_common_arch_doll_v1_doll_proto_msgTypes,
	}.Build()
	File_bilibili_main_common_arch_doll_v1_doll_proto = out.File
	file_bilibili_main_common_arch_doll_v1_doll_proto_rawDesc = nil
	file_bilibili_main_common_arch_doll_v1_doll_proto_goTypes = nil
	file_bilibili_main_common_arch_doll_v1_doll_proto_depIdxs = nil
}
