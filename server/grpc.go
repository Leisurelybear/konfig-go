package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"konfig-go/conf"
	"konfig-go/pb"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedKonfigServer
}

func (s *server) ListConfigs(context.Context, *pb.ListConfigsRequest) (*pb.ListConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigs not implemented")
}
func (s *server) ListCollection(context.Context, *pb.ListCollectionRequest) (*pb.ListCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollection not implemented")
}

func GrpcRun(config conf.Config) {
	lis, err := net.Listen("tcp", config.AppConfig.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterKonfigServer(s, &server{})

	fmt.Println("gRPC server listening on port ", config.AppConfig.GRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
