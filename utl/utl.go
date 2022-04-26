package utl

func TransposeSlices[T any](slice [][]T) [][]T {
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

func Rotate[T any](matrix [][]T) {
	var temp T
	for i := 0; i < 3; i++ {
		for i, n := 0, len(matrix)-1; i <= n/2; i++ {
			for j := i; j < n-i; j++ {
				temp = matrix[j][n-i]
				matrix[j][n-i] = matrix[i][j]
				matrix[i][j] = matrix[n-j][i]
				matrix[n-j][i] = matrix[n-i][n-j]
				matrix[n-i][n-j] = temp
			}
		}
	}
}

func DiagonalView(matrix [][]string) string {
	str := ""
	// Get the size
	var row int = len(matrix)
	var col int = len(matrix[0])
	// First Half
	for i := 0; i < col; i++ {
		for j := i; j >= 0 && i-j < row; j-- {
			str = str + (matrix[i-j][j])
		}
	}
	// Second Half
	for i := 1; i < row; i++ {
		for j, k := col-1, i; j >= 0 && k < row; j-- {
			str = str + (matrix[k][j])
			k++
		}
	}
	return str
}
