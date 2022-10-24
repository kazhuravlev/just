package just

import "errors"

// ErrIsAnyOf return true when at least one expression
// `errors.Is(err, errSlice[N])` return true.
func ErrIsAnyOf(err error, errSlice ...error) bool {
	for i := range errSlice {
		if errors.Is(err, errSlice[i]) {
			return true
		}
	}

	return false
}

// ErrIsNotAnyOf return true when all errors from errSlice is not
// `errors.Is(err, errSlice[N])`.
func ErrIsNotAnyOf(err error, errSlice ...error) bool {
	for i := range errSlice {
		if errors.Is(err, errSlice[i]) {
			return false
		}
	}

	return true
}
