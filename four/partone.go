package four

import (
	"fmt"
)

func PartOne() {
	words := ReadMatrix("./input4")
	total := 0

	total += WordSearch(words, "XMAS")

	fmt.Printf("Solution to part one is: %d\n", total)
}

func PartTwo() {
	words := ReadMatrix("./input4")
	total := 0

	total += CrossSearch(words)

	fmt.Printf("Solution to part two is: %d\n", total)
}
