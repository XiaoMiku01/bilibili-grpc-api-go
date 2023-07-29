// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: pgc/gateway/vega/v1/vega.proto

package v1

import (
	rpc "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/rpc"
	any1 "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type AuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReq.ProtoReflect.Descriptor instead.
func (*AuthReq) Descriptor() ([]byte, []int) {
	return file_pgc_gateway_vega_v1_vega_proto_rawDescGZIP(), []int{0}
}

type AuthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthResp) Reset() {
	*x = AuthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResp) ProtoMessage() {}

func (x *AuthResp) ProtoReflect() protoreflect.Message {
	mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResp.ProtoReflect.Descriptor instead.
func (*AuthResp) Descriptor() ([]byte, []int) {
	return file_pgc_gateway_vega_v1_vega_proto_rawDescGZIP(), []int{1}
}

type FrameOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VegaId int64 `protobuf:"varint,1,opt,name=vega_id,json=vegaId,proto3" json:"vega_id,omitempty"`
	ReqId string `protobuf:"bytes,2,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
	Sequence int64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	IsAck bool `protobuf:"varint,4,opt,name=is_ack,json=isAck,proto3" json:"is_ack,omitempty"`
	Status *rpc.Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	AckOrigin string `protobuf:"bytes,6,opt,name=ack_origin,json=ackOrigin,proto3" json:"ack_origin,omitempty"`
	Mid int64 `protobuf:"varint,7,opt,name=mid,proto3" json:"mid,omitempty"`
}

func (x *FrameOption) Reset() {
	*x = FrameOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameOption) ProtoMessage() {}

func (x *FrameOption) ProtoReflect() protoreflect.Message {
	mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameOption.ProtoReflect.Descriptor instead.
func (*FrameOption) Descriptor() ([]byte, []int) {
	return file_pgc_gateway_vega_v1_vega_proto_rawDescGZIP(), []int{2}
}

func (x *FrameOption) GetVegaId() int64 {
	if x != nil {
		return x.VegaId
	}
	return 0
}

func (x *FrameOption) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *FrameOption) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *FrameOption) GetIsAck() bool {
	if x != nil {
		return x.IsAck
	}
	return false
}

func (x *FrameOption) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *FrameOption) GetAckOrigin() string {
	if x != nil {
		return x.AckOrigin
	}
	return ""
}

func (x *FrameOption) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

type HeartbeatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatReq) Reset() {
	*x = HeartbeatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatReq) ProtoMessage() {}

func (x *HeartbeatReq) ProtoReflect() protoreflect.Message {
	mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatReq.ProtoReflect.Descriptor instead.
func (*HeartbeatReq) Descriptor() ([]byte, []int) {
	return file_pgc_gateway_vega_v1_vega_proto_rawDescGZIP(), []int{3}
}

type HeartbeatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatResp) Reset() {
	*x = HeartbeatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResp) ProtoMessage() {}

func (x *HeartbeatResp) ProtoReflect() protoreflect.Message {
	mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResp.ProtoReflect.Descriptor instead.
func (*HeartbeatResp) Descriptor() ([]byte, []int) {
	return file_pgc_gateway_vega_v1_vega_proto_rawDescGZIP(), []int{4}
}

type MessageAckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VegaId string `protobuf:"bytes,1,opt,name=vega_id,json=vegaId,proto3" json:"vega_id,omitempty"`
	ReqId string `protobuf:"bytes,2,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
	Origin string `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	TargetPath string `protobuf:"bytes,4,opt,name=target_path,json=targetPath,proto3" json:"target_path,omitempty"`
}

func (x *MessageAckReq) Reset() {
	*x = MessageAckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAckReq) ProtoMessage() {}

func (x *MessageAckReq) ProtoReflect() protoreflect.Message {
	mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAckReq.ProtoReflect.Descriptor instead.
func (*MessageAckReq) Descriptor() ([]byte, []int) {
	return file_pgc_gateway_vega_v1_vega_proto_rawDescGZIP(), []int{5}
}

func (x *MessageAckReq) GetVegaId() string {
	if x != nil {
		return x.VegaId
	}
	return ""
}

func (x *MessageAckReq) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *MessageAckReq) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *MessageAckReq) GetTargetPath() string {
	if x != nil {
		return x.TargetPath
	}
	return ""
}

type SubscribeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetPaths []*TargetPath `protobuf:"bytes,1,rep,name=target_paths,json=targetPaths,proto3" json:"target_paths,omitempty"`
}

func (x *SubscribeReq) Reset() {
	*x = SubscribeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReq) ProtoMessage() {}

func (x *SubscribeReq) ProtoReflect() protoreflect.Message {
	mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeReq.ProtoReflect.Descriptor instead.
func (*SubscribeReq) Descriptor() ([]byte, []int) {
	return file_pgc_gateway_vega_v1_vega_proto_rawDescGZIP(), []int{6}
}

func (x *SubscribeReq) GetTargetPaths() []*TargetPath {
	if x != nil {
		return x.TargetPaths
	}
	return nil
}

type TargetPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Subs *any1.Any `protobuf:"bytes,2,opt,name=subs,proto3" json:"subs,omitempty"`
}

func (x *TargetPath) Reset() {
	*x = TargetPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetPath) ProtoMessage() {}

func (x *TargetPath) ProtoReflect() protoreflect.Message {
	mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetPath.ProtoReflect.Descriptor instead.
func (*TargetPath) Descriptor() ([]byte, []int) {
	return file_pgc_gateway_vega_v1_vega_proto_rawDescGZIP(), []int{7}
}

func (x *TargetPath) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TargetPath) GetSubs() *any1.Any {
	if x != nil {
		return x.Subs
	}
	return nil
}

type VegaFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options *FrameOption `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	RoutePath string `protobuf:"bytes,2,opt,name=route_path,json=routePath,proto3" json:"route_path,omitempty"`
	Body *any1.Any `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	SubBiz *any1.Any `protobuf:"bytes,4,opt,name=sub_biz,json=subBiz,proto3" json:"sub_biz,omitempty"`
}

func (x *VegaFrame) Reset() {
	*x = VegaFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VegaFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VegaFrame) ProtoMessage() {}

func (x *VegaFrame) ProtoReflect() protoreflect.Message {
	mi := &file_pgc_gateway_vega_v1_vega_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VegaFrame.ProtoReflect.Descriptor instead.
func (*VegaFrame) Descriptor() ([]byte, []int) {
	return file_pgc_gateway_vega_v1_vega_proto_rawDescGZIP(), []int{8}
}

func (x *VegaFrame) GetOptions() *FrameOption {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *VegaFrame) GetRoutePath() string {
	if x != nil {
		return x.RoutePath
	}
	return ""
}

func (x *VegaFrame) GetBody() *any1.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *VegaFrame) GetSubBiz() *any1.Any {
	if x != nil {
		return x.SubBiz
	}
	return nil
}

var File_pgc_gateway_vega_v1_vega_proto protoreflect.FileDescriptor

var file_pgc_gateway_vega_v1_vega_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x67, 0x63, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x65,
	0x67, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x70, 0x67, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65,
	0x67, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x09, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x22, 0x0a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22,
	0xcf, 0x01, 0x0a, 0x0b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x0a, 0x07, 0x76, 0x65, 0x67, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x76, 0x65, 0x67, 0x61, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x69,
	0x73, 0x5f, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x41,
	0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x6b, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x69,
	0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x22, 0x0f, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x78, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x65, 0x67, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x67, 0x61, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65,
	0x71, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x52, 0x0a, 0x0c,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x12, 0x42, 0x0a, 0x0c,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x67, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x73,
	0x22, 0x48, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x28, 0x0a, 0x04, 0x73, 0x75, 0x62, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x73, 0x75, 0x62, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x09, 0x56,
	0x65, 0x67, 0x61, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x67, 0x63, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x2d, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x5f, 0x62, 0x69, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x73, 0x75, 0x62, 0x42, 0x69, 0x7a, 0x32, 0x56, 0x0a, 0x04,
	0x56, 0x65, 0x67, 0x61, 0x12, 0x4e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1e, 0x2e, 0x70, 0x67, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x67, 0x61, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x1a, 0x1e, 0x2e, 0x70, 0x67, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x67, 0x61, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x32, 0xb9, 0x02, 0x0a, 0x0c, 0x56, 0x65, 0x67, 0x61, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x44, 0x6f, 0x63, 0x12, 0x43, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1c, 0x2e,
	0x70, 0x67, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x67, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x70, 0x67,
	0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x52, 0x0a, 0x09, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x67, 0x63, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x70, 0x67, 0x63,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x48,
	0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x22, 0x2e, 0x70,
	0x67, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x67, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58,
	0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x70,
	0x67, 0x63, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pgc_gateway_vega_v1_vega_proto_rawDescOnce sync.Once
	file_pgc_gateway_vega_v1_vega_proto_rawDescData = file_pgc_gateway_vega_v1_vega_proto_rawDesc
)

func file_pgc_gateway_vega_v1_vega_proto_rawDescGZIP() []byte {
	file_pgc_gateway_vega_v1_vega_proto_rawDescOnce.Do(func() {
		file_pgc_gateway_vega_v1_vega_proto_rawDescData = protoimpl.X.CompressGZIP(file_pgc_gateway_vega_v1_vega_proto_rawDescData)
	})
	return file_pgc_gateway_vega_v1_vega_proto_rawDescData
}

var file_pgc_gateway_vega_v1_vega_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pgc_gateway_vega_v1_vega_proto_goTypes = []interface{}{
	(*AuthReq)(nil),       // 0: pgc.gateway.vega.v1.AuthReq
	(*AuthResp)(nil),      // 1: pgc.gateway.vega.v1.AuthResp
	(*FrameOption)(nil),   // 2: pgc.gateway.vega.v1.FrameOption
	(*HeartbeatReq)(nil),  // 3: pgc.gateway.vega.v1.HeartbeatReq
	(*HeartbeatResp)(nil), // 4: pgc.gateway.vega.v1.HeartbeatResp
	(*MessageAckReq)(nil), // 5: pgc.gateway.vega.v1.MessageAckReq
	(*SubscribeReq)(nil),  // 6: pgc.gateway.vega.v1.SubscribeReq
	(*TargetPath)(nil),    // 7: pgc.gateway.vega.v1.TargetPath
	(*VegaFrame)(nil),     // 8: pgc.gateway.vega.v1.VegaFrame
	(*rpc.Status)(nil),    // 9: bilibili.rpc.Status
	(*any1.Any)(nil),      // 10: google.protobuf.Any
	(*empty.Empty)(nil),   // 11: google.protobuf.Empty
}
var file_pgc_gateway_vega_v1_vega_proto_depIdxs = []int32{
	9,  // 0: pgc.gateway.vega.v1.FrameOption.status:type_name -> bilibili.rpc.Status
	7,  // 1: pgc.gateway.vega.v1.SubscribeReq.target_paths:type_name -> pgc.gateway.vega.v1.TargetPath
	10, // 2: pgc.gateway.vega.v1.TargetPath.subs:type_name -> google.protobuf.Any
	2,  // 3: pgc.gateway.vega.v1.VegaFrame.options:type_name -> pgc.gateway.vega.v1.FrameOption
	10, // 4: pgc.gateway.vega.v1.VegaFrame.body:type_name -> google.protobuf.Any
	10, // 5: pgc.gateway.vega.v1.VegaFrame.sub_biz:type_name -> google.protobuf.Any
	8,  // 6: pgc.gateway.vega.v1.Vega.CreateTunnel:input_type -> pgc.gateway.vega.v1.VegaFrame
	0,  // 7: pgc.gateway.vega.v1.VegaFrameDoc.Auth:input_type -> pgc.gateway.vega.v1.AuthReq
	3,  // 8: pgc.gateway.vega.v1.VegaFrameDoc.Heartbeat:input_type -> pgc.gateway.vega.v1.HeartbeatReq
	5,  // 9: pgc.gateway.vega.v1.VegaFrameDoc.MessageAck:input_type -> pgc.gateway.vega.v1.MessageAckReq
	6,  // 10: pgc.gateway.vega.v1.VegaFrameDoc.Subscribe:input_type -> pgc.gateway.vega.v1.SubscribeReq
	8,  // 11: pgc.gateway.vega.v1.Vega.CreateTunnel:output_type -> pgc.gateway.vega.v1.VegaFrame
	1,  // 12: pgc.gateway.vega.v1.VegaFrameDoc.Auth:output_type -> pgc.gateway.vega.v1.AuthResp
	4,  // 13: pgc.gateway.vega.v1.VegaFrameDoc.Heartbeat:output_type -> pgc.gateway.vega.v1.HeartbeatResp
	11, // 14: pgc.gateway.vega.v1.VegaFrameDoc.MessageAck:output_type -> google.protobuf.Empty
	11, // 15: pgc.gateway.vega.v1.VegaFrameDoc.Subscribe:output_type -> google.protobuf.Empty
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pgc_gateway_vega_v1_vega_proto_init() }
func file_pgc_gateway_vega_v1_vega_proto_init() {
	if File_pgc_gateway_vega_v1_vega_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pgc_gateway_vega_v1_vega_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReq); i {
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
		file_pgc_gateway_vega_v1_vega_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResp); i {
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
		file_pgc_gateway_vega_v1_vega_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrameOption); i {
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
		file_pgc_gateway_vega_v1_vega_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatReq); i {
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
		file_pgc_gateway_vega_v1_vega_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResp); i {
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
		file_pgc_gateway_vega_v1_vega_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAckReq); i {
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
		file_pgc_gateway_vega_v1_vega_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeReq); i {
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
		file_pgc_gateway_vega_v1_vega_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetPath); i {
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
		file_pgc_gateway_vega_v1_vega_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VegaFrame); i {
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
			RawDescriptor: file_pgc_gateway_vega_v1_vega_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pgc_gateway_vega_v1_vega_proto_goTypes,
		DependencyIndexes: file_pgc_gateway_vega_v1_vega_proto_depIdxs,
		MessageInfos:      file_pgc_gateway_vega_v1_vega_proto_msgTypes,
	}.Build()
	File_pgc_gateway_vega_v1_vega_proto = out.File
	file_pgc_gateway_vega_v1_vega_proto_rawDesc = nil
	file_pgc_gateway_vega_v1_vega_proto_goTypes = nil
	file_pgc_gateway_vega_v1_vega_proto_depIdxs = nil
}
