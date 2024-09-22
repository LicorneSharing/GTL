package slices

func Map[T any, U any, S ~[]T](slice S, fn func(T) U) []U {
	result := make([]U, len(slice))

	for i, elem := range slice {
		result[i] = fn(elem)
	}

	return result
}
