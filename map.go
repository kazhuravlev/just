package just

import "golang.org/x/exp/maps"

// MapMerge returns the map which contains all keys from m1, m2, and values
// from `fn(key, m1Value, m2Value)`.
func MapMerge[M ~map[K]V, K comparable, V any](m1, m2 M, fn func(k K, v1, v2 V) V) M {
	m := make(M, len(m1))
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
func MapFilter[M ~map[K]V, K comparable, V any](in M, fn func(k K, v V) bool) M {
	m := make(M, len(in))
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
func MapFilterKeys[M ~map[K]V, K comparable, V any](in M, fn func(k K) bool) M {
	return MapFilter(in, func(k K, _ V) bool {
		return fn(k)
	})
}

// MapFilterValues returns the map which contains elements that
// `fn(value) == true`. That is a simplified version of MapFilter.
func MapFilterValues[M ~map[K]V, K comparable, V any](in M, fn func(v V) bool) M {
	return MapFilter(in, func(_ K, v V) bool {
		return fn(v)
	})
}

// MapGetKeys returns all keys of the map.
func MapGetKeys[M ~map[K]V, K comparable, V any](m M) []K {
	return maps.Keys(m)
}

// MapGetValues returns all values of the map. Not Uniq, unordered.
func MapGetValues[M ~map[K]V, K comparable, V any](m M) []V {
	return maps.Values(m)
}

// KV represents the key-value of the map.
type KV[K comparable, V any] struct {
	Key K
	Val V
}

// MapPairs returns a slice of KV structs that contains key-value pairs.
func MapPairs[M ~map[K]V, K comparable, V any](m M) []KV[K, V] {
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
func MapDefaults[M ~map[K]V, K comparable, V any](m, defaults M) M {
	res := MapCopy(m)
	for k, v := range defaults {
		if _, ok := res[k]; !ok {
			res[k] = v
		}
	}

	return res
}

// MapCopy returns a shallow copy of the map.
func MapCopy[M ~map[K]V, K comparable, V any](m M) M {
	return maps.Clone(m)
}

// MapMap applies fn to all kv pairs from in.
func MapMap[M ~map[K]V, K, K1 comparable, V, V1 any](in M, fn func(K, V) (K1, V1)) map[K1]V1 {
	res := make(map[K1]V1, len(in))
	for k, v := range in {
		k1, v1 := fn(k, v)
		res[k1] = v1
	}

	return res
}

// MapMapErr applies fn to all kv pairs from in.
func MapMapErr[M ~map[K]V, K, K1 comparable, V, V1 any](in M, fn func(K, V) (K1, V1, error)) (map[K1]V1, error) {
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
func MapContainsKey[M ~map[K]V, K comparable, V any](m M, key K) bool {
	_, ok := m[key]

	return ok
}

// MapContainsKeysAny returns true when at least one key exists in the map.
func MapContainsKeysAny[M ~map[K]V, K comparable, V any](m M, keys []K) bool {
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

// MapContainsKeysAll returns true when at all keys exist in the map.
func MapContainsKeysAll[M ~map[K]V, K comparable, V any](m M, keys []K) bool {
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

// MapApply applies fn to each kv pair
func MapApply[M ~map[K]V, K comparable, V any](in M, fn func(k K, v V)) {
	for k, v := range in {
		fn(k, v)
	}
}

// MapJoin will create a new map containing all key-value pairs from app input
// maps. If several maps have duplicate keys - the last write wins.
func MapJoin[M ~map[K]V, K comparable, V any](maps ...M) M {
	res := make(M)
	for i := range maps {
		for k, v := range maps[i] {
			res[k] = v
		}
	}

	return res
}

// MapGetDefault returns a value for a given key or default value if the key
// is not present in the source map.
func MapGetDefault[M ~map[K]V, K comparable, V any](in M, key K, defaultVal V) V {
	v, ok := in[key]
	if !ok {
		return defaultVal
	}

	return v
}

// MapNotNil returns the source map when it is not nil or creates an empty
// instance of this type.
func MapNotNil[T ~map[K]V, K comparable, V any](in T) T {
	if in == nil {
		return make(T, 0)
	}

	return in
}

// MapDropKeys remove all keys from the source map. Map will change in place.
func MapDropKeys[M ~map[K]V, K comparable, V any](in M, keys ...K) {
	if len(keys) == 0 || len(in) == 0 {
		return
	}

	for i := range keys {
		delete(in, keys[i])
	}
}

// MapPopKeyDefault will return value for given key and delete this key from source map.
// In case of key do not presented in map - returns default value.
func MapPopKeyDefault[M ~map[K]V, K comparable, V any](in M, key K, def V) V {
	val, ok := in[key]
	if ok {
		delete(in, key)
		return val
	}

	return def
}
