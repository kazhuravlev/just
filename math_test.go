package just_test

import (
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	t.Parallel()

	t.Run("empty_should_be_panicked", func(t *testing.T) {
		assert.Panics(t, func() {
			just.Max[int]()
		})
	})

	table := []struct {
		name string
		in   []int
		exp  int
	}{
		{
			name: "case1",
			in:   []int{1},
			exp:  1,
		},
		{
			name: "case2",
			in:   []int{-1, 0, 1},
			exp:  1,
		},
		{
			name: "case3",
			in:   []int{-1, -1, -10},
			exp:  -1,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.Max(row.in...)
			assert.Equal(t, row.exp, res)
		})
	}

}

func TestMaxOr(t *testing.T) {
	t.Parallel()

	table := []struct {
		name       string
		defaultVal int
		in         []int
		exp        int
	}{
		{
			name:       "empty_in",
			defaultVal: -1,
			in:         []int{},
			exp:        -1,
		},
		{
			name:       "case1",
			defaultVal: -1,
			in:         []int{1},
			exp:        1,
		},
		{
			name:       "case2",
			defaultVal: -100,
			in:         []int{-1, 0, 1},
			exp:        1,
		},
		{
			name:       "case3",
			defaultVal: -100,
			in:         []int{-1, -1, -10},
			exp:        -1,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.MaxOr(row.defaultVal, row.in...)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestMaxDefault(t *testing.T) {
	assert.Equal(t, 3, just.MaxDefault(1, 2, 3))
	assert.Equal(t, 0, just.MaxDefault[int]())
}

func TestMin(t *testing.T) {
	t.Parallel()

	t.Run("empty_should_be_panicked", func(t *testing.T) {
		assert.Panics(t, func() {
			just.Min[int]()
		})
	})

	table := []struct {
		name string
		in   []int
		exp  int
	}{
		{
			name: "case1",
			in:   []int{1},
			exp:  1,
		},
		{
			name: "case2",
			in:   []int{-1, 0, 1},
			exp:  -1,
		},
		{
			name: "case3",
			in:   []int{-1, -1, -10},
			exp:  -10,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.Min(row.in...)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestMinOr(t *testing.T) {
	t.Parallel()

	table := []struct {
		name       string
		defaultVal int
		in         []int
		exp        int
	}{
		{
			name:       "empty_in",
			defaultVal: -1,
			in:         []int{},
			exp:        -1,
		},
		{
			name:       "case1",
			defaultVal: -1,
			in:         []int{1},
			exp:        1,
		},
		{
			name:       "case2",
			defaultVal: -100,
			in:         []int{-1, 0, 1},
			exp:        -1,
		},
		{
			name:       "case3",
			defaultVal: -100,
			in:         []int{-1, -1, -10},
			exp:        -10,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.MinOr(row.defaultVal, row.in...)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestMinDefault(t *testing.T) {
	assert.Equal(t, 1, just.MinDefault(1, 2, 3))
	assert.Equal(t, 0, just.MinDefault[int]())
}

func TestSum(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   []int
		exp  int
	}{
		{
			name: "empty",
			in:   []int{},
			exp:  0,
		},
		{
			name: "case1",
			in:   []int{1},
			exp:  1,
		},
		{
			name: "case2",
			in:   []int{1, -1},
			exp:  0,
		},
		{
			name: "case3",
			in:   []int{10, -1, 0},
			exp:  9,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.Sum(row.in...)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestAbs(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, just.Abs(1))
	assert.Equal(t, 1, just.Abs(-1))
	assert.Equal(t, 0, just.Abs(0))

	a := math.Copysign(0, -1)
	assert.Equal(t, float64(0), just.Abs(a))
}
