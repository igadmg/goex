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

func MustHave[T any](v T, e bool) T {
	if !e {
		panic(e)
	}
	return v
}

func ShouldOk[T any](v T, ok bool) T {
	if ok {
		return v
	}
	var t T
	return t
}
