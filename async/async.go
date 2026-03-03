package async

type futureResult[T any] struct {
	value T
	err   error
}

type Future[T any] struct {
	ch     chan futureResult[T]
	result futureResult[T]
}

func Go[T any](fn func() (T, error)) Future[T] {
	f := Future[T]{
		ch: make(chan futureResult[T]),
	}

	go func() {
		v, err := fn()
		f.ch <- futureResult[T]{
			value: v,
			err:   err,
		}
	}()

	return f
}

func (f Future[T]) IsValid() bool {
	return f.ch != nil
}

func (f *Future[T]) Poll() bool {
	ok := false
	select {
	case f.result, ok = <-f.ch:
		if ok {
			close(f.ch)
		}
		return true
	default:
		return false
	}
}
