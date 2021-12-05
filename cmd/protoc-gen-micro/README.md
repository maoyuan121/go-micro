# protoc-gen-micro

这是 go-micro 的 protobuf 代码生成。我们使用 protoc-gen-micro 来减少样板代码。

## Install

```
go install go-micro.dev/v4/cmd/protoc-gen-micro@latest
```

同样需要：

- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://google.golang.org/protobuf)

## Usage

定义你的服务 `greeter.proto`

```
syntax = "proto3";

service Greeter {
	rpc Hello(Request) returns (Response) {}
}

message Request {
	string name = 1;
}

message Response {
	string msg = 1;
}
```

生成代码

```
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
```

输出结果应该是：

```
./
    greeter.proto	# 原始的 protobuf 文件
    greeter.pb.go	# protoc-gen-go 自动生成的
    greeter.micro.go	# protoc-gen-micro 自动生成的
```

micro 生成的代码包含了 clients 和 handlers，他们减少了初始脚手架代码。

### 服务端

使用 micro server 注册 handler

```go
type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

proto.RegisterGreeterHandler(service.Server(), &Greeter{})
```

### 客户端

使用你的 micro client 创建一个服务的客户端

```go
client := proto.NewGreeterService("greeter", service.Client())
```

### Errors

如果你看到一个关于 `protoc-gen-micro` 没有被找到或执行的错误，很可能是你的环境没有被正确配置。
如果你已经安装了 `protoc`， `protoc-gen-go` 和 `protoc-gen-micro`，确保你已经将 `$GOPATH/bin` 包含在你的 `PATH` 中。

另一种选择是指定 Go 插件路径作为 `protoc` 命令的参数

```
protoc --plugin=protoc-gen-go=$GOPATH/bin/protoc-gen-go --plugin=protoc-gen-micro=$GOPATH/bin/protoc-gen-micro --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
```

### Endpoint

添加一个直接路由到 RPC 方法的 micor API 端点

用法：

1. 克隆 `github.com/googleapis/googleapis` 使用此功能，因为它需要 http annotation。
2. protoc 命令必须包含 `-I$GOPATH/src/github.com/googleapis/googleapis` 用于注释导入。





```diff
syntax = "proto3";

import "google/api/annotations.proto";

service Greeter {
	rpc Hello(Request) returns (Response) {
		option (google.api.http) = { post: "/hello"; body: "*"; };
	}
}

message Request {
	string name = 1;
}

message Response {
	string msg = 1;
}
```

proto 生成 `RegisterGreeterHandler` 函数其带有一个 [api.Endpoint](https://godoc.org/go-micro.dev/v3/api#Endpoint)。


```diff
func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) error {
	type greeter interface {
		Hello(ctx context.Context, in *Request, out *Response) error
	}
	type Greeter struct {
		greeter
	}
	h := &greeterHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Greeter.Hello",
		Path:    []string{"/hello"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&Greeter{h}, opts...))
}
```

## LICENSE

protoc-gen-micro is a liberal reuse of protoc-gen-go hence we maintain the original license 
