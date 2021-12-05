# Go Micro [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/go-micro.dev/v4?tab=doc)

Go Micro 是一个分布式系统开发框架。

## 概要

Go Micro 提供分布式系统开发的核心需求，包括 RPC 和事件驱动通信。
**Micro** 哲学是一个具有可插拔架构的默认设置。我们提供默认值让您快速开始，但是所有东西都可以很容易地切换插拔。

## 特性

Go Micro 抽象了分布式系统的细节。以下是主要特性。

- **身份验证** - Auth 是一个天生的一等公民。身份验证和授权启用安全零信任网络，为每个服务提供身份和证书。这还包括规则基于访问控制。

- **Dynamic Config** - 从任何地方加载和热重载动态配置。config 接口提供了一种从任何来源加载应用程序级别配置的方式，如 env var，file，etcd。
您可以合并 source，甚至定义 fallbacks。

- **Data Storage** - 一个简单的数据存储接口，用于读取、写入和删除记录。它包括对内存、文件和默认的 CockroachDB。状态和持久性成为原型之外的核心需求，
Micro 希望将其构建到框架中。

- **Service Discovery** - 自动服务注册和名称解析。服务发现是微服务的核心。当服务 A 需要与服务 B 对话时，它需要该服务的位置。
默认发现机制为多播 DNS (mdns)，一种零配置系统。

- **Load Balancing** - 基于服务发现的客户端负载平衡。一旦我们有了任意数量实例的地址我们现在需要一种方法来决定要路由到哪个节点。
我们使用随机散列负载均衡来提供均匀分布，如果有问题，请跨服务重试不同的节点。

- **Message Encoding** - 基于 content-type 的动态消息编码。客户端和服务器将使用编解码器和 content-type 为您无缝地编码和解码 Go 类型。
任何种类的消息都可以从不同的客户端进行编码和发送。客户端服务器默认处理这个。默认情况下包括 protobuf 和 json。

- **RPC Client/Server** - 支持双向流的基于 RPC 的请求/响应。我们为 synchronous 提供了一个抽象沟通。向服务发出的请求将被自动解析、负载平衡、拨号和流处理。

- **Async Messaging** - PubSub 是作为异步通信和事件驱动架构的一级公民而构建的。事件通知是微服务开发中的核心模式。默认的消息传递系统是一个 HTTP 事件消息代理。

- **Event Streaming** - PubSub 非常适合异步通知，但对于更高级的用例，事件流是首选。提供持久存储，consuming from offsets and acking。Go Micro 包含对 NATS Jetstream 和 Redis流的支持。

- **Synchronization** - 分布式系统通常以一种最终一致的方式构建。支持分布式锁定和 leadership 是内置的 Sync 接口。当使用最终一致的数据库或调度时，请使用 Sync 接口。

- **Pluggable Interfaces** - Go Micro 为每个分布式系统抽象使用了 Go 接口。因此，这些接口都是可插拔的，并允许 Go Micro 运行时不可知。你可以添加任何底层技术。

## Getting Started

使用 Go Micro

```golang
import "go-micro.dev/v4"

// 创建一个新服务
service := micro.NewService(
    micro.Name("helloworld"),
)

// initialise flags
service.Init()

// 启动服务
service.Run()
```

更多使用的详细信息请查看 [examples](https://github.com/micro/go-micro/tree/master/examples)。

## 命令行 interface

查看 [cmd/micro](https://github.com/asim/go-micro/tree/master/cmd/micro) 了解更多 command line interface 的信息。

## 代码生成

查看 [cmd/protoc-gen-micro](https://github.com/micro/go-micro/tree/master/cmd/protoc-gen-micro) 了解 protobuf 代码生成。


## 使用示例

使用示例请查看 [examples](https://github.com/micro/go-micro/tree/master/examples) 目录。

## Plugins

所有的插件位于 [plugins](https://github.com/micro/go-micro/tree/master/plugins) 目录。

## Services

第三方服务位于 [services](https://github.com/micro/go-micro/tree/master/services) 目录。

## License

Go Micro is Apache 2.0 licensed.
