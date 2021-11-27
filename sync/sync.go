// Package sync 是分布式同步的接口
package sync

import (
	"errors"
	"time"
)

var (
	ErrLockTimeout = errors.New("lock timeout")
)

// Sync 是分布式同步的接口
type Sync interface {
	// Initialise options
	Init(...Option) error
	// Return the options
	Options() Options
	// 选 leader
	Leader(id string, opts ...LeaderOption) (Leader, error)
	// 获取锁
	Lock(id string, opts ...LockOption) error
	// 释放锁
	Unlock(id string) error
	// Sync 实现名
	String() string
}

// Leader provides leadership election
type Leader interface {
	// resign leadership
	Resign() error
	// status returns when leadership is lost
	Status() chan bool
}

type Options struct {
	Nodes  []string
	Prefix string
}

type Option func(o *Options)

type LeaderOptions struct{}

type LeaderOption func(o *LeaderOptions)

type LockOptions struct {
	TTL  time.Duration
	Wait time.Duration
}

type LockOption func(o *LockOptions)
