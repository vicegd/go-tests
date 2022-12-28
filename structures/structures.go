package main

import (
	"fmt"
	"time"
	utils "go-backend/utils"
)

func main() {
	user := new(utils.Person)
	user.NewPerson(1, "Pedro", time.Now(), true)
	fmt.Println(user)
}