package mathex

func DivideF[T, U Number](v U, a ...T) (r []float32) {
	fv := float32(v)
	r = make([]float32, len(a))

	for i, ai := range a {
		r[i] = float32(ai) / fv
	}

	return
}

func BiggestF[T, U Number](v []U, a []T) (r T) {
	var mv U = 0
	for i, vi := range v {
		if vi > mv {
			r = a[i]
			mv = vi
		}
	}

	return
}

func MaxA[T Number](v []T) (r T) {
	var mv T = 0
	for _, vi := range v {
		if vi > mv {
			r = vi
			mv = vi
		}
	}

	return
}

func MaxAF[T, U Number](v []T, fn func(i int) U) (r T) {
	var mu U = 0
	for i := range v {
		nu := fn(i)
		if nu > mu {
			r = v[i]
			mu = nu
		}
	}

	return
}
