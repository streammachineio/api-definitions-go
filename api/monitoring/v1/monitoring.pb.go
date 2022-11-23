// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: strmprivacy/api/monitoring/v1/monitoring.proto

package monitoring

import (
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

type GetEntityStatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEntityStatesRequest) Reset() {
	*x = GetEntityStatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntityStatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntityStatesRequest) ProtoMessage() {}

func (x *GetEntityStatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntityStatesRequest.ProtoReflect.Descriptor instead.
func (*GetEntityStatesRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDescGZIP(), []int{0}
}

//
//Initial call returns every entity state that's relevant to the caller.
//
//Subsequent streaming gives deltas
type GetEntityStatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States []*GetEntityStatesResponse_State `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *GetEntityStatesResponse) Reset() {
	*x = GetEntityStatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntityStatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntityStatesResponse) ProtoMessage() {}

func (x *GetEntityStatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntityStatesResponse.ProtoReflect.Descriptor instead.
func (*GetEntityStatesResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDescGZIP(), []int{1}
}

func (x *GetEntityStatesResponse) GetStates() []*GetEntityStatesResponse_State {
	if x != nil {
		return x.States
	}
	return nil
}

type GetEntityStatesResponse_State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sos, schema-registry, batch-job X, ...
	EntityType string                       `protobuf:"bytes,1,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	EntityRef  *GetEntityStatesResponse_Ref `protobuf:"bytes,2,opt,name=entity_ref,json=entityRef,proto3" json:"entity_ref,omitempty"`
	// healthy, broken, struggling, REMOVED
	State   string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// whether it's a STRM entity or a customer entity
	CustomerEntity bool `protobuf:"varint,5,opt,name=customer_entity,json=customerEntity,proto3" json:"customer_entity,omitempty"`
}

func (x *GetEntityStatesResponse_State) Reset() {
	*x = GetEntityStatesResponse_State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntityStatesResponse_State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntityStatesResponse_State) ProtoMessage() {}

func (x *GetEntityStatesResponse_State) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntityStatesResponse_State.ProtoReflect.Descriptor instead.
func (*GetEntityStatesResponse_State) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetEntityStatesResponse_State) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *GetEntityStatesResponse_State) GetEntityRef() *GetEntityStatesResponse_Ref {
	if x != nil {
		return x.EntityRef
	}
	return nil
}

func (x *GetEntityStatesResponse_State) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GetEntityStatesResponse_State) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetEntityStatesResponse_State) GetCustomerEntity() bool {
	if x != nil {
		return x.CustomerEntity
	}
	return false
}

type GetEntityStatesResponse_Ref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id        string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEntityStatesResponse_Ref) Reset() {
	*x = GetEntityStatesResponse_Ref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntityStatesResponse_Ref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntityStatesResponse_Ref) ProtoMessage() {}

func (x *GetEntityStatesResponse_Ref) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntityStatesResponse_Ref.ProtoReflect.Descriptor instead.
func (*GetEntityStatesResponse_Ref) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetEntityStatesResponse_Ref) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GetEntityStatesResponse_Ref) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetEntityStatesResponse_Ref) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_strmprivacy_api_monitoring_v1_monitoring_proto protoreflect.FileDescriptor

var file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22,
	0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x98, 0x03, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x1a, 0xdc, 0x01, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x66, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x48, 0x0a, 0x03, 0x52, 0x65,
	0x66, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x98, 0x01, 0x0a, 0x11, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x35,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42,
	0x6f, 0x0a, 0x20, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f,
	0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDescOnce sync.Once
	file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDescData = file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDesc
)

func file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDescData)
	})
	return file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDescData
}

var file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_strmprivacy_api_monitoring_v1_monitoring_proto_goTypes = []interface{}{
	(*GetEntityStatesRequest)(nil),        // 0: strmprivacy.api.monitoring.v1.GetEntityStatesRequest
	(*GetEntityStatesResponse)(nil),       // 1: strmprivacy.api.monitoring.v1.GetEntityStatesResponse
	(*GetEntityStatesResponse_State)(nil), // 2: strmprivacy.api.monitoring.v1.GetEntityStatesResponse.State
	(*GetEntityStatesResponse_Ref)(nil),   // 3: strmprivacy.api.monitoring.v1.GetEntityStatesResponse.Ref
}
var file_strmprivacy_api_monitoring_v1_monitoring_proto_depIdxs = []int32{
	2, // 0: strmprivacy.api.monitoring.v1.GetEntityStatesResponse.states:type_name -> strmprivacy.api.monitoring.v1.GetEntityStatesResponse.State
	3, // 1: strmprivacy.api.monitoring.v1.GetEntityStatesResponse.State.entity_ref:type_name -> strmprivacy.api.monitoring.v1.GetEntityStatesResponse.Ref
	0, // 2: strmprivacy.api.monitoring.v1.MonitoringService.GetEntityStates:input_type -> strmprivacy.api.monitoring.v1.GetEntityStatesRequest
	1, // 3: strmprivacy.api.monitoring.v1.MonitoringService.GetEntityStates:output_type -> strmprivacy.api.monitoring.v1.GetEntityStatesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_monitoring_v1_monitoring_proto_init() }
func file_strmprivacy_api_monitoring_v1_monitoring_proto_init() {
	if File_strmprivacy_api_monitoring_v1_monitoring_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntityStatesRequest); i {
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
		file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntityStatesResponse); i {
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
		file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntityStatesResponse_State); i {
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
		file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntityStatesResponse_Ref); i {
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
			RawDescriptor: file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_monitoring_v1_monitoring_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_monitoring_v1_monitoring_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_monitoring_v1_monitoring_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_monitoring_v1_monitoring_proto = out.File
	file_strmprivacy_api_monitoring_v1_monitoring_proto_rawDesc = nil
	file_strmprivacy_api_monitoring_v1_monitoring_proto_goTypes = nil
	file_strmprivacy_api_monitoring_v1_monitoring_proto_depIdxs = nil
}
