package homestay

import (
	"context"

	"microservices-go-zero/app/travel/cmd/api/internal/svc"
	"microservices-go-zero/app/travel/cmd/api/internal/types"
	"microservices-go-zero/app/travel/model"
	"microservices-go-zero/common/tool"
	"microservices-go-zero/common/xerr"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type HomestayListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayListLogic {
	return &HomestayListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayListLogic) HomestayList(req *types.HomestayListReq) (*types.HomestayListResp, error) {
	whereBuilder := l.svcCtx.HomestayActivityModel.SelectBuilder().Where(squirrel.Eq{
		"row_type":   model.HomestayActivityPreferredType,
		"row_status": model.HomestayActivityUpStatus,
	})
	// returns a list of homestay activities with pagesize number of the specified page and sorted by data_id desc
	homestayActivityList, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, whereBuilder, req.Page, req.PageSize, "data_id desc")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get activity homestay id set fail rowType: %s, err: %v", model.HomestayActivityPreferredType, err)
	}

	var resp []types.Homestay
	if len(homestayActivityList) > 0 { 
		// MapReduce has three main parameters. 
		// The first parameter is generate, which is used to generate data. 
		// The second parameter is mapper, which is used to process data. 
		// The third parameter is reducer, which is used to aggregate and return the mapped data. 
		// The number of concurrent processing threads can also be set through the opts option.
		mr.MapReduceVoid(func(source chan<- interface{}) {
			for _, homestayActivity := range homestayActivityList {
				source <- homestayActivity.DataId
			}
		}, func(item interface{}, writer mr.Writer[*model.Homestay], cancel func(error)) {
			id := item.(int64)

			homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, id)
			if err != nil && err != model.ErrNotFound {
				logx.WithContext(l.ctx).Errorf("ActivityHomestayListLogic ActivityHomestayList get data failed, id: %d, err: %v", id, err)
				return
			}
			writer.Write(homestay)
		}, func(pipe <-chan *model.Homestay, cancel func(error)) {

			for homestay := range pipe {
				var tyHomestay types.Homestay
				_ = copier.Copy(&tyHomestay, homestay)
				tyHomestay.FoodPrice = tool.Fen2Yuan(homestay.FoodPrice)
				tyHomestay.HomestayPrice = tool.Fen2Yuan(homestay.HomestayPrice)
				tyHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestay.MarketHomestayPrice)
				resp = append(resp, tyHomestay)
			}
		})
	}

	return &types.HomestayListResp{
		List: resp,
	}, nil
}
