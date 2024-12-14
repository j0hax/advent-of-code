package fourteen

import "fmt"
import "slices"

func GridSum(grid [][]int) int {
	total := 0
	for r := range grid {
		for c := range grid[r] {
			total += grid[r][c]
		}
	}

	return total
}

func transpose[T any](slice [][]T) [][]T {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]T, xl)
	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func PrintGrid(grid [][]int) {
	g := transpose(grid)
	for r := range g {
		for c := range g[r] {
			if g[r][c] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(g[r][c])
			}
		}
		fmt.Print("\n")
	}
}

func Quadrants(grid [][]int) [][][]int {
	quadrants := make([][][]int, 4)

	skipMidRow := len(grid)%2 == 1
	skipMidCol := len(grid[0])%2 == 1

	rows := len(grid) / 2
	cols := len(grid[0]) / 2

	if skipMidRow {
		grid = slices.Delete(grid, rows, rows+1)
	}

	// Anull mid col
	if skipMidCol {
		for i := 0; i < len(grid); i++ {
			grid[i] = slices.Delete(grid[i], cols, cols+1)
		}
	}

	// Allocate rowws for each quadrant
	quadrants[0] = make([][]int, rows)
	quadrants[1] = make([][]int, rows)
	quadrants[2] = make([][]int, rows)
	quadrants[3] = make([][]int, rows)

	for r := range rows {
		// Allocate cols for each quadrant
		quadrants[0][r] = make([]int, cols)
		quadrants[1][r] = make([]int, cols)
		quadrants[2][r] = make([]int, cols)
		quadrants[3][r] = make([]int, cols)
		for c := range cols {
			quadrants[0][r][c] = grid[r][c]
			quadrants[1][r][c] = grid[r+rows][c]
			quadrants[2][r][c] = grid[r][c+cols]
			quadrants[3][r][c] = grid[r+rows][c+cols]
		}
	}

	return quadrants
}
