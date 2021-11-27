// Package selector is a way to pick a list of service nodes
package selector

import (
	"errors"

	"go-micro.dev/v4/registry"
)

// Selector 建立在注册表 registry 上，作为一种机制来选择节点并标记它们的状态。
// 这允许使用各种算法来构建主机池和其他东西。
type Selector interface {
	Init(opts ...Option) error
	Options() Options
	// Select 返回一个函数，该函数将返回下一个服务节点
	Select(service string, opts ...SelectOption) (Next, error)
	// 标记针对节点设置成功/错误
	Mark(service string, node *registry.Node, err error)
	// Reset 将一个服务的状态返回为零
	Reset(service string)
	// Close 显示选择器不可用
	Close() error
	// 选择器的名称
	String() string
}

// Next 是一个函数，它根据选择器的策略返回下一个节点
type Next func() (*registry.Node, error)

// Filter 用于在选择过程中过滤服务
type Filter func([]*registry.Service) []*registry.Service

// Strategy is a selection strategy e.g random, round robin
type Strategy func([]*registry.Service) Next

var (
	DefaultSelector = NewSelector()

	ErrNotFound      = errors.New("not found")
	ErrNoneAvailable = errors.New("none available")
)
