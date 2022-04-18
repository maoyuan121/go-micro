// store 包提供一个分布式数据存储的接口
// The design document is located at https://github.com/micro/development/blob/master/design/store.md
// 默认内存存储
package store

import (
	"errors"
	"time"
)

var (
	// ErrNotFound 当 key 不存在返回
	ErrNotFound = errors.New("not found")
	// DefaultStore 是内存存储
	DefaultStore Store = NewStore()
)

// Store 是一个数据存储接口
type Store interface {
	// Init 初始化 store。It must perform any required setup on the backing storage implementation and check that it is ready for use, returning any errors.
	Init(...Option) error
	// Options 允许你查看当前的 options
	Options() Options
	// Read takes a single key name and optional ReadOptions. 返回匹配的 []*Record 或者一个错误
	Read(key string, opts ...ReadOption) ([]*Record, error)
	// Write() 将一个 record 写到 store，如果写失败了那么返回一个  error
	Write(r *Record, opts ...WriteOption) error
	// Delete 从 store 中删除指定的 key 的 record
	Delete(key string, opts ...DeleteOption) error
	// List returns any keys that match, or an empty list with no error if none matched.
	List(opts ...ListOption) ([]string, error)
	// Close 关闭 store
	Close() error
	// String 返回实现的名字
	String() string
}

// Record 是存储在 store 中的一条记录
type Record struct {
	// The key to store the record
	Key string `json:"key"`
	// The value within the record
	Value []byte `json:"value"`
	// Any associated metadata for indexing
	Metadata map[string]interface{} `json:"metadata"`
	// Time to expire a record: TODO: change to timestamp
	Expiry time.Duration `json:"expiry,omitempty"`
}

func NewStore(opts ...Option) Store {
	return NewMemoryStore(opts...)
}
