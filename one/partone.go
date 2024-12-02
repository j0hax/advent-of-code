package one

import (
	"fmt"
	"math"
)

func PartOne() {
	left, right := parseList("./input1")

	solution := 0

	// Now we "zip" the values
	for i := 0; i < len(left); i++ {
		s := float64(left[i]) - float64(right[i])
		solution += int(math.Abs(s))
	}

	fmt.Printf("Solution for Part One: %d\n", solution)
}
