package slicesex

import (
	"iter"
	"slices"

	"deedles.dev/xiter"
)

func Reserve[S ~[]E, E any](x S, size int) S {
	if len(x) < size {
		if size > cap(x) {
			x = slices.Grow(x, size)
		}

		x = x[0:size]
	}

	return x
}

func BinarySearchInsert[S ~[]E, E any](x S, item E, cmp func(E, E) int) S {
	i, _ := slices.BinarySearchFunc(x, item, cmp)
	return slices.Insert(x, i, item)
}

func BinarySearchTake[S ~[]E, E any](x S, item E, cmp func(E, E) int) S {
	i, _ := slices.BinarySearchFunc(x, item, cmp)
	return slices.Delete(x, i, 1)
}

func TakeFunc[S ~[]E, E any](x S, cmp func(E) bool) (s S, e E, ok bool) {
	i := slices.IndexFunc(x, cmp)
	if i < 0 {
		ok = false
		s = x

		return
	}
	e = x[i]
	s = slices.Delete(x, i, 1)
	ok = true

	return
}

func Any[S ~[]E, E any](x S, fn func(E) bool) bool {
	return xiter.Any(slices.Values(x), fn)
}

func Transform[S ~[]E, E, V any](x S, fn func(E) V) []V {
	return xiter.CollectSize(xiter.Map(slices.Values(x), fn), len(x))
}

func RepeatFunc[E any](count int, fn func(int) E) []E {
	return xiter.CollectSize(xiter.Map(Range(0, count), fn), count)
}

func Range(start int, count int) iter.Seq[int] {
	return xiter.Limit(xiter.Generate(start, 1), count)
}

func ToMap[T any, K comparable](src []T, key func(T) K) map[K]T {
	var result = make(map[K]T)
	for _, v := range src {
		result[key(v)] = v
	}
	return result
}

func ToMapPtr[T any, K comparable](src []T, key func(*T) K) map[K]*T {
	var result = make(map[K]*T)
	for i := range src {
		result[key(&src[i])] = &src[i]
	}
	return result
}

func Set[S ~[]E, E any](x S, index int, e E) S {
	x = Reserve(x, index+1)
	x[index] = e
	return x
}

func RemoveSwapback[S ~[]E, E any](x S, index int) (S, E) {
	var e E

	if index >= len(x) {
		return x, e
	}

	e = x[index]
	x[index] = x[len(x)-1]
	x = x[:len(x)-1]

	return x, e
}
