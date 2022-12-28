package main

import (
	"fmt"
	"os"
)

func saveFile() {
	file, err := os.Create("../utils/file-create.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Fprintln(file, "Testing testing...")
	}
	file.Close()
}

func saveFile2() {
	file, err := os.OpenFile("../utils/file-create.txt", os.O_APPEND, 600)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		_, err = file.WriteString("New line...")
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func main() {
	saveFile()
	saveFile2()
}