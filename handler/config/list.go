package config

import (
	"context"
	"konfig-go/common/errorutils"
	"konfig-go/common/logger"
	"konfig-go/database/mysql"
	"konfig-go/pb"
)

func List(ctx context.Context, request *pb.ListConfigsRequest) (*pb.ListConfigsResponse, error) {

	collection, err := mysql.CollectionDAO.GetById(request.CollectionID)
	if err != nil {
		logger.Logger.Error(ctx, "error to find collection [cid:%d,err:%v]", request.CollectionID, err)
		return nil, errorutils.Wrap(err)
	}

	configs, err := mysql.ConfigDAO.GetByCollectionID(int32(request.CollectionID))
	if err != nil {
		return nil, errorutils.Wrap(err)
	}

	respConfigs := []*pb.Config{}
	for _, c := range configs {
		respConfigs = append(respConfigs, &pb.Config{
			ID:           int64(c.ID),
			CollectionID: c.CollectionId,
			Name:         c.Name,
			Key:          c.Key,
			Value:        c.Value,
			CreatedBy:    c.CreatedBy,
			UpdatedBy:    c.UpdatedBy,
			CreatedAt:    c.CreatedAt.Unix(),
			UpdatedAt:    c.UpdatedAt.Unix(),
		})
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
