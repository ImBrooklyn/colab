// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: snowflake/snowflake_srv.proto

package snowflake

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
	SnowflakeService_NextID_FullMethodName = "/proto.SnowflakeService/NextID"
)

// SnowflakeServiceClient is the client API for SnowflakeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnowflakeServiceClient interface {
	NextID(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*SnowflakeID, error)
}

type snowflakeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSnowflakeServiceClient(cc grpc.ClientConnInterface) SnowflakeServiceClient {
	return &snowflakeServiceClient{cc}
}

func (c *snowflakeServiceClient) NextID(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*SnowflakeID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SnowflakeID)
	err := c.cc.Invoke(ctx, SnowflakeService_NextID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnowflakeServiceServer is the server API for SnowflakeService service.
// All implementations should embed UnimplementedSnowflakeServiceServer
// for forward compatibility.
type SnowflakeServiceServer interface {
	NextID(context.Context, *common.Empty) (*SnowflakeID, error)
}

// UnimplementedSnowflakeServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSnowflakeServiceServer struct{}

func (UnimplementedSnowflakeServiceServer) NextID(context.Context, *common.Empty) (*SnowflakeID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextID not implemented")
}
func (UnimplementedSnowflakeServiceServer) testEmbeddedByValue() {}

// UnsafeSnowflakeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnowflakeServiceServer will
// resp in compilation errors.
type UnsafeSnowflakeServiceServer interface {
	mustEmbedUnimplementedSnowflakeServiceServer()
}

func RegisterSnowflakeServiceServer(s grpc.ServiceRegistrar, srv SnowflakeServiceServer) {
	// If the following call pancis, it indicates UnimplementedSnowflakeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SnowflakeService_ServiceDesc, srv)
}

func _SnowflakeService_NextID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnowflakeServiceServer).NextID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnowflakeService_NextID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnowflakeServiceServer).NextID(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SnowflakeService_ServiceDesc is the grpc.ServiceDesc for SnowflakeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SnowflakeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SnowflakeService",
	HandlerType: (*SnowflakeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NextID",
			Handler:    _SnowflakeService_NextID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "snowflake/snowflake_srv.proto",
}