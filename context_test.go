package just_test

import (
	"context"
	"testing"
	"time"

	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
)

func TestContextWithTimeout(t *testing.T) {
	t.Parallel()

	fn := func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Second):
			return nil
		}
	}

	err := just.ContextWithTimeout(context.Background(), time.Millisecond, fn)
	assert.ErrorIs(t, context.DeadlineExceeded, err)

	fn2 := func(ctx context.Context) error {
		return ctx.Err()
	}

	err2 := just.ContextWithTimeout(context.Background(), time.Millisecond, fn2)
	assert.NoError(t, err2)
}

func TestContextWithTimeout2(t *testing.T) {
	t.Parallel()

	fn := func(ctx context.Context) (int, error) {
		select {
		case <-ctx.Done():
			return 10, ctx.Err()
		case <-time.After(time.Second):
			return 20, nil
		}
	}

	r, err := just.ContextWithTimeout2(context.Background(), time.Millisecond, fn)
	assert.ErrorIs(t, context.DeadlineExceeded, err)
	assert.Equal(t, 10, r)

	fn2 := func(ctx context.Context) (int, error) {
		return 42, ctx.Err()
	}

	r, err2 := just.ContextWithTimeout2(context.Background(), time.Millisecond, fn2)
	assert.NoError(t, err2)

	assert.Equal(t, 42, r)
}
