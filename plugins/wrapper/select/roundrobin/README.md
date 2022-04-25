# Round Robin Wrapper

轮询包装器是一个有状态的客户端包装器，它为选择器提供了一个真正的轮询策略

## Usage

当你创建你的 service 的时候传入 wrapper

```
wrapper := roundrobin.NewClientWrapper()

service := micro.NewService(
	micro.Name("foo"),
	micro.WrapClient(wrapper),
)
```
