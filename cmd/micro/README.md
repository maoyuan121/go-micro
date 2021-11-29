# Micro CLI

Micro CLI 是用来开发 [Go Micro][1] 的 command line interface。

## Getting Started

[Download][2] 并安装 **Go**。版本需要是 `1.16` 或以上。

使用 [`go install`][3] 命令安装。

```bash
go install go-micro.dev/v4/cmd/micro@master
```

使用 `new` 命令创建一个新的服务。

```bash
micro new service helloworld
```

按照屏幕上的说明操作。接下来，我们可以运行程序。

```bash
cd helloworld
make proto tidy
micro run
```

Finally we can call the service.

```bash
micro call helloworld Helloworld.Call '{"name": "John"}'
```

这就是你开始所需要知道的一切。有关开发服务的更多信息，请参考 [Go Micro][1] 文档。

## 依赖

你需要 protoc-gen-micro 用来生成代码

* [protobuf][4]
* [protoc-gen-go][5]
* [protoc-gen-micro][6]

```bash
# Download latest proto release
# https://github.com/protocolbuffers/protobuf/releases
go get -u google.golang.org/protobuf/proto
go install github.com/golang/protobuf/protoc-gen-go@latest
go install go-micro.dev/v4/cmd/protoc-gen-micro@latest
```

## Creating A Service

使用 `micro new service` 命令创建一个新的服务。

```bash
$ micro new service helloworld
creating service helloworld

download protoc zip packages (protoc-$VERSION-$PLATFORM.zip) and install:

visit https://github.com/protocolbuffers/protobuf/releases/latest

download protobuf for go-micro:

go get -u google.golang.org/protobuf/proto
go install github.com/golang/protobuf/protoc-gen-go@latest
go install go-micro.dev/cmd/protoc-gen-micro/v4@latest

compile the proto file helloworld.proto:

cd helloworld
make proto tidy
```

使用 `micro new function` 命令创建一个新的 fucntion。Functin 和 services 不同的是运行后就退出了。

```bash
$ micro new function helloworld
creating function helloworld

download protoc zip packages (protoc-$VERSION-$PLATFORM.zip) and install:

visit https://github.com/protocolbuffers/protobuf/releases/latest

download protobuf for go-micro:

go get -u google.golang.org/protobuf/proto
go install github.com/golang/protobuf/protoc-gen-go@latest
go install go-micro.dev/cmd/protoc-gen-micro/v4@latest

compile the proto file helloworld.proto:

cd helloworld
make proto tidy
```

### Jaeger

要创建一个集成了 [Jaeger][7] 的服务，需要在使用 `micro new service` 或 `micro new function` 的时候传入 `--jaeger` flag.
你也可以使用 [environment variables][8] 配置 Jaeger 客户端。

```bash
micro new service --jaeger helloworld
```

你可以调用 `trace.NewSpan(context.Context).Finish()` 来嵌套 span。
例如，考虑下面的处理程序实现一个问候程序。

`handler/helloworld.go`

```go
package helloworld

import (
    "context"

    log "go-micro.dev/v4/logger"

    "helloworld/greeter"
    pb "helloworld/proto"
)

type Helloworld struct{}

func (e *Helloworld) Call(ctx context.Context, req pb.CallRequest, rsp *pb.CallResponse) error {
    log.Infof("Received Helloworld.Call request: %v", req)
    rsp.Msg = greeter.Greet(ctx, req.Name)
    return nil
}
```

`greeter/greeter.go`

```go
package greeter

import (
    "context"
    "fmt"

    "go-micro.dev/v4/cmd/micro/debug/trace"
)

func Greet(ctx context.Context, name string) string {
    defer trace.NewSpan(ctx).Finish()
    return fmt.Sprint("Hello " + name)
}
```

### Skaffold

要创建一个集成了 [Skaffold][9] 文件的服务，需要在使用 `micro new service` 或 `micro new function` 的时候传入 `--skaffold` flag.

```bash
micro new service --skaffold helloworld
```

## Running A Service

要运行一个服务，使用 `micro run` 命令来 build 和运行你的服务。

```bash
$ micro run
2021-08-20 14:05:54  file=v3@v3.5.2/service.go:199 level=info Starting [service] helloworld
2021-08-20 14:05:54  file=server/rpc_server.go:820 level=info Transport [http] Listening on [::]:34531
2021-08-20 14:05:54  file=server/rpc_server.go:840 level=info Broker [http] Connected to 127.0.0.1:44975
2021-08-20 14:05:54  file=server/rpc_server.go:654 level=info Registry [mdns] Registering node: helloworld-45f43a6f-5fc0-4b0d-af73-e4a10c36ef54
```

### With Docker

要使用 docker 运行一个服务，构建 docker images 并运行 docker container。

```bash
$ make docker
$ docker run helloworld:latest
2021-08-20 12:07:31  file=v3@v3.5.2/service.go:199 level=info Starting [service] helloworld
2021-08-20 12:07:31  file=server/rpc_server.go:820 level=info Transport [http] Listening on [::]:36037
2021-08-20 12:07:31  file=server/rpc_server.go:840 level=info Broker [http] Connected to 127.0.0.1:46157
2021-08-20 12:07:31  file=server/rpc_server.go:654 level=info Registry [mdns] Registering node: helloworld-31f58714-72f5-4d12-b2eb-98f66aea7a34
```

### With Skaffold

当你使用 `--skaffold` flag 创建了一个服务，你可以使用 `skaffold` 命令运行 Skaffold pipeline。

```bash
skaffold dev
```

## Creating A Client

使用 `micro new client` 命令创建一个新的 client。
该名称是要为其创建客户端项目的服务。

```bash
$ micro new client helloworld
creating client helloworld
cd helloworld-client
make tidy
```

您可以选择传递要为其创建客户端项目的服务的完全限定包名。

```bash
$ micro new client github.com/auditemarlow/helloworld
creating client helloworld
cd helloworld-client
make tidy
```

## Running A Client

要运行一个 client，使用 `micro run` 命令构建并运行。

```bash
$ micro run
2021-09-03 12:52:23  file=helloworld-client/main.go:33 level=info msg:"Hello John"
```

## Generating Files

使用 `micro generate` 命令生成 Go Micro 项目模板文件。它将生成的文件放在当前工作目录中。

```bash
$ micro generate skaffold
skaffold project template files generated
```

## Listing Services

使用 `micro services` 命令列出服务。

```bash
$ micro services
helloworld
```

## Describing A Service

使用 `micro describe service` 命令描述服务。

```bash
$ micro describe service helloworld
{
  "name": "helloworld",
  "version": "latest",
  "metadata": null,
  "endpoints": [
    {
      "name": "Helloworld.Call",
      "request": {
        "name": "CallRequest",
        "type": "CallRequest",
        "values": [
          {
            "name": "name",
            "type": "string",
            "values": null
          }
        ]
      },
      "response": {
        "name": "CallResponse",
        "type": "CallResponse",
        "values": [
          {
            "name": "msg",
            "type": "string",
            "values": null
          }
        ]
      }
    }
  ],
  "nodes": [
    {
      "id": "helloworld-9660f06a-d608-43d9-9f44-e264ff63c554",
      "address": "172.26.165.161:45059",
      "metadata": {
        "broker": "http",
        "protocol": "mucp",
        "registry": "mdns",
        "server": "mucp",
        "transport": "http"
      }
    }
  ]
}
```

可以传入 `--format=yaml` flag 用来输出 YAML 格式对象。

```bash
$ micro describe service --format=yaml helloworld
name: helloworld
version: latest
metadata: {}
endpoints:
- name: Helloworld.Call
  request:
    name: CallRequest
    type: CallRequest
    values:
    - name: name
      type: string
      values: []
  response:
    name: CallResponse
    type: CallResponse
    values:
    - name: msg
      type: string
      values: []
nodes:
- id: helloworld-9660f06a-d608-43d9-9f44-e264ff63c554
  address: 172.26.165.161:45059
  metadata:
    broker: http
    protocol: mucp
    registry: mdns
    server: mucp
    transport: http
```

## Calling A Service

使用 `micro call` 命令调用一个服务。这将发送一个请求并期望得到一个响应。

```bash
$ micro call helloworld Helloworld.Call '{"name": "John"}'
{"msg":"Hello John"}
```

要调用服务的服务器流，使用 `micro stream server` 命令。
这将发送一个请求并等待一个响应流。

```bash
$ micro stream server helloworld Helloworld.ServerStream '{"count": 10}'
{"count":0}
{"count":1}
{"count":2}
{"count":3}
{"count":4}
{"count":5}
{"count":6}
{"count":7}
{"count":8}
{"count":9}
```

要调用服务的双向流，使用 `micro stream bidi` 命令。这将发送一个请求流并期待一个响应流。

```bash
$ micro stream bidi helloworld Helloworld.BidiStream '{"stroke": 1}' '{"stroke": 2}' '{"stroke": 3}'
{"stroke":1}
{"stroke":2}
{"stroke":3}
```

[1]: https://go-micro.dev
[2]: https://golang.org/dl/
[3]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[4]: https://grpc.io/docs/protoc-installation/
[5]: https://micro.mu/github.com/golang/protobuf/protoc-gen-go
[6]: https://go-micro.dev/tree/master/cmd/protoc-gen-micro
[7]: https://www.jaegertracing.io/
[8]: https://github.com/jaegertracing/jaeger-client-go#environment-variables
[9]: https://skaffold.dev/
