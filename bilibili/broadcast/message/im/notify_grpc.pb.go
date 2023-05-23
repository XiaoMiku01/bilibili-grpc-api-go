// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: bilibili/broadcast/message/im/notify.proto

package im

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Notify_WatchNotify_FullMethodName = "/bilibili.broadcast.message.im.Notify/WatchNotify"
)

// NotifyClient is the client API for Notify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifyClient interface {
	WatchNotify(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Notify_WatchNotifyClient, error)
}

type notifyClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifyClient(cc grpc.ClientConnInterface) NotifyClient {
	return &notifyClient{cc}
}

func (c *notifyClient) WatchNotify(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Notify_WatchNotifyClient, error) {
	stream, err := c.cc.NewStream(ctx, &Notify_ServiceDesc.Streams[0], Notify_WatchNotify_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &notifyWatchNotifyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Notify_WatchNotifyClient interface {
	Recv() (*NotifyRsp, error)
	grpc.ClientStream
}

type notifyWatchNotifyClient struct {
	grpc.ClientStream
}

func (x *notifyWatchNotifyClient) Recv() (*NotifyRsp, error) {
	m := new(NotifyRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotifyServer is the server API for Notify service.
// All implementations must embed UnimplementedNotifyServer
// for forward compatibility
type NotifyServer interface {
	WatchNotify(*empty.Empty, Notify_WatchNotifyServer) error
	mustEmbedUnimplementedNotifyServer()
}

// UnimplementedNotifyServer must be embedded to have forward compatible implementations.
type UnimplementedNotifyServer struct {
}

func (UnimplementedNotifyServer) WatchNotify(*empty.Empty, Notify_WatchNotifyServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchNotify not implemented")
}
func (UnimplementedNotifyServer) mustEmbedUnimplementedNotifyServer() {}

// UnsafeNotifyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifyServer will
// result in compilation errors.
type UnsafeNotifyServer interface {
	mustEmbedUnimplementedNotifyServer()
}

func RegisterNotifyServer(s grpc.ServiceRegistrar, srv NotifyServer) {
	s.RegisterService(&Notify_ServiceDesc, srv)
}

func _Notify_WatchNotify_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotifyServer).WatchNotify(m, &notifyWatchNotifyServer{stream})
}

type Notify_WatchNotifyServer interface {
	Send(*NotifyRsp) error
	grpc.ServerStream
}

type notifyWatchNotifyServer struct {
	grpc.ServerStream
}

func (x *notifyWatchNotifyServer) Send(m *NotifyRsp) error {
	return x.ServerStream.SendMsg(m)
}

// Notify_ServiceDesc is the grpc.ServiceDesc for Notify service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notify_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.broadcast.message.im.Notify",
	HandlerType: (*NotifyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchNotify",
			Handler:       _Notify_WatchNotify_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bilibili/broadcast/message/im/notify.proto",
}
