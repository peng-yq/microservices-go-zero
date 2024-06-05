# microservices-go-zero

`microservices-go-zero`基于`go-zero`的民宿微服务系统，包括用户中心服务、民宿服务、订单服务和支付服务。

## Structure

**系统架构**

<img src="https://cdn.jsdelivr.net/gh/peng-yq/Gallery/202406051229362.png">

**业务架构**

<img src="https://cdn.jsdelivr.net/gh/peng-yq/Gallery/202406051201850.png">

## Feature

1. 使用`go-zero`框架进行开发，稳定、高性能。 

2. 使用`Nginx`作为对外网关，前面可接入`SLB`负载均衡器，并配合`go-zero-jwt`进行认证。

3. 采用微服务开发模式，结合`API (HTTP)`和`RPC (gRPC)`。`API`充当聚合服务，将复杂、涉及其他业务调用的逻辑统一放在`RPC`中。项目使用直链方式，放弃服务注册发现中间件带来的麻烦。

4. 使用`go-queue`实现订单支付状态更新消息队列，`asynq`实现支付成功通知用户消息队列和关闭超时未支付订单延时队列。

5. 使用`Filebeat`统一收集系统日志（各组件和业务），采集和转发日志数据到`Kafka`。使用`Go-Stash`替代`Logstash`进行日志过滤，高性能且资源消耗较低。使用`Prometheus`进行系统性能监控。使用`Jaeger`进行链路追踪，帮助微服务架构的监控和故障排查。`Elasticsearch`存储系统日志数据、监控数据和链路追踪数据，并进行数据分析。

6. 数据库使用`mysql`，并使用`redis`作为缓存、消息队列和延时队列。

## How to use (deploy)

1. Git Clone this repo

```shell
git clone https://github.com/peng-yq/microservices-go-zero.git
```

2. Follow the [tutorial](./docs/deploy-project.md)

## Environment

<img src="https://cdn.jsdelivr.net/gh/peng-yq/Gallery/202406051255703.png">

## Doc

1. [deploy-project](./docs/deploy-project.md)
2. [docker-compose](./docs/docker-compose.md)
3. [modd](./docs/modd.md)
4. [how to create microservices](./docs/how-to-create-microservices.md)
5. [message queue and delay queue](./docs/message-queue-delay-queue-timed-queue.md)

## To-do

- [ ] 性能测试
- [ ] 完善业务逻辑
- [ ] 增加单元测试
- [ ] 完善安全措施，例如`RBAC`

## Thanks

- [go-zero](https://github.com/zeromicro/go-zero)
- [go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook)