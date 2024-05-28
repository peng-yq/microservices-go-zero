### 使用`docker compose`创建项目所依赖环境

```shell
docker-compose -f docker-compose-env.yaml up -d
```

> `docker-compose`创建完成后，使用`docker ps`查看是否所有容器都处于`up`状态。如果有一直`restarting`的容器，需要使用`docker logs 容器名`进行排查并重启。

### 创建`kafka topic`

主题（`Topics`）是消息的分类或者说是消息的目的地。每个消息发布到一个特定的主题中，而消费者可以订阅一个或多个主题来接收消息。创建两个`Kafka`主题：`microservices-log`和`payment-update-paystatus-topic`，分别用于日志收集和支付成功后通知所有订阅者。

```shell
docker exec -it kafka /bin/sh
cd /opt/kafka/bin/
./kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 -partitions 1 --topic microservices-log
./kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 -partitions 1 --topic payment-update-paystatus-topic
```

- `--create`: 指示脚本创建一个新主题。
- `--zookeeper`: 指定`Zookeeper`的地址和端口，`Kafka`使用`Zookeeper`来维护集群状态。
- `--replication-factor`: 设置复制因子，为`1`意味着每个消息在`Kafka集`群中只有一个副本。
- `--partitions`: 设置分区数为，分区是`Kafka`的横向扩展机制，但在这里只需要一个分区。
- `--topic`: 指定主题名称。

### 导入`MySql`数据

为`root`设置远程连接权限，以便使用`navicat`操作。

```shell
docker exec -it mysql mysql -uroot -p
##输入密码：PXDN93VRKUm8TeE7
use mysql;
update user set host='%' where user='root';
FLUSH PRIVILEGES;
```

使用`navicat`创建数据库并导入数据（创建数据库后，进入数据库并运行对应`sql`文件）：

- 创建数据库`looklook_order`，导入`deploy/sql/looklook_order.sql`数据。
- 创建数据库`looklook_payment`，导入`deploy/sql/looklook_payment.sql`数据。
- 创建数据库`looklook_travel`，导入`deploy/sql/looklook_travel.sql`数据。
- 创建数据库`looklook_usercenter`，导入`looklook_usercenter.sql`数据。

> 数据库字符集选择`utf8mb4`；排序规则选择`utf8mb4_general_ci`。       

### 启动项目

```shell
docker-compose up -d
```

> `docker-compose`创建完成后，使用`docker ps`查看是否所有容器都处于`up`状态。如果有一直`restarting`的容器，需要使用`docker logs 容器名`进行排查并重启。

<img src="https://cdn.jsdelivr.net/gh/peng-yq/Gallery/202405231303989.png">

### 查看项目运行情况

访问`http://127.0.0.1:9090/`，点击菜单`Status/Targets`，蓝色表示启动成功，红色表示启动失败。

> 项目第一次构建会拉取依赖，所有服务启动成功可能需要一段时间。

<img src="https://cdn.jsdelivr.net/gh/peng-yq/Gallery/202405231245174.png">

| 服务名称      | URL 或 Host                | Port  | 用户名 | 密码              | 备注                                           |
|--------------|----------------------------|-------|--------|-------------------|------------------------------------------------|
| ElasticSearch| `http://127.0.0.1:9200/`   | N/A   | N/A    | N/A               | 启动时间可能较长                                   |
| Jaeger       | `http://127.0.0.1:16686/search` | N/A   | N/A    | N/A          | 依赖ElasticSearch，可能需要重启               |
| Go-Stash     | N/A                        | N/A   | N/A    | N/A               | 如果日志未收集到，重启Go-Stash                 |
| Asynq        | `http://127.0.0.1:8980/`   | N/A   | N/A    | N/A               | 延迟任务、定时任务、消息队列                    |
| Kibana       | `http://127.0.0.1:5601/`   | N/A   | N/A    | N/A               | N/A                                            |
| Prometheus   | `http://127.0.0.1:9090/`   | N/A   | N/A    | N/A               | N/A                                            |
| Grafana      | `http://127.0.0.1:3001/`   | N/A   | admin  | admin             | N/A                                            |
| MySQL        | `127.0.0.1`                | 33069 | root   | PXDN93VRKUm8TeE7  | 使用客户端工具如Navicat查看                   |
| Redis        | `127.0.0.1`                | 36379 | N/A    | G62m50oigInC30sf  | 使用工具如RedisManager查看                    |
| Kafka        | `127.0.0.1`                | 9092  | N/A    | N/A               | 使用客户端工具查看pub/sub                      |
| Nginx        | `http://127.0.0.1:8888/`   | N/A   | N/A    | N/A               | 用于访问API，如用户注册：`/usercenter/v1/user/register` |

### 访问项目

项目使用`nginx`作为网关，`nginx`对外暴露端口为`8888`，通过`8888`端口可访问`api`（`api`内部通信为`rpc`）提供服务：

```shell
curl  -X POST "http://127.0.0.1:8888/usercenter/v1/user/register" -H "Content-Type: application/json" -d "{\"mobile\":\"18888888888\",\"password\":\"123456\"}"
```

服务访问成功将返回`code:200`，同时在`looklook_usercenter`数据库中能存在注册的用户条目。

### 日志收集

项目日志收集流程如下：

```
filebeat收集日志 -> kafka -> go-stash消费kafka日志 -> 输出到es中 -> kibana查看es数据
```

收集日志，需要创建日志索引：

1. 访问`kibana:http://127.0.0.1:5601/`， 创建日志索引
2. 点击左上角菜单，选择`Analytics/discover` 
3. 选择`Create index pattern`，输入`looklook-*`，点击`Next Step`，选择`@timestamp->Create index pattern`
4. 点击左上角菜单，选择`Analytics/discover`，日志显示

<img src="https://cdn.jsdelivr.net/gh/peng-yq/Gallery/202405231439146.png">