package logic

import (
	"context"

	"github.com/hibiken/asynq"

	"microservices-go-zero/app/mqueue/cmd/job/internal/svc"
	"microservices-go-zero/app/mqueue/cmd/job/jobtype"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register job: mux maps a type to a handler
func (l *CronJob) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	// scheduler job
	mux.Handle(jobtype.ScheduleSettleRecord, NewSettleRecordHandler(l.svcCtx))

	// defer job
	mux.Handle(jobtype.DeferCloseHomestayOrder, NewCloseHomestayOrderHandler(l.svcCtx))

	// payment success job
	mux.Handle(jobtype.MsgPaySuccessNotifyUser, NewPaySuccessNotifyUserHandler(l.svcCtx))

	return mux
}


