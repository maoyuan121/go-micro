package broker

import (
	"context"
	"crypto/tls"

	"go-micro.dev/v4/codec"
	"go-micro.dev/v4/registry"
)

type Options struct {
	Addrs  []string
	Secure bool
	Codec  codec.Marshaler

	// broker message 处理中发生错误时执行的处理程序
	ErrorHandler Handler

	TLSConfig *tls.Config
	// Registry used for clustering
	Registry registry.Registry
	// 接口实现的其他选项可以存储在上下文中
	Context context.Context
}

type PublishOptions struct {
	// 接口实现的其他选项可以存储在上下文中
	Context context.Context
}

type SubscribeOptions struct {
	// 默认为 true。当 handler 返回 nil error 消息被 acked。
	AutoAck bool
	// 具有相同队列名称的订阅者将创建共享订阅，其中每个订阅者都接收消息子集。
	Queue string

	// 接口实现的其他选项可以存储在上下文中
	Context context.Context
}

type Option func(*Options)

type PublishOption func(*PublishOptions)

// PublishContext set context
func PublishContext(ctx context.Context) PublishOption {
	return func(o *PublishOptions) {
		o.Context = ctx
	}
}

type SubscribeOption func(*SubscribeOptions)

func NewSubscribeOptions(opts ...SubscribeOption) SubscribeOptions {
	opt := SubscribeOptions{
		AutoAck: true,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Addrs sets the host addresses to be used by the broker
func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

// Codec sets the codec used for encoding/decoding used where
// a broker does not support headers
func Codec(c codec.Marshaler) Option {
	return func(o *Options) {
		o.Codec = c
	}
}

// DisableAutoAck will disable auto acking of messages
// after they have been handled.
func DisableAutoAck() SubscribeOption {
	return func(o *SubscribeOptions) {
		o.AutoAck = false
	}
}

// ErrorHandler will catch all broker errors that cant be handled
// in normal way, for example Codec errors
func ErrorHandler(h Handler) Option {
	return func(o *Options) {
		o.ErrorHandler = h
	}
}

// Queue sets the name of the queue to share messages on
func Queue(name string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Queue = name
	}
}

func Registry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

// Secure communication with the broker
func Secure(b bool) Option {
	return func(o *Options) {
		o.Secure = b
	}
}

// Specify TLS Config
func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}

// SubscribeContext set context
func SubscribeContext(ctx context.Context) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Context = ctx
	}
}
