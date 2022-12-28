package main

import "log"

func main() {
	defer func() {
		reco := recover()
		if reco != nil { //panic occured when reco != nil
			log.Fatalf("An error ocurred; panic was also generated: %s \n", reco)
		}
	}()

	if true == true {
		panic("Se termina")
	}
}