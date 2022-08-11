package just

func Bool[T builtin](v T) bool {
	switch x := interface{}(v).(type) {
	case bool:
		return x
	case uint8:
		return x > zero
	case uint16:
		return x > zero
	case uint32:
		return x > zero
	case uint64:
		return x > zero
	case int8:
		return x > zero
	case int16:
		return x > zero
	case int32:
		return x > zero
	case int64:
		return x > zero
	case float32:
		return x > zero
	case float64:
		return x > zero
	case int:
		return x > zero
	case uint:
		return x > zero
	case uintptr:
		return x > zero
	case string:
		return x != ""
	}

	panic("unknown type")
}
