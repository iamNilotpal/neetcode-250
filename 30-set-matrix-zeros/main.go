package main

import "fmt"

func setZeroesBruteForce(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	for row := range rows {
		for col := range cols {
			if matrix[row][col] == 0 {
				markRow(matrix, row)
				markColumn(matrix, col)
			}
		}
	}

	for row := range rows {
		for col := range cols {
			if matrix[row][col] == -1 {
				matrix[row][col] = 0
			}
		}
	}
}

func markRow(matrix [][]int, row int) {
	cols := len(matrix[0])
	for i := row; i <= row; i++ {
		for col := range cols {
			if matrix[row][col] != 0 {
				matrix[row][col] = -1
			}
		}
	}
}

func markColumn(matrix [][]int, col int) {
	rows := len(matrix)
	for row := range rows {
		if matrix[row][col] != 0 {
			matrix[row][col] = -1
		}
	}
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	setZeroesBruteForce(matrix)
	fmt.Printf("setZeroesBruteForce : %+v \n", matrix)
}
