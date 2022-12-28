package main

import (
	"fmt"
	"math"
)

func main() {
	result := root_square(25)
	fmt.Println(result)

	result1, result2 := root_square2(25)
	fmt.Println(result1, result2)

	result3 := calculate(1, 2, 3, 4, 10, 20, 100)
	fmt.Println(result3)
}

func root_square(num int) float64  {
	return math.Sqrt(float64(num))
}

func root_square2(num int) (float64, bool)  {
	return math.Sqrt(float64(num)), true
}

func calculate(numbers ...int) (r int)  {
	for i, num := range numbers {
		fmt.Println(i, num)
		r += num
	}
	return
}