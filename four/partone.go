package four

import (
	"fmt"
)

func PartOne() {
	words := ReadWordSearch("./input4")
	total := 0

	total += words.Count("XMAS")

	fmt.Printf("Solution to part one is: %d\n", total)
}

func PartTwo() {
	words := ReadWordSearch("./input4")
	total := 0

	total += words.CrossCount("MAS")

	fmt.Printf("Solution to part two is: %d\n", total)
}
