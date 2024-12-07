package six

import (
	"bufio"
	"io"
	"strings"
)

type block int
type status int

const (
	empty block = iota
	wall
	guardUp
	guardDown
	guardLeft
	guardRight
	visited
)

func (p block) String() string {
	switch p {
	case empty:
		return "."
	case wall:
		return "#"
	case guardUp:
		return "^"
	case guardDown:
		return "V"
	case guardLeft:
		return "<"
	case guardRight:
		return ">"
	case visited:
		return "X"
	}

	return ""
}

const (
	sucessful status = iota
	oob
	blocked
)

type grid [][]block

func (g grid) Copy() grid {
	newMatrix := make([][]block, len(g))
	for i, row := range g {
		newMatrix[i] = make([]block, len(row))
		copy(newMatrix[i], row)
	}
	return newMatrix
}

func (g grid) String() string {
	var b strings.Builder

	for r := range g {
		for c := range g[r] {
			b.WriteString(g[r][c].String())
		}
		b.WriteRune('\n')
	}

	return b.String()
}

// GuardLocation returns the row/column location of the guard on the map.
// If the guard is off the map, -1,-1 is returned.
func (g grid) GuardLocation() (int, int) {
	for x := range g {
		for y := range g[x] {
			if g[x][y] == guardDown || g[x][y] == guardUp || g[x][y] == guardLeft || g[x][y] == guardRight {
				return x, y
			}
		}
	}

	return -1, -1
}

func (g grid) InBounds(r, c int) bool {
	return r >= 0 && c >= 0 && r < len(g) && c < len(g[r])
}

func (g grid) Count(p block) int {
	cnt := 0
	for r := range g {
		for c := range g[r] {
			if g[r][c] == p {
				cnt++
			}
		}
	}
	return cnt
}

func (g grid) move(fromr, fromc, tor, toc int) status {
	if !g.InBounds(tor, toc) {
		g[fromr][fromc] = visited
		return oob
	}

	switch g[tor][toc] {
	case wall:
		g[fromr][fromc].rotate()
		return blocked
	case empty, visited:
		g[tor][toc], g[fromr][fromc] = g[fromr][fromc], visited
	}

	return sucessful
}

// Rotates a guard 90 degrees
func (p *block) rotate() {
	switch *p {
	case guardDown:
		*p = guardLeft
	case guardLeft:
		*p = guardUp
	case guardUp:
		*p = guardRight
	case guardRight:
		*p = guardDown
	}
}

// Step moves the guard one move
func (g grid) Step() status {
	r, c := g.GuardLocation()

	if r < 0 {
		return oob
	}

	var s status

	switch g[r][c] {
	case guardDown:
		s = g.move(r, c, r+1, c)
	case guardUp:
		s = g.move(r, c, r-1, c)
	case guardLeft:
		s = g.move(r, c, r, c-1)
	case guardRight:
		s = g.move(r, c, r, c+1)
	}

	return s
}

// Load loads a grid
func Load(r io.Reader) grid {
	var g grid
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// Determine if the line is a rule or an update
		line := scanner.Text()
		gridLine := make([]block, 0)
		for _, c := range line {
			switch c {
			case '.':
				gridLine = append(gridLine, empty)
			case '#':
				gridLine = append(gridLine, wall)
			case '>':
				gridLine = append(gridLine, guardRight)
			case '<':
				gridLine = append(gridLine, guardLeft)
			case 'V':
				gridLine = append(gridLine, guardDown)
			case '^':
				gridLine = append(gridLine, guardUp)
			case 'X':
				gridLine = append(gridLine, visited)
			}
		}
		g = append(g, gridLine)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return g
}
