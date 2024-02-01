// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: generic.proto

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Generic_InitTracer_FullMethodName = "/proto.Generic/InitTracer"
)

// GenericClient is the client API for Generic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenericClient interface {
	InitTracer(ctx context.Context, in *ServiceName, opts ...grpc.CallOption) (*empty.Empty, error)
}

type genericClient struct {
	cc grpc.ClientConnInterface
}

func NewGenericClient(cc grpc.ClientConnInterface) GenericClient {
	return &genericClient{cc}
}

func (c *genericClient) InitTracer(ctx context.Context, in *ServiceName, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Generic_InitTracer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenericServer is the server API for Generic service.
// All implementations must embed UnimplementedGenericServer
// for forward compatibility
type GenericServer interface {
	InitTracer(context.Context, *ServiceName) (*empty.Empty, error)
	mustEmbedUnimplementedGenericServer()
}

// UnimplementedGenericServer must be embedded to have forward compatible implementations.
type UnimplementedGenericServer struct {
}

func (UnimplementedGenericServer) InitTracer(context.Context, *ServiceName) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitTracer not implemented")
}
func (UnimplementedGenericServer) mustEmbedUnimplementedGenericServer() {}

// UnsafeGenericServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenericServer will
// result in compilation errors.
type UnsafeGenericServer interface {
	mustEmbedUnimplementedGenericServer()
}

func RegisterGenericServer(s grpc.ServiceRegistrar, srv GenericServer) {
	s.RegisterService(&Generic_ServiceDesc, srv)
}

func _Generic_InitTracer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServer).InitTracer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Generic_InitTracer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServer).InitTracer(ctx, req.(*ServiceName))
	}
	return interceptor(ctx, in, info, handler)
}

// Generic_ServiceDesc is the grpc.ServiceDesc for Generic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Generic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Generic",
	HandlerType: (*GenericServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitTracer",
			Handler:    _Generic_InitTracer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generic.proto",
}
