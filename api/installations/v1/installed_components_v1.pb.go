// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: strmprivacy/api/installations/v1/installed_components_v1.proto

package installations

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// (-- api-linter: core::0134::request-mask-required=disabled
//     aip.dev/not-precedent: This RPC only updates one field. --)
type UpdateInstalledComponentStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComponentType InstalledComponentType `protobuf:"varint,1,opt,name=component_type,json=componentType,proto3,enum=strmprivacy.api.installations.v1.InstalledComponentType" json:"component_type,omitempty"`
	// A string that represents the instance. The hostname is used for this, which in Kubernetes defaults to the pod name.
	InstanceId              string                   `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	InstalledComponentState *InstalledComponentState `protobuf:"bytes,3,opt,name=installed_component_state,json=installedComponentState,proto3" json:"installed_component_state,omitempty"`
	InstallationId          string                   `protobuf:"bytes,4,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
}

func (x *UpdateInstalledComponentStateRequest) Reset() {
	*x = UpdateInstalledComponentStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInstalledComponentStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInstalledComponentStateRequest) ProtoMessage() {}

func (x *UpdateInstalledComponentStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInstalledComponentStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateInstalledComponentStateRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateInstalledComponentStateRequest) GetComponentType() InstalledComponentType {
	if x != nil {
		return x.ComponentType
	}
	return InstalledComponentType_INSTALLED_COMPONENT_TYPE_UNSPECIFIED
}

func (x *UpdateInstalledComponentStateRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *UpdateInstalledComponentStateRequest) GetInstalledComponentState() *InstalledComponentState {
	if x != nil {
		return x.InstalledComponentState
	}
	return nil
}

func (x *UpdateInstalledComponentStateRequest) GetInstallationId() string {
	if x != nil {
		return x.InstallationId
	}
	return ""
}

type UpdateInstalledComponentStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateInstalledComponentStateResponse) Reset() {
	*x = UpdateInstalledComponentStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInstalledComponentStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInstalledComponentStateResponse) ProtoMessage() {}

func (x *UpdateInstalledComponentStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInstalledComponentStateResponse.ProtoReflect.Descriptor instead.
func (*UpdateInstalledComponentStateResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescGZIP(), []int{1}
}

type GetInstalledComponentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationId         string                 `protobuf:"bytes,1,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	InstalledComponentType InstalledComponentType `protobuf:"varint,2,opt,name=installed_component_type,json=installedComponentType,proto3,enum=strmprivacy.api.installations.v1.InstalledComponentType" json:"installed_component_type,omitempty"`
}

func (x *GetInstalledComponentRequest) Reset() {
	*x = GetInstalledComponentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInstalledComponentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstalledComponentRequest) ProtoMessage() {}

func (x *GetInstalledComponentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstalledComponentRequest.ProtoReflect.Descriptor instead.
func (*GetInstalledComponentRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescGZIP(), []int{2}
}

func (x *GetInstalledComponentRequest) GetInstallationId() string {
	if x != nil {
		return x.InstallationId
	}
	return ""
}

func (x *GetInstalledComponentRequest) GetInstalledComponentType() InstalledComponentType {
	if x != nil {
		return x.InstalledComponentType
	}
	return InstalledComponentType_INSTALLED_COMPONENT_TYPE_UNSPECIFIED
}

type GetInstalledComponentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstalledComponent *InstalledComponent `protobuf:"bytes,1,opt,name=installed_component,json=installedComponent,proto3" json:"installed_component,omitempty"`
}

func (x *GetInstalledComponentResponse) Reset() {
	*x = GetInstalledComponentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInstalledComponentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstalledComponentResponse) ProtoMessage() {}

func (x *GetInstalledComponentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstalledComponentResponse.ProtoReflect.Descriptor instead.
func (*GetInstalledComponentResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescGZIP(), []int{3}
}

func (x *GetInstalledComponentResponse) GetInstalledComponent() *InstalledComponent {
	if x != nil {
		return x.InstalledComponent
	}
	return nil
}

type ListInstalledComponentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationId string `protobuf:"bytes,1,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
}

func (x *ListInstalledComponentsRequest) Reset() {
	*x = ListInstalledComponentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstalledComponentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstalledComponentsRequest) ProtoMessage() {}

func (x *ListInstalledComponentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstalledComponentsRequest.ProtoReflect.Descriptor instead.
func (*ListInstalledComponentsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescGZIP(), []int{4}
}

func (x *ListInstalledComponentsRequest) GetInstallationId() string {
	if x != nil {
		return x.InstallationId
	}
	return ""
}

type ListInstalledComponentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstalledComponents []*InstalledComponent `protobuf:"bytes,1,rep,name=installed_components,json=installedComponents,proto3" json:"installed_components,omitempty"`
}

func (x *ListInstalledComponentsResponse) Reset() {
	*x = ListInstalledComponentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstalledComponentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstalledComponentsResponse) ProtoMessage() {}

func (x *ListInstalledComponentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstalledComponentsResponse.ProtoReflect.Descriptor instead.
func (*ListInstalledComponentsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescGZIP(), []int{5}
}

func (x *ListInstalledComponentsResponse) GetInstalledComponents() []*InstalledComponent {
	if x != nil {
		return x.InstalledComponents
	}
	return nil
}

type ListInstalledComponentsCurrentStatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationId string `protobuf:"bytes,1,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
}

func (x *ListInstalledComponentsCurrentStatesRequest) Reset() {
	*x = ListInstalledComponentsCurrentStatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstalledComponentsCurrentStatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstalledComponentsCurrentStatesRequest) ProtoMessage() {}

func (x *ListInstalledComponentsCurrentStatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstalledComponentsCurrentStatesRequest.ProtoReflect.Descriptor instead.
func (*ListInstalledComponentsCurrentStatesRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescGZIP(), []int{6}
}

func (x *ListInstalledComponentsCurrentStatesRequest) GetInstallationId() string {
	if x != nil {
		return x.InstallationId
	}
	return ""
}

type ListInstalledComponentsCurrentStatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (-- api-linter: core::0132::response-unknown-fields=disabled
	//     aip.dev/not-precedent: field name gets too clunky --)
	InstalledComponents []*InstalledComponentsCurrentState `protobuf:"bytes,1,rep,name=installed_components,json=installedComponents,proto3" json:"installed_components,omitempty"`
}

func (x *ListInstalledComponentsCurrentStatesResponse) Reset() {
	*x = ListInstalledComponentsCurrentStatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstalledComponentsCurrentStatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstalledComponentsCurrentStatesResponse) ProtoMessage() {}

func (x *ListInstalledComponentsCurrentStatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstalledComponentsCurrentStatesResponse.ProtoReflect.Descriptor instead.
func (*ListInstalledComponentsCurrentStatesResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescGZIP(), []int{7}
}

func (x *ListInstalledComponentsCurrentStatesResponse) GetInstalledComponents() []*InstalledComponentsCurrentState {
	if x != nil {
		return x.InstalledComponents
	}
	return nil
}

var File_strmprivacy_api_installations_v1_installed_components_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x32, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xef, 0x02, 0x0a, 0x24, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x38, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82,
	0x01, 0x02, 0x20, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x04, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x7f, 0x0a,
	0x19, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x17, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x31,
	0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x27, 0x0a, 0x25, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x7c,
	0x0a, 0x18, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x38, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82,
	0x01, 0x02, 0x20, 0x00, 0x52, 0x16, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x86, 0x01, 0x0a,
	0x1d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65,
	0x0a, 0x13, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x1f, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67,
	0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x52, 0x13, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x60, 0x0a, 0x2b, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x2c, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x14, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x13, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x32, 0xd3, 0x05, 0x0a, 0x1a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xb0, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x46, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9e, 0x01,
	0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x40, 0x2e, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xc5,
	0x01, 0x0a, 0x24, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x4d, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4e, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x78, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescData = file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDesc
)

func file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDescData
}

var file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_strmprivacy_api_installations_v1_installed_components_v1_proto_goTypes = []interface{}{
	(*UpdateInstalledComponentStateRequest)(nil),         // 0: strmprivacy.api.installations.v1.UpdateInstalledComponentStateRequest
	(*UpdateInstalledComponentStateResponse)(nil),        // 1: strmprivacy.api.installations.v1.UpdateInstalledComponentStateResponse
	(*GetInstalledComponentRequest)(nil),                 // 2: strmprivacy.api.installations.v1.GetInstalledComponentRequest
	(*GetInstalledComponentResponse)(nil),                // 3: strmprivacy.api.installations.v1.GetInstalledComponentResponse
	(*ListInstalledComponentsRequest)(nil),               // 4: strmprivacy.api.installations.v1.ListInstalledComponentsRequest
	(*ListInstalledComponentsResponse)(nil),              // 5: strmprivacy.api.installations.v1.ListInstalledComponentsResponse
	(*ListInstalledComponentsCurrentStatesRequest)(nil),  // 6: strmprivacy.api.installations.v1.ListInstalledComponentsCurrentStatesRequest
	(*ListInstalledComponentsCurrentStatesResponse)(nil), // 7: strmprivacy.api.installations.v1.ListInstalledComponentsCurrentStatesResponse
	(InstalledComponentType)(0),                          // 8: strmprivacy.api.installations.v1.InstalledComponentType
	(*InstalledComponentState)(nil),                      // 9: strmprivacy.api.installations.v1.InstalledComponentState
	(*InstalledComponent)(nil),                           // 10: strmprivacy.api.installations.v1.InstalledComponent
	(*InstalledComponentsCurrentState)(nil),              // 11: strmprivacy.api.installations.v1.InstalledComponentsCurrentState
}
var file_strmprivacy_api_installations_v1_installed_components_v1_proto_depIdxs = []int32{
	8,  // 0: strmprivacy.api.installations.v1.UpdateInstalledComponentStateRequest.component_type:type_name -> strmprivacy.api.installations.v1.InstalledComponentType
	9,  // 1: strmprivacy.api.installations.v1.UpdateInstalledComponentStateRequest.installed_component_state:type_name -> strmprivacy.api.installations.v1.InstalledComponentState
	8,  // 2: strmprivacy.api.installations.v1.GetInstalledComponentRequest.installed_component_type:type_name -> strmprivacy.api.installations.v1.InstalledComponentType
	10, // 3: strmprivacy.api.installations.v1.GetInstalledComponentResponse.installed_component:type_name -> strmprivacy.api.installations.v1.InstalledComponent
	10, // 4: strmprivacy.api.installations.v1.ListInstalledComponentsResponse.installed_components:type_name -> strmprivacy.api.installations.v1.InstalledComponent
	11, // 5: strmprivacy.api.installations.v1.ListInstalledComponentsCurrentStatesResponse.installed_components:type_name -> strmprivacy.api.installations.v1.InstalledComponentsCurrentState
	0,  // 6: strmprivacy.api.installations.v1.InstalledComponentsService.UpdateInstalledComponentState:input_type -> strmprivacy.api.installations.v1.UpdateInstalledComponentStateRequest
	2,  // 7: strmprivacy.api.installations.v1.InstalledComponentsService.GetInstalledComponent:input_type -> strmprivacy.api.installations.v1.GetInstalledComponentRequest
	4,  // 8: strmprivacy.api.installations.v1.InstalledComponentsService.ListInstalledComponents:input_type -> strmprivacy.api.installations.v1.ListInstalledComponentsRequest
	6,  // 9: strmprivacy.api.installations.v1.InstalledComponentsService.ListInstalledComponentsCurrentStates:input_type -> strmprivacy.api.installations.v1.ListInstalledComponentsCurrentStatesRequest
	1,  // 10: strmprivacy.api.installations.v1.InstalledComponentsService.UpdateInstalledComponentState:output_type -> strmprivacy.api.installations.v1.UpdateInstalledComponentStateResponse
	3,  // 11: strmprivacy.api.installations.v1.InstalledComponentsService.GetInstalledComponent:output_type -> strmprivacy.api.installations.v1.GetInstalledComponentResponse
	5,  // 12: strmprivacy.api.installations.v1.InstalledComponentsService.ListInstalledComponents:output_type -> strmprivacy.api.installations.v1.ListInstalledComponentsResponse
	7,  // 13: strmprivacy.api.installations.v1.InstalledComponentsService.ListInstalledComponentsCurrentStates:output_type -> strmprivacy.api.installations.v1.ListInstalledComponentsCurrentStatesResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_installations_v1_installed_components_v1_proto_init() }
func file_strmprivacy_api_installations_v1_installed_components_v1_proto_init() {
	if File_strmprivacy_api_installations_v1_installed_components_v1_proto != nil {
		return
	}
	file_strmprivacy_api_installations_v1_entities_v1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInstalledComponentStateRequest); i {
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
		file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInstalledComponentStateResponse); i {
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
		file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInstalledComponentRequest); i {
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
		file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInstalledComponentResponse); i {
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
		file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstalledComponentsRequest); i {
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
		file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstalledComponentsResponse); i {
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
		file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstalledComponentsCurrentStatesRequest); i {
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
		file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstalledComponentsCurrentStatesResponse); i {
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
			RawDescriptor: file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_installations_v1_installed_components_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_installations_v1_installed_components_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_installations_v1_installed_components_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_installations_v1_installed_components_v1_proto = out.File
	file_strmprivacy_api_installations_v1_installed_components_v1_proto_rawDesc = nil
	file_strmprivacy_api_installations_v1_installed_components_v1_proto_goTypes = nil
	file_strmprivacy_api_installations_v1_installed_components_v1_proto_depIdxs = nil
}
