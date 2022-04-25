package client

import (
	"context"

	"go-micro.dev/v4/registry"
)

// CallFunc 表示一个 call func
type CallFunc func(ctx context.Context, node *registry.Node, req Request, rsp interface{}, opts CallOptions) error

// CallWrapper 是 CallFunc 的一个底层的包装器
type CallWrapper func(CallFunc) CallFunc

// Wrapper 包装一个 client 返回一个 client
type Wrapper func(Client) Client

// StreamWrapper 包装一个 Stream 返回一个 Stream
type StreamWrapper func(Stream) Stream
