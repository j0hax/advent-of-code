package six

import (
	"io"
)

func PartOne(r io.Reader) int {
	world := Load(r)
	for world.Step() != oob {

	}
	return world.Count(visited)
}

type point struct {
	x, y int
}

func PartTwo(r io.Reader) int {
	original := Load(r)

	world := original.Copy()

	loopPossibilities := 0

	for r := range world {
		for c := range world[r] {

			world = original.Copy()

			// Change one litle thing
			orig := world[r][c]

			if orig != empty {
				continue
			}

			world[r][c] = wall

			iterations := 0

			// Begin world simulation
			for {
				if world.Step() != oob {
					iterations++
					if iterations > 20000 {
						loopPossibilities++
						break
					}
				} else {
					break
				}
			}

			// Change it back
			world[r][c] = orig
		}
	}

	return loopPossibilities
}
