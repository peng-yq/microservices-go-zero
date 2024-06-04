package listen

import (
	"context"

	"microservices-go-zero/app/order/cmd/mq/internal/config"
	"microservices-go-zero/app/order/cmd/mq/internal/svc"
	kqMq "microservices-go-zero/app/order/cmd/mq/internal/mqs/kq"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// pub sub use kq (kafka)
// mustnewqueue optional parameters:
// commitInterval: The interval between commits to the kafka broker, the default is 1s
// queueCapacity: The length of the internal queue of kafka
// maxWait: The maximum time to wait for new data to arrive when getting data in batches from kafka.
// metrics: Report the consumption time of each message, 
// which will be initialized internally by default and generally does not need to be specified
func KqMqConsumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		// listening for payment update status consumers
		kq.MustNewQueue(c.PaymentUpdateStatusConf, kqMq.NewPaymentUpdateStatusMq(ctx, svcContext)),
	}

}
