### check-port.sh

`check-port.sh`脚本用于检测容器需要映射的宿主机端口是否被占用。

```shell
#!/bin/bash

# define the list of ports to be checked
ports=(5775 6831 6832 5778 16686 14268 9411 9090 3001 9200 9300 5601 2181 9092 8980 33069 36379)

for port in "${ports[@]}"; do
    # use "ss" command to check
    result=$(sudo ss -tuln | grep ":$port ")
    if [ -z "$result" ]; then
        echo "Port $port is available."
    else
        echo "Port $port is in use:"
        echo "$result"
    fi
done
```

### pull-image.sh

拉取容器镜像需要较大的带宽和一定时间，具体取决于镜像大小和网络速度。直接运行`docker compose`去启动环境并拉取镜像可能会十分漫长，可以使用`pull-image.sh`脚本预先拉取所需镜像。

```shell
#!/bin/bash

# define all the docker images we need
images=(
    "nginx:1.21.5"
    "lyumikael/gomodd:v1.20.3"
    "jaegertracing/all-in-one:1.42.0"
    "prom/prometheus:v2.28.1"
    "grafana/grafana:8.0.6"
    "docker.elastic.co/elasticsearch/elasticsearch:7.13.4"
    "docker.elastic.co/kibana/kibana:7.13.4"
    "kevinwan/go-stash:1.0"
    "elastic/filebeat:7.13.4"
    "wurstmeister/zookeeper"
    "wurstmeister/kafka"
    "hibiken/asynqmon:latest"
    "mysql/mysql-server:8.0.28"
    "redis:6.2.5"
)

for image in "${images[@]}"; do
    echo "Pulling $image..."
    docker pull $image
    echo "$image pulled successfully!"
done

echo "All images have been pulled successfully!"
```

在`Shell`脚本中，可以通过在命令后添加`&`符号来实现并行执行（速度会快很多）。这样，每个`docker pull`命令将在一个新的子进程中运行，允许多个镜像同时下载。需要注意的是脚本较为简单，没有处理可能发生的任何错误，并且多个进程可能同时输出到控制台，因此输出可能会交错在一起。

```shell
# ...
for image in "${images[@]}"; do
    echo "Pulling $image..."
    docker pull $image &
done
wait
# ...
```

如果依旧很慢，请修改镜像源。

```shell
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://yxzrazem.mirror.aliyuncs.com"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker
```

修改`dns server`为阿里云`dns`：`223.5.5.5`

```shell
sudo apt install resolvconf 
sudo systemctl enable --now resolvconf.service
# 编辑 /etc/resolvconf/resolv.conf.d/head 文件，并添加所需的DNS服务器223.5.5.5
# 使DNS配置生效
sudo resolvconf -u
# 验证DNS服务器是否已经修改
nslookup www.baidu.com
```

