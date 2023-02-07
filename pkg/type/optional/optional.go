package optional

import (
	"database/sql/driver"
	"encoding/json"
)

type Optional[T any] struct {
	value T
	isSet bool
}

func Some[T any](v T) Optional[T] {
	return Optional[T]{value: v, isSet: true}
}

func None[T any]() Optional[T] {
	return Optional[T]{isSet: false}
}

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if o.isSet {
		return json.Marshal(o.value)
	}
	return []byte("null"), nil
}

func (o *Optional[T]) UnmarshalJSON(b []byte) error {
	if b == nil || string(b) == "null" {
		*o = None[T]()
		return nil
	}

	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	*o = Some(v)

	return nil
}

func (o Optional[T]) Value() (driver.Value, error) {
	if o.isSet {
		return o.value, nil
	}
	return zero[T](), nil
}

func (o *Optional[T]) Scan(value any) error {
	if value == nil {
		*o = None[T]()
		return nil
	}

	v, ok := value.(T)
	if !ok {
		return ErrOptionalScan
	}

	*o = Some(v)

	return nil
}

func (o Optional[T]) IsSet() bool {
	return o.isSet
}

func (o Optional[T]) Get() T {
	return o.value
}

func zero[T any]() T {
	var o T
	return o
}
