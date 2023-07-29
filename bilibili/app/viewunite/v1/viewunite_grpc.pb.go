// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: bilibili/app/viewunite/v1/viewunite.proto

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
	View_View_FullMethodName         = "/bilibili.app.viewunite.v1.View/View"
	View_ViewProgress_FullMethodName = "/bilibili.app.viewunite.v1.View/ViewProgress"
)

// ViewClient is the client API for View service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViewClient interface {
	View(ctx context.Context, in *ViewReq, opts ...grpc.CallOption) (*ViewReply, error)
	ViewProgress(ctx context.Context, in *ViewProgressReq, opts ...grpc.CallOption) (*ViewProgressReply, error)
}

type viewClient struct {
	cc grpc.ClientConnInterface
}

func NewViewClient(cc grpc.ClientConnInterface) ViewClient {
	return &viewClient{cc}
}

func (c *viewClient) View(ctx context.Context, in *ViewReq, opts ...grpc.CallOption) (*ViewReply, error) {
	out := new(ViewReply)
	err := c.cc.Invoke(ctx, View_View_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewClient) ViewProgress(ctx context.Context, in *ViewProgressReq, opts ...grpc.CallOption) (*ViewProgressReply, error) {
	out := new(ViewProgressReply)
	err := c.cc.Invoke(ctx, View_ViewProgress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViewServer is the server API for View service.
// All implementations must embed UnimplementedViewServer
// for forward compatibility
type ViewServer interface {
	View(context.Context, *ViewReq) (*ViewReply, error)
	ViewProgress(context.Context, *ViewProgressReq) (*ViewProgressReply, error)
	mustEmbedUnimplementedViewServer()
}

// UnimplementedViewServer must be embedded to have forward compatible implementations.
type UnimplementedViewServer struct {
}

func (UnimplementedViewServer) View(context.Context, *ViewReq) (*ViewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedViewServer) ViewProgress(context.Context, *ViewProgressReq) (*ViewProgressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewProgress not implemented")
}
func (UnimplementedViewServer) mustEmbedUnimplementedViewServer() {}

// UnsafeViewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViewServer will
// result in compilation errors.
type UnsafeViewServer interface {
	mustEmbedUnimplementedViewServer()
}

func RegisterViewServer(s grpc.ServiceRegistrar, srv ViewServer) {
	s.RegisterService(&View_ServiceDesc, srv)
}

func _View_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).View(ctx, req.(*ViewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _View_ViewProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewProgressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServer).ViewProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: View_ViewProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServer).ViewProgress(ctx, req.(*ViewProgressReq))
	}
	return interceptor(ctx, in, info, handler)
}

// View_ServiceDesc is the grpc.ServiceDesc for View service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var View_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.viewunite.v1.View",
	HandlerType: (*ViewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "View",
			Handler:    _View_View_Handler,
		},
		{
			MethodName: "ViewProgress",
			Handler:    _View_ViewProgress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/viewunite/v1/viewunite.proto",
}
