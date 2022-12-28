package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"	
)

func readFile() {
	file, err := ioutil.ReadFile("./utils/file.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(string(file))
	}
}

func readFile2() {
	file, err := os.Open("./utils/file.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

func main() {
	readFile()
	readFile2()
}