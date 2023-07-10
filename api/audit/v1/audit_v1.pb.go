// (-- api-linter: core::0133::request-id-field=disabled
//     aip.dev/not-precedent: We need to do this because we typically use entities that contain
//     references inside them, and not an entity id directly in the request. --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: strmprivacy/api/audit/v1/audit_v1.proto

package audit

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

type GetAuditTrailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// entity_ref is the entity to get the audit trail for.
	// If empty, a project id must be provided instead.
	EntityRef *v1.GenericRef `protobuf:"bytes,1,opt,name=entity_ref,json=entityRef,proto3" json:"entity_ref,omitempty"`
	PageSize  int32          `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string         `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// If a project id is specified, an audit trail for all entities in the project is returned.
	// If empty, an entity_ref must be provided instead.
	ProjectId string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetAuditTrailRequest) Reset() {
	*x = GetAuditTrailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuditTrailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuditTrailRequest) ProtoMessage() {}

func (x *GetAuditTrailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuditTrailRequest.ProtoReflect.Descriptor instead.
func (*GetAuditTrailRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_audit_v1_audit_v1_proto_rawDescGZIP(), []int{0}
}

func (x *GetAuditTrailRequest) GetEntityRef() *v1.GenericRef {
	if x != nil {
		return x.EntityRef
	}
	return nil
}

func (x *GetAuditTrailRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetAuditTrailRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *GetAuditTrailRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetAuditTrailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries       []*v1.AuditTrailEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	NextPageToken string                `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *GetAuditTrailResponse) Reset() {
	*x = GetAuditTrailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuditTrailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuditTrailResponse) ProtoMessage() {}

func (x *GetAuditTrailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuditTrailResponse.ProtoReflect.Descriptor instead.
func (*GetAuditTrailResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_audit_v1_audit_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GetAuditTrailResponse) GetEntries() []*v1.AuditTrailEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *GetAuditTrailResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateAuditTrailEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuditTrailEntry *v1.AuditTrailEntry `protobuf:"bytes,1,opt,name=audit_trail_entry,json=auditTrailEntry,proto3" json:"audit_trail_entry,omitempty"`
}

func (x *CreateAuditTrailEntryRequest) Reset() {
	*x = CreateAuditTrailEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuditTrailEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuditTrailEntryRequest) ProtoMessage() {}

func (x *CreateAuditTrailEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuditTrailEntryRequest.ProtoReflect.Descriptor instead.
func (*CreateAuditTrailEntryRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_audit_v1_audit_v1_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAuditTrailEntryRequest) GetAuditTrailEntry() *v1.AuditTrailEntry {
	if x != nil {
		return x.AuditTrailEntry
	}
	return nil
}

type CreateAuditTrailEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuditTrailEntry *v1.AuditTrailEntry `protobuf:"bytes,1,opt,name=audit_trail_entry,json=auditTrailEntry,proto3" json:"audit_trail_entry,omitempty"`
}

func (x *CreateAuditTrailEntryResponse) Reset() {
	*x = CreateAuditTrailEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuditTrailEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuditTrailEntryResponse) ProtoMessage() {}

func (x *CreateAuditTrailEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuditTrailEntryResponse.ProtoReflect.Descriptor instead.
func (*CreateAuditTrailEntryResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_audit_v1_audit_v1_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAuditTrailEntryResponse) GetAuditTrailEntry() *v1.AuditTrailEntry {
	if x != nil {
		return x.AuditTrailEntry
	}
	return nil
}

var File_strmprivacy_api_audit_v1_audit_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_audit_v1_audit_v1_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x2d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x72,
	0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65,
	0x66, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x66, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42,
	0x08, 0x72, 0x06, 0xd0, 0x01, 0x01, 0xb0, 0x01, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x82,
	0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x72,
	0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x62, 0x0a, 0x11, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x72,
	0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x22, 0x79, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x74, 0x72,
	0x61, 0x69, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x32, 0x8b,
	0x02, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x70, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c,
	0x12, 0x2e, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x88, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x36, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x60, 0x0a, 0x1b,
	0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_audit_v1_audit_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_audit_v1_audit_v1_proto_rawDescData = file_strmprivacy_api_audit_v1_audit_v1_proto_rawDesc
)

func file_strmprivacy_api_audit_v1_audit_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_audit_v1_audit_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_audit_v1_audit_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_audit_v1_audit_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_audit_v1_audit_v1_proto_rawDescData
}

var file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_strmprivacy_api_audit_v1_audit_v1_proto_goTypes = []interface{}{
	(*GetAuditTrailRequest)(nil),          // 0: strmprivacy.api.audit.v1.GetAuditTrailRequest
	(*GetAuditTrailResponse)(nil),         // 1: strmprivacy.api.audit.v1.GetAuditTrailResponse
	(*CreateAuditTrailEntryRequest)(nil),  // 2: strmprivacy.api.audit.v1.CreateAuditTrailEntryRequest
	(*CreateAuditTrailEntryResponse)(nil), // 3: strmprivacy.api.audit.v1.CreateAuditTrailEntryResponse
	(*v1.GenericRef)(nil),                 // 4: strmprivacy.api.entities.v1.GenericRef
	(*v1.AuditTrailEntry)(nil),            // 5: strmprivacy.api.entities.v1.AuditTrailEntry
}
var file_strmprivacy_api_audit_v1_audit_v1_proto_depIdxs = []int32{
	4, // 0: strmprivacy.api.audit.v1.GetAuditTrailRequest.entity_ref:type_name -> strmprivacy.api.entities.v1.GenericRef
	5, // 1: strmprivacy.api.audit.v1.GetAuditTrailResponse.entries:type_name -> strmprivacy.api.entities.v1.AuditTrailEntry
	5, // 2: strmprivacy.api.audit.v1.CreateAuditTrailEntryRequest.audit_trail_entry:type_name -> strmprivacy.api.entities.v1.AuditTrailEntry
	5, // 3: strmprivacy.api.audit.v1.CreateAuditTrailEntryResponse.audit_trail_entry:type_name -> strmprivacy.api.entities.v1.AuditTrailEntry
	0, // 4: strmprivacy.api.audit.v1.AuditService.GetAuditTrail:input_type -> strmprivacy.api.audit.v1.GetAuditTrailRequest
	2, // 5: strmprivacy.api.audit.v1.AuditService.CreateAuditTrailEntry:input_type -> strmprivacy.api.audit.v1.CreateAuditTrailEntryRequest
	1, // 6: strmprivacy.api.audit.v1.AuditService.GetAuditTrail:output_type -> strmprivacy.api.audit.v1.GetAuditTrailResponse
	3, // 7: strmprivacy.api.audit.v1.AuditService.CreateAuditTrailEntry:output_type -> strmprivacy.api.audit.v1.CreateAuditTrailEntryResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_audit_v1_audit_v1_proto_init() }
func file_strmprivacy_api_audit_v1_audit_v1_proto_init() {
	if File_strmprivacy_api_audit_v1_audit_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuditTrailRequest); i {
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
		file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuditTrailResponse); i {
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
		file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuditTrailEntryRequest); i {
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
		file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuditTrailEntryResponse); i {
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
			RawDescriptor: file_strmprivacy_api_audit_v1_audit_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_audit_v1_audit_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_audit_v1_audit_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_audit_v1_audit_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_audit_v1_audit_v1_proto = out.File
	file_strmprivacy_api_audit_v1_audit_v1_proto_rawDesc = nil
	file_strmprivacy_api_audit_v1_audit_v1_proto_goTypes = nil
	file_strmprivacy_api_audit_v1_audit_v1_proto_depIdxs = nil
}
