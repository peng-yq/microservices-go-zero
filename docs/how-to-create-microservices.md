## 使用go-zero构建微服务

### 编写api文件，生成基本的api代码

1. [语法](https://go-zero.dev/docs/tutorials)
2. [goctl api生成代码](https://go-zero.dev/docs/tutorials/cli/api)
3. `api`文件只需写明具体的路由、`HTTP`方法、请求体和响应体

### 编写proto文件，生成基本的rpc代码

1. [语法](https://protobuf.dev/getting-started/gotutorial/)
2. [goctl rpc生成代码](https://go-zero.dev/docs/tutorials/cli/rpc)
3. `proto`需要写明消息和`rpc`调用方法

### 根据sql文件或数据库连接，生成数据库model代码

可通过两种方式生成`model`代码，均通过[`goctl model`](https://go-zero.dev/docs/tutorials/cli/model)生成：

1. `sql`文件
2. 数据库连接

生成的`model`代码提供了基础的针对主键和唯一键的数据库（`sql`）操作，更多复杂的`crud`需要在对应的`*_gen.go`中编写并在接口中注册。

此外，`go-zero`的数据库事务操作只能在对应的`model`中进行调用，但我们可以对其进行暴露，使得可以在编写具体的`HTTP/RPC`逻辑时调用。

示例：

```go
type userModel interface {
    // TransactCtx can only be applied locally, so it is encapsulated, registered and exposed
    Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
}

func (m *defaultUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

// 调用示例
UserModel.Trans(ctx, func(ctx context.Context, session sqlx.Session) error {
    // 具体的逻辑
    return nil
}
```

### 补充rpc服务逻辑

1. 根据需要修改`rpc`代码的配置文件`rpc/etc/*.yaml`，例如修改监听的地址和端口，或者增加数据库、`cache`、`log`、`redis`、监控和链路追踪等配置。
2. 在`rpc/internal/config/config.go`中为`config`结构体增加对应的配置字段，`rpc`服务在启动后会将`yaml`文件中的内容解析至`config`中。
3. 在`rpc/internal/svc/serviceContext.go`中为`ServiceContext`结构体增加对应的配置字段，例如`redis`和数据库实例，并在`NewServiceContext`中完成资源和组件的初始化。
4. 在`rpc/internal/logic`中编写具体调用方法的逻辑。

`rpc`服务主程序示例：

```go
var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServer(grpcServer, server.NewUsercenterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
```

主要流程：

1. 解析配置文件到`config.Config`中
2. 根据`config.Config`初始化组件和资源
3. 创建新的`gRPC`服务器实例，服务器会从配置结构`config.Config.RpcServerConf`中读取具体的服务器配置
4. 将具体的服务逻辑注册到`gRPC`服务器中，使得服务器能够知道如何处理特定的`rpc`调用

### 补充HTTP服务逻辑

1. 根据需要修改`rpc`代码的配置文件`api/etc/*.yaml`，例如修改监听的地址和端口，或者增加`rpc`、`log`、监控和链路追踪等配置。
2. 在`api/internal/config/config.go`中为`config`结构体增加对应的配置字段，`api`服务在启动后会将`yaml`文件中的内容解析至`config`中。
3. 在`api/internal/svc/serviceContext.go`中为`ServiceContext`结构体增加对应的配置字段，例如`rpc`实例，并在`NewServiceContext`中完成资源和组件的初始化。
4. 在`api/internal/logic`中编写具体调用方法的逻辑。

`HTTP`服务主程序示例：

```go
var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
```

主要流程：

1. 解析配置文件到`config.Config`中
2. 创建新的`HTTP`服务器实例，服务器会从配置结构`config.Config.RestConf`中读取具体的服务器配置
3. 根据`config.Config`初始化组件和资源
4. 注册所有的路由和对应的`handler`到服务器
5. 服务器启动并监听`http`请求，当服务器收到具体的`http`请求时，调用具体的`handler`函数，`handler`函数将解析请求体，调用具体的逻辑方法，并构造请求体返回