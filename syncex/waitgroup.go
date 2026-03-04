package syncex

import (
	"context"
	"sync"
	"time"
)

// WaitGroupGo runs multiple functions concurrently and calls waitFn when all are done.
func WaitGroupGo(waitFn func(), fn ...func()) {
	if len(fn) == 0 {
		waitFn()
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(fn))

	for _, f := range fn {
		wg.Go(f)
	}

	go func() {
		wg.Wait()
		waitFn()
	}()
}

func Cancellabe(ctx context.Context, fn func(context.Context) bool) func() {
	return func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !fn(ctx) {
					return
				}
			}
		}
	}
}

func CancellabeWithTimeout(ctx context.Context, timeout time.Duration, fn func(context.Context) bool) func() {
	return func() {
		for {
			select {
			case <-time.After(timeout):
				return
			case <-ctx.Done():
				return
			default:
				if !fn(ctx) {
					return
				}
			}
		}
	}
}
