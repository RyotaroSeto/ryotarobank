// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: service_ryotaro_bank.proto

package pb

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
	RyotaroBank_CreateUser_FullMethodName = "/pb.RyotaroBank/CreateUser"
	RyotaroBank_LoginUser_FullMethodName  = "/pb.RyotaroBank/LoginUser"
)

// RyotaroBankClient is the client API for RyotaroBank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RyotaroBankClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type ryotaroBankClient struct {
	cc grpc.ClientConnInterface
}

func NewRyotaroBankClient(cc grpc.ClientConnInterface) RyotaroBankClient {
	return &ryotaroBankClient{cc}
}

func (c *ryotaroBankClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, RyotaroBank_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ryotaroBankClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, RyotaroBank_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RyotaroBankServer is the server API for RyotaroBank service.
// All implementations must embed UnimplementedRyotaroBankServer
// for forward compatibility
type RyotaroBankServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedRyotaroBankServer()
}

// UnimplementedRyotaroBankServer must be embedded to have forward compatible implementations.
type UnimplementedRyotaroBankServer struct {
}

func (UnimplementedRyotaroBankServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedRyotaroBankServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedRyotaroBankServer) mustEmbedUnimplementedRyotaroBankServer() {}

// UnsafeRyotaroBankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RyotaroBankServer will
// result in compilation errors.
type UnsafeRyotaroBankServer interface {
	mustEmbedUnimplementedRyotaroBankServer()
}

func RegisterRyotaroBankServer(s grpc.ServiceRegistrar, srv RyotaroBankServer) {
	s.RegisterService(&RyotaroBank_ServiceDesc, srv)
}

func _RyotaroBank_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RyotaroBankServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RyotaroBank_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RyotaroBankServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RyotaroBank_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RyotaroBankServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RyotaroBank_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RyotaroBankServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RyotaroBank_ServiceDesc is the grpc.ServiceDesc for RyotaroBank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RyotaroBank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RyotaroBank",
	HandlerType: (*RyotaroBankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _RyotaroBank_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _RyotaroBank_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ryotaro_bank.proto",
}
