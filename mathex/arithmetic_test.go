package mathex

import (
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b   int
		result int
	}{
		{48, 18, 6},
		{18, 48, 6},
		{7, 13, 1},
		{0, 5, 5},
		{5, 0, 5},
		{0, 0, 0},
		//{-48, 18, 6},
		//{48, -18, 6},
		//{-48, -18, 6},
	}

	for _, test := range tests {
		if res := GCD(test.a, test.b); res != test.result {
			t.Errorf("GCD(%d, %d) = %d; want %d", test.a, test.b, res, test.result)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		rest   []int
		result int
	}{
		{[]int{25, 5, 13, 7}, 2275},
		{[]int{1, 2, 3, 11, 14, 16}, 3696},
		{[]int{1, 2, 3, 14, 16}, 336},
		{[]int{4, 6}, 12},
		{[]int{6, 4}, 12},
		{[]int{7, 13}, 91},
		{[]int{13, 7}, 91},
		{[]int{0, 5}, 5},
		{[]int{5, 0}, 5},
		{[]int{25, 5}, 25},
		{[]int{5, 25}, 25},
		{[]int{0, 0}, 0},
		{[]int{-4, 6}, 12},
		{[]int{4, -6}, 12},
		{[]int{-4, -6}, 12},
	}

	for _, test := range tests {
		if res := LCMn(test.rest...); res != test.result {
			t.Errorf("LCM(%v) = %d; want %d", test.rest, res, test.result)
		}
	}
}
