package main

import "fmt"

var A bool    
var b int         
var c string = "a"
var d = 3.14 

func main() {
	var e, f int
	g, h := "hello", 2
	e = 3
	var i float32 = float32(e)

	fmt.Println("Boolean: ", A)
	fmt.Println("Integer: ", b, e, f, h)
	fmt.Println("Float:   ", d, i)
	fmt.Println("String:  ", c, g)
}