//go:build debug
// +build debug

package gx

func Assert(test func() bool, msg string) {
	if !test() {
		panic(msg)
	}
}
