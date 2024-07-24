package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

// a person method
func (p Person) PrintName() {
	fmt.Printf("%s %s\n", p.FirstName, p.LastName)
}

func (p Person) PrintDetails() {
	fmt.Printf("[Date of birth %s, Email %s, Location %s]\n", p.Dob.String(), p.Email, p.Location)
}

// a person method with pointer receiver
func (p *Person) ChangeLocation(newLocation string) {
	p.Location = newLocation
}

func main() {
	p := &Person{
		FirstName: "John",
		LastName:  "Doe",
		Dob:       time.Date(1980, time.October, 1, 0, 0, 0, 0, time.UTC),
		Email:     "aoeu@email.com",
		Location:  "Box Hill",
	}
	p.PrintDetails()
	p.ChangeLocation("Sydney")
	p.PrintName()
	p.PrintDetails()
}
