// loader 包管理从多个源加载
package loader

import (
	"context"

	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/source"
)

// Loader manages loading sources
type Loader interface {
	// 关闭加载器
	Close() error
	// 从源中加载
	Load(...source.Source) error
	// 加载配置的快照
	Snapshot() (*Snapshot, error)
	// 强制同步源
	Sync() error
	// 监控修改
	Watch(...string) (Watcher, error)
	// 加载器名
	String() string
}

// 监视器允许您监视源并返回合并的变更集
type Watcher interface {
	// 对 next 的第一个调用可以返回当前的 Snapshot
	// 如果您正在监视一个路径，那么只返回来自该路径的数据。
	Next() (*Snapshot, error)
	// 停止监控修改
	Stop() error
}

// Snapshot 是一个合并的修改集
type Snapshot struct {
	// 合并的修改集
	ChangeSet *source.ChangeSet
	// 快照的确定性和可比版本
	Version string
}

type Options struct {
	Reader reader.Reader
	Source []source.Source

	// for alternative data
	Context context.Context
}

type Option func(o *Options)

// Copy snapshot
func Copy(s *Snapshot) *Snapshot {
	cs := *(s.ChangeSet)

	return &Snapshot{
		ChangeSet: &cs,
		Version:   s.Version,
	}
}
