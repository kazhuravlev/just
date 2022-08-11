package just

func FilterSlice[T any](in []T, fn func(T) bool) []T {
	res := make([]T, 0, len(in))
	for i := range in {
		if !fn(in[i]) {
			continue
		}

		res = append(res, in[i])
	}

	return res
}
