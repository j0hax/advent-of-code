package four

import (
	"bufio"
	"fmt"
	"os"
)

func PrintMatrix[T rune](mat [][]T) {
	for i := range mat {
		for j := range mat[i] {
			fmt.Print(string(mat[i][j]))
		}
		fmt.Print("\n")
	}
}

// ReadMatrix reads a file into a [][]rune
func ReadMatrix(filename string) [][]rune {
	file, err := os.Open("./input4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var matrix [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	return matrix
}

// inBounds simply checks if a given coordinate is within the bounds of the
// matrix
func inBounds[T any](matrix [][]T, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(matrix) && j < len(matrix[i])
}

// remaining recursively checks to see if the substring matches for the rest of
// the matrix in a given direction.
func remaining(matrix [][]rune, i, j, di, dj int, substr string) bool {

	// condition 0:
	// empty substring means we have reached the end of recursion.
	if len(substr) == 0 {
		return true
	}

	// condition 1:
	// coords must be in-bounds
	if !inBounds(matrix, i, j) {
		return false
	}

	// condition 2:
	// (i, j) must match the first letter of the substring
	if matrix[i][j] != rune(substr[0]) {
		return false
	}

	// check the rest
	return remaining(matrix, i+di, j+dj, di, dj, substr[1:])
}

// countDirs checks the matrix at the given position for the substring in each
// direction
func countDirs(matrix [][]rune, i, j int, substr string) int {
	total := 0
	// up
	if remaining(matrix, i, j, 0, 1, substr) {
		total++
	}
	// down
	if remaining(matrix, i, j, 0, -1, substr) {
		total++
	}
	// backwards
	if remaining(matrix, i, j, -1, 0, substr) {
		total++
	}
	// forwards
	if remaining(matrix, i, j, 1, 0, substr) {
		total++
	}

	// diag up
	if remaining(matrix, i, j, 1, 1, substr) {
		total++
	}
	// diag down
	if remaining(matrix, i, j, 1, -1, substr) {
		total++
	}
	// diag backwards up
	if remaining(matrix, i, j, -1, 1, substr) {
		total++
	}
	// diag backwards down
	if remaining(matrix, i, j, -1, -1, substr) {
		total++
	}

	return total
}

// WordSearch searches for the word in all directions in the matrix
func WordSearch(matrix [][]rune, substr string) int {
	cnt := 0
	// Find first letter in the matrix, then check each direction
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == rune(substr[0]) {
				cnt += countDirs(matrix, i, j, substr)
			}
		}
	}
	return cnt
}

// something
func countX(matrix [][]rune, i, j int, substr string) int {
	total := 0

	// left to right
	if remaining(matrix, i-1, j-1, 1, 1, substr) && remaining(matrix, i-1, j+1, 1, -1, substr) {
		total++
	}

	// top to bottom
	if remaining(matrix, i+1, j+1, -1, -1, substr) && remaining(matrix, i-1, j+1, 1, -1, substr) {
		total++
	}

	// bottom to top
	if remaining(matrix, i-1, j-1, 1, 1, substr) && remaining(matrix, i+1, j-1, -1, 1, substr) {
		total++
	}

	// right to left
	if remaining(matrix, i+1, j+1, -1, -1, substr) && remaining(matrix, i+1, j-1, -1, 1, substr) {
		total++
	}

	return total
}

// CrossSeach searches for the crossed word diagonally
func CrossSearch(matrix [][]rune) int {
	cnt := 0
	// Find first letter in the matrix, then check each direction
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 'A' {
				cnt += countX(matrix, i, j, "MAS")
			}
		}
	}

	return cnt
}
