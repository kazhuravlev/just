package just

// SliceUniq returns only unique values from `in`.
func SliceUniq[T comparable](in []T) []T {
	index := Slice2Map(in)

	res := make([]T, 0, len(index))
	for k := range index {
		res = append(res, k)
	}

	return res
}

// SliceMap returns the slice where each element of `in` was handled by `fn`.
func SliceMap[T any, V any](in []T, fn func(T) V) []V {
	res := make([]V, len(in))
	for i := range in {
		res[i] = fn(in[i])
	}

	return res
}

// SliceFilter returns slice of values from `in` where `fn(elem) == true`.
func SliceFilter[T any](in []T, fn func(T) bool) []T {
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
		return []T{}
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

// SliceDifference return the difference between `oldSlice` and `newSlice`.
// Returns only elements which presented in `newSlice` but not presented
// in `oldSlice`.
// Example: [1,2,3], [3,4,5] => [4,5]
func SliceDifference[T comparable](oldSlice, newSlice []T) []T {
	if len(oldSlice) == 0 {
		return newSlice
	}

	if len(newSlice) == 0 {
		return nil
	}

	index := Slice2Map(oldSlice)
	res := make([]T, 0, len(newSlice))
	for i := range newSlice {
		if _, ok := index[newSlice[i]]; ok {
			continue
		}

		res = append(res, newSlice[i])
	}

	return res
}

// SliceWithoutElem returns the slice `in` that not contains `elem`.
func SliceWithoutElem[T comparable](in []T, elem T) []T {
	res := make([]T, 0, len(in))
	for i := range in {
		if in[i] == elem {
			continue
		}

		res[i] = in[i]
	}

	return res
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
		return nil
	}

	maxLen := len(in[0])
	for i := range in {
		if len(in[i]) < maxLen {
			maxLen = len(in[i])
		}
	}

	if maxLen == 0 {
		return nil
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

// SliceChunk split `in` into chunks by fn(index, elem) == true.
func SliceChunk[T any](in []T, fn func(i int, elem T) bool) [][]T {
	if len(in) == 0 {
		return nil
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

// SliceFindFirst return first elem from `in` that fn(index, elem) == true.
// returns index of found elem or -1 if elem not found.
func SliceFindFirst[T any](in []T, fn func(i int, elem T) bool) (T, int) {
	for i := range in {
		if fn(i, in[i]) {
			return in[i], i
		}
	}

	var v T
	return v, -1
}

// SliceFindLast return last elem from `in` that fn(index, elem) == true.
// returns index of found elem or -1 if elem not found.
func SliceFindLast[T any](in []T, fn func(i int, elem T) bool) (T, int) {
	for i := len(in) - 1; i != -1; i-- {
		if fn(i, in[i]) {
			return in[i], i
		}
	}

	var v T
	return v, -1
}

// SliceElem represent element of slice.
type SliceElem[T any] struct {
	// Idx is index of element in slice.
	Idx int
	// Val is value on slice by Idx index.
	Val T
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
