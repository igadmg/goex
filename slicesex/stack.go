package slicesex

import "slices"

func Top[S ~[]E, E any](x S) E {
	return x[len(x)-1]
}

func Push[S ~[]E, E any](x S, v E) S {
	return append(x, v)
}

func Pop[S ~[]E, E any](x S) S {
	return slices.Delete(x, len(x)-1, len(x))
}
