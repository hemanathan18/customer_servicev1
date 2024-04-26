// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: offMail.proto

package OfficialMail

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
	OffMailService_CreateOffMail_FullMethodName = "/OfficialMail.OffMailService/CreateOffMail"
	OffMailService_GetOffMail_FullMethodName    = "/OfficialMail.OffMailService/GetOffMail"
	OffMailService_UpdateOffMail_FullMethodName = "/OfficialMail.OffMailService/UpdateOffMail"
	OffMailService_DeleteOffMail_FullMethodName = "/OfficialMail.OffMailService/DeleteOffMail"
	OffMailService_GetOffMails_FullMethodName   = "/OfficialMail.OffMailService/GetOffMails"
)

// OffMailServiceClient is the client API for OffMailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OffMailServiceClient interface {
	CreateOffMail(ctx context.Context, in *CreateOffMailRequest, opts ...grpc.CallOption) (*CreateOffMailResponse, error)
	GetOffMail(ctx context.Context, in *GetOffMailRequest, opts ...grpc.CallOption) (*GetOffMailResponse, error)
	UpdateOffMail(ctx context.Context, in *UpdateOffMailRequest, opts ...grpc.CallOption) (*UpdateOffMailResponse, error)
	DeleteOffMail(ctx context.Context, in *DeleteOffMailRequest, opts ...grpc.CallOption) (*DeleteOffMailResponse, error)
	GetOffMails(ctx context.Context, in *GetOffMailsRequest, opts ...grpc.CallOption) (*GetOffMailsResponse, error)
}

type offMailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOffMailServiceClient(cc grpc.ClientConnInterface) OffMailServiceClient {
	return &offMailServiceClient{cc}
}

func (c *offMailServiceClient) CreateOffMail(ctx context.Context, in *CreateOffMailRequest, opts ...grpc.CallOption) (*CreateOffMailResponse, error) {
	out := new(CreateOffMailResponse)
	err := c.cc.Invoke(ctx, OffMailService_CreateOffMail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *offMailServiceClient) GetOffMail(ctx context.Context, in *GetOffMailRequest, opts ...grpc.CallOption) (*GetOffMailResponse, error) {
	out := new(GetOffMailResponse)
	err := c.cc.Invoke(ctx, OffMailService_GetOffMail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *offMailServiceClient) UpdateOffMail(ctx context.Context, in *UpdateOffMailRequest, opts ...grpc.CallOption) (*UpdateOffMailResponse, error) {
	out := new(UpdateOffMailResponse)
	err := c.cc.Invoke(ctx, OffMailService_UpdateOffMail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *offMailServiceClient) DeleteOffMail(ctx context.Context, in *DeleteOffMailRequest, opts ...grpc.CallOption) (*DeleteOffMailResponse, error) {
	out := new(DeleteOffMailResponse)
	err := c.cc.Invoke(ctx, OffMailService_DeleteOffMail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *offMailServiceClient) GetOffMails(ctx context.Context, in *GetOffMailsRequest, opts ...grpc.CallOption) (*GetOffMailsResponse, error) {
	out := new(GetOffMailsResponse)
	err := c.cc.Invoke(ctx, OffMailService_GetOffMails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OffMailServiceServer is the server API for OffMailService service.
// All implementations must embed UnimplementedOffMailServiceServer
// for forward compatibility
type OffMailServiceServer interface {
	CreateOffMail(context.Context, *CreateOffMailRequest) (*CreateOffMailResponse, error)
	GetOffMail(context.Context, *GetOffMailRequest) (*GetOffMailResponse, error)
	UpdateOffMail(context.Context, *UpdateOffMailRequest) (*UpdateOffMailResponse, error)
	DeleteOffMail(context.Context, *DeleteOffMailRequest) (*DeleteOffMailResponse, error)
	GetOffMails(context.Context, *GetOffMailsRequest) (*GetOffMailsResponse, error)
	mustEmbedUnimplementedOffMailServiceServer()
}

// UnimplementedOffMailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOffMailServiceServer struct {
}

func (UnimplementedOffMailServiceServer) CreateOffMail(context.Context, *CreateOffMailRequest) (*CreateOffMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOffMail not implemented")
}
func (UnimplementedOffMailServiceServer) GetOffMail(context.Context, *GetOffMailRequest) (*GetOffMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffMail not implemented")
}
func (UnimplementedOffMailServiceServer) UpdateOffMail(context.Context, *UpdateOffMailRequest) (*UpdateOffMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOffMail not implemented")
}
func (UnimplementedOffMailServiceServer) DeleteOffMail(context.Context, *DeleteOffMailRequest) (*DeleteOffMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOffMail not implemented")
}
func (UnimplementedOffMailServiceServer) GetOffMails(context.Context, *GetOffMailsRequest) (*GetOffMailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffMails not implemented")
}
func (UnimplementedOffMailServiceServer) mustEmbedUnimplementedOffMailServiceServer() {}

// UnsafeOffMailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OffMailServiceServer will
// result in compilation errors.
type UnsafeOffMailServiceServer interface {
	mustEmbedUnimplementedOffMailServiceServer()
}

func RegisterOffMailServiceServer(s grpc.ServiceRegistrar, srv OffMailServiceServer) {
	s.RegisterService(&OffMailService_ServiceDesc, srv)
}

func _OffMailService_CreateOffMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOffMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OffMailServiceServer).CreateOffMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OffMailService_CreateOffMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OffMailServiceServer).CreateOffMail(ctx, req.(*CreateOffMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OffMailService_GetOffMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOffMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OffMailServiceServer).GetOffMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OffMailService_GetOffMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OffMailServiceServer).GetOffMail(ctx, req.(*GetOffMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OffMailService_UpdateOffMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOffMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OffMailServiceServer).UpdateOffMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OffMailService_UpdateOffMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OffMailServiceServer).UpdateOffMail(ctx, req.(*UpdateOffMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OffMailService_DeleteOffMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOffMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OffMailServiceServer).DeleteOffMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OffMailService_DeleteOffMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OffMailServiceServer).DeleteOffMail(ctx, req.(*DeleteOffMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OffMailService_GetOffMails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOffMailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OffMailServiceServer).GetOffMails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OffMailService_GetOffMails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OffMailServiceServer).GetOffMails(ctx, req.(*GetOffMailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OffMailService_ServiceDesc is the grpc.ServiceDesc for OffMailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OffMailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OfficialMail.OffMailService",
	HandlerType: (*OffMailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOffMail",
			Handler:    _OffMailService_CreateOffMail_Handler,
		},
		{
			MethodName: "GetOffMail",
			Handler:    _OffMailService_GetOffMail_Handler,
		},
		{
			MethodName: "UpdateOffMail",
			Handler:    _OffMailService_UpdateOffMail_Handler,
		},
		{
			MethodName: "DeleteOffMail",
			Handler:    _OffMailService_DeleteOffMail_Handler,
		},
		{
			MethodName: "GetOffMails",
			Handler:    _OffMailService_GetOffMails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "offMail.proto",
}