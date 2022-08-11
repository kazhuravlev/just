package just

func MapMerge[K comparable, V any](m1, m2 map[K]V, fn func(k K, v1, v2 V) V) map[K]V {
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

func MapFilter[K comparable, V any](in map[K]V, fn func(k K, v V) bool) map[K]V {
	m := make(map[K]V, len(in))
	for k, v := range in {
		if !fn(k, v) {
			continue
		}

		m[k] = v
	}

	return m
}

func MapFilterKeys[K comparable, V any](in map[K]V, fn func(k K) bool) map[K]V {
	return MapFilter(in, func(k K, _ V) bool {
		return fn(k)
	})
}

func MapFilterValues[K comparable, V any](in map[K]V, fn func(v V) bool) map[K]V {
	return MapFilter(in, func(_ K, v V) bool {
		return fn(v)
	})
}
