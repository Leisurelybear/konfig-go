package config

import (
	"context"
	"gorm.io/gorm"
	"konfig-go/common/errorutils"
	"konfig-go/common/logger"
	"konfig-go/database/mysql"
	"konfig-go/database/mysql/collection"
	"konfig-go/database/mysql/config"
	"konfig-go/pb"
)

func Upsert(ctx context.Context, request *pb.UpsertConfigRequest) (*pb.UpsertConfigResponse, error) {
	coll, err := checkRequest(ctx, request)
	if err != nil {
		return nil, errorutils.Wrap(err)
	}

	newConfig := &config.Config{
		Model: gorm.Model{
			ID: uint(request.ID),
		},
		CollectionId: request.CollectionID,
		Name:         request.Name,
		Key:          request.Key,
		Value:        request.Value,
		Version:      collection.Draft,
		UpdatedBy:    "",
		CreatedBy:    "",
	}
	if request.ID > 0 {
		oldConfig, er := mysql.ConfigDAO.GetByID(ctx, uint(request.ID))
		if er != nil {
			logger.Logger.Error(ctx, "error to get old conf[%v]", er)
		} else {
			newConfig.CreatedBy = oldConfig.CreatedBy
		}
	}
	err = mysql.ConfigDAO.Update(ctx, newConfig)
	if err != nil {
		logger.Logger.Error(ctx, "error to upsert config: [cfgID:%v]", newConfig.ID)
		return nil, err
	}

	err = mysql.CollectionDAO.Update(ctx, coll)
	if err != nil {
		logger.Logger.Error(ctx, "error to set collection to draft[er:%v]", err)
	}

	return &pb.UpsertConfigResponse{Config: newConfig.ConvertToDTO()}, nil
}

func checkRequest(ctx context.Context, request *pb.UpsertConfigRequest) (*collection.Collection, error) {
	if request.CollectionID == 0 || request.Key == "" || request.Name == "" {
		logger.Logger.Error(ctx, "parameter is invalid[req:%v]", request)
		return nil, errorutils.ErrInvalidParameters
	}
	coll, err := mysql.CollectionDAO.GetById(ctx, request.CollectionID)
	if err != nil || coll == nil {
		logger.Logger.Error(ctx, "collection is invalid[err:%v]", err)
		return nil, err
	}

	cfg, err := mysql.ConfigDAO.GetByCollectionIDAndKey(ctx, request.CollectionID, request.Key)
	if err != nil {
		logger.Logger.Error(ctx, "query collection failed[err:%v]", err)
		return nil, errorutils.ErrInvalidDuplicated
	}
	if cfg != nil {
		logger.Logger.Error(ctx, "config is duplicated[cid:%v, cname:%s, k:%v]", cfg.ID, cfg.Name, request.Key)
		return nil, errorutils.ErrInvalidDuplicated
	}
	return coll, nil
}
