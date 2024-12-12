package twelve

import (
	"bufio"
	"fmt"
	"io"
)

type Garden [][]rune

type Point struct {
	r, c int
}

type Region struct {
	char         rune
	memberPoints map[Point]struct{}
}

func (r *Region) Perimeter() int {
	totalP := 0

	// Look at each neighbor of the point.
	for p := range r.memberPoints {
		if _, ok := r.memberPoints[Point{p.r + 1, p.c}]; !ok {
			totalP += 1
		}
		if _, ok := r.memberPoints[Point{p.r - 1, p.c}]; !ok {
			totalP += 1
		}
		if _, ok := r.memberPoints[Point{p.r, p.c + 1}]; !ok {
			totalP += 1
		}
		if _, ok := r.memberPoints[Point{p.r, p.c - 1}]; !ok {
			totalP += 1
		}
	}

	return totalP
}

func (r *Region) Area() int {
	return len(r.memberPoints)
}

func (r Region) Price() int {
	a := r.Area()
	p := r.Perimeter()

	price := a * p

	fmt.Printf("[%c] %d * %d = %d\n", r.char, a, p, price)
	return price
}

func (g Garden) InBounds(r, c int) bool {
	return r >= 0 && r < len(g) && c >= 0 && c < len(g[r])
}

func (g Garden) floodFill(re *Region, r, c int) {
	// Ensure the row/col is in bounds and correct type
	if !g.InBounds(r, c) || g[r][c] != re.char {
		return
	}

	// Ensure point is not known (save recursion)
	if _, ok := re.memberPoints[Point{r, c}]; ok {
		return
	}

	// Add to maps
	re.memberPoints[Point{r, c}] = struct{}{}

	// Recurse
	g.floodFill(re, r-1, c)
	g.floodFill(re, r+1, c)
	g.floodFill(re, r, c-1)
	g.floodFill(re, r, c+1)
}

func (g Garden) ParseRegions() []Region {
	knownPoints := make(map[Point]struct{})
	var reg []Region

	// Loop through each part
	for r := range g {
		for c := range g[r] {
			// Skip points in
			if _, ok := knownPoints[Point{r, c}]; ok {
				continue
			}

			// Floodfill with our data structure
			re := &Region{
				char:         g[r][c],
				memberPoints: make(map[Point]struct{}),
			}

			g.floodFill(re, r, c)

			// Copy discovered points
			for k := range re.memberPoints {
				knownPoints[k] = struct{}{}
			}

			reg = append(reg, *re)
		}
	}

	return reg
}

func ParseGarden(r io.Reader) Garden {
	var gard Garden

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()
		gard = append(gard, []rune(t))
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return gard
}
