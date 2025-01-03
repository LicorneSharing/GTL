package slices

func Map[T any, U any, S ~[]T](slice S, fn func(T) U) []U {
	return MapRef(slice, func(e *T) U {
		return fn(*e)
	})
}

func MapRef[E, U any, S ~[]E](slice S, fn func(*E) U) []U {
	result := make([]U, len(slice))

	for i := range slice {
		result[i] = fn(&slice[i])
	}

	return result
}
