## nginx的配置文件

项目使用`nginx`做为网关，用于将不同的请求路径转发到不同的微服务。

我们在`docker-compose.yaml`文件中的`nginx-gateway`部分将该配置文件挂载至`nginx-gateway`容器的`/etc/nginx/conf.d/microservices-gateway.conf`路径进行应用，并将`/var/log/nginx`路径下的日志持久化到`/data/nginx/log`目录，并将该容器`8081`端口映射至主机的`8888`端口。

```
listen 8081;
  access_log /var/log/nginx/microservices.com_access.log;
  error_log /var/log/nginx/microservices.com_error.log;
```

`Nginx`监听的端口号为`8081`（也就是主机的`8888`端口），所有发往这个端口的`HTTP`请求都将由这个服务器处理。

- `access_log，/var/log/nginx/microservices.com_access.log`：记录所有接入请求的详细信息。
- `error_log，/var/log/nginx/microservices.com_error.log`：记录所有错误信息。

每个`location`块定义了不同的请求路径，并指定了如何处理这些路径的请求。

设置传递给代理服务器（`nginx`服务）的请求头，这些头包括：
- `Host $http_host`; 传递原始请求的主机头。
- `X-Real-IP $remote_addr`; 传递请求者的`IP`地址。
- `REMOTE-HOST $remote_addr`; 同上，通常用于日志或其他用途。
- `X-Forwarded-For $proxy_add_x_forwarded_for`; 添加代理服务器的`IP`地址到`X-Forwarded-For`头，用于跟踪请求链。

指定请求应该被转发到的目标地址和端口：
- `/order/`：请求被转发到`http://microservices:1001`
- `/payment/`：请求被转发到`http://microservices:1002`
- `/travel/`：请求被转发到`http://microservices:1003`
- `/usercenter/`：请求被转发到`http://microservices:1004`