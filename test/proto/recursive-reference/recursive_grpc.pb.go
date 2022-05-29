// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package recursive_reference

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

// RecursiveClient is the client API for Recursive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecursiveClient interface {
	RecursiveCall(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error)
}

type recursiveClient struct {
	cc grpc.ClientConnInterface
}

func NewRecursiveClient(cc grpc.ClientConnInterface) RecursiveClient {
	return &recursiveClient{cc}
}

func (c *recursiveClient) RecursiveCall(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error) {
	out := new(FooResponse)
	err := c.cc.Invoke(ctx, "/jsonrpc.gateway.test.proto.recursive_reference.Recursive/RecursiveCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecursiveServer is the server API for Recursive service.
// All implementations should embed UnimplementedRecursiveServer
// for forward compatibility
type RecursiveServer interface {
	RecursiveCall(context.Context, *FooRequest) (*FooResponse, error)
}

// UnimplementedRecursiveServer should be embedded to have forward compatible implementations.
type UnimplementedRecursiveServer struct {
}

func (UnimplementedRecursiveServer) RecursiveCall(context.Context, *FooRequest) (*FooResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecursiveCall not implemented")
}

// UnsafeRecursiveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecursiveServer will
// result in compilation errors.
type UnsafeRecursiveServer interface {
	mustEmbedUnimplementedRecursiveServer()
}

func RegisterRecursiveServer(s grpc.ServiceRegistrar, srv RecursiveServer) {
	s.RegisterService(&Recursive_ServiceDesc, srv)
}

func _Recursive_RecursiveCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecursiveServer).RecursiveCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jsonrpc.gateway.test.proto.recursive_reference.Recursive/RecursiveCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecursiveServer).RecursiveCall(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Recursive_ServiceDesc is the grpc.ServiceDesc for Recursive service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Recursive_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jsonrpc.gateway.test.proto.recursive_reference.Recursive",
	HandlerType: (*RecursiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecursiveCall",
			Handler:    _Recursive_RecursiveCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/proto/recursive-reference/recursive.proto",
}
