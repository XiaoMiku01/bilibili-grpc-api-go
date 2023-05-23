// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: bilibili/app/resource/v1/module.proto

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

type CompressType int32

const (
	CompressType_Unzip    CompressType = 0 // unzip
	CompressType_Original CompressType = 1 // 不操作
)

// Enum value maps for CompressType.
var (
	CompressType_name = map[int32]string{
		0: "Unzip",
		1: "Original",
	}
	CompressType_value = map[string]int32{
		"Unzip":    0,
		"Original": 1,
	}
)

func (x CompressType) Enum() *CompressType {
	p := new(CompressType)
	*p = x
	return p
}

func (x CompressType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompressType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_app_resource_v1_module_proto_enumTypes[0].Descriptor()
}

func (CompressType) Type() protoreflect.EnumType {
	return &file_bilibili_app_resource_v1_module_proto_enumTypes[0]
}

func (x CompressType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompressType.Descriptor instead.
func (CompressType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_app_resource_v1_module_proto_rawDescGZIP(), []int{0}
}

type EnvType int32

const (
	EnvType_Unknown EnvType = 0 //
	EnvType_Release EnvType = 1 //
	EnvType_Test    EnvType = 2 //
)

// Enum value maps for EnvType.
var (
	EnvType_name = map[int32]string{
		0: "Unknown",
		1: "Release",
		2: "Test",
	}
	EnvType_value = map[string]int32{
		"Unknown": 0,
		"Release": 1,
		"Test":    2,
	}
)

func (x EnvType) Enum() *EnvType {
	p := new(EnvType)
	*p = x
	return p
}

func (x EnvType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnvType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_app_resource_v1_module_proto_enumTypes[1].Descriptor()
}

func (EnvType) Type() protoreflect.EnumType {
	return &file_bilibili_app_resource_v1_module_proto_enumTypes[1]
}

func (x EnvType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnvType.Descriptor instead.
func (EnvType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_app_resource_v1_module_proto_rawDescGZIP(), []int{1}
}

type IncrementType int32

const (
	IncrementType_Total       IncrementType = 0 // 全量包
	IncrementType_Incremental IncrementType = 1 // 增量包
)

// Enum value maps for IncrementType.
var (
	IncrementType_name = map[int32]string{
		0: "Total",
		1: "Incremental",
	}
	IncrementType_value = map[string]int32{
		"Total":       0,
		"Incremental": 1,
	}
)

func (x IncrementType) Enum() *IncrementType {
	p := new(IncrementType)
	*p = x
	return p
}

func (x IncrementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IncrementType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_app_resource_v1_module_proto_enumTypes[2].Descriptor()
}

func (IncrementType) Type() protoreflect.EnumType {
	return &file_bilibili_app_resource_v1_module_proto_enumTypes[2]
}

func (x IncrementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IncrementType.Descriptor instead.
func (IncrementType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_app_resource_v1_module_proto_rawDescGZIP(), []int{2}
}

type LevelType int32

const (
	LevelType_Undefined LevelType = 0 //
	LevelType_High      LevelType = 1 // 高 需立即下载
	LevelType_Middle    LevelType = 2 // 中 可以延迟下载
	LevelType_Low       LevelType = 3 // 低 仅在业务方使用到时由业务方手动进行下载
)

// Enum value maps for LevelType.
var (
	LevelType_name = map[int32]string{
		0: "Undefined",
		1: "High",
		2: "Middle",
		3: "Low",
	}
	LevelType_value = map[string]int32{
		"Undefined": 0,
		"High":      1,
		"Middle":    2,
		"Low":       3,
	}
)

func (x LevelType) Enum() *LevelType {
	p := new(LevelType)
	*p = x
	return p
}

func (x LevelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LevelType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_app_resource_v1_module_proto_enumTypes[3].Descriptor()
}

func (LevelType) Type() protoreflect.EnumType {
	return &file_bilibili_app_resource_v1_module_proto_enumTypes[3]
}

func (x LevelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LevelType.Descriptor instead.
func (LevelType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_app_resource_v1_module_proto_rawDescGZIP(), []int{3}
}

type ListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Env string `protobuf:"bytes,1,opt,name=env,proto3" json:"env,omitempty"`
	Pools []*PoolReply `protobuf:"bytes,2,rep,name=pools,proto3" json:"pools,omitempty"`
	ListVersion int64 `protobuf:"varint,3,opt,name=list_version,json=listVersion,proto3" json:"list_version,omitempty"`
}

func (x *ListReply) Reset() {
	*x = ListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReply) ProtoMessage() {}

func (x *ListReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReply.ProtoReflect.Descriptor instead.
func (*ListReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_resource_v1_module_proto_rawDescGZIP(), []int{0}
}

func (x *ListReply) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *ListReply) GetPools() []*PoolReply {
	if x != nil {
		return x.Pools
	}
	return nil
}

func (x *ListReply) GetListVersion() int64 {
	if x != nil {
		return x.ListVersion
	}
	return 0
}

type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolName string `protobuf:"bytes,1,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	ModuleName string `protobuf:"bytes,2,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	VersionList []*VersionListReq `protobuf:"bytes,3,rep,name=version_list,json=versionList,proto3" json:"version_list,omitempty"`
	Env EnvType `protobuf:"varint,4,opt,name=env,proto3,enum=bilibili.app.resource.v1.EnvType" json:"env,omitempty"`
	SysVer int32 `protobuf:"varint,5,opt,name=sys_ver,json=sysVer,proto3" json:"sys_ver,omitempty"`
	Scale int32 `protobuf:"varint,6,opt,name=scale,proto3" json:"scale,omitempty"`
	Arch int32 `protobuf:"varint,7,opt,name=arch,proto3" json:"arch,omitempty"`
	ListVersion int64 `protobuf:"varint,8,opt,name=list_version,json=listVersion,proto3" json:"list_version,omitempty"`
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_resource_v1_module_proto_rawDescGZIP(), []int{1}
}

func (x *ListReq) GetPoolName() string {
	if x != nil {
		return x.PoolName
	}
	return ""
}

func (x *ListReq) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *ListReq) GetVersionList() []*VersionListReq {
	if x != nil {
		return x.VersionList
	}
	return nil
}

func (x *ListReq) GetEnv() EnvType {
	if x != nil {
		return x.Env
	}
	return EnvType_Unknown
}

func (x *ListReq) GetSysVer() int32 {
	if x != nil {
		return x.SysVer
	}
	return 0
}

func (x *ListReq) GetScale() int32 {
	if x != nil {
		return x.Scale
	}
	return 0
}

func (x *ListReq) GetArch() int32 {
	if x != nil {
		return x.Arch
	}
	return 0
}

func (x *ListReq) GetListVersion() int64 {
	if x != nil {
		return x.ListVersion
	}
	return 0
}

type ModuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version int64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Md5 string `protobuf:"bytes,4,opt,name=md5,proto3" json:"md5,omitempty"`
	TotalMd5 string `protobuf:"bytes,5,opt,name=total_md5,json=totalMd5,proto3" json:"total_md5,omitempty"`
	Increment IncrementType `protobuf:"varint,6,opt,name=increment,proto3,enum=bilibili.app.resource.v1.IncrementType" json:"increment,omitempty"`
	IsWifi bool `protobuf:"varint,7,opt,name=is_wifi,json=isWifi,proto3" json:"is_wifi,omitempty"`
	Level LevelType `protobuf:"varint,8,opt,name=level,proto3,enum=bilibili.app.resource.v1.LevelType" json:"level,omitempty"`
	Filename string `protobuf:"bytes,9,opt,name=filename,proto3" json:"filename,omitempty"`
	FileType string `protobuf:"bytes,10,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	FileSize int64 `protobuf:"varint,11,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	Compress CompressType `protobuf:"varint,12,opt,name=compress,proto3,enum=bilibili.app.resource.v1.CompressType" json:"compress,omitempty"`
	PublishTime int64 `protobuf:"varint,13,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	// 上报使用
	PoolId int64 `protobuf:"varint,14,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	ModuleId int64 `protobuf:"varint,15,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	VersionId int64 `protobuf:"varint,16,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	FileId int64 `protobuf:"varint,17,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	ZipCheck bool `protobuf:"varint,18,opt,name=zip_check,json=zipCheck,proto3" json:"zip_check,omitempty"`
}

func (x *ModuleReply) Reset() {
	*x = ModuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleReply) ProtoMessage() {}

func (x *ModuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleReply.ProtoReflect.Descriptor instead.
func (*ModuleReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_resource_v1_module_proto_rawDescGZIP(), []int{2}
}

func (x *ModuleReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModuleReply) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ModuleReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ModuleReply) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *ModuleReply) GetTotalMd5() string {
	if x != nil {
		return x.TotalMd5
	}
	return ""
}

func (x *ModuleReply) GetIncrement() IncrementType {
	if x != nil {
		return x.Increment
	}
	return IncrementType_Total
}

func (x *ModuleReply) GetIsWifi() bool {
	if x != nil {
		return x.IsWifi
	}
	return false
}

func (x *ModuleReply) GetLevel() LevelType {
	if x != nil {
		return x.Level
	}
	return LevelType_Undefined
}

func (x *ModuleReply) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ModuleReply) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *ModuleReply) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *ModuleReply) GetCompress() CompressType {
	if x != nil {
		return x.Compress
	}
	return CompressType_Unzip
}

func (x *ModuleReply) GetPublishTime() int64 {
	if x != nil {
		return x.PublishTime
	}
	return 0
}

func (x *ModuleReply) GetPoolId() int64 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *ModuleReply) GetModuleId() int64 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *ModuleReply) GetVersionId() int64 {
	if x != nil {
		return x.VersionId
	}
	return 0
}

func (x *ModuleReply) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *ModuleReply) GetZipCheck() bool {
	if x != nil {
		return x.ZipCheck
	}
	return false
}

type PoolReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Modules []*ModuleReply `protobuf:"bytes,2,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *PoolReply) Reset() {
	*x = PoolReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoolReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoolReply) ProtoMessage() {}

func (x *PoolReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoolReply.ProtoReflect.Descriptor instead.
func (*PoolReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_resource_v1_module_proto_rawDescGZIP(), []int{3}
}

func (x *PoolReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PoolReply) GetModules() []*ModuleReply {
	if x != nil {
		return x.Modules
	}
	return nil
}

type VersionListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolName string `protobuf:"bytes,1,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	Versions []*VersionReq `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *VersionListReq) Reset() {
	*x = VersionListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionListReq) ProtoMessage() {}

func (x *VersionListReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionListReq.ProtoReflect.Descriptor instead.
func (*VersionListReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_resource_v1_module_proto_rawDescGZIP(), []int{4}
}

func (x *VersionListReq) GetPoolName() string {
	if x != nil {
		return x.PoolName
	}
	return ""
}

func (x *VersionListReq) GetVersions() []*VersionReq {
	if x != nil {
		return x.Versions
	}
	return nil
}

type VersionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	Version int64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionReq) Reset() {
	*x = VersionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionReq) ProtoMessage() {}

func (x *VersionReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_resource_v1_module_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionReq.ProtoReflect.Descriptor instead.
func (*VersionReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_resource_v1_module_proto_rawDescGZIP(), []int{5}
}

func (x *VersionReq) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *VersionReq) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_bilibili_app_resource_v1_module_proto protoreflect.FileDescriptor

var file_bilibili_app_resource_v1_module_proto_rawDesc = []byte{
	0x0a, 0x25, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0x7b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76,
	0x12, 0x39, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xaf,
	0x02, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f,
	0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x79,
	0x73, 0x5f, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x79, 0x73,
	0x56, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x21, 0x0a,
	0x0c, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0xdf, 0x04, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x64, 0x35, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x64, 0x35, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x64, 0x35, 0x12,
	0x45, 0x0a, 0x09, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x77, 0x69, 0x66,
	0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x57, 0x69, 0x66, 0x69, 0x12,
	0x39, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x42, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x26, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x7a, 0x69, 0x70, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x22, 0x60, 0x0a, 0x09, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x73, 0x22, 0x6f, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x52, 0x08, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x47, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x27,
	0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x55, 0x6e, 0x7a, 0x69, 0x70, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x10, 0x01, 0x2a, 0x2d, 0x0a, 0x07, 0x45, 0x6e, 0x76, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x54, 0x65, 0x73, 0x74, 0x10, 0x02, 0x2a, 0x2b, 0x0a, 0x0d, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x10, 0x01, 0x2a, 0x39, 0x0a, 0x09, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0d, 0x0a, 0x09, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x69, 0x67, 0x68, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x6f, 0x77, 0x10, 0x03, 0x32, 0x58,
	0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x21, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30,
	0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_resource_v1_module_proto_rawDescOnce sync.Once
	file_bilibili_app_resource_v1_module_proto_rawDescData = file_bilibili_app_resource_v1_module_proto_rawDesc
)

func file_bilibili_app_resource_v1_module_proto_rawDescGZIP() []byte {
	file_bilibili_app_resource_v1_module_proto_rawDescOnce.Do(func() {
		file_bilibili_app_resource_v1_module_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_resource_v1_module_proto_rawDescData)
	})
	return file_bilibili_app_resource_v1_module_proto_rawDescData
}

var file_bilibili_app_resource_v1_module_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_bilibili_app_resource_v1_module_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bilibili_app_resource_v1_module_proto_goTypes = []interface{}{
	(CompressType)(0),      // 0: bilibili.app.resource.v1.CompressType
	(EnvType)(0),           // 1: bilibili.app.resource.v1.EnvType
	(IncrementType)(0),     // 2: bilibili.app.resource.v1.IncrementType
	(LevelType)(0),         // 3: bilibili.app.resource.v1.LevelType
	(*ListReply)(nil),      // 4: bilibili.app.resource.v1.ListReply
	(*ListReq)(nil),        // 5: bilibili.app.resource.v1.ListReq
	(*ModuleReply)(nil),    // 6: bilibili.app.resource.v1.ModuleReply
	(*PoolReply)(nil),      // 7: bilibili.app.resource.v1.PoolReply
	(*VersionListReq)(nil), // 8: bilibili.app.resource.v1.VersionListReq
	(*VersionReq)(nil),     // 9: bilibili.app.resource.v1.VersionReq
}
var file_bilibili_app_resource_v1_module_proto_depIdxs = []int32{
	7, // 0: bilibili.app.resource.v1.ListReply.pools:type_name -> bilibili.app.resource.v1.PoolReply
	8, // 1: bilibili.app.resource.v1.ListReq.version_list:type_name -> bilibili.app.resource.v1.VersionListReq
	1, // 2: bilibili.app.resource.v1.ListReq.env:type_name -> bilibili.app.resource.v1.EnvType
	2, // 3: bilibili.app.resource.v1.ModuleReply.increment:type_name -> bilibili.app.resource.v1.IncrementType
	3, // 4: bilibili.app.resource.v1.ModuleReply.level:type_name -> bilibili.app.resource.v1.LevelType
	0, // 5: bilibili.app.resource.v1.ModuleReply.compress:type_name -> bilibili.app.resource.v1.CompressType
	6, // 6: bilibili.app.resource.v1.PoolReply.modules:type_name -> bilibili.app.resource.v1.ModuleReply
	9, // 7: bilibili.app.resource.v1.VersionListReq.versions:type_name -> bilibili.app.resource.v1.VersionReq
	5, // 8: bilibili.app.resource.v1.Module.List:input_type -> bilibili.app.resource.v1.ListReq
	4, // 9: bilibili.app.resource.v1.Module.List:output_type -> bilibili.app.resource.v1.ListReply
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_bilibili_app_resource_v1_module_proto_init() }
func file_bilibili_app_resource_v1_module_proto_init() {
	if File_bilibili_app_resource_v1_module_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_resource_v1_module_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReply); i {
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
		file_bilibili_app_resource_v1_module_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
		file_bilibili_app_resource_v1_module_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleReply); i {
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
		file_bilibili_app_resource_v1_module_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoolReply); i {
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
		file_bilibili_app_resource_v1_module_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionListReq); i {
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
		file_bilibili_app_resource_v1_module_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionReq); i {
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
			RawDescriptor: file_bilibili_app_resource_v1_module_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_app_resource_v1_module_proto_goTypes,
		DependencyIndexes: file_bilibili_app_resource_v1_module_proto_depIdxs,
		EnumInfos:         file_bilibili_app_resource_v1_module_proto_enumTypes,
		MessageInfos:      file_bilibili_app_resource_v1_module_proto_msgTypes,
	}.Build()
	File_bilibili_app_resource_v1_module_proto = out.File
	file_bilibili_app_resource_v1_module_proto_rawDesc = nil
	file_bilibili_app_resource_v1_module_proto_goTypes = nil
	file_bilibili_app_resource_v1_module_proto_depIdxs = nil
}
