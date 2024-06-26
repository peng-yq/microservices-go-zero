package svc

import (
	"fmt"
	"time"

	"microservices-go-zero/app/mqueue/cmd/scheduler/internal/config"

	"github.com/hibiken/asynq"
)

// create scheduler
func newScheduler(c config.Config) *asynq.Scheduler {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr: c.Redis.Host,
			Password: c.Redis.Pass,
		}, &asynq.SchedulerOpts{
			Location: location,
			EnqueueErrorHandler: func(task *asynq.Task, opts []asynq.Option, err error) {
				fmt.Printf("Scheduler EnqueueErrorHandler <<<<<<<===>>>>> err: %+v, task: %+v",err,task)
			},
		})
}
