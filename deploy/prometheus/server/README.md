## prometheus的配置文件

`Prometheus`是最初在`SoundCloud`上构建的开源系统监控和警报工具，在`2016`年加入了`Cloud Native Computing Foundation（CNCF）`基金会，是继`Kubernetes`之后该基金会的第二个托管项目。项目通过使用`Promethues`来监控压力测试时服务端的性能。

我们在`docker-compose-env.yaml`文件中的`prometheus`部分将该配置文件挂载至`prometheus`容器的`/etc/prometheus/prometheus.yml`路径进行应用，并且将容器中`/prometheus`路径中的内容持久化到`/data/prometheus/data`。

```yaml
global:
  scrape_interval: 15s  # Default scrape interval
  external_labels:
    monitor: 'codelab-monitor'
```

`scrape_interval`: 这是`Prometheus`默认的抓取（或“采集”）间隔，意味着`Prometheus`将每`15`秒从配置的目标中抓取一次指标数据。如果`scrape_configs`中的具体任务没有指定自己的抓取间隔，就会使用这个全局间隔。
`external_labels`: 这些是添加到所有从此`Prometheus`实例抓取的指标上的额外标签。所有的指标都会被标记为来自'`codelab-monitor`'的监控器。

```yaml
  # Order API service
  - job_name: 'order-api'
    static_configs:
      - targets: ['microservices:4001']
        labels:
          job: order-api
          app: order-api
          env: dev
```

- `job_name`: 任务名称，如`order-api`，用于标识这个任务是用来监控哪个服务的。
- `targets`: 指定`Prometheus`需要从哪些地址抓取数据。
- `labels`: 为从这个任务抓取的指标添加额外的标签，有助于在`Prometheus`中更好地组织和查询数据。标签包括服务名称 (`app`), 任务名称 (`job`), 和环境 (`env`)。

本项目中主要监控的服务包括：

- `prometheus`：`127.0.0.1:9090`
- `order-api`：`microservices:4001`
- `order-rpc`：`microservices:4002`
- `order-mq`：`microservices:4003`
- `payment-api`：`microservices:4004`
- `payment-rpc`：`microservices:4005`
- `travel-api`：`microservices:4006`
- `travel-rpc`：`microservices:4007`
- `usercenter-api`：`microservices:4008`
- `usercenter-rpc`：`microservices:4009`
- `mqueue-job`：`microservices:4010`
- `mqueue-scheduler`：`microservices:4011`