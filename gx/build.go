package gx

// Builds and object by applying options functions to it
func Build[T any](v T, options ...func(T) T) T {
	for _, o := range options {
		v = o(v)
	}

	return v
}
