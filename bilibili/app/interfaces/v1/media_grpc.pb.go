// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: bilibili/app/interfaces/v1/media.proto

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
	Media_MediaTab_FullMethodName      = "/bilibili.app.interface.v1.Media/MediaTab"
	Media_MediaDetail_FullMethodName   = "/bilibili.app.interface.v1.Media/MediaDetail"
	Media_MediaVideo_FullMethodName    = "/bilibili.app.interface.v1.Media/MediaVideo"
	Media_MediaRelation_FullMethodName = "/bilibili.app.interface.v1.Media/MediaRelation"
	Media_MediaFollow_FullMethodName   = "/bilibili.app.interface.v1.Media/MediaFollow"
)

// MediaClient is the client API for Media service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaClient interface {
	MediaTab(ctx context.Context, in *MediaTabReq, opts ...grpc.CallOption) (*MediaTabReply, error)
	MediaDetail(ctx context.Context, in *MediaDetailReq, opts ...grpc.CallOption) (*MediaDetailReply, error)
	MediaVideo(ctx context.Context, in *MediaVideoReq, opts ...grpc.CallOption) (*MediaVideoReply, error)
	MediaRelation(ctx context.Context, in *MediaRelationReq, opts ...grpc.CallOption) (*MediaRelationReply, error)
	MediaFollow(ctx context.Context, in *MediaFollowReq, opts ...grpc.CallOption) (*MediaFollowReply, error)
}

type mediaClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaClient(cc grpc.ClientConnInterface) MediaClient {
	return &mediaClient{cc}
}

func (c *mediaClient) MediaTab(ctx context.Context, in *MediaTabReq, opts ...grpc.CallOption) (*MediaTabReply, error) {
	out := new(MediaTabReply)
	err := c.cc.Invoke(ctx, Media_MediaTab_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) MediaDetail(ctx context.Context, in *MediaDetailReq, opts ...grpc.CallOption) (*MediaDetailReply, error) {
	out := new(MediaDetailReply)
	err := c.cc.Invoke(ctx, Media_MediaDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) MediaVideo(ctx context.Context, in *MediaVideoReq, opts ...grpc.CallOption) (*MediaVideoReply, error) {
	out := new(MediaVideoReply)
	err := c.cc.Invoke(ctx, Media_MediaVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) MediaRelation(ctx context.Context, in *MediaRelationReq, opts ...grpc.CallOption) (*MediaRelationReply, error) {
	out := new(MediaRelationReply)
	err := c.cc.Invoke(ctx, Media_MediaRelation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) MediaFollow(ctx context.Context, in *MediaFollowReq, opts ...grpc.CallOption) (*MediaFollowReply, error) {
	out := new(MediaFollowReply)
	err := c.cc.Invoke(ctx, Media_MediaFollow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaServer is the server API for Media service.
// All implementations must embed UnimplementedMediaServer
// for forward compatibility
type MediaServer interface {
	MediaTab(context.Context, *MediaTabReq) (*MediaTabReply, error)
	MediaDetail(context.Context, *MediaDetailReq) (*MediaDetailReply, error)
	MediaVideo(context.Context, *MediaVideoReq) (*MediaVideoReply, error)
	MediaRelation(context.Context, *MediaRelationReq) (*MediaRelationReply, error)
	MediaFollow(context.Context, *MediaFollowReq) (*MediaFollowReply, error)
	mustEmbedUnimplementedMediaServer()
}

// UnimplementedMediaServer must be embedded to have forward compatible implementations.
type UnimplementedMediaServer struct {
}

func (UnimplementedMediaServer) MediaTab(context.Context, *MediaTabReq) (*MediaTabReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MediaTab not implemented")
}
func (UnimplementedMediaServer) MediaDetail(context.Context, *MediaDetailReq) (*MediaDetailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MediaDetail not implemented")
}
func (UnimplementedMediaServer) MediaVideo(context.Context, *MediaVideoReq) (*MediaVideoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MediaVideo not implemented")
}
func (UnimplementedMediaServer) MediaRelation(context.Context, *MediaRelationReq) (*MediaRelationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MediaRelation not implemented")
}
func (UnimplementedMediaServer) MediaFollow(context.Context, *MediaFollowReq) (*MediaFollowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MediaFollow not implemented")
}
func (UnimplementedMediaServer) mustEmbedUnimplementedMediaServer() {}

// UnsafeMediaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaServer will
// result in compilation errors.
type UnsafeMediaServer interface {
	mustEmbedUnimplementedMediaServer()
}

func RegisterMediaServer(s grpc.ServiceRegistrar, srv MediaServer) {
	s.RegisterService(&Media_ServiceDesc, srv)
}

func _Media_MediaTab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaTabReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).MediaTab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_MediaTab_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).MediaTab(ctx, req.(*MediaTabReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_MediaDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).MediaDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_MediaDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).MediaDetail(ctx, req.(*MediaDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_MediaVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).MediaVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_MediaVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).MediaVideo(ctx, req.(*MediaVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_MediaRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaRelationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).MediaRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_MediaRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).MediaRelation(ctx, req.(*MediaRelationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_MediaFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaFollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).MediaFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_MediaFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).MediaFollow(ctx, req.(*MediaFollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Media_ServiceDesc is the grpc.ServiceDesc for Media service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Media_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.interface.v1.Media",
	HandlerType: (*MediaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MediaTab",
			Handler:    _Media_MediaTab_Handler,
		},
		{
			MethodName: "MediaDetail",
			Handler:    _Media_MediaDetail_Handler,
		},
		{
			MethodName: "MediaVideo",
			Handler:    _Media_MediaVideo_Handler,
		},
		{
			MethodName: "MediaRelation",
			Handler:    _Media_MediaRelation_Handler,
		},
		{
			MethodName: "MediaFollow",
			Handler:    _Media_MediaFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/interfaces/v1/media.proto",
}
