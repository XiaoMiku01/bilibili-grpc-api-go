// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: bilibili/community/service/dm/v1/dm.proto

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

// DMClient is the client API for DM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DMClient interface {
	// 获取分段弹幕
	DmSegMobile(ctx context.Context, in *DmSegMobileReq, opts ...grpc.CallOption) (*DmSegMobileReply, error)
	// 客户端弹幕元数据 字幕、分段、防挡蒙版等
	DmView(ctx context.Context, in *DmViewReq, opts ...grpc.CallOption) (*DmViewReply, error)
	// 修改弹幕配置
	DmPlayerConfig(ctx context.Context, in *DmPlayerConfigReq, opts ...grpc.CallOption) (*Response, error)
	// ott弹幕列表
	DmSegOtt(ctx context.Context, in *DmSegOttReq, opts ...grpc.CallOption) (*DmSegOttReply, error)
	// SDK弹幕列表
	DmSegSDK(ctx context.Context, in *DmSegSDKReq, opts ...grpc.CallOption) (*DmSegSDKReply, error)
	DmExpoReport(ctx context.Context, in *DmExpoReportReq, opts ...grpc.CallOption) (*DmExpoReportRes, error)
}

type dMClient struct {
	cc grpc.ClientConnInterface
}

func NewDMClient(cc grpc.ClientConnInterface) DMClient {
	return &dMClient{cc}
}

func (c *dMClient) DmSegMobile(ctx context.Context, in *DmSegMobileReq, opts ...grpc.CallOption) (*DmSegMobileReply, error) {
	out := new(DmSegMobileReply)
	err := c.cc.Invoke(ctx, "/bilibili.community.service.dm.v1.DM/DmSegMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dMClient) DmView(ctx context.Context, in *DmViewReq, opts ...grpc.CallOption) (*DmViewReply, error) {
	out := new(DmViewReply)
	err := c.cc.Invoke(ctx, "/bilibili.community.service.dm.v1.DM/DmView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dMClient) DmPlayerConfig(ctx context.Context, in *DmPlayerConfigReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/bilibili.community.service.dm.v1.DM/DmPlayerConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dMClient) DmSegOtt(ctx context.Context, in *DmSegOttReq, opts ...grpc.CallOption) (*DmSegOttReply, error) {
	out := new(DmSegOttReply)
	err := c.cc.Invoke(ctx, "/bilibili.community.service.dm.v1.DM/DmSegOtt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dMClient) DmSegSDK(ctx context.Context, in *DmSegSDKReq, opts ...grpc.CallOption) (*DmSegSDKReply, error) {
	out := new(DmSegSDKReply)
	err := c.cc.Invoke(ctx, "/bilibili.community.service.dm.v1.DM/DmSegSDK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dMClient) DmExpoReport(ctx context.Context, in *DmExpoReportReq, opts ...grpc.CallOption) (*DmExpoReportRes, error) {
	out := new(DmExpoReportRes)
	err := c.cc.Invoke(ctx, "/bilibili.community.service.dm.v1.DM/DmExpoReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DMServer is the server API for DM service.
// All implementations must embed UnimplementedDMServer
// for forward compatibility
type DMServer interface {
	// 获取分段弹幕
	DmSegMobile(context.Context, *DmSegMobileReq) (*DmSegMobileReply, error)
	// 客户端弹幕元数据 字幕、分段、防挡蒙版等
	DmView(context.Context, *DmViewReq) (*DmViewReply, error)
	// 修改弹幕配置
	DmPlayerConfig(context.Context, *DmPlayerConfigReq) (*Response, error)
	// ott弹幕列表
	DmSegOtt(context.Context, *DmSegOttReq) (*DmSegOttReply, error)
	// SDK弹幕列表
	DmSegSDK(context.Context, *DmSegSDKReq) (*DmSegSDKReply, error)
	DmExpoReport(context.Context, *DmExpoReportReq) (*DmExpoReportRes, error)
	mustEmbedUnimplementedDMServer()
}

// UnimplementedDMServer must be embedded to have forward compatible implementations.
type UnimplementedDMServer struct {
}

func (UnimplementedDMServer) DmSegMobile(context.Context, *DmSegMobileReq) (*DmSegMobileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DmSegMobile not implemented")
}
func (UnimplementedDMServer) DmView(context.Context, *DmViewReq) (*DmViewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DmView not implemented")
}
func (UnimplementedDMServer) DmPlayerConfig(context.Context, *DmPlayerConfigReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DmPlayerConfig not implemented")
}
func (UnimplementedDMServer) DmSegOtt(context.Context, *DmSegOttReq) (*DmSegOttReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DmSegOtt not implemented")
}
func (UnimplementedDMServer) DmSegSDK(context.Context, *DmSegSDKReq) (*DmSegSDKReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DmSegSDK not implemented")
}
func (UnimplementedDMServer) DmExpoReport(context.Context, *DmExpoReportReq) (*DmExpoReportRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DmExpoReport not implemented")
}
func (UnimplementedDMServer) mustEmbedUnimplementedDMServer() {}

// UnsafeDMServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DMServer will
// result in compilation errors.
type UnsafeDMServer interface {
	mustEmbedUnimplementedDMServer()
}

func RegisterDMServer(s grpc.ServiceRegistrar, srv DMServer) {
	s.RegisterService(&DM_ServiceDesc, srv)
}

func _DM_DmSegMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DmSegMobileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DMServer).DmSegMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.community.service.dm.v1.DM/DmSegMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DMServer).DmSegMobile(ctx, req.(*DmSegMobileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DM_DmView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DmViewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DMServer).DmView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.community.service.dm.v1.DM/DmView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DMServer).DmView(ctx, req.(*DmViewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DM_DmPlayerConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DmPlayerConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DMServer).DmPlayerConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.community.service.dm.v1.DM/DmPlayerConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DMServer).DmPlayerConfig(ctx, req.(*DmPlayerConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DM_DmSegOtt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DmSegOttReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DMServer).DmSegOtt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.community.service.dm.v1.DM/DmSegOtt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DMServer).DmSegOtt(ctx, req.(*DmSegOttReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DM_DmSegSDK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DmSegSDKReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DMServer).DmSegSDK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.community.service.dm.v1.DM/DmSegSDK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DMServer).DmSegSDK(ctx, req.(*DmSegSDKReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DM_DmExpoReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DmExpoReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DMServer).DmExpoReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.community.service.dm.v1.DM/DmExpoReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DMServer).DmExpoReport(ctx, req.(*DmExpoReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DM_ServiceDesc is the grpc.ServiceDesc for DM service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DM_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.community.service.dm.v1.DM",
	HandlerType: (*DMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DmSegMobile",
			Handler:    _DM_DmSegMobile_Handler,
		},
		{
			MethodName: "DmView",
			Handler:    _DM_DmView_Handler,
		},
		{
			MethodName: "DmPlayerConfig",
			Handler:    _DM_DmPlayerConfig_Handler,
		},
		{
			MethodName: "DmSegOtt",
			Handler:    _DM_DmSegOtt_Handler,
		},
		{
			MethodName: "DmSegSDK",
			Handler:    _DM_DmSegSDK_Handler,
		},
		{
			MethodName: "DmExpoReport",
			Handler:    _DM_DmExpoReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/community/service/dm/v1/dm.proto",
}