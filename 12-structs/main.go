package main

import (
	"fmt"
	"strconv"
)

// Person struct definition
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method [value-reciever]
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + "." + "\nMy age is " + strconv.Itoa(p.age) + "."
}

// Has birthday method [pointer-reciver]
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried [pointer-receiver]

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastName = spouseLastName
	} else {
		return
	}
}

func main() {
	// Init Person using struct
	person1 := Person{
		firstName: "Samantha",
		lastName:  "Smith",
		city:      "Boston",
		gender:    "f",
		age:       25,
	}

	// Altenative initialization
	person2 := Person{
		"William",
		"Harvey",
		"Chicago",
		"m",
		28,
	}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.firstName)

	person1.age = 23
	fmt.Println(person1)

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person1.getMarried("McGill")
	fmt.Println(person1.greet())

	person2.getMarried("Thompson")
	fmt.Println(person2.greet())

}
