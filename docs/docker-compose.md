## **启动项目所依赖的环境**

项目所有组件如监控、日志、消息队列、数据库以及核心业务逻辑都通过`docker`部署，并通过`docker network`进行通信和隔离。

### **docker network**

> 直接通过`yaml`文件启动时会自动创建目标`docker network`。

使用`docker network create`命令在`Docker`中创建一个网络，默认将创建`bridge`类型的网络。此外还可以添加显式添加`--driver`选项支持`host`、`container`等其他网络类型。创建网络后，可以使用以下命令列出所有`Docker`网络，以验证网络是否已成功创建：

```bash
docker network ls
```

创建网络后，你可以在运行容器时通过`--network`选项将容器连接到这个网络。例如：

```bash
docker run --name some-app --network microservices_net -d some-image
```

创建`Docker`网络主要用于容器间的隔离和通信。`Docker`自身提供了简单的内置`DNS`服务，用于容器名到`IP`地址的解析，这使得可以更容易地通过容器名称进行互联，而不是通过`IP`地址，

> `Docker`容器每次重启后容器`ip`是会发生变化的。这也意味着如果容器间使用`ip`地址来进行通信的话，一旦有容器重启，重启的容器将不再能被访问到。

[更多关于docker网络介绍](https://peng-yq.github.io/2021/12/10/docker/)

### **docker-compose-env.yaml**

使用`docker-compose`部署项目依赖环境。

服务组件：

- `Jaeger` - 用于链路追踪，帮助监控和故障排查微服务架构的性能问题。
- `Prometheus` - 监控工具，收集和存储指标数据，用于系统的性能监控。
- `Grafana` - 与`Prometheus`配合使用，提供数据可视化界面，用于展示监控数据。
- `Elasticsearch` - 高效的搜索和数据分析引擎，这里用于存储日志和监控数据。
- `Kibana` - `Elasticsearch`的数据可视化前端工具，用于查看和分析在`Elasticsearch`中的数据。
- `Go-Stash` - 日志处理服务，用于处理和转发日志数据到`Elasticsearch`。
- `Filebeat` - 轻量级日志收集器，用于采集和转发日志数据到`Kafka`。
- `Zookeeper` - 分布式服务的协调工具，`Kafka`的依赖服务。
- `Kafka` - 高吞吐量的分布式消息队列，用于日志和事件的传输。
- `Asynqmon` - 提供`Web UI`管理`Asynq`任务队列和监控。
- `MySQL` - 关系型数据库，用于存储应用数据。
- `Redis` - 内存中的`kv`键值数据库，用作数据库、缓存和消息代理。

网络配置：

- `microservices_net`：所有服务都连接到这个自定义`bridge`网络，允许容器间通信和隔离外部网络。网络配置指定了子网为`172.16.0.0/16`，确保容器间网络通信不受外界干扰。

> `Docker`将为`172.16.0.0/16`网络创建一个新的虚拟桥接，并负责为连接到同一网络的容器自动分配IP地址。虽然容器内部使用的是隔离的子网，但它们可以通过`NAT`（网络地址转换）访问外部网络。宿主机可以通过端口映射访问运行在容器内的服务，例如，将容器的80端口映射到宿主机的8080端口，从而从宿主机网络访问容器服务。

其他配置：

- `Volumes`：多个服务使用卷（`volumes`）来持久化和共享数据。例如，`Prometheus`、`Grafana`、`Elasticsearch`等服务的数据都被保存在本地文件系统中，以便在容器重启后数据不会丢失。
- `Ports`：服务如`Prometheus`、`Grafana`、`Kibana`等对外暴露端口，使得可以从宿主机器访问这些服务。
- `Environment`：设置环境变量，如时区设置为上海（`Asia/Shanghai`），以及其他服务特定的配置。

> `Docker Compose`默认会给每个网络前缀一个项目名称，通常是目录名。项目目录名是`microservices-go-zero`，那么实际的网络名自动变成了`microservices-go-zero_microservices_net`。

`docker-compose-env.yaml`

> 如果遇到`es`启动失败，需要修改挂载的卷的权限`chmod 777 data/elasticsearch/data`，`go-stash`和`jaeger`依赖于`es`，将会自动重启成功 。

```yaml
services:
  # Jaeger for tracing
  jaeger:
    image: jaegertracing/all-in-one:1.42.0
    container_name: jaeger
    restart: always
    ports:
      - "5775:5775/udp"
      - "6831:6831/udp"
      - "6832:6832/udp"
      - "5778:5778"
      - "16686:16686"
      - "14268:14268"
      - "9411:9411"
    environment:
      - SPAN_STORAGE_TYPE=elasticsearch
      - ES_SERVER_URLS=http://elasticsearch:9200
      - LOG_LEVEL=debug
    networks:
      - microservices_net

  # Prometheus for monitoring
  prometheus:
    image: prom/prometheus:v2.28.1
    container_name: prometheus
    environment:
      # Time zone Shanghai (Change if needed)
      TZ: Asia/Shanghai
    volumes:
      - ./deploy/prometheus/server/prometheus.yml:/etc/prometheus/prometheus.yml
      - ./data/prometheus/data:/prometheus
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
    restart: always
    user: root
    ports:
      - 9090:9090
    networks:
      - microservices_net

  # Grafana to view Prometheus monitoring data
  grafana:
    image: grafana/grafana:8.0.6
    container_name: grafana
    hostname: grafana
    user: root
    environment:
      # Time zone Shanghai (Change if needed)
      TZ: Asia/Shanghai
    restart: always
    volumes:
        - ./data/grafana/data:/var/lib/grafana
    ports:
        - "3001:3000"
    networks:
        - microservices_net

  # Kafka for collecting business logs and storing Prometheus monitoring data
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:7.13.4
    container_name: elasticsearch
    user: root
    environment:
      - discovery.type=single-node
      - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
      - TZ=Asia/Shanghai
    volumes:
      - ./data/elasticsearch/data:/usr/share/elasticsearch/data
    restart: always
    ports:
    - 9200:9200
    - 9300:9300
    networks:
      - microservices_net

  # Kibana to view Elasticsearch data
  kibana:
    image: docker.elastic.co/kibana/kibana:7.13.4
    container_name: kibana
    environment:
      - elasticsearch.hosts=http://elasticsearch:9200
      - TZ=Asia/Shanghai
    restart: always
    networks:
      - microservices_net
    ports:
      - "5601:5601"
    depends_on:
      - elasticsearch

  # The data output collected by FileBeat in Kafka is output to ES
  go-stash:
    image: kevinwan/go-stash:1.0
    container_name: go-stash
    environment:
      # Time zone Shanghai (Change if needed)
      TZ: Asia/Shanghai
    user: root
    restart: always
    volumes:
      - ./deploy/go-stash/etc:/app/etc
    networks:
      - microservices_net
    depends_on:
      - elasticsearch
      - kafka

  # Collect business data
  filebeat:
    image: elastic/filebeat:7.13.4
    container_name: filebeat
    environment:
      # Time zone Shanghai (Change if needed)
      TZ: Asia/Shanghai
    user: root
    restart: always
    entrypoint: "filebeat -e -strict.perms=false"  # Solving the configuration file permissions
    volumes:
      - ./deploy/filebeat/conf/filebeat.yml:/usr/share/filebeat/filebeat.yml
      - /var/lib/docker/containers:/var/lib/docker/containers
    networks:
      - microservices_net
    depends_on:
      - kafka


  # Zookeeper is the dependencies of Kafka
  zookeeper:
    image: wurstmeister/zookeeper
    container_name: zookeeper
    environment:
      # Time zone Shanghai (Change if needed)
      TZ: Asia/Shanghai
    restart: always
    ports:
      - 2181:2181
    networks:
      - microservices_net

  # Message queue
  kafka:
    image: wurstmeister/kafka
    container_name: kafka
    ports:
      - 9092:9092
    environment:
      - KAFKA_ADVERTISED_HOST_NAME=kafka
      - KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181
      - KAFKA_AUTO_CREATE_TOPICS_ENABLE=false
      - TZ=Asia/Shanghai
    restart: always
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    networks:
      - microservices_net
    depends_on:
      - zookeeper

  # Asynqmon asynq delay queue, timing queue's webUI
  asynqmon:
    image: hibiken/asynqmon:latest
    container_name: asynqmon
    ports:
      - 8980:8080
    command:
      - '--redis-addr=redis:6379'
      - '--redis-password=G62m50oigInC30sf'
    restart: always
    networks:
      - microservices_net
    depends_on:
      - redis

  mysql:
    image: mysql/mysql-server:8.0.28
    container_name: mysql
    environment:
      # Time zone Shanghai (Change if needed)
      TZ: Asia/Shanghai
      # root password
      MYSQL_ROOT_PASSWORD: PXDN93VRKUm8TeE7
    ports:
      - 33069:3306
    volumes:
      # Data mounting
      - ./data/mysql/data:/var/lib/mysql
    command:
      # Modify the Mysql 8.0 default password strategy to the original strategy (MySQL8.0 to change its default strategy will cause the password to be unable to match)
      --default-authentication-plugin=mysql_native_password
      --character-set-server=utf8mb4
      --collation-server=utf8mb4_general_ci
      --explicit_defaults_for_timestamp=true
      --lower_case_table_names=1
    restart: always
    networks:
      - microservices_net

  # Redis container
  redis:
    image: redis:6.2.5
    container_name: redis
    ports:
      - 36379:6379
    environment:
      # Time zone Shanghai (Change if needed)
      TZ: Asia/Shanghai
    volumes:
      - ./data/redis/data:/data:rw
    command: "redis-server --requirepass G62m50oigInC30sf  --appendonly yes"
    privileged: true
    restart: always
    networks:
      - microservices_net

networks:
  microservices_net:
    driver: bridge
    ipam:
      config:
        - subnet: 172.16.0.0/16
```

### **docker-compose.yaml**

- `nginx`：创建一个前端网关，主要负责代理`microservices`服务，但不代理`admin-api`服务。
- `gomodd`：提供`golang`环境运行具体代码逻辑以及`modd`的热加载功能，修改代码即时生效，无需重启容器服务。

```yaml
services:
  # Front-end gateway nginx-gateway (Only agent microservices，admin-api Do not be an agent here)
  nginx-gateway:
    image: nginx:1.21.5
    container_name: nginx-gateway
    restart: always
    privileged: true
    environment:
      - TZ=Asia/Shanghai
    ports:
      - 8888:8081
    volumes:
      - ./deploy/nginx/conf.d:/etc/nginx/conf.d
      - ./data/nginx/log:/var/log/nginx
    networks:
      - microservices_net
    depends_on:
      - microservices

  # Front-end API + business RPC
  microservices:
    # docker-hub : https://hub.docker.com/r/lyumikael/gomodd
    # dockerfile: https://github.com/Mikaelemmmm/gomodd , If you are macOs m1\m2 use dockerfile yourself to build the image
    image: lyumikael/gomodd:v1.20.3
    container_name: microservices
    environment:
      TZ: Asia/Shanghai
      GOPROXY: https://goproxy.cn,direct
    working_dir: /go/microservices
    volumes:
      - .:/go/microservices
    privileged: true
    restart: always
    networks:
      - microservices_net

networks:
  microservices_net:
    driver: bridge
    ipam:
      config:
        - subnet: 172.20.0.0/16
```

