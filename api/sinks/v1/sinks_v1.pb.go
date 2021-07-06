// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.2
// source: streammachine/api/sinks/v1/sinks_v1.proto

package sinks

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

type ListSinksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
	Recursive bool   `protobuf:"varint,2,opt,name=recursive,proto3" json:"recursive,omitempty"`
}

func (x *ListSinksRequest) Reset() {
	*x = ListSinksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSinksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSinksRequest) ProtoMessage() {}

func (x *ListSinksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSinksRequest.ProtoReflect.Descriptor instead.
func (*ListSinksRequest) Descriptor() ([]byte, []int) {
	return file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescGZIP(), []int{0}
}

func (x *ListSinksRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

func (x *ListSinksRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

type ListSinksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sinks []*v1.SinkTree `protobuf:"bytes,1,rep,name=sinks,proto3" json:"sinks,omitempty"`
}

func (x *ListSinksResponse) Reset() {
	*x = ListSinksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSinksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSinksResponse) ProtoMessage() {}

func (x *ListSinksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSinksResponse.ProtoReflect.Descriptor instead.
func (*ListSinksResponse) Descriptor() ([]byte, []int) {
	return file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ListSinksResponse) GetSinks() []*v1.SinkTree {
	if x != nil {
		return x.Sinks
	}
	return nil
}

type DeleteSinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *v1.SinkRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// (-- api-linter: core::0135::request-unknown-fields=disabled
	//     aip.dev/not-precedent: We really need this field. --)
	Recursive bool `protobuf:"varint,2,opt,name=recursive,proto3" json:"recursive,omitempty"`
}

func (x *DeleteSinkRequest) Reset() {
	*x = DeleteSinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSinkRequest) ProtoMessage() {}

func (x *DeleteSinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSinkRequest.ProtoReflect.Descriptor instead.
func (*DeleteSinkRequest) Descriptor() ([]byte, []int) {
	return file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteSinkRequest) GetRef() *v1.SinkRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *DeleteSinkRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

type DeleteSinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSinkResponse) Reset() {
	*x = DeleteSinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSinkResponse) ProtoMessage() {}

func (x *DeleteSinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSinkResponse.ProtoReflect.Descriptor instead.
func (*DeleteSinkResponse) Descriptor() ([]byte, []int) {
	return file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescGZIP(), []int{3}
}

type CreateSinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sink *v1.Sink `protobuf:"bytes,1,opt,name=sink,proto3" json:"sink,omitempty"`
}

func (x *CreateSinkRequest) Reset() {
	*x = CreateSinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSinkRequest) ProtoMessage() {}

func (x *CreateSinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSinkRequest.ProtoReflect.Descriptor instead.
func (*CreateSinkRequest) Descriptor() ([]byte, []int) {
	return file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSinkRequest) GetSink() *v1.Sink {
	if x != nil {
		return x.Sink
	}
	return nil
}

type CreateSinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sink *v1.Sink `protobuf:"bytes,1,opt,name=sink,proto3" json:"sink,omitempty"`
}

func (x *CreateSinkResponse) Reset() {
	*x = CreateSinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSinkResponse) ProtoMessage() {}

func (x *CreateSinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSinkResponse.ProtoReflect.Descriptor instead.
func (*CreateSinkResponse) Descriptor() ([]byte, []int) {
	return file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSinkResponse) GetSink() *v1.Sink {
	if x != nil {
		return x.Sink
	}
	return nil
}

type GetSinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref       *v1.SinkRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Recursive bool        `protobuf:"varint,2,opt,name=recursive,proto3" json:"recursive,omitempty"`
}

func (x *GetSinkRequest) Reset() {
	*x = GetSinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSinkRequest) ProtoMessage() {}

func (x *GetSinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSinkRequest.ProtoReflect.Descriptor instead.
func (*GetSinkRequest) Descriptor() ([]byte, []int) {
	return file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescGZIP(), []int{6}
}

func (x *GetSinkRequest) GetRef() *v1.SinkRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *GetSinkRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

type GetSinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SinkTree *v1.SinkTree `protobuf:"bytes,1,opt,name=sink_tree,json=sinkTree,proto3" json:"sink_tree,omitempty"`
}

func (x *GetSinkResponse) Reset() {
	*x = GetSinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSinkResponse) ProtoMessage() {}

func (x *GetSinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSinkResponse.ProtoReflect.Descriptor instead.
func (*GetSinkResponse) Descriptor() ([]byte, []int) {
	return file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescGZIP(), []int{7}
}

func (x *GetSinkResponse) GetSinkTree() *v1.SinkTree {
	if x != nil {
		return x.SinkTree
	}
	return nil
}

var File_streammachine_api_sinks_v1_sinks_v1_proto protoreflect.FileDescriptor

var file_streammachine_api_sinks_v1_sinks_v1_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x6e,
	0x6b, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x22,
	0x52, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x54, 0x72, 0x65, 0x65, 0x52, 0x05, 0x73, 0x69,
	0x6e, 0x6b, 0x73, 0x22, 0x75, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x21, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72,
	0x73, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x73, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x73, 0x69, 0x6e, 0x6b, 0x22, 0x4d,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x73, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x73, 0x69, 0x6e, 0x6b, 0x22, 0x72, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3d, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x21,
	0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76,
	0x65, 0x22, 0x57, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x73, 0x69, 0x6e, 0x6b, 0x5f, 0x74, 0x72, 0x65,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x54, 0x72, 0x65, 0x65,
	0x52, 0x08, 0x73, 0x69, 0x6e, 0x6b, 0x54, 0x72, 0x65, 0x65, 0x32, 0xb6, 0x03, 0x0a, 0x0c, 0x53,
	0x69, 0x6e, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x2c, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6e,
	0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x6e, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x6b,
	0x12, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6e, 0x6b,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x6e, 0x6b, 0x12, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x63, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x69, 0x6e, 0x6b,
	0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescOnce sync.Once
	file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescData = file_streammachine_api_sinks_v1_sinks_v1_proto_rawDesc
)

func file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescGZIP() []byte {
	file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescOnce.Do(func() {
		file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescData)
	})
	return file_streammachine_api_sinks_v1_sinks_v1_proto_rawDescData
}

var file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_streammachine_api_sinks_v1_sinks_v1_proto_goTypes = []interface{}{
	(*ListSinksRequest)(nil),   // 0: streammachine.api.sinks.v1.ListSinksRequest
	(*ListSinksResponse)(nil),  // 1: streammachine.api.sinks.v1.ListSinksResponse
	(*DeleteSinkRequest)(nil),  // 2: streammachine.api.sinks.v1.DeleteSinkRequest
	(*DeleteSinkResponse)(nil), // 3: streammachine.api.sinks.v1.DeleteSinkResponse
	(*CreateSinkRequest)(nil),  // 4: streammachine.api.sinks.v1.CreateSinkRequest
	(*CreateSinkResponse)(nil), // 5: streammachine.api.sinks.v1.CreateSinkResponse
	(*GetSinkRequest)(nil),     // 6: streammachine.api.sinks.v1.GetSinkRequest
	(*GetSinkResponse)(nil),    // 7: streammachine.api.sinks.v1.GetSinkResponse
	(*v1.SinkTree)(nil),        // 8: streammachine.api.entities.v1.SinkTree
	(*v1.SinkRef)(nil),         // 9: streammachine.api.entities.v1.SinkRef
	(*v1.Sink)(nil),            // 10: streammachine.api.entities.v1.Sink
}
var file_streammachine_api_sinks_v1_sinks_v1_proto_depIdxs = []int32{
	8,  // 0: streammachine.api.sinks.v1.ListSinksResponse.sinks:type_name -> streammachine.api.entities.v1.SinkTree
	9,  // 1: streammachine.api.sinks.v1.DeleteSinkRequest.ref:type_name -> streammachine.api.entities.v1.SinkRef
	10, // 2: streammachine.api.sinks.v1.CreateSinkRequest.sink:type_name -> streammachine.api.entities.v1.Sink
	10, // 3: streammachine.api.sinks.v1.CreateSinkResponse.sink:type_name -> streammachine.api.entities.v1.Sink
	9,  // 4: streammachine.api.sinks.v1.GetSinkRequest.ref:type_name -> streammachine.api.entities.v1.SinkRef
	8,  // 5: streammachine.api.sinks.v1.GetSinkResponse.sink_tree:type_name -> streammachine.api.entities.v1.SinkTree
	0,  // 6: streammachine.api.sinks.v1.SinksService.ListSinks:input_type -> streammachine.api.sinks.v1.ListSinksRequest
	6,  // 7: streammachine.api.sinks.v1.SinksService.GetSink:input_type -> streammachine.api.sinks.v1.GetSinkRequest
	2,  // 8: streammachine.api.sinks.v1.SinksService.DeleteSink:input_type -> streammachine.api.sinks.v1.DeleteSinkRequest
	4,  // 9: streammachine.api.sinks.v1.SinksService.CreateSink:input_type -> streammachine.api.sinks.v1.CreateSinkRequest
	1,  // 10: streammachine.api.sinks.v1.SinksService.ListSinks:output_type -> streammachine.api.sinks.v1.ListSinksResponse
	7,  // 11: streammachine.api.sinks.v1.SinksService.GetSink:output_type -> streammachine.api.sinks.v1.GetSinkResponse
	3,  // 12: streammachine.api.sinks.v1.SinksService.DeleteSink:output_type -> streammachine.api.sinks.v1.DeleteSinkResponse
	5,  // 13: streammachine.api.sinks.v1.SinksService.CreateSink:output_type -> streammachine.api.sinks.v1.CreateSinkResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_streammachine_api_sinks_v1_sinks_v1_proto_init() }
func file_streammachine_api_sinks_v1_sinks_v1_proto_init() {
	if File_streammachine_api_sinks_v1_sinks_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSinksRequest); i {
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
		file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSinksResponse); i {
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
		file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSinkRequest); i {
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
		file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSinkResponse); i {
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
		file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSinkRequest); i {
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
		file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSinkResponse); i {
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
		file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSinkRequest); i {
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
		file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSinkResponse); i {
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
			RawDescriptor: file_streammachine_api_sinks_v1_sinks_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_streammachine_api_sinks_v1_sinks_v1_proto_goTypes,
		DependencyIndexes: file_streammachine_api_sinks_v1_sinks_v1_proto_depIdxs,
		MessageInfos:      file_streammachine_api_sinks_v1_sinks_v1_proto_msgTypes,
	}.Build()
	File_streammachine_api_sinks_v1_sinks_v1_proto = out.File
	file_streammachine_api_sinks_v1_sinks_v1_proto_rawDesc = nil
	file_streammachine_api_sinks_v1_sinks_v1_proto_goTypes = nil
	file_streammachine_api_sinks_v1_sinks_v1_proto_depIdxs = nil
}
