package main

import "fmt"

var age = 20

func main() {
	if age == 20 {
		fmt.Println(age)
	} else if age == 21 {
		fmt.Println("NO")
	} else {
		fmt.Println("ELSE")
	}

	switch age {
	case 10:
		fmt.Println("NO")
	case 20:
		fmt.Println("YES")
	default:
		fmt.Println("DEFAULT")
	}
}