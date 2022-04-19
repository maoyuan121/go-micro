// micro 包是微服务的可插拔框架
package micro

import (
	"context"

	"go-micro.dev/v4/client"
	"go-micro.dev/v4/server"
)

type serviceKey struct{}

// Service 是一个接口，它将底层库封装在 go-micro 中。
// 这是一种构建和初始化服务的方便方法。
type Service interface {
	// 服务名
	Name() string
	// Init 初始化选项
	Init(...Option)
	// Options 返回选项
	Options() Options
	// Client 用来调用服务
	Client() client.Client
	// Server 用来处理请求和事件
	Server() server.Server
	// 运行服务
	Run() error
	// 服务实现名
	String() string
}

// Function 是一个一次性执行的服务
type Function interface {
	// 继承自 service 接口
	Service
	// Done 完成执行的信号
	Done() error
	// Handle 注册一个 RPC 处理程序
	Handle(v interface{}) error
	// Subscribe 注册一个订阅者
	Subscribe(topic string, v interface{}) error
}

// Event 用来发布消息到一个 topic
type Event interface {
	// Publish 发布消息到一个 event topic
	Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error
}

// Type alias to satisfy the deprecation
type Publisher = Event

type Option func(*Options)

// NewService creates and returns a new Service based on the packages within.
func NewService(opts ...Option) Service {
	return newService(opts...)
}

// FromContext 从 Context 中获取一个 Service
func FromContext(ctx context.Context) (Service, bool) {
	s, ok := ctx.Value(serviceKey{}).(Service)
	return s, ok
}

// NewContext 返回一个新的 Context，其中嵌入了 Service
func NewContext(ctx context.Context, s Service) context.Context {
	return context.WithValue(ctx, serviceKey{}, s)
}

// NewFunction returns a new Function for a one time executing Service
func NewFunction(opts ...Option) Function {
	return newFunction(opts...)
}

// NewEvent 创建一个新的 event publisher
func NewEvent(topic string, c client.Client) Event {
	if c == nil {
		c = client.NewClient()
	}
	return &event{c, topic}
}

// Deprecated: NewPublisher returns a new Publisher
func NewPublisher(topic string, c client.Client) Event {
	return NewEvent(topic, c)
}

// RegisterHandler 是注册一个 handler 的语法糖
func RegisterHandler(s server.Server, h interface{}, opts ...server.HandlerOption) error {
	return s.Handle(s.NewHandler(h, opts...))
}

// RegisterSubscriber 是注册一个 subscriber 的语法糖
func RegisterSubscriber(topic string, s server.Server, h interface{}, opts ...server.SubscriberOption) error {
	return s.Subscribe(s.NewSubscriber(topic, h, opts...))
}
