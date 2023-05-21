package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"konfig-go/handler/collection"
	"konfig-go/handler/config"
	"konfig-go/pb"
)

type server struct {
	pb.UnimplementedKonfigServer
}

func (s *server) ListConfigs(ctx context.Context, request *pb.ListConfigsRequest) (*pb.ListConfigsResponse, error) {
	return config.List(ctx, request)
}
func (s *server) ListCollection(ctx context.Context, request *pb.ListCollectionRequest) (*pb.ListCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollection not implemented")
}
func (s *server) CollectionDetail(ctx context.Context, request *pb.CollectionDetailRequest) (*pb.CollectionDetailResponse, error) {
	return collection.Detail(ctx, request)
}
