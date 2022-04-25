package sync

import (
	"time"
)

// Nodes 设置要用的地址
func Nodes(a ...string) Option {
	return func(o *Options) {
		o.Nodes = a
	}
}

// Prefix 为使用的任何锁 id 设置前缀
func Prefix(p string) Option {
	return func(o *Options) {
		o.Prefix = p
	}
}

// LockTTL 设置锁 ttl
func LockTTL(t time.Duration) LockOption {
	return func(o *LockOptions) {
		o.TTL = t
	}
}

// LockWait 设置 wait 时间
func LockWait(t time.Duration) LockOption {
	return func(o *LockOptions) {
		o.Wait = t
	}
}
