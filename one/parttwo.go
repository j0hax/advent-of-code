package one

import (
	"fmt"
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

func PartTwo() {
	left, right := parseList("./input1")

	solution := 0

	// Now we "zip" the values
	for i := 0; i < len(left); i++ {
		occ := countOccurences(left[i], right)
		solution += left[i] * occ
	}

	fmt.Printf("Solution for Part Two: %d\n", solution)
}
