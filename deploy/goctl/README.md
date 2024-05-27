`goctl`是`go-zero`配套的代码生成工具脚手架，其集成了`Go HTTP`服务，`Go gRPC`服务，数据库`model`，`k8s`，`Dockerfile`等生成功能。

如果go版本在1.16以前，则使用如下命令安装：

```shell
GO111MODULE=on go get -u github.com/zeromicro/go-zero/tools/goctl@latest
```

如果go版本在1.16及以后，则使用如下命令安装：

```shell
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

验证安装：

```shell
goctl --version
```

> 如果安装完成却显示不出版本，则需要检查是否将`$GOPATH/bin`加入至环境变量中。

[官方文档](https://go-zero.dev/docs/tutorials/cli/overview)

<img src="https://go-zero.dev/assets/images/goctl-6fcfe4bce2787b1122816329e94db82a.png">