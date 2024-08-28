package gx

func Must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}
