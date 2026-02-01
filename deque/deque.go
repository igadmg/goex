package deque

import (
	"iter"
	"slices"

	"github.com/igadmg/goex/cache"
	"github.com/igadmg/goex/unsafeex"
)

type Of[T any] struct {
	page_size      int
	items_per_page int
	pages          [][]T
}

func Make[T any](page_size int) Of[T] {
	r := Of[T]{
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
		// TODO(iga): need some special optimization for empty types
		r.items_per_page = r.page_size
	}

	return r
}

func MakePerfect[T any]() Of[T] {
	r := Of[T]{
		page_size: cache.PerfectPageSize[T](),
	}
	r.items_per_page = r.page_size / int(unsafeex.Sizeof[T]())

	return r
}

func (d Of[T]) lastPage() []T {
	if len(d.pages) == 0 {
		return nil
	}

	return d.pages[len(d.pages)-1]
}

func (d Of[T]) makePage() []T {
	return make([]T, 0, d.items_per_page)
}

func (d Of[T]) Each() iter.Seq[T] {
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

func (d Of[T]) Item(index int) *T {
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

func (d Of[T]) Last() *T {
	lp := d.lastPage()
	if len(lp) == 0 {
		return nil
	}

	return &lp[len(lp)-1]
}

func (d Of[T]) Reserve(capacity int) Of[T] {
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

func (d Of[T]) Append(v T) Of[T] {
	lastPage := d.lastPage()
	if cap(lastPage) == 0 || len(lastPage) == cap(lastPage) {
		d.pages = append(d.pages, make([]T, 0, d.items_per_page))
	}
	d.pages[len(d.pages)-1] = append(d.pages[len(d.pages)-1], v)
	return d
}

func (d *Of[T]) Push(v T) {
	*d = d.Append(v)
}

func (d *Of[T]) Pop() (v T) {
	lastPage := d.lastPage()
	if len(lastPage) > 0 {
		v = lastPage[len(lastPage)-1]
		lastPage = lastPage[:len(lastPage)-1]
		if len(lastPage) > 0 {
			d.pages[len(d.pages)-1] = lastPage
		} else {
			d.pages = d.pages[:len(d.pages)-1]
		}
	}
	return
}
