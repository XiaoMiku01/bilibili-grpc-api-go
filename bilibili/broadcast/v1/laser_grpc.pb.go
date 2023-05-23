// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: bilibili/broadcast/v1/laser.proto

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
	Laser_WatchLogUploadEvent_FullMethodName = "/bilibili.broadcast.v1.Laser/WatchLogUploadEvent"
)

// LaserClient is the client API for Laser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LaserClient interface {
	// 监听上报事件
	WatchLogUploadEvent(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Laser_WatchLogUploadEventClient, error)
}

type laserClient struct {
	cc grpc.ClientConnInterface
}

func NewLaserClient(cc grpc.ClientConnInterface) LaserClient {
	return &laserClient{cc}
}

func (c *laserClient) WatchLogUploadEvent(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Laser_WatchLogUploadEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &Laser_ServiceDesc.Streams[0], Laser_WatchLogUploadEvent_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &laserWatchLogUploadEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Laser_WatchLogUploadEventClient interface {
	Recv() (*LaserLogUploadResp, error)
	grpc.ClientStream
}

type laserWatchLogUploadEventClient struct {
	grpc.ClientStream
}

func (x *laserWatchLogUploadEventClient) Recv() (*LaserLogUploadResp, error) {
	m := new(LaserLogUploadResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LaserServer is the server API for Laser service.
// All implementations must embed UnimplementedLaserServer
// for forward compatibility
type LaserServer interface {
	// 监听上报事件
	WatchLogUploadEvent(*empty.Empty, Laser_WatchLogUploadEventServer) error
	mustEmbedUnimplementedLaserServer()
}

// UnimplementedLaserServer must be embedded to have forward compatible implementations.
type UnimplementedLaserServer struct {
}

func (UnimplementedLaserServer) WatchLogUploadEvent(*empty.Empty, Laser_WatchLogUploadEventServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchLogUploadEvent not implemented")
}
func (UnimplementedLaserServer) mustEmbedUnimplementedLaserServer() {}

// UnsafeLaserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LaserServer will
// result in compilation errors.
type UnsafeLaserServer interface {
	mustEmbedUnimplementedLaserServer()
}

func RegisterLaserServer(s grpc.ServiceRegistrar, srv LaserServer) {
	s.RegisterService(&Laser_ServiceDesc, srv)
}

func _Laser_WatchLogUploadEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LaserServer).WatchLogUploadEvent(m, &laserWatchLogUploadEventServer{stream})
}

type Laser_WatchLogUploadEventServer interface {
	Send(*LaserLogUploadResp) error
	grpc.ServerStream
}

type laserWatchLogUploadEventServer struct {
	grpc.ServerStream
}

func (x *laserWatchLogUploadEventServer) Send(m *LaserLogUploadResp) error {
	return x.ServerStream.SendMsg(m)
}

// Laser_ServiceDesc is the grpc.ServiceDesc for Laser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Laser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.broadcast.v1.Laser",
	HandlerType: (*LaserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchLogUploadEvent",
			Handler:       _Laser_WatchLogUploadEvent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bilibili/broadcast/v1/laser.proto",
}
