// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: user/server.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	favorite "tinytiktok/user/proto/favorite"
	info2 "tinytiktok/user/proto/info2"
	login "tinytiktok/user/proto/login"
	publish "tinytiktok/user/proto/publish"
	register "tinytiktok/user/proto/register"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_Register_FullMethodName          = "/userServer.UserService/Register"
	UserService_Login_FullMethodName             = "/userServer.UserService/Login"
	UserService_Info_FullMethodName              = "/userServer.UserService/Info"
	UserService_CalcFavoriteCount_FullMethodName = "/userServer.UserService/CalcFavoriteCount"
	UserService_CalcWorkCount_FullMethodName     = "/userServer.UserService/CalcWorkCount"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Register(ctx context.Context, in *register.RegisterRequest, opts ...grpc.CallOption) (*register.RegisterResponse, error)
	Login(ctx context.Context, in *login.LoginRequest, opts ...grpc.CallOption) (*login.LoginResponse, error)
	Info(ctx context.Context, in *info2.UserRequest, opts ...grpc.CallOption) (*info2.UserResponse, error)
	// 用户点赞或取消点赞
	CalcFavoriteCount(ctx context.Context, in *favorite.CalcFavoriteCountRequest, opts ...grpc.CallOption) (*favorite.CalcFavoriteCountResponse, error)
	// 用户发表/删除视频时需要增加/减少用户的作品数量
	CalcWorkCount(ctx context.Context, in *publish.CalcWorkCountRequest, opts ...grpc.CallOption) (*publish.CalcWorkCountResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *register.RegisterRequest, opts ...grpc.CallOption) (*register.RegisterResponse, error) {
	out := new(register.RegisterResponse)
	err := c.cc.Invoke(ctx, UserService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *login.LoginRequest, opts ...grpc.CallOption) (*login.LoginResponse, error) {
	out := new(login.LoginResponse)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Info(ctx context.Context, in *info2.UserRequest, opts ...grpc.CallOption) (*info2.UserResponse, error) {
	out := new(info2.UserResponse)
	err := c.cc.Invoke(ctx, UserService_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CalcFavoriteCount(ctx context.Context, in *favorite.CalcFavoriteCountRequest, opts ...grpc.CallOption) (*favorite.CalcFavoriteCountResponse, error) {
	out := new(favorite.CalcFavoriteCountResponse)
	err := c.cc.Invoke(ctx, UserService_CalcFavoriteCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CalcWorkCount(ctx context.Context, in *publish.CalcWorkCountRequest, opts ...grpc.CallOption) (*publish.CalcWorkCountResponse, error) {
	out := new(publish.CalcWorkCountResponse)
	err := c.cc.Invoke(ctx, UserService_CalcWorkCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Register(context.Context, *register.RegisterRequest) (*register.RegisterResponse, error)
	Login(context.Context, *login.LoginRequest) (*login.LoginResponse, error)
	Info(context.Context, *info2.UserRequest) (*info2.UserResponse, error)
	// 用户点赞或取消点赞
	CalcFavoriteCount(context.Context, *favorite.CalcFavoriteCountRequest) (*favorite.CalcFavoriteCountResponse, error)
	// 用户发表/删除视频时需要增加/减少用户的作品数量
	CalcWorkCount(context.Context, *publish.CalcWorkCountRequest) (*publish.CalcWorkCountResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *register.RegisterRequest) (*register.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *login.LoginRequest) (*login.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) Info(context.Context, *info2.UserRequest) (*info2.UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedUserServiceServer) CalcFavoriteCount(context.Context, *favorite.CalcFavoriteCountRequest) (*favorite.CalcFavoriteCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalcFavoriteCount not implemented")
}
func (UnimplementedUserServiceServer) CalcWorkCount(context.Context, *publish.CalcWorkCountRequest) (*publish.CalcWorkCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalcWorkCount not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(register.RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*register.RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(login.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*login.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(info2.UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Info(ctx, req.(*info2.UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CalcFavoriteCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(favorite.CalcFavoriteCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CalcFavoriteCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CalcFavoriteCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CalcFavoriteCount(ctx, req.(*favorite.CalcFavoriteCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CalcWorkCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(publish.CalcWorkCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CalcWorkCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CalcWorkCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CalcWorkCount(ctx, req.(*publish.CalcWorkCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userServer.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _UserService_Info_Handler,
		},
		{
			MethodName: "CalcFavoriteCount",
			Handler:    _UserService_CalcFavoriteCount_Handler,
		},
		{
			MethodName: "CalcWorkCount",
			Handler:    _UserService_CalcWorkCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/server.proto",
}
