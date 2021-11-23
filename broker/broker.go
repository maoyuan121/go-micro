// Package broker is an interface used for asynchronous messaging
package broker

// Broker 是异步消息使用的一个接口
type Broker interface {
	Init(...Option) error
	Options() Options
	Address() string
	Connect() error
	Disconnect() error
	Publish(topic string, m *Message, opts ...PublishOption) error
	Subscribe(topic string, h Handler, opts ...SubscribeOption) (Subscriber, error)
	String() string
}

// Handler 用于通过订阅主题来处理消息。
// The handler is passed a publication interface which contains the message and optional Ack method to acknowledge receipt of the message.
type Handler func(Event) error

type Message struct {
	Header map[string]string
	Body   []byte
}

// Event 交给订阅处理程序进行处理
type Event interface {
	Topic() string
	Message() *Message
	Ack() error
	Error() error
}

// Subscriber is a convenience return type for the Subscribe method
type Subscriber interface {
	Options() SubscribeOptions
	Topic() string
	Unsubscribe() error
}

var (
	DefaultBroker Broker = NewBroker()
)

func Init(opts ...Option) error {
	return DefaultBroker.Init(opts...)
}

func Connect() error {
	return DefaultBroker.Connect()
}

func Disconnect() error {
	return DefaultBroker.Disconnect()
}

func Publish(topic string, msg *Message, opts ...PublishOption) error {
	return DefaultBroker.Publish(topic, msg, opts...)
}

func Subscribe(topic string, handler Handler, opts ...SubscribeOption) (Subscriber, error) {
	return DefaultBroker.Subscribe(topic, handler, opts...)
}

func String() string {
	return DefaultBroker.String()
}
