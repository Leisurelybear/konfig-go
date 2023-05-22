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
	err := checkRequest(request)
	if err != nil {
		return nil, errorutils.Wrap(err)
	}

	//todo: check collection is valid

	cfg := &config.Config{
		Model: gorm.Model{
			ID: uint(request.ID),
		},
		CollectionId: request.CollectionID,
		Name:         request.Name,
		Key:          request.Key,
		Value:        request.Value,
		Version:      collection.Draft,
		UpdatedBy:    "",
	}
	err = mysql.ConfigDAO.Update(ctx, cfg)
	if err != nil {
		logger.Logger.Error(ctx, "error to upsert config: [cfgID:%v]", cfg.ID)
		return nil, err
	}

	// todo:not update createBy
	// todo:set collection as draft
	return &pb.UpsertConfigResponse{Config: cfg.ConvertToDTO()}, nil
}

func checkRequest(request *pb.UpsertConfigRequest) error {
	return nil
}
