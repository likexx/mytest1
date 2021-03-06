// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tttppp

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

// TtttClient is the client API for Tttt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TtttClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type ttttClient struct {
	cc grpc.ClientConnInterface
}

func NewTtttClient(cc grpc.ClientConnInterface) TtttClient {
	return &ttttClient{cc}
}

func (c *ttttClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/testm.tttppp.tttt/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TtttServer is the server API for Tttt service.
// All implementations must embed UnimplementedTtttServer
// for forward compatibility
type TtttServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedTtttServer()
}

// UnimplementedTtttServer must be embedded to have forward compatible implementations.
type UnimplementedTtttServer struct {
}

func (UnimplementedTtttServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedTtttServer) mustEmbedUnimplementedTtttServer() {}

// UnsafeTtttServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TtttServer will
// result in compilation errors.
type UnsafeTtttServer interface {
	mustEmbedUnimplementedTtttServer()
}

func RegisterTtttServer(s grpc.ServiceRegistrar, srv TtttServer) {
	s.RegisterService(&Tttt_ServiceDesc, srv)
}

func _Tttt_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TtttServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testm.tttppp.tttt/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TtttServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tttt_ServiceDesc is the grpc.ServiceDesc for Tttt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tttt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testm.tttppp.tttt",
	HandlerType: (*TtttServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Tttt_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tttt.proto",
}
