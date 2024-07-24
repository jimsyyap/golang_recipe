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
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

// a person method
func (p Person) PrintDetails() {
	fmt.Printf("[Email: %s, Location: %s, DoB: %s]\n", p.Email, p.Location, p.Dob.String())
}

type Admin struct {
	Person
	Roles []string
}

// overrides PrintDetails
func (a Admin) PrintDetails() {
	a.Person.PrintDetails()
	fmt.Println("Admin Roles:")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}
