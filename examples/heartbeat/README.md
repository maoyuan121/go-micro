# Heartbeat

在服务发现中使用心跳的 demo。

## Rationale

服务在启动时注册服务发现，在关闭时注销。有时这些服务可能会意外死亡或被强行杀害或面临短暂的网络问题。
在这些情况下，陈旧的节点将留在服务发现中。这将是如果服务是自动删除的原因。

## Solution

出于这个原因，Micro 支持 register TTL 和 register 间隔选项。
TTL 指定一个注册应该持续多长时间存在于发现中，之后过期并被删除。间隔是服务重新注册以保持其在服务发现中的注册的时间。

这些选项在 go-micro 中可用，在 micro 工具包中作为标志

## Toolkit

像这样运行工具包的任何组件

```
micro --register_ttl=30 --register_interval=15 api
```

这个例子显示我们将 ttl 设置为 30 秒，重新注册间隔为 15 秒。

## Go Micro

在声明微服务时，可以将选项作为 time.Duration 传入。

```
service := micro.NewService(
	micro.Name("com.example.srv.foo"),
	micro.RegisterTTL(time.Second*30),
	micro.RegisterInterval(time.Second*15),
)
```
