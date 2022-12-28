package main

import "fmt"

func main() {
	var matrix[5][4]int
	matrix[2][3] = 10
	fmt.Println(matrix)

	matrix2 := [3][2]int{{3, 4},{1, 2},{3, 56}}
	fmt.Println(matrix2)

	for i := 0; i < len(matrix2); i++ {
		for j := 0; j < len(matrix2[i]); j++ {
			fmt.Print(matrix2[i][j], " ")
		}
		fmt.Println()
	}
}