## `go-stash`的配置文件

`go-stash`属于`go-zero`生态的一个组件，是一个高效的从`Kafka`获取，根据配置的规则进行处理，然后发送到`ElasticSearch`集群的工具。 `go-stash`有大概`logstash`五倍的吞吐性能，并且部署简单，一个可执行文件即可。本项目中，`go-stash`中`kafka`中收集日志，进行处理过滤后发送至`ElasticSearch`集群中。

我们在`docker-compose-env.yaml`文件中的`go-stash`部分将该配置文件挂载至`go-stash`容器的`/app/etc/config.yml`路径进行应用

```yaml
- Input:
      Kafka:
        Name: gostash
        Brokers:
          - "kafka:9092"
        Topics:
          - microservices-log
        Group: pro
        Consumers: 16
```

设置链接`kafka`的连接，包括地址和主题。这里只设置一个连接（一般来说连接数 <= `cpu`核心数），并设置单个连接中消费者线程数为`16`（总的消费者线程数应该 <= `kafka topic`分片数）。

```yaml
- Action: drop
        Conditions:
          - Key: k8s_container_name
            Value: "-rpc"
            Type: contains
          - Key: level
            Value: info
            Type: match
            Op: and
```

满足特定条件的记录将被丢弃，下面两个条件必须同时满足，记录才会被丢弃：
- 第一个条件检查`k8s_container_name`字段是否包含`-rpc`。
- 第二个条件检查`level`字段是否匹配`info`。

```yaml
- Action: remove_field
        Fields:
          # - message
          - _source
          - _type
          - _score
          - _id
          - "@version"
          - topic
          - index
          - beat
          - docker_container
          - offset
          - prospector
          - source
          - stream
          - "@metadata"
```

从记录中移除`Fields`指定的字段，一些不需要存储到`ElasticSearch`中的元数据或冗余信息。

```yaml
- Action: transfer
        Field: message
        Target: data
```

将`message`字段，重新定义为`data`字段。

```yaml
    Output:
      ElasticSearch:
        Hosts:
          - "http://elasticsearch:9200"
        Index: "microservices-{{yyyy-MM-dd}}"
```

设置将收集过滤的日志发送到`ElasticSearch`集群的信息，包括地址和索引。

