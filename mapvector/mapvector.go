package mapvector

import (
	"iter"
	"unsafe"

	"golang.org/x/exp/constraints"
)

const (
	DefaultPageSize = 256

	DefaultPageBits = 8
)

// MapVector is a paged container where the int key is split into page number (upper 8 bits)
// and index within page (lower 56 bits). Each page is a slice of T.
type MapVector[K constraints.Integer, T any] struct {
	pages     map[K][]T
	pageShift int
	PageSize  int // Public property for page size, set during initialization
}

// MakeMapVector creates a new MapVector with 8 bits for page number.
func MakeMapVector[K constraints.Integer, T any]() MapVector[K, T] {
	return MapVector[K, T]{
		pages:     make(map[K][]T),
		pageShift: int(unsafe.Sizeof(K(0)))*8 - DefaultPageBits,
		PageSize:  DefaultPageSize, // Default page size
	}
}

func (mv MapVector[K, T]) MakeKey(typeID int, index int) K {
	return K((uint64(typeID) << mv.pageShift) | uint64(index))
}

// PageAndIndex splits the key into page and index.
// Assumes key >= 0.
func (mv MapVector[K, T]) PageAndIndex(key K) (page K, index int) {
	ukey := uint64(key)
	page = K(ukey >> mv.pageShift)
	index = int(ukey & ((uint64(1) << mv.pageShift) - 1))
	return
}

// Assign assigns the entire slice at the given key's page.
func (mv MapVector[K, T]) Assign(key K, array []T) {
	page, _ := mv.PageAndIndex(key)
	mv.pages[page] = array
}

// Assign assigns the entire slice at the given key's page.
func (mv MapVector[K, T]) Page(page K) []T {
	slice, ok := mv.pages[page]
	if !ok {
		slice = make([]T, mv.PageSize)
		mv.pages[page] = slice
	}
	return slice
}

// Get retrieves the value at the given key.
// Returns the zero value of T and false if not found.
func (mv MapVector[K, T]) Get(key K) (T, bool) {
	page, index := mv.PageAndIndex(key)
	slice, ok := mv.pages[page]
	if !ok || index >= len(slice) {
		var zero T
		return zero, false
	}
	return slice[index], true
}

// Set sets the value at the given key.
// Creates the page slice if it doesn't exist, of size PageSize.
// Does nothing if index >= PageSize.
func (mv MapVector[K, T]) Set(key K, value T) {
	page, index := mv.PageAndIndex(key)
	slice, ok := mv.pages[page]
	if !ok {
		slice = make([]T, mv.PageSize)
		mv.pages[page] = slice
	}
	if index >= len(slice) {
		return
	}
	slice[index] = value
}

// Delete removes the value at the given key.
// Does nothing if the key doesn't exist.
func (mv MapVector[K, T]) Delete(key K) {
	page, index := mv.PageAndIndex(key)
	slice, ok := mv.pages[page]
	if !ok || index >= len(slice) {
		return
	}
	// To keep it simple, set to zero value. For sparse, perhaps use a map instead.
	var zero T
	slice[index] = zero
	// Optionally, shrink if all after are zero, but complicated.
}

// Len returns the total number of elements across all pages.
// Note: This counts allocated slots, not just set ones.
func (mv MapVector[K, T]) Len() int {
	total := 0
	for _, slice := range mv.pages {
		total += len(slice)
	}
	return total
}

// All returns an iterator over all values in the MapVector.
func (mv MapVector[K, T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, page := range mv.pages {
			for _, v := range page {
				if !yield(v) {
					return
				}
			}
		}
	}
}
