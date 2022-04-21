# Options
 
Go-micro 使用 [functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)。
他是一种设计模式， 允许在不更改方法签名的情况下添加新选项。

每个包都用一个 [Option](https://godoc.org/github.com/asim/go-micro#Option) type

```
type Option func(*Options)
```

Options 就像 [Name](https://godoc.org/github.com/asim/go-micro#Name) 函数设置服务名

实现如下

```
func Name(n string) Option {
	return func(o *Options) {
		o.Server.Init(server.Name(n))
	}
}
```

## Usage

Here's an example at the top level

```
import "go-micro.dev/v4"


service := micro.NewService(
	micro.Name("my.service"),
)
```
