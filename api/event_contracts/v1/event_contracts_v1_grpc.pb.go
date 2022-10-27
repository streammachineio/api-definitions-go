// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/event_contracts/v1/event_contracts_v1.proto

package event_contracts

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

// EventContractsServiceClient is the client API for EventContractsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventContractsServiceClient interface {
	ListEventContracts(ctx context.Context, in *ListEventContractsRequest, opts ...grpc.CallOption) (*ListEventContractsResponse, error)
	GetEventContract(ctx context.Context, in *GetEventContractRequest, opts ...grpc.CallOption) (*GetEventContractResponse, error)
	CreateEventContract(ctx context.Context, in *CreateEventContractRequest, opts ...grpc.CallOption) (*CreateEventContractResponse, error)
	UpdateEventContract(ctx context.Context, in *UpdateEventContractRequest, opts ...grpc.CallOption) (*UpdateEventContractResponse, error)
	ActivateEventContract(ctx context.Context, in *ActivateEventContractRequest, opts ...grpc.CallOption) (*ActivateEventContractResponse, error)
	DeleteEventContract(ctx context.Context, in *DeleteEventContractRequest, opts ...grpc.CallOption) (*DeleteEventContractResponse, error)
	ArchiveEventContract(ctx context.Context, in *ArchiveEventContractRequest, opts ...grpc.CallOption) (*ArchiveEventContractResponse, error)
	ValidateMaskedFields(ctx context.Context, in *ValidateMaskedFieldsRequest, opts ...grpc.CallOption) (*ValidateMaskedFieldsResponse, error)
	GetEventContractAndSchema(ctx context.Context, in *GetEventContractAndSchemaRequest, opts ...grpc.CallOption) (*GetEventContractAndSchemaResponse, error)
}

type eventContractsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventContractsServiceClient(cc grpc.ClientConnInterface) EventContractsServiceClient {
	return &eventContractsServiceClient{cc}
}

func (c *eventContractsServiceClient) ListEventContracts(ctx context.Context, in *ListEventContractsRequest, opts ...grpc.CallOption) (*ListEventContractsResponse, error) {
	out := new(ListEventContractsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.event_contracts.v1.EventContractsService/ListEventContracts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventContractsServiceClient) GetEventContract(ctx context.Context, in *GetEventContractRequest, opts ...grpc.CallOption) (*GetEventContractResponse, error) {
	out := new(GetEventContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.event_contracts.v1.EventContractsService/GetEventContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventContractsServiceClient) CreateEventContract(ctx context.Context, in *CreateEventContractRequest, opts ...grpc.CallOption) (*CreateEventContractResponse, error) {
	out := new(CreateEventContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.event_contracts.v1.EventContractsService/CreateEventContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventContractsServiceClient) UpdateEventContract(ctx context.Context, in *UpdateEventContractRequest, opts ...grpc.CallOption) (*UpdateEventContractResponse, error) {
	out := new(UpdateEventContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.event_contracts.v1.EventContractsService/UpdateEventContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventContractsServiceClient) ActivateEventContract(ctx context.Context, in *ActivateEventContractRequest, opts ...grpc.CallOption) (*ActivateEventContractResponse, error) {
	out := new(ActivateEventContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.event_contracts.v1.EventContractsService/ActivateEventContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventContractsServiceClient) DeleteEventContract(ctx context.Context, in *DeleteEventContractRequest, opts ...grpc.CallOption) (*DeleteEventContractResponse, error) {
	out := new(DeleteEventContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.event_contracts.v1.EventContractsService/DeleteEventContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventContractsServiceClient) ArchiveEventContract(ctx context.Context, in *ArchiveEventContractRequest, opts ...grpc.CallOption) (*ArchiveEventContractResponse, error) {
	out := new(ArchiveEventContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.event_contracts.v1.EventContractsService/ArchiveEventContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventContractsServiceClient) ValidateMaskedFields(ctx context.Context, in *ValidateMaskedFieldsRequest, opts ...grpc.CallOption) (*ValidateMaskedFieldsResponse, error) {
	out := new(ValidateMaskedFieldsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.event_contracts.v1.EventContractsService/ValidateMaskedFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventContractsServiceClient) GetEventContractAndSchema(ctx context.Context, in *GetEventContractAndSchemaRequest, opts ...grpc.CallOption) (*GetEventContractAndSchemaResponse, error) {
	out := new(GetEventContractAndSchemaResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.event_contracts.v1.EventContractsService/GetEventContractAndSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventContractsServiceServer is the server API for EventContractsService service.
// All implementations must embed UnimplementedEventContractsServiceServer
// for forward compatibility
type EventContractsServiceServer interface {
	ListEventContracts(context.Context, *ListEventContractsRequest) (*ListEventContractsResponse, error)
	GetEventContract(context.Context, *GetEventContractRequest) (*GetEventContractResponse, error)
	CreateEventContract(context.Context, *CreateEventContractRequest) (*CreateEventContractResponse, error)
	UpdateEventContract(context.Context, *UpdateEventContractRequest) (*UpdateEventContractResponse, error)
	ActivateEventContract(context.Context, *ActivateEventContractRequest) (*ActivateEventContractResponse, error)
	DeleteEventContract(context.Context, *DeleteEventContractRequest) (*DeleteEventContractResponse, error)
	ArchiveEventContract(context.Context, *ArchiveEventContractRequest) (*ArchiveEventContractResponse, error)
	ValidateMaskedFields(context.Context, *ValidateMaskedFieldsRequest) (*ValidateMaskedFieldsResponse, error)
	GetEventContractAndSchema(context.Context, *GetEventContractAndSchemaRequest) (*GetEventContractAndSchemaResponse, error)
	mustEmbedUnimplementedEventContractsServiceServer()
}

// UnimplementedEventContractsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventContractsServiceServer struct {
}

func (UnimplementedEventContractsServiceServer) ListEventContracts(context.Context, *ListEventContractsRequest) (*ListEventContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventContracts not implemented")
}
func (UnimplementedEventContractsServiceServer) GetEventContract(context.Context, *GetEventContractRequest) (*GetEventContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventContract not implemented")
}
func (UnimplementedEventContractsServiceServer) CreateEventContract(context.Context, *CreateEventContractRequest) (*CreateEventContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEventContract not implemented")
}
func (UnimplementedEventContractsServiceServer) UpdateEventContract(context.Context, *UpdateEventContractRequest) (*UpdateEventContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEventContract not implemented")
}
func (UnimplementedEventContractsServiceServer) ActivateEventContract(context.Context, *ActivateEventContractRequest) (*ActivateEventContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateEventContract not implemented")
}
func (UnimplementedEventContractsServiceServer) DeleteEventContract(context.Context, *DeleteEventContractRequest) (*DeleteEventContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEventContract not implemented")
}
func (UnimplementedEventContractsServiceServer) ArchiveEventContract(context.Context, *ArchiveEventContractRequest) (*ArchiveEventContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveEventContract not implemented")
}
func (UnimplementedEventContractsServiceServer) ValidateMaskedFields(context.Context, *ValidateMaskedFieldsRequest) (*ValidateMaskedFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateMaskedFields not implemented")
}
func (UnimplementedEventContractsServiceServer) GetEventContractAndSchema(context.Context, *GetEventContractAndSchemaRequest) (*GetEventContractAndSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventContractAndSchema not implemented")
}
func (UnimplementedEventContractsServiceServer) mustEmbedUnimplementedEventContractsServiceServer() {}

// UnsafeEventContractsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventContractsServiceServer will
// result in compilation errors.
type UnsafeEventContractsServiceServer interface {
	mustEmbedUnimplementedEventContractsServiceServer()
}

func RegisterEventContractsServiceServer(s grpc.ServiceRegistrar, srv EventContractsServiceServer) {
	s.RegisterService(&EventContractsService_ServiceDesc, srv)
}

func _EventContractsService_ListEventContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventContractsServiceServer).ListEventContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.event_contracts.v1.EventContractsService/ListEventContracts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventContractsServiceServer).ListEventContracts(ctx, req.(*ListEventContractsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventContractsService_GetEventContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventContractsServiceServer).GetEventContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.event_contracts.v1.EventContractsService/GetEventContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventContractsServiceServer).GetEventContract(ctx, req.(*GetEventContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventContractsService_CreateEventContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventContractsServiceServer).CreateEventContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.event_contracts.v1.EventContractsService/CreateEventContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventContractsServiceServer).CreateEventContract(ctx, req.(*CreateEventContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventContractsService_UpdateEventContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventContractsServiceServer).UpdateEventContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.event_contracts.v1.EventContractsService/UpdateEventContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventContractsServiceServer).UpdateEventContract(ctx, req.(*UpdateEventContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventContractsService_ActivateEventContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateEventContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventContractsServiceServer).ActivateEventContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.event_contracts.v1.EventContractsService/ActivateEventContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventContractsServiceServer).ActivateEventContract(ctx, req.(*ActivateEventContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventContractsService_DeleteEventContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventContractsServiceServer).DeleteEventContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.event_contracts.v1.EventContractsService/DeleteEventContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventContractsServiceServer).DeleteEventContract(ctx, req.(*DeleteEventContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventContractsService_ArchiveEventContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveEventContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventContractsServiceServer).ArchiveEventContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.event_contracts.v1.EventContractsService/ArchiveEventContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventContractsServiceServer).ArchiveEventContract(ctx, req.(*ArchiveEventContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventContractsService_ValidateMaskedFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateMaskedFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventContractsServiceServer).ValidateMaskedFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.event_contracts.v1.EventContractsService/ValidateMaskedFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventContractsServiceServer).ValidateMaskedFields(ctx, req.(*ValidateMaskedFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventContractsService_GetEventContractAndSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventContractAndSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventContractsServiceServer).GetEventContractAndSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.event_contracts.v1.EventContractsService/GetEventContractAndSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventContractsServiceServer).GetEventContractAndSchema(ctx, req.(*GetEventContractAndSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventContractsService_ServiceDesc is the grpc.ServiceDesc for EventContractsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventContractsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.event_contracts.v1.EventContractsService",
	HandlerType: (*EventContractsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEventContracts",
			Handler:    _EventContractsService_ListEventContracts_Handler,
		},
		{
			MethodName: "GetEventContract",
			Handler:    _EventContractsService_GetEventContract_Handler,
		},
		{
			MethodName: "CreateEventContract",
			Handler:    _EventContractsService_CreateEventContract_Handler,
		},
		{
			MethodName: "UpdateEventContract",
			Handler:    _EventContractsService_UpdateEventContract_Handler,
		},
		{
			MethodName: "ActivateEventContract",
			Handler:    _EventContractsService_ActivateEventContract_Handler,
		},
		{
			MethodName: "DeleteEventContract",
			Handler:    _EventContractsService_DeleteEventContract_Handler,
		},
		{
			MethodName: "ArchiveEventContract",
			Handler:    _EventContractsService_ArchiveEventContract_Handler,
		},
		{
			MethodName: "ValidateMaskedFields",
			Handler:    _EventContractsService_ValidateMaskedFields_Handler,
		},
		{
			MethodName: "GetEventContractAndSchema",
			Handler:    _EventContractsService_GetEventContractAndSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/event_contracts/v1/event_contracts_v1.proto",
}
