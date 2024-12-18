package eighteen

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Represents a "byte"
type Cell int

const (
	Safe Cell = iota
	Corrupt
	Step
)

const SIZE = 71

type Point struct {
	r, c int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.c, p.r)
}

type QueueItem struct {
	p Point
	d int
}

type Memory struct {
	grid    [][]Cell
	visited [][]bool
	q       []QueueItem
	lastLoc Point
}

func (m Memory) InBounds(p Point) bool {
	return p.r >= 0 && p.r < len(m.grid) && p.c >= 0 && p.c < len(m.grid[p.r])
}

func (m Memory) Valid(p Point) bool {
	if !m.InBounds(p) {
		return false
	}

	return m.grid[p.r][p.c] == Safe && !m.visited[p.r][p.c]
}

func (m *Memory) Visit(qp QueueItem) {
	if m.Valid(qp.p) {
		m.q = append(m.q, qp)
		m.visited[qp.p.r][qp.p.c] = true
	}
}

func (m *Memory) VisitNeighbors(qp QueueItem) {
	dist := qp.d + 1

	c := Point{qp.p.r, qp.p.c - 1}
	m.Visit(QueueItem{p: c, d: dist})

	c = Point{qp.p.r, qp.p.c + 1}
	m.Visit(QueueItem{p: c, d: dist})

	c = Point{qp.p.r - 1, qp.p.c}
	m.Visit(QueueItem{p: c, d: dist})

	c = Point{qp.p.r + 1, qp.p.c}
	m.Visit(QueueItem{p: c, d: dist})
}

func (m *Memory) Solve() int {
	// Add initial position
	m.q = append(m.q, QueueItem{Point{0, 0}, 0})
	m.visited[0][0] = true

	for len(m.q) > 0 {
		curr := m.q[0]
		if curr.p.r == SIZE-1 && curr.p.c == SIZE-1 {
			return curr.d
		}
		m.q = m.q[1:]
		m.VisitNeighbors(curr)
	}

	return -1
}

func (m Memory) String() string {
	var sb strings.Builder
	for r := range m.grid {
		for c := range m.grid[r] {
			switch m.grid[r][c] {
			case Safe:
				sb.WriteRune('.')
			case Corrupt:
				sb.WriteRune('#')
			case Step:
				sb.WriteRune('O')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func ParseRAM(r io.Reader, byteCount int) Memory {
	var points []Point
	count := 0
	// Build a list of points
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		count++
		var p Point
		fmt.Sscanf(scanner.Text(), "%d,%d", &p.c, &p.r)
		points = append(points, p)
		if count >= byteCount {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Create, then populate the Memory
	var m Memory
	m.grid = make([][]Cell, SIZE)
	m.visited = make([][]bool, SIZE)
	for r := range m.grid {
		m.grid[r] = make([]Cell, SIZE)
		m.visited[r] = make([]bool, SIZE)
	}

	for _, p := range points {
		m.grid[p.r][p.c] = Corrupt
	}

	m.lastLoc = points[len(points)-1]

	return m
}
