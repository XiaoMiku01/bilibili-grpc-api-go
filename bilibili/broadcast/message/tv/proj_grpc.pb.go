// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: bilibili/broadcast/message/tv/proj.proto

package tv

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Tv_Proj_FullMethodName       = "/bilibili.broadcast.message.tv.Tv/Proj"
	Tv_LiveStatus_FullMethodName = "/bilibili.broadcast.message.tv.Tv/LiveStatus"
	Tv_Esports_FullMethodName    = "/bilibili.broadcast.message.tv.Tv/Esports"
	Tv_Publicity_FullMethodName  = "/bilibili.broadcast.message.tv.Tv/Publicity"
	Tv_LiveSkip_FullMethodName   = "/bilibili.broadcast.message.tv.Tv/LiveSkip"
)

// TvClient is the client API for Tv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TvClient interface {
	// 投屏
	Proj(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tv_ProjClient, error)
	// 直播状态
	LiveStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tv_LiveStatusClient, error)
	// 赛事比分通知
	Esports(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tv_EsportsClient, error)
	// 直播插卡
	Publicity(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tv_PublicityClient, error)
	// 直转点
	LiveSkip(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tv_LiveSkipClient, error)
}

type tvClient struct {
	cc grpc.ClientConnInterface
}

func NewTvClient(cc grpc.ClientConnInterface) TvClient {
	return &tvClient{cc}
}

func (c *tvClient) Proj(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tv_ProjClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tv_ServiceDesc.Streams[0], Tv_Proj_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &tvProjClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tv_ProjClient interface {
	Recv() (*ProjReply, error)
	grpc.ClientStream
}

type tvProjClient struct {
	grpc.ClientStream
}

func (x *tvProjClient) Recv() (*ProjReply, error) {
	m := new(ProjReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tvClient) LiveStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tv_LiveStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tv_ServiceDesc.Streams[1], Tv_LiveStatus_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &tvLiveStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tv_LiveStatusClient interface {
	Recv() (*LiveStatusNotify, error)
	grpc.ClientStream
}

type tvLiveStatusClient struct {
	grpc.ClientStream
}

func (x *tvLiveStatusClient) Recv() (*LiveStatusNotify, error) {
	m := new(LiveStatusNotify)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tvClient) Esports(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tv_EsportsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tv_ServiceDesc.Streams[2], Tv_Esports_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &tvEsportsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tv_EsportsClient interface {
	Recv() (*EsportsNotify, error)
	grpc.ClientStream
}

type tvEsportsClient struct {
	grpc.ClientStream
}

func (x *tvEsportsClient) Recv() (*EsportsNotify, error) {
	m := new(EsportsNotify)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tvClient) Publicity(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tv_PublicityClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tv_ServiceDesc.Streams[3], Tv_Publicity_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &tvPublicityClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tv_PublicityClient interface {
	Recv() (*PublicityNotify, error)
	grpc.ClientStream
}

type tvPublicityClient struct {
	grpc.ClientStream
}

func (x *tvPublicityClient) Recv() (*PublicityNotify, error) {
	m := new(PublicityNotify)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tvClient) LiveSkip(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Tv_LiveSkipClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tv_ServiceDesc.Streams[4], Tv_LiveSkip_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &tvLiveSkipClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tv_LiveSkipClient interface {
	Recv() (*LiveSkipNotify, error)
	grpc.ClientStream
}

type tvLiveSkipClient struct {
	grpc.ClientStream
}

func (x *tvLiveSkipClient) Recv() (*LiveSkipNotify, error) {
	m := new(LiveSkipNotify)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TvServer is the server API for Tv service.
// All implementations must embed UnimplementedTvServer
// for forward compatibility
type TvServer interface {
	// 投屏
	Proj(*emptypb.Empty, Tv_ProjServer) error
	// 直播状态
	LiveStatus(*emptypb.Empty, Tv_LiveStatusServer) error
	// 赛事比分通知
	Esports(*emptypb.Empty, Tv_EsportsServer) error
	// 直播插卡
	Publicity(*emptypb.Empty, Tv_PublicityServer) error
	// 直转点
	LiveSkip(*emptypb.Empty, Tv_LiveSkipServer) error
	mustEmbedUnimplementedTvServer()
}

// UnimplementedTvServer must be embedded to have forward compatible implementations.
type UnimplementedTvServer struct {
}

func (UnimplementedTvServer) Proj(*emptypb.Empty, Tv_ProjServer) error {
	return status.Errorf(codes.Unimplemented, "method Proj not implemented")
}
func (UnimplementedTvServer) LiveStatus(*emptypb.Empty, Tv_LiveStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method LiveStatus not implemented")
}
func (UnimplementedTvServer) Esports(*emptypb.Empty, Tv_EsportsServer) error {
	return status.Errorf(codes.Unimplemented, "method Esports not implemented")
}
func (UnimplementedTvServer) Publicity(*emptypb.Empty, Tv_PublicityServer) error {
	return status.Errorf(codes.Unimplemented, "method Publicity not implemented")
}
func (UnimplementedTvServer) LiveSkip(*emptypb.Empty, Tv_LiveSkipServer) error {
	return status.Errorf(codes.Unimplemented, "method LiveSkip not implemented")
}
func (UnimplementedTvServer) mustEmbedUnimplementedTvServer() {}

// UnsafeTvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TvServer will
// result in compilation errors.
type UnsafeTvServer interface {
	mustEmbedUnimplementedTvServer()
}

func RegisterTvServer(s grpc.ServiceRegistrar, srv TvServer) {
	s.RegisterService(&Tv_ServiceDesc, srv)
}

func _Tv_Proj_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TvServer).Proj(m, &tvProjServer{stream})
}

type Tv_ProjServer interface {
	Send(*ProjReply) error
	grpc.ServerStream
}

type tvProjServer struct {
	grpc.ServerStream
}

func (x *tvProjServer) Send(m *ProjReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Tv_LiveStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TvServer).LiveStatus(m, &tvLiveStatusServer{stream})
}

type Tv_LiveStatusServer interface {
	Send(*LiveStatusNotify) error
	grpc.ServerStream
}

type tvLiveStatusServer struct {
	grpc.ServerStream
}

func (x *tvLiveStatusServer) Send(m *LiveStatusNotify) error {
	return x.ServerStream.SendMsg(m)
}

func _Tv_Esports_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TvServer).Esports(m, &tvEsportsServer{stream})
}

type Tv_EsportsServer interface {
	Send(*EsportsNotify) error
	grpc.ServerStream
}

type tvEsportsServer struct {
	grpc.ServerStream
}

func (x *tvEsportsServer) Send(m *EsportsNotify) error {
	return x.ServerStream.SendMsg(m)
}

func _Tv_Publicity_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TvServer).Publicity(m, &tvPublicityServer{stream})
}

type Tv_PublicityServer interface {
	Send(*PublicityNotify) error
	grpc.ServerStream
}

type tvPublicityServer struct {
	grpc.ServerStream
}

func (x *tvPublicityServer) Send(m *PublicityNotify) error {
	return x.ServerStream.SendMsg(m)
}

func _Tv_LiveSkip_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TvServer).LiveSkip(m, &tvLiveSkipServer{stream})
}

type Tv_LiveSkipServer interface {
	Send(*LiveSkipNotify) error
	grpc.ServerStream
}

type tvLiveSkipServer struct {
	grpc.ServerStream
}

func (x *tvLiveSkipServer) Send(m *LiveSkipNotify) error {
	return x.ServerStream.SendMsg(m)
}

// Tv_ServiceDesc is the grpc.ServiceDesc for Tv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.broadcast.message.tv.Tv",
	HandlerType: (*TvServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Proj",
			Handler:       _Tv_Proj_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LiveStatus",
			Handler:       _Tv_LiveStatus_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Esports",
			Handler:       _Tv_Esports_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Publicity",
			Handler:       _Tv_Publicity_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LiveSkip",
			Handler:       _Tv_LiveSkip_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bilibili/broadcast/message/tv/proj.proto",
}
