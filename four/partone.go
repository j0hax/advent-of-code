package four

import (
	"io"
)

func PartOne(r io.Reader) int {
	words := ReadWordSearch(r)
	total := 0

	total += words.Count("XMAS")

	return total

}

func PartTwo(r io.Reader) int {
	words := ReadWordSearch(r)
	total := 0

	total += words.CrossCount("MAS")

	return total
}
