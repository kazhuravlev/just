package just

import (
	"database/sql"
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-yaml"
)

// NullVal represent the nullable value for this type.
type NullVal[T any] struct {
	Val   T    `json:"v"`
	Valid bool `json:"ok"`
}

// UnmarshalYAML implements the interface for unmarshalling yaml.
func (nv *NullVal[T]) UnmarshalYAML(bb []byte) error {
	if len(bb) == 0 {
		nv.Valid = false
		nv.Val = *new(T)
		return nil
	}

	if err := yaml.Unmarshal(bb, &nv.Val); err != nil {
		return err
	}

	nv.Valid = true

	return nil
}

// MarshalYAML implements the interface for marshalling yaml.
func (nv NullVal[T]) MarshalYAML() ([]byte, error) {
	if !nv.Valid {
		return []byte("null"), nil
	}

	return yaml.Marshal(nv.Val)
}

// Scan implements the Scanner interface.
func (nv *NullVal[T]) Scan(value any) error {
	if v, ok := any(&nv.Val).(sql.Scanner); ok {
		if err := v.Scan(value); err != nil {
			nv.Val, nv.Valid = *new(T), false
			return err
		}

		nv.Valid = true
		return nil
	}

	if value == nil {
		nv.Val, nv.Valid = *new(T), true
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

	if v, ok := any(nv.Val).(driver.Valuer); ok {
		return v.Value()
	}

	return nv.Val, nil
}

// ValueOk returns the NullVal.Val and NullVal.Valid.
func (nv NullVal[T]) ValueOk() (T, bool) {
	return nv.Val, nv.Valid
}

// SetDefault set value `val` if NullVal.Valid == false.
func (nv *NullVal[T]) SetDefault(val T) bool {
	if nv.Valid {
		return false
	}

	nv.Val = val

	return true
}

// Null returns NullVal for `val` type, which are `NullVal.Valid == true`.
func Null[T any](val T) NullVal[T] {
	return NullVal[T]{
		Val:   val,
		Valid: true,
	}
}

// NullNull returns NullVal, which are `NullVal.Valid == false`.
func NullNull[T any]() NullVal[T] {
	return NullVal[T]{
		Val:   *new(T),
		Valid: false,
	}
}
