// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.2
// source: streammachine/api/key_streams/v1/key_streams.proto

package key_streams

import (
	v1 "github.com/streammachineio/api-definitions-go/api/entities/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListKeyStreamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
}

func (x *ListKeyStreamsRequest) Reset() {
	*x = ListKeyStreamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKeyStreamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKeyStreamsRequest) ProtoMessage() {}

func (x *ListKeyStreamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKeyStreamsRequest.ProtoReflect.Descriptor instead.
func (*ListKeyStreamsRequest) Descriptor() ([]byte, []int) {
	return file_streammachine_api_key_streams_v1_key_streams_proto_rawDescGZIP(), []int{0}
}

func (x *ListKeyStreamsRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

type ListKeyStreamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyStreams []*v1.KeyStream `protobuf:"bytes,1,rep,name=key_streams,json=keyStreams,proto3" json:"key_streams,omitempty"`
}

func (x *ListKeyStreamsResponse) Reset() {
	*x = ListKeyStreamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKeyStreamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKeyStreamsResponse) ProtoMessage() {}

func (x *ListKeyStreamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKeyStreamsResponse.ProtoReflect.Descriptor instead.
func (*ListKeyStreamsResponse) Descriptor() ([]byte, []int) {
	return file_streammachine_api_key_streams_v1_key_streams_proto_rawDescGZIP(), []int{1}
}

func (x *ListKeyStreamsResponse) GetKeyStreams() []*v1.KeyStream {
	if x != nil {
		return x.KeyStreams
	}
	return nil
}

type GetKeyStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *v1.KeyStreamRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *GetKeyStreamRequest) Reset() {
	*x = GetKeyStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyStreamRequest) ProtoMessage() {}

func (x *GetKeyStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyStreamRequest.ProtoReflect.Descriptor instead.
func (*GetKeyStreamRequest) Descriptor() ([]byte, []int) {
	return file_streammachine_api_key_streams_v1_key_streams_proto_rawDescGZIP(), []int{2}
}

func (x *GetKeyStreamRequest) GetRef() *v1.KeyStreamRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type GetKeyStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyStream *v1.KeyStream `protobuf:"bytes,1,opt,name=key_stream,json=keyStream,proto3" json:"key_stream,omitempty"`
}

func (x *GetKeyStreamResponse) Reset() {
	*x = GetKeyStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyStreamResponse) ProtoMessage() {}

func (x *GetKeyStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyStreamResponse.ProtoReflect.Descriptor instead.
func (*GetKeyStreamResponse) Descriptor() ([]byte, []int) {
	return file_streammachine_api_key_streams_v1_key_streams_proto_rawDescGZIP(), []int{3}
}

func (x *GetKeyStreamResponse) GetKeyStream() *v1.KeyStream {
	if x != nil {
		return x.KeyStream
	}
	return nil
}

var File_streammachine_api_key_streams_v1_key_streams_proto protoreflect.FileDescriptor

var file_streammachine_api_key_streams_v1_key_streams_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x0a,
	0x6b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x59, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x42, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b,
	0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0x5f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a,
	0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x09, 0x6b, 0x65, 0x79,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x32, 0x98, 0x02, 0x0a, 0x11, 0x4b, 0x65, 0x79, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x83, 0x01, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12,
	0x37, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x7d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x35, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65,
	0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x75, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65,
	0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x65, 0x79,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_streammachine_api_key_streams_v1_key_streams_proto_rawDescOnce sync.Once
	file_streammachine_api_key_streams_v1_key_streams_proto_rawDescData = file_streammachine_api_key_streams_v1_key_streams_proto_rawDesc
)

func file_streammachine_api_key_streams_v1_key_streams_proto_rawDescGZIP() []byte {
	file_streammachine_api_key_streams_v1_key_streams_proto_rawDescOnce.Do(func() {
		file_streammachine_api_key_streams_v1_key_streams_proto_rawDescData = protoimpl.X.CompressGZIP(file_streammachine_api_key_streams_v1_key_streams_proto_rawDescData)
	})
	return file_streammachine_api_key_streams_v1_key_streams_proto_rawDescData
}

var file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_streammachine_api_key_streams_v1_key_streams_proto_goTypes = []interface{}{
	(*ListKeyStreamsRequest)(nil),  // 0: streammachine.api.key_streams.v1.ListKeyStreamsRequest
	(*ListKeyStreamsResponse)(nil), // 1: streammachine.api.key_streams.v1.ListKeyStreamsResponse
	(*GetKeyStreamRequest)(nil),    // 2: streammachine.api.key_streams.v1.GetKeyStreamRequest
	(*GetKeyStreamResponse)(nil),   // 3: streammachine.api.key_streams.v1.GetKeyStreamResponse
	(*v1.KeyStream)(nil),           // 4: streammachine.api.entities.v1.KeyStream
	(*v1.KeyStreamRef)(nil),        // 5: streammachine.api.entities.v1.KeyStreamRef
}
var file_streammachine_api_key_streams_v1_key_streams_proto_depIdxs = []int32{
	4, // 0: streammachine.api.key_streams.v1.ListKeyStreamsResponse.key_streams:type_name -> streammachine.api.entities.v1.KeyStream
	5, // 1: streammachine.api.key_streams.v1.GetKeyStreamRequest.ref:type_name -> streammachine.api.entities.v1.KeyStreamRef
	4, // 2: streammachine.api.key_streams.v1.GetKeyStreamResponse.key_stream:type_name -> streammachine.api.entities.v1.KeyStream
	0, // 3: streammachine.api.key_streams.v1.KeyStreamsService.ListKeyStreams:input_type -> streammachine.api.key_streams.v1.ListKeyStreamsRequest
	2, // 4: streammachine.api.key_streams.v1.KeyStreamsService.GetKeyStream:input_type -> streammachine.api.key_streams.v1.GetKeyStreamRequest
	1, // 5: streammachine.api.key_streams.v1.KeyStreamsService.ListKeyStreams:output_type -> streammachine.api.key_streams.v1.ListKeyStreamsResponse
	3, // 6: streammachine.api.key_streams.v1.KeyStreamsService.GetKeyStream:output_type -> streammachine.api.key_streams.v1.GetKeyStreamResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_streammachine_api_key_streams_v1_key_streams_proto_init() }
func file_streammachine_api_key_streams_v1_key_streams_proto_init() {
	if File_streammachine_api_key_streams_v1_key_streams_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKeyStreamsRequest); i {
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
		file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKeyStreamsResponse); i {
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
		file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyStreamRequest); i {
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
		file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyStreamResponse); i {
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
			RawDescriptor: file_streammachine_api_key_streams_v1_key_streams_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_streammachine_api_key_streams_v1_key_streams_proto_goTypes,
		DependencyIndexes: file_streammachine_api_key_streams_v1_key_streams_proto_depIdxs,
		MessageInfos:      file_streammachine_api_key_streams_v1_key_streams_proto_msgTypes,
	}.Build()
	File_streammachine_api_key_streams_v1_key_streams_proto = out.File
	file_streammachine_api_key_streams_v1_key_streams_proto_rawDesc = nil
	file_streammachine_api_key_streams_v1_key_streams_proto_goTypes = nil
	file_streammachine_api_key_streams_v1_key_streams_proto_depIdxs = nil
}
