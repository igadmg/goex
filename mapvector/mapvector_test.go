package mapvector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TypeID constants for test data generation
const (
	TypeA TypeID = iota
	TypeB
	TypeC
	TypeD
	TypeE
	TypeF
	// 6 types
)

type TypeID int

// MakeItem creates a test item value based on type and index
func MakeItem(typeID TypeID, index int) int {
	return int(typeID)*1000 + index
}

// MakeStringItem creates a test string item
func MakeStringItem(typeID TypeID, index int) string {
	return string(rune('A'+int(typeID))) + string(rune('0'+index))
}

// TestData represents a key-value pair for testing
type TestData[T any] struct {
	Key  int
	Item T
}

// populateMapVector sets multiple key-value pairs in the MapVector
func populateMapVector[T any](mv MapVector[int, T], data []TestData[T]) {
	for _, d := range data {
		mv.Set(d.Key, d.Item)
	}
}

// verifyMapVector checks that all key-value pairs are correctly set
func verifyMapVector[T any](t *testing.T, mv MapVector[int, T], data []TestData[T]) {
	for _, d := range data {
		val, ok := mv.Get(d.Key)
		assert.True(t, ok, "Expected key %d to exist", d.Key)
		assert.Equal(t, d.Item, val, "Expected item for key %d", d.Key)
	}
}

func TestMakeMapVector(t *testing.T) {
	mv := MakeMapVector[int, int]()
	assert.NotNil(t, mv.pages, "Expected pages map to be initialized")
	assert.Empty(t, mv.pages, "Expected empty pages map")
}

func TestSetAndGet(t *testing.T) {
	mv := MakeMapVector[int, int]()
	testData := []TestData[int]{
		{Key: mv.MakeKey(int(TypeA), 0), Item: MakeItem(TypeA, 0)},
		{Key: mv.MakeKey(int(TypeA), 1), Item: MakeItem(TypeA, 1)},
		{Key: mv.MakeKey(int(TypeA), 10), Item: MakeItem(TypeA, 10)},
		{Key: mv.MakeKey(int(TypeB), 0), Item: MakeItem(TypeB, 0)},
		{Key: mv.MakeKey(int(TypeB), 5), Item: MakeItem(TypeB, 5)},
		{Key: mv.MakeKey(int(TypeC), 0), Item: MakeItem(TypeC, 0)},
		{Key: mv.MakeKey(int(TypeC), 20), Item: MakeItem(TypeC, 20)},
		{Key: mv.MakeKey(int(TypeD), 1), Item: MakeItem(TypeD, 1)},
		{Key: mv.MakeKey(int(TypeE), 2), Item: MakeItem(TypeE, 2)},
		{Key: mv.MakeKey(int(TypeF), 3), Item: MakeItem(TypeF, 3)},
	}

	populateMapVector(mv, testData)
	verifyMapVector(t, mv, testData)
}

func TestGetNonExistent(t *testing.T) {
	mv := MakeMapVector[int, int]()
	val, ok := mv.Get(1)
	assert.False(t, ok, "Expected key 1 to not exist")
	assert.Equal(t, 0, val, "Expected zero value")
}

func TestDelete(t *testing.T) {
	mv := MakeMapVector[int, int]()
	mv.Set(0, 42)
	mv.Delete(0)
	val, ok := mv.Get(0)
	assert.True(t, ok, "Expected key 0 to still exist after delete (set to zero)")
	assert.Equal(t, 0, val, "Expected zero value after delete")
}

func TestDeleteNonExistent(t *testing.T) {
	mv := MakeMapVector[int, int]()
	mv.Delete(1) // Should not panic
}

func TestLen(t *testing.T) {
	mv := MakeMapVector[int, int]()
	assert.Equal(t, 0, mv.Len(), "Expected len 0")

	// Set data across multiple pages
	keys := []int{
		mv.MakeKey(int(TypeA), 0), mv.MakeKey(int(TypeA), 1), mv.MakeKey(int(TypeA), 5),
		mv.MakeKey(int(TypeB), 0), mv.MakeKey(int(TypeB), 1),
		mv.MakeKey(int(TypeC), 0),
	}
	for _, key := range keys {
		mv.Set(key, MakeItem(TypeA, 0)) // item value doesn't matter for len
	}

	// Len() sums len(slice) for each page: 3 pages * PageSize
	expectedLen := 3 * mv.PageSize
	assert.Equal(t, expectedLen, mv.Len(), "Expected len 3*PageSize")
}

func TestMultiplePages(t *testing.T) {
	mv := MakeMapVector[int, int]()
	testData := []TestData[int]{
		{Key: mv.MakeKey(int(TypeA), 0), Item: MakeItem(TypeA, 0)},
		{Key: mv.MakeKey(int(TypeA), 100), Item: MakeItem(TypeA, 100)},
		{Key: mv.MakeKey(int(TypeB), 0), Item: MakeItem(TypeB, 0)},
		{Key: mv.MakeKey(int(TypeB), 10), Item: MakeItem(TypeB, 10)},
		{Key: mv.MakeKey(int(TypeB), 50), Item: MakeItem(TypeB, 50)},
		{Key: mv.MakeKey(int(TypeC), 0), Item: MakeItem(TypeC, 0)},
		{Key: mv.MakeKey(int(TypeC), 25), Item: MakeItem(TypeC, 25)},
		{Key: mv.MakeKey(int(TypeD), 1), Item: MakeItem(TypeD, 1)},
		{Key: mv.MakeKey(int(TypeE), 2), Item: MakeItem(TypeE, 2)},
	}

	populateMapVector(mv, testData)
	verifyMapVector(t, mv, testData)
}

func TestGrowSlice(t *testing.T) {
	mv := MakeMapVector[int, int]()
	// Set at index 10 in page 0
	mv.Set(10, 123)
	val, ok := mv.Get(10)
	assert.True(t, ok, "Failed to set and get at higher index")
	assert.Equal(t, 123, val, "Failed to set and get at higher index")
	// Check len of page 0 slice
	slice := mv.pages[0]
	assert.Equal(t, mv.PageSize, len(slice), "Expected slice len PageSize")
}

func TestOverwrite(t *testing.T) {
	mv := MakeMapVector[int, int]()
	key := mv.MakeKey(int(TypeA), 0)
	originalItem := MakeItem(TypeA, 0)
	mv.Set(key, originalItem)
	val, ok := mv.Get(key)
	assert.True(t, ok)
	assert.Equal(t, originalItem, val)
	// Overwrite with new item
	newItem := MakeItem(TypeB, 1) // Different value
	mv.Set(key, newItem)
	val2, ok2 := mv.Get(key)
	assert.True(t, ok2)
	assert.Equal(t, newItem, val2)
}

func TestStringType(t *testing.T) {
	mv := MakeMapVector[int, string]()
	testData := []TestData[string]{
		{Key: mv.MakeKey(int(TypeA), 0), Item: MakeStringItem(TypeA, 0)},
		{Key: mv.MakeKey(int(TypeA), 5), Item: MakeStringItem(TypeA, 5)},
		{Key: mv.MakeKey(int(TypeB), 1), Item: MakeStringItem(TypeB, 1)},
		{Key: mv.MakeKey(int(TypeB), 10), Item: MakeStringItem(TypeB, 10)},
		{Key: mv.MakeKey(int(TypeC), 5), Item: MakeStringItem(TypeC, 5)},
		{Key: mv.MakeKey(int(TypeD), 2), Item: MakeStringItem(TypeD, 2)},
		{Key: mv.MakeKey(int(TypeE), 3), Item: MakeStringItem(TypeE, 3)},
		{Key: mv.MakeKey(int(TypeF), 4), Item: MakeStringItem(TypeF, 4)},
	}

	populateMapVector(mv, testData)
	verifyMapVector(t, mv, testData)
}
