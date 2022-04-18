package cache

import "time"

// Options 是 cache 的一些选项
type Options struct {
	Expiration time.Duration   // 过期时间
	Items      map[string]Item // 预设缓存项
}

// Option 修改传入的 Options
type Option func(o *Options)

// Expiration 设置 擦车 中存储的 item 的持续时间设置为过期
func Expiration(d time.Duration) Option {
	return func(o *Options) {
		o.Expiration = d
	}
}

// Items 使用预先配置的项初始化缓存
func Items(i map[string]Item) Option {
	return func(o *Options) {
		o.Items = i
	}
}

// NewOptions 返回一个新的 Options 的结构体
func NewOptions(opts ...Option) Options {
	options := Options{
		Expiration: DefaultExpiration,
		Items:      make(map[string]Item),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}
