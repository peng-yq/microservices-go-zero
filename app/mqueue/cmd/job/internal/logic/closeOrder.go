package logic

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"

	"microservices-go-zero/app/mqueue/cmd/job/internal/svc"
	"microservices-go-zero/app/mqueue/cmd/job/jobtype"
	"microservices-go-zero/app/order/cmd/rpc/order"
	"microservices-go-zero/app/order/model"
	"microservices-go-zero/common/xerr"
)

var ErrCloseOrderFal = xerr.NewErrMsg("close order fail")

// CloseHomestayOrderHandler close no pay homestayOrder
type CloseHomestayOrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewCloseHomestayOrderHandler(svcCtx *svc.ServiceContext) *CloseHomestayOrderHandler {
	return &CloseHomestayOrderHandler{
		svcCtx:svcCtx,
	}
}

// defer close no pay homestayOrder: if return err != nil , asynq will retry
func (l *CloseHomestayOrderHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p jobtype.DeferCloseHomestayOrderPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(ErrCloseOrderFal, "closeHomestayOrderStateMqHandler payload err: %v, payLoad: %+v", err, t.Payload())
	}

	resp, err := l.svcCtx.OrderRpc.HomestayOrderDetail(ctx, &order.HomestayOrderDetailReq{
		Sn: p.Sn,
	})
	if err != nil || resp.HomestayOrder == nil {
		return errors.Wrapf(ErrCloseOrderFal, "closeHomestayOrderStateMqHandler get order fail or order no exists err: %v, sn: %s, HomestayOrder: %+v", err, p.Sn, resp.HomestayOrder)
	}

	// update the homestay order trade status to cancel the unpaid order
	if resp.HomestayOrder.TradeState == model.HomestayOrderTradeStateWaitPay {
		_, err := l.svcCtx.OrderRpc.UpdateHomestayOrderTradeState(ctx, &order.UpdateHomestayOrderTradeStateReq{
			Sn:         p.Sn,
			TradeState: model.HomestayOrderTradeStateCancel,
		})
		if err != nil {
			return errors.Wrapf(ErrCloseOrderFal, "CloseHomestayOrderHandler close order fail err: %v, sn: %s ", err, p.Sn)
		}
	}

	return nil
}
