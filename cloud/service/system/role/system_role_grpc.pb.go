// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: role/system_role.proto

// system_role 系统角色

package role

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
	SystemRoleService_SystemRoleCreate_FullMethodName          = "/role.SystemRoleService/SystemRoleCreate"
	SystemRoleService_SystemRoleUpdate_FullMethodName          = "/role.SystemRoleService/SystemRoleUpdate"
	SystemRoleService_SystemRoleDelete_FullMethodName          = "/role.SystemRoleService/SystemRoleDelete"
	SystemRoleService_SystemRole_FullMethodName                = "/role.SystemRoleService/SystemRole"
	SystemRoleService_SystemRoleRecover_FullMethodName         = "/role.SystemRoleService/SystemRoleRecover"
	SystemRoleService_SystemRoleList_FullMethodName            = "/role.SystemRoleService/SystemRoleList"
	SystemRoleService_SystemRoleDataScopeCreate_FullMethodName = "/role.SystemRoleService/SystemRoleDataScopeCreate"
	SystemRoleService_SystemRoleDataScope_FullMethodName       = "/role.SystemRoleService/SystemRoleDataScope"
)

// SystemRoleServiceClient is the client API for SystemRoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemRoleServiceClient interface {
	SystemRoleCreate(ctx context.Context, in *SystemRoleCreateRequest, opts ...grpc.CallOption) (*SystemRoleCreateResponse, error)
	SystemRoleUpdate(ctx context.Context, in *SystemRoleUpdateRequest, opts ...grpc.CallOption) (*SystemRoleUpdateResponse, error)
	SystemRoleDelete(ctx context.Context, in *SystemRoleDeleteRequest, opts ...grpc.CallOption) (*SystemRoleDeleteResponse, error)
	SystemRole(ctx context.Context, in *SystemRoleRequest, opts ...grpc.CallOption) (*SystemRoleResponse, error)
	SystemRoleRecover(ctx context.Context, in *SystemRoleRecoverRequest, opts ...grpc.CallOption) (*SystemRoleRecoverResponse, error)
	SystemRoleList(ctx context.Context, in *SystemRoleListRequest, opts ...grpc.CallOption) (*SystemRoleListResponse, error)
	SystemRoleDataScopeCreate(ctx context.Context, in *SystemRoleDataScopeCreateRequest, opts ...grpc.CallOption) (*SystemRoleDataScopeCreateResponse, error)
	SystemRoleDataScope(ctx context.Context, in *SystemRoleDataScopeRequest, opts ...grpc.CallOption) (*SystemRoleDataScopeResponse, error)
}

type systemRoleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemRoleServiceClient(cc grpc.ClientConnInterface) SystemRoleServiceClient {
	return &systemRoleServiceClient{cc}
}

func (c *systemRoleServiceClient) SystemRoleCreate(ctx context.Context, in *SystemRoleCreateRequest, opts ...grpc.CallOption) (*SystemRoleCreateResponse, error) {
	out := new(SystemRoleCreateResponse)
	err := c.cc.Invoke(ctx, SystemRoleService_SystemRoleCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRoleServiceClient) SystemRoleUpdate(ctx context.Context, in *SystemRoleUpdateRequest, opts ...grpc.CallOption) (*SystemRoleUpdateResponse, error) {
	out := new(SystemRoleUpdateResponse)
	err := c.cc.Invoke(ctx, SystemRoleService_SystemRoleUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRoleServiceClient) SystemRoleDelete(ctx context.Context, in *SystemRoleDeleteRequest, opts ...grpc.CallOption) (*SystemRoleDeleteResponse, error) {
	out := new(SystemRoleDeleteResponse)
	err := c.cc.Invoke(ctx, SystemRoleService_SystemRoleDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRoleServiceClient) SystemRole(ctx context.Context, in *SystemRoleRequest, opts ...grpc.CallOption) (*SystemRoleResponse, error) {
	out := new(SystemRoleResponse)
	err := c.cc.Invoke(ctx, SystemRoleService_SystemRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRoleServiceClient) SystemRoleRecover(ctx context.Context, in *SystemRoleRecoverRequest, opts ...grpc.CallOption) (*SystemRoleRecoverResponse, error) {
	out := new(SystemRoleRecoverResponse)
	err := c.cc.Invoke(ctx, SystemRoleService_SystemRoleRecover_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRoleServiceClient) SystemRoleList(ctx context.Context, in *SystemRoleListRequest, opts ...grpc.CallOption) (*SystemRoleListResponse, error) {
	out := new(SystemRoleListResponse)
	err := c.cc.Invoke(ctx, SystemRoleService_SystemRoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRoleServiceClient) SystemRoleDataScopeCreate(ctx context.Context, in *SystemRoleDataScopeCreateRequest, opts ...grpc.CallOption) (*SystemRoleDataScopeCreateResponse, error) {
	out := new(SystemRoleDataScopeCreateResponse)
	err := c.cc.Invoke(ctx, SystemRoleService_SystemRoleDataScopeCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRoleServiceClient) SystemRoleDataScope(ctx context.Context, in *SystemRoleDataScopeRequest, opts ...grpc.CallOption) (*SystemRoleDataScopeResponse, error) {
	out := new(SystemRoleDataScopeResponse)
	err := c.cc.Invoke(ctx, SystemRoleService_SystemRoleDataScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemRoleServiceServer is the server API for SystemRoleService service.
// All implementations must embed UnimplementedSystemRoleServiceServer
// for forward compatibility
type SystemRoleServiceServer interface {
	SystemRoleCreate(context.Context, *SystemRoleCreateRequest) (*SystemRoleCreateResponse, error)
	SystemRoleUpdate(context.Context, *SystemRoleUpdateRequest) (*SystemRoleUpdateResponse, error)
	SystemRoleDelete(context.Context, *SystemRoleDeleteRequest) (*SystemRoleDeleteResponse, error)
	SystemRole(context.Context, *SystemRoleRequest) (*SystemRoleResponse, error)
	SystemRoleRecover(context.Context, *SystemRoleRecoverRequest) (*SystemRoleRecoverResponse, error)
	SystemRoleList(context.Context, *SystemRoleListRequest) (*SystemRoleListResponse, error)
	SystemRoleDataScopeCreate(context.Context, *SystemRoleDataScopeCreateRequest) (*SystemRoleDataScopeCreateResponse, error)
	SystemRoleDataScope(context.Context, *SystemRoleDataScopeRequest) (*SystemRoleDataScopeResponse, error)
	mustEmbedUnimplementedSystemRoleServiceServer()
}

// UnimplementedSystemRoleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemRoleServiceServer struct {
}

func (UnimplementedSystemRoleServiceServer) SystemRoleCreate(context.Context, *SystemRoleCreateRequest) (*SystemRoleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemRoleCreate not implemented")
}
func (UnimplementedSystemRoleServiceServer) SystemRoleUpdate(context.Context, *SystemRoleUpdateRequest) (*SystemRoleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemRoleUpdate not implemented")
}
func (UnimplementedSystemRoleServiceServer) SystemRoleDelete(context.Context, *SystemRoleDeleteRequest) (*SystemRoleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemRoleDelete not implemented")
}
func (UnimplementedSystemRoleServiceServer) SystemRole(context.Context, *SystemRoleRequest) (*SystemRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemRole not implemented")
}
func (UnimplementedSystemRoleServiceServer) SystemRoleRecover(context.Context, *SystemRoleRecoverRequest) (*SystemRoleRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemRoleRecover not implemented")
}
func (UnimplementedSystemRoleServiceServer) SystemRoleList(context.Context, *SystemRoleListRequest) (*SystemRoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemRoleList not implemented")
}
func (UnimplementedSystemRoleServiceServer) SystemRoleDataScopeCreate(context.Context, *SystemRoleDataScopeCreateRequest) (*SystemRoleDataScopeCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemRoleDataScopeCreate not implemented")
}
func (UnimplementedSystemRoleServiceServer) SystemRoleDataScope(context.Context, *SystemRoleDataScopeRequest) (*SystemRoleDataScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemRoleDataScope not implemented")
}
func (UnimplementedSystemRoleServiceServer) mustEmbedUnimplementedSystemRoleServiceServer() {}

// UnsafeSystemRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemRoleServiceServer will
// result in compilation errors.
type UnsafeSystemRoleServiceServer interface {
	mustEmbedUnimplementedSystemRoleServiceServer()
}

func RegisterSystemRoleServiceServer(s grpc.ServiceRegistrar, srv SystemRoleServiceServer) {
	s.RegisterService(&SystemRoleService_ServiceDesc, srv)
}

func _SystemRoleService_SystemRoleCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemRoleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRoleServiceServer).SystemRoleCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRoleService_SystemRoleCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRoleServiceServer).SystemRoleCreate(ctx, req.(*SystemRoleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRoleService_SystemRoleUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemRoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRoleServiceServer).SystemRoleUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRoleService_SystemRoleUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRoleServiceServer).SystemRoleUpdate(ctx, req.(*SystemRoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRoleService_SystemRoleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemRoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRoleServiceServer).SystemRoleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRoleService_SystemRoleDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRoleServiceServer).SystemRoleDelete(ctx, req.(*SystemRoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRoleService_SystemRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRoleServiceServer).SystemRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRoleService_SystemRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRoleServiceServer).SystemRole(ctx, req.(*SystemRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRoleService_SystemRoleRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemRoleRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRoleServiceServer).SystemRoleRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRoleService_SystemRoleRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRoleServiceServer).SystemRoleRecover(ctx, req.(*SystemRoleRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRoleService_SystemRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemRoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRoleServiceServer).SystemRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRoleService_SystemRoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRoleServiceServer).SystemRoleList(ctx, req.(*SystemRoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRoleService_SystemRoleDataScopeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemRoleDataScopeCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRoleServiceServer).SystemRoleDataScopeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRoleService_SystemRoleDataScopeCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRoleServiceServer).SystemRoleDataScopeCreate(ctx, req.(*SystemRoleDataScopeCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRoleService_SystemRoleDataScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemRoleDataScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRoleServiceServer).SystemRoleDataScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRoleService_SystemRoleDataScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRoleServiceServer).SystemRoleDataScope(ctx, req.(*SystemRoleDataScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemRoleService_ServiceDesc is the grpc.ServiceDesc for SystemRoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemRoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "role.SystemRoleService",
	HandlerType: (*SystemRoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemRoleCreate",
			Handler:    _SystemRoleService_SystemRoleCreate_Handler,
		},
		{
			MethodName: "SystemRoleUpdate",
			Handler:    _SystemRoleService_SystemRoleUpdate_Handler,
		},
		{
			MethodName: "SystemRoleDelete",
			Handler:    _SystemRoleService_SystemRoleDelete_Handler,
		},
		{
			MethodName: "SystemRole",
			Handler:    _SystemRoleService_SystemRole_Handler,
		},
		{
			MethodName: "SystemRoleRecover",
			Handler:    _SystemRoleService_SystemRoleRecover_Handler,
		},
		{
			MethodName: "SystemRoleList",
			Handler:    _SystemRoleService_SystemRoleList_Handler,
		},
		{
			MethodName: "SystemRoleDataScopeCreate",
			Handler:    _SystemRoleService_SystemRoleDataScopeCreate_Handler,
		},
		{
			MethodName: "SystemRoleDataScope",
			Handler:    _SystemRoleService_SystemRoleDataScope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role/system_role.proto",
}
