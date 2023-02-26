package optional

import (
	_ "database/sql"
	"database/sql/driver"
	"encoding/json"
	_ "unsafe"

	"github.com/mehdieidi/freakshake/pkg/xerror"
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

func (o Optional[T]) IsSome() bool {
	return o.isSet
}

func (o Optional[T]) IsNone() bool {
	return !o.isSet
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
	if o.IsNone() {
		return nil, nil
	}

	val, err := driver.DefaultParameterConverter.ConvertValue(o.value)
	xerror.Wrap(&err, "%T.Value()", o)

	return val, err
}

func (o *Optional[T]) Scan(value any) error {
	if value == nil {
		*o = None[T]()
		return nil
	}
	err := convertAssign(&o.value, value)
	xerror.Wrap(&err, "%T.Scan(%v)", *o, value)

	if err == nil {
		o.isSet = true
	}

	return err
}

func (o Optional[T]) Get() T {
	return o.value
}

func (o Optional[T]) GetDefault(x T) T {
	if o.IsNone() {
		return x
	}
	return o.value
}

func (o Optional[T]) GetOr(f func() T) T {
	if o.IsNone() {
		return f()
	}
	return o.value
}

//go:linkname convertAssign database/sql.convertAssign
func convertAssign(dest, src any) error
