package homestay

import (
	"context"

	"microservices-go-zero/app/travel/cmd/api/internal/svc"
	"microservices-go-zero/app/travel/cmd/api/internal/types"
	"microservices-go-zero/common/tool"
	"microservices-go-zero/common/xerr"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessListLogic {
	return &BusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BusinessListLogic) BusinessList(req *types.BusinessListReq) (*types.BusinessListResp, error) {
	// find the first req.PageSize records in the table whose id is less than req.LastId
	whereBuilder := l.svcCtx.HomestayModel.SelectBuilder().Where(squirrel.Eq{"homestay_business_id": req.HomestayBusinessId})
	list, err := l.svcCtx.HomestayModel.FindPageListByIdDESC(l.ctx, whereBuilder, req.LastId, req.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HomestayBusinessId: %d, err: %v", req.HomestayBusinessId, err)
	}

	var resp []types.Homestay
	if len(list) > 0 {
		for _, homestay := range list {
			var typeHomestay types.Homestay
			_ = copier.Copy(&typeHomestay, homestay)
			// convert fen to yuan of the prices
			typeHomestay.FoodPrice = tool.Fen2Yuan(homestay.FoodPrice)
			typeHomestay.HomestayPrice = tool.Fen2Yuan(homestay.HomestayPrice)
			typeHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestay.MarketHomestayPrice)
			resp = append(resp, typeHomestay)
		}
	}

	return &types.BusinessListResp{
		List: resp,
	}, nil
}
