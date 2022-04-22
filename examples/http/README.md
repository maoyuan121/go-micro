# http server and http client demo

一个 http 应用的例子

## Contents

- **srv** - 一个作为 go-micro service 的 http server
- **cli** - 一个 http client 用来调用 http server
- **rpcli** - 一个 http client 用来调用 rpc server


## Run Service
Start http server
```shell
go run srv/main.go
```

## Client

Call http client
```shell
go run cli/main.go

```


## Run rpc Service
Start greeter service
```shell
go run ../greeter/srv/main.go
```

## Client
http client call rpc service
```shell
go run rpccli/main.go
```
