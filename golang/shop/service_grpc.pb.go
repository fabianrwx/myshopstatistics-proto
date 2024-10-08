// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: shop/service.proto

package shop

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ShopService_Ping_FullMethodName    = "/ShopService/Ping"
	ShopService_Create_FullMethodName  = "/ShopService/Create"
	ShopService_GetById_FullMethodName = "/ShopService/GetById"
)

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Shop, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, ShopService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, ShopService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, ShopService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations must embed UnimplementedShopServiceServer
// for forward compatibility
type ShopServiceServer interface {
	Ping(context.Context, *emptypb.Empty) (*PingResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	GetById(context.Context, *GetByIdRequest) (*Shop, error)
	mustEmbedUnimplementedShopServiceServer()
}

// UnimplementedShopServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShopServiceServer struct {
}

func (UnimplementedShopServiceServer) Ping(context.Context, *emptypb.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedShopServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedShopServiceServer) GetById(context.Context, *GetByIdRequest) (*Shop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedShopServiceServer) mustEmbedUnimplementedShopServiceServer() {}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ShopService_Ping_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ShopService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ShopService_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop/service.proto",
}
