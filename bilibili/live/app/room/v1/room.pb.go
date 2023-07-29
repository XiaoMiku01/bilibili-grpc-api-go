// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: bilibili/live/app/room/v1/room.proto

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

type GetStudioListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *GetStudioListReq) Reset() {
	*x = GetStudioListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_live_app_room_v1_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudioListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudioListReq) ProtoMessage() {}

func (x *GetStudioListReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_live_app_room_v1_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudioListReq.ProtoReflect.Descriptor instead.
func (*GetStudioListReq) Descriptor() ([]byte, []int) {
	return file_bilibili_live_app_room_v1_room_proto_rawDescGZIP(), []int{0}
}

func (x *GetStudioListReq) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type GetStudioListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	MasterList []*GetStudioListResp_StudioMaster `protobuf:"bytes,2,rep,name=master_list,json=masterList,proto3" json:"master_list,omitempty"`
}

func (x *GetStudioListResp) Reset() {
	*x = GetStudioListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_live_app_room_v1_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudioListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudioListResp) ProtoMessage() {}

func (x *GetStudioListResp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_live_app_room_v1_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudioListResp.ProtoReflect.Descriptor instead.
func (*GetStudioListResp) Descriptor() ([]byte, []int) {
	return file_bilibili_live_app_room_v1_room_proto_rawDescGZIP(), []int{1}
}

func (x *GetStudioListResp) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetStudioListResp) GetMasterList() []*GetStudioListResp_StudioMaster {
	if x != nil {
		return x.MasterList
	}
	return nil
}

type GetStudioListResp_Pendants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frame *GetStudioListResp_Pendant `protobuf:"bytes,1,opt,name=frame,proto3" json:"frame,omitempty"`
	Badge *GetStudioListResp_Pendant `protobuf:"bytes,2,opt,name=badge,proto3" json:"badge,omitempty"`
}

func (x *GetStudioListResp_Pendants) Reset() {
	*x = GetStudioListResp_Pendants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_live_app_room_v1_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudioListResp_Pendants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudioListResp_Pendants) ProtoMessage() {}

func (x *GetStudioListResp_Pendants) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_live_app_room_v1_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudioListResp_Pendants.ProtoReflect.Descriptor instead.
func (*GetStudioListResp_Pendants) Descriptor() ([]byte, []int) {
	return file_bilibili_live_app_room_v1_room_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetStudioListResp_Pendants) GetFrame() *GetStudioListResp_Pendant {
	if x != nil {
		return x.Frame
	}
	return nil
}

func (x *GetStudioListResp_Pendants) GetBadge() *GetStudioListResp_Pendant {
	if x != nil {
		return x.Badge
	}
	return nil
}

type GetStudioListResp_Pendant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Position int64 `protobuf:"varint,2,opt,name=position,proto3" json:"position,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Desc string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *GetStudioListResp_Pendant) Reset() {
	*x = GetStudioListResp_Pendant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_live_app_room_v1_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudioListResp_Pendant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudioListResp_Pendant) ProtoMessage() {}

func (x *GetStudioListResp_Pendant) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_live_app_room_v1_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudioListResp_Pendant.ProtoReflect.Descriptor instead.
func (*GetStudioListResp_Pendant) Descriptor() ([]byte, []int) {
	return file_bilibili_live_app_room_v1_room_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetStudioListResp_Pendant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetStudioListResp_Pendant) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *GetStudioListResp_Pendant) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetStudioListResp_Pendant) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type GetStudioListResp_StudioMaster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	RoomId int64 `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Uname string `protobuf:"bytes,3,opt,name=uname,proto3" json:"uname,omitempty"`
	Face string `protobuf:"bytes,4,opt,name=face,proto3" json:"face,omitempty"`
	Pendants *GetStudioListResp_Pendants `protobuf:"bytes,5,opt,name=pendants,proto3" json:"pendants,omitempty"`
	Tag string `protobuf:"bytes,6,opt,name=tag,proto3" json:"tag,omitempty"`
	TagType int64 `protobuf:"varint,7,opt,name=tag_type,json=tagType,proto3" json:"tag_type,omitempty"`
}

func (x *GetStudioListResp_StudioMaster) Reset() {
	*x = GetStudioListResp_StudioMaster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_live_app_room_v1_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudioListResp_StudioMaster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudioListResp_StudioMaster) ProtoMessage() {}

func (x *GetStudioListResp_StudioMaster) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_live_app_room_v1_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudioListResp_StudioMaster.ProtoReflect.Descriptor instead.
func (*GetStudioListResp_StudioMaster) Descriptor() ([]byte, []int) {
	return file_bilibili_live_app_room_v1_room_proto_rawDescGZIP(), []int{1, 2}
}

func (x *GetStudioListResp_StudioMaster) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetStudioListResp_StudioMaster) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *GetStudioListResp_StudioMaster) GetUname() string {
	if x != nil {
		return x.Uname
	}
	return ""
}

func (x *GetStudioListResp_StudioMaster) GetFace() string {
	if x != nil {
		return x.Face
	}
	return ""
}

func (x *GetStudioListResp_StudioMaster) GetPendants() *GetStudioListResp_Pendants {
	if x != nil {
		return x.Pendants
	}
	return nil
}

func (x *GetStudioListResp_StudioMaster) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *GetStudioListResp_StudioMaster) GetTagType() int64 {
	if x != nil {
		return x.TagType
	}
	return 0
}

var File_bilibili_live_app_room_v1_room_proto protoreflect.FileDescriptor

var file_bilibili_live_app_room_v1_room_proto_rawDesc = []byte{
	0x0a, 0x24, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76,
	0x31, 0x22, 0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0xf7,
	0x04, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5a, 0x0a, 0x0b,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e,
	0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0xa2, 0x01, 0x0a, 0x08, 0x50, 0x65, 0x6e,
	0x64, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x2e, 0x50, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x12, 0x4a, 0x0a, 0x05, 0x62, 0x61, 0x64, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x50,
	0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x62, 0x61, 0x64, 0x67, 0x65, 0x1a, 0x63, 0x0a,
	0x07, 0x50, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x1a, 0xe3, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x75, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x61, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x70, 0x65, 0x6e, 0x64,
	0x61, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72,
	0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x50, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74,
	0x73, 0x52, 0x08, 0x70, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x61, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x74, 0x61, 0x67, 0x54, 0x79, 0x70, 0x65, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30,
	0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f,
	0x6c, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_live_app_room_v1_room_proto_rawDescOnce sync.Once
	file_bilibili_live_app_room_v1_room_proto_rawDescData = file_bilibili_live_app_room_v1_room_proto_rawDesc
)

func file_bilibili_live_app_room_v1_room_proto_rawDescGZIP() []byte {
	file_bilibili_live_app_room_v1_room_proto_rawDescOnce.Do(func() {
		file_bilibili_live_app_room_v1_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_live_app_room_v1_room_proto_rawDescData)
	})
	return file_bilibili_live_app_room_v1_room_proto_rawDescData
}

var file_bilibili_live_app_room_v1_room_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bilibili_live_app_room_v1_room_proto_goTypes = []interface{}{
	(*GetStudioListReq)(nil),               // 0: bilibili.live.app.room.v1.GetStudioListReq
	(*GetStudioListResp)(nil),              // 1: bilibili.live.app.room.v1.GetStudioListResp
	(*GetStudioListResp_Pendants)(nil),     // 2: bilibili.live.app.room.v1.GetStudioListResp.Pendants
	(*GetStudioListResp_Pendant)(nil),      // 3: bilibili.live.app.room.v1.GetStudioListResp.Pendant
	(*GetStudioListResp_StudioMaster)(nil), // 4: bilibili.live.app.room.v1.GetStudioListResp.StudioMaster
}
var file_bilibili_live_app_room_v1_room_proto_depIdxs = []int32{
	4, // 0: bilibili.live.app.room.v1.GetStudioListResp.master_list:type_name -> bilibili.live.app.room.v1.GetStudioListResp.StudioMaster
	3, // 1: bilibili.live.app.room.v1.GetStudioListResp.Pendants.frame:type_name -> bilibili.live.app.room.v1.GetStudioListResp.Pendant
	3, // 2: bilibili.live.app.room.v1.GetStudioListResp.Pendants.badge:type_name -> bilibili.live.app.room.v1.GetStudioListResp.Pendant
	2, // 3: bilibili.live.app.room.v1.GetStudioListResp.StudioMaster.pendants:type_name -> bilibili.live.app.room.v1.GetStudioListResp.Pendants
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bilibili_live_app_room_v1_room_proto_init() }
func file_bilibili_live_app_room_v1_room_proto_init() {
	if File_bilibili_live_app_room_v1_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_live_app_room_v1_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudioListReq); i {
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
		file_bilibili_live_app_room_v1_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudioListResp); i {
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
		file_bilibili_live_app_room_v1_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudioListResp_Pendants); i {
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
		file_bilibili_live_app_room_v1_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudioListResp_Pendant); i {
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
		file_bilibili_live_app_room_v1_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudioListResp_StudioMaster); i {
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
			RawDescriptor: file_bilibili_live_app_room_v1_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_live_app_room_v1_room_proto_goTypes,
		DependencyIndexes: file_bilibili_live_app_room_v1_room_proto_depIdxs,
		MessageInfos:      file_bilibili_live_app_room_v1_room_proto_msgTypes,
	}.Build()
	File_bilibili_live_app_room_v1_room_proto = out.File
	file_bilibili_live_app_room_v1_room_proto_rawDesc = nil
	file_bilibili_live_app_room_v1_room_proto_goTypes = nil
	file_bilibili_live_app_room_v1_room_proto_depIdxs = nil
}
