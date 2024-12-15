package fifteen

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Block rune

const (
	Empty Block = '.'
	Wall        = '#'
	Box         = 'O'
	Robot       = '@'
)

type Warehouse struct {
	Map   [][]Block
	Moves []Direction
}

type Direction rune

const (
	North Direction = '^'
	South           = 'v'
	East            = '<'
	West            = '>'
)

func (d Direction) String() string {
	return string(rune(d))
}

func (w Warehouse) String() string {
	var sb strings.Builder

	// Write map first
	for r := range w.Map {
		for c := range w.Map[r] {
			sb.WriteRune(rune(w.Map[r][c]))
		}
		sb.WriteRune('\n')
	}

	// Then directions
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("%v", w.Moves))

	return sb.String()
}

// Shift moves each field one unit in the direction, if space is available
// indicated by bool = true if possible
func (w *Warehouse) Shift(r, c int, d Direction) bool {
	// End Case of recursion:
	// We have a wall or a free space
	curr := w.Map[r][c]

	if curr == Empty {
		return true
	} else if curr == Wall {
		return false
	}

	// Otherwise, continue to next neighbor
	nR, nC := r, c
	switch d {
	case North:
		nR--
	case South:
		nR++
	case East:
		nC--
	case West:
		nC++
	}

	possible := w.Shift(nR, nC, d)

	// Move to the next position if not blocked
	if possible {
		w.Map[nR][nC], w.Map[r][c] = curr, Empty
	}

	return possible
}

func (w *Warehouse) RobotLoc() (int, int) {
	for r := range w.Map {
		for c := range w.Map[r] {
			if w.Map[r][c] == Robot {
				return r, c
			}
		}
	}

	return -1, -1
}

func (w *Warehouse) Step() {
	if len(w.Moves) < 1 {
		return
	}

	var d Direction
	// Pop direction
	d, w.Moves = w.Moves[0], w.Moves[1:]
	r, c := w.RobotLoc()
	w.Shift(r, c, d)
}

func (w *Warehouse) SumBoxes() int {
	total := 0
	for r := range w.Map {
		for c := range w.Map[r] {
			if w.Map[r][c] == Box {
				total += (r*100 + c)
			}
		}
	}

	return total
}

func ParseWareHouse(r io.Reader) *Warehouse {
	var m [][]Block
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// Empty line means we have finished
		if len(line) == 0 {
			break
		}

		m = append(m, []Block(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Now we scan directional chars
	var d []Direction
	for scanner.Scan() {
		for _, c := range scanner.Text() {
			d = append(d, Direction(c))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return &Warehouse{
		Map:   m,
		Moves: d,
	}
}
