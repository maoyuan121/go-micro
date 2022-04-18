// Package stats provides runtime stats
package stats

// Stats 提供了一个统计接口
type Stats interface {
	// Read stat snapshot
	Read() ([]*Stat, error)
	// Write a stat snapshot
	Write(*Stat) error
	// Record a request
	Record(error) error
}

// A runtime stat
type Stat struct {
	// Timestamp of recording
	Timestamp int64
	// Start time as unix timestamp
	Started int64
	// Uptime in seconds
	Uptime int64
	// 使用了多少内存，单位 byte
	Memory uint64
	// Threads aka go routines
	Threads uint64
	// Garbage collection in nanoseconds
	GC uint64
	// 总请求数
	Requests uint64
	// 总错误数
	Errors uint64
}

var (
	DefaultStats = NewStats()
)
