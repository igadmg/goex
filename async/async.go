package async

import (
	"context"
	"fmt"
	"sync"

	"github.com/Mishka-Squat/goex/contextex"
)

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

type FutureCtxFn[T any] func(ctx context.Context) (T, error)
type FutureCtxThenFn[T, U any] func(ctx context.Context, v T) (U, error)
type FutureFn[T any] func() (T, error)
type FutureThenFn[T, U any] func(v T) (U, error)

type Future[T any] struct {
	parent poller
	ch     chan futureResult[T]
	result futureResult[T]
	then   func() bool
}

func GoCtx[T any](ctx context.Context, fn FutureCtxFn[T], onErr ...func(error)) *Future[T] {
	f := &Future[T]{
		ch: make(chan futureResult[T]),
	}

	go func() {
		v, err := fn(ctx)
		f.ch <- futureResult[T]{
			value: v,
			err:   err,
		}
	}()

	return f
}

func Go[T any](fn FutureFn[T], onErr ...func(error)) *Future[T] {
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

func Then[T any, U any](f *Future[T], fn FutureThenFn[T, U], onErr ...func(error)) *Future[U] {
	tf := &Future[U]{
		parent: f,
		ch:     make(chan futureResult[U]), // ch == nil is used as done flag
	}

	f.then = func() bool {
		tf.parent = nil
		if f.result.err != nil {
			tf.result.err = f.result.err
			for _, e := range onErr {
				e(f.result.err)
			}

			return true
		} else {
			go func() {
				v, err := fn(f.result.value)
				tf.ch <- futureResult[U]{
					value: v,
					err:   err,
				}
			}()

			return false
		}
	}

	return tf
}

func Error[T any](err error) *Future[T] {
	return &Future[T]{
		result: futureResult[T]{
			err: err,
		},
	}
}

func ErrorF[T any](format string, v ...any) *Future[T] {
	return &Future[T]{
		result: futureResult[T]{
			err: fmt.Errorf(format, v...),
		},
	}
}

//func (f *Future[T]) IsDone() bool {
//	return f != nil && f.ch == nil && f.result.err == nil
//}

func (f *Future[T]) Value() (v T, err error) {
	if f != nil {
		return f.result.value, f.result.err
	}

	err = errNil
	return
}

// Poll - polls for value, id value is ready or error occured returns true.
func (f *Future[T]) Poll() bool {
	if f == nil {
		return false
	}

	if f.parent != nil {
		return f.parent.Poll()
	}

	ok := false
	select {
	case f.result, ok = <-f.ch:
		if ok {
			close(f.ch)
			f.ch = nil
		}
		if f.then == nil {
			return true
		} else {
			return f.then()
		}
	default:
		return false
	}
}

func (f *Future[T]) Await(ctx context.Context) (v T, err error) {
	var wg sync.WaitGroup

	wg.Go(contextex.Cancellabe(ctx, func(ctx context.Context) bool {
		if f.Poll() {
			v, err = f.Value()
			return false
		}

		return true
	}))
	wg.Wait()

	return
}
