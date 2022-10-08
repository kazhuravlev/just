package just_test

import (
	"bytes"
	"fmt"
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"regexp"
	"testing"
)

func ExampleMust() {
	val := just.Must(io.ReadAll(bytes.NewBufferString("this is body!")))
	fmt.Println(string(val))
	// Output:
	// this is body!
}

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
