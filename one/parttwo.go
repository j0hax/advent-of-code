package one

import (
	"io"
)

func countOccurences(target int, s []int) int {
	count := 0
	for _, n := range s {
		if target == n {
			count++
		}
	}

	return count
}

func PartTwo(r io.Reader) int {
	left, right := parseList(r)

	solution := 0

	// Now we "zip" the values
	for i := 0; i < len(left); i++ {
		occ := countOccurences(left[i], right)
		solution += left[i] * occ
	}

	return solution
}
