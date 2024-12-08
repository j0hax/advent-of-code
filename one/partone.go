package one

import (
	"io"
	"math"
)

func PartOne(r io.Reader) int {
	left, right := parseList(r)

	solution := 0

	// Now we "zip" the values
	for i := 0; i < len(left); i++ {
		s := float64(left[i]) - float64(right[i])
		solution += int(math.Abs(s))
	}

	return solution
}
