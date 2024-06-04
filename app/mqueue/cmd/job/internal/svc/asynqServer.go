package svc

import (
	"fmt"
	"github.com/hibiken/asynq"
	"microservices-go-zero/app/mqueue/cmd/job/internal/config"
)

// IsFailure: 
// Predicate function to determine whether the error returned from Handler is a failure.
// If the function returns false, Server will not increment the retried counter for the task,
// and Server won't record the queue stats (processed and failed stats) to avoid skewing the error
// rate of the queue.
//
// By default, if the given error is non-nil the function returns true.

func newAsynqServer(c config.Config) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task IsFailure err: %+v \n",err)
				return true
			},
			Concurrency: 20, // max concurrent process job task num
		},
	)
}

