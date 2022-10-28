package just

// MapMerge returns the map which contains all keys from m1, m2 and values
// from `fn(key, m1Value, m2Value)`.
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

// MapFilter returns the map which contains elements that
// `fn(key, value) == true`.
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

// MapFilterKeys returns the map which contains elements that
// `fn(key) == true`. That is a simplified version of MapFilter.
func MapFilterKeys[K comparable, V any](in map[K]V, fn func(k K) bool) map[K]V {
	return MapFilter(in, func(k K, _ V) bool {
		return fn(k)
	})
}

// MapFilterValues returns the map which contains elements that
// `fn(value) == true`. That is a simplified version of MapFilter.
func MapFilterValues[K comparable, V any](in map[K]V, fn func(v V) bool) map[K]V {
	return MapFilter(in, func(_ K, v V) bool {
		return fn(v)
	})
}

// MapGetKeys returns all keys of map.
func MapGetKeys[K comparable, V any](m map[K]V) []K {
	if len(m) == 0 {
		return nil
	}

	res := make([]K, len(m))
	var i int
	for k := range m {
		res[i] = k
		i++
	}

	return res
}

// MapGetValues returns all values of map. Not Uniq, unordered.
func MapGetValues[K comparable, V any](m map[K]V) []V {
	if len(m) == 0 {
		return nil
	}

	res := make([]V, len(m))
	var i int
	for _, v := range m {
		res[i] = v
		i++
	}

	return res
}

// KV represents the key-value of map.
type KV[K comparable, V any] struct {
	Key K
	Val V
}

// MapPairs returns slice of KV structs that contains ket-value pairs.
func MapPairs[K comparable, V any](m map[K]V) []KV[K, V] {
	if len(m) == 0 {
		return nil
	}

	res := make([]KV[K, V], len(m))
	var i int
	for k, v := range m {
		res[i] = KV[K, V]{
			Key: k,
			Val: v,
		}
		i++
	}

	return res
}

// MapDefaults returns the map `m` after filling in its non-exists keys by
// `defaults`.
// Example: {1:1}, {1:0, 2:2} => {1:1, 2:2}
func MapDefaults[K comparable, V any](m map[K]V, defaults map[K]V) map[K]V {
	res := MapCopy(m)
	for k, v := range defaults {
		if _, ok := res[k]; !ok {
			res[k] = v
		}
	}

	return res
}

// MapCopy returns a shallow copy of map.
func MapCopy[K comparable, V any](m map[K]V) map[K]V {
	res := make(map[K]V, len(m))
	for k, v := range m {
		res[k] = v
	}

	return res
}

// MapMap apply fn to all kv pairs from in.
func MapMap[K, K1 comparable, V, V1 any](in map[K]V, fn func(K, V) (K1, V1)) map[K1]V1 {
	res := make(map[K1]V1, len(in))
	for k, v := range in {
		k1, v1 := fn(k, v)
		res[k1] = v1
	}

	return res
}

// MapMapErr apply fn to all kv pairs from in.
func MapMapErr[K, K1 comparable, V, V1 any](in map[K]V, fn func(K, V) (K1, V1, error)) (map[K1]V1, error) {
	res := make(map[K1]V1, len(in))
	for k, v := range in {
		k1, v1, err := fn(k, v)
		if err != nil {
			return nil, err
		}

		res[k1] = v1
	}

	return res, nil
}

// MapContainsKey returns true if key is exists in the map.
func MapContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]

	return ok
}

// MapContainsKeysAny returns true when at least one key exists in the map.
func MapContainsKeysAny[K comparable, V any](m map[K]V, keys []K) bool {
	if len(keys) == 0 {
		return false
	}

	if len(m) == 0 {
		return false
	}

	for i := range keys {
		if MapContainsKey(m, keys[i]) {
			return true
		}
	}

	return false
}

// MapContainsKeysAll returns true when at all keys exists in the map.
func MapContainsKeysAll[K comparable, V any](m map[K]V, keys []K) bool {
	if len(keys) == 0 {
		return false
	}

	if len(m) == 0 {
		return false
	}

	for i := range keys {
		if !MapContainsKey(m, keys[i]) {
			return false
		}
	}

	return true
}

// MapApply apply fn to each kv pair
func MapApply[K comparable, V any](in map[K]V, fn func(k K, v V)) {
	for k, v := range in {
		fn(k, v)
	}
}

// MapJoin will create a new map containing all key-value pairs from app input
// maps. If several maps have duplicate keys - the last write wins.
func MapJoin[K comparable, V any](maps ...map[K]V) map[K]V {
	res := make(map[K]V)
	for i := range maps {
		for k, v := range maps[i] {
			res[k] = v
		}
	}

	return res
}
