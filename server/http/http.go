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

	// 启用CORS
	handler := enableCORS(gwMux)

	// 启动HTTP服务器
	log.Println("HTTP server listening on port ", config.AppConfig.HttpPort)
	if err := http.ListenAndServe(config.AppConfig.HttpPort, handler); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}

// enableCORS 是一个中间件函数，用于启用CORS支持
func enableCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 允许的来源，可以根据需求进行调整
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// 允许的HTTP方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		// 允许的请求头
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 处理预检请求
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 继续处理请求
		handler.ServeHTTP(w, r)
	})
}
