package osex

import (
	"os"
)

// reports whether the named path is a directory.
func IsDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		return false
	}
	return info.Mode().IsDir()
}

// reports whether the named path is a file.
func IsFile(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		return false
	}
	return info.Mode().IsRegular()
}

// Chdir to newWd and return deferable to switch back to old one.
func Switchdir(newWd string) func() {
	oldWd, err := os.Getwd()
	if err != nil {
		return func() {}
	}

	if err = os.Chdir(newWd); err != nil {
		return func() {}
	}

	return func() {
		os.Chdir(oldWd)
	}
}
