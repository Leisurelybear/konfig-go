package config

import (
	"context"
	"konfig-go/pb"
)

func List(ctx context.Context, request *pb.ListConfigsRequest) (*pb.ListConfigsResponse, error) {

	return &pb.ListConfigsResponse{
		CollectionID:   1,
		CollectionName: "asdasd",
		Sort:           1,
		PageNum:        0,
		PageSize:       0,
		Configs:        nil,
	}, nil
}
