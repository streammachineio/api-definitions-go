// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/credentials/v1/credentials_v1.proto

package credentials

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

// CredentialsServiceClient is the client API for CredentialsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CredentialsServiceClient interface {
	ListCredentials(ctx context.Context, in *ListCredentialsRequest, opts ...grpc.CallOption) (*ListCredentialsResponse, error)
	GetCredential(ctx context.Context, in *GetCredentialRequest, opts ...grpc.CallOption) (*GetCredentialResponse, error)
	CreateCredential(ctx context.Context, in *CreateCredentialRequest, opts ...grpc.CallOption) (*CreateCredentialResponse, error)
	DeleteCredential(ctx context.Context, in *DeleteCredentialRequest, opts ...grpc.CallOption) (*DeleteCredentialResponse, error)
}

type credentialsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCredentialsServiceClient(cc grpc.ClientConnInterface) CredentialsServiceClient {
	return &credentialsServiceClient{cc}
}

func (c *credentialsServiceClient) ListCredentials(ctx context.Context, in *ListCredentialsRequest, opts ...grpc.CallOption) (*ListCredentialsResponse, error) {
	out := new(ListCredentialsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.credentials.v1.CredentialsService/ListCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsServiceClient) GetCredential(ctx context.Context, in *GetCredentialRequest, opts ...grpc.CallOption) (*GetCredentialResponse, error) {
	out := new(GetCredentialResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.credentials.v1.CredentialsService/GetCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsServiceClient) CreateCredential(ctx context.Context, in *CreateCredentialRequest, opts ...grpc.CallOption) (*CreateCredentialResponse, error) {
	out := new(CreateCredentialResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.credentials.v1.CredentialsService/CreateCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsServiceClient) DeleteCredential(ctx context.Context, in *DeleteCredentialRequest, opts ...grpc.CallOption) (*DeleteCredentialResponse, error) {
	out := new(DeleteCredentialResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.credentials.v1.CredentialsService/DeleteCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialsServiceServer is the server API for CredentialsService service.
// All implementations must embed UnimplementedCredentialsServiceServer
// for forward compatibility
type CredentialsServiceServer interface {
	ListCredentials(context.Context, *ListCredentialsRequest) (*ListCredentialsResponse, error)
	GetCredential(context.Context, *GetCredentialRequest) (*GetCredentialResponse, error)
	CreateCredential(context.Context, *CreateCredentialRequest) (*CreateCredentialResponse, error)
	DeleteCredential(context.Context, *DeleteCredentialRequest) (*DeleteCredentialResponse, error)
	mustEmbedUnimplementedCredentialsServiceServer()
}

// UnimplementedCredentialsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCredentialsServiceServer struct {
}

func (UnimplementedCredentialsServiceServer) ListCredentials(context.Context, *ListCredentialsRequest) (*ListCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCredentials not implemented")
}
func (UnimplementedCredentialsServiceServer) GetCredential(context.Context, *GetCredentialRequest) (*GetCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCredential not implemented")
}
func (UnimplementedCredentialsServiceServer) CreateCredential(context.Context, *CreateCredentialRequest) (*CreateCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCredential not implemented")
}
func (UnimplementedCredentialsServiceServer) DeleteCredential(context.Context, *DeleteCredentialRequest) (*DeleteCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCredential not implemented")
}
func (UnimplementedCredentialsServiceServer) mustEmbedUnimplementedCredentialsServiceServer() {}

// UnsafeCredentialsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CredentialsServiceServer will
// result in compilation errors.
type UnsafeCredentialsServiceServer interface {
	mustEmbedUnimplementedCredentialsServiceServer()
}

func RegisterCredentialsServiceServer(s grpc.ServiceRegistrar, srv CredentialsServiceServer) {
	s.RegisterService(&CredentialsService_ServiceDesc, srv)
}

func _CredentialsService_ListCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServiceServer).ListCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.credentials.v1.CredentialsService/ListCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServiceServer).ListCredentials(ctx, req.(*ListCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialsService_GetCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServiceServer).GetCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.credentials.v1.CredentialsService/GetCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServiceServer).GetCredential(ctx, req.(*GetCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialsService_CreateCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServiceServer).CreateCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.credentials.v1.CredentialsService/CreateCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServiceServer).CreateCredential(ctx, req.(*CreateCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialsService_DeleteCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServiceServer).DeleteCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.credentials.v1.CredentialsService/DeleteCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServiceServer).DeleteCredential(ctx, req.(*DeleteCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CredentialsService_ServiceDesc is the grpc.ServiceDesc for CredentialsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CredentialsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.credentials.v1.CredentialsService",
	HandlerType: (*CredentialsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCredentials",
			Handler:    _CredentialsService_ListCredentials_Handler,
		},
		{
			MethodName: "GetCredential",
			Handler:    _CredentialsService_GetCredential_Handler,
		},
		{
			MethodName: "CreateCredential",
			Handler:    _CredentialsService_CreateCredential_Handler,
		},
		{
			MethodName: "DeleteCredential",
			Handler:    _CredentialsService_DeleteCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/credentials/v1/credentials_v1.proto",
}
