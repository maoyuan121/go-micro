# Stream

这个例子中有一个 streaming service 和两个 client，其中一个是 streaming rpc client 另外一个是 websocket client。

## Contents

- server - 是服务
- client - 是 rpc 客户端
- web - 是 websocket 客户端

## 运行例子

运行服务

```shell
go run server/main.go
```

运行 client

```shell
go run client/main.go
```

为 websocket client 运行 micro web reverse proxy

``` shell
micro web
```

运行 websocket client

```shell
cd web # must be in the web directory to serve static files.
go run main.go
```

访问 http://localhost:8082/stream，发送一个请求！

And that's all there is to it.
