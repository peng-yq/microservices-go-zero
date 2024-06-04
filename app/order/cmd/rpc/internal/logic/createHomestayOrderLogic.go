package logic

import (
	"context"
	"time"
	"strings"
	"encoding/json"

	"microservices-go-zero/app/order/cmd/rpc/internal/svc"
	"microservices-go-zero/app/order/cmd/rpc/pb"
	"microservices-go-zero/app/order/model"
	"microservices-go-zero/app/travel/cmd/rpc/travel"
	"microservices-go-zero/app/mqueue/cmd/job/jobtype"
	"microservices-go-zero/common/tool"
	"microservices-go-zero/common/uniqueid"
	"microservices-go-zero/common/xerr"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

// defer close order time (mins)
const CloseOrderTimeMinutes = 30  

type CreateHomestayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateHomestayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHomestayOrderLogic {
	return &CreateHomestayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// create homestay order
func (l *CreateHomestayOrderLogic) CreateHomestayOrder(in *pb.CreateHomestayOrderReq) (*pb.CreateHomestayOrderResp, error) {
	// create Order
	if in.LiveEndTime <= in.LiveStartTime {
		return nil, errors.Wrapf(xerr.NewErrMsg("Stay at least one night"), "Place an order at a B&B. The end time of your stay must be greater than the start time. in: %+v", in)
	}

	resp, err := l.svcCtx.TravelRpc.HomestayDetail(l.ctx, &travel.HomestayDetailReq{
		Id: in.HomestayId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to query the record"), "Failed to query the record rpc HomestayDetail fail, homestayId: %d, err: %v", in.HomestayId, err)
	}
	if resp.Homestay == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("This record does not exist"), "This record does not exist, homestayId: %d ", in.HomestayId)
	}

	var cover string
	if len(resp.Homestay.Banner) > 0 {
		cover = strings.Split(resp.Homestay.Banner, ",")[0]
	}

	order := new(model.HomestayOrder)
	// gen order id
	order.Sn = uniqueid.GenSn(uniqueid.SN_PREFIX_HOMESTAY_ORDER)
	order.UserId = in.UserId
	order.HomestayId = in.HomestayId
	order.Title = resp.Homestay.Title
	order.SubTitle = resp.Homestay.SubTitle
	order.Cover = cover
	order.Info = resp.Homestay.Info
	order.PeopleNum = resp.Homestay.PeopleNum
	order.RowType = resp.Homestay.RowType
	order.HomestayPrice = resp.Homestay.HomestayPrice
	order.MarketHomestayPrice = resp.Homestay.MarketHomestayPrice
	order.HomestayBusinessId = resp.Homestay.HomestayBusinessId
	order.HomestayUserId = resp.Homestay.UserId
	order.LivePeopleNum = in.LivePeopleNum
	order.TradeState = model.HomestayOrderTradeStateWaitPay
	// gen trade code
	order.TradeCode = tool.Krand(8, tool.KC_RAND_KIND_ALL)
	order.Remark = in.Remark
	order.FoodInfo = resp.Homestay.FoodInfo
	order.FoodPrice = resp.Homestay.FoodPrice
	order.LiveStartDate = time.Unix(in.LiveStartTime, 0)
	order.LiveEndDate = time.Unix(in.LiveEndTime, 0)

	// stayed a few days in total
	liveDays := int64(order.LiveEndDate.Sub(order.LiveStartDate).Seconds() / 86400) 

	// calculate the total price of the B&B
	order.HomestayTotalPrice = int64(resp.Homestay.HomestayPrice * liveDays) 
	if in.IsFood {
		order.NeedFood = model.HomestayOrderNeedFoodYes
		// calculate the total price of the meal.
		order.FoodTotalPrice = int64(resp.Homestay.FoodPrice * in.LivePeopleNum * liveDays)
	}

	// calculate total order price
	order.OrderTotalPrice = order.HomestayTotalPrice + order.FoodTotalPrice 

	_, err = l.svcCtx.HomestayOrderModel.Insert(l.ctx, nil, order)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Order Database Exception order: %+v, err: %v", order, err)
	}

	// delayed closing of order tasks.
	payload, err := json.Marshal(jobtype.DeferCloseHomestayOrderPayload{Sn: order.Sn})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create defer close order task json Marshal fail err: %+v, sn: %s", err, order.Sn)
	}else{
		_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseHomestayOrder, payload), asynq.ProcessIn(CloseOrderTimeMinutes * time.Minute))
		if err != nil {
			logx.WithContext(l.ctx).Errorf("create defer close order task insert queue fail err: %+v, sn: %s", err, order.Sn)
		}
	}

	return &pb.CreateHomestayOrderResp{
		Sn: order.Sn,
	}, nil
}
