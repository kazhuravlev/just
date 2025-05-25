//go:build go1.23

package just_test

import (
	"iter"
	"testing"

	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceIter(t *testing.T) {
	t.Parallel()

	f := func(next func() (just.IterContext, int, bool), val, idx, revIdx int, isFirst, isLast bool) {
		t.Helper()

		iterCtx, elem, valid := next()
		require.True(t, valid)
		assert.Equal(t, val, elem)
		assert.Equal(t, idx, iterCtx.Idx())
		assert.Equal(t, revIdx, iterCtx.RevIdx())
		assert.Equal(t, isFirst, iterCtx.IsFirst())
		assert.Equal(t, isLast, iterCtx.IsLast())
	}

	in := []int{10, 20, 30, 40}
	iterator := just.SliceIter(in)
	next, _ := iter.Pull2(iterator)

	f(next, 10, 0, 3, true, false)
	f(next, 20, 1, 2, false, false)
	f(next, 30, 2, 1, false, false)
	f(next, 40, 3, 0, false, true)
}
