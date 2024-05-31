package homestay

import (
	"context"

	"microservices-go-zero/app/travel/cmd/api/internal/svc"
	"microservices-go-zero/app/travel/cmd/api/internal/types"
	"microservices-go-zero/app/travel/cmd/rpc/travel"
	"microservices-go-zero/common/tool"
	"microservices-go-zero/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDetailLogic {
	return &HomestayDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayDetailLogic) HomestayDetail(req *types.HomestayDetailReq) (resp *types.HomestayDetailResp, err error) {
	// use rpc to get the homestay detail of the specified id
	homestayResp, err := l.svcCtx.TravelRpc.HomestayDetail(l.ctx, &travel.HomestayDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get homestay detail fail"), "get homestay detail err, id: %d , err: %v ", req.Id, err)
	}

	var typeHomestay types.Homestay
	if homestayResp.Homestay != nil {
		_ = copier.Copy(&typeHomestay, homestayResp.Homestay)
		typeHomestay.FoodPrice = tool.Fen2Yuan(homestayResp.Homestay.FoodPrice)
		typeHomestay.HomestayPrice = tool.Fen2Yuan(homestayResp.Homestay.HomestayPrice)
		typeHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestayResp.Homestay.MarketHomestayPrice)

	}

	return &types.HomestayDetailResp{
		Homestay: typeHomestay,
	}, nil
}
