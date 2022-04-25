# Version Wrapper

version wrapper 是一个有状态的客户端包装器，它使您能够只选择最新版本的服务。
适合于无需停机即可轻松升级运行的服务。


## Usage

当你创建你的 service 的时候传入 wrapper

```
wrapper := version.NewClientWrapper()

service := micro.NewService(
	micro.Name("foo"),
	micro.WrapClient(wrapper),
)
```
