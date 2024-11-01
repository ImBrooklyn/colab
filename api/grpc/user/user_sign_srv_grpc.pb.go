// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: user/user_sign_srv.proto

package user

import (
	common "colab/api/common"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserSignService_Signup_FullMethodName           = "/proto.UserSignService/Signup"
	UserSignService_SigninByMobile_FullMethodName   = "/proto.UserSignService/SigninByMobile"
	UserSignService_SigninByPassword_FullMethodName = "/proto.UserSignService/SigninByPassword"
)

// UserSignServiceClient is the client API for UserSignService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserSignServiceClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*common.Flag, error)
	SigninByMobile(ctx context.Context, in *SigninByMobileRequest, opts ...grpc.CallOption) (*SigninResponse, error)
	SigninByPassword(ctx context.Context, in *SigninByPasswordRequest, opts ...grpc.CallOption) (*SigninResponse, error)
}

type userSignServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSignServiceClient(cc grpc.ClientConnInterface) UserSignServiceClient {
	return &userSignServiceClient{cc}
}

func (c *userSignServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*common.Flag, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Flag)
	err := c.cc.Invoke(ctx, UserSignService_Signup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSignServiceClient) SigninByMobile(ctx context.Context, in *SigninByMobileRequest, opts ...grpc.CallOption) (*SigninResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SigninResponse)
	err := c.cc.Invoke(ctx, UserSignService_SigninByMobile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSignServiceClient) SigninByPassword(ctx context.Context, in *SigninByPasswordRequest, opts ...grpc.CallOption) (*SigninResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SigninResponse)
	err := c.cc.Invoke(ctx, UserSignService_SigninByPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSignServiceServer is the server API for UserSignService service.
// All implementations should embed UnimplementedUserSignServiceServer
// for forward compatibility.
type UserSignServiceServer interface {
	Signup(context.Context, *SignupRequest) (*common.Flag, error)
	SigninByMobile(context.Context, *SigninByMobileRequest) (*SigninResponse, error)
	SigninByPassword(context.Context, *SigninByPasswordRequest) (*SigninResponse, error)
}

// UnimplementedUserSignServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserSignServiceServer struct{}

func (UnimplementedUserSignServiceServer) Signup(context.Context, *SignupRequest) (*common.Flag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedUserSignServiceServer) SigninByMobile(context.Context, *SigninByMobileRequest) (*SigninResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigninByMobile not implemented")
}
func (UnimplementedUserSignServiceServer) SigninByPassword(context.Context, *SigninByPasswordRequest) (*SigninResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigninByPassword not implemented")
}
func (UnimplementedUserSignServiceServer) testEmbeddedByValue() {}

// UnsafeUserSignServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserSignServiceServer will
// resp in compilation errors.
type UnsafeUserSignServiceServer interface {
	mustEmbedUnimplementedUserSignServiceServer()
}

func RegisterUserSignServiceServer(s grpc.ServiceRegistrar, srv UserSignServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserSignServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserSignService_ServiceDesc, srv)
}

func _UserSignService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSignServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSignService_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSignServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSignService_SigninByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SigninByMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSignServiceServer).SigninByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSignService_SigninByMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSignServiceServer).SigninByMobile(ctx, req.(*SigninByMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSignService_SigninByPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SigninByPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSignServiceServer).SigninByPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSignService_SigninByPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSignServiceServer).SigninByPassword(ctx, req.(*SigninByPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserSignService_ServiceDesc is the grpc.ServiceDesc for UserSignService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserSignService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserSignService",
	HandlerType: (*UserSignServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _UserSignService_Signup_Handler,
		},
		{
			MethodName: "SigninByMobile",
			Handler:    _UserSignService_SigninByMobile_Handler,
		},
		{
			MethodName: "SigninByPassword",
			Handler:    _UserSignService_SigninByPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user_sign_srv.proto",
}