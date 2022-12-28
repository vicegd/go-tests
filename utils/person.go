package utils

import (
	"fmt"
	"time"
)

type Person struct {
	Id int
	Name string
	LastAccess time.Time
	Active bool 
}

func (this *Person) Move() { fmt.Println("MOVE") }
func (this *Person) Stop() { fmt.Println("STOP") }
func (this *Person) Talk() { fmt.Println("TALK") }
func (this *Person) Speed() float32 {
	fmt.Println("SPEED")
	return 9.2
}

func (this *Person) NewPerson(id int, name string, 
	lastAccess time.Time, active bool) {
	this.Id = id
	this.Name = name
	this.LastAccess = lastAccess
	this.Active = active
}