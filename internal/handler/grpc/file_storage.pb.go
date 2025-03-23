// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: internal/handler/grpc/file_storage.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_internal_handler_grpc_file_storage_proto_rawDescGZIP(), []int{0}
}

type GetPhotoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPhotoRequest) Reset() {
	*x = GetPhotoRequest{}
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhotoRequest) ProtoMessage() {}

func (x *GetPhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhotoRequest.ProtoReflect.Descriptor instead.
func (*GetPhotoRequest) Descriptor() ([]byte, []int) {
	return file_internal_handler_grpc_file_storage_proto_rawDescGZIP(), []int{1}
}

func (x *GetPhotoRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPhotosRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Names         []string               `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPhotosRequest) Reset() {
	*x = GetPhotosRequest{}
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPhotosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhotosRequest) ProtoMessage() {}

func (x *GetPhotosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhotosRequest.ProtoReflect.Descriptor instead.
func (*GetPhotosRequest) Descriptor() ([]byte, []int) {
	return file_internal_handler_grpc_file_storage_proto_rawDescGZIP(), []int{2}
}

func (x *GetPhotosRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type GetPhotoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Photo         []byte                 `protobuf:"bytes,1,opt,name=photo,proto3" json:"photo,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPhotoResponse) Reset() {
	*x = GetPhotoResponse{}
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPhotoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhotoResponse) ProtoMessage() {}

func (x *GetPhotoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhotoResponse.ProtoReflect.Descriptor instead.
func (*GetPhotoResponse) Descriptor() ([]byte, []int) {
	return file_internal_handler_grpc_file_storage_proto_rawDescGZIP(), []int{3}
}

func (x *GetPhotoResponse) GetPhoto() []byte {
	if x != nil {
		return x.Photo
	}
	return nil
}

func (x *GetPhotoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MetaData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Created       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Edited        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=edited,proto3" json:"edited,omitempty"`
	Extension     string                 `protobuf:"bytes,5,opt,name=extension,proto3" json:"extension,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetaData) Reset() {
	*x = MetaData{}
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaData) ProtoMessage() {}

func (x *MetaData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaData.ProtoReflect.Descriptor instead.
func (*MetaData) Descriptor() ([]byte, []int) {
	return file_internal_handler_grpc_file_storage_proto_rawDescGZIP(), []int{4}
}

func (x *MetaData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MetaData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetaData) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *MetaData) GetEdited() *timestamppb.Timestamp {
	if x != nil {
		return x.Edited
	}
	return nil
}

func (x *MetaData) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

type GetPhotosInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileInfo      []*MetaData            `protobuf:"bytes,1,rep,name=fileInfo,proto3" json:"fileInfo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPhotosInfoResponse) Reset() {
	*x = GetPhotosInfoResponse{}
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPhotosInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhotosInfoResponse) ProtoMessage() {}

func (x *GetPhotosInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhotosInfoResponse.ProtoReflect.Descriptor instead.
func (*GetPhotosInfoResponse) Descriptor() ([]byte, []int) {
	return file_internal_handler_grpc_file_storage_proto_rawDescGZIP(), []int{5}
}

func (x *GetPhotosInfoResponse) GetFileInfo() []*MetaData {
	if x != nil {
		return x.FileInfo
	}
	return nil
}

type PostPhotoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Extension     string                 `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
	Photo         []byte                 `protobuf:"bytes,3,opt,name=photo,proto3" json:"photo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostPhotoRequest) Reset() {
	*x = PostPhotoRequest{}
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostPhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPhotoRequest) ProtoMessage() {}

func (x *PostPhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPhotoRequest.ProtoReflect.Descriptor instead.
func (*PostPhotoRequest) Descriptor() ([]byte, []int) {
	return file_internal_handler_grpc_file_storage_proto_rawDescGZIP(), []int{6}
}

func (x *PostPhotoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostPhotoRequest) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

func (x *PostPhotoRequest) GetPhoto() []byte {
	if x != nil {
		return x.Photo
	}
	return nil
}

type PostPhotosRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Photos        []*PostPhotoRequest    `protobuf:"bytes,1,rep,name=photos,proto3" json:"photos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostPhotosRequest) Reset() {
	*x = PostPhotosRequest{}
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostPhotosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPhotosRequest) ProtoMessage() {}

func (x *PostPhotosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPhotosRequest.ProtoReflect.Descriptor instead.
func (*PostPhotosRequest) Descriptor() ([]byte, []int) {
	return file_internal_handler_grpc_file_storage_proto_rawDescGZIP(), []int{7}
}

func (x *PostPhotosRequest) GetPhotos() []*PostPhotoRequest {
	if x != nil {
		return x.Photos
	}
	return nil
}

type CreateOrReadyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ok            bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrReadyResponse) Reset() {
	*x = CreateOrReadyResponse{}
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrReadyResponse) ProtoMessage() {}

func (x *CreateOrReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handler_grpc_file_storage_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrReadyResponse.ProtoReflect.Descriptor instead.
func (*CreateOrReadyResponse) Descriptor() ([]byte, []int) {
	return file_internal_handler_grpc_file_storage_proto_rawDescGZIP(), []int{8}
}

func (x *CreateOrReadyResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_internal_handler_grpc_file_storage_proto protoreflect.FileDescriptor

var file_internal_handler_grpc_file_storage_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x06,
	0x65, 0x64, 0x69, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x64, 0x69, 0x74, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x43,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x5a, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x22,
	0x43, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0xf4, 0x01,
	0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0b,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x09, 0x50, 0x6f,
	0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x17, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_internal_handler_grpc_file_storage_proto_rawDescOnce sync.Once
	file_internal_handler_grpc_file_storage_proto_rawDescData []byte
)

func file_internal_handler_grpc_file_storage_proto_rawDescGZIP() []byte {
	file_internal_handler_grpc_file_storage_proto_rawDescOnce.Do(func() {
		file_internal_handler_grpc_file_storage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_handler_grpc_file_storage_proto_rawDesc), len(file_internal_handler_grpc_file_storage_proto_rawDesc)))
	})
	return file_internal_handler_grpc_file_storage_proto_rawDescData
}

var file_internal_handler_grpc_file_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_handler_grpc_file_storage_proto_goTypes = []any{
	(*Empty)(nil),                 // 0: grpc.Empty
	(*GetPhotoRequest)(nil),       // 1: grpc.GetPhotoRequest
	(*GetPhotosRequest)(nil),      // 2: grpc.GetPhotosRequest
	(*GetPhotoResponse)(nil),      // 3: grpc.GetPhotoResponse
	(*MetaData)(nil),              // 4: grpc.MetaData
	(*GetPhotosInfoResponse)(nil), // 5: grpc.GetPhotosInfoResponse
	(*PostPhotoRequest)(nil),      // 6: grpc.PostPhotoRequest
	(*PostPhotosRequest)(nil),     // 7: grpc.PostPhotosRequest
	(*CreateOrReadyResponse)(nil), // 8: grpc.CreateOrReadyResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_internal_handler_grpc_file_storage_proto_depIdxs = []int32{
	9, // 0: grpc.MetaData.created:type_name -> google.protobuf.Timestamp
	9, // 1: grpc.MetaData.edited:type_name -> google.protobuf.Timestamp
	4, // 2: grpc.GetPhotosInfoResponse.fileInfo:type_name -> grpc.MetaData
	6, // 3: grpc.PostPhotosRequest.photos:type_name -> grpc.PostPhotoRequest
	1, // 4: grpc.FileStorage.GetPhoto:input_type -> grpc.GetPhotoRequest
	0, // 5: grpc.FileStorage.GetAllPhotosInfo:input_type -> grpc.Empty
	6, // 6: grpc.FileStorage.PostPhoto:input_type -> grpc.PostPhotoRequest
	7, // 7: grpc.FileStorage.PostPhotos:input_type -> grpc.PostPhotosRequest
	3, // 8: grpc.FileStorage.GetPhoto:output_type -> grpc.GetPhotoResponse
	5, // 9: grpc.FileStorage.GetAllPhotosInfo:output_type -> grpc.GetPhotosInfoResponse
	0, // 10: grpc.FileStorage.PostPhoto:output_type -> grpc.Empty
	0, // 11: grpc.FileStorage.PostPhotos:output_type -> grpc.Empty
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_handler_grpc_file_storage_proto_init() }
func file_internal_handler_grpc_file_storage_proto_init() {
	if File_internal_handler_grpc_file_storage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_handler_grpc_file_storage_proto_rawDesc), len(file_internal_handler_grpc_file_storage_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_handler_grpc_file_storage_proto_goTypes,
		DependencyIndexes: file_internal_handler_grpc_file_storage_proto_depIdxs,
		MessageInfos:      file_internal_handler_grpc_file_storage_proto_msgTypes,
	}.Build()
	File_internal_handler_grpc_file_storage_proto = out.File
	file_internal_handler_grpc_file_storage_proto_goTypes = nil
	file_internal_handler_grpc_file_storage_proto_depIdxs = nil
}
