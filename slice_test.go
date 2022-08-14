package just_test

import (
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestUniq(t *testing.T) {
	table := []struct {
		name string
		in   []int
		exp  []int
	}{
		{
			name: "empty_nil",
			in:   nil,
			exp:  []int{},
		},
		{
			name: "empty_len0",
			in:   []int{},
			exp:  []int{},
		},
		{
			name: "uniq_1",
			in:   []int{1},
			exp:  []int{1},
		},
		{
			name: "uniq_3",
			in:   []int{1, 2, 3},
			exp:  []int{1, 2, 3},
		},
		{
			name: "non_uniq_3",
			in:   []int{1, 1, 1},
			exp:  []int{1},
		},
		{
			name: "non_uniq_6_unordered",
			in:   []int{1, 2, 1, 3, 1, 4},
			exp:  []int{1, 2, 3, 4},
		},
		{
			name: "non_uniq_100",
			in:   make([]int, 100),
			exp:  []int{0},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceUniq(row.in)
			sort.SliceStable(res, func(i, j int) bool {
				return res[i] < res[j]
			})
			require.EqualValues(t, row.exp, res)
		})
	}
}

func TestSliceReverse(t *testing.T) {
	table := []struct {
		name string
		in   []int
		exp  []int
	}{
		{
			name: "empty_nil",
			in:   nil,
			exp:  []int{},
		},
		{
			name: "empty_len0",
			in:   []int{},
			exp:  []int{},
		},
		{
			name: "one_element",
			in:   []int{1},
			exp:  []int{1},
		},
		{
			name: "three_elements",
			in:   []int{1, 2, 3},
			exp:  []int{3, 2, 1},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceReverse(row.in)
			require.Equal(t, row.exp, res)
		})
	}
}

func TestSliceZip(t *testing.T) {
	table := []struct {
		name string
		in   [][]int
		exp  [][]int
	}{
		{
			name: "empty",
			in:   nil,
			exp:  nil,
		},
		{
			name: "empty_len0",
			in:   [][]int{},
			exp:  nil,
		},
		{
			name: "one_slice_in_args",
			in: [][]int{
				{1, 2, 3},
			},
			exp: [][]int{
				{1},
				{2},
				{3},
			},
		},
		{
			name: "two_slice_in_args",
			in: [][]int{
				{10, 11, 12},
				{20, 21, 22},
			},
			exp: [][]int{
				{10, 20},
				{11, 21},
				{12, 22},
			},
		},
		{
			name: "three_slices_diff_len",
			in: [][]int{
				{10},
				{20, 21},
				{30, 31, 32},
			},
			exp: [][]int{
				{10, 20, 30},
			},
		},
		{
			name: "two_slices_one_empty",
			in: [][]int{
				{},
				{20, 21},
			},
			exp: nil,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceZip(row.in...)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceChunk(t *testing.T) {
	table := []struct {
		name string
		in   []int
		fn   func(int, int) bool
		exp  [][]int
	}{
		{
			name: "empty",
			in:   nil,
			fn:   nil,
			exp:  nil,
		},
		{
			name: "split_fn_always_true",
			in:   []int{1, 2, 3, 4},
			fn:   func(i int, v int) bool { return true },
			exp: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			name: "split_fn_always_false",
			in:   []int{1, 2, 3, 4},
			fn:   func(i int, v int) bool { return false },
			exp: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			name: "split_every_2",
			in:   []int{1, 2, 3, 4},
			fn:   func(i int, v int) bool { return i%2 == 0 },
			exp: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "split_every_3",
			in:   []int{1, 2, 3, 4},
			fn:   func(i int, v int) bool { return i%3 == 0 },
			exp: [][]int{
				{1, 2, 3},
				{4},
			},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceChunk(row.in, row.fn)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceFillElem(t *testing.T) {
	res := just.SliceFillElem(3, "Hello")
	assert.Equal(t, []string{"Hello", "Hello", "Hello"}, res)
}
