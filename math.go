package just

// Max returns the max number from `in`.
func Max[T number](in ...T) T {
	if len(in) == 0 {
		panic("cannot find max of nothing")
	}

	if len(in) == 1 {
		return in[0]
	}

	res := in[0]
	for i := 1; i < len(in); i++ {
		if in[i] > res {
			res = in[i]
		}
	}

	return res
}

// MaxOr returns max element from `in` or defaultVal when `in` is empty
func MaxOr[T number](defaultVal T, in ...T) T {
	if len(in) == 0 {
		return defaultVal
	}

	return Max(in...)
}

// MaxDefault returns max element from `in` or default value for specified
// type when `in` is empty.
func MaxDefault[T number](in ...T) T {
	return MaxOr(0, in...)
}

// Min returns the min number from `in`.
func Min[T number](in ...T) T {
	if len(in) == 0 {
		panic("cannot find min of nothing")
	}

	if len(in) == 1 {
		return in[0]
	}

	res := in[0]
	for i := 1; i < len(in); i++ {
		if in[i] < res {
			res = in[i]
		}
	}

	return res
}

// MinOr returns min element from `in` or defaultVal when `in` is empty
func MinOr[T number](defaultVal T, in ...T) T {
	if len(in) == 0 {
		return defaultVal
	}

	return Min(in...)
}

// MinDefault returns min element from `in` or default value for specified
// type when `in` is empty.
func MinDefault[T number](in ...T) T {
	return MinOr(0, in...)
}

// Sum returns the sum of numbers from `in`.
func Sum[T number](in ...T) T {
	var acc T
	for i := range in {
		acc += in[i]
	}

	return acc
}

// Abs returns abs value of v.
func Abs[T number](v T) T {
	if v < 0 {
		return -v
	}

	return v
}
