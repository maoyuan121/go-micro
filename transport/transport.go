// transport 包是一个基于通信的同步连接接口
package transport

import (
	"time"
)

// Transport 是用于服务之间通信的接口。
// 它使用了基于 socket 的连接 send/recv语义，并有各种实现; http、grpc quic。
type Transport interface {
	Init(...Option) error
	Options() Options
	Dial(addr string, opts ...DialOption) (Client, error)
	Listen(addr string, opts ...ListenOption) (Listener, error)
	String() string
}

type Message struct {
	Header map[string]string
	Body   []byte
}

type Socket interface {
	Recv(*Message) error
	Send(*Message) error
	Close() error
	Local() string  // 本地 ip
	Remote() string // 远程 ip
}

type Client interface {
	Socket
}

type Listener interface {
	Addr() string
	Close() error
	Accept(func(Socket)) error
}

type Option func(*Options)

type DialOption func(*DialOptions)

type ListenOption func(*ListenOptions)

var (
	DefaultTransport Transport = NewHTTPTransport()

	DefaultDialTimeout = time.Second * 5
)
