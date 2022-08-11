package just

func Pointer[T any](v T) *T {
	return &v
}
