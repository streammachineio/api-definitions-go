// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/data_sources/v1alpha/data_sources_v1alpha.proto

package data_sources

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

// DataSourcesServiceClient is the client API for DataSourcesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataSourcesServiceClient interface {
	ListDataSources(ctx context.Context, in *ListDataSourcesRequest, opts ...grpc.CallOption) (*ListDataSourcesResponse, error)
	GetDataSource(ctx context.Context, in *GetDataSourceRequest, opts ...grpc.CallOption) (*GetDataSourceResponse, error)
	// Create or update a data source.
	UpsertDataSource(ctx context.Context, in *UpsertDataSourceRequest, opts ...grpc.CallOption) (*UpsertDataSourceResponse, error)
}

type dataSourcesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataSourcesServiceClient(cc grpc.ClientConnInterface) DataSourcesServiceClient {
	return &dataSourcesServiceClient{cc}
}

func (c *dataSourcesServiceClient) ListDataSources(ctx context.Context, in *ListDataSourcesRequest, opts ...grpc.CallOption) (*ListDataSourcesResponse, error) {
	out := new(ListDataSourcesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_sources.v1alpha.DataSourcesService/ListDataSources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourcesServiceClient) GetDataSource(ctx context.Context, in *GetDataSourceRequest, opts ...grpc.CallOption) (*GetDataSourceResponse, error) {
	out := new(GetDataSourceResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_sources.v1alpha.DataSourcesService/GetDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourcesServiceClient) UpsertDataSource(ctx context.Context, in *UpsertDataSourceRequest, opts ...grpc.CallOption) (*UpsertDataSourceResponse, error) {
	out := new(UpsertDataSourceResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_sources.v1alpha.DataSourcesService/UpsertDataSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataSourcesServiceServer is the server API for DataSourcesService service.
// All implementations should embed UnimplementedDataSourcesServiceServer
// for forward compatibility
type DataSourcesServiceServer interface {
	ListDataSources(context.Context, *ListDataSourcesRequest) (*ListDataSourcesResponse, error)
	GetDataSource(context.Context, *GetDataSourceRequest) (*GetDataSourceResponse, error)
	// Create or update a data source.
	UpsertDataSource(context.Context, *UpsertDataSourceRequest) (*UpsertDataSourceResponse, error)
}

// UnimplementedDataSourcesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDataSourcesServiceServer struct {
}

func (UnimplementedDataSourcesServiceServer) ListDataSources(context.Context, *ListDataSourcesRequest) (*ListDataSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataSources not implemented")
}
func (UnimplementedDataSourcesServiceServer) GetDataSource(context.Context, *GetDataSourceRequest) (*GetDataSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataSource not implemented")
}
func (UnimplementedDataSourcesServiceServer) UpsertDataSource(context.Context, *UpsertDataSourceRequest) (*UpsertDataSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertDataSource not implemented")
}

// UnsafeDataSourcesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataSourcesServiceServer will
// result in compilation errors.
type UnsafeDataSourcesServiceServer interface {
	mustEmbedUnimplementedDataSourcesServiceServer()
}

func RegisterDataSourcesServiceServer(s grpc.ServiceRegistrar, srv DataSourcesServiceServer) {
	s.RegisterService(&DataSourcesService_ServiceDesc, srv)
}

func _DataSourcesService_ListDataSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataSourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourcesServiceServer).ListDataSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_sources.v1alpha.DataSourcesService/ListDataSources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourcesServiceServer).ListDataSources(ctx, req.(*ListDataSourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSourcesService_GetDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourcesServiceServer).GetDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_sources.v1alpha.DataSourcesService/GetDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourcesServiceServer).GetDataSource(ctx, req.(*GetDataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSourcesService_UpsertDataSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertDataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourcesServiceServer).UpsertDataSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_sources.v1alpha.DataSourcesService/UpsertDataSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourcesServiceServer).UpsertDataSource(ctx, req.(*UpsertDataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataSourcesService_ServiceDesc is the grpc.ServiceDesc for DataSourcesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataSourcesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.data_sources.v1alpha.DataSourcesService",
	HandlerType: (*DataSourcesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDataSources",
			Handler:    _DataSourcesService_ListDataSources_Handler,
		},
		{
			MethodName: "GetDataSource",
			Handler:    _DataSourcesService_GetDataSource_Handler,
		},
		{
			MethodName: "UpsertDataSource",
			Handler:    _DataSourcesService_UpsertDataSource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/data_sources/v1alpha/data_sources_v1alpha.proto",
}
