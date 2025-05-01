package pkl

import (
	"bytes"
	"encoding/json"
)

// Option is a wrapper around a value, and a flag indicating whether it is present or not.
//
// To obtain the underlying value, call Get.
//
//	value, exists := o.Get()
//	if !exists {
//		handleNotExists()
//	}
type Option[T any] struct {
	Value  T
	Exists bool
}

// Some creates an Option that is populated by Value.
func Some[T any](value T) Option[T] {
	return Option[T]{Value: value, Exists: true}
}

// None creates an Option with no Value.
func None[T any]() Option[T] {
	return Option[T]{Exists: false}
}

// Get returns the Value, and a boolean indicating if the value exists or not.
//
// If the Value does not exist, the first return value is the type's zero value.
func (o *Option[T]) Get() (T, bool) {
	return o.Value, o.Exists
}

var jsonNull = []byte("null")

func (o *Option[T]) MarshalJSON() ([]byte, error) {
	if value, exists := o.Get(); exists {
		return json.Marshal(value)
	}
	return jsonNull, nil
}

func (o *Option[T]) UnmarshalJSON(data []byte) error {
	if len(data) <= 0 || bytes.Equal(data, jsonNull) {
		*o = None[T]()
		return nil
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*o = Some(v)
	return nil
}
