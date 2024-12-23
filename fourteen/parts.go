package fourteen

import (
	"fmt"
	"io"
)

func PartOne(r io.Reader) int {
	f := ParseRobots(101, 103, r)

	// Simulate 100 seconds
	for range 100 {
		f.Step()
	}

	// Get positions/counts
	g := f.RobotGrid()

	quads := Quadrants(g)

	answer := 1
	for q := range quads {
		s := GridSum(quads[q])
		answer *= s
	}

	return answer
}

/*
Because I had no Idea what to expect, part two took a brute-force approach in
which I just dumped 10k little PNGs of the field.

It turns out my input is 7572.

TODO: build something to detect low entropy (=ordered pixels) automatically.
*/
func PartTwo(r io.Reader) int {
	robots := ParseRobots(101, 103, r)
	counter := 0

	for {
		path := fmt.Sprintf("/tmp/step%d.png", counter)
		if err := robots.ToImage(path); err != nil {
			panic(err)
		}
		robots.Step()
		counter++

		if counter == 10000 {
			break
		}
	}

	return 0
}
