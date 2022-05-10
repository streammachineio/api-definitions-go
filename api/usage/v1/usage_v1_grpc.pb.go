// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.18.1
// source: strmprivacy/api/usage/v1/usage_v1.proto

package usage

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

// UsageServiceClient is the client API for UsageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsageServiceClient interface {
	GetStreamEventUsage(ctx context.Context, in *GetStreamEventUsageRequest, opts ...grpc.CallOption) (*GetStreamEventUsageResponse, error)
	StoreUsageEvent(ctx context.Context, in *StoreUsageEventRequest, opts ...grpc.CallOption) (*StoreUsageEventResponse, error)
}

type usageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsageServiceClient(cc grpc.ClientConnInterface) UsageServiceClient {
	return &usageServiceClient{cc}
}

func (c *usageServiceClient) GetStreamEventUsage(ctx context.Context, in *GetStreamEventUsageRequest, opts ...grpc.CallOption) (*GetStreamEventUsageResponse, error) {
	out := new(GetStreamEventUsageResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.usage.v1.UsageService/GetStreamEventUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usageServiceClient) StoreUsageEvent(ctx context.Context, in *StoreUsageEventRequest, opts ...grpc.CallOption) (*StoreUsageEventResponse, error) {
	out := new(StoreUsageEventResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.usage.v1.UsageService/StoreUsageEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsageServiceServer is the server API for UsageService service.
// All implementations must embed UnimplementedUsageServiceServer
// for forward compatibility
type UsageServiceServer interface {
	GetStreamEventUsage(context.Context, *GetStreamEventUsageRequest) (*GetStreamEventUsageResponse, error)
	StoreUsageEvent(context.Context, *StoreUsageEventRequest) (*StoreUsageEventResponse, error)
	mustEmbedUnimplementedUsageServiceServer()
}

// UnimplementedUsageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsageServiceServer struct {
}

func (UnimplementedUsageServiceServer) GetStreamEventUsage(context.Context, *GetStreamEventUsageRequest) (*GetStreamEventUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamEventUsage not implemented")
}
func (UnimplementedUsageServiceServer) StoreUsageEvent(context.Context, *StoreUsageEventRequest) (*StoreUsageEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreUsageEvent not implemented")
}
func (UnimplementedUsageServiceServer) mustEmbedUnimplementedUsageServiceServer() {}

// UnsafeUsageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsageServiceServer will
// result in compilation errors.
type UnsafeUsageServiceServer interface {
	mustEmbedUnimplementedUsageServiceServer()
}

func RegisterUsageServiceServer(s grpc.ServiceRegistrar, srv UsageServiceServer) {
	s.RegisterService(&UsageService_ServiceDesc, srv)
}

func _UsageService_GetStreamEventUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamEventUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageServiceServer).GetStreamEventUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.usage.v1.UsageService/GetStreamEventUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageServiceServer).GetStreamEventUsage(ctx, req.(*GetStreamEventUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsageService_StoreUsageEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreUsageEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageServiceServer).StoreUsageEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.usage.v1.UsageService/StoreUsageEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageServiceServer).StoreUsageEvent(ctx, req.(*StoreUsageEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsageService_ServiceDesc is the grpc.ServiceDesc for UsageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.usage.v1.UsageService",
	HandlerType: (*UsageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStreamEventUsage",
			Handler:    _UsageService_GetStreamEventUsage_Handler,
		},
		{
			MethodName: "StoreUsageEvent",
			Handler:    _UsageService_StoreUsageEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/usage/v1/usage_v1.proto",
}
