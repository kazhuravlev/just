package just

// Bool returns true if element not equal to default value for this type.
func Bool[T builtin](v T) bool {
	switch x := any(v).(type) {
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

	return false
}

// Must will panic on error after calling typical function.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}
