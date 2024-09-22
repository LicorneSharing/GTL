package iter

import "iter"

type Seq[T any] iter.Seq[T]

func (s Seq[T]) Collect() []T {
	var result []T

	for elem := range s {
		result = append(result, elem)
	}

	return result
}
func (s Seq[T]) ToMapSeq() MapSeq[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		s(func(elem T) bool {
			defer func() { i++ }()
			return yield(i, elem)
		})
	}
}

type Seq2[K, V any] iter.Seq2[K, V]

type KeyValue[K, V any] struct {
	Key   K
	Value V
}

func (s Seq2[K, V]) Collect() []KeyValue[K, V] {
	var result []KeyValue[K, V]

	for key, value := range s {
		result = append(result, KeyValue[K, V]{
			Key:   key,
			Value: value,
		})
	}

	return result
}

type MapSeq[K comparable, V any] Seq2[K, V]

func (ms MapSeq[K, V]) Collect() map[K]V {
	result := make(map[K]V)

	for key, value := range ms {
		result[key] = value
	}

	return result
}
