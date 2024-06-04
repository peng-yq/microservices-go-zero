package listen

import (
	"context"
	"microservices-go-zero/app/order/cmd/mq/internal/config"
	"microservices-go-zero/app/order/cmd/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

// back to all consumers
func Mqs(c config.Config) []service.Service {
	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service
	services = append(services, KqMqConsumers(c, ctx, svcContext)...)

	return services
}
