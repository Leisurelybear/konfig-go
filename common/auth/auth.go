package auth

import (
	"context"
	"google.golang.org/grpc"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 在此处进行权限校验的逻辑
	// 可以根据需求从上下文 ctx 中获取用户信息或其他相关信息进行校验

	// 如果校验失败，可以返回错误并拒绝请求
	// return nil, status.Errorf(codes.PermissionDenied, "Permission denied")

	// 如果校验通过，可以继续处理请求
	return handler(ctx, req)
}
