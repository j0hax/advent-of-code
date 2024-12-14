package fourteen

import (
	"fmt"
	"io"
)

func PartOne(r io.Reader) int {
	f := ParseRobots(101, 103, r)

	orig := f.RobotGrid()
	PrintGrid(orig)

	fmt.Println()

	// Simulate 100 seconds
	for i := 0; i < 100; i++ {
		f.Step()
	}

	// Get positions/counts
	g := f.RobotGrid()
	fmt.Println("After stepping:")
	PrintGrid(g)

	quads := Quadrants(g)

	for quad, q := range quads {
		fmt.Printf("Quadrant %d:\n", quad)
		PrintGrid(q)
	}

	answer := 1
	for q := range quads {
		s := GridSum(quads[q])
		fmt.Printf("Sum is %d\n", s)
		answer *= s
	}

	return answer
}

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
