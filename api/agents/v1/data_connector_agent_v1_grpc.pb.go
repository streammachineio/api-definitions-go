// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/agents/v1/data_connector_agent_v1.proto

package agents

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

// DataConnectorsAgentServiceClient is the client API for DataConnectorsAgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataConnectorsAgentServiceClient interface {
	GetDesiredDataConnectors(ctx context.Context, in *GetDesiredDataConnectorsRequest, opts ...grpc.CallOption) (*GetDesiredDataConnectorsResponse, error)
}

type dataConnectorsAgentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataConnectorsAgentServiceClient(cc grpc.ClientConnInterface) DataConnectorsAgentServiceClient {
	return &dataConnectorsAgentServiceClient{cc}
}

func (c *dataConnectorsAgentServiceClient) GetDesiredDataConnectors(ctx context.Context, in *GetDesiredDataConnectorsRequest, opts ...grpc.CallOption) (*GetDesiredDataConnectorsResponse, error) {
	out := new(GetDesiredDataConnectorsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.agents.v1.DataConnectorsAgentService/GetDesiredDataConnectors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataConnectorsAgentServiceServer is the server API for DataConnectorsAgentService service.
// All implementations must embed UnimplementedDataConnectorsAgentServiceServer
// for forward compatibility
type DataConnectorsAgentServiceServer interface {
	GetDesiredDataConnectors(context.Context, *GetDesiredDataConnectorsRequest) (*GetDesiredDataConnectorsResponse, error)
	mustEmbedUnimplementedDataConnectorsAgentServiceServer()
}

// UnimplementedDataConnectorsAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataConnectorsAgentServiceServer struct {
}

func (UnimplementedDataConnectorsAgentServiceServer) GetDesiredDataConnectors(context.Context, *GetDesiredDataConnectorsRequest) (*GetDesiredDataConnectorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDesiredDataConnectors not implemented")
}
func (UnimplementedDataConnectorsAgentServiceServer) mustEmbedUnimplementedDataConnectorsAgentServiceServer() {
}

// UnsafeDataConnectorsAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataConnectorsAgentServiceServer will
// result in compilation errors.
type UnsafeDataConnectorsAgentServiceServer interface {
	mustEmbedUnimplementedDataConnectorsAgentServiceServer()
}

func RegisterDataConnectorsAgentServiceServer(s grpc.ServiceRegistrar, srv DataConnectorsAgentServiceServer) {
	s.RegisterService(&DataConnectorsAgentService_ServiceDesc, srv)
}

func _DataConnectorsAgentService_GetDesiredDataConnectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDesiredDataConnectorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataConnectorsAgentServiceServer).GetDesiredDataConnectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.agents.v1.DataConnectorsAgentService/GetDesiredDataConnectors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataConnectorsAgentServiceServer).GetDesiredDataConnectors(ctx, req.(*GetDesiredDataConnectorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataConnectorsAgentService_ServiceDesc is the grpc.ServiceDesc for DataConnectorsAgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataConnectorsAgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.agents.v1.DataConnectorsAgentService",
	HandlerType: (*DataConnectorsAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDesiredDataConnectors",
			Handler:    _DataConnectorsAgentService_GetDesiredDataConnectors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/agents/v1/data_connector_agent_v1.proto",
}
