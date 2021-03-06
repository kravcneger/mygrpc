// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mygrpc

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

// MyGrpcClient is the client API for MyGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyGrpcClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*StatusCode, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*StatusCode, error)
	ListUsers(ctx context.Context, in *Query, opts ...grpc.CallOption) (MyGrpc_ListUsersClient, error)
}

type myGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewMyGrpcClient(cc grpc.ClientConnInterface) MyGrpcClient {
	return &myGrpcClient{cc}
}

func (c *myGrpcClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := c.cc.Invoke(ctx, "/mygrpc.MyGrpc/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myGrpcClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := c.cc.Invoke(ctx, "/mygrpc.MyGrpc/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myGrpcClient) ListUsers(ctx context.Context, in *Query, opts ...grpc.CallOption) (MyGrpc_ListUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &MyGrpc_ServiceDesc.Streams[0], "/mygrpc.MyGrpc/ListUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &myGrpcListUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MyGrpc_ListUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type myGrpcListUsersClient struct {
	grpc.ClientStream
}

func (x *myGrpcListUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyGrpcServer is the server API for MyGrpc service.
// All implementations must embed UnimplementedMyGrpcServer
// for forward compatibility
type MyGrpcServer interface {
	CreateUser(context.Context, *User) (*StatusCode, error)
	DeleteUser(context.Context, *User) (*StatusCode, error)
	ListUsers(*Query, MyGrpc_ListUsersServer) error
	mustEmbedUnimplementedMyGrpcServer()
}

// UnimplementedMyGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedMyGrpcServer struct {
}

func (UnimplementedMyGrpcServer) CreateUser(context.Context, *User) (*StatusCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMyGrpcServer) DeleteUser(context.Context, *User) (*StatusCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedMyGrpcServer) ListUsers(*Query, MyGrpc_ListUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedMyGrpcServer) mustEmbedUnimplementedMyGrpcServer() {}

// UnsafeMyGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyGrpcServer will
// result in compilation errors.
type UnsafeMyGrpcServer interface {
	mustEmbedUnimplementedMyGrpcServer()
}

func RegisterMyGrpcServer(s grpc.ServiceRegistrar, srv MyGrpcServer) {
	s.RegisterService(&MyGrpc_ServiceDesc, srv)
}

func _MyGrpc_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyGrpcServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mygrpc.MyGrpc/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyGrpcServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyGrpc_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyGrpcServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mygrpc.MyGrpc/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyGrpcServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyGrpc_ListUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MyGrpcServer).ListUsers(m, &myGrpcListUsersServer{stream})
}

type MyGrpc_ListUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type myGrpcListUsersServer struct {
	grpc.ServerStream
}

func (x *myGrpcListUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

// MyGrpc_ServiceDesc is the grpc.ServiceDesc for MyGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mygrpc.MyGrpc",
	HandlerType: (*MyGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _MyGrpc_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _MyGrpc_DeleteUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUsers",
			Handler:       _MyGrpc_ListUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mygrpc/my_grpc.proto",
}
