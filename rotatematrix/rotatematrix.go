package rotatematrix

import (
	"fmt"
)

// RotateMatrix rotate the matrix 90degress clockwise
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/770/
// n * m
// n row length
// m column length
func RotateMatrix(matrix [][]int) {
	n := len(matrix)

	//m := len(matrix[0])
	fmt.Println("N:", n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ", matrix[i][j])
			temp := matrix[j][i]
			matrix[j][i] = matrix[i][j]
			matrix[i][j] = temp

		}
		fmt.Println("")
	}
	fmt.Println("AFTER 1st ")
	for i := 0; i <= n-1; i++ {
		for j := 0; j <= n-1; j++ {
			fmt.Print(" ", matrix[i][j])

		}
		fmt.Println("")
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			fmt.Print(" ", matrix[i][j])
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][n-1-j]
			matrix[i][n-1-j] = temp

		}
		fmt.Println("")
	}

	fmt.Println("AFTER 2nd")
	for i := 0; i <= n-1; i++ {
		for j := 0; j <= n-1; j++ {
			fmt.Print(" ", matrix[i][j])

		}
		fmt.Println("")
	}
}
