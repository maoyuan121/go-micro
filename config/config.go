// config 包是一个动态配置接口
package config

import (
	"context"

	"go-micro.dev/v4/config/loader"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/file"
)

// Config 是动态配置接口
type Config interface {
	// 提供 reader.Values 接口
	reader.Values
	// 初始配置
	Init(opts ...Option) error
	// 配置的选项
	Options() Options
	// 停止 config 的 loader/watcher
	Close() error
	// 加载配置源
	Load(source ...source.Source) error
	// 强制同步源更改集
	Sync() error
	// 监视值的变化
	Watch(path ...string) (Watcher, error)
}

// Watcher 是配置的监视器 is the config watcher
type Watcher interface {
	Next() (reader.Value, error)
	Stop() error
}

type Options struct {
	Loader loader.Loader
	Reader reader.Reader
	Source []source.Source

	// for alternative data
	Context context.Context
}

type Option func(o *Options)

var (
	// Default Config Manager
	DefaultConfig, _ = NewConfig()
)

// NewConfig 返回一个新的 config
func NewConfig(opts ...Option) (Config, error) {
	return newConfig(opts...)
}

// config 作为 byte 返回
func Bytes() []byte {
	return DefaultConfig.Bytes()
}

// config 作为 map 返回
func Map() map[string]interface{} {
	return DefaultConfig.Map()
}

// Scan values to a go type
func Scan(v interface{}) error {
	return DefaultConfig.Scan(v)
}

// Force a source changeset sync
func Sync() error {
	return DefaultConfig.Sync()
}

// 从配置中获取一个值
func Get(path ...string) reader.Value {
	return DefaultConfig.Get(path...)
}

// 加载配置源
func Load(source ...source.Source) error {
	return DefaultConfig.Load(source...)
}

// Watch a value for changes
func Watch(path ...string) (Watcher, error) {
	return DefaultConfig.Watch(path...)
}

// LoadFile 是创建文件源并加载它的简写
func LoadFile(path string) error {
	return Load(file.NewSource(
		file.WithPath(path),
	))
}
