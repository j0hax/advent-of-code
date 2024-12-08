package four

import (
	"bufio"
	"io"
	"strings"
)

type WordSearch [][]rune

func (w WordSearch) String() string {
	var b strings.Builder
	for i := range w {
		for j := range w[i] {
			b.WriteRune(w[i][j])
		}
		b.WriteRune('\n')
	}

	return b.String()
}

// ReadWordSearch reads a file into a [][]rune
func ReadWordSearch(r io.Reader) WordSearch {
	var matrix WordSearch

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	return matrix
}

// InBounds simply checks if a given coordinate is within the bounds of a
// two-dimensional slice ("matrix")
func InBounds[T any](matrix [][]T, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(matrix) && j < len(matrix[i])
}

// remaining recursively checks to see if the substring matches for the rest of
// the matrix in a given direction.
func (matrix WordSearch) remaining(i, j, di, dj int, substr string) bool {
	// condition 0:
	// empty substring means we have reached the end of recursion.
	if len(substr) == 0 {
		return true
	}

	// condition 1:
	// coords must be in-bounds
	if !InBounds(matrix, i, j) {
		return false
	}

	// condition 2:
	// (i, j) must match the first letter of the substring
	if matrix[i][j] != rune(substr[0]) {
		return false
	}

	// check the rest
	return matrix.remaining(i+di, j+dj, di, dj, substr[1:])
}

// countDirs checks the matrix at the given position for the substring in each
// direction
func (matrix WordSearch) countDirs(i, j int, substr string) int {
	total := 0
	// up
	if matrix.remaining(i, j, 0, 1, substr) {
		total++
	}
	// down
	if matrix.remaining(i, j, 0, -1, substr) {
		total++
	}
	// backwards
	if matrix.remaining(i, j, -1, 0, substr) {
		total++
	}
	// forwards
	if matrix.remaining(i, j, 1, 0, substr) {
		total++
	}

	// diag up
	if matrix.remaining(i, j, 1, 1, substr) {
		total++
	}
	// diag down
	if matrix.remaining(i, j, 1, -1, substr) {
		total++
	}
	// diag backwards up
	if matrix.remaining(i, j, -1, 1, substr) {
		total++
	}
	// diag backwards down
	if matrix.remaining(i, j, -1, -1, substr) {
		total++
	}

	return total
}

// Count searches for the word in all directions (left/right, up/down, and
// diagonally in the matrix, returning the number of found strings.
func (w WordSearch) Count(substr string) int {
	cnt := 0
	// Find first letter in the matrix, then check each direction
	for i := range w {
		for j := range w[0] {
			if w[i][j] == rune(substr[0]) {
				cnt += w.countDirs(i, j, substr)
			}
		}
	}
	return cnt
}

// countX is a helper method for CrossCount, returning true if the specified
// indices are at the center of the crossed substring.
func (w WordSearch) countX(i, j int, substr string) bool {
	// left to right
	if w.remaining(i-1, j-1, 1, 1, substr) && w.remaining(i-1, j+1, 1, -1, substr) {
		return true
	}

	// top to bottom
	if w.remaining(i+1, j+1, -1, -1, substr) && w.remaining(i-1, j+1, 1, -1, substr) {
		return true
	}

	// bottom to top
	if w.remaining(i-1, j-1, 1, 1, substr) && w.remaining(i+1, j-1, -1, 1, substr) {
		return true
	}

	// right to left
	if w.remaining(i+1, j+1, -1, -1, substr) && w.remaining(i+1, j-1, -1, 1, substr) {
		return true
	}

	return false
}

// CrossSeach searches for the crossed word diagonally
func (w WordSearch) CrossCount(substr string) int {
	cnt := 0

	// Center rune of string
	c := rune(substr[len(substr)/2])

	// Find first letter in the matrix, then check each direction
	for i := range w {
		for j := range w[i] {
			if w[i][j] == c {
				if w.countX(i, j, substr) {
					cnt++
				}
			}
		}
	}

	return cnt
}
