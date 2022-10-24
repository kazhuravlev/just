package just_test

import (
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointer(t *testing.T) {
	t.Parallel()

	a := just.Pointer(10)
	assert.Equal(t, 10, *a)
}

func TestPointerUnwrap(t *testing.T) {
	t.Parallel()

	n := 10
	a := just.PointerUnwrap(&n)
	assert.Equal(t, n, a)
}
