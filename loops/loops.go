package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := 1
	for i <= 5 {
		fmt.Print(i, " ")
		i++
	}	
	for {
		value := rand.Intn(1000000000)
		if value == 27 {
			break
		}
	}
	fmt.Println()
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
}