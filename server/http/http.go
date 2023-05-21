package http

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"konfig-go/conf"
	"konfig-go/pb"
	"log"
	"net/http"
)

func Run(config conf.Config) {
	//// Listen and Server in 0.0.0.0:8080
	//err := Init().Run(config.AppConfig.HttpPort)
	//if err != nil {
	//	return
	//}

	// 创建HTTP服务器
	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterKonfigHandlerFromEndpoint(context.Background(), gwMux, config.AppConfig.GRPCPort, opts); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	// 启动HTTP服务器
	log.Println("HTTP server listening on port ", config.AppConfig.HttpPort)
	if err := http.ListenAndServe(config.AppConfig.HttpPort, gwMux); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
