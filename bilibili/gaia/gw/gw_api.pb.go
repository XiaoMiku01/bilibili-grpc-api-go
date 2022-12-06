// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bilibili/gaia/gw/gw_api.proto

package gw

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

// 加密方式
type EncryptType int32

const (
	EncryptType_INVALID_ENCRYPT_TYPE EncryptType = 0 // 非法值
	EncryptType_CLIENT_AES           EncryptType = 1 // 同客户端人工约定AES加密私钥，存储在客户端
	EncryptType_SERVER_RSA_AES       EncryptType = 2 // 客户端随机生成一个用于AES加密的私钥，并用服务端下发的RSA公钥来加密
)

// Enum value maps for EncryptType.
var (
	EncryptType_name = map[int32]string{
		0: "INVALID_ENCRYPT_TYPE",
		1: "CLIENT_AES",
		2: "SERVER_RSA_AES",
	}
	EncryptType_value = map[string]int32{
		"INVALID_ENCRYPT_TYPE": 0,
		"CLIENT_AES":           1,
		"SERVER_RSA_AES":       2,
	}
)

func (x EncryptType) Enum() *EncryptType {
	p := new(EncryptType)
	*p = x
	return p
}

func (x EncryptType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncryptType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_gaia_gw_gw_api_proto_enumTypes[0].Descriptor()
}

func (EncryptType) Type() protoreflect.EnumType {
	return &file_bilibili_gaia_gw_gw_api_proto_enumTypes[0]
}

func (x EncryptType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncryptType.Descriptor instead.
func (EncryptType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_gaia_gw_gw_api_proto_rawDescGZIP(), []int{0}
}

// 负载类型
type PayloadType int32

const (
	PayloadType_INVALID_PAYLOAD PayloadType = 0 //非法值
	PayloadType_DEVICE_APP_LIST PayloadType = 1 //设备app列表，对应DeviceAppList
)

// Enum value maps for PayloadType.
var (
	PayloadType_name = map[int32]string{
		0: "INVALID_PAYLOAD",
		1: "DEVICE_APP_LIST",
	}
	PayloadType_value = map[string]int32{
		"INVALID_PAYLOAD": 0,
		"DEVICE_APP_LIST": 1,
	}
)

func (x PayloadType) Enum() *PayloadType {
	p := new(PayloadType)
	*p = x
	return p
}

func (x PayloadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PayloadType) Descriptor() protoreflect.EnumDescriptor {
	return file_bilibili_gaia_gw_gw_api_proto_enumTypes[1].Descriptor()
}

func (PayloadType) Type() protoreflect.EnumType {
	return &file_bilibili_gaia_gw_gw_api_proto_enumTypes[1]
}

func (x PayloadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PayloadType.Descriptor instead.
func (PayloadType) EnumDescriptor() ([]byte, []int) {
	return file_bilibili_gaia_gw_gw_api_proto_rawDescGZIP(), []int{1}
}

// 待加密的pb对象
type DeviceAppList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 上报类型
	// first_installation:首次安装上报 first_open:每日启动上报
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// 安装的系统程序列表
	SystemAppList []string `protobuf:"bytes,2,rep,name=system_app_list,json=systemAppList,proto3" json:"system_app_list,omitempty"`
	// 安装的用户程序列表
	UserAppList []string `protobuf:"bytes,3,rep,name=user_app_list,json=userAppList,proto3" json:"user_app_list,omitempty"`
}

func (x *DeviceAppList) Reset() {
	*x = DeviceAppList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceAppList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceAppList) ProtoMessage() {}

func (x *DeviceAppList) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceAppList.ProtoReflect.Descriptor instead.
func (*DeviceAppList) Descriptor() ([]byte, []int) {
	return file_bilibili_gaia_gw_gw_api_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceAppList) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *DeviceAppList) GetSystemAppList() []string {
	if x != nil {
		return x.SystemAppList
	}
	return nil
}

func (x *DeviceAppList) GetUserAppList() []string {
	if x != nil {
		return x.UserAppList
	}
	return nil
}

type FetchPublicKeyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 版本号
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// RSA公钥
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// 公钥过期时间
	Deadline int64 `protobuf:"varint,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (x *FetchPublicKeyReply) Reset() {
	*x = FetchPublicKeyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchPublicKeyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPublicKeyReply) ProtoMessage() {}

func (x *FetchPublicKeyReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPublicKeyReply.ProtoReflect.Descriptor instead.
func (*FetchPublicKeyReply) Descriptor() ([]byte, []int) {
	return file_bilibili_gaia_gw_gw_api_proto_rawDescGZIP(), []int{1}
}

func (x *FetchPublicKeyReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *FetchPublicKeyReply) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *FetchPublicKeyReply) GetDeadline() int64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

type GaiaDeviceBasicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 平台&应用信息
	Platform string `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`              //android/ios/web/h5;
	Device   string `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`                  //运行设备, 用于区分不同的app, 见客户端传入的对应参数。对于苹果系统，device有效值为phone, pad；安卓无法区分phone和pad，留空即可。
	MobiApp  string `protobuf:"bytes,3,opt,name=mobi_app,json=mobiApp,proto3" json:"mobi_app,omitempty"` //包类型，用于区分不同的app, 见客户端传入的对应参数（mobi_app ）；对于web端请求，请传空
	Origin   string `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`                  //客户端appkey, 用以区分不同的客户端，对应客户端请求参数中的appkey,如果无法获取可传空“”
	AppId    string `protobuf:"bytes,5,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`       //app产品编号 //产品编号，由数据平台分配，粉=1，白=2，蓝=3，直播姬=4，HD=5，海外=6，OTT=7，漫画=8，TV野版=9，小视频=10，网易漫画=11，网易漫画lite=12，网易漫画HD=13,国际版=14
	// 应用的版本信息
	Sdkver         string `protobuf:"bytes,6,opt,name=sdkver,proto3" json:"sdkver,omitempty"`                                         // SDK版本号   "sdkver": "2.6.6"
	AppVersion     string `protobuf:"bytes,7,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`               // app版本  "app_version":"5.36.0"
	AppVersionCode string `protobuf:"bytes,8,opt,name=app_version_code,json=appVersionCode,proto3" json:"app_version_code,omitempty"` // app版本号 "app_version_code":"5360000"
	Build          string `protobuf:"bytes,9,opt,name=build,proto3" json:"build,omitempty"`                                           // app版本号，见客户端传入的对应参数；对于web端请求，请传空
	// 渠道信息
	Channel string `protobuf:"bytes,10,opt,name=channel,proto3" json:"channel,omitempty"` //渠道标识，见客户端传入的对应参数；对于web端请求，请传空；对应chid
	// 机器硬件信息
	Brand     string `protobuf:"bytes,11,opt,name=brand,proto3" json:"brand,omitempty"` //手机品牌，见客户端传入的对应参数；
	Model     string `protobuf:"bytes,12,opt,name=model,proto3" json:"model,omitempty"` //手机型号，见客户端传入的对应参数
	Osver     string `protobuf:"bytes,13,opt,name=osver,proto3" json:"osver,omitempty"` //系统版本，见客户端传入的对应参数
	UserAgent string `protobuf:"bytes,14,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// 设备标识信息
	BuvidLocal string `protobuf:"bytes,15,opt,name=buvid_local,json=buvidLocal,proto3" json:"buvid_local,omitempty"` //本地设备唯一标识
	Buvid      string `protobuf:"bytes,16,opt,name=buvid,proto3" json:"buvid,omitempty"`                             //设备唯一标识
	// 登陆用户信息
	Mid string `protobuf:"bytes,17,opt,name=mid,proto3" json:"mid,omitempty"` //最后一次登陆用户的mid，如果无登陆信息，传0即可
	// 本次启动信息
	Fts   int64 `protobuf:"varint,18,opt,name=fts,proto3" json:"fts,omitempty"`     // app首次启动时间 "fts":1530447775661
	First int32 `protobuf:"varint,19,opt,name=first,proto3" json:"first,omitempty"` // 是否首次启动 是：0 否：1
	// 网络相关的信息
	Network string `protobuf:"bytes,20,opt,name=network,proto3" json:"network,omitempty"` // 网络连接方式, WIFI/CELLULAR/OFFLINE/OTHERNET/ETHERNET "network":"WIFI", ESS_NETWORK_STATE、ACCESS_WIFI_STATE
}

func (x *GaiaDeviceBasicInfo) Reset() {
	*x = GaiaDeviceBasicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaiaDeviceBasicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaiaDeviceBasicInfo) ProtoMessage() {}

func (x *GaiaDeviceBasicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaiaDeviceBasicInfo.ProtoReflect.Descriptor instead.
func (*GaiaDeviceBasicInfo) Descriptor() ([]byte, []int) {
	return file_bilibili_gaia_gw_gw_api_proto_rawDescGZIP(), []int{2}
}

func (x *GaiaDeviceBasicInfo) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetMobiApp() string {
	if x != nil {
		return x.MobiApp
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetSdkver() string {
	if x != nil {
		return x.Sdkver
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetAppVersionCode() string {
	if x != nil {
		return x.AppVersionCode
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetOsver() string {
	if x != nil {
		return x.Osver
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetBuvidLocal() string {
	if x != nil {
		return x.BuvidLocal
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetBuvid() string {
	if x != nil {
		return x.Buvid
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

func (x *GaiaDeviceBasicInfo) GetFts() int64 {
	if x != nil {
		return x.Fts
	}
	return 0
}

func (x *GaiaDeviceBasicInfo) GetFirst() int32 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *GaiaDeviceBasicInfo) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

// 应用列表上报-请求
type GaiaEncryptMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 上报头部
	Header *GaiaMsgHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// 加密数据
	EncryptPayload []byte `protobuf:"bytes,2,opt,name=encrypt_payload,json=encryptPayload,proto3" json:"encrypt_payload,omitempty"`
}

func (x *GaiaEncryptMsgReq) Reset() {
	*x = GaiaEncryptMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaiaEncryptMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaiaEncryptMsgReq) ProtoMessage() {}

func (x *GaiaEncryptMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaiaEncryptMsgReq.ProtoReflect.Descriptor instead.
func (*GaiaEncryptMsgReq) Descriptor() ([]byte, []int) {
	return file_bilibili_gaia_gw_gw_api_proto_rawDescGZIP(), []int{3}
}

func (x *GaiaEncryptMsgReq) GetHeader() *GaiaMsgHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GaiaEncryptMsgReq) GetEncryptPayload() []byte {
	if x != nil {
		return x.EncryptPayload
	}
	return nil
}

// 风控通用消息头
type GaiaMsgHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 加密类型
	EncodeType EncryptType `protobuf:"varint,1,opt,name=encode_type,json=encodeType,proto3,enum=bilibili.gaia.gw.EncryptType" json:"encode_type,omitempty"`
	// 类型
	PayloadType PayloadType `protobuf:"varint,2,opt,name=payload_type,json=payloadType,proto3,enum=bilibili.gaia.gw.PayloadType" json:"payload_type,omitempty"`
	// RAS加密后的aes_key
	EncodedAesKey []byte `protobuf:"bytes,3,opt,name=encoded_aes_key,json=encodedAesKey,proto3" json:"encoded_aes_key,omitempty"`
	// 当前时间戳(ms)
	Ts int64 `protobuf:"varint,4,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *GaiaMsgHeader) Reset() {
	*x = GaiaMsgHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaiaMsgHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaiaMsgHeader) ProtoMessage() {}

func (x *GaiaMsgHeader) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaiaMsgHeader.ProtoReflect.Descriptor instead.
func (*GaiaMsgHeader) Descriptor() ([]byte, []int) {
	return file_bilibili_gaia_gw_gw_api_proto_rawDescGZIP(), []int{4}
}

func (x *GaiaMsgHeader) GetEncodeType() EncryptType {
	if x != nil {
		return x.EncodeType
	}
	return EncryptType_INVALID_ENCRYPT_TYPE
}

func (x *GaiaMsgHeader) GetPayloadType() PayloadType {
	if x != nil {
		return x.PayloadType
	}
	return PayloadType_INVALID_PAYLOAD
}

func (x *GaiaMsgHeader) GetEncodedAesKey() []byte {
	if x != nil {
		return x.EncodedAesKey
	}
	return nil
}

func (x *GaiaMsgHeader) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

// 应用列表上报-响应
type UploadAppListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 上报响应id
	TraceId string `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (x *UploadAppListReply) Reset() {
	*x = UploadAppListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAppListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAppListReply) ProtoMessage() {}

func (x *UploadAppListReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_gaia_gw_gw_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAppListReply.ProtoReflect.Descriptor instead.
func (*UploadAppListReply) Descriptor() ([]byte, []int) {
	return file_bilibili_gaia_gw_gw_api_proto_rawDescGZIP(), []int{5}
}

func (x *UploadAppListReply) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

var File_bilibili_gaia_gw_gw_api_proto protoreflect.FileDescriptor

var file_bilibili_gaia_gw_gw_api_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x67, 0x61, 0x69, 0x61, 0x2f,
	0x67, 0x77, 0x2f, 0x67, 0x77, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x67, 0x61, 0x69, 0x61, 0x2e, 0x67,
	0x77, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73,
	0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22,
	0x92, 0x04, 0x0a, 0x13, 0x47, 0x61, 0x69, 0x61, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x6f, 0x62, 0x69, 0x5f, 0x61, 0x70, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x6f, 0x62, 0x69, 0x41, 0x70, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x64, 0x6b, 0x76, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x64, 0x6b, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x10, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x73, 0x76, 0x65, 0x72, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x73, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x76,
	0x69, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x75, 0x76, 0x69, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75,
	0x76, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x76, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x74, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x66, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x22, 0x75, 0x0a, 0x11, 0x47, 0x61, 0x69, 0x61, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x67, 0x61, 0x69, 0x61, 0x2e, 0x67, 0x77, 0x2e, 0x47, 0x61, 0x69,
	0x61, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xc9, 0x01, 0x0a, 0x0d,
	0x47, 0x61, 0x69, 0x61, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3e, 0x0a,
	0x0b, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x67, 0x61,
	0x69, 0x61, 0x2e, 0x67, 0x77, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0a, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a,
	0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x67,
	0x61, 0x69, 0x61, 0x2e, 0x67, 0x77, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x65, 0x73, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x41, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x22, 0x2f, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x2a, 0x4b, 0x0a, 0x0b, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x45, 0x53, 0x10,
	0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x52, 0x53, 0x41, 0x5f,
	0x41, 0x45, 0x53, 0x10, 0x02, 0x2a, 0x37, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x45, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x01, 0x32, 0xb7,
	0x01, 0x0a, 0x04, 0x47, 0x61, 0x69, 0x61, 0x12, 0x5c, 0x0a, 0x0f, 0x45, 0x78, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x67, 0x61, 0x69, 0x61, 0x2e, 0x67, 0x77, 0x2e, 0x47, 0x61,
	0x69, 0x61, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a,
	0x24, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x67, 0x61, 0x69, 0x61, 0x2e,
	0x67, 0x77, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x51, 0x0a, 0x10, 0x45, 0x78, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x25, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x67, 0x61, 0x69,
	0x61, 0x2e, 0x67, 0x77, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6b, 0x75, 0x30,
	0x31, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f,
	0x67, 0x61, 0x69, 0x61, 0x2f, 0x67, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_gaia_gw_gw_api_proto_rawDescOnce sync.Once
	file_bilibili_gaia_gw_gw_api_proto_rawDescData = file_bilibili_gaia_gw_gw_api_proto_rawDesc
)

func file_bilibili_gaia_gw_gw_api_proto_rawDescGZIP() []byte {
	file_bilibili_gaia_gw_gw_api_proto_rawDescOnce.Do(func() {
		file_bilibili_gaia_gw_gw_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_gaia_gw_gw_api_proto_rawDescData)
	})
	return file_bilibili_gaia_gw_gw_api_proto_rawDescData
}

var file_bilibili_gaia_gw_gw_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bilibili_gaia_gw_gw_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bilibili_gaia_gw_gw_api_proto_goTypes = []interface{}{
	(EncryptType)(0),            // 0: bilibili.gaia.gw.EncryptType
	(PayloadType)(0),            // 1: bilibili.gaia.gw.PayloadType
	(*DeviceAppList)(nil),       // 2: bilibili.gaia.gw.DeviceAppList
	(*FetchPublicKeyReply)(nil), // 3: bilibili.gaia.gw.FetchPublicKeyReply
	(*GaiaDeviceBasicInfo)(nil), // 4: bilibili.gaia.gw.GaiaDeviceBasicInfo
	(*GaiaEncryptMsgReq)(nil),   // 5: bilibili.gaia.gw.GaiaEncryptMsgReq
	(*GaiaMsgHeader)(nil),       // 6: bilibili.gaia.gw.GaiaMsgHeader
	(*UploadAppListReply)(nil),  // 7: bilibili.gaia.gw.UploadAppListReply
	(*emptypb.Empty)(nil),       // 8: google.protobuf.Empty
}
var file_bilibili_gaia_gw_gw_api_proto_depIdxs = []int32{
	6, // 0: bilibili.gaia.gw.GaiaEncryptMsgReq.header:type_name -> bilibili.gaia.gw.GaiaMsgHeader
	0, // 1: bilibili.gaia.gw.GaiaMsgHeader.encode_type:type_name -> bilibili.gaia.gw.EncryptType
	1, // 2: bilibili.gaia.gw.GaiaMsgHeader.payload_type:type_name -> bilibili.gaia.gw.PayloadType
	5, // 3: bilibili.gaia.gw.Gaia.ExUploadAppList:input_type -> bilibili.gaia.gw.GaiaEncryptMsgReq
	8, // 4: bilibili.gaia.gw.Gaia.ExFetchPublicKey:input_type -> google.protobuf.Empty
	7, // 5: bilibili.gaia.gw.Gaia.ExUploadAppList:output_type -> bilibili.gaia.gw.UploadAppListReply
	3, // 6: bilibili.gaia.gw.Gaia.ExFetchPublicKey:output_type -> bilibili.gaia.gw.FetchPublicKeyReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bilibili_gaia_gw_gw_api_proto_init() }
func file_bilibili_gaia_gw_gw_api_proto_init() {
	if File_bilibili_gaia_gw_gw_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_gaia_gw_gw_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceAppList); i {
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
		file_bilibili_gaia_gw_gw_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchPublicKeyReply); i {
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
		file_bilibili_gaia_gw_gw_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaiaDeviceBasicInfo); i {
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
		file_bilibili_gaia_gw_gw_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaiaEncryptMsgReq); i {
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
		file_bilibili_gaia_gw_gw_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaiaMsgHeader); i {
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
		file_bilibili_gaia_gw_gw_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAppListReply); i {
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
			RawDescriptor: file_bilibili_gaia_gw_gw_api_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_gaia_gw_gw_api_proto_goTypes,
		DependencyIndexes: file_bilibili_gaia_gw_gw_api_proto_depIdxs,
		EnumInfos:         file_bilibili_gaia_gw_gw_api_proto_enumTypes,
		MessageInfos:      file_bilibili_gaia_gw_gw_api_proto_msgTypes,
	}.Build()
	File_bilibili_gaia_gw_gw_api_proto = out.File
	file_bilibili_gaia_gw_gw_api_proto_rawDesc = nil
	file_bilibili_gaia_gw_gw_api_proto_goTypes = nil
	file_bilibili_gaia_gw_gw_api_proto_depIdxs = nil
}