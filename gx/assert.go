//go:build !debug
// +build !debug

package gx

func Assert(test func() bool, msg string) {
	// TODO why debug does not work?
	if !test() {
		panic(msg)
	}
}
