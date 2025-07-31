package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	// ?
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = i * j
		}
	}
	return matrix
}

func main() {
	var a = createMatrix(5, 10)
	// Print the whole 2D slice

	// Or row‐by‐row
	for _, row := range a {
		fmt.Println(row)
	}
}
