package iter

func Map[T any, S ~func(func(T) bool), U any](seq S, fn func(T) U) Seq[U] {
	return func(yield func(U) bool) {
		seq(func(elem T) bool {
			return yield(fn(elem))
		})
	}
}

func Map2[K, V any, S ~func(func(K, V) bool), K2, V2 any](seq S, fn func(K, V) (K2, V2)) Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		seq(func(k K, v V) bool {
			return yield(fn(k, v))
		})
	}
}
