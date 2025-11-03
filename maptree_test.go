package goex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapTree_Clone(t *testing.T) {
	tree := MapTree[string, any]{
		"a": 1,
		"b": MapTree[string, any]{
			"c": "value",
		},
	}

	clone := tree.Clone().(MapTree[string, any])
	assert.Equal(t, tree, clone)

	// Modify clone shouldn't affect original
	clone["a"] = 2
	assert.NotEqual(t, tree["a"], clone["a"])
}

func TestMapTree_Contains(t *testing.T) {
	tree := MapTree[string, any]{
		"a": MapTree[string, any]{
			"b": MapTree[string, any]{
				"c": "value",
			},
		},
	}

	tests := []struct {
		path []string
		want bool
	}{
		{[]string{}, true},
		{[]string{"a"}, true},
		{[]string{"a", "b"}, true},
		{[]string{"a", "b", "c"}, true},
		{[]string{"x"}, false},
		{[]string{"a", "x"}, false},
		{[]string{"a", "b", "x"}, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, tree.Contains(tt.path...))
	}
}

func TestMapTree_Get(t *testing.T) {
	tree := MapTree[string, string]{
		"a": "value1",
		"b": MapTree[string, string]{
			"c": "value2",
		},
	}

	tests := []struct {
		path   []string
		want   string
		wantOk bool
	}{
		{[]string{"a"}, "value1", true},
		{[]string{"b", "c"}, "value2", true},
		{[]string{"x"}, "", false},
		{[]string{"b", "x"}, "", false},
	}

	for _, tt := range tests {
		got, ok := tree.Get(tt.path...)
		assert.Equal(t, tt.wantOk, ok)
		if ok {
			assert.Equal(t, tt.want, got)
		}
	}
}

func TestMapTree_GetAny(t *testing.T) {
	tree := MapTree[string, any]{
		"a": 1,
		"b": MapTree[string, any]{
			"c": "value",
		},
	}

	tests := []struct {
		path   []string
		want   any
		wantOk bool
	}{
		{[]string{"a"}, 1, true},
		{[]string{"b", "c"}, "value", true},
		{[]string{"x"}, nil, false},
		{[]string{"b", "x"}, nil, false},
	}

	for _, tt := range tests {
		got, ok := tree.GetAny(tt.path...)
		assert.Equal(t, tt.wantOk, ok)
		if ok {
			assert.Equal(t, tt.want, got)
		}
	}
}

func TestMapTree_Set(t *testing.T) {
	tree := MapTree[string, any]{}

	tests := []struct {
		path  []string
		value any
		want  bool
	}{
		{[]string{"a"}, 1, true},
		{[]string{"b", "c"}, "value", true},
		{[]string{}, nil, false},
	}

	for _, tt := range tests {
		ok := tree.Set(tt.value, tt.path...)
		assert.Equal(t, tt.want, ok)
		if ok {
			v, exists := tree.GetAny(tt.path...)
			assert.True(t, exists)
			assert.Equal(t, tt.value, v)
		}
	}
}

func TestMapTree_Merge(t *testing.T) {
	tree1 := MapTree[string, any]{
		"a": 1,
		"b": MapTree[string, any]{
			"c": "value1",
		},
	}

	tree2 := MapTree[string, any]{
		"a": 2,
		"b": MapTree[string, any]{
			"c": "value2",
			"d": "value3",
		},
	}

	assert.True(t, tree1.Merge(tree2))
	assert.Equal(t, 2, tree1["a"])
	b := tree1["b"].(MapTree[string, any])
	assert.Equal(t, "value2", b["c"])
	assert.Equal(t, "value3", b["d"])

	// Test incompatible merge
	incompatible := MapTree[string, any]{
		"b": "not a map",
	}
	assert.False(t, tree1.Merge(incompatible))
}
