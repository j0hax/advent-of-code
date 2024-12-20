package twenty

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Represents a "byte"
type Block rune

const (
	Empty Block = '.'
	Wall        = '#'
	Start       = 'S'
	End         = 'E'
)

func (b Block) String() string {
	return fmt.Sprintf("%c", b)
}

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

type Track struct {
	grid    [][]Block
	visited [][]bool
	q       []QueueItem
}

func (m Track) CountCheats() map[Point]int {
	// cheatList saves each cheat and the distance saved
	cheatList := make(map[Point]int)

	original := m.SESolve()

	// Now, for each cell, break twice:
	// curr cell + right neighbor,
	// curr cell + bottom neighbor
	for r := range m.grid {
		for c := range m.grid[r] {
			// Can only activate when on normal track
			if m.grid[r][c] != Empty {
				continue
			}

			//start := Point{r, c}

			// Check right
			if m.grid[r][c+1] == Wall {
				m.grid[r][c+1] = Empty
				cheatList[Point{r, c + 1}] = original - m.SESolve()
				m.grid[r][c+1] = Wall
			}

			// Check bottom
			if m.grid[r+1][c] == Wall {
				m.grid[r+1][c] = Empty
				cheatList[Point{r + 1, c}] = original - m.SESolve()
				m.grid[r+1][c] = Wall
			}

			// ******* TOP AND LEFT? ******
			if m.InBounds(Point{r - 1, c}) && m.grid[r-1][c] == Wall {
				m.grid[r-1][c] = Empty
				cheatList[Point{r - 1, c}] = original - m.SESolve()
				m.grid[r-1][c] = Wall
			}

			// Check bottom
			if m.InBounds(Point{r, c - 1}) && m.grid[r][c-1] == Wall {
				m.grid[r][c-1] = Empty
				cheatList[Point{r, c - 1}] = original - m.SESolve()
				m.grid[r][c-1] = Wall
			}
		}
	}

	return cheatList
}

func (m Track) InBounds(p Point) bool {
	return p.r >= 0 && p.r < len(m.grid) && p.c >= 0 && p.c < len(m.grid[p.r])
}

func (m Track) Valid(p Point) bool {
	if !m.InBounds(p) {
		return false
	}

	return m.grid[p.r][p.c] != Wall && !m.visited[p.r][p.c]
}

func (m *Track) Visit(qp QueueItem) {
	if m.Valid(qp.p) {
		m.q = append(m.q, qp)
		m.visited[qp.p.r][qp.p.c] = true
	}
}

func (m *Track) VisitNeighbors(qp QueueItem) {
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

func (m *Track) Solve(start, end Point) int {
	m.Reset()

	// Add initial position
	m.q = append(m.q, QueueItem{start, 0})
	m.visited[start.r][start.c] = true

	for len(m.q) > 0 {
		curr := m.q[0]
		if curr.p == end {
			return curr.d
		}
		m.q = m.q[1:]
		m.VisitNeighbors(curr)
	}

	return 0
}

func (m *Track) SESolve() int {
	startPos := m.FindBlock(Start)
	endPos := m.FindBlock(End)

	return m.Solve(startPos, endPos)
}

func (m Track) FindBlock(b Block) Point {
	for r := range m.grid {
		for c := range m.grid[r] {
			if m.grid[r][c] == b {
				return Point{r, c}
			}
		}
	}

	return Point{-1, -1}
}

func (m Track) String() string {
	var sb strings.Builder
	for r := range m.grid {
		for c := range m.grid[r] {
			sb.WriteRune(rune(m.grid[r][c]))
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

// Reset resets the visited matrix and clears the queue
func (m *Track) Reset() {
	for r := range m.visited {
		for c := range m.visited[r] {
			m.visited[r][c] = false
		}
	}

	m.q = make([]QueueItem, 0)
}

func ParseTrack(r io.Reader) Track {
	var m Track
	// Build a list of points
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		m.grid = append(m.grid, []Block(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Create list of visited locations
	m.visited = make([][]bool, len(m.grid))
	for r := range m.visited {
		m.visited[r] = make([]bool, len(m.grid[r]))
	}

	return m
}
