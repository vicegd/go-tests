package main

import "fmt"

var result int
func main() {
	result = middleware(add)(4, 6)
	fmt.Printf("Result = %d \n", result) 

	result = middleware(subtract)(4, 6)
	fmt.Printf("Result = %d \n", result) 
}

func add(a,b int) int {
	return a+b
}

func subtract(a,b int) int {
	return a-b
}

func middleware(f func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		fmt.Println("Interesting stuff here...")
		return f(x, y)
	}
}