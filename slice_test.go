package just_test

import (
	"fmt"
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var less = func(a, b int) bool { return a < b }

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
			require.EqualValues(t, row.exp, just.SliceSortCopy(res, func(a, b int) bool { return a < b }))
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

func TestSliceChunkEvery(t *testing.T) {
	table := []struct {
		name  string
		in    []int
		every int
		exp   [][]int
	}{
		{
			name:  "empty",
			in:    nil,
			every: 1,
			exp:   nil,
		},
		{
			name:  "split_every_1",
			in:    []int{1, 2, 3, 4},
			every: 1,
			exp: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			name:  "split_every_2",
			in:    []int{1, 2, 3, 4},
			every: 2,
			exp: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name:  "split_every_minus_2",
			in:    []int{1, 2, 3, 4},
			every: -2,
			exp: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name:  "split_every_3",
			in:    []int{1, 2, 3, 4},
			every: 3,
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

			res := just.SliceChunkEvery(row.in, row.every)
			assert.Equal(t, row.exp, res)
		})
	}

	t.Run("split_every_0_invalid", func(t *testing.T) {
		assert.Panics(t, func() {
			just.SliceChunkEvery([]int{1, 2, 3, 4}, 0)
		})
	})
}

func TestSliceFillElem(t *testing.T) {
	res := just.SliceFillElem(3, "Hello")
	assert.Equal(t, []string{"Hello", "Hello", "Hello"}, res)
}

func TestSliceFindFirst(t *testing.T) {
	table := []struct {
		name     string
		in       []int
		fn       func(int, int) bool
		exp      int
		expIndex int
	}{
		{
			name:     "empty",
			in:       nil,
			fn:       nil,
			exp:      0,
			expIndex: -1,
		},
		{
			name: "found_index_0",
			in:   []int{1, 1, 1},
			fn: func(i int, v int) bool {
				return v == 1
			},
			exp:      1,
			expIndex: 0,
		},
		{
			name: "found_index_2",
			in:   []int{3, 2, 1},
			fn: func(i int, v int) bool {
				return v == 1
			},
			exp:      1,
			expIndex: 2,
		},
		{
			name: "not_found",
			in:   []int{3, 2, 1},
			fn: func(i int, v int) bool {
				return v == 42
			},
			exp:      0,
			expIndex: -1,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceFindFirst(row.in, row.fn)
			assert.Equal(t, row.expIndex, res.Idx)
			assert.Equal(t, row.exp, res.Val)
		})
	}
}

func TestSliceFindLast(t *testing.T) {
	table := []struct {
		name     string
		in       []int
		fn       func(int, int) bool
		exp      int
		expIndex int
	}{
		{
			name:     "empty",
			in:       nil,
			fn:       nil,
			exp:      0,
			expIndex: -1,
		},
		{
			name: "found_index_0",
			in:   []int{1, 2, 3},
			fn: func(i int, v int) bool {
				return v == 1
			},
			exp:      1,
			expIndex: 0,
		},
		{
			name: "found_index_2",
			in:   []int{3, 2, 1},
			fn: func(i int, v int) bool {
				return v == 1
			},
			exp:      1,
			expIndex: 2,
		},
		{
			name: "not_found",
			in:   []int{3, 2, 1},
			fn: func(i int, v int) bool {
				return v == 42
			},
			exp:      0,
			expIndex: -1,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceFindLast(row.in, row.fn)
			assert.Equal(t, row.expIndex, res.Idx)
			assert.Equal(t, row.exp, res.Val)
		})
	}
}

func TestSliceFindAllElements(t *testing.T) {
	table := []struct {
		name string
		in   []int
		fn   func(int, int) bool
		exp  []int
	}{
		{
			name: "empty",
			in:   nil,
			fn:   nil,
			exp:  []int{},
		},
		{
			name: "found_gte_2",
			in:   []int{1, 2, 3},
			fn: func(i int, v int) bool {
				return v >= 2
			},
			exp: []int{2, 3},
		},
		{
			name: "not_found",
			in:   []int{3, 2, 1},
			fn: func(i int, v int) bool {
				return v == 42
			},
			exp: []int{},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceFindAllElements(row.in, row.fn)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceFindAllIndexes(t *testing.T) {
	table := []struct {
		name string
		in   []int
		fn   func(int, int) bool
		exp  []int
	}{
		{
			name: "empty",
			in:   nil,
			fn:   nil,
			exp:  []int{},
		},
		{
			name: "found_gte_20",
			in:   []int{11, 21, 31},
			fn: func(i int, v int) bool {
				return v >= 20
			},
			exp: []int{1, 2},
		},
		{
			name: "not_found",
			in:   []int{3, 2, 1},
			fn: func(i int, v int) bool {
				return v == 42
			},
			exp: []int{},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceFindAllIndexes(row.in, row.fn)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceFindAll(t *testing.T) {
	table := []struct {
		name string
		in   []int
		fn   func(int, int) bool
		exp  []just.SliceElem[int]
	}{
		{
			name: "empty",
			in:   nil,
			fn:   nil,
			exp:  []just.SliceElem[int]{},
		},
		{
			name: "found_gte_20",
			in:   []int{11, 21, 31},
			fn: func(i int, v int) bool {
				return v >= 20
			},
			exp: []just.SliceElem[int]{
				{Idx: 1, Val: 21},
				{Idx: 2, Val: 31},
			},
		},
		{
			name: "not_found",
			in:   []int{3, 2, 1},
			fn: func(i int, v int) bool {
				return v == 42
			},
			exp: []just.SliceElem[int]{},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceFindAll(row.in, row.fn)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceRange(t *testing.T) {
	table := []struct {
		name              string
		start, stop, step int
		exp               []int
	}{
		{
			name:  "from_zero_to_zero",
			start: 0,
			stop:  0,
			step:  0,
			exp:   nil,
		},
		{
			name:  "from_0_to_5_by_2",
			start: 0,
			stop:  5,
			step:  2,
			exp:   []int{0, 2, 4},
		},
		{
			name:  "from_0_to_5_by_minus_2",
			start: 0,
			stop:  5,
			step:  -2,
			exp:   nil,
		},
		{
			name:  "from_5_to_0_by_2",
			start: 5,
			stop:  0,
			step:  2,
			exp:   nil,
		},
		{
			name:  "from_minus_5_to_minus_1_by_2",
			start: -5,
			stop:  -1,
			step:  2,
			exp:   []int{-5, -3},
		},
		{
			name:  "from_0_to_minus_10_by_minus_10",
			start: 0,
			stop:  -10,
			step:  -10,
			exp:   []int{0},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceRange(row.start, row.stop, row.step)
			fmt.Println(row.exp, res)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceDifference(t *testing.T) {
	table := []struct {
		name     string
		in1, in2 []int
		exp      []int
	}{
		{
			name: "empty_both",
			in1:  nil,
			in2:  nil,
			exp:  nil,
		},
		{
			name: "empty_first",
			in1:  nil,
			in2:  []int{1, 2, 3},
			exp:  []int{1, 2, 3},
		},
		{
			name: "empty_second",
			in1:  []int{1, 2, 3},
			in2:  nil,
			exp:  nil,
		},
		{
			name: "equal",
			in1:  []int{1, 2, 3},
			in2:  []int{1, 2, 3},
			exp:  []int{},
		},
		{
			name: "has_diff_1",
			in1:  []int{1, 2, 3},
			in2:  []int{1, 2, 3, 4},
			exp:  []int{4},
		},
		{
			name: "has_diff_2",
			in1:  []int{1, 2},
			in2:  []int{1, 2, 3, 4},
			exp:  []int{3, 4},
		},
		{
			name: "has_diff_3_duplicated",
			in1:  []int{1, 2},
			in2:  []int{2, 4, 4, 2, 2, 4},
			exp:  []int{4},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceDifference(row.in1, row.in2)
			assert.Equal(t, just.SliceSortCopy(row.exp, less), just.SliceSortCopy(res, less))
		})
	}
}

func TestSliceIntersection(t *testing.T) {
	table := []struct {
		name     string
		in1, in2 []int
		exp      []int
	}{
		{
			name: "empty_both",
			in1:  nil,
			in2:  nil,
			exp:  nil,
		},
		{
			name: "empty_first",
			in1:  nil,
			in2:  []int{1, 2, 3},
			exp:  nil,
		},
		{
			name: "empty_second",
			in1:  []int{1, 2, 3},
			in2:  nil,
			exp:  nil,
		},
		{
			name: "equal",
			in1:  []int{1, 2, 3},
			in2:  []int{1, 2, 3},
			exp:  []int{1, 2, 3},
		},
		{
			name: "has_diff_1",
			in1:  []int{1, 2, 3},
			in2:  []int{1, 2, 3, 4},
			exp:  []int{1, 2, 3},
		},
		{
			name: "has_diff_2",
			in1:  []int{1, 2},
			in2:  []int{1, 2, 3, 4},
			exp:  []int{1, 2},
		},
		{
			name: "has_diff_3_duplicated",
			in1:  []int{1, 2},
			in2:  []int{2, 4, 4, 2, 2, 4, 1},
			exp:  []int{1, 2},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceIntersection(row.in1, row.in2)
			assert.Equal(t, just.SliceSortCopy(row.exp, less), just.SliceSortCopy(res, less))
		})
	}
}

func TestSliceEqualUnordered(t *testing.T) {
	table := []struct {
		name     string
		in1, in2 []int
		exp      bool
	}{
		{
			name: "empty",
			in1:  nil,
			in2:  nil,
			exp:  true,
		},
		{
			name: "empty_first",
			in1:  nil,
			in2:  []int{1},
			exp:  false,
		},
		{
			name: "empty_second",
			in1:  []int{1},
			in2:  nil,
			exp:  false,
		},
		{
			name: "equal_full",
			in1:  []int{1},
			in2:  []int{1},
			exp:  true,
		},
		{
			name: "equal_dupl",
			in1:  []int{1, 1, 1},
			in2:  []int{1},
			exp:  true,
		},
		{
			name: "equal_dupl2",
			in1:  []int{1, 1},
			in2:  []int{1, 1, 1, 1},
			exp:  true,
		},
		{
			name: "equal2",
			in1:  []int{1, 1, 2, 3, 2, 1},
			in2:  []int{1, 2, 3, 3},
			exp:  true,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceEqualUnordered(row.in1, row.in2)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceChain(t *testing.T) {
	res := just.SliceChain([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)
}
