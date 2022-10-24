package just

// Pointer return a pointer to `v`.
func Pointer[T any](v T) *T {
	return &v
}

// PointerUnwrap returns value from pointer.
func PointerUnwrap[T any](in *T) T {
	return *in
}
