package just_test

import (
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool(t *testing.T) {
	{
		assert.True(t, just.Bool(true))
		assert.False(t, just.Bool(false))
	}

	{
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
	}

	{
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
	}

	{
		assert.True(t, just.Bool(float32(1)))
		assert.True(t, just.Bool(float64(1)))

		assert.False(t, just.Bool(float32(0)))
		assert.False(t, just.Bool(float64(0)))
	}

	{
		assert.True(t, just.Bool(uintptr(1)))

		assert.False(t, just.Bool(uintptr(0)))
	}

	{
		assert.True(t, just.Bool("1"))

		assert.False(t, just.Bool(""))
	}

}
