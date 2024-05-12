package just

import (
	"math/rand"
	"sort"
)

// SliceUniq returns unique values from `in`.
func SliceUniq[T comparable](in []T) []T {
	index := Slice2Map(in)

	res := make([]T, 0, len(index))
	for k := range index {
		res = append(res, k)
	}

	return res
}

// SliceUniqStable returns unique values from `in`. Keep original ordering.
func SliceUniqStable[T comparable](in []T) []T {
	index := make(map[T]struct{}, len(in))

	res := make([]T, 0, len(in))
	for i := range in {
		if MapContainsKey(index, in[i]) {
			continue
		}

		index[in[i]] = struct{}{}

		res = append(res, in[i])
	}

	return res
}

// SliceMap returns the slice where each element of `in` was handled by `fn`.
func SliceMap[T any, V any](in []T, fn func(T) V) []V {
	if len(in) == 0 {
		return make([]V, 0)
	}

	res := make([]V, len(in))
	for i := range in {
		res[i] = fn(in[i])
	}

	return res
}

// SliceFlatMap applies `fn` to each element of `in` and join all output slices.
func SliceFlatMap[T, V any](in []T, fn func(val T) []V) []V {
	if len(in) == 0 {
		return make([]V, 0)
	}

	res := make([]V, 0, len(in))
	for i := range in {
		slice := fn(in[i])
		res = append(res, slice...)
	}

	return res
}

// SliceFlatMap2 does the same as SliceFlatMap but receives the index
// of the element.
func SliceFlatMap2[T, V any](in []T, fn func(i int, val T) []V) []V {
	if len(in) == 0 {
		return make([]V, 0)
	}

	var res []V
	for i := range in {
		slice := fn(i, in[i])
		res = append(res, slice...)
	}

	return res
}

// SliceApply handles all elements from `in` by function `fn`. Function
// applies sequentially.
func SliceApply[T any](in []T, fn func(int, T)) {
	if len(in) == 0 {
		return
	}

	for i := range in {
		fn(i, in[i])
	}
}

// SliceMapErr does the same thing as SliceMap but returns an error when
// an error occurs in fn.
func SliceMapErr[T any, V any](in []T, fn func(T) (V, error)) ([]V, error) {
	if len(in) == 0 {
		return make([]V, 0), nil
	}

	res := make([]V, len(in))
	var err error
	for i := range in {
		res[i], err = fn(in[i])
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

// SliceFilter returns a slice of values from `in` where `fn(elem) == true`.
func SliceFilter[T any](in []T, fn func(T) bool) []T {
	if len(in) == 0 {
		return make([]T, 0)
	}

	res := make([]T, 0, len(in))
	for i := range in {
		if !fn(in[i]) {
			continue
		}

		res = append(res, in[i])
	}

	return res
}

// SliceReverse reverse the slice.
func SliceReverse[T any](in []T) []T {
	if len(in) == 0 {
		return make([]T, 0)
	}

	res := make([]T, len(in))
	for i := range in {
		res[i] = in[len(in)-i-1]
	}

	return res
}

// SliceAny returns true when `fn` returns true for at least one element
// from `in`.
func SliceAny[T any](in []T, fn func(T) bool) bool {
	for i := range in {
		if fn(in[i]) {
			return true
		}
	}

	return false
}

// SliceAll returns true when `fn` returns true for all elements from `in`.
// Returns true when in is empty.
func SliceAll[T any](in []T, fn func(T) bool) bool {
	for i := range in {
		if !fn(in[i]) {
			return false
		}
	}

	return true
}

// SliceContainsElem returns true when `in` contains elem.
func SliceContainsElem[T comparable](in []T, elem T) bool {
	return SliceAny(in, func(v T) bool { return v == elem })
}

// SliceAddNotExists return `in` with `elem` inside when `elem` not exists in
// `in`.
func SliceAddNotExists[T comparable](in []T, elem T) []T {
	for i := range in {
		if in[i] == elem {
			return in
		}
	}

	return append(in, elem)
}

// SliceUnion returns only uniq items from all slices.
func SliceUnion[T comparable](in ...[]T) []T {
	var res []T
	for i := range in {
		res = append(res, in[i]...)
	}

	return SliceUniq[T](res)
}

// Slice2Map make map from slice, which contains all values from `in` as map
// keys.
func Slice2Map[T comparable](in []T) map[T]struct{} {
	res := make(map[T]struct{}, len(in))
	for i := range in {
		res[in[i]] = struct{}{}
	}

	return res
}

// SliceDifference returns the difference between `oldSlice` and `newSlice`.
// Returns only elements presented in `newSlice` but not presented
// in `oldSlice`.
// Example: [1,2,3], [3,4,5,5,5] => [4,5,5,5]
func SliceDifference[T comparable](oldSlice, newSlice []T) []T {
	if len(oldSlice) == 0 {
		return newSlice
	}

	if len(newSlice) == 0 {
		return make([]T, 0)
	}

	index := Slice2Map(oldSlice)
	res := make([]T, 0, len(newSlice))
	for i := range newSlice {
		if _, ok := index[newSlice[i]]; ok {
			continue
		}

		res = append(res, newSlice[i])
	}

	return SliceUniq(res)
}

// SliceIntersection returns elements that are presented in both slices.
// Example: [1,2,3], [2,4,3,3,3] => [2, 3]
func SliceIntersection[T comparable](oldSlice, newSlice []T) []T {
	if len(oldSlice) == 0 {
		return make([]T, 0)
	}

	if len(newSlice) == 0 {
		return make([]T, 0)
	}

	index := Slice2Map(oldSlice)
	res := make([]T, 0, len(newSlice))
	for i := range newSlice {
		if _, ok := index[newSlice[i]]; !ok {
			continue
		}

		res = append(res, newSlice[i])
	}

	return SliceUniq(res)
}

// SliceWithoutElem returns the slice `in` that not contains `elem`.
func SliceWithoutElem[T comparable](in []T, elem T) []T {
	return SliceWithout(in, func(v T) bool {
		return v == elem
	})
}

// SliceWithout returns the slice `in` where fn(elem) == true.
func SliceWithout[T any](in []T, fn func(T) bool) []T {
	return SliceFilter(in, func(elem T) bool {
		return !fn(elem)
	})
}

// SliceZip returns merged together the values of each of the arrays with the
// values at the corresponding position. If the len of `in` is different - will
// use smaller one.
func SliceZip[T any](in ...[]T) [][]T {
	if len(in) == 0 {
		return make([][]T, 0)
	}

	maxLen := len(in[0])
	for i := range in {
		if len(in[i]) < maxLen {
			maxLen = len(in[i])
		}
	}

	if maxLen == 0 {
		return make([][]T, 0)
	}

	res := make([][]T, maxLen)
	for i := 0; i < maxLen; i++ {
		row := make([]T, len(in))
		for j := range in {
			row[j] = in[j][i]
		}

		res[i] = row
	}

	return res
}

// SliceFillElem returns the slice with len `l` where all elements are equal to
// `elem`.
func SliceFillElem[T any](l int, elem T) []T {
	res := make([]T, l)
	for i := 0; i < l; i++ {
		res[i] = elem
	}

	return res
}

// SliceNotNil return source slice when it is not nil or create empty
// instance of this type.
func SliceNotNil[T any](in []T) []T {
	if in == nil {
		return make([]T, 0)
	}

	return in
}

// SliceChunk split `in` into chunks by fn(index, elem) == true.
func SliceChunk[T any](in []T, fn func(i int, elem T) bool) [][]T {
	if len(in) == 0 {
		return make([][]T, 0)
	}

	res := make([][]T, 0, len(in))
	var chunk []T
	for i := range in {
		if fn(i, in[i]) && len(chunk) != 0 {
			res = append(res, chunk)
			chunk = make([]T, 0)
		}

		chunk = append(chunk, in[i])
	}

	if len(chunk) != 0 {
		res = append(res, chunk)
	}

	return res
}

// SliceChunkEvery split `in` into chunks by size `every`
func SliceChunkEvery[T any](in []T, every int) [][]T {
	if every == 0 {
		panic("invalid arg")
	}

	return SliceChunk(in, func(i int, elem T) bool {
		return i%every == 0
	})
}

// SliceElem represent element of slice.
type SliceElem[T any] struct {
	// Idx is index of element in slice.
	Idx int
	// Val is value on slice by Idx index.
	Val T
}

// ValueOk returns the value and true (when element is exists in slice) or false in other case.
// Useful for cases like:
//
//	if elem, ok := SliceFindFirstElem([]int{1,2,3}, 2); ok{
//		fmt.Println(elem)
//	}
func (e SliceElem[T]) ValueOk() (T, bool) {
	return e.Val, e.Idx != -1
}

// Ok returns true if Idx is valid.
func (e SliceElem[T]) Ok() bool {
	return e.Idx != -1
}

// ValueIdx returns value and index as is.
// Useful for this:
//
//	elem, idx := SliceFindFirstElem([]int{1,2,3}, 2).ValueIdx()
func (e SliceElem[T]) ValueIdx() (T, int) {
	return e.Val, e.Idx
}

// SliceFindFirst return first elem from `in` that fn(index, elem) == true.
// returns index of found elem or -1 if elem not found.
func SliceFindFirst[T any](in []T, fn func(i int, elem T) bool) SliceElem[T] {
	for i := range in {
		if fn(i, in[i]) {
			return SliceElem[T]{
				Idx: i,
				Val: in[i],
			}
		}
	}

	return SliceElem[T]{
		Idx: -1,
	}
}

// SliceFindFirstElem return first elem from `in` that equals to `elem`.
func SliceFindFirstElem[T comparable](in []T, elem T) SliceElem[T] {
	return SliceFindFirst(in, func(_ int, e T) bool {
		return e == elem
	})
}

// SliceFindLast return last elem from `in` that fn(index, elem) == true.
// returns index of found elem or -1 if elem not found.
func SliceFindLast[T any](in []T, fn func(i int, elem T) bool) SliceElem[T] {
	for i := len(in) - 1; i != -1; i-- {
		if fn(i, in[i]) {
			return SliceElem[T]{
				Idx: i,
				Val: in[i],
			}
		}
	}

	return SliceElem[T]{
		Idx: -1,
	}
}

// SliceFindLastElem return last elem from `in` that equals to `elem`.
func SliceFindLastElem[T comparable](in []T, elem T) SliceElem[T] {
	return SliceFindLast(in, func(_ int, e T) bool {
		return e == elem
	})
}

// SliceFindAll return all elem and index from `in` that fn(index, elem) == true.
func SliceFindAll[T any](in []T, fn func(i int, elem T) bool) []SliceElem[T] {
	res := make([]SliceElem[T], 0, len(in))
	for i := range in {
		if !fn(i, in[i]) {
			continue
		}

		res = append(res, SliceElem[T]{
			Idx: i,
			Val: in[i],
		})
	}

	return res
}

// SliceFindAllElements return all elem from `in` that fn(index, elem) == true.
func SliceFindAllElements[T any](in []T, fn func(i int, elem T) bool) []T {
	res := make([]T, 0, len(in))
	for i := range in {
		if !fn(i, in[i]) {
			continue
		}

		res = append(res, in[i])
	}

	return res
}

// SliceFindAllIndexes return all indexes from `in` that fn(index, elem) == true.
func SliceFindAllIndexes[T any](in []T, fn func(i int, elem T) bool) []int {
	res := make([]int, 0, len(in))
	for i := range in {
		if !fn(i, in[i]) {
			continue
		}

		res = append(res, i)
	}

	return res
}

// SliceRange produces a sequence of integers from start (inclusive)
// to stop (exclusive) by step.
func SliceRange[T number](start, stop, step T) []T {
	if start == stop {
		return make([]T, 0)
	}

	if step == 0 {
		return make([]T, 0)
	}

	isIncr := start < stop
	if isIncr && step < 0 {
		return make([]T, 0)
	}

	if !isIncr && step > 0 {
		return make([]T, 0)
	}

	res := make([]T, 0, int(Abs((start-stop)/step)))
	e := start
	for {
		if isIncr && e >= stop {
			break
		}

		if !isIncr && e <= stop {
			break
		}

		res = append(res, e)
		e += step
	}

	return res
}

// SliceEqualUnordered returns true when all uniq values from `in1` contains in `in2`.
// Useful in tests for comparing expected and actual slices.
// Examples:
//   - [1,2,3], [2,3,3,3,1,1] => true
//   - [1], [1,1,1] => true
//   - [1], [1] => true
//   - [1], [2] => false
func SliceEqualUnordered[T comparable](in1, in2 []T) bool {
	m1 := Slice2Map(in1)
	m2 := Slice2Map(in2)

	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
	}

	return true
}

// SliceChain returns a slice where all `in` slices id appended to the end. Like
// append(append(in[0], in[1]...), in[2]...).
func SliceChain[T any](in ...[]T) []T {
	if len(in) == 0 {
		return make([]T, 0)
	}

	var l int
	for i := range in {
		l += len(in[i])
	}

	res := make([]T, l)
	var x int
	for i := range in {
		copy(res[x:x+len(in[i])], in[i])
		x += len(in[i])
	}

	return res
}

// SliceSort sort slice inplace.
func SliceSort[T any](in []T, less func(a, b T) bool) {
	sort.SliceStable(in, func(i, j int) bool {
		return less(in[i], in[j])
	})
}

// SliceSortCopy copy and sort slice.
func SliceSortCopy[T any](in []T, less func(a, b T) bool) []T {
	res := make([]T, len(in))
	copy(res, in)
	sort.SliceStable(res, func(i, j int) bool {
		return less(res[i], res[j])
	})

	return res
}

// SliceGroupBy will group all
func SliceGroupBy[K comparable, V any](in []V, fn func(V) K) map[K][]V {
	if len(in) == 0 {
		return make(map[K][]V, 0)
	}

	res := make(map[K][]V, len(in))
	for i := range in {
		key := fn(in[i])
		res[key] = append(res[key], in[i])
	}

	return res
}

// Slice2MapFn apply fn to every elem. fn should return key and value, which
// will be applied to result map.
func Slice2MapFn[T any, K comparable, V any](in []T, fn func(idx int, elem T) (K, V)) map[K]V {
	m := make(map[K]V, len(in))
	for i := range in {
		k, v := fn(i, in[i])
		m[k] = v
	}

	return m
}

// Slice2MapFnErr apply fn to every elem. fn should return key and value, which
// will be applied to result map. return error when at least one fn returns an
// error.
func Slice2MapFnErr[T any, K comparable, V any](in []T, fn func(idx int, elem T) (K, V, error)) (map[K]V, error) {
	m := make(map[K]V, len(in))
	for i := range in {
		k, v, err := fn(i, in[i])
		if err != nil {
			return nil, err
		}

		m[k] = v
	}

	return m, nil
}

// Slice2Chan make chan with specified capacity from source slice.
func Slice2Chan[T any](in []T, capacity int) chan T {
	if len(in) == capacity {
		return Slice2ChanFill(in)
	}

	ch := make(chan T, capacity)
	go func() {
		for i := range in {
			ch <- in[i]
		}
	}()

	return ch
}

// Slice2ChanFill make chan from source slice with will already filled by all
// elements from source slice.
func Slice2ChanFill[T any](in []T) chan T {
	ch := make(chan T, len(in))
	ChanPut(ch, in)

	return ch
}

// SliceFromElem return a slice which contains only one element `elem`.
func SliceFromElem[T any](elem T) []T {
	return []T{elem}
}

// SliceGetFirstN return a subslice of source slice, which will contains not
// more than `maxElems` items.
func SliceGetFirstN[T any](in []T, maxElems int) []T {
	if maxElems < 0 {
		panic("maxElems should be >= 0")
	}

	if len(in) < maxElems {
		return in
	}

	return in[:maxElems]
}

// SliceCopy return a copy of source slice.
func SliceCopy[T any](in []T) []T {
	res := make([]T, len(in))
	copy(res, in)

	return res
}

// SliceReplaceFirst will replace the first element in `in` such
// that fn(index, elem) == true.
// Will do nothing if the element is not found.
func SliceReplaceFirst[T any](in []T, fn func(i int, elem T) bool, newElem T) {
	e := SliceFindFirst(in, fn)
	if e.Idx == -1 {
		return
	}

	in[e.Idx] = newElem
}

// SliceReplaceFirstOrAdd will replace the first element in `in` such
// that fn(index, elem) == true.
// Will add a `newElem` if the element is not found.
func SliceReplaceFirstOrAdd[T any](in []T, fn func(i int, elem T) bool, newElem T) []T {
	e := SliceFindFirst(in, fn)
	if e.Idx == -1 {
		return append(in, newElem)
	}

	in[e.Idx] = newElem

	return in
}

// SliceLastDefault return a last elem from slice or default value in case of zero slice.
func SliceLastDefault[T any](in []T, defaultVal T) T {
	if len(in) == 0 {
		return defaultVal
	}

	return in[len(in)-1]
}

// Slice2Iter create an iterator from slice. Check this docs https://go.dev/ref/spec#For_range.
func Slice2Iter[T any](in []T) func(func(int, T) bool) {
	return func(yield func(int, T) bool) {
		for i := range in {
			if !yield(i, in[i]) {
				return
			}
		}
	}
}

// SliceShuffle will shuffle the slice in-place.
func SliceShuffle[T any](in []T) {
	for i := range in {
		j := rand.Intn(i + 1)
		in[i], in[j] = in[j], in[i]
	}
}

// SliceShuffleCopy will make a copy and shuffle slice.
func SliceShuffleCopy[T any](in []T) []T {
	res := make([]T, len(in))
	copy(res, in)

	SliceShuffle(res)

	return res
}
