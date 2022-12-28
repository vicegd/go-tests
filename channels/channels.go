package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan time.Duration)
	go loop(channel1)
	fmt.Println("Here")
	info := <- channel1 //listen
	fmt.Println(info)
}

func loop(channel1 chan time.Duration) {
	begin := time.Now()
	for i:=0; i<9999999999; i++ {

	}
	end := time.Now()
	channel1 <- end.Sub(begin) //store
}