package just

func SliceUniq[T comparable](in []T) []T {
	m := make(map[T]struct{}, len(in))
	for i := range in {
		m[in[i]] = struct{}{}
	}

	res := make([]T, 0, len(m))
	for k := range m {
		res = append(res, k)
	}

	return res
}

func SliceMap[T any, V any](in []T, fn func(T) V) []V {
	res := make([]V, len(in))
	for i := range in {
		res[i] = fn(in[i])
	}

	return res
}

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

func SliceAny[T any](in []T, fn func(T) bool) bool {
	for i := range in {
		if fn(in[i]) {
			return true
		}
	}

	return false
}

func SliceAll[T any](in []T, fn func(T) bool) bool {
	for i := range in {
		if !fn(in[i]) {
			return false
		}
	}

	return true
}
