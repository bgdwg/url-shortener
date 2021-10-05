// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// UrlShortenerClient is the client API for UrlShortener server.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlShortenerClient interface {
	CreateUrl(ctx context.Context, in *Url, opts ...grpc.CallOption) (*Key, error)
	GetUrl(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Url, error)
}

type urlShortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlShortenerClient(cc grpc.ClientConnInterface) UrlShortenerClient {
	return &urlShortenerClient{cc}
}

func (c *urlShortenerClient) CreateUrl(ctx context.Context, in *Url, opts ...grpc.CallOption) (*Key, error) {
	out := new(Key)
	err := c.cc.Invoke(ctx, "/UrlShortener/CreateUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlShortenerClient) GetUrl(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Url, error) {
	out := new(Url)
	err := c.cc.Invoke(ctx, "/UrlShortener/GetUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlShortenerServer is the server API for UrlShortener server.
// All implementations must embed UnimplementedUrlShortenerServer
// for forward compatibility
type UrlShortenerServer interface {
	CreateUrl(context.Context, *Url) (*Key, error)
	GetUrl(context.Context, *Key) (*Url, error)
	mustEmbedUnimplementedUrlShortenerServer()
}

// UnimplementedUrlShortenerServer must be embedded to have forward compatible implementations.
type UnimplementedUrlShortenerServer struct {
}

func (UnimplementedUrlShortenerServer) CreateUrl(context.Context, *Url) (*Key, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUrl not implemented")
}
func (UnimplementedUrlShortenerServer) GetUrl(context.Context, *Key) (*Url, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrl not implemented")
}
func (UnimplementedUrlShortenerServer) mustEmbedUnimplementedUrlShortenerServer() {}

// UnsafeUrlShortenerServer may be embedded to opt out of forward compatibility for this server.
// Use of this interface is not recommended, as added methods to UrlShortenerServer will
// result in compilation errors.
type UnsafeUrlShortenerServer interface {
	mustEmbedUnimplementedUrlShortenerServer()
}

func RegisterUrlShortenerServer(s grpc.ServiceRegistrar, srv UrlShortenerServer) {
	s.RegisterService(&UrlShortener_ServiceDesc, srv)
}

func _UrlShortener_CreateUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Url)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServer).CreateUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UrlShortener/CreateUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServer).CreateUrl(ctx, req.(*Url))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlShortener_GetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServer).GetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UrlShortener/GetUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServer).GetUrl(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlShortener_ServiceDesc is the grpc.ServiceDesc for UrlShortener server.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlShortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UrlShortener",
	HandlerType: (*UrlShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUrl",
			Handler:    _UrlShortener_CreateUrl_Handler,
		},
		{
			MethodName: "GetUrl",
			Handler:    _UrlShortener_GetUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/server.proto",
}
