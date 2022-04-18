// trace 包为分布式跟踪提供了一个接口
package trace

import (
	"context"
	"time"

	"go-micro.dev/v4/metadata"
)

// Tracer 是分布式跟踪的一个接口
type Tracer interface {
	// 开始跟踪
	Start(ctx context.Context, name string) (context.Context, *Span)
	// 结束追踪
	Finish(*Span) error
	// 读取追踪
	Read(...ReadOption) ([]*Span, error)
}

// SpanType 描述跟踪跨度的性质
type SpanType int

const (
	// SpanTypeRequestInbound 是服务请求时创建的 span
	SpanTypeRequestInbound SpanType = iota
	// SpanTypeRequestOutbound 在进行服务调用时创建的 span
	SpanTypeRequestOutbound
)

// Span 用于记录一个 entry
type Span struct {
	// trace 的 id
	Trace string
	// span 的名
	Name string
	// span 的 id
	Id string
	// 父 span id
	Parent string
	// 开始时间
	Started time.Time
	// Duration 单位纳秒
	Duration time.Duration
	// 关联的数据
	Metadata map[string]string
	// 类型
	Type SpanType
}

const (
	traceIDKey = "Micro-Trace-Id"
	spanIDKey  = "Micro-Span-Id"
)

// FromContext returns a span from context
func FromContext(ctx context.Context) (traceID string, parentSpanID string, isFound bool) {
	traceID, traceOk := metadata.Get(ctx, traceIDKey)
	microID, microOk := metadata.Get(ctx, "Micro-Id")
	if !traceOk && !microOk {
		isFound = false
		return
	}
	if !traceOk {
		traceID = microID
	}
	parentSpanID, ok := metadata.Get(ctx, spanIDKey)
	return traceID, parentSpanID, ok
}

// ToContext saves the trace and span ids in the context
func ToContext(ctx context.Context, traceID, parentSpanID string) context.Context {
	return metadata.MergeContext(ctx, map[string]string{
		traceIDKey: traceID,
		spanIDKey:  parentSpanID,
	}, true)
}

var (
	DefaultTracer Tracer = NewTracer()
)

type noop struct{}

func (n *noop) Init(...Option) error {
	return nil
}

func (n *noop) Start(ctx context.Context, name string) (context.Context, *Span) {
	return nil, nil
}

func (n *noop) Finish(*Span) error {
	return nil
}

func (n *noop) Read(...ReadOption) ([]*Span, error) {
	return nil, nil
}
