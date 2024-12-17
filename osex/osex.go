package osex

import "os"

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
