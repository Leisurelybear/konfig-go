// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.23.1
// source: konfig.proto

package pb

import (
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

type ListConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionID int64  `protobuf:"varint,1,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // fuzzy
	Sort         int64  `protobuf:"varint,3,opt,name=sort,proto3" json:"sort,omitempty"`
	PageNum      int64  `protobuf:"varint,4,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize     int64  `protobuf:"varint,5,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListConfigsRequest) Reset() {
	*x = ListConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConfigsRequest) ProtoMessage() {}

func (x *ListConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_konfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConfigsRequest.ProtoReflect.Descriptor instead.
func (*ListConfigsRequest) Descriptor() ([]byte, []int) {
	return file_konfig_proto_rawDescGZIP(), []int{0}
}

func (x *ListConfigsRequest) GetCollectionID() int64 {
	if x != nil {
		return x.CollectionID
	}
	return 0
}

func (x *ListConfigsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListConfigsRequest) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *ListConfigsRequest) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListConfigsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionID   int64     `protobuf:"varint,1,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	CollectionName string    `protobuf:"bytes,2,opt,name=collectionName,proto3" json:"collectionName,omitempty"` // fuzzy
	Sort           int64     `protobuf:"varint,3,opt,name=sort,proto3" json:"sort,omitempty"`
	PageNum        int64     `protobuf:"varint,4,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize       int64     `protobuf:"varint,5,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Configs        []*Config `protobuf:"bytes,6,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *ListConfigsResponse) Reset() {
	*x = ListConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConfigsResponse) ProtoMessage() {}

func (x *ListConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_konfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConfigsResponse.ProtoReflect.Descriptor instead.
func (*ListConfigsResponse) Descriptor() ([]byte, []int) {
	return file_konfig_proto_rawDescGZIP(), []int{1}
}

func (x *ListConfigsResponse) GetCollectionID() int64 {
	if x != nil {
		return x.CollectionID
	}
	return 0
}

func (x *ListConfigsResponse) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

func (x *ListConfigsResponse) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *ListConfigsResponse) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListConfigsResponse) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListConfigsResponse) GetConfigs() []*Config {
	if x != nil {
		return x.Configs
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CollectionID int64  `protobuf:"varint,2,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Key          string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Value        string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	CreatedBy    string `protobuf:"bytes,6,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	UpdatedBy    string `protobuf:"bytes,7,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	CreatedAt    int64  `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    int64  `protobuf:"varint,9,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_konfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_konfig_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Config) GetCollectionID() int64 {
	if x != nil {
		return x.CollectionID
	}
	return 0
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Config) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Config) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Config) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *Config) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Config) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type ListCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // fuzzy
	Sort      int64  `protobuf:"varint,2,opt,name=sort,proto3" json:"sort,omitempty"`
	PageNum   int64  `protobuf:"varint,3,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize  int64  `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	CreatedBy string `protobuf:"bytes,6,opt,name=createdBy,proto3" json:"createdBy,omitempty"` // created by username, fuzzy
}

func (x *ListCollectionRequest) Reset() {
	*x = ListCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konfig_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionRequest) ProtoMessage() {}

func (x *ListCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_konfig_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionRequest.ProtoReflect.Descriptor instead.
func (*ListCollectionRequest) Descriptor() ([]byte, []int) {
	return file_konfig_proto_rawDescGZIP(), []int{3}
}

func (x *ListCollectionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListCollectionRequest) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *ListCollectionRequest) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListCollectionRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCollectionRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type ListCollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDraft bool   `protobuf:"varint,1,opt,name=isDraft,proto3" json:"isDraft,omitempty"`
	Result  string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ListCollectionResponse) Reset() {
	*x = ListCollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konfig_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionResponse) ProtoMessage() {}

func (x *ListCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_konfig_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionResponse.ProtoReflect.Descriptor instead.
func (*ListCollectionResponse) Descriptor() ([]byte, []int) {
	return file_konfig_proto_rawDescGZIP(), []int{4}
}

func (x *ListCollectionResponse) GetIsDraft() bool {
	if x != nil {
		return x.IsDraft
	}
	return false
}

func (x *ListCollectionResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_konfig_proto protoreflect.FileDescriptor

var file_konfig_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6b, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x96, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xd1, 0x01, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0xf0, 0x01,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x93, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x4a, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x44, 0x72, 0x61, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x73, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0xc3, 0x01, 0x0a, 0x06, 0x4b, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x57, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x60, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_konfig_proto_rawDescOnce sync.Once
	file_konfig_proto_rawDescData = file_konfig_proto_rawDesc
)

func file_konfig_proto_rawDescGZIP() []byte {
	file_konfig_proto_rawDescOnce.Do(func() {
		file_konfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_konfig_proto_rawDescData)
	})
	return file_konfig_proto_rawDescData
}

var file_konfig_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_konfig_proto_goTypes = []interface{}{
	(*ListConfigsRequest)(nil),     // 0: pb.ListConfigsRequest
	(*ListConfigsResponse)(nil),    // 1: pb.ListConfigsResponse
	(*Config)(nil),                 // 2: pb.Config
	(*ListCollectionRequest)(nil),  // 3: pb.ListCollectionRequest
	(*ListCollectionResponse)(nil), // 4: pb.ListCollectionResponse
}
var file_konfig_proto_depIdxs = []int32{
	2, // 0: pb.ListConfigsResponse.configs:type_name -> pb.Config
	0, // 1: pb.Konfig.ListConfigs:input_type -> pb.ListConfigsRequest
	3, // 2: pb.Konfig.ListCollection:input_type -> pb.ListCollectionRequest
	1, // 3: pb.Konfig.ListConfigs:output_type -> pb.ListConfigsResponse
	4, // 4: pb.Konfig.ListCollection:output_type -> pb.ListCollectionResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_konfig_proto_init() }
func file_konfig_proto_init() {
	if File_konfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_konfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConfigsRequest); i {
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
		file_konfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConfigsResponse); i {
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
		file_konfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_konfig_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCollectionRequest); i {
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
		file_konfig_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCollectionResponse); i {
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
			RawDescriptor: file_konfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_konfig_proto_goTypes,
		DependencyIndexes: file_konfig_proto_depIdxs,
		MessageInfos:      file_konfig_proto_msgTypes,
	}.Build()
	File_konfig_proto = out.File
	file_konfig_proto_rawDesc = nil
	file_konfig_proto_goTypes = nil
	file_konfig_proto_depIdxs = nil
}