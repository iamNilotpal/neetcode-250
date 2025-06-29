package main

import "fmt"

func markRow(matrix [][]int, row int) {
	cols := len(matrix[0])
	for col := range cols {
		if matrix[row][col] != 0 {
			matrix[row][col] = -1
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

func setZeroesBetter(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	tempRows := make([]int, rows)
	tempCols := make([]int, cols)

	for row := range rows {
		for col := range cols {
			if matrix[row][col] == 0 {
				tempRows[row] = 1
				tempCols[col] = 1
			}
		}
	}

	for row := range rows {
		for col := range cols {
			if tempRows[row] == 1 || tempCols[col] == 1 {
				matrix[row][col] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroesBruteForce(matrix)
	fmt.Printf("setZeroesBruteForce : %+v \n", matrix)

	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroesBetter(matrix)
	fmt.Printf("setZeroesBetter : %+v \n", matrix)
}
