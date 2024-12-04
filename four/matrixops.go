package four

/*
transpose returns a copy of a matrix with size (i,j) to (j,i), essentially

	a "diagonal" swap of values
*/
func transpose[T any](m [][]T) [][]T {
	// Initialize a new matrix with swapped lengths and heights
	newX := len(m[0])
	newY := len(m)
	n := make([][]T, newX)
	for i := range n {
		n[i] = make([]T, newY)
	}

	// Now copy with swapped indices
	for i := 0; i < newX-1; i++ {
		for j := 0; j < newY-1; j++ {
			n[i][j] = m[j][i]
		}
	}

	return n
}

// reverse returns a copy of the matrix with the order of its columns
// reversed
func reverse[T any](m [][]T) [][]T {
	// Initialize a new matrix
	newX := len(m)
	newY := len(m[0])
	//newY := len(m[0])
	n := make([][]T, newX)

	for i := range m {
		n[i] = make([]T, newY)

		for j := 0; j < newY; j++ {
			n[i][j] = m[i][newY-j-1]
		}
	}

	return n
}

// Rotates a matrix
// (brownie points for making it generic?)
func rotateMatrix[T any](mat [][]T) [][]T {
	/* To rotate a matrix by 90 degrees:
	   1. Transpose the matrix
	   2. Flip the matrix
	*/

	prog := transpose(mat)
	prog = reverse(mat)

	return prog
}
