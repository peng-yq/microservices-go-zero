package main

import (
	"flag"
	
	"microservices-go-zero/app/order/cmd/mq/internal/config"
	"microservices-go-zero/app/order/cmd/mq/internal/listen"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/order.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// init log、prometheus、trace、metricsUrl.
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}
	// Start is a blocking function
	serviceGroup.Start()
}
