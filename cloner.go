package goex

type Cloner interface {
	Clone() any
}

func Clone[T any](o any) (t T, ok bool) {
	co, ok := o.(Cloner)
	if !ok {
		return
	}

	t = co.Clone().(T)
	return
}
