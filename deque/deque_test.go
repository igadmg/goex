package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeDeque(t *testing.T) {
	deque := Make[int](CacheLinePadSize)
	assert.Equal(t, CacheLinePadSize, deque.page_size)
	assert.Equal(t, CacheLinePadSize/8, deque.items_per_page)
	assert.Empty(t, deque.pages)
}

func TestItem(t *testing.T) {
	deque := Make[int](CacheLinePadSize)

	// Test when no items exist
	item := deque.Item(0)
	assert.Nil(t, item)

	// Test when items exist
	deque = deque.Append(1).Append(2).Append(3)
	item = deque.Item(1)
	assert.NotNil(t, item)
	assert.Equal(t, 2, *item)
}

func TestReserve(t *testing.T) {
	deque := Make[int](CacheLinePadSize)

	// Reserve space for capacity 15
	deque = deque.Reserve(16)

	item := deque.Item(15)
	assert.NotNil(t, item)
	assert.Equal(t, 0, *item)

	item = deque.Item(10)
	assert.NotNil(t, item)
	assert.Equal(t, 0, *item)

	item = deque.Item(4)
	assert.NotNil(t, item)
	assert.Equal(t, 0, *item)

	// Reserve space for capacity 31
	deque = deque.Reserve(31)

	item = deque.Item(30)
	assert.NotNil(t, item)
	assert.Equal(t, 0, *item)
}

func TestAppend(t *testing.T) {
	deque := Make[int](CacheLinePadSize)

	// Append items
	deque = deque.Append(1).Append(2).Append(3)
	//assert.Len(t, deque.Pages(), 2) // Two pages should exist
	//assert.Equal(t, []int{1, 2}, deque.Pages()[0])
	//assert.Equal(t, []int{3}, deque.Pages()[1])
}
