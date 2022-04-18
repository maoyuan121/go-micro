// broker 包是用于异步通信的一个接口
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
// 该处理程序被传递一个发布接口，该接口包含消息和可选的Ack方法，用于确认收到消息。
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

// Subscriber 是 Subscribe方法的一种方便的返回类型
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
