package cache

import (
	"context"
	"errors"
	"time"
)

var (
	// DefaultCache 是默认的缓存（内存）
	DefaultCache Cache = NewCache()
	// DefaultExpiration 是默认的过期时间
	DefaultExpiration time.Duration = 0

	// 如果 Cache.Get 找到的 item 过期了，则返回  ErrItemExpired
	ErrItemExpired error = errors.New("item has expired")
	// Cache.Get 和 Cache.Delete 中给定的 Key 没有找到则返回  ErrKeyNotFound
	ErrKeyNotFound error = errors.New("key not found in cache")
)

// Cache 是 cache 的接口
//
// Context 是 cache 指定的  context
// Get 根据 key 返回一个缓存的值
// Put 存储一个键值对到缓存
// Delete 从缓存中删除 key
type Cache interface {
	Context(ctx context.Context) Cache
	Get(key string) (interface{}, time.Time, error)
	Put(key string, val interface{}, d time.Duration) error
	Delete(key string) error
}

// Item 表示存储在 cache 中的一个条目
type Item struct {
	Value      interface{}
	Expiration int64 // 单位纳秒
}

// Expired 表示 item 是否过期
func (i *Item) Expired() bool {
	if i.Expiration == 0 {
		return false
	}

	return time.Now().UnixNano() > i.Expiration
}

// NewCache 返回一个新的 Cache （内存实现）
func NewCache(opts ...Option) Cache {
	options := NewOptions(opts...)
	items := make(map[string]Item)

	if len(options.Items) > 0 {
		items = options.Items
	}

	return &memCache{
		opts:  options,
		items: items,
	}
}
