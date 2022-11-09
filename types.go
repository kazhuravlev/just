package just

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Float | constraints.Signed | constraints.Unsigned
}

type builtin interface {
	number | ~bool | ~string
}

const zero = 0
