package unsafeex

import "unsafe"

func Sizeof[T any]() uintptr {
	var v T
	return unsafe.Sizeof(v)
}
