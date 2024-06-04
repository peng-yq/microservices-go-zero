package svc

import (
	"github.com/hibiken/asynq"
	"microservices-go-zero/app/order/cmd/rpc/internal/config"
)

// create asynq client.
// package asynq provides a framework for Redis based distrubted task queue.
// asynq uses Redis as a message broker. 
// to connect to redis, specify the connection using one of RedisConnOpt types (host, pass).

func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass})
}
