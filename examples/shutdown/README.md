# Shutdown

演示通过 context concellation 优雅的关闭服务

micro.Service 等待  context.Done() 或者 一个 OS kill signal
