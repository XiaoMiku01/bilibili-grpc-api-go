// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: bilibili/relation/interfaces/api.proto

package interfaces

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
	RelationInterface_AtSearch_FullMethodName = "/bilibili.relation.interface.v1.RelationInterface/AtSearch"
)

// RelationInterfaceClient is the client API for RelationInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationInterfaceClient interface {
	// 评论区 At 用户列表 (无需登录鉴权)
	AtSearch(ctx context.Context, in *AtSearchReq, opts ...grpc.CallOption) (*AtSearchReply, error)
}

type relationInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationInterfaceClient(cc grpc.ClientConnInterface) RelationInterfaceClient {
	return &relationInterfaceClient{cc}
}

func (c *relationInterfaceClient) AtSearch(ctx context.Context, in *AtSearchReq, opts ...grpc.CallOption) (*AtSearchReply, error) {
	out := new(AtSearchReply)
	err := c.cc.Invoke(ctx, RelationInterface_AtSearch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationInterfaceServer is the server API for RelationInterface service.
// All implementations must embed UnimplementedRelationInterfaceServer
// for forward compatibility
type RelationInterfaceServer interface {
	// 评论区 At 用户列表 (无需登录鉴权)
	AtSearch(context.Context, *AtSearchReq) (*AtSearchReply, error)
	mustEmbedUnimplementedRelationInterfaceServer()
}

// UnimplementedRelationInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedRelationInterfaceServer struct {
}

func (UnimplementedRelationInterfaceServer) AtSearch(context.Context, *AtSearchReq) (*AtSearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtSearch not implemented")
}
func (UnimplementedRelationInterfaceServer) mustEmbedUnimplementedRelationInterfaceServer() {}

// UnsafeRelationInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationInterfaceServer will
// result in compilation errors.
type UnsafeRelationInterfaceServer interface {
	mustEmbedUnimplementedRelationInterfaceServer()
}

func RegisterRelationInterfaceServer(s grpc.ServiceRegistrar, srv RelationInterfaceServer) {
	s.RegisterService(&RelationInterface_ServiceDesc, srv)
}

func _RelationInterface_AtSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AtSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationInterfaceServer).AtSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationInterface_AtSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationInterfaceServer).AtSearch(ctx, req.(*AtSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationInterface_ServiceDesc is the grpc.ServiceDesc for RelationInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.relation.interface.v1.RelationInterface",
	HandlerType: (*RelationInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AtSearch",
			Handler:    _RelationInterface_AtSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/relation/interfaces/api.proto",
}
