# Function

这是一个创建 micro function 的例子。function 是一次性执行的服务。

## Contents

- main.go - 是 function 的 main 定义
- proto - 包含这个 API 的 protobuf 定义

## Run function

```shell
while true; do
	github.com/asim/go-micro/examples/v4/function
done
```

## Call function

```shell
micro call greeter Greeter.Hello '{"name": "john"}'
```
