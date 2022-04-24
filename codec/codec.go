// codec 是用来编码消息的接口
package codec

import (
	"errors"
	"io"
)

const (
	Error MessageType = iota
	Request
	Response
	Event
)

var (
	ErrInvalidMessage = errors.New("invalid message")
)

type MessageType int

// 接受一个 connection/buffer 返回一个新的  Codec
type NewCodec func(io.ReadWriteCloser) Codec

// Codec encode/decode 在 go-micro 里面使用的各种消息
// ReadHeader and ReadBody are called in pairs to read requests/responses from the connection.
// Close is called when finished with the  connection.
// ReadBody may be called with a nil argument to force the body to be read and discarded.
type Codec interface {
	Reader
	Writer
	Close() error
	String() string
}

type Reader interface {
	ReadHeader(*Message, MessageType) error
	ReadBody(interface{}) error
}

type Writer interface {
	Write(*Message, interface{}) error
}

// Marshaler 是 broker/transport 用的简单的 encoding 接口
// where headers are not supported by the underlying implementation.
type Marshaler interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
	String() string
}

// Message 是通信的信息细节，如 body, 在发生错误的情况下，body 为 nil。
type Message struct {
	Id       string
	Type     MessageType
	Target   string
	Method   string
	Endpoint string
	Error    string

	// The values read from the socket
	Header map[string]string
	Body   []byte
}
