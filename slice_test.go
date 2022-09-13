package just_test

import (
	"errors"
	"fmt"
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
)

var less = func(a, b int) bool { return a < b }

func TestUniq(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
				{20, 21},
				{},
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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

	res := just.SliceFillElem(3, "Hello")
	assert.Equal(t, []string{"Hello", "Hello", "Hello"}, res)
}

func TestSliceFindFirstElem(t *testing.T) {
	t.Parallel()

	table := []struct {
		name     string
		in       []int
		elem     int
		exp      int
		expIndex int
	}{
		{
			name:     "empty",
			in:       nil,
			elem:     0,
			exp:      0,
			expIndex: -1,
		},
		{
			name:     "found_index_0",
			in:       []int{1, 1, 1},
			elem:     1,
			exp:      1,
			expIndex: 0,
		},
		{
			name:     "found_index_2",
			in:       []int{3, 2, 1},
			elem:     1,
			exp:      1,
			expIndex: 2,
		},
		{
			name:     "not_found",
			in:       []int{3, 2, 1},
			elem:     42,
			exp:      0,
			expIndex: -1,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceFindFirstElem(row.in, row.elem)
			assert.Equal(t, row.expIndex, res.Idx)
			assert.Equal(t, row.exp, res.Val)
		})
	}
}

func TestSliceFindLastElem(t *testing.T) {
	t.Parallel()

	table := []struct {
		name     string
		in       []int
		elem     int
		exp      int
		expIndex int
	}{
		{
			name:     "empty",
			in:       nil,
			elem:     0,
			exp:      0,
			expIndex: -1,
		},
		{
			name:     "found_index_0",
			in:       []int{1, 2, 3},
			elem:     1,
			exp:      1,
			expIndex: 0,
		},
		{
			name:     "found_index_2",
			in:       []int{3, 2, 1},
			elem:     1,
			exp:      1,
			expIndex: 2,
		},
		{
			name:     "not_found",
			in:       []int{3, 2, 1},
			elem:     42,
			exp:      0,
			expIndex: -1,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceFindLastElem(row.in, row.elem)
			assert.Equal(t, row.expIndex, res.Idx)
			assert.Equal(t, row.exp, res.Val)
		})
	}
}

func TestSliceFindAllElements(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
		{
			name:  "from_0_to_10_by_0",
			start: 0,
			stop:  10,
			step:  0,
			exp:   nil,
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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
		{
			name: "equal3",
			in1:  []int{1, 2, 3},
			in2:  []int{4, 5, 6},
			exp:  false,
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
	t.Parallel()

	res1 := just.SliceChain[int]()
	assert.Equal(t, []int(nil), res1)

	res := just.SliceChain([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)
}

func TestSliceSort(t *testing.T) {
	t.Parallel()

	a := []int{1, 3, 2}
	just.SliceSort(a, less)
	assert.Equal(t, []int{1, 2, 3}, a)
}

func TestSliceElem(t *testing.T) {
	t.Parallel()

	notExists := just.SliceElem[int]{Idx: -1, Val: 0}
	existsFirst := just.SliceElem[int]{Idx: 0, Val: 10}
	existsSecond := just.SliceElem[int]{Idx: 1, Val: 20}

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		assert.False(t, notExists.Ok())
		assert.True(t, existsFirst.Ok())
		assert.True(t, existsSecond.Ok())
	})

	t.Run("value_ok", func(t *testing.T) {
		t.Parallel()

		var v int
		var ok bool

		v, ok = notExists.ValueOk()
		assert.Equal(t, 0, v)
		assert.False(t, ok)

		v, ok = existsFirst.ValueOk()
		assert.Equal(t, 10, v)
		assert.True(t, ok)

		v, ok = existsSecond.ValueOk()
		assert.Equal(t, 20, v)
		assert.True(t, ok)
	})

	t.Run("value_idx", func(t *testing.T) {
		t.Parallel()

		var v int
		var idx int

		v, idx = notExists.ValueIdx()
		assert.Equal(t, 0, v)
		assert.Equal(t, -1, idx)

		v, idx = existsFirst.ValueIdx()
		assert.Equal(t, 10, v)
		assert.Equal(t, 0, idx)

		v, idx = existsSecond.ValueIdx()
		assert.Equal(t, 20, v)
		assert.Equal(t, 1, idx)
	})
}

func TestSliceWithout(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   []int
		elem int
		exp  []int
	}{
		{
			name: "empty",
			in:   nil,
			elem: 0,
			exp:  nil,
		},
		{
			name: "exclude_two",
			in:   []int{1, 2, 3, 4, 5, 6},
			elem: 2,
			exp:  []int{1, 3, 4, 5, 6},
		},
		{
			name: "exclude_not_found",
			in:   []int{1, 2, 3, 4, 5, 6},
			elem: 10000,
			exp:  []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceWithoutElem(row.in, row.elem)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceUnion(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   [][]int
		exp  []int
	}{
		{
			name: "empty",
			in:   nil,
			exp:  []int{},
		},
		{
			name: "case1",
			in: [][]int{
				{1, 2, 3},
			},
			exp: []int{1, 2, 3},
		},
		{
			name: "case2",
			in: [][]int{
				{1, 2, 3},
				{1, 2, 3, 1, 1, 1},
				{3, 4, 5},
				{4, 5, 1, 12},
			},
			exp: []int{1, 2, 3, 4, 5, 12},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceUnion(row.in...)
			assert.Equal(t, just.SliceSortCopy(row.exp, less), just.SliceSortCopy(res, less))
		})
	}
}

func TestSliceAddNotExists(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   []int
		elem int
		exp  []int
	}{
		{
			name: "empty",
			in:   nil,
			elem: 11,
			exp:  []int{11},
		},
		{
			name: "case1",
			in:   []int{1, 1, 1},
			elem: 1,
			exp:  []int{1, 1, 1},
		},
		{
			name: "case2",
			in:   []int{1, 2, 3},
			elem: 4,
			exp:  []int{1, 2, 3, 4},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceAddNotExists(row.in, row.elem)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceContainsElem(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   []int
		elem int
		exp  bool
	}{
		{
			name: "empty",
			in:   nil,
			elem: 11,
			exp:  false,
		},
		{
			name: "case1",
			in:   []int{1, 1, 1},
			elem: 1,
			exp:  true,
		},
		{
			name: "case2",
			in:   []int{1, 2, 3},
			elem: 4,
			exp:  false,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceContainsElem(row.in, row.elem)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceAll(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   []int
		fn   func(int) bool
		exp  bool
	}{
		{
			name: "true_on_empty",
			in:   nil,
			fn:   func(v int) bool { return true },
			exp:  true,
		},
		{
			name: "case1",
			in:   []int{1, 1, 1},
			fn:   func(v int) bool { return v == 1 },
			exp:  true,
		},
		{
			name: "case2",
			in:   []int{1, 2, 3},
			fn:   func(v int) bool { return v == 1 },
			exp:  false,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceAll(row.in, row.fn)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSlice2MapFn(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   []int
		fn   func(int, int) (string, string)
		exp  map[string]string
	}{
		{
			name: "empty",
			in:   nil,
			fn:   func(k, v int) (string, string) { return strconv.Itoa(k), strconv.Itoa(v) },
			exp:  map[string]string{},
		},
		{
			name: "uniq_values",
			in:   []int{10, 20, 30},
			fn:   func(k, v int) (string, string) { return strconv.Itoa(v), strconv.Itoa(k) },
			exp:  map[string]string{"10": "0", "20": "1", "30": "2"},
		},
		{
			name: "non_uniq_values",
			in:   []int{10, 10, 10},
			fn:   func(k, v int) (string, string) { return strconv.Itoa(v), strconv.Itoa(k) },
			exp:  map[string]string{"10": "2"},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.Slice2MapFn(row.in, row.fn)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSlice2MapFnErr(t *testing.T) {
	t.Parallel()

	atoi := func(k int, v string) (int, int, error) {
		x, err := strconv.Atoi(v)
		return k, x, err
	}

	t.Run("error_case", func(t *testing.T) {
		res, err := just.Slice2MapFnErr([]string{"1", "lol", "2"}, atoi)
		require.Error(t, err)
		require.Empty(t, res)
	})

	table := []struct {
		name string
		in   []string
		fn   func(int, string) (int, int, error)
		exp  map[int]int
	}{
		{
			name: "empty",
			in:   nil,
			fn:   atoi,
			exp:  map[int]int{},
		},
		{
			name: "uniq_values",
			in:   []string{"10", "20", "30"},
			fn:   atoi,
			exp: map[int]int{
				0: 10,
				1: 20,
				2: 30,
			},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res, err := just.Slice2MapFnErr(row.in, row.fn)
			assert.Equal(t, row.exp, res)
			assert.NoError(t, err)
		})
	}
}

func TestSliceMap(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   []int
		fn   func(int) string
		exp  []string
	}{
		{
			name: "empty",
			in:   nil,
			fn:   strconv.Itoa,
			exp:  nil,
		},
		{
			name: "case1",
			in:   []int{1, 1, 1},
			fn:   strconv.Itoa,
			exp:  []string{"1", "1", "1"},
		},
		{
			name: "case2",
			in:   []int{1, 2, 3},
			fn:   strconv.Itoa,
			exp:  []string{"1", "2", "3"},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceMap(row.in, row.fn)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceApply(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		var s int
		just.SliceApply([]int{}, func(idx int, v int) { s += v })
		assert.Equal(t, 0, s)
	})

	t.Run("case1", func(t *testing.T) {
		var s int
		just.SliceApply([]int{1, 2, 3}, func(idx int, v int) { s += v })
		assert.Equal(t, 6, s)
	})
}

func TestSliceMapErr(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   []int
		fn   func(int) (string, error)
		exp  []string
		err  bool
	}{
		{
			name: "empty",
			in:   nil,
			fn: func(v int) (string, error) {
				return "", nil
			},
			exp: nil,
			err: false,
		},
		{
			name: "case1",
			in:   []int{1, 2, 3},
			fn: func(v int) (string, error) {
				return strconv.Itoa(v), nil
			},
			exp: []string{"1", "2", "3"},
			err: false,
		},
		{
			name: "case2",
			in:   []int{1, 2, 3},
			fn: func(v int) (string, error) {
				return "", errors.New("the sky is falling")
			},
			exp: nil,
			err: true,
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res, err := just.SliceMapErr(row.in, row.fn)
			if row.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSliceGroupBy(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   []int
		fn   func(int) string
		exp  map[string][]int
	}{
		{
			name: "empty",
			in:   nil,
			fn: func(v int) string {
				return strconv.Itoa(v % 2)
			},
			exp: nil,
		},
		{
			name: "group_odd_even",
			in:   []int{1, 2, 3, 4},
			fn: func(v int) string {
				return strconv.Itoa(v % 2)
			},
			exp: map[string][]int{
				"0": {2, 4},
				"1": {1, 3},
			},
		},
		{
			name: "group_nothing",
			in:   []int{1, 2, 3, 4},
			fn: func(v int) string {
				return strconv.Itoa(v)
			},
			exp: map[string][]int{
				"1": {1},
				"2": {2},
				"3": {3},
				"4": {4},
			},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.SliceGroupBy(row.in, row.fn)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestSlice2Chan(t *testing.T) {
	t.Run("do_not_run_goroutine_on_capacity_is_equal_to_input_len", func(t *testing.T) {
		in := []int{1, 2, 3}
		capacity := len(in)
		ch := just.Slice2Chan(in, capacity)
		require.Equal(t, len(in), len(ch))

		res := just.ChanReadN(ch, len(in))
		require.Equal(t, in, res)
	})

	t.Run("capacity_is_lt_input_len", func(t *testing.T) {
		in := []int{10, 20, 30}
		capacity := 1
		ch := just.Slice2Chan(in, capacity)
		time.Sleep(100 * time.Microsecond)

		// NOTE(zhuravlev): floating tests
		require.Equal(t, capacity, len(ch))

		res := just.ChanReadN(ch, capacity)
		require.Equal(t, []int{10}, res)
	})

	t.Run("capacity_is_gt_input_len", func(t *testing.T) {
		in := []int{10, 20, 30}
		capacity := 100
		ch := just.Slice2Chan(in, capacity)
		time.Sleep(100 * time.Microsecond)

		// NOTE(zhuravlev): floating tests
		require.Equal(t, len(in), len(ch))
		require.Equal(t, capacity, cap(ch))

		res := just.ChanReadN(ch, len(in))
		require.Equal(t, in, res)
	})
}

func TestSlice2ChanFill(t *testing.T) {
	in := []int{1, 2, 3}
	ch := just.Slice2ChanFill(in)
	require.Equal(t, len(in), len(ch))

	res := just.ChanReadN(ch, len(in))
	require.Equal(t, in, res)
}
