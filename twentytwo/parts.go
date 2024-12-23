package twentytwo

import (
	"io"
	"maps"
)

// Trivial.
func PartOne(r io.Reader) int {
	n := ParseNumbers(r)

	sum := 0
	for _, number := range n {
		for range 2000 {
			number = number.Next()
		}

		sum += int(number)
	}

	return sum
}

func PartTwo(r io.Reader) int {
	n := ParseNumbers(r)
	var p []map[ProfitHistory]int

	for _, number := range n {
		p = append(p, number.ProfitDevelopment())
	}

	// Copy sequence keys
	seqs := make(map[ProfitHistory]struct{})
	for _, profits := range p {
		for key := range maps.Keys(profits) {
			seqs[key] = struct{}{}
		}
	}

	// here things get wonky?
	best := 0
	for seq := range maps.Keys(seqs) {
		total := TotalProfit(p, seq)
		best = max(best, total)
	}

	return best
}
