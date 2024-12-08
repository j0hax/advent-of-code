package eight

import (
	"bufio"
	"fmt"
	"io"
	"slices"
)

type Signal rune

// IsEmpty returns true if the character equals '.'
func (s Signal) IsEmpty() bool {
	return s == '.'
}

// SignalMap represents the map of day 8
type SignalMap [][]Signal

// InBounds checks if the specified row and column fit into the SignalMap
func (s SignalMap) InBounds(p Point) bool {
	return p.r >= 0 && p.r < len(s) && p.c >= 0 && p.c < len(s[p.r])
}

func (s SignalMap) Print(markers []Point) {
	for r := range s {
		for c := range s[r] {
			if slices.Contains(markers, Point{r, c}) {
				fmt.Print("#")
			} else {
				fmt.Print(string(s[r][c]))
			}
		}
		fmt.Print("\n")
	}
}

// Point represents a row and column.
// it is used for hashmap locations
type Point struct {
	r, c int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.r, p.c)
}

// Add adds dr and dc to the row and columns of the point, respectively
func (p *Point) Add(dr, dc int) {
	p.r += dr
	p.c += dc
}

// AntinodesTo calculates the line-of-sight antinodes between two points
func (s SignalMap) AntinodesTo(a, b Point, includeOwn bool) []Point {
	var antinodes []Point

	// Can not be the same point
	if (a.r == b.r) && (a.c == b.c) {
		if includeOwn {
			antinodes = append(antinodes, a)
		}
		return antinodes
	}

	// Calculate the difference vector between two points
	dr := a.r - b.r
	dc := a.c - b.c

	a.Add(dr, dc)

	for s.InBounds(a) {
		antinodes = append(antinodes, a)
		a.Add(dr, dc)
	}

	return antinodes
}

func (s SignalMap) FindAntennas(a Signal) []Point {
	var knownPoints []Point
	for r := range s {
		for c := range s[r] {
			if s[r][c] == a {
				knownPoints = append(knownPoints, Point{r, c})
			}
		}
	}

	return knownPoints
}

func (s SignalMap) CountAntinodes(res bool) int {
	knownLocations := make(map[Point]bool)

	for r := range s {
		for c := range s[r] {
			if s[r][c].IsEmpty() {
				continue
			}
			currPoint := Point{r, c}
			// find all antennas with the same frequency
			inLineAntennas := s.FindAntennas(s[r][c])
			// for each of the in-line antennas, find an antinode
			for _, a := range inLineAntennas {
				antinodes := s.AntinodesTo(currPoint, a, res)

				// Part one runs only once (res = false)
				for _, p := range antinodes {
					knownLocations[p] = true
					if !res {
						break
					}
				}
			}
		}
	}

	keys := make([]Point, len(knownLocations))
	i := 0
	for k := range knownLocations {
		keys[i] = k
		i++
	}

	s.Print(keys)
	fmt.Println("---")

	return len(knownLocations)
}

// ParseMap reads data into a SignalMap
func ParseMap(r io.Reader) SignalMap {
	var smap SignalMap

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()
		smap = append(smap, []Signal(t))
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return smap
}
