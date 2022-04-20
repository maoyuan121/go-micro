package signal

import (
	"os"
	"syscall"
)

// ShutDownSingals 返回所有用来关闭服务的信号
func Shutdown() []os.Signal {
	return []os.Signal{
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL,
	}
}
