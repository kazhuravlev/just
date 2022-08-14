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

// Sum returns the sum of numbers from `in`.
func Sum[T number](in ...T) T {
	var acc T
	for i := range in {
		acc += in[i]
	}

	return acc
}
