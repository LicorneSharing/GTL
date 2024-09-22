package optional

import "encoding/json"

func Nil[T any]() Value[T] {
	return Value[T]{}
}

type Value[T any] struct {
	value *T
}

func Some[T any](value T) Value[T] {
	return Value[T]{
		value: &value,
	}
}
func ZeroValue[T any]() Value[T] {
	return Value[T]{
		value: new(T),
	}
}

// Get may panic if Value is Nil
func (v Value[T]) Get() T {
	return *v.value
}

// GetRef may panic if Value is Nil. If Value is not Nil, then the returned pointer is never nil.
func (v Value[T]) GetRef() *T {
	if v.value == nil {
		_ = *v.value
	}

	return v.value
}

func (v Value[T]) HasValue() bool {
	return v.value != nil
}

func (v Value[T]) LookupValue() (T, bool) {
	if v.HasValue() {
		return v.Get(), true
	}

	var zero T
	return zero, false
}

func (v Value[T]) LooupRef() (*T, bool) {
	if v.HasValue() {
		return v.GetRef(), true
	}

	return nil, false
}

func (v Value[T]) GetValueOr(or T) T {
	if !v.HasValue() {
		return or
	}

	return v.Get()
}

func (v *Value[T]) Set(value T) {
	if v.value == nil {
		v.value = &value
	} else {
		*v.value = value
	}
}

// Assign may set Value to Nil, thus it would lose the reference to the pointer returned by Get
func (v *Value[T]) Assign(value Value[T]) {
	if value.value == nil {
		v.value = nil
	} else {
		v.Set(*value.value)
	}
}

// SetNil will lose the reference returned by Get
func (v *Value[T]) SetNil() {
	v.Assign(Nil[T]())
}

func (v Value[T]) MarshalJSON() ([]byte, error) {
	if v.HasValue() {
		return json.Marshal(v.GetRef())
	}

	return []byte("null"), nil
}

func (v *Value[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		v.SetNil()
		return nil
	}

	v.value = new(T)
	return json.Unmarshal(data, v.value)
}
