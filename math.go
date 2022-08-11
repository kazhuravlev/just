package just

type number interface {
	int | int64 | int32 | int16 | int8 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float64 | float32
}

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

func Sum[T number](in ...T) T {
	var acc T
	for i := range in {
		acc += in[i]
	}

	return acc
}
