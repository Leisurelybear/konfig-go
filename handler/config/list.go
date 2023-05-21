package config

import (
	"context"
	"konfig-go/common/errorutils"
	"konfig-go/common/logger"
	"konfig-go/database/mysql"
	"konfig-go/pb"
)

func List(ctx context.Context, request *pb.ListConfigsRequest) (*pb.ListConfigsResponse, error) {

	collection, err := mysql.CollectionDAO.GetById(ctx, request.CollectionID)
	if err != nil {
		logger.Logger.Error(ctx, "error to find collection [cid:%d,err:%v]", request.CollectionID, err)
		return nil, errorutils.Wrap(err)
	}

	configs, err := mysql.ConfigDAO.SearchWithPaging(ctx, request.CollectionID, request.Name, request.Sort, int(request.PageNum), int(request.PageSize))
	if err != nil {
		logger.Logger.Error(ctx, "error to find collection [cid:%d,err:%v]", request.CollectionID, err)
		return nil, errorutils.Wrap(err)
	}

	respConfigs := []*pb.Config{}
	for _, c := range configs {
		respConfigs = append(respConfigs, c.ConvertToDTO())
	}

	return &pb.ListConfigsResponse{
		CollectionID:   request.CollectionID,
		CollectionName: collection.Name,
		Sort:           request.Sort,
		PageNum:        request.PageNum,
		PageSize:       request.PageSize,
		Configs:        respConfigs,
	}, nil
}
