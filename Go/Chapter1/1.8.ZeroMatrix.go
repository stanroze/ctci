package Chapter_1

// Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0.
func ZeroMatrix(matrix [][]int) [][]int {
	// make sure the first does not have a zero
	var firstRowHasZero,firstColHasZero bool
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			firstRowHasZero = true
			break
		}
	}
	for i := 0; i<len(matrix);i++ {
		if matrix[i][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// mark which cell has a 0, by using the first row and col as placeholders
	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}

	// zero out every row and col that have a zero
	for row := 0; row < len(matrix); row ++ {
		if matrix[row][0] == 0 {
			matrix = zeroRow(matrix, row)
		}
	}

	for col := 0; col < len(matrix[0]); col++ {
		if matrix[0][col] == 0 {
			matrix = zeroCol(matrix, col)
		}
	}

	if firstRowHasZero {
		matrix = zeroRow(matrix, 0)
	}
	if firstColHasZero {
		matrix = zeroCol(matrix, 0)
	}

	return matrix
}

func zeroRow(matrix [][]int, index int) [][]int {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[index][i] = 0
	}

	return matrix
}

func zeroCol(matrix [][]int, index int) [][]int {
	for i := 0; i < len(matrix); i++ {
		matrix[i][index] = 0
	}

	return matrix
}
