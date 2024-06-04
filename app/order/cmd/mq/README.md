## 使用go-zero的ServiceGroup进行消息队列服务管理

### servicegroup

假如项目中需要使用多个消息队列（延时队列、定时队列），我们就需要对其进行服务管理，在`go-zero`里，提供了一个`ServiceGroup`工具，方便管理多个服务的启动和停止。

[进程内优雅管理多个服务-go-zero](https://talkgo.org/t/topic/3720)

`servicegroup`的源码短小精悍，并且`go-zero`框架本身对每个服务（`Restful`, `RPC`, `MQ`）也基本都是通过`ServiceGroup`来管理。要使用`servicegroup`只需实现`start`和`stop`两个方法，就可以加入到`serviceGroup`统一管理。

`kq`已经实现了`start`和`stop`接口，这里用于实现第三方支付订单支付状态更新的消息队列，此处为消费者的实现，从`kafka`特定`topic`中取出第三方支付消息进行消费，并修改订单支付状态。

```go
func (q *kafkaQueue) Start() {
	q.startConsumers()
	q.startProducers()

	q.producerRoutines.Wait()
	close(q.channel)
	q.consumerRoutines.Wait()
}

func (q *kafkaQueue) Stop() {
	q.consumer.Close()
	logx.Close()
}
```

```go
package service

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/core/threading"
)

type (
	// Starter is the interface wraps the Start method.
	Starter interface {
		Start()
	}

	// Stopper is the interface wraps the Stop method.
	Stopper interface {
		Stop()
	}

	// Service is the interface that groups Start and Stop methods.
	Service interface {
		Starter
		Stopper
	}

	// A ServiceGroup is a group of services.
	// Attention: the starting order of the added services is not guaranteed.
	ServiceGroup struct {
		services []Service
		stopOnce func()
	}
)

// NewServiceGroup returns a ServiceGroup.
func NewServiceGroup() *ServiceGroup {
	sg := new(ServiceGroup)
	sg.stopOnce = syncx.Once(sg.doStop)
	return sg
}

// Add adds service into sg.
func (sg *ServiceGroup) Add(service Service) {
	// push front, stop with reverse order.
	sg.services = append([]Service{service}, sg.services...)
}

// Start starts the ServiceGroup.
// There should not be any logic code after calling this method, because this method is a blocking one.
// Also, quitting this method will close the logx output.
func (sg *ServiceGroup) Start() {
	proc.AddShutdownListener(func() {
		logx.Info("Shutting down services in group")
		sg.stopOnce()
	})

	sg.doStart()
}

// Stop stops the ServiceGroup.
func (sg *ServiceGroup) Stop() {
	sg.stopOnce()
}

func (sg *ServiceGroup) doStart() {
	routineGroup := threading.NewRoutineGroup()

	for i := range sg.services {
		service := sg.services[i]
		routineGroup.Run(func() {
			service.Start()
		})
	}

	routineGroup.Wait()
}

func (sg *ServiceGroup) doStop() {
	for _, service := range sg.services {
		service.Stop()
	}
}

// WithStart wraps a start func as a Service.
func WithStart(start func()) Service {
	return startOnlyService{
		start: start,
	}
}

// WithStarter wraps a Starter as a Service.
func WithStarter(start Starter) Service {
	return starterOnlyService{
		Starter: start,
	}
}

type (
	stopper struct{}

	startOnlyService struct {
		start func()
		stopper
	}

	starterOnlyService struct {
		Starter
		stopper
	}
)

func (s stopper) Stop() {
}

func (s startOnlyService) Start() {
	s.start()
}
```
`servicegroup`对不同的服务进行管理和启动本质也是通过`waitGroup`实现。

```go
func (g *RoutineGroup) Run(fn func()) {
	g.waitGroup.Add(1)

	go func() {
		defer g.waitGroup.Done()
		fn()
	}()
}
```

函数在一个独立的`Goroutine`中运行，而`waitGroup`则可以用来等待`Goroutine`执行完毕。

### 代码结构

`go-zero`并未像`api`和`rpc`一样提供对`servicegroup`的代码生成，项目通过模仿`api/rpc`目录结构进行改造（参考的官方用例）。

- `listen`目录用于初始化和监听不同的消费队列消费者
- `mqs`目录用于管理不同的消息队列服务逻辑
- `main.go`用于将消费队列加入`servicegroup`，并启动消费者进行消费

`main.go`代码比较简单，包括初始化配置和服务，将消费队列加入`servicegroup`，并启动消费者进行消费。区别于`api/rpc`生成的`main`代码中使用`mustnewserver`方法来初始化配置，这里通过`setup`来初始。

```go
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
```