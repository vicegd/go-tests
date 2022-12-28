package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("splitName in a synchronous way")
	splitName("Hello")
	fmt.Println("splitName in an unsynchronous way")
	go splitName("Hello")
	fmt.Println("just waiting some time")
	var value string
	fmt.Scan(&value)
}

func splitName(name string) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(time.Second)
		fmt.Println(letter)
	}
}