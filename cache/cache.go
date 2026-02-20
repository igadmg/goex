package cache

import (
	"github.com/Mishka-Squat/goex/mathex"
	"github.com/Mishka-Squat/goex/unsafeex"
	"golang.org/x/sys/cpu"
)

var CacheLinePadSize = int(unsafeex.Sizeof[cpu.CacheLinePad]())

func PerfectPageSize[T any]() int {
	return mathex.LCM(CacheLinePadSize, int(unsafeex.Sizeof[T]()))
}
