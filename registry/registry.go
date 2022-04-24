// Package registry is an interface for service discovery
// registry 包是一个服务发现的接口
// 默认注册中心为 mdns
package registry

import (
	"errors"
)

var (
	DefaultRegistry = NewRegistry()

	// Not found error when GetService is called
	ErrNotFound = errors.New("service not found")
	// Watcher stopped error when watcher is stopped
	ErrWatcherStopped = errors.New("watcher stopped")
)

// registry (注册中心)为服务发现提供了一个接口，并在不同的实现上提供了一个抽象
// {consul, etcd, zookeeper, ...}
type Registry interface {
	Init(...Option) error
	Options() Options
	Register(*Service, ...RegisterOption) error
	Deregister(*Service, ...DeregisterOption) error
	GetService(string, ...GetOption) ([]*Service, error)
	ListServices(...ListOption) ([]*Service, error)
	Watch(...WatchOption) (Watcher, error)
	String() string
}

// 服务
type Service struct {
	Name      string            `json:"name"`      // 服务名
	Version   string            `json:"version"`   // 版本号
	Metadata  map[string]string `json:"metadata"`  // metadata
	Endpoints []*Endpoint       `json:"endpoints"` // 端点
	Nodes     []*Node           `json:"nodes"`     // 节点
}

// 服务节点
type Node struct {
	Id       string            `json:"id"`
	Address  string            `json:"address"`
	Metadata map[string]string `json:"metadata"`
}

// 服务端点
type Endpoint struct {
	Name     string            `json:"name"`
	Request  *Value            `json:"request"`
	Response *Value            `json:"response"`
	Metadata map[string]string `json:"metadata"`
}

type Value struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []*Value `json:"values"`
}

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

type WatchOption func(*WatchOptions)

type DeregisterOption func(*DeregisterOptions)

type GetOption func(*GetOptions)

type ListOption func(*ListOptions)

// 注册服务节点。另外还提供 TTL 等选项。
func Register(s *Service, opts ...RegisterOption) error {
	return DefaultRegistry.Register(s, opts...)
}

// 卸载一个服务节点
func Deregister(s *Service) error {
	return DefaultRegistry.Deregister(s)
}

// 检索服务。因为我们将 Name/Version 分开，所以返回一个 slice。
func GetService(name string) ([]*Service, error) {
	return DefaultRegistry.GetService(name)
}

// 列出所有的服务。只返回服务名
func ListServices() ([]*Service, error) {
	return DefaultRegistry.ListServices()
}

// Watch 返回一个 watcher，允许您跟踪注册中心的更新。
func Watch(opts ...WatchOption) (Watcher, error) {
	return DefaultRegistry.Watch(opts...)
}

// 返回注册中心类型
func String() string {
	return DefaultRegistry.String()
}
