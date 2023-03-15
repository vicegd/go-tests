package main

import (
	utils "go-tests/utils"
	"time"
)

func status(hb utils.HumanBeing) {
	hb.Stop()
	hb.Move()
	hb.Talk()
	hb.Speed()
}

func main() {
	person := new(utils.Person)
	person.NewPerson(1, "Pedro", time.Now(), true)
	status(person)
}
