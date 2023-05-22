// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// KonfigClient is the client API for Konfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KonfigClient interface {
	ListConfigs(ctx context.Context, in *ListConfigsRequest, opts ...grpc.CallOption) (*ListConfigsResponse, error)
	UpsertConfig(ctx context.Context, in *UpsertConfigRequest, opts ...grpc.CallOption) (*UpsertConfigResponse, error)
	ListCollection(ctx context.Context, in *ListCollectionRequest, opts ...grpc.CallOption) (*ListCollectionResponse, error)
	CollectionDetail(ctx context.Context, in *CollectionDetailRequest, opts ...grpc.CallOption) (*CollectionDetailResponse, error)
}

type konfigClient struct {
	cc grpc.ClientConnInterface
}

func NewKonfigClient(cc grpc.ClientConnInterface) KonfigClient {
	return &konfigClient{cc}
}

func (c *konfigClient) ListConfigs(ctx context.Context, in *ListConfigsRequest, opts ...grpc.CallOption) (*ListConfigsResponse, error) {
	out := new(ListConfigsResponse)
	err := c.cc.Invoke(ctx, "/pb.Konfig/ListConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *konfigClient) UpsertConfig(ctx context.Context, in *UpsertConfigRequest, opts ...grpc.CallOption) (*UpsertConfigResponse, error) {
	out := new(UpsertConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.Konfig/UpsertConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *konfigClient) ListCollection(ctx context.Context, in *ListCollectionRequest, opts ...grpc.CallOption) (*ListCollectionResponse, error) {
	out := new(ListCollectionResponse)
	err := c.cc.Invoke(ctx, "/pb.Konfig/ListCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *konfigClient) CollectionDetail(ctx context.Context, in *CollectionDetailRequest, opts ...grpc.CallOption) (*CollectionDetailResponse, error) {
	out := new(CollectionDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.Konfig/CollectionDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KonfigServer is the server API for Konfig service.
// All implementations must embed UnimplementedKonfigServer
// for forward compatibility
type KonfigServer interface {
	ListConfigs(context.Context, *ListConfigsRequest) (*ListConfigsResponse, error)
	UpsertConfig(context.Context, *UpsertConfigRequest) (*UpsertConfigResponse, error)
	ListCollection(context.Context, *ListCollectionRequest) (*ListCollectionResponse, error)
	CollectionDetail(context.Context, *CollectionDetailRequest) (*CollectionDetailResponse, error)
	mustEmbedUnimplementedKonfigServer()
}

// UnimplementedKonfigServer must be embedded to have forward compatible implementations.
type UnimplementedKonfigServer struct {
}

func (UnimplementedKonfigServer) ListConfigs(context.Context, *ListConfigsRequest) (*ListConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigs not implemented")
}
func (UnimplementedKonfigServer) UpsertConfig(context.Context, *UpsertConfigRequest) (*UpsertConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertConfig not implemented")
}
func (UnimplementedKonfigServer) ListCollection(context.Context, *ListCollectionRequest) (*ListCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollection not implemented")
}
func (UnimplementedKonfigServer) CollectionDetail(context.Context, *CollectionDetailRequest) (*CollectionDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionDetail not implemented")
}
func (UnimplementedKonfigServer) mustEmbedUnimplementedKonfigServer() {}

// UnsafeKonfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KonfigServer will
// result in compilation errors.
type UnsafeKonfigServer interface {
	mustEmbedUnimplementedKonfigServer()
}

func RegisterKonfigServer(s grpc.ServiceRegistrar, srv KonfigServer) {
	s.RegisterService(&Konfig_ServiceDesc, srv)
}

func _Konfig_ListConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KonfigServer).ListConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Konfig/ListConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KonfigServer).ListConfigs(ctx, req.(*ListConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Konfig_UpsertConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KonfigServer).UpsertConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Konfig/UpsertConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KonfigServer).UpsertConfig(ctx, req.(*UpsertConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Konfig_ListCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KonfigServer).ListCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Konfig/ListCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KonfigServer).ListCollection(ctx, req.(*ListCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Konfig_CollectionDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KonfigServer).CollectionDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Konfig/CollectionDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KonfigServer).CollectionDetail(ctx, req.(*CollectionDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Konfig_ServiceDesc is the grpc.ServiceDesc for Konfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Konfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Konfig",
	HandlerType: (*KonfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListConfigs",
			Handler:    _Konfig_ListConfigs_Handler,
		},
		{
			MethodName: "UpsertConfig",
			Handler:    _Konfig_UpsertConfig_Handler,
		},
		{
			MethodName: "ListCollection",
			Handler:    _Konfig_ListCollection_Handler,
		},
		{
			MethodName: "CollectionDetail",
			Handler:    _Konfig_CollectionDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "konfig.proto",
}
