package sortex

import "sort"

func Slice[V any](a []V, less func(a, b V) bool) {
	sort.Slice(a, func(i, j int) bool { return less(a[i], a[j]) })
}

func Search[V any](a []V, fn func(v V) bool) int {
	return sort.Search(len(a), func(i int) bool { return fn(a[i]) })
}
