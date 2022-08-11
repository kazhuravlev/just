package just

type number interface {
	int | int64 | int32 | int16 | int8 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float64 | float32
}

type builtin interface {
	number | bool | string
}

const zero = 0
