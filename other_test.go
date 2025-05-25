package just_test

import (
	"bytes"
	"context"
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"regexp"
	"sync"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	t.Parallel()

	t.Run("bool", func(t *testing.T) {
		assert.True(t, just.Bool(true))
		assert.False(t, just.Bool(false))
	})

	t.Run("int", func(t *testing.T) {
		assert.True(t, just.Bool(int(1)))
		assert.True(t, just.Bool(int8(1)))
		assert.True(t, just.Bool(int16(1)))
		assert.True(t, just.Bool(int32(1)))
		assert.True(t, just.Bool(int64(1)))

		assert.False(t, just.Bool(int(0)))
		assert.False(t, just.Bool(int8(0)))
		assert.False(t, just.Bool(int16(0)))
		assert.False(t, just.Bool(int32(0)))
		assert.False(t, just.Bool(int64(0)))
	})

	t.Run("uint", func(t *testing.T) {
		assert.True(t, just.Bool(uint(1)))
		assert.True(t, just.Bool(uint8(1)))
		assert.True(t, just.Bool(uint16(1)))
		assert.True(t, just.Bool(uint32(1)))
		assert.True(t, just.Bool(uint64(1)))

		assert.False(t, just.Bool(uint(0)))
		assert.False(t, just.Bool(uint8(0)))
		assert.False(t, just.Bool(uint16(0)))
		assert.False(t, just.Bool(uint32(0)))
		assert.False(t, just.Bool(uint64(0)))
	})

	t.Run("float", func(t *testing.T) {
		assert.True(t, just.Bool(float32(1)))
		assert.True(t, just.Bool(float64(1)))

		assert.False(t, just.Bool(float32(0)))
		assert.False(t, just.Bool(float64(0)))
	})

	t.Run("uintptr", func(t *testing.T) {
		assert.True(t, just.Bool(uintptr(1)))

		assert.False(t, just.Bool(uintptr(0)))
	})

	t.Run("string", func(t *testing.T) {
		assert.True(t, just.Bool("1"))

		assert.False(t, just.Bool(""))
	})
}

func TestMust(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		const str = "this is body!"

		val := just.Must(io.ReadAll(bytes.NewBufferString(str)))
		assert.Equal(t, str, string(val))
	})

	t.Run("panic", func(t *testing.T) {
		require.Panics(t, func() {
			just.Must(regexp.Compile("["))
		})
	})
}

func TestRunAfter(t *testing.T) {
	t.Parallel()

	t.Run("runs immediately when runNow is true", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		ticker := make(chan time.Time)

		var called bool
		done := make(chan error, 1)
		go func() {
			done <- just.RunAfter(ctx, ticker, true, func(ctx context.Context) error {
				called = true
				return assert.AnError // Return error to exit the loop
			})
		}()

		err := <-done
		assert.Equal(t, assert.AnError, err)
		assert.True(t, called)
	})

	t.Run("does not run immediately when runNow is false", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		ticker := make(chan time.Time)

		var called bool
		done := make(chan struct{})
		go func() {
			_ = just.RunAfter(ctx, ticker, false, func(ctx context.Context) error {
				called = true
				return nil
			})
			close(done)
		}()

		time.Sleep(10 * time.Millisecond)
		assert.False(t, called)
		cancel()
		<-done
	})

	t.Run("runs on ticker", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		ticker := make(chan time.Time)

		var count int
		var mu sync.Mutex
		done := make(chan struct{})
		go func() {
			_ = just.RunAfter(ctx, ticker, false, func(ctx context.Context) error {
				mu.Lock()
				count++
				mu.Unlock()
				return nil
			})
			close(done)
		}()

		ticker <- time.Now()
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		assert.Equal(t, 1, count)
		mu.Unlock()

		ticker <- time.Now()
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		assert.Equal(t, 2, count)
		mu.Unlock()

		cancel()
		<-done
	})

	t.Run("returns error from function", func(t *testing.T) {
		ctx := context.Background()
		ticker := make(chan time.Time)
		defer close(ticker)

		expectedErr := assert.AnError
		err := just.RunAfter(ctx, ticker, true, func(ctx context.Context) error {
			return expectedErr
		})

		assert.Equal(t, expectedErr, err)
	})

	t.Run("stops on context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		ticker := make(chan time.Time)
		defer close(ticker)

		errChan := make(chan error, 1)
		go func() {
			errChan <- just.RunAfter(ctx, ticker, false, func(ctx context.Context) error {
				return nil
			})
		}()

		cancel()
		err := <-errChan
		assert.Equal(t, context.Canceled, err)
	})
}

func TestIf(t *testing.T) {
	t.Parallel()

	t.Run("returns first value when condition is true", func(t *testing.T) {
		result := just.If(true, "yes", "no")
		assert.Equal(t, "yes", result)
	})

	t.Run("returns second value when condition is false", func(t *testing.T) {
		result := just.If(false, "yes", "no")
		assert.Equal(t, "no", result)
	})

	t.Run("works with different types", func(t *testing.T) {
		// int
		intResult := just.If(true, 42, 0)
		assert.Equal(t, 42, intResult)

		// struct
		type person struct {
			name string
			age  int
		}
		alice := person{name: "Alice", age: 30}
		bob := person{name: "Bob", age: 25}
		personResult := just.If(true, alice, bob)
		assert.Equal(t, alice, personResult)

		// pointers
		x, y := 1, 2
		ptrResult := just.If(false, &x, &y)
		assert.Equal(t, &y, ptrResult)
	})
}
