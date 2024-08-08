// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: mailer/service.proto

package mailer

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MailerService_SendVerification_FullMethodName = "/MailerService/SendVerification"
	MailerService_Ping_FullMethodName             = "/MailerService/Ping"
)

// MailerServiceClient is the client API for MailerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Mailer service definition
type MailerServiceClient interface {
	SendVerification(ctx context.Context, in *SendVerificationRequest, opts ...grpc.CallOption) (*SendVerificationResponse, error)
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error)
}

type mailerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMailerServiceClient(cc grpc.ClientConnInterface) MailerServiceClient {
	return &mailerServiceClient{cc}
}

func (c *mailerServiceClient) SendVerification(ctx context.Context, in *SendVerificationRequest, opts ...grpc.CallOption) (*SendVerificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendVerificationResponse)
	err := c.cc.Invoke(ctx, MailerService_SendVerification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailerServiceClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, MailerService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailerServiceServer is the server API for MailerService service.
// All implementations must embed UnimplementedMailerServiceServer
// for forward compatibility.
//
// Mailer service definition
type MailerServiceServer interface {
	SendVerification(context.Context, *SendVerificationRequest) (*SendVerificationResponse, error)
	Ping(context.Context, *empty.Empty) (*PingResponse, error)
	mustEmbedUnimplementedMailerServiceServer()
}

// UnimplementedMailerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMailerServiceServer struct{}

func (UnimplementedMailerServiceServer) SendVerification(context.Context, *SendVerificationRequest) (*SendVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerification not implemented")
}
func (UnimplementedMailerServiceServer) Ping(context.Context, *empty.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedMailerServiceServer) mustEmbedUnimplementedMailerServiceServer() {}
func (UnimplementedMailerServiceServer) testEmbeddedByValue()                       {}

// UnsafeMailerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailerServiceServer will
// result in compilation errors.
type UnsafeMailerServiceServer interface {
	mustEmbedUnimplementedMailerServiceServer()
}

func RegisterMailerServiceServer(s grpc.ServiceRegistrar, srv MailerServiceServer) {
	// If the following call pancis, it indicates UnimplementedMailerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MailerService_ServiceDesc, srv)
}

func _MailerService_SendVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServiceServer).SendVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailerService_SendVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServiceServer).SendVerification(ctx, req.(*SendVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailerService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailerService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServiceServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MailerService_ServiceDesc is the grpc.ServiceDesc for MailerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MailerService",
	HandlerType: (*MailerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendVerification",
			Handler:    _MailerService_SendVerification_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _MailerService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mailer/service.proto",
}
