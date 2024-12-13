package thirteen

import (
	"io"
)

func PartOne(r io.Reader) int {
	machines := ParseMachines(r)

	tokens := 0

	for _, m := range machines {
		cost := m.Win()
		tokens += cost
	}

	return tokens
}

func PartTwo(r io.Reader) int {
	machines := ParseMachines(r)

	tokens := 0

	for _, m := range machines {
		cost := m.AdjWin()
		tokens += cost
	}

	return tokens
}
