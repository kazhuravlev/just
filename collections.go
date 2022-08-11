package just

func MergeMap[K comparable, V any](m1, m2 map[K]V, fn func(k K, v1, v2 V) V) map[K]V {
	m := make(map[K]V, len(m1))
	for k, v := range m1 {
		m[k] = fn(k, v, m2[k])
	}

	var emptyVal V
	for k, v := range m2 {
		if _, ok := m[k]; ok {
			continue
		}

		m[k] = fn(k, emptyVal, v)
	}

	return m
}

func FilterMap[K comparable, V any](in map[K]V, fn func(k K, v V) bool) map[K]V {
	m := make(map[K]V, len(in))
	for k, v := range in {
		if !fn(k, v) {
			continue
		}

		m[k] = v
	}

	return m
}

func FilterMapKeys[K comparable, V any](in map[K]V, fn func(k K) bool) map[K]V {
	return FilterMap(in, func(k K, _ V) bool {
		return fn(k)
	})
}

func FilterMapValues[K comparable, V any](in map[K]V, fn func(v V) bool) map[K]V {
	return FilterMap(in, func(_ K, v V) bool {
		return fn(v)
	})
}

func Uniq[T comparable](in []T) []T {
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
