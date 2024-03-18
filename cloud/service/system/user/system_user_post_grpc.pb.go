// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: user/system_user_post.proto

// system_user_post 系统用户职位

package user

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
	SystemUserPostService_SystemUserPostCreate_FullMethodName = "/user.SystemUserPostService/SystemUserPostCreate"
	SystemUserPostService_SystemUserPostList_FullMethodName   = "/user.SystemUserPostService/SystemUserPostList"
)

// SystemUserPostServiceClient is the client API for SystemUserPostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemUserPostServiceClient interface {
	SystemUserPostCreate(ctx context.Context, in *SystemUserPostCreateRequest, opts ...grpc.CallOption) (*SystemUserPostCreateResponse, error)
	SystemUserPostList(ctx context.Context, in *SystemUserPostListRequest, opts ...grpc.CallOption) (*SystemUserPostListResponse, error)
}

type systemUserPostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemUserPostServiceClient(cc grpc.ClientConnInterface) SystemUserPostServiceClient {
	return &systemUserPostServiceClient{cc}
}

func (c *systemUserPostServiceClient) SystemUserPostCreate(ctx context.Context, in *SystemUserPostCreateRequest, opts ...grpc.CallOption) (*SystemUserPostCreateResponse, error) {
	out := new(SystemUserPostCreateResponse)
	err := c.cc.Invoke(ctx, SystemUserPostService_SystemUserPostCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserPostServiceClient) SystemUserPostList(ctx context.Context, in *SystemUserPostListRequest, opts ...grpc.CallOption) (*SystemUserPostListResponse, error) {
	out := new(SystemUserPostListResponse)
	err := c.cc.Invoke(ctx, SystemUserPostService_SystemUserPostList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemUserPostServiceServer is the server API for SystemUserPostService service.
// All implementations must embed UnimplementedSystemUserPostServiceServer
// for forward compatibility
type SystemUserPostServiceServer interface {
	SystemUserPostCreate(context.Context, *SystemUserPostCreateRequest) (*SystemUserPostCreateResponse, error)
	SystemUserPostList(context.Context, *SystemUserPostListRequest) (*SystemUserPostListResponse, error)
	mustEmbedUnimplementedSystemUserPostServiceServer()
}

// UnimplementedSystemUserPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemUserPostServiceServer struct {
}

func (UnimplementedSystemUserPostServiceServer) SystemUserPostCreate(context.Context, *SystemUserPostCreateRequest) (*SystemUserPostCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserPostCreate not implemented")
}
func (UnimplementedSystemUserPostServiceServer) SystemUserPostList(context.Context, *SystemUserPostListRequest) (*SystemUserPostListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserPostList not implemented")
}
func (UnimplementedSystemUserPostServiceServer) mustEmbedUnimplementedSystemUserPostServiceServer() {}

// UnsafeSystemUserPostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemUserPostServiceServer will
// result in compilation errors.
type UnsafeSystemUserPostServiceServer interface {
	mustEmbedUnimplementedSystemUserPostServiceServer()
}

func RegisterSystemUserPostServiceServer(s grpc.ServiceRegistrar, srv SystemUserPostServiceServer) {
	s.RegisterService(&SystemUserPostService_ServiceDesc, srv)
}

func _SystemUserPostService_SystemUserPostCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserPostCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserPostServiceServer).SystemUserPostCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserPostService_SystemUserPostCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserPostServiceServer).SystemUserPostCreate(ctx, req.(*SystemUserPostCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserPostService_SystemUserPostList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserPostListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserPostServiceServer).SystemUserPostList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserPostService_SystemUserPostList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserPostServiceServer).SystemUserPostList(ctx, req.(*SystemUserPostListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemUserPostService_ServiceDesc is the grpc.ServiceDesc for SystemUserPostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemUserPostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.SystemUserPostService",
	HandlerType: (*SystemUserPostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemUserPostCreate",
			Handler:    _SystemUserPostService_SystemUserPostCreate_Handler,
		},
		{
			MethodName: "SystemUserPostList",
			Handler:    _SystemUserPostService_SystemUserPostList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/system_user_post.proto",
}