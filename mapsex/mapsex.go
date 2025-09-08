package mapsex

// Gets first element from map
func First[K comparable, V any](m map[K]V) (k K, v V, ok bool) {
	for k, v = range m {
		ok = true
		return
	}

	ok = false
	return
}
