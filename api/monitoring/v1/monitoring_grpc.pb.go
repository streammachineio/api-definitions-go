// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/monitoring/v1/monitoring.proto

package monitoring

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

// MonitoringServiceClient is the client API for MonitoringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitoringServiceClient interface {
	//
	// will be called from clients such as cli or console, to retrieve entity states
	// and indicate them to users.
	GetEntityStates(ctx context.Context, in *GetEntityStatesRequest, opts ...grpc.CallOption) (MonitoringService_GetEntityStatesClient, error)
	//
	// will be called from entity agents so that they can send the entity states
	// of items they're responsible for to the monitoring service.
	UpdateEntityStates(ctx context.Context, opts ...grpc.CallOption) (MonitoringService_UpdateEntityStatesClient, error)
}

type monitoringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitoringServiceClient(cc grpc.ClientConnInterface) MonitoringServiceClient {
	return &monitoringServiceClient{cc}
}

func (c *monitoringServiceClient) GetEntityStates(ctx context.Context, in *GetEntityStatesRequest, opts ...grpc.CallOption) (MonitoringService_GetEntityStatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &MonitoringService_ServiceDesc.Streams[0], "/strmprivacy.api.monitoring.v1.MonitoringService/GetEntityStates", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitoringServiceGetEntityStatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MonitoringService_GetEntityStatesClient interface {
	Recv() (*GetEntityStatesResponse, error)
	grpc.ClientStream
}

type monitoringServiceGetEntityStatesClient struct {
	grpc.ClientStream
}

func (x *monitoringServiceGetEntityStatesClient) Recv() (*GetEntityStatesResponse, error) {
	m := new(GetEntityStatesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *monitoringServiceClient) UpdateEntityStates(ctx context.Context, opts ...grpc.CallOption) (MonitoringService_UpdateEntityStatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &MonitoringService_ServiceDesc.Streams[1], "/strmprivacy.api.monitoring.v1.MonitoringService/UpdateEntityStates", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitoringServiceUpdateEntityStatesClient{stream}
	return x, nil
}

type MonitoringService_UpdateEntityStatesClient interface {
	Send(*UpdateEntityStatesRequest) error
	CloseAndRecv() (*UpdateEntityStatesResponse, error)
	grpc.ClientStream
}

type monitoringServiceUpdateEntityStatesClient struct {
	grpc.ClientStream
}

func (x *monitoringServiceUpdateEntityStatesClient) Send(m *UpdateEntityStatesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *monitoringServiceUpdateEntityStatesClient) CloseAndRecv() (*UpdateEntityStatesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpdateEntityStatesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MonitoringServiceServer is the server API for MonitoringService service.
// All implementations must embed UnimplementedMonitoringServiceServer
// for forward compatibility
type MonitoringServiceServer interface {
	//
	// will be called from clients such as cli or console, to retrieve entity states
	// and indicate them to users.
	GetEntityStates(*GetEntityStatesRequest, MonitoringService_GetEntityStatesServer) error
	//
	// will be called from entity agents so that they can send the entity states
	// of items they're responsible for to the monitoring service.
	UpdateEntityStates(MonitoringService_UpdateEntityStatesServer) error
	mustEmbedUnimplementedMonitoringServiceServer()
}

// UnimplementedMonitoringServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMonitoringServiceServer struct {
}

func (UnimplementedMonitoringServiceServer) GetEntityStates(*GetEntityStatesRequest, MonitoringService_GetEntityStatesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEntityStates not implemented")
}
func (UnimplementedMonitoringServiceServer) UpdateEntityStates(MonitoringService_UpdateEntityStatesServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateEntityStates not implemented")
}
func (UnimplementedMonitoringServiceServer) mustEmbedUnimplementedMonitoringServiceServer() {}

// UnsafeMonitoringServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitoringServiceServer will
// result in compilation errors.
type UnsafeMonitoringServiceServer interface {
	mustEmbedUnimplementedMonitoringServiceServer()
}

func RegisterMonitoringServiceServer(s grpc.ServiceRegistrar, srv MonitoringServiceServer) {
	s.RegisterService(&MonitoringService_ServiceDesc, srv)
}

func _MonitoringService_GetEntityStates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEntityStatesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitoringServiceServer).GetEntityStates(m, &monitoringServiceGetEntityStatesServer{stream})
}

type MonitoringService_GetEntityStatesServer interface {
	Send(*GetEntityStatesResponse) error
	grpc.ServerStream
}

type monitoringServiceGetEntityStatesServer struct {
	grpc.ServerStream
}

func (x *monitoringServiceGetEntityStatesServer) Send(m *GetEntityStatesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MonitoringService_UpdateEntityStates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MonitoringServiceServer).UpdateEntityStates(&monitoringServiceUpdateEntityStatesServer{stream})
}

type MonitoringService_UpdateEntityStatesServer interface {
	SendAndClose(*UpdateEntityStatesResponse) error
	Recv() (*UpdateEntityStatesRequest, error)
	grpc.ServerStream
}

type monitoringServiceUpdateEntityStatesServer struct {
	grpc.ServerStream
}

func (x *monitoringServiceUpdateEntityStatesServer) SendAndClose(m *UpdateEntityStatesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *monitoringServiceUpdateEntityStatesServer) Recv() (*UpdateEntityStatesRequest, error) {
	m := new(UpdateEntityStatesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MonitoringService_ServiceDesc is the grpc.ServiceDesc for MonitoringService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MonitoringService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.monitoring.v1.MonitoringService",
	HandlerType: (*MonitoringServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEntityStates",
			Handler:       _MonitoringService_GetEntityStates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateEntityStates",
			Handler:       _MonitoringService_UpdateEntityStates_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "strmprivacy/api/monitoring/v1/monitoring.proto",
}
