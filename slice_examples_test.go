package just_test

import (
	"fmt"
	"github.com/kazhuravlev/just"
	"sort"
	"strconv"
)

func ExampleSliceMap() {
	unsignedIntegers := []uint{1, 2, 3, 4}
	multiply := func(i uint) int { return int(i) * 10 }
	integers := just.SliceMap(unsignedIntegers, multiply)
	fmt.Println(integers)
	// Output: [10 20 30 40]
}

func ExampleSliceUniq() {
	data := []int{1, 1, 2, 2, 3, 3}
	uniqData := just.SliceUniq(data)
	sort.Ints(uniqData)
	fmt.Println(uniqData)
	// Output: [1 2 3]
}

func ExampleSliceFlatMap() {
	input := []int{1, 2, 3}
	result := just.SliceFlatMap(input, func(i int) []string {
		return []string{strconv.Itoa(i), strconv.Itoa(i)}
	})
	fmt.Println(result)
	// Output: [1 1 2 2 3 3]
}

func ExampleSliceApply() {
	input := []int{10, 20, 30}
	just.SliceApply(input, func(i int, v int) {
		fmt.Println(i, v)
	})
	// Output:
	// 0 10
	// 1 20
	// 2 30
}

func ExampleSliceFilter() {
	input := []int{10, 20, 30}
	filtered := just.SliceFilter(input, func(v int) bool {
		return v > 15
	})
	fmt.Println(filtered)
	// Output: [20 30]
}

func ExampleSliceContainsElem() {
	input := []int{10, 20, 30}
	contains20 := just.SliceContainsElem(input, 20)
	fmt.Println(contains20)
	// Output: true
}

func ExampleSliceAddNotExists() {
	input := []int{10, 20, 30}
	result := just.SliceAddNotExists(input, 42)
	fmt.Println(result)
	// Output: [10 20 30 42]
}

func ExampleSlice2Map() {
	input := []int{10, 20, 30, 30}
	result := just.Slice2Map(input)
	fmt.Println(result)
	// Output: map[10:{} 20:{} 30:{}]
}

func ExampleSliceChunkEvery() {
	input := []int{10, 20, 30, 40, 50}
	result := just.SliceChunkEvery(input, 2)
	fmt.Println(result)
	// Output: [[10 20] [30 40] [50]]
}

func ExampleSlice2Chan() {
	input := []int{10, 20, 30, 40, 50}
	ch := just.Slice2Chan(input, 0)
	result := just.ChanReadN(ch, 5)
	fmt.Println(result)
	// Output: [10 20 30 40 50]
}
