// server 包是 micro server 的接口
package server

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"go-micro.dev/v4/codec"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	signalutil "go-micro.dev/v4/util/signal"
)

// Server 是一个简单的 micro server 抽象
type Server interface {
	// 初始化选项
	Init(...Option) error
	// 获取选项
	Options() Options
	// 注册一个 handler
	Handle(Handler) error
	// 创建一个新的 handler
	NewHandler(interface{}, ...HandlerOption) Handler
	// 创建一个新的  subscriber
	NewSubscriber(string, interface{}, ...SubscriberOption) Subscriber
	// 注册一个 subscriber
	Subscribe(Subscriber) error
	// 启动 server
	Start() error
	// 停止 server
	Stop() error
	// 返回 Server 实现名
	String() string
}

// Router 处理服务消息
type Router interface {
	// ProcessMessage 处理一条消息
	ProcessMessage(context.Context, Message) error
	// ServeRequest 处理请求以完成
	ServeRequest(context.Context, Request, Response) error
}

// Message 是异步消息接口
type Message interface {
	// 消息的 topic
	Topic() string
	// The decoded payload value
	Payload() interface{}
	// The content type of the payload
	ContentType() string
	// The raw headers of the message
	Header() map[string]string
	// The raw body of the message
	Body() []byte
	// Codec used to decode the message
	Codec() codec.Reader
}

// Request 是同步请求接口
type Request interface {
	// 请求的服务名
	Service() string
	// 请求的  action
	Method() string
	// 请求的 Endpoint name
	Endpoint() string
	// Content type provided
	ContentType() string
	// Header of the request
	Header() map[string]string
	// Body is the initial decoded value
	Body() interface{}
	// Read the undecoded request body
	Read() ([]byte, error)
	// The encoded message stream
	Codec() codec.Reader
	// Indicates whether its a stream
	Stream() bool
}

// Response is the response writer for unencoded messages
type Response interface {
	// Encoded writer
	Codec() codec.Writer
	// Write the header
	WriteHeader(map[string]string)
	// write a response directly to the client
	Write([]byte) error
}

// Stream represents a stream established with a client.
// A stream can be bidirectional which is indicated by the request.
// The last error will be left in Error().
// EOF indicates end of the stream.
type Stream interface {
	Context() context.Context
	Request() Request
	Send(interface{}) error
	Recv(interface{}) error
	Error() error
	Close() error
}

// Handler interface represents a request handler. It's generated
// by passing any type of public concrete object with endpoints into server.NewHandler.
// Most will pass in a struct.
//
// Example:
//
//      type Greeter struct {}
//
//      func (g *Greeter) Hello(context, request, response) error {
//              return nil
//      }
//
type Handler interface {
	Name() string
	Handler() interface{}
	Endpoints() []*registry.Endpoint
	Options() HandlerOptions
}

// Subscriber interface represents a subscription to a given topic using
// a specific subscriber function or object with endpoints. It mirrors
// the handler in its behaviour.
type Subscriber interface {
	Topic() string
	Subscriber() interface{}
	Endpoints() []*registry.Endpoint
	Options() SubscriberOptions
}

type Option func(*Options)

var (
	DefaultAddress                 = ":0"
	DefaultName                    = "go.micro.server"
	DefaultVersion                 = "latest"
	DefaultId                      = uuid.New().String()
	DefaultServer           Server = newRpcServer()
	DefaultRouter                  = newRpcRouter()
	DefaultRegisterCheck           = func(context.Context) error { return nil }
	DefaultRegisterInterval        = time.Second * 30
	DefaultRegisterTTL             = time.Second * 90

	// NewServer creates a new server
	NewServer func(...Option) Server = newRpcServer
)

// DefaultOptions returns config options for the default service
func DefaultOptions() Options {
	return DefaultServer.Options()
}

// Init initialises the default server with options passed in
func Init(opt ...Option) {
	if DefaultServer == nil {
		DefaultServer = newRpcServer(opt...)
	}
	DefaultServer.Init(opt...)
}

// NewRouter returns a new router
func NewRouter() *router {
	return newRpcRouter()
}

// NewSubscriber creates a new subscriber interface with the given topic
// and handler using the default server
func NewSubscriber(topic string, h interface{}, opts ...SubscriberOption) Subscriber {
	return DefaultServer.NewSubscriber(topic, h, opts...)
}

// NewHandler creates a new handler interface using the default server
// Handlers are required to be a public object with public
// endpoints. Call to a service endpoint such as Foo.Bar expects
// the type:
//
//	type Foo struct {}
//	func (f *Foo) Bar(ctx, req, rsp) error {
//		return nil
//	}
//
func NewHandler(h interface{}, opts ...HandlerOption) Handler {
	return DefaultServer.NewHandler(h, opts...)
}

// Handle registers a handler interface with the default server to
// handle inbound requests
func Handle(h Handler) error {
	return DefaultServer.Handle(h)
}

// Subscribe registers a subscriber interface with the default server
// which subscribes to specified topic with the broker
func Subscribe(s Subscriber) error {
	return DefaultServer.Subscribe(s)
}

// Run starts the default server and waits for a kill
// signal before exiting. Also registers/deregisters the server
func Run() error {
	if err := Start(); err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signalutil.Shutdown()...)

	if logger.V(logger.InfoLevel, logger.DefaultLogger) {
		logger.Infof("Received signal %s", <-ch)
	}
	return Stop()
}

// Start starts the default server
func Start() error {
	config := DefaultServer.Options()
	if logger.V(logger.InfoLevel, logger.DefaultLogger) {
		logger.Infof("Starting server %s id %s", config.Name, config.Id)
	}
	return DefaultServer.Start()
}

// Stop stops the default server
func Stop() error {
	if logger.V(logger.InfoLevel, logger.DefaultLogger) {
		logger.Infof("Stopping server")
	}
	return DefaultServer.Stop()
}

// String returns name of Server implementation
func String() string {
	return DefaultServer.String()
}
