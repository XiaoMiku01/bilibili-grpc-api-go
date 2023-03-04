// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: bilibili/app/view/v1/view.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	View_View_FullMethodName                   = "/bilibili.app.view.v1.View/View"
	View_ViewTag_FullMethodName                = "/bilibili.app.view.v1.View/ViewTag"
	View_ViewMaterial_FullMethodName           = "/bilibili.app.view.v1.View/ViewMaterial"
	View_ViewProgress_FullMethodName           = "/bilibili.app.view.v1.View/ViewProgress"
	View_ShortFormVideoDownload_FullMethodName = "/bilibili.app.view.v1.View/ShortFormVideoDownload"
	View_ClickPlayerCard_FullMethodName        = "/bilibili.app.view.v1.View/ClickPlayerCard"
	View_ClickActivitySeason_FullMethodName    = "/bilibili.app.view.v1.View/ClickActivitySeason"
	View_Season_FullMethodName                 = "/bilibili.app.view.v1.View/Season"
	View_ExposePlayerCard_FullMethodName       = "/bilibili.app.view.v1.View/ExposePlayerCard"
	View_AddContract_FullMethodName            = "/bilibili.app.view.v1.View/AddContract"
	View_ChronosPkg_FullMethodName             = "/bilibili.app.view.v1.View/ChronosPkg"
	View_CacheView_FullMethodName              = "/bilibili.app.view.v1.View/CacheView"
	View_ContinuousPlay_FullMethodName         = "/bilibili.app.view.v1.View/ContinuousPlay"
	View_RelatesFeed_FullMethodName            = "/bilibili.app.view.v1.View/RelatesFeed"
	View_PremiereArchive_FullMethodName        = "/bilibili.app.view.v1.View/PremiereArchive"
	View_Reserve_FullMethodName                = "/bilibili.app.view.v1.View/Reserve"
	View_PlayerRelates_FullMethodName          = "/bilibili.app.view.v1.View/PlayerRelates"
	View_SeasonActivityRecord_FullMethodName   = "/bilibili.app.view.v1.View/SeasonActivityRecord"
	View_SeasonWidgetExpose_FullMethodName     = "/bilibili.app.view.v1.View/SeasonWidgetExpose"
	View_GetArcsPlayer_FullMethodName          = "/bilibili.app.view.v1.View/GetArcsPlayer"
)

// ViewClient is the client API for View service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViewClient interface {
	// 视频页详情页
	View(ctx context.Context, in *ViewReq, opts ...grpc.CallOption) (*ViewReply, error)
	ViewTag(ctx context.Context, in *ViewTagReq, opts ...grpc.CallOption) (*ViewTagReply, error)
	ViewMaterial(ctx context.Context, in *ViewMaterialReq, opts ...grpc.CallOption) (*ViewMaterialReply, error)
	// 视频播放过程中的数据
	ViewProgress(ctx context.Context, in *ViewProgressReq, opts ...grpc.CallOption) (*ViewProgressReply, error)
	// 短视频下载
	ShortFormVideoDownload(ctx context.Context, in *ShortFormVideoDownloadReq, opts ...grpc.CallOption) (*ShortFormVideoDownloadReply, error)
	// 点击播放器卡片事件
	ClickPlayerCard(ctx context.Context, in *ClickPlayerCardReq, opts ...grpc.CallOption) (*NoReply, error)
	// 点击大型活动页预约
	ClickActivitySeason(ctx context.Context, in *ClickActivitySeasonReq, opts ...grpc.CallOption) (*NoReply, error)
	// 合集详情页
	Season(ctx context.Context, in *SeasonReq, opts ...grpc.CallOption) (*SeasonReply, error)
	// 播放器卡片曝光
	ExposePlayerCard(ctx context.Context, in *ExposePlayerCardReq, opts ...grpc.CallOption) (*NoReply, error)
	// 点击签订契约
	AddContract(ctx context.Context, in *AddContractReq, opts ...grpc.CallOption) (*NoReply, error)
	// 资源包
	ChronosPkg(ctx context.Context, in *ChronosPkgReq, opts ...grpc.CallOption) (*Chronos, error)
	CacheView(ctx context.Context, in *CacheViewReq, opts ...grpc.CallOption) (*CacheViewReply, error)
	ContinuousPlay(ctx context.Context, in *ContinuousPlayReq, opts ...grpc.CallOption) (*ContinuousPlayReply, error)
	// 播放页推荐IFS
	RelatesFeed(ctx context.Context, in *RelatesFeedReq, opts ...grpc.CallOption) (*RelatesFeedReply, error)
	PremiereArchive(ctx context.Context, in *PremiereArchiveReq, opts ...grpc.CallOption) (*PremiereArchiveReply, error)
	Reserve(ctx context.Context, in *ReserveReq, opts ...grpc.CallOption) (*ReserveReply, error)
	PlayerRelates(ctx context.Context, in *PlayerRelatesReq, opts ...grpc.CallOption) (*PlayerRelatesReply, error)
	SeasonActivityRecord(ctx context.Context, in *SeasonActivityRecordReq, opts ...grpc.CallOption) (*SeasonActivityRecordReply, error)
	SeasonWidgetExpose(ctx context.Context, in *SeasonWidgetExposeReq, opts ...grpc.CallOption) (*SeasonWidgetExposeReply, error)
	GetArcsPlayer(ctx context.Context, in *GetArcsPlayerReq, opts ...grpc.CallOption) (*GetArcsPlayerReply, error)
}

type viewClient struct {
	cc grpc.ClientConnInterface
}

func NewViewClient(cc grpc.ClientConnInterface) ViewClient {
	return &viewClient{cc}
}

func (c *viewClient) View(ctx context.Context, in *ViewReq, opts ...grpc.CallOption) (*ViewReply, error) {
	out := new(ViewReply)
	err := c.cc.Invoke(ctx, View_View_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) ViewTag(ctx context.Context, in *ViewTagReq, opts ...grpc.CallOption) (*ViewTagReply, error) {
	out := new(ViewTagReply)
	err := c.cc.Invoke(ctx, View_ViewTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) ViewMaterial(ctx context.Context, in *ViewMaterialReq, opts ...grpc.CallOption) (*ViewMaterialReply, error) {
	out := new(ViewMaterialReply)
	err := c.cc.Invoke(ctx, View_ViewMaterial_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) ViewProgress(ctx context.Context, in *ViewProgressReq, opts ...grpc.CallOption) (*ViewProgressReply, error) {
	out := new(ViewProgressReply)
	err := c.cc.Invoke(ctx, View_ViewProgress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) ShortFormVideoDownload(ctx context.Context, in *ShortFormVideoDownloadReq, opts ...grpc.CallOption) (*ShortFormVideoDownloadReply, error) {
	out := new(ShortFormVideoDownloadReply)
	err := c.cc.Invoke(ctx, View_ShortFormVideoDownload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) ClickPlayerCard(ctx context.Context, in *ClickPlayerCardReq, opts ...grpc.CallOption) (*NoReply, error) {
	out := new(NoReply)
	err := c.cc.Invoke(ctx, View_ClickPlayerCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) ClickActivitySeason(ctx context.Context, in *ClickActivitySeasonReq, opts ...grpc.CallOption) (*NoReply, error) {
	out := new(NoReply)
	err := c.cc.Invoke(ctx, View_ClickActivitySeason_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) Season(ctx context.Context, in *SeasonReq, opts ...grpc.CallOption) (*SeasonReply, error) {
	out := new(SeasonReply)
	err := c.cc.Invoke(ctx, View_Season_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) ExposePlayerCard(ctx context.Context, in *ExposePlayerCardReq, opts ...grpc.CallOption) (*NoReply, error) {
	out := new(NoReply)
	err := c.cc.Invoke(ctx, View_ExposePlayerCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) AddContract(ctx context.Context, in *AddContractReq, opts ...grpc.CallOption) (*NoReply, error) {
	out := new(NoReply)
	err := c.cc.Invoke(ctx, View_AddContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) ChronosPkg(ctx context.Context, in *ChronosPkgReq, opts ...grpc.CallOption) (*Chronos, error) {
	out := new(Chronos)
	err := c.cc.Invoke(ctx, View_ChronosPkg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) CacheView(ctx context.Context, in *CacheViewReq, opts ...grpc.CallOption) (*CacheViewReply, error) {
	out := new(CacheViewReply)
	err := c.cc.Invoke(ctx, View_CacheView_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) ContinuousPlay(ctx context.Context, in *ContinuousPlayReq, opts ...grpc.CallOption) (*ContinuousPlayReply, error) {
	out := new(ContinuousPlayReply)
	err := c.cc.Invoke(ctx, View_ContinuousPlay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) RelatesFeed(ctx context.Context, in *RelatesFeedReq, opts ...grpc.CallOption) (*RelatesFeedReply, error) {
	out := new(RelatesFeedReply)
	err := c.cc.Invoke(ctx, View_RelatesFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) PremiereArchive(ctx context.Context, in *PremiereArchiveReq, opts ...grpc.CallOption) (*PremiereArchiveReply, error) {
	out := new(PremiereArchiveReply)
	err := c.cc.Invoke(ctx, View_PremiereArchive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) Reserve(ctx context.Context, in *ReserveReq, opts ...grpc.CallOption) (*ReserveReply, error) {
	out := new(ReserveReply)
	err := c.cc.Invoke(ctx, View_Reserve_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) PlayerRelates(ctx context.Context, in *PlayerRelatesReq, opts ...grpc.CallOption) (*PlayerRelatesReply, error) {
	out := new(PlayerRelatesReply)
	err := c.cc.Invoke(ctx, View_PlayerRelates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) SeasonActivityRecord(ctx context.Context, in *SeasonActivityRecordReq, opts ...grpc.CallOption) (*SeasonActivityRecordReply, error) {
	out := new(SeasonActivityRecordReply)
	err := c.cc.Invoke(ctx, View_SeasonActivityRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) SeasonWidgetExpose(ctx context.Context, in *SeasonWidgetExposeReq, opts ...grpc.CallOption) (*SeasonWidgetExposeReply, error) {
	out := new(SeasonWidgetExposeReply)
	err := c.cc.Invoke(ctx, View_SeasonWidgetExpose_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) GetArcsPlayer(ctx context.Context, in *GetArcsPlayerReq, opts ...grpc.CallOption) (*GetArcsPlayerReply, error) {
	out := new(GetArcsPlayerReply)
	err := c.cc.Invoke(ctx, View_GetArcsPlayer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViewServer is the server API for View service.
// All implementations must embed UnimplementedViewServer
// for forward compatibility
type ViewServer interface {
	// 视频页详情页
	View(context.Context, *ViewReq) (*ViewReply, error)
	ViewTag(context.Context, *ViewTagReq) (*ViewTagReply, error)
	ViewMaterial(context.Context, *ViewMaterialReq) (*ViewMaterialReply, error)
	// 视频播放过程中的数据
	ViewProgress(context.Context, *ViewProgressReq) (*ViewProgressReply, error)
	// 短视频下载
	ShortFormVideoDownload(context.Context, *ShortFormVideoDownloadReq) (*ShortFormVideoDownloadReply, error)
	// 点击播放器卡片事件
	ClickPlayerCard(context.Context, *ClickPlayerCardReq) (*NoReply, error)
	// 点击大型活动页预约
	ClickActivitySeason(context.Context, *ClickActivitySeasonReq) (*NoReply, error)
	// 合集详情页
	Season(context.Context, *SeasonReq) (*SeasonReply, error)
	// 播放器卡片曝光
	ExposePlayerCard(context.Context, *ExposePlayerCardReq) (*NoReply, error)
	// 点击签订契约
	AddContract(context.Context, *AddContractReq) (*NoReply, error)
	// 资源包
	ChronosPkg(context.Context, *ChronosPkgReq) (*Chronos, error)
	CacheView(context.Context, *CacheViewReq) (*CacheViewReply, error)
	ContinuousPlay(context.Context, *ContinuousPlayReq) (*ContinuousPlayReply, error)
	// 播放页推荐IFS
	RelatesFeed(context.Context, *RelatesFeedReq) (*RelatesFeedReply, error)
	PremiereArchive(context.Context, *PremiereArchiveReq) (*PremiereArchiveReply, error)
	Reserve(context.Context, *ReserveReq) (*ReserveReply, error)
	PlayerRelates(context.Context, *PlayerRelatesReq) (*PlayerRelatesReply, error)
	SeasonActivityRecord(context.Context, *SeasonActivityRecordReq) (*SeasonActivityRecordReply, error)
	SeasonWidgetExpose(context.Context, *SeasonWidgetExposeReq) (*SeasonWidgetExposeReply, error)
	GetArcsPlayer(context.Context, *GetArcsPlayerReq) (*GetArcsPlayerReply, error)
	mustEmbedUnimplementedViewServer()
}

// UnimplementedViewServer must be embedded to have forward compatible implementations.
type UnimplementedViewServer struct {
}

func (UnimplementedViewServer) View(context.Context, *ViewReq) (*ViewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedViewServer) ViewTag(context.Context, *ViewTagReq) (*ViewTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewTag not implemented")
}
func (UnimplementedViewServer) ViewMaterial(context.Context, *ViewMaterialReq) (*ViewMaterialReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewMaterial not implemented")
}
func (UnimplementedViewServer) ViewProgress(context.Context, *ViewProgressReq) (*ViewProgressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewProgress not implemented")
}
func (UnimplementedViewServer) ShortFormVideoDownload(context.Context, *ShortFormVideoDownloadReq) (*ShortFormVideoDownloadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortFormVideoDownload not implemented")
}
func (UnimplementedViewServer) ClickPlayerCard(context.Context, *ClickPlayerCardReq) (*NoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClickPlayerCard not implemented")
}
func (UnimplementedViewServer) ClickActivitySeason(context.Context, *ClickActivitySeasonReq) (*NoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClickActivitySeason not implemented")
}
func (UnimplementedViewServer) Season(context.Context, *SeasonReq) (*SeasonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Season not implemented")
}
func (UnimplementedViewServer) ExposePlayerCard(context.Context, *ExposePlayerCardReq) (*NoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExposePlayerCard not implemented")
}
func (UnimplementedViewServer) AddContract(context.Context, *AddContractReq) (*NoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContract not implemented")
}
func (UnimplementedViewServer) ChronosPkg(context.Context, *ChronosPkgReq) (*Chronos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChronosPkg not implemented")
}
func (UnimplementedViewServer) CacheView(context.Context, *CacheViewReq) (*CacheViewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheView not implemented")
}
func (UnimplementedViewServer) ContinuousPlay(context.Context, *ContinuousPlayReq) (*ContinuousPlayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContinuousPlay not implemented")
}
func (UnimplementedViewServer) RelatesFeed(context.Context, *RelatesFeedReq) (*RelatesFeedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelatesFeed not implemented")
}
func (UnimplementedViewServer) PremiereArchive(context.Context, *PremiereArchiveReq) (*PremiereArchiveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PremiereArchive not implemented")
}
func (UnimplementedViewServer) Reserve(context.Context, *ReserveReq) (*ReserveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reserve not implemented")
}
func (UnimplementedViewServer) PlayerRelates(context.Context, *PlayerRelatesReq) (*PlayerRelatesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerRelates not implemented")
}
func (UnimplementedViewServer) SeasonActivityRecord(context.Context, *SeasonActivityRecordReq) (*SeasonActivityRecordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeasonActivityRecord not implemented")
}
func (UnimplementedViewServer) SeasonWidgetExpose(context.Context, *SeasonWidgetExposeReq) (*SeasonWidgetExposeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeasonWidgetExpose not implemented")
}
func (UnimplementedViewServer) GetArcsPlayer(context.Context, *GetArcsPlayerReq) (*GetArcsPlayerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArcsPlayer not implemented")
}
func (UnimplementedViewServer) mustEmbedUnimplementedViewServer() {}

// UnsafeViewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViewServer will
// result in compilation errors.
type UnsafeViewServer interface {
	mustEmbedUnimplementedViewServer()
}

func RegisterViewServer(s grpc.ServiceRegistrar, srv ViewServer) {
	s.RegisterService(&View_ServiceDesc, srv)
}

func _View_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).View(ctx, req.(*ViewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_ViewTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).ViewTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_ViewTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).ViewTag(ctx, req.(*ViewTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_ViewMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewMaterialReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).ViewMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_ViewMaterial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).ViewMaterial(ctx, req.(*ViewMaterialReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_ViewProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewProgressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).ViewProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_ViewProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).ViewProgress(ctx, req.(*ViewProgressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_ShortFormVideoDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortFormVideoDownloadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).ShortFormVideoDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_ShortFormVideoDownload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).ShortFormVideoDownload(ctx, req.(*ShortFormVideoDownloadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_ClickPlayerCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClickPlayerCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).ClickPlayerCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_ClickPlayerCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).ClickPlayerCard(ctx, req.(*ClickPlayerCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_ClickActivitySeason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClickActivitySeasonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).ClickActivitySeason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_ClickActivitySeason_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).ClickActivitySeason(ctx, req.(*ClickActivitySeasonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_Season_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeasonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).Season(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_Season_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).Season(ctx, req.(*SeasonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_ExposePlayerCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExposePlayerCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).ExposePlayerCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_ExposePlayerCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).ExposePlayerCard(ctx, req.(*ExposePlayerCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_AddContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddContractReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).AddContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_AddContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).AddContract(ctx, req.(*AddContractReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_ChronosPkg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChronosPkgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).ChronosPkg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_ChronosPkg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).ChronosPkg(ctx, req.(*ChronosPkgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_CacheView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheViewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).CacheView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_CacheView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).CacheView(ctx, req.(*CacheViewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_ContinuousPlay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContinuousPlayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).ContinuousPlay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_ContinuousPlay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).ContinuousPlay(ctx, req.(*ContinuousPlayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_RelatesFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelatesFeedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).RelatesFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_RelatesFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).RelatesFeed(ctx, req.(*RelatesFeedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_PremiereArchive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PremiereArchiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).PremiereArchive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_PremiereArchive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).PremiereArchive(ctx, req.(*PremiereArchiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_Reserve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).Reserve(ctx, req.(*ReserveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_PlayerRelates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerRelatesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).PlayerRelates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_PlayerRelates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).PlayerRelates(ctx, req.(*PlayerRelatesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_SeasonActivityRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeasonActivityRecordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).SeasonActivityRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_SeasonActivityRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).SeasonActivityRecord(ctx, req.(*SeasonActivityRecordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_SeasonWidgetExpose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeasonWidgetExposeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).SeasonWidgetExpose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_SeasonWidgetExpose_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).SeasonWidgetExpose(ctx, req.(*SeasonWidgetExposeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_GetArcsPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArcsPlayerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).GetArcsPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_GetArcsPlayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).GetArcsPlayer(ctx, req.(*GetArcsPlayerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// View_ServiceDesc is the grpc.ServiceDesc for View service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var View_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.view.v1.View",
	HandlerType: (*ViewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "View",
			Handler:    _View_View_Handler,
		},
		{
			MethodName: "ViewTag",
			Handler:    _View_ViewTag_Handler,
		},
		{
			MethodName: "ViewMaterial",
			Handler:    _View_ViewMaterial_Handler,
		},
		{
			MethodName: "ViewProgress",
			Handler:    _View_ViewProgress_Handler,
		},
		{
			MethodName: "ShortFormVideoDownload",
			Handler:    _View_ShortFormVideoDownload_Handler,
		},
		{
			MethodName: "ClickPlayerCard",
			Handler:    _View_ClickPlayerCard_Handler,
		},
		{
			MethodName: "ClickActivitySeason",
			Handler:    _View_ClickActivitySeason_Handler,
		},
		{
			MethodName: "Season",
			Handler:    _View_Season_Handler,
		},
		{
			MethodName: "ExposePlayerCard",
			Handler:    _View_ExposePlayerCard_Handler,
		},
		{
			MethodName: "AddContract",
			Handler:    _View_AddContract_Handler,
		},
		{
			MethodName: "ChronosPkg",
			Handler:    _View_ChronosPkg_Handler,
		},
		{
			MethodName: "CacheView",
			Handler:    _View_CacheView_Handler,
		},
		{
			MethodName: "ContinuousPlay",
			Handler:    _View_ContinuousPlay_Handler,
		},
		{
			MethodName: "RelatesFeed",
			Handler:    _View_RelatesFeed_Handler,
		},
		{
			MethodName: "PremiereArchive",
			Handler:    _View_PremiereArchive_Handler,
		},
		{
			MethodName: "Reserve",
			Handler:    _View_Reserve_Handler,
		},
		{
			MethodName: "PlayerRelates",
			Handler:    _View_PlayerRelates_Handler,
		},
		{
			MethodName: "SeasonActivityRecord",
			Handler:    _View_SeasonActivityRecord_Handler,
		},
		{
			MethodName: "SeasonWidgetExpose",
			Handler:    _View_SeasonWidgetExpose_Handler,
		},
		{
			MethodName: "GetArcsPlayer",
			Handler:    _View_GetArcsPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/view/v1/view.proto",
}
