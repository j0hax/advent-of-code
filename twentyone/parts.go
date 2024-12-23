package twentyone

import (
	"io"
)

func PartOne(r io.Reader) int {
	codes := ParseCodes(r)

	res := 0
	for _, c := range codes {
		res += c.Solve(2)
	}

	return res
}

func PartTwo(r io.Reader) int {
	codes := ParseCodes(r)

	res := 0
	for _, c := range codes {
		res += c.Solve(25)
	}

	return res
}
