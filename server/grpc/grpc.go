package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"konfig-go/common/auth"
	"konfig-go/conf"
	"konfig-go/pb"
	"log"
	"net"
)

func Run(config conf.Config) {
	lis, err := net.Listen("tcp", config.AppConfig.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(auth.AuthInterceptor), // auth
	)

	pb.RegisterKonfigServer(s, &server{})

	fmt.Println("gRPC server listening on port ", config.AppConfig.GRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
