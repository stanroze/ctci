package Chapter_1

// Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes,
// write a method to rotate the image by 90 degrees. Can you do this in place?
func RotateMatrix(matrix [][]rune) [][]rune {
	// make sure the matrix is NxN
	n := len(matrix)
	if n < 1 || n != len(matrix[0]) {
		return matrix
	}

	for x := 0; x < n / 2; x++ {
		for y := x; y < n - 1 - x; y++ {
			tmp := matrix[x][y]
			matrix[x][y] = matrix[y][n - 1 - x]
			matrix[y][n - 1 - x] = matrix[n - 1 - x][n - 1 - y]
			matrix[n - 1 - x][n - 1 - y] = matrix[n - 1 - y][x]
			matrix[n - 1 - y][x] = tmp
		}
	}

	return matrix
}
