// Package source is the interface for sources
package source

import (
	"errors"
	"time"
)

var (
	// ErrWatcherStopped is returned when source watcher has been stopped
	ErrWatcherStopped = errors.New("watcher stopped")
)

// Source 是加载配置的源
type Source interface {
	Read() (*ChangeSet, error)
	Write(*ChangeSet) error
	Watch() (Watcher, error)
	String() string
}

// ChangeSet 表示源的修改集
type ChangeSet struct {
	Data      []byte
	Checksum  string
	Format    string
	Source    string
	Timestamp time.Time
}

// Watcher 监视源的更改
type Watcher interface {
	Next() (*ChangeSet, error)
	Stop() error
}
