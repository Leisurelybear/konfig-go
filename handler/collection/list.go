package collection

import (
	"context"
	"konfig-go/common/errorutils"
	"konfig-go/common/logger"
	"konfig-go/database/mysql"
	"konfig-go/pb"
)

func List(ctx context.Context, request *pb.ListCollectionRequest) (*pb.ListCollectionResponse, error) {

	colls, err := mysql.CollectionDAO.SearchWithPaging(ctx, request.Name, request.CreatedBy, request.Sort, int(request.PageNum), int(request.PageSize))
	if err != nil {
		logger.Logger.Error(ctx, "error to find collection [name:%s,err:%v]", request.Name, err)
		return nil, errorutils.Wrap(err)
	}

	collIDs := make([]int64, 0, len(colls))
	respColls := []*pb.Collection{}
	for _, c := range colls {
		respColls = append(respColls, c.ConvertToDTO())
		collIDs = append(collIDs, int64(c.ID))
	}

	// get count of configs in each collection
	configCountsMap, err := mysql.ConfigDAO.CountByCollectionIDs(ctx, collIDs)
	if err != nil {
		logger.Logger.Error(ctx, "error to find configCounts [collIDs:%s,err:%v]", collIDs, err)
	}
	for _, coll := range respColls {
		coll.ConfigCount = configCountsMap[coll.ID]
	}

	resp := &pb.ListCollectionResponse{
		Sort:        request.Sort,
		PageNum:     request.PageNum,
		PageSize:    request.PageSize,
		Collections: respColls,
	}
	return resp, nil
}
