package go_koans

import (
	"fmt"
	"os"
)

func divideFourBy(i int) int {
	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
	// panics are exceptional errors at runtime

	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Panic: %v\n", err)
		}
	}()

	n := divideFourBy(__divisor__)
	assert(n == 2) // panics are exceptional errors at runtime
}
