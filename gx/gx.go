package gx

func Must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}

func Should[T any](v T, e error) T {
	if e != nil {
		var d T
		return d
	}
	return v
}
