package main

import "fmt"

func main() {
	var vector [5]int
	vector[0] = 10
	fmt.Println(vector)

	vector2 := [3]int{3, 4, 6}
	fmt.Println(vector2)

	for i := 0; i < len(vector2); i++ {
		fmt.Println(vector2[i])
	}
}
