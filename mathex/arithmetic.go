package mathex

import (
	"golang.org/x/exp/constraints"
)

// GCD вычисляет наибольший общий делитель двух чисел с помощью алгоритма Евклида.
func GCD[T constraints.Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM вычисляет наименьшее общее кратное двух чисел.
func LCMn[T constraints.Integer](rest ...T) T {
	if len(rest) == 0 {
		return 0
	}
	if len(rest) == 1 {
		return rest[0]
	}

		rest[len(rest)-2] = LCM(rest[len(rest)-1], rest[len(rest)-2])
		return LCMn(rest[:len(rest)-1]...)
}

func LCM[T constraints.Integer](a, b T) T {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return (a * b) / GCD(a, b)
}

// LCM вычисляет наименьшее общее кратное двух чисел.
func WeightedLCM[T constraints.Integer](rest ...T) (gdc T, w []float32) {
	gdc = LCMn(rest...)
	w = DivideF(gdc, rest...)
	return
}
