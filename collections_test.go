package just_test

import (
	"github.com/kazhuravlev/just"
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

			res := just.Uniq(row.in)
			sort.SliceStable(res, func(i, j int) bool {
				return res[i] < res[j]
			})
			require.EqualValues(t, row.exp, res)
		})
	}
}

func TestMergeMap(t *testing.T) {
	alwaysTen := func(...int) int { return 10 }

	table := []struct {
		name string
		m1   map[int]int
		m2   map[int]int
		fn   func(...int) int
		exp  map[int]int
	}{
		{
			name: "empty_nil",
			m1:   nil,
			m2:   nil,
			fn:   alwaysTen,
			exp:  map[int]int{},
		},
		{
			name: "empty_len0",
			m1:   map[int]int{},
			m2:   map[int]int{},
			fn:   alwaysTen,
			exp:  map[int]int{},
		},
		{
			name: "merge_all_keys",
			m1:   map[int]int{1: 1},
			m2:   map[int]int{2: 2},
			fn:   alwaysTen,
			exp:  map[int]int{1: 10, 2: 10},
		},
		{
			name: "merge_all_keys_duplicated",
			m1:   map[int]int{1: 1, 2: 2},
			m2:   map[int]int{2: 2, 1: 1},
			fn:   alwaysTen,
			exp:  map[int]int{1: 10, 2: 10},
		},
		{
			name: "merge_all_keys_m1_empty",
			m1:   map[int]int{},
			m2:   map[int]int{2: 2, 1: 1},
			fn:   alwaysTen,
			exp:  map[int]int{1: 10, 2: 10},
		},
		{
			name: "merge_all_keys_m2_empty",
			m1:   map[int]int{2: 2, 1: 1},
			m2:   map[int]int{},
			fn:   alwaysTen,
			exp:  map[int]int{1: 10, 2: 10},
		},
		{
			name: "merge_all_keys_get_biggest",
			m1:   map[int]int{1: 10, 2: 11},
			m2:   map[int]int{1: 11, 2: 10},
			fn:   just.Max[int],
			exp:  map[int]int{1: 11, 2: 11},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.MergeMap(row.m1, row.m2, func(a, b int) int { return row.fn(a, b) })
			require.EqualValues(t, row.exp, res)
		})
	}
}

func TestFilterMap(t *testing.T) {
	alwaysTrue := func(_, _ int) bool { return true }
	alwaysFalse := func(_, _ int) bool { return false }

	table := []struct {
		name string
		m    map[int]int
		fn   func(int, int) bool
		exp  map[int]int
	}{
		{
			name: "empty_nil",
			m:    nil,
			fn:   alwaysTrue,
			exp:  map[int]int{},
		},
		{
			name: "empty_len0",
			m:    map[int]int{},
			fn:   alwaysTrue,
			exp:  map[int]int{},
		},
		{
			name: "should_copy_all_kv",
			m:    map[int]int{1: 1, 2: 2},
			fn:   alwaysTrue,
			exp:  map[int]int{1: 1, 2: 2},
		},
		{
			name: "should_ignore_all_kv",
			m:    map[int]int{1: 1, 2: 2},
			fn:   alwaysFalse,
			exp:  map[int]int{},
		},
		{
			name: "keep_only_values_gt_10",
			m:    map[int]int{1: 10, 2: 2, 3: 100, 4: -1},
			fn: func(_, v int) bool {
				return v > 10
			},
			exp: map[int]int{3: 100},
		},
		{
			name: "keep_only_even_keys",
			m:    map[int]int{1: 10, 2: 2, 3: 100, 4: -1},
			fn: func(k, _ int) bool {
				return k%2 == 0
			},
			exp: map[int]int{2: 2, 4: -1},
		},
	}

	for _, row := range table {
		row := row
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			res := just.FilterMap(row.m, row.fn)
			require.EqualValues(t, row.exp, res)
		})
	}
}
