package yamlex

import (
	"os"

	"go.yaml.in/yaml/v4"
)

func MarshalFile[T any](path string, data T) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	e := yaml.NewEncoder(file)
	defer e.Close()

	if err := e.Encode(data); err != nil {
		return err
	}

	return nil
}

func UnmarshalFile[T any](path string) (data T, err error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	l, err := yaml.NewLoader(file)
	if err != nil {
		return
	}

	err = l.Load(&data)
	return
}
