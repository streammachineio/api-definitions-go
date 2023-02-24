// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/data_catalog/v1/datasets_v1.proto

package data_catalog

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

// DatasetsServiceClient is the client API for DatasetsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasetsServiceClient interface {
	GetDataset(ctx context.Context, in *GetDatasetRequest, opts ...grpc.CallOption) (*GetDatasetResponse, error)
	GetDatasetSample(ctx context.Context, in *GetDatasetSampleRequest, opts ...grpc.CallOption) (*GetDatasetSampleResponse, error)
	GetDatasetMetrics(ctx context.Context, in *GetDatasetMetricsRequest, opts ...grpc.CallOption) (*GetDatasetMetricsResponse, error)
	ListDatasets(ctx context.Context, in *ListDatasetsRequest, opts ...grpc.CallOption) (*ListDatasetsResponse, error)
	UpdateDatasetTags(ctx context.Context, in *UpdateDatasetTagsRequest, opts ...grpc.CallOption) (*UpdateDatasetTagsResponse, error)
	UpdateSchemaFieldTags(ctx context.Context, in *UpdateSchemaFieldTagsRequest, opts ...grpc.CallOption) (*UpdateSchemaFieldTagsResponse, error)
	CreateDatasetContentsConnection(ctx context.Context, in *CreateDatasetContentsConnectionRequest, opts ...grpc.CallOption) (*CreateDatasetContentsConnectionResponse, error)
}

type datasetsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetsServiceClient(cc grpc.ClientConnInterface) DatasetsServiceClient {
	return &datasetsServiceClient{cc}
}

func (c *datasetsServiceClient) GetDataset(ctx context.Context, in *GetDatasetRequest, opts ...grpc.CallOption) (*GetDatasetResponse, error) {
	out := new(GetDatasetResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_catalog.v1.DatasetsService/GetDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetsServiceClient) GetDatasetSample(ctx context.Context, in *GetDatasetSampleRequest, opts ...grpc.CallOption) (*GetDatasetSampleResponse, error) {
	out := new(GetDatasetSampleResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_catalog.v1.DatasetsService/GetDatasetSample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetsServiceClient) GetDatasetMetrics(ctx context.Context, in *GetDatasetMetricsRequest, opts ...grpc.CallOption) (*GetDatasetMetricsResponse, error) {
	out := new(GetDatasetMetricsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_catalog.v1.DatasetsService/GetDatasetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetsServiceClient) ListDatasets(ctx context.Context, in *ListDatasetsRequest, opts ...grpc.CallOption) (*ListDatasetsResponse, error) {
	out := new(ListDatasetsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_catalog.v1.DatasetsService/ListDatasets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetsServiceClient) UpdateDatasetTags(ctx context.Context, in *UpdateDatasetTagsRequest, opts ...grpc.CallOption) (*UpdateDatasetTagsResponse, error) {
	out := new(UpdateDatasetTagsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_catalog.v1.DatasetsService/UpdateDatasetTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetsServiceClient) UpdateSchemaFieldTags(ctx context.Context, in *UpdateSchemaFieldTagsRequest, opts ...grpc.CallOption) (*UpdateSchemaFieldTagsResponse, error) {
	out := new(UpdateSchemaFieldTagsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_catalog.v1.DatasetsService/UpdateSchemaFieldTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetsServiceClient) CreateDatasetContentsConnection(ctx context.Context, in *CreateDatasetContentsConnectionRequest, opts ...grpc.CallOption) (*CreateDatasetContentsConnectionResponse, error) {
	out := new(CreateDatasetContentsConnectionResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_catalog.v1.DatasetsService/CreateDatasetContentsConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetsServiceServer is the server API for DatasetsService service.
// All implementations should embed UnimplementedDatasetsServiceServer
// for forward compatibility
type DatasetsServiceServer interface {
	GetDataset(context.Context, *GetDatasetRequest) (*GetDatasetResponse, error)
	GetDatasetSample(context.Context, *GetDatasetSampleRequest) (*GetDatasetSampleResponse, error)
	GetDatasetMetrics(context.Context, *GetDatasetMetricsRequest) (*GetDatasetMetricsResponse, error)
	ListDatasets(context.Context, *ListDatasetsRequest) (*ListDatasetsResponse, error)
	UpdateDatasetTags(context.Context, *UpdateDatasetTagsRequest) (*UpdateDatasetTagsResponse, error)
	UpdateSchemaFieldTags(context.Context, *UpdateSchemaFieldTagsRequest) (*UpdateSchemaFieldTagsResponse, error)
	CreateDatasetContentsConnection(context.Context, *CreateDatasetContentsConnectionRequest) (*CreateDatasetContentsConnectionResponse, error)
}

// UnimplementedDatasetsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDatasetsServiceServer struct {
}

func (UnimplementedDatasetsServiceServer) GetDataset(context.Context, *GetDatasetRequest) (*GetDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataset not implemented")
}
func (UnimplementedDatasetsServiceServer) GetDatasetSample(context.Context, *GetDatasetSampleRequest) (*GetDatasetSampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetSample not implemented")
}
func (UnimplementedDatasetsServiceServer) GetDatasetMetrics(context.Context, *GetDatasetMetricsRequest) (*GetDatasetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetMetrics not implemented")
}
func (UnimplementedDatasetsServiceServer) ListDatasets(context.Context, *ListDatasetsRequest) (*ListDatasetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatasets not implemented")
}
func (UnimplementedDatasetsServiceServer) UpdateDatasetTags(context.Context, *UpdateDatasetTagsRequest) (*UpdateDatasetTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDatasetTags not implemented")
}
func (UnimplementedDatasetsServiceServer) UpdateSchemaFieldTags(context.Context, *UpdateSchemaFieldTagsRequest) (*UpdateSchemaFieldTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSchemaFieldTags not implemented")
}
func (UnimplementedDatasetsServiceServer) CreateDatasetContentsConnection(context.Context, *CreateDatasetContentsConnectionRequest) (*CreateDatasetContentsConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDatasetContentsConnection not implemented")
}

// UnsafeDatasetsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasetsServiceServer will
// result in compilation errors.
type UnsafeDatasetsServiceServer interface {
	mustEmbedUnimplementedDatasetsServiceServer()
}

func RegisterDatasetsServiceServer(s grpc.ServiceRegistrar, srv DatasetsServiceServer) {
	s.RegisterService(&DatasetsService_ServiceDesc, srv)
}

func _DatasetsService_GetDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetsServiceServer).GetDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_catalog.v1.DatasetsService/GetDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetsServiceServer).GetDataset(ctx, req.(*GetDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetsService_GetDatasetSample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetSampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetsServiceServer).GetDatasetSample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_catalog.v1.DatasetsService/GetDatasetSample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetsServiceServer).GetDatasetSample(ctx, req.(*GetDatasetSampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetsService_GetDatasetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetsServiceServer).GetDatasetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_catalog.v1.DatasetsService/GetDatasetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetsServiceServer).GetDatasetMetrics(ctx, req.(*GetDatasetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetsService_ListDatasets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatasetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetsServiceServer).ListDatasets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_catalog.v1.DatasetsService/ListDatasets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetsServiceServer).ListDatasets(ctx, req.(*ListDatasetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetsService_UpdateDatasetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDatasetTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetsServiceServer).UpdateDatasetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_catalog.v1.DatasetsService/UpdateDatasetTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetsServiceServer).UpdateDatasetTags(ctx, req.(*UpdateDatasetTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetsService_UpdateSchemaFieldTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSchemaFieldTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetsServiceServer).UpdateSchemaFieldTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_catalog.v1.DatasetsService/UpdateSchemaFieldTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetsServiceServer).UpdateSchemaFieldTags(ctx, req.(*UpdateSchemaFieldTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetsService_CreateDatasetContentsConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatasetContentsConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetsServiceServer).CreateDatasetContentsConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_catalog.v1.DatasetsService/CreateDatasetContentsConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetsServiceServer).CreateDatasetContentsConnection(ctx, req.(*CreateDatasetContentsConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatasetsService_ServiceDesc is the grpc.ServiceDesc for DatasetsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatasetsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.data_catalog.v1.DatasetsService",
	HandlerType: (*DatasetsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDataset",
			Handler:    _DatasetsService_GetDataset_Handler,
		},
		{
			MethodName: "GetDatasetSample",
			Handler:    _DatasetsService_GetDatasetSample_Handler,
		},
		{
			MethodName: "GetDatasetMetrics",
			Handler:    _DatasetsService_GetDatasetMetrics_Handler,
		},
		{
			MethodName: "ListDatasets",
			Handler:    _DatasetsService_ListDatasets_Handler,
		},
		{
			MethodName: "UpdateDatasetTags",
			Handler:    _DatasetsService_UpdateDatasetTags_Handler,
		},
		{
			MethodName: "UpdateSchemaFieldTags",
			Handler:    _DatasetsService_UpdateSchemaFieldTags_Handler,
		},
		{
			MethodName: "CreateDatasetContentsConnection",
			Handler:    _DatasetsService_CreateDatasetContentsConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/data_catalog/v1/datasets_v1.proto",
}
