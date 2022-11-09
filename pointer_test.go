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

func TestPointerUnwrapDefault(t *testing.T) {
	t.Parallel()

	table := []struct {
		in         *int
		defaultVal int
		exp        int
	}{
		{
			in:         nil,
			defaultVal: 10,
			exp:        10,
		},
		{
			in:         just.Pointer(6),
			defaultVal: 10,
			exp:        6,
		},
	}

	for _, row := range table {
		t.Run("", func(t *testing.T) {
			res := just.PointerUnwrapDefault(row.in, row.defaultVal)
			assert.Equal(t, row.exp, res)
		})
	}

	type Data struct {
		Value int
	}

	res := just.PointerUnwrapDefault((*Data)(nil), Data{42})
	assert.Equal(t, Data{42}, res)
}
