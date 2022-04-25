# Wrappers

Wrappers 是中间件的一种形式，用来和 go-micro services 使用。
可以包装 client 和 server handler。

## Client Interface

```go
// Wrapper 包装一个 client 返回一个 client
type Wrapper func(Client) Client

// StreamWrapper 包装一个 Streamer 返回一个 Streamer
type StreamWrapper func(Streamer) Streamer
```

## Handler Interface

```go
// HandlerFunc 表示一个 handler 的一个方法。 
// 主要用来包装。
// 传递给实际方法的是具体的请求和响应类型。
type HandlerFunc func(ctx context.Context, req Request, rsp interface{}) error

// SubscriberFunc 表示一个 subscriber 的一个方法。 
// 主要用来包装。 
// What's handed to the actual method is the concrete publication message.
type SubscriberFunc func(ctx context.Context, msg Event) error

// HandlerWrapper 包装一个 HandlerFunc 返回一个 HandlerFunc
type HandlerWrapper func(HandlerFunc) HandlerFunc

// SubscriberWrapper 包装要给 SubscriberFunc 返回一个 SubscriberFunc
type SubscriberWrapper func(SubscriberFunc) SubscriberFunc

// StreamerWrapper wraps a Streamer interface and returns the equivalent.
// Because streams exist for the lifetime of a method invocation this
// is a convenient way to wrap a Stream as its in use for trace, monitoring,
// metrics, etc.
type StreamerWrapper func(Streamer) Streamer
```

## Client Wrapper Usage

下面是一个 client 的基本的 log wrapper

```go
type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}
```


## Handler Wrapper Usage

下面是一个 handler 的基本的 log wrapper

```go
func NewLogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[Log Wrapper] Before serving request method: %v", req.Endpoint())
		err := fn(ctx, req, rsp)
		log.Printf("[Log Wrapper] After serving request")
		return err
	}
}
```
