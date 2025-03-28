// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: internal/handler/grpc/file_storage.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FileStorage_GetPhoto_FullMethodName         = "/grpc.FileStorage/GetPhoto"
	FileStorage_GetAllPhotosInfo_FullMethodName = "/grpc.FileStorage/GetAllPhotosInfo"
	FileStorage_PostPhoto_FullMethodName        = "/grpc.FileStorage/PostPhoto"
	FileStorage_PostPhotos_FullMethodName       = "/grpc.FileStorage/PostPhotos"
)

// FileStorageClient is the client API for FileStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileStorageClient interface {
	GetPhoto(ctx context.Context, in *GetPhotoRequest, opts ...grpc.CallOption) (*GetPhotoResponse, error)
	GetAllPhotosInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPhotosInfoResponse, error)
	PostPhoto(ctx context.Context, in *PostPhotoRequest, opts ...grpc.CallOption) (*Empty, error)
	PostPhotos(ctx context.Context, in *PostPhotosRequest, opts ...grpc.CallOption) (*Empty, error)
}

type fileStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewFileStorageClient(cc grpc.ClientConnInterface) FileStorageClient {
	return &fileStorageClient{cc}
}

func (c *fileStorageClient) GetPhoto(ctx context.Context, in *GetPhotoRequest, opts ...grpc.CallOption) (*GetPhotoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPhotoResponse)
	err := c.cc.Invoke(ctx, FileStorage_GetPhoto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageClient) GetAllPhotosInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPhotosInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPhotosInfoResponse)
	err := c.cc.Invoke(ctx, FileStorage_GetAllPhotosInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageClient) PostPhoto(ctx context.Context, in *PostPhotoRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, FileStorage_PostPhoto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageClient) PostPhotos(ctx context.Context, in *PostPhotosRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, FileStorage_PostPhotos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileStorageServer is the server API for FileStorage service.
// All implementations must embed UnimplementedFileStorageServer
// for forward compatibility.
type FileStorageServer interface {
	GetPhoto(context.Context, *GetPhotoRequest) (*GetPhotoResponse, error)
	GetAllPhotosInfo(context.Context, *Empty) (*GetPhotosInfoResponse, error)
	PostPhoto(context.Context, *PostPhotoRequest) (*Empty, error)
	PostPhotos(context.Context, *PostPhotosRequest) (*Empty, error)
	mustEmbedUnimplementedFileStorageServer()
}

// UnimplementedFileStorageServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileStorageServer struct{}

func (UnimplementedFileStorageServer) GetPhoto(context.Context, *GetPhotoRequest) (*GetPhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhoto not implemented")
}
func (UnimplementedFileStorageServer) GetAllPhotosInfo(context.Context, *Empty) (*GetPhotosInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPhotosInfo not implemented")
}
func (UnimplementedFileStorageServer) PostPhoto(context.Context, *PostPhotoRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostPhoto not implemented")
}
func (UnimplementedFileStorageServer) PostPhotos(context.Context, *PostPhotosRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostPhotos not implemented")
}
func (UnimplementedFileStorageServer) mustEmbedUnimplementedFileStorageServer() {}
func (UnimplementedFileStorageServer) testEmbeddedByValue()                     {}

// UnsafeFileStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileStorageServer will
// result in compilation errors.
type UnsafeFileStorageServer interface {
	mustEmbedUnimplementedFileStorageServer()
}

func RegisterFileStorageServer(s grpc.ServiceRegistrar, srv FileStorageServer) {
	// If the following call pancis, it indicates UnimplementedFileStorageServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileStorage_ServiceDesc, srv)
}

func _FileStorage_GetPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServer).GetPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileStorage_GetPhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServer).GetPhoto(ctx, req.(*GetPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorage_GetAllPhotosInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServer).GetAllPhotosInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileStorage_GetAllPhotosInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServer).GetAllPhotosInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorage_PostPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServer).PostPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileStorage_PostPhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServer).PostPhoto(ctx, req.(*PostPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorage_PostPhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostPhotosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServer).PostPhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileStorage_PostPhotos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServer).PostPhotos(ctx, req.(*PostPhotosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileStorage_ServiceDesc is the grpc.ServiceDesc for FileStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.FileStorage",
	HandlerType: (*FileStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPhoto",
			Handler:    _FileStorage_GetPhoto_Handler,
		},
		{
			MethodName: "GetAllPhotosInfo",
			Handler:    _FileStorage_GetAllPhotosInfo_Handler,
		},
		{
			MethodName: "PostPhoto",
			Handler:    _FileStorage_PostPhoto_Handler,
		},
		{
			MethodName: "PostPhotos",
			Handler:    _FileStorage_PostPhotos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/handler/grpc/file_storage.proto",
}
