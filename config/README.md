# Config [![GoDoc](https://godoc.org/github.com/micro/go-micro/config?status.svg)](https://godoc.org/github.com/micro/go-micro/config)


Config 是一个可插拔的动态配置包

应用程序中的大多数配置都是静态配置的，或者包含从多个源加载的复杂逻辑。
Go Config 让这变得简单、可插拔和可合并。你再也不用以同样的方式处理配置了。




## 特性

- **Dynamic Loading** - 当需要时，从多个源加载配置。Go Config 在后台管理监视配置源，并自动合并和更新内存视图。

- **Pluggable Sources** - 从任意数量的源中选择加载和合并配置。后端源被抽象为一种内部使用并通过编码器解码的标准格式。源可以是 env var、flags、file、etcd、k8s configmap 等。

- **Mergeable Config** - 如果您指定多个配置源，无论格式如何，它们都将被合并并呈现在一个视图。这大大简化了基于环境的优先级顺序加载和更改。

- **Observe Changes** - 可以选择观察配置，以查看对特定值的更改。使用 Go Config 的监视器热加载你的应用程序。您不需要处理 ad-hoc hup 重载或其他任何事情，
只要继续阅读配置并在需要时观察更改即可得到通知。

- **Sane Defaults** - 如果配置加载不良或由于某些未知原因被完全删除，您可以指定 fallback 值，直接访问任何配置值。这确保在出现问题时，您总是能够读取一些合理的默认值。

## Getting Started

有关详细信息或架构、安装和一般用法，请参见[docs](https://micro.mu/docs/go-config.html)