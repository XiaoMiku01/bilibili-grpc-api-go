// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: bilibili/account/fission/v1/fission.proto

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
	Fission_Entrance_FullMethodName = "/bilibili.account.fission.v1.Fission/Entrance"
	Fission_Window_FullMethodName   = "/bilibili.account.fission.v1.Fission/Window"
	Fission_Privacy_FullMethodName  = "/bilibili.account.fission.v1.Fission/Privacy"
)

// FissionClient is the client API for Fission service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FissionClient interface {
	// 活动入口
	Entrance(ctx context.Context, in *EntranceReq, opts ...grpc.CallOption) (*EntranceReply, error)
	// 首页弹窗
	Window(ctx context.Context, in *WindowReq, opts ...grpc.CallOption) (*WindowReply, error)
	Privacy(ctx context.Context, in *PrivacyReq, opts ...grpc.CallOption) (*PrivacyReply, error)
}

type fissionClient struct {
	cc grpc.ClientConnInterface
}

func NewFissionClient(cc grpc.ClientConnInterface) FissionClient {
	return &fissionClient{cc}
}

func (c *fissionClient) Entrance(ctx context.Context, in *EntranceReq, opts ...grpc.CallOption) (*EntranceReply, error) {
	out := new(EntranceReply)
	err := c.cc.Invoke(ctx, Fission_Entrance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fissionClient) Window(ctx context.Context, in *WindowReq, opts ...grpc.CallOption) (*WindowReply, error) {
	out := new(WindowReply)
	err := c.cc.Invoke(ctx, Fission_Window_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fissionClient) Privacy(ctx context.Context, in *PrivacyReq, opts ...grpc.CallOption) (*PrivacyReply, error) {
	out := new(PrivacyReply)
	err := c.cc.Invoke(ctx, Fission_Privacy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FissionServer is the server API for Fission service.
// All implementations must embed UnimplementedFissionServer
// for forward compatibility
type FissionServer interface {
	// 活动入口
	Entrance(context.Context, *EntranceReq) (*EntranceReply, error)
	// 首页弹窗
	Window(context.Context, *WindowReq) (*WindowReply, error)
	Privacy(context.Context, *PrivacyReq) (*PrivacyReply, error)
	mustEmbedUnimplementedFissionServer()
}

// UnimplementedFissionServer must be embedded to have forward compatible implementations.
type UnimplementedFissionServer struct {
}

func (UnimplementedFissionServer) Entrance(context.Context, *EntranceReq) (*EntranceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Entrance not implemented")
}
func (UnimplementedFissionServer) Window(context.Context, *WindowReq) (*WindowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Window not implemented")
}
func (UnimplementedFissionServer) Privacy(context.Context, *PrivacyReq) (*PrivacyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Privacy not implemented")
}
func (UnimplementedFissionServer) mustEmbedUnimplementedFissionServer() {}

// UnsafeFissionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FissionServer will
// result in compilation errors.
type UnsafeFissionServer interface {
	mustEmbedUnimplementedFissionServer()
}

func RegisterFissionServer(s grpc.ServiceRegistrar, srv FissionServer) {
	s.RegisterService(&Fission_ServiceDesc, srv)
}

func _Fission_Entrance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntranceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FissionServer).Entrance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fission_Entrance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FissionServer).Entrance(ctx, req.(*EntranceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fission_Window_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WindowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FissionServer).Window(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fission_Window_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FissionServer).Window(ctx, req.(*WindowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fission_Privacy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivacyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FissionServer).Privacy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fission_Privacy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FissionServer).Privacy(ctx, req.(*PrivacyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Fission_ServiceDesc is the grpc.ServiceDesc for Fission service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fission_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.account.fission.v1.Fission",
	HandlerType: (*FissionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Entrance",
			Handler:    _Fission_Entrance_Handler,
		},
		{
			MethodName: "Window",
			Handler:    _Fission_Window_Handler,
		},
		{
			MethodName: "Privacy",
			Handler:    _Fission_Privacy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/account/fission/v1/fission.proto",
}
