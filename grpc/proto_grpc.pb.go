// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: grpc/proto.proto

package proto

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

// ReceiveClient is the client API for Receive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReceiveClient interface {
	Receive(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type receiveClient struct {
	cc grpc.ClientConnInterface
}

func NewReceiveClient(cc grpc.ClientConnInterface) ReceiveClient {
	return &receiveClient{cc}
}

func (c *receiveClient) Receive(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Peer2peer.receive/receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReceiveServer is the server API for Receive service.
// All implementations must embed UnimplementedReceiveServer
// for forward compatibility
type ReceiveServer interface {
	Receive(context.Context, *Request) (*Reply, error)
	mustEmbedUnimplementedReceiveServer()
}

// UnimplementedReceiveServer must be embedded to have forward compatible implementations.
type UnimplementedReceiveServer struct {
}

func (UnimplementedReceiveServer) Receive(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedReceiveServer) mustEmbedUnimplementedReceiveServer() {}

// UnsafeReceiveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReceiveServer will
// result in compilation errors.
type UnsafeReceiveServer interface {
	mustEmbedUnimplementedReceiveServer()
}

func RegisterReceiveServer(s grpc.ServiceRegistrar, srv ReceiveServer) {
	s.RegisterService(&Receive_ServiceDesc, srv)
}

func _Receive_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiveServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Peer2peer.receive/receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiveServer).Receive(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Receive_ServiceDesc is the grpc.ServiceDesc for Receive service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Receive_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Peer2peer.receive",
	HandlerType: (*ReceiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "receive",
			Handler:    _Receive_Receive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
