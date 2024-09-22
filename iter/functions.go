package iter

import "slices"

func MapSeqSeq[T any, S ~func(func(T) bool), U any](seq S, fn func(T) U) Seq[U] {
	return func(yield func(U) bool) {
		seq(func(elem T) bool {
			return yield(fn(elem))
		})
	}
}

func MapSeq2Seq2[K, V any, S ~func(func(K, V) bool), K2, V2 any](seq S, fn func(K, V) (K2, V2)) Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		seq(func(k K, v V) bool {
			return yield(fn(k, v))
		})
	}
}

func MapSeqSeq2[T any, S ~func(func(T) bool), K, V any](seq S, fn func(T) (K, V)) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(elem T) bool {
			return yield(fn(elem))
		})
	}
}

func MapSeq2Seq[K, V any, S ~func(func(K, V) bool), T any](seq S, fn func(K, V) T) Seq[T] {
	return func(yield func(T) bool) {
		seq(func(k K, v V) bool {
			return yield(fn(k, v))
		})
	}
}

func FilterSeq[T any, S ~func(func(T) bool)](seq S, filter func(T) bool) Seq[T] {
	return func(yield func(T) bool) {
		seq(func(elem T) bool {
			if filter(elem) {
				return yield(elem)
			}

			return true
		})
	}
}

func FilterSeq2[K, V any, S ~func(func(K, V) bool)](seq S, filter func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			if filter(k, v) {
				return yield(k, v)
			}

			return true
		})
	}
}

func FilterMapSeq[K comparable, V any, S ~func(func(K, V) bool)](seq S, filter func(K, V) bool) MapSeq[K, V] {
	return MapSeq[K, V](FilterSeq2[K, V](seq, filter))
}

type Filtering[T any] []T

func (f Filtering[T]) By(filter func(T) bool) Seq[T] {
	return FilterSeq(slices.Values(f), filter)
}
