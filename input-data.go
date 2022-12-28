package main

import "fmt"

var num1 int
var num2 int

func main() {
	fmt.Print("Type a number: ")
	fmt.Scanf("%d ", &num1)

	fmt.Print("Type another number: ")
	fmt.Scanf("%d ", &num2)
	
	fmt.Printf("The result is: %d", num1+num2)
}