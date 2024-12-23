package sixteen

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

// A Tile is a basic block of the warehouse.
type Tile rune

const (
	Empty Tile = '.'
	Wall       = '#'
	Start      = 'S'
	End        = 'E'
	//Left       = '>'
	//Right      = '<'
	//Up         = '^'
	//Down       = 'v'
)

type Direction int

const (
	Up = iota
	Right
	Down
	Left
)

func (d Direction) Left() Direction {
	return Direction(d - 1&3)
}

func (d Direction) Right() Direction {
	return Direction(d + 1&3)
}

// Maze contains the current simulation state.
type Maze struct {
	Map   [][]Tile
	cache [][]int
}

// Used to store row/col
type Point struct {
	r, c int
}

func (p Point) Next(d Direction) Point {
	switch d {
	case Up:
		return Point{p.r - 1, p.c}
	case Right:
		return Point{p.r, p.c + 1}
	case Left:
		return Point{p.r, p.c - 1}
	case Down:
		return Point{p.r + 1, p.c}
	}

	return p
}

func (m Maze) dfs(pos Point, d Direction, cost int) int {
	if m.Map[pos.r][pos.c] == End {
		return cost
	}

	m.cache[pos.r][pos.c] = min(m.cache[pos.r][pos.c], cost)

	fmt.Printf("Start at %d, %d\n", pos.r, pos.c)

	minimum := math.MaxInt

	// Continue straight
	s := pos.Next(d)
	if (m.Map[s.r][s.c] == Empty || m.Map[s.r][s.c] == End) && m.cache[s.r][s.c] > cost-1001 {
		minimum = min(minimum, m.dfs(s, d, cost+1))
	}

	// Turn left
	ld := d.Left()
	l := pos.Next(ld)
	if (m.Map[l.r][l.c] == Empty || m.Map[l.r][l.c] == End) && m.cache[l.r][l.c] > cost {
		minimum = min(minimum, m.dfs(l, ld, cost+1001))
	}

	// Turn right
	rd := d.Right()
	r := pos.Next(rd)
	if (m.Map[r.r][r.c] == Empty || m.Map[r.r][r.c] == End) && m.cache[r.r][r.c] > cost {
		minimum = min(minimum, m.dfs(r, rd, cost+1001))
	}

	return minimum
}

func (m Maze) Solve() int {
	return m.dfs(Point{len(m.Map) - 2, 1}, Right, 0)
}

/*
ParseMaze reads an input containing both the the floor state as well as
instructions for the robot. When wide = true, all blocks are doubled (Part 2)
*/
func ParseMaze(r io.Reader) *Maze {
	var m [][]Tile
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, []Tile(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Populate a full cache
	cache := make([][]int, len(m))
	for r := range cache {
		cache[r] = make([]int, len(m[r]))
		for c := range cache {
			cache[r][c] = math.MaxInt
		}
	}

	return &Maze{
		Map:   m,
		cache: cache,
	}
}
