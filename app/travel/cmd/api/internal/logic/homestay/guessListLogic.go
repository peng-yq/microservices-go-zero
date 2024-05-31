package homestay

import (
	"context"

	"microservices-go-zero/app/travel/cmd/api/internal/svc"
	"microservices-go-zero/app/travel/cmd/api/internal/types"
	"microservices-go-zero/common/tool"
	"microservices-go-zero/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GuessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGuessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GuessListLogic {
	return &GuessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GuessList only return the first 5 records in the table whose id is bigger than 0
// not really guss you like list, need access to recommendation algorithms
func (l *GuessListLogic) GuessList(req *types.GuessListReq) (*types.GuessListResp, error) {
	var resp []types.Homestay

	list, err := l.svcCtx.HomestayModel.FindPageListByIdASC(l.ctx, l.svcCtx.HomestayModel.SelectBuilder(), 0, 5)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GuessList db err req : %+v , err : %v", req, err)
	}

	if len(list) > 0 {
		for _, homestay := range list {
			var typeHomestay types.Homestay
			_ = copier.Copy(&typeHomestay, homestay)
			typeHomestay.FoodPrice = tool.Fen2Yuan(homestay.FoodPrice)
			typeHomestay.HomestayPrice = tool.Fen2Yuan(homestay.HomestayPrice)
			typeHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestay.MarketHomestayPrice)
			resp = append(resp, typeHomestay)
		}
	}

	return &types.GuessListResp{
		List: resp,
	}, nil
}
