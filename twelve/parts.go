package twelve

import (
	"io"
)

func PartOne(r io.Reader) int {
	world := ParseGarden(r)
	regions := world.ParseRegions()

	total := 0

	for _, reg := range regions {
		total += reg.Price()
	}

	return total
}

func PartTwo(r io.Reader) int {
	world := ParseGarden(r)
	regions := world.ParseRegions()

	total := 0

	for _, reg := range regions {
		total += reg.BulkPrice()
	}

	return total
}
