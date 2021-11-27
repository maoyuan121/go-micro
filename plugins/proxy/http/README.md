# HTTP Proxy

这是一个 http 代理插件，它将 RPC 转换为 HTTP 请求

## Overview

`NewService` 返回一个新的 http 代理。它充当微服务并代理到 http 后端。
路由是动态设置的，例如 Foo.Bar路由到 /foo/bar。默认后端是 http:localhost:9090。
可选地指定后端端点 url 或路由。还可以选择注册特定的端点。


## Usage

```
service := NewService(
      micro.Name("greeter"),
      // Sets the default http endpoint
      http.WithBackend("http:localhost:10001"),
)

// Set fixed backend endpoints
// register an endpoint
http.RegisterEndpoint("Hello.World", "/helloworld")

service := NewService(
      micro.Name("greeter"),
      // Set the http endpoint
      http.WithBackend("http:localhost:10001"),
)
```
