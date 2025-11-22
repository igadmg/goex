package syncex

import (
	"sync"
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
