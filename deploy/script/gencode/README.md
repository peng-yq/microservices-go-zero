## 生成业务代码的命令

`goctl api`是`goctl`中的核心模块之一，其可以通过`.api`文件一键快速生成一个`api`服务。使用`goctl`根据`api`文件生成`GO HTTP`代码，需要进入每个服务下的`/cmd/api/desc`目录，并执行以下命令：

```shell
goctl api go --api *.api --dir ../  --style=goZero
```

```shell
goctl api go --help
Generate go files for provided api in api file

Usage:
  goctl api go [flags]

Flags:
      --api string      The api file
      --branch string   The branch of the remote repo, it does work with --remote
      --dir string      The target dir
  -h, --help            help for go
      --home string     The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --remote string   The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                        The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
      --style string    The file naming format, see (default "gozero")
```

## 生产简易api接口文档

使用`goctl`根据`api`文件生成生成`markdown`文档，需要进入每个服务下的`/cmd/api/desc`目录，并执行以下命令：

```shell
goctl api doc --dir .
```

```shell
$ goctl api doc --help
Generate doc files

Usage:
  goctl api doc [flags]

Flags:
      --dir string   The target dir
  -h, --help         help for doc
      --o string     The output markdown directory
```

## 生成rpc业务代码

1. 安装`protoc (>= 3.13.0)`

下载链接[https://github.com/protocolbuffers/protobuf]。

2. 安装`protoc-gen-go`

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

3. 安装`protoc-gen-go-grpc`

```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

4. 如果有要使用`grpc-gateway`，需要额外安装如下两个插件

```shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

`goctl rpc`是`goctl`中的核心模块之一，其可以通过`.proto`文件一键快速生成一个`rpc`服务。用`goctl`根据`protobufer`文件生成`rpc`代码，需要进入每个服务下的`/cmd/rpc/pb`目录，并执行以下命令：

```shell
goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero
``` 

## 去除proto中的json的omitempty

进入每个服务下的`/cmd/rpc/pb`目录，去除`proto`中的`json`的`omitempty`，从而显示所有字段（包括为`0`，`""`或`nil`）。

```shell
sed -i 's/,omitempty//g' *.pb.go
```

## kafka

创建`kafka`的`topic`：

```shell
kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 -partitions 1 --topic {topic}
```

查看消费者组情况：

```shell
kafka-consumer-groups.sh --bootstrap-server kafka:9092 --describe --group {group}
```

命令行消费：

```shell
./kafka-console-consumer.sh  --bootstrap-server kafka:9092  --topic microservices-log   --from-beginning
```

命令生产：

```shell
./kafka-console-producer.sh --bootstrap-server kafka:9092 --topic second
```