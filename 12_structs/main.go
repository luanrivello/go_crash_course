package main

import (
	"fmt"
	"strconv"
)

//Person Coment
type Person struct {
	firstName string
	lastName  string
	city      string
	age       int
}

func (p Person) greeting() string {
	return "Greetings " + p.firstName + " " + p.lastName + " of the age " + strconv.Itoa(p.age)
}

func (p *Person) birthday() {
	p.age++
}

func (p Person) copy() Person {
	return p
}

func main() {

	person1 := Person{
		firstName: "John",
		lastName:  "Doe",
		city:      "New York",
		age:       22}

	//person1 := Person {"John", "Doe", "New York", 22}

	fmt.Println(person1)

	person2 := person1.copy()

	person1.birthday()

	fmt.Println(person2.greeting())
	fmt.Println(person1.greeting())
}
