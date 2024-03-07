// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: tenant/system_tenant_package.proto

// system_tenant_package 租户套餐包

package tenant

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
	SystemTenantPackageService_SystemTenantPackageCreate_FullMethodName  = "/tenant.SystemTenantPackageService/SystemTenantPackageCreate"
	SystemTenantPackageService_SystemTenantPackageUpdate_FullMethodName  = "/tenant.SystemTenantPackageService/SystemTenantPackageUpdate"
	SystemTenantPackageService_SystemTenantPackageDelete_FullMethodName  = "/tenant.SystemTenantPackageService/SystemTenantPackageDelete"
	SystemTenantPackageService_SystemTenantPackage_FullMethodName        = "/tenant.SystemTenantPackageService/SystemTenantPackage"
	SystemTenantPackageService_SystemTenantPackageRecover_FullMethodName = "/tenant.SystemTenantPackageService/SystemTenantPackageRecover"
	SystemTenantPackageService_SystemTenantPackageList_FullMethodName    = "/tenant.SystemTenantPackageService/SystemTenantPackageList"
)

// SystemTenantPackageServiceClient is the client API for SystemTenantPackageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemTenantPackageServiceClient interface {
	SystemTenantPackageCreate(ctx context.Context, in *SystemTenantPackageCreateRequest, opts ...grpc.CallOption) (*SystemTenantPackageCreateResponse, error)
	SystemTenantPackageUpdate(ctx context.Context, in *SystemTenantPackageUpdateRequest, opts ...grpc.CallOption) (*SystemTenantPackageUpdateResponse, error)
	SystemTenantPackageDelete(ctx context.Context, in *SystemTenantPackageDeleteRequest, opts ...grpc.CallOption) (*SystemTenantPackageDeleteResponse, error)
	SystemTenantPackage(ctx context.Context, in *SystemTenantPackageRequest, opts ...grpc.CallOption) (*SystemTenantPackageResponse, error)
	SystemTenantPackageRecover(ctx context.Context, in *SystemTenantPackageRecoverRequest, opts ...grpc.CallOption) (*SystemTenantPackageRecoverResponse, error)
	SystemTenantPackageList(ctx context.Context, in *SystemTenantPackageListRequest, opts ...grpc.CallOption) (*SystemTenantPackageListResponse, error)
}

type systemTenantPackageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemTenantPackageServiceClient(cc grpc.ClientConnInterface) SystemTenantPackageServiceClient {
	return &systemTenantPackageServiceClient{cc}
}

func (c *systemTenantPackageServiceClient) SystemTenantPackageCreate(ctx context.Context, in *SystemTenantPackageCreateRequest, opts ...grpc.CallOption) (*SystemTenantPackageCreateResponse, error) {
	out := new(SystemTenantPackageCreateResponse)
	err := c.cc.Invoke(ctx, SystemTenantPackageService_SystemTenantPackageCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemTenantPackageServiceClient) SystemTenantPackageUpdate(ctx context.Context, in *SystemTenantPackageUpdateRequest, opts ...grpc.CallOption) (*SystemTenantPackageUpdateResponse, error) {
	out := new(SystemTenantPackageUpdateResponse)
	err := c.cc.Invoke(ctx, SystemTenantPackageService_SystemTenantPackageUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemTenantPackageServiceClient) SystemTenantPackageDelete(ctx context.Context, in *SystemTenantPackageDeleteRequest, opts ...grpc.CallOption) (*SystemTenantPackageDeleteResponse, error) {
	out := new(SystemTenantPackageDeleteResponse)
	err := c.cc.Invoke(ctx, SystemTenantPackageService_SystemTenantPackageDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemTenantPackageServiceClient) SystemTenantPackage(ctx context.Context, in *SystemTenantPackageRequest, opts ...grpc.CallOption) (*SystemTenantPackageResponse, error) {
	out := new(SystemTenantPackageResponse)
	err := c.cc.Invoke(ctx, SystemTenantPackageService_SystemTenantPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemTenantPackageServiceClient) SystemTenantPackageRecover(ctx context.Context, in *SystemTenantPackageRecoverRequest, opts ...grpc.CallOption) (*SystemTenantPackageRecoverResponse, error) {
	out := new(SystemTenantPackageRecoverResponse)
	err := c.cc.Invoke(ctx, SystemTenantPackageService_SystemTenantPackageRecover_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemTenantPackageServiceClient) SystemTenantPackageList(ctx context.Context, in *SystemTenantPackageListRequest, opts ...grpc.CallOption) (*SystemTenantPackageListResponse, error) {
	out := new(SystemTenantPackageListResponse)
	err := c.cc.Invoke(ctx, SystemTenantPackageService_SystemTenantPackageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemTenantPackageServiceServer is the server API for SystemTenantPackageService service.
// All implementations must embed UnimplementedSystemTenantPackageServiceServer
// for forward compatibility
type SystemTenantPackageServiceServer interface {
	SystemTenantPackageCreate(context.Context, *SystemTenantPackageCreateRequest) (*SystemTenantPackageCreateResponse, error)
	SystemTenantPackageUpdate(context.Context, *SystemTenantPackageUpdateRequest) (*SystemTenantPackageUpdateResponse, error)
	SystemTenantPackageDelete(context.Context, *SystemTenantPackageDeleteRequest) (*SystemTenantPackageDeleteResponse, error)
	SystemTenantPackage(context.Context, *SystemTenantPackageRequest) (*SystemTenantPackageResponse, error)
	SystemTenantPackageRecover(context.Context, *SystemTenantPackageRecoverRequest) (*SystemTenantPackageRecoverResponse, error)
	SystemTenantPackageList(context.Context, *SystemTenantPackageListRequest) (*SystemTenantPackageListResponse, error)
	mustEmbedUnimplementedSystemTenantPackageServiceServer()
}

// UnimplementedSystemTenantPackageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemTenantPackageServiceServer struct {
}

func (UnimplementedSystemTenantPackageServiceServer) SystemTenantPackageCreate(context.Context, *SystemTenantPackageCreateRequest) (*SystemTenantPackageCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantPackageCreate not implemented")
}
func (UnimplementedSystemTenantPackageServiceServer) SystemTenantPackageUpdate(context.Context, *SystemTenantPackageUpdateRequest) (*SystemTenantPackageUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantPackageUpdate not implemented")
}
func (UnimplementedSystemTenantPackageServiceServer) SystemTenantPackageDelete(context.Context, *SystemTenantPackageDeleteRequest) (*SystemTenantPackageDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantPackageDelete not implemented")
}
func (UnimplementedSystemTenantPackageServiceServer) SystemTenantPackage(context.Context, *SystemTenantPackageRequest) (*SystemTenantPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantPackage not implemented")
}
func (UnimplementedSystemTenantPackageServiceServer) SystemTenantPackageRecover(context.Context, *SystemTenantPackageRecoverRequest) (*SystemTenantPackageRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantPackageRecover not implemented")
}
func (UnimplementedSystemTenantPackageServiceServer) SystemTenantPackageList(context.Context, *SystemTenantPackageListRequest) (*SystemTenantPackageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemTenantPackageList not implemented")
}
func (UnimplementedSystemTenantPackageServiceServer) mustEmbedUnimplementedSystemTenantPackageServiceServer() {
}

// UnsafeSystemTenantPackageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemTenantPackageServiceServer will
// result in compilation errors.
type UnsafeSystemTenantPackageServiceServer interface {
	mustEmbedUnimplementedSystemTenantPackageServiceServer()
}

func RegisterSystemTenantPackageServiceServer(s grpc.ServiceRegistrar, srv SystemTenantPackageServiceServer) {
	s.RegisterService(&SystemTenantPackageService_ServiceDesc, srv)
}

func _SystemTenantPackageService_SystemTenantPackageCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantPackageCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackageCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantPackageService_SystemTenantPackageCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackageCreate(ctx, req.(*SystemTenantPackageCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemTenantPackageService_SystemTenantPackageUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantPackageUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackageUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantPackageService_SystemTenantPackageUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackageUpdate(ctx, req.(*SystemTenantPackageUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemTenantPackageService_SystemTenantPackageDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantPackageDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackageDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantPackageService_SystemTenantPackageDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackageDelete(ctx, req.(*SystemTenantPackageDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemTenantPackageService_SystemTenantPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantPackageService_SystemTenantPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackage(ctx, req.(*SystemTenantPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemTenantPackageService_SystemTenantPackageRecover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantPackageRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackageRecover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantPackageService_SystemTenantPackageRecover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackageRecover(ctx, req.(*SystemTenantPackageRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemTenantPackageService_SystemTenantPackageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemTenantPackageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemTenantPackageService_SystemTenantPackageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemTenantPackageServiceServer).SystemTenantPackageList(ctx, req.(*SystemTenantPackageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemTenantPackageService_ServiceDesc is the grpc.ServiceDesc for SystemTenantPackageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemTenantPackageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tenant.SystemTenantPackageService",
	HandlerType: (*SystemTenantPackageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemTenantPackageCreate",
			Handler:    _SystemTenantPackageService_SystemTenantPackageCreate_Handler,
		},
		{
			MethodName: "SystemTenantPackageUpdate",
			Handler:    _SystemTenantPackageService_SystemTenantPackageUpdate_Handler,
		},
		{
			MethodName: "SystemTenantPackageDelete",
			Handler:    _SystemTenantPackageService_SystemTenantPackageDelete_Handler,
		},
		{
			MethodName: "SystemTenantPackage",
			Handler:    _SystemTenantPackageService_SystemTenantPackage_Handler,
		},
		{
			MethodName: "SystemTenantPackageRecover",
			Handler:    _SystemTenantPackageService_SystemTenantPackageRecover_Handler,
		},
		{
			MethodName: "SystemTenantPackageList",
			Handler:    _SystemTenantPackageService_SystemTenantPackageList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenant/system_tenant_package.proto",
}
