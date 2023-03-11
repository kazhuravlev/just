package just

// Pointer returns a pointer to `v`.
func Pointer[T any](v T) *T {
	return &v
}

// PointerUnwrap returns the value from the pointer.
func PointerUnwrap[T any](in *T) T {
	return *in
}

// PointerUnwrapDefault returns a value from pointer or defaultVal when input
// is an empty pointer.
func PointerUnwrapDefault[T builtin | any](in *T, defaultVal T) T {
	if in == nil {
		return defaultVal
	}

	return *in
}
