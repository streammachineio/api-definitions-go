// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: strmprivacy/api/agents/v1/batch_exporters_agent_v1.proto

package agents

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/strmprivacy/api-definitions-go/v2/api/entities/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDesiredBatchExportersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDesiredBatchExportersRequest) Reset() {
	*x = GetDesiredBatchExportersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDesiredBatchExportersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDesiredBatchExportersRequest) ProtoMessage() {}

func (x *GetDesiredBatchExportersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDesiredBatchExportersRequest.ProtoReflect.Descriptor instead.
func (*GetDesiredBatchExportersRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDescGZIP(), []int{0}
}

type GetDesiredBatchExportersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exporters []*v1.ExtendedBatchExporter `protobuf:"bytes,1,rep,name=exporters,proto3" json:"exporters,omitempty"`
}

func (x *GetDesiredBatchExportersResponse) Reset() {
	*x = GetDesiredBatchExportersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDesiredBatchExportersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDesiredBatchExportersResponse) ProtoMessage() {}

func (x *GetDesiredBatchExportersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDesiredBatchExportersResponse.ProtoReflect.Descriptor instead.
func (*GetDesiredBatchExportersResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GetDesiredBatchExportersResponse) GetExporters() []*v1.ExtendedBatchExporter {
	if x != nil {
		return x.Exporters
	}
	return nil
}

var File_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDesc = []byte{
	0x0a, 0x38, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a,
	0x1f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x3a, 0x03, 0xf8, 0x42, 0x00, 0x22, 0x74, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69,
	0x72, 0x65, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x65, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x32, 0xb2, 0x01, 0x0a, 0x1a,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x63, 0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDescData = file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDesc
)

func file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDescData
}

var file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_goTypes = []interface{}{
	(*GetDesiredBatchExportersRequest)(nil),  // 0: strmprivacy.api.agents.v1.GetDesiredBatchExportersRequest
	(*GetDesiredBatchExportersResponse)(nil), // 1: strmprivacy.api.agents.v1.GetDesiredBatchExportersResponse
	(*v1.ExtendedBatchExporter)(nil),         // 2: strmprivacy.api.entities.v1.ExtendedBatchExporter
}
var file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_depIdxs = []int32{
	2, // 0: strmprivacy.api.agents.v1.GetDesiredBatchExportersResponse.exporters:type_name -> strmprivacy.api.entities.v1.ExtendedBatchExporter
	0, // 1: strmprivacy.api.agents.v1.BatchExportersAgentService.GetDesiredBatchExporters:input_type -> strmprivacy.api.agents.v1.GetDesiredBatchExportersRequest
	1, // 2: strmprivacy.api.agents.v1.BatchExportersAgentService.GetDesiredBatchExporters:output_type -> strmprivacy.api.agents.v1.GetDesiredBatchExportersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_init() }
func file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_init() {
	if File_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDesiredBatchExportersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDesiredBatchExportersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto = out.File
	file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_rawDesc = nil
	file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_goTypes = nil
	file_strmprivacy_api_agents_v1_batch_exporters_agent_v1_proto_depIdxs = nil
}
