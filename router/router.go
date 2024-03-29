package router

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
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
func (s *Server) UpdateConfig(ctx context.Context, request *pb.UpdateConfigRequest) (*pb.UpdateConfigResponse, error) {
	return config.Update(ctx, request)
}
func (s *Server) RemoveConfig(ctx context.Context, request *pb.RemoveConfigRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, config.RemoveConfig(ctx, request)
}
