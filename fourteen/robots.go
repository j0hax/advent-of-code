package fourteen

import (
	"bufio"
	"fmt"
	"io"
)

type Robot struct {
	x, y   int
	vx, vy int
}

type Floor struct {
	robots        []Robot
	width, height int
}

func (f *Floor) RobotGrid() [][]int {
	// Instantiate
	g := make([][]int, f.width)
	for c := range g {
		g[c] = make([]int, f.height)
	}

	for _, r := range f.robots {
		g[r.x][r.y]++
	}

	return g
}

func (r *Robot) Move(w, h int) {
	r.x = (((r.x + r.vx) % w) + w) % w
	r.y = (((r.y + r.vy) % h) + h) % h
}

func (f *Floor) Step() {
	for r := range f.robots {
		f.robots[r].Move(f.width, f.height)
	}
}

func ParseRobots(w, h int, r io.Reader) *Floor {
	var robs []Robot
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var rob Robot
		line := scanner.Text()
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &rob.x, &rob.y, &rob.vx, &rob.vy)
		robs = append(robs, rob)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return &Floor{
		robots: robs,
		width:  w,
		height: h,
	}
}
