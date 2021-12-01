// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// SimpleClient is the client API for Simple service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleClient interface {
	Getinfo(ctx context.Context, in *Simplerequest, opts ...grpc.CallOption) (*Simpleresponse, error)
}

type simpleClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleClient(cc grpc.ClientConnInterface) SimpleClient {
	return &simpleClient{cc}
}

func (c *simpleClient) Getinfo(ctx context.Context, in *Simplerequest, opts ...grpc.CallOption) (*Simpleresponse, error) {
	out := new(Simpleresponse)
	err := c.cc.Invoke(ctx, "/main.simple/getinfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleServer is the server API for Simple service.
// All implementations must embed UnimplementedSimpleServer
// for forward compatibility
type SimpleServer interface {
	Getinfo(context.Context, *Simplerequest) (*Simpleresponse, error)
	mustEmbedUnimplementedSimpleServer()
}

// UnimplementedSimpleServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleServer struct {
}

func (UnimplementedSimpleServer) Getinfo(context.Context, *Simplerequest) (*Simpleresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getinfo not implemented")
}
func (UnimplementedSimpleServer) mustEmbedUnimplementedSimpleServer() {}

// UnsafeSimpleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleServer will
// result in compilation errors.
type UnsafeSimpleServer interface {
	mustEmbedUnimplementedSimpleServer()
}

func RegisterSimpleServer(s grpc.ServiceRegistrar, srv SimpleServer) {
	s.RegisterService(&Simple_ServiceDesc, srv)
}

func _Simple_Getinfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Simplerequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServer).Getinfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.simple/getinfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServer).Getinfo(ctx, req.(*Simplerequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Simple_ServiceDesc is the grpc.ServiceDesc for Simple service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Simple_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.simple",
	HandlerType: (*SimpleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getinfo",
			Handler:    _Simple_Getinfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/simple.proto",
}