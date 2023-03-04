// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: bilibili/im/interfaces/inner-interface/v1/api.proto

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
	InnerInterface_UpdateListInn_FullMethodName = "/bilibili.im.interface.inner.interface.v1.InnerInterface/UpdateListInn"
)

// InnerInterfaceClient is the client API for InnerInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InnerInterfaceClient interface {
	UpdateListInn(ctx context.Context, in *ReqOpBlacklist, opts ...grpc.CallOption) (*RspOpBlacklist, error)
}

type innerInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewInnerInterfaceClient(cc grpc.ClientConnInterface) InnerInterfaceClient {
	return &innerInterfaceClient{cc}
}

func (c *innerInterfaceClient) UpdateListInn(ctx context.Context, in *ReqOpBlacklist, opts ...grpc.CallOption) (*RspOpBlacklist, error) {
	out := new(RspOpBlacklist)
	err := c.cc.Invoke(ctx, InnerInterface_UpdateListInn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InnerInterfaceServer is the server API for InnerInterface service.
// All implementations must embed UnimplementedInnerInterfaceServer
// for forward compatibility
type InnerInterfaceServer interface {
	UpdateListInn(context.Context, *ReqOpBlacklist) (*RspOpBlacklist, error)
	mustEmbedUnimplementedInnerInterfaceServer()
}

// UnimplementedInnerInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedInnerInterfaceServer struct {
}

func (UnimplementedInnerInterfaceServer) UpdateListInn(context.Context, *ReqOpBlacklist) (*RspOpBlacklist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateListInn not implemented")
}
func (UnimplementedInnerInterfaceServer) mustEmbedUnimplementedInnerInterfaceServer() {}

// UnsafeInnerInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InnerInterfaceServer will
// result in compilation errors.
type UnsafeInnerInterfaceServer interface {
	mustEmbedUnimplementedInnerInterfaceServer()
}

func RegisterInnerInterfaceServer(s grpc.ServiceRegistrar, srv InnerInterfaceServer) {
	s.RegisterService(&InnerInterface_ServiceDesc, srv)
}

func _InnerInterface_UpdateListInn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqOpBlacklist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerInterfaceServer).UpdateListInn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InnerInterface_UpdateListInn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerInterfaceServer).UpdateListInn(ctx, req.(*ReqOpBlacklist))
	}
	return interceptor(ctx, in, info, handler)
}

// InnerInterface_ServiceDesc is the grpc.ServiceDesc for InnerInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InnerInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.im.interface.inner.interface.v1.InnerInterface",
	HandlerType: (*InnerInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateListInn",
			Handler:    _InnerInterface_UpdateListInn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/im/interfaces/inner-interface/v1/api.proto",
}
