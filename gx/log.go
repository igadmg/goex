package gx

import "fmt"

func LogS[T any](v T) string {
	return fmt.Sprintf("%+v", v)
}
