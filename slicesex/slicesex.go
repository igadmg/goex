package slicesex

import (
	"iter"
	"slices"

	"github.com/DeedleFake/xiter"
)

func BinarySearchInsert[S ~[]E, E any](x S, item E, cmp func(E, E) int) S {
	i, _ := slices.BinarySearchFunc(x, item, cmp)
	return slices.Insert(x, i, item)
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
