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

// Gets only element from map
func Only[K comparable, V any](m map[K]V) (k K, v V, ok bool) {
	if len(m) == 1 {
		for k, v = range m {
			ok = true
			return
		}
	}

	ok = false
	return
}

//func DeepCopy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
//	for k, v := range src {
//		var av any = v
//		switch vv := av.(type) {
//		case map[K]V:
//
//		}
//	}
//}
