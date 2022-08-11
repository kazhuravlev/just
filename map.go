package just

func Map[T any, V any](in []T, fn func(T) V) []V {
	res := make([]V, len(in))
	for i := range in {
		res[i] = fn(in[i])
	}

	return res
}
