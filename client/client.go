// client 包提供一个 RPC client 的接口
package client

import (
	"context"
	"time"

	"go-micro.dev/v4/codec"
)

// Client 使用来请求 services 的接口
// 它通过 Transport 支持 Request/Response，通过 Broker 支持 Publishing
// 它还支持请求双向流
type Client interface {
	Init(...Option) error
	Options() Options
	NewMessage(topic string, msg interface{}, opts ...MessageOption) Message
	NewRequest(service, endpoint string, req interface{}, reqOpts ...RequestOption) Request
	Call(ctx context.Context, req Request, rsp interface{}, opts ...CallOption) error
	Stream(ctx context.Context, req Request, opts ...CallOption) (Stream, error)
	Publish(ctx context.Context, msg Message, opts ...PublishOption) error
	String() string
}

// Router 管理 request routing
type Router interface {
	SendRequest(context.Context, Request) (Response, error)
}

// Message 是一个接口用来异步发布
type Message interface {
	Topic() string
	Payload() interface{}
	ContentType() string
}

// Request 是 Call 和 Stream 方法用来同步请求的接口
type Request interface {
	// 要调用的 service
	Service() string
	// The action to take
	Method() string
	// 要 invoke 的 endpoint
	Endpoint() string
	// The content type
	ContentType() string
	// The unencoded request body
	Body() interface{}
	// Write to the encoded request writer. This is nil before a call is made
	Codec() codec.Writer
	// indicates whether the request will be a streaming one rather than unary
	Stream() bool
}

// Response 是从一个 service 接收到的 response
type Response interface {
	// Read the response
	Codec() codec.Reader
	// read the header
	Header() map[string]string
	// Read the undecoded response
	Read() ([]byte, error)
}

// Stream is the inteface for a bidirectional synchronous stream
type Stream interface {
	Closer
	// Context for the stream
	Context() context.Context
	// The request made
	Request() Request
	// The response read
	Response() Response
	// Send will encode and send a request
	Send(interface{}) error
	// Recv will decode and read a response
	Recv(interface{}) error
	// Error returns the stream error
	Error() error
	// Close closes the stream
	Close() error
}

// Closer handle client close
type Closer interface {
	// CloseSend closes the send direction of the stream.
	CloseSend() error
}

// Option used by the Client
type Option func(*Options)

// CallOption used by Call or Stream
type CallOption func(*CallOptions)

// PublishOption used by Publish
type PublishOption func(*PublishOptions)

// MessageOption used by NewMessage
type MessageOption func(*MessageOptions)

// RequestOption used by NewRequest
type RequestOption func(*RequestOptions)

var (
	// DefaultClient 是开箱即用的默认的 client
	DefaultClient Client = newRpcClient()
	// DefaultBackoff 是重试的默认 backoff function
	DefaultBackoff = exponentialBackoff
	// DefaultRetry 是重试默认的 check-for-retry function
	DefaultRetry = RetryOnError
	// DefaultRetries 是默认的重试次数
	DefaultRetries = 1
	// DefaultRequestTimeout 是默认的请求超时时间
	DefaultRequestTimeout = time.Second * 5
	// DefaultPoolSize 是默认的连接池大小
	DefaultPoolSize = 100
	// DefaultPoolTTL 是默认的连接池 ttl
	DefaultPoolTTL = time.Minute

	// NewClient 返回一个新 client
	NewClient func(...Option) Client = newRpcClient
)

// Makes a synchronous call to a service using the default client
func Call(ctx context.Context, request Request, response interface{}, opts ...CallOption) error {
	return DefaultClient.Call(ctx, request, response, opts...)
}

// Publishes a publication using the default client. Using the underlying broker
// set within the options.
func Publish(ctx context.Context, msg Message, opts ...PublishOption) error {
	return DefaultClient.Publish(ctx, msg, opts...)
}

// Creates a new message using the default client
func NewMessage(topic string, payload interface{}, opts ...MessageOption) Message {
	return DefaultClient.NewMessage(topic, payload, opts...)
}

// Creates a new request using the default client. Content Type will
// be set to the default within options and use the appropriate codec
func NewRequest(service, endpoint string, request interface{}, reqOpts ...RequestOption) Request {
	return DefaultClient.NewRequest(service, endpoint, request, reqOpts...)
}

// Creates a streaming connection with a service and returns responses on the
// channel passed in. It's up to the user to close the streamer.
func NewStream(ctx context.Context, request Request, opts ...CallOption) (Stream, error) {
	return DefaultClient.Stream(ctx, request, opts...)
}

func String() string {
	return DefaultClient.String()
}
