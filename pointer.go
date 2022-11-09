package just

// Pointer return a pointer to `v`.
func Pointer[T any](v T) *T {
	return &v
}

// PointerUnwrap returns value from pointer.
func PointerUnwrap[T any](in *T) T {
	return *in
}

// PointerUnwrapDefault returns value from pointer or defaultVal when input is
// empty pointer.
func PointerUnwrapDefault[T builtin | any](in *T, defaultVal T) T {
	if in == nil {
		return defaultVal
	}

	return *in
}
