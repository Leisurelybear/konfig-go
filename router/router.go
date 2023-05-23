package router

import (
	"context"
	"konfig-go/handler/collection"
	"konfig-go/handler/config"
	"konfig-go/pb"
)

type Server struct {
	pb.UnimplementedKonfigServer
}

func (s *Server) ListConfigs(ctx context.Context, request *pb.ListConfigsRequest) (*pb.ListConfigsResponse, error) {
	return config.List(ctx, request)
}
func (s *Server) ListCollection(ctx context.Context, request *pb.ListCollectionRequest) (*pb.ListCollectionResponse, error) {
	return collection.List(ctx, request)
}
func (s *Server) CollectionDetail(ctx context.Context, request *pb.CollectionDetailRequest) (*pb.CollectionDetailResponse, error) {
	return collection.Detail(ctx, request)
}
func (s *Server) UpsertConfig(ctx context.Context, request *pb.UpsertConfigRequest) (*pb.UpsertConfigResponse, error) {
	return config.Upsert(ctx, request)
}
