// metadata 包是定义消息头的一种方式
package metadata

import (
	"context"
	"strings"
)

type metadataKey struct{}

// Metadata 是我们内部表示请求头的方式
// 它们在 RPC 级别使用，并从传输头来回转换
type Metadata map[string]string

func (md Metadata) Get(key string) (string, bool) {
	// attempt to get as is
	val, ok := md[key]
	if ok {
		return val, ok
	}

	// attempt to get lower case
	val, ok = md[strings.Title(key)]
	return val, ok
}

func (md Metadata) Set(key, val string) {
	md[key] = val
}

func (md Metadata) Delete(key string) {
	// delete key as-is
	delete(md, key)
	// delete also Title key
	delete(md, strings.Title(key))
}

// Copy makes a copy of the metadata
func Copy(md Metadata) Metadata {
	cmd := make(Metadata, len(md))
	for k, v := range md {
		cmd[k] = v
	}
	return cmd
}

// Delete key from metadata
func Delete(ctx context.Context, k string) context.Context {
	return Set(ctx, k, "")
}

// Set add key with val to metadata
func Set(ctx context.Context, k, v string) context.Context {
	md, ok := FromContext(ctx)
	if !ok {
		md = make(Metadata)
	}
	if v == "" {
		delete(md, k)
	} else {
		md[k] = v
	}
	return context.WithValue(ctx, metadataKey{}, md)
}

// Get returns a single value from metadata in the context
func Get(ctx context.Context, key string) (string, bool) {
	md, ok := FromContext(ctx)
	if !ok {
		return "", ok
	}
	// attempt to get as is
	val, ok := md[key]
	if ok {
		return val, ok
	}

	// attempt to get lower case
	val, ok = md[strings.Title(key)]

	return val, ok
}

// FromContext returns metadata from the given context
func FromContext(ctx context.Context) (Metadata, bool) {
	md, ok := ctx.Value(metadataKey{}).(Metadata)
	if !ok {
		return nil, ok
	}

	// capitalise all values
	newMD := make(Metadata, len(md))
	for k, v := range md {
		newMD[strings.Title(k)] = v
	}

	return newMD, ok
}

// NewContext creates a new context with the given metadata
func NewContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, metadataKey{}, md)
}

// MergeContext merges metadata to existing metadata, overwriting if specified
func MergeContext(ctx context.Context, patchMd Metadata, overwrite bool) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	md, _ := ctx.Value(metadataKey{}).(Metadata)
	cmd := make(Metadata, len(md))
	for k, v := range md {
		cmd[k] = v
	}
	for k, v := range patchMd {
		if _, ok := cmd[k]; ok && !overwrite {
			// skip
		} else if v != "" {
			cmd[k] = v
		} else {
			delete(cmd, k)
		}
	}
	return context.WithValue(ctx, metadataKey{}, cmd)
}
