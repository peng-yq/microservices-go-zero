package homestayBussiness

import (
	"context"

	"microservices-go-zero/app/travel/cmd/api/internal/svc"
	"microservices-go-zero/app/travel/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GoodBossLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodBossLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodBossLogic {
	return &GoodBossLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodBossLogic) GoodBoss(req *types.GoodBossReq) (*types.GoodBossResp, error) {
	// to-do
	var resp []types.HomestayBusinessBoss
	return &types.GoodBossResp{
		List: resp,
	}, nil
}
