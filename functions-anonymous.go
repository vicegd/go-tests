package main

import "fmt"

var Calculation func(int, int) int = func(num1 int, num2 int) int {
	return num1 * num2
}

func main() {
	fmt.Printf("%d * %d = %d \n", 6, 3, Calculation(6, 3))

	Calculation = func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Printf("%d / %d = %d \n", 6, 3, Calculation(6, 3))
}