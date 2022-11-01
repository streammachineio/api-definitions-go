// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/key_streams/v1/key_streams.proto

package key_streams

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

// KeyStreamsServiceClient is the client API for KeyStreamsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyStreamsServiceClient interface {
	ListKeyStreams(ctx context.Context, in *ListKeyStreamsRequest, opts ...grpc.CallOption) (*ListKeyStreamsResponse, error)
	GetKeyStream(ctx context.Context, in *GetKeyStreamRequest, opts ...grpc.CallOption) (*GetKeyStreamResponse, error)
}

type keyStreamsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyStreamsServiceClient(cc grpc.ClientConnInterface) KeyStreamsServiceClient {
	return &keyStreamsServiceClient{cc}
}

func (c *keyStreamsServiceClient) ListKeyStreams(ctx context.Context, in *ListKeyStreamsRequest, opts ...grpc.CallOption) (*ListKeyStreamsResponse, error) {
	out := new(ListKeyStreamsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.key_streams.v1.KeyStreamsService/ListKeyStreams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyStreamsServiceClient) GetKeyStream(ctx context.Context, in *GetKeyStreamRequest, opts ...grpc.CallOption) (*GetKeyStreamResponse, error) {
	out := new(GetKeyStreamResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.key_streams.v1.KeyStreamsService/GetKeyStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyStreamsServiceServer is the server API for KeyStreamsService service.
// All implementations must embed UnimplementedKeyStreamsServiceServer
// for forward compatibility
type KeyStreamsServiceServer interface {
	ListKeyStreams(context.Context, *ListKeyStreamsRequest) (*ListKeyStreamsResponse, error)
	GetKeyStream(context.Context, *GetKeyStreamRequest) (*GetKeyStreamResponse, error)
	mustEmbedUnimplementedKeyStreamsServiceServer()
}

// UnimplementedKeyStreamsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyStreamsServiceServer struct {
}

func (UnimplementedKeyStreamsServiceServer) ListKeyStreams(context.Context, *ListKeyStreamsRequest) (*ListKeyStreamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeyStreams not implemented")
}
func (UnimplementedKeyStreamsServiceServer) GetKeyStream(context.Context, *GetKeyStreamRequest) (*GetKeyStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeyStream not implemented")
}
func (UnimplementedKeyStreamsServiceServer) mustEmbedUnimplementedKeyStreamsServiceServer() {}

// UnsafeKeyStreamsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyStreamsServiceServer will
// result in compilation errors.
type UnsafeKeyStreamsServiceServer interface {
	mustEmbedUnimplementedKeyStreamsServiceServer()
}

func RegisterKeyStreamsServiceServer(s grpc.ServiceRegistrar, srv KeyStreamsServiceServer) {
	s.RegisterService(&KeyStreamsService_ServiceDesc, srv)
}

func _KeyStreamsService_ListKeyStreams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeyStreamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyStreamsServiceServer).ListKeyStreams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.key_streams.v1.KeyStreamsService/ListKeyStreams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyStreamsServiceServer).ListKeyStreams(ctx, req.(*ListKeyStreamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyStreamsService_GetKeyStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeyStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyStreamsServiceServer).GetKeyStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.key_streams.v1.KeyStreamsService/GetKeyStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyStreamsServiceServer).GetKeyStream(ctx, req.(*GetKeyStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyStreamsService_ServiceDesc is the grpc.ServiceDesc for KeyStreamsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyStreamsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.key_streams.v1.KeyStreamsService",
	HandlerType: (*KeyStreamsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKeyStreams",
			Handler:    _KeyStreamsService_ListKeyStreams_Handler,
		},
		{
			MethodName: "GetKeyStream",
			Handler:    _KeyStreamsService_GetKeyStream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/key_streams/v1/key_streams.proto",
}
