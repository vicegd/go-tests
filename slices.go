package main

import "fmt"

func main() {
	slice := []int{1,2,4}
	fmt.Println(slice)

	elements := [10]int{0,1,2,3,4,5,6,7,8,9}
	slice = elements[5:8]
	fmt.Println(slice)

	/*Go reserva 100 espacios (capacidad) para el slice 
	pero solo vamos a tener 4 elementos en principio*/
	slice2 := make([]int, 4, 100)
	fmt.Printf("Length %d, Capacity %d \n", len(slice2), cap(slice2))
	slice2 = append(slice2, 3);
	fmt.Println(slice2)
}