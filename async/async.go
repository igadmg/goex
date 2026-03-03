package async

import "fmt"

var (
	errNil = fmt.Errorf("future is nil")
)

type poller interface {
	Poll() bool
}

type futureResult[T any] struct {
	value T
	err   error
}

type FutureFn[T any] func() (T, error)
type FutureThenFn[T, U any] func(T) (U, error)

type Future[T any] struct {
	parent poller
	ch     chan futureResult[T]
	result futureResult[T]
	then   func()
}

func Go[T any](fn FutureFn[T]) *Future[T] {
	f := &Future[T]{
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

func Then[T any, U any](f *Future[T], fn FutureThenFn[T, U]) *Future[U] {
	tf := &Future[U]{
		parent: f,
	}

	f.then = func() {
		tf.parent = nil
		if f.result.err != nil {
			tf.result.err = f.result.err
		} else {
			tf.ch = make(chan futureResult[U])
			go func() {
				v, err := fn(f.result.value)
				tf.ch <- futureResult[U]{
					value: v,
					err:   err,
				}
			}()
		}
	}

	return tf
}

func (f *Future[T]) IsValid() bool {
	return f != nil && f.ch != nil
}

func (f *Future[T]) Value() (v T, err error) {
	if f != nil {
		return f.result.value, f.result.err
	}

	err = errNil
	return
}

func (f *Future[T]) Poll() bool {
	if f.parent != nil {
		return f.parent.Poll()
	}

	ok := false
	select {
	case f.result, ok = <-f.ch:
		if ok {
			close(f.ch)
		}
		if f.then == nil {
			return true
		} else {
			f.then()
			return false
		}
	default:
		return false
	}
}
