package collection

import (
	"context"
	"konfig-go/common/errorutils"
	"konfig-go/common/logger"
	"konfig-go/database/mysql"
	"konfig-go/pb"
)

func Detail(ctx context.Context, request *pb.CollectionDetailRequest) (*pb.CollectionDetailResponse, error) {
	collection, err := mysql.CollectionDAO.GetById(ctx, request.CollectionID)
	if err != nil {
		logger.Logger.Error(ctx, "error to find collection [cid:%d,err:%v]", request.CollectionID, err)
		return nil, errorutils.Wrap(err)
	}

	configs, err := mysql.ConfigDAO.GetByCollectionID(ctx, int32(request.CollectionID))
	if err != nil {
		logger.Logger.Error(ctx, "error to find collection [cid:%d,err:%v]", request.CollectionID, err)
		return nil, errorutils.Wrap(err)
	}

	respConfs := []*pb.Config{}
	for _, c := range configs {
		respConfs = append(respConfs, c.ConvertToDTO())
	}

	resp := &pb.CollectionDetailResponse{Collection: collection.ConvertToDTO()}
	resp.Collection.Configs = respConfs
	return resp, nil
}
