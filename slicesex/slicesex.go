package slicesex

import (
	"iter"
	"slices"

	"github.com/DeedleFake/xiter"
)

func Reserve[T any](x []T, size int) []T {
	if len(x) < size {
		x = append(x, make([]T, size-len(x))...)
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
