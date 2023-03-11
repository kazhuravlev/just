package just

import (
	"context"
	"time"
)

// ContextWithTimeout will create a new context with a specified timeout and
// call the function with this context.
func ContextWithTimeout(ctx context.Context, d time.Duration, fn func(context.Context) error) error {
	ctx2, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	return fn(ctx2)
}

// ContextWithTimeout2 will do the same as ContextWithTimeout but returns
// 2 arguments from the function callback.
func ContextWithTimeout2[T any](ctx context.Context, d time.Duration, fn func(context.Context) (T, error)) (T, error) {
	ctx2, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	return fn(ctx2)
}
