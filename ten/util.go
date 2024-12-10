package ten

import (
	"bufio"
	"io"
	"strings"
)

// Map represents the topographical map
type Map [][]int

// Point is used in a hashmap to track previously visited points on a map
type Point struct {
	r, c int
}

func (m Map) String() string {
	var sb strings.Builder

	for _, r := range m {
		for _, c := range r {
			sb.WriteRune(rune(c + '0'))
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

// InBounds checks if a row/col coordinate is within Map boundaries
func (m Map) InBounds(r, c int) bool {
	return r >= 0 && r < len(m) && c >= 0 && c < len(m[r])
}

/*
evalPath recursively counts paths in up, down, left, right directions.

The score (unique, non-crossing) of paths is calculated by setting unique = true
and passing a hashmap to track.

The rating of paths is calculated by setting unique = false, in which all
possible 0...9 trails are calculated.
*/
func (m Map) evalPath(want, r, c int, unique bool, known map[Point]struct{}) int {
	if !m.InBounds(r, c) {
		return 0
	}

	height := m[r][c]

	// Check if height increments correctly
	if height != want {
		return 0
	}

	if unique {
		// Skip known points
		if _, ok := known[Point{r, c}]; ok {
			return 0
		}
		known[Point{r, c}] = struct{}{}
	}

	// Check if end
	if height == 9 {
		return 1
	}

	// Spread out search where the height is one higher
	paths := 0

	// left/right/up/down
	paths += m.evalPath(want+1, r-1, c, unique, known)
	paths += m.evalPath(want+1, r+1, c, unique, known)
	paths += m.evalPath(want+1, r, c+1, unique, known)
	paths += m.evalPath(want+1, r, c-1, unique, known)

	return paths
}

/*
ScoreTrails returns the cumulative "scores", i.e. the number of 9-heights
reachable from a 0 with a single step in height between them.
*/
func (m Map) ScoreTrails() int {
	paths := 0

	for r := range m {
		for c := range m[r] {
			if m[r][c] == 0 {
				known := make(map[Point]struct{})
				score := m.evalPath(0, r, c, true, known)
				paths += score
			}
		}
	}

	return paths
}

/*
GradeTrails returns the cumulative "grades", i.e. the total number of 0...9
paths. In contrast to ScoreTrails, this grade also allows for crossing paths.
*/
func (m Map) GradeTrails() int {
	paths := 0

	for r := range m {
		for c := range m[r] {
			if m[r][c] == 0 {
				score := m.evalPath(0, r, c, false, nil)
				paths += score
			}
		}
	}

	return paths
}

// parseLine is a helper function with returns a slice of integers from a string
func parseLine(line string) []int {
	var nums []int

	for _, c := range line {
		num := int(c - '0')
		nums = append(nums, num)
	}

	return nums
}

// ParseMap returns the parsed Map data structure from an input.
func ParseMap(r io.Reader) Map {
	scanner := bufio.NewScanner(r)

	var m Map

	for scanner.Scan() {
		m = append(m, parseLine(scanner.Text()))
	}

	return m
}
