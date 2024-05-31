package main

import (
	"flag"
	"fmt"

	"microservices-go-zero/app/travel/cmd/rpc/internal/config"
	"microservices-go-zero/app/travel/cmd/rpc/internal/server"
	"microservices-go-zero/app/travel/cmd/rpc/internal/svc"
	"microservices-go-zero/app/travel/cmd/rpc/pb"
	"microservices-go-zero/common/interceptor/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/travel.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterTravelServer(grpcServer, server.NewTravelServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
