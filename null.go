package just

import (
	"database/sql/driver"
	"errors"
)

type NullVal[T any] struct {
	Val   T
	Valid bool
}

// Scan implements the Scanner interface.
func (nv *NullVal[T]) Scan(value any) error {
	if value == nil {
		var val T
		nv.Val, nv.Valid = val, false
		return nil
	}

	v, ok := value.(T)
	if !ok {
		return errors.New("unexpected value type")
	}

	nv.Valid = true
	nv.Val = v
	return nil
}

// Value implements the driver Valuer interface.
func (nv NullVal[T]) Value() (driver.Value, error) {
	if !nv.Valid {
		return nil, nil
	}

	return nv.Val, nil
}

func (nv NullVal[T]) ValueOk() (T, bool) {
	return nv.Val, nv.Valid
}

func (nv NullVal[T]) SetDefault(val T) bool {
	if nv.Valid {
		return false
	}

	nv.Val = val

	return true
}

func Null[T any](val T) NullVal[T] {
	return NullVal[T]{
		Val:   val,
		Valid: true,
	}
}

func NullNull[T any]() NullVal[T] {
	var val T
	return NullVal[T]{
		Val:   val,
		Valid: false,
	}
}
