package cache

import (
	"github.com/igadmg/goex/mathex"
	"github.com/igadmg/goex/unsafeex"
	"golang.org/x/sys/cpu"
)

var CacheLinePadSize = int(unsafeex.Sizeof[cpu.CacheLinePad]())

func PerfectPageSize[T any]() int {
	return mathex.LCM(CacheLinePadSize, int(unsafeex.Sizeof[T]()))
}
