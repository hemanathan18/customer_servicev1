// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: tempAddress.proto

package TempAddress

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
	TempAddressService_CreateTempAddress_FullMethodName = "/TempAddress.TempAddressService/CreateTempAddress"
	TempAddressService_GetTempAddress_FullMethodName    = "/TempAddress.TempAddressService/GetTempAddress"
	TempAddressService_UpdateTempAddress_FullMethodName = "/TempAddress.TempAddressService/UpdateTempAddress"
	TempAddressService_DeleteTempAddress_FullMethodName = "/TempAddress.TempAddressService/DeleteTempAddress"
	TempAddressService_GetTempAddresses_FullMethodName  = "/TempAddress.TempAddressService/GetTempAddresses"
)

// TempAddressServiceClient is the client API for TempAddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TempAddressServiceClient interface {
	CreateTempAddress(ctx context.Context, in *CreateTempAddressRequest, opts ...grpc.CallOption) (*CreateTempAddressResponse, error)
	GetTempAddress(ctx context.Context, in *GetTempAddressRequest, opts ...grpc.CallOption) (*GetTempAddressResponse, error)
	UpdateTempAddress(ctx context.Context, in *UpdateTempAddressRequest, opts ...grpc.CallOption) (*UpdateTempAddressResponse, error)
	DeleteTempAddress(ctx context.Context, in *DeleteTempAddressRequest, opts ...grpc.CallOption) (*DeleteTempAddressResponse, error)
	GetTempAddresses(ctx context.Context, in *GetTempAddressesRequest, opts ...grpc.CallOption) (*GetTempAddressesResponse, error)
}

type tempAddressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTempAddressServiceClient(cc grpc.ClientConnInterface) TempAddressServiceClient {
	return &tempAddressServiceClient{cc}
}

func (c *tempAddressServiceClient) CreateTempAddress(ctx context.Context, in *CreateTempAddressRequest, opts ...grpc.CallOption) (*CreateTempAddressResponse, error) {
	out := new(CreateTempAddressResponse)
	err := c.cc.Invoke(ctx, TempAddressService_CreateTempAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tempAddressServiceClient) GetTempAddress(ctx context.Context, in *GetTempAddressRequest, opts ...grpc.CallOption) (*GetTempAddressResponse, error) {
	out := new(GetTempAddressResponse)
	err := c.cc.Invoke(ctx, TempAddressService_GetTempAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tempAddressServiceClient) UpdateTempAddress(ctx context.Context, in *UpdateTempAddressRequest, opts ...grpc.CallOption) (*UpdateTempAddressResponse, error) {
	out := new(UpdateTempAddressResponse)
	err := c.cc.Invoke(ctx, TempAddressService_UpdateTempAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tempAddressServiceClient) DeleteTempAddress(ctx context.Context, in *DeleteTempAddressRequest, opts ...grpc.CallOption) (*DeleteTempAddressResponse, error) {
	out := new(DeleteTempAddressResponse)
	err := c.cc.Invoke(ctx, TempAddressService_DeleteTempAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tempAddressServiceClient) GetTempAddresses(ctx context.Context, in *GetTempAddressesRequest, opts ...grpc.CallOption) (*GetTempAddressesResponse, error) {
	out := new(GetTempAddressesResponse)
	err := c.cc.Invoke(ctx, TempAddressService_GetTempAddresses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TempAddressServiceServer is the server API for TempAddressService service.
// All implementations must embed UnimplementedTempAddressServiceServer
// for forward compatibility
type TempAddressServiceServer interface {
	CreateTempAddress(context.Context, *CreateTempAddressRequest) (*CreateTempAddressResponse, error)
	GetTempAddress(context.Context, *GetTempAddressRequest) (*GetTempAddressResponse, error)
	UpdateTempAddress(context.Context, *UpdateTempAddressRequest) (*UpdateTempAddressResponse, error)
	DeleteTempAddress(context.Context, *DeleteTempAddressRequest) (*DeleteTempAddressResponse, error)
	GetTempAddresses(context.Context, *GetTempAddressesRequest) (*GetTempAddressesResponse, error)
	mustEmbedUnimplementedTempAddressServiceServer()
}

// UnimplementedTempAddressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTempAddressServiceServer struct {
}

func (UnimplementedTempAddressServiceServer) CreateTempAddress(context.Context, *CreateTempAddressRequest) (*CreateTempAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTempAddress not implemented")
}
func (UnimplementedTempAddressServiceServer) GetTempAddress(context.Context, *GetTempAddressRequest) (*GetTempAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTempAddress not implemented")
}
func (UnimplementedTempAddressServiceServer) UpdateTempAddress(context.Context, *UpdateTempAddressRequest) (*UpdateTempAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTempAddress not implemented")
}
func (UnimplementedTempAddressServiceServer) DeleteTempAddress(context.Context, *DeleteTempAddressRequest) (*DeleteTempAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTempAddress not implemented")
}
func (UnimplementedTempAddressServiceServer) GetTempAddresses(context.Context, *GetTempAddressesRequest) (*GetTempAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTempAddresses not implemented")
}
func (UnimplementedTempAddressServiceServer) mustEmbedUnimplementedTempAddressServiceServer() {}

// UnsafeTempAddressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TempAddressServiceServer will
// result in compilation errors.
type UnsafeTempAddressServiceServer interface {
	mustEmbedUnimplementedTempAddressServiceServer()
}

func RegisterTempAddressServiceServer(s grpc.ServiceRegistrar, srv TempAddressServiceServer) {
	s.RegisterService(&TempAddressService_ServiceDesc, srv)
}

func _TempAddressService_CreateTempAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTempAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TempAddressServiceServer).CreateTempAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TempAddressService_CreateTempAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TempAddressServiceServer).CreateTempAddress(ctx, req.(*CreateTempAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TempAddressService_GetTempAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTempAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TempAddressServiceServer).GetTempAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TempAddressService_GetTempAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TempAddressServiceServer).GetTempAddress(ctx, req.(*GetTempAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TempAddressService_UpdateTempAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTempAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TempAddressServiceServer).UpdateTempAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TempAddressService_UpdateTempAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TempAddressServiceServer).UpdateTempAddress(ctx, req.(*UpdateTempAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TempAddressService_DeleteTempAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTempAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TempAddressServiceServer).DeleteTempAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TempAddressService_DeleteTempAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TempAddressServiceServer).DeleteTempAddress(ctx, req.(*DeleteTempAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TempAddressService_GetTempAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTempAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TempAddressServiceServer).GetTempAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TempAddressService_GetTempAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TempAddressServiceServer).GetTempAddresses(ctx, req.(*GetTempAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TempAddressService_ServiceDesc is the grpc.ServiceDesc for TempAddressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TempAddressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TempAddress.TempAddressService",
	HandlerType: (*TempAddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTempAddress",
			Handler:    _TempAddressService_CreateTempAddress_Handler,
		},
		{
			MethodName: "GetTempAddress",
			Handler:    _TempAddressService_GetTempAddress_Handler,
		},
		{
			MethodName: "UpdateTempAddress",
			Handler:    _TempAddressService_UpdateTempAddress_Handler,
		},
		{
			MethodName: "DeleteTempAddress",
			Handler:    _TempAddressService_DeleteTempAddress_Handler,
		},
		{
			MethodName: "GetTempAddresses",
			Handler:    _TempAddressService_GetTempAddresses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tempAddress.proto",
}