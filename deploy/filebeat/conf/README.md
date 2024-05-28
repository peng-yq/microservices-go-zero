## `Filebeat`的配置文件

`Filebeat`是`Elastic Stack`的一部分，主要用于轻量级日志采集器。它可以将日志或文件数据发送到`Elasticsearch`或`Logstash`，然后再进行分析和存储。本项目中，`Filebeat`从`Docker`容器中收集日志，并将这些日志发送到`Kafka`。

我们在`docker-compose-env.yaml`文件中的`filebeat`部分将该配置文件挂载至`filebeat`容器的`/usr/share/filebeat/filebeat.yml`路径进行应用，并且将主机的`/var/lib/docker/containers`路径挂载到`filebeat`容器的相同路径，使其能读取所有主机所有容器的日志，并发送至`kafka`中。

### 详细解释


```yaml
filebeat.inputs:
  - type: log
    enabled: true
    paths:
      - /var/lib/docker/containers/*/*-json.log
```

- `type: log`：指定输入类型为日志文件。
- `enabled: true`：启用这个输入配置。
- `paths`：指定`Filebeat`需要监控的文件或目录的路径。这里配置的路径是`/var/lib/docker/containers/*/*-json.log`，它会匹配所有`Docker`容器的日志文件，通常这些文件以`容器的ID-json.log`命名（包含容器的`stdout`和`stderr`）。

```yaml
filebeat.config:
  modules:
    path: ${path.config}/modules.d/*.yml
    reload.enabled: false
```

- `path: ${path.config}/modules.d/*.yml`：指定`Filebeat`模块配置文件的路径。
- `reload.enabled: false`：禁用自动重载配置文件，意味着`Filebeat`在运行期间不会自动检测配置文件的变化。

```yaml
processors:
  - add_cloud_metadata: ~
  - add_docker_metadata: ~
```

- `add_cloud_metadata`：自动添加云提供商的元数据到收集的日志中，例如`AWS、Azure、GCP`的实例`ID`和区域信息。
- `add_docker_metadata`：自动添加`Docker`容器相关的元数据，如容器`ID`、名称、镜像信息等。

```yaml
output.kafka:
  enabled: true
  hosts: ["kafka:9092"]
  topic: "microservices-log"
  partition.hash:
    reachable_only: true
  compression: gzip
  max_message_bytes: 1000000
  required_acks: 1
```

- `enabled: true`：启用`Kafka`输出。
- `hosts: ["kafka:9092"]`：`Kafka`服务器的地址和端口。
- `topic: "microservices-log"`：指定`Kafka`主题名称，`Filebeat`将日志数据发送到这个主题。
- `partition.hash.reachable_only: true`：确保只将消息发送到可达的分区。
- `compression: gzip`：启用`gzip`压缩，减少网络传输数据量。
- `max_message_bytes: 1000000`：设置`Kafka`消息的最大字节数。
- `required_acks: 1`：设置`Kafka`生产者需要从服务器接收的确认数量，`1`表示至少要从一个服务器得到消息接收的确认。

