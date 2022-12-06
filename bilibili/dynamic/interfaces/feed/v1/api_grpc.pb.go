// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bilibili/dynamic/interfaces/feed/v1/api.proto

package v1

import (
	context "context"
	common "github.com/XiaoMiku01/bilibili-grpc-api-go/bilibili/dynamic/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FeedClient is the client API for Feed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedClient interface {
	// 发布页预校验
	CreateInitCheck(ctx context.Context, in *CreateInitCheckReq, opts ...grpc.CallOption) (*common.CreateCheckResp, error)
	SubmitCheck(ctx context.Context, in *SubmitCheckReq, opts ...grpc.CallOption) (*SubmitCheckRsp, error)
	// 创建动态
	CreateDyn(ctx context.Context, in *CreateDynReq, opts ...grpc.CallOption) (*common.CreateResp, error)
	// 根据name取uid
	GetUidByName(ctx context.Context, in *common.GetUidByNameReq, opts ...grpc.CallOption) (*common.GetUidByNameRsp, error)
	// at用户推荐列表
	AtList(ctx context.Context, in *common.AtListReq, opts ...grpc.CallOption) (*common.AtListRsp, error)
	// at用户搜索列表
	AtSearch(ctx context.Context, in *common.AtSearchReq, opts ...grpc.CallOption) (*common.AtListRsp, error)
	ReserveButtonClick(ctx context.Context, in *ReserveButtonClickReq, opts ...grpc.CallOption) (*ReserveButtonClickResp, error)
	CreatePlusButtonClick(ctx context.Context, in *CreatePlusButtonClickReq, opts ...grpc.CallOption) (*CreatePlusButtonClickRsp, error)
	HotSearch(ctx context.Context, in *HotSearchReq, opts ...grpc.CallOption) (*HotSearchRsp, error)
	Suggest(ctx context.Context, in *SuggestReq, opts ...grpc.CallOption) (*SuggestRsp, error)
	DynamicButtonClick(ctx context.Context, in *DynamicButtonClickReq, opts ...grpc.CallOption) (*DynamicButtonClickRsp, error)
	CreatePermissionButtonClick(ctx context.Context, in *CreatePermissionButtonClickReq, opts ...grpc.CallOption) (*CreatePermissionButtonClickRsp, error)
	CreatePageInfos(ctx context.Context, in *CreatePageInfosReq, opts ...grpc.CallOption) (*CreatePageInfosRsp, error)
}

type feedClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedClient(cc grpc.ClientConnInterface) FeedClient {
	return &feedClient{cc}
}

func (c *feedClient) CreateInitCheck(ctx context.Context, in *CreateInitCheckReq, opts ...grpc.CallOption) (*common.CreateCheckResp, error) {
	out := new(common.CreateCheckResp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/CreateInitCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) SubmitCheck(ctx context.Context, in *SubmitCheckReq, opts ...grpc.CallOption) (*SubmitCheckRsp, error) {
	out := new(SubmitCheckRsp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/SubmitCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) CreateDyn(ctx context.Context, in *CreateDynReq, opts ...grpc.CallOption) (*common.CreateResp, error) {
	out := new(common.CreateResp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/CreateDyn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) GetUidByName(ctx context.Context, in *common.GetUidByNameReq, opts ...grpc.CallOption) (*common.GetUidByNameRsp, error) {
	out := new(common.GetUidByNameRsp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/GetUidByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) AtList(ctx context.Context, in *common.AtListReq, opts ...grpc.CallOption) (*common.AtListRsp, error) {
	out := new(common.AtListRsp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/AtList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) AtSearch(ctx context.Context, in *common.AtSearchReq, opts ...grpc.CallOption) (*common.AtListRsp, error) {
	out := new(common.AtListRsp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/AtSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) ReserveButtonClick(ctx context.Context, in *ReserveButtonClickReq, opts ...grpc.CallOption) (*ReserveButtonClickResp, error) {
	out := new(ReserveButtonClickResp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/ReserveButtonClick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) CreatePlusButtonClick(ctx context.Context, in *CreatePlusButtonClickReq, opts ...grpc.CallOption) (*CreatePlusButtonClickRsp, error) {
	out := new(CreatePlusButtonClickRsp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/CreatePlusButtonClick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) HotSearch(ctx context.Context, in *HotSearchReq, opts ...grpc.CallOption) (*HotSearchRsp, error) {
	out := new(HotSearchRsp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/HotSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) Suggest(ctx context.Context, in *SuggestReq, opts ...grpc.CallOption) (*SuggestRsp, error) {
	out := new(SuggestRsp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/Suggest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) DynamicButtonClick(ctx context.Context, in *DynamicButtonClickReq, opts ...grpc.CallOption) (*DynamicButtonClickRsp, error) {
	out := new(DynamicButtonClickRsp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/DynamicButtonClick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) CreatePermissionButtonClick(ctx context.Context, in *CreatePermissionButtonClickReq, opts ...grpc.CallOption) (*CreatePermissionButtonClickRsp, error) {
	out := new(CreatePermissionButtonClickRsp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/CreatePermissionButtonClick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) CreatePageInfos(ctx context.Context, in *CreatePageInfosReq, opts ...grpc.CallOption) (*CreatePageInfosRsp, error) {
	out := new(CreatePageInfosRsp)
	err := c.cc.Invoke(ctx, "/bilibili.main.dynamic.feed.v1.Feed/CreatePageInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServer is the server API for Feed service.
// All implementations must embed UnimplementedFeedServer
// for forward compatibility
type FeedServer interface {
	// 发布页预校验
	CreateInitCheck(context.Context, *CreateInitCheckReq) (*common.CreateCheckResp, error)
	SubmitCheck(context.Context, *SubmitCheckReq) (*SubmitCheckRsp, error)
	// 创建动态
	CreateDyn(context.Context, *CreateDynReq) (*common.CreateResp, error)
	// 根据name取uid
	GetUidByName(context.Context, *common.GetUidByNameReq) (*common.GetUidByNameRsp, error)
	// at用户推荐列表
	AtList(context.Context, *common.AtListReq) (*common.AtListRsp, error)
	// at用户搜索列表
	AtSearch(context.Context, *common.AtSearchReq) (*common.AtListRsp, error)
	ReserveButtonClick(context.Context, *ReserveButtonClickReq) (*ReserveButtonClickResp, error)
	CreatePlusButtonClick(context.Context, *CreatePlusButtonClickReq) (*CreatePlusButtonClickRsp, error)
	HotSearch(context.Context, *HotSearchReq) (*HotSearchRsp, error)
	Suggest(context.Context, *SuggestReq) (*SuggestRsp, error)
	DynamicButtonClick(context.Context, *DynamicButtonClickReq) (*DynamicButtonClickRsp, error)
	CreatePermissionButtonClick(context.Context, *CreatePermissionButtonClickReq) (*CreatePermissionButtonClickRsp, error)
	CreatePageInfos(context.Context, *CreatePageInfosReq) (*CreatePageInfosRsp, error)
	mustEmbedUnimplementedFeedServer()
}

// UnimplementedFeedServer must be embedded to have forward compatible implementations.
type UnimplementedFeedServer struct {
}

func (UnimplementedFeedServer) CreateInitCheck(context.Context, *CreateInitCheckReq) (*common.CreateCheckResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInitCheck not implemented")
}
func (UnimplementedFeedServer) SubmitCheck(context.Context, *SubmitCheckReq) (*SubmitCheckRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitCheck not implemented")
}
func (UnimplementedFeedServer) CreateDyn(context.Context, *CreateDynReq) (*common.CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDyn not implemented")
}
func (UnimplementedFeedServer) GetUidByName(context.Context, *common.GetUidByNameReq) (*common.GetUidByNameRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUidByName not implemented")
}
func (UnimplementedFeedServer) AtList(context.Context, *common.AtListReq) (*common.AtListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtList not implemented")
}
func (UnimplementedFeedServer) AtSearch(context.Context, *common.AtSearchReq) (*common.AtListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtSearch not implemented")
}
func (UnimplementedFeedServer) ReserveButtonClick(context.Context, *ReserveButtonClickReq) (*ReserveButtonClickResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReserveButtonClick not implemented")
}
func (UnimplementedFeedServer) CreatePlusButtonClick(context.Context, *CreatePlusButtonClickReq) (*CreatePlusButtonClickRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlusButtonClick not implemented")
}
func (UnimplementedFeedServer) HotSearch(context.Context, *HotSearchReq) (*HotSearchRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HotSearch not implemented")
}
func (UnimplementedFeedServer) Suggest(context.Context, *SuggestReq) (*SuggestRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggest not implemented")
}
func (UnimplementedFeedServer) DynamicButtonClick(context.Context, *DynamicButtonClickReq) (*DynamicButtonClickRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DynamicButtonClick not implemented")
}
func (UnimplementedFeedServer) CreatePermissionButtonClick(context.Context, *CreatePermissionButtonClickReq) (*CreatePermissionButtonClickRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermissionButtonClick not implemented")
}
func (UnimplementedFeedServer) CreatePageInfos(context.Context, *CreatePageInfosReq) (*CreatePageInfosRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePageInfos not implemented")
}
func (UnimplementedFeedServer) mustEmbedUnimplementedFeedServer() {}

// UnsafeFeedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServer will
// result in compilation errors.
type UnsafeFeedServer interface {
	mustEmbedUnimplementedFeedServer()
}

func RegisterFeedServer(s grpc.ServiceRegistrar, srv FeedServer) {
	s.RegisterService(&Feed_ServiceDesc, srv)
}

func _Feed_CreateInitCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInitCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).CreateInitCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/CreateInitCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).CreateInitCheck(ctx, req.(*CreateInitCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_SubmitCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).SubmitCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/SubmitCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).SubmitCheck(ctx, req.(*SubmitCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_CreateDyn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDynReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).CreateDyn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/CreateDyn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).CreateDyn(ctx, req.(*CreateDynReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_GetUidByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.GetUidByNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).GetUidByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/GetUidByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).GetUidByName(ctx, req.(*common.GetUidByNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_AtList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.AtListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).AtList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/AtList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).AtList(ctx, req.(*common.AtListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_AtSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.AtSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).AtSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/AtSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).AtSearch(ctx, req.(*common.AtSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_ReserveButtonClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveButtonClickReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).ReserveButtonClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/ReserveButtonClick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).ReserveButtonClick(ctx, req.(*ReserveButtonClickReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_CreatePlusButtonClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlusButtonClickReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).CreatePlusButtonClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/CreatePlusButtonClick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).CreatePlusButtonClick(ctx, req.(*CreatePlusButtonClickReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_HotSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).HotSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/HotSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).HotSearch(ctx, req.(*HotSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_Suggest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).Suggest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/Suggest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).Suggest(ctx, req.(*SuggestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_DynamicButtonClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DynamicButtonClickReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).DynamicButtonClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/DynamicButtonClick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).DynamicButtonClick(ctx, req.(*DynamicButtonClickReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_CreatePermissionButtonClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionButtonClickReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).CreatePermissionButtonClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/CreatePermissionButtonClick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).CreatePermissionButtonClick(ctx, req.(*CreatePermissionButtonClickReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_CreatePageInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePageInfosReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).CreatePageInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.main.dynamic.feed.v1.Feed/CreatePageInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).CreatePageInfos(ctx, req.(*CreatePageInfosReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Feed_ServiceDesc is the grpc.ServiceDesc for Feed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Feed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.main.dynamic.feed.v1.Feed",
	HandlerType: (*FeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInitCheck",
			Handler:    _Feed_CreateInitCheck_Handler,
		},
		{
			MethodName: "SubmitCheck",
			Handler:    _Feed_SubmitCheck_Handler,
		},
		{
			MethodName: "CreateDyn",
			Handler:    _Feed_CreateDyn_Handler,
		},
		{
			MethodName: "GetUidByName",
			Handler:    _Feed_GetUidByName_Handler,
		},
		{
			MethodName: "AtList",
			Handler:    _Feed_AtList_Handler,
		},
		{
			MethodName: "AtSearch",
			Handler:    _Feed_AtSearch_Handler,
		},
		{
			MethodName: "ReserveButtonClick",
			Handler:    _Feed_ReserveButtonClick_Handler,
		},
		{
			MethodName: "CreatePlusButtonClick",
			Handler:    _Feed_CreatePlusButtonClick_Handler,
		},
		{
			MethodName: "HotSearch",
			Handler:    _Feed_HotSearch_Handler,
		},
		{
			MethodName: "Suggest",
			Handler:    _Feed_Suggest_Handler,
		},
		{
			MethodName: "DynamicButtonClick",
			Handler:    _Feed_DynamicButtonClick_Handler,
		},
		{
			MethodName: "CreatePermissionButtonClick",
			Handler:    _Feed_CreatePermissionButtonClick_Handler,
		},
		{
			MethodName: "CreatePageInfos",
			Handler:    _Feed_CreatePageInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/dynamic/interfaces/feed/v1/api.proto",
}