package logger

import "context"

type loggerKey struct{}

func FromContext(ctx context.Context) (Logger, bool) {
	l, ok := ctx.Value(loggerKey{}).(Logger)
	return l, ok
}

func NewContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, l)
}

type logger1Key struct {
}

func FromContext1(ctx context.Context) (Logger, bool) {
	l, ok := ctx.Value(logger1Key{}).(Logger)
	return l, ok
}

func NewContext1(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, logger1Key{}, l)
}
