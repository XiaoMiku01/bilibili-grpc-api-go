// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: bilibili/vega/deneb/v1/deneb.proto

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
	VegaDenebRPC_MessagePulls_FullMethodName = "/bilibili.vega.deneb.v1.VegaDenebRPC/MessagePulls"
)

// VegaDenebRPCClient is the client API for VegaDenebRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VegaDenebRPCClient interface {
	MessagePulls(ctx context.Context, in *MessagePullsReq, opts ...grpc.CallOption) (*MessagePullsReply, error)
}

type vegaDenebRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewVegaDenebRPCClient(cc grpc.ClientConnInterface) VegaDenebRPCClient {
	return &vegaDenebRPCClient{cc}
}

func (c *vegaDenebRPCClient) MessagePulls(ctx context.Context, in *MessagePullsReq, opts ...grpc.CallOption) (*MessagePullsReply, error) {
	out := new(MessagePullsReply)
	err := c.cc.Invoke(ctx, VegaDenebRPC_MessagePulls_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VegaDenebRPCServer is the server API for VegaDenebRPC service.
// All implementations must embed UnimplementedVegaDenebRPCServer
// for forward compatibility
type VegaDenebRPCServer interface {
	MessagePulls(context.Context, *MessagePullsReq) (*MessagePullsReply, error)
	mustEmbedUnimplementedVegaDenebRPCServer()
}

// UnimplementedVegaDenebRPCServer must be embedded to have forward compatible implementations.
type UnimplementedVegaDenebRPCServer struct {
}

func (UnimplementedVegaDenebRPCServer) MessagePulls(context.Context, *MessagePullsReq) (*MessagePullsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessagePulls not implemented")
}
func (UnimplementedVegaDenebRPCServer) mustEmbedUnimplementedVegaDenebRPCServer() {}

// UnsafeVegaDenebRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VegaDenebRPCServer will
// result in compilation errors.
type UnsafeVegaDenebRPCServer interface {
	mustEmbedUnimplementedVegaDenebRPCServer()
}

func RegisterVegaDenebRPCServer(s grpc.ServiceRegistrar, srv VegaDenebRPCServer) {
	s.RegisterService(&VegaDenebRPC_ServiceDesc, srv)
}

func _VegaDenebRPC_MessagePulls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessagePullsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VegaDenebRPCServer).MessagePulls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VegaDenebRPC_MessagePulls_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VegaDenebRPCServer).MessagePulls(ctx, req.(*MessagePullsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VegaDenebRPC_ServiceDesc is the grpc.ServiceDesc for VegaDenebRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VegaDenebRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.vega.deneb.v1.VegaDenebRPC",
	HandlerType: (*VegaDenebRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessagePulls",
			Handler:    _VegaDenebRPC_MessagePulls_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/vega/deneb/v1/deneb.proto",
}
