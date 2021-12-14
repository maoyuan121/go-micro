# Client

## Contents

- main.go - calls each of the go.micro.srv.example handlers and includes the use of the streaming handler
- codegen - 演示如何使用 code generation 移除 boilerplate code
- dc_filter - 演示在一个 call wrapper 使用 Select filters filtering to the local DC
- dc_selector - 和 dc_filter 一样，但是是一个 Selector 的实现
- pub - 使用 Publish 方法发布消息。默认的编码是 protobuf
- selector - 如何编写和加载你自己的 selector
- wrapper - 如何使用  client wrappers（中间件）

