// log 包提供了一个 log 接口
package logger

var (
	// 默认 logger
	DefaultLogger Logger = NewLogger()
)

// Logge 是一个 logging 接口
type Logger interface {
	// Init 初始化选项
	Init(options ...Option) error
	// 获取选项
	Options() Options
	// Fields 设置总是被记录的字段
	Fields(fields map[string]interface{}) Logger
	// Log 写一个日志条目
	Log(level Level, v ...interface{})
	// Logf writes a formatted log entry
	Logf(level Level, format string, v ...interface{})
	// String 返回 logger 的实现名
	String() string
}

func Init(opts ...Option) error {
	return DefaultLogger.Init(opts...)
}

func Fields(fields map[string]interface{}) Logger {
	return DefaultLogger.Fields(fields)
}

func Log(level Level, v ...interface{}) {
	DefaultLogger.Log(level, v...)
}

func Logf(level Level, format string, v ...interface{}) {
	DefaultLogger.Logf(level, format, v...)
}

func String() string {
	return DefaultLogger.String()
}
