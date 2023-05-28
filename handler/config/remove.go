package config

import (
	"context"
	"konfig-go/common/errorutils"
	"konfig-go/common/logger"
	"konfig-go/database/mysql"
	"konfig-go/pb"
)

func RemoveConfig(ctx context.Context, request *pb.RemoveConfigRequest) error {
	cfg, err := mysql.ConfigDAO.GetByID(ctx, uint(request.ID))
	if err != nil || cfg == nil {
		logger.Logger.Error(ctx, "error to find config: [cfgID:%v, err:%v]", request.ID, err)
		return errorutils.ErrNotFound
	}
	err = mysql.ConfigDAO.Delete(ctx, cfg)
	if err != nil {
		logger.Logger.Error(ctx, "error to delete config: [cfgID:%v, err:%v]", request.ID, err)
		return err
	}
	return nil
}
