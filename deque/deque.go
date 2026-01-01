package deque

import (
	"iter"
	"slices"

	"github.com/igadmg/goex/mathex"
	"github.com/igadmg/goex/unsafeex"
	"golang.org/x/sys/cpu"
)

var CacheLinePadSize = int(unsafeex.Sizeof[cpu.CacheLinePad]())

func PerfectPageSize[T any]() int {
	return mathex.LCM(CacheLinePadSize, int(unsafeex.Sizeof[T]()))
}

type Deque[T any] struct {
	page_size      int
	items_per_page int
	pages          [][]T
}

func Make[T any](page_size int) Deque[T] {
	r := Deque[T]{
		page_size: page_size,
	}
	tsizeof := int(unsafeex.Sizeof[T]())

	if tsizeof != 0 {
		r.items_per_page = page_size / tsizeof
		for r.items_per_page == 0 {
			page_size *= 2
			r.items_per_page = page_size / tsizeof
		}
	} else {
		// TODO need some special optimization for empty types
		r.items_per_page = r.page_size
	}

	return r
}

func MakePerfect[T any]() Deque[T] {
	r := Deque[T]{
		page_size: PerfectPageSize[T](),
	}
	r.items_per_page = r.page_size

	return r
}

func (d Deque[T]) lastPage() []T {
	if len(d.pages) == 0 {
		return []T{}
	}

	return d.pages[len(d.pages)-1]
}

func (d Deque[T]) makePage() []T {
	return make([]T, 0, d.items_per_page)
}

func (d Deque[T]) Each() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, page := range d.pages {
			for _, i := range page {
				if !yield(i) {
					return
				}
			}
		}
	}
}

func (d Deque[T]) Item(index int) *T {
	if d.items_per_page == 0 {
		return nil
	}

	page_index := index / d.items_per_page
	item_index := index % d.items_per_page

	if len(d.pages) <= page_index {
		return nil
	}

	page := d.pages[page_index]
	return &page[item_index]
}

func (d Deque[T]) Last() *T {
	lp := d.lastPage()
	if len(lp) == 0 {
		return nil
	}

	return &lp[len(lp)-1]
}

func (d Deque[T]) Reserve(capacity int) Deque[T] {
	page_index := capacity / d.items_per_page
	item_index := capacity % d.items_per_page
	if item_index != 0 {
		page_index++
	}

	if len(d.pages) < page_index {
		d.pages = slices.Grow(d.pages, page_index)
		for i := len(d.pages); i < page_index-1; i++ {
			d.pages = append(d.pages, d.makePage()[0:d.items_per_page])
		}
		if item_index == 0 {
			d.pages = append(d.pages, d.makePage()[0:d.items_per_page])
		} else {
			d.pages = append(d.pages, d.makePage()[0:min(d.items_per_page, item_index)])
		}
	}

	last_index := capacity - 1
	page_index = last_index / d.items_per_page
	item_index = last_index % d.items_per_page

	page := d.pages[page_index]
	d.pages[page_index] = page[0:max(len(page), item_index+1)]
	return d
}

func (d Deque[T]) Append(v T) Deque[T] {
	lastPage := d.lastPage()
	if len(lastPage) == cap(lastPage) {
		d.pages = append(d.pages, make([]T, 0, d.items_per_page))
	}
	d.pages[len(d.pages)-1] = append(d.pages[len(d.pages)-1], v)
	return d
}
