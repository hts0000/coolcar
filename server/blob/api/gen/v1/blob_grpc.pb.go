// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: blob/api/gen/v1/blob.proto

package blobpb

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

// BlobServiceClient is the client API for BlobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlobServiceClient interface {
	CreateBlob(ctx context.Context, in *CreateBlobRequest, opts ...grpc.CallOption) (*CreateBlobResponse, error)
	GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (*GetBlobResponse, error)
	GetBlobURL(ctx context.Context, in *GetBlobURLRequest, opts ...grpc.CallOption) (*GetBlobURLResponse, error)
}

type blobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlobServiceClient(cc grpc.ClientConnInterface) BlobServiceClient {
	return &blobServiceClient{cc}
}

func (c *blobServiceClient) CreateBlob(ctx context.Context, in *CreateBlobRequest, opts ...grpc.CallOption) (*CreateBlobResponse, error) {
	out := new(CreateBlobResponse)
	err := c.cc.Invoke(ctx, "/blob.v1.BlobService/CreateBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobServiceClient) GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (*GetBlobResponse, error) {
	out := new(GetBlobResponse)
	err := c.cc.Invoke(ctx, "/blob.v1.BlobService/GetBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobServiceClient) GetBlobURL(ctx context.Context, in *GetBlobURLRequest, opts ...grpc.CallOption) (*GetBlobURLResponse, error) {
	out := new(GetBlobURLResponse)
	err := c.cc.Invoke(ctx, "/blob.v1.BlobService/GetBlobURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlobServiceServer is the server API for BlobService service.
// All implementations must embed UnimplementedBlobServiceServer
// for forward compatibility
type BlobServiceServer interface {
	CreateBlob(context.Context, *CreateBlobRequest) (*CreateBlobResponse, error)
	GetBlob(context.Context, *GetBlobRequest) (*GetBlobResponse, error)
	GetBlobURL(context.Context, *GetBlobURLRequest) (*GetBlobURLResponse, error)
	mustEmbedUnimplementedBlobServiceServer()
}

// UnimplementedBlobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlobServiceServer struct {
}

func (UnimplementedBlobServiceServer) CreateBlob(context.Context, *CreateBlobRequest) (*CreateBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlob not implemented")
}
func (UnimplementedBlobServiceServer) GetBlob(context.Context, *GetBlobRequest) (*GetBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlob not implemented")
}
func (UnimplementedBlobServiceServer) GetBlobURL(context.Context, *GetBlobURLRequest) (*GetBlobURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlobURL not implemented")
}
func (UnimplementedBlobServiceServer) mustEmbedUnimplementedBlobServiceServer() {}

// UnsafeBlobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlobServiceServer will
// result in compilation errors.
type UnsafeBlobServiceServer interface {
	mustEmbedUnimplementedBlobServiceServer()
}

func RegisterBlobServiceServer(s grpc.ServiceRegistrar, srv BlobServiceServer) {
	s.RegisterService(&BlobService_ServiceDesc, srv)
}

func _BlobService_CreateBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).CreateBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blob.v1.BlobService/CreateBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).CreateBlob(ctx, req.(*CreateBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobService_GetBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).GetBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blob.v1.BlobService/GetBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).GetBlob(ctx, req.(*GetBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobService_GetBlobURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlobURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).GetBlobURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blob.v1.BlobService/GetBlobURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).GetBlobURL(ctx, req.(*GetBlobURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlobService_ServiceDesc is the grpc.ServiceDesc for BlobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blob.v1.BlobService",
	HandlerType: (*BlobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlob",
			Handler:    _BlobService_CreateBlob_Handler,
		},
		{
			MethodName: "GetBlob",
			Handler:    _BlobService_GetBlob_Handler,
		},
		{
			MethodName: "GetBlobURL",
			Handler:    _BlobService_GetBlobURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blob/api/gen/v1/blob.proto",
}
