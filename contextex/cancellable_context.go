package contextex

import (
	"context"
	"time"
)

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

type CancelContext struct {
	context.Context
	cancel context.CancelFunc
}

func (g *CancelContext) Prime(ctx context.Context) {
	g.Context, g.cancel = context.WithCancel(ctx)
}

func (g CancelContext) AfterFunc(fn func()) {
	context.AfterFunc(g.Context, fn)
}

func (g CancelContext) Cancel() CancelContext {
	if g.cancel != nil {
		g.cancel()
		//g.Context = nil // no need to clear context here, because that crashes waiting goroutines
		g.cancel = nil
	}

	return g
}

func (g *CancelContext) Cancellabe(fn func(context.Context) bool) func() {
	return Cancellabe(g, fn)
}

func (g *CancelContext) CancellabeWithTimeout(timeout time.Duration, fn func(context.Context) bool) func() {
	return CancellabeWithTimeout(g, timeout, fn)
}
