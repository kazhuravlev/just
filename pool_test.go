package just_test

import (
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPool(t *testing.T) {
	t.Run("with_reset", func(t *testing.T) {
		var isCalled bool
		p := just.NewPool(
			func() *[]byte { return just.Pointer(make([]byte, 8)) },
			func(b *[]byte) { isCalled = true },
		)

		assert.False(t, isCalled)

		bb := p.Get()
		assert.Equal(t, 8, len(*bb))
		assert.False(t, isCalled)

		p.Put(bb)
		assert.True(t, isCalled)
	})

	t.Run("without_reset", func(t *testing.T) {
		p := just.NewPool(
			func() *[]byte { return just.Pointer(make([]byte, 8)) },
			nil,
		)

		bb := p.Get()
		assert.Equal(t, 8, len(*bb))

		p.Put(bb)
	})
}

func BenchmarkPoolReset(b *testing.B) {
	// BenchmarkPoolReset-8   	130386861	         8.977 ns/op	       0 B/op	       0 allocs/op
	p := just.NewPool(
		func() *[]byte { return just.Pointer(make([]byte, 0)) },
		func(b *[]byte) { *b = (*b)[:0] },
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bb := p.Get()
		_ = bb
		p.Put(bb)
	}
}
