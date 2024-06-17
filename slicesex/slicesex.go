package slicesex

import (
	"slices"

	"github.com/DeedleFake/xiter"
)

func BinarySearchInsert[S ~[]E, E any](x S, item E, cmp func(E, E) int) S {
	i, _ := slices.BinarySearchFunc(x, item, cmp)
	return slices.Insert(x, i, item)
}

func Transform[S ~[]E, E, V any](x S, fn func(E) V) []V {
	result := make([]V, len(x))
	for i, t := range x {
		result[i] = fn(t)
	}
	return result
}

func Range(start int, count int) xiter.Seq[int] {
	return xiter.Limit(xiter.Generate(start, 1), count)
}
