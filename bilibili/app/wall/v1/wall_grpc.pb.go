// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: bilibili/app/wall/v1/wall.proto

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
	Wall_RuleInfo_FullMethodName = "/bilibili.app.wall.v1.Wall/RuleInfo"
)

// WallClient is the client API for Wall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WallClient interface {
	// 获取免流规则信息
	RuleInfo(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*RulesReply, error)
}

type wallClient struct {
	cc grpc.ClientConnInterface
}

func NewWallClient(cc grpc.ClientConnInterface) WallClient {
	return &wallClient{cc}
}

func (c *wallClient) RuleInfo(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*RulesReply, error) {
	out := new(RulesReply)
	err := c.cc.Invoke(ctx, Wall_RuleInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WallServer is the server API for Wall service.
// All implementations must embed UnimplementedWallServer
// for forward compatibility
type WallServer interface {
	// 获取免流规则信息
	RuleInfo(context.Context, *RuleRequest) (*RulesReply, error)
	mustEmbedUnimplementedWallServer()
}

// UnimplementedWallServer must be embedded to have forward compatible implementations.
type UnimplementedWallServer struct {
}

func (UnimplementedWallServer) RuleInfo(context.Context, *RuleRequest) (*RulesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RuleInfo not implemented")
}
func (UnimplementedWallServer) mustEmbedUnimplementedWallServer() {}

// UnsafeWallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WallServer will
// result in compilation errors.
type UnsafeWallServer interface {
	mustEmbedUnimplementedWallServer()
}

func RegisterWallServer(s grpc.ServiceRegistrar, srv WallServer) {
	s.RegisterService(&Wall_ServiceDesc, srv)
}

func _Wall_RuleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WallServer).RuleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wall_RuleInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WallServer).RuleInfo(ctx, req.(*RuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Wall_ServiceDesc is the grpc.ServiceDesc for Wall service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wall_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.wall.v1.Wall",
	HandlerType: (*WallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RuleInfo",
			Handler:    _Wall_RuleInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/wall/v1/wall.proto",
}
