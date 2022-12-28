package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.OpenFile("./utils/file-does-not-exist.txt", os.O_APPEND, 600)
	file.WriteString("New line...")
	defer file.Close()
	defer fmt.Println("Ending...")
}