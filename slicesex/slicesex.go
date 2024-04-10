package slicesex

import "slices"

func BinarySearchInsert[S ~[]E, E any](x S, item E, cmp func(E, E) int) S {
	i, _ := slices.BinarySearchFunc(x, item, cmp)
	return slices.Insert(x, i, item)
}

func Any[S ~[]E, E any](x S, fn func(E) bool) bool {
	for _, t := range x {
		if fn(t) {
			return true
		}
	}
	return false
}

func Select[S ~[]E, E, V any](x S, fn func(E) V) []V {
	result := make([]V, len(x))
	for i, t := range x {
		result[i] = fn(t)
	}
	return result
}

func WhereSelect[S ~[]E, E, V any](x S, whereFn func(E) bool, selectFn func(E) V) []V {
	i := 0
	result := make([]V, 0, len(x))
	for _, t := range x {
		if whereFn(t) {
			result[i] = selectFn(t)
			i++
		}
	}
	return result
}

func Range(start int, count int) []int {
	nums := make([]int, count)
	for i := range nums {
		nums[i] = start + i
	}
	return nums
}
