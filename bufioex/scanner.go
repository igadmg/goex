package bufioex

import (
	"bufio"
	"io"
)

// Use bufio.Scanner to read from r and send each scanned line to the provided channel.
func ScanToChan(r io.Reader, ch chan<- string) func() {
	return func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}
}
