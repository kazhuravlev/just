package just

func Pointer[T any](v T) *T {
	return &v
}

func PointerUnwrap[T any](in *T) T {
	return *in
}
