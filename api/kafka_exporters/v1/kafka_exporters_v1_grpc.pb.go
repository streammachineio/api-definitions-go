// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kafka_exporters

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

// KafkaExportersServiceClient is the client API for KafkaExportersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KafkaExportersServiceClient interface {
	ListKafkaExporters(ctx context.Context, in *ListKafkaExportersRequest, opts ...grpc.CallOption) (*ListKafkaExportersResponse, error)
	GetKafkaExporter(ctx context.Context, in *GetKafkaExporterRequest, opts ...grpc.CallOption) (*GetKafkaExporterResponse, error)
	DeleteKafkaExporter(ctx context.Context, in *DeleteKafkaExporterRequest, opts ...grpc.CallOption) (*DeleteKafkaExporterResponse, error)
	CreateKafkaExporter(ctx context.Context, in *CreateKafkaExporterRequest, opts ...grpc.CallOption) (*CreateKafkaExporterResponse, error)
}

type kafkaExportersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKafkaExportersServiceClient(cc grpc.ClientConnInterface) KafkaExportersServiceClient {
	return &kafkaExportersServiceClient{cc}
}

func (c *kafkaExportersServiceClient) ListKafkaExporters(ctx context.Context, in *ListKafkaExportersRequest, opts ...grpc.CallOption) (*ListKafkaExportersResponse, error) {
	out := new(ListKafkaExportersResponse)
	err := c.cc.Invoke(ctx, "/streammachine.api.kafka_exporters.v1.KafkaExportersService/ListKafkaExporters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaExportersServiceClient) GetKafkaExporter(ctx context.Context, in *GetKafkaExporterRequest, opts ...grpc.CallOption) (*GetKafkaExporterResponse, error) {
	out := new(GetKafkaExporterResponse)
	err := c.cc.Invoke(ctx, "/streammachine.api.kafka_exporters.v1.KafkaExportersService/GetKafkaExporter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaExportersServiceClient) DeleteKafkaExporter(ctx context.Context, in *DeleteKafkaExporterRequest, opts ...grpc.CallOption) (*DeleteKafkaExporterResponse, error) {
	out := new(DeleteKafkaExporterResponse)
	err := c.cc.Invoke(ctx, "/streammachine.api.kafka_exporters.v1.KafkaExportersService/DeleteKafkaExporter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaExportersServiceClient) CreateKafkaExporter(ctx context.Context, in *CreateKafkaExporterRequest, opts ...grpc.CallOption) (*CreateKafkaExporterResponse, error) {
	out := new(CreateKafkaExporterResponse)
	err := c.cc.Invoke(ctx, "/streammachine.api.kafka_exporters.v1.KafkaExportersService/CreateKafkaExporter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaExportersServiceServer is the server API for KafkaExportersService service.
// All implementations must embed UnimplementedKafkaExportersServiceServer
// for forward compatibility
type KafkaExportersServiceServer interface {
	ListKafkaExporters(context.Context, *ListKafkaExportersRequest) (*ListKafkaExportersResponse, error)
	GetKafkaExporter(context.Context, *GetKafkaExporterRequest) (*GetKafkaExporterResponse, error)
	DeleteKafkaExporter(context.Context, *DeleteKafkaExporterRequest) (*DeleteKafkaExporterResponse, error)
	CreateKafkaExporter(context.Context, *CreateKafkaExporterRequest) (*CreateKafkaExporterResponse, error)
	mustEmbedUnimplementedKafkaExportersServiceServer()
}

// UnimplementedKafkaExportersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKafkaExportersServiceServer struct {
}

func (UnimplementedKafkaExportersServiceServer) ListKafkaExporters(context.Context, *ListKafkaExportersRequest) (*ListKafkaExportersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKafkaExporters not implemented")
}
func (UnimplementedKafkaExportersServiceServer) GetKafkaExporter(context.Context, *GetKafkaExporterRequest) (*GetKafkaExporterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKafkaExporter not implemented")
}
func (UnimplementedKafkaExportersServiceServer) DeleteKafkaExporter(context.Context, *DeleteKafkaExporterRequest) (*DeleteKafkaExporterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKafkaExporter not implemented")
}
func (UnimplementedKafkaExportersServiceServer) CreateKafkaExporter(context.Context, *CreateKafkaExporterRequest) (*CreateKafkaExporterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKafkaExporter not implemented")
}
func (UnimplementedKafkaExportersServiceServer) mustEmbedUnimplementedKafkaExportersServiceServer() {}

// UnsafeKafkaExportersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KafkaExportersServiceServer will
// result in compilation errors.
type UnsafeKafkaExportersServiceServer interface {
	mustEmbedUnimplementedKafkaExportersServiceServer()
}

func RegisterKafkaExportersServiceServer(s grpc.ServiceRegistrar, srv KafkaExportersServiceServer) {
	s.RegisterService(&KafkaExportersService_ServiceDesc, srv)
}

func _KafkaExportersService_ListKafkaExporters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKafkaExportersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaExportersServiceServer).ListKafkaExporters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streammachine.api.kafka_exporters.v1.KafkaExportersService/ListKafkaExporters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaExportersServiceServer).ListKafkaExporters(ctx, req.(*ListKafkaExportersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaExportersService_GetKafkaExporter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKafkaExporterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaExportersServiceServer).GetKafkaExporter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streammachine.api.kafka_exporters.v1.KafkaExportersService/GetKafkaExporter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaExportersServiceServer).GetKafkaExporter(ctx, req.(*GetKafkaExporterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaExportersService_DeleteKafkaExporter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKafkaExporterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaExportersServiceServer).DeleteKafkaExporter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streammachine.api.kafka_exporters.v1.KafkaExportersService/DeleteKafkaExporter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaExportersServiceServer).DeleteKafkaExporter(ctx, req.(*DeleteKafkaExporterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaExportersService_CreateKafkaExporter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKafkaExporterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaExportersServiceServer).CreateKafkaExporter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streammachine.api.kafka_exporters.v1.KafkaExportersService/CreateKafkaExporter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaExportersServiceServer).CreateKafkaExporter(ctx, req.(*CreateKafkaExporterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KafkaExportersService_ServiceDesc is the grpc.ServiceDesc for KafkaExportersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KafkaExportersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streammachine.api.kafka_exporters.v1.KafkaExportersService",
	HandlerType: (*KafkaExportersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKafkaExporters",
			Handler:    _KafkaExportersService_ListKafkaExporters_Handler,
		},
		{
			MethodName: "GetKafkaExporter",
			Handler:    _KafkaExportersService_GetKafkaExporter_Handler,
		},
		{
			MethodName: "DeleteKafkaExporter",
			Handler:    _KafkaExportersService_DeleteKafkaExporter_Handler,
		},
		{
			MethodName: "CreateKafkaExporter",
			Handler:    _KafkaExportersService_CreateKafkaExporter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "streammachine/api/kafka_exporters/v1/kafka_exporters_v1.proto",
}
