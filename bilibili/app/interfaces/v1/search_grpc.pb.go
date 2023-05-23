// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: bilibili/app/interfaces/v1/search.proto

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
	Search_Suggest3_FullMethodName     = "/bilibili.app.interfaces.v1.Search/Suggest3"
	Search_DefaultWords_FullMethodName = "/bilibili.app.interfaces.v1.Search/DefaultWords"
)

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchClient interface {
	// 获取搜索建议
	Suggest3(ctx context.Context, in *SuggestionResult3Req, opts ...grpc.CallOption) (*SuggestionResult3Reply, error)
	DefaultWords(ctx context.Context, in *DefaultWordsReq, opts ...grpc.CallOption) (*DefaultWordsReply, error)
}

type searchClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchClient(cc grpc.ClientConnInterface) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) Suggest3(ctx context.Context, in *SuggestionResult3Req, opts ...grpc.CallOption) (*SuggestionResult3Reply, error) {
	out := new(SuggestionResult3Reply)
	err := c.cc.Invoke(ctx, Search_Suggest3_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) DefaultWords(ctx context.Context, in *DefaultWordsReq, opts ...grpc.CallOption) (*DefaultWordsReply, error) {
	out := new(DefaultWordsReply)
	err := c.cc.Invoke(ctx, Search_DefaultWords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServer is the server API for Search service.
// All implementations must embed UnimplementedSearchServer
// for forward compatibility
type SearchServer interface {
	// 获取搜索建议
	Suggest3(context.Context, *SuggestionResult3Req) (*SuggestionResult3Reply, error)
	DefaultWords(context.Context, *DefaultWordsReq) (*DefaultWordsReply, error)
	mustEmbedUnimplementedSearchServer()
}

// UnimplementedSearchServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServer struct {
}

func (UnimplementedSearchServer) Suggest3(context.Context, *SuggestionResult3Req) (*SuggestionResult3Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggest3 not implemented")
}
func (UnimplementedSearchServer) DefaultWords(context.Context, *DefaultWordsReq) (*DefaultWordsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefaultWords not implemented")
}
func (UnimplementedSearchServer) mustEmbedUnimplementedSearchServer() {}

// UnsafeSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServer will
// result in compilation errors.
type UnsafeSearchServer interface {
	mustEmbedUnimplementedSearchServer()
}

func RegisterSearchServer(s grpc.ServiceRegistrar, srv SearchServer) {
	s.RegisterService(&Search_ServiceDesc, srv)
}

func _Search_Suggest3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestionResult3Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).Suggest3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Search_Suggest3_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).Suggest3(ctx, req.(*SuggestionResult3Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_DefaultWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultWordsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).DefaultWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Search_DefaultWords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).DefaultWords(ctx, req.(*DefaultWordsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Search_ServiceDesc is the grpc.ServiceDesc for Search service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Search_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.interfaces.v1.Search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Suggest3",
			Handler:    _Search_Suggest3_Handler,
		},
		{
			MethodName: "DefaultWords",
			Handler:    _Search_DefaultWords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/interfaces/v1/search.proto",
}

const (
	SearchTest_NotExist_FullMethodName = "/bilibili.app.interfaces.v1.SearchTest/NotExist"
)

// SearchTestClient is the client API for SearchTest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchTestClient interface {
	NotExist(ctx context.Context, in *SuggestionResult3Req, opts ...grpc.CallOption) (*SuggestionResult3Reply, error)
}

type searchTestClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchTestClient(cc grpc.ClientConnInterface) SearchTestClient {
	return &searchTestClient{cc}
}

func (c *searchTestClient) NotExist(ctx context.Context, in *SuggestionResult3Req, opts ...grpc.CallOption) (*SuggestionResult3Reply, error) {
	out := new(SuggestionResult3Reply)
	err := c.cc.Invoke(ctx, SearchTest_NotExist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchTestServer is the server API for SearchTest service.
// All implementations must embed UnimplementedSearchTestServer
// for forward compatibility
type SearchTestServer interface {
	NotExist(context.Context, *SuggestionResult3Req) (*SuggestionResult3Reply, error)
	mustEmbedUnimplementedSearchTestServer()
}

// UnimplementedSearchTestServer must be embedded to have forward compatible implementations.
type UnimplementedSearchTestServer struct {
}

func (UnimplementedSearchTestServer) NotExist(context.Context, *SuggestionResult3Req) (*SuggestionResult3Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotExist not implemented")
}
func (UnimplementedSearchTestServer) mustEmbedUnimplementedSearchTestServer() {}

// UnsafeSearchTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchTestServer will
// result in compilation errors.
type UnsafeSearchTestServer interface {
	mustEmbedUnimplementedSearchTestServer()
}

func RegisterSearchTestServer(s grpc.ServiceRegistrar, srv SearchTestServer) {
	s.RegisterService(&SearchTest_ServiceDesc, srv)
}

func _SearchTest_NotExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestionResult3Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchTestServer).NotExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchTest_NotExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchTestServer).NotExist(ctx, req.(*SuggestionResult3Req))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchTest_ServiceDesc is the grpc.ServiceDesc for SearchTest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchTest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.interfaces.v1.SearchTest",
	HandlerType: (*SearchTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotExist",
			Handler:    _SearchTest_NotExist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/interfaces/v1/search.proto",
}
