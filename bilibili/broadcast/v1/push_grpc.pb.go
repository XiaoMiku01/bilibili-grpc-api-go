// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: bilibili/broadcast/v1/push.proto

package v1

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
	Push_WatchMessage_FullMethodName = "/bilibili.broadcast.v1.Push/WatchMessage"
)

// PushClient is the client API for Push service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushClient interface {
	WatchMessage(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Push_WatchMessageClient, error)
}

type pushClient struct {
	cc grpc.ClientConnInterface
}

func NewPushClient(cc grpc.ClientConnInterface) PushClient {
	return &pushClient{cc}
}

func (c *pushClient) WatchMessage(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Push_WatchMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Push_ServiceDesc.Streams[0], Push_WatchMessage_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pushWatchMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Push_WatchMessageClient interface {
	Recv() (*PushMessageResp, error)
	grpc.ClientStream
}

type pushWatchMessageClient struct {
	grpc.ClientStream
}

func (x *pushWatchMessageClient) Recv() (*PushMessageResp, error) {
	m := new(PushMessageResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PushServer is the server API for Push service.
// All implementations must embed UnimplementedPushServer
// for forward compatibility
type PushServer interface {
	WatchMessage(*empty.Empty, Push_WatchMessageServer) error
	mustEmbedUnimplementedPushServer()
}

// UnimplementedPushServer must be embedded to have forward compatible implementations.
type UnimplementedPushServer struct {
}

func (UnimplementedPushServer) WatchMessage(*empty.Empty, Push_WatchMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchMessage not implemented")
}
func (UnimplementedPushServer) mustEmbedUnimplementedPushServer() {}

// UnsafePushServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushServer will
// result in compilation errors.
type UnsafePushServer interface {
	mustEmbedUnimplementedPushServer()
}

func RegisterPushServer(s grpc.ServiceRegistrar, srv PushServer) {
	s.RegisterService(&Push_ServiceDesc, srv)
}

func _Push_WatchMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushServer).WatchMessage(m, &pushWatchMessageServer{stream})
}

type Push_WatchMessageServer interface {
	Send(*PushMessageResp) error
	grpc.ServerStream
}

type pushWatchMessageServer struct {
	grpc.ServerStream
}

func (x *pushWatchMessageServer) Send(m *PushMessageResp) error {
	return x.ServerStream.SendMsg(m)
}

// Push_ServiceDesc is the grpc.ServiceDesc for Push service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Push_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.broadcast.v1.Push",
	HandlerType: (*PushServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchMessage",
			Handler:       _Push_WatchMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bilibili/broadcast/v1/push.proto",
}
