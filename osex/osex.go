package osex

import (
	"log"
	"os"
)

func CopyFile(source, destination string) error {
	r, err := os.Open(source)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer w.Close()

	w.ReadFrom(r)
	return nil
}

// isDirectory reports whether the named file is a directory.
func IsDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}
