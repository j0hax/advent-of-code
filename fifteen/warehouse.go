package fifteen

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"
)

// A Tile is a basic block of the warehouse.
type Tile rune

const (
	Empty     Tile = '.'
	Wall           = '#'
	Box            = 'O'
	Robot          = '@'
	WideLeft       = '['
	WideRight      = ']'
)

// Warehouse contains the current simulation state.
type Warehouse struct {
	Map   [][]Tile
	Moves []Direction
}

// Direction is the orientation of the robot when pushing boxes
type Direction rune

const (
	Up    Direction = '^'
	Down            = 'v'
	Left            = '<'
	Right           = '>'
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

// Shift performs a simple (non-wide) recursive shift of boxes.
func Shift(m *[][]Tile, r, c int, d Direction) bool {
	// End Case of recursion:
	// We have a wall or a free space
	curr := (*m)[r][c]

	if curr == Empty {
		return true
	} else if curr == Wall {
		return false
	}

	// Otherwise, continue to next neighbor
	nR, nC := r, c
	switch d {
	case Up:
		nR--
	case Down:
		nR++
	case Left:
		nC--
	case Right:
		nC++
	}

	possible := Shift(m, nR, nC, d)

	// Move to the next position if not blocked
	if possible {
		(*m)[nR][nC], (*m)[r][c] = curr, Empty
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

// Step has the robot perform one move on a simple (i.e. not wide) warehouse.
func (w *Warehouse) Step() {
	if len(w.Moves) < 1 {
		return
	}

	var d Direction
	// Pop direction
	d, w.Moves = w.Moves[0], w.Moves[1:]
	r, c := w.RobotLoc()
	Shift(&w.Map, r, c, d)
}

// WideShift recursively moves wide (i.e. two char) boxes up or down.
func WideShift(m *[][]Tile, r, c int, d Direction) bool {
	// Calculate next coord
	nR, nC := r, c
	switch d {
	case Up:
		nR--
	case Down:
		nR++
	case Left:
		nC--
	case Right:
		nC++
	}

	switch (*m)[nR][nC] {
	case Wall:
		// End recursion, a wall cannot be moved.
		return false
	case Empty:
		// Move current element up and allow all previous calls to do as well
		(*m)[nR][nC], (*m)[r][c] = (*m)[r][c], Empty
		return true
	case WideLeft:
		// If we found a left hand box, try moving it, then the current spot
		if WideShift(m, nR, nC+1, d) && WideShift(m, nR, nC, d) {
			return WideShift(m, r, c, d)
		}
		return false
	case WideRight:
		// Same as left goes for right.
		if WideShift(m, nR, nC-1, d) && WideShift(m, nR, nC, d) {
			return WideShift(m, r, c, d)
		}
		return false
	}

	return false
}

/*
CopyMap returns a copy of the warehouse's floor.

This is useful to restore a previous state, for example, when working with wide
floors, which may alter the floor's state but determine that that box(es) cannot
be moved later.
*/
func (w *Warehouse) CopyMap() [][]Tile {
	fresh := make([][]Tile, len(w.Map))
	for i := range w.Map {
		fresh[i] = slices.Clone(w.Map[i])
	}
	return fresh
}

// WideStep consumes one move from the queue, moving any boxes in a wide
// warehouse.
func (w *Warehouse) WideStep() {
	if len(w.Moves) < 1 {
		return
	}

	var d Direction
	// Pop direction
	d, w.Moves = w.Moves[0], w.Moves[1:]
	r, c := w.RobotLoc()

	backup := w.CopyMap()
	if !WideShift(&w.Map, r, c, d) {
		// Restore backup
		w.Map = backup
	}
}

// SumBoxes counts the sumf of "GPS" coordinates of all normal or wide boxes
// contained in the warehouse.
func (w *Warehouse) SumBoxes() int {
	total := 0
	for r := range w.Map {
		for c := range w.Map[r] {
			if w.Map[r][c] == Box || w.Map[r][c] == WideLeft {
				total += (r*100 + c)
			}
		}
	}

	return total
}

/*
ParseWareHouse reads an input containing both the the floor state as well as
instructions for the robot. When wide = true, all blocks are doubled (Part 2)
*/
func ParseWareHouse(r io.Reader, wide bool) *Warehouse {
	var m [][]Tile
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// Empty line means we have finished
		if len(line) == 0 {
			break
		}

		// Normal and wide cases
		if !wide {
			m = append(m, []Tile(line))
		} else {
			var l []Tile
			for _, c := range line {
				switch Tile(c) {
				case Wall:
					l = append(l, Tile(Wall), Tile(Wall))
				case Box:
					l = append(l, Tile(WideLeft), Tile(WideRight))
				case Empty:
					l = append(l, Tile(Empty), Tile(Empty))
				case Robot:
					l = append(l, Tile(Robot), Tile(Empty))
				}
			}
			m = append(m, l)
		}
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
