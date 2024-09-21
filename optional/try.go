package optional

func Try[T any](fn func() (T, error)) Value[T] {
	return WrapResult(fn())
}

func WrapResult[T any](value T, err error) Value[T] {
	if err != nil {
		return Nil[T]()
	}

	return Value[T]{value: &value}
}
